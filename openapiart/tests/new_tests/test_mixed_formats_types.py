import pytest
import importlib

module = importlib.import_module("sanity")


def test_formats_sanity(api):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.string_param = "asdf"
    mixed_object.integer = 88
    mixed_object.float = 22.3
    mixed_object.double = 2342.222
    mixed_object.mac = "00:00:fa:ce:fa:ce"
    mixed_object.ipv4 = "1.1.1.1"
    mixed_object.ipv6 = "::02"
    mixed_object.hex = "0102030405060708090a0b0c0d0e0f"
    yaml = config.serialize(encoding=config.YAML)
    config.deserialize(yaml)
    print(yaml)


@pytest.mark.parametrize("value", [33.4, 100])
def test_formats_bad_string(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.string_param = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail(
            "Value {value} was successfully validated".format(value=value)
        )
    except TypeError:
        pass


@pytest.mark.parametrize("value", [33.4, "asdf"])
def test_formats_bad_integer(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.integer = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", [6, 100, -20])
def test_formats_integer_to_be_removed(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    try:
        mixed_object.integer = value
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1.1.1.1", "01.002.003.4", "0.0.0.0"])
def test_formats_good_ipv4(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.ipv4 = value
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
def test_formats_bad_ipv4(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.ipv4 = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1.1", "1.1.1", " 1.1.1.1 "])
def test_formats_ipv4_to_be_removed(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    try:
        mixed_object.ipv4 = value
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
def test_formats_bad_ipv6(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.ipv6 = value
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
def test_formats_bad_mac(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.mac = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize(
    "value", [1, 2.2, "1.1.1.1", "::01", "00:00:fa:ce:fa:ce:01"]
)
def test_formats_bad_hex(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.hex = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


@pytest.mark.parametrize("value", ["1234567", "12"])
def test_string_length(api, value):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.str_len = value
    try:
        config.deserialize(config.serialize(encoding=config.YAML))
        pytest.fail("Value {} was successfully validated".format(value))
    except TypeError:
        pass


def test_int64_list(api):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    mixed_object.integer64_1 = 10
    mixed_object.integer64_list = [1, 2]
    data = config.serialize("dict")
    assert isinstance(
        data["native_features"]["mixed_object"]["integer64_1"], str
    )
    for d in data["native_features"]["mixed_object"]["integer64_list"]:
        assert isinstance(d, str)
    config.deserialize(data)
    assert isinstance(mixed_object.integer64_1, int)
    for d in mixed_object.integer64_list:
        assert isinstance(d, int)


def test_int64_type(api):
    config = api.test_config()

    mixed_object = config.native_features.mixed_object
    value1 = mixed_object._TYPES.get("integer64_2")
    value2 = mixed_object._TYPES.get("integer64_list")
    assert value1.get("format") is not None
    assert value2.get("itemformat") is not None
    assert value1.get("format") == "int64"
    assert value2.get("itemformat") == "int64"
    mixed_object.integer64_2 = 100
    mixed_object.integer64_list = [2000]
    data = config.serialize("dict")
    assert isinstance(
        data["native_features"]["mixed_object"]["integer64_2"], str
    )
    assert isinstance(
        data["native_features"]["mixed_object"]["integer64_list"][0], str
    )

    config = module.Api().test_config()
    config.deserialize(data)
    assert isinstance(mixed_object.integer64_2, int)
    assert isinstance(mixed_object.integer64_list[0], int)


def test_integer_datatypes(api):

    config = api.test_config()

    number_type_object = config.native_features.number_type_object
    value = number_type_object._TYPES.get("validate_uint32_1")
    assert value.get("format") is not None
    # TODO: Uncomment when the backward compatibility support is uncommented from openapiartprotobuf
    # and remove the existing validation with int64
    # assert value.get("format") == "uint32"
    assert value.get("format") == "int64"
    number_type_object.validate_uint32_1 = 4261412865

    value = number_type_object._TYPES.get("validate_uint32_2")
    assert value.get("format") is not None
    assert value.get("format") == "uint32"
    number_type_object.validate_uint32_2 = 4261412865

    value = number_type_object._TYPES.get("validate_uint64_1")
    assert value.get("format") is not None
    # TODO: Uncomment when the backward compatibility support is uncommented from openapiartprotobuf
    # and remove the existing validation with int64
    # assert value.get("format") == "uint64"
    assert value.get("format") == "int64"
    number_type_object.validate_uint64_1 = 9223372036854775807

    value = number_type_object._TYPES.get("validate_uint64_2")
    assert value.get("format") is not None
    assert value.get("format") == "uint64"
    number_type_object.validate_uint64_2 = 9223372036854775807

    value = number_type_object._TYPES.get("validate_int32_1")
    # assert value.get("format") is not None
    # assert value.get("format") == "int32"
    number_type_object.validate_int32_1 = 2147483646

    value = number_type_object._TYPES.get("validate_int32_2")
    assert value.get("format") is not None
    assert value.get("format") == "int32"
    number_type_object.validate_int32_2 = 2147483647

    value = number_type_object._TYPES.get("validate_int64_1")
    assert value.get("format") is not None
    assert value.get("format") == "int64"
    number_type_object.validate_int64_1 = 9223372036854775807

    value = number_type_object._TYPES.get("validate_int64_2")
    assert value.get("format") is not None
    assert value.get("format") == "int64"
    number_type_object.validate_int64_2 = 9223372036854775807

    data = config.serialize("dict")
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_uint32_1"], str
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_uint32_2"], int
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_uint64_1"], str
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_uint64_2"], str
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_int32_1"], int
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_int32_2"], int
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_int64_1"], str
    )
    assert isinstance(
        data["native_features"]["number_type_object"]["validate_int64_2"], str
    )

    config = module.Api().test_config()
    config.deserialize(data)
    assert isinstance(number_type_object.validate_uint32_1, int)
    assert isinstance(number_type_object.validate_uint32_1, int)
    assert isinstance(number_type_object.validate_uint64_1, int)
    assert isinstance(number_type_object.validate_uint64_1, int)
    assert isinstance(number_type_object.validate_int32_1, int)
    assert isinstance(number_type_object.validate_int32_1, int)
    assert isinstance(number_type_object.validate_int64_1, int)
    assert isinstance(number_type_object.validate_int64_1, int)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
