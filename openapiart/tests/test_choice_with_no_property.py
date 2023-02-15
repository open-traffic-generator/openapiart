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

    # encode and decode should have no problem
    s_f_obj = f_obj.serialize()
    f_obj.deserialize(s_f_obj)
    assert f_obj._properties.get("choice", None) == "f_c"
    len(f_obj._properties) == 1


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
