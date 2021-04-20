import sys
import os
import importlib


class OpenApiArt(object):
    """Bundle and generate artifacts from OpenAPI files.

    Args
    ----
    - api_files (list[str]): list of OpenAPI files that contain info and/or path
      keywords
    - python_module_name (str): name of the consolidated python file that will be generated
    - output_dir (str): directory where artifacts will be created. 
      Unless otherwise specified the default directory for generated artifacts 
      is `current working directory/.output`.
      The artifacts that will be generated are:
      - openapi.yaml
      - openapi.json
      - static openapi.html documentation (if redoc-cli has been installed)
      - python file
    """
    def __init__(self, api_files, python_module_name='sample', output_dir=None):
        self._python_module_name = python_module_name
        if output_dir is None:
            output_dir = os.path.join(os.getcwd(), '.output')
        self._output_dir = os.path.abspath(output_dir)
        self._api_files = api_files
        module = importlib.import_module('openapiart.bundler')
        bundler = getattr(module, 'Bundler')(api_files=api_files,
                                             output_dir=self._output_dir)
        bundler.bundle()
        module = importlib.import_module('openapiart.generator')
        generator = getattr(module, 'Generator')(bundler.openapi_filepath,
                                                 self._python_module_name,
                                                 output_dir=self._output_dir)
        generator.generate()

    @property
    def output_dir(self):
        return self._output_dir

    @property
    def python_module_name(self):
        return self._python_module_name
