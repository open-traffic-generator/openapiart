package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func TestPrefixConfigRequired(t *testing.T) {
	object := goapi.NewPrefixConfig()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "RequiredObject", "A", "B", "C")
	assert.Contains(t, err1.Error(), "RequiredObject", "A", "B", "C")
	assert.Contains(t, err2.Error(), "RequiredObject", "A", "B", "C")
}

//	func TestEObjectRequired(t *testing.T) {
//		object := goapi.NewEObject()
//		opts := protojson.MarshalOptions{
//			UseProtoNames:   true,
//			AllowPartial:    true,
//			EmitUnpopulated: false,
//			Indent:          "  ",
//		}
//		data, _ := opts.Marshal(object.Msg())
//		err := object.FromJson(string(data))
//		err1 := object.FromYaml(string(data))
//		err2 := object.FromPbText(proto.MarshalTextString(object.Msg()))
//		assert.Contains(t, err.Error(), "EA", "EB")
//		assert.Contains(t, err1.Error(), "EA", "EB")
//		assert.Contains(t, err2.Error(), "EA", "EB")
//	}
func TestMandateRequired(t *testing.T) {
	object := goapi.NewMandate()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "RequiredParam")
	assert.Contains(t, err1.Error(), "RequiredParam")
	assert.Contains(t, err2.Error(), "RequiredParam")
}
func TestMObjectRequired(t *testing.T) {
	object := goapi.NewMObject()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "StringParam", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
	assert.Contains(t, err1.Error(), "StringParam", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
	assert.Contains(t, err2.Error(), "StringParam", "Integer", "Float", "Double", "Mac", "Ipv4", "Ipv6", "Hex")
}
func TestPortMetricRequired(t *testing.T) {
	object := goapi.NewPortMetric()
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, _ := opts.Marshal(object.Msg())
	err := object.FromJson(string(data))
	err1 := object.FromYaml(string(data))
	protoMarshal, _ := proto.Marshal(object.Msg())
	err2 := object.FromPbText(string(protoMarshal))
	assert.Contains(t, err.Error(), "Name", "TxFrames", "RxFrames")
	assert.Contains(t, err1.Error(), "Name", "TxFrames", "RxFrames")
	assert.Contains(t, err2.Error(), "Name", "TxFrames", "RxFrames")
}
