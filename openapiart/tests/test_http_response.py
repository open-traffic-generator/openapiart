import json
import pytest


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
