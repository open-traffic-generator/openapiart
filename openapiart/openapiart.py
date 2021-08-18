import sys
import os
import importlib
import shutil
import yaml
import subprocess
import requests


class OpenApiArt(object):
    """Bundle and generate artifacts from OpenAPI files.

    Args
    ----
    - api_files (list[str]): list of OpenAPI files that contain info and/or path
      keywords
    - python_module_name (str): name of the consolidated python file that will be generated
    - output_dir (str): directory where artifacts will be created.
      Unless otherwise specified the default directory for generated artifacts
      is `current working directory/.output`.
      The artifacts that will be generated are:
      - openapi.yaml
      - openapi.json
      - openapi.html (static documentation, if redoc-cli has been installed)
      - python package
      - protobuf file
      - python grpc
    """

    def __init__(
        self,
        api_files,
        python_module_name=None,
        protobuf_package_name=None,
        go_sdk_package_dir=None,
        go_sdk_package_name=None,
        output_dir=None,
        extension_prefix=None,
    ):
        self._python_module_name = python_module_name
        self._protobuf_package_name = protobuf_package_name
        self._go_sdk_package_dir = go_sdk_package_dir
        self._go_sdk_package_name = go_sdk_package_name
        self._extension_prefix = extension_prefix
        self._output_dir = os.path.abspath(output_dir)
        print("Artifact output directory: {output_dir}".format(output_dir=self._output_dir))
        shutil.rmtree(self._output_dir, ignore_errors=True)
        self._api_files = api_files
        self._bundle()
        self._get_license()
        self._get_info()
        self._document()
        self._generate()

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
        """Try documenting the openapi using redoc-cli"""
        try:
            process_args = [
                "redoc-cli",
                "bundle",
                self._bundler.openapi_filepath,
                "--output",
                os.path.join(self._output_dir, "openapi.html"),
            ]
            subprocess.check_call(process_args, shell=True)
        except Exception as e:
            print("Bypassed creation of static documentation [missing redoc-cli]: {}".format(e))

    def _generate(self):
        # this generates the python ux module
        if self._python_module_name is not None:
            module = importlib.import_module("openapiart.generator")
            python_ux = getattr(module, "Generator")(
                self._bundler.openapi_filepath,
                self._python_module_name,
                output_dir=self._output_dir,
                extension_prefix=self._extension_prefix,
            )
            python_ux.generate()

        # this generates protobuf definitions into the output_dir
        if self._protobuf_package_name and self._go_sdk_package_dir:
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

        # this generates the python stubs into the output dir/python module
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

        # this generates the go stubs
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

    @property
    def output_dir(self):
        return self._output_dir

    @property
    def python_module_name(self):
        return self._python_module_name
