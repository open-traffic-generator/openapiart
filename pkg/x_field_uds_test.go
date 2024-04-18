package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFieldUdsDefault(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	m := config.FieldUdsMac()
	assert.Equal(t, m.Mac().Value(), "00:00:00:00:00:00")
	assert.Equal(t, m.Mac().Mask(), "ffffffffffff")
	_, err := m.Marshal().ToJson()
	assert.Nil(t, err)

	v4 := config.FieldUdsIpv4()
	assert.Equal(t, v4.Ipv4().Value(), "0.0.0.0")
	assert.Equal(t, v4.Ipv4().Mask(), "ffffffff")
	_, err = v4.Marshal().ToJson()
	assert.Nil(t, err)

	v6 := config.FieldUdsIpv6()
	assert.Equal(t, v6.Ipv6().Value(), "::0")
	assert.Equal(t, v6.Ipv6().Mask(), "ffffffffffffffffffffffffffffffff")
	_, err = v6.Marshal().ToJson()
	assert.Nil(t, err)

	i := config.FieldUdsInt()
	assert.Equal(t, i.Integer().Value(), uint32(0))
	assert.Equal(t, i.Integer().Mask(), "ff")
	_, err = i.Marshal().ToJson()
	assert.Nil(t, err)

}

func TestFieldUdsErrors(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	m := config.FieldUdsMac()
	m.Mac().SetMask("ffffffffffffffffffff")
	_, err := m.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "1 <= length of FilterMacUdsPatternMac.Mask <= 12 but Got 20")

	v4 := config.FieldUdsIpv4()
	v4.Ipv4().SetMask("")
	_, err = v4.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "1 <= length of FilterIpv4UdsPatternIpv4.Mask <= 8 but Got 0")

	v6 := config.FieldUdsIpv6()
	v6.Ipv6().SetValue("::3")
	_, err = v6.Marshal().ToJson()
	assert.Nil(t, err)

	i := config.FieldUdsInt()
	i.Integer().SetValue(456)
	i.Integer().SetMask("fff")
	_, err = i.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "0 <= FilterIntegerUdsPatternInteger.Value <= 3 but Got 456")
	assert.Contains(t, err.Error(), "1 <= length of FilterIntegerUdsPatternInteger.Mask <= 2 but Got 3")
}
