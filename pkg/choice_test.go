package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
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
}
