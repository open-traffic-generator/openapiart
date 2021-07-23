import pytest
import jsonpath_ng as jp
import json


def test_add(api):
    config = api.prefix_config()
    assert config.a == config._DEFAULTS["a"]
    assert config.b == config._DEFAULTS["b"]
    g1 = config.g.add(name="unique list name", g_a="dkdkd", g_b=3, g_c=22.2)
    g1.g_d = "gdgdgd"
    j = config.j.add()
    j.j_b.f_a = "a"
    print(config)
    assert config.g[0].choice == "g_d"
    yaml = config.serialize(encoding=config.YAML)
    config.deserialize(yaml)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
