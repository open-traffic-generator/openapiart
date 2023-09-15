package goapi_test

import (
	"testing"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	sanity "github.com/open-traffic-generator/goapi/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestToAndFromProtoMsg(t *testing.T) {
	fObj_proto := &sanity.FObject{}
	fObj := goapi.NewFObject()

	obj, err := fObj.FromProto(fObj_proto)
	assert.Nil(t, err)

	assert.Equal(t, obj.FA(), "some string")

	proto_obj, err1 := obj.ToProto()
	d := "some string"
	var fa *string = &d

	assert.Nil(t, err1)
	assert.Equal(t, *proto_obj.FA, *fa)
	assert.Equal(t, fObj_proto.GetFA(), "")

}
