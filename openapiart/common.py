import importlib
import logging
import json
import yaml
import requests
import urllib3
import io
import sys
import time
import re

try:
    from typing import Union, Dict, List, Any, Literal
except ImportError:
    from typing_extensions import Literal

if sys.version_info[0] == 3:
    unicode = str


def api(location=None, verify=True, logger=None, loglevel=logging.INFO, ext=None):
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
    if ext is None:
        return HttpApi(**params)
    try:
        lib = importlib.import_module("{}_{}".format(__name__, ext))
        return lib.Api(**params)
    except ImportError as err:
        msg = "Extension %s is not installed or invalid: %s"
        raise Exception(msg % (ext, err))


class HttpTransport(object):
    def __init__(self, **kwargs):
        """Use args from api() method to instantiate an HTTP transport"""
        self.location = kwargs["location"] if "location" in kwargs else "https://localhost"
        self.verify = kwargs["verify"] if "verify" in kwargs else False
        self.logger = kwargs["logger"] if "logger" in kwargs else None
        self.loglevel = kwargs["loglevel"] if "loglevel" in kwargs else logging.DEBUG
        if self.logger is None:
            stdout_handler = logging.StreamHandler(sys.stdout)
            formatter = logging.Formatter(fmt="%(asctime)s [%(name)s] [%(levelname)s] %(message)s", datefmt="%Y-%m-%d %H:%M:%S")
            formatter.converter = time.gmtime
            stdout_handler.setFormatter(formatter)
            self.logger = logging.Logger(self.__module__, level=self.loglevel)
            self.logger.addHandler(stdout_handler)
        self.logger.debug("HttpTransport args: {}".format(", ".join(["{}={!r}".format(k, v) for k, v in kwargs.items()])))
        if self.verify is False:
            urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
            self.logger.warning("Certificate verification is disabled")
        self._session = requests.Session()

    def send_recv(self, method, relative_url, payload=None, return_object=None):
        url = "%s%s" % (self.location, relative_url)
        data = None
        if payload is not None:
            if isinstance(payload, (str, unicode)):
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
            headers={"Content-Type": "application/json"},
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
            elif "application/octet-stream" in response.headers["content-type"]:
                return io.BytesIO(response.content)
            else:
                # TODO: for now, return bare response object for unknown
                # content types
                return response
        else:
            raise Exception(response.status_code, yaml.safe_load(response.text))


class OpenApiBase(object):
    """Base class for all generated classes"""

    JSON = "json"
    YAML = "yaml"
    DICT = "dict"

    __slots__ = ()

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
        if encoding == OpenApiBase.JSON:
            return json.dumps(self._encode(), indent=2)
        elif encoding == OpenApiBase.YAML:
            return yaml.safe_dump(self._encode())
        elif encoding == OpenApiBase.DICT:
            return self._encode()
        else:
            raise NotImplementedError("Encoding %s not supported" % encoding)

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
        if isinstance(serialized_object, (str, unicode)):
            serialized_object = yaml.safe_load(serialized_object)
        self._decode(serialized_object)
        return self

    def _decode(self, dict_object):
        raise NotImplementedError()


class OpenApiValidator(object):

    _MAC_REGEX = re.compile(
        r"^([\da-fA-F]{2}[:]){5}[\da-fA-F]{2}$")
    _IPV6_REP1 = re.compile(r"^:[\da-fA-F].+")
    _IPV6_REP2 = re.compile(r".+[\da-fA-F]:$")
    _IPV6_REP3 = re.compile(r"^" +
        r"[\da-fA-F]{1,4}:" *7 + r"[\da-fA-F]{1,4}$")
    _HEX_REGEX = re.compile(r"^0?x?[\da-fA-F]+$")

    __slots__ = ()

    def __init__(self):
        pass

    def validate_mac(self, mac):
        if mac is None:
            return False
        if isinstance(mac, list):
            return all([
                True if self._MAC_REGEX.match(m) else False
                for m in mac
            ])
        if self._MAC_REGEX.match(mac):
            return True
        return False

    def validate_ipv4(self, ip):
        if ip is None:
            return False
        try:
            if isinstance(ip, list):
                return all([
                    all([0 <= int(oct) <= 255 for oct in i.split(".", 3)])
                    for i in ip
                ])
            else:
                return all([0 <= int(oct) <= 255 for oct in ip.split(".", 3)])
        except Exception:
            return False

    def _validate_ipv6(self, ip):
        if ip is None:
            return False
        if self._IPV6_REP1.match(ip) or self._IPV6_REP2.match(ip):
            return False
        if ip.count("::") == 0:
            if self._IPV6_REP3.match(ip):
                return True
            else:
                return False
        if ip.count(":") > 7 or ip.count("::") > 1 or ip.count(":::") > 0:
            return False
        return True

    def validate_ipv6(self, ip):
        if isinstance(ip, list):
            return all([
                self._validate_ipv6(i) for i in ip
            ])
        return self._validate_ipv6(ip)

    def validate_hex(self, hex):
        if isinstance(hex, list):
            return all([
                self._HEX_REGEX(h)
                for h in hex
            ])
        if self._HEX_REGEX.match(hex):
            return True
        return False

    def validate_integer(self, value):
        if not isinstance(value, list):
            value = [value]
        return all([isinstance(i, int) for i in value])

    def validate_float(self, value):
        if not isinstance(value, list):
            value = [value]
        return all([
            isinstance(i, (float, int)) for i in value
        ])

    def validate_string(self, value):
        if not isinstance(value, list):
            value = [value]
        return all([
            isinstance(i, (str, unicode)) for i in value
        ])
    
    def validate_bool(self, value):
        if not isinstance(value, list):
            value = [value]
        return all([
            isinstance(i, bool) for i in value
        ])

    def validate_list(self, value):
        # TODO pending item validation
        return isinstance(value, list)
    
    def validate_binary(self, value):
        if not isinstance(value, list):
            value = [value]
        return all([
            all([
                True if int(b) == 0 or int(b) == 1 else False
                for b in binary
            ])
            for binary in value
        ])
    
    def types_validation(self, value, type_, err_msg):
        type_map = {
            int: "integer", str: "string",
            float: "float", bool: "bool",
            list: "list", "int64": "integer",
            "int32": "integer", "double": "float"
        }
        if type_ in type_map:
            type_ = type_map[type_]
        v_obj = getattr(self, "validate_{}".format(type_))
        if v_obj is None:
            msg = "{} is not a valid or unsupported format".format(type_)
            raise TypeError(msg)
        if v_obj(value) is False:
            raise TypeError(err_msg)


class OpenApiObject(OpenApiBase, OpenApiValidator):
    """Base class for any /components/schemas object

    Every OpenApiObject is reuseable within the schema so it can
    exist in multiple locations within the hierarchy.
    That means it can exist in multiple locations as a
    leaf, parent/choice or parent.
    """

    __slots__ = ("_properties", "_parent", "_choice")

    def __init__(self, parent=None, choice=None):
        super(OpenApiObject, self).__init__()
        self._parent = parent
        self._choice = choice
        self._properties = {}

    @property
    def parent(self):
        return self._parent

    def _set_choice(self, name):
        if "choice" in dir(self) and "_TYPES" in dir(self) \
            and "choice" in self._TYPES and name in self._TYPES["choice"]["enum"]:
            for enum in self._TYPES["choice"]["enum"]:
                if enum in self._properties and name != enum:
                    self._properties.pop(enum)
            self._properties["choice"] = name
        
    def _get_property(self, name, default_value=None, parent=None, choice=None):
        if name in self._properties and self._properties[name] is not None:
            return self._properties[name]
        if isinstance(default_value, type) is True:
            self._set_choice(name)
            self._properties[name] = default_value(parent=parent, choice=choice)

            if "_DEFAULTS" in dir(self._properties[name]) and\
                "choice" in self._properties[name]._DEFAULTS:
                getattr(self._properties[name], self._properties[name]._DEFAULTS["choice"])
        else:
            if default_value is None and name in self._DEFAULTS:
                self._set_choice(name)
                self._properties[name] = self._DEFAULTS[name]
            else:
                self._properties[name] = default_value
        return self._properties[name]

    def _set_property(self, name, value, choice=None):
        if name in self._DEFAULTS and value is None:
            self._set_choice(name)
            self._properties[name] = self._DEFAULTS[name]
        else:
            self._set_choice(name)
            self._properties[name] = value
        if self._parent is not None and self._choice is not None and value is not None:
            self._parent._set_property("choice", self._choice)

    def _encode(self):
        """Helper method for serialization"""
        output = {}
        self._validate_required()
        for key, value in self._properties.items():
            self._validate_types(key, value)
            if isinstance(value, (OpenApiObject, OpenApiIter)):
                output[key] = value._encode()
            elif value is not None:
                output[key] = value
        return output

    def _decode(self, obj):
        openapi_names = dir(self)
        dtypes = [list, str, int, float, bool]
        for property_name, property_value in obj.items():
            if property_name in openapi_names:
                if isinstance(property_value, dict):
                    child = self._get_child_class(property_name)
                    if "_choice" in dir(child[1]) and "_parent" in dir(child[1]):
                        property_value = child[1](self, property_name)._decode(property_value)
                    else:
                        property_value = child[1]()._decode(property_value)
                elif isinstance(property_value, list) and \
                    property_name in self._TYPES and \
                        self._TYPES[property_name]["type"] not in dtypes:
                    child = self._get_child_class(property_name, True)
                    openapi_list = child[0]()
                    for item in property_value:
                        item = child[1]()._decode(item)
                        openapi_list._items.append(item)
                    property_value = openapi_list
                elif property_name in self._DEFAULTS and property_value is None:
                    if isinstance(self._DEFAULTS[property_name], tuple(dtypes)):
                        property_value = self._DEFAULTS[property_name]
                self._properties[property_name] = property_value
            self._validate_types(property_name, property_value)
        self._validate_required()
        return self

    def _get_child_class(self, property_name, is_property_list=False):
        list_class = None
        class_name = self._TYPES[property_name]["type"]
        module = importlib.import_module(self.__module__)
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
        """
        if getattr(self, "_REQUIRED", None) is None:
            return
        for p in self._REQUIRED:
            if p in self._properties and self._properties[p] is not None:
                continue
            msg = "{} is a mandatory property of {}"\
                " and should not be set to None".format(
                p, self.__class__
            )
            raise ValueError(msg)

    def _validate_types(self, property_name, property_value):
        common_data_types = [list, str, int, float, bool]
        if property_name not in self._TYPES:
            raise ValueError("Invalid Property {}".format(property_name))
        details = self._TYPES[property_name]
        if property_value is None and property_name not in self._DEFAULTS and \
            property_name not in self._REQUIRED:
            return
        if "enum" in details and property_value not in details["enum"]:
            msg = "property {} shall be one of these" \
                " {} enum, but got {} at {}"
            raise TypeError(msg.format(
                property_name, details["enum"], property_value, self.__class__
            ))
        if details["type"] in common_data_types and \
                "format" not in details:
            msg = "property {} shall be of type {}, but got {} at {}".format(
                    property_name, details["type"], type(property_value), self.__class__
                )
            self.types_validation(property_value, details["type"], msg)

        if details["type"] not in common_data_types:
            class_name = details["type"]
            # TODO Need to revisit importlib
            module = importlib.import_module(self.__module__)
            object_class = getattr(module, class_name)
            if not isinstance(property_value, object_class):
                msg = "property {} shall be of type {}," \
                    " but got {} at {}"
                raise TypeError(msg.format(
                    property_name, class_name, type(property_value),
                    self.__class__
                ))
        if "format" in details:
            msg = "Invalid {} format, expected {} at {}".format(
                property_value, details["format"], self.__class__
            )
            self.types_validation(property_value, details["format"], msg)

    def validate(self):
        self._validate_required()
        for key, value in self._properties.items():
            self._validate_types(key, value)
    
    def get(self, name, with_default=False):
        """ 
        getattr for openapi object
        """
        if self._properties.get(name) is not None:
            return self._properties[name]
        elif with_default:
            # TODO need to find a way to avoid getattr
            choice = self._properties.get("choice")\
                    if "choice" in dir(self) else None
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
        if self._GETITEM_RETURNS_CHOICE_OBJECT is True and found._properties.get("choice") is not None:
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

    def _add(self, item):
        self._items.append(item)
        self._index = len(self._items) - 1

    def append(self, item):
        """Append an item to the end of OpenApiIter
        TBD: type check, raise error on mismatch
        """
        if isinstance(item, OpenApiObject) is False:
            raise Exception("Item is not an instance of OpenApiObject")
        self._add(item)
        return self

    def clear(self):
        del self._items[:]
        self._index = -1

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
        raise NotImplementedError("Shallow copy of OpenApiIter objects is not supported")

    def __deepcopy__(self, memo):
        raise NotImplementedError("Deep copy of OpenApiIter objects is not supported")

    def __str__(self):
        return yaml.safe_dump(self._encode())

    def __eq__(self, other):
        return self.__str__() == other.__str__()
