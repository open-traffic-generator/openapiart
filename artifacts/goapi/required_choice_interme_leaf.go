package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredChoiceIntermeLeaf *****
type requiredChoiceIntermeLeaf struct {
	validation
	obj          *openapi.RequiredChoiceIntermeLeaf
	marshaller   marshalRequiredChoiceIntermeLeaf
	unMarshaller unMarshalRequiredChoiceIntermeLeaf
}

func NewRequiredChoiceIntermeLeaf() RequiredChoiceIntermeLeaf {
	obj := requiredChoiceIntermeLeaf{obj: &openapi.RequiredChoiceIntermeLeaf{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceIntermeLeaf) msg() *openapi.RequiredChoiceIntermeLeaf {
	return obj.obj
}

func (obj *requiredChoiceIntermeLeaf) setMsg(msg *openapi.RequiredChoiceIntermeLeaf) RequiredChoiceIntermeLeaf {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredChoiceIntermeLeaf struct {
	obj *requiredChoiceIntermeLeaf
}

type marshalRequiredChoiceIntermeLeaf interface {
	// ToProto marshals RequiredChoiceIntermeLeaf to protobuf object *openapi.RequiredChoiceIntermeLeaf
	ToProto() (*openapi.RequiredChoiceIntermeLeaf, error)
	// ToPbText marshals RequiredChoiceIntermeLeaf to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceIntermeLeaf to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceIntermeLeaf to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredChoiceIntermeLeaf struct {
	obj *requiredChoiceIntermeLeaf
}

type unMarshalRequiredChoiceIntermeLeaf interface {
	// FromProto unmarshals RequiredChoiceIntermeLeaf from protobuf object *openapi.RequiredChoiceIntermeLeaf
	FromProto(msg *openapi.RequiredChoiceIntermeLeaf) (RequiredChoiceIntermeLeaf, error)
	// FromPbText unmarshals RequiredChoiceIntermeLeaf from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceIntermeLeaf from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceIntermeLeaf from JSON text
	FromJson(value string) error
}

func (obj *requiredChoiceIntermeLeaf) Marshal() marshalRequiredChoiceIntermeLeaf {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredChoiceIntermeLeaf{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredChoiceIntermeLeaf) Unmarshal() unMarshalRequiredChoiceIntermeLeaf {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredChoiceIntermeLeaf{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredChoiceIntermeLeaf) ToProto() (*openapi.RequiredChoiceIntermeLeaf, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredChoiceIntermeLeaf) FromProto(msg *openapi.RequiredChoiceIntermeLeaf) (RequiredChoiceIntermeLeaf, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredChoiceIntermeLeaf) ToPbText() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermeLeaf) FromPbText(value string) error {
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

func (m *marshalrequiredChoiceIntermeLeaf) ToYaml() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermeLeaf) FromYaml(value string) error {
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

func (m *marshalrequiredChoiceIntermeLeaf) ToJson() (string, error) {
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

func (m *unMarshalrequiredChoiceIntermeLeaf) FromJson(value string) error {
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

func (obj *requiredChoiceIntermeLeaf) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermeLeaf) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermeLeaf) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceIntermeLeaf) Clone() (RequiredChoiceIntermeLeaf, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceIntermeLeaf()
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

// RequiredChoiceIntermeLeaf is description is TBD
type RequiredChoiceIntermeLeaf interface {
	Validation
	// msg marshals RequiredChoiceIntermeLeaf to protobuf object *openapi.RequiredChoiceIntermeLeaf
	// and doesn't set defaults
	msg() *openapi.RequiredChoiceIntermeLeaf
	// setMsg unmarshals RequiredChoiceIntermeLeaf from protobuf object *openapi.RequiredChoiceIntermeLeaf
	// and doesn't set defaults
	setMsg(*openapi.RequiredChoiceIntermeLeaf) RequiredChoiceIntermeLeaf
	// provides marshal interface
	Marshal() marshalRequiredChoiceIntermeLeaf
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredChoiceIntermeLeaf
	// validate validates RequiredChoiceIntermeLeaf
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceIntermeLeaf, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RequiredChoiceIntermeLeaf.
	Name() string
	// SetName assigns string provided by user to RequiredChoiceIntermeLeaf
	SetName(value string) RequiredChoiceIntermeLeaf
	// HasName checks if Name has been set in RequiredChoiceIntermeLeaf
	HasName() bool
}

// description is TBD
// Name returns a string
func (obj *requiredChoiceIntermeLeaf) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *requiredChoiceIntermeLeaf) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the RequiredChoiceIntermeLeaf object
func (obj *requiredChoiceIntermeLeaf) SetName(value string) RequiredChoiceIntermeLeaf {

	obj.obj.Name = &value
	return obj
}

func (obj *requiredChoiceIntermeLeaf) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *requiredChoiceIntermeLeaf) setDefault() {

}
