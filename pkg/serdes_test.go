package openapiart_test

import (
	"encoding/json"
	"fmt"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

func NewFullyPopulatedPrefixConfig(api openapiart.OpenapiartApi) openapiart.PrefixConfig {
	config := api.NewPrefixConfig()
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
	l := config.L()
	l.SetStringParam("test")
	l.SetInteger(80)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a")
	l.SetIpv4("1.1.1.1")
	l.SetIpv6("2000::1")
	l.SetHex("0x12")
	config.SetListOfStringValues([]string{"first string", "second string", "third string"})
	level := config.Level()
	level.L1P1().L2P1().SetL3P1("test")
	level.L1P2().L4P1().L1P2().L4P1().L1P1().L2P1().SetL3P1("l3p1")
	config.Mandatory().SetRequiredParam("required")
	config.Ipv4Pattern().Ipv4().SetValue("1.1.1.1")
	config.Ipv4Pattern().Ipv4().SetValues([]string{"10.10.10.10", "20.20.20.20"})
	config.Ipv4Pattern().Ipv4().Increment().SetStart("1.1.1.1").SetStep("0.0.0.1").SetCount(100)
	config.Ipv4Pattern().Ipv4().Decrement().SetStart("1.1.1.1").SetStep("0.0.0.1").SetCount(100)
	config.Ipv6Pattern().Ipv6().SetValue("2000::1")
	config.Ipv6Pattern().Ipv6().SetValues([]string{"2000::1", "2001::2"})
	config.Ipv6Pattern().Ipv6().Increment().SetStart("2000::1").SetStep("::1").SetCount(100)
	config.Ipv6Pattern().Ipv6().Decrement().SetStart("3000::1").SetStep("::1").SetCount(100)
	config.IntegerPattern().Integer().SetValue(1)
	config.IntegerPattern().Integer().SetValues([]int32{1, 2, 3})
	config.IntegerPattern().Integer().Increment().SetStart(1).SetStart(1).SetCount(100)
	config.IntegerPattern().Integer().Decrement().SetStart(1).SetStart(1).SetCount(100)
	config.MacPattern().Mac().SetValue("00:00:00:00:00:0a")
	config.MacPattern().Mac().SetValues([]string{"00:00:00:00:00:0a", "00:00:00:00:00:0b", "00:00:00:00:00:0c"})
	config.MacPattern().Mac().Increment().SetStart("00:00:00:00:00:0a").SetStep("00:00:00:00:00:01").SetCount(100)
	config.MacPattern().Mac().Decrement().SetStart("00:00:00:00:00:0a").SetStep("00:00:00:00:00:01").SetCount(100)
	config.ChecksumPattern().Checksum().SetCustom(64)
	return config
}

func TestPrefixConfigYamlSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := NewFullyPopulatedPrefixConfig(api)

	yaml1, err := c1.ToYaml()
	assert.Nil(t, err)
	c2 := api.NewPrefixConfig()
	yaml_err := c2.FromYaml(yaml1)
	assert.Nil(t, yaml_err)
	yaml2, err := c2.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, yaml1, yaml2)
}

func TestPrefixConfigJsonSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := NewFullyPopulatedPrefixConfig(api)

	json1, err := c1.ToJson()
	assert.Nil(t, err)
	c2 := api.NewPrefixConfig()
	json_err := c2.FromJson(json1)
	assert.Nil(t, json_err)
	json2, err := c2.ToJson()
	assert.Nil(t, err)
	assert.Equal(t, json1, json2)
}

func TestPartialSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := NewFullyPopulatedPrefixConfig(api)

	// convert the configuration to a map[string]interface{}
	var jsonMap map[string]interface{}
	c1json, err := c1.ToJson()
	assert.Nil(t, err)
	unmarsh_err := json.Unmarshal([]byte(c1json), &jsonMap)
	assert.Nil(t, unmarsh_err)

	// extract just the e object
	data1, _ := json.Marshal(jsonMap["e"])

	// extract the first object in the g array
	data2, _ := json.Marshal(jsonMap["g"].([]interface{})[0].(map[string]interface{}))

	// create a new config that consists of just the e object and the g object
	c2 := api.NewPrefixConfig()
	json_err := c2.E().FromJson(string(data1))
	assert.Nil(t, json_err)
	json_err1 := c2.G().Add().FromJson(string(data2))
	assert.Nil(t, json_err1)
	yaml1, err := c2.E().ToYaml()
	assert.Nil(t, err)
	fmt.Println(yaml1)
	yaml2, err := c2.G().Add().ToYaml()
	assert.Nil(t, err)
	fmt.Println(yaml2)
}

func TestPrefixConfigPbTextSerDes(t *testing.T) {
	api := openapiart.NewApi()
	c1 := NewFullyPopulatedPrefixConfig(api)
	pbString, err := c1.ToPbText()
	assert.Nil(t, err)
	c2 := api.NewPrefixConfig()
	pbtext_err := c2.FromPbText(pbString)
	assert.Nil(t, pbtext_err)
	c1json, err := c1.ToJson()
	assert.Nil(t, err)
	c2json, err := c2.ToJson()
	assert.Nil(t, err)
	assert.Equal(t, c1json, c2json)
}

func TestArrayOfStringsSetGet(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	values := config.ListOfStringValues()
	assert.Equal(t, 0, len(values))
	values = config.SetListOfStringValues([]string{"one", "two", "three"}).ListOfStringValues()
	assert.Equal(t, 3, len(values))
}

func TestArrayOfEnumsSetGet(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	values := config.DValues()
	assert.Equal(t, 0, len(values))
	enums := []openapiart.PrefixConfigDValuesEnum{
		openapiart.PrefixConfigDValues.A,
		openapiart.PrefixConfigDValues.B,
		openapiart.PrefixConfigDValues.C,
	}
	values = config.SetDValues(enums).DValues()
	assert.Equal(t, 3, len(values))
}

func TestArrayOfIntegersSetGet(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	values := config.ListOfIntegerValues()
	assert.Equal(t, 0, len(values))
	values = config.SetListOfIntegerValues([]int32{1, 5, 23, 6}).ListOfIntegerValues()
	assert.Equal(t, 4, len(values))
}

func TestValidJsonDecode(t *testing.T) {
	// Valid FromJson
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a":"ixia", "b" : 8.8, "c" : 1, "response" : "status_200", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.Nil(t, err)
}

func TestBadKeyJsonDecode(t *testing.T) {
	// Valid Wrong key
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a":"ixia", "bz" : 8.8, "c" : 1, "response" : "status_200", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:14): unknown field "bz"`)
}

func TestBadEnumJsonDecode(t *testing.T) {
	// Valid Wrong key
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a":"ixia", "b" : 8.8, "c" : 1, "response" : "status_800", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:47): invalid value for enum type: "status_800"`)
}

func TestBadDatatypeJsonDecode(t *testing.T) {
	// Valid Wrong data type. configure "b" with string
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a":"ixia", "b" : "abc", "c" : 1, "response" : "status_200", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:20): invalid value for float type: "abc"`)
}

func TestBadDatastructureJsonDecode(t *testing.T) {
	// Valid Wrong data structure. configure "a" with array
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a":["ixia"], "b" : 9.9, "c" : 1, "response" : "status_200", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:6): invalid value for string type: [`)
}

func TestWithoutValueJsonDecode(t *testing.T) {
	// Valid without value
	api := openapiart.NewApi()
	c1 := api.NewPrefixConfig()
	input_str := `{"a": "ixia", "b" : 8.8, "c" : "", "response" : "status_200", "required_object" : {"e_a": 1, "e_b": 2}}`
	err := c1.FromJson(input_str)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:32): invalid value for int32 type: ""`)
}

func TestValidYamlDecode(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	var data = `a: Easy
b: 12.2
c: 2
h: true
required_object:
  e_a: 1
  e_b: 2
response: status_200
`
	err := config.FromYaml(data)
	assert.Nil(t, err)
	configYaml, err := config.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, data, configYaml)
}

func TestBadKeyYamlDecode(t *testing.T) {
	// Valid Wrong key
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	var data = `a: Easy
bz: 12.2
c: 2
response: status_200
required_object:
  e_a: 1
  e_b: 2
`
	err := config.FromYaml(data)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:13): unknown field "bz"`)
}

func TestBadEnumYamlDecode(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	var data = `a: Easy
b: 12.2
c: 2
h: true
required_object:
  e_a: 1
  e_b: 2
response: status_800
`
	err := config.FromYaml(data)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:84): invalid value for enum type: "status_800"`)
}

func TestBadDatatypeYamlDecode(t *testing.T) {
	// Valid Wrong data type. configure "b" with string
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	var data = `a: Easy
b: abc
c: 2
response: status_200
required_object:
  e_a: 1
  e_b: 2
`
	err := config.FromYaml(data)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:17): invalid value for float type: "abc"`)
}

func TestBadDatastructureYamlDecode(t *testing.T) {
	// Valid Wrong data structure. configure "a" with array
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	var data = `a: [Make It Easy]
b: 9.9
c: 2
response: status_200
required_object:
  e_a: 1
  e_b: 2
`
	err := config.FromYaml(data)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), `unmarshal error (line 1:6): invalid value for string type: [`)
}

func TestSetMsg(t *testing.T) {
	api := openapiart.NewApi()
	config := NewFullyPopulatedPrefixConfig(api)
	copy := openapiart.NewApi().NewPrefixConfig()
	copy.SetMsg(config.Msg())
	configYaml, err := config.ToYaml()
	assert.Nil(t, err)
	copyYaml, err := copy.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, configYaml, copyYaml)
}

func TestNestedSetMsg(t *testing.T) {
	api := openapiart.NewApi()
	eObject := openapiart.NewApi().NewPrefixConfig().K().EObject()
	eObject.SetEA(23423.22)
	eObject.SetEB(10.24)
	eObject.SetName("asdfasdf")
	config := api.NewPrefixConfig()
	config.K().EObject().SetMsg(eObject.Msg())
	yaml1, err := config.K().EObject().ToYaml()
	assert.Nil(t, err)
	yaml2, err := eObject.ToYaml()
	assert.Nil(t, err)
	assert.Equal(t, yaml1, yaml2)
}

func TestAuto(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1)
	config.RequiredObject().SetEA(1).SetEB(2)
	assert.Equal(
		t,
		openapiart.PatternPrefixConfigAutoFieldTestChoiceEnum("auto"),
		config.AutoFieldTest().Choice())
	assert.Equal(t, int32(0), config.AutoFieldTest().Auto())

}
