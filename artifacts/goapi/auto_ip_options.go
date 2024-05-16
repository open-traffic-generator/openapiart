package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** AutoIpOptions *****
type autoIpOptions struct {
	validation
	obj          *openapi.AutoIpOptions
	marshaller   marshalAutoIpOptions
	unMarshaller unMarshalAutoIpOptions
}

func NewAutoIpOptions() AutoIpOptions {
	obj := autoIpOptions{obj: &openapi.AutoIpOptions{}}
	obj.setDefault()
	return &obj
}

func (obj *autoIpOptions) msg() *openapi.AutoIpOptions {
	return obj.obj
}

func (obj *autoIpOptions) setMsg(msg *openapi.AutoIpOptions) AutoIpOptions {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalautoIpOptions struct {
	obj *autoIpOptions
}

type marshalAutoIpOptions interface {
	// ToProto marshals AutoIpOptions to protobuf object *openapi.AutoIpOptions
	ToProto() (*openapi.AutoIpOptions, error)
	// ToPbText marshals AutoIpOptions to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals AutoIpOptions to YAML text
	ToYaml() (string, error)
	// ToJson marshals AutoIpOptions to JSON text
	ToJson() (string, error)
}

type unMarshalautoIpOptions struct {
	obj *autoIpOptions
}

type unMarshalAutoIpOptions interface {
	// FromProto unmarshals AutoIpOptions from protobuf object *openapi.AutoIpOptions
	FromProto(msg *openapi.AutoIpOptions) (AutoIpOptions, error)
	// FromPbText unmarshals AutoIpOptions from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals AutoIpOptions from YAML text
	FromYaml(value string) error
	// FromJson unmarshals AutoIpOptions from JSON text
	FromJson(value string) error
}

func (obj *autoIpOptions) Marshal() marshalAutoIpOptions {
	if obj.marshaller == nil {
		obj.marshaller = &marshalautoIpOptions{obj: obj}
	}
	return obj.marshaller
}

func (obj *autoIpOptions) Unmarshal() unMarshalAutoIpOptions {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalautoIpOptions{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalautoIpOptions) ToProto() (*openapi.AutoIpOptions, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalautoIpOptions) FromProto(msg *openapi.AutoIpOptions) (AutoIpOptions, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalautoIpOptions) ToPbText() (string, error) {
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

func (m *unMarshalautoIpOptions) FromPbText(value string) error {
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

func (m *marshalautoIpOptions) ToYaml() (string, error) {
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

func (m *unMarshalautoIpOptions) FromYaml(value string) error {
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

func (m *marshalautoIpOptions) ToJson() (string, error) {
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

func (m *unMarshalautoIpOptions) FromJson(value string) error {
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

func (obj *autoIpOptions) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *autoIpOptions) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *autoIpOptions) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *autoIpOptions) Clone() (AutoIpOptions, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewAutoIpOptions()
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

// AutoIpOptions is the OTG implementation can provide a system generated,
// value for this property. If the OTG is unable to generate a value,
// the default value must be used.
type AutoIpOptions interface {
	Validation
	// msg marshals AutoIpOptions to protobuf object *openapi.AutoIpOptions
	// and doesn't set defaults
	msg() *openapi.AutoIpOptions
	// setMsg unmarshals AutoIpOptions from protobuf object *openapi.AutoIpOptions
	// and doesn't set defaults
	setMsg(*openapi.AutoIpOptions) AutoIpOptions
	// provides marshal interface
	Marshal() marshalAutoIpOptions
	// provides unmarshal interface
	Unmarshal() unMarshalAutoIpOptions
	// validate validates AutoIpOptions
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (AutoIpOptions, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns AutoIpOptionsChoiceEnum, set in AutoIpOptions
	Choice() AutoIpOptionsChoiceEnum
	// setChoice assigns AutoIpOptionsChoiceEnum provided by user to AutoIpOptions
	setChoice(value AutoIpOptionsChoiceEnum) AutoIpOptions
	// getter for Dhcp to set choice.
	Dhcp()
	// getter for Static to set choice.
	Static()
}

type AutoIpOptionsChoiceEnum string

// Enum of Choice on AutoIpOptions
var AutoIpOptionsChoice = struct {
	STATIC AutoIpOptionsChoiceEnum
	DHCP   AutoIpOptionsChoiceEnum
}{
	STATIC: AutoIpOptionsChoiceEnum("static"),
	DHCP:   AutoIpOptionsChoiceEnum("dhcp"),
}

func (obj *autoIpOptions) Choice() AutoIpOptionsChoiceEnum {
	return AutoIpOptionsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Dhcp to set choice
func (obj *autoIpOptions) Dhcp() {
	obj.setChoice(AutoIpOptionsChoice.DHCP)
}

// getter for Static to set choice
func (obj *autoIpOptions) Static() {
	obj.setChoice(AutoIpOptionsChoice.STATIC)
}

func (obj *autoIpOptions) setChoice(value AutoIpOptionsChoiceEnum) AutoIpOptions {
	intValue, ok := openapi.AutoIpOptions_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on AutoIpOptionsChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.AutoIpOptions_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

func (obj *autoIpOptions) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface AutoIpOptions")
	}
}

func (obj *autoIpOptions) setDefault() {
	var choices_set int = 0
	var choice AutoIpOptionsChoiceEnum
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in AutoIpOptions")
			}
		} else {
			intVal := openapi.AutoIpOptions_Choice_Enum_value[string(choice)]
			enumValue := openapi.AutoIpOptions_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
