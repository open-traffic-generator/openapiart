package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChoiceValWithNoProperties *****
type choiceValWithNoProperties struct {
	validation
	obj                   *openapi.ChoiceValWithNoProperties
	marshaller            marshalChoiceValWithNoProperties
	unMarshaller          unMarshalChoiceValWithNoProperties
	intermediateObjHolder RequiredChoice
}

func NewChoiceValWithNoProperties() ChoiceValWithNoProperties {
	obj := choiceValWithNoProperties{obj: &openapi.ChoiceValWithNoProperties{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceValWithNoProperties) msg() *openapi.ChoiceValWithNoProperties {
	return obj.obj
}

func (obj *choiceValWithNoProperties) setMsg(msg *openapi.ChoiceValWithNoProperties) ChoiceValWithNoProperties {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchoiceValWithNoProperties struct {
	obj *choiceValWithNoProperties
}

type marshalChoiceValWithNoProperties interface {
	// ToProto marshals ChoiceValWithNoProperties to protobuf object *openapi.ChoiceValWithNoProperties
	ToProto() (*openapi.ChoiceValWithNoProperties, error)
	// ToPbText marshals ChoiceValWithNoProperties to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceValWithNoProperties to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceValWithNoProperties to JSON text
	ToJson() (string, error)
}

type unMarshalchoiceValWithNoProperties struct {
	obj *choiceValWithNoProperties
}

type unMarshalChoiceValWithNoProperties interface {
	// FromProto unmarshals ChoiceValWithNoProperties from protobuf object *openapi.ChoiceValWithNoProperties
	FromProto(msg *openapi.ChoiceValWithNoProperties) (ChoiceValWithNoProperties, error)
	// FromPbText unmarshals ChoiceValWithNoProperties from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceValWithNoProperties from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceValWithNoProperties from JSON text
	FromJson(value string) error
}

func (obj *choiceValWithNoProperties) Marshal() marshalChoiceValWithNoProperties {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchoiceValWithNoProperties{obj: obj}
	}
	return obj.marshaller
}

func (obj *choiceValWithNoProperties) Unmarshal() unMarshalChoiceValWithNoProperties {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchoiceValWithNoProperties{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchoiceValWithNoProperties) ToProto() (*openapi.ChoiceValWithNoProperties, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchoiceValWithNoProperties) FromProto(msg *openapi.ChoiceValWithNoProperties) (ChoiceValWithNoProperties, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchoiceValWithNoProperties) ToPbText() (string, error) {
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

func (m *unMarshalchoiceValWithNoProperties) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalchoiceValWithNoProperties) ToYaml() (string, error) {
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

func (m *unMarshalchoiceValWithNoProperties) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalchoiceValWithNoProperties) ToJson() (string, error) {
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

func (m *unMarshalchoiceValWithNoProperties) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *choiceValWithNoProperties) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceValWithNoProperties) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceValWithNoProperties) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceValWithNoProperties) Clone() (ChoiceValWithNoProperties, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceValWithNoProperties()
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

func (obj *choiceValWithNoProperties) setNil() {
	obj.intermediateObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChoiceValWithNoProperties is description is TBD
type ChoiceValWithNoProperties interface {
	Validation
	// msg marshals ChoiceValWithNoProperties to protobuf object *openapi.ChoiceValWithNoProperties
	// and doesn't set defaults
	msg() *openapi.ChoiceValWithNoProperties
	// setMsg unmarshals ChoiceValWithNoProperties from protobuf object *openapi.ChoiceValWithNoProperties
	// and doesn't set defaults
	setMsg(*openapi.ChoiceValWithNoProperties) ChoiceValWithNoProperties
	// provides marshal interface
	Marshal() marshalChoiceValWithNoProperties
	// provides unmarshal interface
	Unmarshal() unMarshalChoiceValWithNoProperties
	// validate validates ChoiceValWithNoProperties
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceValWithNoProperties, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ChoiceValWithNoPropertiesChoiceEnum, set in ChoiceValWithNoProperties
	Choice() ChoiceValWithNoPropertiesChoiceEnum
	// setChoice assigns ChoiceValWithNoPropertiesChoiceEnum provided by user to ChoiceValWithNoProperties
	setChoice(value ChoiceValWithNoPropertiesChoiceEnum) ChoiceValWithNoProperties
	// getter for NoObj to set choice.
	NoObj()
	// IntermediateObj returns RequiredChoice, set in ChoiceValWithNoProperties.
	// RequiredChoice is description is TBD
	IntermediateObj() RequiredChoice
	// SetIntermediateObj assigns RequiredChoice provided by user to ChoiceValWithNoProperties.
	// RequiredChoice is description is TBD
	SetIntermediateObj(value RequiredChoice) ChoiceValWithNoProperties
	// HasIntermediateObj checks if IntermediateObj has been set in ChoiceValWithNoProperties
	HasIntermediateObj() bool
	setNil()
}

type ChoiceValWithNoPropertiesChoiceEnum string

// Enum of Choice on ChoiceValWithNoProperties
var ChoiceValWithNoPropertiesChoice = struct {
	INTERMEDIATE_OBJ ChoiceValWithNoPropertiesChoiceEnum
	NO_OBJ           ChoiceValWithNoPropertiesChoiceEnum
}{
	INTERMEDIATE_OBJ: ChoiceValWithNoPropertiesChoiceEnum("intermediate_obj"),
	NO_OBJ:           ChoiceValWithNoPropertiesChoiceEnum("no_obj"),
}

func (obj *choiceValWithNoProperties) Choice() ChoiceValWithNoPropertiesChoiceEnum {
	return ChoiceValWithNoPropertiesChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoObj to set choice
func (obj *choiceValWithNoProperties) NoObj() {
	obj.setChoice(ChoiceValWithNoPropertiesChoice.NO_OBJ)
}

func (obj *choiceValWithNoProperties) setChoice(value ChoiceValWithNoPropertiesChoiceEnum) ChoiceValWithNoProperties {
	intValue, ok := openapi.ChoiceValWithNoProperties_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ChoiceValWithNoPropertiesChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.ChoiceValWithNoProperties_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IntermediateObj = nil
	obj.intermediateObjHolder = nil

	if value == ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ {
		obj.obj.IntermediateObj = NewRequiredChoice().msg()
	}

	return obj
}

// description is TBD
// IntermediateObj returns a RequiredChoice
func (obj *choiceValWithNoProperties) IntermediateObj() RequiredChoice {
	if obj.obj.IntermediateObj == nil {
		obj.setChoice(ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	}
	if obj.intermediateObjHolder == nil {
		obj.intermediateObjHolder = &requiredChoice{obj: obj.obj.IntermediateObj}
	}
	return obj.intermediateObjHolder
}

// description is TBD
// IntermediateObj returns a RequiredChoice
func (obj *choiceValWithNoProperties) HasIntermediateObj() bool {
	return obj.obj.IntermediateObj != nil
}

// description is TBD
// SetIntermediateObj sets the RequiredChoice value in the ChoiceValWithNoProperties object
func (obj *choiceValWithNoProperties) SetIntermediateObj(value RequiredChoice) ChoiceValWithNoProperties {
	obj.setChoice(ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ)
	obj.intermediateObjHolder = nil
	obj.obj.IntermediateObj = value.msg()

	return obj
}

func (obj *choiceValWithNoProperties) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ChoiceValWithNoProperties")
	}

	if obj.obj.IntermediateObj != nil {

		obj.IntermediateObj().validateObj(vObj, set_default)
	}

}

func (obj *choiceValWithNoProperties) setDefault() {
	var choices_set int = 0
	var choice ChoiceValWithNoPropertiesChoiceEnum

	if obj.obj.IntermediateObj != nil {
		choices_set += 1
		choice = ChoiceValWithNoPropertiesChoice.INTERMEDIATE_OBJ
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ChoiceValWithNoProperties")
			}
		} else {
			intVal := openapi.ChoiceValWithNoProperties_Choice_Enum_value[string(choice)]
			enumValue := openapi.ChoiceValWithNoProperties_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
