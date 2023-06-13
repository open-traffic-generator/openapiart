package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"

	"github.com/stretchr/testify/assert"
)

func TestPropertyWithNameString(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	config.SetStringVal("abc")
	assert.Equal(t, config.HasStringVal(), true)
	assert.Equal(t, config.StringVal(), "abc")

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.NO_OBJ)

	// check functionality of string property as choice
	choiceObj.SetStringVal("abc")
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.STRING)
	assert.Equal(t, choiceObj.HasStringVal(), true)
	assert.Equal(t, choiceObj.StringVal(), "abc")

	err := choiceObj.Validate()
	assert.Nil(t, err)
}
