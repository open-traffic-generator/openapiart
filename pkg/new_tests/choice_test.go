package openapiart_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestChoiceValSchema(t *testing.T) {

	// This test checks the values in choice val schema.
	// Objective is to check if values are set properly.

	config := goapi.NewTestConfig()
	mVal := config.ExtendedFeatures().ChoiceVal().MixedVal()

	mVal.SetIntVal(34)

	assert.Equal(t, mVal.IntVal(), int32(34))
	assert.Equal(t, mVal.Choice(), goapi.MixedValChoice.INT_VAL)

	mVal.SetNumVal(50.05)
	assert.Equal(t, mVal.NumVal(), float32(50.05))
	assert.Equal(t, mVal.Choice(), goapi.MixedValChoice.NUM_VAL)

	mVal.SetStrVal("str1")
	assert.Equal(t, mVal.StrVal(), "str1")
	assert.Equal(t, mVal.Choice(), goapi.MixedValChoice.STR_VAL)

	mVal.SetBoolVal(true)
	assert.Equal(t, mVal.BoolVal(), true)
	assert.Equal(t, mVal.Choice(), goapi.MixedValChoice.BOOL_VAL)

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestChoiceHeirarchy(t *testing.T) {

	// This test checks choices at different heirarchy

	config := goapi.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	val.IntermediateObj().Leaf().SetName("str1").SetValue(3)

	assert.Equal(t, val.Choice(), goapi.ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	assert.Equal(t, val.IntermediateObj().Choice(), goapi.RequiredChoiceChoice.LEAF)
	assert.Equal(t, val.IntermediateObj().Leaf().Name(), "str1")
	assert.Equal(t, val.IntermediateObj().Leaf().Value(), int32(3))

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestChoiceWithNoProperties(t *testing.T) {

	// This test checks choices with no properties has no issues reagrding set and get

	config := goapi.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	val.NoObj()
	assert.Equal(t, val.Choice(), goapi.ChoiceValWithNoPropertiesChoice.NO_OBJ)

	_, err := config.Marshal().ToYaml()
	assert.Nil(t, err)

}

func TestChoiceWithRequiredFeild(t *testing.T) {

	// This set checks choices which are defined as required

	config := goapi.NewTestConfig()
	val := config.ExtendedFeatures().ChoiceValNoProperties()

	// check error
	_, err := config.Marshal().ToYaml()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Choice is required field on interface ChoiceValWithNoProperties")

	val.IntermediateObj().SetStrVal("str1")

	assert.Equal(t, val.Choice(), goapi.ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	assert.Equal(t, val.IntermediateObj().Choice(), goapi.RequiredChoiceChoice.STR_VAL)
	assert.Equal(t, val.IntermediateObj().StrVal(), "str1")

	_, err = config.Marshal().ToYaml()
	assert.Nil(t, err)

}
