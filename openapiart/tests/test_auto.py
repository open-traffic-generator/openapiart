import pytest


def test_auto(config):
    try:
        config.auto_field_test.auto = 10
        pytest.fail("able to set the auto field")
    except Exception:
        pass

    assert config.auto_field_test.auto == 0

    dt = config.serialize(config.DICT)

    assert dt.get("auto_field_test", {}).get("choice") == "auto"
    assert dt.get("auto_field_test", {}).get("auto") == 0

    config.auto_field_test.value = 20
    assert config.auto_field_test.choice == "value"

    config.auto_field_test.auto
    assert config.auto_field_test.choice == "auto"


def test_auto_dhcp(config):
    try:
        config.auto_dhcp_pattern.dhcp.auto_dhcp = "asd"
        pytest.fail("able to set the auto field")
    except Exception:
        pass

    assert config.auto_dhcp_pattern.dhcp.auto_dhcp == "0.0.0.0"

    dt = config.serialize(config.DICT)

    assert (
        dt.get("auto_dhcp_pattern", {}).get("dhcp").get("choice")
        == "auto_dhcp"
    )
    assert (
        dt.get("auto_dhcp_pattern", {}).get("dhcp").get("auto_dhcp")
        == "0.0.0.0"
    )

    config.auto_dhcp_pattern.dhcp.value = 20
    assert config.auto_dhcp_pattern.dhcp.choice == "value"

    config.auto_dhcp_pattern.dhcp.auto_dhcp
    assert config.auto_dhcp_pattern.dhcp.choice == "auto_dhcp"
