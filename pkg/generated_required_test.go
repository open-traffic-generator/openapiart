package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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
	str, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(str))
	assert.Contains(t, err.Error(), "PrefixConfig.required_object")
	assert.Contains(t, err.Error(), "PrefixConfig.a")
	assert.Contains(t, err1.Error(), "PrefixConfig.required_object")
	assert.Contains(t, err1.Error(), "PrefixConfig.a")
	assert.Contains(t, err2.Error(), "PrefixConfig.required_object")
	assert.Contains(t, err2.Error(), "PrefixConfig.a")
	// assert.Contains(t, err2.Error(), "PrefixConfig.b")
	// assert.Contains(t, err2.Error(), "PrefixConfig.c")
}

// func TestEObjectRequired(t *testing.T) {
// 	object := openapiart.NewEObject()
// 	opts := protojson.MarshalOptions{
// 		UseProtoNames:   true,
// 		AllowPartial:    true,
// 		EmitUnpopulated: false,
// 		Indent:          "  ",
// 	}
// 	data, _ := opts.Marshal(object.Msg())
// 	err := object.FromJson(string(data))
// 	err1 := object.FromYaml(string(data))
// 	err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
// 	assert.Contains(t, err.Error(), "EA", "EB")
// 	assert.Contains(t, err1.Error(), "EA", "EB")
// 	assert.Contains(t, err2.Error(), "EA", "EB")
// }
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
	str, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(str))
	assert.Contains(t, err.Error(), "Mandate.required_param")
	assert.Contains(t, err1.Error(), "Mandate.required_param")
	assert.Contains(t, err2.Error(), "Mandate.required_param")
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
	str, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(str))
	assert.Contains(t, err.Error(), "MObject.string")
	// assert.Contains(t, err.Error(), "MObject.integer")
	assert.Contains(t, err.Error(), "MObject.ipv4")
	assert.Contains(t, err.Error(), "MObject.mac")
	// assert.Contains(t, err.Error(), "MObject.float")
	// assert.Contains(t, err.Error(), "MObject.double")
	assert.Contains(t, err.Error(), "MObject.ipv6")
	assert.Contains(t, err.Error(), "MObject.hex")
	assert.Contains(t, err1.Error(), "MObject.string")
	// assert.Contains(t, err1.Error(), "MObject.integer")
	assert.Contains(t, err1.Error(), "MObject.ipv4")
	assert.Contains(t, err1.Error(), "MObject.mac")
	// assert.Contains(t, err1.Error(), "MObject.float")
	// assert.Contains(t, err1.Error(), "MObject.double")
	assert.Contains(t, err1.Error(), "MObject.ipv6")
	assert.Contains(t, err1.Error(), "MObject.hex")
	assert.Contains(t, err2.Error(), "MObject.string")
	// assert.Contains(t, err2.Error(), "MObject.integer")
	assert.Contains(t, err2.Error(), "MObject.ipv4")
	assert.Contains(t, err2.Error(), "MObject.mac")
	// assert.Contains(t, err2.Error(), "MObject.float")
	// assert.Contains(t, err2.Error(), "MObject.double")
	assert.Contains(t, err2.Error(), "MObject.ipv6")
	assert.Contains(t, err2.Error(), "MObject.hex")
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
	str, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(str))
	assert.Contains(t, err.Error(), "PortMetric.name")
	assert.Contains(t, err1.Error(), "PortMetric.name")
	assert.Contains(t, err2.Error(), "PortMetric.name")
}
