import pytest


def test_req_choice_obj(api):
    config = api.prefix_config()
    req_obj = config.required_choice_object
    int_obj = req_obj.intermediate_obj
    int_obj.leaf.name = "some string"


def test_req_choice_validation_errors(api):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.required_choice_object
    parent_exception = "choice is a mandatory property of <class 'sanity.sanity.RequiredChoiceParent'> and should not be set to None"

    with pytest.raises(Exception) as execinfo:
        conf.serialize(conf.DICT)
    assert execinfo.value.args[0] == parent_exception


def test_req_choice_serilize_deserialize(api):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    conf.required_choice_object.intermediate_obj.leaf.name = "some string"

    s_obj = conf.serialize(conf.DICT)

    conf.deserialize(s_obj)
    assert (
        conf.required_choice_object.intermediate_obj.leaf.name == "some string"
    )


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
