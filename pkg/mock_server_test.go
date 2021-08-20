package openapiart_test

import (
	"context"
	"fmt"
	"log"
	"net"

	. "github.com/open-traffic-generator/openapiart/pkg/sanity"

	"google.golang.org/grpc"
)

var (
	testPort   uint         = 40051
	testServer *grpc.Server = nil
)

type server struct {
	UnimplementedOpenapiServer
}

func StartMockServer() error {
	if testServer != nil {
		log.Print("MockServer: Server already running")
		return nil
	}

	addr := fmt.Sprintf("[::]:%d", testPort)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(fmt.Sprintf("MockServer: Server failed to listen on address %s", addr))
	}

	svr := grpc.NewServer()
	log.Print(fmt.Sprintf("MockServer: Server started and listening on address %s", addr))

	RegisterOpenapiServer(svr, &server{})
	log.Print("MockServer: Server subscribed with gRPC Protocol Service")

	go func() {
		if err := svr.Serve(lis); err != nil {
			log.Fatal("MockServer: Server failed to serve for incoming gRPC request.")
		}
	}()

	testServer = svr
	return nil
}

func (s *server) SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResponse, error) {
	var resp *SetConfigResponse
	switch req.PrefixConfig.Response.Enum().Number() {
	case PrefixConfig_Response_status_400.Enum().Number():
		resp = &SetConfigResponse{
			StatusCode_400: &SetConfigResponse_StatusCode400{
				ErrorDetails: &ErrorDetails{
					Errors: []string{"SetConfig has detected configuration errors"},
				},
			},
		}
	case PrefixConfig_Response_status_500.Enum().Number():
		resp = &SetConfigResponse{
			StatusCode_500: &SetConfigResponse_StatusCode500{
				Error: &Error{
					Errors: []string{"SetConfig has encountered a server error"},
				},
			},
		}
	case PrefixConfig_Response_status_200.Enum().Number():
	default:
		resp = &SetConfigResponse{
			StatusCode_200: &SetConfigResponse_StatusCode200{
				Bytes: []byte("SetConfig has completed successfully"),
			},
		}
	}
	return resp, nil
}
