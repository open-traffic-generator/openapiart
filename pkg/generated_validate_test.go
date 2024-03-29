package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestConfigGeneratedValidate(t *testing.T) {
	config := openapiart.NewPrefixConfig()

	v1 := config.RequiredObject()
	_, err := v1.Marshal().ToYaml()
	assert.NotNil(t, err)
	v2 := config.OptionalObject()
	_, err = v2.Marshal().ToJson()
	assert.NotNil(t, err)
	v3 := config.E()
	_, err = v3.Marshal().ToYaml()
	assert.NotNil(t, err)
	v4 := config.F()
	_, err = v4.Marshal().ToYaml()
	assert.Nil(t, err)
	v5 := config.G().Add()
	_, err = v5.Marshal().ToYaml()
	assert.Nil(t, err)
	v6 := config.J().Add()
	v7 := v6.JA()
	_, err = v7.Marshal().ToYaml()
	assert.NotNil(t, err)
	v8 := v6.JB()
	_, err = v8.Marshal().ToYaml()
	assert.Nil(t, err)
	_, err = v6.Marshal().ToYaml()
	assert.Nil(t, err)
	v9 := config.K()
	v10 := v9.EObject()
	_, err = v10.Marshal().ToYaml()
	assert.NotNil(t, err)
	v11 := v9.FObject()
	_, err = v11.Marshal().ToYaml()
	assert.Nil(t, err)
	_, err = v9.Marshal().ToYaml()
	assert.NotNil(t, err)
	v12 := config.L()
	v12.SetInteger(9).SetMac("00:00:00:00:00:0a::0b").SetIpv4("1.1.1.1.1").SetIpv6("2000::1::1").SetHex("0x12JKL")
	_, err = v12.Marshal().ToProto()
	assert.NotNil(t, err)
	v13 := config.Level()
	v14 := v13.L1P1()
	v15 := v14.L2P1()
	_, err = v15.Marshal().ToYaml()
	assert.Nil(t, err)
	_, err = v13.Marshal().ToYaml()
	assert.Nil(t, err)
	_, err = v14.Marshal().ToYaml()
	assert.Nil(t, err)
	v16 := config.Mandatory()
	_, err = v16.Marshal().ToYaml()
	assert.NotNil(t, err)
	v17 := config.Ipv4Pattern()
	v18 := v17.Ipv4()
	v18.SetValue("1.1.1.1.1").SetValues([]string{"1.1.1.1.1"})
	v19 := v18.Increment()
	v19.SetStart("1.1.1.1.1").SetStep("1.1.1.1.1")
	_, err = v19.Marshal().ToProto()
	assert.NotNil(t, err)
	v20 := v18.Decrement()
	v20.SetStart("1.1.1.1.1").SetStep("1.1.1.1.1")
	_, err = v20.Marshal().ToProto()
	assert.NotNil(t, err)
	_, err = v18.Marshal().ToProto()
	assert.NotNil(t, err)
	v21 := config.Ipv6Pattern()
	v22 := v21.Ipv6()
	v22.SetValue("2000::1::1").SetValues([]string{"2000::1::2"})
	v23 := v22.Increment()
	v23.SetStart("2000::1::1").SetStep("2000::1::1")
	_, err = v23.Marshal().ToProto()
	assert.NotNil(t, err)
	v24 := v22.Decrement()
	v24.SetStart("2000::1::1").SetStep("2000::1::1")
	_, err = v24.Marshal().ToProto()
	assert.NotNil(t, err)
	_, err = v22.Marshal().ToProto()
	assert.NotNil(t, err)
	v25 := config.MacPattern()
	v26 := v25.Mac()
	v26.SetValue("00:00:00:00:00:0a::0b").SetValues([]string{"00:00:00:00:00:0a::b"})
	v27 := v26.Increment()
	v27.SetStart("00:00:00:00:00:0a::0b").SetStep("00:00:00:00:00:0a::0b")
	_, err = v27.Marshal().ToProto()
	assert.NotNil(t, err)
	v28 := v26.Decrement()
	v28.SetStart("00:00:00:00:00:0a::0b").SetStep("00:00:00:00:00:0a::0b")
	_, err = v28.Marshal().ToProto()
	assert.NotNil(t, err)
	_, err = v26.Marshal().ToProto()
	assert.NotNil(t, err)
	_, err = v25.Marshal().ToYaml()
	assert.NotNil(t, err)
	v29 := config.IntegerPattern()
	v30 := v29.Integer()
	v30.SetValue(1).SetValues([]uint32{1})
	v31 := v30.Increment()
	v31.SetStart(500).SetStep(500)
	_, err = v31.Marshal().ToProto()
	assert.NotNil(t, err)
	v32 := v30.Decrement()
	v32.SetStart(500).SetStep(500)
	_, err = v32.Marshal().ToProto()
	assert.NotNil(t, err)
	_, err = v30.Marshal().ToProto()
	assert.NotNil(t, err)
	v33 := config.ChecksumPattern()
	v34 := v33.Checksum()
	v34.SetCustom(500)
	_, err = v34.Marshal().ToProto()
	assert.NotNil(t, err)
	v35 := config.MObject()
	v35.SetInteger(9).SetMac("00:00:00:00:00:0a::0b").SetIpv4("1.1.1.1.1").SetIpv6("2000::1::1").SetHex("0x12JKL")
	_, err = v35.Marshal().ToProto()
	assert.NotNil(t, err)
	v36 := config.HeaderChecksum()
	v36.SetCustom(12345678)
	_, err = v36.Marshal().ToProto()
	assert.NotNil(t, err)

}
