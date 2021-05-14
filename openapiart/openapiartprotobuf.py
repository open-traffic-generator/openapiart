import os
import jsonpath_ng
from .openapiartplugin import OpenApiArtPlugin


class OpenApiArtProtobuf(OpenApiArtPlugin):
    def __init__(self, **kwargs):
        super(OpenApiArtProtobuf, self).__init__(**kwargs)
        self._filename = os.path.normpath(os.path.join(
            self._output_dir, 
            '{}.proto'.format(self._protobuf_file_name)
        ))
        self.default_indent = '  '
        self._custom_id = 50000
        self._init_fp(self._filename)

    def generate(self, openapi):
        self._openapi = openapi
        self._write_header(self._openapi['info'])
        for name, schema_object in self._openapi['components']['schemas'].items():
            self._write_msg(name, schema_object)
        for name, schema_object in self._openapi['components']['responses'].items():
            self._write_msg(name, schema_object)
        self._write_service()

    def _next_custom_id(self):
        self._custom_id += 1
        return self._custom_id

    def _write_header(self, info_object):
        license_path = os.path.join(os.path.dirname(__file__), '..', 'LICENSE')
        with open(license_path) as fp:
            for line in fp.readlines():
                self._write('// {}'.format(line), newline=False)
        self._write()
        self._write('syntax = "proto3";')
        self._write()
        self._write('package {};'.format(self._protobuf_file_name))
        self._write()
        self._write('import "google/protobuf/descriptor.proto";')
        self._write('import "google/protobuf/empty.proto";')
        self._write()
        self._write('message OpenApiMsgOpt {')
        self._write('string description = 10;', indent=1)
        self._write('}')
        self._write('extend google.protobuf.MessageOptions {')
        self._write('optional OpenApiMsgOpt msg_meta = 50000;', indent=1)
        self._write('}')
        self._write()
        self._write('message OpenApiFldOpt {')
        self._write('oneof type {', indent=1)
        self._write('bool object = 1;', indent=2)
        self._write('bool string = 2;', indent=2)
        self._write('bool bool = 3;', indent=2)
        self._write('bool enum = 4;', indent=2)
        self._write('bool float = 5;', indent=2)
        self._write('bool double = 6;', indent=2)
        self._write('bool int32 = 7;', indent=2)
        self._write('bool bytes = 18;', indent=2)
        self._write('bool none = 19;', indent=2)
        self._write('}', indent=1)
        self._write('string default = 20;', indent=1)
        self._write('bool required = 30;', indent=1)
        self._write('string description = 40;', indent=1)
        self._write('}')
        self._write('extend google.protobuf.FieldOptions {')
        self._write('optional OpenApiFldOpt fld_meta = 50001;', indent=1)
        self._write('}')

    def _get_field_type(self, property_name, openapi_object):
        """Convert openapi type -> protobuf type

        - type:number -> float
        - type:number [format: float] -> float
        - type:number [format: double] -> double
        - type:integer -> int32
        - type:integer [format:int32] -> int32
        - type:integer [format:int64] -> int64
        - type:boolean -> bool
        - type:string -> string
        - type:string [format:binary] -> bytes
        """
        if 'type' in openapi_object:
            type = openapi_object['type']
            if type == 'boolean': return 'bool'
            if type == 'string': 
                if 'format' in openapi_object:
                    if openapi_object['format'] == 'binary':
                        return 'bytes'
                elif 'enum' in openapi_object:
                    enum_msg = self._camelcase('{}.Enum'.format(property_name))
                    self._write_enum_msg(enum_msg, openapi_object['enum'])
                    return enum_msg
                return 'string'
            if type == 'integer': return 'int32'
            if type == 'number': return 'double'
            if type == 'array': 
                return 'repeated ' + self._get_field_type(property_name, openapi_object['items'])
        elif '$ref' in openapi_object:
            return openapi_object['$ref'].split('/')[-1].replace('.', '')

    def _camelcase(self, value):
        return '{}{}'.format(value[0].upper(), value[1:])

    def _write_description(self, openapi_object, indent=0):
        if 'description' in openapi_object:
            description = []
            for line in openapi_object['description'].split('\n'):
                line = '// {}'.format(line.strip(' '))
                self._write(line, indent=indent)
        else:
            self._write('// Description missing in models', indent=indent)

    def _get_description(self, openapi_object):
        if 'description' in openapi_object:
            return openapi_object['description'].replace('\n', '\\n')
        else:
            return 'Description missing in models'

    def _write_enum_msg(self, enum_msg_name, enums):
        enum_msg = enum_msg_name.split('.')
        self._write('message {} {{'.format(enum_msg[0]), indent=1)
        self._write('enum Enum {', indent=2)
        id = 0
        for enum in enums:
            self._write('{} = {};'.format(enum, id), indent=3)
            id += 1
        self._write('}', indent=2)
        self._write('}', indent=1)

    def _write_msg(self, name, schema_object):
        msg_name = name.replace('.', '')
        print('writing msg {}'.format(msg_name))
        self._write()
        self._write_description(schema_object, indent=0)
        self._write('message {} {{'.format(msg_name), indent=0)
        self._write('option (msg_meta).description = "{}";'.format(self._get_description(schema_object)), indent=1)
        self._write_msg_fields(name, schema_object)
        self._write('}')

    def _write_msg_fields(self, name, schema_object):
        if 'properties' not in schema_object:
            return
        id = 0
        for property_name, property_object in schema_object['properties'].items():
            id += 1
            if id > 1: self._write()
            property_type = self._get_field_type(property_name, property_object)
            default = ''
            if 'default' in property_object:
                default = property_object['default']
            type = property_type.split(' ')[-1]
            if hasattr(self, '_refparse') is False:
                self._refparse = jsonpath_ng.parse("$..'$ref'")
            if len(self._refparse.find(property_object)) > 0:
                type = 'object'
                default = property_type.split(' ')[-1]
            elif property_type.endswith('.Enum'):
                type = 'enum'
                default = '{}.{}'.format(property_type, default)
            required = ('required' in schema_object and property_name in schema_object['required'])
            self._write('{} {} = {} ['.format(property_type, property_name.lower(), id), indent=1)
            self._write('(fld_meta).{} = true,'.format(type), indent=2)
            self._write('(fld_meta).default = "{}",'.format(default), indent=2)
            self._write('(fld_meta).required = {},'.format(str(required).lower()), indent=2)
            self._write('(fld_meta).description = "{}"'.format(self._get_description(property_object)), indent=2)
            self._write('];', indent=1)

    def _write_service(self):
        self._write()
        self._write('service Openapi {')
        for url, path_object in self._openapi['paths'].items():
            for method, path_item_object in path_object.items():
                if method in ['get', 'patch', 'post', 'delete']:
                    self._write_rpc(url, method, path_item_object)
        self._write('}')

    def _write_rpc(self, url, method, path_item_object):
        self._write_description(path_item_object, indent=1)
        operationId = ''
        for piece in path_item_object['operationId'].split('_'):
            operationId += piece[0].upper()
            if len(piece) > 1: operationId += piece[1:]
        params = []
        for response in path_item_object['responses']:
            # write a components/schemas/SetConfigResponse
            # contains choice enum[status_200, status_400, status_500]
            # status_220 has a reference to payload
            # httptransport.send_recv should return status_2xx or 
            # throw an error on status_4xx or status_5xx
            pass
        if len(params) == 0:
            params.append('google.protobuf.Empty')
        returns = None
        if returns is None:
            returns = 'google.protobuf.Empty'
        print('writing rpc {}'.format(operationId))
        self._write('rpc {}({}) returns({}) {{'.format(operationId, ', '.join(params), returns), indent=1)
        self._write('}', indent=1)



        