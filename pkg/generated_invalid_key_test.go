package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
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

	object := goapi.NewPrefixConfig()
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

	object := goapi.NewEObject()
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

	object := goapi.NewFObject()
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

	object := goapi.NewGObject()
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

	object := goapi.NewJObject()
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

	object := goapi.NewKObject()
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

	object := goapi.NewLObject()
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

	object := goapi.NewLevelOne()
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

	object := goapi.NewMandate()
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

	object := goapi.NewIpv4Pattern()
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

	object := goapi.NewIpv6Pattern()
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

	object := goapi.NewMacPattern()
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

	object := goapi.NewIntegerPattern()
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

	object := goapi.NewChecksumPattern()
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

	object := goapi.NewLayer1Ieee802X()
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

	object := goapi.NewMObject()
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

	object := goapi.NewPatternPrefixConfigHeaderChecksum()
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

	object := goapi.NewUpdateConfig()
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

	object := goapi.NewSetConfigResponse()
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

	object := goapi.NewError()
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

	object := goapi.NewError()
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

	object := goapi.NewUpdateConfigurationResponse()
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

	object := goapi.NewGetConfigResponse()
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

	object := goapi.NewGetMetricsResponse()
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

	object := goapi.NewMetrics()
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

	object := goapi.NewGetWarningsResponse()
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

	object := goapi.NewWarningDetails()
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

	object := goapi.NewClearWarningsResponse()
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

	object := goapi.NewLevelTwo()
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

	object := goapi.NewLevelFour()
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

	object := goapi.NewPatternIpv4PatternIpv4()
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

	object := goapi.NewPatternIpv6PatternIpv6()
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

	object := goapi.NewPatternMacPatternMac()
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

	object := goapi.NewPatternIntegerPatternInteger()
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

	object := goapi.NewPatternChecksumPatternChecksum()
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

	object := goapi.NewPortMetric()
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

	object := goapi.NewLevelThree()
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

	object := goapi.NewPatternIpv4PatternIpv4Counter()
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

	object := goapi.NewPatternIpv6PatternIpv6Counter()
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

	object := goapi.NewPatternMacPatternMacCounter()
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

	object := goapi.NewPatternIntegerPatternIntegerCounter()
	assert.NotNil(t, object.FromYaml(incorrect_key))
	assert.NotNil(t, object.FromJson(incorrect_key))
	assert.NotNil(t, object.FromPbText(incorrect_key))
}
