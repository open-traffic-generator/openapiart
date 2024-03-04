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


def test_validate_pattern():
    error_msgs = [
        "components.schemas.Config.properties.integer.x-field-pattern property using x-field-pattern with format integer must contain length property",
        "components.schemas.Config.properties.wrong.x-field-pattern has unspported format random , valid formats are ['integer', 'ipv4', 'ipv6', 'mac', 'checksum', 'oid']",
        "components.schemas.Config.properties.int_128.x-field-pattern property using x-field-pattern with format integer cannot have length greater than 64",
        "signed property can only be used if the format is set to integer in property components.schemas.Config.properties.signed_value_without_int.x-field-pattern",
        "invalid value 45 in components.schemas.Config.properties.wrong_int_signed_value.x-field-pattern, signed property can either be true or false",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./pattern/invalid-pattern.yaml",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        assert str_compare(msg, error_value)
