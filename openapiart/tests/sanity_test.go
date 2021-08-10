package openapiart_tests

import (
	"fmt"
	"testing"

	sanityux "./art/sanityux"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	api := sanityux.NewApi()
	assert.NotNil(t, api)
}

func TestApiGrpc(t *testing.T) {
	api := sanityux.NewApi().
		SetTransport(sanityux.ApiTransport.GRPC).
		SetGrpcLocation("127.0.0.1:5050").
		SetGrpcRequestTimeout(10)
	assert.NotNil(t, api.Transport())
	assert.NotNil(t, api.GrpcLocation())
}

func TestNewPrefixConfig(t *testing.T) {
	api := sanityux.NewApi()
	config := api.NewPrefixConfig()
	assert.NotNil(t, config)
}

func TestPrefixConfigSetName(t *testing.T) {
	config := sanityux.NewApi().NewPrefixConfig()
	name := "asdfasdf"
	config.SetName(name)
	assert.Equal(t, name, config.Name())
}

func TestYamlSerialization(t *testing.T) {
	api := sanityux.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("simple string")
	yaml := config.Yaml()
	yamlLength := len(yaml)
	assert.True(t, yamlLength > 10)
}

func TestNewPrefixConfigSimpleTypes(t *testing.T) {
	api := sanityux.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("simple string")
	config.SetB(12.2)
	config.SetC(-33)
	config.SetH(true)
	config.SetI([]byte("a simple byte string"))
	config.SetName("name string")
	assert.NotNil(t, config)
	fmt.Println(config.Yaml())
}

// Test setting/getting a nested object field
// func TestEObject(t *testing.T) {
// 	config := sanityux.NewApi().NewPrefixConfig()
// 	e := config.SetName("PrefixConfig Name").E().SetName("E Name")
// 	assert.NotNil(t, config.E().Name())
// 	assert.NotNil(t, e.Name())
// 	assert.Equal(t, e.Name(), config.E().Name())
// 	fmt.Println(e)
// }
