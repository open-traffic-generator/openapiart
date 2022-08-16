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


def test_validate_xinclude():
    include_error = "x-include shall be a path of any property or response"
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class, file_name="./field_uid/xinclude.yaml"
        )
    error_value = execinfo.value.args[0]
    assert str_compare(include_error, error_value)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
