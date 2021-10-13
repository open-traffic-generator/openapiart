import typing
import sys
import os
import importlib
import shutil
import yaml
import subprocess
import requests
import platform
import shutil
import openapiart.goserver.generator_context as ctx
from openapiart.goserver.go_interface_generator import GoServerInterfaceGenerator
from openapiart.goserver.go_controller_generator import GoServerControllerGenerator

class GoServerGenerator(object):
    def __init__(
        self, 
        openapi, # openapi.yaml.yaml
        output_root_path: str,
        module_path: str,
        models_prefix: str = '',
        models_path: str = ''
        ):
        self._output_root_path = output_root_path
        self._openapi = openapi
        self._context = ctx.GeneratorContext()
        self._context.output_path = os.path.join(output_root_path, 'httpapi')
        if len(models_prefix) > 0:
            models_prefix = models_prefix + '.'
        if len(models_path) == 0:
            models_path = module_path
        self._context.module_path = module_path
        self._context.models_prefix = models_prefix
        self._context.models_path = models_path
        print(f'GoServer output directory: {self._context.output_path}')

    def generate(self):
        self._loadyaml()
        self._copy_static_files()
        GoServerInterfaceGenerator(self._context).generate()
        GoServerControllerGenerator(self._context).generate()
        self._tidy_mod_file()

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
            c = ctx.Component(componentname, componentobj)
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

    def _copy_static_files(self):
        output_path = self._context.output_path
        if os.path.exists(output_path) is False:
            os.makedirs(output_path)
        srcfolder = os.path.dirname(__file__)
        name = 'http_setup.go'
        shutil.copyfile(os.path.join(srcfolder, name), os.path.join(output_path, name))
        name = 'response.go'
        shutil.copyfile(os.path.join(srcfolder, name), os.path.join(output_path, name))

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
            process = subprocess.Popen(process_args, cwd=self._output_root_path, shell=False, env=os.environ)
            process.wait()
        except Exception as e:
            print("Bypassed tidying the generated mod file: {}".format(e))



