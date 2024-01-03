package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestOptionalValSchema(t *testing.T) {

	// This test checks the values in optional schema.
	// Objective is to check if default values are set properly.

	config := openapiart.NewTestConfig()
	oVal := config.NativeFeatures().OptionalVal()
	assert.Equal(t, oVal.IntVal(), int32(50))
	assert.Equal(t, oVal.NumVal(), float32(50.05))
	assert.Equal(t, oVal.StrVal(), "default_str_val")
	assert.Equal(t, oVal.BoolVal(), true)

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestOptionalArrayValSchema(t *testing.T) {

	// This test checks the values in optional array schema.
	// Objective is to check if default values are set properly.

	config := openapiart.NewTestConfig()
	oVal := config.NativeFeatures().OptionalValArray()
	assert.Equal(t, oVal.IntVals(), []int32{10, 20})
	assert.Equal(t, oVal.NumVals(), []float32{10.01, 20.02})
	assert.Equal(t, oVal.StrVals(), []string{"first_str", "second_str"})
	assert.Equal(t, oVal.BoolVals(), []bool{})

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
}
