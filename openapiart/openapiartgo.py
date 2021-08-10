from .openapiartplugin import OpenApiArtPlugin
import os
import subprocess


class FluentStructure(object):
    def __init__(self):
        self.internal_struct_name = None
        self.external_interface_name = None
        self.external_new_methods = []
        self.external_rpc_methods = []
        self.components = {}


class FluentRpc(object):
    """SetConfig(config Config) error"""

    def __init__(self):
        self.operation_name = None
        self.method = None
        self.request = None


class FluentNew(object):
    """New<external_interface_name> <external_interface_name>"""

    def __init__(self):
        self.generated = False
        self.interface = None
        self.struct = None
        self.method = None
        self.schema_name = None
        self.schema_object = None
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
        self.type = None
        self.isPointer = True
        self.struct = None
        self.external_struct = None
        self.setter_method = None
        self.getter_method = None
        self.adder_method = None


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
            "number": "float32",
            "numberfloat": "float32",
            "numberdouble": "float64",
            # "stringmac": "StringMac",
            # "stringipv4": "StringIpv4",
            # "stringipv6": "StringIpv6",
            # "stringhex": "StringHex",
            "stringbinary": "[]byte",
        }

    def generate(self, openapi):
        self._openapi = openapi
        os.mkdir(os.path.normpath(os.path.join(self._output_dir, self._go_module_name)))
        self._structs = {}
        self._write_mod_file()
        self._write_go_file()
        self._format_go_file()
        self._tidy_mod_file()

    def _write_mod_file(self):
        self._filename = os.path.normpath(os.path.join(self._output_dir, self._go_module_name, "go.mod"))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write("module {}".format(self._go_module_name))
        self._write()
        self._write("go 1.16")
        self._close_fp()

    def _write_go_file(self):
        self._filename = os.path.normpath(os.path.join(self._output_dir, self._go_module_name, "{}.go".format(self._go_module_name)))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write_package_docstring(self._openapi["info"])
        self._write_package()
        self._write_common_code()
        self._write_types()
        self._build_api_interface()
        self._build_request_interfaces()
        self._build_component_interfaces()
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
        self._write(f"package {self._go_module_name}")
        self._write()

    def _write_common_code(self):
        """Writes the base wrapper go code"""
        # with open(os.path.join(os.path.dirname(__file__), "common.go")) as fp:
        #     self._write(fp.read().strip().strip("\n"))
        # self._write()
        self._write(f'''import {self._protobuf_package_name} "../{self._protobuf_package_name}"''')
        self._write('import "google.golang.org/grpc"')
        self._write(
            r"""
        import (
            "encoding/json"
            "fmt"
            "time"

            "gopkg.in/yaml.v3"
        )

        type ApiTransportEnum string

        var ApiTransport = struct {
            GRPC ApiTransportEnum
            HTTP ApiTransportEnum
        }{
            GRPC: "grpc",
            HTTP: "http",
        }

        type api struct {
            transport          ApiTransportEnum
            grpcLocation       string
            grpcRequestTimeout time.Duration
        }

        type Api interface {
            SetTransport(value ApiTransportEnum) *api
            SetGrpcLocation(value string) *api
            SetGrpcRequestTimeout(value time.Duration) *api
        }

        // Transport returns the active transport
        func (api *api) Transport() string {
            return string(api.transport)
        }

        // SetTransport sets the active type of transport to be used
        func (api *api) SetTransport(value ApiTransportEnum) *api {
            api.transport = value
            return api
        }

        func (api *api) GrpcLocation() string {
            return api.grpcLocation
        }

        // SetGrpcLocation
        func (api *api) SetGrpcLocation(value string) *api {
            api.grpcLocation = value
            return api
        }

        func (api *api) GrpcRequestTimeout() time.Duration {
            return api.grpcRequestTimeout
        }

        // SetGrpcRequestTimeout contains the timeout value in seconds for a grpc request
        func (api *api) SetGrpcRequestTimeout(value int) *api {
            api.grpcRequestTimeout = time.Duration(value) * time.Second
            return api
        }

        // All methods that perform validation will add errors here
        // All api rpcs MUST call Validate
        var validation []string

        func Validate() {
            if len(validation) > 0 {
                for _, item := range validation {
                    fmt.Println(item)
                }
                panic("validation errors")
            }
        }
        """
        )

    def _write_types(self):
        for _, go_type in self._oapi_go_types.items():
            if go_type.startswith("String"):
                self._write(f"type {go_type} string")
        self._write()

    def _get_internal_name(self, openapi_name):
        return openapi_name[0].lower() + openapi_name[1:].replace("_", "").replace(".", "")

    def _get_external_name(self, openapi_name):
        pieces = openapi_name.replace(".", "").split("_")
        external_name = ""
        for piece in pieces:
            external_name += piece[0].upper()
            if len(external_name) > 0:
                external_name += piece[1:]
        if external_name in ["String"]:
            external_name += "_"
        return external_name

    def _get_external_field_name(self, openapi_name):
        return self._get_external_name(openapi_name) + "_"

    def _build_api_interface(self):
        self._api.internal_struct_name = f"""{self._get_internal_name(self._go_module_name)}Api"""
        self._api.external_interface_name = f"""{self._get_external_name(self._go_module_name)}Api"""
        for _, path_object in self._openapi["paths"].items():
            for _, path_item_object in path_object.items():
                ref = self._get_parser("$..requestBody..'$ref'").find(path_item_object)
                if len(ref) == 1:
                    new = FluentNew()
                    new.schema_name = self._get_schema_object_name_from_ref(ref[0].value)
                    new.schema_object = self._get_schema_object_from_ref(ref[0].value)
                    new.interface = self._get_external_name(new.schema_name)
                    new.struct = self._get_internal_name(new.schema_name)
                    new.method = f"""New{new.interface}() {new.interface}"""
                    self._api.external_new_methods.append(new)
                    rpc = FluentRpc()
                    rpc.operation_name = self._get_external_name(path_item_object["operationId"])
                    rpc.method = f"""{rpc.operation_name}({new.struct} {new.interface}) error"""
                    rpc.request = f"""{self._protobuf_package_name}.{rpc.operation_name}Request{{{new.interface}: {new.struct}.msg()}}"""
                    self._api.external_rpc_methods.append(rpc)

        # write the go code
        self._write(
            f"""type {self._api.internal_struct_name} struct {{
                api
                grpcClient {self._protobuf_package_name}.OpenapiClient
            }}

            // grpcConnect builds up a grpc connection
            func (api *{self._api.internal_struct_name}) grpcConnect() error {{
                if api.grpcClient == nil {{
                    conn, err := grpc.Dial(api.grpcLocation, grpc.WithInsecure())
                    if err != nil {{
                        return err
                    }}
                    api.grpcClient = {self._protobuf_package_name}.NewOpenapiClient(conn)
                }}
                return nil
            }}

            // NewApi returns a new instance of the top level interface hierarchy
            func NewApi() *{self._api.internal_struct_name} {{
                api := {self._api.internal_struct_name}{{}}
                api.transport = ApiTransport.GRPC
                api.grpcLocation = "127.0.0.1:5050"
                api.grpcRequestTimeout = 10 * time.Second
                api.grpcClient = nil
                return &api
            }}
            """
        )
        methods = []
        for new in set(self._api.external_new_methods):
            methods.append(new.method)
        for rpc in set(self._api.external_rpc_methods):
            methods.append(rpc.method)
        method_signatures = "\n".join(methods)
        self._write(
            f"""type {self._api.external_interface_name} interface {{
                Api
                {method_signatures}
            }}
            """
        )
        for new in self._api.external_new_methods:
            self._write(
                f"""func (api *{self._api.internal_struct_name}) {new.method} {{
                    return &{new.struct}{{obj: &{self._protobuf_package_name}.{new.interface}{{}}}}
                }}
                """
            )
        for rpc in self._api.external_rpc_methods:
            self._write(
                f"""func (api *{self._api.internal_struct_name}) {rpc.method} {{
                    if err := api.grpcConnect(); err != nil {{
                        return err
                    }}
                    request := {rpc.request}
                    ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpcRequestTimeout)
                    defer cancelFunc()
                    client, err := api.grpcClient.{rpc.operation_name}(ctx, &request)
                    if err != nil {{
                        return err
                    }}
                    resp, _ := client.Recv()
                    if resp.GetStatusCode_200() == nil {{
                        return fmt.Errorf("fail")
                    }}
                    return nil
                }}
                """
            )

    def _build_request_interfaces(self):
        for new in self._api.external_new_methods:
            self._build_interface(new)

    def _build_component_interfaces(self):
        while True:
            components = [component for _, component in self._api.components.items() if component.generated is False]
            if len(components) == 0:
                break
            for component in components:
                self._build_interface(component)

    def _build_interface(self, new):
        self._write(
            f"""type {new.struct} struct {{
                obj *{self._protobuf_package_name}.{new.interface}
            }}
            
            func (obj *{new.struct}) msg() *{self._protobuf_package_name}.{new.interface} {{
                return obj.obj
            }}

            func (obj *{new.struct}) Yaml() string {{
                data, _ := yaml.Marshal(obj.msg())
                return string(data)
            }}

            func (obj *{new.struct}) Json() string {{
                data, _ := json.Marshal(obj.msg())
                return string(data)
            }}
        """
        )
        self._build_setters_getters(new)
        interfaces = ["Yaml() string", "Json() string"]
        for field in new.interface_fields:
            interfaces.append(field.getter_method)
            if field.setter_method is not None:
                interfaces.append(field.setter_method)
            if field.adder_method is not None:
                interfaces.append(field.adder_method)
        interface_signatures = "\n".join(interfaces)
        self._write(
            f"""type {new.interface} interface {{
                msg() *{self._protobuf_package_name}.{new.interface}
                {interface_signatures}
            }}
        """
        )
        for field in new.interface_fields:
            self._write_field_getter(new, field)
            self._write_field_setter(new, field)
            self._write_field_adder(new, field)
        new.generated = True

    def _write_field_getter(self, new, field):
        if field.getter_method is None:
            return
        if field.struct is not None:
            body = f"""if obj.obj.{field.name} == nil {{
                    obj.obj.{field.name} = &{self._protobuf_package_name}.{field.external_struct}{{}}
                }}
                return &{field.struct}{{obj: obj.obj.{field.name}}}
            """
        elif field.isPointer:
            body = f"""return *obj.obj.{field.name}"""
        else:
            body = f"""return obj.obj.{field.name}"""
        self._write(
            f"""func (obj *{new.struct}) {field.getter_method} {{
                {body}
            }}
            """
        )

    def _write_field_setter(self, new, field):
        if field.setter_method is None:
            return
        if field.isPointer:
            body = f"""obj.obj.{field.name} = &value"""
        else:
            body = f"""obj.obj.{field.name} = value"""
        self._write(
            f"""func (obj *{new.struct}) {field.setter_method} {{
                {body}
                return obj
            }}
            """
        )

    def _write_field_adder(self, new, field):
        if field.adder_method is None:
            return
        self._write(
            f"""func (obj *{new.struct}) {field.adder_method} {{
                if obj.obj.{field.name} == nil {{
                    obj.obj.{field.name} = &{field.type}{{}}
                }}
                slice := append(*obj.obj.{field.name}, value)
                obj.obj.{field.name} = &slice
                return &slice[len(slice)-1]
            }}
            """
        )

    def _build_setters_getters(self, fluent_new):
        """Add new FluentField objects for each interface field"""
        if "properties" not in fluent_new.schema_object:
            return
        for property_name, property_schema in fluent_new.schema_object["properties"].items():
            if len(self._get_parser("$..enum").find(property_schema)) > 0:  # temporary
                continue
            field = FluentField()
            field.name = self._get_external_name(property_name)
            field.type = self._get_struct_field_type(property_schema)
            if field.type.startswith("["):
                # field pointer cannot be done on array(slice)
                field.isPointer = False
            else:
                field.isPointer = fluent_new.isOptional(property_name)
            field.getter_method = f"{field.name}() {field.type}"
            if field.type not in self._oapi_go_types.values() and "$ref" not in property_schema:
                continue
            if "$ref" in property_schema:
                schema_name = self._get_schema_object_name_from_ref(property_schema["$ref"])
                field.struct = self._get_internal_name(schema_name)
                field.external_struct = self._get_external_name(schema_name)
            if field.type in self._oapi_go_types.values():
                field.setter_method = f"Set{field.name}(value {field.type}) {fluent_new.interface}"
            elif field.type.startswith("["):
                if field.type.split("]")[-1] in self._oapi_go_types.values():
                    field.adder_method = f"Add{field.name}(value {field.type}) {fluent_new.interface}"
                else:
                    field.adder_method = f"Add{field.name}() {field.type[2:]}"
            fluent_new.interface_fields.append(field)

    def _get_schema_object_name_from_ref(self, ref):
        final_piece = ref.split("/")[-1]
        return final_piece.replace(".", "")

    def _get_schema_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf

    def _get_struct_field_type(self, property_schema, field=None):
        """Convert openapi type, format, items, $ref keywords to a go type"""
        go_type = ""
        if "type" in property_schema:
            oapi_type = property_schema["type"]
            if oapi_type.lower() in self._oapi_go_types:
                go_type = f"{self._oapi_go_types[oapi_type.lower()]}"
            if oapi_type == "array":
                go_type += f"[]" + self._get_struct_field_type(property_schema["items"]).replace("*", "")
            if "format" in property_schema:
                format_type = (oapi_type + property_schema["format"]).lower()
                if format_type.lower() in self._oapi_go_types:
                    go_type = f"{self._oapi_go_types[format_type.lower()]}"
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
                new.interface = self._get_external_name(schema_object_name)
                self._api.components[new.schema_name] = new
            go_type = new.interface
        else:
            raise Exception(f"No type or $ref keyword present in property schema: {property_schema}")
        return go_type

    def _get_description(self, function_name, openapi_object):
        description = f"// {function_name} TBD"
        if "description" in openapi_object:
            description = function_name + "\n" + openapi_object["description"].strip("\n")
            description = "/* {}\n*/".format(description)
        return description

    def _format_go_file(self):
        """Format the generated go code"""
        try:
            process_args = [
                "goimports",
                "-w",
                self._filename,
            ]
            print("Formatting generated go ux file: {}".format(" ".join(process_args)))
            process = subprocess.Popen(process_args, shell=True)
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
            self._mod_dir = os.path.normpath(os.path.join(self._output_dir, self._go_module_name))
            os.environ["GO111MODULE"] = "on"
            print("Tidying the generated go mod file: {}".format(" ".join(process_args)))
            process = subprocess.Popen(process_args, cwd=self._mod_dir, shell=True, env=os.environ)
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))
