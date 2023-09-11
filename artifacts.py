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
import shutil


def create_openapi_artifacts(openapiart_class, sdk=None):

    base_dir = os.path.abspath(os.path.dirname(__file__))
    artifacts_dir = os.path.join(base_dir, "artifacts")

    open_api = openapiart_class(
        api_files=[
            os.path.join(
                os.path.dirname(__file__), "./openapiart/tests/api/info.yaml"
            ),
            os.path.join(
                os.path.dirname(__file__),
                "./openapiart/tests/common/common.yaml",
            ),
            os.path.join(
                os.path.dirname(__file__), "./openapiart/tests/api/api.yaml"
            ),
            # os.path.join(os.path.dirname(__file__), "./openapiart/goserver/api/api.yaml"),
            os.path.join(
                os.path.dirname(__file__),
                "./openapiart/goserver/api/service_a.api.yaml",
            ),
            os.path.join(
                os.path.dirname(__file__),
                "./openapiart/goserver/api/service_b.api.yaml",
            ),
        ],
        artifact_dir=os.path.join(os.path.dirname(__file__), "artifacts"),
        extension_prefix="sanity",
        proto_service="Openapi",
        generate_version_api=True,
    )
    if sdk == "proto" or sdk is None or sdk == "all":
        open_api.GenerateProtoDef(package_name="sanity")

    if sdk == "python" or sdk is None or sdk == "all":
        open_api.GeneratePythonSdk(package_name="sanity")
        shutil.move(
            os.path.join(artifacts_dir, "requirements.txt"),
            os.path.join(artifacts_dir, "sanity"),
        )

    if sdk == "go" or sdk is None or sdk == "all":
        open_api.GenerateGoSdk(
            package_dir="github.com/open-traffic-generator/openapiart/pkg",
            package_name="openapiart",
            sdk_version="0.0.1",
        )
        open_api.GenerateGoServer(
            module_path="github.com/open-traffic-generator/openapiart/pkg",
            models_prefix="openapiart",
            models_path="github.com/open-traffic-generator/openapiart/pkg",
        )
        open_api.GoTidy(
            relative_package_dir="pkg",
        )
        # copy all the files to artifacts/openapiart
        go_path = os.path.join(artifacts_dir, "openapiart_go")
        shutil.copytree(
            os.path.join(base_dir, "pkg", "httpapi"),
            os.path.join(go_path, "httpapi"),
            dirs_exist_ok=True,
        )
        shutil.copytree(
            os.path.join(base_dir, "pkg", "sanity"),
            os.path.join(go_path, "sanity"),
            dirs_exist_ok=True,
        )
        # files = ["openapiart.go", "go.mod", "go.sum"]
        # shutil.copy(os.path.join)


if __name__ == "__main__":
    sdk = None
    # import pdb; pdb.set_trace()
    if len(sys.argv) >= 2:
        sdk = sys.argv[1]
    if len(sys.argv) == 3:
        cicd = sys.argv
    try:
        from openapiart.openapiart import OpenApiArt as openapiart_class
    except:
        if not cicd:
            sys.path.append(
                os.path.normpath(os.path.join(os.path.dirname(__file__), ".."))
            )
            module = importlib.import_module("openapiart.openapiart")
            openapiart_class = getattr(module, "OpenApiArt")
        else:
            raise Exception(
                "Error: Not able to import openapiart module with the generated sdk"
            )
    create_openapi_artifacts(openapiart_class, sdk)
