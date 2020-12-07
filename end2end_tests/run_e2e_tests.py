"""
Script for running end to end tests for GotwockAppServer.

The tests are rather simple sanity checks that the server is able to return data and response codes.
Note, that I am not running the entire test suite (all forbidden methods and multiple payload variants) as those things
are unit tested on the Go side.

Example usage:
    python3 -m pip install requests
    python3 end2end_tests/run_e2e_tests.py
"""
import requests

EMPTY_LOCATIONS_MESSAGE = '{"locations":null}'
DEFAULT_SERVER_ADDRESS = "http://127.0.0.1:9100"
HEADERS = {"Content-Type": "application/json"}


def _test_basic_request():
    """Send an example request to the server and verify response."""
    expected_status_code = 200

    payload = {"Latitude": 52.0989711, "Longitude": 21.2715719, "MaxDistance": 5.1}
    response = requests.post(DEFAULT_SERVER_ADDRESS, json=payload, headers=HEADERS)

    assert response.text != EMPTY_LOCATIONS_MESSAGE
    assert response.status_code == expected_status_code

    print("[OK] Test example basic request.")


def _test_another_request():
    """Send a different example request to the server and verify response."""
    expected_status_code = 200

    payload = {"Latitude": 52.1101533, "Longitude": 21.2567803, "MaxDistance": 3.0}
    response = requests.post(DEFAULT_SERVER_ADDRESS, json=payload, headers=HEADERS)

    assert response.text != EMPTY_LOCATIONS_MESSAGE
    assert response.status_code == expected_status_code

    print("[OK] Test another basic request.")


def _test_request_for_zero_locations():
    """Send request asking for locations with maxDistance equal to zero."""
    expected_status_code = 200

    payload = {"Latitude": 52.1101533, "Longitude": 21.2567803, "MaxDistance": 0.0}
    response = requests.post(DEFAULT_SERVER_ADDRESS, json=payload, headers=HEADERS)

    assert response.text == EMPTY_LOCATIONS_MESSAGE
    assert response.status_code == expected_status_code

    print("[OK] Test zero locations request.")


def _test_forbidden_request():
    """Send a request with forbidden DELETE method."""
    expected_status_code = 405

    payload = {"Latitude": 52.1101533, "Longitude": 21.2567803, "MaxDistance": 0}
    response = requests.delete(DEFAULT_SERVER_ADDRESS, json=payload, headers=HEADERS)

    assert response.status_code == expected_status_code

    print("[OK] Test forbidden request deny.")


def main():
    try:
        _test_basic_request()
        _test_another_request()
        _test_request_for_zero_locations()
        _test_forbidden_request()

        print("All test passed.")
    except requests.exceptions.ConnectionError:
        assert (
            False
        ), "The server is not running. Did you remember to start it on a separate process?"


if __name__ == "__main__":
    main()
