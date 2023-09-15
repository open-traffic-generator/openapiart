package goapi_test

import (
	"fmt"
	"testing"
	"time"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/stretchr/testify/assert"
)

func TestPerf(t *testing.T) {
	start := time.Now()
	api := goapi.NewApi()
	config := api.NewPrefixConfig()
	config.SetA("asdf").SetB(12.2).SetC(1).SetH(true).SetI([]byte{1, 0, 0, 1, 0, 0, 1, 1})
	config.SetCase(goapi.NewLayer1Ieee802X().SetFlowControl(true))
	config.SetSpace1(10)
	config.OptionalObject().SetEA(10.1).SetEB(0.001)
	config.RequiredObject().SetEA(1).SetEB(2)
	config.SetIeee8021Qbb(true)
	config.SetFullDuplex100Mb(2)
	config.SetResponse(goapi.PrefixConfigResponse.STATUS_200)
	config.E().SetEA(1.1).SetEB(1.2).SetMParam1("Mparam1").SetMParam2("Mparam2")
	config.F().SetFB(3.0)
	config.G().Add().SetGA("a g_a value").SetGB(6).SetGC(77.7).SetGE(3.0)
	config.J().Add().JA().SetEA(1.0).SetEB(2.0)
	config.K().EObject().SetEA(77.7).SetEB(2.0).SetName("An EB name")
	config.K().FObject().SetFA("asdf").SetFB(44.32232)
	l := config.L()

	l.SetStringParam("test")
	l.SetInteger(80)
	l.SetFloat(100.11)
	l.SetDouble(1.7976931348623157e+308)
	l.SetMac("00:00:00:00:00:0a")
	l.SetIpv4("1.1.1.1")
	l.SetIpv6("2000::1")
	l.SetHex("0x12")
	config.SetListOfIntegerValues([]int32{1, 2, 3})
	config.SetListOfStringValues([]string{"first string", "second string", "third string"})
	level := config.Level()
	level.L1P1().L2P1().SetL3P1("test")
	level.L1P2().L4P1().L1P2().L4P1().L1P1().L2P1().SetL3P1("l3p1")
	config.Mandatory().SetRequiredParam("required")
	config.Ipv4Pattern().Ipv4().SetValue("1.1.1.1")
	config.Ipv6Pattern().Ipv6().SetValues([]string{"2000::1", "2001::2"})
	config.IntegerPattern().Integer().Increment().SetStart(1).SetStart(1).SetCount(100)
	config.MacPattern().Mac().Decrement().SetStart("00:00:00:00:00:0a").SetStep("00:00:00:00:00:01").SetCount(100)
	config.ChecksumPattern().Checksum().SetCustom(64)
	// config.Validate()
	end := time.Now()

	fmt.Printf("Time elapsed for manual configuration %d ms\n", (end.Nanosecond()-start.Nanosecond())/1000)

	jStart := time.Now()
	json, j_err := config.ToJson()
	jEnd := time.Now()
	yStart := time.Now()
	yaml, y_err := config.ToYaml()
	yEnd := time.Now()
	pStart := time.Now()
	pbText, p_err := config.ToPbText()
	pEnd := time.Now()

	fmt.Printf("Time elapsed to serialize to Json %d ms \n", (jEnd.Nanosecond()-jStart.Nanosecond())/1000)
	fmt.Printf("Time elapsed to serialize to Yaml %d ms \n", (yEnd.Nanosecond()-yStart.Nanosecond())/1000)
	fmt.Printf("Time elapsed to serialize to PbText %d ms \n", (pEnd.Nanosecond()-pStart.Nanosecond())/1000)

	assert.Nil(t, j_err)
	assert.Nil(t, y_err)
	assert.Nil(t, p_err)

	jDStart := time.Now()
	jsonconf := api.NewPrefixConfig()
	jdErr := jsonconf.FromJson(json)
	assert.Nil(t, jdErr)
	jDEnd := time.Now()

	yDStart := time.Now()
	yamlconf := api.NewPrefixConfig()
	ydErr := yamlconf.FromYaml(yaml)
	assert.Nil(t, ydErr)
	yDEnd := time.Now()

	pDStart := time.Now()
	pbConf := api.NewPrefixConfig()
	pdErr := pbConf.FromPbText(pbText)
	assert.Nil(t, pdErr)
	pDEnd := time.Now()

	fmt.Printf("Time elapsed to deserialize to Json %d ms \n", (jDEnd.Nanosecond()-jDStart.Nanosecond())/1000)
	fmt.Printf("Time elapsed to deserialize to Yaml %d ms \n", (yDEnd.Nanosecond()-yDStart.Nanosecond())/1000)
	fmt.Printf("Time elapsed to deserialize to PbText %d ms \n", (pDEnd.Nanosecond()-pDStart.Nanosecond())/1000)

	callStart := time.Now()
	_, apiErr := apis[0].SetConfig(config)
	assert.Nil(t, apiErr)
	callEnd := time.Now()
	fmt.Printf("Time elapsed to Call SetConfig %d ms \n", (callEnd.Nanosecond()-callStart.Nanosecond())/1000)
}
