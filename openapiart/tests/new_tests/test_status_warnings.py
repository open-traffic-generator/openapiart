import pytest


def test_status_warning_for_primitive_attr(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.decprecated_property_2 = 34
    status_obj.under_review_property_2 = 56
    s_obj = config.serialize()

    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: decprecated_property_2 property in schema XStatusObject is deprecated, test deprecated"
        in err
    )
    assert (
        "[WARNING]: under_review_property_2 property in schema XStatusObject is under_review, test under_review"
        in err
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "[WARNING]: decprecated_property_2 property in schema XStatusObject is deprecated, test deprecated"
        in err
    )
    assert (
        "[WARNING]: under_review_property_2 property in schema XStatusObject is under_review, test under_review"
        in err
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
        in err
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "x_status_object property in schema ExtendedFeatures is under_review, test under_review"
        in err
    )


def test_warnings_for_x_enmu_attr(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.enum_property = "decprecated_property_1"

    s_obj = config.serialize(config.DICT)
    out, err = capsys.readouterr()

    assert out == ""
    assert (
        "DECPRECATED_PROPERTY_1 enum in property enum_property is deprecated, test deprecated"
        in err
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "DECPRECATED_PROPERTY_1 enum in property enum_property is deprecated, test deprecated"
        in err
    )


def test_warnings_for_schema(api, capsys):
    config = api.test_config()
    status_obj = config.extended_features.x_status_object
    status_obj.enum_property = "decprecated_property_1"

    s_obj = config.serialize(config.DICT)
    out, err = capsys.readouterr()

    assert out == ""
    assert (
        "TestConfig is under_review, the whole schema is being reviewed" in err
    )

    config.deserialize(s_obj)
    out, err = capsys.readouterr()
    assert out == ""
    assert (
        "TestConfig is under_review, the whole schema is being reviewed" in err
    )


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
