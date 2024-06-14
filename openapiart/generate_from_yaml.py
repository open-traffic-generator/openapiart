import os

# import sys
import yaml

# sys.path.insert(0, os.getcwd())
# print(os.getcwd())
# print(sys.path)
from .openapiart import OpenApiArt as openapiart_class


class GenerateFromYaml(object):
    """
    This class takes in yaml file which basically acts as a config file according
    to which openapiart operations are truggered.
    example yaml:

    api_files:
        - openapiart/tests/api/info.yaml
        - openapiart/tests/api/api.yaml
        - openapiart/goserver/api/service_a.api.yaml
        - openapiart/goserver/api/service_b.api.yaml
    artifact_dir: artifacts
    generate_version_api: true
    languages:
        python:
            package_name: pyapi
        go:
            sdk:
                package_dir: github.com/open-traffic-generator/goapi/pkg
                package_name: goapi
                sdk_version: 0.0.1
            server:
                module_path: github.com/open-traffic-generator/goapi/pkg
                models_prefix: openapi
                models_path: github.com/open-traffic-generator/goapi/pkg
            tidy:
                relative_package_dir: pkg

    """

    def __init__(self, yaml_file):
        """
        Takes in yaml config ,does basic checking initiates generation.
        """

        self._file = yaml_file

        self._file = os.path.normpath(os.path.abspath(self._file))

        if not os.path.exists(self._file):
            raise Exception("the file %s does not exsist" % self._file)

        with open(self._file) as fp:
            self._config = yaml.safe_load(fp.read())

        self._initiate_generation()

    def _initiate_generation(self):
        """
        common place to initial openapiart operations
        """

        self._initiate_bundling()
        if self._config.get("languages") is not None:
            langs = self._config.get("languages")
            if langs.get("python") is not None:
                self._generate_python(langs["python"])
            if langs.get("go") is not None:
                self._generate_go(langs["go"])

    def _initiate_bundling(self):
        """
        mainly handles the bunling and json, yaml generation
        """

        if self._config.get("api_files") is None:
            raise Exception("api_files is a mandatory property in yaml")

        if self._config.get("artifact_dir") is None:
            raise Exception("artifact_dir is a mandatory property in yaml")

        print("\n\nStarting Bundling  Process:\n\n")

        files = self._config.get("api_files")
        validated_files = []
        for file in files:
            file_name = os.path.normpath(os.path.abspath(file))
            if not os.path.exists(file_name):
                raise Exception(
                    "%s file in api_files does not exsists" % file_name
                )
            validated_files.append(file_name)

        artifact_dir = self._config.get("artifact_dir", "art")
        artifact_dir = os.path.normpath(os.path.abspath(artifact_dir))
        proto_service = self._config.get("proto_service", "Openapi")
        protobuf_name = self._config.get("protobuf_name", "openapi")
        extension_prefix = self._config.get("extension_prefix", "openapi")
        generate_version_api = self._config.get("generate_version_api", True)

        self._openapiart = openapiart_class(
            api_files=validated_files,
            protobuf_name=protobuf_name,
            artifact_dir=artifact_dir,
            extension_prefix=extension_prefix,
            proto_service=proto_service,
            generate_version_api=generate_version_api,
        )

    def _generate_python(self, config):
        """
        Initiates python sdk geenration.
        """
        print("\n\nStarting Python SDK generation \n\n")
        if config.get("package_name") is None:
            raise Exception(
                "package_name is a mandatory parameter to generate python sdk"
            )

        self._openapiart.GeneratePythonSdk(
            package_name=config.get("package_name")
        )

    def _generate_go(self, config):
        """
        Initiates go sdk and server generations
        """
        if config.get("sdk") is not None:
            print("\n\nStarting Go SDK generation\n\n")
            go_sdk = config.get("sdk")

            if "package_dir" not in go_sdk or "package_name" not in go_sdk:
                raise Exception(
                    "package_dir and package_name are manadatory for go sdk generation"
                )

            package_dir = go_sdk.get("package_dir")
            package_name = go_sdk.get("package_name")
            sdk_version = go_sdk.get("sdk_version")

            self._openapiart.GenerateGoSdk(
                package_dir=package_dir,
                package_name=package_name,
                sdk_version=sdk_version,
            )

        if config.get("server") is not None:
            print("\n\nStarting Go Server generation\n\n")
            go_server = config.get("server")

            if (
                "module_path" not in go_server
                or "models_prefix" not in go_server
                or "models_path" not in go_server
            ):
                raise Exception(
                    "module_path, models_prefix and models_path are manadatory for go server generation"
                )

            module_path = go_server.get("module_path")
            models_prefix = go_server.get("models_prefix")
            models_path = go_server.get("models_path")

            self._openapiart.GenerateGoServer(
                module_path=module_path,
                models_prefix=models_prefix,
                models_path=models_path,
            )

        if config.get("tidy") is not None:
            print("\n\nStarting Go Tidy Process\n\n")
            go_tidy = config.get("tidy")

            if "relative_package_dir" not in go_tidy:
                raise Exception(
                    "relative_package_dir is a mandatory for go tidy"
                )

            self._openapiart.GoTidy(
                relative_package_dir=go_tidy["relative_package_dir"]
            )
