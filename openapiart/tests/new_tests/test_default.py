import pytest


def test_option_val_schema(api):

    #  This test checks the values in optional schema.
    #  Objective is to check if default values are set properly.

    config = api.test_config()
    o_val = config.native_features.optional_val

    assert o_val.int_val == 50
    assert o_val.num_val == 50.05
    assert o_val.str_val == "default_str_val"
    assert o_val.bool_val

    s_obj = config.serialize()
    config.deserialize(s_obj)
    o_val = config.native_features.optional_val

    assert o_val.int_val == 50
    assert o_val.num_val == 50.05
    assert o_val.str_val == "default_str_val"
    assert o_val.bool_val


def test_option_array_val_schema(api):

    #  This test checks the values in optional array schema.
    #  Objective is to check if default values are set properly.

    config = api.test_config()
    o_val = config.native_features.optional_val_array

    assert o_val.int_vals == [10, 20]
    assert o_val.num_vals == [10.01, 20.02]
    assert o_val.str_vals == ["first_str", "second_str"]
    assert o_val.bool_vals is None

    s_obj = config.serialize()
    config.deserialize(s_obj)
    o_val = config.native_features.optional_val_array

    assert o_val.int_vals == [10, 20]
    assert o_val.num_vals == [10.01, 20.02]
    assert o_val.str_vals == ["first_str", "second_str"]
    assert o_val.bool_vals is None


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
