import os, re
from jsonpath_ng import parse
import openapiart.goserver.string_util as util
import openapiart.goserver.generator_context as ctx
from openapiart.goserver.writer import Writer

class GoServerControllerGenerator(object):

    def __init__(self, ctx: ctx.GeneratorContext):
        self._indent = '\t'
        self._root_package = ctx.module_path
        self._package_name = "controllers"
        self._ctx = ctx
        self._output_path = os.path.join(ctx.output_path, 'controllers')
    
    def generate(self):
        self._write_controllers()

    def _write_controllers(self):
        if not os.path.exists(self._output_path):
            os.makedirs(self._output_path)
        for ctrl in self._ctx.controllers:
            self._write_controller(ctrl)

    def _write_controller(self, ctrl: ctx.Controller):
        filename = ctrl.yamlname.lower() + "_controller.go"
        fullname = os.path.join(self._output_path, filename)
        w = Writer(self._indent)
        self._write_header(w)
        self._write_import(w)
        self._write_controller_struct(w, ctrl)
        self._write_newcontroller(w, ctrl)
        self._write_routes(w, ctrl)
        self._write_methods(w, ctrl)
        with open(fullname, 'w') as file:
            print(f"Controller: {fullname}")
            for line in w.strings:
                file.write(line + '\n')
            pass
        pass

    def _write_header(self, w: Writer):
        w.write_line(
            "// This file is autogenerated. Do not modify",
            f"package {self._package_name}",
            ""
        )

    def _write_import(self, w: Writer):
        w.write_line(
            "import ("
        ).push_indent(
        ).write_line(
            '"io/ioutil"',
            '"net/http"',
            '"google.golang.org/protobuf/encoding/protojson"',
            f'"{self._root_package}/httpapi"',
            f'"{self._root_package}/httpapi/interfaces"',
            f'{re.sub("[.]", "", self._ctx.models_prefix)} "{self._ctx.models_path}"',

            # f'"{self._root_package}/models"',
        ).pop_indent(
        ).write_line(
            ")",
            ""
        )

    def _struct_name(self, ctrl: ctx.Controller) -> str:
        return util.camel_case(ctrl.controller_name)
    
    def _write_controller_struct(self, w: Writer, ctrl: ctx.Controller):
        w.write_line(
            f"type {self._struct_name(ctrl)} struct {{",
        )
        w.push_indent()
        w.write_line(
            f"handler interfaces.{ctrl.service_handler_name}",
        )
        w.pop_indent()
        w.write_line(
            "}",
            ""
        )
        pass

    def _write_newcontroller(self, w: Writer, ctrl: ctx.Controller):
        w.write_line(
            f"func NewHttp{ctrl.controller_name}(handler interfaces.{ctrl.service_handler_name}) interfaces.{ctrl.controller_name} {{",
        ).push_indent()
        w.write_line(
            f"return &{self._struct_name(ctrl)}{{handler}}",
        ).pop_indent()
        w.write_line(
            "}",
            ""
        )
        pass

    def _write_routes(self, w: Writer, ctrl: ctx.Controller):
        w.write_line(
            f"func (ctrl *{self._struct_name(ctrl)}) Routes() []httpapi.Route {{",
        ).push_indent()
        w.write_line(
            "return [] httpapi.Route {",
        ).push_indent()
        for r in ctrl.routes:
            w.write_line(
                f'{{ Path: "{r.url}\", Method: \"{r.method}\", Name: "{r.operation_name}", Handler: ctrl.{r.operation_name}}},',
            )
        w.pop_indent()
        w.write_line(
            "}",
        ).pop_indent()
        w.write_line(
            "}",
            ""
        )
        pass

    def _write_methods(self, w: Writer, ctrl: ctx.Controller):
        for route in ctrl.routes:
            self._write_method(w, ctrl, route)
    
    def _write_method(self, w: Writer, ctrl: ctx.Controller, route: ctx.ControllerRoute):
        w.write_line("/*")
        w.write_line(f"{route.operation_name}: {route.method} {route.url}")
        w.write_line("Description: " + route.description)
        w.write_line("*/")
        w.write_line(
            f"func (ctrl *{self._struct_name(ctrl)}) {route.operation_name}(w http.ResponseWriter, r *http.Request) {{",
        )
        w.push_indent()
        request_body: Component = route.requestBody()
        rsp_400_error = "response{}400".format(route.operation_name)
        rsp_500_error = "response{}500".format(route.operation_name)
        if request_body != None:
            modelname = request_body.model_name
            full_modelname = request_body.full_model_name
            new_modelname = f"{self._ctx.models_prefix}New{modelname}"

            w.write_line(
                f"var item {full_modelname}",
                "if r.Body != nil {",
                "    body, readError := ioutil.ReadAll(r.Body)",
                "    if body != nil {",
                f"        item = {new_modelname}()",
                "        err := item.FromJson(string(body))",
                "        if err != nil {",
                f"            ctrl.{rsp_400_error}(w, err)",
                "            return",
                "        }",
                "    } else {",
                f"        ctrl.{rsp_400_error}(w, readError)",
                "        return"
                "    }"  
                "} else {",
                "    bodyError := errors.New(\"Request do not have any body\")",
                f"    ctrl.{rsp_400_error}(w, bodyError)",
                "    return",
                "}",
                f"result := ctrl.handler.{route.operation_name}(item, r)",
            )
        else:
            w.write_line(
                f"result := ctrl.handler.{route.operation_name}(r)",
            )

        error_responses = []
        for response in route.responses:
            # This is require as workaround of https://github.com/open-traffic-generator/openapiart/issues/220
            if self._need_warnning_check(route, response, ctrl):
                self._handle_warnning(w, rsp_400_error)
                continue

            if int(response.response_value) in [400, 500]:
                error_responses.append(response)
            w.write_line(
                f"if result.HasStatusCode{response.response_value}() {{",
            ).push_indent()
            # print(response_obj)
            # no response content defined, return as 'any'
            write_method = None
            if response.has_json:
                write_method = "WriteJSONResponse"
            elif response.has_binary:
                write_method = "WriteByteResponse"
            else:
                write_method = "WriteAnyResponse"
            w.write_line(
                "httpapi.{write_method}(w, {response_value}, result.StatusCode{response_value}())".format(
                    write_method=write_method,
                    response_value=response.response_value
                )
            )
            w.write_line("return")
            w.pop_indent()
            w.write_line(
                "}",
            )

        # w.push_indent()
        # for r in ctrl.routes:
        #     w.write_line(
        #         f"httpapi.Route(\"{r.url}\", ctrl.{r.operation_name}, \"{r.method}\"),",
        #     )
        # w.pop_indent()
        w.write_line("httpapi.WriteDefaultResponse(w, http.StatusInternalServerError)")
        w.pop_indent()
        w.write_line(
            "}",
            ""
        )

        for err_rsp in error_responses:
            result_status_code = "StatusCode{}".format(str(err_rsp.response_value))
            w.write_line("""func (ctrl *{struct_name}) {method_name}(w http.ResponseWriter, rsp_err error) {{
                result := {models_prefix}New{response_model_name}()
                result.StatusCode{response_value}().SetErrors([]string{{rsp_err.Error()}})
                httpapi.WriteJSONResponse(w, {response_value}, result.{result_status_code}())
            }}
            """.format(
                struct_name=self._struct_name(ctrl),
                method_name=rsp_400_error if int(err_rsp.response_value) == 400 else rsp_500_error,
                models_prefix=self._ctx.models_prefix,
                response_model_name=route.response_model_name,
                response_value=err_rsp.response_value,
                result_status_code=result_status_code
            ))

        pass

    def _need_warnning_check(self, route, response, ctrl):
        parse_schema = parse("$..schema").find(response.response_obj)
        schema = [s.value for s in parse_schema]
        if len(schema) == 0:
            return False
        schema = schema[0]
        if "$ref" in schema:
            schema = self._ctx.get_object_from_ref(schema["$ref"])
        parse_warnings = parse("$..warnings").find(schema)
        if route.method in ["PUT", "POST"] and \
                int(response.response_value) == 200 and \
                response.has_json and \
                len(parse_warnings) > 0:
            return True
        return False


    def _handle_warnning(self, w, rsp_400_error):
        w.write_line("""if result.HasStatusCode200() {{
                opts := protojson.MarshalOptions{{
                    UseProtoNames:   true,
                    AllowPartial:    true,
                    EmitUnpopulated: true,
                    Indent:          "  ",
                }}
                data, err := opts.Marshal(result.StatusCode200().Msg())
                if err != nil {{
                    ctrl.{rsp_400_error}(w, err)
                }}
                httpapi.WriteCustomJSONResponse(w, 200, data)
                return
            }}
            """.format(
                rsp_400_error=rsp_400_error
        ))

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


