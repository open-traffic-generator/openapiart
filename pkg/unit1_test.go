package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"
	"github.com/stretchr/testify/assert"
)

func TestToAndFromProtoMsg(t *testing.T) {
	fObj_proto := &sanity.FObject{}
	fObj := openapiart.NewFObject()

	obj, err := fObj.Marshaller().FromProto(fObj_proto)
	assert.Nil(t, err)

	assert.Equal(t, obj.FA(), "some string")

	proto_obj, err1 := obj.Marshaller().ToProto()
	d := "some string"
	var fa *string = &d

	assert.Nil(t, err1)
	assert.Equal(t, *proto_obj.FA, *fa)
	assert.Equal(t, fObj_proto.GetFA(), "")

}
