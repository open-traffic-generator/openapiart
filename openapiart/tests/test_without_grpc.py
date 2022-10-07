import os
import sys
import time
import pytest
import importlib
import logging
from openapiart.openapiart import OpenApiArt as openapiart_class
from .server import app

pytest.withoutgrpc_module = "withoutgrpc"


def create_without_grpc_artifacts(openapiart_class):
    open_api = openapiart_class(
        api_files=[
            os.path.join(os.path.dirname(__file__), "./api/info.yaml"),
            os.path.join(os.path.dirname(__file__), "./common/common.yaml"),
            os.path.join(os.path.dirname(__file__), "./api/api.yaml"),
        ],
        artifact_dir=os.path.join(
            os.path.dirname(__file__), "..", "..", "art_without_grpc"
        ),
        extension_prefix=pytest.withoutgrpc_module,
    )
    open_api.GeneratePythonSdk(
        package_name=pytest.withoutgrpc_module,
        generate_grpc=False,
    )


def test_module():
    create_without_grpc_artifacts(openapiart_class)
    sys.path.append(
        os.path.join(
            os.path.dirname(__file__),
            "..",
            "..",
            "art_without_grpc",
            pytest.withoutgrpc_module,
        )
    )
    module = importlib.import_module(pytest.withoutgrpc_module)
    api = module.api()
    assert api.__module__ == pytest.withoutgrpc_module
    with pytest.raises(Exception) as execinfo:
        importlib.import_module(pytest.withoutgrpc_module + "_pb2")
    grpc_error = "No module named 'withoutgrpc_pb2'"
    assert execinfo.value.args[0] == grpc_error


def test_http_client():
    module = importlib.import_module(pytest.withoutgrpc_module)
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
    assert api.__module__ == pytest.withoutgrpc_module

    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    config.level.l1_p1.l2_p1.l3_p1 = "test"
    config.level.l1_p2.l4_p1.l1_p2.l4_p1.l1_p1.l2_p1.l3_p1 = "test"
    api.set_config(config)
    _config = api.get_config()
    assert config.serialize(config.DICT) == _config.serialize(_config.DICT)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
