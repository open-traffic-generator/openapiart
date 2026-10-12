import json
import importlib
import pytest
import yaml


def _reset_rpc_logger(module):
    impl = importlib.import_module("%s.%s" % (module.__name__, module.__name__))
    impl.RpcLogger.close()


def test_valid_version_check(api):
    try:
        api._version_check = True
        config = api.prefix_config()
        config.a = "asdf"
        config.b = 1.1
        config.c = 1
        config.required_object.e_a = 1.1
        config.required_object.e_b = 1.2
        config.d_values = [config.A, config.B, config.C]
        config.level.l1_p1.l2_p1.l3_p1 = "test"
        config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
        api.set_config(config)
    finally:
        api._version_check = False


def test_invalid_version_check(api):
    try:
        api.get_local_version().api_spec_version = "0.2.1"
        api._version_check = True
        config = api.prefix_config()
        config.a = "asdf"
        config.b = 1.1
        config.c = 1
        config.required_object.e_a = 1.1
        config.required_object.e_b = 1.2
        config.d_values = [config.A, config.B, config.C]
        config.level.l1_p1.l2_p1.l3_p1 = "test"
        config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
        api.set_config(config)
        raise Exception("expected version error")
    except Exception:
        pass
    finally:
        api.get_local_version().api_spec_version = "0.1.0"
        api._version_check = False


def test_error_for_non_okay_error_codes(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 500
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    config.level.l1_p1.l2_p1.l3_p1 = "test"
    config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
    with pytest.raises(Exception) as execinfo:
        api.set_config(config)

    e = execinfo.value.args[0]
    e.code == 500
    assert str(e.errors[0]) == "{'detail': 'invalid data type'}"
    err = api.from_exception(execinfo.value)
    assert err is not None
    assert err.code == 500
    assert str(err.errors[0]) == "{'detail': 'invalid data type'}"


def test_error_structure_for_non_okay_error_codes(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 400
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    config.level.l1_p1.l2_p1.l3_p1 = "test"
    config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
    with pytest.raises(Exception) as execinfo:
        api.set_config(config)

    e = execinfo.value.args[0]
    e.code == 400
    assert e.kind == "validation"
    assert e.errors[0] == "err for validation"
    err = api.from_exception(execinfo.value)
    assert err is not None
    assert err.code == 400
    assert err.kind == "validation"
    assert err.errors[0] == "err for validation"


def test_http_accepts_yaml_str(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 50
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    config.level.l1_p1.l2_p1.l3_p1 = "test"
    config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"

    s_obj = config.serialize()
    api.set_config(s_obj)


def test_error_incorrect_json_str(api):
    json_str = """
        {
            "abc": 456,
            "bcd": "fgh"
        }
    """

    with pytest.raises(Exception) as execinfo:
        api.set_config(json_str)

    print(execinfo.value)


def test_append_config(api):
    config = api.config_append()
    f1 = config.config_append_list.add().flows.add()
    f1.name = "f1"
    f1.rate = 23
    f2 = config.config_append_list.add().flows.add()
    f2.name = "f2"
    f2.rate = 32
    res = api.append_config(config.serialize())
    assert res.warnings == ["w1", "w2"]


def test_upload_config(api):
    bts = b"Hello\n123\nHello\n456!!@###"
    res = api.upload_config(bts)
    assert res.warnings == ["w1", "w2"]


def test_http_rpc_logging(monkeypatch, tmp_path, api, default_config):
    log_path = tmp_path / "http-rpc.yaml"
    monkeypatch.setenv("SANITY_RPC_LOG", str(log_path))
    monkeypatch.delenv("OPENAPIART_RPC_LOG", raising=False)
    _reset_rpc_logger(pytest.module)
    try:
        api.set_config(default_config)
    finally:
        _reset_rpc_logger(pytest.module)

    docs = list(yaml.safe_load_all(log_path.read_text()))
    entry = docs[-1]
    assert entry["transport"] == "http"
    assert entry["method"] == "POST /api/config"
    assert entry["request"]["a"] == "asdf"
    assert entry["response"]["status"] == 200


def test_http_rpc_logging_fallback(monkeypatch, tmp_path, api, default_config):
    log_path = tmp_path / "http-rpc-fallback.yaml"
    monkeypatch.delenv("SANITY_RPC_LOG", raising=False)
    monkeypatch.setenv("OPENAPIART_RPC_LOG", str(log_path))
    _reset_rpc_logger(pytest.module)
    try:
        api.set_config(default_config)
    finally:
        _reset_rpc_logger(pytest.module)

    docs = list(yaml.safe_load_all(log_path.read_text()))
    entry = docs[-1]
    assert entry["transport"] == "http"
    assert entry["method"] == "POST /api/config"
