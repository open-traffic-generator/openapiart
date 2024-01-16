import importlib
import pytest

module = importlib.import_module("sanity")


def test_oid_in_normal_attr(default_config):
    default_config.m_object.string_param = "abc"
    default_config.m_object.integer = 44
    default_config.m_object.float = 1.23
    default_config.m_object.double = 3.45
    default_config.m_object.mac = "00:00:00:00:00:00"
    default_config.m_object.ipv4 = "1.2.3.4"
    default_config.m_object.ipv6 = "::"
    default_config.m_object.hex = "0f"
    default_config.m_object.oid = "."

    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "Invalid . format, expected oid" in error_value

    default_config.m_object.oid = "1."
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "Invalid 1. format, expected oid" in error_value

    default_config.m_object.oid = "0.42949672967.22.44.55"
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "Invalid 0.42949672967.22.44.55 format, expected oid" in error_value

    default_config.m_object.oid = "1.2.abcd"
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "Invalid 1.2.abcd format, expected oid" in error_value

    default_config.m_object.oid = "1.2.3.4.5"
    data = default_config.serialize("dict")

    data["m_object"]["oid"] = "."
    with pytest.raises(Exception) as execinfo:
        default_config.deserialize(data)
    error_value = execinfo.value.args[0]
    assert "Invalid . format, expected oid" in error_value

    data["m_object"]["oid"] = "1."
    with pytest.raises(Exception) as execinfo:
        default_config.deserialize(data)
    error_value = execinfo.value.args[0]
    assert "Invalid 1. format, expected oid" in error_value

    data["m_object"]["oid"] = "0.42949672967.22.44.55"
    with pytest.raises(Exception) as execinfo:
        default_config.deserialize(data)
    error_value = execinfo.value.args[0]
    assert "Invalid 0.42949672967.22.44.55 format, expected oid" in error_value

    data["m_object"]["oid"] = "1.2.3.4"
    default_config.deserialize(data)


def test_oid_x_field_pattern(default_config):
    default_config.oid_pattern.oid.value = "1."

    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "Invalid 1. format, expected oid" in error_value

    default_config.oid_pattern.oid.value = "1.2.3.4"
    data = default_config.serialize("dict")

    data["oid_pattern"]["oid"]["value"] = "1."
    with pytest.raises(Exception) as execinfo:
        default_config.deserialize(data)
    error_value = execinfo.value.args[0]
    assert "Invalid 1. format, expected oid" in error_value

    default_config.oid_pattern.oid.values = [
        "0.42949672967.22.44.55",
        "1.2.3.4.5",
        "0.",
        ".",
        "44.55555.77777",
    ]
    with pytest.raises(Exception) as execinfo:
        default_config.serialize("dict")
    error_value = execinfo.value.args[0]
    assert "['0.42949672967.22.44.55', '0.', '.'] are not valid" in error_value

    default_config.oid_pattern.oid.values = ["1.2.3.4.5", "66.789.6789.56789"]
    data = default_config.serialize("dict")

    data["oid_pattern"]["oid"]["values"] = [
        "0.42949672967.22.44.55",
        "1.2.3.4.5",
        "0.",
        ".",
        "44.55555.77777",
    ]
    with pytest.raises(Exception) as execinfo:
        default_config.deserialize(data)
    error_value = execinfo.value.args[0]
    assert "['0.42949672967.22.44.55', '0.', '.'] are not valid" in error_value

    data["oid_pattern"]["oid"]["values"] = ["1.2.3.4.5", "66.789.6789.56789"]
    default_config.deserialize(data)
