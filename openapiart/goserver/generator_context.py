import re
import openapiart.goserver.string_util as util
from jsonpath_ng import parse

class Server(object):
    @property
    def basepath(self):
        # type: () -> [str]
        return self._basepath

    def __init__(self, serverobj):
        self._obj = serverobj
        self._basepath = ''
        try:
            self._basepath = self._obj['variables']['basePath']['default']
            if self._basepath.startswith("/") == False:
                self._basepath = "/" + self._basepath
        except KeyError:
            pass


class Component(object):
    @property
    def yaml_name(self):
        # type: () -> [str]
        return self._yamlname

    @property
    def model_name(self):
        # type: () -> str
        return re.sub('[.]', '', self._yamlname)

    @property
    def full_model_name(self):
        # type: () -> str
        _ctx = self._ctx
        return "{models_prefix}{model_name}".format(
            models_prefix=_ctx.models_prefix,
            model_name=self.model_name
        )

    def __init__(
        self,
        yamlname,
        componentobj,
        ctx
        ):
        self._ctx = ctx  # type: GeneratorContext
        self._yamlname = yamlname
        self._obj = componentobj


class Responses(object):
    @property
    def response_value(self):
        return self._response_value

    @property
    def has_json(self):
        return self._has_json

    @property
    def has_binary(self):
        return self._has_binary

    @property
    def response_obj(self):
        return self._response_obj

    def __init__(self, response_value, response_obj, ctx):
        self._response_value = response_value
        self._response_obj = response_obj
        self._ctx = ctx # type: GeneratorContext
        self._has_json = False
        self._has_binary = False
        self._check_content()

    def _check_content(self):
        if "$ref" in self._response_obj:
            self._response_obj = self._ctx.get_object_from_ref(self._response_obj["$ref"])
        if "content" in self._response_obj:
            content = self._response_obj["content"]
            if 'application/json' in content:
                self._has_json = True
            else:
                parse_schema = parse("$..schema").find(self._response_obj)
                schema = [s.value for s in parse_schema][0]
                if "$ref" in schema:
                    schema = self._ctx.get_object_from_ref(schema["$ref"])
                if "format" in schema and schema["format"] == "binary":
                    self._has_binary = True


class ControllerRoute(object):
    @property
    def description(self):
        # type: () -> str
        if "description" in self._obj:
            return self._obj["description"]
        return ""

    @property
    def responses(self):
        return self._responses

    @property
    def url(self):
        # type: () -> str
        return self.full_url()

    @property
    def method(self):
        # type: () -> str
        return self._method

    @property
    def operation_name(self):
        # type: () -> str
        name = self._obj['operationId']
        name = util.pascal_case(name)
        return name

    @property
    def route_parameters(self):
        # type: () -> [str]
        return self._parameters

    @property
    def response_model_name(self):
        # type: () -> str
        return self.operation_name + 'Response'

    @property
    def full_responsename(self):
        # type: () -> str
        _ctx = self._ctx # type: GeneratorContext
        return """{models_prefix}{response_model_name}""".format(
            models_prefix=_ctx.models_prefix,
            response_model_name=self.response_model_name
        )

    def __init__(
        self,
        url,
        method,
        methodobj,
        ctx
    ):
        self._ctx = ctx
        self._url = url
        self._method = method.upper()
        self._obj = methodobj
        self._parameters = []
        self._extract_parameters()
        self._responses = []
        self._extract_responses()

    def requestBody(self):
        # type: () -> Component
        _ctx = self._ctx # type: GeneratorContext
        try:
            ref = self._obj['requestBody']['content']['application/json']['schema']['$ref']
            yamlname = ref.split('/')[-1]
            for component in _ctx.components:
                if component.yaml_name == yamlname:
                    return component
            return None
        except KeyError:
            return None
    def full_url(self):
        _ctx = self._ctx  # type: GeneratorContext
        server = _ctx.servers[0]
        if server is None:
            return self._url
        return server.basepath + self._url

    def _extract_parameters(self):
        if "parameters" in self._obj:
            for param in self._obj["parameters"]:
                self._parameters.append(param["name"])

    def _extract_responses(self):
        for response_value, response_obj in self._obj["responses"].items():
            self._responses.append(Responses(response_value, response_obj, self._ctx))


class Controller(object):
    @property
    def controller_name(self):
        # type: () -> str
        name = util.pascal_case(self.yamlname) + 'Controller'
        return name

    @property
    def service_handler_name(self):
        # type: () -> str
        name = util.pascal_case(self.yamlname) + 'Handler'
        return name

    @property
    def yamlname(self):
        # type: () -> str
        return self._yamlname

    def __init__(self, yamlname, ctx):
        self._ctx = ctx
        self.routes = []  # type: [GeneratorContext]
        self._yamlname = yamlname

    def add_route(self, url, method, methodobj):
        self.routes.append(ControllerRoute(url, method, methodobj, self._ctx))


class GeneratorContext(object):
    def __init__(self, openapi):
        self._openapi = openapi
        self.module_path = str()  # type: str
        self.models_prefix = str()  # type: str
        self.models_path = str()  # type: str
        self.output_path = str()  # type: str
        self.servers = []   # type: [Server]
        self.components = []  # type: [Component]
        self.controllers = []  # type: [Controller]

    def find_controller(self, yamlname):
        #  type:(str) -> Controller
        ctrl = None  #type: Controller
        for c in self.controllers:
            if c.yamlname == yamlname:
                ctrl = c
                break
        if ctrl is None:
            ctrl = Controller(yamlname, self)
            self.controllers.append(ctrl)
        return ctrl

    def get_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf


