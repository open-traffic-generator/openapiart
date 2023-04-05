import grpc
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


def test_grpc_valid_version_check(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    try:
        grpc_api._version_check = True
        result = grpc_api.set_config(payload)
        assert result.read() == b"success"
    finally:
        grpc_api._version_check = False


def test_grpc_valid_inversion_check(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    try:
        grpc_api.get_local_version().api_spec_version = "0.2.1"
        grpc_api._version_check = True
        grpc_api.set_config(payload)
        raise Exception("expected version error")
    except Exception:
        pass
    finally:
        grpc_api.get_local_version().api_spec_version = "0.1.0"
        grpc_api._version_check = False


def test_grpc_set_config_error_struct(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    payload["l"]["integer"] = 100
    try:
        grpc_api.set_config(payload)
    except grpc.RpcError as e:
        err_obj = grpc_api.from_exception(e)
        assert err_obj.code == 13
        assert len(err_obj.errors) == 2
        assert err_obj.errors[1] == "err2"


def test_grpc_set_config_error_str(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    payload["l"]["integer"] = -3
    try:
        grpc_api.set_config(payload)
    except grpc.RpcError as e:
        err_obj = grpc_api.from_exception(e)
        assert err_obj.code == 2
        assert len(err_obj.errors) == 1
        assert err_obj.errors[0] == "some random error!"


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
