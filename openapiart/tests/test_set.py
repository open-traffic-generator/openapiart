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

    if config_data["ieee_802_1qbb"] != True:
        raise ValueError("ieee_802_1qbb value is not set correctly")
    if config_data["space_1"] != 10:
        raise ValueError("space_1 value is not set correctly")
    if config_data["full_duplex_100_mb"] != 20:
        raise ValueError("full_duplex_100_mb value is not set correctly")
    if config_data["response"] != config.STATUS_200:
        raise ValueError("response value is not set correctly")
    if config_data["a"] != "abc":
        raise ValueError("a value is not set correctly")
    if config_data["b"] != 0.2:
        raise ValueError("b value is not set correctly")
    if config_data["c"] != 20:
        raise ValueError("c value is not set correctly")
    d_values = [config.A]
    i = 0
    for value in config_data["d_values"]:
        if value != d_values[i]:
            raise ValueError("d values are not set correctly")
        i = i + 1
    if config_data["h"] != False:
        raise ValueError("h value is not set correctly")
    if config_data["i"] != "1010100":
        raise ValueError("i value is not set correctly")
    list_of_string_values = ["abc", "efg"]
    i = 0
    for value in config_data["list_of_string_values"]:
        if value != list_of_string_values[i]:
            raise ValueError(
                "list_of_string_values values are not set correctly"
            )
        i = i + 1
    list_of_integer_values = [1, 2, 3, 4]
    i = 0
    for value in config_data["list_of_integer_values"]:
        if value != list_of_integer_values[i]:
            raise ValueError(
                "list_of_integer_values values are not set correctly"
            )
        i = i + 1

    config_e_data = config.e._properties
    if config_e_data["e_a"] != 1.1:
        raise ValueError("e_a value is not set correctly")
    if config_e_data["e_b"] != 2.1:
        raise ValueError("e_b value is not set correctly")

    config_l_data = config.l._properties
    if config_l_data["ipv4"] != "10.1.1.1":
        raise ValueError("ipv4 value is not set correctly")
    if config_l_data["ipv6"] != "::":
        raise ValueError("ipv6 value is not set correctly")
    if config_l_data["mac"] != "aa:bb:cc:11:22:33":
        raise ValueError("mac value is not set correctly")
    if config_l_data["string_param"] != "abcd":
        raise ValueError("string_param value is not set correctly")

    config_mac_pattern_mac_increment_data = (
        config.mac_pattern.mac.increment._properties
    )
    if config_mac_pattern_mac_increment_data["start"] != "aa:aa:bb:bb:cc:cc":
        raise ValueError("start value is not set correctly")
    if config_mac_pattern_mac_increment_data["step"] != "00:01:00:00:10:00":
        raise ValueError("step value is not set correctly")
    if config_mac_pattern_mac_increment_data["count"] != 10:
        raise ValueError("count value is not set correctly")

    config_integer_pattern_integer_decrement = (
        config.integer_pattern.integer.decrement._properties
    )
    if config_integer_pattern_integer_decrement["start"] != 200:
        raise ValueError("start value is not set correctly")
    if config_integer_pattern_integer_decrement["step"] != 2:
        raise ValueError("step value is not set correctly")
    if config_integer_pattern_integer_decrement["count"] != 100:
        raise ValueError("count value is not set correctly")
