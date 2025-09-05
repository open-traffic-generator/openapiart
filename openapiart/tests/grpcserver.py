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

    def streamSetConfig(self, request_iterator, context):
        self._log("Executing streamSetConfig")
        full_str = b""
        for data in request_iterator:
            print("received chunk: ", data.chunk_size)
            full_str += data.datum
            self._log("received ")

        self._log("received all chunks ")
        self._log(full_str)
        obj = pb2.PrefixConfig()
        obj.ParseFromString(full_str)
        self._log(obj)

        response_200 = """
            {
                "response_bytes" : "%s"
            }
        """ % base64.b64encode(
            b"success"
        ).decode(
            "utf-8"
        )

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
        v.app_version = "1.2.3"
        response_200 = {"version": v.serialize(v.DICT)}
        res_obj = json_format.Parse(
            json.dumps(response_200), pb2.GetVersionResponse()
        )
        return res_obj

    def AppendConfig(self, request, context):
        self._log("Executing AppendConfig")
        response = {
            "warning_details": {
                "warnings": ["w1", "w2"],
            }
        }
        res_obj = json_format.Parse(
            json.dumps(response), pb2.AppendConfigResponse()
        )
        return res_obj

    def streamGetMetrics(self, request, context):
        print(request)
        metrics = op.api().metrics()
        p = metrics.ports.add()
        p.name = "p1"
        p.tx_frames = 1.23
        p.rx_frames = 3.45
        f = metrics.flows.add()
        f.name = "f1"
        f.tx_frames = 4.23
        f.rx_frames = 6.45
        js = metrics.serialize()
        obj = json_format.Parse(js, pb2.Metrics())
        bts = obj.SerializeToString()
        chunk_size = 10
        for i in range(0, len(bts), chunk_size):
            if i + chunk_size > len(bts):
                chunk = bts[i : len(bts)]
            else:
                chunk = bts[i : i + chunk_size]
            print("sending ", chunk)
            yield pb2.Data(datum=chunk)
        print("finished sending all")

    def streamGetConfig(self, request, context):
        config = op.api().prefix_config()
        config.a = "asdf"
        config.b = 1.1
        config.c = 50
        config.required_object.e_a = 1.1
        config.required_object.e_b = 1.2
        config.d_values = [config.A, config.B, config.C]
        res_obj = json_format.Parse(config.serialize(), pb2.PrefixConfig())
        bts = res_obj.SerializeToString()
        chunk_size = 10
        for i in range(0, len(bts), chunk_size):
            if i + chunk_size > len(bts):
                chunk = bts[i : len(bts)]
            else:
                chunk = bts[i : i + chunk_size]
            print("sending ", chunk)
            yield pb2.Data(datum=chunk)
        print("finished sending all")

    def UploadConfig(self, request, context):
        print(request)
        response = {
            "warning_details": {
                "warnings": ["w1", "w2"],
            }
        }
        res_obj = json_format.Parse(
            json.dumps(response), pb2.UploadConfigResponse()
        )
        return res_obj

    def streamUploadConfig(self, request_iterator, context):
        self._log("Executing streamUploadConfig")
        full_str = b""
        for data in request_iterator:
            print("received chunk: ", data.chunk_size)
            full_str += data.datum
            self._log("received ")

        self._log("received all chunks ")
        self._log(full_str)
        response = {
            "warning_details": {
                "warnings": ["w11", "w22"],
            }
        }
        res_obj = json_format.Parse(
            json.dumps(response), pb2.UploadConfigResponse()
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
