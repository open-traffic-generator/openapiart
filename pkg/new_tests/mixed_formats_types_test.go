package openapiart_test

import (
	"log"
	"strings"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFormatsSanity(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	mixedVal.SetStringParam("asdf")
	mixedVal.SetInteger(88)
	mixedVal.SetFloat(22.3)
	mixedVal.SetDouble(2342.222)
	mixedVal.SetMac("00:00:fa:ce:fa:ce")
	mixedVal.SetIpv4("1.1.1.1")
	mixedVal.SetIpv6("::02")
	mixedVal.SetHex("0102030405060708090a0b0c0d0e0f")
	mixedVal.SetInteger641(9223372036854775807)
	log.Print(mixedVal.Integer641())
	mixedVal.SetInteger642(4261412864)
	log.Print(mixedVal.Integer642())
	mixedVal.SetInteger64List([]int64{4261412864, 2})
	log.Print(mixedVal.Integer64List())
	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
	configJson, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	log.Print(configJson)
}

var GoodIpv4 = []string{"1.1.1.1", "255.255.255.255"}
var BadIpv4 = []string{"1.1. 1.1", "33.4", "asdf", "100", "-20", "::01", "1.1.1.1.1", "256.256.256.256", "-255.-255.-255.-255"}

func TestGoodIpv4Validation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	for _, ip := range GoodIpv4 {
		ipv4 := mixedVal.SetIpv4(ip)
		_, err := ipv4.Marshal().ToYaml()
		assert.Nil(t, err)
	}
}

func TestBadIpv4Validation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	for _, ip := range BadIpv4 {
		ipv4 := mixedVal.SetIpv4(ip)
		_, err := ipv4.Marshal().ToYaml()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Ipv4")
		}
	}
}

var GoodIpv6 = []string{"::", "1::", "abcd::1234", "aa:00bd:a:b:c:d:f:abcd"}
var BadIpv6 = []string{"33.4", "asdf", "1.1.1.1", "100", "-20", "65535::65535", "ab: :ab", "ab:ab:ab", "ffff0::ffff0"}

func TestGoodIpv6Validation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	for _, ip := range GoodIpv6 {
		ipv6 := mixedVal.SetIpv6(ip)
		_, err := ipv6.Marshal().ToYaml()
		assert.Nil(t, err)
	}
}

func TestBadIpv6Validation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	for _, ip := range BadIpv6 {
		ipv6 := mixedVal.SetIpv6(ip)
		_, err := ipv6.Marshal().ToYaml()
		if assert.Error(t, err) {
			assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
		}
	}
}

var GoodMac = []string{"ab:ab:10:12:ff:ff"}
var BadMac = []string{
	"1", "2.2", "1.1.1.1", "::01", "00:00:00", "00:00:00:00:gg:00", "00:00:fa:ce:fa:ce:01", "255:255:255:255:255:255",
}

func TestGoodMacValidation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	mac := mixedVal.SetMac(GoodMac[0])
	_, err := mac.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestBadMacValidation(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	for _, mac := range BadMac {
		macObj := mixedVal.SetMac(mac)
		_, err := macObj.Marshal().ToYaml()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Mac")
		}
	}
}

func TestGoodHex(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	mixedVal.SetHex("200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	_, err := mixedVal.Marshal().ToYaml()
	assert.Nil(t, err)
	mixedVal.SetHex("0x00200000000000000b00000000200000000000000b00000000200000000000000b00000000")
	_, err1 := mixedVal.Marshal().ToYaml()
	assert.Nil(t, err1)
	mixedVal.SetHex("")
	_, err2 := mixedVal.Marshal().ToYaml()
	assert.NotNil(t, err2)
	mixedVal.SetHex("0x00200000000000000b00000000200000000000000b00000000200000000000000b0000000x0")
	_, err3 := mixedVal.Marshal().ToYaml()
	assert.NotNil(t, err3)
	mixedVal.SetHex("0x00")
	_, err4 := mixedVal.Marshal().ToYaml()
	assert.Nil(t, err4)
	mixedVal.SetHex("0XAF12")
	_, err5 := mixedVal.Marshal().ToYaml()
	assert.Nil(t, err5)
}

func TestBadHex(t *testing.T) {

	config := openapiart.NewTestConfig()
	mixedVal := config.NativeFeatures().MixedObject()
	mixedVal.SetHex("1.1.1.1")
	_, err := mixedVal.Marshal().ToYaml()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid hex")
	}
	mixedVal.SetHex("::01")
	_, err1 := mixedVal.Marshal().ToYaml()
	if assert.Error(t, err1) {
		assert.Contains(t, err1.Error(), "Invalid hex")
	}
	mixedVal.SetHex("00:00:fa:ce:fa:ce:01")
	_, err2 := mixedVal.Marshal().ToYaml()
	if assert.Error(t, err2) {
		assert.Contains(t, err2.Error(), "Invalid hex")
	}
}

// The below test needs tobe revistited once the proper support of data types is added in go
func TestIntegerDatatypes(t *testing.T) {

	config := openapiart.NewTestConfig()
	numberTypeVal := config.NativeFeatures().NumberTypeObject()
	numberTypeVal.SetValidateUint321(2147483646)
	log.Print(numberTypeVal.ValidateUint321())
	numberTypeVal.SetValidateUint322(2147483646)
	log.Print(numberTypeVal.ValidateUint322())
	numberTypeVal.SetValidateUint641(4261412865)
	log.Print(numberTypeVal.ValidateUint641())
	numberTypeVal.SetValidateUint642(4261412865)
	log.Print(numberTypeVal.ValidateUint642())
	numberTypeVal.SetValidateInt321(2147483646)
	log.Print(numberTypeVal.ValidateInt321())
	numberTypeVal.SetValidateInt322(2147483646)
	log.Print(numberTypeVal.ValidateInt322())
	numberTypeVal.SetValidateInt641(4261412865)
	log.Print(numberTypeVal.ValidateInt641())
	numberTypeVal.SetValidateInt642(4261412865)
	log.Print(numberTypeVal.ValidateInt642())

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
	configJson, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	log.Print(configJson)
}
