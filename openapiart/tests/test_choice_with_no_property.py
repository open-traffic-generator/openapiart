import pytest


def test_choice_with_leaf_nodes(api):
    config = api.prefix_config()
    f_obj = config.f

    # check default
    assert f_obj.choice == "f_a"
    assert f_obj._properties.get("f_a", None) is not None

    # setting other choices work as before
    f_obj.f_b = 3.45
    assert f_obj.choice == "f_b"
    assert f_obj._properties.get("f_b", None) == 3.45
    f_obj.f_a = "str2"
    assert f_obj.choice == "f_a"
    assert f_obj._properties.get("f_a", None) == "str2"

    # setting choices with no properties
    f_obj.choice = f_obj.F_C
    assert f_obj._properties.get("choice", None) == "f_c"
    len(f_obj._properties) == 1

    # serialize and deserialize should have no problem
    s_f_obj = f_obj.serialize()
    f_obj.deserialize(s_f_obj)
    assert f_obj._properties.get("choice", None) == "f_c"
    len(f_obj._properties) == 1


def test_choice_with_iter_objects(api):
    config = api.prefix_config()

    # default choice with no properties should be set properly
    c_obj = config.choice_object.add()
    assert c_obj.choice == "no_obj"
    assert len(c_obj._properties) == 1

    # acesing of objects with choice set to choice with no property should work
    c_obj = config.choice_object[0]

    # setting of other properties should not have a problem
    c_obj.e_obj.e_a = 1.1
    assert c_obj.choice == "e_obj"
    assert c_obj._properties.get("e_obj") is not None

    c_obj.f_obj.f_b = 1.1
    assert c_obj.choice == "f_obj"
    assert c_obj._properties.get("f_obj") is not None

    c_obj.choice = "no_obj"
    assert c_obj.choice == "no_obj"
    assert len(c_obj._properties) == 1

    # serialize and deserialize should have no problem
    s_c_obj = c_obj.serialize()
    c_obj.deserialize(s_c_obj)
    assert c_obj._properties.get("choice") == "no_obj"
    len(c_obj._properties) == 1


def test_choice_in_choice_heirarchy(api):
    config = api.prefix_config()

    # default choice with no properties should be set properly
    c_obj = config.choice_object.add()
    assert c_obj.choice == "no_obj"
    assert len(c_obj._properties) == 1

    # acesing of objects with choice set to choice with no property should work
    c_obj = config.choice_object[0]

    f_obj = c_obj.f_obj

    # check default in child
    assert f_obj.choice == "f_a"
    assert f_obj._properties.get("f_a", None) is not None

    # setting choice with no properties in child as well
    f_obj.choice = "f_c"
    assert f_obj._properties.get("choice", None) == "f_c"
    len(f_obj._properties) == 1

    # serialize and deserialize should have no problem
    s_c_obj = c_obj.serialize()
    c_obj.deserialize(s_c_obj)
    assert c_obj.choice == "f_obj"
    assert c_obj._properties.get("f_obj") is not None
    assert c_obj.f_obj.choice == "f_c"


def test_choice_with_invalid_enum_and_none_value(api):
    config = api.prefix_config()
    f_obj = config.f

    # check default
    assert f_obj.choice == "f_a"
    assert f_obj._properties.get("f_a", None) is not None

    # setting it to a valid value
    f_obj.choice = "f_b"
    assert f_obj.choice == "f_b"
    assert f_obj._properties.get("f_a", None) is None

    # setting None should set to default value
    f_obj.choice = None
    assert f_obj.choice == "f_a"
    assert f_obj._properties.get("f_b", None) is None

    # setting invalid value should result in exception
    choice_error = (
        "random is not a valid choice, valid choices are f_a, f_b, f_c"
    )
    with pytest.raises(Exception) as execinfo:
        f_obj.choice = "random"
    assert execinfo.value.args[0] == choice_error


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
