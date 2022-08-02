import json
import sys
import os
from .server import app

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
import sanity


def test_log_info():
    api = sanity.api("http://127.0.0.1:{}".format(app.PORT))
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    api.set_config(config)


def test_log(api):
    config = api.prefix_config()
    config.a = "asdf"
    config.b = 1.1
    config.c = 1
    config.required_object.e_a = 1.1
    config.required_object.e_b = 1.2
    config.d_values = [config.A, config.B, config.C]
    api.set_config(config)


def test_grpc_log(utils, grpc_api):
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)
    grpc_api.set_config(payload)
