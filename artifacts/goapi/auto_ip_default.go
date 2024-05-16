package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** AutoIpDefault *****
type autoIpDefault struct {
	validation
	obj          *openapi.AutoIpDefault
	marshaller   marshalAutoIpDefault
	unMarshaller unMarshalAutoIpDefault
}

func NewAutoIpDefault() AutoIpDefault {
	obj := autoIpDefault{obj: &openapi.AutoIpDefault{}}
	obj.setDefault()
	return &obj
}

func (obj *autoIpDefault) msg() *openapi.AutoIpDefault {
	return obj.obj
}

func (obj *autoIpDefault) setMsg(msg *openapi.AutoIpDefault) AutoIpDefault {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalautoIpDefault struct {
	obj *autoIpDefault
}

type marshalAutoIpDefault interface {
	// ToProto marshals AutoIpDefault to protobuf object *openapi.AutoIpDefault
	ToProto() (*openapi.AutoIpDefault, error)
	// ToPbText marshals AutoIpDefault to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals AutoIpDefault to YAML text
	ToYaml() (string, error)
	// ToJson marshals AutoIpDefault to JSON text
	ToJson() (string, error)
}

type unMarshalautoIpDefault struct {
	obj *autoIpDefault
}

type unMarshalAutoIpDefault interface {
	// FromProto unmarshals AutoIpDefault from protobuf object *openapi.AutoIpDefault
	FromProto(msg *openapi.AutoIpDefault) (AutoIpDefault, error)
	// FromPbText unmarshals AutoIpDefault from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals AutoIpDefault from YAML text
	FromYaml(value string) error
	// FromJson unmarshals AutoIpDefault from JSON text
	FromJson(value string) error
}

func (obj *autoIpDefault) Marshal() marshalAutoIpDefault {
	if obj.marshaller == nil {
		obj.marshaller = &marshalautoIpDefault{obj: obj}
	}
	return obj.marshaller
}

func (obj *autoIpDefault) Unmarshal() unMarshalAutoIpDefault {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalautoIpDefault{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalautoIpDefault) ToProto() (*openapi.AutoIpDefault, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalautoIpDefault) FromProto(msg *openapi.AutoIpDefault) (AutoIpDefault, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalautoIpDefault) ToPbText() (string, error) {
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

func (m *unMarshalautoIpDefault) FromPbText(value string) error {
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

func (m *marshalautoIpDefault) ToYaml() (string, error) {
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

func (m *unMarshalautoIpDefault) FromYaml(value string) error {
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

func (m *marshalautoIpDefault) ToJson() (string, error) {
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

func (m *unMarshalautoIpDefault) FromJson(value string) error {
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

func (obj *autoIpDefault) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *autoIpDefault) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *autoIpDefault) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *autoIpDefault) Clone() (AutoIpDefault, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewAutoIpDefault()
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

// AutoIpDefault is the OTG implementation can provide a system generated,
// value for this property. If the OTG is unable to generate a value,
// the default value must be used.
type AutoIpDefault interface {
	Validation
	// msg marshals AutoIpDefault to protobuf object *openapi.AutoIpDefault
	// and doesn't set defaults
	msg() *openapi.AutoIpDefault
	// setMsg unmarshals AutoIpDefault from protobuf object *openapi.AutoIpDefault
	// and doesn't set defaults
	setMsg(*openapi.AutoIpDefault) AutoIpDefault
	// provides marshal interface
	Marshal() marshalAutoIpDefault
	// provides unmarshal interface
	Unmarshal() unMarshalAutoIpDefault
	// validate validates AutoIpDefault
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (AutoIpDefault, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns AutoIpDefaultChoiceEnum, set in AutoIpDefault
	Choice() AutoIpDefaultChoiceEnum
	// setChoice assigns AutoIpDefaultChoiceEnum provided by user to AutoIpDefault
	setChoice(value AutoIpDefaultChoiceEnum) AutoIpDefault
	// HasChoice checks if Choice has been set in AutoIpDefault
	HasChoice() bool
	// getter for Dhcp to set choice.
	Dhcp()
	// getter for Static to set choice.
	Static()
}

type AutoIpDefaultChoiceEnum string

// Enum of Choice on AutoIpDefault
var AutoIpDefaultChoice = struct {
	STATIC AutoIpDefaultChoiceEnum
	DHCP   AutoIpDefaultChoiceEnum
}{
	STATIC: AutoIpDefaultChoiceEnum("static"),
	DHCP:   AutoIpDefaultChoiceEnum("dhcp"),
}

func (obj *autoIpDefault) Choice() AutoIpDefaultChoiceEnum {
	return AutoIpDefaultChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Dhcp to set choice
func (obj *autoIpDefault) Dhcp() {
	obj.setChoice(AutoIpDefaultChoice.DHCP)
}

// getter for Static to set choice
func (obj *autoIpDefault) Static() {
	obj.setChoice(AutoIpDefaultChoice.STATIC)
}

// description is TBD
// Choice returns a string
func (obj *autoIpDefault) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *autoIpDefault) setChoice(value AutoIpDefaultChoiceEnum) AutoIpDefault {
	intValue, ok := openapi.AutoIpDefault_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on AutoIpDefaultChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.AutoIpDefault_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *autoIpDefault) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *autoIpDefault) setDefault() {
	var choices_set int = 0
	var choice AutoIpDefaultChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(AutoIpDefaultChoice.DHCP)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in AutoIpDefault")
			}
		} else {
			intVal := openapi.AutoIpDefault_Choice_Enum_value[string(choice)]
			enumValue := openapi.AutoIpDefault_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
