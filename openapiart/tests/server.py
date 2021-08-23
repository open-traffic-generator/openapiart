from flask import Flask, request, Response
import threading
import json
import time
import os
import importlib
import sys

import pytest


app = Flask(__name__)
app.CONFIG = None
app.PACKAGE = None


@app.route("/config", methods=["POST"])
def set_config():
    config = app.PACKAGE.Api().prefix_config()
    config.deserialize(request.data.decode("utf-8"))
    test = config.h
    if test is not None and isinstance(test, bool) is False:
        return Response(status=590, response=json.dumps({"detail": "invalid data type"}), headers={"Content-Type": "application/json"})
    else:
        app.CONFIG = config
        return Response(status=200)


@app.route("/config", methods=["GET"])
def get_config():
    serialized_config = app.CONFIG.serialize()
    return Response(serialized_config, mimetype="application/json", status=200)


@app.after_request
def after_request(resp):
    print(request.method, request.url, " -> ", resp.status)
    return resp


def web_server():
    app.run(port=80, debug=True, use_reloader=False)


class OpenApiServer(object):
    def __init__(self, package):
        # TODO Shall change the below sanity path to be dynamic
        pkg_name = "sanity"
        lib_path = "../../.output/openapiart/%s" % pkg_name
        sys.path.append(os.path.join(os.path.join(os.path.dirname(__file__), lib_path)))
        app.PACKAGE = importlib.import_module(pkg_name)
        app.CONFIG = app.PACKAGE.Api().prefix_config()

    def start(self):
        self._web_server_thread = threading.Thread(target=web_server)
        self._web_server_thread.setDaemon(True)
        self._web_server_thread.start()
        self._wait_until_ready()
        return self

    def _wait_until_ready(self):
        api = app.PACKAGE.api(location="http://127.0.0.1:80")
        while True:
            try:
                api.get_config()
                break
            except Exception as e:
                print(e)
            time.sleep(0.1)
