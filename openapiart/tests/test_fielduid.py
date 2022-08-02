import os
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(openapiart_class, sdk=None, file_name=None):
    openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
            os.path.join(os.path.dirname(__file__), file_name),
        ],
        artifact_dir=os.path.join(os.path.dirname(__file__), "fielduid"),
        extension_prefix="field",
        proto_service="fieldapi",
    )


def str_compare(validte_str, entire_str):
    return validte_str in entire_str


def test_validate_field_uid():
    dup_error = "Field.Config contain duplicate [1] x-field-uid"
    reserved_error = "x-field-uid 2 of Field.Config:usereserved should not conflict with x-reserved-field-uids"
    missing_error = "x-field-uid is missing in Field.Config:missinguid"
    min_range_error = (
        "x-field-uid -1 of Field.Config:minrange not in range (1 to 2^29)"
    )
    max_range_error = "x-field-uid 536870912 of Field.Config:maxrange not in range (1 to 2^29)"
    dup_enum_error = "Field.Config contain duplicate [1] x-field-uid. x-field-uid should be unique."
    reserved_enum_error = "x-field-uid 4 within enum fieldenum:conflictenum conflict with x-reserved-field-uids"
    missing_enum_error = "x-field-uid is missing in missingenum"
    min_enum_range_error = (
        "x-field-uid -3 of fieldenum:minenum not in range (1 to 2^29)"
    )
    max_enum_range_error = (
        "x-field-uid 536870912 of fieldenum:maxenum not in range (1 to 2^29)"
    )
    missing_rsp_error = "x-field-uid is missing in /config:post:500 response"
    dup_rsp_error = "/config:post contain duplicate [2] x-field-uid. x-field-uid should be unique."
    min_rsp_error = (
        "x-field-uid -4 of /config:post:501 not in range (1 to 2^29)"
    )
    max_rsp_error = (
        "x-field-uid 536870998 of /config:post:502 not in range (1 to 2^29)"
    )
    reserved_rsp_error = "x-field-uid 5 of /config:post:503 should not conflict with x-reserved-field-uids"

    with pytest.raises(Exception) as execinfo:
        create_openapi_artifacts(
            openapiart_class, file_name="./field_uid/fielduid.yaml"
        )
    error_value = execinfo.value.args[0]
    assert str_compare(dup_error, error_value)
    assert str_compare(reserved_error, error_value)
    assert str_compare(missing_error, error_value)
    assert str_compare(dup_enum_error, error_value)
    assert str_compare(reserved_enum_error, error_value)
    assert str_compare(missing_enum_error, error_value)
    assert str_compare(min_range_error, error_value)
    assert str_compare(max_range_error, error_value)
    assert str_compare(min_enum_range_error, error_value)
    assert str_compare(max_enum_range_error, error_value)
    assert str_compare(missing_rsp_error, error_value)
    assert str_compare(dup_rsp_error, error_value)
    assert str_compare(min_rsp_error, error_value)
    assert str_compare(max_rsp_error, error_value)
    assert str_compare(reserved_rsp_error, error_value)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
