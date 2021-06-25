import pytest
import jsonpath_ng as jp
import json


def test_config(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.d_values = [config.A, config.B, config.C]
    config.e.e_a = 1.1
    config.e.e_b = 1.2
    config.f.f_a = "a"
    # g1 = config.g.append(name="unique list name", g_a="dkdkd", g_b=3, g_c=22.2)
    # g1.g_d = "gdgdgd"
    config.h = False
    config.i = "11011011"
    # j = config.j.append()
    # j.j_b.f_a = "a"
    djson = json.loads(config.serialize(config.JSON))
    assert jp.parse("$.a").find(djson)[0].value == config.a
    assert jp.parse("$.f.f_a").find(djson)[0].value == config.f.f_a
    print(config)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
