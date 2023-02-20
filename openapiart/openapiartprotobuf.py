import subprocess
import os
from .openapiartplugin import OpenApiArtPlugin


class OpenApiArtProtobuf(OpenApiArtPlugin):
    def __init__(self, **kwargs):
        self._errors = []
        super(OpenApiArtProtobuf, self).__init__(**kwargs)
        self._filename = os.path.normpath(
            os.path.join(
                self._output_dir, "{}.proto".format(self._protobuf_file_name)
            )
        )
        self.default_indent = "  "
        self.proto_service_name = kwargs.get("proto_service", "Openapi")
        self.doc_dir = kwargs.get("doc_dir")
        self._init_fp(self._filename)

    def generate(self, openapi):
        self._errors = []
        self._openapi = openapi
        self._operations = {}
        self._write_header(self._openapi["info"])
        for name, schema_object in self._openapi["components"][
            "schemas"
        ].items():
            self._write_msg(name, schema_object)
        for name, response_object in self._openapi["components"][
            "responses"
        ].items():
            self._write_msg(name, response_object)
        for _, path_object in self._openapi["paths"].items():
            self._write_request_msg(path_object)
            self._write_response_msg(path_object)
        self._validate_error()
        self._write_service()
        self._close_fp()
        self.generate_doc()

    def generate_doc(self):
        if self.doc_dir is None:
            return
        process_args = [
            "protoc",
            "--doc_out={}".format(self.doc_dir),
            "--doc_opt=html,index.html",
            "--proto_path={}".format(self._output_dir),
            self._filename,
        ]
        cmd = " ".join(process_args)
        try:
            process = subprocess.Popen(cmd, shell=True)
            process.wait()
        except Exception:
            print("Bypassed generating proto document")
        # protoc --plugin=protoc-gen-doc=./protoc-gen-doc --doc_out=./doc --doc_opt=html,index.html sanity.proto

    def _validate_error(self):
        if len(self._errors) > 0:
            raise TypeError("\n".join(self._errors))

    def _get_operation(self, path_item_object):
        if "operationId" in path_item_object:
            operation_id = path_item_object["operationId"]
            if operation_id not in self._operations:
                operation = lambda x: None
                operation.rpc = self._get_camel_case(operation_id)
                operation.request = "google.protobuf.Empty"
                operation.response = "{}Response".format(operation.rpc)
                operation.stream = (
                    len(
                        self._get_parser('$.."application/octet-stream"').find(
                            path_item_object
                        )
                    )
                    > 0
                )
                self._operations[operation_id] = operation
            return self._operations[operation_id]
        return None

    def _write_request_msg(self, path_object):
        for _, path_item_object in path_object.items():
            operation = self._get_operation(path_item_object)
            if operation is None:
                continue
            if (
                len(self._get_parser("$..requestBody").find(path_item_object))
                > 0
            ):
                operation.request = "{}Request".format(operation.rpc)
                self._write()
                self._write("message {} {{".format(operation.request))
                for ref in self._get_parser('$..requestBody.."$ref"').find(
                    path_item_object
                ):
                    message = self._get_message_name(ref.value)
                    field_type = message.replace(".", "")
                    field_name = self._lowercase(field_type)
                    self._write(
                        "{} {} = 1;".format(field_type, field_name), indent=1
                    )
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
            for response in self._get_parser("$..responses").find(
                path_item_object
            ):
                response_fields = []
                for code, code_schema in response.value.items():
                    response_field = lambda: None
                    response_field.type = None
                    response_field.name = "status_code_{}".format(code)
                    response_field.field_uid = code_schema["x-field-uid"]
                    schema = self._get_parser("$..schema").find(
                        code_schema
                    )  # finds the first instance of schema in responses
                    if len(schema) > 0:
                        schema_ref = self._get_parser("$..'$ref'").find(
                            schema[0].value
                        )  # gets a ref
                    else:
                        schema_ref = self._get_parser("$..'$ref'").find(
                            code_schema
                        )  # gets a ref
                    if len(schema_ref) > 0:
                        schema = schema_ref[0].value
                    elif len(schema) > 0:
                        schema = schema[0].value
                    if "#/components/responses" in schema:
                        # lookup the response object and use the schema or ref in that object
                        jsonpath = "$.{}..schema".format(
                            schema[2:].replace("/", ".")
                        )
                        schema = (
                            self._get_parser(jsonpath)
                            .find(self._openapi)[0]
                            .value
                        )
                        ref = self._get_parser("$..'$ref'").find(schema)
                        if len(ref) > 0:
                            schema = ref[0].value
                    if "$ref" in schema:
                        response_field.type = self._get_message_name(
                            schema["$ref"]
                        ).replace(".", "")
                    elif "type" in schema:
                        response_field.type = self._get_message_name(
                            schema["type"]
                        ).replace(".", "")
                        if "format" in schema and schema["format"] == "binary":
                            response_field.type = "bytes"
                    elif len(schema) > 0:
                        response_field.type = self._get_message_name(
                            schema
                        ).replace(".", "")
                    else:
                        response_field.type = "string"
                    response_fields.append(response_field)
            self._write("message {} {{".format(operation.response))
            for response_field in response_fields:
                self._write(
                    "optional {} {} = {};".format(
                        response_field.type,
                        response_field.name,
                        response_field.field_uid,
                    ),
                    indent=1,
                )
            self._write("}")
            self._write()

    def _get_message_name(self, ref):
        return ref.split("/")[-1]

    def _write_header(self, info_object):
        self._write(
            self._justify_desc(
                self._info + "\n{}".format(self._license), use_multi=True
            )
        )
        self._write()
        self._write('syntax = "proto3";')
        self._write()
        self._write("package {};".format(self._protobuf_package_name))
        self._write()

        if self._go_sdk_package_dir is None:
            option_go_pkg = 'option go_package = "./{};{}";'.format(
                self._protobuf_package_name, self._protobuf_package_name
            )

        else:
            option_go_pkg = 'option go_package = "{}/{}";'.format(
                self._go_sdk_package_dir, self._protobuf_package_name
            )
        self._write(option_go_pkg)
        self._write()

        self._write('import "google/protobuf/descriptor.proto";')
        self._write('import "google/protobuf/empty.proto";')

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
                elif "x-enum" in openapi_object:
                    enum_msg = self._camelcase("{}".format(property_name))
                    self._write_x_enum_msg(
                        enum_msg,
                        openapi_object["x-enum"],
                        property_name,
                        openapi_object,
                    )
                    return enum_msg + ".Enum"
                return "string"
            if type == "integer":
                format = openapi_object.get("format")
                min = openapi_object.get("minimum")
                max = openapi_object.get("maximum")
                if (min is not None and min > 2147483647) or (
                    max is not None and max > 2147483647
                ):
                    return "int64"
                if format is not None and "int64" in format:
                    return "int64"
                return "int32"
            if type == "number":
                if "format" in openapi_object:
                    if openapi_object["format"] == "double":
                        return "double"
                    elif openapi_object["format"] == "float":
                        return "float"
                return "float"
            if type == "array":
                item_type = self._get_field_type(
                    property_name, openapi_object["items"]
                )
                format = openapi_object.get("format")
                min = openapi_object.get("minimum")
                max = openapi_object.get("maximum")
                if (min is not None and min > 2147483647) or (
                    max is not None and max > 2147483647
                ):
                    item_type = item_type.replace("32", "64")
                if format is not None and "int64" in format:
                    item_type = item_type.replace("32", "64")
                return "repeated " + item_type
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
            return openapi_object["description"].replace('"', "")
        else:
            return "Description missing in models"

    def _write_x_enum_msg(
        self, enum_msg_name, enums, property_name, property_object
    ):
        """Follow google developers style guide for enums
        - reference: https://developers.google.com/protocol-buffers/docs/style#enums
        """
        self._write(
            "message {} {{".format(enum_msg_name.replace(".", "")), indent=1
        )
        self._write("enum Enum {", indent=2)
        if "unspecified" not in enums:
            self._write("{} = {};".format("unspecified", 0), indent=3)
        for key, value in enums.items():
            field_uid = value["x-field-uid"]
            self._write("{} = {};".format(key.lower(), field_uid), indent=3)
        self._write("}", indent=2)
        self._write("}", indent=1)

    def _write_msg(self, name, schema_object):
        msg_name = name.replace(".", "")
        print("writing msg {}".format(msg_name))
        self._write()
        self._write(self._justify_desc(self._get_description(schema_object)))
        self._write("message {} {{".format(msg_name), indent=0)
        if "content" in schema_object:
            # when accessing components/responses
            self._write_response_fields(msg_name, schema_object)
        else:
            # when accessing components/schemas
            self._write_msg_fields(name, schema_object)
        self._write("}")

    def _write_response_fields(self, msg_name, schema_object):
        try:
            ref = schema_object["content"]["application/json"]["schema"][
                "$ref"
            ]
            field_type = self._get_message_name(ref).replace(".", "")
            self._write(
                "{} {} = 1;".format(field_type, self._lowercase(field_type)),
                indent=1,
            )
        except AttributeError as err:
            print("Failed writing response {}: {}".format(msg_name, err))

    def _write_msg_fields(self, name, schema_object):
        if "properties" not in schema_object:
            return
        for property_name, property_object in schema_object[
            "properties"
        ].items():
            self._write()
            property_type = self._get_field_type(
                property_name, property_object
            )
            default = None
            if "default" in property_object:
                default = property_object["default"]
            if property_type.endswith(".Enum"):
                if default is not None:
                    default = "{}.{}".format(
                        property_type.split(" ")[-1], default.lower()
                    )
            if (
                "required" in schema_object
                and property_name in schema_object["required"]
                or property_type.startswith("repeated")
            ):
                optional = ""
            else:
                optional = "optional "
            desc = self._get_description(property_object)
            if default is not None:
                desc += "\ndefault = {}".format(default)
            if (
                optional == ""
                and property_type.startswith("repeated") is not True
            ):
                desc += "\nrequired = true"
            self._write(self._justify_desc(desc, indent=1))
            field_uid = property_object["x-field-uid"]
            self._write(
                "{}{} {} = {};".format(
                    optional, property_type, property_name.lower(), field_uid
                ),
                indent=1,
            )

    def _write_service(self):
        self._write()
        paths_object = self._openapi["paths"]
        self._write(
            self._justify_desc(self._get_description(paths_object), indent=1)
        )
        self._write("service {name} {{".format(name=self.proto_service_name))
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
        self._write(
            self._justify_desc(
                self._get_description(path_item_object), indent=1
            )
        )
        line = "rpc {}({}) returns ({}{});".format(
            operation.rpc, operation.request, "", operation.response
        )
        self._write(line, indent=1)
