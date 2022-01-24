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
	assert.Contains(t, err.Error(), "prefix_config.required_object")
	assert.Contains(t, err.Error(), "prefix_config.a")
	assert.Contains(t, err1.Error(), "prefix_config.required_object")
	assert.Contains(t, err1.Error(), "prefix_config.a")
	assert.Contains(t, err2.Error(), "prefix_config.required_object")
	assert.Contains(t, err2.Error(), "prefix_config.a")
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
	assert.Contains(t, err.Error(), "mandate.required_param")
	assert.Contains(t, err1.Error(), "mandate.required_param")
	assert.Contains(t, err2.Error(), "mandate.required_param")
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
	assert.Contains(t, err.Error(), "mObject.string_param")
	// assert.Contains(t, err.Error(), "MObject.integer")
	assert.Contains(t, err.Error(), "mObject.ipv4")
	assert.Contains(t, err.Error(), "mObject.mac")
	// assert.Contains(t, err.Error(), "MObject.float")
	// assert.Contains(t, err.Error(), "MObject.double")
	assert.Contains(t, err.Error(), "mObject.ipv6")
	assert.Contains(t, err.Error(), "mObject.hex")
	assert.Contains(t, err1.Error(), "mObject.string")
	// assert.Contains(t, err1.Error(), "MObject.integer")
	assert.Contains(t, err1.Error(), "mObject.ipv4")
	assert.Contains(t, err1.Error(), "mObject.mac")
	// assert.Contains(t, err1.Error(), "MObject.float")
	// assert.Contains(t, err1.Error(), "MObject.double")
	assert.Contains(t, err1.Error(), "mObject.ipv6")
	assert.Contains(t, err1.Error(), "mObject.hex")
	assert.Contains(t, err2.Error(), "mObject.string")
	// assert.Contains(t, err2.Error(), "MObject.integer")
	assert.Contains(t, err2.Error(), "mObject.ipv4")
	assert.Contains(t, err2.Error(), "mObject.mac")
	// assert.Contains(t, err2.Error(), "MObject.float")
	// assert.Contains(t, err2.Error(), "MObject.double")
	assert.Contains(t, err2.Error(), "mObject.ipv6")
	assert.Contains(t, err2.Error(), "mObject.hex")
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
	assert.Contains(t, err.Error(), "port_metric.name")
	assert.Contains(t, err1.Error(), "port_metric.name")
	assert.Contains(t, err2.Error(), "port_metric.name")
}
