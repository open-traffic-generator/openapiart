import pytest


def test_normal_working_of_enums(api):
    config = api.prefix_config()
    config.required_object.e_a = 10
    config.required_object.e_b = 20
    config.a = "abc"
    config.b = 10.2
    config.c = 30
    config.response = "status_400"

    # should not cause any problem
    s_obj = config.serialize()
    config.deserialize(s_obj)


def test_working_of_list_of_enums(api):
    config = api.prefix_config()
    config.required_object.e_a = 10
    config.required_object.e_b = 20
    config.a = "abc"
    config.b = 10.2
    config.c = 30

    # should accept a subset of enums
    config.d_values = ["a"]
    s_obj = config.serialize()
    config.deserialize(s_obj)
    assert config.d_values == ["a"]

    config.d_values = ["a", "b"]
    s_obj = config.serialize()
    config.deserialize(s_obj)
    assert config.d_values == ["a", "b"]

    config.d_values = ["a", "b", "c"]
    s_obj = config.serialize()
    config.deserialize(s_obj)
    assert config.d_values == ["a", "b", "c"]

    # should allow duplicate enum values
    config.d_values = ["a", "a"]
    s_obj = config.serialize()
    config.deserialize(s_obj)
    assert config.d_values == ["a", "a"]

    config.d_values = ["a", "a", "b", "b", "c", "c"]
    s_obj = config.serialize()
    config.deserialize(s_obj)
    assert config.d_values == ["a", "a", "b", "b", "c", "c"]


def test_error_for_enums(api):
    config = api.prefix_config()
    config.required_object.e_a = 10
    config.required_object.e_b = 20
    config.a = "abc"
    config.b = 10.2
    config.c = 30
    config.response = "status_123"

    # wrong enum value for simple enum should fail
    expected_exception = "property response shall be one of these ['status_200', 'status_400', 'status_404', 'status_500'] enum, but got status_123 at <class 'sanity.sanity.PrefixConfig'>"
    with pytest.raises(Exception) as execinfo:
        config.serialize()
    assert execinfo.value.args[0] == expected_exception

    # wrong enum for list of enums should fail
    config.response = "status_400"
    config.d_values = ["a", "error"]
    expected_exception = "property d_values shall be one of these ['a', 'b', 'c'] enum, but got ['a', 'error'] at <class 'sanity.sanity.PrefixConfig'>"
    with pytest.raises(Exception) as execinfo:
        config.serialize()
    assert execinfo.value.args[0] == expected_exception


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
