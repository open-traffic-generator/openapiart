package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Mandate *****
type mandate struct {
	validation
	obj          *openapi.Mandate
	marshaller   marshalMandate
	unMarshaller unMarshalMandate
}

func NewMandate() Mandate {
	obj := mandate{obj: &openapi.Mandate{}}
	obj.setDefault()
	return &obj
}

func (obj *mandate) msg() *openapi.Mandate {
	return obj.obj
}

func (obj *mandate) setMsg(msg *openapi.Mandate) Mandate {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmandate struct {
	obj *mandate
}

type marshalMandate interface {
	// ToProto marshals Mandate to protobuf object *openapi.Mandate
	ToProto() (*openapi.Mandate, error)
	// ToPbText marshals Mandate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Mandate to YAML text
	ToYaml() (string, error)
	// ToJson marshals Mandate to JSON text
	ToJson() (string, error)
}

type unMarshalmandate struct {
	obj *mandate
}

type unMarshalMandate interface {
	// FromProto unmarshals Mandate from protobuf object *openapi.Mandate
	FromProto(msg *openapi.Mandate) (Mandate, error)
	// FromPbText unmarshals Mandate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Mandate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Mandate from JSON text
	FromJson(value string) error
}

func (obj *mandate) Marshal() marshalMandate {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmandate{obj: obj}
	}
	return obj.marshaller
}

func (obj *mandate) Unmarshal() unMarshalMandate {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmandate{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmandate) ToProto() (*openapi.Mandate, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmandate) FromProto(msg *openapi.Mandate) (Mandate, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmandate) ToPbText() (string, error) {
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

func (m *unMarshalmandate) FromPbText(value string) error {
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

func (m *marshalmandate) ToYaml() (string, error) {
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

func (m *unMarshalmandate) FromYaml(value string) error {
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

func (m *marshalmandate) ToJson() (string, error) {
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

func (m *unMarshalmandate) FromJson(value string) error {
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

func (obj *mandate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mandate) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mandate) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mandate) Clone() (Mandate, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMandate()
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

// Mandate is object to Test required Parameter
type Mandate interface {
	Validation
	// msg marshals Mandate to protobuf object *openapi.Mandate
	// and doesn't set defaults
	msg() *openapi.Mandate
	// setMsg unmarshals Mandate from protobuf object *openapi.Mandate
	// and doesn't set defaults
	setMsg(*openapi.Mandate) Mandate
	// provides marshal interface
	Marshal() marshalMandate
	// provides unmarshal interface
	Unmarshal() unMarshalMandate
	// validate validates Mandate
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Mandate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequiredParam returns string, set in Mandate.
	RequiredParam() string
	// SetRequiredParam assigns string provided by user to Mandate
	SetRequiredParam(value string) Mandate
}

// description is TBD
// RequiredParam returns a string
func (obj *mandate) RequiredParam() string {

	return *obj.obj.RequiredParam

}

// description is TBD
// SetRequiredParam sets the string value in the Mandate object
func (obj *mandate) SetRequiredParam(value string) Mandate {

	obj.obj.RequiredParam = &value
	return obj
}

func (obj *mandate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// RequiredParam is required
	if obj.obj.RequiredParam == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RequiredParam is required field on interface Mandate")
	}
}

func (obj *mandate) setDefault() {

}
