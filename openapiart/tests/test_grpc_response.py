import pytest
import grpc
from google.protobuf import json_format
import json
import base64


def test_grpc_response(utils, pb2, pb2_grpc):
    """
    Send a protobuf request to the mock server and
    validate the response with expected response
    """
    expected_response = """{
            "status_code_200": {
                "bytes": "%s"
            }
        }""" % base64.b64encode(b"success").decode('utf-8')

    # load the json from a file
    with open(utils.get_test_config_path("config.json")) as f:
        payload = json.load(f)

    # open a gRPC channel
    channel = grpc.insecure_channel('localhost:50051')

    # create a stub (client)
    stub = pb2_grpc.OpenapiStub(channel)

    # create a valid request message
    pb_obj = json_format.Parse(json.dumps(payload), pb2.PrefixConfig())

    # make the call
    req_obj = pb2.SetConfigRequest(prefix_config=pb_obj)

    # get the responses
    response = stub.SetConfig(req_obj)
    resp = json_format.MessageToJson(
        response, preserving_proto_field_name=True)
    assert utils.compare_json(resp, expected_response)


if __name__ == "__main__":
    pytest.main(["-v", "-s", __file__])
