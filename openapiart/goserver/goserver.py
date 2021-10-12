import typing
import sys
import os
import importlib
import shutil
import yaml
import subprocess
import requests
import platform

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
    def __init__(
        self,
        url: str,
        method: str,
        methodobj
        ):
        self._url = url
        self._method = method
        self._obj = methodobj
        pass

class Controller(object):
    routes: [ControllerRoute] = []
    @property
    def yamlname(self) -> str:
        return self._yamlname

    def __init__(self, yamlname: str):
        self._yamlname = yamlname

    def add_route(self, url: str, method: str, methodobj):
        self.routes.append(ControllerRoute(url, method, methodobj))
        pass

class GeneratorContext(object):
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


class GoServerGenerator(object):
    def __init__(
        self, 
        openapi, # openapi.yaml.yaml
        output_root_path: str
        ):
        self._openapi = openapi
        self.output_path = os.path.join(output_root_path, 'internal')
        self._context = GeneratorContext()
        print(f'GoServer output directory: {self.output_path}')

    def generate(self):
        self._loadyaml()

    def _loadyaml(self):
        # load components
        if "components" in self._openapi:
            components = self._openapi["components"]
            if "schemas" in components:
                self._load_components(components["schemas"])
        # load routes
        if "paths" in self._openapi:
            for url, pathobj in self._openapi["paths"].items():
                self._loadroute(url, pathobj)
        pass
    
    def _load_components(self, components):
        for componentname, componentobj in components.items():
            c = Component(componentname, componentobj)
            self._context.components.append(c)
        pass
    
    def _loadroute(self, url: str, pathobj):
        for methodname, methodobj in pathobj.items():
            if "tags" not in methodobj:
                raise AttributeError(f"controller name missing from '{url} - {methodname}:'\nUse tags: [<name>]")
            controllername = methodobj["tags"][0]
            ctrl = self._context.find_controller(controllername)
            ctrl.add_route(url, methodname, methodobj)
        pass



