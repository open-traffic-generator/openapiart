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
import sanity_pb2_grpc as pb2_grpc
import sanity_pb2 as pb2

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
        lib = importlib.import_module("{}_{}".format(__name__, ext))
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

    def send_recv(
        self,
        method,
        relative_url,
        payload=None,
        return_object=None,
        headers=None,
    ):
        url = "%s%s" % (self.location, relative_url)
        data = None
        headers = headers or {"Content-Type": "application/json"}
        if payload is not None:
            if isinstance(payload, bytes):
                data = payload
                headers["Content-Type"] = "application/octet-stream"
            elif isinstance(payload, (str, unicode)):
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
            raise Exception(
                response.status_code, yaml.safe_load(response.text)
            )


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

    def validate_integer(self, value, min, max):
        if value is None or not isinstance(value, int):
            return False
        if value < 0:
            return False
        if min is not None and value < min:
            return False
        if max is not None and value > max:
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
            "double": "float",
        }
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
            verdict = v_obj(value, min, max)
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
        self._validate_required()
        for key, value in self._properties.items():
            self._validate_types(key, value)
            # TODO: restore behavior
            # self._validate_unique_and_name(key, value, True)
            # self._validate_constraint(key, value, True)
            if isinstance(value, (OpenApiObject, OpenApiIter)):
                output[key] = value._encode()
            elif value is not None:
                if self._TYPES.get(key, {}).get("format", "") == "int64":
                    value = str(value)
                elif self._TYPES.get(key, {}).get("itemformat", "") == "int64":
                    value = [str(v) for v in value]
                output[key] = value
                # TODO: restore behavior
                # OpenApiStatus.warn(
                #     "{}.{}".format(type(self).__name__, key), self
                # )
        return output

    def _decode(self, obj):
        dtypes = [list, str, int, float, bool]
        for property_name, property_value in obj.items():
            if property_name in self._TYPES:
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
                elif (
                    property_name in self._DEFAULTS and property_value is None
                ):
                    if isinstance(
                        self._DEFAULTS[property_name], tuple(dtypes)
                    ):
                        property_value = self._DEFAULTS[property_name]
                self._set_choice(property_name)
                # convert int64(will be string on wire) to to int
                if self._TYPES[property_name].get("format", "") == "int64":
                    property_value = int(property_value)
                elif (
                    self._TYPES[property_name].get("itemformat", "") == "int64"
                ):
                    property_value = [int(v) for v in property_value]
                self._properties[property_name] = property_value
                # TODO: restore behavior
                # OpenApiStatus.warn(
                #     "{}.{}".format(type(self).__name__, property_name), self
                # )
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
            self.types_validation(
                property_value,
                details["type"],
                msg,
                details.get("itemtype"),
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
