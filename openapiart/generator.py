"""Generator

Given an openapi.yaml file that has been produced by the Bundler class in the
bundler.py file the Generator class will produce an enhanced python ux file.

TBD: 
- packet slicing using constants
- docstrings
- type checking
"""
import sys
import yaml
import os
import subprocess
import re
import requests
import pkgutil
import importlib
from jsonpath_ng import parse
from .openapiartplugin import OpenApiArtPlugin

MODELS_RELEASE = "v0.3.3"


class Generator(object):
    """Generates python classes based on an openapi.yaml file produced by the
    bundler.py infrastructure.
    """

    def __init__(self, openapi_filename, package_name, output_dir=None, extension_prefix=None):
        self._parsers = {}
        self._generated_methods = []
        self._generated_classes = []
        self._generated_top_level_factories = []
        self._openapi_filename = openapi_filename
        self._extension_prefix = extension_prefix
        self.__python = os.path.normpath(sys.executable)
        self.__python_dir = os.path.dirname(self.__python)
        self._src_dir = output_dir
        self._output_dir = os.path.join(output_dir, package_name)
        if os.path.exists(self._output_dir) is False:
            os.mkdir(self._output_dir)
        self._package_name = package_name
        self._output_file = package_name
        self._docs_dir = os.path.join(self._src_dir, "..", "docs")
        self._get_openapi_file()
        # self._plugins = self._load_plugins()

    def _get_parser(self, pattern):
        if pattern not in self._parsers:
            parser = parse(pattern)
            self._parsers[pattern] = parser
        else:
            parser = self._parsers[pattern]
        return parser

    def _load_plugins(self):
        plugins = []
        pkg_dir = os.path.dirname(__file__)
        for (_, name, _) in pkgutil.iter_modules([pkg_dir]):
            module_name = "openapiart." + name
            importlib.import_module(module_name)
            obj = sys.modules[module_name]
            for dir_name in dir(obj):
                if dir_name.startswith("_"):
                    continue
                dir_obj = getattr(obj, dir_name)
                print(dir_obj)
                if issubclass(dir_obj.__class__, OpenApiArtPlugin):
                    plugins.append(dir_obj)
        return plugins

    def _get_openapi_file(self):
        if self._openapi_filename is None:
            OPENAPI_URL = ("https://github.com/open-traffic-generator/models/releases" "/download/%s/openapi.yaml") % MODELS_RELEASE
            response = requests.request("GET", OPENAPI_URL, allow_redirects=True)
            if response.status_code != 200:
                raise Exception("Unable to retrieve the Open Traffic Generator openapi.yaml" " file [%s]" % response.content)
            openapi_content = response.content

            project_dir = os.path.dirname(os.path.dirname(__file__))
            with open(os.path.join(project_dir, "models-release"), "w") as out:
                out.write(MODELS_RELEASE)
        else:
            with open(self._openapi_filename, "rb") as fp:
                openapi_content = fp.read()
        self._openapi = yaml.safe_load(openapi_content)
        self._openapi_version = self._openapi["info"]["version"]
        print("generating using model version %s" % self._openapi_version)

    def generate(self):
        self._api_filename = os.path.join(self._output_dir, self._output_file + ".py")
        with open(os.path.join(os.path.dirname(__file__), "common.py"), "r") as fp:
            common_content = fp.read()
            if re.search(r"def[\s+]api\(", common_content) is not None:
                self._generated_top_level_factories.append("api")
            if self._extension_prefix is not None:
                common_content = common_content.replace(
                    r'"{}_{}".format(__name__, ext)', r'"' + self._extension_prefix + r"_{}." + self._package_name + r'_api".format(ext)'
                )
        with open(self._api_filename, "w") as self._fid:
            self._fid.write(common_content)
        methods, factories = self._get_methods_and_factories()
        self._write_api_class(methods, factories)
        self._write_http_api_class(methods)
        self._write_init()
        return self

    def _write_init(self):
        filename = os.path.join(self._output_dir, "__init__.py")
        with open(filename, "w") as self._fid:
            for class_name in self._generated_classes:
                self._write(0, "from .%s import %s" % (self._output_file, class_name))
            for factory_name in self._generated_top_level_factories:
                self._write(0, "from .%s import %s" % (self._output_file, factory_name))

    def _find(self, path, schema_object):
        finds = self._get_parser(path).find(schema_object)
        for find in finds:
            yield find.value
            self._get_parser(path).find(find.value)

    def _get_methods_and_factories(self):
        """
        Parse methods and top level objects from yaml file to be later used in
        code generation.
        """
        methods = []
        factories = []
        refs = []
        self._top_level_schema_refs = []

        # parse methods
        for path in self._get_api_paths():
            operation = path["operation"]
            method_name = operation["operationId"].replace(".", "_").lower()
            if method_name in self._generated_methods:
                continue
            self._generated_methods.append(method_name)
            print("found method %s" % method_name)

            request = self._get_parser("$..requestBody..schema").find(operation)
            for req in request:
                _, _, _, ref = self._get_object_property_class_names(req.value)
                if ref:
                    refs.append(ref)

            response = self._get_parser("$..responses..schema").find(operation)
            response_type = None
            if len(response) == 0:
                # since some responses currently directly $ref to a schema
                # stored someplace else, we need to go one level deeper to
                # get actual response type (currently extracting only for 200)
                response = self._get_parser('$..responses.."200"').find(operation)
                response_name, _, _, _ = self._get_object_property_class_names(response[0].value)
                if response_name is not None:
                    response = self._get_parser('$.."$ref"').find(self._openapi["components"]["responses"][response_name])
                    if len(response) > 0:
                        _, response_type, _, ref = self._get_object_property_class_names(response[0].value)
                        if ref:
                            refs.append(ref)
            else:
                _, response_type, _, ref = self._get_object_property_class_names(response[0].value)
                if ref:
                    refs.append(ref)

            if response_type is None:
                # TODO: response type is usually None for schema which does not
                # contain any ref (e.g. response of POST /results/capture)
                pass

            methods.append(
                {
                    "name": method_name,
                    "args": ["self"] if len(request) == 0 else ["self", "payload"],
                    "http_method": path["method"],
                    "url": path["url"],
                    "description": self._get_description(operation),
                    "response_type": response_type,
                }
            )

        # parse top level objects (arguments for API requests)
        for ref in refs:
            if ref in self._generated_methods:
                continue
            self._generated_methods.append(ref)
            ret = self._get_object_property_class_names(ref)
            _, property_name, class_name, _ = ret
            schema_object = self._get_object_from_ref(ref)
            if "type" not in schema_object:
                continue
            print("found top level factory method %s" % property_name)
            if schema_object["type"] == "array":
                ref = schema_object["items"]["$ref"]
                _, _, class_name, _ = self._get_object_property_class_names(ref)
                class_name = "%sIter" % class_name
                self._top_level_schema_refs.append((ref, property_name))
            self._top_level_schema_refs.append((ref, None))

            factories.append({"name": property_name, "class_name": class_name})

        for ref, property_name in self._top_level_schema_refs:
            if property_name is None:
                self._write_openapi_object(ref)
            else:
                self._write_openapi_list(ref, property_name)

        return methods, factories

    def _write_http_api_class(self, methods):
        self._generated_classes.append("HttpApi")
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class HttpApi(Api):")
            self._write(1, '"""%s' % "OpenAPI HTTP Api")
            self._write(1, '"""')
            self._write(1, "def __init__(self, **kwargs):")
            self._write(2, "super(HttpApi, self).__init__(**kwargs)")
            self._write(2, "self._transport = HttpTransport(**kwargs)")

            for method in methods:
                print("generating method %s" % method["name"])
                self._write()
                self._write(1, "def %s(%s):" % (method["name"], ", ".join(method["args"])))
                self._write(2, '"""%s %s' % (method["http_method"].upper(), method["url"]))
                self._write(0)
                self._write(2, "%s" % method["description"])
                self._write(0)
                self._write(2, "Return: %s" % method["response_type"])
                self._write(2, '"""')

                self._write(2, "return self._transport.send_recv(")
                self._write(3, '"%s",' % method["http_method"])
                self._write(3, '"%s",' % method["url"])
                self._write(3, "payload=%s," % (method["args"][1] if len(method["args"]) > 1 else "None"))
                self._write(3, "return_object=%s," % ("self." + method["response_type"] + "()" if method["response_type"] else "None"))
                self._write(2, ")")

    def _write_api_class(self, methods, factories):
        self._generated_classes.append("Api")
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class Api(object):")
            self._write(1, '"""%s' % "OpenApi Abstract API")
            self._write(1, '"""')
            self._write()
            self._write(1, "def __init__(self, **kwargs):")
            self._write(2, "pass")

            for method in methods:
                print("generating method %s" % method["name"])
                self._write()
                self._write(1, "def %s(%s):" % (method["name"], ", ".join(method["args"])))
                self._write(2, '"""%s %s' % (method["http_method"].upper(), method["url"]))
                self._write(0)
                self._write(2, "%s" % method["description"])
                self._write(0)
                self._write(2, "Return: %s" % method["response_type"])
                self._write(2, '"""')
                self._write(2, 'raise NotImplementedError("%s")' % method["name"])

            for factory in factories:
                print("generating top level factory method %s" % factory["name"])
                self._write()
                self._write(1, "def %s(self):" % factory["name"])
                self._write(2, '"""Factory method that creates an instance of %s' % (factory["class_name"]))
                self._write()
                self._write(2, "Return: %s" % factory["class_name"])
                self._write(2, '"""')
                self._write(2, "return %s()" % factory["class_name"])

    def _get_object_property_class_names(self, ref):
        """Returns: `Tuple(object_name, property_name, class_name, ref_name)`"""
        object_name = None
        property_name = None
        class_name = None
        ref_name = None
        if isinstance(ref, dict) is True and "$ref" in ref:
            ref_name = ref["$ref"]
        elif isinstance(ref, str) is True:
            ref_name = ref
        if ref_name is not None:
            object_name = ref_name.split("/")[-1]
            property_name = object_name.lower().replace(".", "_")
            class_name = object_name.replace(".", "")
        return (object_name, property_name, class_name, ref_name)

    def _write_openapi_object(self, ref, choice_method_name=None):
        schema_object = self._get_object_from_ref(ref)
        ref_name = ref.split("/")[-1]
        class_name = ref_name.replace(".", "")
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        print("generating class %s" % (class_name))
        refs = []
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class %s(OpenApiObject):" % class_name)
            slots = ["'_parent'"]
            if "choice" in self._get_choice_names(schema_object):
                slots.append("'_choice'")
            self._write(1, "__slots__ = (%s)" % ",".join(slots))
            self._write()

            # write _TYPES definition
            # TODO: this func won't detect whether $ref for a given property is
            # a list because it relies on 'type' attribute to do so
            openapi_types = self._get_openapi_types(schema_object)
            if len(openapi_types) > 0:
                self._write(1, "_TYPES = {")
                for name, value in openapi_types:
                    if len(value) == 1:
                        self._write(2, "'%s': {'%s': %s}," % (name, list(value.keys())[0], list(value.values())[0]))
                        continue
                    self._write(2, "'%s': %s" % (name, "{"))
                    for n, v in value.items():
                        if isinstance(v, list):
                            self._write(3, "'%s': [" % n)
                            for i in v:
                                self._write(4, "'%s'," % i)
                            self._write(3, "],")
                            continue
                        self._write(3, "'%s': %s," % (n, v))
                    self._write(2, "},")
                self._write(1, "} # type: Dict[str, str]")
                self._write()
            else:
                # TODO: provide empty types as workaround because deserializer
                # in common.py currently expects it
                self._write(1, "_TYPES = {} # type: Dict[str, str]")
                self._write()

            required, defaults = self._get_required_and_defaults(schema_object)

            if len(required) > 0:
                self._write(1, "_REQUIRED = {} # type: tuple(str)".format(required))
                self._write()
            else:
                self._write(1, "_REQUIRED= () # type: tuple(str)")
                self._write()

            if len(defaults) > 0:
                self._write(1, "_DEFAULTS = {")
                for name, value in defaults:
                    if isinstance(value, (list, bool, int, float, tuple)):
                        self._write(2, "'%s': %s," % (name, value))
                    else:
                        self._write(2, "'%s': '%s'," % (name, value))
                self._write(1, "} # type: Dict[str, Union(type)]")
                self._write()
            else:
                self._write(1, "_DEFAULTS= {} # type: Dict[str, Union(type)]")
                self._write()

            # write constants
            # search for all simple properties with enum or
            # x-constant and add them here
            for enum in self._get_parser("$..enum | x-constants").find(schema_object):
                for name in enum.value:
                    value = name
                    value_type = "string"
                    if isinstance(enum.value, dict):
                        value = enum.value[name]
                        value_type = enum.context.value["type"] if "type" in enum.context.value else "string"
                    if value_type == "string":
                        self._write(1, "%s = '%s' # type: str" % (name.upper(), value))
                    else:
                        self._write(1, "%s = %s #" % (name.upper(), value))
                if len(enum.value) > 0:
                    self._write()

            # write def __init__(self)
            params = "self, parent=None"
            if "choice" in self._get_choice_names(schema_object):
                params += ", choice=None"
            init_params, properties, _ = self._get_property_param_string(schema_object)
            params = params if len(init_params) == 0 else ", ".join([params, init_params])
            self._write(1, "def __init__(%s):" % (params))
            self._write(2, "super(%s, self).__init__()" % class_name)
            self._write(2, "self._parent = parent")
            for property_name in properties:
                self._write(2, "self._set_property('%s', %s)" % (property_name, property_name))
            if "choice" in self._get_choice_names(schema_object):
                self._write(2, "if 'choice' in self._DEFAULTS and choice is None:")
                self._write(3, "getattr(self, self._DEFAULTS['choice'])")
                self._write(2, "else:")
                self._write(3, "self.choice = choice")

            # process properties - TBD use this one level up to process
            # schema, in requestBody, Response and also
            refs = self._process_properties(class_name, schema_object, choice_child=choice_method_name is not None)

        # descend into child properties
        for ref in refs:
            self._write_openapi_object(ref[0], ref[3])
            if ref[1] is True:
                self._write_openapi_list(ref[0], ref[2])

    def _get_simple_type_names(self, schema_object):
        simple_type_names = []
        if "properties" in schema_object:
            choice_names = self._get_choice_names(schema_object)
            for name in schema_object["properties"]:
                if name in choice_names:
                    continue
                ref = self._get_parser("$..'$ref'").find(schema_object["properties"][name])
                if len(ref) == 0:
                    simple_type_names.append(name)
        return simple_type_names

    def _get_choice_names(self, schema_object):
        choice_names = []
        if "properties" in schema_object and "choice" in schema_object["properties"]:
            choice_names = schema_object["properties"]["choice"]["enum"][:]
            choice_names.append("choice")
        return choice_names

    def _process_properties(self, class_name=None, schema_object=None, choice_child=False):
        """Process all properties of a /component/schema object
        Write a factory method for all choice
        If there are no properties then the schema_object is a primitive or array type
        """
        refs = []
        if "properties" in schema_object:
            choice_names = self._get_choice_names(schema_object)
            excluded_property_names = []
            for choice_name in choice_names:
                if "$ref" not in schema_object["properties"][choice_name]:
                    continue
                ref = schema_object["properties"][choice_name]["$ref"]
                self._write_factory_method(None, choice_name, ref)
                excluded_property_names.append(choice_name)
            for property_name in schema_object["properties"]:
                if property_name in excluded_property_names:
                    continue
                property = schema_object["properties"][property_name]
                write_set_choice = property_name in choice_names and property_name != "choice"
                self._write_openapi_property(schema_object, property_name, property, write_set_choice)
            for property_name, property in schema_object["properties"].items():
                ref = self._get_parser("$..'$ref'").find(property)
                if len(ref) > 0:
                    restriction = self._get_type_restriction(property)
                    choice_name = property_name if property_name in excluded_property_names else None
                    refs.append((ref[0].value, restriction.startswith("List["), property_name, choice_name))
        return refs

    def _write_openapi_list(self, ref, property_name):
        """This is the class writer for schema object properties that are of
        type array with a ref to an object.  The class should provide a factory
        method for the encapsulated ref.
        ```
        properties:
          ports:
            type: array
            items:
              $ref: '#/components/schema/...'
        ```

        If the schema object has a property named choice, that property needs
        to be brought forward so that the generated class can provide factory
        methods for objects for each of the choice $refs (if any).

        if choice exists:
            for each choice enum that is a $ref:
                generate a factory method named after the choice
                in the method set the choice property
        """
        yobject = self._get_object_from_ref(ref)
        ref_name = ref.split("/")[-1]
        contained_class_name = ref_name.replace(".", "")
        class_name = "%sIter" % contained_class_name
        if class_name in self._generated_classes:
            return
        self._generated_classes.append(class_name)

        self._imports = []
        print("generating class %s" % (class_name))
        with open(self._api_filename, "a") as self._fid:
            self._write()
            self._write()
            self._write(0, "class %s(OpenApiIter):" % class_name)
            self._write(1, "__slots__ = ('_parent', '_choice')")
            self._write()

            # if all choice(s) are $ref, the getitem should return the actual choice object
            # the _GETITEM_RETURNS_CHOICE_OBJECT class static allows the OpenApiIter to
            # correctly return the selected choice if any
            get_item_returns_choice = True
            if "properties" in yobject and "choice" in yobject["properties"]:
                for property, item in yobject["properties"].items():
                    if property == "choice":
                        continue
                    elif property in yobject["properties"]["choice"]["enum"] and len(self._get_parser('$.."$ref"').find(item)) > 0:
                        continue
                    get_item_returns_choice = False
                    break
            else:
                get_item_returns_choice = False
            self._write(1, "_GETITEM_RETURNS_CHOICE_OBJECT = {}".format(get_item_returns_choice))

            self._write()
            self._write(1, "def __init__(self, parent=None, choice=None):")
            self._write(2, "super(%s, self).__init__()" % class_name)
            self._write(2, "self._parent = parent")
            self._write(2, "self._choice = choice")

            # write container emulation methods __getitem__, __iter__, __next__
            self._write_openapilist_special_methods(contained_class_name, yobject)

            # write a factory method for the schema object in the list that returns the container
            self._write_factory_method(contained_class_name, ref_name.lower().split(".")[-1], ref, True, False)

            # write an add method for the schema object in the list that creates and returns the new object
            self._write_add_method(yobject, ref, False, class_name, contained_class_name, class_name)

            # write choice factory methods if the only properties are choice properties
            if get_item_returns_choice is True:
                for property, item in yobject["properties"].items():
                    if property == "choice":
                        continue
                    self._write_factory_method(contained_class_name, property, item["$ref"], True, True)

        return class_name

    def _write_openapilist_special_methods(self, contained_class_name, schema_object):
        get_item_class_names = [contained_class_name]
        if "properties" in schema_object and "choice" in schema_object["properties"]:
            for property in schema_object["properties"]:
                if property in schema_object["properties"]["choice"]["enum"]:
                    if "$ref" in schema_object["properties"][property]:
                        ref = schema_object["properties"][property]["$ref"]
                        _, _, choice_class_name, _ = self._get_object_property_class_names(ref)
                        if choice_class_name not in get_item_class_names:
                            get_item_class_names.append(choice_class_name)
        get_item_class_names.sort()
        self._write()
        self._write(1, "def __getitem__(self, key):")
        self._write(2, "# type: (str) -> Union[%s]" % (", ".join(get_item_class_names)))
        self._write(2, "return self._getitem(key)")
        self._write()
        self._write(1, "def __iter__(self):")
        self._write(2, "# type: () -> %sIter" % contained_class_name)
        self._write(2, "return self._iter()")
        self._write()
        self._write(1, "def __next__(self):")
        self._write(2, "# type: () -> %s" % contained_class_name)
        self._write(2, "return self._next()")
        self._write()
        self._write(1, "def next(self):")
        self._write(2, "# type: () -> %s" % contained_class_name)
        self._write(2, "return self._next()")

    def _write_factory_method(self, contained_class_name, method_name, ref, openapi_list=False, choice_method=False):
        yobject = self._get_object_from_ref(ref)
        _, _, class_name, _ = self._get_object_property_class_names(ref)
        param_string, properties, type_string = self._get_property_param_string(yobject)
        self._write()
        if openapi_list is True:
            self._imports.append("from .%s import %s" % (class_name.lower(), class_name))
            params = "self" if len(param_string) == 0 else ", ".join(["self", param_string])
            self._write(1, "def %s(%s):" % (method_name, params))
            return_class_name = class_name
            if contained_class_name is not None:
                return_class_name = "{}Iter".format(contained_class_name)
            self._write(2, "# type: (%s) -> %s" % (type_string, return_class_name))
            self._write(2, '"""Factory method that creates an instance of the %s class' % (class_name))
            self._write()
            self._write(2, "%s" % self._get_description(yobject))
            self._write()
            self._write(2, "Returns: %s" % (return_class_name))
            self._write(2, '"""')
            if choice_method is True:
                self._write(2, "item = %s()" % (contained_class_name))
                self._write(2, "item.%s" % (method_name))
                self._write(2, "item.choice = '%s'" % (method_name))
            else:
                params = ["parent=self._parent"]
                if "properties" in yobject and "choice" in yobject["properties"]:
                    params.append("choice=self._choice")
                params.extend(["%s=%s" % (name, name) for name in properties])
                self._write(2, "item = %s(%s)" % (class_name, ", ".join(params)))
            self._write(2, "self._add(item)")
            self._write(2, "return self")
            self._write()
        else:
            self._write(1, "@property")
            self._write(1, "def %s(self):" % (method_name))
            self._write(2, "# type: () -> %s" % (class_name))
            self._write(2, '"""Factory property that returns an instance of the %s class' % (class_name))
            self._write()
            self._write(2, "%s" % self._get_description(yobject))
            self._write()
            self._write(2, "Returns: %s" % (class_name))
            self._write(2, '"""')
            self._write(2, "return self._get_property('%s', %s, self, '%s')" % (method_name, class_name, method_name))

    def _write_add_method(self, yobject, ref, choice_method, class_name, contained_class_name, return_class_name):
        """Writes an add method"""
        method_name = ref.lower().split("/")[-1]
        self._imports.append("from .%s import %s" % (contained_class_name.lower(), contained_class_name))
        param_string, properties, type_string = self._get_property_param_string(yobject)
        params = "self" if len(param_string) == 0 else ", ".join(["self", param_string])
        self._write(1, "def add(%s):" % (params))
        self._write(2, "# type: (%s) -> %s" % (type_string, contained_class_name))
        self._write(2, '"""Add method that creates and returns an instance of the %s class' % (contained_class_name))
        self._write()
        self._write(2, "%s" % self._get_description(yobject))
        self._write()
        self._write(2, "Returns: %s" % (contained_class_name))
        self._write(2, '"""')
        if choice_method is True:
            self._write(2, "item = self.%s()" % (method_name))
            self._write(2, "item.%s" % (contained_class_name))
            self._write(2, "item.choice = '%s'" % (contained_class_name))
        else:
            params = ["parent=self._parent"]
            if "properties" in yobject and "choice" in yobject["properties"]:
                params.append("choice=self._choice")
            params.extend(["%s=%s" % (name, name) for name in properties])
            self._write(2, "item = %s(%s)" % (contained_class_name, ", ".join(params)))
        self._write(2, "self._add(item)")
        self._write(2, "return item")

    def _get_property_param_string(self, yobject):
        property_param_string = []
        property_type_string = []
        properties = []
        if "properties" in yobject:
            for name, property in yobject["properties"].items():
                if name == "choice":
                    continue
                default = None
                type_string = self._get_type_restriction(property)
                if "obj" not in type_string:
                    if "default" in property:
                        default = property["default"]
                    if name == "choice":
                        val = "None"
                    elif property["type"] in ["number", "integer", "boolean", "array"]:
                        val = "None" if default is None else default
                    else:
                        val = "None" if default is None else "'{}'".format(default.strip())
                    properties.append(name)
                    property_param_string.append("%s=%s" % (name, val))
                    property_type_string.append(type_string)
        types = ",".join(property_type_string)
        return (", ".join(property_param_string), properties, types)

    def _write_openapi_property(self, schema_object, name, property, write_set_choice=False):
        ref = self._get_parser("$..'$ref'").find(property)
        restriction = self._get_type_restriction(property)
        if len(ref) > 0:
            object_name = ref[0].value.split("/")[-1]
            class_name = object_name.replace(".", "")
            if restriction.startswith("List["):
                type_name = "%sIter" % class_name
            else:
                type_name = class_name
        else:
            type_name = restriction
        self._write()
        self._write(1, "@property")
        self._write(1, "def %s(self):" % name)
        self._write(2, "# type: () -> %s" % (type_name))
        self._write(2, '"""%s getter' % (name))
        self._write()
        self._write(2, self._get_description(property))
        self._write()
        self._write(2, "Returns: %s" % type_name)
        self._write(2, '"""')
        if len(self._get_parser("$..'type'").find(property)) > 0 and len(ref) == 0:
            self._write(2, "return self._get_property('%s')" % (name))
            self._write()
            self._write(1, "@%s.setter" % name)
            self._write(1, "def %s(self, value):" % name)
            self._write(2, '"""%s setter' % (name))
            self._write()
            self._write(2, self._get_description(property))
            self._write()
            self._write(2, "value: %s" % restriction)
            self._write(2, '"""')
            if write_set_choice is True:
                self._write(2, "self._set_property('%s', value, '%s')" % (name, name))
            else:
                self._write(2, "self._set_property('%s', value)" % (name))
        elif len(ref) > 0:
            if restriction.startswith("List["):
                self._write(2, "return self._get_property('%s', %sIter, self._parent, self._choice)" % (name, class_name))
            else:
                self._write(2, "return self._get_property('%s', %s)" % (name, class_name))

    def _get_description(self, yobject):
        if "description" not in yobject:
            yobject["description"] = "TBD"
        # remove tabs, multiple spaces
        description = re.sub(r"\n", ". ", yobject["description"])
        description = re.sub(r"\s+", " ", description)
        return description
        # doc_string = []
        # for line in re.split('\. ', description):
        #     line = re.sub('\.$', '', line)
        #     if len(line) > 0:
        #         doc_string.append('%s  ' % line)
        # return doc_string

    def _get_data_types(self, yproperty):
        data_type_map = {"integer": "int", "string": "str", "boolean": "bool", "array": "list", "number": "float", "float": "float", "double": "float"}
        if yproperty["type"] in data_type_map:
            return data_type_map[yproperty["type"]]
        else:
            return yproperty["type"]

    def _get_openapi_types(self, yobject):
        types = []
        if "properties" in yobject:
            for name in yobject["properties"]:
                yproperty = yobject["properties"][name]
                ref = self._get_parser("$..'$ref'").find(yproperty)
                pt = {}
                if "type" in yproperty:
                    pt.update({"type": self._get_data_types(yproperty)})
                    pt.update({"enum": yproperty["enum"]}) if "enum" in yproperty else None
                    pt.update({"format": "'%s'" % yproperty["format"]}) if "format" in yproperty else None
                if len(ref) > 0:
                    object_name = ref[0].value.split('/')[-1]
                    class_name = object_name.replace('.', '')
                    if 'type' in yproperty and yproperty['type'] == 'array':
                        class_name += 'Iter'
                    pt.update({'type': "\'%s\'" % class_name})
                if len(ref) == 0 and 'items' in yproperty and 'type' in yproperty['items']:
                    pt.update({'itemtype': self._get_data_types(yproperty['items'])})
                if len(ref) == 0 and 'minimum' in yproperty:
                    pt.update({"minimum": yproperty['minimum']})
                if len(ref) == 0 and 'maximum' in yproperty:
                    pt.update({"maximum": yproperty['maximum']})
                if len(pt) > 0:
                    types.append((name, pt))

        return types

    def _get_required_and_defaults(self, yobject):
        required = []
        defaults = []
        if "required" in yobject:
            required = yobject["required"]
        if "properties" in yobject:
            for name in yobject["properties"]:
                yproperty = yobject["properties"][name]
                if "default" in yproperty:
                    default = yproperty["default"]
                    if "type" in yproperty and yproperty["type"] == "number":
                        default = float(default)
                    defaults.append((name, default))
        return (tuple(required), defaults)

    def _get_default_value(self, property):
        if "default" in property:
            return property["default"]
        property_type = property["type"]
        if property_type == "array":
            return "[]"
        if property_type == "string":
            return "''"
        if property_type == "integer":
            return 0
        if property_type == "number":
            return 0
        if property_type == "bool":
            return False
        raise Exception("Missing handler for property type `%s`" % property_type)

    def _get_api_paths(self):
        paths = []
        for url, yobject in self._openapi["paths"].items():
            for method in yobject:
                if method.lower() in ["get", "post", "put", "patch", "delete"]:
                    paths.append(
                        {
                            "url": url,
                            "method": method,
                            "operation": yobject[method],
                        }
                    )
        return paths

    def _write_data_properties(self, schema, classname, choice_tuples):
        if len(choice_tuples) > 0:
            choices = []
            for choice_tuple in choice_tuples:
                choices.append(choice_tuple[0])
            self._write(2, "if isinstance(choice, (%s)) is False:" % (", ".join(choices)))
            self._write(3, "raise TypeError('choice must be of type: %s')" % (", ".join(choices)))
            self._write(2, "self.__setattr__('choice', %s._CHOICE_MAP[type(choice).__name__])" % classname)
            self._write(2, "self.__setattr__(%s._CHOICE_MAP[type(choice).__name__], choice)" % classname)

        if "properties" in schema:
            for name, property in schema["properties"].items():
                if len([item for item in choice_tuples if item[1] == name]) == 0 and name != "choice":
                    restriction = self._get_isinstance_restriction(schema, name, property)
                    self._write(2, "if isinstance(%s, %s) is True:" % (name, restriction))
                    if restriction == "(list, type(None))":
                        self._write(3, "self.%s = [] if %s is None else list(%s)" % (name, name, name))
                    else:
                        if "pattern" in property:
                            self._write(3, "import re")
                            self._write(3, "assert(bool(re.match(r'%s', %s)) is True)" % (property["pattern"], name))
                        self._write(3, "self.%s = %s" % (name, name))
                    self._write(2, "else:")
                    self._write(3, "raise TypeError('%s must be an instance of %s')" % (name, restriction))

    def _get_isinstance_restriction(self, schema, name, property):
        type_none = ", type(None)"
        if "required" in schema and name in schema["required"]:
            type_none = ""
        if "$ref" in property:
            return "(%s%s)" % (self._get_classname_from_ref(property["$ref"]), type_none)
        elif name == "additionalProperties":
            return "**additional_properties"
        elif property["type"] in ["number", "integer"]:
            return "(float, int%s)" % type_none
        elif property["type"] == "string":
            return "(str%s)" % type_none
        elif property["type"] == "array":
            return "(list%s)" % type_none
        elif property["type"] == "boolean":
            return "(bool%s)" % type_none

    def _get_type_restriction(self, property):
        try:
            if "$ref" in property:
                ref_obj = self._get_object_from_ref(property["$ref"])
                description = ""
                if "description" in ref_obj:
                    description = ref_obj["description"]
                if "description" in property:
                    description += property["description"]
                property["description"] = description
                class_name = property["$ref"].split("/")[-1].replace(".", "")
                return "obj(%s)" % class_name
            elif "oneOf" in property:
                return "Union[%s]" % ",".join([item["type"] for item in property["oneOf"]])
            elif property["type"] == "number":
                return "float"
            elif property["type"] == "integer":
                return "int"
            elif property["type"] == "string":
                if "enum" in property:
                    values = property["enum"]
                    values.sort()
                    values = ['Literal["{}"]'.format(s) for s in values]
                    return "Union[%s]" % ", ".join(values)
                else:
                    return "str"
            elif property["type"] == "array":
                return "List[%s]" % self._get_type_restriction(property["items"])
            elif property["type"] == "boolean":
                return "bool"
        except Exception as e:
            print("Error ", property, e)
            raise e

    def _get_object_from_ref(self, ref):
        leaf = self._openapi
        for attr in ref.split("/")[1:]:
            leaf = leaf[attr]
        return leaf

    def _get_classname_from_ref(self, ref):
        final_piece = ref.split("/")[-1]
        if "." in final_piece:
            return final_piece.split(".")[-1]
        else:
            return final_piece

    def _write(self, indent=0, line=""):
        self._fid.write("    " * indent + line + "\n")
