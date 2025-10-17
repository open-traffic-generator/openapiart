import os
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(
    openapiart_class, sdk=None, file_name=None, description_check=None
):
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
        strict_description_validation=description_check,
    )
    if sdk == "python" or sdk is None:
        open_api.GeneratePythonSdk(package_name="fieldapi")


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_validate_descriptions_in_properties():
    error_msgs = [
        "Property Field.Config:prop1 is missing description field",
        "Property Field.Config2:prop_b is missing description field",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/without_description.yaml",
            description_check="properties",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        str_compare(msg, error_value)


def test_validate_descriptions_in_objects():
    error_msgs = [
        "Schema object Field.Config is missing description field",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/without_description.yaml",
            description_check="objects",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        str_compare(msg, error_value)


def test_validate_descriptions_in_all_nodes():
    error_msgs = [
        "Schema object Field.Config is missing description field",
        "Property Field.Config:prop1 is missing description field",
        "Property Field.Config2:prop_b is missing description field",
    ]
    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class,
            file_name="./field_uid/without_description.yaml",
            description_check="all",
        )
    error_value = execinfo.value.args[0]
    for msg in error_msgs:
        str_compare(msg, error_value)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
