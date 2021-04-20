import pytest
import sys
import os
import importlib


@pytest.fixture(scope='session')
def generator():
    api_files = [
        os.path.join(os.path.dirname(__file__), './api/info.yaml'),
        os.path.join(os.path.dirname(__file__), './api/api.yaml')
    ]
    sys.path.append(os.path.normpath(os.path.join(os.path.dirname(__file__), '..', '..')))
    module = importlib.import_module('openapiart.openapiart')
    generator = getattr(module, 'OpenApiArt')(api_files=api_files)
    return generator


@pytest.fixture
def api(generator):
    sys.path.append(generator.output_dir)
    module = importlib.import_module(generator.python_module_name)
    api = getattr(module, 'api')()
    return api
