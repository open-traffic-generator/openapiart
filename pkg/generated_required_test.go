package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPrefixConfigRequired(t *testing.T) {
	object := openapiart.NewPrefixConfig()
	err := object.Validate()
	assert.Contains(t, err.Error(), "RequiredObject", "Response", "A", "B", "C")
}
func TestEObjectRequired(t *testing.T) {
	object := openapiart.NewEObject()
	err := object.Validate()
	assert.Contains(t, err.Error(), "EA", "EB")
}
func TestMandateRequired(t *testing.T) {
	object := openapiart.NewMandate()
	err := object.Validate()
	assert.Contains(t, err.Error(), "RequiredParam")
}
func TestMObjectRequired(t *testing.T) {
	object := openapiart.NewMObject()
	err := object.Validate()
	assert.Contains(t, err.Error(), "String_", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
}
func TestPortMetricRequired(t *testing.T) {
	object := openapiart.NewPortMetric()
	err := object.Validate()
	assert.Contains(t, err.Error(), "Name", "TxFrames", "RxFrames")
}
