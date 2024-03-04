from .openapiartplugin import OpenApiArtPlugin, type_limits
import os
import re
import subprocess


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
        self.status = {}


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
        self.status = None
        self.status_info = None

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
        self.status = None
        self.status_msg = None
        self.x_enum_status = {}
        self.x_constraints = []
        self.x_unique = None
        self.iter_name = None
        self.choice_with_no_prop = (
            []
        )  # maintain a list of choices with no properties for adding getter methods


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
        self._base_url = ""
        self._proto_service = kwargs.get("proto_service")
        self._oapi_go_types = {
            "string": "string",
            "boolean": "bool",
            "integer": "int32",
            "int32": "int32",
            "int64": "int64",
            "uint32": "uint32",
            "uint64": "uint64",
            "number": "float32",
            "numberfloat": "float32",
            "numberdouble": "float64",
            "stringbinary": "[]byte",
        }

    def generate(self, openapi):
        self._base_url = ""
        self._openapi = openapi
        self._ux_path = os.path.normpath(
            os.path.join(
                self._output_dir,
                "..",
                os.path.split(self._go_sdk_package_dir)[-1],
            )
        )
        self._protoc_path = os.path.normpath(
            os.path.join(self._ux_path, self._protobuf_package_name)
        )
        self._structs = {}
        self._get_base_url()
        self._write_mod_file()
        self._write_go_file()
        self._format_go_file()
        self._tidy_mod_file()

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

    def _write_mod_file(self):
        self._filename = os.path.normpath(
            os.path.join(self._ux_path, "go.mod")
        )
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write("module {}".format(self._go_sdk_package_dir))
        self._write()
        self._write("go 1.16")
        self._close_fp()

    def _write_go_file(self):
        self._filename = os.path.normpath(
            os.path.join(
                self._ux_path, "{}.go".format(self._go_sdk_package_name)
            )
        )
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
        self._write(
            self._justify_desc(
                self._info + "\n" + self._license, use_multi=True
            )
        )
        self._write()

    def _write_package(self):
        self._write(
            "package {go_sdk_package_name}".format(
                go_sdk_package_name=self._go_sdk_package_name
            )
        )
        self._write()

    def _write_common_code(self):
        """Writes the base wrapper go code"""
        line = 'import {pb_pkg_name} "{go_sdk_pkg_dir}/{pb_pkg_name}"'.format(
            pb_pkg_name=self._protobuf_package_name,
            go_sdk_pkg_dir=self._go_sdk_package_dir,
        )
        self._write(line)
        self._write('import "google.golang.org/protobuf/types/known/emptypb"')
        self._write('import "google.golang.org/grpc"')
        self._write('import "google.golang.org/grpc/credentials/insecure"')
        self._write('import "github.com/ghodss/yaml"')
        self._write('import "google.golang.org/protobuf/encoding/protojson"')
        self._write('import "google.golang.org/protobuf/proto"')
        go_pkg_fp = self._fp
        go_pkg_filename = self._filename
        self._filename = os.path.normpath(
            os.path.join(self._ux_path, "common.go")
        )
        self._init_fp(self._filename)
        self._write_package()
        with open(os.path.join(os.path.dirname(__file__), "common.go")) as fp:
            self._write(fp.read().strip().strip("\n"))
        self._write()

        self._filename = os.path.normpath(
            os.path.join(self._ux_path, "common_test.go")
        )
        self._init_fp(self._filename)
        self._write_package()
        with open(
            os.path.join(os.path.dirname(__file__), "common_test.go")
        ) as fp:
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
            jsonpath = "$.{}..schema".format(
                parser_result[0].value[2:].replace("/", ".")
            )
            schema = self._get_parser(jsonpath).find(self._openapi)[0].value
            response_component_ref = self._get_parser("$..'$ref'").find(schema)
            return response_component_ref
        return parser_result

    def _build_api_interface(self):
        self._api.internal_struct_name = """{internal_name}Api""".format(
            internal_name=self._get_internal_name(self._go_sdk_package_name)
        )
        self._api.external_interface_name = """{external_name}Api""".format(
            external_name=self._get_external_struct_name(
                self._go_sdk_package_name
            )
        )
        for url, path_object in self._openapi["paths"].items():
            http_url = self._base_url + url
            if http_url.startswith("/"):
                http_url = http_url[1:]
            for operation_id in self._get_parser("$..operationId").find(
                path_object
            ):
                path_item_object = operation_id.context.value
                rpc = FluentRpc()
                http = FluentHttp()
                rpc.operation_name = self._get_external_struct_name(
                    operation_id.value
                )
                rpc.description = self._get_description(path_item_object, True)
                http.operation_name = self._get_external_struct_name(
                    operation_id.value
                )
                if path_item_object.get("x-status", {}).get("status") in [
                    "deprecated",
                    "under_review",
                ]:
                    rpc.status = path_item_object.get("x-status", {})
                    rpc.status["status"] = rpc.status["status"].replace(
                        "-", "_"
                    )
                http.description = self._get_description(path_item_object)
                if (
                    len(
                        [
                            m
                            for m in self._api.external_rpc_methods
                            if m.operation_name == rpc.operation_name
                        ]
                    )
                    == 0
                ):
                    self._api.external_rpc_methods.append(rpc)
                if (
                    len(
                        [
                            m
                            for m in self._api.internal_http_methods
                            if m.operation_name == http.operation_name
                        ]
                    )
                    == 0
                ):
                    self._api.internal_http_methods.append(http)
                rpc.request_return_type = (
                    "{operation_response_name}Response".format(
                        operation_response_name=self._get_external_struct_name(
                            rpc.operation_name
                        ),
                    )
                )
                binary_type = self._get_parser(
                    "$..responses..'200'..schema..format"
                ).find(path_item_object)
                ref_type = self._get_parser(
                    "$..responses..'200'..'$ref'"
                ).find(path_item_object)
                if len(binary_type) == 1:
                    rpc.request_return_type = "[]byte"
                elif len(ref_type) == 1:
                    request_return_type = (
                        self._get_schema_object_name_from_ref(
                            self._resolve_response(ref_type)[0].value
                        )
                    )
                    rpc.request_return_type = self._get_external_struct_name(
                        request_return_type
                    )
                else:
                    rpc.request_return_type = "*string"

                http.request_return_type = rpc.request_return_type
                ref = self._get_parser("$..requestBody..'$ref'").find(
                    path_item_object
                )
                if len(ref) == 1:
                    new = FluentNew()
                    new.schema_name = self._get_schema_object_name_from_ref(
                        ref[0].value
                    )
                    new.schema_object = self._get_schema_object_from_ref(
                        ref[0].value
                    )
                    new.interface = self._get_external_struct_name(
                        new.schema_name
                    )
                    new.struct = self._get_internal_name(new.schema_name)
                    new.description = self._get_description(
                        new.schema_object, True
                    )
                    new.method_description = """// New{interface} returns a new instance of {interface}.""".format(
                        interface=new.interface
                    )

                    description = "// {} is {}".format(
                        new.interface,
                        self._get_description(new.schema_object, True).lstrip(
                            "// "
                        ),
                    )

                    new.method_description = (
                        description + "\n" + new.method_description
                    )

                    new.method = """New{interface}() {interface}""".format(
                        interface=new.interface
                    )
                    self._populate_status(new)
                    if (
                        len(
                            [
                                m
                                for m in self._api.external_new_methods
                                if m.schema_name == new.schema_name
                            ]
                        )
                        == 0
                    ):
                        self._api.external_new_methods.append(new)
                    rpc.request = "{pb_pkg_name}.{operation_name}Request{{{interface}: {struct}.msg()}}".format(
                        pb_pkg_name=self._protobuf_package_name,
                        operation_name=rpc.operation_name,
                        interface=new.interface,
                        struct=new.struct,
                    )
                    rpc.description = "// {} {}".format(
                        rpc.operation_name, rpc.description.lstrip("// ")
                    )
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
                        struct=new.struct,
                        interface=new.interface,
                        request_return_type=rpc.request_return_type,
                    )
                    rpc.validate = """
                        if err := {struct}.validate(); err != nil {{
                            return nil, err
                        }}
                    """.format(
                        struct=new.struct
                    )
                    rpc.http_call = (
                        """return api.http{operation_name}({struct})""".format(
                            operation_name=rpc.operation_name,
                            struct=new.struct,
                        )
                    )
                    if url.startswith("/"):
                        url = url[1:]
                    http.request = """{struct}Json, err := {struct}.Marshal().ToJson()
                    if err != nil {{return nil, err}}
                    resp, err := api.httpSendRecv("{url}", {struct}Json, "{method}")
                    """.format(
                        url=http_url,
                        struct=new.struct,
                        method=str(
                            operation_id.context.path.fields[0]
                        ).upper(),
                    )
                    http.method = """http{rpc_method}""".format(
                        rpc_method=rpc.method
                    )
                else:
                    rpc.description = "// {} {}".format(
                        rpc.operation_name, rpc.description.lstrip("// ")
                    )
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
                        request_return_type=rpc.request_return_type,
                    )
                    rpc.http_call = (
                        """return api.http{operation_name}()""".format(
                            operation_name=rpc.operation_name,
                        )
                    )
                    http.request = """resp, err := api.httpSendRecv("{url}", "", "{method}")""".format(
                        url=http_url,
                        method=str(
                            operation_id.context.path.fields[0]
                        ).upper(),
                    )
                    http.method = """http{rpc_method}""".format(
                        rpc_method=rpc.method
                    )
                for ref in self._get_parser("$..responses").find(
                    path_item_object
                ):
                    for status_code, response_object in ref.value.items():
                        response = FluentRpcResponse()
                        response.status_code = status_code
                        response.request_return_type = (
                            """New{operation_name}Response""".format(
                                operation_name=rpc.operation_name,
                            )
                        )
                        ref = self._get_parser("$..'$ref'").find(
                            response_object
                        )
                        if len(ref) > 0:
                            response.schema = {
                                "$ref": self._resolve_response(ref)[0].value
                            }
                        else:
                            schema = self._get_parser("$..schema").find(
                                response_object
                            )
                            if len(schema) > 0:
                                response.schema = (
                                    self._get_parser("$..schema")
                                    .find(response_object)[0]
                                    .value
                                )
                            else:
                                response.schema = {"type": "string"}
                        rpc.responses.append(response)
                        http.responses.append(response)

        self._build_response_interfaces()

        # write the go code
        self._write(
            """
            // function related to error handling
            func FromError(err error) (Error, bool) {{
                if rErr, ok := err.(Error); ok {{
                    return rErr, true
                }}

                rErr := NewError()
                if err := rErr.Unmarshal().FromJson(err.Error()); err == nil {{
                    return rErr, true
                }}

                return fromGrpcError(err)
            }}

            func setResponseErr(obj Error, code int32, message string) {{
                errors := []string{{}}
                errors = append(errors, message)
                obj.msg().Code = &code
                obj.msg().Errors = errors
            }}

            // parses and return errors for grpc response
            func fromGrpcError(err error) (Error, bool) {{
                st, ok := status.FromError(err)
                if ok {{
                    rErr := NewError()
                    if err := rErr.Unmarshal().FromJson(st.Message()); err == nil {{
                        var code = int32(st.Code())
                        rErr.msg().Code = &code
                        return rErr, true
                    }}

                    setResponseErr(rErr, int32(st.Code()), st.Message())
                    return rErr, true
                }}

                return nil, false
            }}

            // parses and return errors for http responses
            func fromHttpError(statusCode int, body []byte) Error {{
                rErr := NewError()
                bStr := string(body)
                if err := rErr.Unmarshal().FromJson(bStr); err == nil {{
                    return rErr
                }}

                setResponseErr(rErr, int32(statusCode), bStr)

                return rErr
            }}

            type versionMeta struct {{
                checkVersion  bool
                localVersion  Version
                remoteVersion Version
                checkError    error
            }}
            type {internal_struct_name} struct {{
                apiSt
                grpcClient {pb_pkg_name}.{proto_service}Client
                httpClient httpClient
                versionMeta *versionMeta
            }}

            // grpcConnect builds up a grpc connection
            func (api *{internal_struct_name}) grpcConnect() error {{
                if api.grpcClient == nil {{
                    if api.grpc.clientConnection == nil {{
                        ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.dialTimeout)
                        defer cancelFunc()
                        conn, err := grpc.DialContext(ctx, api.grpc.location, grpc.WithTransportCredentials(insecure.NewCredentials()))
                        if err != nil {{
                            return err
                        }}
                        api.grpcClient = {pb_pkg_name}.New{proto_service}Client(conn)
                        api.grpc.clientConnection = conn
                    }} else {{
                        api.grpcClient = {pb_pkg_name}.New{proto_service}Client(api.grpc.clientConnection)
                    }}
                }}
                return nil
            }}

            func (api *{internal_struct_name}) grpcClose() error {{
                if api.grpc != nil {{
                    if api.grpc.clientConnection != nil {{
                        err := api.grpc.clientConnection.Close()
                        if err != nil {{
                            return err
                        }}
                    }}
                }}
                api.grpcClient = nil
                api.grpc = nil
                return nil
            }}

            func (api *{internal_struct_name}) Close() error {{
                if api.hasGrpcTransport() {{
                    err := api.grpcClose()
                    return err
                }}
                if api.hasHttpTransport() {{
                    err := api.http.conn.(*net.TCPConn).SetLinger(0)
                    api.http.conn.Close()
                    api.http.conn = nil
                    api.http = nil
                    api.httpClient.client = nil
                    return err
                }}
                return nil
            }}

            //  NewApi returns a new instance of the top level interface hierarchy
            func NewApi() Api {{
                api := {internal_struct_name}{{}}
                api.versionMeta = &versionMeta{{checkVersion: false}}
                return &api
            }}

            // httpConnect builds up a http connection
            func (api *{internal_struct_name}) httpConnect() error {{
                if api.httpClient.client == nil {{
                    tr := http.Transport{{
                        DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {{
                            tcpConn, err := (&net.Dialer{{}}).DialContext(ctx, network, addr)
                            if err != nil {{
                                return nil, err
                            }}
                            tlsConn := tls.Client(tcpConn, &tls.Config{{InsecureSkipVerify: !api.http.verify}})
                            err = tlsConn.Handshake()
                            if err != nil {{
                                return nil, err
                            }}
                            api.http.conn = tcpConn
                            return tlsConn, nil
                        }},
                        DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {{
                            tcpConn, err := (&net.Dialer{{}}).DialContext(ctx, network, addr)
                            if err != nil {{
                                return nil, err
                            }}
                            api.http.conn = tcpConn
                            return tcpConn, nil
                        }},
                    }}
                    client := httpClient{{
                        client: &http.Client{{
                            Transport: &tr,
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
                response, err := httpClient.client.Do(req)
                return response, err
            }}
            """.format(
                internal_struct_name=self._api.internal_struct_name,
                pb_pkg_name=self._protobuf_package_name,
                proto_service=self._proto_service,
            )
        )
        methods = []
        # remove new methopds from interface
        # for new in self._api.external_new_methods:
        #     methods.append(new.method_description)
        #     methods.append(new.method)
        for rpc in self._api.external_rpc_methods:
            methods.append(rpc.description)
            methods.append(rpc.method)
            # descriptions.append("(*{}).{}".format(self._api.external_interface_name, rpc.method_description))
        if self._generate_version_api:
            methods.extend(self._get_version_api_interface_method_signatures())
        method_signatures = "\n".join(methods)
        self._write(
            """
            {description}
            type Api interface {{
                api
                {method_signatures}
            }}
            """.format(
                method_signatures=method_signatures,
                description="// {} {}".format(
                    self._api.external_interface_name,
                    self._get_description(self._openapi["info"], True).lstrip(
                        "// "
                    ),
                ),
            )
        )
        # remove new methods from api
        # for new in self._api.external_new_methods:
        #     self._write(
        #         """func (api *{internal_struct_name}) {method} {{
        #             return New{interface}()
        #         }}
        #         """.format(
        #             internal_struct_name=self._api.internal_struct_name,
        #             method=new.method,
        #             interface=new.interface,
        #         )
        #     )
        if self._generate_version_api:
            self._write(
                self._get_version_api_interface_method_impl(
                    self._api.internal_struct_name
                )
            )
        for rpc in self._api.external_rpc_methods:
            if rpc.request_return_type == "[]byte":
                return_value = """if resp.ResponseBytes != nil {
                        return resp.ResponseBytes, nil
                    }
                    return nil, nil"""
            elif rpc.request_return_type == "*string":
                return_value = """if resp.GetString_() != "" {
                        status_code_value := resp.GetString_()
                        return &status_code_value, nil
                    }
                    return nil, nil"""
            else:
                return_value = """ret := New{struct}()
                    if resp.Get{struct}() != nil {{
                        return ret.setMsg(resp.Get{struct}()), nil
                    }}

                    return ret, nil""".format(
                    struct=self._get_external_struct_name(
                        rpc.request_return_type
                    ),
                )

            info = rpc.status.get("information")
            status_type = rpc.status.get("status")
            status_str = ""
            if status_type is not None:
                status_str = self._get_status_msg(
                    rpc.operation_name, status_type, info, "api"
                )

            if self._generate_version_api:
                version_check = """
                    if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
                        return nil, err
                    }"""
            else:
                version_check = ""
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    {status}
                    {validate}
                    {version_check}
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
                        if er, ok := fromGrpcError(err); ok {{
                            return nil, er
                        }}
                        return nil, err
                    }}
                    {return_value}
                }}
                """.format(
                    internal_struct_name=self._api.internal_struct_name,
                    method=rpc.method,
                    status=status_str,
                    request=rpc.request,
                    operation_name=rpc.operation_name,
                    return_value=return_value,
                    http_call=rpc.http_call,
                    validate=getattr(rpc, "validate", ""),
                    version_check=""
                    if rpc.method == "GetVersion() (Version, error)"
                    else version_check,
                )
            )

        for http in self._api.internal_http_methods:
            error_handling = ""
            success_method = None
            for response in http.responses:
                if response.status_code.startswith("2"):
                    success_method = response.request_return_type
                else:
                    error_handling += (
                        "return nil, fromHttpError(resp.StatusCode, bodyBytes)"
                    )

            if http.request_return_type == "[]byte":
                success_handling = """return bodyBytes, nil"""
            elif http.request_return_type == "*string":
                success_handling = """bodyString := string(bodyBytes)
                return &bodyString, nil"""
            else:
                success_handling = """obj := {success_method}().{struct}()
                    if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {{
                        return nil, err
                    }}
                    return obj, nil""".format(
                    success_method=success_method,
                    struct=http.request_return_type,
                )
            # TODO: do not hardcode 200 status code
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    {request}
                    if err != nil {{
                        return nil, err
                    }}
                    bodyBytes, err := io.ReadAll(resp.Body)
                    defer resp.Body.Close()
                    if err != nil {{
                        return nil, err
                    }}
                    if resp.StatusCode == 200 {{
                        {success_handling}
                    }} else {{
                        {error_handling}
                    }}
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

    def _get_version_api_interface_method_signatures(self):
        return [
            "// GetLocalVersion provides version details of local client",
            "GetLocalVersion() Version",
            "// GetRemoteVersion provides version details received from remote server",
            "GetRemoteVersion() (Version, error)",
            "// SetVersionCompatibilityCheck allows enabling or disabling automatic version",
            "// compatibility check between client and server API spec version upon API call",
            "SetVersionCompatibilityCheck(bool)",
            "// CheckVersionCompatibility compares API spec version for local client and remote server,",
            "// and returns an error if they are not compatible according to Semantic Versioning 2.0.0",
            "CheckVersionCompatibility() error",
        ]

    def _get_version_api_interface_method_impl(self, struct_name):
        return """
            func (api *{0}) GetLocalVersion() Version {{
                if api.versionMeta.localVersion == nil {{
                    api.versionMeta.localVersion = NewVersion().SetApiSpecVersion("{1}").SetSdkVersion("{2}")
                }}

                return api.versionMeta.localVersion
            }}

            func (api *{0}) GetRemoteVersion() (Version, error) {{
                if api.versionMeta.remoteVersion == nil {{
                    v, err := api.GetVersion()
                    if err != nil {{
                        return nil, fmt.Errorf("could not fetch remote version: %v", err)
                    }}

                    api.versionMeta.remoteVersion = v
                }}

                return api.versionMeta.remoteVersion, nil
            }}

            func (api *{0}) SetVersionCompatibilityCheck(v bool) {{
                api.versionMeta.checkVersion = v
            }}

            func (api *{0}) checkLocalRemoteVersionCompatibility() (error, error) {{
                localVer := api.GetLocalVersion()
                remoteVer, err := api.GetRemoteVersion()
                if err != nil {{
                    return nil, err
                }}
                err = checkClientServerVersionCompatibility(localVer.ApiSpecVersion(), remoteVer.ApiSpecVersion(), "API spec")
                if err != nil {{
                    return fmt.Errorf(
                        "client SDK version '%s' is not compatible with server SDK version '%s': %v",
                        localVer.SdkVersion(), remoteVer.SdkVersion(), err,
                    ), nil
                }}

                return nil, nil
            }}

            func (api *{0}) checkLocalRemoteVersionCompatibilityOnce() error {{
                if !api.versionMeta.checkVersion {{
                    return nil
                }}

                if api.versionMeta.checkError != nil {{
                    return api.versionMeta.checkError
                }}

                compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
                if compatErr != nil {{
                    api.versionMeta.checkError = compatErr
                    return compatErr
                }}
                if apiErr != nil {{
                    api.versionMeta.checkError = nil
                    return apiErr
                }}

                api.versionMeta.checkVersion = false
                api.versionMeta.checkError = nil
                return nil
            }}

            func (api *{0}) CheckVersionCompatibility() error {{
                compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
                if compatErr != nil {{
                    return fmt.Errorf("version error: %v", compatErr)
                }}
                if apiErr != nil {{
                    return apiErr
                }}

                return nil
            }}
        """.format(
            struct_name, self._api_version, self._sdk_version
        )

    def _write_component_interfaces(self):
        while True:
            components = [
                component
                for _, component in self._api.components.items()
                if component.generated is False
            ]
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
                properties[
                    "status_code_{}".format(response.status_code)
                ] = response.schema
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
            new.method_description = """// New{interface} returns a new instance of {interface}.""".format(
                interface=new.interface
            )

            description = "// {} is {}".format(
                new.interface,
                self._get_description(new.schema_object, True).lstrip("// "),
            )

            new.method_description = (
                description + "\n" + new.method_description
            )
            new.schema_name = self._get_external_struct_name(new.interface)
            self._populate_status(new)
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
        internal_items_nil = []
        for field in new.interface_fields:
            if field.struct and field.isArray is False:
                internal_items.append(
                    "{} {}".format(self._get_holder_name(field), field.type)
                )
                internal_items_nil.append(
                    "obj.{} = nil".format(self._get_holder_name(field))
                )
            if field.adder_method is not None and field.isArray:
                internal_items.append(
                    "{} {}".format(
                        self._get_holder_name(field),
                        field.iter_name,
                    )
                )
                internal_items_nil.append(
                    "obj.{} = nil".format(self._get_holder_name(field))
                )
        self._write(
            """
            // ***** {interface} *****
            type {struct} struct {{
                validation
                obj *{pb_pkg_name}.{interface}
                marshaller marshal{interface}
                unMarshaller unMarshal{interface}
                {internal_items}
            }}

            func New{interface}() {interface} {{
                obj := {struct}{{obj: &{pb_pkg_name}.{interface}{{}}}}
                obj.setDefault()
                return &obj
            }}

            func (obj *{struct}) msg() *{pb_pkg_name}.{interface} {{
                return obj.obj
            }}

            func (obj *{struct}) setMsg(msg *{pb_pkg_name}.{interface}) {interface} {{
                {msg_nil_call}
                proto.Merge(obj.obj, msg)
                return obj
            }}

            type marshal{struct} struct {{
                obj *{struct}
            }}

            type marshal{interface} interface {{
                // ToProto marshals {interface} to protobuf object *{pb_pkg_name}.{interface}
                ToProto() (*{pb_pkg_name}.{interface}, error)
                // ToPbText marshals {interface} to protobuf text
                ToPbText() (string, error)
                // ToYaml marshals {interface} to YAML text
                ToYaml() (string, error)
                // ToJson marshals {interface} to JSON text
                ToJson() (string, error)
            }}

            type unMarshal{struct} struct {{
                obj *{struct}
            }}

            type unMarshal{interface} interface {{
                // FromProto unmarshals {interface} from protobuf object *{pb_pkg_name}.{interface}
                FromProto(msg *{pb_pkg_name}.{interface}) ({interface}, error)
                // FromPbText unmarshals {interface} from protobuf text
                FromPbText(value string) error
                // FromYaml unmarshals {interface} from YAML text
                FromYaml(value string) error
                // FromJson unmarshals {interface} from JSON text
                FromJson(value string) error
            }}

            func (obj *{struct}) Marshal() marshal{interface} {{
                if obj.marshaller == nil {{
                    obj.marshaller = &marshal{struct}{{obj: obj}}
                }}
                return obj.marshaller
            }}

            func (obj *{struct}) Unmarshal() unMarshal{interface} {{
                if obj.unMarshaller == nil {{
                    obj.unMarshaller = &unMarshal{struct}{{obj: obj}}
                }}
                return obj.unMarshaller
            }}

            func (m *marshal{struct}) ToProto() (*{pb_pkg_name}.{interface}, error) {{
                err := m.obj.validateToAndFrom()
                if err != nil {{
                    return nil, err
                }}
                return m.obj.msg(), nil
            }}

            func (m *unMarshal{struct}) FromProto(msg *{pb_pkg_name}.{interface}) ({interface}, error) {{
                newObj := m.obj.setMsg(msg)
                err := newObj.validateToAndFrom()
                if err != nil {{
                    return nil, err
                }}
                return newObj, nil
            }}

            func (m *marshal{struct}) ToPbText() (string, error) {{
                vErr := m.obj.validateToAndFrom()
                if vErr != nil {{
                    return "", vErr
                }}
                protoMarshal, err := proto.Marshal(m.obj.msg())
                if err != nil {{
                    return "", err
                }}
                return string(protoMarshal), nil
            }}

            func (m *unMarshal{struct}) FromPbText(value string) error {{
                retObj := proto.Unmarshal([]byte(value), m.obj.msg())
                if retObj != nil {{
                    return retObj
                }}
                {nil_call}
                vErr := m.obj.validateToAndFrom()
                if vErr != nil {{
                    return vErr
                }}
                return retObj
            }}

            func (m *marshal{struct}) ToYaml() (string, error) {{
                vErr := m.obj.validateToAndFrom()
                if vErr != nil {{
                    return "", vErr
                }}
                opts := protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: false,
                }}
                data, err := opts.Marshal(m.obj.msg())
                if err != nil {{return "", err}}
                data, err = yaml.JSONToYAML(data)
                if err != nil {{
                    return "", err
                }}
                return string(data), nil
            }}

            func (m *unMarshal{struct}) FromYaml(value string) error {{
                if value == "" {{value = "{{}}"}}
                data, err := yaml.YAMLToJSON([]byte(value))
                if err != nil {{
                    return err
                }}
                opts := protojson.UnmarshalOptions{{
                    AllowPartial: true,
                    DiscardUnknown: false,
                }}
                uError := opts.Unmarshal([]byte(data), m.obj.msg())
                if uError != nil {{
                    return fmt.Errorf("unmarshal error %s", strings.Replace(
                        uError.Error(), "\\u00a0", " ", -1)[7:])
                }}
                {nil_call}
                vErr := m.obj.validateToAndFrom()
                if vErr != nil {{
                    return vErr
                }}
                return nil
            }}

            func (m *marshal{struct}) ToJson() (string, error) {{
                vErr := m.obj.validateToAndFrom()
                if vErr != nil {{
                    return "", vErr
                }}
                opts := protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: false,
                    Indent:          "  ",
                }}
                data, err := opts.Marshal(m.obj.msg())
                if err != nil {{
                    return "", err
                }}
                return string(data), nil
            }}

            func (m *unMarshal{struct}) FromJson(value string) error {{
                opts := protojson.UnmarshalOptions{{
                    AllowPartial: true,
                    DiscardUnknown: false,
                }}
                if value == "" {{value = "{{}}"}}
                uError := opts.Unmarshal([]byte(value), m.obj.msg())
                if uError != nil {{
                    return fmt.Errorf("unmarshal error %s", strings.Replace(
                        uError.Error(), "\\u00a0", " ", -1)[7:])
                }}
                {nil_call}
                err := m.obj.validateToAndFrom()
                if err != nil {{
                    return err
                }}
                return nil
            }}

            func (obj *{struct}) validateToAndFrom() error {{
                // emptyVars()
                obj.validateObj(&obj.validation, true)
                return obj.validationResult()
            }}

            func (obj *{struct}) validate() error {{
                // emptyVars()
                obj.validateObj(&obj.validation, false)
                return obj.validationResult()
            }}

            func (obj *{struct}) String() string {{
                str, err := obj.Marshal().ToYaml()
                if err != nil {{
                    return err.Error()
                }}
                return str
            }}

            func (obj *{struct}) Clone() ({interface}, error) {{
                vErr := obj.validate()
                if vErr != nil {{
                    return nil, vErr
                }}
                newObj := New{interface}()
                data, err :=  proto.Marshal(obj.msg())
                if err != nil {{
                    return nil, err
                }}
                pbErr := proto.Unmarshal(data, newObj.msg())
                if pbErr != nil {{
                    return nil, pbErr
                }}
                return newObj, nil
            }}
        """.format(
                struct=new.struct,
                pb_pkg_name=self._protobuf_package_name,
                interface=new.interface,
                internal_items=""
                if len(internal_items) == 0
                else "\n".join(internal_items),
                nil_call="m.obj.setNil()"
                if len(internal_items_nil) > 0
                else "",
                msg_nil_call="obj.setNil()"
                if len(internal_items_nil) > 0
                else "",
            )
        )
        if len(internal_items_nil) > 0:
            self._write(
                """
                func (obj *{struct}) setNil() {{
                    {nil_items}
                    obj.validationErrors = nil
                    obj.warnings = nil
                    obj.constraints = make(map[string]map[string]Constraints)
                }}
            """.format(
                    nil_items="\n".join(internal_items_nil), struct=new.struct
                )
            )

        interfaces = [
            "// msg marshals {interface} to protobuf object *{pb_pkg_name}.{interface}",
            "// and doesn't set defaults",
            "msg() *{pb_pkg_name}.{interface}",
            "// setMsg unmarshals {interface} from protobuf object *{pb_pkg_name}.{interface}",
            "// and doesn't set defaults",
            "setMsg(*{pb_pkg_name}.{interface}) {interface}",
            "// provides marshal interface",
            "Marshal() marshal{interface}",
            "// provides unmarshal interface",
            "Unmarshal() unMarshal{interface}",
            "// validate validates {interface}",
            "validate() error",
            "// A stringer function",
            "String() string",
            "// Clones the object",
            "Clone() ({interface}, error)",
            "validateToAndFrom() error",
            "validateObj(vObj *validation, set_default bool)",
            "setDefault()",
        ]
        for field in new.interface_fields:
            interfaces.append(
                "// {}".format(
                    self._escaped_str(field.getter_method_description)
                )
            )
            interfaces.append(field.getter_method)
            if field.setter_method is not None:
                description = field.setter_method_description
                method = field.setter_method
                if field.name == "Choice":
                    description = field.setter_method_description.replace(
                        "SetChoice", "setChoice"
                    )
                    method = field.setter_method.replace(
                        "SetChoice", "setChoice"
                    )
                interfaces.append(
                    "// {}".format(self._escaped_str(description))
                )

                interfaces.append(method)
            if field.has_method is not None:
                interfaces.append(
                    "// {}".format(
                        self._escaped_str(field.has_method_description)
                    )
                )
                interfaces.append(field.has_method)
            for prop in field.choice_with_no_prop:
                interfaces.append(
                    "// getter for {fieldName} to set choice.\n{fieldName}()".format(
                        fieldName=self._get_external_field_name(prop)
                    )
                )
        if new.interface == "Error":
            interfaces.append(
                "// implement Error function for implementingnative Error Interface. \n Error() string"
            )
        interface_signatures = "\n".join(interfaces)
        self._write(
            """
            {description}
            type {interface} interface {{
                Validation
                {interface_signatures}
                {nil_call}
            }}
        """.format(
                interface=new.interface,
                interface_signatures=interface_signatures.format(
                    interface=new.interface,
                    pb_pkg_name=self._protobuf_package_name,
                ),
                description=""
                if new.description is None
                else "// {} is {}".format(
                    new.interface, new.description.strip("// ")
                ),
                nil_call="setNil()" if len(internal_items_nil) > 0 else "",
            )
        )

        # error-ux change for implement error fucntion inside Error struct
        if new.interface == "Error":
            self._write(
                """
                func (obj *_error) Error() string {
                    json, err := obj.Marshal().ToJson()
                    if err != nil {
                        return fmt.Sprintf("could not convert Error to JSON: %v", err)
                    }
                    return json
                }
                """
            )

        for field in new.interface_fields:
            self._write_field_getter(new, field)
            self._write_field_has(new, field)
            self._write_field_setter(new, field, len(internal_items_nil) > 0)
            self._write_field_adder(new, field)
        # TODO: restore behavior
        # self._write_value_of(new)
        self._write_validate_method(new)
        self._write_default_method(new)

    def _escaped_str(self, val):
        val = val.replace("{", "{{")
        return val.replace("}", "}}")

    def _write_field_getter(self, new, field):
        if field.getter_method is None:
            return
        if field.isArray and field.isEnum is False:
            if field.struct:
                block = (
                    "obj.obj.{name} = []*{pb_pkg_name}.{pb_struct}{{}}".format(
                        name=field.name,
                        pb_pkg_name=self._protobuf_package_name,
                        pb_struct=field.external_struct,
                    )
                )
                if field.setChoiceValue is not None:
                    block = "obj.setChoice({interface}Choice.{enum})".format(
                        interface=new.interface, enum=field.setChoiceValue
                    )
                body = """if len(obj.obj.{name}) == 0 {{
                        {block}
                    }}
                    if obj.{internal_name} == nil {{
                        obj.{internal_name} = new{iter_name}(&obj.obj.{name}).setMsg(obj)
                    }}
                    return obj.{internal_name}""".format(
                    name=field.name,
                    iter_name=field.iter_name,
                    block=block,
                    internal_name=self._get_holder_name(field),
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
                        block = (
                            """obj.Set{name}({type}{{{default}}})""".format(
                                name=field.name, type=field.type, default=value
                            )
                        )
                    else:
                        block = """
                            obj.setChoice({interface}Choice.{enum})
                        """.format(
                            interface=new.interface,
                            enum=field.setChoiceValue,
                        )
                body = """if obj.obj.{name} == nil {{
                        {block}
                    }}
                    return obj.obj.{name}""".format(
                    name=field.name, block=block
                )
        elif field.struct is not None:
            # at this time proto generation ignores the optional keyword
            # if the type is an object
            set_choice_or_new = (
                "obj.obj.{name} = New{pb_struct}().msg()".format(
                    name=field.name,
                    pb_struct=field.external_struct,
                )
            )
            if field.setChoiceValue is not None:
                set_choice_or_new = (
                    """obj.setChoice({interface}Choice.{enum})""".format(
                        interface=new.interface,
                        enum=field.setChoiceValue,
                    )
                )
            body = """if obj.obj.{name} == nil {{
                    {set_choice_or_new}
                }}
                if obj.{internal_name} == nil {{
                    obj.{internal_name} = &{struct}{{obj: obj.obj.{name}}}
                }}
                return obj.{internal_name}""".format(
                name=field.name,
                struct=field.struct,
                set_choice_or_new=set_choice_or_new,
                internal_name=self._get_holder_name(field),
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
            for prop in field.choice_with_no_prop:
                self._write(
                    """
                    // getter for {fieldname} to set choice
                    func (obj *{struct}) {fieldname}() {{
                    obj.setChoice({interface}Choice.{enum})
                }}
                """.format(
                        struct=new.struct,
                        interface=new.interface,
                        fieldname=self._get_external_field_name(prop),
                        enum=prop.upper(),
                    )
                )
            return
        elif field.isPointer:
            set_enum_choice = None
            if field.setChoiceValue is not None:
                set_enum_choice = """
                    if obj.obj.{fieldname} == nil {{
                        obj.setChoice({interface}Choice.{enum})
                    }}
                """.format(
                    fieldname=field.name,
                    interface=new.interface,
                    enum=field.setChoiceValue,
                )
            body = ""
            # TODO: restore behavior
            # if field.type == "string":
            #     body = """
            #     if obj.obj.{fieldname} == nil {{
            #         return ""
            #     }}
            #     """.format(
            #         fieldname=field.name
            #     )
            body += """{set_enum_choice}
                return *obj.obj.{fieldname}
                """.format(
                fieldname=field.name,
                set_enum_choice=set_enum_choice
                if set_enum_choice is not None
                else "",
            )
        else:
            set_enum_choice = None
            if field.setChoiceValue is not None:
                set_enum_choice = """
                    if obj.obj.{fieldname} == nil {{
                        obj.setChoice({interface}Choice.{enum})
                    }}
                """.format(
                    fieldname=field.name,
                    interface=new.interface,
                    enum=field.setChoiceValue,
                )
            body = """{set_enum_choice}\n return obj.obj.{fieldname}""".format(
                fieldname=field.name,
                set_enum_choice=set_enum_choice
                if set_enum_choice is not None
                else "",
            )
        if field.name == "ResponseString":
            self._write(
                """
                {description}\n// {fieldname} returns a {fieldtype}
                func (obj *{struct}) {getter_method} {{
                    return obj.obj.String_
                }}
                """.format(
                    fieldname=self._get_external_struct_name(field.name),
                    struct=new.struct,
                    getter_method=field.getter_method,
                    description=field.description,
                    fieldtype=field.type,
                )
            )
        else:
            self._write(
                """
                {description}\n// {fieldname} returns a {fieldtype}
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
                    # TODO: restore behavior
                    # status=""
                    # if field.status is None
                    # else "obj.{func}(`{msg}`)".format(
                    #     func=field.status, msg=field.status_msg
                    # ),
                )
            )

    def _write_field_setter(self, new, field, set_nil):
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
            enum_set = {
                self._get_external_field_name(name): name
                for name in field.enums
            }
            enum_body = []
            for enum_field in new.interface_fields:
                if enum_set.get(enum_field.name) is None:
                    continue
                if enum_field.isEnum:
                    enum_body.append(
                        "obj.obj.{name} = {pb_pkg_name}.{interface}_{name}_unspecified.Enum()".format(
                            name=enum_field.name,
                            pb_pkg_name=self._protobuf_package_name,
                            interface=new.interface,
                        )
                    )
                    continue
                if (
                    enum_field.struct is not None
                    and enum_field.isArray is False
                ):
                    enum_body.append(
                        """
                        if value == {interface}{name}.{enumupper} {{
                            obj.obj.{enumname} = New{struct}().msg()
                        }}
                    """.format(
                            interface=new.interface,
                            name=field.name,
                            enumname=enum_field.name,
                            struct=self._get_external_struct_name(
                                enum_field.struct
                            ),
                            enumupper=enum_set.get(enum_field.name).upper(),
                        )
                    )
                elif enum_field.struct is not None and enum_field.isArray:
                    enum_body.append(
                        """
                        if value == {interface}{name}.{enumupper} {{
                            obj.obj.{enumname} = []*{pb_pkg}.{struct}{{}}
                        }}
                    """.format(
                            interface=new.interface,
                            name=field.name,
                            enumname=enum_field.name,
                            struct=self._get_external_struct_name(
                                enum_field.struct
                            ),
                            enumupper=enum_set.get(enum_field.name).upper(),
                            pb_pkg=self._protobuf_package_name,
                        )
                    )
                elif enum_field.default is not None:
                    default_value = (
                        ""
                        if enum_field.default is None
                        else enum_field.default
                    )
                    if enum_field.isArray:
                        default_value = "{type}{{{value}}}".format(
                            type=enum_field.type,
                            value='"{}"'.format('", "'.join(default_value))
                            if "string" in enum_field.type
                            else ",".join([str(e) for e in default_value]),
                        )
                    else:
                        default_value = (
                            '"{}"'.format(default_value)
                            if enum_field.type == "string"
                            else "{}({})".format(
                                enum_field.type, default_value
                            )
                        )
                    enum_body.append(
                        """
                        if value == {interface}{name}.{enumupper} {{
                            defaultValue := {default_value}
                            obj.obj.{enumname} = {point}defaultValue
                        }}
                    """.format(
                            interface=new.interface,
                            name=field.name,
                            enumname=enum_field.name,
                            default_value=default_value,
                            enumupper=enum_set.get(enum_field.name).upper(),
                            point="" if enum_field.isArray else "&",
                        )
                    )
                enum_body.insert(
                    0, "obj.obj.{name} = nil".format(name=enum_field.name)
                )
                if enum_field.struct is not None:
                    enum_body.insert(
                        1,
                        "obj.{name} = nil".format(
                            name=self._get_holder_name(enum_field)
                        ),
                    )

            self._write(
                """func (obj* {struct}) {set_str}{fieldname}(value {interface}{fieldname}Enum) {interface} {{
                intValue, ok := {pb_pkg_name}.{interface}_{fieldname}_Enum_value[string(value)]
                if !ok {{
                    obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
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
                    enum_set="\n".join(enum_body)
                    if field.name == "Choice"
                    else "",
                    set_str="set" if field.name == "Choice" else "Set",
                )
            )
            return
        elif field.struct is not None:
            body = """{set_nil} = nil
            obj.obj.{name} = value.msg()
            """.format(
                set_nil="obj.{}".format(self._get_holder_name(field))
                if set_nil is True
                else "",
                name=field.name,
            )
        elif field.isPointer:
            body = """obj.obj.{fieldname} = &value""".format(
                fieldname=field.name
            )
        else:
            body = """obj.obj.{fieldname} = value""".format(
                fieldname=field.name
            )
        set_choice = ""
        if field.setChoiceValue is not None:
            set_choice = """obj.setChoice({interface}Choice.{enum})""".format(
                interface=new.interface,
                enum=field.setChoiceValue,
            )
        if field.name == "ResponseString":
            self._write(
                """
                {description}\n // Set{fieldname} sets the {fieldtype} value in the {fieldstruct} object
                func (obj *{newstruct}) {setter_method} {{
                    obj.obj.String_ = value
                    return obj
                }}
                """.format(
                    fieldname=self._get_external_struct_name(field.name),
                    newstruct=new.struct,
                    setter_method=field.setter_method,
                    description=field.description,
                    fieldtype=field.type,
                    fieldstruct=new.interface,
                )
            )
        else:
            self._write(
                """
                {description}\n // Set{fieldname} sets the {fieldtype} value in the {fieldstruct} object
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
                    # TODO: restore behavior
                    # status=""
                    # if field.status is None
                    # else "obj.{func}(`{msg}`)".format(
                    #     func=field.status, msg=field.status_msg
                    # ),
                )
            )

    def _write_field_adder(self, new, field):
        if field.adder_method is None:
            return
        interface_name = field.iter_name
        if interface_name in self._api.components:
            return
        new_iter = FluentNew()
        new_iter.schema_name = interface_name
        new_iter.interface = interface_name
        new_iter.internal_struct = (
            interface_name[0].lower() + interface_name[1:]
        )
        new_iter.generated = True
        self._populate_status(new_iter)
        self._api.components[interface_name] = new_iter
        self._write(
            """
            type {internal_struct} struct {{
                obj *{parent_internal_struct}
                {internal_items_name} {field_type}
                fieldPtr *[]*{pb_pkg_name}.{field_external_struct}
            }}

            func new{interface}(ptr *[]*{pb_pkg_name}.{field_external_struct}) {interface} {{
                return &{internal_struct}{{fieldPtr: ptr}}
            }}

            type {interface} interface {{
                setMsg(*{parent_internal_struct}) {interface}
                Items() {field_type}
                Add() {field_external_struct}
                Append(items ...{field_external_struct}) {interface}
                Set(index int, newObj {field_external_struct}) {interface}
                Clear() {interface}
                clearHolderSlice() {interface}
                appendHolderSlice(item {field_external_struct}) {interface}
            }}

            func (obj *{internal_struct}) setMsg(msg *{parent_internal_struct}) {interface} {{
                obj.clearHolderSlice()
                for _, val := range *obj.fieldPtr {{
                    obj.appendHolderSlice(&{field_internal_struct}{{obj: val}})
                }}
                obj.obj = msg
                return obj
            }}

            func (obj *{internal_struct}) Items() {field_type} {{
                return obj.{internal_items_name}
            }}

            func (obj *{internal_struct}) Add() {field_external_struct} {{
                newObj := &{pb_pkg_name}.{field_external_struct}{{}}
                *obj.fieldPtr = append(*obj.fieldPtr, newObj)
                newLibObj := &{field_internal_struct}{{obj: newObj}}
                newLibObj.setDefault()
                obj.{internal_items_name} = append(obj.{internal_items_name}, newLibObj)
                return newLibObj
            }}

            func (obj *{internal_struct}) Append(items ...{field_external_struct}) {interface} {{
                for _, item := range items {{
                    newObj := item.msg()
                    *obj.fieldPtr = append(*obj.fieldPtr, newObj)
                    obj.{internal_items_name} = append(obj.{internal_items_name}, item)
                }}
                return obj
            }}

            func (obj *{internal_struct}) Set(index int, newObj {field_external_struct}) {interface} {{
                (*obj.fieldPtr)[index] = newObj.msg()
                obj.{internal_items_name}[index] = newObj
                return obj
            }}
            func (obj *{internal_struct}) Clear()  {interface} {{
                if len(*obj.fieldPtr) > 0 {{
                    *obj.fieldPtr = []*{pb_pkg_name}.{field_external_struct}{{}}
                    obj.{internal_items_name} = {field_type}{{}}
                }}
                return obj
            }}
            func (obj *{internal_struct}) clearHolderSlice() {interface} {{
                if len(obj.{internal_items_name}) > 0 {{
                    obj.{internal_items_name} = {field_type}{{}}
                }}
                return obj
            }}
            func (obj *{internal_struct}) appendHolderSlice(item {field_external_struct}) {interface} {{
                obj.{internal_items_name} = append(obj.{internal_items_name}, item)
                return obj
            }}
            """.format(
                internal_struct=new_iter.internal_struct,
                interface=new_iter.interface,
                field_internal_struct=field.struct,
                parent_internal_struct=new.struct,
                field_external_struct=field.external_struct,
                pb_pkg_name=self._protobuf_package_name,
                field_type=field.type,
                internal_items_name=self._get_holder_name(field, True),
            )
        )

    def _write_field_has(self, new, field):
        if field.has_method is None:
            return
        if field.name == "ResponseString":
            self._write(
                """
                {description}\n// {fieldname} returns a {fieldtype}
                func (obj *{struct}) Has{fieldname}() bool {{
                    return obj.obj.String_ != ""
                }}
                """.format(
                    fieldname=self._get_external_struct_name(field.name),
                    struct=new.struct,
                    description=field.description,
                    fieldtype=field.type,
                )
            )
        else:
            self._write(
                """
                {description}\n// {fieldname} returns a {fieldtype}
                func (obj *{struct}) Has{fieldname}() bool {{
                    return obj.obj.{internal_field_name} != nil
                }}
                """.format(
                    fieldname=self._get_external_struct_name(field.name),
                    struct=new.struct,
                    description=field.description,
                    fieldtype=field.type,
                    internal_field_name=field.name,
                )
            )

    def _build_setters_getters(self, fluent_new):
        """Add new FluentField objects for each interface field"""
        if "properties" not in fluent_new.schema_object:
            schema = self._get_parser("$..schema").find(
                fluent_new.schema_object
            )
            if len(schema) > 0:
                schema = schema[0].value
                schema_name = self._get_schema_object_name_from_ref(
                    schema["$ref"]
                )
                fluent_new.schema_object = {
                    "properties": {
                        self._get_external_struct_name(schema_name): schema,
                    },
                }
            else:
                return
        choice_enums = self._get_parser("$..choice..enum").find(
            fluent_new.schema_object["properties"]
        )
        prop_names = [
            key for key in fluent_new.schema_object["properties"].keys()
        ]
        for property_name, property_schema in fluent_new.schema_object[
            "properties"
        ].items():
            field = FluentField()
            field.schema = property_schema
            field.description = self._get_description(property_schema)
            field.name = self._get_external_field_name(property_name)
            field.type = self._get_struct_field_type(property_schema, field)

            if property_name == "status_code_default":
                continue

            if property_schema.get("x-status", {}).get("status") in [
                "deprecated",
                "under_review",
            ]:
                field.status = property_schema["x-status"]["status"].replace(
                    "-", "_"
                )
                field.status_msg = property_schema["x-status"].get(
                    "information"
                )

            # for x-enum properties we need go to into each x-enum and
            # retrieve x-status values from there
            enums = property_schema.get("x-enum")
            if enums is not None:
                for idx, (enum_name, enum_property) in enumerate(
                    enums.items()
                ):
                    x_status_info = self._get_x_status(
                        enum_property,
                        enum_name.upper(),
                        field.name,
                    )
                    if x_status_info is not None:
                        field.x_enum_status[idx + 1] = x_status_info

            # TODO: restore behavior
            # self._parse_x_constraints(field, property_schema)
            # self._parse_x_unique(field, property_schema)
            if (
                len(choice_enums) == 1
                and property_name in choice_enums[0].value
            ):
                field.setChoiceValue = property_name.upper()
            else:
                field.setChoiceValue = None
            field.isEnum = (
                len(self._get_parser("$..enum").find(property_schema)) > 0
            )
            field.isArray = (
                "type" in property_schema
                and property_schema["type"] == "array"
            )
            minmax_schema = property_schema
            if field.isArray:
                if "items" in property_schema:
                    minmax_schema = property_schema["items"]
                    if (
                        "minimum" not in minmax_schema
                        and "minimum" in property_schema
                    ):
                        minmax_schema["minimum"] = property_schema["minimum"]
                    if (
                        "maximum" not in minmax_schema
                        and "maximum" in property_schema
                    ):
                        minmax_schema["maximum"] = property_schema["maximum"]

            field.hasminmax = (
                "minimum" in minmax_schema or "maximum" in minmax_schema
            )
            field.hasminmaxlength = (
                "minLength" in minmax_schema or "maxLength" in minmax_schema
            )
            if field.isEnum:
                field.enums = (
                    self._get_parser("$..enum").find(property_schema)[0].value
                )
                if "unspecified" in field.enums:
                    field.enums.remove("unspecified")
                if property_name == "choice":
                    prop_names.remove("choice")
                    diff = set(field.enums).difference(set(prop_names))
                    if len(diff) > 0:
                        field.choice_with_no_prop = list(diff)

            if field.hasminmax:
                field.min = (
                    None
                    if "minimum" not in minmax_schema
                    else minmax_schema["minimum"]
                )
                field.max = (
                    None
                    if "maximum" not in minmax_schema
                    else minmax_schema["maximum"]
                )
            if field.hasminmaxlength:
                field.min_length = (
                    None
                    if "minLength" not in property_schema
                    else property_schema["minLength"]
                )
                field.max_length = (
                    None
                    if "maxLength" not in property_schema
                    else property_schema["maxLength"]
                )
            if fluent_new.isRpcResponse:
                if field.type == "[]byte":
                    field.name = "Bytes"
                elif "$ref" in property_schema:
                    schema_name = self._get_schema_object_name_from_ref(
                        property_schema["$ref"]
                    )
                    field.name = self._get_external_struct_name(schema_name)
            field.isOptional = fluent_new.isOptional(property_name)
            field.isPointer = not field.type.startswith("[")
            if field.isArray and field.isEnum:
                field.getter_method = (
                    "{fieldname}() []{interface}{fieldname}Enum".format(
                        fieldname=self._get_external_struct_name(field.name),
                        interface=fluent_new.interface,
                    )
                )
                field.getter_method_description = "{fieldname} returns []{interface}{fieldname}Enum, set in {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
            elif field.isEnum:
                field.getter_method = (
                    "{fieldname}() {interface}{fieldname}Enum".format(
                        fieldname=self._get_external_struct_name(field.name),
                        interface=fluent_new.interface,
                    )
                )
                field.getter_method_description = "{fieldname} returns {interface}{fieldname}Enum, set in {interface}".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
                )
            else:
                if field.name == "StatusCode_200":
                    if field.type == "[]byte":
                        field.name = "ResponseBytes"
                    elif field.type == "string":
                        field.name = "ResponseString"
                    elif "$ref" in property_schema:
                        schema_name = self._get_schema_object_name_from_ref(
                            property_schema["$ref"]
                        )
                        field.name = self._get_external_struct_name(
                            schema_name
                        )

                field.getter_method = "{name}() {ftype}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                )
                field.getter_method_description = (
                    "{name} returns {ftype}, set in {interface}.".format(
                        name=self._get_external_struct_name(field.name),
                        ftype=field.type,
                        interface=fluent_new.interface,
                    )
                )
                if field.type in self._api.components:
                    field.getter_method_description = (
                        field.getter_method_description
                        + """\n// {ftype} is {desc}""".format(
                            ftype=field.type,
                            desc=self._api.components[
                                field.type
                            ].description.lstrip("// "),
                        )
                    )

            if "$ref" in property_schema:
                schema_name = self._get_schema_object_name_from_ref(
                    property_schema["$ref"]
                )
                field.struct = self._get_internal_name(schema_name)
                field.external_struct = self._get_external_struct_name(
                    schema_name
                )
                field.setter_method = (
                    "Set{fieldname}(value {fieldstruct}) {interface}".format(
                        fieldname=self._get_external_struct_name(field.name),
                        fieldstruct=self._get_external_struct_name(
                            field.struct
                        ),
                        interface=fluent_new.interface,
                    )
                )
                field.setter_method_description = "Set{fieldname} assigns {fieldstruct} provided by user to {interface}.".format(
                    fieldname=self._get_external_struct_name(field.name),
                    fieldstruct=self._get_external_struct_name(field.struct),
                    interface=fluent_new.interface,
                )
                fieldstruct = self._get_external_struct_name(field.struct)
                if fieldstruct in self._api.components:
                    field.setter_method_description = (
                        field.setter_method_description
                        + """\n // {fieldstruct} is {desc}""".format(
                            fieldstruct=fieldstruct,
                            desc=self._api.components[
                                fieldstruct
                            ].description.lstrip("// "),
                        )
                    )
            if (
                field.isOptional
                and field.isPointer
                or "status_code" in property_name
            ):
                field.has_method = """Has{fieldname}() bool""".format(
                    fieldname=self._get_external_struct_name(field.name),
                )
                field.has_method_description = """Has{fieldname} checks if {fieldname} has been set in {interface}""".format(
                    fieldname=self._get_external_struct_name(field.name),
                    interface=fluent_new.interface,
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
                field.setter_method = (
                    "Set{name}(value {ftype}) {interface}".format(
                        name=self._get_external_struct_name(field.name),
                        ftype=field.type,
                        interface=fluent_new.interface,
                    )
                )
                field.setter_method_description = "Set{name} assigns {ftype} provided by user to {interface}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                    interface=fluent_new.interface,
                )
            elif field.isArray:
                field.isPointer = False
                if "$ref" in property_schema["items"]:
                    schema_name = self._get_schema_object_name_from_ref(
                        property_schema["items"]["$ref"]
                    )
                    field.isArray = True
                    field.struct = self._get_internal_name(schema_name)
                    field.external_struct = self._get_external_struct_name(
                        schema_name
                    )
                    field.iter_name = (
                        fluent_new.interface + field.external_struct + "Iter"
                    )
                    field.adder_method = "Add() {name}".format(
                        name=field.iter_name
                    )
                    field.isOptional = False
                    field.getter_method = "{name}() {iter_name}".format(
                        name=self._get_external_struct_name(field.name),
                        iter_name=field.iter_name,
                    )
                    field.getter_method_description = "{name} returns {iter_name}Iter, set in {parent}".format(
                        name=self._get_external_struct_name(field.name),
                        parent=fluent_new.interface,
                        iter_name=field.iter_name,
                    )
                else:
                    field.setter_method = (
                        "Set{name}(value {ftype}) {interface}".format(
                            name=self._get_external_struct_name(field.name),
                            ftype=field.type,
                            interface=fluent_new.interface,
                        )
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
                    print(
                        "Warning: Default should not accept for this property ",
                        property_name,
                    )
            if field.name.lower() == "auto":
                field.setter_method = None

            fluent_new.interface_fields.append(field)

    def _parse_x_constraints(self, field, schema):
        if "x-constraint" not in schema:
            return
        for con in schema["x-constraint"]:
            ref, prop = con.split("/properties/")
            ref = self._get_schema_object_name_from_ref(ref)
            prop = self._get_external_field_name(prop.strip("/"))
            field.x_constraints.append((self._get_internal_name(ref), prop))

    def _parse_x_unique(self, field, schema):
        if "x-unique" not in schema:
            return
        field.x_unique = schema["x-unique"]

    def _validate_x_constraint(self, field):
        body = ""
        if field.x_constraints == []:
            return body
        body = """
        xCons := []string{{
            {data}
        }}
        if !vObj.validateConstraint(xCons, obj.{name}()) {{
            vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s is not a valid {cons} type", obj.{name}()))
        }}
        """.format(
            data='"'
            + '", "'.join([".".join(c) for c in field.x_constraints])
            + '",',
            name=field.name,
            cons="|".join([".".join(c) for c in field.x_constraints]),
        )
        return body

    def _validate_unique(self, new, field):
        body = ""
        if field.x_unique is not None:
            body = """if !vObj.isUnique("{struct}", obj.{name}(), obj) {{
                vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("{name} with %s already exists", obj.{name}()))
            }}""".format(
                struct=new.struct, name=field.name
            )
        return body

    def _validate_types(self, new, field):
        body = ""
        if field.isPointer or "[]" in field.type:
            value = "nil"
        elif field.type == "string":
            value = '''""'''
        else:
            value = 0

        # The below code specifically raises warning for x-status in x-enums:
        if len(field.x_enum_status) > 0:
            validate_body = ""

            for enum, msg in field.x_enum_status.items():
                validate_body += """
                if obj.obj.{property}.Number() == {enum_number} {{
                    obj.addWarnings("{message}")
                }}
                """.format(
                    property=field.name,
                    enum_number=enum,
                    message=msg,
                )

            body = "{} \n {}".format(body, validate_body)

        # The below if deals with raising warning for x-status
        status_body = ""
        status_msg = ""
        if field.status is not None:
            status_msg = self._get_status_msg(
                field.name,
                field.status,
                field.status_msg,
                "obj",
                "property",
                new.schema_name,
            )

            status_body = """
            // {name} is {func}
            if obj.obj.{name}{enum} != {value} {{
                {msg}
            }}
            """.format(
                name=field.name,
                value=0 if field.isEnum and field.isArray is False else value,
                enum=".Number()"
                if field.isEnum and field.isArray is False
                else "",
                msg=status_msg,
                func=field.status,
            )

        if (
            field.isOptional is False
            and field.type in self._oapi_go_types.values()
        ):
            body = """
            // {name} is required
            if obj.obj.{name} == {value} {{
                vObj.validationErrors = append(vObj.validationErrors, "{name} is required field on interface {interface}")
            }} """.format(
                name=field.name,
                interface=new.interface,
                value="nil"
                if field.isEnum and field.isArray is False
                else value,
            )
            # TODO: restore behavior
            # unique = self._validate_unique(new, field)
            # body += "else " + unique if unique != "" else unique
        # TODO: restore behavior
        # if field.isOptional is True:
        #     body += self._validate_unique(new, field)
        # body += self._validate_x_constraint(field)
        inner_body = ""
        if field.hasminmax and ("int" in field.type or "float" in field.type):
            line = []
            if "int" in field.type:
                type_min, type_max = type_limits.limits.get(
                    field.type, (None, None)
                )
                if field.min is None and type_min is not None:
                    field.min = type_min
                if field.max is None and type_max is not None:
                    field.max = type_max
            if field.min is not None:
                if field.min == 0 and (
                    field.type.startswith("uint")
                    or field.type.startswith("[]uint")
                ):
                    pass
                else:
                    line.append("{pointer}{value} < {min}")
            if field.max is not None:
                line.append("{pointer}{value} > {max}")
            inner_body += (
                "if "
                + " || ".join(line)
                + """ {{
                    vObj.validationErrors = append(
                        vObj.validationErrors,
                        fmt.Sprintf("{min} <= {interface}.{name} <= {max} but Got {form}", {pointer}{value}))
                    }}
                """
            ).format(
                name=field.name,
                interface=new.interface,
                max="max({})".format(field.type.lstrip("[]"))
                if field.max is None
                else field.max,
                pointer="*" if field.isPointer else "",
                min="min({})".format(field.type.lstrip("[]"))
                if field.min is None
                else field.min,
                value="item"
                if field.isArray
                else "obj.obj.{name}".format(name=field.name),
                form="%d" if "int" in field.type else "%f",
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
                    vObj.validationErrors = append(
                        vObj.validationErrors,
                        fmt.Sprintf(
                            "{min_length} <= length of {interface}.{name} <= {max_length} but Got %d",
                            len({pointer}{value})))
                }}
                """
                ).format(
                    name=field.name,
                    interface=new.interface,
                    max_length="any"
                    if field.max_length is None
                    else field.max_length,
                    pointer="*" if field.isPointer else "",
                    min_length=field.min_length
                    if field.min_length is None
                    else field.min_length,
                    value="item"
                    if field.isArray
                    else "obj.obj.{name}".format(name=field.name),
                )
                if field.isArray:
                    inner_body = """
                        for _, item := range obj.obj.{name} {{
                            {body}
                        }}
                    """.format(
                        body=inner_body, name=field.name
                    )
            elif field.itemformat in [
                "mac",
                "ipv4",
                "ipv6",
                "hex",
                "oid",
            ] or field.format in ["mac", "ipv4", "ipv6", "hex", "oid"]:
                if field.format is None:
                    field.format = field.itemformat
                inner_body = """
                    err := obj.validate{format}(obj.{name}())
                    if err != nil {{
                        vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on {interface}.{name}"))
                    }}
                """.format(
                    name=field.name,
                    interface=new.interface,
                    format=field.format.capitalize()
                    if field.isArray is False
                    else field.format.capitalize() + "Slice",
                )

        # if there is no inner body then add status body or else
        # just the message would do
        if inner_body == "":
            if status_body != "":
                body = "{} \n {}".format(body, status_body)
            return body

        body += """
        if obj.obj.{name} != {value} {{
            {status}
            {body}
        }}
        """.format(
            name=field.name, value=value, body=inner_body, status=status_msg
        )
        return body

    def _validate_struct(self, new, field):
        body = ""
        if field.isOptional is False and field.isArray is False:
            body = """
                // {name} is required
                if obj.obj.{name} == nil {{
                    vObj.validationErrors = append(vObj.validationErrors, "{name} is required field on interface {interface}")
                }}
            """.format(
                name=field.name, interface=new.interface
            )

        inner_body = (
            "obj.{external_name}().validateObj(vObj, set_default)".format(
                external_name=self._get_external_struct_name(field.name)
            )
        )
        if field.isArray:
            inner_body = """
                 if set_default {{
                    obj.{name}().clearHolderSlice()
                    for _, item := range obj.obj.{name} {{
                        obj.{name}().appendHolderSlice(&{field_internal_struct}{{obj: item}})
                    }}
                 }}
                for _, item := range obj.{name}().Items() {{
                    item.validateObj(vObj, set_default)
                }}
            """.format(
                name=field.name,
                field_internal_struct=field.struct,
            )

        #  This part of code is for raising warning for x-status
        status_str = ""
        if field.status is not None:
            status_str = self._get_status_msg(
                field.name,
                field.status,
                field.status_msg,
                "obj",
                "property",
                new.schema_name,
            )
        body += """
            if {condition} {{
                {msg}
                {body}
            }}
        """.format(
            body=inner_body,
            condition="len(obj.obj.{name}) != 0".format(name=field.name)
            if field.isArray is True
            else "obj.obj.{name} != nil".format(name=field.name),
            msg=status_str,
        )
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

        status_str = ""
        if new.status is not None:
            status_str = self._get_status_msg(
                new.schema_name, new.status, new.status_info, "obj", "schema"
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
        if status_str != "":
            body = "\n%s\n%s" % (status_str, body)

        self._write(
            """func (obj *{struct}) validateObj(vObj *validation, set_default bool) {{
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
        enum_map = {}
        choice_enum_map = {}
        for index, field in enumerate(interface_fields):
            if field.name == "Choice":
                for enum in field.enums:
                    enum_str = self._get_external_struct_name(enum)
                    enum_map[enum_str] = ""
                    enum_value = """{struct}Choice.{value}""".format(
                        struct=self._get_external_struct_name(new.struct),
                        value=enum.upper(),
                    )
                    choice_enum_map[enum_str] = enum_value
                if field.default is not None:
                    choice_enums = [
                        self._get_external_struct_name(e)
                        for e in field.enums
                        if e != field.default
                    ]
                    hasChoiceConfig = [
                        "Choice",
                        self._get_external_struct_name(field.default),
                    ]
                interface_fields.insert(0, interface_fields.pop(index))
                break

        choice_body = None
        enum_fields = []
        for field in interface_fields:
            # if hasChoiceConfig != [] and field.name not in hasChoiceConfig:
            #     continue
            ext_name = self._get_external_struct_name(field.name)
            if field.name in enum_map or ext_name in enum_map:
                if field.name != ext_name and ext_name in enum_map:
                    del enum_map[ext_name]
                    enum_map[field.name] = ""
                    choice_enum_map[field.name] = choice_enum_map[ext_name]
                    del choice_enum_map[ext_name]
                type = ""
                if field.isEnum:
                    type = "enum"
                elif field.isPointer:
                    type = "pointer"
                else:
                    type = field.type
                enum_map[field.name] = type

            if (
                field.default is None
                or field.isOptional is False
                or field.name in choice_enums
            ):
                if field.name not in hasChoiceConfig:
                    continue
            if field.struct is not None:
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.{external_name}()".format(
                            external_name=self._get_external_struct_name(
                                field.name
                            )
                        )
                    )
                else:
                    body += """if obj.obj.{name} == nil {{
                        obj.{external_name}()
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(
                            field.name
                        ),
                        # enum_check="&& {}".format(enum_check) if enum_check is not None and field.name in choice_enums
                        # else ""
                    )
            elif field.isArray:
                if "string" in field.type:
                    values = (
                        '"{0}"'.format('", "'.join(field.default))
                        if field.default != []
                        else ""
                    )
                else:
                    values = str(field.default)[1:-1]
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({type}{{{values}}})".format(
                            external_name=self._get_external_struct_name(
                                field.name
                            ),
                            type=field.type,
                            values=values,
                        )
                    )
                else:
                    body += """if obj.obj.{name} == nil {{
                        obj.Set{external_name}({type}{{{values}}})
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(
                            field.name
                        ),
                        type=field.type,
                        values=values,
                    )
            elif field.isEnum:
                enum_value = """{struct}{name}.{value}""".format(
                    struct=self._get_external_struct_name(new.struct),
                    name=field.name,
                    value=field.default.upper(),
                )
                if field.isPointer:
                    cnd_check = """obj.obj.{name} == nil""".format(
                        name=field.name
                    )
                else:
                    cnd_check = """obj.obj.{name}.Number() == 0""".format(
                        name=field.name
                    )
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
                    choice_body = (
                        body1
                        if choice_body is None
                        else choice_body.replace("<choice_fields>", body1)
                    )
                else:
                    body = body + body1.replace("<choice_fields>", "")
            elif field.isPointer:
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({value})".format(
                            external_name=self._get_external_struct_name(
                                field.name
                            ),
                            value='"{0}"'.format(field.default)
                            if field.type == "string"
                            else field.default,
                        )
                    )
                else:
                    choice_cond = ""
                    if field.name in choice_enum_map:
                        choice_cond += (
                            "&& choice == %s " % choice_enum_map[field.name]
                        )
                    body += """if obj.obj.{name} == nil {choice_check}{{
                        obj.Set{external_name}({value})
                    }}
                    """.format(
                        name=field.name,
                        external_name=self._get_external_struct_name(
                            field.name
                        ),
                        value='"{0}"'.format(field.default)
                        if field.type == "string"
                        else field.default,
                        choice_check=choice_cond,
                    )
            else:
                if field.name in hasChoiceConfig:
                    enum_fields.append(
                        "obj.Set{external_name}({value})".format(
                            external_name=self._get_external_struct_name(
                                field.name
                            ),
                            value='"{0}"'.format(field.default)
                            if field.type == "string"
                            else field.default,
                        )
                    )
                else:
                    body += """if obj.obj.{name} == {check_value} {{
                        obj.Set{external_name}({value})
                    }}
                    """.format(
                        name=field.name,
                        check_value='""' if field.type == "string" else "0",
                        external_name=self._get_external_struct_name(
                            field.name
                        ),
                        value='"{0}"'.format(field.default)
                        if field.type == "string"
                        else field.default,
                    )

        # write default case if object has choice property
        choice_code = ""
        # TODO: we need to propagate error from setdefault along the whole heirarchy
        if len(enum_map) > 0:
            choice_code = (
                "var choices_set int = 0\nvar choice %sChoiceEnum\n"
                % new.interface
            )
            for enum in enum_map:
                field_type = enum_map[enum]
                value = "nil"
                if field_type.startswith("int") or field_type.startswith(
                    "float"
                ):
                    value = "0"
                elif field_type == "string":
                    value = '""'
                elif field_type == "pointer":
                    value = "nil"
                elif field_type == "":
                    # signifies choice with no property
                    continue

                if field_type.startswith("[]"):
                    choice_code += """
                    if len(obj.obj.{prop}) > 0{{
                        choices_set += 1
                        choice = {choice_val}
                    }}
                    """.format(
                        prop=enum,
                        choice_val=choice_enum_map[enum],
                    )
                else:
                    enum_check_code = " && obj.obj.%s.Number() != 0 " % enum
                    choice_code += """
                    if obj.obj.{prop} != {val}{enum_check}{{
                        choices_set += 1
                        choice = {choice_val}
                    }}
                    """.format(
                        prop=enum,
                        val=value,
                        choice_val=choice_enum_map[enum],
                        enum_check=enum_check_code
                        if field_type == "enum"
                        else "",
                    )

            # TODO: we need to throw error if more that one choice properties are set
            # choice_code += """
            # if choices_set > 1 {{
            #     obj.validationErrors = append(obj.validationErrors, "more than one choices are set in Interface {intf}")
            # }}""".format(
            #     intf=new.interface
            # )

            if choice_body is not None:
                choice_code += """if choices_set == 0 {{
                    {body}
                }}""".format(
                    body=choice_body.replace("<choice_fields>", "")
                )
                # enum_fields = []
                # body = (
                #     choice_body.replace(
                #         "<choice_fields>",
                #         "\n".join(enum_fields) if enum_fields != [] else "",
                #     )
                #     + body
                # )

            choice_code += """{el}if choices_set == 1 && choice != "" {{
                if obj.obj.Choice != nil {{
                    if obj.Choice() != choice {{
                        obj.validationErrors = append(obj.validationErrors, "choice not matching with property in {intf}")
                    }}
                }} else {{
                    intVal := {pb_pkg_name}.{intf}_Choice_Enum_value[string(choice)]
                    enumValue := {pb_pkg_name}.{intf}_Choice_Enum(intVal)
                    obj.obj.Choice = &enumValue
                }}
            }}
            """.format(
                intf=new.interface,
                pb_pkg_name=self._protobuf_package_name,
                el=" else " if choice_body is not None else "",
            )

            body = choice_code + "\n" + body

        body = body.replace("SetChoice", "setChoice")
        self._write(
            """func (obj *{struct}) setDefault() {{
                {body}
            }}""".format(
                struct=new.struct, body=body
            )
        )

    def _write_value_of(self, new):
        body = ""
        for field in new.interface_fields:
            body += """
            if name == "{name}" {{
                return obj.obj.{name}
            }}
            """.format(
                name=field.name
            )
        self._write(
            """func (obj *{struct}) ValueOf(name string) interface{{}} {{
                {body}
                return nil
            }}
            """.format(
                body=body, struct=new.struct
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

    def _get_struct_field_type(
        self, property_schema, fluent_field=None, min=None, max=None
    ):
        """Convert openapi type, format, items, $ref keywords to a go type"""
        go_type = ""
        if "type" in property_schema:
            oapi_type = property_schema["type"]
            if oapi_type.lower() in self._oapi_go_types:
                if property_schema["type"] == "integer":
                    go_type = type_limits._get_integer_format(
                        property_schema.get("format"),
                        property_schema.get("minimum", min),
                        property_schema.get("maximum", max),
                    )
                else:
                    go_type = "{oapi_go_types}".format(
                        oapi_go_types=self._oapi_go_types[oapi_type.lower()]
                    )
            if oapi_type == "array":
                go_type += "[]" + self._get_struct_field_type(
                    property_schema["items"],
                    fluent_field,
                    property_schema.get("minimum"),
                    property_schema.get("maximum"),
                ).replace("*", "")
                if "format" in property_schema["items"]:
                    fluent_field.itemformat = property_schema["items"][
                        "format"
                    ]
            if "format" in property_schema:
                type_format = (oapi_type + property_schema["format"]).lower()
                if type_format.lower() in self._oapi_go_types:
                    go_type = "{oapi_go_type}".format(
                        oapi_go_type=self._oapi_go_types[type_format.lower()]
                    )
                elif (
                    property_schema["format"].lower()
                    not in self._oapi_go_types
                ):
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
                new.interface = self._get_external_struct_name(
                    schema_object_name
                )
                new.description = self._get_description(schema_object, True)
                self._populate_status(new)
                self._api.components[new.schema_name] = new
            go_type = new.interface
        else:
            raise Exception(
                "No type or $ref keyword present in property schema: {property_schema}".format(
                    property_schema=property_schema
                )
            )
        return go_type

    def _get_description(self, openapi_object, noCap=False):
        description = "// description is TBD"
        if "description" in openapi_object:
            description = ""
            for ind, line in enumerate(
                openapi_object["description"].split("\n")
            ):
                if noCap and ind == 0 and line != "":
                    line = line[0].lower() + line[1:]
                description += "// {line}\n".format(line=line.strip())
        return description.strip("\n")

    def _populate_status(self, new_fluent):
        if new_fluent is None:
            return

        if new_fluent.schema_object is not None:
            if "x-status" in new_fluent.schema_object:
                status_val = new_fluent.schema_object["x-status"]
                new_fluent.status = status_val["status"]
                new_fluent.status_info = status_val["information"]

    def _get_holder_name(self, field, isIter=False):
        if isIter:
            return "{}Slice".format(field.struct)
        return "{}Holder".format(field.name[0].lower() + field.name[1:])

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
            print(
                "Tidying the generated go mod file: {}".format(
                    " ".join(process_args)
                )
            )
            process = subprocess.Popen(
                process_args, cwd=self._ux_path, shell=False, env=os.environ
            )
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))

    def _get_status_msg(
        self,
        name,
        status_type,
        status_msg,
        prefix,
        property_type=None,
        parent_schema=None,
    ):
        """
        This function basically returns the warning message for x-status
        """
        msg = ""
        if status_type == "deprecated":
            msg = "is deprecated"
        elif status_type == "under_review":
            msg = "is under review"
        else:
            raise NotImplementedError(
                "%s status is not implemented" % status_type
            )

        if property_type == "property":
            msg = "%s property in schema %s %s" % (name, parent_schema, msg)
        elif property_type == "schema":
            msg = "%s %s" % (name, msg)

        initial = ""
        if prefix == "api":
            initial = prefix
            msg = "%s api %s" % (name, msg)
        elif prefix == "x-enum":
            return "%s enum in property %s %s, %s" % (
                name,
                parent_schema,
                msg,
                status_msg,
            )
        else:
            initial = "obj"

        msg = '%s.addWarnings("%s, %s")' % (
            initial,
            msg,
            status_msg,
        )
        return msg

    def _get_x_status(self, enum_schema, enum_name=None, property_name=None):
        if enum_schema.get("x-status", {}).get("status") in [
            "deprecated",
            "under_review",
        ]:
            status = enum_schema["x-status"]["status"].replace("-", "_")
            status_msg = enum_schema["x-status"].get("information")
            if enum_name is not None:
                status_msg = self._get_status_msg(
                    enum_name,
                    status,
                    status_msg,
                    "x-enum",
                    parent_schema=property_name,
                )
            return status_msg

    def _handle_response_fields(self, field, property_name, property_schema):
        pass
