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
        "invalid x-stream value abc present in set_config valid values are server and client",
        "x-stream value def present in get_config valid values are server and client",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./api/invalid_api.yaml",
        )
    error_value = execinfo.value.args[0]

    for msg in error_msgs:
        str_compare(msg, error_value)
