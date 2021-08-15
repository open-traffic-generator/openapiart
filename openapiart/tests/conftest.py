import pytest
import sys
import os
import importlib
import logging
from .utils import common as utl

# TBD: fix this hardcoding
# artifacts should not be generated from here as these tests are run as sudo
pytest.module_name = "sanity"
sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art", pytest.module_name))


@pytest.fixture(scope="session")
def api():
    """Return an instance of the top level Api class from the generated package"""
    from .server import OpenApiServer

    module = importlib.import_module(pytest.module_name)

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


@pytest.fixture
def pb_config(pb2):
    """Returns an instance of config of *_pb2 package"""
    return pb2.PrefixConfig()


@pytest.fixture
def utils():
    return utl


@pytest.fixture(scope="session")
def pb2():
    """Returns pb2 package"""
    return importlib.import_module(pytest.module_name + "_pb2")


@pytest.fixture(scope="session")
def pb2_grpc():
    """Returns pb2_grpc package"""
    return importlib.import_module(pytest.module_name + "_pb2_grpc")


@pytest.fixture(scope="session")
def grpc_api(pb2, pb2_grpc):
    """grpc API"""
    from .grpcserver import web_server

    grpc_api = web_server(pb2, pb2_grpc).start()
    yield grpc_api
