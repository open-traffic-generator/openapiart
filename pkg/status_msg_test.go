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
	config.SetB(3.45)
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
	assert.Equal(t, warns[0], "A is under review, Information TBD")
	assert.Equal(t, warns[1], "B is deprecated, Information TBD")
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

func TestStatusMsgInEnumAttrs(t *testing.T) {
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
	assert.Equal(t, warns[0], "j_b is deprecated, use j_a instead")

}
