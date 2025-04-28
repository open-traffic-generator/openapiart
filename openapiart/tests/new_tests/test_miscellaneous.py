import importlib
import pytest
import os
from openapiart.openapiart import OpenApiArt as openapiart_class

module = importlib.import_module("sanity")

# Add all miscellaneous tests in this test script


def test_set(api):
    config = api.test_config()

    r_val = config.native_features.required_val
    r_val.set(int_val=40, num_val=4.56, str_val="str", bool_val=True)

    s_obj = config.serialize()
    config.deserialize(s_obj)
    r_val = config.native_features.required_val

    assert r_val.int_val == 40
    assert r_val.num_val == 4.56
    assert r_val.str_val == "str"
    assert r_val.bool_val

    config_data = config.native_features.required_val._properties
    assert config_data["int_val"], "int_val value is not set correctly"
    assert config_data["num_val"], "num_val value is not set correctly"
    assert config_data["str_val"], "str_val value is not set correctly"
    assert config_data["bool_val"], "bool_val value is not set correctly"

    r_val_array = config.native_features.required_val_array
    r_val_array.set(
        int_vals=[40, 50],
        num_vals=[4.35, 1.23],
        str_vals=["str1", "str2"],
        bool_vals=[True, True],
    )

    s_obj = config.serialize()
    config.deserialize(s_obj)
    r_val_array = config.native_features.required_val_array

    assert r_val_array.int_vals == [40, 50]
    assert r_val_array.num_vals == [4.35, 1.23]
    assert r_val_array.str_vals == ["str1", "str2"]
    assert r_val_array.bool_vals == [True, True]

    config_data = config.native_features.required_val_array._properties
    assert config_data["int_vals"], "int_vals value is not set correctly"
    assert config_data["num_vals"], "num_vals value is not set correctly"
    assert config_data["str_vals"], "str_vals value is not set correctly"
    assert config_data["bool_vals"], "bool_vals value is not set correctly"

    o_val = config.native_features.optional_val
    o_val.set(
        int_val=150, num_val=510.05, str_val="new_str_val", bool_val=True
    )

    s_obj = config.serialize()
    config.deserialize(s_obj)
    o_val = config.native_features.optional_val

    assert o_val.int_val == 150
    assert o_val.num_val == 510.05
    assert o_val.str_val == "new_str_val"
    assert o_val.bool_val

    config_data = config.native_features.optional_val._properties
    assert config_data["int_val"], "int_val value is not set correctly"
    assert config_data["num_val"], "num_val value is not set correctly"
    assert config_data["str_val"], "str_val value is not set correctly"
    assert config_data["bool_val"], "bool_val value is not set correctly"

    o_val_array = config.native_features.optional_val_array
    o_val_array.set(
        int_vals=[20, 10],
        num_vals=[210.01, 120.02],
        str_vals=["first_new_str", "second_new_str"],
        bool_vals=[True, False],
    )

    s_obj = config.serialize()
    config.deserialize(s_obj)
    o_val_array = config.native_features.optional_val_array

    assert o_val_array.int_vals == [20, 10]
    assert o_val_array.num_vals == [210.01, 120.02]
    assert o_val_array.str_vals == ["first_new_str", "second_new_str"]
    assert o_val_array.bool_vals == [True, False]

    config_data = config.native_features.optional_val_array._properties
    assert config_data["int_vals"], "int_vals value is not set correctly"
    assert config_data["num_vals"], "num_vals value is not set correctly"
    assert config_data["str_vals"], "str_vals value is not set correctly"
    assert config_data["bool_vals"], "bool_vals value is not set correctly"

    b_val = config.native_features.boundary_val

    b_val.set(int_val=100, num_val=100.0)

    s_obj = config.serialize()
    config.deserialize(s_obj)
    b_val = config.native_features.boundary_val

    assert b_val.int_val == 100
    assert b_val.num_val == 100.0

    config_data = config.native_features.boundary_val._properties
    assert config_data["int_val"], "int_val value is not set correctly"
    assert config_data["num_val"], "num_val value is not set correctly"

    b_val_array = config.native_features.boundary_val_array

    b_val_array.set(int_vals=[50, 51], num_vals=[50.05, 78.90])

    s_obj = config.serialize()
    config.deserialize(s_obj)
    b_val_array = config.native_features.boundary_val_array

    assert b_val_array.int_vals == [50, 51]
    assert b_val_array.num_vals == [50.05, 78.90]

    config_data = config.native_features.boundary_val_array._properties
    assert config_data["int_vals"], "int_vals value is not set correctly"
    assert config_data["num_vals"], "num_vals value is not set correctly"

    try:
        config.serialize()
    except Exception as e:
        pytest.fail(str(e))


def create_openapi_artifacts(openapiart_class, sdk=None, file_name=None):
    open_api = openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "../api/info.yaml"),
            os.path.join(os.path.dirname(__file__), file_name),
        ],
        artifact_dir=os.path.join(
            os.path.dirname(__file__), "..", "..", "fielduid"
        ),
        extension_prefix="field",
        proto_service="fieldapi",
    )
    if sdk == "python" or sdk is None:
        open_api.GeneratePythonSdk(package_name="fieldapi")


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_validate_property_name_camel_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="../field_uid/property_name_camel_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


def test_validate_property_name_pascal_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="../field_uid/property_name_pascal_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


def test_validate_property_name_upper_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="../field_uid/property_name_upper_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
