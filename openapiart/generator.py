"""Generator

Given an openapi.yaml file that has been produced by the Bundler class in the
bundler.py file the Generator class will produce an enhanced python ux file.

TBD:
- packet slicing using constants
- docstrings
- type checking
"""
import sys
import yaml
import os
import re
import requests
import pkgutil
import importlib
from jsonpath_ng import parse
from .openapiartplugin import OpenApiArtPlugin

MODELS_RELEASE = "v0.3.3"


class FluentRpc(object):
    """SetConfig(config Config) error"""

    def __init__(self):
        self.method = None
        self.operation_name = None
        self.request_class = None
        self.request_property = None
        self.good_response_type = None  # considering 200-ok
        self.bad_responses = []  # != 200-ok
        self.description = None
        self.http_method = None
        self.good_response_property = None
        self.x_status = None


class Generator:
    """Generates python classes based on an openapi.yaml file produced by the
    bundler.py infrastructure.
    """

    def __init__(
        self,
        openapi_filename,
        package_name,
        protobuf_package_name,
        output_dir=None,
        extension_prefix=None,
        generate_version_api=None,
        api_version=None,
        sdk_version=None,
    ):
        self._parsers = {}
        self._base_url = ""
        self._generated_methods = []
        self._generated_classes = []
        self._generated_top_level_factories = []
        self._openapi_filename = openapi_filename
        self._extension_prefix = extension_prefix
        self.__python = os.path.normpath(sys.executable)
        self.__python_dir = os.path.dirname(self.__python)
        self._src_dir = output_dir
        self._output_dir = os.path.join(output_dir, package_name)
        if os.path.exists(self._output_dir) is False:
            os.mkdir(self._output_dir)
        self._package_name = package_name
        self._protobuf_package_name = protobuf_package_name
        self._output_file = package_name
        self._docs_dir = os.path.join(self._src_dir, "..", "docs")
        self._deprecated_properties = {}
        self._generate_version_api = generate_version_api
        if self._generate_version_api is None:
            self._generate_version_api = False
        self._api_version = api_version
        if self._api_version is None:
            self._api_version = ""
        self._sdk_version = sdk_version
        if self._sdk_version is None:
            self._sdk_version = ""
        self._get_openapi_file()
        # self._plugins = self._load_plugins()

    def _get_parser(self, pattern):
        if pattern not in self._parsers:
            parser = parse(pattern)
            self._parsers[pattern] = parser
        else:
            parser = self._parsers[pattern]
        return parser

    def _load_plugins(self):
        plugins = []
        pkg_dir = os.path.dirname(__file__)
        for (_, name, _) in pkgutil.iter_modules([pkg_dir]):
            module_name = "openapiart." + name
            importlib.import_module(module_name)
            obj = sys.modules[module_name]
            for dir_name in dir(obj):
                if dir_name.startswith("_"):
                    continue
                dir_obj = getattr(obj, dir_name)
                print(dir_obj)
                if issubclass(dir_obj.__class__, OpenApiArtPlugin):
                    plugins.append(dir_obj)
        return plugins

    def _get_openapi_file(self):
        if self._openapi_filename is None:
            OPENAPI_URL = (
                "https://github.com/open-traffic-generator/models/releases"
                "/download/%s/openapi.yaml"
            ) % MODELS_RELEASE
            response = requests.request(
                "GET", OPENAPI_URL, allow_redirects=True
            )
            if response.status_code != 200:
                raise Exception(
                    "Unable to retrieve the Open Traffic Generator openapi.yaml"
                    " file [%s]" % response.content
                )
            openapi_content = response.content

            project_dir = os.path.dirname(os.path.dirname(__file__))
            with open(os.path.join(project_dir, "models-release"), "w") as out:
                out.write(MODELS_RELEASE)
        else:
            with open(self._openapi_filename, "rb") as fp:
                openapi_content = fp.read()
        self._openapi = yaml.safe_load(openapi_content)
        self._openapi_version = self._openapi["info"]["version"]
        print("generating using model version %s" % self._openapi_version)

    def generate(self):
        self._base_url = ""
        self._get_base_url()
        self._api_filename = os.path.join(
            self._output_dir, self._output_file + ".py"
        )
        with open(self._api_filename, "w") as self._fid:
            self._fid.write(
                "# {} {}\n".format(
                    self._openapi["info"]["title"],
                    self._openapi_version,
                )
            )
            self._fid.write(
                "# License: {}\n".format(
                    self._openapi["info"]["license"]["name"]
                ),
            )
            self._fid.write("\n")
        with open(
            os.path.join(os.path.dirname(__file__), "common.py"), "r"
        ) as fp:
            common_content = fp.read()
            cnf_text = "import sanity_pb2_grpc as pb2_grpc"
            modify_text = "try:\n    from {pkg_name} {text}\nexcept ImportError:\n    {text}".format(
                pkg_name=self._package_name,
                text=cnf_text.replace("sanity", self._protobuf_package_name),
            )
            common_content = common_content.replace(cnf_text, modify_text)

            cnf_text = "import sanity_pb2 as pb2"
            modify_text = "try:\n    from {pkg_name} {text}\nexcept ImportError:\n    {text}".format(
                pkg_name=self._package_name,
                text=cnf_text.replace("sanity", self._protobuf_package_name),
            )
            common_content = common_content.replace(cnf_text, modify_text)

            if re.search(r"def[\s+]api\(", common_content) is not None:
                self._generated_top_level_factories.append("api")
            if self._extension_prefix is not None:
                common_content = common_content.replace(
                    r'"{}_{}".format(__name__, ext)',
                    r'"'
                    + self._extension_prefix
                    + r"_{}."
                    + self._package_name
                    + r'_api".format(ext)',
                )
        with open(self._api_filename, "a") as self._fid:
            self._fid.write(common_content)
        methods, factories, rpc_methods = self._get_methods_and_factories()
        self._write_api_class(methods, factories)
        self._write_http_api_class(methods)
        self._write_rpc_api_class(rpc_methods)
        self._write_init()
        # TODO: restore behavior
        # self._write_deprecator()
        return self

    def _get_base_url(self):
        self._base_url = ""
        if "servers" in self._openapi:
            server = self._openapi["servers"][0]
            try:
                self._base_url = server["variables"]["basePath"]["default"]
                if not self._base_url.startswith("/"):
                    self._base_url = "/" + self._base_url
            except KeyError:
                pass

    def _write_init(self):
        filename = os.path.join(self._output_dir, "__init__.py")
        with open(filename, "w") as self._fid:
            for class_name in self._generated_classes:
                self._write(
                    0, "from .%s import %s" % (self._output_file, class_name)
                )
            for factory_name in self._generated_top_level_factories:
                self._write(
                    0, "from .%s import %s" % (self._output_file, factory_name)
                )

    def _find(self, path, schema_object):
        finds = self._get_parser(path).find(schema_object)
        for find in finds:
            yield find.value
            self._get_parser(path).find(find.value)

    def _get_methods_and_factories(self):
        """
        Parse methods and top level objects from yaml file to be later used in
        code generation.
        """
        methods = []
        rpc_methods = []
        factories = []
        refs = []
        self._top_level_schema_refs = []

        # parse methods
        for path in self._get_api_paths():
            operation = path["operation"]
            method_name = operation["operationId"].replace(".", "_").lower()
            if method_name in self._generated_methods:
                continue
            self._generated_methods.append(method_name)
            print("found method %s" % method_name)
            rpc = FluentRpc()
            rpc.method = method_name
            rpc.operation_name = self._get_external_struct_name(
                operation["operationId"].replace(".", "_")
            )
            rpc.description = self._get_description(operation)
            request = self._get_parser("$..requestBody..schema").find(
                operation
            )
            for req in request:
                (
                    _,
                    property_name,
                    class_name,
                    ref,
                ) = self._get_object_property_class_names(req.value)
                if ref:
                    refs.append(ref)
                    rpc.request_property = property_name
                    rpc.request_class = class_name

            response_type = None
            response_list = operation.get("responses")
            if response_list is None:
                raise Exception("{} should have responses".format(method_name))
            for response_code, response_property in response_list.items():
                if int(response_code) == 200:
                    rpc.good_response_property = response_property
                    schema_obj = self._get_parser("$..schema").find(
                        response_property
                    )
                    if len(schema_obj) == 0:
                        (
                            response_name,
                            _,
                            _,
                            ref,
                        ) = self._get_object_property_class_names(
                            response_property
                        )
                        if response_name is not None:
                            response = self._get_parser('$.."$ref"').find(
                                self._get_object_from_ref(ref)
                            )
                            if len(response) > 0:
                                (
                                    _,
                                    response_type,
                                    _,
                                    ref,
                                ) = self._get_object_property_class_names(
                                    response[0].value
                                )
                                if ref:
                                    refs.append(ref)
                    else:
                        (
                            _,
                            response_type,
                            _,
                            ref,
                        ) = self._get_object_property_class_names(
                            schema_obj[0].value
                        )
                        if ref:
                            refs.append(ref)
                else:
                    rpc.bad_responses.append(str(response_code))

            if response_type is None:
                # TODO: response type is usually None for schema which does not
                # contain any ref (e.g. response of POST /results/capture)
                pass

            rpc.good_response_type = response_type
            rpc.http_method = path["method"]
            # TODO: restore behavior
            # if "x-status" in path["operation"] and path["operation"][
            #     "x-status"
            # ].get("status") in ["deprecated", "under-review"]:
            #     rpc.x_status = (
            #         path["operation"]["x-status"]["status"],
            #         path["operation"]["x-status"]["additional_information"],
            #     )
            methods.append(
                {
                    "name": method_name,
                    "args": ["self"]
                    if len(request) == 0
                    else ["self", "payload"],
                    "http_method": path["method"],
                    "url": self._base_url + path["url"],
                    "description": self._get_description(operation),
                    "response_type": response_type,
                    # TODO: restore behavior
                    # "x_status": (
                    #     path["operation"]["x-status"]["status"],
                    #     path["operation"]["x-status"][
                    #         "additional_information"
                    #     ],
                    # )
                    # if path["operation"].get("x-status", {}).get("status")
                    # in ["deprecated", "under-review"]
                    # else None,
                }
            )
            rpc_methods.append(rpc)

        # parse top level objects (arguments for API requests)
        for ref in refs:
            if ref in self._generated_methods:
                continue
            self._generated_methods.append(ref)
            ret = self._get_object_property_class_names(ref)
            _, property_name, class_name, _ = ret
            schema_object = self._get_object_from_ref(ref)
            if "type" not in schema_object:
                continue
            print("found top level factory method %s" % property_name)
            if schema_object["type"] == "array":
                ref = schema_object["items"]["$ref"]
                _, _, class_name, _ = self._get_object_property_class_names(
                    ref
                )
                class_name = "%sIter" % class_name
                self._top_level_schema_refs.append((ref, property_name))
            self._top_level_schema_refs.append((ref, None))

            factories.append({"name": property_name, "class_name": class_name})

        for ref, property_name in self._top_level_schema_refs:
            if property_name is None:
                self._write_openapi_object(ref)
            else:
                self._write_openapi_list(ref, property_name)

        return methods, factories, rpc_methods

    def _write_rpc_api_class(self, rpc_methods):
        class_code = """class GrpcApi(Api):
    # OpenAPI gRPC Api
    def __init__(self, **kwargs):
        super(GrpcApi, self).__init__(**kwargs)
        self._stub = None
        self._channel = None
        self._request_timeout = 10
        self._keep_alive_timeout = 10 * 1000
        self._location = (
            kwargs["location"]
            if "location" in kwargs and kwargs["location"] is not None
            else "localhost:50051"
        )
        self._transport = kwargs["transport"] if "transport" in kwargs else None
        self._logger = kwargs["logger"] if "logger" in kwargs else None
        self._loglevel = kwargs["loglevel"] if "loglevel" in kwargs else logging.DEBUG
        if self._logger is None:
            stdout_handler = logging.StreamHandler(sys.stdout)
            formatter = logging.Formatter(fmt="%(asctime)s [%(name)s] [%(levelname)s] %(message)s", datefmt="%Y-%m-%d %H:%M:%S")
            formatter.converter = time.gmtime
            stdout_handler.setFormatter(formatter)
            self._logger = logging.Logger(self.__module__, level=self._loglevel)
            self._logger.addHandler(stdout_handler)
        self._logger.debug("gRPCTransport args: {}".format(", ".join(["{}={!r}".format(k, v) for k, v in kwargs.items()])))

    def _get_stub(self):
        if self._stub is None:
            CHANNEL_OPTIONS = [('grpc.enable_retries', 0),
                               ('grpc.keepalive_timeout_ms', self._keep_alive_timeout)]
            self._channel = grpc.insecure_channel(self._location, options=CHANNEL_OPTIONS)
            self._stub = pb2_grpc.OpenapiStub(self._channel)
        return self._stub

    def _serialize_payload(self, payload):
        if not isinstance(payload, (str, dict, OpenApiBase)):
            raise Exception("We are supporting [str, dict, OpenApiBase] object")
        if isinstance(payload, OpenApiBase):
            payload = payload.serialize()
        if isinstance(payload, dict):
            payload = json.dumps(payload)
        return payload

    @property
    def request_timeout(self):
        \"\"\"duration of time in seconds to allow for the RPC.\"\"\"
        return self._request_timeout

    @request_timeout.setter
    def request_timeout(self, timeout):
        self._request_timeout = timeout

    @property
    def keep_alive_timeout(self):
        return self._keep_alive_timeout

    @keep_alive_timeout.setter
    def keep_alive_timeout(self, timeout):
        self._keep_alive_timeout = timeout * 1000

    def close(self):
        if self._channel is not None:
            self._channel.close()
            self._channel = None
            self._stub = None"""

        self._generated_classes.append("Transport")
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, class_code)
            for rpc_method in rpc_methods:
                self._write()
                # TODO: restore behavior
                # if rpc_method.x_status is not None:
                #     self._write(
                #         1,
                #         "@OpenApiStatus.{func}".format(
                #             func=rpc_method.x_status[0].replace("-", "_")
                #         ),
                #     )
                #     key = "{}.{}".format("GrpcApi", rpc_method.method)
                #     self._deprecated_properties[key] = rpc_method.x_status[1]
                if rpc_method.request_class is None:
                    self._write(1, "def %s(self):" % rpc_method.method)
                    self._write(2, "stub = self._get_stub()")
                    self._write(
                        2,
                        "empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()",
                    )
                    self._write(
                        2,
                        "res_obj = stub.%s(empty, timeout=self._request_timeout)"
                        % rpc_method.operation_name,
                    )
                else:
                    self._write(
                        1, "def %s(self, payload):" % rpc_method.method
                    )
                    self._write(2, "pb_obj = json_format.Parse(")
                    self._write(3, "self._serialize_payload(payload),")
                    self._write(3, "pb2.%s()" % rpc_method.request_class)
                    self._write(2, ")")
                    if (
                        self._generate_version_api
                        and rpc_method.method != "get_version"
                    ):
                        self._write(2, "self._do_version_check_once()")
                    "%s=pb_obj" % rpc_method.request_property
                    self._write(
                        2,
                        "req_obj = pb2.{operation_name}Request({request_property}=pb_obj)".format(
                            operation_name=rpc_method.operation_name,
                            request_property=rpc_method.request_property,
                        ),
                    )
                    self._write(2, "stub = self._get_stub()")
                    self._write(
                        2,
                        "res_obj = stub.%s(req_obj, timeout=self._request_timeout)"
                        % rpc_method.operation_name,
                    )
                including_default, return_byte = self._process_good_response(
                    rpc_method
                )
                self._write(2, "response = json_format.MessageToDict(")
                self._write(3, "res_obj, preserving_proto_field_name=True")
                self._write(2, ")")
                self._write(
                    2, 'status_code_200 = response.get("status_code_200")'
                )
                self._write(2, "if status_code_200 is not None:")
                if return_byte:
                    self._write(
                        3, "return io.BytesIO(res_obj.status_code_200)"
                    )
                elif rpc_method.good_response_type:
                    if including_default:
                        self._write(3, "if len(status_code_200) == 0:")
                        self._write(
                            4, "status_code_200 = json_format.MessageToDict("
                        )
                        self._write(5, "res_obj.status_code_200,")
                        self._write(5, "preserving_proto_field_name=True,")
                        self._write(5, "including_default_value_fields=True")
                        self._write(4, ")")
                    self._write(
                        3,
                        "return self.%s().deserialize("
                        % rpc_method.good_response_type,
                    )
                    self._write(4, "status_code_200")
                    self._write(3, ")")
                else:
                    self._write(3, 'return response.get("status_code_200")')
                for rsp_code in rpc_method.bad_responses:
                    self._write(
                        2,
                        """if response.get("status_code_{code}") is not None:""".format(
                            code=rsp_code
                        ),
                    )
                    self._write(
                        3,
                        """raise Exception({code}, response.get("status_code_{code}"))""".format(
                            code=rsp_code
                        ),
                    )

    def _process_good_response(self, rpc_method):
        including_default = False
        return_byte = False
        property = rpc_method.good_response_property
        _, _, _, ref = self._get_object_property_class_names(property)
        if ref is not None:
            property = self._get_object_from_ref(ref)
        content = self._get_parser("$..content").find(property)
        if len(content) > 0 and "application/octet-stream" in content[0].value:
            return_byte = True

        parse_warnings = False
        if len(content) > 0:
            value = content[0].value
            if "application/json" in value:
                schema = value["application/json"].get("schema")
                if schema is not None:
                    _, _, _, ref = self._get_object_property_class_names(
                        schema
                    )
                    if ref is not None:
                        schema = self._get_object_from_ref(ref)
                    if (
                        schema.get("properties") is not None
                        and schema.get("properties").get("warnings")
                        is not None
                    ):
                        parse_warnings = True
        if (
            rpc_method.http_method in ["put", "post", "patch"]
            and parse_warnings
        ):
            including_default = True
        return including_default, return_byte

    def _write_http_api_class(self, methods):
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class HttpApi(Api):")
            self._write(1, '"""%s' % "OpenAPI HTTP Api")
            self._write(1, '"""')
            self._write(1, "def __init__(self, **kwargs):")
            self._write(2, "super(HttpApi, self).__init__(**kwargs)")
            self._write(2, "self._transport = HttpTransport(**kwargs)")
            self._write()
            self._write(1, "@property")
            self._write(1, "def verify(self):")
            self._write(2, "return self._transport.verify")
            self._write()
            self._write(1, "@verify.setter")
            self._write(1, "def verify(self, value):")
            self._write(2, "self._transport.set_verify(value)")

            for method in methods:
                print("generating method %s" % method["name"])
                self._write()
                # TODO: restore behavior
                # if method["x_status"] is not None:
                #     self._write(
                #         1,
                #         "@OpenApiStatus.{func}".format(
                #             func=method["x_status"][0].replace("-", "_")
                #         ),
                #     )
                #     key = "{}.{}".format("HttpApi", method["name"])
                #     self._deprecated_properties[key] = method["x_status"][1]
                self._write(
                    1,
                    "def %s(%s):"
                    % (method["name"], ", ".join(method["args"])),
                )
                self._write(
                    2,
                    '"""%s %s'
                    % (method["http_method"].upper(), method["url"]),
                )
                self._write(0)
                self._write(2, "%s" % method["description"])
                self._write(0)
                self._write(2, "Return: %s" % method["response_type"])
                self._write(2, '"""')

                if (
                    self._generate_version_api
                    and method["name"] != "get_version"
                ):
                    self._write(2, "self._do_version_check_once()")

                self._write(2, "return self._transport.send_recv(")
                self._write(3, '"%s",' % method["http_method"])
                self._write(3, '"%s",' % method["url"])
                self._write(
                    3,
                    "payload=%s,"
                    % (
                        method["args"][1]
                        if len(method["args"]) > 1
                        else "None"
                    ),
                )
                self._write(
                    3,
                    "return_object=%s,"
                    % (
                        "self." + method["response_type"] + "()"
                        if method["response_type"]
                        else "None"
                    ),
                )
                self._write(2, ")")

    def _write_api_class(self, methods, factories):
        self._generated_classes.append("Api")
        factory_class_name = "Api"
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class %s(object):" % factory_class_name)
            self._write(1, '"""%s' % "OpenApi Abstract API")
            self._write(1, '"""')
            self._write()
            self._write(1, "__warnings__ = []")
            self._write(1, "def __init__(self, **kwargs):")
            if self._generate_version_api:
                self._write(2, "self._version_meta = self.version()")
                self._write(
                    2,
                    'self._version_meta.api_spec_version = "{}"'.format(
                        self._api_version
                    ),
                )
                self._write(
                    2,
                    'self._version_meta.sdk_version = "{}"'.format(
                        self._sdk_version
                    ),
                )
                self._write(
                    2, 'self._version_check = kwargs.get("version_check")'
                )
                self._write(2, "if self._version_check is None:")
                self._write(3, "self._version_check = False")
                self._write(2, "self._version_check_err = None")
            else:
                self._write(2, "pass")
            for method in methods:
                print("generating method %s" % method["name"])
                self._write()
                # TODO: restore behavior
                # if method["x_status"] is not None:
                #     self._write(
                #         1,
                #         "@OpenApiStatus.{func}".format(
                #             func=method["x_status"][0].replace("-", "_")
                #         ),
                #     )
                #     key = "{}.{}".format(
                #         "%s" % factory_class_name, method["name"]
                #     )
                #     self._deprecated_properties[key] = method["x_status"][1]
                self._write(
                    1,
                    "def %s(%s):"
                    % (method["name"], ", ".join(method["args"])),
                )
                self._write(
                    2,
                    '"""%s %s'
                    % (method["http_method"].upper(), method["url"]),
                )
                self._write(0)
                self._write(2, "%s" % method["description"])
                self._write(0)
                self._write(2, "Return: %s" % method["response_type"])
                self._write(2, '"""')
                self._write(
                    2, 'raise NotImplementedError("%s")' % method["name"]
                )

            for factory in factories:
                print(
                    "generating top level factory method %s" % factory["name"]
                )
                self._write()
                self._write(1, "def %s(self):" % factory["name"])
                self._write(
                    2,
                    '"""Factory method that creates an instance of %s'
                    % (factory["class_name"]),
                )
                self._write()
                self._write(2, "Return: %s" % factory["class_name"])
                self._write(2, '"""')
                self._write(2, "return %s()" % factory["class_name"])

            self._write()
            self._write(1, "def close(self):")
            self._write(2, "pass")
            # self._write()
            # self._write(1, "def get_api_warnings(self):")
            # self._write(2, "return openapi_warnings")
            # self._write()
            # self._write(1, "def clear_api_warnings(self):")
            # self._write(
            #     2, 'if "2.7" in platform.python_version().rsplit(".", 1)[0]:'
            # )
            # self._write(3, "del openapi_warnings[:]")
            # self._write(2, "else:")
            # self._write(3, "openapi_warnings.clear()")
            self._generate_version_api_methods()

    def _generate_version_api_methods(self):
        if not self._generate_version_api:
            return

        self._write(
            indent=1,
            line="""def _check_client_server_version_compatibility(
        self, client_ver, server_ver, component_name
    ):
        try:
            c = semantic_version.Version(client_ver)
        except Exception as e:
            raise AssertionError(
                "Client {} version '{}' is not a valid semver: {}".format(
                    component_name, client_ver, e
                )
            )

        try:
            s = semantic_version.SimpleSpec(server_ver)
        except Exception as e:
            raise AssertionError(
                "Server {} version '{}' is not a valid semver: {}".format(
                    component_name, server_ver, e
                )
            )


        err = "Client {} version '{}' is not semver compatible with Server {} version '{}'".format(
            component_name, client_ver, component_name, server_ver
        )

        if not s.match(c):
            raise Exception(err)

    def get_local_version(self):
        return self._version_meta

    def get_remote_version(self):
        return self.get_version()

    def check_version_compatibility(self):
        comp_err, api_err = self._do_version_check()
        if comp_err is not None:
            raise comp_err
        if api_err is not None:
            raise api_err

    def _do_version_check(self):
        local = self.get_local_version()
        try:
            remote = self.get_remote_version()
        except Exception as e:
            return None, e

        try:
            self._check_client_server_version_compatibility(
                local.api_spec_version, remote.api_spec_version, "API spec"
            )
        except Exception as e:
            msg = "client SDK version '{}' is not compatible with server SDK version '{}'".format(
                local.sdk_version, remote.sdk_version
            )
            return Exception("{}: {}".format(msg, str(e))), None

        return None, None

    def _do_version_check_once(self):
        if not self._version_check:
            return

        if self._version_check_err is not None:
            raise self._version_check_err

        comp_err, api_err = self._do_version_check()
        if comp_err is not None:
            self._version_check_err = comp_err
            raise comp_err
        if api_err is not None:
            self._version_check_err = None
            raise api_err

        self._version_check = False
        self._version_check_err = None""",
        )

    def _get_object_property_class_names(self, ref):
        """Returns: `Tuple(object_name, property_name, class_name, ref_name)`"""
        object_name = None
        property_name = None
        class_name = None
        ref_name = None
        if isinstance(ref, dict) is True and "$ref" in ref:
            ref_name = ref["$ref"]
        elif isinstance(ref, str) is True:
            ref_name = ref
        if ref_name is not None:
            object_name = ref_name.split("/")[-1]
            property_name = object_name.lower().replace(".", "_")
            class_name = object_name.replace(".", "")
        return (object_name, property_name, class_name, ref_name)

    def _get_external_struct_name(self, openapi_name):
        return self._get_external_field_name(openapi_name).replace("_", "")

    def _get_external_field_name(self, openapi_name):
        """convert openapi fieldname to protobuf fieldname

        - reference: https://developers.google.com/protocol-buffers/docs/reference/go-generated#fields

        Note that the generated Go field names always use camel-case naming,
        even if the field name in the .proto file uses lower-case with underscores (as it should).
        The case-conversion works as follows:
        - The first letter is capitalized for export.
        - NOTE: this is ignored as OpenAPIArt doesn't allow fieldnames to start with an underscore
            - If the first character is an underscore, it is removed and a capital X is prepended.
        - If an interior underscore is followed by a lower-case letter, the underscore is removed, and the following letter is capitalized.
        - NOTE: This isn't documented, if a number is followed by a lower-case letter the following letter is capitalized.
        - Thus, the proto field foo_bar_baz becomes FooBarBaz in Go, and _my_field_name_2 becomes XMyFieldName_2.
        """
        external = ""
        name = openapi_name.replace(".", "")
        for i in range(len(name)):
            if i == 0:
                if name[i] == "_":
                    pass
                else:
                    external += name[i].upper()
            elif name[i] == "_":
                pass
            elif name[i - 1] == "_":
                if name[i].isdigit() or name[i].isupper():
                    external += "_" + name[i]
                else:
                    external += name[i].upper()
            elif name[i - 1].isdigit() and name[i].islower():
                external += name[i].upper()
            else:
                external += name[i]
        if external in ["String"]:
            external += "_"
        return external

    def _write_openapi_object(self, ref, choice_method_name=None):
        schema_object = self._get_object_from_ref(ref)
        ref_name = ref.split("/")[-1]
        class_name = ref_name.replace(".", "")
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        print("generating class %s" % (class_name))
        refs = []
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class %s(OpenApiObject):" % class_name)
            slots = ["'_parent'"]
            if "choice" in self._get_choice_names(schema_object):
                slots.append("'_choice'")
            self._write(1, "__slots__ = (%s)" % ",".join(slots))
            self._write()

            # write _TYPES definition
            # TODO: this func won't detect whether $ref for a given property is
            # a list because it relies on 'type' attribute to do so
            openapi_types = self._get_openapi_types(schema_object)
            if len(openapi_types) > 0:
                self._write(1, "_TYPES = {")
                for name, value in openapi_types:
                    if len(value) == 1:
                        self._write(
                            2,
                            "'%s': {'%s': %s},"
                            % (
                                name,
                                list(value.keys())[0],
                                list(value.values())[0],
                            ),
                        )
                        continue
                    self._write(2, "'%s': %s" % (name, "{"))
                    for n, v in value.items():
                        if isinstance(v, list):
                            self._write(3, "'%s': [" % n)
                            for i in v:
                                self._write(4, "'%s'," % i)
                            self._write(3, "],")
                            continue
                        self._write(3, "'%s': %s," % (n, v))
                    self._write(2, "},")
                self._write(1, "} # type: Dict[str, str]")
                self._write()
            else:
                # TODO: provide empty types as workaround because deserializer
                # in common.py currently expects it
                self._write(1, "_TYPES = {} # type: Dict[str, str]")
                self._write()

            required, defaults = self._get_required_and_defaults(schema_object)

            if len(required) > 0:
                self._write(
                    1, "_REQUIRED = {} # type: tuple(str)".format(required)
                )
                self._write()
            else:
                self._write(1, "_REQUIRED= () # type: tuple(str)")
                self._write()

            if len(defaults) > 0:
                self._write(1, "_DEFAULTS = {")
                for name, value in defaults:
                    if isinstance(value, (list, bool, int, float, tuple)):
                        self._write(2, "'%s': %s," % (name, value))
                    else:
                        self._write(2, "'%s': '%s'," % (name, value))
                self._write(1, "} # type: Dict[str, Union(type)]")
                self._write()
            else:
                self._write(1, "_DEFAULTS= {} # type: Dict[str, Union(type)]")
                self._write()

            # write constants
            # search for all simple properties with enum or
            # x-constant and add them here
            for enum in self._get_parser("$..enum | x-constants").find(
                schema_object
            ):
                for name in enum.value:
                    value = name
                    value_type = "string"
                    if isinstance(enum.value, dict):
                        value = enum.value[name]
                        value_type = (
                            enum.context.value["type"]
                            if "type" in enum.context.value
                            else "string"
                        )
                    if value_type == "string":
                        self._write(
                            1, "%s = '%s' # type: str" % (name.upper(), value)
                        )
                    else:
                        self._write(1, "%s = %s #" % (name.upper(), value))
                if len(enum.value) > 0:
                    self._write()

            # write def __init__(self)
            params = "self, parent=None"
            if "choice" in self._get_choice_names(schema_object):
                params += ", choice=None"
            init_params, properties, _ = self._get_property_param_string(
                schema_object
            )
            params = (
                params
                if len(init_params) == 0
                else ", ".join([params, init_params])
            )
            self._write(1, "def __init__(%s):" % (params))
            self._write(2, "super(%s, self).__init__()" % class_name)
            self._write(2, "self._parent = parent")
            for property_name in properties:
                self._write(
                    2,
                    "self._set_property('%s', %s)"
                    % (property_name, property_name),
                )
            if "choice" in self._get_choice_names(schema_object):
                self._write(
                    2,
                    "if 'choice' in self._DEFAULTS and choice is None and self._DEFAULTS['choice'] in self._TYPES:",
                )
                self._write(3, "getattr(self, self._DEFAULTS['choice'])")
                self._write(2, "else:")
                self._write(3, "self.choice = choice")

            # write def set(self)
            self._write_set_method(schema_object)

            # process properties - TBD use this one level up to process
            # schema, in requestBody, Response and also
            refs = self._process_properties(
                class_name,
                schema_object,
                choice_child=choice_method_name is not None,
            )

        # descend into child properties
        for ref in refs:
            self._write_openapi_object(ref[0], ref[3])
            if ref[1] is True:
                self._write_openapi_list(ref[0], ref[2])

    def _write_set_method(self, schema_object):
        write_set = False
        if "choice" in self._get_choice_names(schema_object):
            write_set = False
        init_params, properties, _ = self._get_property_param_string(
            schema_object
        )
        if len(init_params) > 0:
            write_set = True
        params = ["self"]
        for property in properties:
            str = property + "=None"
            params.append(str)
        params = params if len(init_params) == 0 else ", ".join(params)
        if write_set:
            self._write(1, "def set(%s):" % (params))
            self._write(
                2, "for property_name, property_value in locals().items():"
            )
            self._write(
                3,
                "if property_name != 'self' and property_value is not None:",
            )
            self._write(4, "self._set_property(property_name, property_value)")

    def _get_simple_type_names(self, schema_object):
        simple_type_names = []
        if "properties" in schema_object:
            choice_names = self._get_choice_names(schema_object)
            for name in schema_object["properties"]:
                if name in choice_names:
                    continue
                ref = self._get_parser("$..'$ref'").find(
                    schema_object["properties"][name]
                )
                if len(ref) == 0:
                    simple_type_names.append(name)
        return simple_type_names

    def _get_choice_names(self, schema_object):
        choice_names = []
        if (
            "properties" in schema_object
            and "choice" in schema_object["properties"]
        ):
            choice_names = schema_object["properties"]["choice"]["enum"][:]
            choice_names.append("choice")
        return choice_names

    def _process_properties(
        self, class_name=None, schema_object=None, choice_child=False
    ):
        """Process all properties of a /component/schema object
        Write a factory method for all choice
        If there are no properties then the schema_object is a primitive or array type
        """
        refs = []
        if "properties" in schema_object:
            choice_names = self._get_choice_names(schema_object)
            excluded_property_names = []
            for choice_name in choice_names:

                # this code is to allow choices with no properties
                if choice_name not in schema_object["properties"]:
                    excluded_property_names.append(choice_name)
                    continue

                if "$ref" not in schema_object["properties"][choice_name]:
                    continue
                ref = schema_object["properties"][choice_name]["$ref"]
                # TODO: restore behavior
                # status = schema_object["properties"][choice_name].get(
                #     "x-status"
                # )
                self._write_factory_method(
                    None, choice_name, ref, property_status=None
                )
                excluded_property_names.append(choice_name)
            for property_name in schema_object["properties"]:
                if property_name in excluded_property_names:
                    continue
                property = schema_object["properties"][property_name]
                write_set_choice = (
                    property_name in choice_names and property_name != "choice"
                )
                self._write_openapi_property(
                    schema_object,
                    property_name,
                    property,
                    class_name,
                    write_set_choice,
                )
            for property_name, property in schema_object["properties"].items():
                ref = self._get_parser("$..'$ref'").find(property)
                if len(ref) > 0:
                    restriction = self._get_type_restriction(property)
                    choice_name = (
                        property_name
                        if property_name in excluded_property_names
                        else None
                    )
                    refs.append(
                        (
                            ref[0].value,
                            restriction.startswith("List["),
                            property_name,
                            choice_name,
                        )
                    )
        return refs

    def _write_openapi_list(self, ref, property_name):
        """This is the class writer for schema object properties that are of
        type array with a ref to an object.  The class should provide a factory
        method for the encapsulated ref.
        ```
        properties:
          ports:
            type: array
            items:
              $ref: '#/components/schema/...'
        ```

        If the schema object has a property named choice, that property needs
        to be brought forward so that the generated class can provide factory
        methods for objects for each of the choice $refs (if any).

        if choice exists:
            for each choice enum that is a $ref:
                generate a factory method named after the choice
                in the method set the choice property
        """
        yobject = self._get_object_from_ref(ref)
        ref_name = ref.split("/")[-1]
        contained_class_name = ref_name.replace(".", "")
        class_name = "%sIter" % contained_class_name
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        self._imports = []
        print("generating class %s" % (class_name))
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class %s(OpenApiIter):" % class_name)
            self._write(1, "__slots__ = ('_parent', '_choice')")
            self._write()

            # if all choice(s) are $ref, the getitem should return the actual choice object
            # the _GETITEM_RETURNS_CHOICE_OBJECT class static allows the OpenApiIter to
            # correctly return the selected choice if any
            get_item_returns_choice = True
            if "properties" in yobject and "choice" in yobject["properties"]:
                for property, item in yobject["properties"].items():
                    if property == "choice":
                        continue
                    elif (
                        property in yobject["properties"]["choice"]["enum"]
                        and len(self._get_parser('$.."$ref"').find(item)) > 0
                    ):
                        continue
                    get_item_returns_choice = False
                    break
            else:
                get_item_returns_choice = False
            self._write(
                1,
                "_GETITEM_RETURNS_CHOICE_OBJECT = {}".format(
                    get_item_returns_choice
                ),
            )

            self._write()
            self._write(1, "def __init__(self, parent=None, choice=None):")
            self._write(2, "super(%s, self).__init__()" % class_name)
            self._write(2, "self._parent = parent")
            self._write(2, "self._choice = choice")

            # write container emulation methods __getitem__, __iter__, __next__
            self._write_openapilist_special_methods(
                contained_class_name, yobject
            )

            # write a factory method for the schema object in the list that returns the container
            self._write_factory_method(
                contained_class_name,
                ref_name.lower().split(".")[-1],
                ref,
                True,
                False,
            )

            # write an add method for the schema object in the list that creates and returns the new object
            self._write_add_method(
                yobject,
                ref,
                False,
                class_name,
                contained_class_name,
                class_name,
            )

            # write choice factory methods if the only properties are choice properties
            if get_item_returns_choice is True:
                for property, item in yobject["properties"].items():
                    if property == "choice":
                        continue
                    self._write_factory_method(
                        contained_class_name,
                        property,
                        item["$ref"],
                        True,
                        True,
                    )

        return class_name

    def _write_openapilist_special_methods(
        self, contained_class_name, schema_object
    ):
        get_item_class_names = [contained_class_name]
        if (
            "properties" in schema_object
            and "choice" in schema_object["properties"]
        ):
            for property in schema_object["properties"]:
                if property in schema_object["properties"]["choice"]["enum"]:
                    if "$ref" in schema_object["properties"][property]:
                        ref = schema_object["properties"][property]["$ref"]
                        (
                            _,
                            _,
                            choice_class_name,
                            _,
                        ) = self._get_object_property_class_names(ref)
                        if choice_class_name not in get_item_class_names:
                            get_item_class_names.append(choice_class_name)
        get_item_class_names.sort()
        self._write()
        self._write(1, "def __getitem__(self, key):")
        self._write(
            2, "# type: (str) -> Union[%s]" % (", ".join(get_item_class_names))
        )
        self._write(2, "return self._getitem(key)")
        self._write()
        self._write(1, "def __iter__(self):")
        self._write(2, "# type: () -> %sIter" % contained_class_name)
        self._write(2, "return self._iter()")
        self._write()
        self._write(1, "def __next__(self):")
        self._write(2, "# type: () -> %s" % contained_class_name)
        self._write(2, "return self._next()")
        self._write()
        self._write(1, "def next(self):")
        self._write(2, "# type: () -> %s" % contained_class_name)
        self._write(2, "return self._next()")
        self._write()
        self._write(1, "def _instanceOf(self, item):")
        self._write(2, "if not isinstance(item, %s):" % (contained_class_name))
        self._write(
            3,
            'raise Exception("Item is not an instance of %s")'
            % (contained_class_name),
        )

    def _write_factory_method(
        self,
        contained_class_name,
        method_name,
        ref,
        openapi_list=False,
        choice_method=False,
        property_status=None,
    ):
        yobject = self._get_object_from_ref(ref)
        _, _, class_name, _ = self._get_object_property_class_names(ref)
        (
            param_string,
            properties,
            type_string,
        ) = self._get_property_param_string(yobject)
        self._write()
        if openapi_list is True:
            self._imports.append(
                "from .%s import %s" % (class_name.lower(), class_name)
            )
            params = (
                "self"
                if len(param_string) == 0
                else ", ".join(["self", param_string])
            )
            self._write(1, "def %s(%s):" % (method_name, params))
            return_class_name = class_name
            if contained_class_name is not None:
                return_class_name = "{}Iter".format(contained_class_name)
            self._write(
                2, "# type: (%s) -> %s" % (type_string, return_class_name)
            )
            self._write(
                2,
                '"""Factory method that creates an instance of the %s class'
                % (class_name),
            )
            self._write()
            self._write(2, "%s" % self._get_description(yobject))
            self._write()
            self._write(2, "Returns: %s" % (return_class_name))
            self._write(2, '"""')
            if choice_method is True:
                self._write(2, "item = %s()" % (contained_class_name))
                self._write(2, "item.%s" % (method_name))
                self._write(2, "item.choice = '%s'" % (method_name))
            else:
                params = ["parent=self._parent"]
                if (
                    "properties" in yobject
                    and "choice" in yobject["properties"]
                ):
                    params.append("choice=self._choice")
                params.extend(["%s=%s" % (name, name) for name in properties])
                self._write(
                    2, "item = %s(%s)" % (class_name, ", ".join(params))
                )
            self._write(2, "self._add(item)")
            self._write(2, "return self")
            self._write()
        else:
            self._write(1, "@property")
            # TODO: restore behavior
            # if property_status is not None and property_status in [
            #     "deprecated",
            #     "under-review",
            # ]:
            #     self._write(
            #         1,
            #         "@OpenApiStatus.{func}".format(
            #             func=property_status.replace("-", "_")
            #         ),
            #     )
            #     key = "{}.{}".format(class_name, method_name)
            #     self._deprecated_properties[key] = property["x-status"][
            #         "additional_information"
            #     ]
            self._write(1, "def %s(self):" % (method_name))
            self._write(2, "# type: () -> %s" % (class_name))
            self._write(
                2,
                '"""Factory property that returns an instance of the %s class'
                % (class_name),
            )
            self._write()
            self._write(2, "%s" % self._get_description(yobject))
            self._write()
            self._write(2, "Returns: %s" % (class_name))
            self._write(2, '"""')
            self._write(
                2,
                "return self._get_property('%s', %s, self, '%s')"
                % (method_name, class_name, method_name),
            )

    def _write_add_method(
        self,
        yobject,
        ref,
        choice_method,
        class_name,
        contained_class_name,
        return_class_name,
    ):
        """Writes an add method"""
        method_name = ref.lower().split("/")[-1]
        self._imports.append(
            "from .%s import %s"
            % (contained_class_name.lower(), contained_class_name)
        )
        (
            param_string,
            properties,
            type_string,
        ) = self._get_property_param_string(yobject)
        params = (
            "self"
            if len(param_string) == 0
            else ", ".join(["self", param_string])
        )
        self._write(1, "def add(%s):" % (params))
        self._write(
            2, "# type: (%s) -> %s" % (type_string, contained_class_name)
        )
        self._write(
            2,
            '"""Add method that creates and returns an instance of the %s class'
            % (contained_class_name),
        )
        self._write()
        self._write(2, "%s" % self._get_description(yobject))
        self._write()
        self._write(2, "Returns: %s" % (contained_class_name))
        self._write(2, '"""')
        if choice_method is True:
            self._write(2, "item = self.%s()" % (method_name))
            self._write(2, "item.%s" % (contained_class_name))
            self._write(2, "item.choice = '%s'" % (contained_class_name))
        else:
            params = ["parent=self._parent"]
            if "properties" in yobject and "choice" in yobject["properties"]:
                params.append("choice=self._choice")
            params.extend(["%s=%s" % (name, name) for name in properties])
            self._write(
                2, "item = %s(%s)" % (contained_class_name, ", ".join(params))
            )
        self._write(2, "self._add(item)")
        self._write(2, "return item")

    def _get_property_param_string(self, yobject):
        property_param_string = []
        property_type_string = []
        properties = []
        if "properties" in yobject:
            for name, property in yobject["properties"].items():
                if name == "choice":
                    continue
                default = None
                type_string = self._get_type_restriction(property)
                if "obj" not in type_string:
                    if "default" in property:
                        default = property["default"]
                    if name == "choice":
                        val = "None"
                    elif property["type"] in [
                        "number",
                        "integer",
                        "boolean",
                        "array",
                    ]:
                        val = "None" if default is None else default
                    else:
                        val = (
                            "None"
                            if default is None
                            else "'{}'".format(default.strip())
                        )
                    properties.append(name)
                    property_param_string.append("%s=%s" % (name, val))
                    property_type_string.append(type_string)
        types = ",".join(property_type_string)
        return (", ".join(property_param_string), properties, types)

    def _write_openapi_property(
        self, schema_object, name, property, klass_name, write_set_choice=False
    ):
        ref = self._get_parser("$..'$ref'").find(property)
        restriction = self._get_type_restriction(property)
        if len(ref) > 0:
            object_name = ref[0].value.split("/")[-1]
            class_name = object_name.replace(".", "")
            if restriction.startswith("List["):
                type_name = "%sIter" % class_name
            else:
                type_name = class_name
        else:
            type_name = restriction
        self._write()
        self._write(1, "@property")
        # TODO: restore behavior
        # if property.get("x-status", {}).get("status") in [
        #     "deprecated",
        #     "under-review",
        # ]:
        #     func = property["x-status"]["status"].replace("-", "_")
        #     self._write(1, "@OpenApiStatus.{func}".format(func=func))
        #     key = "{}.{}".format(klass_name, name)
        #     self._deprecated_properties[key] = property["x-status"][
        #         "additional_information"
        #     ]
        self._write(1, "def %s(self):" % name)
        self._write(2, "# type: () -> %s" % (type_name))
        self._write(2, '"""%s getter' % (name))
        self._write()
        self._write(2, self._get_description(property))
        self._write()
        self._write(2, "Returns: %s" % type_name)
        self._write(2, '"""')
        if (
            len(self._get_parser("$..'type'").find(property)) > 0
            and len(ref) == 0
        ):
            self._write(2, "return self._get_property('%s')" % (name))
            self._write()
            if name == "auto":
                return
            self._write(1, "@%s.setter" % name)
            # TODO: restore behavior
            # if property.get("x-status", {}).get("status") in [
            #     "deprecated",
            #     "under-review",
            # ]:
            #     func = property["x-status"]["status"].replace("-", "_")
            #     self._write(1, "@OpenApiStatus.{func}".format(func=func))
            #     key = "{}.{}".format(klass_name, name)
            #     self._deprecated_properties[key] = property["x-status"][
            #         "additional_information"
            #     ]
            self._write(1, "def %s(self, value):" % name)
            self._write(2, '"""%s setter' % (name))
            self._write()
            self._write(2, self._get_description(property))
            self._write()
            self._write(2, "value: %s" % restriction)
            self._write(2, '"""')
            required, defaults = self._get_required_and_defaults(schema_object)
            if len(required) > 0:
                if name in required:
                    self._write(2, "if value is None:")
                    self._write(
                        3,
                        "raise TypeError('Cannot set required property %s as None')"
                        % name,
                    )
            if write_set_choice is True:
                self._write(
                    2, "self._set_property('%s', value, '%s')" % (name, name)
                )
            else:
                self._write(2, "self._set_property('%s', value)" % (name))
        elif len(ref) > 0:
            if restriction.startswith("List["):
                self._write(
                    2,
                    "return self._get_property('%s', %sIter, self._parent, self._choice)"
                    % (name, class_name),
                )
            else:
                self._write(
                    2,
                    "return self._get_property('%s', %s)" % (name, class_name),
                )

    def _get_description(self, yobject):
        if "description" not in yobject:
            yobject["description"] = "TBD"
        # remove tabs, multiple spaces
        description = re.sub(r"\n", ". ", yobject["description"])
        description = re.sub(r"\s+", " ", description)
        return description
        # doc_string = []
        # for line in re.split('\. ', description):
        #     line = re.sub('\.$', '', line)
        #     if len(line) > 0:
        #         doc_string.append('%s  ' % line)
        # return doc_string

    def _get_data_types(self, yproperty):
        data_type_map = {
            "integer": "int",
            "string": "str",
            "boolean": "bool",
            "array": "list",
            "number": "float",
            "float": "float",
            "double": "float",
        }
        if yproperty["type"] in data_type_map:
            return data_type_map[yproperty["type"]]
        else:
            return yproperty["type"]

    def _get_openapi_types(self, yobject):
        types = []
        if "properties" in yobject:
            for name in yobject["properties"]:
                yproperty = yobject["properties"][name]
                ref = self._get_parser("$..'$ref'").find(yproperty)
                pt = {}
                if "type" in yproperty:
                    pt.update({"type": self._get_data_types(yproperty)})
                    pt.update(
                        {"enum": yproperty["enum"]}
                    ) if "enum" in yproperty else None
                    pt.update(
                        {"format": "'%s'" % yproperty["format"]}
                    ) if "format" in yproperty else None
                if len(ref) > 0:
                    object_name = ref[0].value.split("/")[-1]
                    class_name = object_name.replace(".", "")
                    if "type" in yproperty and yproperty["type"] == "array":
                        class_name += "Iter"
                    pt.update({"type": "'%s'" % class_name})
                if (
                    len(ref) == 0
                    and "items" in yproperty
                    and "type" in yproperty["items"]
                ):
                    pt.update(
                        {"itemtype": self._get_data_types(yproperty["items"])}
                    )
                    if "format" in yproperty["items"]:
                        pt.update(
                            {
                                "itemformat": "'%s'"
                                % yproperty["items"]["format"]
                            }
                        )
                min_max = yproperty.get("maximum", yproperty.get("minimum", 0))
                key = (
                    "itemformat"
                    if pt.get("itemtype") is not None
                    else "format"
                )
                if min_max > 2147483647:
                    pt.update({key: r"'int64'"})
                if len(ref) == 0 and "minimum" in yproperty:
                    pt.update({"minimum": yproperty["minimum"]})
                if len(ref) == 0 and "maximum" in yproperty:
                    pt.update({"maximum": yproperty["maximum"]})
                if len(ref) == 0 and "minLength" in yproperty:
                    pt.update({"minLength": yproperty["minLength"]})
                if len(ref) == 0 and "maxLength" in yproperty:
                    pt.update({"maxLength": yproperty["maxLength"]})
                if len(pt) > 0:
                    types.append((name, pt))
                # TODO: restore behavior
                # if "x-constraint" in yproperty:
                #     cons_lst = []
                #     for cons in yproperty["x-constraint"]:
                #         ref, prop = cons.split("/properties/")
                #         klass = self._get_classname_from_ref(ref)
                #         cons_lst.append("%s.%s" % (klass, prop.strip("/")))
                #     if cons_lst != []:
                #         pt.update({"constraint": cons_lst})
                # if "x-unique" in yproperty:
                #     pt.update({"unique": '"%s"' % yproperty["x-unique"]})
        return types

    def _get_required_and_defaults(self, yobject):
        required = []
        defaults = []
        if "required" in yobject:
            required = yobject["required"]
        if "properties" in yobject:
            for name in yobject["properties"]:
                yproperty = yobject["properties"][name]
                if "default" in yproperty:
                    default = yproperty["default"]
                    if "type" in yproperty and yproperty["type"] == "number":
                        default = float(default)
                    defaults.append((name, default))
        return (tuple(required), defaults)

    def _get_default_value(self, property):
        if "default" in property:
            return property["default"]
        property_type = property["type"]
        if property_type == "array":
            return "[]"
        if property_type == "string":
            return "''"
        if property_type == "integer":
            return 0
        if property_type == "number":
            return 0
        if property_type == "bool":
            return False
        raise Exception(
            "Missing handler for property type `%s`" % property_type
        )

    def _get_api_paths(self):
        paths = []
        for url, yobject in self._openapi["paths"].items():
            for method in yobject:
                if method.lower() in ["get", "post", "put", "patch", "delete"]:
                    paths.append(
                        {
                            "url": url,
                            "method": method,
                            "operation": yobject[method],
                        }
                    )
        return paths

    def _write_data_properties(self, schema, classname, choice_tuples):
        if len(choice_tuples) > 0:
            choices = []
            for choice_tuple in choice_tuples:
                choices.append(choice_tuple[0])
            self._write(
                2,
                "if isinstance(choice, (%s)) is False:" % (", ".join(choices)),
            )
            self._write(
                3,
                "raise TypeError('choice must be of type: %s')"
                % (", ".join(choices)),
            )
            self._write(
                2,
                "self.__setattr__('choice', %s._CHOICE_MAP[type(choice).__name__])"
                % classname,
            )
            self._write(
                2,
                "self.__setattr__(%s._CHOICE_MAP[type(choice).__name__], choice)"
                % classname,
            )

        if "properties" in schema:
            for name, property in schema["properties"].items():
                if (
                    len([item for item in choice_tuples if item[1] == name])
                    == 0
                    and name != "choice"
                ):
                    restriction = self._get_isinstance_restriction(
                        schema, name, property
                    )
                    self._write(
                        2,
                        "if isinstance(%s, %s) is True:" % (name, restriction),
                    )
                    if restriction == "(list, type(None))":
                        self._write(
                            3,
                            "self.%s = [] if %s is None else list(%s)"
                            % (name, name, name),
                        )
                    else:
                        if "pattern" in property:
                            self._write(3, "import re")
                            self._write(
                                3,
                                "assert(bool(re.match(r'%s', %s)) is True)"
                                % (property["pattern"], name),
                            )
                        self._write(3, "self.%s = %s" % (name, name))
                    self._write(2, "else:")
                    self._write(
                        3,
                        "raise TypeError('%s must be an instance of %s')"
                        % (name, restriction),
                    )

    def _get_isinstance_restriction(self, schema, name, property):
        type_none = ", type(None)"
        if "required" in schema and name in schema["required"]:
            type_none = ""
        if "$ref" in property:
            return "(%s%s)" % (
                self._get_classname_from_ref(property["$ref"]),
                type_none,
            )
        elif name == "additionalProperties":
            return "**additional_properties"
        elif property["type"] in ["number", "integer"]:
            return "(float, int%s)" % type_none
        elif property["type"] == "string":
            return "(str%s)" % type_none
        elif property["type"] == "array":
            return "(list%s)" % type_none
        elif property["type"] == "boolean":
            return "(bool%s)" % type_none

    def _get_type_restriction(self, property):
        try:
            if "$ref" in property:
                ref_obj = self._get_object_from_ref(property["$ref"])
                description = ""
                if "description" in ref_obj:
                    description = ref_obj["description"]
                if "description" in property:
                    description += property["description"]
                property["description"] = description
                class_name = property["$ref"].split("/")[-1].replace(".", "")
                return "obj(%s)" % class_name
            elif "oneOf" in property:
                return "Union[%s]" % ",".join(
                    [item["type"] for item in property["oneOf"]]
                )
            elif property["type"] == "number":
                return "float"
            elif property["type"] == "integer":
                return "int"
            elif property["type"] == "string":
                if "enum" in property:
                    values = property["enum"]
                    values.sort()
                    values = ['Literal["{}"]'.format(s) for s in values]
                    return "Union[%s]" % ", ".join(values)
                else:
                    return "str"
            elif property["type"] == "array":
                return "List[%s]" % self._get_type_restriction(
                    property["items"]
                )
            elif property["type"] == "boolean":
                return "bool"
        except Exception as e:
            print("Error ", property, e)
            raise e

    def _get_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf

    def _get_classname_from_ref(self, ref):
        final_piece = ref.split("/")[-1]
        if "." in final_piece:
            return final_piece.split(".")[-1]
        else:
            return final_piece

    def _write(self, indent=0, line=""):
        self._fid.write("    " * indent + line + "\n")

    def _write_deprecator(self):
        with open(self._api_filename, "a") as self._fid:
            self._write(0, "OpenApiStatus.messages = {")
            for klass, msg in self._deprecated_properties.items():
                if "\n" in msg:
                    print(msg)
                    self._write(1, '"{}" : """{}""",'.format(klass, msg))
                else:
                    self._write(1, '"{}" : "{}",'.format(klass, msg))
            self._write(1, "}")
