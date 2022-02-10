import pytest
import json


def test_grpc_set_config(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = grpc_api.set_config(payload)
    assert result.read() == b"success"


def test_grpc_get_config(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = grpc_api.get_config()
    assert result.a == payload.get("a")


def test_invalid_transport():
    with pytest.raises(Exception) as execinfo:
        pytest.module.api(location="localhost", transport="NetConf")
    assert (
        execinfo.value.args[0]
        == """NetConf is not within valid transport types ['http', 'grpc']"""
    )


def test_transport_with_ext():
    with pytest.raises(Exception) as execinfo:
        pytest.module.api(
            location="localhost",
            transport=pytest.module.Transport.GRPC,
            ext="ixnetwork",
        )
    assert (
        execinfo.value.args[0]
        == """ext and transport are not mutually exclusive. Please configure one of them."""
    )


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
