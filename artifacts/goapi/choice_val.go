package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChoiceVal *****
type choiceVal struct {
	validation
	obj            *openapi.ChoiceVal
	marshaller     marshalChoiceVal
	unMarshaller   unMarshalChoiceVal
	mixedValHolder MixedVal
}

func NewChoiceVal() ChoiceVal {
	obj := choiceVal{obj: &openapi.ChoiceVal{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceVal) msg() *openapi.ChoiceVal {
	return obj.obj
}

func (obj *choiceVal) setMsg(msg *openapi.ChoiceVal) ChoiceVal {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchoiceVal struct {
	obj *choiceVal
}

type marshalChoiceVal interface {
	// ToProto marshals ChoiceVal to protobuf object *openapi.ChoiceVal
	ToProto() (*openapi.ChoiceVal, error)
	// ToPbText marshals ChoiceVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceVal to JSON text
	ToJson() (string, error)
}

type unMarshalchoiceVal struct {
	obj *choiceVal
}

type unMarshalChoiceVal interface {
	// FromProto unmarshals ChoiceVal from protobuf object *openapi.ChoiceVal
	FromProto(msg *openapi.ChoiceVal) (ChoiceVal, error)
	// FromPbText unmarshals ChoiceVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceVal from JSON text
	FromJson(value string) error
}

func (obj *choiceVal) Marshal() marshalChoiceVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchoiceVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *choiceVal) Unmarshal() unMarshalChoiceVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchoiceVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchoiceVal) ToProto() (*openapi.ChoiceVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchoiceVal) FromProto(msg *openapi.ChoiceVal) (ChoiceVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchoiceVal) ToPbText() (string, error) {
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

func (m *unMarshalchoiceVal) FromPbText(value string) error {
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

func (m *marshalchoiceVal) ToYaml() (string, error) {
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

func (m *unMarshalchoiceVal) FromYaml(value string) error {
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

func (m *marshalchoiceVal) ToJson() (string, error) {
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

func (m *unMarshalchoiceVal) FromJson(value string) error {
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

func (obj *choiceVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceVal) Clone() (ChoiceVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceVal()
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

func (obj *choiceVal) setNil() {
	obj.mixedValHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChoiceVal is description is TBD
type ChoiceVal interface {
	Validation
	// msg marshals ChoiceVal to protobuf object *openapi.ChoiceVal
	// and doesn't set defaults
	msg() *openapi.ChoiceVal
	// setMsg unmarshals ChoiceVal from protobuf object *openapi.ChoiceVal
	// and doesn't set defaults
	setMsg(*openapi.ChoiceVal) ChoiceVal
	// provides marshal interface
	Marshal() marshalChoiceVal
	// provides unmarshal interface
	Unmarshal() unMarshalChoiceVal
	// validate validates ChoiceVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// MixedVal returns MixedVal, set in ChoiceVal.
	// MixedVal is description is TBD
	MixedVal() MixedVal
	// SetMixedVal assigns MixedVal provided by user to ChoiceVal.
	// MixedVal is description is TBD
	SetMixedVal(value MixedVal) ChoiceVal
	// HasMixedVal checks if MixedVal has been set in ChoiceVal
	HasMixedVal() bool
	setNil()
}

// description is TBD
// MixedVal returns a MixedVal
func (obj *choiceVal) MixedVal() MixedVal {
	if obj.obj.MixedVal == nil {
		obj.obj.MixedVal = NewMixedVal().msg()
	}
	if obj.mixedValHolder == nil {
		obj.mixedValHolder = &mixedVal{obj: obj.obj.MixedVal}
	}
	return obj.mixedValHolder
}

// description is TBD
// MixedVal returns a MixedVal
func (obj *choiceVal) HasMixedVal() bool {
	return obj.obj.MixedVal != nil
}

// description is TBD
// SetMixedVal sets the MixedVal value in the ChoiceVal object
func (obj *choiceVal) SetMixedVal(value MixedVal) ChoiceVal {

	obj.mixedValHolder = nil
	obj.obj.MixedVal = value.msg()

	return obj
}

func (obj *choiceVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.MixedVal != nil {

		obj.MixedVal().validateObj(vObj, set_default)
	}

}

func (obj *choiceVal) setDefault() {

}
