import os
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(openapiart_class, file_name):
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


def test_invalid_type():
    status_error = "invalid x-unique type boolean in components.schemas.Config.properties.invalid_type.x-unique, x-unique is only allowed on string values"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            "./x_unique/invalid_type_x_unique.yaml",
        )
    error_value = execinfo.value.args[0]
    assert error_value == status_error


def test_invalid_value():
    status_error = "invalid value scoped for x-unique in components.schemas.Config.properties.invalid_value.x-unique, x-unique can only have one of the following ['global', 'local']"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            "./x_unique/invalid_value_x_unique.yaml",
        )
    error_value = execinfo.value.args[0]
    assert error_value == status_error


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
