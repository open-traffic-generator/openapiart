import sys
import os
import pytest

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
import sanity


def test_set(api):
    config = sanity.Api().prefix_config()

    # Use the new set API to set normal property values of each class
    config.set(
        ieee_802_1qbb=True,
        space_1=10,
        full_duplex_100_mb=20,
        response=config.STATUS_200,
        a="abc",
        b=0.2,
        c=20,
        d_values=[config.A],
        h=False,
        i="1010100",
        list_of_string_values=["abc", "efg"],
        list_of_integer_values=[1, 2, 3, 4],
    )
    config.required_object.set(e_a=3.1, e_b=4.1)
    config.e.set(e_a=1.1, e_b=2.1)
    config.l.set(
        ipv4="10.1.1.1",
        ipv6="::",
        mac="aa:bb:cc:11:22:33",
        string_param="abcd",
    )
    # Use the new set API to set pattern property values of each class
    config.mac_pattern.mac.increment.set(
        start="aa:aa:bb:bb:cc:cc", step="00:01:00:00:10:00", count=10
    )
    config.integer_pattern.integer.decrement.set(start=200, step=2, count=100)

    config_data = config._properties

    try:
        config.serialize()
    except Exception as e:
        pytest.fail(str(e))

    assert config_data[
        "ieee_802_1qbb"
    ], "ieee_802_1qbb value is not set correctly"
    assert config_data["space_1"] == 10, "space_1 value is not set correctly"
    assert (
        config_data["full_duplex_100_mb"] == 20
    ), "full_duplex_100_mb value is not set correctly"
    assert (
        config_data["response"] == config.STATUS_200
    ), "response value is not set correctly"
    assert config_data["a"] == "abc", "a value is not set correctly"
    assert config_data["b"] == 0.2, "b value is not set correctly"
    assert config_data["c"] == 20, "c value is not set correctly"
    d_values = [config.A]
    i = 0
    for value in config_data["d_values"]:
        assert value == d_values[i], "d values are not set correctly"
        i = i + 1
    assert not config_data["h"], "h value is not set correctly"
    assert config_data["i"] == "1010100", "i value is not set correctly"
    list_of_string_values = ["abc", "efg"]
    i = 0
    for value in config_data["list_of_string_values"]:
        assert (
            value == list_of_string_values[i]
        ), "list_of_string_values values are not set correctly"
        i = i + 1
    list_of_integer_values = [1, 2, 3, 4]
    i = 0
    for value in config_data["list_of_integer_values"]:
        assert (
            value == list_of_integer_values[i]
        ), "list_of_integer_values values are not set correctly"
        i = i + 1

    config_required_object_data = config.required_object._properties
    assert (
        config_required_object_data["e_a"] == 3.1
    ), "e_a value is not set correctly"
    assert (
        config_required_object_data["e_b"] == 4.1
    ), "e_b value is not set correctly"

    config_e_data = config.e._properties
    assert config_e_data["e_a"] == 1.1, "e_a value is not set correctly"
    assert config_e_data["e_b"] == 2.1, "e_b value is not set correctly"

    config_l_data = config.l._properties
    assert (
        config_l_data["ipv4"] == "10.1.1.1"
    ), "ipv4 value is not set correctly"
    assert config_l_data["ipv6"] == "::", "ipv6 value is not set correctly"
    assert (
        config_l_data["mac"] == "aa:bb:cc:11:22:33"
    ), "mac value is not set correctly"
    assert (
        config_l_data["string_param"] == "abcd"
    ), "string_param value is not set correctly"

    config_mac_pattern_mac_increment_data = (
        config.mac_pattern.mac.increment._properties
    )
    assert (
        config_mac_pattern_mac_increment_data["start"] == "aa:aa:bb:bb:cc:cc"
    ), "start value is not set correctly"
    assert (
        config_mac_pattern_mac_increment_data["step"] == "00:01:00:00:10:00"
    ), "step value is not set correctly"
    assert (
        config_mac_pattern_mac_increment_data["count"] == 10
    ), "count value is not set correctly"

    config_integer_pattern_integer_decrement = (
        config.integer_pattern.integer.decrement._properties
    )
    assert (
        config_integer_pattern_integer_decrement["start"] == 200
    ), "start value is not set correctly"
    assert (
        config_integer_pattern_integer_decrement["step"] == 2
    ), "step value is not set correctly"
    assert (
        config_integer_pattern_integer_decrement["count"] == 100
    ), "count value is not set correctly"

    # Negative Test 1 - Not setting required object
    new_config_1 = sanity.Api().prefix_config()
    new_config_1.set(
        ieee_802_1qbb=True,
        space_1=10,
        full_duplex_100_mb=20,
        response=config.STATUS_200,
        a="abc",
        b=0.2,
        c=20,
        d_values=[config.A],
        h=False,
        i="1010100",
        list_of_string_values=["abc", "efg"],
        list_of_integer_values=[1, 2, 3, 4],
    )
    new_config_1.e.set(e_a=1.1, e_b=2.1)
    new_config_1.l.set(
        ipv4="10.1.1.1",
        ipv6="::",
        mac="aa:bb:cc:11:22:33",
        string_param="abcd",
    )
    new_config_1.mac_pattern.mac.increment.set(
        start="aa:aa:bb:bb:cc:cc", step="00:01:00:00:10:00", count=10
    )
    new_config_1.integer_pattern.integer.decrement.set(
        start=200, step=2, count=100
    )

    try:
        new_config_1.serialize()
    except Exception as e:
        if "required_object is a mandatory property of" not in str(e):
            pytest.fail("required_object validation failed")

    # Negative Test 2 - Not setting the correct value to property a
    new_config_2 = sanity.Api().prefix_config()
    new_config_2.set(
        ieee_802_1qbb=True,
        space_1=10,
        full_duplex_100_mb=20,
        response=config.STATUS_200,
        a=10,
        b=0.2,
        c=20,
        d_values=[config.A],
        h=False,
        i="1010100",
        list_of_string_values=["abc", "efg"],
        list_of_integer_values=[1, 2, 3, 4],
    )
    new_config_2.required_object.set(e_a=3.1, e_b=4.1)
    new_config_2.e.set(e_a=1.1, e_b=2.1)
    new_config_2.l.set(
        ipv4="10.1.1.1",
        ipv6="::",
        mac="aa:bb:cc:11:22:33",
        string_param="abcd",
    )
    new_config_2.mac_pattern.mac.increment.set(
        start="aa:aa:bb:bb:cc:cc", step="00:01:00:00:10:00", count=10
    )
    new_config_2.integer_pattern.integer.decrement.set(
        start=200, step=2, count=100
    )

    try:
        new_config_2.serialize()
    except Exception as e:
        if "property a shall be of type" not in str(e):
            pytest.fail("a validation failed")
