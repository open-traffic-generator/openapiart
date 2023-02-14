import importlib
import pytest

module = importlib.import_module("sanity")


@pytest.mark.skip(reason="shall be restored")
def test_x_constraints(default_config):
    default_config.name = "pc1"
    default_config.w_list.wobject(w_name="wObj1")
    default_config.w_list.wobject(w_name="wObj2")
    default_config.z_object.name = "zObj"
    # set the non existing name to y_object
    default_config.y_object.y_name = "wObj3"
    try:
        default_config.validate()
        pytest.fail("validation passed when setting y_name with wObj3")
    except Exception as err:
        if "wObj3 is not a valid type of" not in str(err):
            pytest.fail("Exception is not valid")

    # set the name with invalid object name
    default_config.y_object.y_name = "pc1"
    try:
        default_config.validate()
        pytest.fail("validation passed when setting y_name with pc1")
    except Exception as err:
        if "pc1 is not a valid type of" not in str(err):
            pytest.fail("Exception is not valid")

    # validate with valid data
    default_config.y_object.y_name = "wObj1"
    default_config.validate()

    # serialize with non existing name
    default_config.y_object.y_name = "wObj3"
    try:
        data = default_config.serialize("dict")
        pytest.fail("validation passed at serialize with wObj3")
    except Exception as err:
        if "wObj3 is not a valid type of" not in str(err):
            pytest.fail("Exception not valid at serialize wObj3")

    # serialize with valid data
    default_config.y_object.y_name = "wObj1"
    data = default_config.serialize("dict")

    # deserialize with valid data
    config = module.Api().prefix_config()
    config.deserialize(data)

    # deserialize with invalid name
    data["y_object"]["y_name"] = "pc1"
    try:
        config.deserialize(data)
        pytest.fail("deserialize passed with pc1 name")
    except Exception as err:
        if "pc1 is not a valid type of" not in str(err):
            pytest.fail("Exception is not valid at deserialize with pc1")
