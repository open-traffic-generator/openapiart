package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
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

	object := openapiart.NewPrefixConfig()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewEObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewFObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewGObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewJObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewKObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLevelOne()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewMandate()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewIpv4Pattern()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewIpv6Pattern()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewMacPattern()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewIntegerPattern()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewChecksumPattern()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLayer1Ieee802X()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewMObject()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternPrefixConfigHeaderChecksum()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewUpdateConfig()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewSetConfigResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewError()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewError()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewUpdateConfigurationResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewGetConfigResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewGetMetricsResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewMetrics()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewGetWarningsResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewWarningDetails()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewClearWarningsResponse()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLevelTwo()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLevelFour()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIpv4PatternIpv4()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIpv6PatternIpv6()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternMacPatternMac()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIntegerPatternInteger()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternChecksumPatternChecksum()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPortMetric()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewLevelThree()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIpv4PatternIpv4Counter()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIpv6PatternIpv6Counter()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternMacPatternMacCounter()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
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

	object := openapiart.NewPatternIntegerPatternIntegerCounter()
	assert.NotNil(t, object.Marshaller().FromYaml(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromJson(incorrect_format))
	assert.NotNil(t, object.Marshaller().FromPbText(incorrect_format))
}
