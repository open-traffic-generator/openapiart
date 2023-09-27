package openapiart_test

import (
	"fmt"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChoiceWithNoPropertiesForLeafNode(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	fObj := config.F()

	// test default choice and values
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "some string")

	// setting of other choices should work as usual
	fObj.SetFB(5.67)
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_B)
	assert.True(t, fObj.HasFB())
	assert.Equal(t, fObj.FB(), 5.67)

	fObj.SetFA("str1")
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "str1")

	// setting choice with no property
	fObj.SetChoice(openapiart.FObjectChoice.F_C)
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_C)

	err := fObj.Validate()
	assert.Nil(t, err)
}

func TestChoiceWithNoPropertiesForIterNode(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.NO_OBJ)
	err := choiceObj.Validate()
	assert.Nil(t, err)

	// rest of operation should not be impacted
	assert.Contains(t, config.ChoiceObject().Items()[0].Choice(), "no_obj")
	assert.Len(t, config.ChoiceObject().Items(), 1)

	choiceObj.EObj().SetEA(1.23)
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.E_OBJ)

	choiceObj.FObj().SetFA("str1")
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.F_OBJ)

	config.ChoiceObject().Append(openapiart.NewChoiceObject())

	config.ChoiceObject().Set(1, openapiart.NewChoiceObject().SetChoice("e_obj"))
	assert.Len(t, config.ChoiceObject().Items(), 2)

	config.ChoiceObject().Clear()
	assert.Len(t, config.ChoiceObject().Items(), 0)
}

func TestChoiceWithNoPropertiesForChoiceHeirarchy(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	choiceObj := config.ChoiceObject().Add()

	// check default should be no_obj
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.NO_OBJ)
	err := choiceObj.Validate()
	assert.Nil(t, err)

	fObj := choiceObj.FObj()

	// check default for child obj
	assert.Equal(t, choiceObj.Choice(), openapiart.ChoiceObjectChoice.F_OBJ)
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_A)
	assert.True(t, fObj.HasFA())
	assert.Equal(t, fObj.FA(), "some string")

	// set choice with no properties in child obj
	fObj.SetChoice(openapiart.FObjectChoice.F_C)
	assert.Equal(t, fObj.Choice(), openapiart.FObjectChoice.F_C)
	assert.False(t, fObj.HasFA())
	assert.False(t, fObj.HasFB())

	// validate the whole object
	err = choiceObj.Validate()
	assert.Nil(t, err)
}

func TestChoiceMarshall(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asd").RequiredObject()

	choiceObj := config.ChoiceTest()
	choiceObj.EObj().SetEA(1.23).SetEB(2.34)
	s, err := config.ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json := `{
			"required_object": {},
			"response": "status_200",
			"a": "asd",
			"h": true,
			"choice_test": {
			  "choice": "e_obj",
			  "e_obj": {
				"e_a": 1.23,
				"e_b": 2.34
			  }
			}
		  }`
	require.JSONEq(t, exp_json, s)

	choiceObj.FObj().SetFA("s1").SetFB(34.5678)
	s, err = config.ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json = `{
			"required_object": {},
			"response": "status_200",
			"a": "asd",
			"h": true,
			"choice_test": {
			  "choice": "f_obj",
			  "f_obj": {
				"choice": "f_b",
				"f_b": 34.5678
			  }
			}
		  }`
	require.JSONEq(t, exp_json, s)

	choiceObj.SetChoice(openapiart.ChoiceTestObjChoice.NO_OBJ)
	s, err = config.ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json = `{
			"required_object": {},
			"response": "status_200",
			"a": "asd",
			"h": true,
			"choice_test": {
			  "choice": "no_obj"
			}
		  }`
	require.JSONEq(t, exp_json, s)
}

func TestChoiceUnMarshall(t *testing.T) {
	exp_json := `{
			"a": "asd",
			"required_object": {},
			"choice_test": {
				"choice": "e_obj",
				"e_obj": {
					"e_a": 1.23,
					"e_b": 22.3456
				}
			}
		}`

	api := openapiart.NewApi()
	config := api.NewPrefixConfig()
	err := config.FromJson(exp_json)
	assert.Nil(t, err)
	cObj := config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.E_OBJ)
	assert.Equal(t, cObj.EObj().EA(), float32(1.23))
	assert.Equal(t, cObj.EObj().EB(), float64(22.3456))

	exp_json = `{
			"a": "asd",
			"required_object": {},
			"choice_test": {
				"choice": "f_obj",
				"f_obj": {
					"f_a": "s1",
					"f_b": 22.3456
				}
			}
		}`

	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.F_OBJ)

	exp_json = `{
			"a": "asd",
			"required_object": {},
			"choice_test": {
				"choice": "no_obj"
			}
		}`

	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.NO_OBJ)

	// json without choice
	json := `{
		"a": "asd",
		"required_object": {},
		"choice_test": {
			"e_obj": {
				"e_a": 1.23,
				"e_b": 22.3456
			}
		}
	}`

	config = api.NewPrefixConfig()
	err = config.FromJson(json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	fmt.Println(config)
	assert.Equal(t, cObj.EObj().EA(), float32(1.23))
	assert.Equal(t, cObj.EObj().EB(), float64(22.3456))
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.E_OBJ)

	//json without choice in hierarchy with non-primitive type
	exp_json = `{
		"a": "asd",
		"required_object": {},
		"required_choice_object": {
			"intermediate_obj": {
				"leaf": {
					"name": "name1"
				}
			}
		}
	}`

	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	r := config.RequiredChoiceObject()
	fmt.Println(r)
	assert.Equal(t, r.Choice(), openapiart.RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	ir := r.IntermediateObj()
	assert.Equal(t, ir.Choice(), openapiart.RequiredChoiceIntermediateChoice.LEAF)
	assert.Equal(t, ir.Leaf().Name(), "name1")

	// json without choice in hierarchy with primitive type
	exp_json = `{
		"a": "asd",
		"required_object": {},
		"required_choice_object": {
			"intermediate_obj": {
				"f_a": "name1"
			}
		}
	}`

	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	r = config.RequiredChoiceObject()
	assert.Equal(t, r.Choice(), openapiart.RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	ir = r.IntermediateObj()
	assert.Equal(t, ir.Choice(), openapiart.RequiredChoiceIntermediateChoice.F_A)
	assert.Equal(t, ir.FA(), "name1")

	// json without choice for checksum pattern with enum choice properties
	config = api.NewPrefixConfig()
	config.SetA("asd").RequiredObject()
	config.HeaderChecksum().SetCustom(123)
	fmt.Println(config)
	exp_json = `{
		"a": "asd",
		"required_object": {},
		"header_checksum": {
			"custom": 12345,
			"generated": "unspecified"
		}
	}`
	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	hc := config.HeaderChecksum()
	assert.Equal(t, hc.Choice(), openapiart.PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	assert.Equal(t, hc.Custom(), uint32(12345))

	// json without choice for checksum pattern with enum choice properties
	exp_json = `{
		"a": "asd",
		"required_object": {},
		"header_checksum": {
			"generated": "bad"
		}
	}`
	config = api.NewPrefixConfig()
	err = config.FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	hc = config.HeaderChecksum()
	assert.Equal(t, hc.Choice(), openapiart.PatternPrefixConfigHeaderChecksumChoice.GENERATED)
	assert.Equal(t, hc.Generated(), openapiart.PatternPrefixConfigHeaderChecksumGenerated.BAD)

}
