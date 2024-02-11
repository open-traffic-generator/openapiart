package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredChoiceParent *****
type requiredChoiceParent struct {
	validation
	obj                   *openapi.RequiredChoiceParent
	marshaller            marshalRequiredChoiceParent
	unMarshaller          unMarshalRequiredChoiceParent
	intermediateObjHolder RequiredChoiceIntermediate
}

func NewRequiredChoiceParent() RequiredChoiceParent {
	obj := requiredChoiceParent{obj: &openapi.RequiredChoiceParent{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceParent) msg() *openapi.RequiredChoiceParent {
	return obj.obj
}

func (obj *requiredChoiceParent) setMsg(msg *openapi.RequiredChoiceParent) RequiredChoiceParent {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredChoiceParent struct {
	obj *requiredChoiceParent
}

type marshalRequiredChoiceParent interface {
	// ToProto marshals RequiredChoiceParent to protobuf object *openapi.RequiredChoiceParent
	ToProto() (*openapi.RequiredChoiceParent, error)
	// ToPbText marshals RequiredChoiceParent to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceParent to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceParent to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredChoiceParent struct {
	obj *requiredChoiceParent
}

type unMarshalRequiredChoiceParent interface {
	// FromProto unmarshals RequiredChoiceParent from protobuf object *openapi.RequiredChoiceParent
	FromProto(msg *openapi.RequiredChoiceParent) (RequiredChoiceParent, error)
	// FromPbText unmarshals RequiredChoiceParent from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceParent from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceParent from JSON text
	FromJson(value string) error
}

func (obj *requiredChoiceParent) Marshal() marshalRequiredChoiceParent {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredChoiceParent{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredChoiceParent) Unmarshal() unMarshalRequiredChoiceParent {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredChoiceParent{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredChoiceParent) ToProto() (*openapi.RequiredChoiceParent, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredChoiceParent) FromProto(msg *openapi.RequiredChoiceParent) (RequiredChoiceParent, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredChoiceParent) ToPbText() (string, error) {
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

func (m *unMarshalrequiredChoiceParent) FromPbText(value string) error {
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

func (m *marshalrequiredChoiceParent) ToYaml() (string, error) {
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

func (m *unMarshalrequiredChoiceParent) FromYaml(value string) error {
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

func (m *marshalrequiredChoiceParent) ToJson() (string, error) {
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

func (m *unMarshalrequiredChoiceParent) FromJson(value string) error {
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

func (obj *requiredChoiceParent) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceParent) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceParent) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceParent) Clone() (RequiredChoiceParent, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceParent()
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

func (obj *requiredChoiceParent) setNil() {
	obj.intermediateObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RequiredChoiceParent is description is TBD
type RequiredChoiceParent interface {
	Validation
	// msg marshals RequiredChoiceParent to protobuf object *openapi.RequiredChoiceParent
	// and doesn't set defaults
	msg() *openapi.RequiredChoiceParent
	// setMsg unmarshals RequiredChoiceParent from protobuf object *openapi.RequiredChoiceParent
	// and doesn't set defaults
	setMsg(*openapi.RequiredChoiceParent) RequiredChoiceParent
	// provides marshal interface
	Marshal() marshalRequiredChoiceParent
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredChoiceParent
	// validate validates RequiredChoiceParent
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceParent, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns RequiredChoiceParentChoiceEnum, set in RequiredChoiceParent
	Choice() RequiredChoiceParentChoiceEnum
	// setChoice assigns RequiredChoiceParentChoiceEnum provided by user to RequiredChoiceParent
	setChoice(value RequiredChoiceParentChoiceEnum) RequiredChoiceParent
	// getter for NoObj to set choice.
	NoObj()
	// IntermediateObj returns RequiredChoiceIntermediate, set in RequiredChoiceParent.
	// RequiredChoiceIntermediate is description is TBD
	IntermediateObj() RequiredChoiceIntermediate
	// SetIntermediateObj assigns RequiredChoiceIntermediate provided by user to RequiredChoiceParent.
	// RequiredChoiceIntermediate is description is TBD
	SetIntermediateObj(value RequiredChoiceIntermediate) RequiredChoiceParent
	// HasIntermediateObj checks if IntermediateObj has been set in RequiredChoiceParent
	HasIntermediateObj() bool
	setNil()
}

type RequiredChoiceParentChoiceEnum string

// Enum of Choice on RequiredChoiceParent
var RequiredChoiceParentChoice = struct {
	INTERMEDIATE_OBJ RequiredChoiceParentChoiceEnum
	NO_OBJ           RequiredChoiceParentChoiceEnum
}{
	INTERMEDIATE_OBJ: RequiredChoiceParentChoiceEnum("intermediate_obj"),
	NO_OBJ:           RequiredChoiceParentChoiceEnum("no_obj"),
}

func (obj *requiredChoiceParent) Choice() RequiredChoiceParentChoiceEnum {
	return RequiredChoiceParentChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoObj to set choice
func (obj *requiredChoiceParent) NoObj() {
	obj.setChoice(RequiredChoiceParentChoice.NO_OBJ)
}

func (obj *requiredChoiceParent) setChoice(value RequiredChoiceParentChoiceEnum) RequiredChoiceParent {
	intValue, ok := openapi.RequiredChoiceParent_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RequiredChoiceParentChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.RequiredChoiceParent_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.IntermediateObj = nil
	obj.intermediateObjHolder = nil

	if value == RequiredChoiceParentChoice.INTERMEDIATE_OBJ {
		obj.obj.IntermediateObj = NewRequiredChoiceIntermediate().msg()
	}

	return obj
}

// description is TBD
// IntermediateObj returns a RequiredChoiceIntermediate
func (obj *requiredChoiceParent) IntermediateObj() RequiredChoiceIntermediate {
	if obj.obj.IntermediateObj == nil {
		obj.setChoice(RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	}
	if obj.intermediateObjHolder == nil {
		obj.intermediateObjHolder = &requiredChoiceIntermediate{obj: obj.obj.IntermediateObj}
	}
	return obj.intermediateObjHolder
}

// description is TBD
// IntermediateObj returns a RequiredChoiceIntermediate
func (obj *requiredChoiceParent) HasIntermediateObj() bool {
	return obj.obj.IntermediateObj != nil
}

// description is TBD
// SetIntermediateObj sets the RequiredChoiceIntermediate value in the RequiredChoiceParent object
func (obj *requiredChoiceParent) SetIntermediateObj(value RequiredChoiceIntermediate) RequiredChoiceParent {
	obj.setChoice(RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	obj.intermediateObjHolder = nil
	obj.obj.IntermediateObj = value.msg()

	return obj
}

func (obj *requiredChoiceParent) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface RequiredChoiceParent")
	}

	if obj.obj.IntermediateObj != nil {

		obj.IntermediateObj().validateObj(vObj, set_default)
	}

}

func (obj *requiredChoiceParent) setDefault() {

}
