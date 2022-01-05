import pytest
import importlib
import logging
import json
import base64

def test_grpc_set_config(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = grpc_api.set_config(payload)
    assert base64.b64decode(result) == b"success"

def test_grpc_get_config(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = grpc_api.get_config()
    assert result.a == payload.get('a')


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])