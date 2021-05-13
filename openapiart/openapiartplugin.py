"""Abstract plugin class
"""


class OpenApiArtPlugin(object):
    """Abstract class for creating a plugin generator
    """
    def __init__(self, **kwargs):
        self._fp = None
        self._output_dir = kwargs['output_dir']
        self._python_module_name = kwargs['python_module_name']
        self._protobuf_file_name = kwargs['protobuf_file_name']
        self.default_indent = '    '

    def _init_fp(self, filename):
        self._filename = filename
        self._fp = open(self._filename, 'wb')

    def _write(self, line='', indent=0, newline=True):
        line = '{}{}{}'.format(
            self.default_indent * indent, 
            line,
            '\n' if newline else '')
        self._fp.write(line.encode())

    def pre_process(self):
        pass

    def post_process(self):
        if self._fp is not None:
            close(self._fp)

    def info(self, info_object):
        pass

    def path(self, path_object):
        pass

    def component_schema(self, schema_object):
        pass

    def component_response(self, response_object):
        pass


