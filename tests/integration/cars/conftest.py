import hashlib

import pytest


@pytest.fixture
def car_payload(request):
    suffix = hashlib.sha256(request.node.name.encode()).hexdigest()[:12]
    make = f"Integration Make {suffix}"

    return {
        "make": make,
        "model": f"Integration Model {suffix}",
        "vin": f"TEST{suffix}"[:17].upper(),
        "year": 2022,
        "current_mileage": 42_000,
    }
