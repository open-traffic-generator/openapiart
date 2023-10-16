import os
import re
from jsonpath_ng import parse
import openapiart.goserver.string_util as util
import openapiart.goserver.generator_context as ctx
from openapiart.goserver.writer import Writer


class GoServerControllerGenerator(object):
    def __init__(self, context):
        # type: (ctx.GeneratorContext) -> None
        self._indent = "\t"
        self._root_package = context.module_path
        self._package_name = "controllers"
        self._ctx = context
        self._output_path = os.path.join(context.output_path, "controllers")

    def generate(self):
        self._write_controllers()

    def _write_controllers(self):
        if not os.path.exists(self._output_path):
            os.makedirs(self._output_path)
        for ctrl in self._ctx.controllers:
            self._write_controller(ctrl)

    def _write_controller(self, ctrl):
        # type: (ctx.Controller) -> None
        filename = ctrl.yamlname.lower() + "_controller.go"
        fullname = os.path.join(self._output_path, filename)
        w = Writer(self._indent)
        self._write_header(w)
        self._write_import(w)
        self._write_controller_struct(w, ctrl)
        self._write_newcontroller(w, ctrl)
        self._write_routes(w, ctrl)
        self._write_methods(w, ctrl)
        with open(fullname, "w") as file:
            print("Controller: {}".format(fullname))
            for line in w.strings:
                file.write(line + "\n")

    def _write_header(self, w):
        # type: (Writer) -> None
        w.write_line(
            """// This file is autogenerated. Do not modify
            package {package_name}
            """.format(
                package_name=self._package_name
            )
        )

    def _write_import(self, w):
        # type: (Writer) -> None
        w.write_line("import (").push_indent().write_line(
            """\"io"
            "net/http"
            "google.golang.org/protobuf/encoding/protojson"
            "google.golang.org/grpc/status"
            "{root_package}/httpapi"
            "{root_package}/httpapi/interfaces"
            {models_prefix} "{models_path}\"""".format(
                root_package=self._root_package,
                models_prefix=re.sub("[.]", "", self._ctx.models_prefix),
                models_path=self._ctx.models_path,
            )
        ).pop_indent().write_line(")", "")

    def _struct_name(self, ctrl):
        # type: (ctx.Controller) -> str
        return util.camel_case(ctrl.controller_name)

    def _write_controller_struct(self, w, ctrl):
        # type: (Writer, ctx.Controller) -> None
        w.write_line(
            "type {name} struct {{".format(name=self._struct_name(ctrl))
        )
        w.push_indent()
        w.write_line(
            "handler interfaces.{name}".format(name=ctrl.service_handler_name)
        )
        w.pop_indent()
        w.write_line("}", "")

    def _write_newcontroller(self, w, ctrl):
        # type: (Writer, ctx.Controller) -> None
        w.write_line(
            "func NewHttp{ctl_name}(handler interfaces.{handle_name}) interfaces.{ctl_name} {{".format(
                ctl_name=ctrl.controller_name,
                handle_name=ctrl.service_handler_name,
            )
        ).push_indent()
        w.write_line(
            "return &{name}{{handler}}".format(name=self._struct_name(ctrl))
        ).pop_indent()
        w.write_line("}", "")

    def _write_routes(self, w, ctrl):
        # type: (Writer, ctx.Controller) -> None
        w.write_line(
            "func (ctrl *{name}) Routes() []httpapi.Route {{".format(
                name=self._struct_name(ctrl)
            )
        ).push_indent()
        w.write_line("return [] httpapi.Route {").push_indent()
        for r in ctrl.routes:
            w.write_line(
                """{{ Path: "{url}", Method: "{method}", Name: "{operation_name}", Handler: ctrl.{operation_name}}},""".format(
                    url=r.url, method=r.method, operation_name=r.operation_name
                )
            )
        w.pop_indent()
        w.write_line(
            "}",
        ).pop_indent()
        w.write_line("}", "")

    def _write_methods(self, w, ctrl):
        # type: (Writer, ctx.Controller) -> None
        self._has_warning_check = False
        for route in ctrl.routes:
            self._write_method(w, ctrl, route)
        if self._has_warning_check:
            w.write_line(
                """var {name}MrlOpts = protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: true,
                    Indent:          "  ",
                }}""".format(
                    name=util.camel_case(ctrl.yamlname)
                )
            )

    def _write_method(self, w, ctrl, route):
        # type: (Writer, ctx.Controller,ctx.ControllerRoute) -> None
        w.write_line("/*")
        w.write_line(
            "{operation_name}: {method} {url}".format(
                operation_name=route.operation_name,
                method=route.method,
                url=route.url,
            )
        )
        w.write_line("Description: " + route.description)
        w.write_line("*/")
        w.write_line(
            "func (ctrl *{name}) {opt_name}(w http.ResponseWriter, r *http.Request) {{".format(
                name=self._struct_name(ctrl), opt_name=route.operation_name
            )
        )
        w.push_indent()
        request_body = route.requestBody()  # type: ctx.Component
        rsp_error = "response{}Error".format(route.operation_name)
        rsp_errors = ["default"]
        if request_body is not None:
            modelname = request_body.model_name
            full_modelname = request_body.full_model_name
            new_modelname = self._ctx.models_prefix + "New" + modelname

            w.write_line(
                """var item {full_modelname}
                if r.Body != nil {{
                    body, readError := io.ReadAll(r.Body)
                    if body != nil {{
                        item = {new_modelname}()
                        err := item.FromJson(string(body))
                        if err != nil {{
                            ctrl.{rsp_400_error}(w, "validation", err)
                            return
                        }}
                    }} else {{
                        ctrl.{rsp_400_error}(w, "validation", readError)
                        return
                    }}
                }} else {{
                    bodyError := errors.New(\"Request does not have a body\")
                    ctrl.{rsp_500_error}(w, "validation", bodyError)
                    return
                }}
                result, err := ctrl.handler.{operation_name}(item, r)""".format(
                    full_modelname=full_modelname,
                    new_modelname=new_modelname,
                    rsp_400_error=rsp_error,
                    rsp_500_error=rsp_error,
                    operation_name=route.operation_name,
                )
            )
        else:
            w.write_line(
                "result, err := ctrl.handler.{name}(r)".format(
                    name=route.operation_name
                )
            )

        w.write_line(
            """if err != nil {{
                ctrl.{rsp_error}(w, "internal", err)
                return
            }}
            """.format(
                rsp_error=rsp_error
            )
        )

        error_responses = []
        struct_name = ""
        for response in route.responses:
            if response.response_value in rsp_errors:
                error_responses.append(response)

            # no response content defined, return as 'any'
            if response.has_json:
                write_method = "WriteJSONResponse"
                struct_name = self._get_struct_name_from_ref(
                    response.response_obj
                )
            elif response.has_binary:
                write_method = "WriteByteResponse"
                struct_name = "ResponseBytes"
            else:
                write_method = "WriteAnyResponse"
                struct_name = "ResponseString"

            # This is require as workaround of https://github.com/open-traffic-generator/openapiart/issues/220
            if self._need_warning_check(route, response):
                rsp_section = """data, err := {mrl_name}MrlOpts.Marshal(result.{struct}().Msg())
                        if err != nil {{
                            ctrl.{rsp_400_error}(w, "validation", err)
                        }}
                        httpapi.WriteCustomJSONResponse(w, 200, data)
                    """.format(
                    mrl_name=util.camel_case(ctrl.yamlname),
                    struct=struct_name,
                    rsp_400_error=rsp_error,
                )
            elif response.response_value == "default":
                # we dont want to check for default and stuff
                continue
            else:
                rsp_section = """if _, err := httpapi.{write_method}(w, {response_value}, result.{struct_name}()); err != nil {{
                            log.Print(err.Error())
                    }}""".format(
                    write_method=write_method,
                    response_value=response.response_value,
                    struct_name=struct_name,
                )

            w.write_line(
                """if result.Has{response_value}() {{
                               {rsp_section}
                               return
                           }}""".format(
                    response_value=struct_name,
                    rsp_section=rsp_section,
                )
            )
        w.write_line(
            'ctrl.{rsp_500_error}(w, "internal", errors.New("Unknown error"))'.format(
                rsp_500_error=rsp_error
            )
        )
        w.pop_indent()
        w.write_line("}", "")

        for err_rsp in error_responses:
            schema = parse("$..schema").find(err_rsp.response_obj)[0].value
            schema_name = ""
            if "$ref" in schema:
                schema_name = self._get_external_struct_name(
                    schema["$ref"].split("/")[-1]
                )
                schema = self._ctx.get_object_from_ref(schema["$ref"])
            set_errors = """
                if rErr, ok := rsp_err.({models_prefix}Error); ok {{
                    result = rErr
                }} else {{
                    result = {models_prefix}NewError()
                    err := result.FromJson(rsp_err.Error())
                    if err != nil {{
                        result.Msg().Code = &statusCode
                        err = result.SetKind(errorKind)
                        if err != nil {{
                            log.Print(err.Error())
                        }}
                        result.Msg().Errors = []string{{rsp_err.Error()}}
                    }}
                }}
            """.format(
                models_prefix=self._ctx.models_prefix,
            )

            w.write_line(
                """func (ctrl *{struct_name}) {method_name}(w http.ResponseWriter, errorKind {models_prefix}ErrorKindEnum, rsp_err error) {{
                var result {models_prefix}{schema}
                var statusCode int32
                if errorKind == "validation" {{
                    statusCode = 400
                }} else if errorKind == "internal" {{
                    statusCode = 500
                }}
                {set_errors}
                if _, err := httpapi.WriteJSONResponse(w, int(result.Code()), result); err != nil {{
                    log.Print(err.Error())
                }}
            }}
            """.format(
                    struct_name=self._struct_name(ctrl),
                    method_name=rsp_error,
                    set_errors=set_errors,
                    models_prefix=self._ctx.models_prefix,
                    schema=schema_name if schema_name != "" else "Error",
                )
            )

    def _need_warning_check(self, route, response):
        parse_schema = parse("$..schema").find(response.response_obj)
        schema = [s.value for s in parse_schema]
        if len(schema) == 0:
            return False
        schema = schema[0]
        if "$ref" in schema:
            schema = self._ctx.get_object_from_ref(schema["$ref"])
        parse_warnings = parse("$..warnings").find(schema)
        if (
            route.method in ["PUT", "POST"]
            and response.response_value != "default"
            and response.has_json
            and len(parse_warnings) > 0
        ):
            self._has_warning_check = True
            return True
        return False

    def _get_struct_name_from_ref(self, resp_obj):
        if "schema" in resp_obj["content"]["application/json"]:
            if "$ref" in resp_obj["content"]["application/json"]["schema"]:
                refStr = resp_obj["content"]["application/json"]["schema"][
                    "$ref"
                ].split("/")[-1]
                return self._get_external_struct_name(refStr)
        return ""

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

    # def _write_servicehandler_interface(self, w: Writer, ctrl: ctx.Controller):
    #     w.write_line(
    #         f"type {ctrl.service_handler_name} interface {{",
    #     )
    #     w.push_indent()
    #     w.write_line(
    #         f"GetController() {ctrl.controller_name}",
    #     )
    #     for r in ctrl.routes:
    #         response_model_name = r.operation_name + 'Response'
    #         w.write_line(
    #             f"{r.operation_name}(r *http.Request) models.{response_model_name}",
    #         )
    #     w.pop_indent()
    #     w.write_line(
    #         "}",
    #         ""
    #     )
    #     pass
