import copy
import os

import jsonpath_ng
import pytest
import yaml

from openapiart.bundler import Bundler
from openapiart.description import render_description


def _bundler_with_content(content):
    bundler = Bundler(api_files=[])
    bundler._content = content
    return bundler


def test_resolve_x_status_does_not_alter_description():
    schema = {
        "description": "Source description",
        "x-status": {"status": "under-review", "information": "Review note"},
    }
    bundler = _bundler_with_content({"components": {"schemas": {"A": schema}}})

    bundler._resolve_x_status()

    assert schema["description"] == "Source description"
    assert schema["x-status"]["status"] == "under_review"
    assert schema["x-status"]["information"] == "Review note"


def test_resolve_x_status_fills_missing_information(capsys):
    schema = {
        "description": "Source description",
        "x-status": {"status": "deprecated"},
    }
    bundler = _bundler_with_content({"components": {"schemas": {"A": schema}}})

    bundler._resolve_x_status()

    assert schema["description"] == "Source description"
    assert schema["x-status"]["information"] == "Information TBD"
    # jsonpath_ng's Fields.__str__/full_path formatting differs across
    # versions (e.g. "a.b.c" vs "((a.b).c)"), so match on the stable
    # pieces rather than the exact rendered path.
    warning = capsys.readouterr().out
    assert "[WARNING]:" in warning
    assert "information missing" in warning
    assert "components" in warning
    assert "schemas" in warning
    assert "x-status" in warning


def test_resolve_x_status_rejects_invalid_status():
    schema = {
        "description": "Source description",
        "x-status": {"status": "test"},
    }
    bundler = _bundler_with_content({"components": {"schemas": {"A": schema}}})

    with pytest.raises(Exception) as execinfo:
        bundler._resolve_x_status()

    assert (
        execinfo.value.args[0]
        == "Invalid value for x-status.status=test provided; Valid values are ['deprecated', 'under_review']"
    )


def test_render_description_plain_description():
    assert (
        render_description({"description": "Source description"})
        == "Source description"
    )


def test_render_description_status_prefix():
    rendered = render_description(
        {
            "description": "Source description",
            "x-status": {
                "status": "under_review",
                "information": "Review note",
            },
        }
    )

    assert rendered == "Under Review: Review note\n\nSource description"


def test_render_description_constraint_block():
    rendered = render_description(
        {
            "description": "Source description",
            "x-constraint": [
                "/components/schemas/A/properties/a",
                "/components/schemas/B/properties/b",
            ],
        }
    )

    assert (
        rendered
        == "Source description\n\nx-constraint:\n- /components/schemas/A/properties/a\n- /components/schemas/B/properties/b\n"
    )


def test_render_description_does_not_mutate_input():
    openapi_object = {
        "description": "Source description",
        "x-status": {"status": "deprecated", "information": "Old"},
        "x-constraint": ["/components/schemas/A/properties/a"],
    }
    original = copy.deepcopy(openapi_object)

    render_description(openapi_object)

    assert openapi_object == original


def test_bundled_openapi_keeps_authored_descriptions(tmp_path):
    api_dir = os.path.join(os.path.dirname(__file__), "api")
    bundler = Bundler(
        api_files=[
            os.path.join(api_dir, "info.yaml"),
            os.path.join(api_dir, "api.yaml"),
        ],
        output_dir=str(tmp_path),
    )

    bundler.bundle()

    with open(os.path.join(str(tmp_path), "openapi.yaml")) as fp:
        bundled = yaml.safe_load(fp.read())

    assert (
        bundled["paths"]["/config"]["patch"]["description"]
        == "Sets configuration resources."
    )
    assert (
        bundled["paths"]["/config"]["patch"]["x-status"]["status"]
        == "deprecated"
    )
    y_name = bundled["components"]["schemas"]["YObject"]["properties"][
        "y_name"
    ]
    assert "description" not in y_name
    assert y_name["x-constraint"] == [
        "/components/schemas/ZObject/properties/name",
        "/components/schemas/WObject/properties/w_name",
    ]


def test_auto_feature(openapi_yaml):
    property = jsonpath_ng.parse("$..auto").find(
        openapi_yaml.get("components").get("schemas")
    )
    for auto_field in property:
        if (
            len(auto_field.value) == 1
            and "x-field-uid" in auto_field.value
            or "$ref" in auto_field.value
        ):
            continue
        assert auto_field.value.get("description") is not None
        assert auto_field.value.get("type") is not None
        assert auto_field.value.get("default") is not None
        if auto_field.value.get("type") == "integer":
            assert auto_field.value.get("minimum") is None
            assert auto_field.value.get("maximum") is not None


def test_auto_in_config(config):
    assert config.auto_field_test.choice == "auto"
    assert config.auto_field_test.auto == 0
    assert config.auto_field_test._TYPES.get("auto").get("minimum") is None
    assert config.auto_field_test._TYPES.get("auto").get("maximum") == 255
