import os
import sys
import grpc
import json
import base64
import threading
import importlib
from concurrent import futures
from google.protobuf import json_format

sys.path.append(os.path.join(os.path.dirname(__file__), "..", "..", "art"))
sys.path.append(
    os.path.join(os.path.dirname(__file__), "..", "..", "art", "sanity")
)
pb2_grpc = importlib.import_module("sanity_pb2_grpc")
pb2 = importlib.import_module("sanity_pb2")
op = importlib.import_module("sanity")

GRPC_PORT = 40052
SECURE_GRPC_PORT = 40055


class OpenapiServicer(pb2_grpc.OpenapiServicer):
    def __init__(self):
        self._prefix_config = None
        super(OpenapiServicer, self).__init__()

    def _log(self, value):
        print("gRPC Server: %s" % value)

    def SetConfig(self, request, context):
        self._log("Executing SetConfig")
        response_200 = """
            {
                "response_bytes" : "%s"
            }
        """ % base64.b64encode(
            b"success"
        ).decode(
            "utf-8"
        )

        test = request.prefix_config.l.integer
        self._prefix_config = json_format.MessageToDict(
            request.prefix_config, preserving_proto_field_name=True
        )
        if test is not None and (test > 90):
            err = op.api().error()
            err.code = 13
            err.errors = ["err1", "err2"]
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(err.serialize())
            res_obj = pb2.SetConfigResponse()
        elif test is not None and (test < 0):
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("some random error!")
            res_obj = pb2.SetConfigResponse()
        else:
            res_obj = json_format.Parse(response_200, pb2.SetConfigResponse())

        return res_obj

    def GetConfig(self, request, context):
        self._log("Executing GetConfig")
        response_200 = {"prefix_config": self._prefix_config}
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetConfigResponse()
        )
        return res_obj

    def GetVersion(self, request, context):
        self._log("Executing GetVersion")
        v = op.api().get_local_version()
        response_200 = {"version": v.serialize(v.DICT)}
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetVersionResponse()
        )
        return res_obj


def gRpcServer(secure):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_OpenapiServicer_to_server(OpenapiServicer(), server)
    if secure:
        print("connection is secure")
        print(
            "gRPC Server: Starting server. Listening on port %s."
            % SECURE_GRPC_PORT
        )
        base_dir = os.path.dirname(__file__)
        key_location = os.path.join(base_dir, "./credentials/localhost.key")
        cert_location = os.path.join(base_dir, "./credentials/localhost.crt")
        key = open(key_location, "rb").read()
        cert = open(cert_location, "rb").read()
        server_credentials = grpc.ssl_server_credentials(
            (
                (
                    key,
                    cert,
                ),
            )
        )
        server.add_secure_port(
            "[::]:{}".format(SECURE_GRPC_PORT), server_credentials
        )
    else:
        print("connection is insecure")
        print(
            "gRPC Server: Starting server. Listening on port %s." % GRPC_PORT
        )
        server.add_insecure_port("[::]:{}".format(GRPC_PORT))
    server.start()

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(5)
        print("Server shutdown gracefully")


def grpc_server(secure=False):
    web_server_thread = threading.Thread(target=gRpcServer, args=(secure,))
    web_server_thread.setDaemon(True)
    web_server_thread.start()


if __name__ == "__main__":
    gRpcServer(True)
