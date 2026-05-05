import os
import yaml
import tempfile
import pytest
from openapiart.openapiart import OpenApiArt as openapiart_class


def create_openapi_artifacts(artifact_dir):
    openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
            os.path.join(
                os.path.dirname(__file__), "./security/security.yaml"
            ),
        ],
        artifact_dir=artifact_dir,
        extension_prefix="status",
        proto_service="statusapi",
    )


def test_security_definitions_preserved():
    with tempfile.TemporaryDirectory() as artifact_dir:
        create_openapi_artifacts(artifact_dir)

        bundled_path = os.path.join(artifact_dir, "openapi.yaml")
        with open(bundled_path) as f:
            bundled = yaml.safe_load(f)

        _assert_root_security(bundled)
        _assert_security_schemes(bundled)
        _assert_operation_level_security(bundled)


def _assert_root_security(bundled):
    assert (
        "security" in bundled
    ), "root-level 'security' key missing from bundled output"
    security = bundled["security"]
    assert isinstance(security, list)
    assert {"bearerAuth": []} in security
    assert {"apiKeyAuth": []} in security


def _assert_security_schemes(bundled):
    assert "securitySchemes" in bundled.get(
        "components", {}
    ), "'securitySchemes' missing from bundled components"
    schemes = bundled["components"]["securitySchemes"]

    # http bearer
    assert "bearerAuth" in schemes
    bearer = schemes["bearerAuth"]
    assert bearer["type"] == "http"
    assert bearer["scheme"] == "bearer"
    assert bearer["bearerFormat"] == "JWT"

    # apiKey
    assert "apiKeyAuth" in schemes
    apikey = schemes["apiKeyAuth"]
    assert apikey["type"] == "apiKey"
    assert apikey["in"] == "header"
    assert apikey["name"] == "X-API-Key"

    # oauth2 with authorizationCode flow
    assert "oauth2Auth" in schemes
    oauth2 = schemes["oauth2Auth"]
    assert oauth2["type"] == "oauth2"
    flow = oauth2["flows"]["authorizationCode"]
    assert flow["authorizationUrl"] == "https://example.com/oauth/authorize"
    assert flow["tokenUrl"] == "https://example.com/oauth/token"
    assert "read:config" in flow["scopes"]
    assert "write:config" in flow["scopes"]

    # openIdConnect
    assert "openIdAuth" in schemes
    oidc = schemes["openIdAuth"]
    assert oidc["type"] == "openIdConnect"
    assert (
        oidc["openIdConnectUrl"]
        == "https://example.com/.well-known/openid-configuration"
    )


def _assert_operation_level_security(bundled):
    paths = bundled.get("paths", {})

    # /status overrides root security with oauth2 + scope
    status_security = paths["/status"]["get"]["security"]
    assert isinstance(status_security, list)
    assert len(status_security) == 1
    assert "oauth2Auth" in status_security[0]
    assert "read:config" in status_security[0]["oauth2Auth"]

    # /public explicitly disables security with an empty list
    public_security = paths["/public"]["get"]["security"]
    assert isinstance(public_security, list)
    assert len(public_security) == 0
