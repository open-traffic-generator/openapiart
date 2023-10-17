package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestStatusApi(t *testing.T) {
	err := StartMockGrpcServer()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	// create grpc API
	grpcApi := openapiart.NewApi()
	grpcApi.NewGrpcTransport().SetLocation(grpcServer.Location)

	config := openapiart.NewUpdateConfig()
	config.G().Add().SetGA("str1")

	// check warning for grpc API
	_, err = grpcApi.UpdateConfiguration(config)
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
}

func TestStatusMsgInPrimitiveAttrs(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	// setting all the primitive values which has x-status set
	config.SetA("test")
	config.SetSpace1(32)
	enums := []openapiart.PrefixConfigDValuesEnum{
		openapiart.PrefixConfigDValues.A,
		openapiart.PrefixConfigDValues.B,
		openapiart.PrefixConfigDValues.C,
	}
	config.SetDValues(enums)
	config.SetStrLen("1234")
	config.SetHexSlice([]string{"str1", "str2"})

	// validating the warnings
	_, err := config.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}

	warns := config.Warnings()
	assert.Equal(t, len(warns), 5)
	assert.Equal(t, warns[0], "Space_1 property in schema PrefixConfig is deprecated, Information TBD")
	assert.Equal(t, warns[1], "A property in schema PrefixConfig is under review, Information TBD")
	assert.Equal(t, warns[2], "DValues property in schema PrefixConfig is deprecated, Information TBD")
	assert.Equal(t, warns[3], "StrLen property in schema PrefixConfig is under review, Information TBD")
	assert.Equal(t, warns[4], "HexSlice property in schema PrefixConfig is under review, Information TBD")
}

func TestStatusMsgInStructAttrs(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	// setting a non primitive property with x-status
	config.E().SetEA(4.56)

	// validating the warnings
	_, err := config.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "E property in schema PrefixConfig is deprecated, Information TBD")
}

func TestStatusMsgInChoiceAttrs(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	j := config.J().Add()
	j.JB()
	_, err := j.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := j.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "J_B enum in property Choice is deprecated, use j_a instead")
}

func TestStatusMsgInXEnumAttrs(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_404)

	// validating the warnings
	_, err := config.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "STATUS_404 enum in property Response is deprecated, new code will be coming soon")

	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)

	// validating the warnings
	_, err = config.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns = config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "STATUS_500 enum in property Response is under review, 500 can change to other values")
}

func TestStatusMsgInIterattrs(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	list := config.G()
	list.Append(openapiart.NewGObject().SetGC(5.67))
	list.Append(openapiart.NewGObject().SetGC(7.67))
	list.Append(openapiart.NewGObject().SetGC(8.67))
	assert.Len(t, list.Items(), 3)

	for _, item := range list.Items() {
		_, err := item.Marshaller().ToYaml()
		if err != nil {
			t.Logf("error: %s", err.Error())
		}
		warns := item.Warnings()
		assert.Equal(t, len(warns), 2)
		assert.Equal(t, warns[1], "GC property in schema GObject is deprecated, Information TBD")
	}
}

func TestStatusMsgInSchemaObjects(t *testing.T) {
	config := openapiart.NewUpdateConfig()

	_, err := config.Marshaller().ToYaml()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "UpdateConfig is under review, the whole schema is being reviewed")

	list := config.G()
	list.Append(openapiart.NewGObject().SetGC(5.67))
	list.Append(openapiart.NewGObject().SetGC(7.67))
	list.Append(openapiart.NewGObject().SetGC(8.67))
	assert.Len(t, list.Items(), 3)

	for _, item := range list.Items() {
		_, err := item.Marshaller().ToYaml()
		if err != nil {
			t.Logf("error: %s", err.Error())
		}
		warns := item.Warnings()
		assert.Equal(t, len(warns), 2)
		assert.Equal(t, warns[0], "GObject is deprecated, new schema Jobject to be used")
	}
}
