import pytest


def test_constraints_and_unique(config):
    config.name = "global_unique"
    config.w_list.wobject(w_name="global_unique")

    try:
        config.validate()
    except Exception as e:
        if "global_unique already exists" not in str(e):
            pytest.fail("global_unique validation failed")
    config.name = "global_unique1"

    config.z_object.name = "local_unique"
    config.y_object.y_name = "123"
    config.x_list.zobject(name="local_unique")
    try:
        config.validate()
    except Exception as e:
        if "local_unique already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")
    try:
        config.serialize()
    except Exception as e:
        if "local_unique already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")

    config.x_list[0].name = "local_unique_1"
    config.y_object.y_name = "local_unique_1"
    data = config.serialize(config.DICT)
    data["y_object"]["y_name"] = "local_unique_decode"
    data["z_object"]["name"] = "local_unique_decode"
    try:
        config.deserialize(data)
    except Exception as e:
        if "local_unique_decode already exists" not in str(
            e
        ) or "y_name is not a valid type" not in str(e):
            pytest.fail("validation failed")
    
    config.name = "global_local_same_name_check"
    config.x_list.zobject(name="global_local_same_name_check")
    try:
        config.validate()
    except Exception as e:
        pytest.fail("validation failed\n {e}".format(e=e))
