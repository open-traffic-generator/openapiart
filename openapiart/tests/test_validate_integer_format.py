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
            os.path.dirname(__file__), "..", "..", "check_type"
        ),
        extension_prefix="status",
        proto_service="statusapi",
    )


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_integer_format():
    error_msgs = [
        "components.schemas.Config.properties.u8 has type integer of unsporrted format uint8, supported formats are ['int64', 'int32', 'uint32', 'uint64']",
        "components.schemas.Config.properties.int has type integer of unsporrted format int, supported formats are ['int64', 'int32', 'uint32', 'uint64']",
        "components.schemas.Config.properties.u16 has type integer of unsporrted format uint16, supported formats are ['int64', 'int32', 'uint32', 'uint64']",
        "components.schemas.Config.properties.typo has type integer of unsporrted format usnit64, supported formats are ['int64', 'int32', 'uint32', 'uint64']",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./check_type/invalid_integer_type.yaml",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        str_compare(msg, error_value)
