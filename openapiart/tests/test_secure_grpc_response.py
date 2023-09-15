from traceback import format_exception
import grpc
import pytest
import json


def test_grpc_set_config(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = secure_grpc_api.set_config(payload)
    assert result.read() == b"success"


def test_grpc_get_config(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    result = secure_grpc_api.get_config()
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


def test_grpc_valid_version_check(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    try:
        secure_grpc_api._version_check = True
        result = secure_grpc_api.set_config(payload)
        assert result.read() == b"success"
    finally:
        secure_grpc_api._version_check = False


def test_grpc_valid_inversion_check(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    try:
        secure_grpc_api.get_local_version().api_spec_version = "0.2.1"
        secure_grpc_api._version_check = True
        secure_grpc_api.set_config(payload)
        raise Exception("expected version error")
    except Exception:
        pass
    finally:
        secure_grpc_api.get_local_version().api_spec_version = "0.1.0"
        secure_grpc_api._version_check = False


def test_grpc_set_config_error_struct(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    payload["l"]["integer"] = 100
    try:
        secure_grpc_api.set_config(payload)
    except Exception as e:
        e_obj = e.args[0]
        assert e_obj.code == 13
        assert e_obj.errors[1] == "err2"
        err_obj = secure_grpc_api.from_exception(e)
        assert err_obj is not None
        assert err_obj.code == 13
        assert err_obj.errors[1] == "err2"


def test_grpc_set_config_error_str(utils, secure_grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    payload["l"]["integer"] = -3
    try:
        secure_grpc_api.set_config(payload)
    except Exception as e:
        e_obj = e.args[0]
        assert e_obj.code == 13
        assert e_obj.errors[0] == "some random error!"
        err_obj = secure_grpc_api.from_exception(e)
        assert err_obj is not None
        assert err_obj.code == 13
        assert err_obj.errors[0] == "some random error!"


def test_grpc_accept_yaml(secure_grpc_api):
    config = secure_grpc_api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 50
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]

    s_obj = config.serialize(encoding="yaml")
    secure_grpc_api.set_config(s_obj)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
