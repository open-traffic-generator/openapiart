import os
import jsonpath_ng
from .openapiartplugin import OpenApiArtPlugin


class OpenApiArtProtobuf(OpenApiArtPlugin):
    def __init__(self, **kwargs):
        super(OpenApiArtProtobuf, self).__init__(**kwargs)
        self._filename = os.path.normpath(os.path.join(self._output_dir, "{}.proto".format(self._protobuf_file_name)))
        self.default_indent = "  "
        self._custom_id = 60000
        self._init_fp(self._filename)

    def generate(self, openapi):
        self._openapi = openapi
        self._operations = {}
        self._write_header(self._openapi["info"])
        for name, schema_object in self._openapi["components"]["schemas"].items():
            self._write_msg(name, schema_object)
        for name, response_object in self._openapi["components"]["responses"].items():
            self._write_msg(name, response_object)
        for _, path_object in self._openapi["paths"].items():
            self._write_request_msg(path_object)
            self._write_response_msg(path_object)
        self._write_service()
        self._close_fp()

    def _get_operation(self, path_item_object):
        if "operationId" in path_item_object:
            operation_id = path_item_object["operationId"]
            if operation_id not in self._operations:
                operation = lambda x: None
                operation.rpc = self._get_camel_case(operation_id)
                operation.request = "google.protobuf.Empty"
                operation.response = "{}Response".format(operation.rpc)
                operation.stream = len(self._get_parser('$.."application/octet-stream"').find(path_item_object)) > 0
                self._operations[operation_id] = operation
            return self._operations[operation_id]
        return None

    def _write_request_msg(self, path_object):
        for _, path_item_object in path_object.items():
            operation = self._get_operation(path_item_object)
            if operation is None:
                continue
            if len(self._get_parser("$..requestBody").find(path_item_object)) > 0:
                operation.request = "{}Request".format(operation.rpc)
                self._write()
                self._write("message {} {{".format(operation.request))
                for ref in self._get_parser('$..requestBody.."$ref"').find(path_item_object):
                    message = self._get_message_name(ref.value)
                    field_type = message.replace(".", "")
                    field_name = self._lowercase(field_type)
                    self._write("{} {} = 1;".format(field_type, field_name), indent=1)
                self._write("}")

    def _write_response_msg(self, path_object):
        """
        application/octet-stream -> response (stream <operationId>Response)
        application/json -> response (<operationId>Response)
        """
        for _, path_item_object in path_object.items():
            operation = self._get_operation(path_item_object)
            if operation is None:
                continue
            self._write()
            self._write("message {} {{".format(operation.response))
            for ref in self._get_parser("$..responses").find(path_item_object):
                detail_messages = []
                for code, response in ref.value.items():
                    detail_message = "StatusCode{}".format(code)
                    detail_messages.append(detail_message)
                    self._write("message {} {{".format(detail_message), indent=1)
                    schema = self._get_parser("$..schema").find(response)
                    if len(schema) == 0:
                        if "$ref" in response:
                            field_type = self._get_message_name(response["$ref"]).replace(".", "")
                            self._write("{} {} = 1;".format(field_type, self._lowercase(field_type)), indent=2)
                    else:
                        schema = schema[0].value
                        if "$ref" in schema:
                            field_type = self._get_message_name(schema["$ref"]).replace(".", "")
                            self._write("{} {} = 1;".format(field_type, self._lowercase(field_type)), indent=2)
                        elif "type" in schema:
                            field_type = self._get_message_name(schema["type"]).replace(".", "")
                            if "format" in schema and schema["format"] == "binary":
                                field_type = "bytes"
                            self._write("{} {} = 1;".format(field_type, self._lowercase(field_type)), indent=2)
                    self._write("}", indent=1)
            self._write()
            id = 1
            for detail_message in detail_messages:
                field_type = detail_message.replace(".", "")
                field_name = self._lowercase(field_type)
                self._write("optional {} {} = {};".format(field_type, field_name, id), indent=1)
                id += 1
            self._write("}")

    def _get_message_name(self, ref):
        return ref.split("/")[-1]

    def _next_custom_id(self):
        self._custom_id += 1
        return self._custom_id

    def _write_header(self, info_object):
        self._write("// {}".format(self._info))
        for line in self._license.split("\n"):
            self._write("// {}".format(line))
        self._write()
        self._write('syntax = "proto3";')
        self._write()
        self._write("package {};".format(self._protobuf_package_name))
        self._write()
        self._write('option go_package = "{}/{}";'.format(self._go_sdk_package_dir, self._protobuf_package_name))
        self._write()
        self._write('import "google/protobuf/descriptor.proto";')
        self._write('import "google/protobuf/empty.proto";')
        self._write()
        self._write("message OpenApiMsgOpt {")
        self._write("string description = 10;", indent=1)
        self._write("}")
        self._write("extend google.protobuf.MessageOptions {")
        self._write("optional OpenApiMsgOpt msg_meta = {};".format(self._next_custom_id()), indent=1)
        self._write("}")
        self._write()
        self._write("message OpenApiFldOpt {")
        self._write("string default = 10;", indent=1)
        self._write("string description = 20;", indent=1)
        self._write("}")
        self._write("extend google.protobuf.FieldOptions {")
        self._write("optional OpenApiFldOpt fld_meta = {};".format(self._next_custom_id()), indent=1)
        self._write("}")
        self._write()
        self._write("message OpenApiSvcOpt {")
        self._write("string description = 10;", indent=1)
        self._write("}")
        self._write("extend google.protobuf.ServiceOptions {")
        self._write("optional OpenApiSvcOpt svc_meta = {};".format(self._next_custom_id()), indent=1)
        self._write("}")
        self._write()
        self._write("message OpenApiRpcOpt {")
        self._write("string description = 10;", indent=1)
        self._write("}")
        self._write("extend google.protobuf.MethodOptions {")
        self._write("optional OpenApiRpcOpt rpc_meta = {};".format(self._next_custom_id()), indent=1)
        self._write("}")

    def _get_field_type(self, property_name, openapi_object):
        """Convert openapi type -> protobuf type

        - type:number -> float
        - type:number [format: float] -> float
        - type:number [format: double] -> double
        - type:integer -> int32
        - type:integer [format:int32] -> int32
        - type:integer [format:int64] -> int64
        - type:boolean -> bool
        - type:string -> string
        - type:string [format:binary] -> bytes
        """
        if "type" in openapi_object:
            type = openapi_object["type"]
            if type == "boolean":
                return "bool"
            if type == "string":
                if "format" in openapi_object:
                    if openapi_object["format"] == "binary":
                        return "bytes"
                elif "enum" in openapi_object:
                    enum_msg = self._camelcase("{}".format(property_name))
                    self._write_enum_msg(enum_msg, openapi_object["enum"])
                    return enum_msg + ".Enum"
                return "string"
            if type == "integer":
                return "int32"
            if type == "number":
                if "format" in openapi_object:
                    if openapi_object["format"] == "double":
                        return "double"
                    elif openapi_object["format"] == "float":
                        return "float"
                return "float"
            if type == "array":
                return "repeated " + self._get_field_type(property_name, openapi_object["items"])
        elif "$ref" in openapi_object:
            return openapi_object["$ref"].split("/")[-1].replace(".", "")

    def _camelcase(self, value):
        camel_case = ""
        for piece in value.split("_"):
            camel_case += "{}{}".format(piece[0].upper(), piece[1:])
        return camel_case

    def _camelcase_to_snakecase(self, value, lower=False):
        word = ""
        insert_underscore = False

        for c in value:
            if c.isupper() or c.isdigit():
                if insert_underscore:
                    word += "_" + (c.lower() if lower else c)
                    insert_underscore = False
                else:
                    word += c.lower() if lower else c
            else:
                word += c if lower else c.upper()
                insert_underscore = True

        return word

    def _uppercase(self, value):
        return self._camelcase_to_snakecase(value, lower=False)

    def _lowercase(self, value):
        return self._camelcase_to_snakecase(value, lower=True)

    def _get_description(self, openapi_object):
        if "description" in openapi_object:
            return openapi_object["description"].replace("\n", "\\n").replace('"', "")
        else:
            return "Description missing in models"

    def _write_enum_msg(self, enum_msg_name, enums):
        """Follow google developers style guide for enums
        - reference: https://developers.google.com/protocol-buffers/docs/style#enums
        """
        self._write("message {} {{".format(enum_msg_name.replace(".", "")), indent=1)
        self._write("enum Enum {", indent=2)
        enums.insert(0, "unspecified")
        id = 0
        for enum in enums:
            self._write("{} = {};".format(enum.lower(), id), indent=3)
            id += 1
        self._write("}", indent=2)
        self._write("}", indent=1)

    def _write_msg(self, name, schema_object):
        msg_name = name.replace(".", "")
        print("writing msg {}".format(msg_name))
        self._write()
        self._write("message {} {{".format(msg_name), indent=0)
        self._write('option (msg_meta).description = "{}";'.format(self._get_description(schema_object)), indent=1)
        if "content" in schema_object:
            # when accessing components/responses
            self._write_response_fields(msg_name, schema_object)
        else:
            # when accessing components/schemas
            self._write_msg_fields(name, schema_object)
        self._write("}")

    def _write_response_fields(self, msg_name, schema_object):
        try:
            ref = schema_object["content"]["application/json"]["schema"]["$ref"]
            field_type = self._get_message_name(ref).replace(".", "")
            self._write("{} {} = 1;".format(field_type, self._lowercase(field_type)), indent=1)
        except AttributeError as err:
            print("Failed writing response {}: {}".format(msg_name, err))

    def _write_msg_fields(self, name, schema_object):
        if "properties" not in schema_object:
            return
        id = 0
        for property_name, property_object in schema_object["properties"].items():
            id += 1
            self._write()
            property_type = self._get_field_type(property_name, property_object)
            default = None
            if "default" in property_object:
                default = property_object["default"]
            if property_type.endswith(".Enum"):
                if default is not None:
                    default = "{}.{}".format(property_type.split(" ")[-1], default.lower())
            if "required" in schema_object and property_name in schema_object["required"] or property_type.startswith("repeated"):
                optional = ""
            else:
                optional = "optional "
            self._write("{}{} {} = {} [".format(optional, property_type, property_name.lower(), id), indent=1)
            if default is not None:
                self._write('(fld_meta).default = "{}",'.format(default), indent=2)
            self._write('(fld_meta).description = "{}"'.format(self._get_description(property_object)), indent=2)
            self._write("];", indent=1)

    def _write_service(self):
        self._write()
        self._write("service Openapi {")
        paths_object = self._openapi["paths"]
        self._write('option (svc_meta).description = "{}";'.format(self._get_description(paths_object)), indent=1)
        self._write()
        for url, path_object in self._openapi["paths"].items():
            for method, path_item_object in path_object.items():
                if method in ["get", "patch", "post", "delete"]:
                    self._write_rpc(url, method, path_item_object)
        self._write("}")

    def _write_rpc(self, url, method, path_item_object):
        """ """
        operation = self._get_operation(path_item_object)
        print("writing rpc {}".format(operation.rpc))
        line = "rpc {}({}) returns ({}{}) {{".format(operation.rpc, operation.request, "", operation.response)
        self._write(line, indent=1)
        self._write('option (rpc_meta).description = "{}";'.format(self._get_description(path_item_object)), indent=2)
        self._write("}", indent=1)
