package openapiart_test

import (
	"log"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

var apis []openapiart.OpenapiartApi

func init() {
	StartMockGrpcServer()
	StartMockHttpServer()
	grpcApi := openapiart.NewApi()
	grpcApi.NewGrpcTransport().SetLocation(grpcServer.Location)
	apis = append(apis, grpcApi)
	httpApi := openapiart.NewApi()
	httpApi.NewHttpTransport().SetLocation(httpServer.Location)
	apis = append(apis, httpApi)
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

func TestGetConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		api.SetConfig(config)
		resp, err := api.GetConfig()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestUpdateConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		api.SetConfig(config)
		c := api.NewUpdateConfig()
		c.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
		updatedConfig, err := api.UpdateConfig(c)
		assert.Nil(t, err)
		assert.NotNil(t, updatedConfig)
	}
}

func TestGetMetrics(t *testing.T) {
	for _, api := range apis {
		metrics, err := api.GetMetrics()
		assert.Nil(t, err)
		assert.NotNil(t, metrics)
		for _, row := range metrics.Ports().Items() {
			log.Println(row.ToYaml())
		}
	}
}
