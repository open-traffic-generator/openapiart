"""Using the model files in the api directory create the following artifacts:
"""
import sys
import os
import importlib


def create_go_server_artifacts(openapiart_class):
    openapiart_class(
        api_files=[
        os.path.join(os.path.dirname(__file__), "./openapiart/goserver/api/service_a.api.yaml"),
        os.path.join(os.path.dirname(__file__), "./openapiart/goserver/api/service_b.api.yaml"),
        os.path.join(os.path.dirname(__file__), "./openapiart/goserver/api/api.yaml"),
    ],
        artifact_dir=os.path.join(os.path.dirname(__file__), "art_go/doc"),
        protobuf_name="models_pb"
    ).GenerateGoSdk(
        # package_dir="github.com/open-traffic-generator/openapiart/arg_go/models",
        package_dir="localdev/art_go/models",
        package_name="models"
    ).GenerateGoServer(
        module_path="localdev/art_go/pkg",
        models_path="localdev/art_go/models"
    )
    models_folder = os.path.join(os.path.dirname(__file__), "art_go/models")
    os.remove(os.path.join(models_folder, 'go.mod'))
    os.remove(os.path.join(models_folder, 'go.sum'))




if __name__ == "__main__":
    sys.path.append(os.path.normpath(os.path.join(os.path.dirname(__file__), "..")))
    module = importlib.import_module("openapiart.openapiart")
    openapiart_class = getattr(module, "OpenApiArt")
    create_go_server_artifacts(openapiart_class)
