import pytest


def test_remove(api):
    config = api.prefix_config()
    config.g.add(name="a")
    config.g.add(name="b")
    config.g.add(name="c")
    assert len(config.g) == 3
    assert config.g[0].name == "a"
    assert config.g[1].name == "b"
    assert config.g[2].name == "c"

    config.g.remove(1)
    config.g.add(name="d")
    assert len(config.g) == 3
    assert config.g[0].name == "a"
    assert config.g[1].name == "c"
    assert config.g[2].name == "d"


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
