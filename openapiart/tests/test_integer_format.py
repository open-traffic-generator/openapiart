import importlib
import pytest

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

    # negative test case
    error_msg = "Invalid -12319 format, expected int64 at <class 'sanity.sanity.PrefixConfig'>"
    default_config.full_duplex_100_mb = -12319
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value

    error_msg = "property integer64_list shall be of type <class 'list'> at <class 'sanity.sanity.PrefixConfig'> \n [-2000] are not valid"
    default_config.full_duplex_100_mb = 100
    default_config.integer64_list = [-2000]
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value


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


def test_format_signed_integer_pattern(default_config):
    pat = default_config.signed_integer_pattern.integer
    val = pat._TYPES.get("value")
    assert val.get("format") == "int32"
    assert val.get("minimum") == -128
    assert val.get("maximum") == 127
    inc = pat.increment
    assert inc._TYPES.get("start").get("format") == "int32"
    assert inc._TYPES.get("step").get("format") == "int32"
    assert inc._TYPES.get("count").get("format") == "int32"
    dec = pat.decrement
    assert dec._TYPES.get("start").get("format") == "int32"
    assert dec._TYPES.get("step").get("format") == "int32"
    assert dec._TYPES.get("count").get("format") == "int32"
    pat.value = -456
    error_msg = (
        "got -456 of type <class 'int'> , expected min -128, expected max 127"
    )
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value


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


def test_int32(default_config):
    value1 = default_config._TYPES.get("int32_param")
    value2 = default_config._TYPES.get("int32_list_param")
    assert value1.get("format") is not None
    assert value2.get("itemformat") is not None
    assert value1.get("format") == "int32"
    assert value2.get("itemformat") == "int32"

    default_config.int32_param = -100
    default_config.int32_list_param = [-2000]
    data = default_config.serialize("dict")

    config = module.Api().prefix_config()
    config.deserialize(data)
    assert isinstance(config.int32_param, int)
    assert isinstance(config.int32_list_param[0], int)

    # negative test case
    error_msg = "Invalid 2147483648 format, expected int32 at <class 'sanity.sanity.PrefixConfig'>"
    default_config.int32_param = 2147483648
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value

    error_msg = "property int32_list_param shall be of type <class 'list'> at <class 'sanity.sanity.PrefixConfig'> \n [-2147483649] are not valid"
    default_config.int32_param = 100
    default_config.int32_list_param = [-2147483649]
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value


def test_uint32(default_config):
    value1 = default_config._TYPES.get("uint32_param")
    value2 = default_config._TYPES.get("uint32_list_param")
    assert value1.get("format") is not None
    assert value2.get("itemformat") is not None
    assert value1.get("format") == "uint32"
    assert value2.get("itemformat") == "uint32"

    default_config.uint32_param = 2147483648
    default_config.uint32_list_param = [2147483648]
    data = default_config.serialize("dict")

    config = module.Api().prefix_config()
    config.deserialize(data)
    assert isinstance(config.uint32_param, int)
    assert isinstance(config.uint32_list_param[0], int)

    # negative test case
    error_msg = "Invalid -1 format, expected uint32 at <class 'sanity.sanity.PrefixConfig'>"
    default_config.uint32_param = -1
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value

    error_msg = "property uint32_list_param shall be of type <class 'list'> at <class 'sanity.sanity.PrefixConfig'> \n [4294967300] are not valid"
    default_config.uint32_param = 100
    default_config.uint32_list_param = [4294967300]
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value


def test_uint64(default_config):
    value1 = default_config._TYPES.get("uint64_param")
    value2 = default_config._TYPES.get("uint64_list_param")
    assert value1.get("format") is not None
    assert value2.get("itemformat") is not None
    assert value1.get("format") == "uint64"
    assert value2.get("itemformat") == "uint64"

    default_config.uint64_param = 42949672955
    default_config.uint64_list_param = [42949672955]
    data = default_config.serialize("dict")
    assert isinstance(data["uint64_param"], str)
    assert isinstance(data["uint64_list_param"][0], str)

    config = module.Api().prefix_config()
    config.deserialize(data)
    assert isinstance(config.uint64_param, int)
    assert isinstance(config.uint64_list_param[0], int)

    # negative test case
    error_msg = "Invalid -1 format, expected uint64 at <class 'sanity.sanity.PrefixConfig'>"
    default_config.uint64_param = -1
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert error_msg in error_value


def test_auto_assign_for_integers(default_config):
    value5 = default_config._TYPES.get("auto_int32_param")
    value6 = default_config._TYPES.get("auto_int32_list_param")

    assert value5.get("format") == "int32"
    assert value6.get("itemformat") == "int32"

    default_config.auto_int32_param = 5675
    default_config.auto_int32_list_param = [5675]
    # check serialise and deserialise
    data = default_config.serialize("dict")
    config = module.Api().prefix_config()
    config.deserialize(data)
