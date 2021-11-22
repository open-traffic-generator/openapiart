import pytest
import jsonpath_ng as jp
import json


def test_config(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.d_values = [config.A, config.B, config.C]
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.e.e_a = 1.1
    config.e.e_b = 1.2
    config.f.f_a = "a"
    config.g.add(g_a="a g_a value")
    config.h = False
    config.i = "11011011"
    _, g1, g2 = config.g.gobject().gobject()
    assert g1 == g2
    config.k.e_object.e_a = 77.7
    config.k.e_object.e_b = 77.7
    config.k.f_object.f_a = "asdf"
    # assert j1.choice == j1._DEFAULTS["choice"]
    djson = json.loads(config.serialize(config.JSON))
    assert jp.parse("$.a").find(djson)[0].value == config.a
    assert jp.parse("$.f.f_a").find(djson)[0].value == config.f.f_a
    yaml = config.serialize(encoding=config.JSON)
    config.deserialize(yaml)
    print(yaml)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
