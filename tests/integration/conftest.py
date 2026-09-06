import json
import os
import re
import time
from datetime import datetime

import pytest
import requests


def pytest_addoption(parser):
    parser.addoption(
        "--canonize",
        action="store_true",
        help="rewrite canonical JSON results with actual test output",
    )


class Canon:
    def __init__(self, directory, canonize):
        self.directory = directory
        self.canonize = canonize

    def response(self, response):
        location = response.headers.get("Location")
        if location is not None:
            location = re.sub(r"/\d+$", "/<ID>", location)

        result = {
            "status_code": response.status_code,
            "body": self._normalize(response.json()),
        }
        if location is not None:
            result["headers"] = {"Location": location}
        return result

    def assert_response(self, response, name):
        self.assert_data(self.response(response), name)

    def assert_data(self, actual, name):
        path = self.directory / f"{name}.json"
        if self.canonize:
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text(
                json.dumps(actual, ensure_ascii=False, indent=2, sort_keys=True) + "\n",
                encoding="utf-8",
            )

        if not path.exists():
            pytest.fail(
                f"Canonical result {path} does not exist. "
                "Run pytest --canonize and review the generated file."
            )

        expected = json.loads(path.read_text(encoding="utf-8"))
        assert actual == expected

    def _normalize(self, value, key=None):
        if key == "id" and isinstance(value, int):
            return "<ID>"
        if key is not None and key.endswith("_at") and self._is_timestamp(value):
            return "<TIMESTAMP>"
        if isinstance(value, dict):
            return {name: self._normalize(item, name) for name, item in value.items()}
        if isinstance(value, list):
            return [self._normalize(item) for item in value]
        return value

    @staticmethod
    def _is_timestamp(value):
        if not isinstance(value, str):
            return False
        try:
            datetime.fromisoformat(value.replace("Z", "+00:00"))
        except ValueError:
            return False
        return True


@pytest.fixture
def canon(request):
    return Canon(
        directory=request.path.parent / "canondata",
        canonize=request.config.getoption("--canonize"),
    )


@pytest.fixture(scope="session")
def api_url():
    return os.environ.get("CARLOG_API_URL", "http://localhost:8080")


@pytest.fixture(scope="session")
def api(api_url):
    session = requests.Session()

    deadline = time.monotonic() + 30
    while time.monotonic() < deadline:
        try:
            response = session.get(f"{api_url}/health", timeout=1)
            if response.status_code == 200:
                break
        except requests.RequestException:
            pass
        time.sleep(0.5)
    else:
        pytest.fail(f"API did not become healthy at {api_url}")

    yield session
    session.close()
