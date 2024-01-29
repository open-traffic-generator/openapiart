package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestChoiceWithNoPropertiesForLeafNode(t *testing.T) {

	config := goapi.NewPrefixConfig()
	fObj := config.F()

	// test default choice and values
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "some string")

	// setting of other choices should work as usual
	fObj.SetFB(5.67)
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_B)
	assert.True(t, fObj.HasFB())
	assert.Equal(t, fObj.FB(), 5.67)

	fObj.SetFA("str1")
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "str1")

	// setting choice with no property
	fObj.FC()
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_C)

	_, err := fObj.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestChoiceWithNoPropertiesForIterNode(t *testing.T) {
	config := goapi.NewPrefixConfig()

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.NO_OBJ)
	_, err := choiceObj.Marshal().ToYaml()
	assert.Nil(t, err)

	// rest of operation should not be impacted
	assert.Contains(t, config.ChoiceObject().Items()[0].Choice(), "no_obj")
	assert.Len(t, config.ChoiceObject().Items(), 1)

	choiceObj.EObj().SetEA(1.23)
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.E_OBJ)

	choiceObj.FObj().SetFA("str1")
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.F_OBJ)

	config.ChoiceObject().Append(goapi.NewChoiceObject())

	chObj := goapi.NewChoiceObject()
	chObj.EObj()
	config.ChoiceObject().Set(1, chObj)
	assert.Len(t, config.ChoiceObject().Items(), 2)

	config.ChoiceObject().Clear()
	assert.Len(t, config.ChoiceObject().Items(), 0)
}

func TestChoiceWithNoPropertiesForChoiceHeirarchy(t *testing.T) {
	config := goapi.NewPrefixConfig()

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.NO_OBJ)
	_, err := choiceObj.Marshal().ToYaml()
	assert.Nil(t, err)

	fObj := choiceObj.FObj()

	// check default for child obj
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.F_OBJ)
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "some string")

	// set choice with no properties in child obj
	fObj.FC()
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_C)
	assert.False(t, fObj.HasFA())
	assert.False(t, fObj.HasFB())

	// validate the whole object
	_, err = choiceObj.Marshal().ToYaml()
	assert.Nil(t, err)
}

func TestChoiceWithNoPropertiesForChoiceDefault(t *testing.T) {
	config := goapi.NewPrefixConfig()

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.NO_OBJ)
	_, err := choiceObj.Marshal().ToYaml()
	assert.Nil(t, err)

	fObj := choiceObj.FObj()

	// check default for child obj
	assert.Equal(t, choiceObj.Choice(), goapi.ChoiceObjectChoice.F_OBJ)
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "some string")

	// set choice with no properties in child obj
	fObj.FC()
	assert.Equal(t, fObj.Choice(), goapi.FObjectChoice.F_C)
	assert.False(t, fObj.HasFA())
	assert.False(t, fObj.HasFB())

	// validate the whole object
	_, err = choiceObj.Marshal().ToYaml()
	assert.Nil(t, err)
}
