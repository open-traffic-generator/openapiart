import pytest


def test_field_uds_default(default_config):
    mac = default_config.field_uds_mac.mac
    assert mac._TYPES["value"]["format"] == "mac"
    assert mac._TYPES["value"]["type"] == str
    assert mac._TYPES["mask"]["format"] == "hex"
    assert mac._TYPES["mask"]["type"] == str
    assert mac._TYPES["mask"]["minLength"] == 1
    assert mac._TYPES["mask"]["maxLength"] == 12
    assert mac._DEFAULTS["value"] == "00:00:00:00:00:00"
    assert mac._DEFAULTS["mask"] == "ffffffffffff"

    ipv4 = default_config.field_uds_ipv4.ipv4
    assert ipv4._TYPES["value"]["format"] == "ipv4"
    assert ipv4._TYPES["value"]["type"] == str
    assert ipv4._TYPES["mask"]["format"] == "hex"
    assert ipv4._TYPES["mask"]["type"] == str
    assert ipv4._TYPES["mask"]["minLength"] == 1
    assert ipv4._TYPES["mask"]["maxLength"] == 8
    assert ipv4._DEFAULTS["value"] == "0.0.0.0"
    assert ipv4._DEFAULTS["mask"] == "ffffffff"

    ipv6 = default_config.field_uds_ipv6.ipv6
    assert ipv6._TYPES["value"]["format"] == "ipv6"
    assert ipv6._TYPES["value"]["type"] == str
    assert ipv6._TYPES["mask"]["format"] == "hex"
    assert ipv6._TYPES["mask"]["type"] == str
    assert ipv6._TYPES["mask"]["minLength"] == 1
    assert ipv6._TYPES["mask"]["maxLength"] == 32
    assert ipv6._DEFAULTS["value"] == "::0"
    assert ipv6._DEFAULTS["mask"] == "ffffffffffffffffffffffffffffffff"

    integer = default_config.field_uds_int.integer
    assert integer._TYPES["value"]["format"] == "uint32"
    assert integer._TYPES["value"]["type"] == int
    assert integer._TYPES["value"]["maximum"] == 3
    assert integer._TYPES["mask"]["format"] == "hex"
    assert integer._TYPES["mask"]["type"] == str
    assert integer._TYPES["mask"]["minLength"] == 1
    assert integer._TYPES["mask"]["maxLength"] == 2
    assert integer._DEFAULTS["value"] == 0
    assert integer._DEFAULTS["mask"] == "ff"

    default_config.serialize()


def test_field_uds_restrictions(default_config):
    default_config.field_uds_mac.mac.mask = "zzz"

    with pytest.raises(Exception) as execinfo:
        default_config.serialize()
    error_value = execinfo.value.args[0]
    assert (
        "Invalid zzz format, expected hex at <class 'sanity.sanity.FilterMacUdsPatternMac'>"
        in error_value
    )
    default_config.field_uds_mac.mac.mask = "fff"
    default_config.field_uds_ipv4.ipv4.mask = "ghi"
    with pytest.raises(Exception) as execinfo:
        default_config.serialize()
    error_value = execinfo.value.args[0]
    assert (
        "Invalid ghi format, expected hex at <class 'sanity.sanity.FilterIpv4UdsPatternIpv4'>"
        in error_value
    )

    default_config.field_uds_ipv4.ipv4.mask = "ffffffff"
    default_config.field_uds_ipv6.ipv6.value = "1:1"
    with pytest.raises(Exception) as execinfo:
        default_config.serialize()
    error_value = execinfo.value.args[0]
    assert (
        "Invalid 1:1 format, expected ipv6 at <class 'sanity.sanity.FilterIpv6UdsPatternIpv6'>"
        in error_value
    )
    default_config.field_uds_ipv6.ipv6.value = "::3"
    default_config.field_uds_int.integer.value = -3
    with pytest.raises(Exception) as execinfo:
        default_config.serialize()
    error_value = execinfo.value.args[0]
    assert "got -3 of type <class 'int'> , expected max 3" in error_value


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
