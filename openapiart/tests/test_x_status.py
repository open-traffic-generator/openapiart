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
            os.path.dirname(__file__), "..", "..", "x-status"
        ),
        extension_prefix="status",
        proto_service="statusapi",
    )


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_x_status_with_invalid_status():
    status_error = "Invalid value for x-status.status=test provided; Valid values are ['deprecated', 'under_review']"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./x-status/invalid_x_status.yaml",
        )
    error_value = execinfo.value.args[0]
    assert error_value == status_error


def test_x_status_with_required_object():
    status_error = "Property status within schema Config have both required as well as deprecated status"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./x-status/x_status_with_required.yaml",
        )
    error_value = execinfo.value.args[0]
    assert error_value == status_error


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
