import importlib
import pytest

module = importlib.import_module("sanity")


def test_iter_set_method(default_config):
    default_config.j.add()
    default_config.j.set(0, module.JObject())
    assert len(default_config.j) == 1
    try:
        default_config.j.append(module.FObject())
        pytest.fail("appending an invalid object is not throwing exception")
    except Exception:
        pass
    try:
        default_config.j.set(0, module.FObject())
        pytest.fail("setting an invalid object is not throwing exception")
    except Exception:
        pass

    assert isinstance(default_config.j[0], module.EObject)


def test_validation_errors():
    p = module.Api().prefix_config()
    p.e
    try:
        p.validate()
        pytest.fail
    except Exception as e:
        assert "required field `prefix_config.a` must not be empty" in str(e)
        assert "required field `prefix_config.b` must not be empty" in str(e)
        assert "required field `prefix_config.c` must not be empty" in str(e)
        assert (
            "required field `prefix_config.required_object` must not be empty"
            in str(e)
        )
        assert "required field `prefix_config.e.e_a` must not be empty" in str(
            e
        )
        assert "required field `prefix_config.e.e_b` must not be empty" in str(
            e
        )

    p.e.e_a = "abc"
    try:
        p.validate()
    except Exception as e:
        print(e)
        assert (
            "value of `prefix_config.e.e_a` must be a valid float type, instead of `abc`"
            in str(e)
        )
    p.a = "abc"
    p.b = 10.1
    p.c = 20
    p.required_object.e_a = 10.1
    p.required_object.e_b = 20
    p.j.add().j_a
    p.mac_pattern.mac.values = ["1", "20"]
    p.ipv4_pattern.ipv4.value = "1.1"
    errors = p._validate(p._JSON_NAME, True)
    assert len([True for e in errors if ".e_b` must not be empty" in e]) == 2
    assert (
        "required field `prefix_config.j[0].e_a` must not be empty" in errors
    )
    assert "required field `prefix_config.e.e_b` must not be empty" in errors
    assert (
        "value of `prefix_config.e.e_a` must be a valid float type, instead of `abc`"
        in errors
    )
    assert (
        "required field `prefix_config.j[0].e_a` must not be empty" in errors
    )
    assert (
        "required field `prefix_config.j[0].e_b` must not be empty" in errors
    )
    assert (
        "value of `prefix_config.mac_pattern.mac.values[0]` must be a valid mac string, instead of `1`"
        in errors
    )
    assert (
        "value of `prefix_config.mac_pattern.mac.values[1]` must be a valid mac string, instead of `20`"
        in errors
    )
    assert (
        "value of `prefix_config.ipv4_pattern.ipv4.value` must be a valid ipv4 string, instead of `1.1`"
        in errors
    )


def test_enum_setter():
    p = module.Api().prefix_config()
    p.response = "abc"
    errors = p._validate(p._JSON_NAME, True)
    assert "abc is not a valid enum for property response" in errors
