from .openapiartplugin import OpenApiArtPlugin
import os
import subprocess
import shutil


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
        self.request = "emptypb.Empty{}"
        self.responses = []


class FluentRpcResponse(object):
    """<operation_name>StatusCode<status_code>
    status_code is the http status code
    fluent_new is the 2xx response, all other status codes this should be None
    """

    def __init__(self):
        self.status_code = None
        self.schema = None


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
        self.description = None
        self.type = None
        self.isOptional = True
        self.isPointer = True
        self.isArray = False
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
        self._ux_path = os.path.normpath(os.path.join(self._output_dir, "..", os.path.split(self._go_sdk_package_dir)[-1]))
        self._protoc_path = os.path.normpath(os.path.join(self._ux_path, self._protobuf_package_name))
        self._structs = {}
        self._write_mod_file()
        self._write_go_file()
        self._format_go_file()
        self._tidy_mod_file()

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
        self._write_response_interfaces()
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
        line = 'import {pb_pkg_name} "{go_sdk_pkg_dir}/{pb_pkg_name}"'.format(pb_pkg_name=self._protobuf_package_name, go_sdk_pkg_dir=self._go_sdk_package_dir)
        self._write(line)
        self._write('import "google.golang.org/protobuf/types/known/emptypb"')
        self._write('import "google.golang.org/grpc"')
        self._write('import "github.com/ghodss/yaml"')
        with open(os.path.join(os.path.dirname(__file__), "common.go")) as fp:
            self._write(fp.read().strip().strip("\n"))
        self._write()

    def _write_types(self):
        for _, go_type in self._oapi_go_types.items():
            if go_type.startswith("String"):
                self._write("type {go_type} string".format(go_type=go_type))
        self._write()

    def _get_internal_name(self, openapi_name):
        name = self._get_external_struct_name(openapi_name)
        return name[0].lower() + name[1:]

    def _get_external_field_name(self, openapi_name):
        external_name = ""
        for piece in openapi_name.replace(".", "").split("_"):
            for i in range(len(piece)):
                if i == 0 and piece[i].isdigit() is False:
                    external_name += piece[i].upper()
                elif piece[i].isdigit():
                    if "_" + piece[i] in openapi_name:
                        external_name += "_" + piece[i]
                    else:
                        external_name += piece[i]
                elif external_name[-1].isdigit():
                    external_name += piece[i].upper()
                else:
                    external_name += piece[i]
        if external_name in ["String"]:
            external_name += "_"
        return external_name

    def _get_external_struct_name(self, openapi_name):
        return self._get_external_field_name(openapi_name).replace("_", "")

    def _build_api_interface(self):
        self._api.internal_struct_name = """{internal_name}Api""".format(internal_name=self._get_internal_name(self._go_sdk_package_name))
        self._api.external_interface_name = """{external_name}Api""".format(external_name=self._get_external_struct_name(self._go_sdk_package_name))
        for operation_id in self._get_parser("$.paths..operationId").find(self._openapi):
            path_item_object = operation_id.context.value
            rpc = FluentRpc()
            rpc.operation_name = self._get_external_struct_name(operation_id.value)
            if len([m for m in self._api.external_rpc_methods if m.operation_name == rpc.operation_name]) == 0:
                self._api.external_rpc_methods.append(rpc)
            ref = self._get_parser("$..requestBody..'$ref'").find(path_item_object)
            if len(ref) == 1:
                new = FluentNew()
                new.schema_name = self._get_schema_object_name_from_ref(ref[0].value)
                new.schema_object = self._get_schema_object_from_ref(ref[0].value)
                new.interface = self._get_external_struct_name(new.schema_name)
                new.struct = self._get_internal_name(new.schema_name)
                new.method = """New{interface}() {interface}""".format(interface=new.interface)
                if len([m for m in self._api.external_new_methods if m.schema_name == new.schema_name]) == 0:
                    self._api.external_new_methods.append(new)
                rpc.request = "{pb_pkg_name}.{operation_name}Request{{{interface}: {struct}.msg()}}".format(
                    pb_pkg_name=self._protobuf_package_name,
                    operation_name=rpc.operation_name,
                    interface=new.interface,
                    struct=new.struct,
                )
                rpc.method = """{operation_name}({struct} {interface}) ({operation_response_name}Response_StatusCode200, error)""".format(
                    operation_name=rpc.operation_name,
                    operation_response_name=self._get_external_struct_name(rpc.operation_name),
                    struct=new.struct,
                    interface=new.interface,
                )
            else:
                rpc.method = """{operation_name}() ({operation_response_name}Response_StatusCode200, error)""".format(
                    operation_name=rpc.operation_name,
                    operation_response_name=self._get_external_struct_name(rpc.operation_name),
                )
            for ref in self._get_parser("$..responses").find(path_item_object):
                for status_code, schema in ref.value.items():
                    response = FluentRpcResponse()
                    response.status_code = status_code
                    response.schema = schema
                    rpc.responses.append(response)

        # write the go code
        self._write(
            """type {internal_struct_name} struct {{
                api
                grpcClient {pb_pkg_name}.OpenapiClient
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

            // NewApi returns a new instance of the top level interface hierarchy
            func NewApi() *{internal_struct_name} {{
                api := {internal_struct_name}{{}}
                return &api
            }}
            """.format(
                internal_struct_name=self._api.internal_struct_name,
                pb_pkg_name=self._protobuf_package_name,
            )
        )
        methods = []
        for new in self._api.external_new_methods:
            methods.append(new.method)
        for rpc in self._api.external_rpc_methods:
            methods.append(rpc.method)
        method_signatures = "\n".join(methods)
        self._write(
            """type {external_interface_name} interface {{
                Api
                {method_signatures}
            }}
            """.format(
                external_interface_name=self._api.external_interface_name,
                method_signatures=method_signatures,
            )
        )
        for new in self._api.external_new_methods:
            self._write(
                """func (api *{internal_struct_name}) {method} {{
                    return &{struct}{{obj: &{pb_pkg_name}.{interface}{{}}}}
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
                        data, _ := yaml.Marshal(resp.GetStatusCode_400())
                        return nil, fmt.Errorf(string(data))
                    }}
                    """.format(
                    status_code=response.status_code,
                )
            error_handling += 'return nil, fmt.Errorf("Response not implemented")'
            self._write(
                """func (api *{internal_struct_name}) {method} {{
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
                    if resp.GetStatusCode_200() != nil {{
                        return &{operation_response_name}ResponseStatusCode200{{obj: resp.GetStatusCode_200()}}, nil
                    }}
                    {error_handling}
                }}
                """.format(
                    internal_struct_name=self._api.internal_struct_name,
                    method=rpc.method,
                    request=rpc.request,
                    operation_name=rpc.operation_name,
                    operation_response_name=self._get_internal_name(rpc.operation_name),
                    error_handling=error_handling,
                )
            )

    def _build_request_interfaces(self):
        for new in self._api.external_new_methods:
            self._build_interface(new)

    def _write_component_interfaces(self):
        while True:
            components = [component for _, component in self._api.components.items() if component.generated is False]
            if len(components) == 0:
                break
            for component in components:
                self._build_interface(component)

    def _write_response_interfaces(self):
        for rpc in self._api.external_rpc_methods:
            for response in rpc.responses:
                if response.status_code.startswith("2") is False:
                    continue
                new = FluentNew()
                new.schema_object = response.schema
                new.struct = "{operation_name}ResponseStatusCode{status_code}".format(
                    operation_name=self._get_internal_name(rpc.operation_name),
                    status_code=response.status_code,
                )
                new.interface = "{operation_name}Response_StatusCode{status_code}".format(
                    operation_name=rpc.operation_name,
                    status_code=response.status_code,
                )
                self._build_interface(new)

    def _build_interface(self, new):
        self._write(
            """type {struct} struct {{
                obj *{pb_pkg_name}.{interface}
            }}
            
            func (obj *{struct}) msg() *{pb_pkg_name}.{interface} {{
                return obj.obj
            }}

            func (obj *{struct}) ToYaml() string {{
                data, _ := json.Marshal(obj.msg())
                data, _ = yaml.JSONToYAML(data)
                return string(data)
            }}

            func (obj *{struct}) FromYaml(value string) error {{
                data, _ := yaml.YAMLToJSON([]byte(value))
                return json.Unmarshal([]byte(data), obj.msg())
            }}

            func (obj *{struct}) ToJson() string {{
                data, _ := json.Marshal(obj.msg())
                return string(data)
            }}

            func (obj *{struct}) FromJson(value string) error {{
                return json.Unmarshal([]byte(value), obj.msg())
            }}
        """.format(
                struct=new.struct,
                pb_pkg_name=self._protobuf_package_name,
                interface=new.interface,
            )
        )
        self._build_setters_getters(new)
        interfaces = ["ToYaml() string", "ToJson() string"]
        for field in new.interface_fields:
            interfaces.append(field.getter_method)
            if field.setter_method is not None:
                interfaces.append(field.setter_method)
        interface_signatures = "\n".join(interfaces)
        self._write(
            """type {interface} interface {{
                msg() *{pb_pkg_name}.{interface}
                {interface_signatures}
            }}
        """.format(
                interface=new.interface,
                pb_pkg_name=self._protobuf_package_name,
                interface_signatures=interface_signatures,
            )
        )
        for field in new.interface_fields:
            self._write_field_getter(new, field)
            self._write_field_setter(new, field)
            self._write_field_adder(new, field)
        new.generated = True

    def _write_field_getter(self, new, field):
        if field.getter_method is None:
            return
        elif field.isArray:
            if field.struct:
                body = """if obj.obj.{name} == nil {{
                        obj.obj.{name} = []*{pb_pkg_name}.{pb_struct}{{}}
                    }}
                    return &{parent}{interface}Iter{{obj: obj}}
                """.format(
                    name=field.name,
                    pb_pkg_name=self._protobuf_package_name,
                    pb_struct=field.external_struct,
                    interface=field.external_struct,
                    parent=new.struct,
                )
            else:
                body = """if obj.obj.{name} == nil {{
                        obj.obj.{name} = make({type}, 0)
                    }}
                    for _, item := range value {{
                        obj.obj.{name} = append(obj.obj.{name}, item)
                    }}
                    return obj
                """.format(
                    name=field.name,
                    type=field.type,
                )
        elif field.struct is not None:
            if field.isPointer:
                body = """if obj.obj.{name} == nil {{
                        obj.obj.{name} = &{pb_pkg_name}.{pb_struct}{{}}
                    }}
                    return &{struct}{{obj: obj.obj.{name}}}
                """.format(
                    name=field.name,
                    pb_pkg_name=self._protobuf_package_name,
                    pb_struct=field.external_struct,
                    struct=field.struct,
                )
            else:
                body = "return &{struct}{{obj: obj.obj.{name}}}".format(
                    struct=field.struct,
                    name=field.name,
                )
        elif field.isPointer:
            body = """return *obj.obj.{name}""".format(name=field.name)
        else:
            body = """return obj.obj.{name}""".format(name=field.name)
        self._write(
            """
            // {fieldname} returns a {fieldtype}\n{description}
            func (obj *{struct}) {getter_method} {{
                {body}
            }}
            """.format(
                fieldname=field.name,
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
        if field.isArray:
            body = """if obj.obj.{fieldname} == nil {{
                    obj.obj.{fieldname} = make({fieldtype}, 0)
                }}
                for _, item := range value {{
                    obj.obj.{field.name} = append(obj.obj.{field.name}, item)
                }}
            }}
            """.format(
                fieldname=field.name,
                fieldtype=field.type,
            )
            return
        elif field.isPointer:
            body = """obj.obj.{fieldname} = &value""".format(fieldname=field.name)
        else:
            body = """obj.obj.{fieldname} = value""".format(fieldname=field.name)
        self._write(
            """
            // Set{fieldname} sets the {fieldtype} value in the {fieldstruct} object\n{description}
            func (obj *{newstruct}) {setter_method} {{
                {body}
                return obj
            }}
            """.format(
                fieldname=field.name,
                newstruct=new.struct,
                setter_method=field.setter_method,
                body=body,
                description=field.description,
                fieldtype=field.type,
                fieldstruct=field.external_struct,
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
                Add() {field_external_struct}
                Items() {field_type}
            }}

            func (obj *{internal_struct}) Add() {field_external_struct} {{
                newObj := &{pb_pkg_name}.{field_external_struct}{{}}
                obj.obj.obj.{field_name} = append(obj.obj.obj.{field_name}, newObj)
                return &{field_internal_struct}{{obj: newObj}}
            }}

            func (obj *{internal_struct}) Items() {field_type} {{
                slice := {field_type}{{}}
                for _, item := range obj.obj.obj.{field_name} {{
                    slice = append(slice, &{field_internal_struct}{{obj: item}})
                }}
                return slice
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
            )
        )

    def _build_setters_getters(self, fluent_new):
        """Add new FluentField objects for each interface field"""
        if "properties" not in fluent_new.schema_object:
            return
        for property_name, property_schema in fluent_new.schema_object["properties"].items():
            if len(self._get_parser("$..enum").find(property_schema)) > 0:  # temporary
                continue
            field = FluentField()
            field.schema = property_schema
            field.description = self._get_description(property_schema)
            field.name = self._get_external_field_name(property_name)
            field.type = self._get_struct_field_type(property_schema)
            field.isOptional = fluent_new.isOptional(property_name)
            field.isPointer = False if field.type.startswith("[") else field.isOptional
            field.getter_method = "{name}() {ftype}".format(
                name=self._get_external_struct_name(field.name),
                ftype=field.type,
            )
            if "$ref" in property_schema:
                schema_name = self._get_schema_object_name_from_ref(property_schema["$ref"])
                field.struct = self._get_internal_name(schema_name)
                field.external_struct = self._get_external_struct_name(schema_name)
            if field.type in self._oapi_go_types.values():
                field.setter_method = "Set{name}(value {ftype}) {interface}".format(
                    name=self._get_external_struct_name(field.name),
                    ftype=field.type,
                    interface=fluent_new.interface,
                )
            elif "type" in property_schema and property_schema["type"] == "array":
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
                else:
                    field.setter_method = "Set{name}(value {ftype}) {interface}".format(
                        name=self._get_external_struct_name(field.name),
                        ftype=field.type,
                        interface=fluent_new.interface,
                    )
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
                go_type = "{oapi_go_types}".format(oapi_go_types=self._oapi_go_types[oapi_type.lower()])
            if oapi_type == "array":
                go_type += "[]" + self._get_struct_field_type(property_schema["items"]).replace("*", "")
            if "format" in property_schema:
                format_type = (oapi_type + property_schema["format"]).lower()
                if format_type.lower() in self._oapi_go_types:
                    go_type = "{oapi_go_type}".format(oapi_go_type=self._oapi_go_types[format_type.lower()])
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
                self._api.components[new.schema_name] = new
            go_type = new.interface
        else:
            raise Exception("No type or $ref keyword present in property schema: {property_schema}".format(property_schema=property_schema))
        return go_type

    def _get_description(self, openapi_object):
        description = "//  description is TBD"
        if "description" in openapi_object:
            description = ""
            for line in openapi_object["description"].split("\n"):
                description += "//  {line}\n".format(line=line.strip())
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
