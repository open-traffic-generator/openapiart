import importlib
import pytest
module = importlib.import_module("sanity")


def test_iter_set_method(default_config):
    default_config.j.add()
    default_config.j.set(0, module.JObject())
    assert len(default_config.j) == 1
    # below 2 statements shall raise exception need to modify generator
    # if exception is raised, it would conflict with below mentioned
    # comments about flow.packet
    default_config.j.append(module.FObject())
    default_config.j.set(1, module.FObject())

    default_config.j.set(1, module.JObject())
    # below statement is a replica of snappi flow.packet Iter
    # indexing will return protocol header obj instead of FlowHeader
    # which is actually set with set or append or add methods
    assert isinstance(default_config.j[1], module.EObject)


def test_validation_errors():
    p = module.Api().prefix_config()
    p.e
    try:
        p.validate()
        pytest.fail
    except Exception as e:
        assert "a is a mandatory property" in str(e)
        assert "b is a mandatory property" in str(e)
        assert "c is a mandatory property" in str(e)
        assert "required_object is a mandatory property" in str(e)
        assert "e_a is a mandatory property" in str(e)
        assert "e_b is a mandatory property" in str(e)
    
    p.e.e_a = "abc"
    try:
        p.validate()
    except Exception as e:
        assert "Invalid abc format on property e_a" in str(e)
    p.a = "abc"
    p.b = 10.1
    p.c = 20
    p.required_object.e_a = 10.1
    p.required_object.e_b = 20
    p.j.add().j_a
    errors = p.validate(True)
    assert len([True for e in errors if "e_b is a mandatory property" in e]) == 2
    

def test_enum_setter():
    p = module.Api().prefix_config()
    p.response = "abc"
    errors = p.validate(True)
    assert "abc is not a valid enum for property response" in errors
