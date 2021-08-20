package openapiart_test

import (
	"fmt"
	"log"
	"testing"

	. "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

func init() {
	if err := StartMockServer(); err != nil {
		log.Fatal("Mock Server Init failed")
	}
}

func TestApi(t *testing.T) {
	api := NewApi()
	assert.NotNil(t, api)
}

func TestGrpcTransport(t *testing.T) {
	location := "127.0.0.1:5050"
	timeout := 10
	api := NewApi()
	transport := api.NewGrpcTransport().SetLocation("127.0.0.1:5050").SetRequestTimeout(10)
	assert.NotNil(t, transport)
	assert.NotNil(t, transport.Location(), location)
	assert.NotNil(t, transport.RequestTimeout(), timeout)
}

func TestHttpTransport(t *testing.T) {
	location := "https://127.0.0.1:5050"
	verify := false
	api := NewApi()
	transport := api.NewHttpTransport().SetLocation(location).SetVerify(verify)
	assert.NotNil(t, transport)
	assert.NotNil(t, transport.Location(), location)
	assert.NotNil(t, transport.Verify(), verify)
}

func TestNewPrefixConfig(t *testing.T) {
	api := NewApi()
	config := api.NewPrefixConfig()
	assert.NotNil(t, config)
}

func TestPrefixConfigSetName(t *testing.T) {
	config := NewApi().NewPrefixConfig()
	name := "asdfasdf"
	config.SetName(name)
	assert.Equal(t, name, config.Name())
}

func TestYamlSerialization(t *testing.T) {
	api := NewApi()
	config := api.NewPrefixConfig()
	config.SetA("simple string")
	yaml := config.ToJson()
	yamlLength := len(yaml)
	assert.True(t, yamlLength > 10)
}

func TestNewPrefixConfigSimpleTypes(t *testing.T) {
	api := NewApi()
	config := api.NewPrefixConfig()
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
	config := NewApi().NewPrefixConfig()
	e := config.SetName("PrefixConfig Name").E().SetName("E Name")
	f := config.F().SetFA("a f_a value")
	assert.NotNil(t, config.E().Name())
	assert.NotNil(t, e.Name())
	assert.Equal(t, e.Name(), config.E().Name())
	assert.Equal(t, f.FA(), config.F().FA())
	fmt.Println(config.ToYaml())
}

func TestAddObject(t *testing.T) {
	config := NewApi().NewPrefixConfig()
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
	api := NewApi()
	api.NewGrpcTransport().SetLocation(fmt.Sprintf("127.0.0.1:%d", testPort))
	c := api.NewPrefixConfig()
	c.SetA("asdfasdf").SetB(22.2).SetC(33).E().SetEA(44.4)
	resp, err := api.SetConfig(c)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
