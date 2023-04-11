import importlib

module = importlib.import_module("sanity")


def test_integer_datatypes(default_config):
    value = default_config._TYPES.get("validate_uint32_1")
    assert value.get("format") is not None
    # TODO: Uncomment when the backward compatibility support is uncommented from openapiartprotobuf
    # and remove the existing validation with int64
    # assert value.get("format") == "uint32"
    assert value.get("format") == "int64"
    default_config.validate_uint32_1 = 4261412865

    value = default_config._TYPES.get("validate_uint32_2")
    assert value.get("format") is not None
    assert value.get("format") == "uint32"
    default_config.validate_uint32_2 = 4261412865

    value = default_config._TYPES.get("validate_uint64_1")
    assert value.get("format") is not None
    # TODO: Uncomment when the backward compatibility support is uncommented from openapiartprotobuf
    # and remove the existing validation with int64
    # assert value.get("format") == "uint64"
    assert value.get("format") == "int64"
    default_config.validate_uint64_1 = 9223372036854775807

    value = default_config._TYPES.get("validate_uint64_2")
    assert value.get("format") is not None
    assert value.get("format") == "uint64"
    default_config.validate_uint64_2 = 18446744073709551614

    value = default_config._TYPES.get("validate_int32_1")
    # assert value.get("format") is not None
    # assert value.get("format") == "int32"
    default_config.validate_int32_1 = 2147483646

    value = default_config._TYPES.get("validate_int32_2")
    assert value.get("format") is not None
    assert value.get("format") == "int32"
    default_config.validate_int32_2 = 2147483647

    value = default_config._TYPES.get("validate_int64_1")
    assert value.get("format") is not None
    assert value.get("format") == "int64"
    default_config.validate_int64_1 = 9223372036854775807

    value = default_config._TYPES.get("validate_int64_2")
    assert value.get("format") is not None
    assert value.get("format") == "int64"
    default_config.validate_int64_2 = 9223372036854775807

    data = default_config.serialize("dict")
    assert isinstance(data["validate_uint32_1"], str)
    assert isinstance(data["validate_uint32_2"], int)
    assert isinstance(data["validate_uint64_1"], str)
    assert isinstance(data["validate_uint64_2"], str)
    assert isinstance(data["validate_int32_1"], int)
    assert isinstance(data["validate_int32_2"], int)
    assert isinstance(data["validate_int64_1"], str)
    assert isinstance(data["validate_int64_2"], str)

    config = module.Api().prefix_config()
    config.deserialize(data)
    assert isinstance(config.validate_uint32_1, int)
    assert isinstance(config.validate_uint32_1, int)
    assert isinstance(config.validate_uint64_1, int)
    assert isinstance(config.validate_uint64_1, int)
    assert isinstance(config.validate_int32_1, int)
    assert isinstance(config.validate_int32_1, int)
    assert isinstance(config.validate_int64_1, int)
    assert isinstance(config.validate_int64_1, int)
