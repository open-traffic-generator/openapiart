import pytest
import importlib
import logging
import json
import base64

def test_grpc_set_config(utils):
    module = importlib.import_module(pytest.module_name)
    grpc_api = module.api(
        location="localhost:50051",
        transport="grpc",
        logger=None,
        loglevel=logging.DEBUG,
    )
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = grpc_api.set_config(payload)
    assert base64.b64decode(result) == b"success"

if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])