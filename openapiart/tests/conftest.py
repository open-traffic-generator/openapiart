import pytest
import sys
import os
import importlib
import logging


@pytest.fixture(scope='session')
def openapiart():
    """Return an instance of the OpenApiArt class

    Instantiating the OpenApiArt class generates OpenAPI artifacts and an 
    enhanced ux python package.
    """
    api_files = [
        os.path.join(os.path.dirname(__file__), './api/info.yaml'),
        os.path.join(os.path.dirname(__file__), './api/api.yaml')
    ]
    sys.path.append(os.path.normpath(os.path.join(os.path.dirname(__file__), '..', '..')))
    module = importlib.import_module('openapiart.openapiart')
    openapiart = getattr(module, 'OpenApiArt')(api_files=api_files)
    return openapiart


@pytest.fixture
def api(openapiart):
    """Return an instance of the top level Api class from the generated package
    """
    sys.path.append(openapiart.output_dir)
    module = importlib.import_module(openapiart.python_module_name)
    package = getattr(module, openapiart.python_module_name)
    return package.api(location=None, verify=False, logger=None, loglevel=logging.DEBUG)
