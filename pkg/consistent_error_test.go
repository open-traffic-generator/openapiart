package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestRequiredValidationErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	config.SetB(1.23)
	config.SetC(32)
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "RequiredObject is required field on interface PrefixConfig")
	assert.Contains(t, err.Error(), "A is required field on interface PrefixConfig")
}

func TestMinMaxValidationErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	config.SetB(1.23)
	config.SetC(32)
	config.RequiredObject()
	config.SetA("asd")
	config.SetStrLen("asdasdasdasdasdasd")
	vals := []int64{-101, 100, -202, 200}
	config.SetInteger64List(vals)
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "3 <= length of PrefixConfig.StrLen <= 6 but Got 18")
	assert.Contains(t, err.Error(), "-12 <= PrefixConfig.Integer64List <= 4261412864 but Got -101")
}

func TestSpecialFormatErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetB(1.23)
	config.SetC(32)
	config.RequiredObject()
	config.SetA("asd")
	config.L().SetIpv4("1222.2.3.4")
	config.L().SetIpv6("2::::2")
	config.L().SetMac("00:00:123:00:00:12")
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid Mac address at 2nd octet in 00:00:123:00:00:12 mac on PrefixConfig.L.Mac")
	assert.Contains(t, err.Error(), "Invalid Ipv4 address at 1st octet in 1222.2.3.4 ipv4 on PrefixConfig.L.Ipv4")
	assert.Contains(t, err.Error(), "Invalid ipv6 address 2::::2 on PrefixConfig.L.Ipv6")
}

func TestPatternErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetB(1.23)
	config.SetC(32)
	config.RequiredObject()
	config.SetA("asd")
	config.Ipv4Pattern().Ipv4().Increment().SetStart("1.2.3")
	config.MacPattern().Mac().Decrement().SetStep("00:34:45:222:11:01")
	config.IntegerPattern().Integer().Increment().SetStep(0)
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid Ipv4 address 1.2.3 on PrefixConfig.Ipv4Pattern.Ipv4.Increment.Start")
	assert.Contains(t, err.Error(), "Invalid Mac address at 3rd octet in 00:34:45:222:11:01 mac on PrefixConfig.MacPattern.Mac.Decrement.Step")
}

func TestIteratorErrors(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetB(1.23)
	config.SetC(32)
	config.RequiredObject()
	config.SetA("asd")
	config.XList().Add()
	config.XList().Add().SetName("Item2")
	config.XList().Add()
	config.WList().Add().SetWName("Item1")
	config.WList().Add()
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "WName is required field on interface PrefixConfig.WList[1]")
	assert.Contains(t, err.Error(), "Name is required field on interface PrefixConfig.XList[0]")
	assert.Contains(t, err.Error(), "Name is required field on interface PrefixConfig.XList[2]")
}
