package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestOid(t *testing.T) {
	config := goapi.NewPrefixConfig()
	m := config.MObject()
	m.SetDouble(1.23)
	m.SetFloat(3.45)
	m.SetHex("0f")
	m.SetIpv4("1.2.3.4")
	m.SetIpv6("::")
	m.SetMac("00:00:00:00:00:00")
	m.SetStringParam("abcd")
	m.SetInteger(34)

	m.SetOid("1.abc")
	_, err := m.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid oid value 1.abc on MObject.Oid")

	m.SetOid("1.")
	_, err = m.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid oid value 1. on MObject.Oid")

	m.SetOid("1.-1.33.44.5678.9876")
	_, err = m.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid oid value 1.-1.33.44.5678.9876 on MObject.Oid")

	m.SetOid("1.2.3.4.5.6.7")
	_, err = m.Marshal().ToJson()
	assert.Nil(t, err)
}

func TestOidSlice(t *testing.T) {
	config := goapi.NewPrefixConfig()
	oid := config.OidPattern().Oid()

	oid.SetValues([]string{"1.2.3.4", "3.4.5.6"})
	_, err := oid.Marshal().ToJson()
	assert.Nil(t, err)

	oid.SetValues([]string{"1.2.3.4", "3.4.5.6", "-1.3.4.5", "1", ".", "11111.33333", "abcd.23", "1.2.3.4294967298"})
	_, err = oid.Marshal().ToJson()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Invalid oid addresses at indices 2,3,4,6,7 on PatternOidPatternOid.Values")
}
