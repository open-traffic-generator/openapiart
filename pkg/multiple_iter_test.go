package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestMultipleIter(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetA("test")
	config.SetB(1.234)
	config.SetC(32)
	config.RequiredObject().SetEA(3.45).SetEB(6.78)
	enums := []openapiart.PrefixConfigDValuesEnum{
		openapiart.PrefixConfigDValues.A,
		openapiart.PrefixConfigDValues.B,
		openapiart.PrefixConfigDValues.C,
	}
	config.SetDValues(enums)
	config.SetStrLen("1234")
	config.G().Add().SetGA("this is g")
	assert.Equal(t, len(config.G().Items()), 1)
	config.G1().Add().SetGA("this is g1")
	assert.Equal(t, len(config.G1().Items()), 1)
	config.G2().Add().SetGA("this is g2")
	assert.Equal(t, len(config.G2().Items()), 1)
	t.Log(config)
	// validating the warnings
	err := config.Validate()
	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}
