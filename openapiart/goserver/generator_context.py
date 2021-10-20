import re
import openapiart.goserver.string_util as util

class Server(object):
    
    @property
    def basepath(self) -> [str]:
        return self._basepath

    def __init__(self, serverobj):
        self._obj = serverobj
        self._basepath = ''
        try:
            self._basepath = self._obj['variables']['basePath']['default']
            if self._basepath.startswith("/") == False:
                self._basepath = "/" + self._basepath
        except:
            pass

class Component(object):
    @property
    def yaml_name(self) -> [str]:
        return self._yamlname

    @property
    def model_name(self) -> str:
        return re.sub('[.]', '', self._yamlname)

    @property
    def full_model_name(self) -> str:
        _ctx: GeneratorContext = self._ctx
        return f"{_ctx.models_prefix}{self.model_name}"

    def __init__(
        self,
        yamlname: str,
        componentobj,
        ctx
        ):
        self._ctx: GeneratorContext = ctx
        self._yamlname = yamlname
        self._obj = componentobj


class ControllerRoute(object):
    @property
    def description(self) -> str:
        if "description" in self._obj:
            return self._obj["description"]
        return ""
    @property
    def url(self) -> str:
        return self.full_url()
    @property
    def method(self) -> str:
        return self._method
    @property
    def operation_name(self) -> str:
        name = self._obj['operationId']
        name = util.pascal_case(name)
        return name
    @property
    def route_parameters(self) -> [str]:
        return self._parameters

    @property
    def response_model_name(self) -> str:
        return self.operation_name + 'Response'

    @property
    def full_responsename(self) -> Component:
        _ctx: GeneratorContext = self._ctx
        return f"{_ctx.models_prefix}{self.response_model_name}"

    def __init__(
        self,
        url: str,
        method: str,
        methodobj,
        ctx
        ):
        self._ctx: GeneratorContext = ctx
        self._url = url
        self._method = method.upper()
        self._obj = methodobj
        self._parameters: [str] = []
        self._extract_parameters()

    def requestBody(self) -> Component:
        _ctx: GeneratorContext = self._ctx
        try:
            ref = self._obj['requestBody']['content']['application/json']['schema']['$ref']
            yamlname = ref.split('/')[-1]
            for component in _ctx.components:
                if component.yaml_name == yamlname:
                    return component
            return None
        except:
            return None
    def full_url(self):
        _ctx: GeneratorContext = self._ctx
        server = _ctx.servers[0]
        if server == None:
            return self._url
        return server.basepath + self._url

    def _extract_parameters(self):
        if "parameters" in self._obj:
            for param in self._obj["parameters"]:
                self._parameters.append(param["name"])

class Controller(object):
    @property
    def controller_name(self) -> str:
        name = util.pascal_case(self.yamlname) + 'Controller'
        return name

    @property
    def service_handler_name(self) -> str:
        name = util.pascal_case(self.yamlname) + 'Handler'
        return name

    @property
    def yamlname(self) -> str:
        return self._yamlname

    def __init__(self, yamlname: str, ctx):
        self._ctx = ctx
        self.routes: [ControllerRoute] = []
        self._yamlname = yamlname

    def add_route(self, url: str, method: str, methodobj):
        self.routes.append(ControllerRoute(url, method, methodobj, self._ctx))
        pass

class GeneratorContext(object):
    def __init__(self):
        self.module_path: str
        self.models_prefix: str
        self.models_path: str
        self.output_path: str
        self.servers: [Server] = []
        self.components: [Component] = []
        self.controllers: [Controller] = []

    def find_controller(self, yamlname: str) -> Controller:
        ctrl: Controller = None
        for c in self.controllers:
            if c.yamlname == yamlname:
                ctrl = c
                break
        if ctrl == None:
            ctrl = Controller(yamlname, self)
            self.controllers.append(ctrl)
        return ctrl


