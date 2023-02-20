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
        self._python_module_name = (
            None
            if "python_module_name" not in kwargs
            else kwargs["python_module_name"]
        )
        self._protobuf_package_name = kwargs["protobuf_package_name"]
        self._protobuf_file_name = kwargs["protobuf_package_name"]
        self._go_sdk_package_dir = kwargs["go_sdk_package_dir"]
        self._generate_version_api = kwargs.get("generate_version_api")
        if self._generate_version_api is None:
            self._generate_version_api = False
        self._api_version = kwargs.get("api_version")
        if self._api_version is None:
            self._api_version = ""
        self._sdk_version = kwargs.get("sdk_version")
        if self._sdk_version is None:
            self._sdk_version = ""
        self._go_sdk_package_name = (
            None
            if "go_sdk_package_name" not in kwargs
            else kwargs["go_sdk_package_name"]
        )
        self.default_indent = "    "
        self._parsers = {}

    def _init_fp(self, filename):
        self._filename = filename
        self._fp = open(self._filename, "wb")

    def _close_fp(self):
        self._fp.close()

    def _write(self, line="", indent=0, newline=True):
        line = "{}{}{}".format(
            self.default_indent * indent, line, "\n" if newline else ""
        )
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

    def _justify_desc(self, text, indent=0, use_multi=False):
        indent = " " * (indent * 2)
        lines = []
        text = text.split("\n")
        for line in text:
            char_80 = ""
            for word in line.split(" "):
                if len(char_80) <= 80:
                    char_80 += word + " "
                    continue
                lines.append(char_80.strip())
                char_80 = word + " "
            if char_80 != "":
                lines.append(char_80.strip())
            # lines.append("\n{}{}".format(indent, comment).join(each_line))
        if use_multi is True:
            return (
                "{}/* ".format(indent)
                + "\n{} * ".format(indent).join(lines)
                + " */"
            )
        return "{}// ".format(indent) + "\n{}// ".format(indent).join(lines)
