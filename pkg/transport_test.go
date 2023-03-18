package openapiart_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
	"testing"
	"time"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"runtime"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var apis []openapiart.OpenapiartApi

func init() {
	err := StartMockGrpcServer()
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	StartMockHttpServer()
	grpcApi := openapiart.NewApi()
	grpcApi.NewGrpcTransport().SetLocation(grpcServer.Location)
	apis = append(apis, grpcApi)

	httpApi := openapiart.NewApi()
	httpApi.NewHttpTransport().SetLocation(httpServer.Location)
	apis = append(apis, httpApi)

	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, grpcServer.Location, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed grpc dialcontext due to %s", err.Error())
	}
	clientConnApi := openapiart.NewApi()
	clientConnApi.NewGrpcTransport().SetClientConnection(conn)
	apis = append(apis, clientConnApi)
}

func TestSetConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		resp, err := api.SetConfig(config)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestSetConfig400(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
		resp, err := api.SetConfig(config)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
		log.Println(err)
	}
}

func TestSetConfig404(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_404)
		resp, err := api.SetConfig(config)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
	}
}

func TestSetConfig500(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)
		resp, err := api.SetConfig(config)
		assert.Nil(t, resp)
		assert.NotNil(t, err)
	}
}

func TestGetConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		_, err := api.SetConfig(config)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
		resp, err := api.GetConfig()
		fmt.Println(resp)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestUpdateConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config1 := NewFullyPopulatedPrefixConfig(api)
		config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		_, err := api.SetConfig(config1)
		if err != nil {
			log.Printf("error: %s", err.Error())
		}
		config2 := api.NewUpdateConfig()
		config2.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
		config3, err := api.UpdateConfiguration(config2)
		assert.Nil(t, err)
		assert.NotNil(t, config3)
	}
}

func TestGetMetrics(t *testing.T) {
	for _, api := range apis {
		metReq := openapiart.NewMetricsRequest()
		metReq.SetChoice(openapiart.MetricsRequestChoice.PORT)
		metrics, err := api.GetMetrics(metReq)
		assert.Nil(t, err)
		assert.NotNil(t, metrics)
		assert.Len(t, metrics.Ports().Items(), 2)
		m_err := metrics.Validate()
		assert.Nil(t, m_err)
		assert.Equal(t, openapiart.MetricsChoice.PORTS, metrics.Choice())
		for _, row := range metrics.Ports().Items() {
			log.Println(row.ToYaml())
		}
		metReqflow := openapiart.NewMetricsRequest()
		metReqflow.SetChoice(openapiart.MetricsRequestChoice.FLOW)
		metResp, err := api.GetMetrics(metReqflow)
		assert.Nil(t, err)
		assert.NotNil(t, metResp)
		assert.Len(t, metResp.Flows().Items(), 2)
		m_err1 := metResp.Validate()
		assert.Nil(t, m_err1)
		assert.Equal(t, openapiart.MetricsChoice.FLOWS, metResp.Choice())
		for _, row := range metResp.Flows().Items() {
			log.Println(row.ToYaml())
		}
	}
}

func TestGetWarnings(t *testing.T) {
	for _, api := range apis {
		resp, err := api.GetWarnings()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		log.Println(resp.ToYaml())
	}
}

func TestClearWarnings(t *testing.T) {
	for _, api := range apis {
		api.NewClearWarningsResponse()
		res, err := api.ClearWarnings()
		assert.Nil(t, err)
		assert.NotNil(t, res)
		log.Println(*res)
	}
}

func NetStat(t *testing.T) []string {
	var grep string
	grep = "grep"
	if runtime.GOOS == "windows" {
		grep = "findstr"
	}
	c1 := exec.Command("netstat", "-n")
	c2 := exec.Command(grep, "127.0.0.1:8444")
	r1, w1 := io.Pipe()

	c1.Stdout = w1
	c2.Stdin = r1
	var b2 bytes.Buffer
	c2.Stdout = &b2
	e1 := c1.Start()
	e2 := c2.Start()
	e4 := c1.Wait()
	e5 := w1.Close()
	_ = c2.Wait()

	assert.Nil(t, e1)
	assert.Nil(t, e2)
	assert.Nil(t, e4)
	assert.Nil(t, e5)

	var data []string
	for _, val := range strings.Split(b2.String(), "\n") {
		if val != "" {
			data = append(data, val)
		}
	}
	return data
}

func TestConnectionClose(t *testing.T) {
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(grpcServer.Location)
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	httpApi := openapiart.NewApi()
	httpApi.NewHttpTransport().SetLocation(httpServer.Location)
	config1 := NewFullyPopulatedPrefixConfig(httpApi)
	config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp1, err1 := httpApi.SetConfig(config1)
	assert.Nil(t, err1)
	assert.NotNil(t, resp1)

	err2 := api.Close()
	assert.Nil(t, err2)
	data := NetStat(t)
	fmt.Println(len(data))
	fmt.Println(data)
	assert.NotEqual(t, len(data), 0)
	err3 := httpApi.Close()
	assert.Nil(t, err3)
	// time.Sleep(10 * time.Second)
	data1 := NetStat(t)
	fmt.Println(len(data1))
	fmt.Println(data1)
	assert.Equal(t, len(data1), len(data)-2)
}

func TestGrpcClientConnection(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, grpcServer.Location, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed grpc dialcontext due to %s", err.Error())
	}
	api := openapiart.NewApi()
	grpc := api.NewGrpcTransport()
	grpc.SetClientConnection(conn)
	assert.NotNil(t, grpc.ClientConnection())
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestValidVersionCheckHttp(t *testing.T) {
	api := openapiart.NewApi()
	api.SetVersionCompatibilityCheck(true)
	api.NewHttpTransport().SetLocation(httpServer.Location)

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	log.Println(resp)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestInvalidVersionCheckHttp(t *testing.T) {
	api := openapiart.NewApi()
	api.SetVersionCompatibilityCheck(true)
	api.NewHttpTransport().SetLocation(httpServer.Location)
	api.GetLocalVersion().SetApiSpecVersion("0.2.0")

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestValidVersionCheckGrpc(t *testing.T) {
	api := openapiart.NewApi()
	api.SetVersionCompatibilityCheck(true)
	api.NewGrpcTransport().SetLocation(grpcServer.Location)

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestInvalidVersionCheckGrpc(t *testing.T) {
	api := openapiart.NewApi()
	api.SetVersionCompatibilityCheck(true)
	api.NewGrpcTransport().SetLocation(grpcServer.Location)
	api.GetLocalVersion().SetApiSpecVersion("0.2.0")

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGrpcErrorStructSetConfig(t *testing.T) {
	api := apis[0]
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_404)
	resp, err := api.SetConfig(config)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(404))
	assert.False(t, errSt.HasKind())
	assert.Equal(t, errSt.Errors()[0], "returning err1")
	assert.Equal(t, errSt.Errors()[1], "returning err2")
}

func TestHttpErrorStructSetConfig(t *testing.T) {
	api := apis[1]
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)
	resp, err := api.SetConfig(config)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(500))
	assert.Equal(t, errSt.Kind(), openapiart.ErrorKind.INTERNAL)
	assert.Equal(t, errSt.Errors()[0], "internal err 1")
	assert.Equal(t, errSt.Errors()[1], "internal err 2")
	assert.Equal(t, errSt.Errors()[2], "internal err 3")
}

func TestGrpcErrorStringSetConfig(t *testing.T) {
	api := apis[0]
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
	resp, err := api.SetConfig(config)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(2))
	assert.False(t, errSt.HasKind())
	assert.Equal(t, errSt.Errors()[0], "SetConfig has detected configuration errors")
}

func TestHttpErrorStringSetConfig(t *testing.T) {
	api := apis[1]
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
	resp, err := api.SetConfig(config)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(500))
	assert.False(t, errSt.HasKind())
	assert.Equal(t, errSt.Errors()[0], "client error !!!!")
}

func TestGrpcErrorkindSetConfig(t *testing.T) {
	api := apis[0]
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)
	resp, err := api.SetConfig(config)
	assert.Nil(t, resp)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(500))
	assert.Equal(t, errSt.Kind(), openapiart.ErrorKind.INTERNAL)
	assert.Equal(t, errSt.Errors()[0], "internal err 1")
}

func TestGrpcErrorStringUpdate(t *testing.T) {
	api := apis[2]
	config1 := NewFullyPopulatedPrefixConfig(api)
	config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	_, err := api.SetConfig(config1)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	config2 := api.NewUpdateConfig()
	config2.G().Add().SetName("ErStr").SetGA("ga string").SetGB(232)
	config3, err := api.UpdateConfiguration(config2)
	assert.Nil(t, config3)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(2))
	assert.False(t, errSt.HasKind())
	assert.Equal(t, errSt.Errors()[0], "unit test error")
}

func TestGrpcErrorStructUpdate(t *testing.T) {
	api := apis[2]
	config1 := NewFullyPopulatedPrefixConfig(api)
	config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	_, err := api.SetConfig(config1)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	config2 := api.NewUpdateConfig()
	config2.G().Add().SetName("Erkind").SetGA("ga string").SetGB(232)
	config3, err := api.UpdateConfiguration(config2)
	assert.Nil(t, config3)
	assert.NotNil(t, err)

	// if user wants to get the json now
	errSt := api.FromError(err)
	assert.Equal(t, errSt.Code(), int32(404))
	assert.Equal(t, errSt.Kind(), openapiart.ErrorKind.VALIDATION)
	assert.Equal(t, errSt.Errors()[0], "invalid1")
}
