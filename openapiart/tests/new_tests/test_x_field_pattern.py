import pytest


@pytest.mark.parametrize("ipv4", ["1.1.1.1", "01.002.003.4", "0.0.0.0"])
def test_pattern_ipv4_good_value(api, ipv4):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.ipv4_pattern.ipv4.value = ipv4
    try:
        pattern_obj.ipv4_pattern.serialize(config.DICT)
    except Exception:
        pytest.fail("ipv4 pattern failed to serialize")


@pytest.mark.parametrize(
    "ipv6",
    [
        "::",
        " :: ",
        "1::",
        "::1",
        "abcd::abcd",
        "abcd:abcd:abcd:abcd:abcd:abcd:abcd:abcd",
    ],
)
def test_pattern_ipv6_good_value(api, ipv6):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.ipv6_pattern.ipv6.value = ipv6
    try:
        pattern_obj.ipv6_pattern.serialize(config.DICT)
    except Exception:
        pytest.fail("ipv6 pattern failed to serialize")


@pytest.mark.parametrize("mac", ["ab:ab:ab:ab:ab:ab"])
def test_pattern_mac_good_value(api, mac):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.mac_pattern.mac.value = mac
    try:
        pattern_obj.mac_pattern.serialize(config.DICT)
    except Exception:
        pytest.fail("mac pattern failed to serialize")


@pytest.mark.parametrize("integer", [0, 100, 255])
def test_pattern_integer_good_value(api, integer):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.integer_pattern.integer.value = integer
    try:
        pattern_obj.integer_pattern.serialize(pattern_obj.DICT)
    except Exception:
        pytest.fail("integer pattern failed to serialize")


@pytest.mark.parametrize("ipv4", [["1.1.1.1", "-255.-255.-255.-255"]])
def test_pattern_ipv4_good_and_bad_list(api, ipv4):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.ipv4_pattern.ipv4.values = ipv4
    try:
        pattern_obj.ipv4_pattern.serialize(config.DICT)
        pytest.fail("ipv4 values got serialize")
    except TypeError as e:
        if "['-255.-255.-255.-255']" not in str(e):
            pytest.fail("Invalid ipv4 list is not proper in error message")


@pytest.mark.parametrize("ipv6", [["::", ":", "abcd::abcd::"]])
def test_pattern_ipv6_good_and_bad_list(api, ipv6):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.ipv6_pattern.ipv6.values = ipv6
    try:
        pattern_obj.ipv6_pattern.serialize(config.DICT)
        pytest.fail("ipv6 values got serialize")
    except TypeError as e:
        if "[':', 'abcd::abcd::']" not in str(e):
            pytest.fail("Invalid ipv6 list is not proper in error message")


@pytest.mark.parametrize("mac", [[":", "abcd::abcd::", "ab:ab:ab:ab:ab:ab"]])
def test_pattern_mac_good_and_bad_list(api, mac):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.mac_pattern.mac.values = mac
    try:
        pattern_obj.mac_pattern.serialize(config.DICT)
        pytest.fail("mac values got serialize")
    except TypeError as e:
        if "[':', 'abcd::abcd::']" not in str(e):
            pytest.fail("Invalid mac list is not proper in error message")


@pytest.mark.parametrize(
    "integer", [["abcd::abcd::", 10, 3, 255, 256, "ab:ab:ab:ab:ab:ab"]]
)
def test_pattern_integer_good_and_bad_list(api, integer):
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    pattern_obj.integer_pattern.integer.values = integer
    try:
        pattern_obj.integer_pattern.serialize(config.DICT)
        pytest.fail("integer values got serialize")
    except TypeError as e:
        if "['abcd::abcd::', 256, 'ab:ab:ab:ab:ab:ab']" not in str(e):
            pytest.fail("Invalid integer list is not proper in error message")


@pytest.mark.parametrize("direction", ["increment", "decrement"])
@pytest.mark.parametrize("index", list(range(4)))
def test_pattern_good_inc_dec(api, index, direction):
    start = ["1.1.1.1", "2000::1", "00:00:00:00:00:0a", 10]
    step = ["0.1.1.0", "::1", "00:00:00:00:00:01", 1]
    count = [100, 100, 100, 100]
    enum = ["ipv4", "ipv6", "mac", "integer"]
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    enum_obj = getattr(
        getattr(pattern_obj, "%s_pattern" % enum[index]), enum[index]
    )
    dir_obj = getattr(enum_obj, direction)
    dir_obj.start = start[index]
    dir_obj.step = step[index]
    dir_obj.count = count[index]
    try:
        pattern_obj.serialize(pattern_obj.DICT)
    except TypeError:
        pytest.fail("%s with %s Failed to serialize" % (enum, direction))


@pytest.mark.parametrize("direction", ["increment", "decrement"])
@pytest.mark.parametrize("index", list(range(4)))
def test_pattern_bad_inc_dec(api, index, direction):
    start = ["1.1.1.1", "2000::1", "00:00:00:00:00:00", 10]
    step = [1, 10, "00:00:00:00:0001", -1]
    count = [100, 100, 100, 100]
    enum = ["ipv4", "ipv6", "mac", "integer"]
    config = api.test_config()
    pattern_obj = config.extended_features.x_field_pattern_object
    enum_obj = getattr(
        getattr(pattern_obj, "%s_pattern" % enum[index]), enum[index]
    )
    dir_obj = getattr(enum_obj, direction)
    dir_obj.start = start[index]
    dir_obj.step = step[index]
    dir_obj.count = count[index]
    try:
        pattern_obj.serialize(pattern_obj.DICT)
        pytest.fail("%s with %s got serialized" % (enum, direction))
    except TypeError as e:
        print(e)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
