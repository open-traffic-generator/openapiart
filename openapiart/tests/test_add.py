import pytest


def test_add(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    assert config.f.f_a == config.f._DEFAULTS["f_a"]
    assert config.f.f_b == config.f._DEFAULTS["f_b"]
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
