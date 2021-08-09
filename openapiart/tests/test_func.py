import pytest


def test_defaults_and_stale_choices(api):
    config = api.prefix_config()
    config.f
    f = {
        'choice': 'f_a', 'f_a': 'some string'
    }
    assert config.f.serialize(config.DICT) == f

    f = {
        'choice': 'f_b', 'f_b': 4.25
    }
    config = api.prefix_config()
    config.f.deserialize(f)
    assert config.f._properties == f
    assert config.f.serialize(config.DICT) == f


def test_required(api):
    config = api.prefix_config()
    try:
        config.mandatory
        config.mandatory.serialize()
        pytest.fail("config got validated")
    except ValueError:
        pass


def test_serialize_deserialize(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.d_values = [config.A, config.B, config.C]
    config.level.l1_p1.l2_p1.l3_p1 = "test"
    config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
    api.set_config(config)
    _config = api.get_config()
    assert config.serialize(config.DICT) == _config.serialize(_config.DICT)


def test_x_include(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.e.e_a = 0.1
    config.e.e_b = 0.2
    assert "name" in config.e._TYPES
    assert "m_param1" in config.e._TYPES
    assert "m_param2" in config.e._TYPES


@pytest.mark.parametrize('ipv4', ["1.1.1.1", "01.002.003.4", "0.0.0.0"])
def test_x_pattern_ipv4_good_value(default_config, ipv4):
    default_config.ipv4_pattern.ipv4.value = ipv4
    try:
        default_config.ipv4_pattern.serialize(default_config.DICT)
    except Exception:
        pytest.fail("ipv4 pattern failed to serialize")


@pytest.mark.parametrize('ipv6', ["::", " :: ", "1::", "::1", "abcd::abcd", "abcd:abcd:abcd:abcd:abcd:abcd:abcd:abcd"])
def test_x_pattern_ipv6_good_value(default_config, ipv6):
    default_config.ipv6_pattern.ipv6.value = ipv6
    try:
        default_config.ipv6_pattern.serialize(default_config.DICT)
    except Exception:
        pytest.fail("ipv6 pattern failed to serialize")


@pytest.mark.parametrize('mac', ["ab:ab:ab:ab:ab:ab"])
def test_x_pattern_mac_good_value(default_config, mac):
    default_config.mac_pattern.mac.value = mac
    try:
        default_config.mac_pattern.serialize(default_config.DICT)
    except Exception:
        pytest.fail("mac pattern failed to serialize")


@pytest.mark.parametrize('integer', [0, 100, 255])
def test_x_pattern_integer_good_value(default_config, integer):
    default_config.integer_pattern.integer.value = integer
    try:
        default_config.integer_pattern.serialize(default_config.DICT)
    except Exception:
        pytest.fail("integer pattern failed to serialize")


@pytest.mark.parametrize("ipv4", [["1.1.1.1", "-255.-255.-255.-255"]])
def test_x_pattern_ipv4_good_and_bad_list(default_config, ipv4):
    default_config.ipv4_pattern.ipv4.values = ipv4
    try:
        default_config.ipv4_pattern.serialize(default_config.DICT)
        pytest.fail("ipv4 values got serialize")
    except TypeError as e:
        if "['-255.-255.-255.-255']" not in str(e):
            pytest.fail("Invalid ipv4 list is not proper in error message")


@pytest.mark.parametrize("ipv6", [["::", ":", "abcd::abcd::"]])
def test_x_pattern_ipv6_good_and_bad_list(default_config, ipv6):
    default_config.ipv6_pattern.ipv6.values = ipv6
    try:
        default_config.ipv6_pattern.serialize(default_config.DICT)
        pytest.fail("ipv6 values got serialize")
    except TypeError as e:
        if "[':', 'abcd::abcd::']" not in str(e):
            pytest.fail("Invalid ipv6 list is not proper in error message")


@pytest.mark.parametrize("mac", [[":", "abcd::abcd::", "ab:ab:ab:ab:ab:ab"]])
def test_x_pattern_mac_good_and_bad_list(default_config, mac):
    default_config.mac_pattern.mac.values = mac
    try:
        default_config.mac_pattern.serialize(default_config.DICT)
        pytest.fail("mac values got serialize")
    except TypeError as e:
        if "[':', 'abcd::abcd::']" not in str(e):
            pytest.fail("Invalid mac list is not proper in error message")


@pytest.mark.parametrize("integer", [["abcd::abcd::", 10, 3, 255, 256, "ab:ab:ab:ab:ab:ab"]])
def test_x_pattern_integer_good_and_bad_list(default_config, integer):
    default_config.integer_pattern.integer.values = integer
    try:
        default_config.integer_pattern.serialize(default_config.DICT)
        pytest.fail("integer values got serialize")
    except TypeError as e:
        if "['abcd::abcd::', 256, 'ab:ab:ab:ab:ab:ab']" not in str(e):
            pytest.fail("Invalid integer list is not proper in error message")


@pytest.mark.parametrize("direction", ["increment", "decrement"])
@pytest.mark.parametrize("index", list(range(4)))
def test_x_pattern_good_inc_dec(default_config, index, direction):
    start = ['1.1.1.1', '2000::1', '00:00:00:00:00:0a', 10]
    step = ['0.1.1.0', '::1', '00:00:00:00:00:01', 1]
    count = [100, 100, 100, 100]
    enum = ['ipv4', 'ipv6', 'mac', 'integer']
    enum_obj = getattr(getattr(default_config, '%s_pattern' % enum[index]), enum[index])
    dir_obj = getattr(enum_obj, direction)
    dir_obj.start = start[index]
    dir_obj.step = step[index]
    dir_obj.count = count[index]
    try:
        default_config.serialize(default_config.DICT)
    except TypeError as e:
        pytest.fail("%s with %s Failed to serialize" % (enum, direction))


@pytest.mark.parametrize("direction", ["increment", "decrement"])
@pytest.mark.parametrize("index", list(range(4)))
def test_x_pattern_bad_inc_dec(default_config, index, direction):
    start = ['1.1.1.1', '2000::1', '00:00:00:00:00:00', 10]
    step = [1, 10, '00:00:00:00:0001', -1]
    count = [100, 100, 100, 100]
    enum = ['ipv4', 'ipv6', 'mac', 'integer']
    enum_obj = getattr(getattr(default_config, '%s_pattern' % enum[index]), enum[index])
    dir_obj = getattr(enum_obj, direction)
    dir_obj.start = start[index]
    dir_obj.step = step[index]
    dir_obj.count = count[index]
    try:
        default_config.serialize(default_config.DICT)
        pytest.fail("%s with %s got serialized" % (enum, direction))
    except TypeError as e:
        print(e)