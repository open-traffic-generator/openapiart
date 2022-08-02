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
	config.RequiredObject().SetEA(3.0).SetEB(47.234)
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.F().SetFB(3.0)
	config.G().Add().SetGA("a g_a value").SetGB(6).SetGC(77.7).SetGE(3.0)
	config.J().Add().JA().SetEA(1.0).SetEB(2.0)
	config.K().EObject().SetEA(77.7).SetEB(2.0)
	config.K().FObject().SetFA("asdf")
	l := config.L()
	l.SetStringParam("test")
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
	yaml, err := config.ToYaml()
	assert.Nil(t, err)
	log.Print(yaml)
}

func TestNewAndSet(t *testing.T) {
	c := openapiart.NewPrefixConfig()
	c.SetE(openapiart.NewEObject().SetEA(123.456).SetEB(453.123))
	c.SetF(openapiart.NewFObject().SetFA("fa string"))
	yaml1, err := c.E().ToYaml()
	assert.Nil(t, err)
	yaml2, err := c.F().ToYaml()
	assert.Nil(t, err)
	log.Println(yaml1)
	log.Println(yaml2)
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

var gaValues = []string{"1111", "2222"}
var gbValues = []int32{11, 22}
var gcValues = []float32{11.11, 22.22}

func TestIterAdd(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.G().Add().SetGA("1111").SetGB(11).SetGC(11.11)
	config.G().Add().SetGA("2222").SetGB(22).SetGC(22.22)

	assert.Equal(t, len(config.G().Items()), 2)
	for idx, gObj := range config.G().Items() {
		assert.Equal(t, gaValues[idx], gObj.GA())
		assert.Equal(t, gbValues[idx], gObj.GB())
		assert.Equal(t, gcValues[idx], gObj.GC())
	}
}

func TestIterAppend(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.G().Add().SetGA("1111").SetGB(11).SetGC(11.11)
	g := config.G().Append(openapiart.NewGObject().SetGA("2222").SetGB(22).SetGC(22.22))

	assert.Equal(t, len(g.Items()), 2)
	for idx, gObj := range config.G().Items() {
		assert.Equal(t, gaValues[idx], gObj.GA())
		assert.Equal(t, gbValues[idx], gObj.GB())
		assert.Equal(t, gcValues[idx], gObj.GC())
	}
}

func TestIterSet(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			errValue := "runtime error: index out of range [3] with length 2"
			assert.Equal(t, errValue, fmt.Sprintf("%v", err))
		}
	}()
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	name := "new name set on slice"
	config.G().Add().SetName("original name set on add")
	config.G().Add()
	g := config.G().Set(1, openapiart.NewGObject().SetName(name))

	assert.Equal(t, name, g.Items()[1].Name())
	assert.Equal(t, len(g.Items()), 2)

	config.G().Set(3, openapiart.NewGObject().SetName(name))
}

func TestListWrapFromJson(t *testing.T) {
	var listWrap = `{
		"required_object":  {
		  "e_a":  3,
		  "e_b":  47.234
		},
		"response":  "status_200",
		"a":  "asdfg",
		"b":  12.2,
		"c":  1,
		"g":  [
		  {
			"g_a":  "1111",
			"g_b":  11,
			"g_c":  11.11,
			"choice":  "g_d",
			"g_d":  "some string",
			"g_f":  "a"
		  }
		],
		"h":  true
	  }`
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	assert.Nil(t, config.FromJson(listWrap))
	assert.Equal(t, len(config.G().Items()), 1)
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
	assert.Len(t, config.G().Items(), 2)
	for i, G := range config.G().Items() {
		assert.Equal(t, ga[i], G.GA())
		assert.Equal(t, gb[i], G.GB())
		assert.Equal(t, gc[i], G.GC())
		assert.Equal(t, ge[i], G.GE())
	}
	g1json, err := g1.ToJson()
	assert.Nil(t, err)
	g1yaml, err := g1.ToYaml()
	assert.Nil(t, err)
	log.Print(g1json, g1yaml)
}
func TestGObjectAppendMultiple(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	items := []openapiart.GObject{
		openapiart.NewGObject().SetGA("g_1"),
		openapiart.NewGObject().SetGA("g_2"),
		openapiart.NewGObject().SetGA("g_3"),
	}
	config.G().Append(items...)
	assert.Len(t, config.G().Items(), 3)
	item := config.G().Items()[1]
	assert.Equal(t, item.GA(), "g_2")
}

func TestGObjectAppend(t *testing.T) {
	ga := []string{"g_1", "g_2"}
	gb := []int32{1, 2}
	gc := []float32{11.1, 22.2}
	ge := []float64{1.0, 2.0}
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	g1 := openapiart.NewGObject()
	g1.SetGA("g_1").SetGB(1).SetGC(11.1).SetGE(1.0)
	g2 := openapiart.NewGObject()
	g2.SetGA("g_2").SetGB(2).SetGC(22.2).SetGE(2.0)
	config.G().Append(g1, g2)
	for i, G := range config.G().Items() {
		assert.Equal(t, ga[i], G.GA())
		assert.Equal(t, gb[i], G.GB())
		assert.Equal(t, gc[i], G.GC())
		assert.Equal(t, ge[i], G.GE())
	}
	// log.Print(g1.ToJson(), g1.ToYaml())
}

func TestLObject(t *testing.T) {
	var int_ int32 = 80
	var float_ float32 = 100.11
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	l := config.L()
	l.SetStringParam("test")
	l.SetInteger(80)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a")
	l.SetIpv4("1.1.1.1")
	l.SetIpv6("2000::1")
	l.SetHex("0x12")
	assert.Equal(t, "test", config.L().StringParam())
	assert.Equal(t, int_, config.L().Integer())
	assert.Equal(t, float_, config.L().Float())
	assert.Equal(t, 1.7976931348623157e+308, config.L().Double())
	assert.Equal(t, "00:00:00:00:00:0a", config.L().Mac())
	assert.Equal(t, "1.1.1.1", config.L().Ipv4())
	assert.Equal(t, "2000::1", config.L().Ipv6())
	assert.Equal(t, "0x12", config.L().Hex())
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
	fJson, err := f.ToJson()
	assert.Nil(t, err)
	fmt.Println(fJson)
	f.SetFA("a fa string")
	assert.Equal(t, f.Choice(), openapiart.FObjectChoice.F_A)

	j := config.J().Add()
	j.JA().SetEA(22.2).SetEB(44.32)
	assert.Equal(t, j.Choice(), openapiart.JObjectChoice.J_A)
	j.JB()
	assert.Equal(t, j.Choice(), openapiart.JObjectChoice.J_B)

	configYaml, err := config.ToYaml()
	assert.Nil(t, err)
	fmt.Println(configYaml)
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
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetA("asdf")
	config.SetB(65)
	config.SetC(33)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	actual_result, err := config.ToJson()
	assert.Nil(t, err)
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
	config.E().SetEA(1).SetEB(2)
	actual_result, err := config.E().ToJson()
	assert.Nil(t, err)
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
	actual_result, err := config.F().ToJson()
	assert.Nil(t, err)
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
	config.RequiredObject().SetEA(10.2).SetEB(20)
	config.SetA("abc")
	config.SetB(10.32)
	config.SetC(20)
	config.MObject().
		SetStringParam("asdf").
		SetInteger(63).
		SetDouble(55.4).
		SetFloat(33.2).
		SetHex("00AABB").
		SetMac("00:AA:BB:CC:DD:EE").
		SetIpv6("2001::1").
		SetIpv4("1.1.1.1")
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
	err := config.Validate()
	assert.Nil(t, err)
}

func TestHexPattern(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	l := config.L()
	l.SetHex("200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	err := l.Validate()
	assert.Nil(t, err)
	l.SetHex("0x00200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	err1 := l.Validate()
	assert.Nil(t, err1)
	l.SetHex("")
	err2 := l.Validate()
	assert.NotNil(t, err2)
	l.SetHex("0x00200000000000000b00000000200000000000000b00000000200000000000000b0000000x0")
	err3 := l.Validate()
	assert.NotNil(t, err3)
	l.SetHex("0x00")
	err4 := l.Validate()
	assert.Nil(t, err4)
	l.SetHex("0XAF12")
	err5 := l.Validate()
	assert.Nil(t, err5)
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
	configFjson, err := config.F().ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, configFjson, json)
	json2 := `{
		"choice": "f_a",
		"f_a": "this is f string"
	}`
	f := config.F().FromJson(json2)
	assert.Nil(t, f)
	configFjson2, err := config.F().ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, configFjson2, json2)
	fmt.Println(configFjson2)
}

func TestRequiredField(t *testing.T) {
	mandate := openapiart.NewMandate()
	err := mandate.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "RequiredParam is required field")
}

func TestOptionalDefault(t *testing.T) {
	gObject := openapiart.NewGObject()
	gJson := `{
		"g_a":  "asdf",
		"g_b":  6,
		"g_c":  77.7,
		"choice":  "g_d",
		"g_d":  "some string",
		"g_f":  "a"
	  }`
	gObjectJson, err := gObject.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, gObjectJson, gJson)
}

func TestInterger64(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	int_64 := `{
		"a":"asdf", 
		"b" : 65, 
		"c" : 33,
		"response" : "status_200", 
		"required_object" : {
			"e_a" : 1, 
			"e_b" : 2
		},
		"integer64": 100
	}`
	err := config.FromJson(int_64)
	fmt.Println(config.Integer64())
	assert.Nil(t, err)
	int_64_str := `{
		"a":"asdf", 
		"b" : 65, 
		"c" : 33,
		"response" : "status_200", 
		"required_object" : {
			"e_a" : 1, 
			"e_b" : 2
		},
		"integer64": "100"
	}`
	err1 := config.FromJson(int_64_str)
	fmt.Println(config.Integer64())
	assert.Nil(t, err1)
}

func TestFromJsonToCleanObject(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetA("abcd")
	config.SetB(100)
	config.SetC(4000)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)
	config.SetRequiredObject(openapiart.NewEObject().SetEA(10.1).SetEB(30.234))
	config.SetInteger64(200645)
	assert.Nil(t, config.Validate())
	new_json := `{
		"a":"asdf", 
		"b" : 65, 
		"c" : 33,
		"response" : "status_200", 
		"required_object" : {
			"e_a" : 1, 
			"e_b" : 2
		},
		"h": false
	}`
	err := config.FromJson(new_json)
	assert.Nil(t, err)
	configJson, err := config.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, new_json, configJson)
	new_json1 := `{
		"b" : 65, 
		"c" : 33,
		"response" : "status_200", 
		"required_object" : {
			"e_a" : 1, 
			"e_b" : 2
		},
		"h": false
	}`
	err1 := config.FromJson(new_json1)
	assert.NotNil(t, err1)
	assert.Contains(t, err1.Error(), "A is required field")
}

func TestChoiceStale(t *testing.T) {
	fObject := openapiart.NewFObject()
	fObject.SetFA("This is A Value")
	expected_json := `{
		"choice": "f_a",
		"f_a": "This is A Value"
	}`
	fObjectJson, err := fObject.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, expected_json, fObjectJson)
	fObject.SetFB(30.245)
	expected_json1 := `{
		"choice": "f_b",
		"f_b": 30.245
	}`
	fObjectJson2, err := fObject.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, expected_json1, fObjectJson2)
}

func TestChoice2(t *testing.T) {
	expected_json := `{
		"required_object": {
		  "e_a": 1,
		  "e_b": 2
		},
		"response": "status_200",
		"a": "asdf",
		"b": 12.2,
		"c": 1,
		"e": {
		  "e_a": 1.1,
		  "e_b": 1.2,
		  "m_param1": "Mparam1",
		  "m_param2": "Mparam2"
		},
		"h": true,
		"j": [
		  {
			"choice": "j_a",
			"j_a": {
			  "e_a": 1,
			  "e_b": 2
			}
		  },
		  {
			"choice": "j_b",
			"j_b": {
			  "choice": "f_a",
			  "f_a": "asf"
			}
		  }
		],
		"k": {
		  "f_object": {
			"choice": "f_a",
			"f_a": "asf"
		  }
		}
	  }`
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1)
	config.RequiredObject().SetEA(1).SetEB(2)
	config.K().FObject().SetFA("asf")
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.J().Add().JA().SetEA(1.0).SetEB(2.0)
	config.J().Add().JB().SetFA("asf")
	configJson, err := config.ToJson()
	assert.Nil(t, err)
	log.Print(configJson)
	require.JSONEq(t, expected_json, configJson)
}

func TestGetter(t *testing.T) {
	fObject := openapiart.NewFObject()
	fObject.FA()
	expected_json := `{
		"choice": "f_a",
		"f_a": "some string"
	}`
	fObjectJson, err := fObject.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, expected_json, fObjectJson)

	fObject1 := openapiart.NewFObject()
	fObject1.Choice()
	fObject1Json, err := fObject1.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, expected_json, fObject1Json)

	pattern := openapiart.NewIpv4Pattern()
	pattern.Ipv4()
	exp_ipv4 := `{
		"ipv4":  {
			"choice":  "value",
			"value":  "0.0.0.0"
		}
	}`
	patternJson, err := pattern.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, exp_ipv4, patternJson)
	pattern.Ipv4().SetValue("10.1.1.1")
	assert.Equal(t, "10.1.1.1", pattern.Ipv4().Value())
	pattern.Ipv4().Values()
	exp_ipv41 := `{
		"ipv4": {
			"choice": "values",
			"values": [
				"0.0.0.0"
			]
		}
	}`
	patternJson1, err := pattern.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, exp_ipv41, patternJson1)
	pattern.Ipv4().SetValues([]string{"20.1.1.1"})
	assert.Equal(t, []string{"20.1.1.1"}, pattern.Ipv4().Values())
	checksum := openapiart.NewChecksumPattern().Checksum()
	ch_json := `{
		"choice": "generated",
		"generated": "good"
	}`
	checksumJson, err := checksum.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, ch_json, checksumJson)
	fmt.Println(checksumJson)
}

func TestStringLength(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.SetStrLen("123456")
	configJson, err := config.ToJson()
	assert.Nil(t, err)
	log.Print(configJson)
}

func TestListClear(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	list := config.G()
	list.Append(openapiart.NewGObject().SetGA("a1"))
	list.Append(openapiart.NewGObject().SetGA("a2"))
	list.Append(openapiart.NewGObject().SetGA("a3"))
	assert.Len(t, list.Items(), 3)
	list.Clear()
	assert.Len(t, list.Items(), 0)
	list.Append(openapiart.NewGObject().SetGA("b1"))
	list.Append(openapiart.NewGObject().SetGA("b2"))
	assert.Len(t, list.Items(), 2)
	assert.Equal(t, list.Items()[1].GA(), "b2")

	list1 := []openapiart.GObject{
		openapiart.NewGObject().SetGA("c_1"),
		openapiart.NewGObject().SetGA("c_2"),
		openapiart.NewGObject().SetGA("c_3"),
	}
	list.Clear().Append(list1...)
	assert.Len(t, list.Items(), 3)
	list2 := []openapiart.GObject{
		openapiart.NewGObject().SetGA("d_1"),
		openapiart.NewGObject().SetGA("d_1"),
	}
	list.Clear().Append(list2...)
	assert.Len(t, list.Items(), 2)
}

func TestConfigHas200Result(t *testing.T) {
	// https://github.com/open-traffic-generator/openapiart/issues/242
	cfg := openapiart.NewSetConfigResponse()
	cfg.SetStatusCode200([]byte("anything"))
	assert.True(t, cfg.HasStatusCode200())
}

func TestFromJsonErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	json := `{
		"abc": "test"
	}`
	err := config.FromJson(json)
	assert.Contains(t, err.Error(), "unmarshal error")
	json1 := `{
		"choice": "g_e",
		"g_e": "10",
		"g_b": "20"
	}`
	gObj := openapiart.NewGObject()
	err1 := gObj.FromJson(json1)
	assert.Nil(t, err1)
	json2 := `{
		"choice": "f_t"
	}`
	fObj := openapiart.NewFObject()
	err2 := fObj.FromJson(json2)
	assert.Contains(t, err2.Error(), "unmarshal error")
}

func TestStringLengthError(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1}).SetName("config")
	config.SetSpace1(1)
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse("status_200")
	config.SetStrLen("12345678")
	config.StrLen()
	config.Space1()
	config.Name()
	err := config.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "3 <= length of prefixconfig.strlen <= 6 but got 8")
	}
}

func TestIncorrectChoiceEnum(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse("status_600")
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.F().SetFB(3.0)
	config.Ieee8021Qbb()
	config.FullDuplex100Mb()
	config.Response()
	err := config.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "status_600 is not a valid choice")
	}
}

// func TestEObjectValidation(t *testing.T) {
// 	eObject := openapiart.NewEObject()
// 	err := eObject.Validate()
// 	if assert.Error(t, err) {
// 		assert.Contains(t, strings.ToLower(err.Error()), "ea is required field on interface eobject\neb is required field on interface eobject\nvalidation errors")
// 	}
// }

func TestMObjectValidation(t *testing.T) {
	mObject := openapiart.NewMObject()
	err := mObject.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "required field on interface mobject")
	}
}

func TestMobjectValidationError(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.RequiredObject().SetEA(10.2).SetEB(20)
	config.SetA("abc")
	config.SetB(10.32)
	config.SetC(20)
	config.MObject().
		SetStringParam("asdf").
		SetInteger(120).
		SetDouble(55.4).
		SetFloat(33.2).
		SetHex("00AABBCCCIJ").
		SetMac("00:AA:BB:CC:DD:EE:AA").
		SetIpv6("2001::1::1").
		SetIpv4("1.1.1.1.2")
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_400)
	err := config.Validate()
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()),
			"invalid mac address",
			"invalid ipv4 address",
			"invalid hex value",
			"invalid ipv6 address")
	}
}

func TestLObjectError(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	l := config.L()
	l.SetStringParam("test")
	l.SetInteger(180)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a:22")
	l.SetIpv4("1.1.1.1.1.1")
	l.SetIpv6("2000::1:::4")
	l.SetHex("0x12KJN")
	err := config.Validate()
	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()),
			"invalid mac address",
			"invalid ipv4 address",
			"invalid hex value",
			"invalid ipv6 address")
	}
}

func TestIeee802x(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetCase(openapiart.NewLayer1Ieee802X())
	config.Case()
	l1 := openapiart.NewLayer1Ieee802X()
	l1.SetFlowControl(true)
	l1.SetMsg(l1.Msg())
	l1.FlowControl()
	l1.HasFlowControl()
	l1json, err := l1.ToJson()
	assert.Nil(t, err)
	l1yaml, err := l1.ToYaml()
	assert.Nil(t, err)
	l1pbText, err := l1.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, l1.FromJson(l1json))
	assert.Nil(t, l1.FromYaml(l1yaml))
	assert.Nil(t, l1.FromPbText(l1pbText))
}

func TestLevelFour(t *testing.T) {
	new_level_four := openapiart.NewLevelFour()
	new_level_four.Msg()
	new_level_four.SetMsg(new_level_four.Msg())
	new_level_four.HasL4P1()
	new_level_four.SetL4P1(new_level_four.L4P1())
	fourJson, err := new_level_four.ToJson()
	assert.Nil(t, err)
	fourYaml, err := new_level_four.ToYaml()
	assert.Nil(t, err)
	fourPbText, err := new_level_four.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_level_four.FromJson(fourJson))
	assert.Nil(t, new_level_four.FromYaml(fourYaml))
	assert.Nil(t, new_level_four.FromPbText(fourPbText))
}

func TestIterAppendJObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.J().Add()
	j := config.J().Append(openapiart.NewJObject())

	assert.Equal(t, len(j.Items()), 2)
}

func TestIterSetJObject(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.J().Add().JA().SetEA(100)
	config.J().Add()
	j := config.J().Set(1, openapiart.NewJObject().SetChoice("j_b"))

	assert.Contains(t, j.Items()[1].Choice(), "j_b")
	assert.Len(t, config.J().Items(), 2)
	config.J().Clear()
	assert.Len(t, config.J().Items(), 0)
}

func TestIterAppendGObject(t *testing.T) {
	config := openapiart.NewUpdateConfig()
	config.G().Add()
	g := config.G().Append(openapiart.NewGObject())

	assert.Equal(t, len(g.Items()), 2)
}

func TestIterSetGObject(t *testing.T) {
	config := openapiart.NewUpdateConfig()
	name := "new name set on slice"
	config.G().Add().SetName("original name set on add")
	config.G().Add()
	g := config.G().Set(1, openapiart.NewGObject().SetName(name))

	assert.Equal(t, name, g.Items()[1].Name())
	assert.Len(t, g.Items(), 2)
	g.Clear()
	assert.Len(t, g.Items(), 0)

}

func TestIterAppendPortMetrics(t *testing.T) {
	config := openapiart.NewMetrics()
	config.Ports().Add()
	p := config.Ports().Append(openapiart.NewPortMetric())

	assert.Equal(t, len(p.Items()), 2)
}

func TestIterSetPortMetrics(t *testing.T) {
	config := openapiart.NewMetrics()
	name := "new port metric"
	config.Ports().Add().SetName("original name set on add")
	config.Ports().Add()
	p := config.Ports().Set(1, openapiart.NewPortMetric().SetName(name))

	assert.Equal(t, name, p.Items()[1].Name())
	assert.Len(t, p.Items(), 2)
	p.Clear()
	assert.Len(t, p.Items(), 0)
}

func TestFromYamlErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2

	}`
	assert.NotNil(t, config.FromYaml(incorrect_format))
	incorrect_key := `{
		"a":"asdf",
		"z" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" : {
			"e_a" : 1,
			"e_b" : 2
		}
	}`
	assert.NotNil(t, config.FromYaml(incorrect_key))
	incorrect_value := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"str_len" : "abcdefg",
		"required_object" : {
			"e_a" : 1,
			"e_b" : 2
		}
	}`
	assert.NotNil(t, config.FromYaml(incorrect_value))
}

func TestFromPbTextErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2

	}`
	assert.NotNil(t, config.FromPbText(incorrect_format))
	incorrect_key := `{
		"a":"asdf",
		"z" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" : {
			"e_a" : 1,
			"e_b" : 2
		}
	}`
	assert.NotNil(t, config.FromPbText(incorrect_key))
}

func TestUpdateConfig(t *testing.T) {
	for _, api := range apis {
		config1 := NewFullyPopulatedPrefixConfig(api)
		config1.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
		_, set_err := api.SetConfig(config1)
		assert.Nil(t, set_err)
		config2 := api.NewUpdateConfig()
		config2.G().Add().SetName("G1").SetGA("ga string").SetGB(232)
		config2PbText, err := config2.ToPbText()
		assert.Nil(t, err)
		config2Json, err := config2.ToJson()
		assert.Nil(t, err)
		config2Yaml, err := config2.ToYaml()
		assert.Nil(t, err)
		assert.Nil(t, config2.FromJson(config2Json))
		assert.Nil(t, config2.FromYaml(config2Yaml))
		assert.Nil(t, config2.FromPbText(config2PbText))
		// config2.SetMsg(config2.Msg())
		config3, err := api.UpdateConfiguration(config2)
		assert.Nil(t, err)
		assert.NotNil(t, config3)
		config3Yaml, err := config3.ToYaml()
		assert.Nil(t, err)
		log.Println(config3Yaml)
	}
}

func TestNewSetConfigResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewSetConfigResponse()
	new_resp.SetStatusCode200([]byte{0, 1})
	new_resp.SetStatusCode400(new_resp.StatusCode400())
	new_resp.SetStatusCode500(new_resp.StatusCode500())
	new_resp.SetStatusCode404(new_resp.StatusCode400())
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode400()
	new_resp.HasStatusCode404()
	new_resp.HasStatusCode500()
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewUpdateConfigResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewUpdateConfigurationResponse()
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode200()
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewGetConfigResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewGetConfigResponse()
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode200()
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewGetMetricsResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewGetMetricsResponse()
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode200()
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewGetWarningsResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewGetWarningsResponse()
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode200()
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewClearWarningsResponse(t *testing.T) {
	api := openapiart.NewApi()
	new_resp := api.NewClearWarningsResponse()
	new_resp.SetMsg(new_resp.Msg())
	new_resp.Msg()
	new_resp.HasStatusCode200()
	new_resp.SetStatusCode200("success")
	respJson, err := new_resp.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_resp.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_resp.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_resp.FromJson(respJson))
	assert.Nil(t, new_resp.FromYaml(respYaml))
	assert.Nil(t, new_resp.FromPbText(respPbText))
}

func TestNewErrorDetails(t *testing.T) {
	new_err := openapiart.NewErrorDetails()
	new_err.SetMsg(new_err.Msg())
	new_err.Msg()
	respJson, err := new_err.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_err.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_err.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_err.FromJson(respJson))
	assert.Nil(t, new_err.FromYaml(respYaml))
	assert.Nil(t, new_err.FromPbText(respPbText))
}

func TestNewError(t *testing.T) {
	new_err := openapiart.NewError()
	new_err.SetMsg(new_err.Msg())
	new_err.Msg()
	respJson, err := new_err.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_err.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_err.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_err.FromJson(respJson))
	assert.Nil(t, new_err.FromYaml(respYaml))
	assert.Nil(t, new_err.FromPbText(respPbText))
	new_err.SetErrors(new_err.Errors())
}

func TestNewMetrics(t *testing.T) {
	new_metrics := openapiart.NewMetrics()
	new_metrics.SetMsg(new_metrics.Msg())
	new_metrics.Msg()
	respJson, err := new_metrics.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_metrics.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_metrics.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_metrics.FromJson(respJson))
	assert.Nil(t, new_metrics.FromYaml(respYaml))
	assert.Nil(t, new_metrics.FromPbText(respPbText))
}

func TestNewWarningDetails(t *testing.T) {
	new_warnings := openapiart.NewWarningDetails()
	new_warnings.SetMsg(new_warnings.Msg())
	new_warnings.Msg()
	respJson, err := new_warnings.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_warnings.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_warnings.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_warnings.FromJson(respJson))
	assert.Nil(t, new_warnings.FromYaml(respYaml))
	assert.Nil(t, new_warnings.FromPbText(respPbText))
}

func TestNewPortMetric(t *testing.T) {
	new_port_metric := openapiart.NewPortMetric()
	new_port_metric.SetName("portmetric")
	new_port_metric.SetTxFrames(1000)
	new_port_metric.SetRxFrames(2000)
	new_port_metric.SetMsg(new_port_metric.Msg())
	new_port_metric.Msg()
	respJson, err := new_port_metric.ToJson()
	assert.Nil(t, err)
	respYaml, err := new_port_metric.ToYaml()
	assert.Nil(t, err)
	respPbText, err := new_port_metric.ToPbText()
	assert.Nil(t, err)
	assert.Nil(t, new_port_metric.FromJson(respJson))
	assert.Nil(t, new_port_metric.FromYaml(respYaml))
	assert.Nil(t, new_port_metric.FromPbText(respPbText))
	new_port_metric.Name()
	new_port_metric.RxFrames()
	new_port_metric.TxFrames()
	assert.Nil(t, new_port_metric.Validate())
}

func TestItemsMethod(t *testing.T) {
	api := openapiart.NewApi()
	config1 := NewFullyPopulatedPrefixConfig(api)
	config1.G().Add().SetGA("this is GA string")
	assert.Equal(t, config1.G(), config1.G())
	config2 := api.NewPrefixConfig()
	config1Json, err := config1.ToJson()
	assert.Nil(t, err)
	json_err := config2.FromJson(config1Json)
	assert.Nil(t, json_err)
	assert.Len(t, config1.G().Items(), 2)
	assert.Len(t, config2.G().Items(), 2)
	for ind, obj := range config1.G().Items() {
		objJson, err := obj.ToJson()
		assert.Nil(t, err)
		gJson, err := config2.G().Items()[ind].ToJson()
		assert.Nil(t, err)
		assert.Equal(t, objJson, gJson)
	}
	configJson1, err := config1.ToJson()
	assert.Nil(t, err)
	config2Json, err := config2.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, configJson1, config2Json)
	config2.G().Add().SetGB(200000)
	assert.Len(t, config2.G().Items(), 3)
	for ind, obj := range config1.G().Items() {
		objJson1, err := obj.ToJson()
		assert.Nil(t, err)
		gJson1, err := config2.G().Items()[ind].ToJson()
		assert.Nil(t, err)
		assert.Equal(t, objJson1, gJson1)
	}
}

func TestStructGetterMethod(t *testing.T) {
	jObject := openapiart.NewJObject()
	val := jObject.JA()
	val.SetEA(1.45)
	val.SetEB(1.456)
	assert.Equal(t, val, jObject.JA())
	jObject.JA().SetEA(0.23495)
	assert.Equal(t, val, jObject.JA())

	jObject1 := openapiart.NewJObject()
	jOject1json, err := jObject.ToJson()
	assert.Nil(t, err)
	err1 := jObject1.FromJson(jOject1json)
	assert.Nil(t, err1)
	assert.Equal(t, jObject1.JA(), jObject1.JA())

	jObject2 := openapiart.NewJObject()
	val2 := jObject2.JA()
	val2.SetEA(0.23495).SetEB(1.456)
	jObject2Json, err := jObject.ToJson()
	assert.Nil(t, err)
	err2 := jObject2.FromJson(jObject2Json)
	assert.Nil(t, err2)
	assert.NotEqual(t, val2, jObject2.JA())
}

func TestFromJsonEmpty(t *testing.T) {
	fObject := openapiart.NewFObject()
	value1, err := fObject.ToJson()
	assert.Nil(t, err)
	value2, err := fObject.ToYaml()
	assert.Nil(t, err)
	value3, err := fObject.ToPbText()
	assert.Nil(t, err)
	for i, v := range []string{"", ``, `{}`, "{}"} {
		err1 := fObject.FromJson(v)
		assert.Nil(t, err1)
		err2 := fObject.FromYaml(v)
		assert.Nil(t, err2)
		if i < 2 {
			err3 := fObject.FromPbText(v)
			assert.Nil(t, err3)
		}
	}

	fObjectJson, err := fObject.ToJson()
	assert.Nil(t, err)
	fObjectYaml, err := fObject.ToYaml()
	assert.Nil(t, err)
	fObjectPbText, err := fObject.ToPbText()
	assert.Nil(t, err)
	require.JSONEq(t, value1, fObjectJson)
	require.Equal(t, value2, fObjectYaml)
	require.Equal(t, value3, fObjectPbText)
}

func TestChoiceDefaults(t *testing.T) {
	jObject := openapiart.NewJObject()
	json := `
	{
		"choice": "j_a",
		"j_a": {}
	}`
	j, err0 := jObject.ToJson()
	assert.Nil(t, err0)
	require.JSONEq(t, json, j)
	jObject.SetChoice(openapiart.JObjectChoice.J_B)
	json1 := `
	{
		"choice": "j_b",
		"j_b": {
			"choice": "f_a",
			"f_a": "some string"
		}
	}`
	j1, err1 := jObject.ToJson()
	assert.Nil(t, err1)
	require.JSONEq(t, json1, j1)
	jObject.JB().FB()
	json2 := `
	{
		"choice": "j_b",
		"j_b": {
			"choice": "f_b",
			"f_b": 3
		}
	}`
	j2, err2 := jObject.ToJson()
	assert.Nil(t, err2)
	require.JSONEq(t, json2, j2)
	integer := openapiart.NewIntegerPattern()
	integer.Integer().Values()
	json3 := `
	{
		"integer":  {
		  "choice":  "values",
		  "values":  [
			0
		  ]
		}
	}`
	j3, err3 := integer.ToJson()
	assert.Nil(t, err3)
	require.JSONEq(t, json3, j3)
	integer.Integer().SetValues([]int32{1, 2, 3})
	json4 := `
	{
		"integer":  {
		  "choice":  "values",
		  "values":  [
			1, 2, 3
		  ]
		}
	}`
	j4, err4 := integer.ToJson()
	assert.Nil(t, err4)
	require.JSONEq(t, json4, j4)
}

func TestSetterWrapperHolder(t *testing.T) {
	metricsResp := openapiart.NewGetMetricsResponse()
	metricsResp.SetStatusCode200(openapiart.NewMetrics())
	json1 := `{
		"status_code_200":  {
			"choice": "ports"
		}
	}`
	metricsrespJson, err := metricsResp.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, json1, metricsrespJson)
	fmt.Println(metricsrespJson)
	metricsResp.StatusCode200().Ports().Add().SetName("abc").SetRxFrames(100)
	json := `{
		"status_code_200":  {
		  "choice": "ports",
		  "ports":  [
			{
			  "name":  "abc",
			  "rx_frames":  100
			}
		  ]
		}
	}`
	metricsrespJson1, err := metricsResp.ToJson()
	assert.Nil(t, err)
	fmt.Println(metricsrespJson1)
	require.JSONEq(t, json, metricsrespJson1)
	metricsResp.SetStatusCode200(openapiart.NewMetrics())
	metricsrespJson2, err := metricsResp.ToJson()
	assert.Nil(t, err)
	require.JSONEq(t, json1, metricsrespJson2)
	fmt.Println(metricsrespJson2)
}

func TestStringer(t *testing.T) {
	lObj := openapiart.NewLObject()
	lObj.SetDouble(10.1).SetFloat(20.1).SetStringParam("abc")
	fmt.Println(lObj)
}

func TestClone(t *testing.T) {
	lObj1 := openapiart.NewLObject()
	lObj2, err := lObj1.Clone()
	assert.Nil(t, err)
	fmt.Println(&lObj1, &lObj2)
	assert.NotSame(t, &lObj1, &lObj2)
}

func TestDeprecationWarning(t *testing.T) {

	// Warning by config
	api := openapiart.NewApi()
	api.NewGrpcTransport().SetLocation(grpcServer.Location)
	config := api.NewPrefixConfig()
	config.RequiredObject().SetEA(10).SetEB(20)
	config.SetA("abc")
	config.SetB(20)
	config.SetC(30)

	warnings := api.GetApiWarnings()

	t.Log(warnings)

	assert.NotNil(t, warnings)
	assert.Len(t, warnings, 2)
	api.ClearApiWarnings()

	assert.Len(t, api.GetApiWarnings(), 0)

	// Warning by ToJson
	data, err := config.ToJson()

	assert.Nil(t, err)
	warnings1 := api.GetApiWarnings()

	t.Log(warnings1)

	assert.NotNil(t, warnings1)
	assert.Len(t, warnings1, 2)
	api.ClearApiWarnings()

	assert.Len(t, api.GetApiWarnings(), 0)

	config1 := api.NewPrefixConfig()

	// Warning by FromJson
	err1 := config1.FromJson(data)
	assert.Nil(t, err1)
	warnings2 := api.GetApiWarnings()

	t.Log(warnings)

	assert.NotNil(t, warnings2)
	assert.Len(t, warnings2, 2)
	api.ClearApiWarnings()

	assert.Len(t, api.GetApiWarnings(), 0)

	u_config := api.NewUpdateConfig()
	u_config.G().Add().SetGA("abcd")
	api.UpdateConfiguration(u_config)

	warnings = api.GetApiWarnings()

	t.Log(warnings)

	assert.NotNil(t, warnings)
	assert.Len(t, warnings, 1)
	api.ClearApiWarnings()

}

func TestConstraintAndUnique(t *testing.T) {
	prefix := openapiart.NewPrefixConfig()
	prefix.SetA("abc").SetB(10).SetC(32).RequiredObject().SetEA(20).SetEB(10)

	// *************** global unique ****************
	// Two similar objects with same Name.
	prefix.WList().Add().SetWName("global_unique_similar_obj")
	prefix.WList().Add().SetWName("global_unique_similar_obj")
	_, err := prefix.ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "global_unique_similar_obj already exists")

	// Two similar objects with different name
	prefix.WList().Items()[1].SetWName("global_unique_similar_obj1")
	_, err = prefix.ToJson()
	assert.Nil(t, err)

	// Two different objects with same name
	prefix.SetName("global_unique")
	prefix.WList().Add().SetWName("global_unique")
	_, err = prefix.ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "global_unique already exists")

	// Two different objects with different name
	prefix.SetName("global_unique1")
	_, err = prefix.ToJson()
	assert.Nil(t, err)
	// ********************************************

	// *************** local unique ****************

	// prefix.ZObject().SetName("local_unique")
	// Two similar objects with same Name.
	prefix.XList().Add().SetName("local_unique")
	prefix.XList().Add().SetName("local_unique")
	_, err = prefix.ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "local_unique already exists")

	// Two similar objects with different name
	prefix.XList().Items()[0].SetName("local_unique1")
	_, err = prefix.ToJson()
	assert.Nil(t, err)

	// Two different objects with same name
	prefix.SetName("local_global_mix")
	prefix.ZObject().SetName("local_global_mix")
	_, err = prefix.ToJson()
	assert.NotNil(t, err)
	// ************************************************

	// prefix.YObject().SetYName("123")
	// _, err = prefix.ToJson()
	// assert.NotNil(t, err)
	// assert.Contains(t, err.Error(), "123 is not a valid")

	// prefix.YObject().SetYName("local_unique")
	// prefix.SetName("global_and_local_same_name")
	// prefix.XList().Add().SetName("global_and_local_same_name")
	// data, err1 := prefix.ToJson()
	// assert.Nil(t, err1)
	// prefix1 := openapiart.NewPrefixConfig()
	// res := make(map[string]interface{})
	// err = json.Unmarshal([]byte(data), &res)
	// assert.Nil(t, err)
	// err = prefix1.FromJson(data)
	// assert.Nil(t, err)
}
