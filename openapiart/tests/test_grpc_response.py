from traceback import format_exception
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
    except Exception as e:
        e_obj = e.args[0]
        assert e_obj.code == 13
        assert e_obj.errors[1] == "err2"
        err_obj = grpc_api.from_exception(e)
        assert err_obj is not None
        assert err_obj.code == 13
        assert err_obj.errors[1] == "err2"


def test_grpc_set_config_error_str(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    payload["l"]["integer"] = -3
    try:
        grpc_api.set_config(payload)
    except Exception as e:
        e_obj = e.args[0]
        assert e_obj.code == 13
        assert e_obj.errors[0] == "some random error!"
        err_obj = grpc_api.from_exception(e)
        assert err_obj is not None
        assert err_obj.code == 13
        assert err_obj.errors[0] == "some random error!"


def test_grpc_accept_yaml(grpc_api):
    config = grpc_api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 50
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]

    s_obj = config.serialize(encoding="yaml")
    grpc_api.set_config(s_obj)


def test_version_mismatch_error(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    try:
        grpc_api.get_local_version().api_spec_version = "2.0.1"
        grpc_api._version_check = True
        grpc_api._version_check_err = None
        grpc_api.set_component_info(
            "keng-controller", "1.8.0", "protocol-engine"
        )
        grpc_api.set_config(payload)
        raise Exception("expected version error")
    except Exception as e:
        assert (
            str(e)
            == "keng-controller 1.8.0 is not compatible with protocol-engine 1.2.3"
        )
    finally:
        grpc_api.get_local_version().api_spec_version = "0.1.0"
        grpc_api._version_check = False
    #
    # print(grpc_api.__dict__)


def test_stream_config(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    grpc_api.enable_grpc_streaming = True
    grpc_api.chunk_size = 200
    result = grpc_api.set_config(payload)
    assert result.read() == b"success"
    grpc_api.enable_grpc_streaming = False


def test_grpc_append_config(grpc_api):
    config = grpc_api.config_append()
    f1 = config.config_append_list.add().flows.add()
    f1.name = "f1"
    f1.rate = 23
    f2 = config.config_append_list.add().flows.add()
    f2.name = "f2"
    f2.rate = 32
    result = grpc_api.append_config(config)
    assert result.warnings == ["w1", "w2"]


def test_grpc_stream_get_metrics(grpc_api):
    grpc_api.enable_grpc_streaming = True
    grpc_api.chunk_size = 200
    mr = grpc_api.metrics_request()
    mr.port = "p1"
    result = grpc_api.get_metrics(mr)
    print(result)
    grpc_api.enable_grpc_streaming = False


def test_grpc_stream_get_config(grpc_api):
    grpc_api.enable_grpc_streaming = True
    result = grpc_api.get_config()
    assert result.b == 1.1
    assert result.d_values == ["a", "b", "c"]
    grpc_api.enable_grpc_streaming = False


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
