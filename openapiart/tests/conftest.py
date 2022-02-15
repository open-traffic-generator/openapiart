import pytest
import sys
import os
import importlib
import logging
import time
import yaml
from .utils import common as utl
from .server import OpenApiServer
from .grpcserver import grpc_server, GRPC_PORT
from .server import app


# TBD: fix this hardcoding
# artifacts should not be generated from here as these tests are run as sudo
pytest.module_name = "sanity"
pytest.artifacts_path = os.path.join(
    os.path.dirname(__file__), "..", "..", "art"
)
sys.path.append(pytest.artifacts_path)
sys.path.append(
    os.path.join(
        os.path.dirname(__file__), "..", "..", "art", pytest.module_name
    )
)

pytest.module = importlib.import_module(pytest.module_name)
pytest.http_server = OpenApiServer(pytest.module).start()
pytest.pb2_module = importlib.import_module(pytest.module_name + "_pb2")
pytest.pb2_grpc_module = importlib.import_module(
    pytest.module_name + "_pb2_grpc"
)
pytest.grpc_server = grpc_server()


@pytest.fixture(scope="session")
def api():
    """Return an instance of the top level Api class from the generated package"""
    module = importlib.import_module(pytest.module_name)
    api = module.api(
        location="http://127.0.0.1:{}".format(app.PORT),
        verify=False,
        logger=None,
        loglevel=logging.DEBUG,
    )
    # verify http server is up
    attempts = 1
    while True:
        try:
            api.get_config()
            break
        except Exception as e:
            print(e)
            if attempts > 5:
                raise (e)
        time.sleep(0.5)
        attempts += 1
    return api


@pytest.fixture(scope="session")
def grpc_api():
    """Return an instance of the top level gRPC Api class from the generated package"""
    return pytest.module.api(
        location="localhost:{}".format(GRPC_PORT),
        transport=pytest.module.Transport.GRPC,
        logger=None,
        loglevel=logging.DEBUG,
    )


@pytest.fixture(scope="session")
def proto_file_name():
    art_dir = os.path.join(os.path.dirname(__file__), "..", "..", "art")
    proto_file = os.path.join(art_dir, "{}.proto".format(pytest.module_name))
    return proto_file


@pytest.fixture
def config(api):
    """Return a new instance of an empty config"""
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    return config


@pytest.fixture
def default_config(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
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


@pytest.fixture
def openapi_yaml():
    path = os.path.join(pytest.artifacts_path, "openapi.yaml")
    _openapi = None
    with open(path) as fp:
        _openapi = yaml.safe_load(fp.read())
    return _openapi
