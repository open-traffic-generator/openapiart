package openapiart_test

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	sanity.UnimplementedOpenapiServer
	Location string
	Server   *grpc.Server
	Config   *sanity.PrefixConfig
}

var (
	grpcServer GrpcServer = GrpcServer{
		Location: "[::]:40052",
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
		log.Fatalf("MockGrpcServer: Server failed to listen on address %s", grpcServer.Location)
	}

	grpcServer.Server = grpc.NewServer()
	log.Printf("MockGrpcServer: Server started and listening on address %s", grpcServer.Location)

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
			StatusCode_400: &sanity.ErrorDetails{
				Errors: []string{"SetConfig has detected configuration errors"},
			},
		}
	case sanity.PrefixConfig_Response_status_500.Enum().Number():
		resp = &sanity.SetConfigResponse{
			StatusCode_500: &sanity.Error{
				Errors: []string{"SetConfig has encountered a server error"},
			},
		}
	case sanity.PrefixConfig_Response_status_200.Enum().Number():
		s.Config = req.PrefixConfig
		resp = &sanity.SetConfigResponse{
			StatusCode_200: []byte("SetConfig has completed successfully"),
		}
	case sanity.PrefixConfig_Response_status_404.Enum().Number():
		s.Config = req.PrefixConfig
		resp = &sanity.SetConfigResponse{
			StatusCode_404: &sanity.ErrorDetails{
				Errors: []string{"Not found error"},
			},
		}
	}
	return resp, nil
}

func (s *GrpcServer) GetConfig(ctx context.Context, req *empty.Empty) (*sanity.GetConfigResponse, error) {
	resp := &sanity.GetConfigResponse{
		StatusCode_200: s.Config,
	}
	return resp, nil
}

func (s *GrpcServer) GetVersion(ctx context.Context, req *empty.Empty) (*sanity.GetVersionResponse, error) {
	resp := &sanity.GetVersionResponse{
		StatusCode_200: openapiart.NewApi().GetLocalVersion().Msg(),
	}
	return resp, nil
}

func (s *GrpcServer) UpdateConfiguration(ctx context.Context, req *sanity.UpdateConfigurationRequest) (*sanity.UpdateConfigurationResponse, error) {
	resp := &sanity.UpdateConfigurationResponse{
		StatusCode_200: s.Config,
	}
	return resp, nil
}

func (s *GrpcServer) GetMetrics(ctx context.Context, req *sanity.GetMetricsRequest) (*sanity.GetMetricsResponse, error) {
	choice := req.MetricsRequest.Choice.String()
	switch choice {
	case "port":
		choice_val := sanity.Metrics_Choice_Enum(sanity.Metrics_Choice_ports)
		resp := &sanity.GetMetricsResponse{
			StatusCode_200: &sanity.Metrics{
				Choice: &choice_val,
				Ports: []*sanity.PortMetric{
					{
						Name:     "P2",
						TxFrames: 3000,
						RxFrames: 2788,
					},
					{
						Name:     "P1",
						TxFrames: 2323,
						RxFrames: 2000,
					},
				},
			},
		}
		return resp, nil
	case "flow":
		choice_val := sanity.Metrics_Choice_Enum(sanity.Metrics_Choice_flows)
		resp := &sanity.GetMetricsResponse{
			StatusCode_200: &sanity.Metrics{
				Choice: &choice_val,
				Flows: []*sanity.FlowMetric{
					{
						Name:     "F2",
						TxFrames: 4000,
						RxFrames: 2000,
					},
					{
						Name:     "F1",
						TxFrames: 2000,
						RxFrames: 2000,
					},
				},
			},
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("Invalid choice")
	}
}

func (s *GrpcServer) GetWarnings(ctx context.Context, empty *empty.Empty) (*sanity.GetWarningsResponse, error) {
	resp := &sanity.GetWarningsResponse{
		StatusCode_200: &sanity.WarningDetails{
			Warnings: []string{"This is your first warning", "Your last warning"},
		},
	}
	return resp, nil
}

func (s *GrpcServer) ClearWarnings(ctx context.Context, empty *empty.Empty) (*sanity.ClearWarningsResponse, error) {
	value := "warnings cleared"
	resp := &sanity.ClearWarningsResponse{
		StatusCode_200: &value,
	}
	return resp, nil
}
