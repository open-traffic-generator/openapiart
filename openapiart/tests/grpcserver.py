import grpc
from concurrent import futures
from google.protobuf import json_format
import threading
import base64


def grpc_server(pb2, pb2_grpc):
    class OpenapiServicer(pb2_grpc.OpenapiServicer):
        def __init__(self):
            super().__init__()

        def SetConfig(self, request, context):
            response_400 = """
                {
                    "status_code_400" : {
                        "error_details" : {
                            "errors" : ["invalid value"]
                        }
                    }
                }
                """

            response_200 = """
                {
                    "status_code_200" : {
                        "bytes" : "%s"
                    }
                }
            """ % base64.b64encode(
                b"success"
            ).decode(
                "utf-8"
            )

            test = request.prefix_config.l.integer
            if test is not None and (test < 10 or test > 90):
                res_obj = json_format.Parse(response_400, pb2.SetConfigResponse())
            else:
                res_obj = json_format.Parse(response_200, pb2.SetConfigResponse())

            return res_obj

        def start(self):
            self._web_server_thread = threading.Thread(target=local_web_server)
            self._web_server_thread.setDaemon(True)
            self._web_server_thread.start()
            return self

    def local_web_server():
        server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        pb2_grpc.add_OpenapiServicer_to_server(OpenapiServicer(), server)

        print("Starting server. Listening on port 50051.")
        server.add_insecure_port("[::]:50051")
        server.start()

    return OpenapiServicer()
