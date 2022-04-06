def test_deprecate(api):
    conf = api.prefix_config()
    conf.required_object.e_a = 10
    conf.required_object.e_b = 20
    conf.a = "abc"
    conf.b = 10.2
    conf.c = 30

    msg = (
        "Deprecated message and use instead information and "
        "deprecated from version and lib names can be provided"
    )

    def strip_data(message):
        return " ".join([ln.strip() for ln in message.split("\n")]).strip()

    assert api.get_api_warnings() != []
    assert len(api.get_api_warnings()) == 1
    assert strip_data(api.get_api_warnings()[0]) == msg

    api.clear_api_warnings()
    assert api.get_api_warnings() == []

    data = conf.serialize(conf.DICT)
    assert api.get_api_warnings() != []
    assert len(api.get_api_warnings()) == 1
    assert strip_data(api.get_api_warnings()[0]) == msg

    api.clear_api_warnings()
    assert api.get_api_warnings() == []

    conf1 = api.prefix_config()
    conf1.deserialize(data)
    assert api.get_api_warnings() != []
    assert len(api.get_api_warnings()) == 1
    assert strip_data(api.get_api_warnings()[0]) == msg

    api.clear_api_warnings()
    assert api.get_api_warnings() == []
