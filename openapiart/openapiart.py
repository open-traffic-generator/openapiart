import sys
import os
import importlib
import shutil
import yaml
import subprocess
import requests
import platform


class OpenApiArt(object):
    """Bundle and generate artifacts from OpenAPI files.

    Args
    ----
    - api_files (list[str]): list of OpenAPI files that contain info and/or path
      keywords
    - protobuf_name (str): name of the .proto file that will be generated
    - artifact_dir (str): directory where artifacts will be created.
      Unless otherwise specified the default directory for generated artifacts
      is `current working directory/.art`.
    - extension_prefix (str): name of the python extension
    """

    def __init__(
        self,
        api_files,
        protobuf_name=None,
        artifact_dir=None,
        extension_prefix=None,
    ):
        self._output_dir = os.path.abspath(artifact_dir if artifact_dir is not None else "art")
        self._go_sdk_package_dir = None
        self._protobuf_package_name = protobuf_name if protobuf_name is not None else "sanity"
        self._extension_prefix = extension_prefix if extension_prefix is not None else "sanity"

        print("Artifact output directory: {output_dir}".format(output_dir=self._output_dir))
        shutil.rmtree(self._output_dir, ignore_errors=True)
        self._api_files = api_files
        self._bundle()
        self._get_license()
        self._get_info()
        self._document()

    def _get_license(self):
        try:
            license_name = self._bundler._content["info"]["license"]["name"]
            self._license = "License: {}".format(license_name)
        except:
            raise Exception("The /info/license/name is a REQUIRED property and must be present in the schema.")

    def _get_info(self):
        try:
            self._info = "{} {}".format(self._bundler._content["info"]["title"], self._bundler._content["info"]["version"])
        except Exception as e:
            self._info = "OpenAPI info error [{}]".format(e)

    def _bundle(self):
        # bundle the yaml files
        module = importlib.import_module("openapiart.bundler")
        bundler_class = getattr(module, "Bundler")
        self._bundler = bundler_class(api_files=self._api_files, output_dir=self._output_dir)
        self._bundler.bundle()
        # read the entire openapi file
        with open(self._bundler.openapi_filepath) as fp:
            self._openapi = yaml.safe_load(fp.read())

    def _document(self):
        """Try documenting the openapi using redoc-cli

        - Requires nodejs, npm, redoc-cli
        - npm install -g redoc-cli

        If the requirements are not installed/reachable then the document step
        will be bypassed but the generation will not fail.
        """
        try:
            process_args = [
                "redoc-cli",
                "bundle",
                self._bundler.openapi_filepath,
                "--output",
                os.path.join(self._output_dir, "openapi.html"),
            ]
            subprocess.check_call(process_args, shell=platform.system() == "Windows")
        except Exception as e:
            print("Bypassed creation of static documentation: {}".format(e))

    def GeneratePythonSdk(self, package_name):
        """Generates a Python UX Sdk
        Args
        ----
        - package_name (str): the name of the python module

        Example
        -------
        ```
        art = Openapiart(api_files=["<list of open_api_file_path>"], artifact_dir="./")
        art.GeneratePythonSdk(package_name="sanity")
        ```

        Output
        ------
        ```
        ./sanity
                |_ __init__.py
                |_ sanity.py
                |_ sanity_pb.py
                |_ sanity_grpc_pb.py
        ```
        """
        self._python_module_name = package_name
        self._generate_proto_file()
        if self._python_module_name is not None:
            module = importlib.import_module("openapiart.generator")
            python_ux = getattr(module, "Generator")(
                self._bundler.openapi_filepath,
                self._python_module_name,
                output_dir=self._output_dir,
                extension_prefix=self._extension_prefix,
            )
            python_ux.generate()
        try:
            python_sdk_dir = os.path.normpath(os.path.join(self._output_dir, self._python_module_name))
            process_args = [
                sys.executable,
                "-m",
                "grpc_tools.protoc",
                "--python_out={}".format(python_sdk_dir),
                "--grpc_python_out={}".format(python_sdk_dir),
                "--proto_path={}".format(self._output_dir),
                "{}.proto".format(self._protobuf_package_name),
            ]
            print("Generating python grpc stubs: {}".format(" ".join(process_args)))
            subprocess.check_call(process_args, shell=False)
        except Exception as e:
            print("Bypassed creation of python stubs: {}".format(e))
        return self

    def GenerateGoSdk(self, package_dir, package_name):
        """Generates a Go UX Sdk

        Args
        ----
        - package_dir: Go mod package dir published under go.mod
        - package_name: Name of the Go package to generate

        Example
        -------
        ```
        art = Openapiart(api_files=["list of open_api_file_path"], artifact_dir="./")
        art.GenerateGoSdk(
            package_dir="github.com/<path to repo>/$package_name",
            package_name="sanity"
        )
        ```

        Output
        ------
        ```
        ./sanity
                |_ sanity.go
                |_ go.mod
                |_ go.sum
                |_ sanitypb
                            |_ sanitypb.pb.go
                            |_ sanitypb_grpc.go
        ```
        """

        self._go_sdk_package_dir = package_dir
        self._go_sdk_package_name = package_name
        self._generate_proto_file()
        if self._go_sdk_package_dir and self._protobuf_package_name:
            go_sdk_output_dir = os.path.normpath(os.path.join(self._output_dir, "..", os.path.split(self._go_sdk_package_dir)[-1]))
            go_protobuffer_out_dir = os.path.normpath(os.path.join(go_sdk_output_dir, self._protobuf_package_name))
            if not os.path.exists(go_protobuffer_out_dir):
                os.makedirs(go_protobuffer_out_dir)
            process_args = [
                "protoc",
                "--go_opt=paths=source_relative",
                "--go-grpc_opt=paths=source_relative",
                "--go_out={}".format(go_protobuffer_out_dir),
                "--go-grpc_out={}".format(go_protobuffer_out_dir),
                "--proto_path={}".format(self._output_dir),
                "--experimental_allow_proto3_optional",
                "{}.proto".format(self._protobuf_package_name),
            ]
            cmd = " ".join(process_args)
            print("Generating go gRPC stubs: {}".format(cmd))
            subprocess.check_call(cmd, shell=True)

        # this generates the go ux module
        if self._protobuf_package_name and self._go_sdk_package_dir:
            module = importlib.import_module("openapiart.openapiartgo")
            go_ux = getattr(module, "OpenApiArtGo")(
                **{
                    "info": self._info,
                    "license": self._license,
                    "protobuf_package_name": self._protobuf_package_name,
                    "go_sdk_package_dir": self._go_sdk_package_dir,
                    "go_sdk_package_name": self._go_sdk_package_name,
                    "output_dir": self._output_dir,
                }
            )
            print("Generating go ux sdk: {}".format(" ".join(process_args)))
            go_ux.generate(self._openapi)
        return self

    def _generate_proto_file(self):
        if self._protobuf_package_name is None:
            self._protobuf_package_name = "default"
        module = importlib.import_module("openapiart.openapiartprotobuf")
        protobuf = getattr(module, "OpenApiArtProtobuf")(
            **{
                "info": self._info,
                "license": self._license,
                "protobuf_package_name": self._protobuf_package_name,
                "go_sdk_package_dir": self._go_sdk_package_dir,
                "output_dir": self._output_dir,
            }
        )
        protobuf.generate(self._openapi)

    @property
    def output_dir(self):
        return self._output_dir

    @property
    def python_module_name(self):
        return self._python_module_name
