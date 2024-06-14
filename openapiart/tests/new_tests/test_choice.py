import pytest


def test_choice_val_schema(api):

    #  This test checks the values in choice schema.
    #  Objective is to check if choices are set properly or not.

    config = api.test_config()
    c = config.extended_features.choice_val

    c.mixed_val.int_val = 5

    assert c.mixed_val.int_val == 5
    assert c.mixed_val.choice == "int_val"

    c.mixed_val.num_val = 60.70

    assert c.mixed_val.num_val == 60.70
    assert c.mixed_val.choice == "num_val"

    c.mixed_val.str_val = "str1"

    assert c.mixed_val.str_val == "str1"
    assert c.mixed_val.choice == "str_val"

    c.mixed_val.bool_val = False

    assert c.mixed_val.bool_val is False
    assert c.mixed_val.choice == "bool_val"

    s_obj = config.serialize()
    config.deserialize(s_obj)
    c = config.extended_features.choice_val

    assert c.mixed_val.bool_val is False
    assert c.mixed_val.choice == "bool_val"


def test_wrong_choice_error(api):

    # This test checks the behaviour when a wrong choice is passed

    config = api.test_config()
    c = config.extended_features.choice_val

    with pytest.raises(Exception) as execinfo:
        c.mixed_val.choice = "random"

    assert "random is not a valid choice" in execinfo.value.args[0]


def test_choice_heirarchy(api):

    # This test basically checks choice works in object hierarchy

    config = api.test_config()
    c = config.extended_features.choice_val_no_properties
    i_obj = c.intermediate_obj
    lv = i_obj.leaf
    lv.name = "leaf_node"
    lv.value = 3

    assert c.choice == "intermediate_obj"
    assert i_obj.choice == "leaf"
    assert lv.name == "leaf_node"
    assert lv.value == 3

    s_obj = config.serialize()
    config.deserialize(s_obj)
    c = config.extended_features.choice_val_no_properties

    assert c.choice == "intermediate_obj"
    assert c.intermediate_obj.choice == "leaf"
    assert c.intermediate_obj.leaf.name == "leaf_node"
    assert c.intermediate_obj.leaf.value == 3


def test_choice_with_no_properties(api):

    # This test checks setting choice with no properties works properly

    config = api.test_config()
    c = config.extended_features.choice_val_no_properties

    c.choice = "no_obj"

    s_obj = config.serialize()
    config.deserialize(s_obj)
    c = config.extended_features.choice_val_no_properties

    assert c.choice == "no_obj"


def test_choice_with_required_field(api):

    # This test basically choices that are required, errors as well as behaviour

    config = api.test_config()
    c = config.extended_features.choice_val_no_properties

    # should throw error if choice not provided
    with pytest.raises(Exception) as execinfo:
        c.serialize()

    assert (
        execinfo.value.args[0]
        == "choice is a mandatory property of <class 'pyapi.pyapi.ChoiceValWithNoProperties'> and should not be set to None"
    )

    c.intermediate_obj.str_val = "str1"

    assert c.choice == "intermediate_obj"
    assert c.intermediate_obj.choice == "str_val"
    assert c.intermediate_obj.str_val == "str1"

    s_obj = config.serialize()
    config.deserialize(s_obj)
    c = config.extended_features.choice_val_no_properties

    assert c.choice == "intermediate_obj"
    assert c.intermediate_obj.choice == "str_val"
    assert c.intermediate_obj.str_val == "str1"


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
