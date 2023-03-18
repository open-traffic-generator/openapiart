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
	var err error
	switch req.PrefixConfig.Response.Enum().Number() {
	case sanity.PrefixConfig_Response_status_400.Enum().Number():
		resp = nil
		err = fmt.Errorf("SetConfig has detected configuration errors")
	case sanity.PrefixConfig_Response_status_500.Enum().Number():
		resp = nil
		errObj := openapiart.NewError()
		errObj.Msg().Code = 500
		tmp := errObj.SetKind("internal")
		fmt.Println(tmp)
		errObj.Msg().Errors = []string{"internal err 1"}
		jsonStr, e := errObj.ToJson()
		if e != nil {
			return resp, e
		}
		err = fmt.Errorf(jsonStr)
	case sanity.PrefixConfig_Response_status_200.Enum().Number():
		s.Config = req.PrefixConfig
		resp = &sanity.SetConfigResponse{
			ResponseBytes: []byte("SetConfig has completed successfully"),
		}
		err = nil
	case sanity.PrefixConfig_Response_status_404.Enum().Number():
		s.Config = req.PrefixConfig
		errObj := openapiart.NewError()
		errObj.Msg().Code = 404
		errObj.Msg().Errors = []string{"returning err1", "returning err2"}
		jsonStr, e := errObj.ToJson()
		if e != nil {
			return resp, e
		}
		err = fmt.Errorf(jsonStr)
	}
	return resp, err
}

func (s *GrpcServer) GetConfig(ctx context.Context, req *empty.Empty) (*sanity.GetConfigResponse, error) {
	resp := &sanity.GetConfigResponse{
		PrefixConfig: s.Config,
	}
	return resp, nil
}

func (s *GrpcServer) GetVersion(ctx context.Context, req *empty.Empty) (*sanity.GetVersionResponse, error) {
	resp := &sanity.GetVersionResponse{
		Version: openapiart.NewApi().GetLocalVersion().Msg(),
	}
	return resp, nil
}

func (s *GrpcServer) UpdateConfiguration(ctx context.Context, req *sanity.UpdateConfigurationRequest) (*sanity.UpdateConfigurationResponse, error) {

	if len(req.UpdateConfig.G) > 0 {
		if req.UpdateConfig.G[0].GetName() == "ErStr" {
			return nil, fmt.Errorf("unit test error")
		} else if req.UpdateConfig.G[0].GetName() == "Erkind" {
			errObj := openapiart.NewError()
			errObj.Msg().Code = 404
			tmp := errObj.SetKind("validation")
			fmt.Println(tmp)
			errObj.Msg().Errors = []string{"invalid1", "invalid2"}
			jsonStr, e := errObj.ToJson()
			if e != nil {
				return nil, e
			}
			return nil, fmt.Errorf(jsonStr)
		} else {
			resp := &sanity.UpdateConfigurationResponse{
				PrefixConfig: s.Config,
			}
			return resp, nil
		}
	} else {
		return nil, nil
	}
}

func (s *GrpcServer) GetMetrics(ctx context.Context, req *sanity.GetMetricsRequest) (*sanity.GetMetricsResponse, error) {
	choice := req.MetricsRequest.Choice.String()
	switch choice {
	case "port":
		choice_val := sanity.Metrics_Choice_Enum(sanity.Metrics_Choice_ports)
		resp := &sanity.GetMetricsResponse{
			Metrics: &sanity.Metrics{
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
			Metrics: &sanity.Metrics{
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
		WarningDetails: &sanity.WarningDetails{
			Warnings: []string{"This is your first warning", "Your last warning"},
		},
	}
	return resp, nil
}

func (s *GrpcServer) ClearWarnings(ctx context.Context, empty *empty.Empty) (*sanity.ClearWarningsResponse, error) {
	value := "warnings cleared"
	resp := &sanity.ClearWarningsResponse{
		String_: value,
	}
	return resp, nil
}
