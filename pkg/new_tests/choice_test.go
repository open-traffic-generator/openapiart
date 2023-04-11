package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestChoiceValSchema(t *testing.T) {

	// This test checks the values in choice val schema.
	// Objective is to check if values are set properly.

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	mVal := config.ExtendedFeatures().ChoiceVal().MixedVal()

	mVal.SetIntVal(34)

	assert.Equal(t, mVal.IntVal(), int32(34))
	assert.Equal(t, mVal.Choice(), openapiart.MixedValChoice.INT_VAL)

	mVal.SetNumVal(50.05)
	assert.Equal(t, mVal.NumVal(), float32(50.05))
	assert.Equal(t, mVal.Choice(), openapiart.MixedValChoice.NUM_VAL)

	mVal.SetStrVal("str1")
	assert.Equal(t, mVal.StrVal(), "str1")
	assert.Equal(t, mVal.Choice(), openapiart.MixedValChoice.STR_VAL)

	mVal.SetBoolVal(true)
	assert.Equal(t, mVal.BoolVal(), true)
	assert.Equal(t, mVal.Choice(), openapiart.MixedValChoice.BOOL_VAL)

	err := config.Validate()
	assert.Nil(t, err)
}

func TestChoiceHeirarchy(t *testing.T) {

	// This test checks choices at different heirarchy

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	val.IntermediateObj().Leaf().SetName("str1").SetValue(3)

	assert.Equal(t, val.Choice(), openapiart.ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	assert.Equal(t, val.IntermediateObj().Choice(), openapiart.RequiredChoiceChoice.LEAF)
	assert.Equal(t, val.IntermediateObj().Leaf().Name(), "str1")
	assert.Equal(t, val.IntermediateObj().Leaf().Value(), int32(3))

	err := config.Validate()
	assert.Nil(t, err)
}

func TestChoiceWithNoProperties(t *testing.T) {

	// This test checks choices with no properties has no issues reagrding set and get

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	val.SetChoice(openapiart.ChoiceValWithNoPropertiesChoice.NO_OBJ)
	assert.Equal(t, val.Choice(), openapiart.ChoiceValWithNoPropertiesChoice.NO_OBJ)

	err := config.Validate()
	assert.Nil(t, err)

}

func TestChoiceWithRequiredFeild(t *testing.T) {

	// This set checks choices which are defined as required

	api := openapiart.NewApi()
	config := api.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	// check error
	err := config.Validate()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Choice is required field on interface ChoiceValWithNoProperties")

	val.IntermediateObj().SetStrVal("str1")

	assert.Equal(t, val.Choice(), openapiart.ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	assert.Equal(t, val.IntermediateObj().Choice(), openapiart.RequiredChoiceChoice.STR_VAL)
	assert.Equal(t, val.IntermediateObj().StrVal(), "str1")

	err = config.Validate()
	assert.Nil(t, err)

}
