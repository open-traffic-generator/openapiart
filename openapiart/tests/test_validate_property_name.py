import os
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(openapiart_class, sdk=None, file_name=None):
    open_api = openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
            os.path.join(os.path.dirname(__file__), file_name),
        ],
        artifact_dir=os.path.join(
            os.path.dirname(__file__), "..", "..", "fielduid"
        ),
        extension_prefix="field",
        proto_service="fieldapi",
    )
    if sdk == "python" or sdk is None:
        open_api.GeneratePythonSdk(package_name="fieldapi")


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_validate_property_name_camel_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/property_name_camel_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


def test_validate_property_name_pascal_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/property_name_pascal_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


def test_validate_property_name_upper_case():
    include_error = "is invalid. Only lower case letters separated with an underscore is allowed"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/property_name_upper_case.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
