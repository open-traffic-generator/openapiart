import jsonpath_ng


def test_auto_feature(openapi_yaml):
    property = parser = jsonpath_ng.parse("$..auto").find(
        openapi_yaml.get("components").get("schemas")
    )
    for auto_field in property:
        assert auto_field.value.get("description") is not None
        assert auto_field.value.get("type") is not None
        assert auto_field.value.get("default") is not None
        if auto_field.value.get("type") == "integer":
            assert auto_field.value.get("minimum") is not None
            assert auto_field.value.get("maximum") is not None

