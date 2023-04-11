import pytest


def test_status_warning_for_primitive_attr(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.decprecated_property_2 = 34
    status_obj.under_review_property_2 = 56
    s_obj = config.serialize()

    out, err = capsys.readouterr()
    print("wohoooo")
    print(out)
    assert err == ""
    assert (
        "[WARNING]: decprecated_property_2 property in schema XStatusObject is deprecated, test deprecated"
        in out
    )
    assert (
        "[WARNING]: under_review_property_2 property in schema XStatusObject is under_review, test under_review"
        in out
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert err == ""
    assert err == ""
    assert (
        "[WARNING]: decprecated_property_2 property in schema XStatusObject is deprecated, test deprecated"
        in out
    )
    assert (
        "[WARNING]: under_review_property_2 property in schema XStatusObject is under_review, test under_review"
        in out
    )


def test_warnings_for_non_primitive_attr(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.decprecated_property_2 = 34

    s_obj = config.serialize(config.DICT)
    out, err = capsys.readouterr()

    assert err == ""
    assert (
        "x_status_object property in schema ExtendedFeatures is under_review, test under_review"
        in out
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert err == ""
    assert (
        "x_status_object property in schema ExtendedFeatures is under_review, test under_review"
        in out
    )


def test_warnings_for_x_enmu_attr(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.enum_property = "decprecated_property_1"

    s_obj = config.serialize(config.DICT)
    out, err = capsys.readouterr()

    assert err == ""
    assert (
        "DECPRECATED_PROPERTY_1 enum in property enum_property is deprecated, test deprecated"
        in out
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert err == ""
    assert (
        "DECPRECATED_PROPERTY_1 enum in property enum_property is deprecated, test deprecated"
        in out
    )


def test_warnings_for_schema(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.enum_property = "decprecated_property_1"

    s_obj = config.serialize(config.DICT)
    out, err = capsys.readouterr()

    assert err == ""
    assert (
        "TestConfig is under_review, the whole schema is being reviewed" in out
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert err == ""
    assert (
        "TestConfig is under_review, the whole schema is being reviewed" in out
    )


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
