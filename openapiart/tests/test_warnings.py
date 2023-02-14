import pytest


@pytest.mark.skip(reason="shall be restored")
def test_warnings(api):
    # api.clear_api_warnings()
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    msg = (
        "Property b is being deprecated from the sdk version x.x.x "
        "and property x shall be used instead"
    )

    def strip_data(message):
        return " ".join([ln.strip() for ln in message.split("\n")]).strip()

    assert conf.__warnings__ != []
    assert len(conf.__warnings__) == 2
    assert strip_data(conf.__warnings__[0]) == "property under review"
    assert strip_data(conf.__warnings__[1]) == msg

    conf.warnings()

    data = conf.serialize(conf.DICT)
    assert conf.__warnings__ != []
    assert len(conf.__warnings__) == 2
    assert strip_data(conf.__warnings__[0]) == "property under review"
    assert strip_data(conf.__warnings__[1]) == msg

    conf.warnings()
    assert conf.__warnings__ == []

    conf1 = api.prefix_config()
    conf1.deserialize(data)
    assert conf1.__warnings__ != []
    assert len(conf1.__warnings__) == 2
    assert strip_data(conf1.__warnings__[0]) == "property under review"
    assert strip_data(conf1.__warnings__[1]) == msg

    conf1.warnings()
    assert conf1.__warnings__ == []

    u_conf = api.update_config()
    u_conf.g.gobject(g_a="abcd")

    api.update_configuration(u_conf)
    assert api.__warnings__ != []
    assert len(api.__warnings__) == 1
