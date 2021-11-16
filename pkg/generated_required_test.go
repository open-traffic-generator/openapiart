package openapiart_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestPrefixConfigRequired(t *testing.T) {
	object := openapiart.NewPrefixConfig()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
	assert.Contains(t, err.Error(), "RequiredObject", "Response", "A", "B", "C")
	assert.Contains(t, err1.Error(), "RequiredObject", "Response", "A", "B", "C")
	assert.Contains(t, err2.Error(), "RequiredObject", "Response", "A", "B", "C")
}
func TestEObjectRequired(t *testing.T) {
	object := openapiart.NewEObject()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
	assert.Contains(t, err.Error(), "EA", "EB")
	assert.Contains(t, err1.Error(), "EA", "EB")
	assert.Contains(t, err2.Error(), "EA", "EB")
}
func TestMandateRequired(t *testing.T) {
	object := openapiart.NewMandate()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
	assert.Contains(t, err.Error(), "RequiredParam")
	assert.Contains(t, err1.Error(), "RequiredParam")
	assert.Contains(t, err2.Error(), "RequiredParam")
}
func TestMObjectRequired(t *testing.T) {
	object := openapiart.NewMObject()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
	assert.Contains(t, err.Error(), "String_", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
	assert.Contains(t, err1.Error(), "String_", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
	assert.Contains(t, err2.Error(), "String_", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
}
func TestPortMetricRequired(t *testing.T) {
	object := openapiart.NewPortMetric()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
	assert.Contains(t, err.Error(), "Name", "TxFrames", "RxFrames")
	assert.Contains(t, err1.Error(), "Name", "TxFrames", "RxFrames")
	assert.Contains(t, err2.Error(), "Name", "TxFrames", "RxFrames")
}
