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
