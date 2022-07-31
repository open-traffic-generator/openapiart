import jsonpath_ng


def test_auto_feature(openapi_yaml):
    property = jsonpath_ng.parse("$..auto").find(
        openapi_yaml.get("components").get("schemas")
    )
    for auto_field in property:
        if len(auto_field.value) == 1 and "x-field-uid" in auto_field.value:
            continue
        assert auto_field.value.get("description") is not None
        assert auto_field.value.get("type") is not None
        assert auto_field.value.get("default") is not None
        if auto_field.value.get("type") == "integer":
            assert auto_field.value.get("minimum") is not None
            assert auto_field.value.get("maximum") is not None


def test_auto_in_config(config):
    assert config.auto_field_test.choice == "auto"
    assert config.auto_field_test.auto == 0
    assert config.auto_field_test._TYPES.get("auto").get("minimum") == 0
    assert config.auto_field_test._TYPES.get("auto").get("maximum") == 255
