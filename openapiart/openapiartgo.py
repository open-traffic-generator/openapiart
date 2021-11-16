from .openapiartplugin import OpenApiArtPlugin
import os
import subprocess
import random


class FluentStructure(object):
    def __init__(self):
        self.internal_struct_name = None
        self.external_interface_name = None
        self.external_new_methods = []
        self.external_rpc_methods = []
        self.internal_http_methods = []
        self.components = {}


class FluentRpc(object):
    """SetConfig(config Config) error"""

    def __init__(self):
        self.operation_name = None
        self.method = None
        self.request = "emptypb.Empty{}"
        self.responses = []
        self.http_call = None
        self.method_description = None


class FluentRpcResponse(object):
    """<operation_name>StatusCode<status_code>
    status_code is the http status code
    fluent_new is the 2xx response, all other status codes this should be None
    """

    def __init__(self):
        self.status_code = None
        self.schema = None
        self.request_return_type = None
        self.description = None
        self.method_description = None


class FluentHttp(object):
    """httpSetConfig(config Config) error"""

    def __init__(self):
        self.operation_name = None
        self.method = None
        self.request = None
        self.request_return_type = None
        self.responses = []
        self.description = None


class FluentNew(object):
    """New<external_interface_name> <external_interface_name>"""

    def __init__(self):
        self.generated = False
        self.interface = None
        self.struct = None
        self.method = None
        self.schema_name = None
        self.schema_object = None
        self.isRpcResponse = False
        self.description = None
        self.method_description = None
        self.interface_fields = []

    def isOptional(self, property_name):
        if self.schema_object is None:
            return True
        if "required" not in self.schema_object:
            return True
        if property_name not in self.schema_object["required"]:
            return True
        return False


class FluentField(object):
    def __init__(self):
        self.name = None
        self.description = None
        self.getter_method_description = None
        self.setter_method_description = None
        self.has_method_description = None
        self.iter_method_description = None
        self.method_description = None
        self.type = None
        self.isOptional = True
        self.isPointer = True
        self.isArray = False
        self.struct = None
        self.external_struct = None
        self.setter_method = None
        self.getter_method = None
        self.adder_method = None
        self.has_method = None
        self.format = None  # only for mac, ipv4, ipv6 and hex validation
        self.default = None
        self.itemformat = None  # only for mac, ipv4, ipv6 and hex validation
        self.hasminmax = False
        self.min = None
        self.max = None
        self.hasminmaxlength = False
        self.min_length = None
        self.max_length = None


class OpenApiArtGo(OpenApiArtPlugin):
    """Generates a fluent interface go package that encapsulates protoc
    generated .pg.go and _grpc.pb.go content

    Toolchain
    ---------
    api/**.yaml | OpenAPIArt.bundler
    out: openapi.yaml | openapiartprotobuf
    out: sanity.proto | protoc
    out: sanity.pb.go, sanity_grpc.pb.go | openapiartgo (fluent wrapper around pb.go, _grpc.pb.go)
    output: currently handwritten poc.go (want generated sanity.go)

    - Update the openartprotobuf.py to generate grpc-gateway stubs
        import "google/api/annotations.proto";

        rpc SetConfig(SetConfigRequest) SetConfigResponse {
            option (google.api.http) = {
                post: "/config",
                body: "*"
            };
        }

    - Generate the .proto file using artifacts.py

    - From the .proto generate the .pb.go and _grpc.pb.go
        protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --experimental_allow_proto3_optional .proto

    - From the bundled openapi.yaml generate a custom fluent interface using
    openapiartgo.py that encapsulates the generated .pb.go and _grpc.pb.go
        - features:
            - top level api, api interface
            - api transport (grpc, http) ->
            - path item object operationId -> api interface member
            - path item object input schema -> part of api interface, internal struct, external interface
            - all components/schemas internal struct, external interface


    Contents should be <pkgname>.go, go.mod, go.sum under a specific directory

    Workflow/action should install go
        - https://github.com/actions/setup-go
    Which needs to copy the generated go file to the test dir
    The action then needs to run the go test to ensure all tests pass
    """

    def __init__(self, **kwargs):
        super(OpenApiArtGo, self).__init__(**kwargs)
        self._api = FluentStructure()
        self._api_interface_methods = []
        self._oapi_go_types = {
            "string": "string",
            "boolean": "bool",
            "integer": "int32",
            "int64": "int64",
            "number": "float32",
            "numberfloat": "float32",
            "numberdouble": "float64",
            # "stringmac": "StringMac",
            # "stringipv4": "StringIpv4",
            # "stringipv6": "StringIpv6",
            # "stringhex": "StringHex",
            "stringbinary": "[]byte",
        }

    def generate(self, openapi, generate_unit_test=False):
        self._openapi = openapi
        self._ux_path = os.path.normpath(os.path.join(self._output_dir, "..", os.path.split(self._go_sdk_package_dir)[-1]))
        self._protoc_path = os.path.normpath(os.path.join(self._ux_path, self._protobuf_package_name))
        self._structs = {}
        self._write_mod_file()
        self._write_go_file()
        self._format_go_file()
        self._tidy_mod_file()
        if generate_unit_test:
            self._generate_go_unit_test()

    def _write_mod_file(self):
        self._filename = os.path.normpath(os.path.join(self._ux_path, "go.mod"))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write("module {}".format(self._go_sdk_package_dir))
        self._write()
        self._write("go 1.16")
        self._close_fp()

    def _write_go_file(self):
        self._filename = os.path.normpath(os.path.join(self._ux_path, "{}.go".format(self._go_sdk_package_name)))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write_package_docstring(self._openapi["info"])
        self._write_package()
        self._write_common_code()
        self._write_types()
        self._build_api_interface()
        self._build_request_interfaces()
        self._write_component_interfaces()
        self._close_fp()

    def _write_package_docstring(self, info_object):
        """Write the header of the generated go code file which consists of:
        - license, version and description
        - package name
        - common custom code
        """
        self._write("// {}".format(self._info))
        for line in self._license.split("\n"):
            self._write("// {}".format(line))
        self._write()

    def _write_package(self):
        self._write("package {go_sdk_package_name}".format(go_sdk_package_name=self._go_sdk_package_name))
        self._write()

    def _write_common_code(self):
        """Writes the base wrapper go code"""
        line = 'import {pb_pkg_name} "{go_sdk_pkg_dir}/{pb_pkg_name}"'.format(
            pb_pkg_name=self._protobuf_package_name, go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write(line)
        self._write('import "google.golang.org/protobuf/types/known/emptypb"')
        self._write('import "google.golang.org/grpc"')
        self._write('import "github.com/ghodss/yaml"')
        self._write('import "google.golang.org/protobuf/encoding/protojson"')
        self._write('import "github.com/golang/protobuf/proto"')
        go_pkg_fp = self._fp
        go_pkg_filename = self._filename
        self._filename = os.path.normpath(os.path.join(self._ux_path, "common.go"))
        self._init_fp(self._filename)
        self._write_package()
        with open(os.path.join(os.path.dirname(__file__), "common.go")) as fp:
            self._write(fp.read().strip().strip("\n"))
        self._write()
        self._fp = go_pkg_fp
        self._filename = go_pkg_filename

    def _write_types(self):
        for _, go_type in self._oapi_go_types.items():
            if go_type.startswith("String"):
                self._write("type {go_type} string".format(go_type=go_type))
        self._write()

    def _get_internal_name(self, openapi_name):
        name = self._get_external_struct_name(openapi_name)
        name = name[0].lower() + name[1:]
        if name in ["error"]:
            name = "_" + name
        return name

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

    def _get_external_struct_name(self, openapi_name):
        return self._get_external_field_name(openapi_name).replace("_", "")

    def _resolve_response(self, parser_result):
        """returns the inner response type if any"""
        if "/components/responses" in parser_result[0].value:
            jsonpath = "$.{}..schema".format(parser_result[0].value[2:].replace("/", "."))
            schema = self._get_parser(jsonpath).find(self._openapi)[0].value
            response_component_ref = self._get_parser("$..'$ref'").find(schema)
            return response_component_ref
        return parser_result

    def _build_api_interface(self):
        self._api.internal_struct_name = """{internal_name}Api""".format(
            internal_name=self._get_internal_name(self._go_sdk_package_name)
        )
        self._api.external_interface_name = """{external_name}Api""".format(
            external_name=self._get_external_struct_name(self._go_sdk_package_name)
        )
        for url, path_object in self._openapi["paths"].items():
            for operation_id in self._get_parser("$..operationId").find(path_object):
                path_item_object = operation_id.context.value
                rpc = FluentRpc()
                http = FluentHttp()
                rpc.operation_name = self._get_external_struct_name(operation_id.value)
                rpc.description = self._get_description(path_item_object, True)
                http.operation_name = self._get_external_struct_name(operation_id.value)
                http.description = self._get_description(path_item_object)
                if len([m for m in self._api.external_rpc_methods if m.operation_name == rpc.operation_name]) == 0:
                    self._api.external_rpc_methods.append(rpc)
                if len([m for m in self._api.internal_http_methods if m.operation_name == http.operation_name]) == 0:
                    self._api.internal_http_methods.append(http)
                rpc.request_return_type = "{operation_response_name}Response".format(
                    operation_response_name=self._get_external_struct_name(rpc.operation_name),
                )
                binary_type = self._get_parser("$..responses..'200'..schema..format").find(path_item_object)
                ref_type = self._get_parser("$..responses..'200'..'$ref'").find(path_item_object)
                if len(binary_type) == 1:
                    rpc.request_return_type = "[]byte"
                elif len(ref_type) == 1:
                    request_return_type = self._get_schema_object_name_from_ref(self._resolve_response(ref_type)[0].value)
                    rpc.request_return_type = self._get_external_struct_name(request_return_type)
                else:
                    rpc.request_return_type = "*string"

                http.request_return_type = rpc.request_return_type
                ref = self._get_parser("$..requestBody..'$ref'").find(path_item_object)
                if len(ref) == 1:
                    new = FluentNew()
                    new.schema_name = self._get_schema_object_name_from_ref(ref[0].value)
                    new.schema_object = self._get_schema_object_from_ref(ref[0].value)
                    new.interface = self._get_external_struct_name(new.schema_name)
                    new.struct = self._get_internal_name(new.schema_name)
                    new.description = self._get_description(new.schema_object, True)
                    new.method_description = """// New{interface} returns a new instance of {interface}.
                    """.format(
                        interface=new.interface) + "// {} is {}".format(
                            new.interface, self._get_description(new.schema_object, True).lstrip("// ")
                        )
                    new.method = """New{interface}() {interface}""".format(interface=new.interface)
                    if len([m for m in self._api.external_new_methods if m.schema_name == new.schema_name]) == 0:
                        self._api.external_new_methods.append(new)
                    rpc.request = "{pb_pkg_name}.{operation_name}Request{{{interface}: {struct}.Msg()}}".format(
                        pb_pkg_name=self._protobuf_package_name,
                        operation_name=rpc.operation_name,
                        interface=new.interface,
                        struct=new.struct,
                    )
                    rpc.description = "// {} {}".format(rpc.operation_name, rpc.description.lstrip("// ")) 
                    # """
                    #     // Performs {operation_name} on user provided {interface} and returns {request_return_type}
                    #     // or returns error on failure
                    #     """.format(
                    #     operation_name=rpc.operation_name,
                    #     operation_response_name=self._get_external_struct_name(rpc.operation_name),
                    #     struct=new.struct,
                    #     interface=new.interface,
                    #     request_return_type=rpc.request_return_type,
                    # )
                    rpc.method = """{operation_name}({struct} {interface}) ({request_return_type}, error)""".format(
                        operation_name=rpc.operation_name,
                        operation_response_name=self._get_external_struct_name(rpc.operation_name),
                        struct=new.struct,
                        interface=new.interface,
                        request_return_type=rpc.request_return_type,
                    )
                    rpc.validate = """
                        err := {struct}.Validate()
                        if err != nil {{
                            return nil, err
                        }}
                    """.format(
                        struct=new.struct
                    )
                    rpc.http_call = """return api.http{operation_name}({struct})""".format(
                        operation_name=rpc.operation_name,
                        struct=new.struct,
                    )
                    if url.startswith("/"):
                        url = url[1:]
                    http.request = """api.httpSendRecv("{url}", {struct}.ToJson(), "{method}")""".format(
                        operation_name=http.operation_name,
                        url=url,
                        struct=new.struct,
                        method=str(operation_id.context.path.fields[0]).upper(),
                    )
                    http.method = """http{rpc_method}""".format(rpc_method=rpc.method)
                else:
                    rpc.description = "// {} {}".format(rpc.operation_name, rpc.description.lstrip("// "))
                    # """
                    # // Perform {operation_name} and returns {request_return_type} on success
                    # // or error on failure
                    # """.format(
                    #     operation_name=rpc.operation_name,
                    #     operation_response_name=self._get_external_struct_name(rpc.operation_name),
                    #     request_return_type=rpc.request_return_type,
                    # )
                    rpc.method = """{operation_name}() ({request_return_type}, error)""".format(
                        operation_name=rpc.operation_name,
                        operation_response_name=self._get_external_struct_name(rpc.operation_name),
                        request_return_type=rpc.request_return_type,
                    )
                    rpc.http_call = """return api.http{operation_name}()""".format(
                        operation_name=rpc.operation_name,
                    )
                    http.request = """api.httpSendRecv("{url}", "", "{method}")""".format(
                        url=url, method=str(operation_id.context.path.fields[0]).upper()
                    )
                    http.method = """http{rpc_method}""".format(rpc_method=rpc.method)
                for ref in self._get_parser("$..responses").find(path_item_object):
                    for status_code, response_object in ref.value.items():
                        response = FluentRpcResponse()
                        response.status_code = status_code
                        response.request_return_type = """New{operation_name}Response""".format(
                            operation_name=rpc.operation_name,
                        )
                        ref = self._get_parser("$..'$ref'").find(response_object)
                        if len(ref) > 0:
                            response.schema = {"$ref": self._resolve_response(ref)[0].value}
                        else:
                            schema = self._get_parser("$..schema").find(response_object)
                            if len(schema) > 0:
                                response.schema = self._get_parser("$..schema").find(response_object)[0].value
                            else:
                                response.schema = {"type": "string"}
                        rpc.responses.append(response)
                        http.responses.append(response)

        self._build_response_interfaces()

        # write the go code
        self._write(
            """type {internal_struct_name} struct {{
                api
                grpcClient {pb_pkg_name}.OpenapiClient
                httpClient httpClient
            }}

            // grpcConnect builds up a grpc connection
            func (api *{internal_struct_name}) grpcConnect() error {{
                if api.grpcClient == nil {{
                    conn, err := grpc.Dial(api.grpc.location, grpc.WithInsecure())
                    if err != nil {{
                        return err
                    }}
                    api.grpcClient = {pb_pkg_name}.NewOpenapiClient(conn)
                }}
                return nil
            }}

            //  NewApi returns a new instance of the top level interface hierarchy
            func NewApi() {interface} {{
                api := {internal_struct_name}{{}}
                return &api
            }}
            
            // httpConnect builds up a http connection
            func (api *{internal_struct_name}) httpConnect() error {{
                if api.httpClient.client == nil {{
                    var verify = !api.http.verify
                    client := httpClient{{
                        client: &http.Client{{
                            Transport: &http.Transport{{
                                TLSClientConfig: &tls.Config{{InsecureSkipVerify: verify}},
                            }},
                        }},
                        ctx: context.Background(),
                    }}
                    api.httpClient = client
                }}
                return nil
            }}
            
            func (api *{internal_struct_name}) httpSendRecv(urlPath string, jsonBody string, method string) (*http.Response, error) {{
                err := api.httpConnect()
                if err != nil {{
                    return nil, err
                }}
                httpClient := api.httpClient
                var bodyReader = bytes.NewReader([]byte(jsonBody))
                queryUrl, err := url.Parse(api.http.location)
                if err != nil {{
                    return nil, err
                }}
                queryUrl, _ = queryUrl.Parse(urlPath)
                req, _ := http.NewRequest(method, queryUrl.String(), bodyReader)
                req.Header.Set("Content-Type", "application/json")
                req = req.WithContext(httpClient.ctx)
                return httpClient.client.Do(req)
            }}            
            """.format(
                internal_struct_name=self._api.internal_struct_name,
                interface=self._api.external_interface_name,
                pb_pkg_name=self._protobuf_package_name,
            )
        )
        methods = []
        for new in self._api.external_new_methods:
            methods.append(new.method_description)
            methods.append(new.method)
        for rpc in self._api.external_rpc_methods:
            methods.append(rpc.description)
            methods.append(rpc.method)
            # descriptions.append("(*{}).{}".format(self._api.external_interface_name, rpc.method_description))
        method_signatures = "\n".join(methods)
        self._write(
            """
            {description}
            type {external_interface_name} interface {{
                Api
                {method_signatures}
            }}
            """.format(
                external_interface_name=self._api.external_interface_name,
                method_signatures=method_signatures,
                description="// {} {}".format(
                    self._api.external_interface_name, 
                    self._get_description(self._openapi["info"], True).lstrip("// "))
            )
        )
        for new in self._api.external_new_methods:
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    newObj := &{struct}{{obj: &{pb_pkg_name}.{interface}{{}}}}
                    newObj.setDefault()
                    return newObj
                }}
                """.format(
                    internal_struct_name=self._api.internal_struct_name,
                    method=new.method,
                    struct=new.struct,
                    pb_pkg_name=self._protobuf_package_name,
                    interface=new.interface,
                )
            )
        for rpc in self._api.external_rpc_methods:
            error_handling = ""
            for response in rpc.responses:
                if response.status_code.startswith("2"):
                    continue
                error_handling += """if resp.GetStatusCode_{status_code}() != nil {{
                        data, _ := yaml.Marshal(resp.GetStatusCode_{status_code}())
                        return nil, fmt.Errorf(string(data))
                    }}
                    """.format(
                    status_code=response.status_code,
                )
            error_handling += 'return nil, fmt.Errorf("response of 200, 400, 500 has not been implemented")'
            if rpc.request_return_type == "[]byte":
                return_value = """if resp.GetStatusCode_200() != nil {
                        return resp.GetStatusCode_200(), nil
                    }"""
            elif rpc.request_return_type == "*string":
                return_value = """if resp.GetStatusCode_200() != "" {
                        status_code_value := resp.GetStatusCode_200()
                        return &status_code_value, nil
                    }"""
            else:
                return_value = """if resp.GetStatusCode_200() != nil {{
                        return &{struct}{{obj: resp.GetStatusCode_200()}}, nil
                    }}""".format(
                    struct=self._get_internal_name(rpc.request_return_type),
                    request_return_type=rpc.request_return_type,
                )
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    {validate}
                    if api.hasHttpTransport() {{
                            {http_call}
                    }}
                    
                    if err := api.grpcConnect(); err != nil {{
                        return nil, err
                    }}
                    request := {request}
                    ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
                    defer cancelFunc()
                    resp, err := api.grpcClient.{operation_name}(ctx, &request)
                    if err != nil {{
                        return nil, err
                    }}
                    {return_value}
                    {error_handling}
                }}
                """.format(
                    internal_struct_name=self._api.internal_struct_name,
                    method=rpc.method,
                    request=rpc.request,
                    operation_name=rpc.operation_name,
                    operation_response_name=self._get_internal_name(rpc.operation_name),
                    error_handling=error_handling,
                    return_value=return_value,
                    http_call=rpc.http_call,
                    validate=getattr(rpc, "validate", ""),
                )
            )

        for http in self._api.internal_http_methods:
            error_handling = ""
            success_method = None
            for response in http.responses:
                if response.status_code.startswith("2"):
                    success_method = response.request_return_type
                else:
                    error_handling += """if resp.StatusCode == {status_code} {{
                            return nil, fmt.Errorf(string(bodyBytes))
                        }}
                        """.format(
                        status_code=response.status_code,
                    )
            error_handling += 'return nil, fmt.Errorf("response not implemented")'
            if http.request_return_type == "[]byte":
                success_handling = """return bodyBytes, nil"""
            elif http.request_return_type == "*string":
                success_handling = """bodyString := string(bodyBytes)
                return &bodyString, nil"""
            else:
                success_handling = """obj := api.{success_method}().StatusCode200()
                    if err := obj.FromJson(string(bodyBytes)); err != nil {{
                        return nil, err
                    }}
                    if err != nil {{
                        return nil, err
                    }}
                    return obj, nil""".format(
                    success_method=success_method,
                    pb_pkg_name=self._protobuf_package_name,
                )
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    resp, err := {request}
                    if err != nil {{
                        return nil, err
                    }}
                    bodyBytes, err := ioutil.ReadAll(resp.Body)
                    defer resp.Body.Close()
                    if err != nil {{
                        return nil, err
                    }}
                    if resp.StatusCode == 200 {{
                        {success_handling}
                    }}
                    {error_handling}
                }}
                """.format(
                    internal_struct_name=self._api.internal_struct_name,
                    method=http.method,
                    request=http.request,
                    success_handling=success_handling,
                    error_handling=error_handling,
                )
            )

    def _build_request_interfaces(self):
        for new in self._api.external_new_methods:
            self._write_interface(new)

    def _write_component_interfaces(self):
        while True:
            components = [component for _, component in self._api.components.items() if component.generated is False]
            if len(components) == 0:
                break
            for component in components:
                self._write_interface(component)

    def _build_response_interfaces(self):
        for rpc in self._api.external_rpc_methods:
            new = FluentNew()
            new.schema_object = {
                "type": "object",
                "properties": {},
            }
            properties = {}
            for response in rpc.responses:
                properties["status_code_{}".format(response.status_code)] = response.schema
            new.schema_object["properties"] = properties
            new.struct = "{operation_name}Response".format(
                operation_name=self._get_internal_name(rpc.operation_name),
            )
            new.interface = "{operation_name}Response".format(
                operation_name=rpc.operation_name,
            )
            new.method = "New{interface}() {interface}".format(
                interface=new.interface,
            )
            new.description = self._get_description(new.schema_object, True)
            new.method_description = """// New{interface} returns a new instance of {interface}.
            """.format(
                interface=new.interface
            ) + "// {} is {}".format(
                new.interface, self._get_description(new.schema_object, True).lstrip("// "))
            new.schema_name = self._get_external_struct_name(new.interface)
            # new.isRpcResponse = True
            self._api.external_new_methods.append(new)

    def _write_interface(self, new):
        if new.schema_name in self._api.components:
            new = self._api.components[new.schema_name]
        else:
            self._api.components[new.schema_name] = new
        if new.generated is True:
            return
        else:
            new.generated = True

        self._build_setters_getters(new)
        internal_items = []
        for field in new.interface_fields:
            if field.adder_method is not None and field.isArray is True:
                internal_items.append("{}s {}".format(
                    field.struct, field.type
                ))
        self._write(
            """type {struct} struct {{
                obj *{pb_pkg_name}.{interface}
                {internal_items}
            }}
            
            func New{interface}() {interface} {{
                obj := {struct}{{obj: &{pb_pkg_name}.{interface}{{}}}}
                obj.setDefault()
                return &obj
            }}

            func (obj *{struct}) Msg() *{pb_pkg_name}.{interface} {{
                return obj.obj
            }}

            func (obj *{struct}) SetMsg(msg *{pb_pkg_name}.{interface}) {interface} {{
                proto.Merge(obj.obj, msg)
                return obj
            }}

            func (obj *{struct}) ToPbText() string {{
                vErr := obj.Validate()
                if vErr != nil {{
                    panic(vErr)
                }}
                return proto.MarshalTextString(obj.Msg())
            }}

            func (obj *{struct}) FromPbText(value string) error {{
                retObj := proto.UnmarshalText(value, obj.Msg())
                if retObj != nil {{
                    return retObj
                }}
                vErr := obj.validateFromText()
                if vErr != nil {{
                    return vErr
                }}
                return retObj
            }}

            func (obj *{struct}) ToYaml() string {{
                vErr := obj.Validate()
                if vErr != nil {{
                    panic(vErr)
                }}
                opts := protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: false,
                }}
                data, err := opts.Marshal(obj.Msg())
                if err != nil {{panic(err)}}
                data, err = yaml.JSONToYAML(data)
                if err != nil {{
                    panic(err)
                }}
                return string(data)
            }}

            func (obj *{struct}) FromYaml(value string) error {{
                data, err := yaml.YAMLToJSON([]byte(value))
                if err != nil {{
                    return err
                }}
                opts := protojson.UnmarshalOptions{{
                    AllowPartial: true,
                    DiscardUnknown: false,
                }}
                uError := opts.Unmarshal([]byte(data), obj.Msg())
                if uError != nil {{
                    return fmt.Errorf("unmarshal error %s", strings.Replace(
                        uError.Error(), "\\u00a0", " ", -1)[7:])
                }}                
                vErr := obj.validateFromText()
                if vErr != nil {{
                    return vErr
                }}
                return nil
            }}

            func (obj *{struct}) ToJson() string {{
                vErr := obj.Validate()
                if vErr != nil {{
                    panic(vErr)
                }}
                opts := protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: false,
                    Indent:          "  ",
                }}
                data, err := opts.Marshal(obj.Msg())
                if err != nil {{
                    panic(err)
                }}
                return string(data)
            }}

            func (obj *{struct}) FromJson(value string) error {{
                opts := protojson.UnmarshalOptions{{
                    AllowPartial: true,
                    DiscardUnknown: false,
                }}
                uError := opts.Unmarshal([]byte(value), obj.Msg())
                if uError != nil {{
                    return fmt.Errorf("unmarshal error %s", strings.Replace(
                        uError.Error(), "\\u00a0", " ", -1)[7:])
                }}
                err := obj.validateFromText()
                if err != nil {{
                    return err
                }}
                return nil
            }}

            func (obj *{struct}) validateFromText() error {{
                obj.validateObj(true)
                return validationResult()
            }}

            func (obj *{struct}) Validate() error {{
                obj.validateObj(false)
                return validationResult()
            }}
        """.format(
                struct=new.struct,
                pb_pkg_name=self._protobuf_package_name,
                interface=new.interface,
                internal_items="" if len(internal_items) == 0 else "\n".join(internal_items)
            )
        )

        interfaces = [
            "// ToPbText marshals {interface} to protobuf text",
            "ToPbText() string",
            "// ToYaml marshals {interface} to YAML text",
            "ToYaml() string",
            "// ToJson marshals {interface} to JSON text",
            "ToJson() string",
            "// FromPbText unmarshals {interface} from protobuf text",
            "FromPbText(value string) error",
            "// FromYaml unmarshals {interface} from YAML text",
            "FromYaml(value string) error",
            "// FromJson unmarshals {interface} from JSON text",
            "FromJson(value string) error",
            "// Validate validates {interface}",
            "Validate() error",
            "validateFromText() error",
            "validateObj(set_default bool)",
            "setDefault()",
        ]
        for field in new.interface_fields:
            interfaces.append("// {}".format(field.getter_method_description))
            interfaces.append(field.getter_method)
            if field.setter_method is not None:
                interfaces.append("// {}".format(field.setter_method_description))
                interfaces.append(field.setter_method)
            if field.has_method is not None:
                interfaces.append("// {}".format(field.has_method_description))
                interfaces.append(field.has_method)
        interface_signatures = "\n".join(interfaces)
        self._write(
            """
            {description}
            type {interface} interface {{
                Msg() *{pb_pkg_name}.{interface}
                SetMsg(*{pb_pkg_name}.{interface}) {interface}
                {interface_signatures}
            }}
        """.format(
                interface=new.interface,
                pb_pkg_name=self._protobuf_package_name,
                interface_signatures=interface_signatures.format(interface=new.interface),
                description="" if new.description is None else "// {} is {}".format(
                    new.interface, new.description.strip("// ")),
            )
        )
        for field in new.interface_fields:
            self._write_field_getter(new, field)
            self._write_field_has(new, field)
            self._write_field_setter(new, field)
            self._write_field_adder(new, field)
        self._write_validate_method(new)
        self._write_default_method(new)

    def _write_field_getter(self, new, field):
        if field.getter_method is None:
            return
        if field.isArray and field.isEnum is False:
            if field.struct:
                block = ""
                if field.setChoiceValue is not None:
                    block = "obj.SetChoice({interface}Choice.{enum})".format(
                        interface=new.interface, enum=field.setChoiceValue
                    )
                body = """if obj.obj.{name} == nil {{
                        {block}
                        obj.obj.{name} = []*{pb_pkg_name}.{pb_struct}{{}}
                    }}
                    return &{parent}{interface}Iter{{obj: obj}}""".format(
                    name=field.name,
                    pb_pkg_name=self._protobuf_package_name,
                    pb_struct=field.external_struct,
                    interface=field.external_struct,
                    parent=new.struct,
                    block=block
                )
            else:
                block = "obj.obj.{name} = make({type}, 0)".format(
                    name=field.name, type=field.type
                )
                if field.setChoiceValue is not None:
                    if field.default is not None and field.default != "":
                        if field.type == "[]string":
                            value = '"{}"'.format('", "'.join(field.default))
                        else:
                            value = ",".join([str(i) for i in field.default])
                        block = """obj.Set{name}({type}{{{default}}})""".format(
                            name=field.name,
                            type=field.type,
                            default=value
                        )
                    else:
                        block = """
                            obj.SetChoice({interface}Choice.{enum})
                            obj.obj.{name} = make({type}, 0)
                        """.format(
                            interface=new.interface,
                            enum=field.setChoiceValue,
                            name=field.name,
                            type=field.type,
                        )
                body = """if obj.obj.{name} == nil {{
                        {block}
                    }}
                    return obj.obj.{name}""".format(
                    name=field.name,
                    block=block
                )
        elif field.struct is not None:
            # at this time proto generation ignores the optional keyword
            # if the type is an object
            set_enum_choice = None
            if field.setChoiceValue is not None:
                set_enum_choice = """obj.SetChoice({interface}Choice.{enum})""".format(
                    interface=new.interface,
                    enum=field.setChoiceValue,
                )
            body = """if obj.obj.{name} == nil {{
                    {set_enum_choice}
                    obj.obj.{name} = New{pb_struct}().Msg()
                }}
                return &{struct}{{obj: obj.obj.{name}}}""".format(
                name=field.name,
                pb_struct=field.external_struct,
                struct=field.struct,
                set_enum_choice=set_enum_choice if set_enum_choice else ""
            )
        elif field.isEnum:
            enum_types = []
            for enum in field.enums:
                enum_types.append(
                    "{enumupper} {interface}{fieldname}Enum".format(
                        enumupper=enum.upper(),
                        interface=new.interface,
                        fieldname=field.name,
                    )
                )
            enum_values = []
            for enum in field.enums:
                enum_values.append(
                    '{enumupper}: {interface}{fieldname}Enum("{enum}")'.format(
                        enumupper=enum.upper(),
                        interface=new.interface,
                        fieldname=field.name,
                        enum=enum,
                    )
                )
            self._write(
                """type {interface}{fieldname}Enum string
                //  Enum of {fieldname} on {interface}
                var {interface}{fieldname} = struct {{
                    {enum_types}
                }} {{
                    {enum_values},
                }}
                """.format(
                    interface=new.interface,
                    fieldname=field.name,
                    enum_types="\n".join(enum_types),
                    enum_values=",\n".join(enum_values),
                )
            )
            if field.isArray:
                self._write(
                    """func (obj *{struct}) {fieldname}() []{interface}{fieldname}Enum {{
                        items := []{interface}{fieldname}Enum{{}}
                        for _, item := range obj.obj.{fieldname} {{
                            items = append(items, {interface}{fieldname}Enum(item.String()))
                        }}
                    return items
                }}
                """.format(
                        struct=new.struct,
                        interface=new.interface,
                        fieldname=field.name,
                    )
                )
            else:
                self._write(
                    """func (obj *{struct}) {fieldname}() {interface}{fieldname}Enum {{
                    return {interface}{fieldname}Enum(obj.obj.{fieldname}.Enum().String())
                }}
                """.format(
                        struct=new.struct,
                        interface=new.interface,
                        fieldname=field.name,
                    )
                )
            return
        elif field.isPointer:
            set_enum_choice = None
            default = None
            if field.setChoiceValue is not None:
                if field.default is not None:
                    if field.type == "string":
                        value = '"{}"'.format(field.default)
                    elif field.type in ["int32", "int64", "float32", "float64"]:
                        value = "{fieldtype}({value})".format(fieldtype=field.type, value=field.default)
                    else:
                        value = field.default
                    default = """value := {value}
                            obj.obj.{fieldname} = &value""".format(value=value, fieldname=field.name)
                set_enum_choice = """
                    if obj.obj.{fieldname} == nil {{
                        obj.SetChoice({interface}Choice.{enum})
                        {default}
                    }}
                """.format(
                    fieldname=field.name,
                    default=default if default else "",
                    interface=new.interface, enum=field.setChoiceValue
                )
            body = """{set_enum_choice}
                return *obj.obj.{fieldname}
                """.format(
                fieldname=field.name,
                set_enum_choice=set_enum_choice if set_enum_choice is not None else ""
            )
        else:
            set_enum_choice = None
            default = None
            if field.setChoiceValue is not None:
                if field.default is not None:
                    default = "obj.obj.{fieldname} = {value}".format(
                        value='"{}"'.format(field.default) if "string" in field.type else field.default,
                        fieldname=field.name
                    )
                set_enum_choice = """
                    if obj.obj.{fieldname} == nil {{
                        obj.SetChoice({interface}Choice.{enum})
                        {default}
                    }}
                """.format(
                    fieldname=field.name,
                    interface=new.interface, enum=field.setChoiceValue,
                    default=default if default else ""
                )
            body = """{set_enum_choice}\n return obj.obj.{fieldname}""".format(
                fieldname=field.name,
                set_enum_choice=set_enum_choice if set_enum_choice is not None else ""
            )
        self._write(
            """
            // {fieldname} returns a {fieldtype}\n{description}
            func (obj *{struct}) {getter_method} {{
                {body}
            }}
            """.format(
                fieldname=self._get_external_struct_name(field.name),
                struct=new.struct,
                getter_method=field.getter_method,
                body=body,
                description=field.description,
                fieldtype=field.type,
            )
        )

    def _write_field_setter(self, new, field):
        if field.setter_method is None:
            return

        if field.isArray and field.isEnum:
            body = """items := []{pb_pkg_name}.{interface}_{fieldname}_Enum{{}}
                for _, item:= range value {{
                    intValue := {pb_pkg_name}.{interface}_{fieldname}_Enum_value[string(item)]
                    items = append(items, {pb_pkg_name}.{interface}_{fieldname}_Enum(intValue))
                }}
                obj.obj.{fieldname} = items""".format(
                interface=new.interface,
                struct=new.struct,
                fieldname=field.name,
                pb_pkg_name=self._protobuf_package_name,
            )
        elif field.isArray:
            body = """if obj.obj.{fieldname} == nil {{
                    obj.obj.{fieldname} = make({fieldtype}, 0)
                }}
                obj.obj.{fieldname} = value
                """.format(
                fieldname=field.name,
                fieldtype=field.type,
            )
        elif field.isEnum:
            if field.isPointer:
                body = """enumValue := {pb_pkg_name}.{interface}_{fieldname}_Enum(intValue)
                obj.obj.{fieldname} = &enumValue""".format(
                    pb_pkg_name=self._protobuf_package_name,
                    interface=new.interface,
                    fieldname=field.name,
                )
            else:
                body = """obj.obj.{fieldname} = {pb_pkg_name}.{interface}_{fieldname}_Enum(intValue)""".format(
                    pb_pkg_name=self._protobuf_package_name,
                    interface=new.interface,
                    fieldname=field.name,
                )
            enum_set = [self._get_external_field_name(name) for name in field.enums]
            enum_body = []
            for enum_field in new.interface_fields:
                if enum_field.name not in enum_set:
                    continue
                if enum_field.isEnum:
                    enum_body.append(
                        "obj.obj.{name} = {pb_pkg_name}.{interface}_{name}_unspecified.Enum()".format(
                            name=enum_field.name, pb_pkg_name=self._protobuf_package_name,
                            interface=new.interface, fieldname=enum_field.name
                        )
                    )
                    continue
                enum_body.append("obj.obj.{name} = nil".format(name=enum_field.name))
            self._write(
                """func (obj* {struct}) Set{fieldname}(value {interface}{fieldname}Enum) {interface} {{
                intValue, ok := {pb_pkg_name}.{interface}_{fieldname}_Enum_value[string(value)]
                if !ok {{
                    validation = append(validation, fmt.Sprintf(
                        "%s is not a valid choice on {interface}{fieldname}Enum", string(value)))
                    return obj
                }}
                {body}
                {enum_set}
                return obj
            }}
            """.format(
                    pb_pkg_name=self._protobuf_package_name,
                    interface=new.interface,
                    struct=new.struct,
                    fieldname=field.name,
                    body=body,
                    enum_set="\n".join(enum_body) if field.name == "Choice" else "",
                )
            )
            return
        elif field.struct is not None:
            body = """obj.{fieldname}().SetMsg(value.Msg())""".format(
                fieldname=self._get_external_struct_name(field.name),
            )
        elif field.isPointer:
            body = """obj.obj.{fieldname} = &value""".format(fieldname=field.name)
        else:
            body = """obj.obj.{fieldname} = value""".format(fieldname=field.name)
        set_choice = ""
        if field.setChoiceValue is not None:
            set_choice = """obj.SetChoice({interface}Choice.{enum})""".format(
                interface=new.interface,
                enum=field.setChoiceValue,
            )
        self._write(
            """
            // Set{fieldname} sets the {fieldtype} value in the {fieldstruct} object\n{description}
            func (obj *{newstruct}) {setter_method} {{
                {set_choice}
                {body}
                return obj
            }}
            """.format(
                fieldname=self._get_external_struct_name(field.name),
                newstruct=new.struct,
                setter_method=field.setter_method,
                body=body,
                description=field.description,
                fieldtype=field.type,
                fieldstruct=new.interface,
                set_choice=set_choice,
            )
        )

    def _write_field_adder(self, new, field):
        if field.adder_method is None:
            return
        interface_name = new.interface + field.external_struct + "Iter"
        if interface_name in self._api.components:
            return
        new_iter = FluentNew()
        new_iter.schema_name = interface_name
        new_iter.interface = interface_name
        new_iter.internal_struct = interface_name[0].lower() + interface_name[1:]
        new_iter.generated = True
        self._api.components[interface_name] = new_iter
        self._write(
            """
            type {internal_struct} struct {{
                obj *{parent_internal_struct}
            }}

            type {interface} interface {{
                Items() {field_type}
                Add() {field_external_struct}
                Append(items ...{field_external_struct}) {interface}
                Set(index int, newObj {field_external_struct}) {interface}
                Clear() {interface}
            }}

            func (obj *{internal_struct}) Items() {field_type} {{
                return obj.obj.{internal_items_name}
            }}

            func (obj *{internal_struct}) Add() {field_external_struct} {{
                newObj := &{pb_pkg_name}.{field_external_struct}{{}}
                obj.obj.obj.{field_name} = append(obj.obj.obj.{field_name}, newObj)
                newLibObj := &{field_internal_struct}{{obj: newObj}}
                newLibObj.setDefault()
                obj.obj.{internal_items_name} = append(obj.obj.{internal_items_name}, newLibObj)
                return newLibObj
            }}

            func (obj *{internal_struct}) Append(items ...{field_external_struct}) {interface} {{
                for _, item := range items {{
                    newObj := item.Msg()
                    obj.obj.obj.{field_name} = append(obj.obj.obj.{field_name}, newObj)
                    obj.obj.{internal_items_name} = append(obj.obj.{internal_items_name}, item)
                }}
                return obj
            }}

            func (obj *{internal_struct}) Set(index int, newObj {field_external_struct}) {interface} {{
                obj.obj.obj.{field_name}[index] = newObj.Msg()
                obj.obj.{internal_items_name}[index] = newObj
                return obj
            }}
            func (obj *{internal_struct}) Clear()  {interface} {{
                if obj.obj.obj.{field_name} != nil {{
                    obj.obj.obj.{field_name} = nil
                    obj.obj.{internal_items_name} = {field_type}{{}}
                }}
                return obj
            }}
            """.format(
                internal_struct=new_iter.internal_struct,
                interface=new_iter.interface,
                field_internal_struct=field.struct,
                parent_internal_struct=new.struct,
                field_external_struct=field.external_struct,
                field_name=field.name,
                pb_pkg_name=self._protobuf_package_name,
                field_type=field.type,
                internal_items_name="{}s".format(field.struct)
            )
        )

    def _write_field_has(self, new, field):
        if field.has_method is None:
            return
        self._write(
            """
            // {fieldname} returns a {fieldtype}\n{description}
            func (obj *{struct}) Has{fieldname}() bool {{
                return obj.obj.{internal_field_name} != nil
            }}
            """.format(
                fieldname=self._get_external_struct_name(field.name),
                struct=new.struct,
                getter_method=field.getter_method,
                description=field.description,
                fieldtype=field.type,
                internal_field_name=field.name,
            )
        )

    def _build_setters_getters(self, fluent_new):
        """Add new FluentField objects for each interface field"""
        if "properties" not in fluent_new.schema_object:
            schema = self._get_parser("$..schema").find(fluent_new.schema_object)
            if len(schema) > 0:
                schema = schema[0].value
                schema_name = self._get_schema_object_name_from_ref(schema["$ref"])
                fluent_new.schema_object = {
                    "properties": {
                        self._get_external_struct_name(schema_name): schema,
                    },
                }
            else:
                return
        choice_enums = self._get_parser("$..choice..enum").find(fluent_new.schema_object["properties"])
        for property_name, property_schema in fluent_new.schema_object["properties"].items():
            field = FluentField()
            field.schema = property_schema
            field.description = self._get_description(property_schema)
            field.name = self._get_external_field_name(property_name)
            field.type = self._get_struct_field_type(property_schema, field)
            if len(choice_enums) == 1 and property_name in choice_enums[0].value:
                field.setChoiceValue = property_name.upper()
            else:
                field.setChoiceValue = None
            field.isEnum = len(self._get_parser("$..enum").find(property_schema)) > 0
            field.hasminmax = "minimum" in property_schema or "maximum" in property_schema
            field.hasminmaxlength = "minLength" in property_schema or "maxLength" in property_schema
            field.isArray = "type" in property_schema and property_schema["type"] == "array"
            if field.isEnum:
                field.enums = self._get_parser("$..enum").find(property_schema)[0].value
                if "unspecified" in field.enums:
                    field.enums.remove("unspecified")
            if field.hasminmax:
                field.min = None if "minimum" not in property_schema else property_schema["minimum"]
                field.max = None if "maximum" not in property_schema else property_schema["maximum"]
                if (
                    (field.min is not None and field.min > 2147483647)
                    or (field.max is not None and field.max > 2147483647)
                    and "int" in field.type
                ):
                    field.type = field.type.replace("32", "64")
            if field.hasminmaxlength:
                field.min_length = None if "minLength" not in property_schema else property_schema["minLength"]
                field.max_length = None if "maxLength" not in property_schema else property_schema["maxLength"]
            if fluent_new.isRpcResponse:
                if field.type == "[]byte":
                    field.name = "Bytes"
                elif "$ref" in property_schema:
                    schema_name = self._get_schema_object_name_from_ref(property_schema["$ref"])
                    field.name = self._get_external_struct_name(schema_name)
            field.isOptional = fluent_new.isOptional(property_name)
            field.isPointer = False if field.type.startswith("[") else field.isOptional
            if field.isArray and field.isEnum:
                field.getter_method = "{fieldname}() []{interface}{fieldname}Enum".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
                field.getter_method_description = "{fieldname} returns []{interface}{fieldname}Enum, set in {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
            elif field.isEnum:
                field.getter_method = "{fieldname}() {interface}{fieldname}Enum".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
                field.getter_method_description = "{fieldname} returns {interface}{fieldname}Enum, set in {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
            else:
                field.getter_method = "{name}() {ftype}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                )
                field.getter_method_description = "{name} returns {ftype}, set in {interface}.".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                    interface=fluent_new.interface
                )
                if field.type in self._api.components:
                    field.getter_method_description = field.getter_method_description + \
                        """\n// {ftype} is {desc}""".format(
                            ftype=field.type,
                            desc=self._api.components[field.type].description.lstrip("// ")
                        )

            if "$ref" in property_schema:
                schema_name = self._get_schema_object_name_from_ref(property_schema["$ref"])
                field.struct = self._get_internal_name(schema_name)
                field.external_struct = self._get_external_struct_name(schema_name)
                field.setter_method = "Set{fieldname}(value {fieldstruct}) {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    fieldstruct=self._get_external_struct_name(field.struct),
                    interface=fluent_new.interface,
                )
                field.setter_method_description = "Set{fieldname} assigns {fieldstruct} provided by user to {interface}.".format(
                    fieldname=self._get_external_struct_name(field.name),
                    fieldstruct=self._get_external_struct_name(field.struct),
                    interface=fluent_new.interface,
                )
                fieldstruct = self._get_external_struct_name(field.struct)
                if fieldstruct in self._api.components:
                    field.setter_method_description = field.setter_method_description + \
                        """\n // {fieldstruct} is {desc}""".format(
                            fieldstruct=fieldstruct,
                            desc=self._api.components[fieldstruct].description.lstrip("// ")
                        )
            if field.isOptional and field.isPointer or "StatusCode" in field.name:
                field.has_method = """Has{fieldname}() bool""".format(
                    fieldname=self._get_external_struct_name(field.name),
                )
                field.has_method_description = """Has{fieldname} checks if {fieldname} has been set in {interface}""".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface
                )
            if field.isArray and field.isEnum:
                field.setter_method = "Set{fieldname}(value []{interface}{fieldname}Enum) {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
                field.setter_method_description = "Set{fieldname} assigns []{interface}{fieldname}Enum provided by user to {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )

            elif field.isEnum:
                field.setter_method = "Set{fieldname}(value {interface}{fieldname}Enum) {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
                field.setter_method_description = "Set{fieldname} assigns {interface}{fieldname}Enum provided by user to {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
            elif field.type in self._oapi_go_types.values():
                field.setter_method = "Set{name}(value {ftype}) {interface}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                    interface=fluent_new.interface,
                )
                field.setter_method_description = "Set{name} assigns {ftype} provided by user to {interface}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                    interface=fluent_new.interface,
                )
            elif field.isArray:
                field.isPointer = False
                if "$ref" in property_schema["items"]:
                    schema_name = self._get_schema_object_name_from_ref(property_schema["items"]["$ref"])
                    field.isArray = True
                    field.struct = self._get_internal_name(schema_name)
                    field.external_struct = self._get_external_struct_name(schema_name)
                    field.adder_method = "Add() {parent}{external_struct}Iter".format(
                        parent=fluent_new.interface,
                        external_struct=field.external_struct,
                    )
                    field.isOptional = False
                    field.getter_method = "{name}() {parent}{external_struct}Iter".format(
                        name=self._get_external_struct_name(field.name),
                        parent=fluent_new.interface,
                        external_struct=field.external_struct,
                    )
                    field.getter_method_description = "{name} returns {parent}{external_struct}Iter, set in {parent}".format(
                        name=self._get_external_struct_name(field.name),
                        parent=fluent_new.interface,
                        external_struct=field.external_struct,
                    )
                else:
                    field.setter_method = "Set{name}(value {ftype}) {interface}".format(
                        name=self._get_external_struct_name(field.name),
                        ftype=field.type,
                        interface=fluent_new.interface,
                    )
                    field.setter_method_description = "Set{name} assigns {ftype} provided by user to {interface}".format(
                        name=self._get_external_struct_name(field.name),
                        ftype=field.type,
                        interface=fluent_new.interface,
                    )
            default = property_schema.get("default")
            if default is not None:
                type = field.type
                if field.isArray:
                    type = field.type.lstrip("[]")
                if type in self._oapi_go_types.values():
                    if field.type == "number":
                        default = float(default)
                    if field.type == "bool":
                        default = str(default).lower()
                    field.default = default
                else:
                    print("Warning: Default should not accept for this property ", property_name)

            fluent_new.interface_fields.append(field)

    def _validate_types(self, new, field):
        body = ""
        if field.isPointer or "[]" in field.type:
            value = "nil"
        elif field.type == "string":
            value = '''""'''
        else:
            value = 0
        if field.isOptional is False:
            body = """
            // {name} is required
            if obj.obj.{name}{enum} == {value} {{
                validation = append(validation, "{name} is required field on interface {interface}")
            }} """.format(
                name=field.name, interface=new.interface,
                value=0 if field.isEnum and field.isArray is False else value,
                enum=".Number()" if field.isEnum and field.isArray is False else ""
            )

        inner_body = ""
        if field.hasminmax and ("int" in field.type or "float" in field.type):
            line = []
            if field.min is None and "int" in field.type:
                field.min = -(2**31 if "32" in field.type else 2**63)
            if field.max is None and "int" in field.type:
                field.max = (2**31 if "32" in field.type else 2**63) - 1
            if field.min is not None:
                line.append("{pointer}{value} < {min}")
            if field.max is not None:
                line.append("{pointer}{value} > {max}")
            inner_body += (
                "if "
                + " || ".join(line)
                + """ {{
                    validation = append(
                        validation,
                        fmt.Sprintf("{min} <= {interface}.{name} <= {max} but Got {form}", {pointer}{value}))
                    }}
                """
            ).format(
                name=field.name,
                interface=new.interface,
                max="max({})".format(field.type.lstrip("[]")) if field.max is None else field.max,
                pointer="*" if field.isPointer else "",
                min="min({})".format(field.type.lstrip("[]")) if field.min is None else field.min,
                value="item" if field.isArray else "obj.obj.{name}".format(name=field.name),
                form="%d" if "int" in field.type else "%f"
            )
            if field.isArray:
                inner_body = """
                    for _, item := range obj.obj.{name} {{
                        {body}
                    }}
                """.format(
                        body=inner_body, name=field.name
                    )
        elif "string" in field.type:
            if field.hasminmaxlength and "string" in field.type:
                line = []
                if field.min_length is not None:
                    line.append("len({pointer}{value}) < {min_length}")
                if field.max_length is not None:
                    line.append("len({pointer}{value}) > {max_length}")
                inner_body = (
                    "if "
                    + " || ".join(line)
                    + """ {{
                    validation = append(
                        validation,
                        fmt.Sprintf(
                            "{min_length} <= length of {interface}.{name} <= {max_length} but Got %d",
                            len({pointer}{value})))
                }}
                """
                ).format(
                    name=field.name,
                    interface=new.interface,
                    max_length="any" if field.max_length is None else field.max_length,
                    pointer="*" if field.isPointer else "",
                    min_length=field.min_length if field.min_length is None else field.min_length,
                    value="item" if field.isArray else "obj.obj.{name}".format(name=field.name),
                )
                if field.isArray:
                    inner_body = """
                        for _, item := range obj.obj.{name} {{
                            {body}
                        }}
                    """.format(
                        body=inner_body, name=field.name
                    )
            elif field.itemformat in ["mac", "ipv4", "ipv6", "hex"] or field.format in ["mac", "ipv4", "ipv6", "hex"]:
                if field.format is None:
                    field.format = field.itemformat
                inner_body = """
                    err := validate{format}(obj.{name}())
                    if err != nil {{
                        validation = append(validation, fmt.Sprintf("%s %s", err.Error(), "on {interface}.{name}"))
                    }}
                """.format(
                        name=field.name, interface=new.interface,
                        format=field.format.capitalize() if field.isArray is False
                        else field.format.capitalize() + "Slice",
                    )
        # Enum will be handled via wrapper lib
        if inner_body == "":
            return body
        body += """
        if obj.obj.{name} != {value} {{
            {body}
        }}
        """.format(name=field.name, value=value, body=inner_body)
        return body

    def _validate_struct(self, new, field):
        body = ""
        if field.isOptional is False and field.isArray is False:
            body = """
                // {name} is required
                if obj.obj.{name} == nil {{
                    validation = append(validation, "{name} is required field on interface {interface}")
                }}
            """.format(name=field.name, interface=new.interface)

        inner_body = "obj.{external_name}().validateObj(set_default)".format(
            external_name=self._get_external_struct_name(field.name)
        )
        if field.isArray:
            inner_body = """if set_default {{
                    obj.{internal_items_name} = {field_type}{{}}
                    for _, item := range obj.obj.{name} {{
                        obj.{internal_items_name} = append(obj.{internal_items_name}, &{field_internal_struct}{{obj: item}})
                    }}
                }}
                for _, item := range obj.{name}().Items() {{
                    item.validateObj(set_default)
                }}
            """.format(
                name=field.name,
                field_type=field.type,
                internal_items_name="{}s".format(field.struct),
                field_internal_struct=field.struct
            )
        body += """
            if obj.obj.{name} != nil {{
                {body}
            }}
        """.format(name=field.name, body=inner_body)
        return body

    def _write_validate_method(self, new):
        statements = []
        def p():
            if valid == 0:
                print(
                    "{field} of type {ftype} and {req} is not set for validation on interface {interface}".format(
                        field=field.name,
                        interface=new.interface,
                        ftype=field.type,
                        req="Optional" if field.isOptional else "required",
                    )
                )
        for field in new.interface_fields:
            valid = 0
            if field.type.lstrip("[]") in self._oapi_go_types.values():
                block = self._validate_types(new, field)
                if block is None or block.strip() == "":
                    p()
                    continue
                valid += 1
                statements.append(block)
            elif field.struct:
                block = self._validate_struct(new, field)
                if block is None or block.strip() == "":
                    p()
                    continue
                valid += 1
                statements.append(block)
            p()
        body = "\n".join(statements)
        self._write(
            """func (obj *{struct}) validateObj(set_default bool) {{
                if set_default {{
                    obj.setDefault()
                }}
                {body}
            }}
            """.format(
                struct=new.struct, body=body
            )
        )

    def _write_default_method(self, new):
        body = ""
        interface_fields = new.interface_fields
        hasChoiceConfig = []
        choice_enums = []
        for index, field in enumerate(interface_fields):
            if field.default is None:
                continue
            if field.name == "Choice":
                choice_enums = [self._get_external_struct_name(e) for e in field.enums if e != field.default]
                hasChoiceConfig = ["Choice", self._get_external_struct_name(field.default)]
                interface_fields.insert(0, interface_fields.pop(index))
                break

        choice_body = None
        enum_fields = []
        for field in interface_fields:
            # if hasChoiceConfig != [] and field.name not in hasChoiceConfig:
            #     continue
            if field.default is None or field.isOptional is False or field.name in choice_enums:
                if field.name not in hasChoiceConfig:
                    continue
            if field.struct is not None:
                if field.name in hasChoiceConfig:
                    enum_fields.append("obj.{external_name}()".format(external_name=self._get_external_struct_name(field.name)))
                else:
                    body += """if obj.obj.{name} == nil {{
                        obj.{external_name}()
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(field.name),
                        # enum_check="&& {}".format(enum_check) if enum_check is not None and field.name in choice_enums
                        # else ""
                    )
            elif field.isArray:
                if "string" in field.type:
                    values = '"{0}"'.format('", "'.join(field.default)) if field.default != [] else ""
                else:
                    values = str(field.default)[1:-1]
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({type}{{{values}}})".format(
                            external_name=self._get_external_struct_name(field.name), type=field.type, values=values
                        )
                    )
                else:
                    body += """if obj.obj.{name} == nil {{
                        obj.Set{external_name}({type}{{{values}}})
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(field.name),
                        type=field.type,
                        values=values,
                    )
            elif field.isEnum:
                enum_value = """{struct}{name}.{value}""".format(
                    struct=self._get_external_struct_name(new.struct), name=field.name, value=field.default.upper()
                )
                if field.isPointer:
                    cnd_check = """obj.obj.{name} == nil""".format(name=field.name)
                else:
                    cnd_check = """obj.obj.{name}.Number() == 0""".format(name=field.name)
                body1 = """if {cnd_check} {{
                    obj.Set{external_name}({enum_value})
                    <choice_fields>
                }}
                """.format(
                    cnd_check=cnd_check,
                    external_name=self._get_external_struct_name(field.name),
                    enum_value=enum_value,
                )
                if field.name in hasChoiceConfig:
                    if choice_body is not None:
                        body1 = body1.replace(" == nil", ".Number() == 0")
                    choice_body = body1 if choice_body is None else choice_body.replace("<choice_fields>", body1)
                else:
                    body = body + body1.replace("<choice_fields>", "")
            elif field.isPointer:
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({value})".format(
                            external_name=self._get_external_struct_name(field.name),
                            value='"{0}"'.format(field.default) if field.type == "string" else field.default,
                        )
                    )
                else:
                    body += """if obj.obj.{name} == nil {{
                        obj.Set{external_name}({value})
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(field.name),
                        value='"{0}"'.format(field.default) if field.type == "string" else field.default,
                    )
            else:
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({value})".format(
                            external_name=self._get_external_struct_name(field.name),
                            value='"{0}"'.format(field.default) if field.type == "string" else field.default,
                        )
                    )
                else:
                    body += """if obj.obj.{name} == {check_value} {{
                        obj.Set{external_name}({value})
                    }}
                    """.format(
                        name=field.name,
                        check_value='""' if field.type == "string" else "0",
                        external_name=self._get_external_struct_name(field.name),
                        value='"{0}"'.format(field.default) if field.type == "string" else field.default,
                    )
        if choice_body is not None:
            body = choice_body.replace("<choice_fields>", "\n".join(enum_fields) if enum_fields != [] else "") + body

        self._write(
            """func (obj *{struct}) setDefault() {{
                {body}
            }}""".format(
                struct=new.struct, body=body
            )
        )

    def _get_schema_object_name_from_ref(self, ref):
        final_piece = ref.split("/")[-1]
        return final_piece.replace(".", "")

    def _get_schema_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf

    def _get_struct_field_type(self, property_schema, fluent_field=None):
        """Convert openapi type, format, items, $ref keywords to a go type"""
        go_type = ""
        if "type" in property_schema:
            oapi_type = property_schema["type"]
            if oapi_type.lower() in self._oapi_go_types:
                go_type = "{oapi_go_types}".format(oapi_go_types=self._oapi_go_types[oapi_type.lower()])
            if oapi_type == "array":
                go_type += "[]" + self._get_struct_field_type(property_schema["items"], fluent_field).replace("*", "")
                if "format" in property_schema["items"]:
                    fluent_field.itemformat = property_schema["items"]["format"]
            if "format" in property_schema:
                format_type = (oapi_type + property_schema["format"]).lower()
                if format_type.lower() in self._oapi_go_types:
                    go_type = "{oapi_go_type}".format(oapi_go_type=self._oapi_go_types[format_type.lower()])
                elif property_schema["format"].lower() in self._oapi_go_types:
                    go_type = "{oapi_go_type}".format(oapi_go_type=self._oapi_go_types[property_schema["format"].lower()])
                else:
                    fluent_field.format = property_schema["format"].lower()
        elif "$ref" in property_schema:
            ref = property_schema["$ref"]
            schema_object_name = self._get_schema_object_name_from_ref(ref)
            schema_object = self._get_schema_object_from_ref(ref)
            new = None
            if schema_object_name in self._api.components:
                new = self._api.components[schema_object_name]
            else:
                new = FluentNew()
                new.schema_object = schema_object
                new.schema_name = schema_object_name
                new.struct = self._get_internal_name(schema_object_name)
                new.interface = self._get_external_struct_name(schema_object_name)
                new.description = self._get_description(schema_object, True)
                self._api.components[new.schema_name] = new
            go_type = new.interface
        else:
            raise Exception(
                "No type or $ref keyword present in property schema: {property_schema}".format(property_schema=property_schema)
            )
        return go_type

    def _get_description(self, openapi_object, noCap=False):
        description = "// description is TBD"
        if "description" in openapi_object:
            description = ""
            for ind, line in enumerate(openapi_object["description"].split("\n")):
                if noCap and ind == 0 and line != "":
                    line = line[0].lower() + line[1:]
                description += "// {line}\n".format(line=line.strip())
        return description.strip("\n")

    def _format_go_file(self):
        """Format the generated go code"""
        try:
            process_args = [
                "goimports",
                "-w",
                self._filename,
            ]
            cmd = " ".join(process_args)
            print("Formatting generated go ux file: {}".format(cmd))
            process = subprocess.Popen(cmd, cwd=self._ux_path, shell=True)
            process.wait()
        except Exception as e:
            print("Bypassed formatting of generated go ux file: {}".format(e))

    def _tidy_mod_file(self):
        """Tidy the mod file"""
        try:
            process_args = [
                "go",
                "mod",
                "tidy",
            ]
            os.environ["GO111MODULE"] = "on"
            print("Tidying the generated go mod file: {}".format(" ".join(process_args)))
            process = subprocess.Popen(process_args, cwd=self._ux_path, shell=False, env=os.environ)
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))

    def _generate_interface_info(self):
        interface_dict = {}
        enum_dict = {}
        for interface in self._api.components.keys():
            if interface == "LevelFour":
                continue
            method_and_type = []
            enum_and_type = []
            has_and_type = []
            for field in self._api.components[interface].interface_fields:
                if field.isEnum:
                    if field.default:
                        enum = interface + field.name + 'Enum'
                        type = field.default
                    else:
                        enum = interface + field.name + 'Enum'
                        type = field.enums[0]
                    enum_and_type.append(tuple([enum, type]))
                if field.default:
                    method, type = field.name, field.type
                    method = method.replace("_", "")
                    method_and_type.append(tuple([method, type]))
                if 'Iter' in field.getter_method:
                    method = field.getter_method.split("()")[0]
                    method = method.replace("_", "")
                    type = field.getter_method.split("()")[1].lstrip()
                    method_and_type.append(tuple([method, type]))
                if field.setter_method is not None:
                    if field.format in ["mac", "ipv4", "ipv6", "hex"]:
                        if field.isArray:
                            method = 'Set' + field.name
                            type = (field.format + "_array")
                        else:
                            method, type = 'Set' + field.name, field.format
                    elif field.hasminmax:
                        method = 'Set' + field.name
                        type = "min" + str(field.min) + "max" + str(field.max)
                    elif field.isEnum:
                        method = 'Set' + field.name
                        type = interface + field.name + 'Enum'
                    else:
                        method, type = 'Set' + field.name, field.type
                    if type != "LevelFour":
                        method = method.replace("_", "")
                        method_and_type.append(tuple([method, type]))
                if field.has_method:
                    has, type = field.has_method.split()
                    has_and_type.append(tuple([has, type]))
            if has_and_type:
                method_and_type.extend(has_and_type)
            if method_and_type:
                interface_dict[interface] = method_and_type
            if enum_and_type:
                enum_dict[interface] = enum_and_type
        return interface_dict, enum_dict

    def _write_in_go_test(self, fp, line="", indent=0, newline=True):
        default_indent = "    "
        line = "{}{}{}".format(
            default_indent * indent, line, "\n" if newline else "")
        fp.write(line.encode())

    def _generate_go_unit_test(self):
        interface_dict, enum_dict = self._generate_interface_info()
        self._write_go_unit_test(interface_dict, enum_dict)
        self._write_go_validate_test(interface_dict)
        self._write_go_required_test()
        self._write_go_invalid_format_test()
        self._write_go_invalid_key_test()

    def _write_go_unit_test(self, interface_dict, enum_dict):
        # TODO: Have to add a more dynamic way of picking up the entry interface
        if self._go_sdk_package_name == "openapiart":
            entry_int = "PrefixConfig"
        elif self._go_sdk_package_name == "gosnappi":
            entry_int = "Config"

        filename = os.path.normpath(
            os.path.join(self._ux_path, "generated_test.go"))
        fp = open(filename, "wb")

        self._write_in_go_test(fp,
                               "package {go_sdk_package_name}_test".format(
                                   go_sdk_package_name=self._go_sdk_package_name))
        line = 'import {go_sdk_package_name} "{go_sdk_pkg_dir}"'.format(
            go_sdk_package_name=self._go_sdk_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write_in_go_test(fp, line)
        self._write_in_go_test(fp, 'import "testing"')
        self._write_in_go_test(fp, 'import "github.com/stretchr/testify/assert"')
        initial_fill = """func TestConfigGenerated(t *testing.T) {{
        api := {}.NewApi()
        config := api.New{}()""".format(self._go_sdk_package_name, entry_int)
        self._write_in_go_test(fp, initial_fill)

        lines = self._call_field("config",
                                 interface_dict,
                                 enum_dict,
                                 entry_int,
                                 1,
                                 [])
        if "." in lines[0]:
            first_level = lines[0]
        else:
            first_level = ""
        recur_level = [value for value in lines[1] if "." in value]
        self._write_in_go_test(fp, first_level)
        self._write_in_go_test(fp, "\n".join(recur_level))
        self._write_in_go_test(fp, "\n".join(lines[2]))
        self._write_in_go_test(fp, """config.ToJson()
        config.ToYaml()
        config.ToPbText()
        assert.Nil(t, config.FromJson(config.ToJson()))
        assert.Nil(t, config.FromYaml(config.ToYaml()))
        assert.Nil(t, config.FromPbText(config.ToPbText()))""")
        self._write_in_go_test(fp, "}")

    def _call_field(self, main_interface, interface_dict, enum_dict, key, var, var_lst):
        type_dict = {
            'string': '"abc"',
            'int32': 1,
            'int64': 10000000000000000,
            'float32': 100.11,
            'float64': 1.7976931348623157e+308,
            '[]int32': "[]int32{1, 2, 3}",
            '[]int64': "[]int64{1, 2, 3}",
            '[]string': """[]string{"a","b","c"}""",
            'bool': "true",
            'ipv4': '"1.1.1.1"',
            'ipv6': '"2000::1"',
            'mac': '"00:00:00:00:00:0a"',
            'hex': '"0x12"',
            '[]byte': "[]byte{1, 0, 0, 1, 0, 0, 1, 1}",
            'ipv4_array': '[]string{"1.1.1.1"}',
            'ipv6_array': '[]string{"2000::1"}',
            'mac_array': '[]string{"00:00:00:00:00:0a"}',
            'hex_array': '[]string{"0x12"}',
        }
        first_level = main_interface
        first_level_get = []
        recur_level = []
        for interface_field in interface_dict[key]:
            if 'Iter' in interface_field[1]:
                type = (interface_field[1]
                        .replace('Iter', "")
                        .replace(key, "", 1))
                if var in var_lst:
                    var = max(var_lst) + 1
                var_lst.append(var)
                new_var = ("v{}:={}.{}().Add()".format(str(var),
                                                       main_interface,
                                                       interface_field[0]))
                recur_level.append(new_var)
                updated_interface = "v{}".format(str(var))
                lines = self._call_field(updated_interface,
                                         interface_dict,
                                         enum_dict,
                                         type,
                                         var,
                                         var_lst)
                recur_level.append(lines[0])
                recur_level.extend(lines[1])
                to_json = "{inter}.ToJson()".format(inter=updated_interface)
                to_yaml = "{inter}.ToYaml()".format(inter=updated_interface)
                to_pb = "{inter}.ToPbText()".format(inter=updated_interface)
                from_json = "assert.Nil(t, {inter}.FromJson({inter}.ToJson()))".format(inter=updated_interface)
                from_yaml = "assert.Nil(t, {inter}.FromYaml({inter}.ToYaml()))".format(inter=updated_interface)
                from_pb = "assert.Nil(t, {inter}.FromPbText({inter}.ToPbText()))".format(inter=updated_interface)
                recur_level.extend([to_json, to_yaml, to_pb, from_json, from_yaml, from_pb])
            else:
                # TODO: Skipping Layer1 & SetDvalues for now
                if ('Ieee' in interface_field[1] or
                        'SetChoice' in interface_field[0] or
                        'SetDValues' in interface_field[0]):
                    continue
                if interface_field[1] not in interface_dict.keys():
                    if 'Has' in interface_field[0]:
                        recur_level.append("{}.{}".format(main_interface, interface_field[0]))
                    else:
                        if 'Enum' in interface_field[1]:
                            val = [ele[1]
                                   for ele in enum_dict[key]
                                   if ele[0] == interface_field[1] and ele[1]]
                            if val:
                                val = '"{}"'.format(str(val[0]))
                        else:
                            if 'max' in interface_field[1]:
                                if 'maxNone' in interface_field[1]:
                                    minimum = interface_field[1].split("max")[0]
                                    val = int(minimum.split("min")[1])
                                else:
                                    minimum = (interface_field[1].split("max")[0].split("min")[1])
                                    maximum = interface_field[1].split("max")[1]
                                    val = random.randint(int(minimum), int(maximum))
                            else:
                                val = type_dict[interface_field[1]]
                            if (interface_field[0] == 'SetValues' and "max" in interface_field[1]):
                                max_value = int(interface_field[1].split("max")[1])
                                if max_value > 2147483647:
                                    val = "[]int64" + "{" + str(val) + "}"
                                else:
                                    val = "[]int32" + "{" + str(val) + "}"
                        if val is not None and interface_field[0].startswith("Set"):
                            first_level = "{}.{}({})".format(
                                first_level, interface_field[0], str(val))
                        else:
                            first_level_get.append(
                                main_interface + "." + interface_field[0] + "()")
                else:
                    method = interface_field[0].split("Set", 1)[1]
                    if var in var_lst:
                        var = max(var_lst) + 1
                    updated_interface = "v{}".format(str(var))
                    new_var = "v{}:={}.{}()".format(str(var), main_interface, method)
                    recur_level.append(new_var)
                    var_lst.append(var)
                    lines = self._call_field(updated_interface,
                                             interface_dict,
                                             enum_dict,
                                             interface_field[1],
                                             var,
                                             var_lst)
                    recur_level.append(lines[0])
                    recur_level.extend(lines[1] + lines[2])
                    to_json = "{inter}.ToJson()".format(inter=updated_interface)
                    to_yaml = "{inter}.ToYaml()".format(inter=updated_interface)
                    to_pb = "{inter}.ToPbText()".format(inter=updated_interface)
                    from_json = "assert.Nil(t, {inter}.FromJson({inter}.ToJson()))".format(inter=updated_interface)
                    from_yaml = "assert.Nil(t, {inter}.FromYaml({inter}.ToYaml()))".format(inter=updated_interface)
                    from_pb = "assert.Nil(t, {inter}.FromPbText({inter}.ToPbText()))".format(inter=updated_interface)
                    recur_level.extend([to_json, to_yaml, to_pb, from_json, from_yaml, from_pb])
                    recur_level.append("{}.Set{}({})".format(main_interface, method, updated_interface))
        return first_level, recur_level, first_level_get

    def _write_go_validate_test(self, interface_dict):
        # TODO: Have to add a more dynamic way of picking up the entry interface
        if self._go_sdk_package_name == "openapiart":
            entry_int = "PrefixConfig"
        elif self._go_sdk_package_name == "gosnappi":
            entry_int = "Config"

        filename = os.path.normpath(
            os.path.join(self._ux_path, "generated_validate_test.go"))
        fp = open(filename, "wb")

        self._write_in_go_test(fp,
                               "package {go_sdk_package_name}_test".format(
                                   go_sdk_package_name=self._go_sdk_package_name))
        line = 'import {go_sdk_package_name} "{go_sdk_pkg_dir}"'.format(
            go_sdk_package_name=self._go_sdk_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write_in_go_test(fp, line)
        self._write_in_go_test(fp, 'import "testing"')
        self._write_in_go_test(fp, 'import "github.com/stretchr/testify/assert"')
        self._write_in_go_test(fp, 'import "google.golang.org/protobuf/encoding/protojson"')
        self._write_in_go_test(fp, 'import "github.com/golang/protobuf/proto"')
        initial_fill = """func TestConfigGeneratedValidate(t *testing.T) {{
            opts := protojson.MarshalOptions{{
            UseProtoNames:   true,
            AllowPartial:    true,
            EmitUnpopulated: false,
            Indent:          "  ",
            }}
        api := {}.NewApi()
        config := api.New{}()""".format(self._go_sdk_package_name, entry_int)
        self._write_in_go_test(fp, initial_fill)

        lines = self._call_field_on_inter("config",
                                          interface_dict,
                                          entry_int,
                                          1,
                                          [])
        if "." in lines[0]:
            first_level = lines[0]
        else:
            first_level = ""
        recur_level = [value for value in lines[1] if "." in value]
        self._write_in_go_test(fp, first_level)
        self._write_in_go_test(fp, "\n".join(recur_level))
        self._write_in_go_test(fp, "\n".join(lines[2]))
        self._write_in_go_test(fp, "}")

    def _call_field_on_inter(self, main_interface, interface_dict, key, var, var_lst):
        type_dict = {
            'ipv4': '"1.1.1.1.1"',
            'ipv6': '"2000::1::1"',
            'mac': '"00:00:00:00:00:0a::0b"',
            'hex': '"0x12JKL"',
            'ipv4_array': '[]string{"1.1.1.1.1"}',
            'ipv6_array': '[]string{"2000::1::2"}',
            'mac_array': '[]string{"00:00:00:00:00:0a::b"}',
            'hex_array': '[]string{"0x12HKLJ"}',
        }
        first_level = main_interface
        first_level_get = []
        recur_level = []
        for interface_field in interface_dict[key]:
            if 'Iter' in interface_field[1]:
                type = (interface_field[1]
                        .replace('Iter', "")
                        .replace(key, "", 1))
                if var in var_lst:
                    var = max(var_lst) + 1
                var_lst.append(var)
                new_var = ("v{}:={}.{}().Add()".format(str(var),
                                                       main_interface,
                                                       interface_field[0]))
                recur_level.append(new_var)
                updated_interface = "v{}".format(str(var))
                lines = self._call_field_on_inter(updated_interface,
                                                  interface_dict,
                                                  type,
                                                  var,
                                                  var_lst)
                recur_level.append(lines[0])
                recur_level.extend(lines[1])
                if "Set" in lines[0]:
                    validate = """data{var}, _ := opts.Marshal({inter}.Msg())
                    assert.NotNil(t, {inter}.FromJson(string(data{var})))
                    assert.NotNil(t, {inter}.FromYaml(string(data{var})))
                    assert.NotNil(t, {inter}.FromPbText(proto.MarshalTextString({inter}.Msg())))""".format(var=var, inter=updated_interface)
                else:
                    validate = "_ = {inter}.Validate()".format(inter=updated_interface)
                recur_level.append(validate)
            else:
                # TODO: Skipping Layer1 & SetDvalues for now
                if ('Ieee' in interface_field[1] or
                        'SetChoice' in interface_field[0] or
                        'SetDValues' in interface_field[0]):
                    continue
                if interface_field[1] not in interface_dict.keys():
                    if 'Has' in interface_field[0]:
                        val = None
                    else:
                        if 'Enum' in interface_field[1]:
                            val = None
                        else:
                            if 'max' in interface_field[1]:
                                if 'maxNone' in interface_field[1]:
                                    minimum = interface_field[1].split("max")[0]
                                    val = int(minimum.split("min")[1]) - 1
                                else:
                                    minimum = (interface_field[1].split("max")[0].split("min")[1])
                                    val = int(minimum) - 1
                            else:
                                if interface_field[1] in type_dict.keys():
                                    val = type_dict[interface_field[1]]
                                else:
                                    val = None
                            if (interface_field[0] == 'SetValues' and "max" in interface_field[1]):
                                max_value = int(interface_field[1].split("max")[1])
                                if max_value > 2147483647:
                                    val = "[]int64" + "{" + str(val) + "}"
                                else:
                                    val = "[]int32" + "{" + str(val) + "}"
                        if val is not None and interface_field[0].startswith("Set"):
                            if val == 0:
                                val = -1
                            first_level = "{}.{}({})".format(
                                first_level, interface_field[0], str(val))
                else:
                    method = interface_field[0].split("Set", 1)[1]
                    if var in var_lst:
                        var = max(var_lst) + 1
                    updated_interface = "v{}".format(str(var))
                    new_var = "v{}:={}.{}()".format(str(var), main_interface, method)
                    recur_level.append(new_var)
                    var_lst.append(var)
                    lines = self._call_field_on_inter(updated_interface,
                                                      interface_dict,
                                                      interface_field[1],
                                                      var,
                                                      var_lst)
                    recur_level.append(lines[0])
                    recur_level.extend(lines[1] + lines[2])
                    if "Set" in lines[0]:
                        validate = """data{var}, _ := opts.Marshal({inter}.Msg())
                        assert.NotNil(t, {inter}.FromJson(string(data{var})))
                        assert.NotNil(t, {inter}.FromYaml(string(data{var})))
                        assert.NotNil(t, {inter}.FromPbText(proto.MarshalTextString({inter}.Msg())))""".format(var=var, inter=updated_interface)
                    else:
                        validate = "_ = {inter}.Validate()".format(inter=updated_interface)
                    recur_level.append(validate)
        return first_level, recur_level, first_level_get

    def _write_go_required_test(self):
        filename = os.path.normpath(
            os.path.join(self._ux_path, "generated_required_test.go"))
        fp = open(filename, "wb")

        self._write_in_go_test(fp,
                               "package {go_sdk_package_name}_test".format(
                                   go_sdk_package_name=self._go_sdk_package_name))
        line = 'import {go_sdk_package_name} "{go_sdk_pkg_dir}"'.format(
            go_sdk_package_name=self._go_sdk_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write_in_go_test(fp, line)
        self._write_in_go_test(fp, 'import "testing"')
        self._write_in_go_test(fp, 'import "github.com/stretchr/testify/assert"')
        self._write_in_go_test(fp, 'import "google.golang.org/protobuf/encoding/protojson"')
        self._write_in_go_test(fp, 'import "github.com/golang/protobuf/proto"')
        for component in self._api.components:
            new = self._api.components[component]
            if 'Iter' in new.interface:
                continue
            else:
                self._write_test(fp, new)

    def _write_test(self, fp, new):
        required_fields = []
        for interface_field in new.interface_fields:
            if interface_field.isOptional is False and interface_field.isArray is False:
                if interface_field.name != 'Choice':
                    required_fields.append(interface_field.name)
        if len(required_fields) > 0:
            initial_fill = """func Test{x}Required(t *testing.T) {{
            object := {pkg}.New{x}()
            opts := protojson.MarshalOptions{{
            UseProtoNames:   true,
            AllowPartial:    true,
            EmitUnpopulated: false,
            Indent:          "  ",
            }}
            data, _ := opts.Marshal(object.Msg())
            err := object.FromJson(string(data))
            err1 := object.FromYaml(string(data))
            err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))""".format(pkg=self._go_sdk_package_name, x=new.interface)
            self._write_in_go_test(fp, initial_fill)
            required_fields = ', '.join('"{0}"'.format(w) for w in required_fields)
            self._write_in_go_test(fp, "assert.Contains(t, err.Error(),{})".format(required_fields))
            self._write_in_go_test(fp, "assert.Contains(t, err1.Error(),{})".format(required_fields))
            self._write_in_go_test(fp, "assert.Contains(t, err2.Error(),{})".format(required_fields))
            self._write_in_go_test(fp, "}")

    def _write_go_invalid_format_test(self):
        var = """incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`
        """
        filename = os.path.normpath(
            os.path.join(self._ux_path, "generated_invalid_format_test.go"))
        fp = open(filename, "wb")

        self._write_in_go_test(fp,
                               "package {go_sdk_package_name}_test".format(
                                   go_sdk_package_name=self._go_sdk_package_name))
        line = 'import {go_sdk_package_name} "{go_sdk_pkg_dir}"'.format(
            go_sdk_package_name=self._go_sdk_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write_in_go_test(fp, line)
        self._write_in_go_test(fp, 'import "testing"')
        self._write_in_go_test(fp, 'import "github.com/stretchr/testify/assert"')
        for component in self._api.components:
            new = self._api.components[component]
            if 'Iter' in new.interface:
                continue
            else:
                initial_fill = """func Test{x}IncorrectFormat(t *testing.T) {{\n{var}\n
                object := {pkg}.New{x}()""".format(pkg=self._go_sdk_package_name, x=new.interface, var=var)
                self._write_in_go_test(fp, initial_fill)
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromYaml(incorrect_format))")
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromJson(incorrect_format))")
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromPbText(incorrect_format))")
                self._write_in_go_test(fp, "}")

    def _write_go_invalid_key_test(self):
        var = """incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`
        """
        filename = os.path.normpath(
            os.path.join(self._ux_path, "generated_invalid_key_test.go"))
        fp = open(filename, "wb")

        self._write_in_go_test(fp,
                               "package {go_sdk_package_name}_test".format(
                                   go_sdk_package_name=self._go_sdk_package_name))
        line = 'import {go_sdk_package_name} "{go_sdk_pkg_dir}"'.format(
            go_sdk_package_name=self._go_sdk_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir
        )
        self._write_in_go_test(fp, line)
        self._write_in_go_test(fp, 'import "testing"')
        self._write_in_go_test(fp, 'import "github.com/stretchr/testify/assert"')
        for component in self._api.components:
            new = self._api.components[component]
            if 'Iter' in new.interface:
                continue
            else:
                initial_fill = """func Test{x}IncorrectKey(t *testing.T) {{\n{var}\n
                object := {pkg}.New{x}()""".format(pkg=self._go_sdk_package_name, x=new.interface, var=var)
                self._write_in_go_test(fp, initial_fill)
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromYaml(incorrect_key))")
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromJson(incorrect_key))")
                self._write_in_go_test(fp, "assert.NotNil(t, object.FromPbText(incorrect_key))")
                self._write_in_go_test(fp, "}")
