import pytest


def test_constraints_and_unique(config):
    config.z_object.name = "abc"
    config.y_object.y_name = "123"
    config.x_list.zobject(name="abc")
    try:
        config.validate()
    except Exception as e:
        if "abc already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")
    try:
        config.serialize()
    except Exception as e:
        if "abc already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")

    config.x_list[0].name = "bcd"
    config.y_object.y_name = "bcd"
    data = config.serialize(config.DICT)
    data["y_object"]["y_name"] = "123"
    data["z_object"]["name"] = "bcd"
    try:
        config.deserialize(data)
    except Exception as e:
        if "bcd already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")
