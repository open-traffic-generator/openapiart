package openapiart_test

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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
		var code int32 = 500
		_ = errObj.SetCode(code)
		tmp := errObj.SetKind("internal")
		fmt.Println(tmp)
		_ = errObj.SetErrors([]string{"internal err 1"})
		jsonStr, e := errObj.Marshal().ToJson()
		if e != nil {
			return resp, e
		}
		err = status.Error(codes.InvalidArgument, jsonStr)
	case sanity.PrefixConfig_Response_status_200.Enum().Number():
		s.Config = req.PrefixConfig
		resp = &sanity.SetConfigResponse{
			ResponseBytes: []byte("SetConfig has completed successfully"),
		}
		err = nil
	case sanity.PrefixConfig_Response_status_404.Enum().Number():
		s.Config = req.PrefixConfig
		errObj := openapiart.NewError()
		var code int32 = 404
		_ = errObj.SetCode(code)
		_ = errObj.SetErrors([]string{"returning err1", "returning err2"})
		jsonStr, e := errObj.Marshal().ToJson()
		if e != nil {
			return resp, e
		}
		err = status.Error(codes.Internal, jsonStr)
	}
	return resp, err
}

func (s *GrpcServer) StreamSetConfig(srv sanity.Openapi_StreamSetConfigServer) error {
	var blob []byte
	idx := 0
	for {
		data, err := srv.Recv()
		if data != nil {
			fmt.Println("chunk size is:", data.ChunkSize)
			fmt.Println("Receiving chunk ", idx, time.Now().String())
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Transfer of %d bytes successful\n", len(blob))
				// log.Println(string(blob))

				resp := &sanity.SetConfigResponse{
					ResponseBytes: []byte("StreamConfig has completed successfully"),
				}
				config := openapiart.NewPrefixConfig()
				err := config.Unmarshal().FromPbText(string(blob))
				if err != nil {
					return err
				}
				fmt.Println(config.Marshal().ToYaml())
				return srv.SendAndClose(resp)
			}

			return err
		}
		idx++
		blob = append(blob, data.Datum...)

	}
}

func (s *GrpcServer) GetConfig(ctx context.Context, req *empty.Empty) (*sanity.GetConfigResponse, error) {
	resp := &sanity.GetConfigResponse{
		PrefixConfig: s.Config,
	}
	return resp, nil
}

func (s *GrpcServer) StreamGetConfig(req *empty.Empty, srv sanity.Openapi_StreamGetConfigServer) error {
	config := openapiart.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.F().SetFB(3.0)
	config.G().Add().SetGA("a g_a value").SetGB(6).SetGC(77.7).SetGE(3.0)
	config.J().Add().JA().SetEA(1.0).SetEB(2.0)
	config.K().EObject().SetEA(77.7).SetEB(2.0).SetName("An EB name")
	config.K().FObject().SetFA("asdf").SetFB(44.32232)
	text, err := config.Marshal().ToPbText()
	bytes := []byte(text)
	if err != nil {
		return err
	}
	chunkSize := 50
	for i := 0; i < len(bytes); i += chunkSize {
		data := &sanity.Data{}
		if i+chunkSize > len(bytes) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := srv.Send(data); err != nil {
			fmt.Printf("Failed to send: %v\n", err)
			return err
		}
		fmt.Printf("Sent: %v\n", data)
	}

	fmt.Println("Finished streaming")
	return nil
}

func (s *GrpcServer) GetVersion(ctx context.Context, req *empty.Empty) (*sanity.GetVersionResponse, error) {
	ver, _ := openapiart.NewApi().GetLocalVersion().SetAppVersion("1.2.3").Marshal().ToProto()
	resp := &sanity.GetVersionResponse{
		Version: ver,
	}
	return resp, nil
}

func (s *GrpcServer) UpdateConfiguration(ctx context.Context, req *sanity.UpdateConfigurationRequest) (*sanity.UpdateConfigurationResponse, error) {

	if len(req.UpdateConfig.G) > 0 {
		if req.UpdateConfig.G[0].GetName() == "ErStr" {
			return nil, fmt.Errorf("unit test error")
		} else if req.UpdateConfig.G[0].GetName() == "Erkind" {
			errObj := openapiart.NewError()
			var code int32 = 404
			_ = errObj.SetCode(code)
			tmp := errObj.SetKind("validation")
			fmt.Println(tmp)
			_ = errObj.SetErrors([]string{"invalid1", "invalid2"})
			jsonStr, e := errObj.Marshal().ToJson()
			if e != nil {
				return nil, e
			}
			return nil, status.Error(codes.AlreadyExists, jsonStr)
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

func (s *GrpcServer) AppendConfig(ctx context.Context, req *sanity.AppendConfigRequest) (*sanity.AppendConfigResponse, error) {
	if len(req.ConfigAppend.ConfigAppendList) < 2 {
		return nil, fmt.Errorf("length of configAppendList should be two")
	}

	resp := &sanity.AppendConfigResponse{
		WarningDetails: &sanity.WarningDetails{
			Warnings: []string{"This is your first warning", "Your last warning"},
		},
	}
	return resp, nil
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
						Name:     s.GetStringPtr("P2"),
						TxFrames: s.GetFloatPtr(3000),
						RxFrames: s.GetFloatPtr(2788),
					},
					{
						Name:     s.GetStringPtr("P1"),
						TxFrames: s.GetFloatPtr(2323),
						RxFrames: s.GetFloatPtr(2000),
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
						Name:     s.GetStringPtr("F2"),
						TxFrames: s.GetFloatPtr(4000),
						RxFrames: s.GetFloatPtr(2000),
					},
					{
						Name:     s.GetStringPtr("F1"),
						TxFrames: s.GetFloatPtr(2000),
						RxFrames: s.GetFloatPtr(2000),
					},
				},
			},
		}
		return resp, nil
	default:
		return nil, fmt.Errorf("Invalid choice")
	}
}

func (s *GrpcServer) StreamGetMetrics(req *sanity.GetMetricsRequest, srv sanity.Openapi_StreamGetMetricsServer) error {
	var resp *sanity.Metrics
	choice := req.MetricsRequest.Choice.String()
	switch choice {
	case "port":
		choice_val := sanity.Metrics_Choice_Enum(sanity.Metrics_Choice_ports)
		resp = &sanity.Metrics{
			Choice: &choice_val,
			Ports: []*sanity.PortMetric{
				{
					Name:     s.GetStringPtr("P2"),
					TxFrames: s.GetFloatPtr(3000),
					RxFrames: s.GetFloatPtr(2788),
				},
				{
					Name:     s.GetStringPtr("P1"),
					TxFrames: s.GetFloatPtr(2323),
					RxFrames: s.GetFloatPtr(2000),
				},
			},
		}
	case "flow":
		choice_val := sanity.Metrics_Choice_Enum(sanity.Metrics_Choice_flows)
		resp = &sanity.Metrics{
			Choice: &choice_val,
			Flows: []*sanity.FlowMetric{
				{
					Name:     s.GetStringPtr("F2"),
					TxFrames: s.GetFloatPtr(4000),
					RxFrames: s.GetFloatPtr(2000),
				},
				{
					Name:     s.GetStringPtr("F1"),
					TxFrames: s.GetFloatPtr(2000),
					RxFrames: s.GetFloatPtr(2000),
				},
			},
		}
	default:
		return fmt.Errorf("Invalid choice")
	}

	bytes, err := proto.Marshal(resp)
	if err != nil {
		return err
	}
	chunkSize := 20
	for i := 0; i < len(bytes); i += chunkSize {
		data := &sanity.Data{}
		if i+chunkSize > len(bytes) {
			data.Datum = bytes[i:]
		} else {
			data.Datum = bytes[i : i+chunkSize]
		}
		if err := srv.Send(data); err != nil {
			fmt.Printf("Failed to send: %v\n", err)
			return err
		}
		fmt.Printf("Sent: %v\n", data)
	}

	fmt.Println("Finished streaming")
	return nil

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

func (s *GrpcServer) GetStringPtr(value string) *string {
	return &value
}

func (s *GrpcServer) GetFloatPtr(value float64) *float64 {
	return &value
}

func (s *GrpcServer) UploadConfig(ctx context.Context, req *sanity.UploadConfigRequest) (*sanity.UploadConfigResponse, error) {
	if len(req.RequestBytes) < 2 {
		return nil, fmt.Errorf("length of bytes should be more than 2")
	}

	resp := &sanity.UploadConfigResponse{
		WarningDetails: &sanity.WarningDetails{
			Warnings: []string{"This is your first warning", "Your second warning"},
		},
	}
	return resp, nil
}

func (s *GrpcServer) StreamUploadConfig(srv sanity.Openapi_StreamUploadConfigServer) error {
	var blob []byte
	idx := 0
	for {
		data, err := srv.Recv()
		if data != nil {
			fmt.Println("chunk size is:", data.ChunkSize)
			fmt.Println("Receiving chunk ", idx, time.Now().String())
		}
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Transfer of %d bytes successful\n", len(blob))
				log.Println(string(blob))

				resp := &sanity.UploadConfigResponse{
					WarningDetails: &sanity.WarningDetails{
						Warnings: []string{"StreamuploadConfig has completed successfully"},
					},
				}
				return srv.SendAndClose(resp)
			}

			return err
		}
		idx++
		blob = append(blob, data.Datum...)

	}
}
