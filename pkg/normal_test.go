package gopine_test

import (
	"encoding/xml"
	"fmt"
	"testing"

	gopine "github.com/open-traffic-generator/openapiart/pkg"
)

func TestRpfToXML(t *testing.T) {
	config := gopine.NewRequestPrepareForStartTx()
	config.PortId().SetId(0)
	config.SetIsTransmitting(false)
	// fmt.Println(config)
	ret_xml, _ := config.ToRpfXml()
	bytes, err := xml.MarshalIndent(ret_xml, "", "\t")
	fmt.Println(string(bytes))
	fmt.Printf("\n\n%+v %+v\n\n", ret_xml, err)

	cfg := gopine.NewRequestGetTimedActionTimestamp()
	cfg.PortId().SetId(0)
	// fmt.Println(config)
	ret_xml, _ = cfg.ToRpfXml()
	bytes, err = xml.MarshalIndent(ret_xml, "", "\t")
	fmt.Println(string(bytes))
	// fmt.Printf("\n\n%+v %+v\n\n", ret_xml, err)

	cfg2 := gopine.NewRequestScheduleTimedActions()
	cfg2.PortId().SetId(1)
	cfg2.ActionList().TimedActionList().Add().SetWhen(1671846795195).SetAction(gopine.ServerTimedActionAction.KSTARTTX)
	// fmt.Println(cfg2)
	ret_xml, _ = cfg2.ToRpfXml()
	bytes, err = xml.MarshalIndent(ret_xml, "", "\t")
	fmt.Println(string(bytes))
	// fmt.Printf("\n\n%+v %+v\n\n", ret_xml, err)
}
