from .openapiartplugin import OpenApiArtPlugin
import os
import subprocess


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
        self._api_interface_methods = []
        self._oapi_go_types = {
            "string": "string",
            "boolean": "bool",
            "integer": "int32",
            "number": "float32",
            "numberfloat": "float32",
            "numberdouble": "float64",
            "stringmac": "StringMac",
            "stringipv4": "StringIpv4",
            "stringipv6": "StringIpv6",
            "stringhex": "StringHex",
            "stringbinary": "[]byte",
        }

    def generate(self, openapi):
        self._openapi = openapi
        self._structs = {}
        self._write_mod_file()
        self._write_go_file()
        self._format_go_file()
        self._tidy_mod_file()

    def _write_mod_file(self):
        self._filename = os.path.normpath(os.path.join(self._output_dir, self._python_module_name, "go.mod"))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write("module {}".format(self._go_module_name))
        self._write()
        self._write("go 1.16")
        self._close_fp()

    def _write_go_file(self):
        self._filename = os.path.normpath(os.path.join(self._output_dir, self._python_module_name, "{}.go".format(self._python_module_name)))
        self.default_indent = "    "
        self._init_fp(self._filename)
        self._write_package_docstring(self._openapi["info"])
        self._write_package()
        self._write_common_code()
        self._write_api_interface()
        # self._write_response_msg(path_object)
        # for name, schema_object in self._openapi["components"]["schemas"].items():
        #     self._write_struct(name, schema_object)
        # for name, response_object in self._openapi["components"]["responses"].items():
        #     self._write_msg(name, response_object)
        # self._write_types()
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
        self._write(f'''import {self._protobuf_package_name} "."''')
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

        func MessageToJson(obj interface{}) string {
            data, _ := json.Marshal(obj)
            return string(data)
        }

        func MessageToYaml(obj interface{}) string {
            data, _ := yaml.Marshal(obj)
            return string(data)
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
        return external_name

    def _get_external_field_name(self, openapi_name):
        return self._get_external_name(openapi_name) + "_"

    def _write_api_interface(self):
        self._internal_api_struct_name = f"""{self._get_internal_name(self._go_module_name)}Api"""
        self._write(
            f"""type {self._internal_api_struct_name} struct {{
                api
                grpcClient {self._protobuf_package_name}.OpenapiClient
            }}

            // grpcConnect builds up a grpc connection
            func (api *{self._internal_api_struct_name}) grpcConnect() error {{
                if api.grpcClient == nil {{
                    conn, err := grpc.Dial(api.grpcLocation, grpc.WithInsecure())
                    if err != nil {{
                        return err
                    }}
                    api.grpcClient = {self._protobuf_package_name}.NewOpenapiClient(conn)
                }}
                return nil
            }}
            func NewApi() *{self._internal_api_struct_name} {{
                api := {self._internal_api_struct_name}{{}}
                api.transport = ApiTransport.GRPC
                api.grpcLocation = "127.0.0.1:5050"
                api.grpcRequestTimeout = 10 * time.Second
                api.grpcClient = nil
                return &api
            }}
            """
        )
        api_interface_methods = []
        delayed_writes = []
        child_refs = []
        for _, path_object in self._openapi["paths"].items():
            self._write_request_struct(path_object, api_interface_methods, delayed_writes, child_refs)
        methods = "\n".join(api_interface_methods)
        self._external_api_interface_name = f"""{self._get_external_name(self._go_module_name)}Api"""
        self._write(
            f"""type {self._external_api_interface_name} interface {{
                Api
                {methods}
            }}
            """
        )
        for item in delayed_writes:
            self._write(item)
        for ref in child_refs:
            delayed_writes = []
            schema_object = self._get_schema_object_from_ref(ref)
            self._write_struct_and_interface(schema_object, delayed_writes)
            for delay in delayed_writes:
                self._write(delay)

    def _write_request_struct(self, path_object, api_interface_methods, delayed_writes, child_refs):
        """Write the operation struct and func, add the func signature to api_interfaces,
        return the api_interfaces
        type <InternalStructName> struct {
            <self._protobuf_package_name>.<ExternalStructName>
        }
        func (api *api) <operationid>(request *<structname>) response *structname {
        }
        """
        for _, path_item_object in path_object.items():
            operation_id = path_item_object["operationId"]
            ref = self._get_parser("$..requestBody..'$ref'").find(path_item_object)
            if len(ref) == 1:
                schema_object_name = self._get_schema_object_name_from_ref(ref[0].value)
                internal_struct_name = self._get_internal_name(schema_object_name)
                external_struct_name = self._get_external_name(schema_object_name)
                delayed_writes.append(
                    f"""type {internal_struct_name} struct {{
                        {self._protobuf_package_name}.{external_struct_name}
                    }}
                """
                )
                # this writes the interface which needs to contain all the properties
                # foreach property generate the struct/interface and add it to delayed writes
                # TBD: have write_struct_and_interface add any refs that have been generated to delayed_refs
                schema_object = self._get_schema_object_from_ref(ref[0].value)
                interface_methods = self._write_struct_and_interface(schema_object, delayed_writes, child_refs)
                interfaces = "\n".join(interface_methods)
                delayed_writes.append(
                    f"""type {external_struct_name} interface {{
                        msg() *{self._protobuf_package_name}.{external_struct_name}
                        {interfaces}
                    }}
                """
                )
                delayed_writes.append(
                    f"""func (obj *{internal_struct_name}) msg() *{self._protobuf_package_name}.{external_struct_name} {{
                        return &obj.{external_struct_name}
                    }}
                """
                )
                external_operation_name = self._get_external_name(operation_id)
                external_method = f"""{external_operation_name}({internal_struct_name} {external_struct_name}) error"""
                delayed_writes.append(
                    f"""func (api *{self._go_module_name}Api) {external_method} {{
                        if err := api.grpcConnect(); err != nil {{
                            return err
                        }}
                        request := {self._protobuf_package_name}.{external_operation_name}Request{{{external_struct_name}: {internal_struct_name}.msg()}}
                        ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpcRequestTimeout)
                        defer cancelFunc()
                        client, err := api.grpcClient.{external_operation_name}(ctx, &request)
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
                api_interface_methods.append(f"New{external_struct_name}() {external_struct_name}")
                api_interface_methods.append(external_method)

    def _write_struct_and_interface(self, schema_object, delayed_writes, child_refs):
        """Return"""
        interface_methods = []
        for property_name, property_schema in schema_object["properties"].items():
            external_struct_name = self._get_external_name(property_name)
            interface_methods.append(f"{external_struct_name}() string")
            interface_methods.append(f"Set{external_struct_name}(value string)")
        return interface_methods

    def _get_schema_object_name_from_ref(self, ref):
        final_piece = ref.split("/")[-1]
        if "." in final_piece:
            return final_piece.split(".")[-1]
        else:
            return final_piece

    def _get_schema_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf

    def _write_struct(self, schema_name, schema_object):
        """Write a public type <external structname> interface for the schema_object
        and write a private type <structname> struct
        Transform schema_name into the structname
        The structname should be lower case, remove underscores, periods

        All structs should be accessible via the parent
        Need to identify the parent for any schema object in order to generate
        func (parent *<parentstructname>) <Structname>() *structname {
            return &<structname>{}
        }
        """
        struct_name = self._get_internal_name(schema_name)
        self._write("type %s struct {" % struct_name)
        self._write_struct_fields(schema_object)
        self._write("}")
        self._write()
        interface_name = self._get_external_name(schema_name)
        self._write("type %s interface {" % interface_name)
        self._write_interface_fields(interface_name, schema_object)
        self._write("}")
        self._write()
        self._write_struct_field_getters(struct_name, schema_object)
        self._write_struct_field_setters(struct_name, schema_object)
        self._write_struct_field_adders(struct_name, schema_object)
        self._structs[struct_name] = schema_name

    def _write_struct_fields(self, schema_object):
        """Write the fields for the <structname> struct
        The fieldnames should be lower case
        """
        if "properties" not in schema_object:
            return
        for name, property in schema_object["properties"].items():
            field_name = self._get_external_field_name(name)
            required = False
            if "required" in schema_object and name in schema_object["required"]:
                required = True
            field_type = self._get_struct_field_type(property, required)
            line = f'{field_name} {field_type} `proto:"{name}" json:"{name}" yaml:"{name}"`'
            self._write(line, 1)

    def _write_interface_fields(self, interface_name, schema_object):
        """Write the fields for the <structname> interface
        Name() <fieldtype>
        SetName(value <fieldtype>) <interfacename>
        """
        if "properties" not in schema_object:
            return
        for name, property in schema_object["properties"].items():
            field_name = self._get_external_name(name)
            required = False
            if "required" in schema_object and name in schema_object["required"]:
                required = True
            field_type = self._get_struct_field_type(property, required)
            self._write(f"{field_name}() {field_type}", 1)
            self._write(f"Set{field_name}(value {field_type}) {interface_name}", 1)

    def _get_struct_field_type(self, property_schema, required=False):
        """Convert openapi type, format, items, $ref keywords to a go type"""
        pointer = "" if required else "*"
        go_type = ""
        if "type" in property_schema:
            oapi_type = property_schema["type"]
            if oapi_type == "array":
                go_type += f"{pointer}[]" + self._get_struct_field_type(property_schema["items"]).replace("*", "")
            else:
                if "format" in property_schema:
                    oapi_type += property_schema["format"]
                go_type = f"{pointer}{self._oapi_go_types[oapi_type.lower()]}"
        elif "$ref" in property_schema:
            ref = property_schema["$ref"].split("/")[-1]
            go_type = f"{pointer}{self._get_internal_name(ref)}"
        else:
            raise Exception(f"No type or $ref keyword present in property schema: {property_schema}")
        return go_type

    def _write_struct_field_getters(self, struct_name, schema_object):
        """Write the getters for each <structname> field

        Getter content should be of the form:

        func (obj *<structname>) <Fieldname> <fieldtype> {
            return obj.<fieldname>
        }
        """
        if "properties" not in schema_object:
            return
        for name, property in schema_object["properties"].items():
            required = "required" in schema_object and name in schema_object["required"]
            field_type = self._get_struct_field_type(property, required)
            field_name = self._get_external_field_name(name)
            function_name = self._get_external_name(name)
            self._write(self._get_description(function_name, property))
            self._write(f"func (obj *{struct_name}) {function_name}() {field_type} {{")
            self._write(f"return obj.{field_name}", 1)
            self._write("}")

    def _write_struct_field_setters(self, struct_name, schema_object):
        """Write the setters/getters for each <structname> field
        The fieldnames should start with an upper case
        func (obj *<structname>) Set<fieldname>(value <fieldtype>) *<structname> {
            obj.<fieldname> = value
            return obj
        }
        """
        if "properties" not in schema_object:
            return
        for name, property in schema_object["properties"].items():
            field_type = self._get_struct_field_type(property, True)
            if "[]" not in field_type:
                address_of = "" if "required" in schema_object and name in schema_object["required"] else "&"
                field_name = self._get_external_field_name(name)
                function_name = f"Set{self._get_external_name(name)}"
                self._write(self._get_description(function_name, property))
                self._write(f"func (obj *{struct_name}) {function_name}(value {field_type}) *{struct_name} {{")
                self._write(f"obj.{field_name} = {address_of}value", 1)
                self._write(f"return obj", 1)
                self._write("}")
                self._write()

    def _write_struct_field_adders(self, struct_name, schema_object):
        """Given a struct as follows write an Add<Fieldname> function.
        type error struct {
            errors *[]string
        }
        api.NewError().Errors().Add("bad address")

        type error1 struct {
            a string
            b string
        }
        api.NewError().Errors().Add("bad address")

        // AddError TBD
        func (obj *error) AddError(value string) *string {
            if obj.errors == nil {
                obj.errors = &[]string{}
            }
            slice := append(*obj.errors, value)
            obj.errors = &slice
            return &slice[len(slice)-1]
        }
        """
        if "properties" not in schema_object:
            return
        for name, property in schema_object["properties"].items():
            pass

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
            self._mod_dir = os.path.normpath(os.path.join(self._output_dir, self._python_module_name))
            os.environ["GO111MODULE"] = "on"
            print("Tidying the generated go mod file: {}".format(" ".join(process_args)))
            process = subprocess.Popen(process_args, cwd=self._mod_dir, shell=True, env=os.environ)
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))
