package openapiart_test

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcServer struct {
	sanity.UnimplementedOpenapiServer
	Location string
	Server   *grpc.Server
	Config   *sanity.PrefixConfig
}

var (
	grpcServer GrpcServer = GrpcServer{
		Location: "[::]:40051",
		Server:   nil,
	}
)

func StartMockGrpcServer() error {
	if grpcServer.Server != nil {
		log.Print("MockGrpcServer: Server already running")
		return nil
	}

	lis, err := net.Listen("tcp", grpcServer.Location)
	if err != nil {
		log.Fatal(fmt.Sprintf("MockGrpcServer: Server failed to listen on address %s", grpcServer.Location))
	}

	grpcServer.Server = grpc.NewServer()
	log.Print(fmt.Sprintf("MockGrpcServer: Server started and listening on address %s", grpcServer.Location))

	sanity.RegisterOpenapiServer(grpcServer.Server, &grpcServer)
	log.Print("MockGrpcServer: Server subscribed with gRPC Protocol Service")

	go func() {
		if err := grpcServer.Server.Serve(lis); err != nil {
			log.Fatal("MockGrpcServer: Server failed to serve for incoming gRPC request.")
		}
	}()

	return nil
}

func (s *GrpcServer) SetConfig(ctx context.Context, req *sanity.SetConfigRequest) (*sanity.SetConfigResponse, error) {
	var resp *sanity.SetConfigResponse
	switch req.PrefixConfig.Response.Enum().Number() {
	case sanity.PrefixConfig_Response_status_400.Enum().Number():
		resp = &sanity.SetConfigResponse{
			StatusCode_400: &sanity.SetConfigResponse_StatusCode400{
				ErrorDetails: &sanity.ErrorDetails{
					Errors: []string{"SetConfig has detected configuration errors"},
				},
			},
		}
	case sanity.PrefixConfig_Response_status_500.Enum().Number():
		resp = &sanity.SetConfigResponse{
			StatusCode_500: &sanity.SetConfigResponse_StatusCode500{
				Error: &sanity.Error{
					Errors: []string{"SetConfig has encountered a server error"},
				},
			},
		}
	case sanity.PrefixConfig_Response_status_200.Enum().Number():
		s.Config = req.PrefixConfig
		resp = &sanity.SetConfigResponse{
			StatusCode_200: &sanity.SetConfigResponse_StatusCode200{
				Bytes: []byte("SetConfig has completed successfully"),
			},
		}
	}
	return resp, nil
}

func (s *GrpcServer) GetConfig(ctx context.Context, req *empty.Empty) (*sanity.GetConfigResponse, error) {
	resp := &sanity.GetConfigResponse{
		StatusCode_200: &sanity.GetConfigResponse_StatusCode200{
			PrefixConfig: s.Config,
		},
	}
	return resp, nil
}

func (s *GrpcServer) UpdateConfig(ctx context.Context, req *sanity.UpdateConfigRequest) (*sanity.UpdateConfigResponse, error) {
	resp := &sanity.UpdateConfigResponse{
		StatusCode_200: &sanity.UpdateConfigResponse_StatusCode200{
			PrefixConfig: s.Config,
		},
	}
	return resp, nil
}

func (s *GrpcServer) GetMetrics(ctx context.Context, empty *emptypb.Empty) (*sanity.GetMetricsResponse, error) {
	metrics := sanity.Metrics{
		Ports: []*sanity.PortMetric{},
	}
	p1 := sanity.PortMetric{
		Name:     "P1",
		TxFrames: 2323,
		RxFrames: 2000,
	}
	p2 := sanity.PortMetric{
		Name:     "P2",
		TxFrames: 3000,
		RxFrames: 2788,
	}
	metrics.Ports = append(metrics.Ports, &p1)
	metrics.Ports = append(metrics.Ports, &p2)
	resp := &sanity.GetMetricsResponse{
		StatusCode_200: &sanity.GetMetricsResponse_StatusCode200{
			Metrics: &metrics,
		},
	}
	return resp, nil
}

func (s *GrpcServer) GetWarnings(ctx context.Context, empty *empty.Empty) (*sanity.GetWarningsResponse, error) {
	resp := &sanity.GetWarningsResponse{
		StatusCode_200: &sanity.GetWarningsResponse_StatusCode200{
			WarningDetails: &sanity.WarningDetails{
				Warnings: []string{"Warning number 1", "Your last warning"},
			},
		},
	}
	return resp, nil
}
