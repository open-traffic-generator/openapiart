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

	config := grpcApi.NewUpdateConfig()
	warnStr := "UpdateConfiguration is deprecated, please use post instead"

	// check warning for grpc API
	_, err = grpcApi.UpdateConfiguration(config)
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	assert.Equal(t, grpcApi.Warnings(), warnStr)
}

func TestStatusMsgInPrimitiveAttrs(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

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
	err := config.Validate()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}

	warns := config.Warnings()
	assert.Equal(t, len(warns), 5)
	assert.Equal(t, warns[0], "Space_1 is deprecated, Information TBD")
	assert.Equal(t, warns[1], "A is under review, Information TBD")
	assert.Equal(t, warns[2], "DValues is deprecated, Information TBD")
	assert.Equal(t, warns[3], "StrLen is under review, Information TBD")
	assert.Equal(t, warns[4], "HexSlice is under review, Information TBD")
}

func TestStatusMsgInStructAttrs(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	// setting a non primitive property with x-status
	config.E().SetEA(4.56)

	// validating the warnings
	err := config.Validate()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "E is deprecated, Information TBD")
}

func TestStatusMsgInChoiceAttrs(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	j := config.J().Add()
	j.JB()
	err := j.Validate()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := j.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "J_B is deprecated, use j_a instead")
}

func TestStatusMsgInXEnumAttrs(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_404)

	// validating the warnings
	err := config.Validate()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns := config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "STATUS_404 is deprecated, new code will be coming soon")

	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_500)

	// validating the warnings
	err = config.Validate()
	if err != nil {
		t.Logf("error: %s", err.Error())
	}
	warns = config.Warnings()
	assert.Equal(t, len(warns), 1)
	assert.Equal(t, warns[0], "STATUS_500 is under review, 500 can change to other values")
}

func TestStatusMsgInIterattrs(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	list := config.G()
	list.Append(openapiart.NewGObject().SetGC(5.67))
	list.Append(openapiart.NewGObject().SetGC(7.67))
	list.Append(openapiart.NewGObject().SetGC(8.67))
	assert.Len(t, list.Items(), 3)

	for _, item := range list.Items() {
		err := item.Validate()
		if err != nil {
			t.Logf("error: %s", err.Error())
		}
		warns := item.Warnings()
		assert.Equal(t, len(warns), 2)
		assert.Equal(t, warns[1], "GC is deprecated, Information TBD")
	}
}

func TestStatusMsgInSchemaObjects(t *testing.T) {
	api := openapiart.NewApi()
	config := api.NewUpdateConfig()

	err := config.Validate()
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
		err := item.Validate()
		if err != nil {
			t.Logf("error: %s", err.Error())
		}
		warns := item.Warnings()
		assert.Equal(t, len(warns), 2)
		assert.Equal(t, warns[0], "GObject is deprecated, new schema Jobject to be used")
	}
}
