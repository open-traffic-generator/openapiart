package openapiart_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"

	. "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

type ResponseWarning struct {
	Warnings []string
}

func StartMockHttpServer() {
	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		data := ResponseWarning{Warnings: []string{"Nothing bad..."}}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	log.Fatal(http.ListenAndServe(":50051", nil))
}

func init() {
	go StartMockHttpServer()
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
	location := "http://127.0.0.1:50051"
	verify := false
	api := NewApi()
	transport := api.NewHttpTransport().SetLocation(location).SetVerify(verify)
	config := api.NewPrefixConfig()
	config.SetA("simple string")
	config.SetB(12.2)
	config.SetC(-33)
	config.SetH(true)
	config.SetI([]byte("a simple byte string"))
	config.SetName("name string")
	assert.NotNil(t, config)
	assert.NotNil(t, transport)
	assert.NotNil(t, transport.Location(), location)
	assert.NotNil(t, transport.Verify(), verify)
	err := api.SetConfig(config)
	fmt.Println("HTTP Transport Error", err)
	assert.Nil(t, err)
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
	yaml := config.Yaml()
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
	fmt.Println(config.Yaml())
}

func TestGetObject(t *testing.T) {
	config := NewApi().NewPrefixConfig()
	e := config.SetName("PrefixConfig Name").E().SetName("E Name")
	f := config.F().SetFA("a f_a value")
	assert.NotNil(t, config.E().Name())
	assert.NotNil(t, e.Name())
	assert.Equal(t, e.Name(), config.E().Name())
	assert.Equal(t, f.FA(), config.F().FA())
	fmt.Println(config.Yaml())
}

func TestAddObject(t *testing.T) {
	config := NewApi().NewPrefixConfig()
	g := config.NewG()
	g.SetName("G-1")
	fmt.Println(config.Yaml())
}

// func TestChoiceObject(t *testing.T) {
// 	config := openapiart.NewApi().NewPrefixConfig()
// 	fmt.Println(config.Yaml())
// }
