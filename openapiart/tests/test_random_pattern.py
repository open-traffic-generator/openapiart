import importlib
import pytest

module = importlib.import_module("sanity")


def test_random_pattern_integer_format(default_config):
    pat = default_config.integer_pattern.integer
    val = pat._TYPES.get("value")
    assert val.get("format") == "uint32"
    rnd = pat.random

    assert rnd._TYPES.get("min").get("format") == "uint32"
    assert rnd._TYPES.get("min").get("type") == int
    assert rnd._TYPES.get("min").get("maximum") == 255
    assert rnd._DEFAULTS.get("min") == 0

    assert rnd._TYPES.get("max").get("format") == "uint32"
    assert rnd._TYPES.get("max").get("type") == int
    assert rnd._TYPES.get("max").get("maximum") == 255
    assert rnd._DEFAULTS.get("max") == 255

    assert rnd._TYPES.get("count").get("format") == "uint32"
    assert rnd._TYPES.get("count").get("type") == int
    assert rnd._TYPES.get("count").get("maximum") == 255
    assert rnd._DEFAULTS.get("count") == 1

    assert rnd._TYPES.get("seed").get("format") == "uint32"
    assert rnd._TYPES.get("seed").get("type") == int
    assert rnd._TYPES.get("seed").get("maximum") == 255
    assert rnd._DEFAULTS.get("seed") == 1

    data = default_config.serialize("dict")
    intp = data["integer_pattern"]["integer"]
    assert intp["choice"] == "random"
    pat = intp["random"]
    assert pat["count"] == 1
    assert pat["max"] == 255
    assert pat["min"] == 0
    assert pat["seed"] == 1

    config = module.Api().prefix_config()
    config.deserialize(data)
    intp = config.integer_pattern.integer
    assert intp.choice == "random"
    pat = intp.random
    assert pat.count == 1
    assert pat.seed == 1
    assert pat.max == 255
    assert pat.min == 0


def test_random_pattern_ipv4_format(default_config):
    pat = default_config.ipv4_pattern.ipv4
    val = pat._TYPES.get("value")
    assert val.get("type") == str
    rnd = pat.random

    assert rnd._TYPES.get("min").get("type") == str
    assert rnd._TYPES.get("min").get("format") == "ipv4"
    assert rnd._DEFAULTS.get("min") == "0.0.0.0"

    assert rnd._TYPES.get("max").get("format") == "ipv4"
    assert rnd._TYPES.get("max").get("type") == str
    assert rnd._DEFAULTS.get("max") == "255.255.255.255"

    assert rnd._TYPES.get("count").get("format") == "uint32"
    assert rnd._TYPES.get("count").get("type") == int
    assert rnd._DEFAULTS.get("count") == 1

    assert rnd._TYPES.get("seed").get("format") == "uint32"
    assert rnd._TYPES.get("seed").get("type") == int
    assert rnd._DEFAULTS.get("seed") == 1

    data = default_config.serialize("dict")
    intp = data["ipv4_pattern"]["ipv4"]
    assert intp["choice"] == "random"
    pat = intp["random"]
    assert pat["count"] == 1
    assert pat["max"] == "255.255.255.255"
    assert pat["min"] == "0.0.0.0"
    assert pat["seed"] == 1

    config = module.Api().prefix_config()
    config.deserialize(data)
    intp = config.ipv4_pattern.ipv4
    assert intp.choice == "random"
    pat = intp.random
    assert pat.count == 1
    assert pat.seed == 1
    assert pat.max == "255.255.255.255"
    assert pat.min == "0.0.0.0"


def test_random_pattern_ipv6_format(default_config):
    pat = default_config.ipv6_pattern.ipv6
    val = pat._TYPES.get("value")
    assert val.get("type") == str
    rnd = pat.random

    assert rnd._TYPES.get("min").get("type") == str
    assert rnd._TYPES.get("min").get("format") == "ipv6"
    assert rnd._DEFAULTS.get("min") == "::0"

    assert rnd._TYPES.get("max").get("format") == "ipv6"
    assert rnd._TYPES.get("max").get("type") == str
    assert rnd._DEFAULTS.get("max") == "::"

    assert rnd._TYPES.get("count").get("format") == "uint32"
    assert rnd._TYPES.get("count").get("type") == int
    assert rnd._DEFAULTS.get("count") == 1

    assert rnd._TYPES.get("seed").get("format") == "uint32"
    assert rnd._TYPES.get("seed").get("type") == int
    assert rnd._DEFAULTS.get("seed") == 1

    data = default_config.serialize("dict")
    intp = data["ipv6_pattern"]["ipv6"]
    assert intp["choice"] == "random"
    pat = intp["random"]
    assert pat["count"] == 1
    assert pat["max"] == "::"
    assert pat["min"] == "::0"
    assert pat["seed"] == 1

    config = module.Api().prefix_config()
    config.deserialize(data)
    intp = config.ipv6_pattern.ipv6
    assert intp.choice == "random"
    pat = intp.random
    assert pat.count == 1
    assert pat.seed == 1
    assert pat.max == "::"
    assert pat.min == "::0"


def test_random_pattern_mac_format(default_config):
    pat = default_config.mac_pattern.mac
    val = pat._TYPES.get("value")
    assert val.get("type") == str
    rnd = pat.random

    assert rnd._TYPES.get("min").get("type") == str
    assert rnd._TYPES.get("min").get("format") == "mac"
    assert rnd._DEFAULTS.get("min") == "00:00:00:00:00:00"

    assert rnd._TYPES.get("max").get("format") == "mac"
    assert rnd._TYPES.get("max").get("type") == str
    assert rnd._DEFAULTS.get("max") == "ff:ff:ff:ff:ff:ff"

    assert rnd._TYPES.get("count").get("format") == "uint32"
    assert rnd._TYPES.get("count").get("type") == int
    assert rnd._DEFAULTS.get("count") == 1

    assert rnd._TYPES.get("seed").get("format") == "uint32"
    assert rnd._TYPES.get("seed").get("type") == int
    assert rnd._DEFAULTS.get("seed") == 1

    data = default_config.serialize("dict")
    intp = data["mac_pattern"]["mac"]
    assert intp["choice"] == "random"
    pat = intp["random"]
    assert pat["count"] == 1
    assert pat["max"] == "ff:ff:ff:ff:ff:ff"
    assert pat["min"] == "00:00:00:00:00:00"
    assert pat["seed"] == 1

    config = module.Api().prefix_config()
    config.deserialize(data)
    intp = config.mac_pattern.mac
    assert intp.choice == "random"
    pat = intp.random
    assert pat.count == 1
    assert pat.seed == 1
    assert pat.max == "ff:ff:ff:ff:ff:ff"
    assert pat.min == "00:00:00:00:00:00"
