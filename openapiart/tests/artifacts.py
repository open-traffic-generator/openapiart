import pytest
import sys
import os
import importlib
import logging


def create_openapi_artifacts():
    """Using the model files in the api directory create the following artifacts:
    - ./art/openapi.yaml (bundled/validated model file)
    - ./art/openapi.json (json version of openapi.yaml)
    - ./art/sanity.proto (protobuf version of openapi.yaml)
    - ./art/openapi.html (html doc of openapi.yaml if node and redoc-cli is installed)
    - ./art/sanity/__init__.py (python ux package)
    - ./art/sanity/sanity.py (python ux package)
    - ./art/sanity/sanity_pb2.py (python protobuf stubs)
    - ./art/sanity/sanity_pb2_grpc.py (python grpc stubs)
    - ./art/sanity/sanity_pb2.go (go protobuf stubs)
    - ./art/sanity/sanity_pb2_grpc.go (go grpc stubs)
    - ./art/sanity/go.mod (go ux package)
    - ./art/sanity/sanity.go (go ux package)
    """
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
        go_module_name="sanityux",
        extension_prefix="sanity",
    )
    return openapiart


if __name__ == "__main__":
    create_openapi_artifacts()
