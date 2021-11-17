"""Using the model files in the api directory create the following artifacts:
- ./art/openapi.yaml (bundled/validated model file)
- ./art/openapi.json (json version of openapi.yaml)
- ./art/openapi.html (html doc of openapi.yaml if node and redoc-cli is installed)
- ./art/sanity/__init__.py (python ux package)
- ./art/sanity/sanity.py (python ux package)
- ./art/sanity/sanity_pb2.py (python protobuf stubs)
- ./art/sanity/sanity_pb2_grpc.py (python grpc stubs)

- ./art/go/<protobuf_package_name>.proto (proto version of openapi.yaml)
- ./art/go/<go_module_name>/<protobuf_package_name>/sanity_pb2.go (go protobuf stubs)
- ./art/go/<go_module_name>/<protobuf_package_name>/sanity_pb2_grpc.go (go grpc stubs)
- ./art/go/<go_module_name>/go.mod (go ux package)
- ./art/go/<go_module_name>/sanity.go (go ux package)
"""
import sys
import os
import importlib


def create_openapi_artifacts(openapiart_class):
    openapiart_class(
        api_files=[
        os.path.join(os.path.dirname(__file__), "./openapiart/tests/api/info.yaml"),
        os.path.join(os.path.dirname(__file__), "./openapiart/tests/common/common.yaml"),
        os.path.join(os.path.dirname(__file__), "./openapiart/tests/api/api.yaml"),
    ],
        artifact_dir=os.path.join(os.path.dirname(__file__), "art"),
        extension_prefix="sanity",
    ).GeneratePythonSdk(
        package_name="sanity"
    ).GenerateGoSdk(
        package_dir="github.com/open-traffic-generator/openapiart/pkg", package_name="openapiart"
    )


if __name__ == "__main__":
    sys.path.append(os.path.normpath(os.path.join(os.path.dirname(__file__), "..")))
    module = importlib.import_module("openapiart.openapiart")
    openapiart_class = getattr(module, "OpenApiArt")
    create_openapi_artifacts(openapiart_class)
