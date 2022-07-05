import pytest


def test_constraints(config):
    config.z_object.name = "abc"
    try:
        config.y_object.y_name = "123"
    except Exception as e:
        if "y_name is not a valid type" not in str(e):
            pytest.fail("constraint validation failed")
    try:
        config.x_list.zobject(name="abc")
    except Exception as e:
        if "name with abc already exists" not in str(e):
            pytest.fail("unique validation failed")
