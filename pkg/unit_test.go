package openapiart_test

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func TestJsonSerialization(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.RequiredObject()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.F().SetFB(3.0)
	config.G().Add().SetGA("a g_a value").SetGB(6).SetGC(77.7).SetGE(3.0)
	config.J().Add().JA().SetEA(1.0).SetEB(2.0)
	config.K().EObject().SetEA(77.7).SetEB(2.0)
	config.K().FObject().SetFA("asdf")
	l := config.L()
	l.SetString("test")
	l.SetInteger(80)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a")
	l.SetIpv4("1.1.1.1")
	l.SetIpv6("2000::1")
	l.SetHex("0x12")
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
	config.MacPattern().Mac().Increment().SetStart("00:00:00:00:00:0a").SetStart("00:00:00:00:00:01").SetCount(100)
	config.MacPattern().Mac().Decrement().SetStart("00:00:00:00:00:0a").SetStart("00:00:00:00:00:01").SetCount(100)
	config.ChecksumPattern().Checksum().SetCustom(64)
	fmt.Println(config.ToJson())

	// TBD: this needs to be fixed as order of json keys is not guaranteed to be the same
	// out := config.ToJson()
	// actualJson := []byte(out)
	// bs, err := ioutil.ReadFile("expected.json")
	// if err != nil {
	// 	log.Println("Error occured while reading config")
	// 	return
	// }
	// expectedJson := bs
	// eq, _ := JSONBytesEqual(actualJson, expectedJson)
	// assert.Equal(t, eq, true)
	yaml := config.ToYaml()
	log.Print(yaml)
}

func TestNewAndSet(t *testing.T) {
	c := openapiart.NewPrefixConfig()
	c.SetE(openapiart.NewEObject().SetEA(123.456))
	c.SetF(openapiart.NewFObject().SetFA("fa string"))
	log.Println(c.E().ToYaml())
	log.Println(c.F().ToYaml())
}

func TestSimpleTypes(t *testing.T) {
	a := "asdfg"
	var b float32 = 12.2
	var c int32 = 1
	h := true
	i := []byte("sample string")
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdfg").SetB(12.2).SetC(1).SetH(true).SetI([]byte("sample string"))
	assert.Equal(t, a, config.A())
	assert.Equal(t, b, config.B())
	assert.Equal(t, c, config.C())
	assert.Equal(t, h, config.H())
	assert.Equal(t, i, config.I())
}

func TestEObject(t *testing.T) {
	var ea float32 = 1.1
	eb := 1.2
	mparam1 := "Mparam1"
	maparam2 := "Mparam2"
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	assert.Equal(t, ea, config.E().EA())
	assert.Equal(t, eb, config.E().EB())
	assert.Equal(t, mparam1, config.E().MParam1())
	assert.Equal(t, maparam2, config.E().MParam2())
	log.Print(config.E().ToJson(), config.E().ToYaml())
}

func TestGObject(t *testing.T) {
	ga := []string{"g_1", "g_2"}
	gb := []int32{1, 2}
	gc := []float32{11.1, 22.2}
	ge := []float64{1.0, 2.0}
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	g1 := config.G().Add()
	g1.SetGA("g_1").SetGB(1).SetGC(11.1).SetGE(1.0)
	g2 := config.G().Add()
	g2.SetGA("g_2").SetGB(2).SetGC(22.2).SetGE(2.0)
	for i, G := range config.G().Items() {
		assert.Equal(t, ga[i], G.GA())
		assert.Equal(t, gb[i], G.GB())
		assert.Equal(t, gc[i], G.GC())
		assert.Equal(t, ge[i], G.GE())
	}
	log.Print(g1.ToJson(), g1.ToYaml())
}

func TestLObject(t *testing.T) {
	var int_ int32 = 80
	var float_ float32 = 100.11
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	l := config.L()
	l.SetString("test")
	l.SetInteger(80)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a")
	l.SetIpv4("1.1.1.1")
	l.SetIpv6("2000::1")
	l.SetHex("0x12")
	assert.Equal(t, "test", config.L().String())
	assert.Equal(t, int_, config.L().Integer())
	assert.Equal(t, float_, config.L().Float())
	assert.Equal(t, 1.7976931348623157e+308, config.L().Double())
	assert.Equal(t, "00:00:00:00:00:0a", config.L().Mac())
	assert.Equal(t, "1.1.1.1", config.L().Ipv4())
	assert.Equal(t, "2000::1", config.L().Ipv6())
	assert.Equal(t, "0x12", config.L().Hex())
	log.Print(l.ToJson(), l.ToYaml())
}

// TestRequiredObject
//  This test MUST create the underlying protobuf EObject
//  The generated wrapper get accessor must create the underlying protobuf EObject
//  Confirm the underlying protobuf EObject has been created by setting the
//  properties of the returned RequiredObject
func TestRequiredObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	r := config.RequiredObject()
	r.SetEA(22.2)
	r.SetEB(66.1)
}

// TestOptionalObject
//  This test MUST create the underlying protobuf EObject
//  The generated wrapper get accessor must create the underlying protobuf EObject
//  Confirm the underlying protobuf EObject has been created by setting the
//  properties of the returned OptionalObject
func TestOptionalObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	r := config.OptionalObject()
	r.SetEA(22.2)
	r.SetEB(66.1)
}

func TestResponseEnum(t *testing.T) {
	// UNCOMMENT the following when github workflow supports go 1.17
	// flds := reflect.VisibleFields(reflect.TypeOf(openapiart.PrefixConfigResponse))
	// for _, fld := range flds {
	// 	assert.NotEqual(t, fld.Name, "UNSPECIFIED")
	// }
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
	assert.Equal(t, config.Response(), openapiart.PrefixConfigResponse.STATUS_400)
	fmt.Println("response: ", config.Response())
}

func TestChoice(t *testing.T) {
	api := openapiart.NewApi()
	config := NewFullyPopulatedPrefixConfig(api)

	f := config.F()
	fmt.Println(f.ToJson())
	f.SetFA("a fa string")
	assert.Equal(t, f.Choice(), openapiart.FObjectChoice.F_A)

	j := config.J().Add()
	j.JA().SetEA(22.2)
	assert.Equal(t, j.Choice(), openapiart.JObjectChoice.J_A)
	j.JB()
	assert.Equal(t, j.Choice(), openapiart.JObjectChoice.J_B)

	fmt.Println(config.ToYaml())
}

func TestHas(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	assert.False(t, config.HasE())
	assert.False(t, config.HasF())
	assert.False(t, config.HasChecksumPattern())
	assert.False(t, config.HasFullDuplex100Mb())
	assert.False(t, config.HasIeee8021Qbb())
	assert.False(t, config.HasOptionalObject())
}

var GoodMac = []string{"ab:ab:10:12:ff:ff"}
var BadMac = []string{
	"1", "2.2", "1.1.1.1", "::01", "00:00:00", "00:00:00:00:gg:00", "00:00:fa:ce:fa:ce:01", "255:255:255:255:255:255",
}

func TestGoodMacValidation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	mac := config.MacPattern().Mac().SetValue(GoodMac[0])
	err := mac.Validate()
	assert.Nil(t, err)
}

func TestBadMacValidation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	for _, mac := range BadMac {
		macObj := config.MacPattern().Mac().SetValue(mac)
		err := macObj.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Mac")
		}
	}
}

func TestGoodMacValues(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	mac := config.MacPattern().Mac().SetValues(GoodMac)
	err := mac.Validate()
	assert.Nil(t, err)
}

func TestBadMacValues(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	mac := config.MacPattern().Mac().SetValues(BadMac)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

func TestBadMacIncrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	mac := config.MacPattern().Mac().Increment().SetStart(GoodMac[0])
	mac.SetStep(BadMac[0])
	mac.SetCount(10)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

func TestBadMacDecrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	mac := config.MacPattern().Mac().Decrement().SetStart(BadMac[0])
	mac.SetStep(GoodMac[0])
	mac.SetCount(10)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

var GoodIpv4 = []string{"1.1.1.1", "255.255.255.255"}
var BadIpv4 = []string{"1.1. 1.1", "33.4", "asdf", "100", "-20", "::01", "1.1.1.1.1", "256.256.256.256", "-255.-255.-255.-255"}

func TestGoodIpv4Validation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().SetValue(GoodIpv4[0])
	err := ipv4.Validate()
	assert.Nil(t, err)
}

func TestBadIpv4Validation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	for _, ip := range BadIpv4 {
		ipv4 := config.Ipv4Pattern().Ipv4().SetValue(ip)
		err := ipv4.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Ipv4")
		}
	}
}

func TestBadIpv4Values(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().SetValues(BadIpv4)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid ipv4 addresses")
	}
}

func TestBadIpv4Increment(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().Increment().SetStart(GoodIpv4[0])
	ipv4.SetStep(BadIpv4[0])
	ipv4.SetCount(10)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid Ipv4")
	}
}

func TestBadIpv4Decrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().Decrement().SetStart(GoodIpv4[0])
	ipv4.SetStep(BadIpv4[0])
	ipv4.SetCount(10)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid Ipv4")
	}
}

var GoodIpv6 = []string{"::", "1::", ": :", "abcd::1234", "aa:00bd:a:b:c:d:f:abcd"}
var BadIpv6 = []string{"33.4", "asdf", "1.1.1.1", "100", "-20", "65535::65535", "ab: :ab", "ab:ab:ab", "ffff0::ffff0"}

func TestGoodIpv6Validation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv6 := config.Ipv6Pattern().Ipv6().SetValue(GoodIpv6[0])
	err := ipv6.Validate()
	assert.Nil(t, err)
}

func TestBadIpv6Validation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	for _, ip := range BadIpv6 {
		ipv6 := config.Ipv6Pattern().Ipv6().SetValue(ip)
		err := ipv6.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
		}
	}
}

func TestBadIpv6Values(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv6 := config.Ipv6Pattern().Ipv6().SetValues(BadIpv6)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6 address")
	}
}

func TestBadIpv6Increment(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv6 := config.Ipv6Pattern().Ipv6().Increment().SetStart(GoodIpv6[0])
	ipv6.SetStep(BadIpv6[0])
	ipv6.SetCount(10)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
	}
}

func TestBadIpv6Decrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv6 := config.Ipv6Pattern().Ipv6().Decrement().SetStart(GoodIpv6[0])
	ipv6.SetStep(BadIpv6[0])
	ipv6.SetCount(10)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
	}
}

func TestDefaultSimpleTypes(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.RequiredObject()
	actual_result := config.ToJson()
	expected_result := `{
		"a":"asdf", 
		"b" : 65, 
		"c" : 33,  
		"h": true, 
		"response" : "status_200", 
		"required_object" : {
			"e_a" : 1, 
			"e_b" : 2
		}
	}`
	require.JSONEq(t, actual_result, expected_result)
}

func TestDefaultEObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.E()
	actual_result := config.E().ToJson()
	expected_result := `
	{
		"e_a":  1,
		"e_b":  2
	}`
	require.JSONEq(t, actual_result, expected_result)
}

func TestDefaultFObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.F()
	actual_result := config.F().ToJson()
	expected_result := `
	{
		"choice": "f_a",
		"f_a": "some string"
	}`
	require.JSONEq(t, actual_result, expected_result)
}

func TestRequiredValidation(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.RequiredObject()
	config.MObject().
		SetString("asdf").
		SetInteger(63).
		SetDouble(55.4).
		SetFloat(33.2).
		SetHex("00AABB").
		SetMac("00:AA:BB:CC:DD:EE").
		SetIpv6("2001::1").
		SetIpv4("1.1.1.1")
	err := config.Validate()
	assert.Nil(t, err)
}

func TestHexPattern(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	l := config.L()
	l.SetHex("200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	err := l.Validate()
	fmt.Println(err)
	assert.Nil(t, err)
	l.SetHex("0x00200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	err1 := l.Validate()
	fmt.Println(err1)
	assert.Nil(t, err1)
	l.SetHex("")
	err2 := l.Validate()
	assert.NotNil(t, err2)
	fmt.Println(err2)
}

func TestChoice1(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	json := `{
		"choice": "f_b",
		"f_b": 30.0
	}`
	g := config.F().FromJson(json)
	assert.Nil(t, g)
	fmt.Println(config.F().ToJson())
	require.JSONEq(t, config.F().ToJson(), json)
	json2 := `{
		"choice": "f_a",
		"f_a": "this is f string"
	}`
	f := config.F().FromJson(json2)
	assert.Nil(t, f)
	require.JSONEq(t, config.F().ToJson(), json2)
	fmt.Println(config.F().ToJson())
}
