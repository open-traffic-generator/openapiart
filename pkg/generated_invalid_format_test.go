package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPrefixConfigIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPrefixConfig()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestEObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewEObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestFObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewFObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestGObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewGObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestJObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewJObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestKObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewKObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLevelOneIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLevelOne()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestMandateIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewMandate()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestIpv4PatternIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewIpv4Pattern()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestIpv6PatternIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewIpv6Pattern()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestMacPatternIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewMacPattern()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestIntegerPatternIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewIntegerPattern()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestChecksumPatternIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewChecksumPattern()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLayer1Ieee802XIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLayer1Ieee802X()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestMObjectIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewMObject()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternPrefixConfigHeaderChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternPrefixConfigHeaderChecksum()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestUpdateConfigIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewUpdateConfig()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestSetConfigResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewSetConfigResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestErrorDetailsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewError()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestErrorIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewError()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestUpdateConfigResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewUpdateConfigurationResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestGetConfigResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewGetConfigResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestGetMetricsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewGetMetricsResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestMetricsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewMetrics()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestGetWarningsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewGetWarningsResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestWarningDetailsIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewWarningDetails()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestClearWarningsResponseIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewClearWarningsResponse()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLevelTwoIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLevelTwo()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLevelFourIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLevelFour()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIpv4PatternIpv4IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIpv4PatternIpv4()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIpv6PatternIpv6IncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIpv6PatternIpv6()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternMacPatternMacIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternMacPatternMac()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIntegerPatternIntegerIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIntegerPatternInteger()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternChecksumPatternChecksumIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternChecksumPatternChecksum()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPortMetricIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPortMetric()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestLevelThreeIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewLevelThree()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIpv4PatternIpv4CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIpv4PatternIpv4Counter()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIpv6PatternIpv6CounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIpv6PatternIpv6Counter()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternMacPatternMacCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternMacPatternMacCounter()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
func TestPatternIntegerPatternIntegerCounterIncorrectFormat(t *testing.T) {
	incorrect_format := `{
		"a":"asdf",
		"b" : 65,
		"c" : 33,
		"h": true,
		"response" : "status_200",
		"required_object" :
			"e_a" : 1,
			"e_b" : 2
	    }`

	object := goapi.NewPatternIntegerPatternIntegerCounter()
	assert.NotNil(t, object.Unmarshal().FromYaml(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromJson(incorrect_format))
	assert.NotNil(t, object.Unmarshal().FromPbText(incorrect_format))
}
