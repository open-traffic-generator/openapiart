# OpenAPIArt Test API 0.0.1
# License: NO-LICENSE-PRESENT

import importlib
import logging
import json
import platform
import yaml
import requests
import urllib3
import io
import sys
import time
import grpc
import semantic_version
import types
import platform
from google.protobuf import json_format

try:
    from pyapi import openapi_pb2_grpc as pb2_grpc
except ImportError:
    import openapi_pb2_grpc as pb2_grpc
try:
    from pyapi import openapi_pb2 as pb2
except ImportError:
    import openapi_pb2 as pb2

try:
    from typing import Union, Dict, List, Any, Literal
except ImportError:
    from typing_extensions import Literal


if sys.version_info[0] == 3:
    unicode = str


openapi_warnings = []


class Transport:
    HTTP = "http"
    GRPC = "grpc"


def api(
    location=None,
    transport=None,
    verify=True,
    logger=None,
    loglevel=logging.INFO,
    ext=None,
    version_check=False,
):
    """Create an instance of an Api class

    generator.Generator outputs a base Api class with the following:
    - an abstract method for each OpenAPI path item object
    - a concrete properties for each unique OpenAPI path item parameter.

    generator.Generator also outputs an HttpApi class that inherits the base
    Api class, implements the abstract methods and uses the common HttpTransport
    class send_recv method to communicate with a REST based server.

    Args
    ----
    - location (str): The location of an Open Traffic Generator server.
    - transport (enum["http", "grpc"]): Transport Type
    - verify (bool): Verify the server's TLS certificate, or a string, in which
      case it must be a path to a CA bundle to use. Defaults to `True`.
      When set to `False`, requests will accept any TLS certificate presented by
      the server, and will ignore hostname mismatches and/or expired
      certificates, which will make your application vulnerable to
      man-in-the-middle (MitM) attacks. Setting verify to `False`
      may be useful during local development or testing.
    - logger (logging.Logger): A user defined logging.logger, if none is provided
      then a default logger with a stdout handler will be provided
    - loglevel (logging.loglevel): The logging package log level.
      The default loglevel is logging.INFO
    - ext (str): Name of an extension package
    """
    params = locals()
    transport_types = ["http", "grpc"]
    if ext is None:
        transport = "http" if transport is None else transport
        if transport not in transport_types:
            raise Exception(
                "{transport} is not within valid transport types {transport_types}".format(
                    transport=transport, transport_types=transport_types
                )
            )
        if transport == "http":
            return HttpApi(**params)
        else:
            return GrpcApi(**params)
    try:
        if transport is not None:
            raise Exception(
                "ext and transport are not mutually exclusive. Please configure one of them."
            )
        lib = importlib.import_module("openapi_{}.pyapi_api".format(ext))
        return lib.Api(**params)
    except ImportError as err:
        msg = "Extension %s is not installed or invalid: %s"
        raise Exception(msg % (ext, err))


class HttpTransport(object):
    def __init__(self, **kwargs):
        """Use args from api() method to instantiate an HTTP transport"""
        self.location = (
            kwargs["location"]
            if "location" in kwargs and kwargs["location"] is not None
            else "https://localhost:443"
        )
        self.verify = kwargs["verify"] if "verify" in kwargs else False
        self.logger = kwargs["logger"] if "logger" in kwargs else None
        self.loglevel = (
            kwargs["loglevel"] if "loglevel" in kwargs else logging.DEBUG
        )
        if self.logger is None:
            stdout_handler = logging.StreamHandler(sys.stdout)
            formatter = logging.Formatter(
                fmt="%(asctime)s [%(name)s] [%(levelname)s] %(message)s",
                datefmt="%Y-%m-%d %H:%M:%S",
            )
            formatter.converter = time.gmtime
            stdout_handler.setFormatter(formatter)
            self.logger = logging.Logger(self.__module__, level=self.loglevel)
            self.logger.addHandler(stdout_handler)
        self.logger.debug(
            "HttpTransport args: {}".format(
                ", ".join(["{}={!r}".format(k, v) for k, v in kwargs.items()])
            )
        )
        self.set_verify(self.verify)
        self._session = requests.Session()

    def set_verify(self, verify):
        self.verify = verify
        if self.verify is False:
            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
            self.logger.warning("Certificate verification is disabled")

    def _parse_response_error(self, response_code, response_text):
        error_response = ""
        try:
            error_response = yaml.safe_load(response_text)
        except Exception as _:
            error_response = response_text

        err_obj = Error()
        try:
            err_obj.deserialize(error_response)
        except Exception as _:
            err_obj.code = response_code
            err_obj.errors = [str(error_response)]

        raise Exception(err_obj)

    def send_recv(
        self,
        method,
        relative_url,
        payload=None,
        return_object=None,
        headers=None,
        request_class=None,
    ):
        url = "%s%s" % (self.location, relative_url)
        data = None
        headers = headers or {"Content-Type": "application/json"}
        if payload is not None:
            if isinstance(payload, bytes):
                data = payload
                headers["Content-Type"] = "application/octet-stream"
            elif isinstance(payload, (str, unicode)):
                if request_class is not None:
                    request_class().deserialize(payload)
                data = payload
            elif isinstance(payload, OpenApiBase):
                data = payload.serialize()
            else:
                raise Exception("Type of payload provided is unknown")
        response = self._session.request(
            method=method,
            url=url,
            data=data,
            verify=False,
            allow_redirects=True,
            # TODO: add a timeout here
            headers=headers,
        )
        if response.ok:
            if "application/json" in response.headers["content-type"]:
                # TODO: we might want to check for utf-8 charset and decode
                # accordingly, but current impl works for now
                response_dict = yaml.safe_load(response.text)
                if return_object is None:
                    # if response type is not provided, return dictionary
                    # instead of python object
                    return response_dict
                else:
                    return return_object.deserialize(response_dict)
            elif (
                "application/octet-stream" in response.headers["content-type"]
            ):
                return io.BytesIO(response.content)
            else:
                # TODO: for now, return bare response object for unknown
                # content types
                return response
        else:
            self._parse_response_error(response.status_code, response.text)


class OpenApiStatus:
    messages = {}
    # logger = logging.getLogger(__module__)

    @classmethod
    def warn(cls, key, object):
        if cls.messages.get(key) is not None:
            if cls.messages[key] in object.__warnings__:
                return
            # cls.logger.warning(cls.messages[key])
            logging.warning(cls.messages[key])
            object.__warnings__.append(cls.messages[key])
            # openapi_warnings.append(cls.messages[key])

    @staticmethod
    def deprecated(func_or_data):
        def inner(self, *args, **kwargs):
            OpenApiStatus.warn(
                "{}.{}".format(type(self).__name__, func_or_data.__name__),
                self,
            )
            return func_or_data(self, *args, **kwargs)

        if isinstance(func_or_data, types.FunctionType):
            return inner
        OpenApiStatus.warn(func_or_data)

    @staticmethod
    def under_review(func_or_data):
        def inner(self, *args, **kwargs):
            OpenApiStatus.warn(
                "{}.{}".format(type(self).__name__, func_or_data.__name__),
                self,
            )
            return func_or_data(self, *args, **kwargs)

        if isinstance(func_or_data, types.FunctionType):
            return inner
        OpenApiStatus.warn(func_or_data)


class OpenApiBase(object):
    """Base class for all generated classes"""

    JSON = "json"
    YAML = "yaml"
    DICT = "dict"

    __slots__ = ()

    __constraints__ = {"global": []}
    __validate_latter__ = {"unique": [], "constraint": []}

    def __init__(self):
        pass

    def serialize(self, encoding=JSON):
        """Serialize the current object according to a specified encoding.

        Args
        ----
        - encoding (str[json, yaml, dict]): The object will be recursively
            serialized according to the specified encoding.
            The supported encodings are json, yaml and python dict.

        Returns
        -------
        - obj(Union[str, dict]): A str or dict object depending on the specified
            encoding. The json and yaml encodings will return a str object and
            the dict encoding will return a python dict object.
        """
        # TODO: restore behavior
        # self._clear_globals()
        if encoding == OpenApiBase.JSON:
            data = json.dumps(self._encode(), indent=2, sort_keys=True)
        elif encoding == OpenApiBase.YAML:
            data = yaml.safe_dump(self._encode())
        elif encoding == OpenApiBase.DICT:
            data = self._encode()
        else:
            raise NotImplementedError("Encoding %s not supported" % encoding)
        # TODO: restore behavior
        # self._validate_coded()
        return data

    def _encode(self):
        raise NotImplementedError()

    def deserialize(self, serialized_object):
        """Deserialize a python object into the current object.

        If the input `serialized_object` does not match the current
        openapi object an exception will be raised.

        Args
        ----
        - serialized_object (Union[str, dict]): The object to deserialize.
            If the serialized_object is of type str then the internal encoding
            of the serialized_object must be json or yaml.

        Returns
        -------
        - obj(OpenApiObject): This object with all the
            serialized_object deserialized within.
        """
        # TODO: restore behavior
        # self._clear_globals()
        if isinstance(serialized_object, (str, unicode)):
            serialized_object = yaml.safe_load(serialized_object)
        self._decode(serialized_object)
        # TODO: restore behavior
        # self._validate_coded()
        return self

    def _decode(self, dict_object):
        raise NotImplementedError()

    def warnings(self):
        warns = list(self.__warnings__)
        if "2.7" in platform.python_version().rsplit(".", 1)[0]:
            del self.__warnings__[:]
        else:
            self.__warnings__.clear()
        return warns


class OpenApiValidator(object):

    __slots__ = ()

    _validation_errors = []

    def __init__(self):
        pass

    def _clear_errors(self):
        if "2.7" in platform.python_version().rsplit(".", 1)[0]:
            del self._validation_errors[:]
        else:
            self._validation_errors.clear()

    def validate_mac(self, mac):
        if (
            mac is None
            or not isinstance(mac, (str, unicode))
            or mac.count(" ") != 0
        ):
            return False
        try:
            if len(mac) != 17:
                return False
            return all([0 <= int(oct, 16) <= 255 for oct in mac.split(":")])
        except Exception:
            return False

    def validate_ipv4(self, ip):
        if (
            ip is None
            or not isinstance(ip, (str, unicode))
            or ip.count(" ") != 0
        ):
            return False
        if len(ip.split(".")) != 4:
            return False
        try:
            return all([0 <= int(oct) <= 255 for oct in ip.split(".", 3)])
        except Exception:
            return False

    def validate_ipv6(self, ip):
        if ip is None or not isinstance(ip, (str, unicode)):
            return False
        ip = ip.strip()
        if (
            ip.count(" ") > 0
            or ip.count(":") > 7
            or ip.count("::") > 1
            or ip.count(":::") > 0
        ):
            return False
        if (ip[0] == ":" and ip[:2] != "::") or (
            ip[-1] == ":" and ip[-2:] != "::"
        ):
            return False
        if ip.count("::") == 0 and ip.count(":") != 7:
            return False
        if ip == "::":
            return True
        if ip[:2] == "::":
            ip = ip.replace("::", "0:")
        elif ip[-2:] == "::":
            ip = ip.replace("::", ":0")
        else:
            ip = ip.replace("::", ":0:")
        try:
            return all(
                [
                    True
                    if (0 <= int(oct, 16) <= 65535) and (1 <= len(oct) <= 4)
                    else False
                    for oct in ip.split(":")
                ]
            )
        except Exception:
            return False

    def validate_hex(self, hex):
        if hex is None or not isinstance(hex, (str, unicode)):
            return False
        try:
            int(hex, 16)
            return True
        except Exception:
            return False

    def validate_integer(self, value, min, max, type_format=None):
        if value is None or not isinstance(value, int):
            return False
        if min is not None and value < min:
            return False
        if max is not None and value > max:
            return False
        if type_format is not None:
            if type_format == "uint32" and (value < 0 or value > 4294967295):
                return False
            elif type_format == "uint64" and (
                value < 0 or value > 18446744073709551615
            ):
                return False
            elif type_format == "int32" and (
                value < -2147483648 or value > 2147483647
            ):
                return False
            elif type_format == "int64" and (
                value < -9223372036854775808 or value > 9223372036854775807
            ):
                return False
        return True

    def validate_float(self, value):
        return isinstance(value, (int, float))

    def validate_string(self, value, min_length, max_length):
        if value is None or not isinstance(value, (str, unicode)):
            return False
        if min_length is not None and len(value) < min_length:
            return False
        if max_length is not None and len(value) > max_length:
            return False
        return True

    def validate_bool(self, value):
        return isinstance(value, bool)

    def validate_list(self, value, itemtype, min, max, min_length, max_length):
        if value is None or not isinstance(value, list):
            return False
        v_obj = getattr(self, "validate_{}".format(itemtype), None)
        if v_obj is None:
            raise AttributeError(
                "{} is not a valid attribute".format(itemtype)
            )
        v_obj_lst = []
        for item in value:
            if itemtype == "integer":
                v_obj_lst.append(v_obj(item, min, max))
            elif itemtype == "string":
                v_obj_lst.append(v_obj(item, min_length, max_length))
            else:
                v_obj_lst.append(v_obj(item))
        return v_obj_lst

    def validate_binary(self, value):
        if value is None or not isinstance(value, (str, unicode)):
            return False
        return all(
            [
                True if int(bin) == 0 or int(bin) == 1 else False
                for bin in value
            ]
        )

    def types_validation(
        self,
        value,
        type_,
        err_msg,
        itemtype=None,
        min=None,
        max=None,
        min_length=None,
        max_length=None,
    ):
        type_map = {
            int: "integer",
            str: "string",
            float: "float",
            bool: "bool",
            list: "list",
            "int64": "integer",
            "int32": "integer",
            "uint64": "integer",
            "uint32": "integer",
            "double": "float",
        }
        type_format = type_
        if type_ in type_map:
            type_ = type_map[type_]
        if itemtype is not None and itemtype in type_map:
            itemtype = type_map[itemtype]
        v_obj = getattr(self, "validate_{}".format(type_), None)
        if v_obj is None:
            msg = "{} is not a valid or unsupported format".format(type_)
            raise TypeError(msg)
        if type_ == "list":
            verdict = v_obj(value, itemtype, min, max, min_length, max_length)
            if all(verdict) is True:
                return
            err_msg = "{} \n {} are not valid".format(
                err_msg,
                [
                    value[index]
                    for index, item in enumerate(verdict)
                    if item is False
                ],
            )
            verdict = False
        elif type_ == "integer":
            verdict = v_obj(value, min, max, type_format)
            if verdict is True:
                return
            min_max = ""
            if min is not None:
                min_max = ", expected min {}".format(min)
            if max is not None:
                min_max = min_max + ", expected max {}".format(max)
            err_msg = "{} \n got {} of type {} {}".format(
                err_msg, value, type(value), min_max
            )
        elif type_ == "string":
            verdict = v_obj(value, min_length, max_length)
            if verdict is True:
                return
            msg = ""
            if min_length is not None:
                msg = ", expected min {}".format(min_length)
            if max_length is not None:
                msg = msg + ", expected max {}".format(max_length)
            err_msg = "{} \n got {} of type {} {}".format(
                err_msg, value, type(value), msg
            )
        else:
            verdict = v_obj(value)
        if verdict is False:
            raise TypeError(err_msg)

    def _validate_unique_and_name(self, name, value, latter=False):
        if self._TYPES[name].get("unique") is None or value is None:
            return
        if latter is True:
            self.__validate_latter__["unique"].append(
                (self._validate_unique_and_name, name, value)
            )
            return
        class_name = type(self).__name__
        unique_type = self._TYPES[name]["unique"]
        if class_name not in self.__constraints__:
            self.__constraints__[class_name] = dict()
        if unique_type == "global":
            values = self.__constraints__["global"]
        else:
            values = self.__constraints__[class_name]
        if value in values:
            self._validation_errors.append(
                "{} with {} already exists".format(name, value)
            )
            return
        if isinstance(values, list):
            values.append(value)
        self.__constraints__[class_name].update({value: self})

    def _validate_constraint(self, name, value, latter=False):
        cons = self._TYPES[name].get("constraint")
        if cons is None or value is None:
            return
        if latter is True:
            self.__validate_latter__["constraint"].append(
                (self._validate_constraint, name, value)
            )
            return
        found = False
        for c in cons:
            klass, prop = c.split(".")
            names = self.__constraints__.get(klass, {})
            props = [obj._properties.get(prop) for obj in names.values()]
            if value in props:
                found = True
                break
        if found is not True:
            self._validation_errors.append(
                "{} is not a valid type of {}".format(value, "||".join(cons))
            )
            return

    def _validate_coded(self):
        for item in self.__validate_latter__["unique"]:
            item[0](item[1], item[2])
        for item in self.__validate_latter__["constraint"]:
            item[0](item[1], item[2])
        self._clear_vars()
        if len(self._validation_errors) > 0:
            errors = "\n".join(self._validation_errors)
            self._clear_errors()
            raise Exception(errors)

    def _clear_vars(self):
        if platform.python_version_tuple()[0] == "2":
            self.__validate_latter__["unique"] = []
            self.__validate_latter__["constraint"] = []
        else:
            self.__validate_latter__["unique"].clear()
            self.__validate_latter__["constraint"].clear()

    def _clear_globals(self):
        keys = list(self.__constraints__.keys())
        for k in keys:
            if k == "global":
                self.__constraints__["global"] = []
                continue
            del self.__constraints__[k]


class OpenApiObject(OpenApiBase, OpenApiValidator):
    """Base class for any /components/schemas object

    Every OpenApiObject is reuseable within the schema so it can
    exist in multiple locations within the hierarchy.
    That means it can exist in multiple locations as a
    leaf, parent/choice or parent.
    """

    __slots__ = ("__warnings__", "_properties", "_parent", "_choice")
    _DEFAULTS = {}
    _TYPES = {}
    _REQUIRED = []
    _STATUS = {}

    def __init__(self, parent=None, choice=None):
        super(OpenApiObject, self).__init__()
        self._parent = parent
        self._choice = choice
        self._properties = {}
        self.__warnings__ = []

    @property
    def parent(self):
        return self._parent

    def _set_choice(self, name):
        if self._has_choice(name):
            for enum in self._TYPES["choice"]["enum"]:
                if enum in self._properties and name != enum:
                    self._properties.pop(enum)
            self._properties["choice"] = name

    def _has_choice(self, name):
        if (
            "choice" in dir(self)
            and "_TYPES" in dir(self)
            and "choice" in self._TYPES
            and name in self._TYPES["choice"]["enum"]
        ):
            return True
        else:
            return False

    def _get_property(
        self, name, default_value=None, parent=None, choice=None
    ):
        if name in self._properties and self._properties[name] is not None:
            return self._properties[name]
        if isinstance(default_value, type) is True:
            self._set_choice(name)
            if "_choice" in default_value.__slots__:
                self._properties[name] = default_value(
                    parent=parent, choice=choice
                )
            else:
                self._properties[name] = default_value(parent=parent)
            if (
                "_DEFAULTS" in dir(self._properties[name])
                and "choice" in self._properties[name]._DEFAULTS
            ):
                getattr(
                    self._properties[name],
                    self._properties[name]._DEFAULTS["choice"],
                )
        else:
            if default_value is None and name in self._DEFAULTS:
                self._set_choice(name)
                self._properties[name] = self._DEFAULTS[name]
            else:
                self._properties[name] = default_value
        return self._properties[name]

    def _set_property(self, name, value, choice=None):
        if name == "choice":

            if (
                self.parent is None
                and value is not None
                and value not in self._TYPES["choice"]["enum"]
            ):
                raise Exception(
                    "%s is not a valid choice, valid choices are %s"
                    % (value, ", ".join(self._TYPES["choice"]["enum"]))
                )

            self._set_choice(value)
            if name in self._DEFAULTS and value is None:
                self._properties[name] = self._DEFAULTS[name]
        elif name in self._DEFAULTS and value is None:
            self._set_choice(name)
            self._properties[name] = self._DEFAULTS[name]
        else:
            self._set_choice(name)
            self._properties[name] = value
        # TODO: restore behavior
        # self._validate_unique_and_name(name, value)
        # self._validate_constraint(name, value)
        if (
            self._parent is not None
            and self._choice is not None
            and value is not None
        ):
            self._parent._set_property("choice", self._choice)

    def _encode(self):
        """Helper method for serialization"""
        output = {}
        self._raise_status_warnings(self, None)
        self._validate_required()
        for key, value in self._properties.items():
            self._validate_types(key, value)
            # TODO: restore behavior
            # self._validate_unique_and_name(key, value, True)
            # self._validate_constraint(key, value, True)
            if isinstance(value, (OpenApiObject, OpenApiIter)):
                output[key] = value._encode()
                if isinstance(value, OpenApiObject):
                    self._raise_status_warnings(key, value)
            elif value is not None:
                if (
                    self._TYPES.get(key, {}).get("format", "") == "int64"
                    or self._TYPES.get(key, {}).get("format", "") == "uint64"
                ):
                    value = str(value)
                elif (
                    self._TYPES.get(key, {}).get("itemformat", "") == "int64"
                    or self._TYPES.get(key, {}).get("itemformat", "")
                    == "uint64"
                ):
                    value = [str(v) for v in value]
                output[key] = value
                self._raise_status_warnings(key, value)
        return output

    def _decode(self, obj):
        dtypes = [list, str, int, float, bool]
        self._raise_status_warnings(self, None)
        for property_name, property_value in obj.items():
            if property_name in self._TYPES:
                ignore_warnings = False
                if isinstance(property_value, dict):
                    child = self._get_child_class(property_name)
                    if (
                        "choice" in child[1]._TYPES
                        and "_parent" in child[1].__slots__
                    ):
                        property_value = child[1](self, property_name)._decode(
                            property_value
                        )
                    elif "_parent" in child[1].__slots__:
                        property_value = child[1](self)._decode(property_value)
                    else:
                        property_value = child[1]()._decode(property_value)
                elif (
                    isinstance(property_value, list)
                    and property_name in self._TYPES
                    and self._TYPES[property_name]["type"] not in dtypes
                ):
                    child = self._get_child_class(property_name, True)
                    openapi_list = child[0]()
                    for item in property_value:
                        item = child[1]()._decode(item)
                        openapi_list._items.append(item)
                    property_value = openapi_list
                    ignore_warnings = True
                elif (
                    property_name in self._DEFAULTS and property_value is None
                ):
                    if isinstance(
                        self._DEFAULTS[property_name], tuple(dtypes)
                    ):
                        property_value = self._DEFAULTS[property_name]
                self._set_choice(property_name)
                # convert int64(will be string on wire) to to int
                if (
                    self._TYPES[property_name].get("format", "") == "int64"
                    or self._TYPES[property_name].get("format", "") == "uint64"
                ):
                    property_value = int(property_value)
                elif (
                    self._TYPES[property_name].get("itemformat", "") == "int64"
                    or self._TYPES[property_name].get("itemformat", "")
                    == "uint64"
                ):
                    property_value = [int(v) for v in property_value]
                self._properties[property_name] = property_value
                # TODO: restore behavior
                # OpenApiStatus.warn(
                #     "{}.{}".format(type(self).__name__, property_name), self
                # )
                if not ignore_warnings:
                    self._raise_status_warnings(property_name, property_value)
            self._validate_types(property_name, property_value)
            # TODO: restore behavior
            # self._validate_unique_and_name(property_name, property_value, True)
            # self._validate_constraint(property_name, property_value, True)
        self._validate_required()
        return self

    def _get_child_class(self, property_name, is_property_list=False):
        list_class = None
        class_name = self._TYPES[property_name]["type"]
        module = globals().get(self.__module__)
        if module is None:
            module = importlib.import_module(self.__module__)
            globals()[self.__module__] = module
        object_class = getattr(module, class_name)
        if is_property_list is True:
            list_class = object_class
            object_class = getattr(module, class_name[0:-4])
        return (list_class, object_class)

    def __str__(self):
        return self.serialize(encoding=self.YAML)

    def __deepcopy__(self, memo):
        """Creates a deep copy of the current object"""
        return self.__class__().deserialize(self.serialize())

    def __copy__(self):
        """Creates a deep copy of the current object"""
        return self.__deepcopy__(None)

    def __eq__(self, other):
        return self.__str__() == other.__str__()

    def clone(self):
        """Creates a deep copy of the current object"""
        return self.__deepcopy__(None)

    def _validate_required(self):
        """Validates the required properties are set
        Use getattr as it will set any defaults prior to validating
        """
        if getattr(self, "_REQUIRED", None) is None:
            return
        for name in self._REQUIRED:
            if self._properties.get(name) is None:
                msg = (
                    "{} is a mandatory property of {}"
                    " and should not be set to None".format(
                        name,
                        self.__class__,
                    )
                )
                raise ValueError(msg)

    def _validate_types(self, property_name, property_value):
        common_data_types = [list, str, int, float, bool]
        if property_name not in self._TYPES:
            # raise ValueError("Invalid Property {}".format(property_name))
            return
        details = self._TYPES[property_name]
        if (
            property_value is None
            and property_name not in self._DEFAULTS
            and property_name not in self._REQUIRED
        ):
            return
        if "enum" in details and property_value not in details["enum"]:
            raise_error = False
            if isinstance(property_value, list):
                for value in property_value:
                    if value not in details["enum"]:
                        raise_error = True
                        break
            elif property_value not in details["enum"]:
                raise_error = True

            if raise_error is True:
                msg = (
                    "property {} shall be one of these"
                    " {} enum, but got {} at {}"
                )
                raise TypeError(
                    msg.format(
                        property_name,
                        details["enum"],
                        property_value,
                        self.__class__,
                    )
                )
        if details["type"] in common_data_types and "format" not in details:
            msg = "property {} shall be of type {} at {}".format(
                property_name, details["type"], self.__class__
            )

            itemtype = (
                details.get("itemformat")
                if "itemformat" in details
                else details.get("itemtype")
            )
            self.types_validation(
                property_value,
                details["type"],
                msg,
                itemtype,
                details.get("minimum"),
                details.get("maximum"),
                details.get("minLength"),
                details.get("maxLength"),
            )

        if details["type"] not in common_data_types:
            class_name = details["type"]
            # TODO Need to revisit importlib
            module = importlib.import_module(self.__module__)
            object_class = getattr(module, class_name)
            if not isinstance(property_value, object_class):
                msg = "property {} shall be of type {}," " but got {} at {}"
                raise TypeError(
                    msg.format(
                        property_name,
                        class_name,
                        type(property_value),
                        self.__class__,
                    )
                )
        if "format" in details:
            msg = "Invalid {} format, expected {} at {}".format(
                property_value, details["format"], self.__class__
            )
            _type = (
                details["type"]
                if details["type"] is list
                else details["format"]
            )
            self.types_validation(
                property_value,
                _type,
                msg,
                details["format"],
                details.get("minimum"),
                details.get("maximum"),
                details.get("minLength"),
                details.get("maxLength"),
            )

    def validate(self):
        self._validate_required()
        for key, value in self._properties.items():
            self._validate_types(key, value)
        # TODO: restore behavior
        # self._validate_coded()

    def get(self, name, with_default=False):
        """
        getattr for openapi object
        """
        if self._properties.get(name) is not None:
            return self._properties[name]
        elif with_default:
            # TODO need to find a way to avoid getattr
            choice = (
                self._properties.get("choice")
                if "choice" in dir(self)
                else None
            )
            getattr(self, name)
            if "choice" in dir(self):
                if choice is None and "choice" in self._properties:
                    self._properties.pop("choice")
                else:
                    self._properties["choice"] = choice
            return self._properties.pop(name)
        return None

    def _raise_status_warnings(self, property_name, property_value):
        if len(self._STATUS) > 0:

            if isinstance(property_name, OpenApiObject):
                if "self" in self._STATUS and property_value is None:
                    print("[WARNING]: %s" % self._STATUS["self"])

                return

            enum_key = "%s.%s" % (property_name, property_value)
            if property_name in self._STATUS:
                print("[WARNING]: %s" % self._STATUS[property_name])
            elif enum_key in self._STATUS:
                print("[WARNING]: %s" % self._STATUS[enum_key])


class OpenApiIter(OpenApiBase):
    """Container class for OpenApiObject

    Inheriting classes contain 0..n instances of an OpenAPI components/schemas
    object.
    - config.flows.flow(name="1").flow(name="2").flow(name="3")

    The __getitem__ method allows getting an instance using ordinal.
    - config.flows[0]
    - config.flows[1:]
    - config.flows[0:1]
    - f1, f2, f3 = config.flows

    The __iter__ method allows for iterating across the encapsulated contents
    - for flow in config.flows:
    """

    __slots__ = ("_index", "_items")
    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self):
        super(OpenApiIter, self).__init__()
        self._index = -1
        self._items = []

    def __len__(self):
        return len(self._items)

    def _getitem(self, key):
        found = None
        if isinstance(key, int):
            found = self._items[key]
        elif isinstance(key, slice) is True:
            start, stop, step = key.indices(len(self))
            sliced = self.__class__()
            for i in range(start, stop, step):
                sliced._items.append(self._items[i])
            return sliced
        elif isinstance(key, str):
            for item in self._items:
                if item.name == key:
                    found = item
        if found is None:
            raise IndexError()
        if (
            self._GETITEM_RETURNS_CHOICE_OBJECT is True
            and found._properties.get("choice") is not None
            and found._properties.get(found._properties["choice"]) is not None
        ):
            return found._properties[found._properties["choice"]]
        return found

    def _iter(self):
        self._index = -1
        return self

    def _next(self):
        if self._index + 1 >= len(self._items):
            raise StopIteration
        else:
            self._index += 1
        return self.__getitem__(self._index)

    def __getitem__(self, key):
        raise NotImplementedError("This should be overridden by the generator")

    def _add(self, item):
        self._items.append(item)
        self._index = len(self._items) - 1

    def remove(self, index):
        del self._items[index]
        self._index = len(self._items) - 1

    def append(self, item):
        """Append an item to the end of OpenApiIter
        TBD: type check, raise error on mismatch
        """
        self._instanceOf(item)
        self._add(item)
        return self

    def clear(self):
        del self._items[:]
        self._index = -1

    def set(self, index, item):
        self._instanceOf(item)
        self._items[index] = item
        return self

    def _encode(self):
        return [item._encode() for item in self._items]

    def _decode(self, encoded_list):
        item_class_name = self.__class__.__name__.replace("Iter", "")
        module = importlib.import_module(self.__module__)
        object_class = getattr(module, item_class_name)
        self.clear()
        for item in encoded_list:
            self._add(object_class()._decode(item))

    def __copy__(self):
        raise NotImplementedError(
            "Shallow copy of OpenApiIter objects is not supported"
        )

    def __deepcopy__(self, memo):
        raise NotImplementedError(
            "Deep copy of OpenApiIter objects is not supported"
        )

    def __str__(self):
        return yaml.safe_dump(self._encode())

    def __eq__(self, other):
        return self.__str__() == other.__str__()

    def _instanceOf(self, item):
        raise NotImplementedError(
            "validating an OpenApiIter object is not supported"
        )


class PrefixConfig(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "required_object": {"type": "EObject"},
        "optional_object": {"type": "EObject"},
        "ieee_802_1qbb": {"type": bool},
        "space_1": {
            "type": int,
            "format": "int32",
        },
        "full_duplex_100_mb": {
            "type": int,
            "format": "int64",
            "minimum": -10,
            "maximum": 4261412864,
        },
        "response": {
            "type": str,
            "enum": [
                "status_200",
                "status_400",
                "status_404",
                "status_500",
            ],
        },
        "a": {"type": str},
        "b": {
            "type": float,
            "format": "float",
        },
        "c": {
            "type": int,
            "format": "int32",
        },
        "d_values": {
            "type": list,
            "enum": [
                "a",
                "b",
                "c",
            ],
            "itemtype": str,
        },
        "e": {"type": "EObject"},
        "f": {"type": "FObject"},
        "g": {"type": "GObjectIter"},
        "h": {"type": bool},
        "i": {
            "type": str,
            "format": "binary",
        },
        "j": {"type": "JObjectIter"},
        "k": {"type": "KObject"},
        "l": {"type": "LObject"},
        "list_of_string_values": {
            "type": list,
            "itemtype": str,
        },
        "list_of_integer_values": {
            "type": list,
            "itemtype": int,
            "itemformat": "int32",
        },
        "level": {"type": "LevelOne"},
        "mandatory": {"type": "Mandate"},
        "ipv4_pattern": {"type": "Ipv4Pattern"},
        "ipv6_pattern": {"type": "Ipv6Pattern"},
        "mac_pattern": {"type": "MacPattern"},
        "integer_pattern": {"type": "IntegerPattern"},
        "checksum_pattern": {"type": "ChecksumPattern"},
        "case": {"type": "Layer1Ieee802x"},
        "m_object": {"type": "MObject"},
        "integer64": {
            "type": int,
            "format": "int64",
        },
        "integer64_list": {
            "type": list,
            "itemtype": int,
            "itemformat": "int64",
            "minimum": -12,
            "maximum": 4261412864,
        },
        "header_checksum": {"type": "PatternPrefixConfigHeaderChecksum"},
        "str_len": {
            "type": str,
            "minLength": 3,
            "maxLength": 6,
        },
        "hex_slice": {
            "type": list,
            "itemtype": str,
            "itemformat": "hex",
        },
        "auto_field_test": {"type": "PatternPrefixConfigAutoFieldTest"},
        "name": {"type": str},
        "w_list": {"type": "WObjectIter"},
        "x_list": {"type": "ZObjectIter"},
        "z_object": {"type": "ZObject"},
        "y_object": {"type": "YObject"},
        "choice_object": {"type": "ChoiceObjectIter"},
        "required_choice_object": {"type": "RequiredChoiceParent"},
        "g1": {"type": "GObjectIter"},
        "g2": {"type": "GObjectIter"},
        "int32_param": {
            "type": int,
            "format": "int32",
        },
        "int32_list_param": {
            "type": list,
            "itemtype": int,
            "itemformat": "int32",
            "minimum": -23456,
            "maximum": 23456,
        },
        "uint32_param": {
            "type": int,
            "format": "uint32",
        },
        "uint32_list_param": {
            "type": list,
            "itemtype": int,
            "itemformat": "uint32",
            "minimum": 0,
            "maximum": 4294967293,
        },
        "uint64_param": {
            "type": int,
            "format": "uint64",
        },
        "uint64_list_param": {
            "type": list,
            "itemtype": int,
            "itemformat": "uint64",
        },
        "auto_int32_param": {
            "type": int,
            "format": "int32",
            "minimum": 64,
            "maximum": 9000,
        },
        "auto_int32_list_param": {
            "type": list,
            "itemtype": int,
            "itemformat": "int32",
            "minimum": 64,
            "maximum": 9000,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ("a", "b", "c", "required_object")  # type: tuple(str)

    _DEFAULTS = {
        "response": "status_200",
        "h": True,
    }  # type: Dict[str, Union(type)]

    STATUS_200 = "status_200"  # type: str
    STATUS_400 = "status_400"  # type: str
    STATUS_404 = "status_404"  # type: str
    STATUS_500 = "status_500"  # type: str

    A = "a"  # type: str
    B = "b"  # type: str
    C = "c"  # type: str

    _STATUS = {
        "space_1": "space_1 property in schema PrefixConfig is deprecated, Information TBD",
        "response.status_404": "STATUS_404 enum in property response is deprecated, new code will be coming soon",
        "response.status_500": "STATUS_500 enum in property response is under_review, 500 can change to other values",
        "a": "a property in schema PrefixConfig is under_review, Information TBD",
        "d_values": "d_values property in schema PrefixConfig is deprecated, Information TBD",
        "e": "e property in schema PrefixConfig is deprecated, Information TBD",
        "str_len": "str_len property in schema PrefixConfig is under_review, Information TBD",
        "hex_slice": "hex_slice property in schema PrefixConfig is under_review, Information TBD",
    }  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        ieee_802_1qbb=None,
        space_1=None,
        full_duplex_100_mb=None,
        response="status_200",
        a=None,
        b=None,
        c=None,
        d_values=None,
        h=True,
        i=None,
        list_of_string_values=None,
        list_of_integer_values=None,
        integer64=None,
        integer64_list=None,
        str_len=None,
        hex_slice=None,
        name=None,
        int32_param=None,
        int32_list_param=None,
        uint32_param=None,
        uint32_list_param=None,
        uint64_param=None,
        uint64_list_param=None,
        auto_int32_param=None,
        auto_int32_list_param=None,
    ):
        super(PrefixConfig, self).__init__()
        self._parent = parent
        self._set_property("ieee_802_1qbb", ieee_802_1qbb)
        self._set_property("space_1", space_1)
        self._set_property("full_duplex_100_mb", full_duplex_100_mb)
        self._set_property("response", response)
        self._set_property("a", a)
        self._set_property("b", b)
        self._set_property("c", c)
        self._set_property("d_values", d_values)
        self._set_property("h", h)
        self._set_property("i", i)
        self._set_property("list_of_string_values", list_of_string_values)
        self._set_property("list_of_integer_values", list_of_integer_values)
        self._set_property("integer64", integer64)
        self._set_property("integer64_list", integer64_list)
        self._set_property("str_len", str_len)
        self._set_property("hex_slice", hex_slice)
        self._set_property("name", name)
        self._set_property("int32_param", int32_param)
        self._set_property("int32_list_param", int32_list_param)
        self._set_property("uint32_param", uint32_param)
        self._set_property("uint32_list_param", uint32_list_param)
        self._set_property("uint64_param", uint64_param)
        self._set_property("uint64_list_param", uint64_list_param)
        self._set_property("auto_int32_param", auto_int32_param)
        self._set_property("auto_int32_list_param", auto_int32_list_param)

    def set(
        self,
        ieee_802_1qbb=None,
        space_1=None,
        full_duplex_100_mb=None,
        response=None,
        a=None,
        b=None,
        c=None,
        d_values=None,
        h=None,
        i=None,
        list_of_string_values=None,
        list_of_integer_values=None,
        integer64=None,
        integer64_list=None,
        str_len=None,
        hex_slice=None,
        name=None,
        int32_param=None,
        int32_list_param=None,
        uint32_param=None,
        uint32_list_param=None,
        uint64_param=None,
        uint64_list_param=None,
        auto_int32_param=None,
        auto_int32_list_param=None,
    ):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def required_object(self):
        # type: () -> EObject
        """required_object getter

        A required object that MUST be generated as such.

        Returns: EObject
        """
        return self._get_property("required_object", EObject)

    @property
    def optional_object(self):
        # type: () -> EObject
        """optional_object getter

        An optional object that MUST be generated as such.

        Returns: EObject
        """
        return self._get_property("optional_object", EObject)

    @property
    def ieee_802_1qbb(self):
        # type: () -> bool
        """ieee_802_1qbb getter

        TBD

        Returns: bool
        """
        return self._get_property("ieee_802_1qbb")

    @ieee_802_1qbb.setter
    def ieee_802_1qbb(self, value):
        """ieee_802_1qbb setter

        TBD

        value: bool
        """
        self._set_property("ieee_802_1qbb", value)

    @property
    def space_1(self):
        # type: () -> int
        """space_1 getter

        Deprecated: Information TBD. Description TBD

        Returns: int
        """
        return self._get_property("space_1")

    @space_1.setter
    def space_1(self, value):
        """space_1 setter

        Deprecated: Information TBD. Description TBD

        value: int
        """
        self._set_property("space_1", value)

    @property
    def full_duplex_100_mb(self):
        # type: () -> int
        """full_duplex_100_mb getter

        TBD

        Returns: int
        """
        return self._get_property("full_duplex_100_mb")

    @full_duplex_100_mb.setter
    def full_duplex_100_mb(self, value):
        """full_duplex_100_mb setter

        TBD

        value: int
        """
        self._set_property("full_duplex_100_mb", value)

    @property
    def response(self):
        # type: () -> Union[Literal["status_200"], Literal["status_400"], Literal["status_404"], Literal["status_500"]]
        """response getter

        Indicate to the server what response should be returned

        Returns: Union[Literal["status_200"], Literal["status_400"], Literal["status_404"], Literal["status_500"]]
        """
        return self._get_property("response")

    @response.setter
    def response(self, value):
        """response setter

        Indicate to the server what response should be returned

        value: Union[Literal["status_200"], Literal["status_400"], Literal["status_404"], Literal["status_500"]]
        """
        self._set_property("response", value)

    @property
    def a(self):
        # type: () -> str
        """a getter

        Under Review: Information TBD. Small single line description

        Returns: str
        """
        return self._get_property("a")

    @a.setter
    def a(self, value):
        """a setter

        Under Review: Information TBD. Small single line description

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property a as None")
        self._set_property("a", value)

    @property
    def b(self):
        # type: () -> float
        """b getter

        Longer multi-line description. Second line is here. Third line

        Returns: float
        """
        return self._get_property("b")

    @b.setter
    def b(self, value):
        """b setter

        Longer multi-line description. Second line is here. Third line

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property b as None")
        self._set_property("b", value)

    @property
    def c(self):
        # type: () -> int
        """c getter

        TBD

        Returns: int
        """
        return self._get_property("c")

    @c.setter
    def c(self, value):
        """c setter

        TBD

        value: int
        """
        if value is None:
            raise TypeError("Cannot set required property c as None")
        self._set_property("c", value)

    @property
    def d_values(self):
        # type: () -> List[Union[Literal["a"], Literal["b"], Literal["c"]]]
        """d_values getter

        Deprecated: Information TBD. A list of enum values

        Returns: List[Union[Literal["a"], Literal["b"], Literal["c"]]]
        """
        return self._get_property("d_values")

    @d_values.setter
    def d_values(self, value):
        """d_values setter

        Deprecated: Information TBD. A list of enum values

        value: List[Union[Literal["a"], Literal["b"], Literal["c"]]]
        """
        self._set_property("d_values", value)

    @property
    def e(self):
        # type: () -> EObject
        """e getter

        Deprecated: Information TBD. A child object

        Returns: EObject
        """
        return self._get_property("e", EObject)

    @property
    def f(self):
        # type: () -> FObject
        """f getter

        An object with only choice(s)

        Returns: FObject
        """
        return self._get_property("f", FObject)

    @property
    def g(self):
        # type: () -> GObjectIter
        """g getter

        A list of objects with choice and properties

        Returns: GObjectIter
        """
        return self._get_property("g", GObjectIter, self._parent, self._choice)

    @property
    def h(self):
        # type: () -> bool
        """h getter

        A boolean value

        Returns: bool
        """
        return self._get_property("h")

    @h.setter
    def h(self, value):
        """h setter

        A boolean value

        value: bool
        """
        self._set_property("h", value)

    @property
    def i(self):
        # type: () -> str
        """i getter

        A byte string

        Returns: str
        """
        return self._get_property("i")

    @i.setter
    def i(self, value):
        """i setter

        A byte string

        value: str
        """
        self._set_property("i", value)

    @property
    def j(self):
        # type: () -> JObjectIter
        """j getter

        A list of objects with only choice

        Returns: JObjectIter
        """
        return self._get_property("j", JObjectIter, self._parent, self._choice)

    @property
    def k(self):
        # type: () -> KObject
        """k getter

        A nested object with only one property which is choice object

        Returns: KObject
        """
        return self._get_property("k", KObject)

    @property
    def l(self):
        # type: () -> LObject
        """l getter

        Format validation objectFormat validation objectFormat validation object

        Returns: LObject
        """
        return self._get_property("l", LObject)

    @property
    def list_of_string_values(self):
        # type: () -> List[str]
        """list_of_string_values getter

        A list of string values

        Returns: List[str]
        """
        return self._get_property("list_of_string_values")

    @list_of_string_values.setter
    def list_of_string_values(self, value):
        """list_of_string_values setter

        A list of string values

        value: List[str]
        """
        self._set_property("list_of_string_values", value)

    @property
    def list_of_integer_values(self):
        # type: () -> List[int]
        """list_of_integer_values getter

        A list of integer values

        Returns: List[int]
        """
        return self._get_property("list_of_integer_values")

    @list_of_integer_values.setter
    def list_of_integer_values(self, value):
        """list_of_integer_values setter

        A list of integer values

        value: List[int]
        """
        self._set_property("list_of_integer_values", value)

    @property
    def level(self):
        # type: () -> LevelOne
        """level getter

        To Test Multi level non-primitive typesTo Test Multi level non-primitive typesTo Test Multi level non-primitive types

        Returns: LevelOne
        """
        return self._get_property("level", LevelOne)

    @property
    def mandatory(self):
        # type: () -> Mandate
        """mandatory getter

        Object to Test required ParameterObject to Test required ParameterObject to Test required Parameter

        Returns: Mandate
        """
        return self._get_property("mandatory", Mandate)

    @property
    def ipv4_pattern(self):
        # type: () -> Ipv4Pattern
        """ipv4_pattern getter

        Test ipv4 patternTest ipv4 patternTest ipv4 pattern

        Returns: Ipv4Pattern
        """
        return self._get_property("ipv4_pattern", Ipv4Pattern)

    @property
    def ipv6_pattern(self):
        # type: () -> Ipv6Pattern
        """ipv6_pattern getter

        Test ipv6 patternTest ipv6 patternTest ipv6 pattern

        Returns: Ipv6Pattern
        """
        return self._get_property("ipv6_pattern", Ipv6Pattern)

    @property
    def mac_pattern(self):
        # type: () -> MacPattern
        """mac_pattern getter

        Test mac patternTest mac patternTest mac pattern

        Returns: MacPattern
        """
        return self._get_property("mac_pattern", MacPattern)

    @property
    def integer_pattern(self):
        # type: () -> IntegerPattern
        """integer_pattern getter

        Test integer patternTest integer patternTest integer pattern

        Returns: IntegerPattern
        """
        return self._get_property("integer_pattern", IntegerPattern)

    @property
    def checksum_pattern(self):
        # type: () -> ChecksumPattern
        """checksum_pattern getter

        Test checksum patternTest checksum patternTest checksum pattern

        Returns: ChecksumPattern
        """
        return self._get_property("checksum_pattern", ChecksumPattern)

    @property
    def case(self):
        # type: () -> Layer1Ieee802x
        """case getter



        Returns: Layer1Ieee802x
        """
        return self._get_property("case", Layer1Ieee802x)

    @property
    def m_object(self):
        # type: () -> MObject
        """m_object getter

        Required format validation objectRequired format validation objectRequired format validation object

        Returns: MObject
        """
        return self._get_property("m_object", MObject)

    @property
    def integer64(self):
        # type: () -> int
        """integer64 getter

        int64 type

        Returns: int
        """
        return self._get_property("integer64")

    @integer64.setter
    def integer64(self, value):
        """integer64 setter

        int64 type

        value: int
        """
        self._set_property("integer64", value)

    @property
    def integer64_list(self):
        # type: () -> List[int]
        """integer64_list getter

        int64 type list

        Returns: List[int]
        """
        return self._get_property("integer64_list")

    @integer64_list.setter
    def integer64_list(self, value):
        """integer64_list setter

        int64 type list

        value: List[int]
        """
        self._set_property("integer64_list", value)

    @property
    def header_checksum(self):
        # type: () -> PatternPrefixConfigHeaderChecksum
        """header_checksum getter

        Header checksumHeader checksumHeader checksum

        Returns: PatternPrefixConfigHeaderChecksum
        """
        return self._get_property(
            "header_checksum", PatternPrefixConfigHeaderChecksum
        )

    @property
    def str_len(self):
        # type: () -> str
        """str_len getter

        Under Review: Information TBD. string minimum&maximum Length

        Returns: str
        """
        return self._get_property("str_len")

    @str_len.setter
    def str_len(self, value):
        """str_len setter

        Under Review: Information TBD. string minimum&maximum Length

        value: str
        """
        self._set_property("str_len", value)

    @property
    def hex_slice(self):
        # type: () -> List[str]
        """hex_slice getter

        Under Review: Information TBD. Array of Hex

        Returns: List[str]
        """
        return self._get_property("hex_slice")

    @hex_slice.setter
    def hex_slice(self, value):
        """hex_slice setter

        Under Review: Information TBD. Array of Hex

        value: List[str]
        """
        self._set_property("hex_slice", value)

    @property
    def auto_field_test(self):
        # type: () -> PatternPrefixConfigAutoFieldTest
        """auto_field_test getter

        TBDTBDTBD

        Returns: PatternPrefixConfigAutoFieldTest
        """
        return self._get_property(
            "auto_field_test", PatternPrefixConfigAutoFieldTest
        )

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        self._set_property("name", value)

    @property
    def w_list(self):
        # type: () -> WObjectIter
        """w_list getter

        TBD

        Returns: WObjectIter
        """
        return self._get_property(
            "w_list", WObjectIter, self._parent, self._choice
        )

    @property
    def x_list(self):
        # type: () -> ZObjectIter
        """x_list getter

        TBD

        Returns: ZObjectIter
        """
        return self._get_property(
            "x_list", ZObjectIter, self._parent, self._choice
        )

    @property
    def z_object(self):
        # type: () -> ZObject
        """z_object getter



        Returns: ZObject
        """
        return self._get_property("z_object", ZObject)

    @property
    def y_object(self):
        # type: () -> YObject
        """y_object getter



        Returns: YObject
        """
        return self._get_property("y_object", YObject)

    @property
    def choice_object(self):
        # type: () -> ChoiceObjectIter
        """choice_object getter

        A list of objects with choice with and without properties

        Returns: ChoiceObjectIter
        """
        return self._get_property(
            "choice_object", ChoiceObjectIter, self._parent, self._choice
        )

    @property
    def required_choice_object(self):
        # type: () -> RequiredChoiceParent
        """required_choice_object getter



        Returns: RequiredChoiceParent
        """
        return self._get_property(
            "required_choice_object", RequiredChoiceParent
        )

    @property
    def g1(self):
        # type: () -> GObjectIter
        """g1 getter

        A list of objects with choice and properties

        Returns: GObjectIter
        """
        return self._get_property(
            "g1", GObjectIter, self._parent, self._choice
        )

    @property
    def g2(self):
        # type: () -> GObjectIter
        """g2 getter

        A list of objects with choice and properties

        Returns: GObjectIter
        """
        return self._get_property(
            "g2", GObjectIter, self._parent, self._choice
        )

    @property
    def int32_param(self):
        # type: () -> int
        """int32_param getter

        int32 type

        Returns: int
        """
        return self._get_property("int32_param")

    @int32_param.setter
    def int32_param(self, value):
        """int32_param setter

        int32 type

        value: int
        """
        self._set_property("int32_param", value)

    @property
    def int32_list_param(self):
        # type: () -> List[int]
        """int32_list_param getter

        int32 type list

        Returns: List[int]
        """
        return self._get_property("int32_list_param")

    @int32_list_param.setter
    def int32_list_param(self, value):
        """int32_list_param setter

        int32 type list

        value: List[int]
        """
        self._set_property("int32_list_param", value)

    @property
    def uint32_param(self):
        # type: () -> int
        """uint32_param getter

        uint32 type

        Returns: int
        """
        return self._get_property("uint32_param")

    @uint32_param.setter
    def uint32_param(self, value):
        """uint32_param setter

        uint32 type

        value: int
        """
        self._set_property("uint32_param", value)

    @property
    def uint32_list_param(self):
        # type: () -> List[int]
        """uint32_list_param getter

        uint32 type list

        Returns: List[int]
        """
        return self._get_property("uint32_list_param")

    @uint32_list_param.setter
    def uint32_list_param(self, value):
        """uint32_list_param setter

        uint32 type list

        value: List[int]
        """
        self._set_property("uint32_list_param", value)

    @property
    def uint64_param(self):
        # type: () -> int
        """uint64_param getter

        uint64 type

        Returns: int
        """
        return self._get_property("uint64_param")

    @uint64_param.setter
    def uint64_param(self, value):
        """uint64_param setter

        uint64 type

        value: int
        """
        self._set_property("uint64_param", value)

    @property
    def uint64_list_param(self):
        # type: () -> List[int]
        """uint64_list_param getter

        uint64 type list

        Returns: List[int]
        """
        return self._get_property("uint64_list_param")

    @uint64_list_param.setter
    def uint64_list_param(self, value):
        """uint64_list_param setter

        uint64 type list

        value: List[int]
        """
        self._set_property("uint64_list_param", value)

    @property
    def auto_int32_param(self):
        # type: () -> int
        """auto_int32_param getter

        should automatically set type to int32

        Returns: int
        """
        return self._get_property("auto_int32_param")

    @auto_int32_param.setter
    def auto_int32_param(self, value):
        """auto_int32_param setter

        should automatically set type to int32

        value: int
        """
        self._set_property("auto_int32_param", value)

    @property
    def auto_int32_list_param(self):
        # type: () -> List[int]
        """auto_int32_list_param getter

        should automatically set type to []int32

        Returns: List[int]
        """
        return self._get_property("auto_int32_list_param")

    @auto_int32_list_param.setter
    def auto_int32_list_param(self, value):
        """auto_int32_list_param setter

        should automatically set type to []int32

        value: List[int]
        """
        self._set_property("auto_int32_list_param", value)


class EObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "e_a": {
            "type": float,
            "format": "float",
        },
        "e_b": {
            "type": float,
            "format": "double",
        },
        "name": {"type": str},
        "m_param1": {"type": str},
        "m_param2": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ("e_a", "e_b")  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        e_a=None,
        e_b=None,
        name=None,
        m_param1=None,
        m_param2=None,
    ):
        super(EObject, self).__init__()
        self._parent = parent
        self._set_property("e_a", e_a)
        self._set_property("e_b", e_b)
        self._set_property("name", name)
        self._set_property("m_param1", m_param1)
        self._set_property("m_param2", m_param2)

    def set(self, e_a=None, e_b=None, name=None, m_param1=None, m_param2=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def e_a(self):
        # type: () -> float
        """e_a getter

        TBD

        Returns: float
        """
        return self._get_property("e_a")

    @e_a.setter
    def e_a(self, value):
        """e_a setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property e_a as None")
        self._set_property("e_a", value)

    @property
    def e_b(self):
        # type: () -> float
        """e_b getter

        TBD

        Returns: float
        """
        return self._get_property("e_b")

    @e_b.setter
    def e_b(self, value):
        """e_b setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property e_b as None")
        self._set_property("e_b", value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        self._set_property("name", value)

    @property
    def m_param1(self):
        # type: () -> str
        """m_param1 getter

        TBD

        Returns: str
        """
        return self._get_property("m_param1")

    @m_param1.setter
    def m_param1(self, value):
        """m_param1 setter

        TBD

        value: str
        """
        self._set_property("m_param1", value)

    @property
    def m_param2(self):
        # type: () -> str
        """m_param2 getter

        TBD

        Returns: str
        """
        return self._get_property("m_param2")

    @m_param2.setter
    def m_param2(self, value):
        """m_param2 setter

        TBD

        value: str
        """
        self._set_property("m_param2", value)


class FObject(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "f_a",
                "f_b",
                "f_c",
            ],
        },
        "f_a": {"type": str},
        "f_b": {
            "type": float,
            "format": "double",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "f_a",
        "f_a": "some string",
        "f_b": 3.0,
    }  # type: Dict[str, Union(type)]

    F_A = "f_a"  # type: str
    F_B = "f_b"  # type: str
    F_C = "f_c"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, f_a="some string", f_b=3.0):
        super(FObject, self).__init__()
        self._parent = parent
        self._set_property("f_a", f_a)
        self._set_property("f_b", f_b)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, f_a=None, f_b=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def choice(self):
        # type: () -> Union[Literal["f_a"], Literal["f_b"], Literal["f_c"]]
        """choice getter

        TBD

        Returns: Union[Literal["f_a"], Literal["f_b"], Literal["f_c"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["f_a"], Literal["f_b"], Literal["f_c"]]
        """
        self._set_property("choice", value)

    @property
    def f_a(self):
        # type: () -> str
        """f_a getter

        TBD

        Returns: str
        """
        return self._get_property("f_a")

    @f_a.setter
    def f_a(self, value):
        """f_a setter

        TBD

        value: str
        """
        self._set_property("f_a", value, "f_a")

    @property
    def f_b(self):
        # type: () -> float
        """f_b getter

        TBD

        Returns: float
        """
        return self._get_property("f_b")

    @f_b.setter
    def f_b(self, value):
        """f_b setter

        TBD

        value: float
        """
        self._set_property("f_b", value, "f_b")


class GObject(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "g_a": {"type": str},
        "g_b": {
            "type": int,
            "format": "int32",
        },
        "g_c": {"type": float},
        "choice": {
            "type": str,
            "enum": [
                "g_d",
                "g_e",
            ],
        },
        "g_d": {"type": str},
        "g_e": {
            "type": float,
            "format": "double",
        },
        "g_f": {
            "type": str,
            "enum": [
                "a",
                "b",
                "c",
            ],
        },
        "name": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "g_a": "asdf",
        "g_b": 6,
        "g_c": 77.7,
        "choice": "g_d",
        "g_d": "some string",
        "g_e": 3.0,
        "g_f": "a",
    }  # type: Dict[str, Union(type)]

    G_D = "g_d"  # type: str
    G_E = "g_e"  # type: str

    A = "a"  # type: str
    B = "b"  # type: str
    C = "c"  # type: str

    _STATUS = {
        "self": "GObject is deprecated, new schema Jobject to be used",
        "g_c": "g_c property in schema GObject is deprecated, Information TBD",
    }  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        choice=None,
        g_a="asdf",
        g_b=6,
        g_c=77.7,
        g_d="some string",
        g_e=3.0,
        g_f="a",
        name=None,
    ):
        super(GObject, self).__init__()
        self._parent = parent
        self._set_property("g_a", g_a)
        self._set_property("g_b", g_b)
        self._set_property("g_c", g_c)
        self._set_property("g_d", g_d)
        self._set_property("g_e", g_e)
        self._set_property("g_f", g_f)
        self._set_property("name", name)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(
        self,
        g_a=None,
        g_b=None,
        g_c=None,
        g_d=None,
        g_e=None,
        g_f=None,
        name=None,
    ):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def g_a(self):
        # type: () -> str
        """g_a getter

        TBD

        Returns: str
        """
        return self._get_property("g_a")

    @g_a.setter
    def g_a(self, value):
        """g_a setter

        TBD

        value: str
        """
        self._set_property("g_a", value)

    @property
    def g_b(self):
        # type: () -> int
        """g_b getter

        TBD

        Returns: int
        """
        return self._get_property("g_b")

    @g_b.setter
    def g_b(self, value):
        """g_b setter

        TBD

        value: int
        """
        self._set_property("g_b", value)

    @property
    def g_c(self):
        # type: () -> float
        """g_c getter

        Deprecated: Information TBD. Description TBD

        Returns: float
        """
        return self._get_property("g_c")

    @g_c.setter
    def g_c(self, value):
        """g_c setter

        Deprecated: Information TBD. Description TBD

        value: float
        """
        self._set_property("g_c", value)

    @property
    def choice(self):
        # type: () -> Union[Literal["g_d"], Literal["g_e"]]
        """choice getter

        TBD

        Returns: Union[Literal["g_d"], Literal["g_e"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["g_d"], Literal["g_e"]]
        """
        self._set_property("choice", value)

    @property
    def g_d(self):
        # type: () -> str
        """g_d getter

        TBD

        Returns: str
        """
        return self._get_property("g_d")

    @g_d.setter
    def g_d(self, value):
        """g_d setter

        TBD

        value: str
        """
        self._set_property("g_d", value, "g_d")

    @property
    def g_e(self):
        # type: () -> float
        """g_e getter

        TBD

        Returns: float
        """
        return self._get_property("g_e")

    @g_e.setter
    def g_e(self, value):
        """g_e setter

        TBD

        value: float
        """
        self._set_property("g_e", value, "g_e")

    @property
    def g_f(self):
        # type: () -> Union[Literal["a"], Literal["b"], Literal["c"]]
        """g_f getter

        Another enum to test protbuf enum generation

        Returns: Union[Literal["a"], Literal["b"], Literal["c"]]
        """
        return self._get_property("g_f")

    @g_f.setter
    def g_f(self, value):
        """g_f setter

        Another enum to test protbuf enum generation

        value: Union[Literal["a"], Literal["b"], Literal["c"]]
        """
        self._set_property("g_f", value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        self._set_property("name", value)


class GObjectIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(GObjectIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[GObject]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> GObjectIter
        return self._iter()

    def __next__(self):
        # type: () -> GObject
        return self._next()

    def next(self):
        # type: () -> GObject
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, GObject):
            raise Exception("Item is not an instance of GObject")

    def gobject(
        self,
        g_a="asdf",
        g_b=6,
        g_c=77.7,
        g_d="some string",
        g_e=3.0,
        g_f="a",
        name=None,
    ):
        # type: (str,int,float,str,float,Union[Literal["a"], Literal["b"], Literal["c"]],str) -> GObjectIter
        """Factory method that creates an instance of the GObject class

        Deprecated: new schema Jobject to be used. Description TBD

        Returns: GObjectIter
        """
        item = GObject(
            parent=self._parent,
            choice=self._choice,
            g_a=g_a,
            g_b=g_b,
            g_c=g_c,
            g_d=g_d,
            g_e=g_e,
            g_f=g_f,
            name=name,
        )
        self._add(item)
        return self

    def add(
        self,
        g_a="asdf",
        g_b=6,
        g_c=77.7,
        g_d="some string",
        g_e=3.0,
        g_f="a",
        name=None,
    ):
        # type: (str,int,float,str,float,Union[Literal["a"], Literal["b"], Literal["c"]],str) -> GObject
        """Add method that creates and returns an instance of the GObject class

        Deprecated: new schema Jobject to be used. Description TBD

        Returns: GObject
        """
        item = GObject(
            parent=self._parent,
            choice=self._choice,
            g_a=g_a,
            g_b=g_b,
            g_c=g_c,
            g_d=g_d,
            g_e=g_e,
            g_f=g_f,
            name=name,
        )
        self._add(item)
        return item


class JObject(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "j_a",
                "j_b",
            ],
        },
        "j_a": {"type": "EObject"},
        "j_b": {"type": "FObject"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "j_a",
    }  # type: Dict[str, Union(type)]

    J_A = "j_a"  # type: str
    J_B = "j_b"  # type: str

    _STATUS = {
        "choice.j_b": "J_B enum in property choice is deprecated, use j_a instead",
    }  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None):
        super(JObject, self).__init__()
        self._parent = parent
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    @property
    def j_a(self):
        # type: () -> EObject
        """Factory property that returns an instance of the EObject class

        TBD

        Returns: EObject
        """
        return self._get_property("j_a", EObject, self, "j_a")

    @property
    def j_b(self):
        # type: () -> FObject
        """Factory property that returns an instance of the FObject class

        TBD

        Returns: FObject
        """
        return self._get_property("j_b", FObject, self, "j_b")

    @property
    def choice(self):
        # type: () -> Union[Literal["j_a"], Literal["j_b"]]
        """choice getter

        TBD

        Returns: Union[Literal["j_a"], Literal["j_b"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["j_a"], Literal["j_b"]]
        """
        self._set_property("choice", value)


class JObjectIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = True

    def __init__(self, parent=None, choice=None):
        super(JObjectIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[EObject, FObject, JObject]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> JObjectIter
        return self._iter()

    def __next__(self):
        # type: () -> JObject
        return self._next()

    def next(self):
        # type: () -> JObject
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, JObject):
            raise Exception("Item is not an instance of JObject")

    def jobject(self):
        # type: () -> JObjectIter
        """Factory method that creates an instance of the JObject class

        TBD

        Returns: JObjectIter
        """
        item = JObject(parent=self._parent, choice=self._choice)
        self._add(item)
        return self

    def add(self):
        # type: () -> JObject
        """Add method that creates and returns an instance of the JObject class

        TBD

        Returns: JObject
        """
        item = JObject(parent=self._parent, choice=self._choice)
        self._add(item)
        return item

    def j_a(self, e_a=None, e_b=None, name=None, m_param1=None, m_param2=None):
        # type: (float,float,str,str,str) -> JObjectIter
        """Factory method that creates an instance of the EObject class

        TBD

        Returns: JObjectIter
        """
        item = JObject()
        item.j_a
        item.choice = "j_a"
        self._add(item)
        return self

    def j_b(self, f_a="some string", f_b=3.0):
        # type: (str,float) -> JObjectIter
        """Factory method that creates an instance of the FObject class

        TBD

        Returns: JObjectIter
        """
        item = JObject()
        item.j_b
        item.choice = "j_b"
        self._add(item)
        return self


class KObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "e_object": {"type": "EObject"},
        "f_object": {"type": "FObject"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(KObject, self).__init__()
        self._parent = parent

    @property
    def e_object(self):
        # type: () -> EObject
        """e_object getter

        TBDTBDTBD

        Returns: EObject
        """
        return self._get_property("e_object", EObject)

    @property
    def f_object(self):
        # type: () -> FObject
        """f_object getter

        TBDTBDTBD

        Returns: FObject
        """
        return self._get_property("f_object", FObject)


class LObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "string_param": {"type": str},
        "integer": {
            "type": int,
            "format": "int32",
            "minimum": -10,
            "maximum": 90,
        },
        "float": {
            "type": float,
            "format": "float",
        },
        "double": {
            "type": float,
            "format": "double",
        },
        "mac": {
            "type": str,
            "format": "mac",
        },
        "ipv4": {
            "type": str,
            "format": "ipv4",
        },
        "ipv6": {
            "type": str,
            "format": "ipv6",
        },
        "hex": {
            "type": str,
            "format": "hex",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        string_param=None,
        integer=None,
        float=None,
        double=None,
        mac=None,
        ipv4=None,
        ipv6=None,
        hex=None,
    ):
        super(LObject, self).__init__()
        self._parent = parent
        self._set_property("string_param", string_param)
        self._set_property("integer", integer)
        self._set_property("float", float)
        self._set_property("double", double)
        self._set_property("mac", mac)
        self._set_property("ipv4", ipv4)
        self._set_property("ipv6", ipv6)
        self._set_property("hex", hex)

    def set(
        self,
        string_param=None,
        integer=None,
        float=None,
        double=None,
        mac=None,
        ipv4=None,
        ipv6=None,
        hex=None,
    ):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def string_param(self):
        # type: () -> str
        """string_param getter

        TBD

        Returns: str
        """
        return self._get_property("string_param")

    @string_param.setter
    def string_param(self, value):
        """string_param setter

        TBD

        value: str
        """
        self._set_property("string_param", value)

    @property
    def integer(self):
        # type: () -> int
        """integer getter

        TBD

        Returns: int
        """
        return self._get_property("integer")

    @integer.setter
    def integer(self, value):
        """integer setter

        TBD

        value: int
        """
        self._set_property("integer", value)

    @property
    def float(self):
        # type: () -> float
        """float getter

        TBD

        Returns: float
        """
        return self._get_property("float")

    @float.setter
    def float(self, value):
        """float setter

        TBD

        value: float
        """
        self._set_property("float", value)

    @property
    def double(self):
        # type: () -> float
        """double getter

        TBD

        Returns: float
        """
        return self._get_property("double")

    @double.setter
    def double(self, value):
        """double setter

        TBD

        value: float
        """
        self._set_property("double", value)

    @property
    def mac(self):
        # type: () -> str
        """mac getter

        TBD

        Returns: str
        """
        return self._get_property("mac")

    @mac.setter
    def mac(self, value):
        """mac setter

        TBD

        value: str
        """
        self._set_property("mac", value)

    @property
    def ipv4(self):
        # type: () -> str
        """ipv4 getter

        TBD

        Returns: str
        """
        return self._get_property("ipv4")

    @ipv4.setter
    def ipv4(self, value):
        """ipv4 setter

        TBD

        value: str
        """
        self._set_property("ipv4", value)

    @property
    def ipv6(self):
        # type: () -> str
        """ipv6 getter

        TBD

        Returns: str
        """
        return self._get_property("ipv6")

    @ipv6.setter
    def ipv6(self, value):
        """ipv6 setter

        TBD

        value: str
        """
        self._set_property("ipv6", value)

    @property
    def hex(self):
        # type: () -> str
        """hex getter

        TBD

        Returns: str
        """
        return self._get_property("hex")

    @hex.setter
    def hex(self, value):
        """hex setter

        TBD

        value: str
        """
        self._set_property("hex", value)


class LevelOne(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "l1_p1": {"type": "LevelTwo"},
        "l1_p2": {"type": "LevelFour"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(LevelOne, self).__init__()
        self._parent = parent

    @property
    def l1_p1(self):
        # type: () -> LevelTwo
        """l1_p1 getter

        Test Level 2Test Level 2Test Level 2Level one

        Returns: LevelTwo
        """
        return self._get_property("l1_p1", LevelTwo)

    @property
    def l1_p2(self):
        # type: () -> LevelFour
        """l1_p2 getter

        Test level4 redundant junk testingTest level4 redundant junk testingTest level4 redundant junk testingLevel one to four

        Returns: LevelFour
        """
        return self._get_property("l1_p2", LevelFour)


class LevelTwo(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "l2_p1": {"type": "LevelThree"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(LevelTwo, self).__init__()
        self._parent = parent

    @property
    def l2_p1(self):
        # type: () -> LevelThree
        """l2_p1 getter

        Test Level3Test Level3Test Level3Level Two

        Returns: LevelThree
        """
        return self._get_property("l2_p1", LevelThree)


class LevelThree(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "l3_p1": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, l3_p1=None):
        super(LevelThree, self).__init__()
        self._parent = parent
        self._set_property("l3_p1", l3_p1)

    def set(self, l3_p1=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def l3_p1(self):
        # type: () -> str
        """l3_p1 getter

        Set value at Level 3

        Returns: str
        """
        return self._get_property("l3_p1")

    @l3_p1.setter
    def l3_p1(self, value):
        """l3_p1 setter

        Set value at Level 3

        value: str
        """
        self._set_property("l3_p1", value)


class LevelFour(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "l4_p1": {"type": "LevelOne"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(LevelFour, self).__init__()
        self._parent = parent

    @property
    def l4_p1(self):
        # type: () -> LevelOne
        """l4_p1 getter

        To Test Multi level non-primitive typesTo Test Multi level non-primitive typesTo Test Multi level non-primitive typesloop over level 1

        Returns: LevelOne
        """
        return self._get_property("l4_p1", LevelOne)


class Mandate(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "required_param": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ("required_param",)  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, required_param=None):
        super(Mandate, self).__init__()
        self._parent = parent
        self._set_property("required_param", required_param)

    def set(self, required_param=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def required_param(self):
        # type: () -> str
        """required_param getter

        TBD

        Returns: str
        """
        return self._get_property("required_param")

    @required_param.setter
    def required_param(self, value):
        """required_param setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError(
                "Cannot set required property required_param as None"
            )
        self._set_property("required_param", value)


class Ipv4Pattern(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "ipv4": {"type": "PatternIpv4PatternIpv4"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(Ipv4Pattern, self).__init__()
        self._parent = parent

    @property
    def ipv4(self):
        # type: () -> PatternIpv4PatternIpv4
        """ipv4 getter

        TBDTBDTBD

        Returns: PatternIpv4PatternIpv4
        """
        return self._get_property("ipv4", PatternIpv4PatternIpv4)


class PatternIpv4PatternIpv4(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "value",
                "values",
                "increment",
                "decrement",
            ],
        },
        "value": {
            "type": str,
            "format": "ipv4",
        },
        "values": {
            "type": list,
            "itemtype": str,
            "itemformat": "ipv4",
        },
        "increment": {"type": "PatternIpv4PatternIpv4Counter"},
        "decrement": {"type": "PatternIpv4PatternIpv4Counter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "value",
        "value": "0.0.0.0",
        "values": ["0.0.0.0"],
    }  # type: Dict[str, Union(type)]

    VALUE = "value"  # type: str
    VALUES = "values"  # type: str
    INCREMENT = "increment"  # type: str
    DECREMENT = "decrement"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self, parent=None, choice=None, value="0.0.0.0", values=["0.0.0.0"]
    ):
        super(PatternIpv4PatternIpv4, self).__init__()
        self._parent = parent
        self._set_property("value", value)
        self._set_property("values", values)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, value=None, values=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def increment(self):
        # type: () -> PatternIpv4PatternIpv4Counter
        """Factory property that returns an instance of the PatternIpv4PatternIpv4Counter class

        ipv4 counter pattern

        Returns: PatternIpv4PatternIpv4Counter
        """
        return self._get_property(
            "increment", PatternIpv4PatternIpv4Counter, self, "increment"
        )

    @property
    def decrement(self):
        # type: () -> PatternIpv4PatternIpv4Counter
        """Factory property that returns an instance of the PatternIpv4PatternIpv4Counter class

        ipv4 counter pattern

        Returns: PatternIpv4PatternIpv4Counter
        """
        return self._get_property(
            "decrement", PatternIpv4PatternIpv4Counter, self, "decrement"
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """choice getter

        TBD

        Returns: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        self._set_property("choice", value)

    @property
    def value(self):
        # type: () -> str
        """value getter

        TBD

        Returns: str
        """
        return self._get_property("value")

    @value.setter
    def value(self, value):
        """value setter

        TBD

        value: str
        """
        self._set_property("value", value, "value")

    @property
    def values(self):
        # type: () -> List[str]
        """values getter

        TBD

        Returns: List[str]
        """
        return self._get_property("values")

    @values.setter
    def values(self, value):
        """values setter

        TBD

        value: List[str]
        """
        self._set_property("values", value, "values")


class PatternIpv4PatternIpv4Counter(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "start": {
            "type": str,
            "format": "ipv4",
        },
        "step": {
            "type": str,
            "format": "ipv4",
        },
        "count": {
            "type": int,
            "format": "uint32",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "start": "0.0.0.0",
        "step": "0.0.0.1",
        "count": 1,
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, start="0.0.0.0", step="0.0.0.1", count=1):
        super(PatternIpv4PatternIpv4Counter, self).__init__()
        self._parent = parent
        self._set_property("start", start)
        self._set_property("step", step)
        self._set_property("count", count)

    def set(self, start=None, step=None, count=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def start(self):
        # type: () -> str
        """start getter

        TBD

        Returns: str
        """
        return self._get_property("start")

    @start.setter
    def start(self, value):
        """start setter

        TBD

        value: str
        """
        self._set_property("start", value)

    @property
    def step(self):
        # type: () -> str
        """step getter

        TBD

        Returns: str
        """
        return self._get_property("step")

    @step.setter
    def step(self, value):
        """step setter

        TBD

        value: str
        """
        self._set_property("step", value)

    @property
    def count(self):
        # type: () -> int
        """count getter

        TBD

        Returns: int
        """
        return self._get_property("count")

    @count.setter
    def count(self, value):
        """count setter

        TBD

        value: int
        """
        self._set_property("count", value)


class Ipv6Pattern(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "ipv6": {"type": "PatternIpv6PatternIpv6"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(Ipv6Pattern, self).__init__()
        self._parent = parent

    @property
    def ipv6(self):
        # type: () -> PatternIpv6PatternIpv6
        """ipv6 getter

        TBDTBDTBD

        Returns: PatternIpv6PatternIpv6
        """
        return self._get_property("ipv6", PatternIpv6PatternIpv6)


class PatternIpv6PatternIpv6(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "value",
                "values",
                "increment",
                "decrement",
            ],
        },
        "value": {
            "type": str,
            "format": "ipv6",
        },
        "values": {
            "type": list,
            "itemtype": str,
            "itemformat": "ipv6",
        },
        "increment": {"type": "PatternIpv6PatternIpv6Counter"},
        "decrement": {"type": "PatternIpv6PatternIpv6Counter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "value",
        "value": "::",
        "values": ["::"],
    }  # type: Dict[str, Union(type)]

    VALUE = "value"  # type: str
    VALUES = "values"  # type: str
    INCREMENT = "increment"  # type: str
    DECREMENT = "decrement"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, value="::", values=["::"]):
        super(PatternIpv6PatternIpv6, self).__init__()
        self._parent = parent
        self._set_property("value", value)
        self._set_property("values", values)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, value=None, values=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def increment(self):
        # type: () -> PatternIpv6PatternIpv6Counter
        """Factory property that returns an instance of the PatternIpv6PatternIpv6Counter class

        ipv6 counter pattern

        Returns: PatternIpv6PatternIpv6Counter
        """
        return self._get_property(
            "increment", PatternIpv6PatternIpv6Counter, self, "increment"
        )

    @property
    def decrement(self):
        # type: () -> PatternIpv6PatternIpv6Counter
        """Factory property that returns an instance of the PatternIpv6PatternIpv6Counter class

        ipv6 counter pattern

        Returns: PatternIpv6PatternIpv6Counter
        """
        return self._get_property(
            "decrement", PatternIpv6PatternIpv6Counter, self, "decrement"
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """choice getter

        TBD

        Returns: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        self._set_property("choice", value)

    @property
    def value(self):
        # type: () -> str
        """value getter

        TBD

        Returns: str
        """
        return self._get_property("value")

    @value.setter
    def value(self, value):
        """value setter

        TBD

        value: str
        """
        self._set_property("value", value, "value")

    @property
    def values(self):
        # type: () -> List[str]
        """values getter

        TBD

        Returns: List[str]
        """
        return self._get_property("values")

    @values.setter
    def values(self, value):
        """values setter

        TBD

        value: List[str]
        """
        self._set_property("values", value, "values")


class PatternIpv6PatternIpv6Counter(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "start": {
            "type": str,
            "format": "ipv6",
        },
        "step": {
            "type": str,
            "format": "ipv6",
        },
        "count": {
            "type": int,
            "format": "uint32",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "start": "::",
        "step": "::1",
        "count": 1,
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, start="::", step="::1", count=1):
        super(PatternIpv6PatternIpv6Counter, self).__init__()
        self._parent = parent
        self._set_property("start", start)
        self._set_property("step", step)
        self._set_property("count", count)

    def set(self, start=None, step=None, count=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def start(self):
        # type: () -> str
        """start getter

        TBD

        Returns: str
        """
        return self._get_property("start")

    @start.setter
    def start(self, value):
        """start setter

        TBD

        value: str
        """
        self._set_property("start", value)

    @property
    def step(self):
        # type: () -> str
        """step getter

        TBD

        Returns: str
        """
        return self._get_property("step")

    @step.setter
    def step(self, value):
        """step setter

        TBD

        value: str
        """
        self._set_property("step", value)

    @property
    def count(self):
        # type: () -> int
        """count getter

        TBD

        Returns: int
        """
        return self._get_property("count")

    @count.setter
    def count(self, value):
        """count setter

        TBD

        value: int
        """
        self._set_property("count", value)


class MacPattern(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "mac": {"type": "PatternMacPatternMac"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(MacPattern, self).__init__()
        self._parent = parent

    @property
    def mac(self):
        # type: () -> PatternMacPatternMac
        """mac getter

        TBDTBDTBD

        Returns: PatternMacPatternMac
        """
        return self._get_property("mac", PatternMacPatternMac)


class PatternMacPatternMac(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "value",
                "values",
                "auto",
                "increment",
                "decrement",
            ],
        },
        "value": {
            "type": str,
            "format": "mac",
        },
        "values": {
            "type": list,
            "itemtype": str,
            "itemformat": "mac",
        },
        "auto": {
            "type": str,
            "format": "mac",
        },
        "increment": {"type": "PatternMacPatternMacCounter"},
        "decrement": {"type": "PatternMacPatternMacCounter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "auto",
        "value": "00:00:00:00:00:00",
        "values": ["00:00:00:00:00:00"],
        "auto": "00:00:00:00:00:00",
    }  # type: Dict[str, Union(type)]

    VALUE = "value"  # type: str
    VALUES = "values"  # type: str
    AUTO = "auto"  # type: str
    INCREMENT = "increment"  # type: str
    DECREMENT = "decrement"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        choice=None,
        value="00:00:00:00:00:00",
        values=["00:00:00:00:00:00"],
        auto="00:00:00:00:00:00",
    ):
        super(PatternMacPatternMac, self).__init__()
        self._parent = parent
        self._set_property("value", value)
        self._set_property("values", values)
        self._set_property("auto", auto)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, value=None, values=None, auto=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def increment(self):
        # type: () -> PatternMacPatternMacCounter
        """Factory property that returns an instance of the PatternMacPatternMacCounter class

        mac counter pattern

        Returns: PatternMacPatternMacCounter
        """
        return self._get_property(
            "increment", PatternMacPatternMacCounter, self, "increment"
        )

    @property
    def decrement(self):
        # type: () -> PatternMacPatternMacCounter
        """Factory property that returns an instance of the PatternMacPatternMacCounter class

        mac counter pattern

        Returns: PatternMacPatternMacCounter
        """
        return self._get_property(
            "decrement", PatternMacPatternMacCounter, self, "decrement"
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """choice getter

        TBD

        Returns: Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        self._set_property("choice", value)

    @property
    def value(self):
        # type: () -> str
        """value getter

        TBD

        Returns: str
        """
        return self._get_property("value")

    @value.setter
    def value(self, value):
        """value setter

        TBD

        value: str
        """
        self._set_property("value", value, "value")

    @property
    def values(self):
        # type: () -> List[str]
        """values getter

        TBD

        Returns: List[str]
        """
        return self._get_property("values")

    @values.setter
    def values(self, value):
        """values setter

        TBD

        value: List[str]
        """
        self._set_property("values", value, "values")

    @property
    def auto(self):
        # type: () -> str
        """auto getter

        The OTG implementation can provide system generated. value for this property. If the OTG is unable to generate value. the default value must be used.

        Returns: str
        """
        return self._get_property("auto")


class PatternMacPatternMacCounter(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "start": {
            "type": str,
            "format": "mac",
        },
        "step": {
            "type": str,
            "format": "mac",
        },
        "count": {
            "type": int,
            "format": "uint32",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "start": "00:00:00:00:00:00",
        "step": "00:00:00:00:00:01",
        "count": 1,
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        start="00:00:00:00:00:00",
        step="00:00:00:00:00:01",
        count=1,
    ):
        super(PatternMacPatternMacCounter, self).__init__()
        self._parent = parent
        self._set_property("start", start)
        self._set_property("step", step)
        self._set_property("count", count)

    def set(self, start=None, step=None, count=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def start(self):
        # type: () -> str
        """start getter

        TBD

        Returns: str
        """
        return self._get_property("start")

    @start.setter
    def start(self, value):
        """start setter

        TBD

        value: str
        """
        self._set_property("start", value)

    @property
    def step(self):
        # type: () -> str
        """step getter

        TBD

        Returns: str
        """
        return self._get_property("step")

    @step.setter
    def step(self, value):
        """step setter

        TBD

        value: str
        """
        self._set_property("step", value)

    @property
    def count(self):
        # type: () -> int
        """count getter

        TBD

        Returns: int
        """
        return self._get_property("count")

    @count.setter
    def count(self, value):
        """count setter

        TBD

        value: int
        """
        self._set_property("count", value)


class IntegerPattern(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "integer": {"type": "PatternIntegerPatternInteger"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(IntegerPattern, self).__init__()
        self._parent = parent

    @property
    def integer(self):
        # type: () -> PatternIntegerPatternInteger
        """integer getter

        TBDTBDTBD

        Returns: PatternIntegerPatternInteger
        """
        return self._get_property("integer", PatternIntegerPatternInteger)


class PatternIntegerPatternInteger(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "value",
                "values",
                "increment",
                "decrement",
            ],
        },
        "value": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "values": {
            "type": list,
            "itemtype": int,
            "itemformat": "uint32",
            "maximum": 255,
        },
        "increment": {"type": "PatternIntegerPatternIntegerCounter"},
        "decrement": {"type": "PatternIntegerPatternIntegerCounter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "value",
        "value": 0,
        "values": [0],
    }  # type: Dict[str, Union(type)]

    VALUE = "value"  # type: str
    VALUES = "values"  # type: str
    INCREMENT = "increment"  # type: str
    DECREMENT = "decrement"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, value=0, values=[0]):
        super(PatternIntegerPatternInteger, self).__init__()
        self._parent = parent
        self._set_property("value", value)
        self._set_property("values", values)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, value=None, values=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def increment(self):
        # type: () -> PatternIntegerPatternIntegerCounter
        """Factory property that returns an instance of the PatternIntegerPatternIntegerCounter class

        integer counter pattern

        Returns: PatternIntegerPatternIntegerCounter
        """
        return self._get_property(
            "increment", PatternIntegerPatternIntegerCounter, self, "increment"
        )

    @property
    def decrement(self):
        # type: () -> PatternIntegerPatternIntegerCounter
        """Factory property that returns an instance of the PatternIntegerPatternIntegerCounter class

        integer counter pattern

        Returns: PatternIntegerPatternIntegerCounter
        """
        return self._get_property(
            "decrement", PatternIntegerPatternIntegerCounter, self, "decrement"
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """choice getter

        TBD

        Returns: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        self._set_property("choice", value)

    @property
    def value(self):
        # type: () -> int
        """value getter

        TBD

        Returns: int
        """
        return self._get_property("value")

    @value.setter
    def value(self, value):
        """value setter

        TBD

        value: int
        """
        self._set_property("value", value, "value")

    @property
    def values(self):
        # type: () -> List[int]
        """values getter

        TBD

        Returns: List[int]
        """
        return self._get_property("values")

    @values.setter
    def values(self, value):
        """values setter

        TBD

        value: List[int]
        """
        self._set_property("values", value, "values")


class PatternIntegerPatternIntegerCounter(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "start": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "step": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "count": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "start": 0,
        "step": 1,
        "count": 1,
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, start=0, step=1, count=1):
        super(PatternIntegerPatternIntegerCounter, self).__init__()
        self._parent = parent
        self._set_property("start", start)
        self._set_property("step", step)
        self._set_property("count", count)

    def set(self, start=None, step=None, count=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def start(self):
        # type: () -> int
        """start getter

        TBD

        Returns: int
        """
        return self._get_property("start")

    @start.setter
    def start(self, value):
        """start setter

        TBD

        value: int
        """
        self._set_property("start", value)

    @property
    def step(self):
        # type: () -> int
        """step getter

        TBD

        Returns: int
        """
        return self._get_property("step")

    @step.setter
    def step(self, value):
        """step setter

        TBD

        value: int
        """
        self._set_property("step", value)

    @property
    def count(self):
        # type: () -> int
        """count getter

        TBD

        Returns: int
        """
        return self._get_property("count")

    @count.setter
    def count(self, value):
        """count setter

        TBD

        value: int
        """
        self._set_property("count", value)


class ChecksumPattern(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "checksum": {"type": "PatternChecksumPatternChecksum"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(ChecksumPattern, self).__init__()
        self._parent = parent

    @property
    def checksum(self):
        # type: () -> PatternChecksumPatternChecksum
        """checksum getter

        TBDTBDTBD

        Returns: PatternChecksumPatternChecksum
        """
        return self._get_property("checksum", PatternChecksumPatternChecksum)


class PatternChecksumPatternChecksum(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "generated",
                "custom",
            ],
        },
        "generated": {
            "type": str,
            "enum": [
                "good",
                "bad",
            ],
        },
        "custom": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "generated",
        "generated": "good",
    }  # type: Dict[str, Union(type)]

    GENERATED = "generated"  # type: str
    CUSTOM = "custom"  # type: str

    GOOD = "good"  # type: str
    BAD = "bad"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self, parent=None, choice=None, generated="good", custom=None
    ):
        super(PatternChecksumPatternChecksum, self).__init__()
        self._parent = parent
        self._set_property("generated", generated)
        self._set_property("custom", custom)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, generated=None, custom=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def choice(self):
        # type: () -> Union[Literal["custom"], Literal["generated"]]
        """choice getter

        The type of checksum

        Returns: Union[Literal["custom"], Literal["generated"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        The type of checksum

        value: Union[Literal["custom"], Literal["generated"]]
        """
        self._set_property("choice", value)

    @property
    def generated(self):
        # type: () -> Union[Literal["bad"], Literal["good"]]
        """generated getter

        A system generated checksum value

        Returns: Union[Literal["bad"], Literal["good"]]
        """
        return self._get_property("generated")

    @generated.setter
    def generated(self, value):
        """generated setter

        A system generated checksum value

        value: Union[Literal["bad"], Literal["good"]]
        """
        self._set_property("generated", value, "generated")

    @property
    def custom(self):
        # type: () -> int
        """custom getter

        A custom checksum value

        Returns: int
        """
        return self._get_property("custom")

    @custom.setter
    def custom(self, value):
        """custom setter

        A custom checksum value

        value: int
        """
        self._set_property("custom", value, "custom")


class Layer1Ieee802x(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "flow_control": {"type": bool},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, flow_control=None):
        super(Layer1Ieee802x, self).__init__()
        self._parent = parent
        self._set_property("flow_control", flow_control)

    def set(self, flow_control=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def flow_control(self):
        # type: () -> bool
        """flow_control getter

        TBD

        Returns: bool
        """
        return self._get_property("flow_control")

    @flow_control.setter
    def flow_control(self, value):
        """flow_control setter

        TBD

        value: bool
        """
        self._set_property("flow_control", value)


class MObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "string_param": {"type": str},
        "integer": {
            "type": int,
            "format": "int32",
            "minimum": -10,
            "maximum": 90,
        },
        "float": {
            "type": float,
            "format": "float",
        },
        "double": {
            "type": float,
            "format": "double",
        },
        "mac": {
            "type": str,
            "format": "mac",
        },
        "ipv4": {
            "type": str,
            "format": "ipv4",
        },
        "ipv6": {
            "type": str,
            "format": "ipv6",
        },
        "hex": {
            "type": str,
            "format": "hex",
        },
    }  # type: Dict[str, str]

    _REQUIRED = (
        "string_param",
        "integer",
        "float",
        "double",
        "mac",
        "ipv4",
        "ipv6",
        "hex",
    )  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        string_param=None,
        integer=None,
        float=None,
        double=None,
        mac=None,
        ipv4=None,
        ipv6=None,
        hex=None,
    ):
        super(MObject, self).__init__()
        self._parent = parent
        self._set_property("string_param", string_param)
        self._set_property("integer", integer)
        self._set_property("float", float)
        self._set_property("double", double)
        self._set_property("mac", mac)
        self._set_property("ipv4", ipv4)
        self._set_property("ipv6", ipv6)
        self._set_property("hex", hex)

    def set(
        self,
        string_param=None,
        integer=None,
        float=None,
        double=None,
        mac=None,
        ipv4=None,
        ipv6=None,
        hex=None,
    ):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def string_param(self):
        # type: () -> str
        """string_param getter

        TBD

        Returns: str
        """
        return self._get_property("string_param")

    @string_param.setter
    def string_param(self, value):
        """string_param setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError(
                "Cannot set required property string_param as None"
            )
        self._set_property("string_param", value)

    @property
    def integer(self):
        # type: () -> int
        """integer getter

        TBD

        Returns: int
        """
        return self._get_property("integer")

    @integer.setter
    def integer(self, value):
        """integer setter

        TBD

        value: int
        """
        if value is None:
            raise TypeError("Cannot set required property integer as None")
        self._set_property("integer", value)

    @property
    def float(self):
        # type: () -> float
        """float getter

        TBD

        Returns: float
        """
        return self._get_property("float")

    @float.setter
    def float(self, value):
        """float setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property float as None")
        self._set_property("float", value)

    @property
    def double(self):
        # type: () -> float
        """double getter

        TBD

        Returns: float
        """
        return self._get_property("double")

    @double.setter
    def double(self, value):
        """double setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property double as None")
        self._set_property("double", value)

    @property
    def mac(self):
        # type: () -> str
        """mac getter

        TBD

        Returns: str
        """
        return self._get_property("mac")

    @mac.setter
    def mac(self, value):
        """mac setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property mac as None")
        self._set_property("mac", value)

    @property
    def ipv4(self):
        # type: () -> str
        """ipv4 getter

        TBD

        Returns: str
        """
        return self._get_property("ipv4")

    @ipv4.setter
    def ipv4(self, value):
        """ipv4 setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property ipv4 as None")
        self._set_property("ipv4", value)

    @property
    def ipv6(self):
        # type: () -> str
        """ipv6 getter

        TBD

        Returns: str
        """
        return self._get_property("ipv6")

    @ipv6.setter
    def ipv6(self, value):
        """ipv6 setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property ipv6 as None")
        self._set_property("ipv6", value)

    @property
    def hex(self):
        # type: () -> str
        """hex getter

        TBD

        Returns: str
        """
        return self._get_property("hex")

    @hex.setter
    def hex(self, value):
        """hex setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property hex as None")
        self._set_property("hex", value)


class PatternPrefixConfigHeaderChecksum(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "generated",
                "custom",
            ],
        },
        "generated": {
            "type": str,
            "enum": [
                "good",
                "bad",
            ],
        },
        "custom": {
            "type": int,
            "format": "uint32",
            "maximum": 65535,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "generated",
        "generated": "good",
    }  # type: Dict[str, Union(type)]

    GENERATED = "generated"  # type: str
    CUSTOM = "custom"  # type: str

    GOOD = "good"  # type: str
    BAD = "bad"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self, parent=None, choice=None, generated="good", custom=None
    ):
        super(PatternPrefixConfigHeaderChecksum, self).__init__()
        self._parent = parent
        self._set_property("generated", generated)
        self._set_property("custom", custom)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, generated=None, custom=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def choice(self):
        # type: () -> Union[Literal["custom"], Literal["generated"]]
        """choice getter

        The type of checksum

        Returns: Union[Literal["custom"], Literal["generated"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        The type of checksum

        value: Union[Literal["custom"], Literal["generated"]]
        """
        self._set_property("choice", value)

    @property
    def generated(self):
        # type: () -> Union[Literal["bad"], Literal["good"]]
        """generated getter

        A system generated checksum value

        Returns: Union[Literal["bad"], Literal["good"]]
        """
        return self._get_property("generated")

    @generated.setter
    def generated(self, value):
        """generated setter

        A system generated checksum value

        value: Union[Literal["bad"], Literal["good"]]
        """
        self._set_property("generated", value, "generated")

    @property
    def custom(self):
        # type: () -> int
        """custom getter

        A custom checksum value

        Returns: int
        """
        return self._get_property("custom")

    @custom.setter
    def custom(self, value):
        """custom setter

        A custom checksum value

        value: int
        """
        self._set_property("custom", value, "custom")


class PatternPrefixConfigAutoFieldTest(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "value",
                "values",
                "auto",
                "increment",
                "decrement",
            ],
        },
        "value": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "values": {
            "type": list,
            "itemtype": int,
            "itemformat": "uint32",
            "maximum": 255,
        },
        "auto": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "increment": {"type": "PatternPrefixConfigAutoFieldTestCounter"},
        "decrement": {"type": "PatternPrefixConfigAutoFieldTestCounter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "auto",
        "value": 0,
        "values": [0],
        "auto": 0,
    }  # type: Dict[str, Union(type)]

    VALUE = "value"  # type: str
    VALUES = "values"  # type: str
    AUTO = "auto"  # type: str
    INCREMENT = "increment"  # type: str
    DECREMENT = "decrement"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, value=0, values=[0], auto=0):
        super(PatternPrefixConfigAutoFieldTest, self).__init__()
        self._parent = parent
        self._set_property("value", value)
        self._set_property("values", values)
        self._set_property("auto", auto)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, value=None, values=None, auto=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def increment(self):
        # type: () -> PatternPrefixConfigAutoFieldTestCounter
        """Factory property that returns an instance of the PatternPrefixConfigAutoFieldTestCounter class

        integer counter pattern

        Returns: PatternPrefixConfigAutoFieldTestCounter
        """
        return self._get_property(
            "increment",
            PatternPrefixConfigAutoFieldTestCounter,
            self,
            "increment",
        )

    @property
    def decrement(self):
        # type: () -> PatternPrefixConfigAutoFieldTestCounter
        """Factory property that returns an instance of the PatternPrefixConfigAutoFieldTestCounter class

        integer counter pattern

        Returns: PatternPrefixConfigAutoFieldTestCounter
        """
        return self._get_property(
            "decrement",
            PatternPrefixConfigAutoFieldTestCounter,
            self,
            "decrement",
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """choice getter

        TBD

        Returns: Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["auto"], Literal["decrement"], Literal["increment"], Literal["value"], Literal["values"]]
        """
        self._set_property("choice", value)

    @property
    def value(self):
        # type: () -> int
        """value getter

        TBD

        Returns: int
        """
        return self._get_property("value")

    @value.setter
    def value(self, value):
        """value setter

        TBD

        value: int
        """
        self._set_property("value", value, "value")

    @property
    def values(self):
        # type: () -> List[int]
        """values getter

        TBD

        Returns: List[int]
        """
        return self._get_property("values")

    @values.setter
    def values(self, value):
        """values setter

        TBD

        value: List[int]
        """
        self._set_property("values", value, "values")

    @property
    def auto(self):
        # type: () -> int
        """auto getter

        The OTG implementation can provide system generated. value for this property. If the OTG is unable to generate value. the default value must be used.

        Returns: int
        """
        return self._get_property("auto")


class PatternPrefixConfigAutoFieldTestCounter(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "start": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "step": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
        "count": {
            "type": int,
            "format": "uint32",
            "maximum": 255,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "start": 0,
        "step": 1,
        "count": 1,
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, start=0, step=1, count=1):
        super(PatternPrefixConfigAutoFieldTestCounter, self).__init__()
        self._parent = parent
        self._set_property("start", start)
        self._set_property("step", step)
        self._set_property("count", count)

    def set(self, start=None, step=None, count=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def start(self):
        # type: () -> int
        """start getter

        TBD

        Returns: int
        """
        return self._get_property("start")

    @start.setter
    def start(self, value):
        """start setter

        TBD

        value: int
        """
        self._set_property("start", value)

    @property
    def step(self):
        # type: () -> int
        """step getter

        TBD

        Returns: int
        """
        return self._get_property("step")

    @step.setter
    def step(self, value):
        """step setter

        TBD

        value: int
        """
        self._set_property("step", value)

    @property
    def count(self):
        # type: () -> int
        """count getter

        TBD

        Returns: int
        """
        return self._get_property("count")

    @count.setter
    def count(self, value):
        """count setter

        TBD

        value: int
        """
        self._set_property("count", value)


class WObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "w_name": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ("w_name",)  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, w_name=None):
        super(WObject, self).__init__()
        self._parent = parent
        self._set_property("w_name", w_name)

    def set(self, w_name=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def w_name(self):
        # type: () -> str
        """w_name getter

        TBD

        Returns: str
        """
        return self._get_property("w_name")

    @w_name.setter
    def w_name(self, value):
        """w_name setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property w_name as None")
        self._set_property("w_name", value)


class WObjectIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(WObjectIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[WObject]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> WObjectIter
        return self._iter()

    def __next__(self):
        # type: () -> WObject
        return self._next()

    def next(self):
        # type: () -> WObject
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, WObject):
            raise Exception("Item is not an instance of WObject")

    def wobject(self, w_name=None):
        # type: (str) -> WObjectIter
        """Factory method that creates an instance of the WObject class

        TBD

        Returns: WObjectIter
        """
        item = WObject(parent=self._parent, w_name=w_name)
        self._add(item)
        return self

    def add(self, w_name=None):
        # type: (str) -> WObject
        """Add method that creates and returns an instance of the WObject class

        TBD

        Returns: WObject
        """
        item = WObject(parent=self._parent, w_name=w_name)
        self._add(item)
        return item


class ZObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "name": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ("name",)  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, name=None):
        super(ZObject, self).__init__()
        self._parent = parent
        self._set_property("name", name)

    def set(self, name=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property name as None")
        self._set_property("name", value)


class ZObjectIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(ZObjectIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[ZObject]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> ZObjectIter
        return self._iter()

    def __next__(self):
        # type: () -> ZObject
        return self._next()

    def next(self):
        # type: () -> ZObject
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, ZObject):
            raise Exception("Item is not an instance of ZObject")

    def zobject(self, name=None):
        # type: (str) -> ZObjectIter
        """Factory method that creates an instance of the ZObject class

        TBD

        Returns: ZObjectIter
        """
        item = ZObject(parent=self._parent, name=name)
        self._add(item)
        return self

    def add(self, name=None):
        # type: (str) -> ZObject
        """Add method that creates and returns an instance of the ZObject class

        TBD

        Returns: ZObject
        """
        item = ZObject(parent=self._parent, name=name)
        self._add(item)
        return item


class YObject(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "y_name": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, y_name=None):
        super(YObject, self).__init__()
        self._parent = parent
        self._set_property("y_name", y_name)

    def set(self, y_name=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def y_name(self):
        # type: () -> str
        """y_name getter

        TBD. x-constraint:. /components/schemas/ZObject/properties/name. /components/schemas/WObject/properties/w_name.

        Returns: str
        """
        return self._get_property("y_name")

    @y_name.setter
    def y_name(self, value):
        """y_name setter

        TBD. x-constraint:. /components/schemas/ZObject/properties/name. /components/schemas/WObject/properties/w_name.

        value: str
        """
        self._set_property("y_name", value)


class ChoiceObject(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "e_obj",
                "f_obj",
                "no_obj",
            ],
        },
        "e_obj": {"type": "EObject"},
        "f_obj": {"type": "FObject"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "no_obj",
    }  # type: Dict[str, Union(type)]

    E_OBJ = "e_obj"  # type: str
    F_OBJ = "f_obj"  # type: str
    NO_OBJ = "no_obj"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None):
        super(ChoiceObject, self).__init__()
        self._parent = parent
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    @property
    def e_obj(self):
        # type: () -> EObject
        """Factory property that returns an instance of the EObject class

        TBD

        Returns: EObject
        """
        return self._get_property("e_obj", EObject, self, "e_obj")

    @property
    def f_obj(self):
        # type: () -> FObject
        """Factory property that returns an instance of the FObject class

        TBD

        Returns: FObject
        """
        return self._get_property("f_obj", FObject, self, "f_obj")

    @property
    def choice(self):
        # type: () -> Union[Literal["e_obj"], Literal["f_obj"], Literal["no_obj"]]
        """choice getter

        TBD

        Returns: Union[Literal["e_obj"], Literal["f_obj"], Literal["no_obj"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["e_obj"], Literal["f_obj"], Literal["no_obj"]]
        """
        self._set_property("choice", value)


class ChoiceObjectIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = True

    def __init__(self, parent=None, choice=None):
        super(ChoiceObjectIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[ChoiceObject, EObject, FObject]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> ChoiceObjectIter
        return self._iter()

    def __next__(self):
        # type: () -> ChoiceObject
        return self._next()

    def next(self):
        # type: () -> ChoiceObject
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, ChoiceObject):
            raise Exception("Item is not an instance of ChoiceObject")

    def choiceobject(self):
        # type: () -> ChoiceObjectIter
        """Factory method that creates an instance of the ChoiceObject class

        TBD

        Returns: ChoiceObjectIter
        """
        item = ChoiceObject(parent=self._parent, choice=self._choice)
        self._add(item)
        return self

    def add(self):
        # type: () -> ChoiceObject
        """Add method that creates and returns an instance of the ChoiceObject class

        TBD

        Returns: ChoiceObject
        """
        item = ChoiceObject(parent=self._parent, choice=self._choice)
        self._add(item)
        return item

    def e_obj(
        self, e_a=None, e_b=None, name=None, m_param1=None, m_param2=None
    ):
        # type: (float,float,str,str,str) -> ChoiceObjectIter
        """Factory method that creates an instance of the EObject class

        TBD

        Returns: ChoiceObjectIter
        """
        item = ChoiceObject()
        item.e_obj
        item.choice = "e_obj"
        self._add(item)
        return self

    def f_obj(self, f_a="some string", f_b=3.0):
        # type: (str,float) -> ChoiceObjectIter
        """Factory method that creates an instance of the FObject class

        TBD

        Returns: ChoiceObjectIter
        """
        item = ChoiceObject()
        item.f_obj
        item.choice = "f_obj"
        self._add(item)
        return self


class RequiredChoiceParent(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "intermediate_obj",
                "no_obj",
            ],
        },
        "intermediate_obj": {"type": "RequiredChoiceIntermediate"},
    }  # type: Dict[str, str]

    _REQUIRED = ("choice",)  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    INTERMEDIATE_OBJ = "intermediate_obj"  # type: str
    NO_OBJ = "no_obj"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None):
        super(RequiredChoiceParent, self).__init__()
        self._parent = parent
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    @property
    def intermediate_obj(self):
        # type: () -> RequiredChoiceIntermediate
        """Factory property that returns an instance of the RequiredChoiceIntermediate class

        TBD

        Returns: RequiredChoiceIntermediate
        """
        return self._get_property(
            "intermediate_obj",
            RequiredChoiceIntermediate,
            self,
            "intermediate_obj",
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["intermediate_obj"], Literal["no_obj"]]
        """choice getter

        TBD

        Returns: Union[Literal["intermediate_obj"], Literal["no_obj"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["intermediate_obj"], Literal["no_obj"]]
        """
        if value is None:
            raise TypeError("Cannot set required property choice as None")
        self._set_property("choice", value)


class RequiredChoiceIntermediate(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "f_a",
                "leaf",
            ],
        },
        "f_a": {"type": str},
        "leaf": {"type": "RequiredChoiceIntermeLeaf"},
    }  # type: Dict[str, str]

    _REQUIRED = ("choice",)  # type: tuple(str)

    _DEFAULTS = {
        "f_a": "some string",
    }  # type: Dict[str, Union(type)]

    F_A = "f_a"  # type: str
    LEAF = "leaf"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, f_a="some string"):
        super(RequiredChoiceIntermediate, self).__init__()
        self._parent = parent
        self._set_property("f_a", f_a)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, f_a=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def leaf(self):
        # type: () -> RequiredChoiceIntermeLeaf
        """Factory property that returns an instance of the RequiredChoiceIntermeLeaf class

        TBD

        Returns: RequiredChoiceIntermeLeaf
        """
        return self._get_property(
            "leaf", RequiredChoiceIntermeLeaf, self, "leaf"
        )

    @property
    def choice(self):
        # type: () -> Union[Literal["f_a"], Literal["leaf"]]
        """choice getter

        TBD

        Returns: Union[Literal["f_a"], Literal["leaf"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["f_a"], Literal["leaf"]]
        """
        if value is None:
            raise TypeError("Cannot set required property choice as None")
        self._set_property("choice", value)

    @property
    def f_a(self):
        # type: () -> str
        """f_a getter

        TBD

        Returns: str
        """
        return self._get_property("f_a")

    @f_a.setter
    def f_a(self, value):
        """f_a setter

        TBD

        value: str
        """
        self._set_property("f_a", value, "f_a")


class RequiredChoiceIntermeLeaf(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "name": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, name=None):
        super(RequiredChoiceIntermeLeaf, self).__init__()
        self._parent = parent
        self._set_property("name", name)

    def set(self, name=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        self._set_property("name", value)


class Error(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "code": {
            "type": int,
            "format": "int32",
        },
        "kind": {
            "type": str,
            "enum": [
                "transport",
                "validation",
                "internal",
            ],
        },
        "errors": {
            "type": list,
            "itemtype": str,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ("code", "errors")  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    TRANSPORT = "transport"  # type: str
    VALIDATION = "validation"  # type: str
    INTERNAL = "internal"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, code=None, kind=None, errors=None):
        super(Error, self).__init__()
        self._parent = parent
        self._set_property("code", code)
        self._set_property("kind", kind)
        self._set_property("errors", errors)

    def set(self, code=None, kind=None, errors=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def code(self):
        # type: () -> int
        """code getter

        Numeric status code based on underlying transport being used.

        Returns: int
        """
        return self._get_property("code")

    @code.setter
    def code(self, value):
        """code setter

        Numeric status code based on underlying transport being used.

        value: int
        """
        if value is None:
            raise TypeError("Cannot set required property code as None")
        self._set_property("code", value)

    @property
    def kind(self):
        # type: () -> Union[Literal["internal"], Literal["transport"], Literal["validation"]]
        """kind getter

        Kind of error message.

        Returns: Union[Literal["internal"], Literal["transport"], Literal["validation"]]
        """
        return self._get_property("kind")

    @kind.setter
    def kind(self, value):
        """kind setter

        Kind of error message.

        value: Union[Literal["internal"], Literal["transport"], Literal["validation"]]
        """
        self._set_property("kind", value)

    @property
    def errors(self):
        # type: () -> List[str]
        """errors getter

        List of error messages generated while serving API request.

        Returns: List[str]
        """
        return self._get_property("errors")

    @errors.setter
    def errors(self, value):
        """errors setter

        List of error messages generated while serving API request.

        value: List[str]
        """
        if value is None:
            raise TypeError("Cannot set required property errors as None")
        self._set_property("errors", value)


class UpdateConfig(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "g": {"type": "GObjectIter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {
        "self": "UpdateConfig is under_review, the whole schema is being reviewed",
    }  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(UpdateConfig, self).__init__()
        self._parent = parent

    @property
    def g(self):
        # type: () -> GObjectIter
        """g getter

        A list of objects with choice and properties

        Returns: GObjectIter
        """
        return self._get_property("g", GObjectIter, self._parent, self._choice)


class MetricsRequest(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "port",
                "flow",
            ],
        },
        "port": {"type": str},
        "flow": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "port",
    }  # type: Dict[str, Union(type)]

    PORT = "port"  # type: str
    FLOW = "flow"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None, port=None, flow=None):
        super(MetricsRequest, self).__init__()
        self._parent = parent
        self._set_property("port", port)
        self._set_property("flow", flow)
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    def set(self, port=None, flow=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def choice(self):
        # type: () -> Union[Literal["flow"], Literal["port"]]
        """choice getter

        TBD

        Returns: Union[Literal["flow"], Literal["port"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["flow"], Literal["port"]]
        """
        self._set_property("choice", value)

    @property
    def port(self):
        # type: () -> str
        """port getter

        TBD

        Returns: str
        """
        return self._get_property("port")

    @port.setter
    def port(self, value):
        """port setter

        TBD

        value: str
        """
        self._set_property("port", value, "port")

    @property
    def flow(self):
        # type: () -> str
        """flow getter

        TBD

        Returns: str
        """
        return self._get_property("flow")

    @flow.setter
    def flow(self, value):
        """flow setter

        TBD

        value: str
        """
        self._set_property("flow", value, "flow")


class Metrics(OpenApiObject):
    __slots__ = ("_parent", "_choice")

    _TYPES = {
        "choice": {
            "type": str,
            "enum": [
                "ports",
                "flows",
            ],
        },
        "ports": {"type": "PortMetricIter"},
        "flows": {"type": "FlowMetricIter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "choice": "ports",
    }  # type: Dict[str, Union(type)]

    PORTS = "ports"  # type: str
    FLOWS = "flows"  # type: str

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, choice=None):
        super(Metrics, self).__init__()
        self._parent = parent
        if (
            "choice" in self._DEFAULTS
            and choice is None
            and self._DEFAULTS["choice"] in self._TYPES
        ):
            getattr(self, self._DEFAULTS["choice"])
        else:
            self._set_property("choice", choice)

    @property
    def choice(self):
        # type: () -> Union[Literal["flows"], Literal["ports"]]
        """choice getter

        TBD

        Returns: Union[Literal["flows"], Literal["ports"]]
        """
        return self._get_property("choice")

    @choice.setter
    def choice(self, value):
        """choice setter

        TBD

        value: Union[Literal["flows"], Literal["ports"]]
        """
        self._set_property("choice", value)

    @property
    def ports(self):
        # type: () -> PortMetricIter
        """ports getter

        TBD

        Returns: PortMetricIter
        """
        return self._get_property(
            "ports", PortMetricIter, self._parent, self._choice
        )

    @property
    def flows(self):
        # type: () -> FlowMetricIter
        """flows getter

        TBD

        Returns: FlowMetricIter
        """
        return self._get_property(
            "flows", FlowMetricIter, self._parent, self._choice
        )


class PortMetric(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "name": {"type": str},
        "tx_frames": {
            "type": float,
            "format": "double",
        },
        "rx_frames": {
            "type": float,
            "format": "double",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ("name", "tx_frames", "rx_frames")  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, name=None, tx_frames=None, rx_frames=None):
        super(PortMetric, self).__init__()
        self._parent = parent
        self._set_property("name", name)
        self._set_property("tx_frames", tx_frames)
        self._set_property("rx_frames", rx_frames)

    def set(self, name=None, tx_frames=None, rx_frames=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property name as None")
        self._set_property("name", value)

    @property
    def tx_frames(self):
        # type: () -> float
        """tx_frames getter

        TBD

        Returns: float
        """
        return self._get_property("tx_frames")

    @tx_frames.setter
    def tx_frames(self, value):
        """tx_frames setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property tx_frames as None")
        self._set_property("tx_frames", value)

    @property
    def rx_frames(self):
        # type: () -> float
        """rx_frames getter

        TBD

        Returns: float
        """
        return self._get_property("rx_frames")

    @rx_frames.setter
    def rx_frames(self, value):
        """rx_frames setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property rx_frames as None")
        self._set_property("rx_frames", value)


class PortMetricIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(PortMetricIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[PortMetric]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> PortMetricIter
        return self._iter()

    def __next__(self):
        # type: () -> PortMetric
        return self._next()

    def next(self):
        # type: () -> PortMetric
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, PortMetric):
            raise Exception("Item is not an instance of PortMetric")

    def metric(self, name=None, tx_frames=None, rx_frames=None):
        # type: (str,float,float) -> PortMetricIter
        """Factory method that creates an instance of the PortMetric class

        TBD

        Returns: PortMetricIter
        """
        item = PortMetric(
            parent=self._parent,
            name=name,
            tx_frames=tx_frames,
            rx_frames=rx_frames,
        )
        self._add(item)
        return self

    def add(self, name=None, tx_frames=None, rx_frames=None):
        # type: (str,float,float) -> PortMetric
        """Add method that creates and returns an instance of the PortMetric class

        TBD

        Returns: PortMetric
        """
        item = PortMetric(
            parent=self._parent,
            name=name,
            tx_frames=tx_frames,
            rx_frames=rx_frames,
        )
        self._add(item)
        return item


class FlowMetric(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "name": {"type": str},
        "tx_frames": {
            "type": float,
            "format": "double",
        },
        "rx_frames": {
            "type": float,
            "format": "double",
        },
    }  # type: Dict[str, str]

    _REQUIRED = ("name", "tx_frames", "rx_frames")  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, name=None, tx_frames=None, rx_frames=None):
        super(FlowMetric, self).__init__()
        self._parent = parent
        self._set_property("name", name)
        self._set_property("tx_frames", tx_frames)
        self._set_property("rx_frames", rx_frames)

    def set(self, name=None, tx_frames=None, rx_frames=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def name(self):
        # type: () -> str
        """name getter

        TBD

        Returns: str
        """
        return self._get_property("name")

    @name.setter
    def name(self, value):
        """name setter

        TBD

        value: str
        """
        if value is None:
            raise TypeError("Cannot set required property name as None")
        self._set_property("name", value)

    @property
    def tx_frames(self):
        # type: () -> float
        """tx_frames getter

        TBD

        Returns: float
        """
        return self._get_property("tx_frames")

    @tx_frames.setter
    def tx_frames(self, value):
        """tx_frames setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property tx_frames as None")
        self._set_property("tx_frames", value)

    @property
    def rx_frames(self):
        # type: () -> float
        """rx_frames getter

        TBD

        Returns: float
        """
        return self._get_property("rx_frames")

    @rx_frames.setter
    def rx_frames(self, value):
        """rx_frames setter

        TBD

        value: float
        """
        if value is None:
            raise TypeError("Cannot set required property rx_frames as None")
        self._set_property("rx_frames", value)


class FlowMetricIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(FlowMetricIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[FlowMetric]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> FlowMetricIter
        return self._iter()

    def __next__(self):
        # type: () -> FlowMetric
        return self._next()

    def next(self):
        # type: () -> FlowMetric
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, FlowMetric):
            raise Exception("Item is not an instance of FlowMetric")

    def metric(self, name=None, tx_frames=None, rx_frames=None):
        # type: (str,float,float) -> FlowMetricIter
        """Factory method that creates an instance of the FlowMetric class

        TBD

        Returns: FlowMetricIter
        """
        item = FlowMetric(
            parent=self._parent,
            name=name,
            tx_frames=tx_frames,
            rx_frames=rx_frames,
        )
        self._add(item)
        return self

    def add(self, name=None, tx_frames=None, rx_frames=None):
        # type: (str,float,float) -> FlowMetric
        """Add method that creates and returns an instance of the FlowMetric class

        TBD

        Returns: FlowMetric
        """
        item = FlowMetric(
            parent=self._parent,
            name=name,
            tx_frames=tx_frames,
            rx_frames=rx_frames,
        )
        self._add(item)
        return item


class WarningDetails(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "warnings": {
            "type": list,
            "itemtype": str,
        },
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, warnings=None):
        super(WarningDetails, self).__init__()
        self._parent = parent
        self._set_property("warnings", warnings)

    def set(self, warnings=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def warnings(self):
        # type: () -> List[str]
        """warnings getter

        TBD

        Returns: List[str]
        """
        return self._get_property("warnings")

    @warnings.setter
    def warnings(self, value):
        """warnings setter

        TBD

        value: List[str]
        """
        self._set_property("warnings", value)


class CommonResponseSuccess(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "message": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, message=None):
        super(CommonResponseSuccess, self).__init__()
        self._parent = parent
        self._set_property("message", message)

    def set(self, message=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def message(self):
        # type: () -> str
        """message getter

        TBD

        Returns: str
        """
        return self._get_property("message")

    @message.setter
    def message(self, value):
        """message setter

        TBD

        value: str
        """
        self._set_property("message", value)


class ApiTestInputBody(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "some_string": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None, some_string=None):
        super(ApiTestInputBody, self).__init__()
        self._parent = parent
        self._set_property("some_string", some_string)

    def set(self, some_string=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def some_string(self):
        # type: () -> str
        """some_string getter

        TBD

        Returns: str
        """
        return self._get_property("some_string")

    @some_string.setter
    def some_string(self, value):
        """some_string setter

        TBD

        value: str
        """
        self._set_property("some_string", value)


class ServiceAbcItemList(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "items": {"type": "ServiceAbcItemIter"},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(self, parent=None):
        super(ServiceAbcItemList, self).__init__()
        self._parent = parent

    @property
    def items(self):
        # type: () -> ServiceAbcItemIter
        """items getter

        TBD

        Returns: ServiceAbcItemIter
        """
        return self._get_property(
            "items", ServiceAbcItemIter, self._parent, self._choice
        )


class ServiceAbcItem(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "some_id": {"type": str},
        "some_string": {"type": str},
        "path_id": {"type": str},
        "level_2": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {}  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self,
        parent=None,
        some_id=None,
        some_string=None,
        path_id=None,
        level_2=None,
    ):
        super(ServiceAbcItem, self).__init__()
        self._parent = parent
        self._set_property("some_id", some_id)
        self._set_property("some_string", some_string)
        self._set_property("path_id", path_id)
        self._set_property("level_2", level_2)

    def set(self, some_id=None, some_string=None, path_id=None, level_2=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def some_id(self):
        # type: () -> str
        """some_id getter

        TBD

        Returns: str
        """
        return self._get_property("some_id")

    @some_id.setter
    def some_id(self, value):
        """some_id setter

        TBD

        value: str
        """
        self._set_property("some_id", value)

    @property
    def some_string(self):
        # type: () -> str
        """some_string getter

        TBD

        Returns: str
        """
        return self._get_property("some_string")

    @some_string.setter
    def some_string(self, value):
        """some_string setter

        TBD

        value: str
        """
        self._set_property("some_string", value)

    @property
    def path_id(self):
        # type: () -> str
        """path_id getter

        TBD

        Returns: str
        """
        return self._get_property("path_id")

    @path_id.setter
    def path_id(self, value):
        """path_id setter

        TBD

        value: str
        """
        self._set_property("path_id", value)

    @property
    def level_2(self):
        # type: () -> str
        """level_2 getter

        TBD

        Returns: str
        """
        return self._get_property("level_2")

    @level_2.setter
    def level_2(self, value):
        """level_2 setter

        TBD

        value: str
        """
        self._set_property("level_2", value)


class ServiceAbcItemIter(OpenApiIter):
    __slots__ = ("_parent", "_choice")

    _GETITEM_RETURNS_CHOICE_OBJECT = False

    def __init__(self, parent=None, choice=None):
        super(ServiceAbcItemIter, self).__init__()
        self._parent = parent
        self._choice = choice

    def __getitem__(self, key):
        # type: (str) -> Union[ServiceAbcItem]
        return self._getitem(key)

    def __iter__(self):
        # type: () -> ServiceAbcItemIter
        return self._iter()

    def __next__(self):
        # type: () -> ServiceAbcItem
        return self._next()

    def next(self):
        # type: () -> ServiceAbcItem
        return self._next()

    def _instanceOf(self, item):
        if not isinstance(item, ServiceAbcItem):
            raise Exception("Item is not an instance of ServiceAbcItem")

    def item(self, some_id=None, some_string=None, path_id=None, level_2=None):
        # type: (str,str,str,str) -> ServiceAbcItemIter
        """Factory method that creates an instance of the ServiceAbcItem class

        TBD

        Returns: ServiceAbcItemIter
        """
        item = ServiceAbcItem(
            parent=self._parent,
            some_id=some_id,
            some_string=some_string,
            path_id=path_id,
            level_2=level_2,
        )
        self._add(item)
        return self

    def add(self, some_id=None, some_string=None, path_id=None, level_2=None):
        # type: (str,str,str,str) -> ServiceAbcItem
        """Add method that creates and returns an instance of the ServiceAbcItem class

        TBD

        Returns: ServiceAbcItem
        """
        item = ServiceAbcItem(
            parent=self._parent,
            some_id=some_id,
            some_string=some_string,
            path_id=path_id,
            level_2=level_2,
        )
        self._add(item)
        return item


class Version(OpenApiObject):
    __slots__ = "_parent"

    _TYPES = {
        "api_spec_version": {"type": str},
        "sdk_version": {"type": str},
        "app_version": {"type": str},
    }  # type: Dict[str, str]

    _REQUIRED = ()  # type: tuple(str)

    _DEFAULTS = {
        "api_spec_version": "",
        "sdk_version": "",
        "app_version": "",
    }  # type: Dict[str, Union(type)]

    _STATUS = {}  # type: Dict[str, Union(type)]

    def __init__(
        self, parent=None, api_spec_version="", sdk_version="", app_version=""
    ):
        super(Version, self).__init__()
        self._parent = parent
        self._set_property("api_spec_version", api_spec_version)
        self._set_property("sdk_version", sdk_version)
        self._set_property("app_version", app_version)

    def set(self, api_spec_version=None, sdk_version=None, app_version=None):
        for property_name, property_value in locals().items():
            if property_name != "self" and property_value is not None:
                self._set_property(property_name, property_value)

    @property
    def api_spec_version(self):
        # type: () -> str
        """api_spec_version getter

        Version of API specification

        Returns: str
        """
        return self._get_property("api_spec_version")

    @api_spec_version.setter
    def api_spec_version(self, value):
        """api_spec_version setter

        Version of API specification

        value: str
        """
        self._set_property("api_spec_version", value)

    @property
    def sdk_version(self):
        # type: () -> str
        """sdk_version getter

        Version of SDK generated from API specification

        Returns: str
        """
        return self._get_property("sdk_version")

    @sdk_version.setter
    def sdk_version(self, value):
        """sdk_version setter

        Version of SDK generated from API specification

        value: str
        """
        self._set_property("sdk_version", value)

    @property
    def app_version(self):
        # type: () -> str
        """app_version getter

        Version of application consuming or serving the API

        Returns: str
        """
        return self._get_property("app_version")

    @app_version.setter
    def app_version(self, value):
        """app_version setter

        Version of application consuming or serving the API

        value: str
        """
        self._set_property("app_version", value)


class Api(object):
    """OpenApi Abstract API"""

    __warnings__ = []

    def __init__(self, **kwargs):
        self._version_meta = self.version()
        self._version_meta.api_spec_version = "0.0.1"
        self._version_meta.sdk_version = ""
        self._version_check = kwargs.get("version_check")
        if self._version_check is None:
            self._version_check = False
        self._version_check_err = None

    def add_warnings(self, msg):
        print("[WARNING]: %s" % msg)
        self.__warnings__.append(msg)

    def _deserialize_error(self, err_string):
        # type: (str) -> Union[Error, None]
        err = self.error()
        try:
            err.deserialize(err_string)
        except Exception:
            err = None
        return err

    def from_exception(self, error):
        # type: (Exception) -> Union[Error, None]
        if isinstance(error, Error):
            return error
        elif isinstance(error, grpc.RpcError):
            err = self._deserialize_error(error.details())
            if err is not None:
                return err
            err = self.error()
            err.code = error.code().value[0]
            err.errors = [error.details()]
            return err
        elif isinstance(error, Exception):
            if len(error.args) != 1:
                return None
            if isinstance(error.args[0], Error):
                return error.args[0]
            elif isinstance(error.args[0], str):
                return self._deserialize_error(error.args[0])

    def set_config(self, payload):
        """POST /api/config

        Sets configuration resources.

        Return: None
        """
        raise NotImplementedError("set_config")

    def update_configuration(self, payload):
        """PATCH /api/config

        Deprecated: please use post instead. Sets configuration resources.

        Return: prefix_config
        """
        raise NotImplementedError("update_configuration")

    def get_config(self):
        """GET /api/config

        Gets the configuration resources.

        Return: prefix_config
        """
        raise NotImplementedError("get_config")

    def get_metrics(self, payload):
        """GET /api/metrics

        Gets metrics.

        Return: metrics
        """
        raise NotImplementedError("get_metrics")

    def get_warnings(self):
        """GET /api/warnings

        Gets warnings.

        Return: warning_details
        """
        raise NotImplementedError("get_warnings")

    def clear_warnings(self):
        """DELETE /api/warnings

        Clears warnings.

        Return: None
        """
        raise NotImplementedError("clear_warnings")

    def getrootresponse(self):
        """GET /api/apitest

        simple GET api with single return type

        Return: common_responsesuccess
        """
        raise NotImplementedError("getrootresponse")

    def dummyresponsetest(self):
        """DELETE /api/apitest

        TBD

        Return: None
        """
        raise NotImplementedError("dummyresponsetest")

    def postrootresponse(self, payload):
        """POST /api/apitest

        simple POST api with single return type

        Return: common_responsesuccess
        """
        raise NotImplementedError("postrootresponse")

    def getallitems(self):
        """GET /api/serviceb

        return list of some items

        Return: serviceabc_itemlist
        """
        raise NotImplementedError("getallitems")

    def getsingleitem(self):
        """GET /api/serviceb/{item_id}

        return single item

        Return: serviceabc_item
        """
        raise NotImplementedError("getsingleitem")

    def getsingleitemlevel2(self):
        """GET /api/serviceb/{item_id}/{level_2}

        return single item

        Return: serviceabc_item
        """
        raise NotImplementedError("getsingleitemlevel2")

    def get_version(self):
        """GET /api/capabilities/version

        TBD

        Return: version
        """
        raise NotImplementedError("get_version")

    def prefix_config(self):
        """Factory method that creates an instance of PrefixConfig

        Return: PrefixConfig
        """
        return PrefixConfig()

    def error(self):
        """Factory method that creates an instance of Error

        Return: Error
        """
        return Error()

    def update_config(self):
        """Factory method that creates an instance of UpdateConfig

        Return: UpdateConfig
        """
        return UpdateConfig()

    def metrics_request(self):
        """Factory method that creates an instance of MetricsRequest

        Return: MetricsRequest
        """
        return MetricsRequest()

    def metrics(self):
        """Factory method that creates an instance of Metrics

        Return: Metrics
        """
        return Metrics()

    def warning_details(self):
        """Factory method that creates an instance of WarningDetails

        Return: WarningDetails
        """
        return WarningDetails()

    def common_responsesuccess(self):
        """Factory method that creates an instance of CommonResponseSuccess

        Return: CommonResponseSuccess
        """
        return CommonResponseSuccess()

    def apitest_inputbody(self):
        """Factory method that creates an instance of ApiTestInputBody

        Return: ApiTestInputBody
        """
        return ApiTestInputBody()

    def serviceabc_itemlist(self):
        """Factory method that creates an instance of ServiceAbcItemList

        Return: ServiceAbcItemList
        """
        return ServiceAbcItemList()

    def serviceabc_item(self):
        """Factory method that creates an instance of ServiceAbcItem

        Return: ServiceAbcItem
        """
        return ServiceAbcItem()

    def version(self):
        """Factory method that creates an instance of Version

        Return: Version
        """
        return Version()

    def close(self):
        pass

    def _check_client_server_version_compatibility(
        self, client_ver, server_ver, component_name
    ):
        try:
            c = semantic_version.Version(client_ver)
        except Exception as e:
            raise AssertionError(
                "Client {} version '{}' is not a valid semver: {}".format(
                    component_name, client_ver, e
                )
            )

        try:
            s = semantic_version.SimpleSpec(server_ver)
        except Exception as e:
            raise AssertionError(
                "Server {} version '{}' is not a valid semver: {}".format(
                    component_name, server_ver, e
                )
            )

        err = "Client {} version '{}' is not semver compatible with Server {} version '{}'".format(
            component_name, client_ver, component_name, server_ver
        )

        if not s.match(c):
            raise Exception(err)

    def get_local_version(self):
        return self._version_meta

    def get_remote_version(self):
        return self.get_version()

    def check_version_compatibility(self):
        comp_err, api_err = self._do_version_check()
        if comp_err is not None:
            raise comp_err
        if api_err is not None:
            raise api_err

    def _do_version_check(self):
        local = self.get_local_version()
        try:
            remote = self.get_remote_version()
        except Exception as e:
            return None, e

        try:
            self._check_client_server_version_compatibility(
                local.api_spec_version, remote.api_spec_version, "API spec"
            )
        except Exception as e:
            msg = "client SDK version '{}' is not compatible with server SDK version '{}'".format(
                local.sdk_version, remote.sdk_version
            )
            return Exception("{}: {}".format(msg, str(e))), None

        return None, None

    def _do_version_check_once(self):
        if not self._version_check:
            return

        if self._version_check_err is not None:
            raise self._version_check_err

        comp_err, api_err = self._do_version_check()
        if comp_err is not None:
            self._version_check_err = comp_err
            raise comp_err
        if api_err is not None:
            self._version_check_err = None
            raise api_err

        self._version_check = False
        self._version_check_err = None


class HttpApi(Api):
    """OpenAPI HTTP Api"""

    def __init__(self, **kwargs):
        super(HttpApi, self).__init__(**kwargs)
        self._transport = HttpTransport(**kwargs)

    @property
    def verify(self):
        return self._transport.verify

    @verify.setter
    def verify(self, value):
        self._transport.set_verify(value)

    def set_config(self, payload):
        """POST /api/config

        Sets configuration resources.

        Return: None
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "post",
            "/api/config",
            payload=payload,
            return_object=None,
            request_class=PrefixConfig,
        )

    def update_configuration(self, payload):
        """PATCH /api/config

        Deprecated: please use post instead. Sets configuration resources.

        Return: prefix_config
        """
        self.add_warnings(
            "update_configuration api is deprecated, please use post instead"
        )
        self._do_version_check_once()
        return self._transport.send_recv(
            "patch",
            "/api/config",
            payload=payload,
            return_object=self.prefix_config(),
            request_class=UpdateConfig,
        )

    def get_config(self):
        """GET /api/config

        Gets the configuration resources.

        Return: prefix_config
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/config",
            payload=None,
            return_object=self.prefix_config(),
        )

    def get_metrics(self, payload):
        """GET /api/metrics

        Gets metrics.

        Return: metrics
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/metrics",
            payload=payload,
            return_object=self.metrics(),
            request_class=MetricsRequest,
        )

    def get_warnings(self):
        """GET /api/warnings

        Gets warnings.

        Return: warning_details
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/warnings",
            payload=None,
            return_object=self.warning_details(),
        )

    def clear_warnings(self):
        """DELETE /api/warnings

        Clears warnings.

        Return: None
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "delete",
            "/api/warnings",
            payload=None,
            return_object=None,
        )

    def getrootresponse(self):
        """GET /api/apitest

        simple GET api with single return type

        Return: common_responsesuccess
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/apitest",
            payload=None,
            return_object=self.common_responsesuccess(),
        )

    def dummyresponsetest(self):
        """DELETE /api/apitest

        TBD

        Return: None
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "delete",
            "/api/apitest",
            payload=None,
            return_object=None,
        )

    def postrootresponse(self, payload):
        """POST /api/apitest

        simple POST api with single return type

        Return: common_responsesuccess
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "post",
            "/api/apitest",
            payload=payload,
            return_object=self.common_responsesuccess(),
            request_class=ApiTestInputBody,
        )

    def getallitems(self):
        """GET /api/serviceb

        return list of some items

        Return: serviceabc_itemlist
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/serviceb",
            payload=None,
            return_object=self.serviceabc_itemlist(),
        )

    def getsingleitem(self):
        """GET /api/serviceb/{item_id}

        return single item

        Return: serviceabc_item
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/serviceb/{item_id}",
            payload=None,
            return_object=self.serviceabc_item(),
        )

    def getsingleitemlevel2(self):
        """GET /api/serviceb/{item_id}/{level_2}

        return single item

        Return: serviceabc_item
        """
        self._do_version_check_once()
        return self._transport.send_recv(
            "get",
            "/api/serviceb/{item_id}/{level_2}",
            payload=None,
            return_object=self.serviceabc_item(),
        )

    def get_version(self):
        """GET /api/capabilities/version

        TBD

        Return: version
        """
        return self._transport.send_recv(
            "get",
            "/api/capabilities/version",
            payload=None,
            return_object=self.version(),
        )


class GrpcApi(Api):
    # OpenAPI gRPC Api
    def __init__(self, **kwargs):
        super(GrpcApi, self).__init__(**kwargs)
        self._stub = None
        self._channel = None
        self._cert = None
        self._cert_domain = None
        self._request_timeout = 10
        self._keep_alive_timeout = 10 * 1000
        self._location = (
            kwargs["location"]
            if "location" in kwargs and kwargs["location"] is not None
            else "localhost:50051"
        )
        self._transport = (
            kwargs["transport"] if "transport" in kwargs else None
        )
        self._logger = kwargs["logger"] if "logger" in kwargs else None
        self._loglevel = (
            kwargs["loglevel"] if "loglevel" in kwargs else logging.DEBUG
        )
        if self._logger is None:
            stdout_handler = logging.StreamHandler(sys.stdout)
            formatter = logging.Formatter(
                fmt="%(asctime)s [%(name)s] [%(levelname)s] %(message)s",
                datefmt="%Y-%m-%d %H:%M:%S",
            )
            formatter.converter = time.gmtime
            stdout_handler.setFormatter(formatter)
            self._logger = logging.Logger(
                self.__module__, level=self._loglevel
            )
            self._logger.addHandler(stdout_handler)
        self._logger.debug(
            "gRPCTransport args: {}".format(
                ", ".join(["{}={!r}".format(k, v) for k, v in kwargs.items()])
            )
        )

    def _use_secure_connection(self, cert_path, cert_domain=None):
        """Accepts certificate and host_name for SSL Connection."""
        if cert_path is None:
            raise Exception("path to certificate cannot be None")
        self._cert = cert_path
        self._cert_domain = cert_domain

    def _get_stub(self):
        if self._stub is None:
            CHANNEL_OPTIONS = [
                ("grpc.enable_retries", 0),
                ("grpc.keepalive_timeout_ms", self._keep_alive_timeout),
            ]
            if self._cert is None:
                self._channel = grpc.insecure_channel(
                    self._location, options=CHANNEL_OPTIONS
                )
            else:
                crt = open(self._cert, "rb").read()
                creds = grpc.ssl_channel_credentials(crt)
                if self._cert_domain is not None:
                    CHANNEL_OPTIONS.append(
                        ("grpc.ssl_target_name_override", self._cert_domain)
                    )
                self._channel = grpc.secure_channel(
                    self._location, credentials=creds, options=CHANNEL_OPTIONS
                )
            self._stub = pb2_grpc.OpenapiStub(self._channel)
        return self._stub

    def _serialize_payload(self, payload):
        if not isinstance(payload, (str, dict, OpenApiBase)):
            raise Exception(
                "We are supporting [str, dict, OpenApiBase] object"
            )
        if isinstance(payload, OpenApiBase):
            payload = payload.serialize()
        if isinstance(payload, dict):
            payload = json.dumps(payload)
        elif isinstance(payload, (str, unicode)):
            payload = json.dumps(yaml.safe_load(payload))
        return payload

    def _raise_exception(self, grpc_error):
        err = self.error()
        try:
            err.deserialize(grpc_error.details())
        except Exception as _:
            err.code = grpc_error.code().value[0]
            err.errors = [grpc_error.details()]
        raise Exception(err)

    @property
    def request_timeout(self):
        """duration of time in seconds to allow for the RPC."""
        return self._request_timeout

    @request_timeout.setter
    def request_timeout(self, timeout):
        self._request_timeout = timeout

    @property
    def keep_alive_timeout(self):
        return self._keep_alive_timeout

    @keep_alive_timeout.setter
    def keep_alive_timeout(self, timeout):
        self._keep_alive_timeout = timeout * 1000

    def close(self):
        if self._channel is not None:
            self._channel.close()
            self._channel = None
            self._stub = None

    def set_config(self, payload):
        pb_obj = json_format.Parse(
            self._serialize_payload(payload), pb2.PrefixConfig()
        )
        self._do_version_check_once()
        req_obj = pb2.SetConfigRequest(prefix_config=pb_obj)
        stub = self._get_stub()
        try:
            res_obj = stub.SetConfig(req_obj, timeout=self._request_timeout)
        except grpc.RpcError as grpc_error:
            self._raise_exception(grpc_error)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        bytes = response.get("response_bytes")
        if bytes is not None:
            return io.BytesIO(res_obj.response_bytes)

    def update_configuration(self, payload):
        self.add_warnings(
            "update_configuration api is deprecated, please use post instead"
        )
        pb_obj = json_format.Parse(
            self._serialize_payload(payload), pb2.UpdateConfig()
        )
        self._do_version_check_once()
        req_obj = pb2.UpdateConfigurationRequest(update_config=pb_obj)
        stub = self._get_stub()
        try:
            res_obj = stub.UpdateConfiguration(
                req_obj, timeout=self._request_timeout
            )
        except grpc.RpcError as grpc_error:
            self._raise_exception(grpc_error)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("prefix_config")
        if result is not None:
            return self.prefix_config().deserialize(result)

    def get_config(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetConfig(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("prefix_config")
        if result is not None:
            return self.prefix_config().deserialize(result)

    def get_metrics(self, payload):
        pb_obj = json_format.Parse(
            self._serialize_payload(payload), pb2.MetricsRequest()
        )
        self._do_version_check_once()
        req_obj = pb2.GetMetricsRequest(metrics_request=pb_obj)
        stub = self._get_stub()
        try:
            res_obj = stub.GetMetrics(req_obj, timeout=self._request_timeout)
        except grpc.RpcError as grpc_error:
            self._raise_exception(grpc_error)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("metrics")
        if result is not None:
            return self.metrics().deserialize(result)

    def get_warnings(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetWarnings(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("warning_details")
        if result is not None:
            return self.warning_details().deserialize(result)

    def clear_warnings(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.ClearWarnings(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        resp_str = response.get("string")
        if resp_str is not None:
            return response.get("string")

    def getrootresponse(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetRootResponse(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("common_response_success")
        if result is not None:
            return self.common_responsesuccess().deserialize(result)

    def dummyresponsetest(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.DummyResponseTest(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        resp_str = response.get("string")
        if resp_str is not None:
            return response.get("string")

    def postrootresponse(self, payload):
        pb_obj = json_format.Parse(
            self._serialize_payload(payload), pb2.ApiTestInputBody()
        )
        self._do_version_check_once()
        req_obj = pb2.PostRootResponseRequest(apitest_inputbody=pb_obj)
        stub = self._get_stub()
        try:
            res_obj = stub.PostRootResponse(
                req_obj, timeout=self._request_timeout
            )
        except grpc.RpcError as grpc_error:
            self._raise_exception(grpc_error)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("common_response_success")
        if result is not None:
            return self.common_responsesuccess().deserialize(result)

    def getallitems(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetAllItems(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("service_abc_item_list")
        if result is not None:
            return self.serviceabc_itemlist().deserialize(result)

    def getsingleitem(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetSingleItem(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("service_abc_item")
        if result is not None:
            return self.serviceabc_item().deserialize(result)

    def getsingleitemlevel2(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetSingleItemLevel2(
            empty, timeout=self._request_timeout
        )
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("service_abc_item")
        if result is not None:
            return self.serviceabc_item().deserialize(result)

    def get_version(self):
        stub = self._get_stub()
        empty = pb2_grpc.google_dot_protobuf_dot_empty__pb2.Empty()
        res_obj = stub.GetVersion(empty, timeout=self._request_timeout)
        response = json_format.MessageToDict(
            res_obj, preserving_proto_field_name=True
        )
        result = response.get("version")
        if result is not None:
            return self.version().deserialize(result)
