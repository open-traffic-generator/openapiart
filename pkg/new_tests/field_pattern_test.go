package openapiart_test

import (
	"fmt"
	"strings"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

var ValidMac = []string{"ab:ab:10:12:ff:ff"}
var InvalidMac = []string{
	"1", "2.2", "1.1.1.1", "::01", "00:00:00", "00:00:00:00:gg:00", "00:00:fa:ce:fa:ce:01", "255:255:255:255:255:255",
}

func TestValidMacPattern(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	mac := config.MacPattern().Mac().SetValue(GoodMac[0])
	err := mac.Validate()
	assert.Nil(t, err)
}

func TestInValidMacPattern(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	for _, mac := range InvalidMac {
		macObj := config.MacPattern().Mac().SetValue(mac)
		err := macObj.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Mac")
		}
	}
}

func TestValidMacValues(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	mac := config.MacPattern().Mac().SetValues(ValidMac)
	err := mac.Validate()
	assert.Nil(t, err)
}

func TestInValidMacValues(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	mac := config.MacPattern().Mac().SetValues(InvalidMac)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

func TestInValidMacIncrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	mac := config.MacPattern().Mac().Increment().SetStart(ValidMac[0])
	mac.SetStep(InvalidMac[0])
	mac.SetCount(10)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

func TestInValidMacDecrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	mac := config.MacPattern().Mac().Decrement().SetStart(InvalidMac[0])
	mac.SetStep(ValidMac[0])
	mac.SetCount(10)
	err := mac.Validate()
	fmt.Println(err.Error())
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid mac address")
	}
}

var ValidIpv4 = []string{"1.1.1.1", "255.255.255.255"}
var InValidIpv4 = []string{"1.1. 1.1", "33.4", "asdf", "100", "-20", "::01", "1.1.1.1.1", "256.256.256.256", "-255.-255.-255.-255"}

func TestValidIpv4Value(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv4 := config.Ipv4Pattern().Ipv4().SetValue(ValidIpv4[0])
	err := ipv4.Validate()
	assert.Nil(t, err)
}

func TestInValidIpv4Value(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	for _, ip := range InValidIpv4 {
		ipv4 := config.Ipv4Pattern().Ipv4().SetValue(ip)
		err := ipv4.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "Invalid Ipv4")
		}
	}
}

func TestInValidIpv4Values(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv4 := config.Ipv4Pattern().Ipv4().SetValues(InValidIpv4)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid ipv4 addresses")
	}
}

func TestInValidIpv4Increment(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().Increment().SetStart(ValidIpv4[0])
	ipv4.SetStep(InValidIpv4[0])
	ipv4.SetCount(10)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid Ipv4")
	}
}

func TestInValidIpv4Decrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	ipv4 := config.Ipv4Pattern().Ipv4().Decrement().SetStart(ValidIpv4[0])
	ipv4.SetStep(InValidIpv4[0])
	ipv4.SetCount(10)
	err := ipv4.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "Invalid Ipv4")
	}
}

var ValidIpv6 = []string{"::", "1::", ": :", "abcd::1234", "aa:00bd:a:b:c:d:f:abcd"}
var InValidIpv6 = []string{"33.4", "asdf", "1.1.1.1", "100", "-20", "65535::65535", "ab: :ab", "ab:ab:ab", "ffff0::ffff0"}

func TestValidIpv6Value(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv6 := config.Ipv6Pattern().Ipv6().SetValue(ValidIpv6[0])
	err := ipv6.Validate()
	assert.Nil(t, err)
}

func TestInValidIpv6Value(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	for _, ip := range InValidIpv6 {
		ipv6 := config.Ipv6Pattern().Ipv6().SetValue(ip)
		err := ipv6.Validate()
		if assert.Error(t, err) {
			assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
		}
	}
}

func TestInvalidIpv6Values(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv6 := config.Ipv6Pattern().Ipv6().SetValues(InValidIpv6)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6 address")
	}
}

func TestInValidIpv6Increment(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv6 := config.Ipv6Pattern().Ipv6().Increment().SetStart(ValidIpv6[0])
	ipv6.SetStep(InValidIpv6[0])
	ipv6.SetCount(10)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
	}
}

func TestInValidIpv6Decrement(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewTestConfig().ExtendedFeatures().XFieldPatternObject()
	ipv6 := config.Ipv6Pattern().Ipv6().Decrement().SetStart(ValidIpv6[0])
	ipv6.SetStep(InValidIpv6[0])
	ipv6.SetCount(10)
	err := ipv6.Validate()
	if assert.Error(t, err) {
		assert.Contains(t, strings.ToLower(err.Error()), "invalid ipv6")
	}
}
