package openapiart_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

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
		log.Fatal(fmt.Sprintf("failed grpc dialcontext due to %s", err.Error()))
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
		resp, err := api.ClearWarnings()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestConnectionClose(t *testing.T) {
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(grpcServer.Location)
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	err1 := api.Close()
	assert.Nil(t, err1)
}

func TestGrpcClientConnection(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, grpcServer.Location, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(fmt.Sprintf("failed grpc dialcontext due to %s", err.Error()))
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
