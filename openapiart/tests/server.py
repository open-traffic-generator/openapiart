from flask import Flask, request, Response
import threading
import json
import os
import importlib
import sys


app = Flask(__name__)
app.CONFIG = None
app.PACKAGE = None
app.UPDATE_CONFIG = None
app.PORT = 8444
app.HOST = "0.0.0.0"


@app.route("/api/config", methods=["POST"])
def set_config():
    config = app.PACKAGE.Api().prefix_config()
    config.deserialize(request.data.decode("utf-8"))
    test = config.h
    if test is not None and isinstance(test, bool) is False:
        return Response(
            status=590,
            response=json.dumps({"detail": "invalid data type"}),
            headers={"Content-Type": "application/json"},
        )
    else:
        app.CONFIG = config
        return Response(status=200)


@app.route("/api/config", methods=["PATCH"])
def update_configuration():
    config = app.PACKAGE.Api().update_config()
    config.deserialize(request.data.decode("utf-8"))
    app.UPDATE_CONFIG = config
    return Response(status=200)


@app.route("/api/config", methods=["GET"])
def get_config():
    app.CONFIG.a = "asdf"
    app.CONFIG.b = 1.1
    app.CONFIG.c = 1
    app.CONFIG.required_object.e_a = 1.1
    app.CONFIG.required_object.e_b = 1.2
    serialized_config = app.CONFIG.serialize()
    return Response(serialized_config, mimetype="application/json", status=200)


@app.route("/api/capabilities/version", methods=["GET"])
def get_version():
    return Response(
        app.PACKAGE.Api().get_local_version().serialize(),
        mimetype="application/json",
        status=200,
    )


@app.after_request
def after_request(resp):
    print(request.method, request.url, " -> ", resp.status)
    return resp


class OpenApiServer(object):
    def __init__(self, package):
        # TODO Shall change the below sanity path to be dynamic
        pkg_name = "sanity"
        lib_path = "../../art/%s" % pkg_name
        sys.path.append(
            os.path.join(os.path.join(os.path.dirname(__file__), lib_path))
        )
        app.PACKAGE = importlib.import_module(pkg_name)
        app.CONFIG = app.PACKAGE.Api().prefix_config()

    @staticmethod
    def web_server():
        app.run(host=app.HOST, port=app.PORT, debug=False, use_reloader=False)

    def start(self):
        self._web_server_thread = threading.Thread(
            target=OpenApiServer.web_server
        )
        self._web_server_thread.setDaemon(True)
        self._web_server_thread.start()
        return self
