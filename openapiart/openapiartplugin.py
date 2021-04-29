"""Abstract plugin class
"""


class OpenApiArtPlugin(object):
    """Abstract class for creating a plugin generator
    """

    def pre_init(self):
        pass

    def post_init(self):
        pass

    def info(self, info_object):
        pass

    def path(self, path_object):
        pass

    def component_schema(self, schema_object):
        pass

    def component_response(self, response_object):
        pass


