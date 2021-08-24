package openapiart_test

import (
	"encoding/json"
	"fmt"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

func TestPrefixConfigYamlSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	c1.SetA("a string").
		SetB(22.2).
		SetC(12).
		SetC(50).
		ChecksumPattern().
		Checksum().
		SetCustom(55)
	c1.G().Add().SetGA("a ga string")
	c1.E().SetEA(67.1)
	yaml1 := c1.ToYaml()
	c2 := api.NewPrefixConfig()
	c2.FromYaml(yaml1)
	yaml2 := c2.ToYaml()
	assert.Equal(t, yaml1, yaml2)
}

func TestPrefixConfigJsonSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	c1.SetA("a string").
		SetB(22.2).
		SetC(12).
		SetC(50).
		ChecksumPattern().
		Checksum().
		SetCustom(55)
	c1.G().Add().SetGA("a ga string")
	c1.E().SetEA(67.1)
	json1 := c1.ToJson()
	c2 := api.NewPrefixConfig()
	c2.FromJson(json1)
	json2 := c2.ToJson()
	assert.Equal(t, json1, json2)
}

func TestPartialSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	c1.SetA("a string").
		SetB(22.2).
		SetC(12).
		SetC(50).
		ChecksumPattern().
		Checksum().
		SetCustom(55)
	c1.G().Add().SetGA("a ga string")
	c1.E().SetEA(67.1)

	// convert the configuration to a map[string]interface{}
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(c1.ToJson()), &jsonMap)

	// extract just the e object
	data1, _ := json.Marshal(jsonMap["e"])

	// extract the first object in the g array
	data2, _ := json.Marshal(jsonMap["g"].([]interface{})[0].(map[string]interface{}))

	// create a new config that consists of just the e object and the g object
	c2 := api.NewPrefixConfig()
	c2.E().FromJson(string(data1))
	c2.G().Add().FromJson(string(data2))
	fmt.Println(c2.ToYaml())
}
