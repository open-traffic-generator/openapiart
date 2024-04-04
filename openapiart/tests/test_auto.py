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


def test_auto_hierarchy(config):
    at = config.auto_pattern.auto_ip
    try:
        at.auto = 10
        pytest.fail("able to set the auto field")
    except Exception:
        pass

    dt = config.serialize(config.DICT)

    assert dt.get("auto_pattern", {}).get("auto_ip").get("choice") == "value"

    at.auto.choice = "static"
    assert config.auto_pattern.auto_ip.auto.choice == "static"

    at.auto.choice = "dhcp"
    assert config.auto_pattern.auto_ip.auto.choice == "dhcp"

    dt = config.serialize(config.DICT)
    assert (
        dt.get("auto_pattern", {}).get("auto_ip").get("auto").get("choice")
        == "dhcp"
    )

    config.deserialize(dt)
    assert config.auto_pattern.auto_ip.choice == "auto"
    assert config.auto_pattern.auto_ip.auto.choice == "dhcp"
