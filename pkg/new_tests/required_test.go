package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestRequiredValSchema(t *testing.T) {

	// This test checks the values in required schema.
	// Objective is validation should not a problem.

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	rVal := config.NativeFeatures().RequiredVal()
	rVal.SetIntVal(40)
	rVal.SetNumVal(5.67)
	rVal.SetBoolVal(false)
	rVal.SetStrVal("str1")

	err := config.Validate()
	assert.Nil(t, err)
}

func TestRequiredErr(t *testing.T) {

	// This test checks error returned by SDK for required properties.
	// Objective is that the SDK return proper validation error.

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().RequiredVal()

	err := config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "StrVal is required field on interface RequiredVal\nvalidation errors")
}

func TestRequiredArraySchema(t *testing.T) {

	// This test checks the values in required array schema.
	// Objective is validation should not a problem.

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	rVal := config.NativeFeatures().RequiredValArray()
	rVal.SetIntVals([]int32{40, 50})
	rVal.SetNumVals([]float32{5.67})
	rVal.SetBoolVals([]bool{false, true})
	rVal.SetStrVals([]string{"s1", "s2"})

	err := config.Validate()
	assert.Nil(t, err)

}

func TestRequiredArrayErr(t *testing.T) {

	// This test checks error returned by SDK for required array properties.
	// Objective is that the SDK return proper validation error.

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	config.NativeFeatures().RequiredValArray()

	err := config.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "StrVals is required field on interface RequiredValArray\nvalidation errors")

}
