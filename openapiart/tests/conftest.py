import pytest
import sys
import os
import importlib
import logging
from .utils import common as utl
from .server import OpenApiServer
from .grpcserver import grpc_server

# TBD: fix this hardcoding
# artifacts should not be generated from here as these tests are run as sudo
pytest.module_name = "sanity"
sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art", pytest.module_name))
pytest.module = importlib.import_module(pytest.module_name)
pytest.http_server = OpenApiServer(pytest.module).start()
pytest.pb2_module = importlib.import_module(pytest.module_name + "_pb2")
pytest.pb2_grpc_module = importlib.import_module(pytest.module_name + "_pb2_grpc")
pytest.grpc_server = grpc_server(pytest.pb2_module, pytest.pb2_grpc_module).start()


@pytest.fixture(scope="session")
def api():
    """Return an instance of the top level Api class from the generated package"""
    module = importlib.import_module(pytest.module_name)
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
    return pytest.pb2_module


@pytest.fixture(scope="session")
def pb2_grpc():
    """Returns pb2_grpc package"""
    return pytest.pb2_grpc_module
