package goapi_test

import (
	"fmt"
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestChoiceMarshall(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetA("asd").SetB(1).SetC(22).RequiredObject().SetEA(1.23).SetEB(2.34)

	choiceObj := config.ChoiceTest()
	choiceObj.EObj().SetEA(1.23).SetEB(2.34)
	s, err := config.Marshal().ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json := `{
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"response": "status_200",
			"a": "asd",
			"b": 1,
			"c": 22,
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
	s, err = config.Marshal().ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json = `{
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"response": "status_200",
			"a": "asd",
			"b": 1,
			"c": 22,
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

	choiceObj.NoObj()
	s, err = config.Marshal().ToJson()
	fmt.Println(s)
	assert.Nil(t, err)
	exp_json = `{
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"response": "status_200",
			"a": "asd",
			"b": 1,
			"c": 22,
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
			"b": 1,
			"c": 22,
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"choice_test": {
				"choice": "e_obj",
				"e_obj": {
					"e_a": 1.23,
					"e_b": 22.3456
				}
			}
		}`

	config := openapiart.NewPrefixConfig()
	err := config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	cObj := config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.E_OBJ)
	assert.Equal(t, cObj.EObj().EA(), float32(1.23))
	assert.Equal(t, cObj.EObj().EB(), float64(22.3456))

	exp_json = `{
			"a": "asd",
			"b": 1,
			"c": 22,
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"choice_test": {
				"choice": "f_obj",
				"f_obj": {
					"f_a": "s1",
					"f_b": 22.3456
				}
			}
		}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.F_OBJ)

	exp_json = `{
			"a": "asd",
			"b": 1,
			"c": 22,
			"required_object": {
				"e_a": 1.23,
				"e_b": 2.34
			},
			"choice_test": {
				"choice": "no_obj"
			}
		}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.NO_OBJ)

	// json without choice
	json := `{
		"a": "asd",
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"choice_test": {
			"e_obj": {
				"e_a": 1.23,
				"e_b": 22.3456
			}
		}
	}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	fmt.Println(config)
	assert.Equal(t, cObj.EObj().EA(), float32(1.23))
	assert.Equal(t, cObj.EObj().EB(), float64(22.3456))
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.E_OBJ)

	// json without choice
	json = `{
		"a": "asd",
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"choice_test": {
			"ieee_802_1qbb": "hello!"
		}
	}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(json)
	assert.Nil(t, err)
	cObj = config.ChoiceTest()
	fmt.Println(config)
	assert.Equal(t, cObj.Choice(), openapiart.ChoiceTestObjChoice.IEEE_802_1QBB)
	assert.Equal(t, cObj.Ieee8021Qbb(), "hello!")

	//json without choice in hierarchy with non-primitive type
	exp_json = `{
		"a": "asd",
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"required_choice_object": {
			"intermediate_obj": {
				"leaf": {
					"name": "name1"
				}
			}
		}
	}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
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
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"required_choice_object": {
			"intermediate_obj": {
				"f_a": "name1"
			}
		}
	}`

	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	r = config.RequiredChoiceObject()
	assert.Equal(t, r.Choice(), openapiart.RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	ir = r.IntermediateObj()
	assert.Equal(t, ir.Choice(), openapiart.RequiredChoiceIntermediateChoice.F_A)
	assert.Equal(t, ir.FA(), "name1")

	// json without choice for checksum pattern with enum choice properties
	config = openapiart.NewPrefixConfig()
	config.SetA("asd").RequiredObject()
	config.HeaderChecksum().SetCustom(123)
	fmt.Println(config)
	exp_json = `{
		"a": "asd",
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"header_checksum": {
			"custom": 12345,
			"generated": "unspecified"
		}
	}`
	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	hc := config.HeaderChecksum()
	assert.Equal(t, hc.Choice(), openapiart.PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	assert.Equal(t, hc.Custom(), uint32(12345))

	// json without choice for checksum pattern with enum choice properties
	exp_json = `{
		"a": "asd",
		"b": 1,
		"c": 22,
		"required_object": {
			"e_a": 1.23,
			"e_b": 2.34
		},
		"header_checksum": {
			"generated": "bad"
		}
	}`
	config = openapiart.NewPrefixConfig()
	err = config.Unmarshal().FromJson(exp_json)
	assert.Nil(t, err)
	fmt.Println(config)
	hc = config.HeaderChecksum()
	assert.Equal(t, hc.Choice(), openapiart.PatternPrefixConfigHeaderChecksumChoice.GENERATED)
	assert.Equal(t, hc.Generated(), openapiart.PatternPrefixConfigHeaderChecksumGenerated.BAD)

}

func TestDefaultChoiceOverwrite(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	crd := config.ChoiceRequiredDefault()
	crd.Ipv4()

	// test default choice and values
	fmt.Println(crd)
	assert.Equal(t, crd.Choice(), openapiart.ChoiceRequiredAndDefaultChoice.IPV4)
	assert.True(t, crd.HasIpv4())
	assert.Equal(t, crd.Ipv4(), "0.0.0.0")

	// setting of other choices should work as usual
	crd.SetIpv6([]string{"1::2"})

	c, err := crd.Marshal().ToYaml()
	fmt.Println(c)
	assert.Nil(t, err)

	assert.Equal(t, crd.Choice(), openapiart.ChoiceRequiredAndDefaultChoice.IPV6)
	assert.True(t, len(crd.Ipv6()) > 0)
	assert.Equal(t, crd.Ipv6()[0], "1::2")

	crd.SetIpv4("1.2.3.4")

	c, err = crd.Marshal().ToYaml()
	fmt.Println(c)
	assert.Nil(t, err)

	assert.Equal(t, crd.Choice(), openapiart.ChoiceRequiredAndDefaultChoice.IPV4)
	assert.True(t, crd.HasIpv4())
	assert.Equal(t, crd.Ipv4(), "1.2.3.4")
}
