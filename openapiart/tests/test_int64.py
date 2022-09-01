import importlib

module = importlib.import_module("sanity")


def test_int64(default_config):
    value1 = default_config._TYPES.get("full_duplex_100_mb")
    value2 = default_config._TYPES.get("integer64_list")
    assert value1.get("format") is not None
    assert value2.get("itemformat") is not None
    assert value1.get("format") == "int64"
    assert value2.get("itemformat") == "int64"
    default_config.full_duplex_100_mb = 100
    default_config.integer64_list = [2000]
    data = default_config.serialize("dict")
    assert isinstance(data["full_duplex_100_mb"], str)
    assert isinstance(data["integer64_list"][0], str)

    config = module.Api().prefix_config()
    config.deserialize(data)
    assert isinstance(config.full_duplex_100_mb, int)
    assert isinstance(config.integer64_list[0], int)
