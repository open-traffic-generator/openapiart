import openapiart.goserver.string_util as util

class Component(object):
    def __init__(
        self,
        yamlname: str,
        componentobj
        ):
        self._yamlname = yamlname
        self._obj = componentobj
        pass

class ControllerRoute(object):
    @property
    def description(self) -> str:
        if "description" in self._obj:
            return self._obj["description"]
        return ""
    @property
    def url(self) -> str:
        return self._url
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

    def __init__(
        self,
        url: str,
        method: str,
        methodobj
        ):
        self._url = url
        self._method = method.upper()
        self._obj = methodobj
        self._parameters: [str] = []
        self._extract_parameters()
        pass

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
        name = util.pascal_case(self.yamlname) + 'ServiceHandler'
        return name

    @property
    def yamlname(self) -> str:
        return self._yamlname

    def __init__(self, yamlname: str):
        self.routes: [ControllerRoute] = []
        self._yamlname = yamlname

    def add_route(self, url: str, method: str, methodobj):
        self.routes.append(ControllerRoute(url, method, methodobj))
        pass

class GeneratorContext(object):
    module_path: str
    output_path: str
    components: [Component] = []
    controllers: [Controller] = []
    def find_controller(self, yamlname: str) -> Controller:
        ctrl: Controller = None
        for c in self.controllers:
            if c.yamlname == yamlname:
                ctrl = c
                break
        if ctrl == None:
            ctrl = Controller(yamlname)
            self.controllers.append(ctrl)
        return ctrl


