package openapiart_test

import (
	"testing"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/stretchr/testify/assert"
)

func TestConfigGenerated(t *testing.T) {
	config := openapiart.NewPrefixConfig()
	config.SetIeee8021Qbb(true).SetSpace1(1).SetFullDuplex100Mb(1).SetResponse("status_200").SetA("abc").SetB(100.11).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1}).SetListOfStringValues([]string{"a", "b", "c"}).SetListOfIntegerValues([]int32{1, 2, 3}).SetInteger64(10000000000000000).SetStrLen("abc").SetName("abc1")
	v1 := config.RequiredObject()
	v1.SetEA(100.11).SetEB(1.7976931348623157e+308).SetName("abc2").SetMParam1("abc").SetMParam2("abc")
	v1.HasName()
	v1.HasMParam1()
	v1.HasMParam2()
	v1Json, v1JsonError := v1.ToJson()
	assert.Nil(t, v1JsonError)
	v1Yaml, v1YamlError := v1.ToYaml()
	assert.Nil(t, v1YamlError)
	v1PbText, v1PbTextError := v1.ToPbText()
	assert.Nil(t, v1PbTextError)
	assert.Nil(t, v1.FromJson(v1Json))
	assert.Nil(t, v1.FromYaml(v1Yaml))
	assert.Nil(t, v1.FromPbText(v1PbText))
	config.SetRequiredObject(v1)
	v2 := config.OptionalObject()
	v2.SetEA(100.11).SetEB(1.7976931348623157e+308).SetName("abc3").SetMParam1("abc").SetMParam2("abc")
	v2.HasName()
	v2.HasMParam1()
	v2.HasMParam2()
	v2Json, v2JsonError := v2.ToJson()
	assert.Nil(t, v2JsonError)
	v2Yaml, v2YamlError := v2.ToYaml()
	assert.Nil(t, v2YamlError)
	v2PbText, v2PbTextError := v2.ToPbText()
	assert.Nil(t, v2PbTextError)
	assert.Nil(t, v2.FromJson(v2Json))
	assert.Nil(t, v2.FromYaml(v2Yaml))
	assert.Nil(t, v2.FromPbText(v2PbText))
	config.SetOptionalObject(v2)
	v3 := config.E()
	v3.SetEA(100.11).SetEB(1.7976931348623157e+308).SetName("abc4").SetMParam1("abc").SetMParam2("abc")
	v3.HasName()
	v3.HasMParam1()
	v3.HasMParam2()
	v3Json, v3JsonError := v3.ToJson()
	assert.Nil(t, v3JsonError)
	v3Yaml, v3YamlError := v3.ToYaml()
	assert.Nil(t, v3YamlError)
	v3PbText, v3PbTextError := v3.ToPbText()
	assert.Nil(t, v3PbTextError)
	assert.Nil(t, v3.FromJson(v3Json))
	assert.Nil(t, v3.FromYaml(v3Yaml))
	assert.Nil(t, v3.FromPbText(v3PbText))
	config.SetE(v3)
	v4 := config.F()
	v4.SetFA("abc").SetFB(1.7976931348623157e+308)
	v4.HasChoice()
	v4.HasFA()
	v4.HasFB()
	v4.Choice()
	v4.FA()
	v4.FB()
	v4Json, v4JsonError := v4.ToJson()
	assert.Nil(t, v4JsonError)
	v4Yaml, v4YamlError := v4.ToYaml()
	assert.Nil(t, v4YamlError)
	v4PbText, v4PbTextError := v4.ToPbText()
	assert.Nil(t, v4PbTextError)
	assert.Nil(t, v4.FromJson(v4Json))
	assert.Nil(t, v4.FromYaml(v4Yaml))
	assert.Nil(t, v4.FromPbText(v4PbText))
	config.SetF(v4)
	v5 := config.G().Add()
	v5.SetGA("abc").SetGB(1).SetGC(100.11).SetGD("abc").SetGE(1.7976931348623157e+308).SetGF("a").SetName("abc5")
	v5.HasChoice()
	v5.HasGA()
	v5.HasGB()
	v5.HasGC()
	v5.HasGD()
	v5.HasGE()
	v5.HasGF()
	v5.HasName()
	v5Json, v5JsonError := v5.ToJson()
	assert.Nil(t, v5JsonError)
	v5Yaml, v5YamlError := v5.ToYaml()
	assert.Nil(t, v5YamlError)
	v5PbText, v5PbTextError := v5.ToPbText()
	assert.Nil(t, v5PbTextError)
	assert.Nil(t, v5.FromJson(v5Json))
	assert.Nil(t, v5.FromYaml(v5Yaml))
	assert.Nil(t, v5.FromPbText(v5PbText))
	v6 := config.J().Add()
	v7 := v6.JA()
	v7.SetEA(100.11).SetEB(1.7976931348623157e+308).SetName("abc6").SetMParam1("abc").SetMParam2("abc")
	v7.HasName()
	v7.HasMParam1()
	v7.HasMParam2()
	v7Json, v7JsonError := v7.ToJson()
	assert.Nil(t, v7JsonError)
	v7Yaml, v7YamlError := v7.ToYaml()
	assert.Nil(t, v7YamlError)
	v7PbText, v7PbTextError := v7.ToPbText()
	assert.Nil(t, v7PbTextError)
	assert.Nil(t, v7.FromJson(v7Json))
	assert.Nil(t, v7.FromYaml(v7Yaml))
	assert.Nil(t, v7.FromPbText(v7PbText))
	v6.SetJA(v7)
	v8 := v6.JB()
	v8.SetFA("abc").SetFB(1.7976931348623157e+308)
	v8.HasChoice()
	v8.HasFA()
	v8.HasFB()
	v8.Choice()
	v8.FA()
	v8.FB()
	v8Json, v8JsonError := v8.ToJson()
	assert.Nil(t, v8JsonError)
	v8Yaml, v8YamlError := v8.ToYaml()
	assert.Nil(t, v8YamlError)
	v8PbText, v8PbTextError := v8.ToPbText()
	assert.Nil(t, v8PbTextError)
	assert.Nil(t, v8.FromJson(v8Json))
	assert.Nil(t, v8.FromYaml(v8Yaml))
	assert.Nil(t, v8.FromPbText(v8PbText))
	v6.SetJB(v8)
	v6.HasChoice()
	v6.HasJA()
	v6.HasJB()
	v6Json, v6JsonError := v6.ToJson()
	assert.Nil(t, v6JsonError)
	v6Yaml, v6YamlError := v6.ToYaml()
	assert.Nil(t, v6YamlError)
	v6PbText, v6PbTextError := v6.ToPbText()
	assert.Nil(t, v6PbTextError)
	assert.Nil(t, v6.FromJson(v6Json))
	assert.Nil(t, v6.FromYaml(v6Yaml))
	assert.Nil(t, v6.FromPbText(v6PbText))
	v9 := config.K()
	v10 := v9.EObject()
	v10.SetEA(100.11).SetEB(1.7976931348623157e+308).SetName("abc7").SetMParam1("abc").SetMParam2("abc")
	v10.HasName()
	v10.HasMParam1()
	v10.HasMParam2()
	v10Json, v10JsonError := v10.ToJson()
	assert.Nil(t, v10JsonError)
	v10Yaml, v10YamlError := v10.ToYaml()
	assert.Nil(t, v10YamlError)
	v10PbText, v10PbTextError := v10.ToPbText()
	assert.Nil(t, v10PbTextError)
	assert.Nil(t, v10.FromJson(v10Json))
	assert.Nil(t, v10.FromYaml(v10Yaml))
	assert.Nil(t, v10.FromPbText(v10PbText))
	v9.SetEObject(v10)
	v11 := v9.FObject()
	v11.SetFA("abc").SetFB(1.7976931348623157e+308)
	v11.HasChoice()
	v11.HasFA()
	v11.HasFB()
	v11.Choice()
	v11.FA()
	v11.FB()
	v11Json, v11JsonError := v11.ToJson()
	assert.Nil(t, v11JsonError)
	v11Yaml, v11YamlError := v11.ToYaml()
	assert.Nil(t, v11YamlError)
	v11PbText, v11PbTextError := v11.ToPbText()
	assert.Nil(t, v11PbTextError)
	assert.Nil(t, v11.FromJson(v11Json))
	assert.Nil(t, v11.FromYaml(v11Yaml))
	assert.Nil(t, v11.FromPbText(v11PbText))
	v9.SetFObject(v11)
	v9.HasEObject()
	v9.HasFObject()
	v9Json, v9JsonError := v9.ToJson()
	assert.Nil(t, v9JsonError)
	v9Yaml, v9YamlError := v9.ToYaml()
	assert.Nil(t, v9YamlError)
	v9PbText, v9PbTextError := v9.ToPbText()
	assert.Nil(t, v9PbTextError)
	assert.Nil(t, v9.FromJson(v9Json))
	assert.Nil(t, v9.FromYaml(v9Yaml))
	assert.Nil(t, v9.FromPbText(v9PbText))
	config.SetK(v9)
	v12 := config.L()
	v12.SetStringParam("abc").SetInteger(61).SetFloat(100.11).SetDouble(1.7976931348623157e+308).SetMac("00:00:00:00:00:0a").SetIpv4("1.1.1.1").SetIpv6("2000::1").SetHex("0x12")
	v12.HasStringParam()
	v12.HasInteger()
	v12.HasFloat()
	v12.HasDouble()
	v12.HasMac()
	v12.HasIpv4()
	v12.HasIpv6()
	v12.HasHex()
	v12Json, v12JsonError := v12.ToJson()
	assert.Nil(t, v12JsonError)
	v12Yaml, v12YamlError := v12.ToYaml()
	assert.Nil(t, v12YamlError)
	v12PbText, v12PbTextError := v12.ToPbText()
	assert.Nil(t, v12PbTextError)
	assert.Nil(t, v12.FromJson(v12Json))
	assert.Nil(t, v12.FromYaml(v12Yaml))
	assert.Nil(t, v12.FromPbText(v12PbText))
	config.SetL(v12)
	v13 := config.Level()
	v14 := v13.L1P1()
	v15 := v14.L2P1()
	v15.SetL3P1("abc")
	v15.HasL3P1()
	v15Json, v15JsonError := v15.ToJson()
	assert.Nil(t, v15JsonError)
	v15Yaml, v15YamlError := v15.ToYaml()
	assert.Nil(t, v15YamlError)
	v15PbText, v15PbTextError := v15.ToPbText()
	assert.Nil(t, v15PbTextError)
	assert.Nil(t, v15.FromJson(v15Json))
	assert.Nil(t, v15.FromYaml(v15Yaml))
	assert.Nil(t, v15.FromPbText(v15PbText))
	v14.SetL2P1(v15)
	v14.HasL2P1()
	v14Json, v14JsonError := v14.ToJson()
	assert.Nil(t, v14JsonError)
	v14Yaml, v14YamlError := v14.ToYaml()
	assert.Nil(t, v14YamlError)
	v14PbText, v14PbTextError := v14.ToPbText()
	assert.Nil(t, v14PbTextError)
	assert.Nil(t, v14.FromJson(v14Json))
	assert.Nil(t, v14.FromYaml(v14Yaml))
	assert.Nil(t, v14.FromPbText(v14PbText))
	v13.SetL1P1(v14)
	v13.HasL1P1()
	v13.HasL1P2()
	v13Json, v13JsonError := v13.ToJson()
	assert.Nil(t, v13JsonError)
	v13Yaml, v13YamlError := v13.ToYaml()
	assert.Nil(t, v13YamlError)
	v13PbText, v13PbTextError := v13.ToPbText()
	assert.Nil(t, v13PbTextError)
	assert.Nil(t, v13.FromJson(v13Json))
	assert.Nil(t, v13.FromYaml(v13Yaml))
	assert.Nil(t, v13.FromPbText(v13PbText))
	config.SetLevel(v13)
	v16 := config.Mandatory()
	v16.SetRequiredParam("abc")
	v16Json, v16JsonError := v16.ToJson()
	assert.Nil(t, v16JsonError)
	v16Yaml, v16YamlError := v16.ToYaml()
	assert.Nil(t, v16YamlError)
	v16PbText, v16PbTextError := v16.ToPbText()
	assert.Nil(t, v16PbTextError)
	assert.Nil(t, v16.FromJson(v16Json))
	assert.Nil(t, v16.FromYaml(v16Yaml))
	assert.Nil(t, v16.FromPbText(v16PbText))
	config.SetMandatory(v16)
	v17 := config.Ipv4Pattern()
	v18 := v17.Ipv4()
	v18.SetValue("1.1.1.1").SetValues([]string{"1.1.1.1"})
	v19 := v18.Increment()
	v19.SetStart("1.1.1.1").SetStep("1.1.1.1").SetCount(1)
	v19.HasStart()
	v19.HasStep()
	v19.HasCount()
	v19.Start()
	v19.Step()
	v19.Count()
	v19Json, v19JsonError := v19.ToJson()
	assert.Nil(t, v19JsonError)
	v19Yaml, v19YamlError := v19.ToYaml()
	assert.Nil(t, v19YamlError)
	v19PbText, v19PbTextError := v19.ToPbText()
	assert.Nil(t, v19PbTextError)
	assert.Nil(t, v19.FromJson(v19Json))
	assert.Nil(t, v19.FromYaml(v19Yaml))
	assert.Nil(t, v19.FromPbText(v19PbText))
	v18.SetIncrement(v19)
	v20 := v18.Decrement()
	v20.SetStart("1.1.1.1").SetStep("1.1.1.1").SetCount(1)
	v20.HasStart()
	v20.HasStep()
	v20.HasCount()
	v20.Start()
	v20.Step()
	v20.Count()
	v20Json, v20JsonError := v20.ToJson()
	assert.Nil(t, v20JsonError)
	v20Yaml, v20YamlError := v20.ToYaml()
	assert.Nil(t, v20YamlError)
	v20PbText, v20PbTextError := v20.ToPbText()
	assert.Nil(t, v20PbTextError)
	assert.Nil(t, v20.FromJson(v20Json))
	assert.Nil(t, v20.FromYaml(v20Yaml))
	assert.Nil(t, v20.FromPbText(v20PbText))
	v18.SetDecrement(v20)
	v18.HasChoice()
	v18.HasValue()
	v18.HasIncrement()
	v18.HasDecrement()
	v18.Choice()
	v18.Value()
	v18.Values()
	v18Json, v18JsonError := v18.ToJson()
	assert.Nil(t, v18JsonError)
	v18Yaml, v18YamlError := v18.ToYaml()
	assert.Nil(t, v18YamlError)
	v18PbText, v18PbTextError := v18.ToPbText()
	assert.Nil(t, v18PbTextError)
	assert.Nil(t, v18.FromJson(v18Json))
	assert.Nil(t, v18.FromYaml(v18Yaml))
	assert.Nil(t, v18.FromPbText(v18PbText))
	v17.SetIpv4(v18)
	v17.HasIpv4()
	v17Json, v17JsonError := v17.ToJson()
	assert.Nil(t, v17JsonError)
	v17Yaml, v17YamlError := v17.ToYaml()
	assert.Nil(t, v17YamlError)
	v17PbText, v17PbTextError := v17.ToPbText()
	assert.Nil(t, v17PbTextError)
	assert.Nil(t, v17.FromJson(v17Json))
	assert.Nil(t, v17.FromYaml(v17Yaml))
	assert.Nil(t, v17.FromPbText(v17PbText))
	config.SetIpv4Pattern(v17)
	v21 := config.Ipv6Pattern()
	v22 := v21.Ipv6()
	v22.SetValue("2000::1").SetValues([]string{"2000::1"})
	v23 := v22.Increment()
	v23.SetStart("2000::1").SetStep("2000::1").SetCount(1)
	v23.HasStart()
	v23.HasStep()
	v23.HasCount()
	v23.Start()
	v23.Step()
	v23.Count()
	v23Json, v23JsonError := v23.ToJson()
	assert.Nil(t, v23JsonError)
	v23Yaml, v23YamlError := v23.ToYaml()
	assert.Nil(t, v23YamlError)
	v23PbText, v23PbTextError := v23.ToPbText()
	assert.Nil(t, v23PbTextError)
	assert.Nil(t, v23.FromJson(v23Json))
	assert.Nil(t, v23.FromYaml(v23Yaml))
	assert.Nil(t, v23.FromPbText(v23PbText))
	v22.SetIncrement(v23)
	v24 := v22.Decrement()
	v24.SetStart("2000::1").SetStep("2000::1").SetCount(1)
	v24.HasStart()
	v24.HasStep()
	v24.HasCount()
	v24.Start()
	v24.Step()
	v24.Count()
	v24Json, v24JsonError := v24.ToJson()
	assert.Nil(t, v24JsonError)
	v24Yaml, v24YamlError := v24.ToYaml()
	assert.Nil(t, v24YamlError)
	v24PbText, v24PbTextError := v24.ToPbText()
	assert.Nil(t, v24PbTextError)
	assert.Nil(t, v24.FromJson(v24Json))
	assert.Nil(t, v24.FromYaml(v24Yaml))
	assert.Nil(t, v24.FromPbText(v24PbText))
	v22.SetDecrement(v24)
	v22.HasChoice()
	v22.HasValue()
	v22.HasIncrement()
	v22.HasDecrement()
	v22.Choice()
	v22.Value()
	v22.Values()
	v22Json, v22JsonError := v22.ToJson()
	assert.Nil(t, v22JsonError)
	v22Yaml, v22YamlError := v22.ToYaml()
	assert.Nil(t, v22YamlError)
	v22PbText, v22PbTextError := v22.ToPbText()
	assert.Nil(t, v22PbTextError)
	assert.Nil(t, v22.FromJson(v22Json))
	assert.Nil(t, v22.FromYaml(v22Yaml))
	assert.Nil(t, v22.FromPbText(v22PbText))
	v21.SetIpv6(v22)
	v21.HasIpv6()
	v21Json, v21JsonError := v21.ToJson()
	assert.Nil(t, v21JsonError)
	v21Yaml, v21YamlError := v21.ToYaml()
	assert.Nil(t, v21YamlError)
	v21PbText, v21PbTextError := v21.ToPbText()
	assert.Nil(t, v21PbTextError)
	assert.Nil(t, v21.FromJson(v21Json))
	assert.Nil(t, v21.FromYaml(v21Yaml))
	assert.Nil(t, v21.FromPbText(v21PbText))
	config.SetIpv6Pattern(v21)
	v25 := config.MacPattern()
	v26 := v25.Mac()
	v26.SetValue("00:00:00:00:00:0a").SetValues([]string{"00:00:00:00:00:0a"})
	v27 := v26.Increment()
	v27.SetStart("00:00:00:00:00:0a").SetStep("00:00:00:00:00:0a").SetCount(1)
	v27.HasStart()
	v27.HasStep()
	v27.HasCount()
	v27.Start()
	v27.Step()
	v27.Count()
	v27Json, v27JsonError := v27.ToJson()
	assert.Nil(t, v27JsonError)
	v27Yaml, v27YamlError := v27.ToYaml()
	assert.Nil(t, v27YamlError)
	v27PbText, v27PbTextError := v27.ToPbText()
	assert.Nil(t, v27PbTextError)
	assert.Nil(t, v27.FromJson(v27Json))
	assert.Nil(t, v27.FromYaml(v27Yaml))
	assert.Nil(t, v27.FromPbText(v27PbText))
	v26.SetIncrement(v27)
	v28 := v26.Decrement()
	v28.SetStart("00:00:00:00:00:0a").SetStep("00:00:00:00:00:0a").SetCount(1)
	v28.HasStart()
	v28.HasStep()
	v28.HasCount()
	v28.Start()
	v28.Step()
	v28.Count()
	v28Json, v28JsonError := v28.ToJson()
	assert.Nil(t, v28JsonError)
	v28Yaml, v28YamlError := v28.ToYaml()
	assert.Nil(t, v28YamlError)
	v28PbText, v28PbTextError := v28.ToPbText()
	assert.Nil(t, v28PbTextError)
	assert.Nil(t, v28.FromJson(v28Json))
	assert.Nil(t, v28.FromYaml(v28Yaml))
	assert.Nil(t, v28.FromPbText(v28PbText))
	v26.SetDecrement(v28)
	v26.HasChoice()
	v26.HasValue()
	v26.HasIncrement()
	v26.HasDecrement()
	v26.Choice()
	v26.Value()
	v26.Values()
	v26Json, v26JsonError := v26.ToJson()
	assert.Nil(t, v26JsonError)
	v26Yaml, v26YamlError := v26.ToYaml()
	assert.Nil(t, v26YamlError)
	v26PbText, v26PbTextError := v26.ToPbText()
	assert.Nil(t, v26PbTextError)
	assert.Nil(t, v26.FromJson(v26Json))
	assert.Nil(t, v26.FromYaml(v26Yaml))
	assert.Nil(t, v26.FromPbText(v26PbText))
	v25.SetMac(v26)
	v25.HasMac()
	v25Json, v25JsonError := v25.ToJson()
	assert.Nil(t, v25JsonError)
	v25Yaml, v25YamlError := v25.ToYaml()
	assert.Nil(t, v25YamlError)
	v25PbText, v25PbTextError := v25.ToPbText()
	assert.Nil(t, v25PbTextError)
	assert.Nil(t, v25.FromJson(v25Json))
	assert.Nil(t, v25.FromYaml(v25Yaml))
	assert.Nil(t, v25.FromPbText(v25PbText))
	config.SetMacPattern(v25)
	v29 := config.IntegerPattern()
	v30 := v29.Integer()
	v30.SetValue(248).SetValues([]uint32{65})
	v31 := v30.Increment()
	v31.SetStart(229).SetStep(169).SetCount(1)
	v31.HasStart()
	v31.HasStep()
	v31.HasCount()
	v31.Step()
	v31.Count()
	v31Json, v31JsonError := v31.ToJson()
	assert.Nil(t, v31JsonError)
	v31Yaml, v31YamlError := v31.ToYaml()
	assert.Nil(t, v31YamlError)
	v31PbText, v31PbTextError := v31.ToPbText()
	assert.Nil(t, v31PbTextError)
	assert.Nil(t, v31.FromJson(v31Json))
	assert.Nil(t, v31.FromYaml(v31Yaml))
	assert.Nil(t, v31.FromPbText(v31PbText))
	v30.SetIncrement(v31)
	v32 := v30.Decrement()
	v32.SetStart(49).SetStep(221).SetCount(1)
	v32.HasStart()
	v32.HasStep()
	v32.HasCount()
	v32.Step()
	v32.Count()
	v32Json, v32JsonError := v32.ToJson()
	assert.Nil(t, v32JsonError)
	v32Yaml, v32YamlError := v32.ToYaml()
	assert.Nil(t, v32YamlError)
	v32PbText, v32PbTextError := v32.ToPbText()
	assert.Nil(t, v32PbTextError)
	assert.Nil(t, v32.FromJson(v32Json))
	assert.Nil(t, v32.FromYaml(v32Yaml))
	assert.Nil(t, v32.FromPbText(v32PbText))
	v30.SetDecrement(v32)
	v30.HasChoice()
	v30.HasValue()
	v30.HasIncrement()
	v30.HasDecrement()
	v30.Choice()
	v30.Values()
	v30Json, v30JsonError := v30.ToJson()
	assert.Nil(t, v30JsonError)
	v30Yaml, v30YamlError := v30.ToYaml()
	assert.Nil(t, v30YamlError)
	v30PbText, v30PbTextError := v30.ToPbText()
	assert.Nil(t, v30PbTextError)
	assert.Nil(t, v30.FromJson(v30Json))
	assert.Nil(t, v30.FromYaml(v30Yaml))
	assert.Nil(t, v30.FromPbText(v30PbText))
	v29.SetInteger(v30)
	v29.HasInteger()
	v29Json, v29JsonError := v29.ToJson()
	assert.Nil(t, v29JsonError)
	v29Yaml, v29YamlError := v29.ToYaml()
	assert.Nil(t, v29YamlError)
	v29PbText, v29PbTextError := v29.ToPbText()
	assert.Nil(t, v29PbTextError)
	assert.Nil(t, v29.FromJson(v29Json))
	assert.Nil(t, v29.FromYaml(v29Yaml))
	assert.Nil(t, v29.FromPbText(v29PbText))
	config.SetIntegerPattern(v29)
	v33 := config.ChecksumPattern()
	v34 := v33.Checksum()
	v34.SetGenerated("good").SetCustom(237)
	v34.HasChoice()
	v34.HasGenerated()
	v34.HasCustom()
	v34.Choice()
	v34.Generated()
	v34Json, v34JsonError := v34.ToJson()
	assert.Nil(t, v34JsonError)
	v34Yaml, v34YamlError := v34.ToYaml()
	assert.Nil(t, v34YamlError)
	v34PbText, v34PbTextError := v34.ToPbText()
	assert.Nil(t, v34PbTextError)
	assert.Nil(t, v34.FromJson(v34Json))
	assert.Nil(t, v34.FromYaml(v34Yaml))
	assert.Nil(t, v34.FromPbText(v34PbText))
	v33.SetChecksum(v34)
	v33.HasChecksum()
	v33Json, v33JsonError := v33.ToJson()
	assert.Nil(t, v33JsonError)
	v33Yaml, v33YamlError := v33.ToYaml()
	assert.Nil(t, v33YamlError)
	v33PbText, v33PbTextError := v33.ToPbText()
	assert.Nil(t, v33PbTextError)
	assert.Nil(t, v33.FromJson(v33Json))
	assert.Nil(t, v33.FromYaml(v33Yaml))
	assert.Nil(t, v33.FromPbText(v33PbText))
	config.SetChecksumPattern(v33)
	v35 := config.MObject()
	v35.SetStringParam("abc").SetInteger(24).SetFloat(100.11).SetDouble(1.7976931348623157e+308).SetMac("00:00:00:00:00:0a").SetIpv4("1.1.1.1").SetIpv6("2000::1").SetHex("0x12")
	v35Json, v35JsonError := v35.ToJson()
	assert.Nil(t, v35JsonError)
	v35Yaml, v35YamlError := v35.ToYaml()
	assert.Nil(t, v35YamlError)
	v35PbText, v35PbTextError := v35.ToPbText()
	assert.Nil(t, v35PbTextError)
	assert.Nil(t, v35.FromJson(v35Json))
	assert.Nil(t, v35.FromYaml(v35Yaml))
	assert.Nil(t, v35.FromPbText(v35PbText))
	config.SetMObject(v35)
	v36 := config.HeaderChecksum()
	v36.SetGenerated("good").SetCustom(32949)
	v36.HasChoice()
	v36.HasGenerated()
	v36.HasCustom()
	v36.Choice()
	v36.Generated()
	v36Json, v36JsonError := v36.ToJson()
	assert.Nil(t, v36JsonError)
	v36Yaml, v36YamlError := v36.ToYaml()
	assert.Nil(t, v36YamlError)
	v36PbText, v36PbTextError := v36.ToPbText()
	assert.Nil(t, v36PbTextError)
	assert.Nil(t, v36.FromJson(v36Json))
	assert.Nil(t, v36.FromYaml(v36Yaml))
	assert.Nil(t, v36.FromPbText(v36PbText))
	config.SetHeaderChecksum(v36)
	config.HasOptionalObject()
	config.HasIeee8021Qbb()
	config.HasSpace1()
	config.HasFullDuplex100Mb()
	config.HasResponse()
	config.HasE()
	config.HasF()
	config.HasH()
	config.HasK()
	config.HasL()
	config.HasLevel()
	config.HasMandatory()
	config.HasIpv4Pattern()
	config.HasIpv6Pattern()
	config.HasMacPattern()
	config.HasIntegerPattern()
	config.HasChecksumPattern()
	config.HasCase()
	config.HasMObject()
	config.HasInteger64()
	config.HasHeaderChecksum()
	config.HasStrLen()
	config.HasName()
	config.Response()
	config.H()
	configJson, jerr := config.ToJson()
	assert.Nil(t, jerr)
	configYaml, yerr := config.ToYaml()
	assert.Nil(t, yerr)
	configPbText, pberr := config.ToPbText()
	assert.Nil(t, pberr)
	assert.Nil(t, config.FromJson(configJson))
	assert.Nil(t, config.FromYaml(configYaml))
	assert.Nil(t, config.FromPbText(configPbText))
}
