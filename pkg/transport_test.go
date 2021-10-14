package openapiart_test

import (
	"log"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

var apis []openapiart.OpenapiartApi

func init() {
	StartHttpServer()
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

func TestSetConfig400(t *testing.T) {
	for _, api := range apis {
		config := NewFullyPopulatedPrefixConfig(api)
		config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
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
		api.SetConfig(config)
		resp, err := api.GetConfig()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestUpdateConfigSuccess(t *testing.T) {
	for _, api := range apis {
		config1 := NewFullyPopulatedPrefixConfig(api)
		config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		api.SetConfig(config1)
		config2 := api.NewUpdateConfig()
		config2.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
		config3, err := api.UpdateConfig(config2)
		assert.Nil(t, err)
		assert.NotNil(t, config3)
		log.Println(config3.ToYaml())
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

func TestGetWarnings(t *testing.T) {
	for _, api := range apis {
		resp, err := api.GetWarnings()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		log.Println(resp.ToYaml())
	}
}
