package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredChoice *****
type requiredChoice struct {
	validation
	obj          *openapi.RequiredChoice
	marshaller   marshalRequiredChoice
	unMarshaller unMarshalRequiredChoice
	leafHolder   LeafVal
}

func NewRequiredChoice() RequiredChoice {
	obj := requiredChoice{obj: &openapi.RequiredChoice{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoice) msg() *openapi.RequiredChoice {
	return obj.obj
}

func (obj *requiredChoice) setMsg(msg *openapi.RequiredChoice) RequiredChoice {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredChoice struct {
	obj *requiredChoice
}

type marshalRequiredChoice interface {
	// ToProto marshals RequiredChoice to protobuf object *openapi.RequiredChoice
	ToProto() (*openapi.RequiredChoice, error)
	// ToPbText marshals RequiredChoice to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoice to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoice to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredChoice struct {
	obj *requiredChoice
}

type unMarshalRequiredChoice interface {
	// FromProto unmarshals RequiredChoice from protobuf object *openapi.RequiredChoice
	FromProto(msg *openapi.RequiredChoice) (RequiredChoice, error)
	// FromPbText unmarshals RequiredChoice from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoice from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoice from JSON text
	FromJson(value string) error
}

func (obj *requiredChoice) Marshal() marshalRequiredChoice {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredChoice{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredChoice) Unmarshal() unMarshalRequiredChoice {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredChoice{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredChoice) ToProto() (*openapi.RequiredChoice, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredChoice) FromProto(msg *openapi.RequiredChoice) (RequiredChoice, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredChoice) ToPbText() (string, error) {
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

func (m *unMarshalrequiredChoice) FromPbText(value string) error {
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

func (m *marshalrequiredChoice) ToYaml() (string, error) {
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

func (m *unMarshalrequiredChoice) FromYaml(value string) error {
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

func (m *marshalrequiredChoice) ToJson() (string, error) {
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

func (m *unMarshalrequiredChoice) FromJson(value string) error {
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

func (obj *requiredChoice) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoice) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoice) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoice) Clone() (RequiredChoice, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoice()
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

func (obj *requiredChoice) setNil() {
	obj.leafHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RequiredChoice is description is TBD
type RequiredChoice interface {
	Validation
	// msg marshals RequiredChoice to protobuf object *openapi.RequiredChoice
	// and doesn't set defaults
	msg() *openapi.RequiredChoice
	// setMsg unmarshals RequiredChoice from protobuf object *openapi.RequiredChoice
	// and doesn't set defaults
	setMsg(*openapi.RequiredChoice) RequiredChoice
	// provides marshal interface
	Marshal() marshalRequiredChoice
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredChoice
	// validate validates RequiredChoice
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoice, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns RequiredChoiceChoiceEnum, set in RequiredChoice
	Choice() RequiredChoiceChoiceEnum
	// setChoice assigns RequiredChoiceChoiceEnum provided by user to RequiredChoice
	setChoice(value RequiredChoiceChoiceEnum) RequiredChoice
	// StrVal returns string, set in RequiredChoice.
	StrVal() string
	// SetStrVal assigns string provided by user to RequiredChoice
	SetStrVal(value string) RequiredChoice
	// HasStrVal checks if StrVal has been set in RequiredChoice
	HasStrVal() bool
	// Leaf returns LeafVal, set in RequiredChoice.
	// LeafVal is description is TBD
	Leaf() LeafVal
	// SetLeaf assigns LeafVal provided by user to RequiredChoice.
	// LeafVal is description is TBD
	SetLeaf(value LeafVal) RequiredChoice
	// HasLeaf checks if Leaf has been set in RequiredChoice
	HasLeaf() bool
	setNil()
}

type RequiredChoiceChoiceEnum string

// Enum of Choice on RequiredChoice
var RequiredChoiceChoice = struct {
	STR_VAL RequiredChoiceChoiceEnum
	LEAF    RequiredChoiceChoiceEnum
}{
	STR_VAL: RequiredChoiceChoiceEnum("str_val"),
	LEAF:    RequiredChoiceChoiceEnum("leaf"),
}

func (obj *requiredChoice) Choice() RequiredChoiceChoiceEnum {
	return RequiredChoiceChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *requiredChoice) setChoice(value RequiredChoiceChoiceEnum) RequiredChoice {
	intValue, ok := openapi.RequiredChoice_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RequiredChoiceChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.RequiredChoice_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Leaf = nil
	obj.leafHolder = nil
	obj.obj.StrVal = nil

	if value == RequiredChoiceChoice.STR_VAL {
		defaultValue := "some string"
		obj.obj.StrVal = &defaultValue
	}

	if value == RequiredChoiceChoice.LEAF {
		obj.obj.Leaf = NewLeafVal().msg()
	}

	return obj
}

// description is TBD
// StrVal returns a string
func (obj *requiredChoice) StrVal() string {

	if obj.obj.StrVal == nil {
		obj.setChoice(RequiredChoiceChoice.STR_VAL)
	}

	return *obj.obj.StrVal

}

// description is TBD
// StrVal returns a string
func (obj *requiredChoice) HasStrVal() bool {
	return obj.obj.StrVal != nil
}

// description is TBD
// SetStrVal sets the string value in the RequiredChoice object
func (obj *requiredChoice) SetStrVal(value string) RequiredChoice {
	obj.setChoice(RequiredChoiceChoice.STR_VAL)
	obj.obj.StrVal = &value
	return obj
}

// description is TBD
// Leaf returns a LeafVal
func (obj *requiredChoice) Leaf() LeafVal {
	if obj.obj.Leaf == nil {
		obj.setChoice(RequiredChoiceChoice.LEAF)
	}
	if obj.leafHolder == nil {
		obj.leafHolder = &leafVal{obj: obj.obj.Leaf}
	}
	return obj.leafHolder
}

// description is TBD
// Leaf returns a LeafVal
func (obj *requiredChoice) HasLeaf() bool {
	return obj.obj.Leaf != nil
}

// description is TBD
// SetLeaf sets the LeafVal value in the RequiredChoice object
func (obj *requiredChoice) SetLeaf(value LeafVal) RequiredChoice {
	obj.setChoice(RequiredChoiceChoice.LEAF)
	obj.leafHolder = nil
	obj.obj.Leaf = value.msg()

	return obj
}

func (obj *requiredChoice) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface RequiredChoice")
	}

	if obj.obj.Leaf != nil {

		obj.Leaf().validateObj(vObj, set_default)
	}

}

func (obj *requiredChoice) setDefault() {
	if obj.obj.StrVal == nil {
		obj.SetStrVal("some string")
	}

}
