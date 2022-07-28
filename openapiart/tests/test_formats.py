import pytest


def test_formats_sanity(config):
    config.l.string_param = "asdf"
    config.l.integer = 88
    config.l.float = 22.3
    config.l.double = 2342.222
    config.l.mac = "00:00:fa:ce:fa:ce"
    config.l.ipv4 = "1.1.1.1"
    config.l.ipv6 = "::02"
    config.l.hex = "0102030405060708090a0b0c0d0e0f"
    yaml = config.serialize(encoding=config.YAML)
    config.deserialize(yaml)
    print(yaml)


@pytest.mark.parametrize("value", [33.4, 100])
def test_formats_bad_string(config, value):
    config.l.string_param = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail(
            "Value {value} was successfully validated".format(value=value)
        )
    except TypeError:
        pass


@pytest.mark.parametrize("value", [33.4, "asdf"])
def test_formats_bad_integer(config, value):
    config.l.integer = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", [6, 100, -20])
def test_formats_integer_to_be_removed(config, value):
    try:
        config.l.integer = value
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1.1.1.1", "01.002.003.4", "0.0.0.0"])
def test_formats_good_ipv4(config, value):
    config.l.ipv4 = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
    except TypeError:
        pytest.fail("Value {} was not valid".format(value))


@pytest.mark.parametrize(
    "value",
    [
        "1.1. 1.1",
        33.4,
        "asdf",
        100,
        -20,
        "::01",
        "1.1.1.1.1",
        "256.256.256.256",
        "-255.-255.-255.-255",
    ],
)
def test_formats_bad_ipv4(config, value):
    config.l.ipv4 = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1.1", "1.1.1", " 1.1.1.1 "])
def test_formats_ipv4_to_be_removed(config, value):
    try:
        config.l.ipv4 = value
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize(
    "value",
    [
        33.4,
        "asdf",
        "1.1.1.1",
        100,
        -20,
        "65535::65535",
        "ab: :ab",
        "ab:ab:ab",
        "ffff0::ffff0",
    ],
)
def test_formats_bad_ipv6(config, value):
    config.l.ipv6 = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize(
    "value",
    [
        1,
        2.2,
        "1.1.1.1",
        "::01",
        "00:00:00",
        "00:00:00:00:gg:00",
        "00:00:fa:ce:fa:ce:01",
        "255:255:255:255:255:255",
    ],
)
def test_formats_bad_mac(config, value):
    config.l.mac = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize(
    "value", [1, 2.2, "1.1.1.1", "::01", "00:00:fa:ce:fa:ce:01"]
)
def test_formats_bad_hex(config, value):
    config.l.hex = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1234567", "12"])
def test_string_length(config, value):
    config.str_len = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


def test_int64_list(config, default_config):
    config.integer64 = 10
    config.integer64_list = [1, 2]
    data = config.serialize("dict")
    assert isinstance(data["integer64"], str)
    for d in data["integer64_list"]:
        assert isinstance(d, str)
    default_config.deserialize(data)
    assert isinstance(default_config.integer64, int)
    for d in default_config.integer64_list:
        assert isinstance(d, int)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
