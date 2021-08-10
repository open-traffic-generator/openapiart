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
import pytest
import sys
import os
import importlib
import logging


def create_openapi_artifacts():
    sys.path.append(os.path.normpath(os.path.join(os.path.dirname(__file__), "..")))
    api_files = [
        os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
        os.path.join(os.path.dirname(__file__), "./common/common.yaml"),
        os.path.join(os.path.dirname(__file__), "./api/api.yaml"),
    ]
    module = importlib.import_module("openapiart.openapiart")
    openapiart_class = getattr(module, "OpenApiArt")
    openapiart = openapiart_class(
        api_files=api_files,
        output_dir="./art",
        python_module_name="sanity",
        protobuf_file_name="sanity",
        protobuf_package_name="sanity",
        go_module_name="openapiart",
        extension_prefix="sanity",
    )
    return openapiart


def create_snappi_artifacts():
    import openapiart

    openapiart.OpenApiArt(
        api_files=[
            "../../../models/api/info.yaml",
            "../../../models/api/api.yaml",
        ],
        output_dir="./art",
        python_module_name="snappi",
        protobuf_file_name="otg",
        protobuf_package_name="otg",
        go_module_name="snappi",
        extension_prefix="snappi",
    )


if __name__ == "__main__":
    create_openapi_artifacts()
    # create_snappi_artifacts()
