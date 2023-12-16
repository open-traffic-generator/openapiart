package gopine_test

import (
	"testing"

	gopine "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFromRpfXML(t *testing.T) {
	testResponse := []byte("<object-response request-id=\"40002\"><retval type=\"int64\" actualType=\"int64\">1671845795195</retval></object-response>")
	cfg1 := gopine.NewResponseGetTimedActionTimestamp()
	err := cfg1.FromRpfXml(testResponse)
	assert.Nil(t, err)
	assert.Equal(t, cfg1.HasValue(), true)
	assert.Equal(t, cfg1.Value(), int64(1671845795195))

	testResponse = []byte("<object-response request-id=\"40003\">\n\t<retval type=\"bool\" actualType=\"bool\">1</retval>\n\t<retval type=\"string\" actualType=\"string\">TxMazuma10GTimedCommandDriver::scheduleCommand: Time Request Fault.</retval>\n</object-response>\n")
	cfg2 := gopine.NewResponseScheduleTimedActions()
	err = cfg2.FromRpfXml(testResponse)
	assert.Nil(t, err)
	assert.Nil(t, err)
	assert.Equal(t, cfg2.HasError(), true)
	assert.Equal(t, cfg2.Error(), true)
	assert.Equal(t, cfg2.HasErrorMessage(), true)
	assert.Equal(t, cfg2.ErrorMessage(), "TxMazuma10GTimedCommandDriver::scheduleCommand: Time Request Fault.")
}
