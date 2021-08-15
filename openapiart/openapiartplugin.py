"""Abstract plugin class
"""
import jsonpath_ng


class OpenApiArtPlugin(object):
    """Abstract class for creating a plugin generator"""

    def __init__(self, **kwargs):
        self._fp = None
        self._license = kwargs["license"]
        self._info = kwargs["info"]
        self._output_dir = kwargs["output_dir"]
        self._python_module_name = None if "python_module_name" not in kwargs else kwargs["python_module_name"]
        self._protobuf_package_name = kwargs["protobuf_package_name"]
        self._protobuf_file_name = kwargs["protobuf_package_name"]
        self._go_sdk_package_dir = kwargs["go_sdk_package_dir"]
        self._go_sdk_package_name = None if "go_sdk_package_name" not in kwargs else kwargs["go_sdk_package_name"]
        self.default_indent = "    "
        self._parsers = {}

    def _init_fp(self, filename):
        self._filename = filename
        self._fp = open(self._filename, "wb")

    def _close_fp(self):
        self._fp.close()

    def _write(self, line="", indent=0, newline=True):
        line = "{}{}{}".format(self.default_indent * indent, line, "\n" if newline else "")
        self._fp.write(line.encode())

    def _get_parser(self, pattern):
        if pattern not in self._parsers:
            parser = jsonpath_ng.parse(pattern)
            self._parsers[pattern] = parser
        else:
            parser = self._parsers[pattern]
        return parser

    def _get_camel_case(self, value):
        camel_case = ""
        for piece in value.split("_"):
            camel_case += piece[0].upper()
            if len(piece) > 1:
                camel_case += piece[1:]
        return camel_case
