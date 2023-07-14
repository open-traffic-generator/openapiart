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


def test_format_integer_pattern(default_config):
    pat = default_config.integer_pattern.integer
    val = pat._TYPES.get("value")
    assert val.get("format") == "uint32"
    inc = pat.increment
    assert inc._TYPES.get("start").get("format") == "uint32"
    assert inc._TYPES.get("step").get("format") == "uint32"
    assert inc._TYPES.get("count").get("format") == "uint32"
    dec = pat.decrement
    assert dec._TYPES.get("start").get("format") == "uint32"
    assert dec._TYPES.get("step").get("format") == "uint32"
    assert dec._TYPES.get("count").get("format") == "uint32"


def test_format_count_pattern(default_config):
    ipv4 = default_config.ipv4_pattern.ipv4
    assert ipv4.increment._TYPES.get("count").get("format") == "uint32"
    assert ipv4.decrement._TYPES.get("count").get("format") == "uint32"
    ipv6 = default_config.ipv6_pattern.ipv6
    assert ipv6.increment._TYPES.get("count").get("format") == "uint32"
    assert ipv4.decrement._TYPES.get("count").get("format") == "uint32"
    mac = default_config.mac_pattern.mac
    assert mac.increment._TYPES.get("count").get("format") == "uint32"
    assert ipv4.decrement._TYPES.get("count").get("format") == "uint32"
