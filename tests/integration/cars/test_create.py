import pytest


pytestmark = pytest.mark.integration


def test_create_car(api, api_url, car_payload, canon):
    response = api.post(f"{api_url}/cars", json=car_payload, timeout=5)

    canon.assert_response(response, "create_success")


def test_reject_duplicate_vin(api, api_url, car_payload, canon):
    first_response = api.post(f"{api_url}/cars", json=car_payload, timeout=5)
    duplicate = car_payload | {"vin": car_payload["vin"].lower()}
    duplicate_response = api.post(f"{api_url}/cars", json=duplicate, timeout=5)

    canon.assert_data(
        {
            "first": canon.response(first_response),
            "duplicate": canon.response(duplicate_response),
        },
        "duplicate_vin",
    )


def test_reject_empty_make(api, api_url, car_payload, canon):
    payload = car_payload | {"make": "   "}

    response = api.post(f"{api_url}/cars", json=payload, timeout=5)

    canon.assert_response(response, "empty_make")
