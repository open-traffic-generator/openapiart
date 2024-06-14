import pytest


def test_required_val_schema(api):

    # This test checks the values in required schema.
    # Objective is that serialize and deserialize should not have a problem.

    config = api.test_config()
    r_val = config.native_features.required_val
    r_val.int_val = 40
    r_val.num_val = 4.56
    r_val.str_val = "str"
    r_val.bool_val = True

    s_obj = config.serialize()
    config.deserialize(s_obj)
    r_val = config.native_features.required_val

    assert r_val.int_val == 40
    assert r_val.num_val == 4.56
    assert r_val.str_val == "str"
    assert r_val.bool_val


def test_reuired_error(api):

    # This test checks error raised by SDK for required properties.
    # Objective is that the SDK should raise proper exception.

    config = api.test_config()
    err_msg1 = "int_val is a mandatory property of <class 'pyapi.pyapi.RequiredVal'> and should not be set to None"
    with pytest.raises(Exception) as execinfo:
        r_val = config.native_features.required_val
        config.serialize()
    assert execinfo.value.args[0] == err_msg1

    r_val.int_val = 40
    r_val.str_val = "str"
    r_val.bool_val = False

    err_msg2 = "num_val is a mandatory property of <class 'pyapi.pyapi.RequiredVal'> and should not be set to None"
    with pytest.raises(Exception) as execinfo:
        config.serialize()

    assert execinfo.value.args[0] == err_msg2


def test_required_array_val_schema(api):

    # This test checks the values in required array schema.
    # Objective is that serialize and deserialize should have not a problem.

    config = api.test_config()
    r_val = config.native_features.required_val_array
    r_val.int_vals = [40, 50]
    r_val.num_vals = [4.35, 1.23]
    r_val.str_vals = ["str1", "str2"]
    r_val.bool_vals = [True, True]

    s_obj = config.serialize()
    config.deserialize(s_obj)
    r_val = config.native_features.required_val_array

    assert r_val.int_vals == [40, 50]
    assert r_val.num_vals == [4.35, 1.23]
    assert r_val.str_vals == ["str1", "str2"]
    assert r_val.bool_vals == [True, True]


def test_required_array_val_errors(api):

    # This test checks error raised by SDK for required array properties.
    # Objective is that the SDK should raise proper exception.

    config = api.test_config()
    err_msg1 = "int_vals is a mandatory property of <class 'pyapi.pyapi.RequiredValArray'> and should not be set to None"
    with pytest.raises(Exception) as execinfo:
        r_val = config.native_features.required_val_array
        config.serialize()
    assert execinfo.value.args[0] == err_msg1

    r_val.int_vals = [40, 50]
    r_val.num_vals = [4.35, 1.23]
    r_val.str_vals = ["str1", "str2"]

    err_msg2 = "bool_vals is a mandatory property of <class 'pyapi.pyapi.RequiredValArray'> and should not be set to None"
    with pytest.raises(Exception) as execinfo:
        config.serialize()

    assert execinfo.value.args[0] == err_msg2


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
