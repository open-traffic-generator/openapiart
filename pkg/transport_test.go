package openapiart_test

import (
	"fmt"
	"log"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

func init() {
	if err := StartMockServer(); err != nil {
		log.Fatal("Mock Server Init failed")
	}
	go StartMockHttpServer()
}

func TestApi(t *testing.T) {
	api := openapiart.NewApi()
	assert.NotNil(t, api)
}

func TestGrpcTransport(t *testing.T) {
	location := "127.0.0.1:5050"
	timeout := 10
	api := openapiart.NewApi()
	transport := api.NewGrpcTransport().SetLocation("127.0.0.1:5050").SetRequestTimeout(10)
	assert.NotNil(t, transport)
	assert.NotNil(t, transport.Location(), location)
	assert.NotNil(t, transport.RequestTimeout(), timeout)
}

func TestHttpTransport(t *testing.T) {
	location := "https://127.0.0.1:5050"
	verify := false
	api := openapiart.NewApi()
	transport := api.NewHttpTransport().SetLocation(location).SetVerify(verify)
	assert.NotNil(t, transport)
	assert.NotNil(t, transport.Location(), location)
	assert.NotNil(t, transport.Verify(), verify)
}

func TestNewPrefixConfig(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	assert.NotNil(t, config)
}

func TestPrefixConfigSetName(t *testing.T) {
	config := openapiart.NewApi().NewPrefixConfig()
	name := "asdfasdf"
	config.SetName(name)
	assert.Equal(t, name, config.Name())
}

func TestNewPrefixConfigSimpleTypes(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetA("simple string")
	config.SetB(12.2)
	config.SetC(-33)
	config.SetH(true)
	config.SetI([]byte("a simple byte string"))
	config.SetName("name string")
	assert.NotNil(t, config)
	fmt.Println(config.ToYaml())
}

func TestGetObject(t *testing.T) {
	config := openapiart.NewApi().NewPrefixConfig()
	e := config.SetName("PrefixConfig Name").E().SetName("E Name")
	f := config.F().SetFA("a f_a value")
	assert.NotNil(t, config.E().Name())
	assert.NotNil(t, e.Name())
	assert.Equal(t, e.Name(), config.E().Name())
	assert.Equal(t, f.FA(), config.F().FA())
	fmt.Println(config.ToYaml())
}

func TestAddObject(t *testing.T) {
	config := openapiart.NewApi().NewPrefixConfig()
	config.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
	config.G().Add().SetName("G2")
	config.G().Add().SetName("G3").SetGA("3 is not 2 or 1 or none")
	config.J().Add().JB().SetFA("a FA string")
	config.J().Add().JA().SetEA(22.2)
	for _, item := range config.G().Items() {
		fmt.Println(item.Name())
	}
	name := "new persistent name"
	config.G().Items()[1].SetName(name)
	assert.Equal(t, len(config.G().Items()), 3)
	assert.Equal(t, config.G().Items()[1].Name(), name)
	fmt.Println(config.ToYaml())
}

func TestSetConfigSuccess(t *testing.T) {
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testPort))
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetConfigSuccess(t *testing.T) {
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testPort))
	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	api.SetConfig(config)
	resp, err := api.GetConfig()
	fmt.Println(resp.ToYaml())
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestUpdateConfigSuccess(t *testing.T) {
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testPort))

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	api.SetConfig(config)
	c := api.NewUpdateConfig()
	c.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
	updatedConfig, err := api.UpdateConfig(c)
	assert.Nil(t, err)
	assert.NotNil(t, updatedConfig)
	fmt.Println(updatedConfig.ToYaml())
}

func TestHttpSetConfigSuccess(t *testing.T) {
	location := fmt.Sprintf("127.0.0.1:%d", httpTestPort)
	verify := false
	api := openapiart.NewApi()
	api.NewHttpTransport().SetLocation(location).SetVerify(verify)

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	resp, err := api.SetConfig(config)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestHttpGetConfigSuccess(t *testing.T) {
	location := fmt.Sprintf("127.0.0.1:%d", httpTestPort)
	verify := false
	api := openapiart.NewApi()
	api.NewHttpTransport().SetLocation(location).SetVerify(verify)

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	api.SetConfig(config)
	resp, err := api.GetConfig()
	fmt.Println(resp.ToYaml())
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestHttpUpdateConfigSuccess(t *testing.T) {
	location := fmt.Sprintf("127.0.0.1:%d", httpTestPort)
	verify := false
	api := openapiart.NewApi()
	api.NewHttpTransport().SetLocation(location).SetVerify(verify)

	config := NewFullyPopulatedPrefixConfig(api)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	api.SetConfig(config)
	c := api.NewUpdateConfig()
	c.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
	updatedConfig, err := api.UpdateConfig(c)
	assert.Nil(t, err)
	assert.NotNil(t, updatedConfig)
	fmt.Println(updatedConfig.ToYaml())
}
