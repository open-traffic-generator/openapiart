import os
import sys
import yaml

sys.path.insert(0, os.getcwd())
print(os.getcwd())
print(sys.path)


class GenerateFromYaml(object):
    def __init__(self, yaml_file):
        pass


from openapiart import OpenApiArt as openapiart_class

openapiart_class(
    api_files=[
        os.path.join(
            "/home/otg/openapiart", "./openapiart/tests/api/info.yaml"
        ),
        os.path.join(
            "/home/otg/openapiart",
            "./openapiart/tests/common/common.yaml",
        ),
        os.path.join(
            "/home/otg/openapiart", "./openapiart/tests/api/api.yaml"
        ),
        # os.path.join(/home/otg/openapiart, "./openapiart/goserver/api/api.yaml"),
        os.path.join(
            "/home/otg/openapiart",
            "./openapiart/goserver/api/service_a.api.yaml",
        ),
        os.path.join(
            "/home/otg/openapiart",
            "./openapiart/goserver/api/service_b.api.yaml",
        ),
    ],
    artifact_dir=os.path.join("/home/otg/openapiart", "art"),
    extension_prefix="sanity",
    proto_service="Openapi",
    generate_version_api=True,
)
