import os
import shutil
import openapiart.goserver.generator_context as ctx
from openapiart.goserver.go_interface_generator import (
    GoServerInterfaceGenerator,
)
from openapiart.goserver.go_controller_generator import (
    GoServerControllerGenerator,
)


class GoServerGenerator(object):
    def __init__(
        self,
        openapi,  # openapi.yaml.yaml
        output_root_path: str,
        module_path: str,
        models_prefix: str = "",
        models_path: str = "",
    ):
        self._output_root_path = output_root_path
        self._openapi = openapi
        self._context = ctx.GeneratorContext(openapi)
        self._context.output_path = os.path.join(output_root_path, "httpapi")
        if len(models_prefix) > 0:
            models_prefix = models_prefix + "."
        if len(models_path) == 0:
            models_path = module_path
        self._context.module_path = module_path
        self._context.models_prefix = models_prefix
        self._context.models_path = models_path
        print(
            "GoServer output directory: {}".format(self._context.output_path)
        )

    def generate(self):
        self._loadyaml()
        self._copy_static_files()
        GoServerInterfaceGenerator(self._context).generate()
        GoServerControllerGenerator(self._context).generate()

    def _loadyaml(self):
        # load servers
        if "servers" in self._openapi:
            servers = self._openapi["servers"]
            self._load_servers(servers)
        # load components
        if "components" in self._openapi:
            components = self._openapi["components"]
            if "schemas" in components:
                self._load_components(components["schemas"])
        # load routes
        if "paths" in self._openapi:
            for url, pathobj in self._openapi["paths"].items():
                self._loadroute(url, pathobj)

    def _load_servers(self, servers):
        for server in servers:
            s = ctx.Server(server)
            self._context.servers.append(s)

    def _load_components(self, components):
        for componentname, componentobj in components.items():
            c = ctx.Component(componentname, componentobj, self._context)
            self._context.components.append(c)

    def _loadroute(self, url: str, pathobj):
        http_methods = ["get", "post", "put", "delete", "head", "patch"]
        for methodname, methodobj in pathobj.items():
            if methodname not in http_methods:
                continue
            if "tags" not in methodobj:
                raise AttributeError(
                    "controller name missing from '{url} - {methodname}:'\nUse tags: [<name>]".format(
                        url=url, methodname=methodname
                    )
                )
            controllername = methodobj["tags"][0]
            ctrl = self._context.find_controller(controllername)
            ctrl.add_route(url, methodname, methodobj)

    def _copy_static_files(self):
        output_path = self._context.output_path
        if os.path.exists(output_path) is False:
            os.makedirs(output_path)
        srcfolder = os.path.dirname(__file__)
        name = "http_setup.go"
        shutil.copyfile(
            os.path.join(srcfolder, name), os.path.join(output_path, name)
        )
        print("copy: " + os.path.join(output_path, name))
        name = "response.go"
        shutil.copyfile(
            os.path.join(srcfolder, name), os.path.join(output_path, name)
        )
        print("copy: " + os.path.join(output_path, name))
