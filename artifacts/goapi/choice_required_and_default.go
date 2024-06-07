package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChoiceRequiredAndDefault *****
type choiceRequiredAndDefault struct {
	validation
	obj          *openapi.ChoiceRequiredAndDefault
	marshaller   marshalChoiceRequiredAndDefault
	unMarshaller unMarshalChoiceRequiredAndDefault
}

func NewChoiceRequiredAndDefault() ChoiceRequiredAndDefault {
	obj := choiceRequiredAndDefault{obj: &openapi.ChoiceRequiredAndDefault{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceRequiredAndDefault) msg() *openapi.ChoiceRequiredAndDefault {
	return obj.obj
}

func (obj *choiceRequiredAndDefault) setMsg(msg *openapi.ChoiceRequiredAndDefault) ChoiceRequiredAndDefault {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchoiceRequiredAndDefault struct {
	obj *choiceRequiredAndDefault
}

type marshalChoiceRequiredAndDefault interface {
	// ToProto marshals ChoiceRequiredAndDefault to protobuf object *openapi.ChoiceRequiredAndDefault
	ToProto() (*openapi.ChoiceRequiredAndDefault, error)
	// ToPbText marshals ChoiceRequiredAndDefault to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceRequiredAndDefault to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceRequiredAndDefault to JSON text
	ToJson() (string, error)
}

type unMarshalchoiceRequiredAndDefault struct {
	obj *choiceRequiredAndDefault
}

type unMarshalChoiceRequiredAndDefault interface {
	// FromProto unmarshals ChoiceRequiredAndDefault from protobuf object *openapi.ChoiceRequiredAndDefault
	FromProto(msg *openapi.ChoiceRequiredAndDefault) (ChoiceRequiredAndDefault, error)
	// FromPbText unmarshals ChoiceRequiredAndDefault from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceRequiredAndDefault from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceRequiredAndDefault from JSON text
	FromJson(value string) error
}

func (obj *choiceRequiredAndDefault) Marshal() marshalChoiceRequiredAndDefault {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchoiceRequiredAndDefault{obj: obj}
	}
	return obj.marshaller
}

func (obj *choiceRequiredAndDefault) Unmarshal() unMarshalChoiceRequiredAndDefault {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchoiceRequiredAndDefault{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchoiceRequiredAndDefault) ToProto() (*openapi.ChoiceRequiredAndDefault, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchoiceRequiredAndDefault) FromProto(msg *openapi.ChoiceRequiredAndDefault) (ChoiceRequiredAndDefault, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchoiceRequiredAndDefault) ToPbText() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (m *unMarshalchoiceRequiredAndDefault) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalchoiceRequiredAndDefault) ToYaml() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalchoiceRequiredAndDefault) FromYaml(value string) error {
	if value == "" {
		value = "{}"
	}
	data, err := yaml.YAMLToJSON([]byte(value))
	if err != nil {
		return err
	}
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	uError := opts.Unmarshal([]byte(data), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalchoiceRequiredAndDefault) ToJson() (string, error) {
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(m.obj.msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (m *unMarshalchoiceRequiredAndDefault) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), m.obj.msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *choiceRequiredAndDefault) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceRequiredAndDefault) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceRequiredAndDefault) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceRequiredAndDefault) Clone() (ChoiceRequiredAndDefault, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceRequiredAndDefault()
	data, err := proto.Marshal(obj.msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// ChoiceRequiredAndDefault is description is TBD
type ChoiceRequiredAndDefault interface {
	Validation
	// msg marshals ChoiceRequiredAndDefault to protobuf object *openapi.ChoiceRequiredAndDefault
	// and doesn't set defaults
	msg() *openapi.ChoiceRequiredAndDefault
	// setMsg unmarshals ChoiceRequiredAndDefault from protobuf object *openapi.ChoiceRequiredAndDefault
	// and doesn't set defaults
	setMsg(*openapi.ChoiceRequiredAndDefault) ChoiceRequiredAndDefault
	// provides marshal interface
	Marshal() marshalChoiceRequiredAndDefault
	// provides unmarshal interface
	Unmarshal() unMarshalChoiceRequiredAndDefault
	// validate validates ChoiceRequiredAndDefault
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceRequiredAndDefault, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ChoiceRequiredAndDefaultChoiceEnum, set in ChoiceRequiredAndDefault
	Choice() ChoiceRequiredAndDefaultChoiceEnum
	// setChoice assigns ChoiceRequiredAndDefaultChoiceEnum provided by user to ChoiceRequiredAndDefault
	setChoice(value ChoiceRequiredAndDefaultChoiceEnum) ChoiceRequiredAndDefault
	// Ipv4 returns string, set in ChoiceRequiredAndDefault.
	Ipv4() string
	// SetIpv4 assigns string provided by user to ChoiceRequiredAndDefault
	SetIpv4(value string) ChoiceRequiredAndDefault
	// HasIpv4 checks if Ipv4 has been set in ChoiceRequiredAndDefault
	HasIpv4() bool
	// Ipv6 returns []string, set in ChoiceRequiredAndDefault.
	Ipv6() []string
	// SetIpv6 assigns []string provided by user to ChoiceRequiredAndDefault
	SetIpv6(value []string) ChoiceRequiredAndDefault
}

type ChoiceRequiredAndDefaultChoiceEnum string

// Enum of Choice on ChoiceRequiredAndDefault
var ChoiceRequiredAndDefaultChoice = struct {
	IPV4 ChoiceRequiredAndDefaultChoiceEnum
	IPV6 ChoiceRequiredAndDefaultChoiceEnum
}{
	IPV4: ChoiceRequiredAndDefaultChoiceEnum("ipv4"),
	IPV6: ChoiceRequiredAndDefaultChoiceEnum("ipv6"),
}

func (obj *choiceRequiredAndDefault) Choice() ChoiceRequiredAndDefaultChoiceEnum {
	return ChoiceRequiredAndDefaultChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *choiceRequiredAndDefault) setChoice(value ChoiceRequiredAndDefaultChoiceEnum) ChoiceRequiredAndDefault {
	intValue, ok := openapi.ChoiceRequiredAndDefault_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ChoiceRequiredAndDefaultChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.ChoiceRequiredAndDefault_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ipv6 = nil
	obj.obj.Ipv4 = nil

	if value == ChoiceRequiredAndDefaultChoice.IPV4 {
		defaultValue := "0.0.0.0"
		obj.obj.Ipv4 = &defaultValue
	}

	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *choiceRequiredAndDefault) Ipv4() string {

	if obj.obj.Ipv4 == nil {
		obj.setChoice(ChoiceRequiredAndDefaultChoice.IPV4)
	}

	return *obj.obj.Ipv4

}

// description is TBD
// Ipv4 returns a string
func (obj *choiceRequiredAndDefault) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the string value in the ChoiceRequiredAndDefault object
func (obj *choiceRequiredAndDefault) SetIpv4(value string) ChoiceRequiredAndDefault {
	obj.setChoice(ChoiceRequiredAndDefaultChoice.IPV4)
	obj.obj.Ipv4 = &value
	return obj
}

// A list of ipv6
// Ipv6 returns a []string
func (obj *choiceRequiredAndDefault) Ipv6() []string {
	if obj.obj.Ipv6 == nil {

		obj.setChoice(ChoiceRequiredAndDefaultChoice.IPV6)

	}
	return obj.obj.Ipv6
}

// A list of ipv6
// SetIpv6 sets the []string value in the ChoiceRequiredAndDefault object
func (obj *choiceRequiredAndDefault) SetIpv6(value []string) ChoiceRequiredAndDefault {
	obj.setChoice(ChoiceRequiredAndDefaultChoice.IPV6)
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = make([]string, 0)
	}
	obj.obj.Ipv6 = value

	return obj
}

func (obj *choiceRequiredAndDefault) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ChoiceRequiredAndDefault")
	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ChoiceRequiredAndDefault.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6Slice(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ChoiceRequiredAndDefault.Ipv6"))
		}

	}

}

func (obj *choiceRequiredAndDefault) setDefault() {
	var choices_set int = 0
	var choice ChoiceRequiredAndDefaultChoiceEnum

	if obj.obj.Ipv4 != nil {
		choices_set += 1
		choice = ChoiceRequiredAndDefaultChoice.IPV4
	}

	if len(obj.obj.Ipv6) > 0 {
		choices_set += 1
		choice = ChoiceRequiredAndDefaultChoice.IPV6
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ChoiceRequiredAndDefault")
			}
		} else {
			intVal := openapi.ChoiceRequiredAndDefault_Choice_Enum_value[string(choice)]
			enumValue := openapi.ChoiceRequiredAndDefault_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Ipv4 == nil && choice == ChoiceRequiredAndDefaultChoice.IPV4 {
		obj.SetIpv4("0.0.0.0")
	}

}
