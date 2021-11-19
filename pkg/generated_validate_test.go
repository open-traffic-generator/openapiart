package openapiart_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestConfigGeneratedValidate(t *testing.T) {
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	api := openapiart.NewApi()
	config := api.NewPrefixConfig()

	v1 := config.RequiredObject()
	_ = v1.Validate()
	v2 := config.OptionalObject()
	_ = v2.Validate()
	v3 := config.E()
	_ = v3.Validate()
	v4 := config.F()
	_ = v4.Validate()
	v5 := config.G().Add()
	_ = v5.Validate()
	v6 := config.J().Add()
	v7 := v6.JA()
	_ = v7.Validate()
	v8 := v6.JB()
	_ = v8.Validate()
	_ = v6.Validate()
	v9 := config.K()
	v10 := v9.EObject()
	_ = v10.Validate()
	v11 := v9.FObject()
	_ = v11.Validate()
	_ = v9.Validate()
	v12 := config.L()
	v12.SetInteger(9).SetMac("00:00:00:00:00:0a::0b").SetIpv4("1.1.1.1.1").SetIpv6("2000::1::1").SetHex("0x12JKL")
	data12, _ := opts.Marshal(v12.Msg())
	assert.NotNil(t, v12.FromJson(string(data12)))
	assert.NotNil(t, v12.FromYaml(string(data12)))
	assert.NotNil(t, v12.FromPbText(proto.MarshalTextString(v12.Msg())))
	v13 := config.Level()
	v14 := v13.L1P1()
	v15 := v14.L2P1()
	_ = v15.Validate()
	_ = v14.Validate()
	_ = v13.Validate()
	v16 := config.Mandatory()
	_ = v16.Validate()
	v17 := config.Ipv4Pattern()
	v18 := v17.Ipv4()
	v18.SetValue("1.1.1.1.1").SetValues([]string{"1.1.1.1.1"})
	v19 := v18.Increment()
	v19.SetStart("1.1.1.1.1").SetStep("1.1.1.1.1")
	data19, _ := opts.Marshal(v19.Msg())
	assert.NotNil(t, v19.FromJson(string(data19)))
	assert.NotNil(t, v19.FromYaml(string(data19)))
	assert.NotNil(t, v19.FromPbText(proto.MarshalTextString(v19.Msg())))
	v20 := v18.Decrement()
	v20.SetStart("1.1.1.1.1").SetStep("1.1.1.1.1")
	data20, _ := opts.Marshal(v20.Msg())
	assert.NotNil(t, v20.FromJson(string(data20)))
	assert.NotNil(t, v20.FromYaml(string(data20)))
	assert.NotNil(t, v20.FromPbText(proto.MarshalTextString(v20.Msg())))
	data18, _ := opts.Marshal(v18.Msg())
	assert.NotNil(t, v18.FromJson(string(data18)))
	assert.NotNil(t, v18.FromYaml(string(data18)))
	assert.NotNil(t, v18.FromPbText(proto.MarshalTextString(v18.Msg())))
	_ = v17.Validate()
	v21 := config.Ipv6Pattern()
	v22 := v21.Ipv6()
	v22.SetValue("2000::1::1").SetValues([]string{"2000::1::2"})
	v23 := v22.Increment()
	v23.SetStart("2000::1::1").SetStep("2000::1::1")
	data23, _ := opts.Marshal(v23.Msg())
	assert.NotNil(t, v23.FromJson(string(data23)))
	assert.NotNil(t, v23.FromYaml(string(data23)))
	assert.NotNil(t, v23.FromPbText(proto.MarshalTextString(v23.Msg())))
	v24 := v22.Decrement()
	v24.SetStart("2000::1::1").SetStep("2000::1::1")
	data24, _ := opts.Marshal(v24.Msg())
	assert.NotNil(t, v24.FromJson(string(data24)))
	assert.NotNil(t, v24.FromYaml(string(data24)))
	assert.NotNil(t, v24.FromPbText(proto.MarshalTextString(v24.Msg())))
	data22, _ := opts.Marshal(v22.Msg())
	assert.NotNil(t, v22.FromJson(string(data22)))
	assert.NotNil(t, v22.FromYaml(string(data22)))
	assert.NotNil(t, v22.FromPbText(proto.MarshalTextString(v22.Msg())))
	_ = v21.Validate()
	v25 := config.MacPattern()
	v26 := v25.Mac()
	v26.SetValue("00:00:00:00:00:0a::0b").SetValues([]string{"00:00:00:00:00:0a::b"})
	v27 := v26.Increment()
	v27.SetStart("00:00:00:00:00:0a::0b").SetStep("00:00:00:00:00:0a::0b")
	data27, _ := opts.Marshal(v27.Msg())
	assert.NotNil(t, v27.FromJson(string(data27)))
	assert.NotNil(t, v27.FromYaml(string(data27)))
	assert.NotNil(t, v27.FromPbText(proto.MarshalTextString(v27.Msg())))
	v28 := v26.Decrement()
	v28.SetStart("00:00:00:00:00:0a::0b").SetStep("00:00:00:00:00:0a::0b")
	data28, _ := opts.Marshal(v28.Msg())
	assert.NotNil(t, v28.FromJson(string(data28)))
	assert.NotNil(t, v28.FromYaml(string(data28)))
	assert.NotNil(t, v28.FromPbText(proto.MarshalTextString(v28.Msg())))
	data26, _ := opts.Marshal(v26.Msg())
	assert.NotNil(t, v26.FromJson(string(data26)))
	assert.NotNil(t, v26.FromYaml(string(data26)))
	assert.NotNil(t, v26.FromPbText(proto.MarshalTextString(v26.Msg())))
	_ = v25.Validate()
	v29 := config.IntegerPattern()
	v30 := v29.Integer()
	v30.SetValue(-1).SetValues([]int32{-1})
	v31 := v30.Increment()
	v31.SetStart(-1).SetStep(-1)
	data31, _ := opts.Marshal(v31.Msg())
	assert.NotNil(t, v31.FromJson(string(data31)))
	assert.NotNil(t, v31.FromYaml(string(data31)))
	assert.NotNil(t, v31.FromPbText(proto.MarshalTextString(v31.Msg())))
	v32 := v30.Decrement()
	v32.SetStart(-1).SetStep(-1)
	data32, _ := opts.Marshal(v32.Msg())
	assert.NotNil(t, v32.FromJson(string(data32)))
	assert.NotNil(t, v32.FromYaml(string(data32)))
	assert.NotNil(t, v32.FromPbText(proto.MarshalTextString(v32.Msg())))
	data30, _ := opts.Marshal(v30.Msg())
	assert.NotNil(t, v30.FromJson(string(data30)))
	assert.NotNil(t, v30.FromYaml(string(data30)))
	assert.NotNil(t, v30.FromPbText(proto.MarshalTextString(v30.Msg())))
	_ = v29.Validate()
	v33 := config.ChecksumPattern()
	v34 := v33.Checksum()
	v34.SetCustom(-1)
	data34, _ := opts.Marshal(v34.Msg())
	assert.NotNil(t, v34.FromJson(string(data34)))
	assert.NotNil(t, v34.FromYaml(string(data34)))
	assert.NotNil(t, v34.FromPbText(proto.MarshalTextString(v34.Msg())))
	_ = v33.Validate()
	v35 := config.MObject()
	v35.SetInteger(9).SetMac("00:00:00:00:00:0a::0b").SetIpv4("1.1.1.1.1").SetIpv6("2000::1::1").SetHex("0x12JKL")
	data35, _ := opts.Marshal(v35.Msg())
	assert.NotNil(t, v35.FromJson(string(data35)))
	assert.NotNil(t, v35.FromYaml(string(data35)))
	assert.NotNil(t, v35.FromPbText(proto.MarshalTextString(v35.Msg())))
	v36 := config.HeaderChecksum()
	v36.SetCustom(-1)
	data36, _ := opts.Marshal(v36.Msg())
	assert.NotNil(t, v36.FromJson(string(data36)))
	assert.NotNil(t, v36.FromYaml(string(data36)))
	assert.NotNil(t, v36.FromPbText(proto.MarshalTextString(v36.Msg())))

}