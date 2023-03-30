import os
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(openapiart_class, sdk=None, file_name=None):
    openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
            os.path.join(os.path.dirname(__file__), file_name),
        ],
        artifact_dir=os.path.join(
            os.path.dirname(__file__), "..", "..", "response"
        ),
        extension_prefix="status",
        proto_service="statusapi",
    )


def str_compare(validte_str, entire_str, item):
    return validte_str in entire_str and item in entire_str


def test_validate_response_default():
    default_error = "paths./config.post.responses: is missing the following required responses:"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./response/response_default_error.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(default_error, error_value, "default")


def test_validate_response_200():
    default_error = "paths./config.post.responses: is missing the following required responses"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./response/response_200_error.yaml",
        )
    error_value = execinfo.value.args[0]
    assert str_compare(default_error, error_value, "200")


def test_required_fields_in_error():
    default_error = (
        "Error schema must have ['code', 'errors'] as required properties"
    )
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./response/response_required_error.yaml",
        )
    error_value = execinfo.value.args[0]
    assert default_error == error_value


def test_error_for_missing_required():
    error_msg = "Error schema in components.schemas.Error must have the required field in it"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./response/response_missing_required_in_error.yaml",
        )
    error_value = execinfo.value.args[0]
    assert error_msg == error_value


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
