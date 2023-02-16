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


class OpenapiServicer(pb2_grpc.OpenapiServicer):
    def __init__(self):
        self._prefix_config = None
        super(OpenapiServicer, self).__init__()

    def _log(self, value):
        print("gRPC Server: %s" % value)

    def SetConfig(self, request, context):
        self._log("Executing SetConfig")
        response_400 = """
            {
                "status_code_400" : {
                    "errors" : ["invalid value"]
                }
            }
            """

        response_200 = """
            {
                "status_code_200" : "%s"
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
        if test is not None and (test < 10 or test > 90):
            res_obj = json_format.Parse(response_400, pb2.SetConfigResponse())
        else:
            res_obj = json_format.Parse(response_200, pb2.SetConfigResponse())

        return res_obj

    def GetConfig(self, request, context):
        self._log("Executing GetConfig")
        response_200 = {"status_code_200": self._prefix_config}
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetConfigResponse()
        )
        return res_obj

    def GetVersion(self, request, context):
        self._log("Executing GetVersion")
        v = op.api().get_local_version()
        response_200 = {"status_code_200": v.serialize(v.DICT)}
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetVersionResponse()
        )
        return res_obj


def gRpcServer():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_OpenapiServicer_to_server(OpenapiServicer(), server)
    print("gRPC Server: Starting server. Listening on port %s." % GRPC_PORT)
    server.add_insecure_port("[::]:{}".format(GRPC_PORT))
    server.start()

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(5)
        print("Server shutdown gracefully")


def grpc_server():
    web_server_thread = threading.Thread(target=gRpcServer)
    web_server_thread.setDaemon(True)
    web_server_thread.start()


if __name__ == "__main__":
    gRpcServer()
