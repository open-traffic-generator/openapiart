import pytest
import sys
import os
import importlib
import logging


@pytest.fixture(scope="session")
def api():
    """Return an instance of the top level Api class from the generated package"""
    from .server import OpenApiServer

    # TBD: fix this hardcoding
    # artifacts should not be generated from here as these tests are run as sudo
    sys.path.append("./art")
    module = importlib.import_module("sanity")

    pytest.server = OpenApiServer(module).start()
    return module.api(location="http://127.0.0.1:80", verify=False, logger=None, loglevel=logging.DEBUG)


@pytest.fixture
def config(api):
    """Return a new instance of an empty config"""
    return api.prefix_config()


@pytest.fixture
def default_config(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    return config
