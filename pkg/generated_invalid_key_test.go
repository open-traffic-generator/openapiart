package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPrefixConfigIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPrefixConfig()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestEObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewEObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestFObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewFObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewGObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestJObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewJObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestKObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewKObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLevelOneIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLevelOne()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMandateIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewMandate()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIpv4PatternIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewIpv4Pattern()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIpv6PatternIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewIpv6Pattern()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMacPatternIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewMacPattern()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestIntegerPatternIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewIntegerPattern()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestChecksumPatternIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewChecksumPattern()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLayer1Ieee802XIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLayer1Ieee802X()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMObjectIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewMObject()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternPrefixConfigHeaderChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternPrefixConfigHeaderChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestUpdateConfigIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewUpdateConfig()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestSetConfigResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewSetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestErrorDetailsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewErrorDetails()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestErrorIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewError()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestUpdateConfigResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewUpdateConfigurationResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetConfigResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewGetConfigResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetMetricsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewGetMetricsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestMetricsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewMetrics()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestGetWarningsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewGetWarningsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestWarningDetailsIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewWarningDetails()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestClearWarningsResponseIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewClearWarningsResponse()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLevelTwoIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLevelTwo()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLevelFourIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLevelFour()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIpv4PatternIpv4IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIpv4PatternIpv4()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIpv6PatternIpv6IncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIpv6PatternIpv6()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternMacPatternMacIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternMacPatternMac()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIntegerPatternIntegerIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIntegerPatternInteger()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternChecksumPatternChecksumIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternChecksumPatternChecksum()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPortMetricIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPortMetric()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestLevelThreeIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewLevelThree()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIpv4PatternIpv4CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIpv4PatternIpv4Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIpv6PatternIpv6CounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIpv6PatternIpv6Counter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternMacPatternMacCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternMacPatternMacCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
func TestPatternIntegerPatternIntegerCounterIncorrectKey(t *testing.T) {
	incorrect_key := `{
            "a":"asdf",
            "zxnvh" : 65,
            "c" : 33,
            "h": true,
            "response" : "status_200",
            "required_object" : {
                "e_a" : 1,
                "e_b" : 2
            }
        }`

	object := openapiart.NewPatternIntegerPatternIntegerCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
