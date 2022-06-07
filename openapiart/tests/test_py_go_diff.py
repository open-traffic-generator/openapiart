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
