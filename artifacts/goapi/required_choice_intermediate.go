package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredChoiceIntermediate *****
type requiredChoiceIntermediate struct {
	validation
	obj          *openapi.RequiredChoiceIntermediate
	marshaller   marshalRequiredChoiceIntermediate
	unMarshaller unMarshalRequiredChoiceIntermediate
	leafHolder   RequiredChoiceIntermeLeaf
}

func NewRequiredChoiceIntermediate() RequiredChoiceIntermediate {
	obj := requiredChoiceIntermediate{obj: &openapi.RequiredChoiceIntermediate{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceIntermediate) msg() *openapi.RequiredChoiceIntermediate {
	return obj.obj
}

func (obj *requiredChoiceIntermediate) setMsg(msg *openapi.RequiredChoiceIntermediate) RequiredChoiceIntermediate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredChoiceIntermediate struct {
	obj *requiredChoiceIntermediate
}

type marshalRequiredChoiceIntermediate interface {
	// ToProto marshals RequiredChoiceIntermediate to protobuf object *openapi.RequiredChoiceIntermediate
	ToProto() (*openapi.RequiredChoiceIntermediate, error)
	// ToPbText marshals RequiredChoiceIntermediate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceIntermediate to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceIntermediate to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredChoiceIntermediate struct {
	obj *requiredChoiceIntermediate
}

type unMarshalRequiredChoiceIntermediate interface {
	// FromProto unmarshals RequiredChoiceIntermediate from protobuf object *openapi.RequiredChoiceIntermediate
	FromProto(msg *openapi.RequiredChoiceIntermediate) (RequiredChoiceIntermediate, error)
	// FromPbText unmarshals RequiredChoiceIntermediate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceIntermediate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceIntermediate from JSON text
	FromJson(value string) error
}

func (obj *requiredChoiceIntermediate) Marshal() marshalRequiredChoiceIntermediate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredChoiceIntermediate{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredChoiceIntermediate) Unmarshal() unMarshalRequiredChoiceIntermediate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredChoiceIntermediate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredChoiceIntermediate) ToProto() (*openapi.RequiredChoiceIntermediate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredChoiceIntermediate) FromProto(msg *openapi.RequiredChoiceIntermediate) (RequiredChoiceIntermediate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredChoiceIntermediate) ToPbText() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermediate) FromPbText(value string) error {
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

func (m *marshalrequiredChoiceIntermediate) ToYaml() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermediate) FromYaml(value string) error {
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

func (m *marshalrequiredChoiceIntermediate) ToJson() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermediate) FromJson(value string) error {
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

func (obj *requiredChoiceIntermediate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermediate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermediate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceIntermediate) Clone() (RequiredChoiceIntermediate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceIntermediate()
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

func (obj *requiredChoiceIntermediate) setNil() {
	obj.leafHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RequiredChoiceIntermediate is description is TBD
type RequiredChoiceIntermediate interface {
	Validation
	// msg marshals RequiredChoiceIntermediate to protobuf object *openapi.RequiredChoiceIntermediate
	// and doesn't set defaults
	msg() *openapi.RequiredChoiceIntermediate
	// setMsg unmarshals RequiredChoiceIntermediate from protobuf object *openapi.RequiredChoiceIntermediate
	// and doesn't set defaults
	setMsg(*openapi.RequiredChoiceIntermediate) RequiredChoiceIntermediate
	// provides marshal interface
	Marshal() marshalRequiredChoiceIntermediate
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredChoiceIntermediate
	// validate validates RequiredChoiceIntermediate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceIntermediate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns RequiredChoiceIntermediateChoiceEnum, set in RequiredChoiceIntermediate
	Choice() RequiredChoiceIntermediateChoiceEnum
	// setChoice assigns RequiredChoiceIntermediateChoiceEnum provided by user to RequiredChoiceIntermediate
	setChoice(value RequiredChoiceIntermediateChoiceEnum) RequiredChoiceIntermediate
	// FA returns string, set in RequiredChoiceIntermediate.
	FA() string
	// SetFA assigns string provided by user to RequiredChoiceIntermediate
	SetFA(value string) RequiredChoiceIntermediate
	// HasFA checks if FA has been set in RequiredChoiceIntermediate
	HasFA() bool
	// Leaf returns RequiredChoiceIntermeLeaf, set in RequiredChoiceIntermediate.
	// RequiredChoiceIntermeLeaf is description is TBD
	Leaf() RequiredChoiceIntermeLeaf
	// SetLeaf assigns RequiredChoiceIntermeLeaf provided by user to RequiredChoiceIntermediate.
	// RequiredChoiceIntermeLeaf is description is TBD
	SetLeaf(value RequiredChoiceIntermeLeaf) RequiredChoiceIntermediate
	// HasLeaf checks if Leaf has been set in RequiredChoiceIntermediate
	HasLeaf() bool
	setNil()
}

type RequiredChoiceIntermediateChoiceEnum string

// Enum of Choice on RequiredChoiceIntermediate
var RequiredChoiceIntermediateChoice = struct {
	F_A  RequiredChoiceIntermediateChoiceEnum
	LEAF RequiredChoiceIntermediateChoiceEnum
}{
	F_A:  RequiredChoiceIntermediateChoiceEnum("f_a"),
	LEAF: RequiredChoiceIntermediateChoiceEnum("leaf"),
}

func (obj *requiredChoiceIntermediate) Choice() RequiredChoiceIntermediateChoiceEnum {
	return RequiredChoiceIntermediateChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *requiredChoiceIntermediate) setChoice(value RequiredChoiceIntermediateChoiceEnum) RequiredChoiceIntermediate {
	intValue, ok := openapi.RequiredChoiceIntermediate_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RequiredChoiceIntermediateChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.RequiredChoiceIntermediate_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Leaf = nil
	obj.leafHolder = nil
	obj.obj.FA = nil

	if value == RequiredChoiceIntermediateChoice.LEAF {
		obj.obj.Leaf = NewRequiredChoiceIntermeLeaf().msg()
	}

	return obj
}

// description is TBD
// FA returns a string
func (obj *requiredChoiceIntermediate) FA() string {

	if obj.obj.FA == nil {
		obj.setChoice(RequiredChoiceIntermediateChoice.F_A)
	}

	return *obj.obj.FA

}

// description is TBD
// FA returns a string
func (obj *requiredChoiceIntermediate) HasFA() bool {
	return obj.obj.FA != nil
}

// description is TBD
// SetFA sets the string value in the RequiredChoiceIntermediate object
func (obj *requiredChoiceIntermediate) SetFA(value string) RequiredChoiceIntermediate {
	obj.setChoice(RequiredChoiceIntermediateChoice.F_A)
	obj.obj.FA = &value
	return obj
}

// description is TBD
// Leaf returns a RequiredChoiceIntermeLeaf
func (obj *requiredChoiceIntermediate) Leaf() RequiredChoiceIntermeLeaf {
	if obj.obj.Leaf == nil {
		obj.setChoice(RequiredChoiceIntermediateChoice.LEAF)
	}
	if obj.leafHolder == nil {
		obj.leafHolder = &requiredChoiceIntermeLeaf{obj: obj.obj.Leaf}
	}
	return obj.leafHolder
}

// description is TBD
// Leaf returns a RequiredChoiceIntermeLeaf
func (obj *requiredChoiceIntermediate) HasLeaf() bool {
	return obj.obj.Leaf != nil
}

// description is TBD
// SetLeaf sets the RequiredChoiceIntermeLeaf value in the RequiredChoiceIntermediate object
func (obj *requiredChoiceIntermediate) SetLeaf(value RequiredChoiceIntermeLeaf) RequiredChoiceIntermediate {
	obj.setChoice(RequiredChoiceIntermediateChoice.LEAF)
	obj.leafHolder = nil
	obj.obj.Leaf = value.msg()

	return obj
}

func (obj *requiredChoiceIntermediate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface RequiredChoiceIntermediate")
	}

	if obj.obj.Leaf != nil {

		obj.Leaf().validateObj(vObj, set_default)
	}

}

func (obj *requiredChoiceIntermediate) setDefault() {
	var choices_set int = 0
	var choice RequiredChoiceIntermediateChoiceEnum

	if obj.obj.FA != nil {
		choices_set += 1
		choice = RequiredChoiceIntermediateChoice.F_A
	}

	if obj.obj.Leaf != nil {
		choices_set += 1
		choice = RequiredChoiceIntermediateChoice.LEAF
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in RequiredChoiceIntermediate")
			}
		} else {
			intVal := openapi.RequiredChoiceIntermediate_Choice_Enum_value[string(choice)]
			enumValue := openapi.RequiredChoiceIntermediate_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
