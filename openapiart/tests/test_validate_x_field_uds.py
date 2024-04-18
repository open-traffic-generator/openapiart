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
            os.path.dirname(__file__), "..", "..", "pattern"
        ),
        extension_prefix="status",
        proto_service="statusapi",
    )


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_validate_uds():
    error_msgs = [
        "components.schemas.Config.properties.wrong_uds.x-field-uds has unspported format random , valid formats are ['integer', 'ipv4', 'ipv6', 'mac']",
        "components.schemas.Config.properties.wrong_integer.x-field-uds property using x-field-uds with format integer must contain length property",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./pattern/invalid-uds.yaml",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        assert str_compare(msg, error_value)
