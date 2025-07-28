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

	obj, err := fObj.Unmarshal().FromProto(fObj_proto)
	assert.Nil(t, err)

	assert.Equal(t, obj.FA(), "some string")

	proto_obj, err1 := obj.Marshal().ToProto()
	d := "some string"
	var fa *string = &d

	assert.Nil(t, err1)
	assert.Equal(t, *proto_obj.FA, *fa)
	assert.Equal(t, fObj_proto.GetFA(), "")

}

func TestRandomPatternForAllFormat(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.IntegerPattern().Integer().Random()
	config.Ipv4Pattern().Ipv4().Random()
	config.Ipv6Pattern().Ipv6().Random()
	config.MacPattern().Mac().Random()

	json, err := config.Marshal().ToJson()
	assert.Nil(t, err)

	config2 := openapiart.NewPrefixConfig()
	err = config2.Unmarshal().FromJson(json)
	assert.Nil(t, err)

	pat1 := config2.IntegerPattern().Integer().Random()
	assert.Equal(t, pat1.Min(), uint32(0))
	assert.Equal(t, pat1.Max(), uint32(255))
	assert.Equal(t, pat1.Count(), uint32(1))
	assert.Equal(t, pat1.Seed(), uint32(1))

	pat2 := config2.Ipv4Pattern().Ipv4().Random()
	assert.Equal(t, pat2.Min(), "0.0.0.0")
	assert.Equal(t, pat2.Max(), "255.255.255.255")
	assert.Equal(t, pat2.Count(), uint32(1))
	assert.Equal(t, pat2.Seed(), uint32(1))

	pat3 := config2.Ipv6Pattern().Ipv6().Random()
	assert.Equal(t, pat3.Min(), "::")
	assert.Equal(t, pat3.Max(), "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	assert.Equal(t, pat3.Count(), uint32(1))
	assert.Equal(t, pat3.Seed(), uint32(1))

	pat4 := config2.MacPattern().Mac().Random()
	assert.Equal(t, pat4.Min(), "00:00:00:00:00:00")
	assert.Equal(t, pat4.Max(), "ff:ff:ff:ff:ff:ff")
	assert.Equal(t, pat4.Count(), uint32(1))
	assert.Equal(t, pat4.Seed(), uint32(1))

	pat5 := config2.Ipv6PatternWithoutDefault().Ipv6().Random()
	assert.Equal(t, pat5.HasMin(), false)
	assert.Equal(t, pat5.Max(), "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")
	assert.Equal(t, pat5.Count(), uint32(1))
	assert.Equal(t, pat5.Seed(), uint32(1))
}

func TestBinaryDataSerialization(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse(openapiart.PrefixConfigResponse.STATUS_200)
	config.SetBinaryData([]byte("Hello\nbyebye!@#$"))
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	json, err := config.Marshal().ToJson()
	assert.Nil(t, err)
	_, err = config.Marshal().ToYaml()
	assert.Nil(t, err)
	_, err = config.Marshal().ToPbText()
	assert.Nil(t, err)
	_, err = config.Marshal().ToProto()
	assert.Nil(t, err)

	config2 := openapiart.NewPrefixConfig()
	err = config2.Unmarshal().FromJson(json)
	assert.Nil(t, err)
	assert.Equal(t, string(config2.BinaryData()), "Hello\nbyebye!@#$")
}
