package openapiart_test

import (
	"fmt"
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestStatuswarningInPrimitiveAttrs(t *testing.T) {
	config := openapiart.NewTestConfig()

	// setting all the primitive values which has x-status set
	xSt := config.ExtendedFeatures().XStatusObject()
	xSt.SetDecprecatedProperty2(45).SetUnderReviewProperty2(65).SetEnumProperty(openapiart.XStatusObjectEnumProperty.DECPRECATED_PROPERTY_1)
	// validating the warnings
	_, err := xSt.Marshal().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}

	warns := xSt.Warnings()
	assert.Equal(t, len(warns), 3)
	assert.Equal(t, warns[1], "DecprecatedProperty_2 property in schema XStatusObject is deprecated, test deprecated")
	assert.Equal(t, warns[2], "UnderReviewProperty_2 property in schema XStatusObject is under review, test under_review")
}

func TestStatusWarningsInStructAttrs(t *testing.T) {
	config := openapiart.NewTestConfig()

	// setting all the primitive values which has x-status set
	eF := config.ExtendedFeatures()
	eF.XStatusObject().SetEnumProperty(openapiart.XStatusObjectEnumProperty.DECPRECATED_PROPERTY_1)

	// validating the warnings
	_, err := config.Marshal().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := eF.Warnings()
	fmt.Println(warns)
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "XStatusObject property in schema ExtendedFeatures is under review, test under_review")
}

func TestStatusWarningsInXEnumAttrs(t *testing.T) {
	config := openapiart.NewTestConfig()

	// setting all the primitive values which has x-status set
	xSt := config.ExtendedFeatures().XStatusObject()
	xSt.SetEnumProperty(openapiart.XStatusObjectEnumProperty.DECPRECATED_PROPERTY_1)
	// validating the warnings
	_, err := xSt.Marshal().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}

	warns := xSt.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "DECPRECATED_PROPERTY_1 enum in property EnumProperty is deprecated, test deprecated")
}

func TestStatusWarningsInSchemaObjects(t *testing.T) {
	config := openapiart.NewTestConfig()

	// setting all the primitive values which has x-status set
	xSt := config.ExtendedFeatures().XStatusObject()
	xSt.SetEnumProperty(openapiart.XStatusObjectEnumProperty.DECPRECATED_PROPERTY_1)
	// validating the warnings
	_, err := config.Marshal().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}

	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "TestConfig is under review, the whole schema is being reviewed")
}
