package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** CommonResponseSuccess *****
type commonResponseSuccess struct {
	validation
	obj          *openapi.CommonResponseSuccess
	marshaller   marshalCommonResponseSuccess
	unMarshaller unMarshalCommonResponseSuccess
}

func NewCommonResponseSuccess() CommonResponseSuccess {
	obj := commonResponseSuccess{obj: &openapi.CommonResponseSuccess{}}
	obj.setDefault()
	return &obj
}

func (obj *commonResponseSuccess) msg() *openapi.CommonResponseSuccess {
	return obj.obj
}

func (obj *commonResponseSuccess) setMsg(msg *openapi.CommonResponseSuccess) CommonResponseSuccess {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalcommonResponseSuccess struct {
	obj *commonResponseSuccess
}

type marshalCommonResponseSuccess interface {
	// ToProto marshals CommonResponseSuccess to protobuf object *openapi.CommonResponseSuccess
	ToProto() (*openapi.CommonResponseSuccess, error)
	// ToPbText marshals CommonResponseSuccess to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CommonResponseSuccess to YAML text
	ToYaml() (string, error)
	// ToJson marshals CommonResponseSuccess to JSON text
	ToJson() (string, error)
}

type unMarshalcommonResponseSuccess struct {
	obj *commonResponseSuccess
}

type unMarshalCommonResponseSuccess interface {
	// FromProto unmarshals CommonResponseSuccess from protobuf object *openapi.CommonResponseSuccess
	FromProto(msg *openapi.CommonResponseSuccess) (CommonResponseSuccess, error)
	// FromPbText unmarshals CommonResponseSuccess from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CommonResponseSuccess from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CommonResponseSuccess from JSON text
	FromJson(value string) error
}

func (obj *commonResponseSuccess) Marshal() marshalCommonResponseSuccess {
	if obj.marshaller == nil {
		obj.marshaller = &marshalcommonResponseSuccess{obj: obj}
	}
	return obj.marshaller
}

func (obj *commonResponseSuccess) Unmarshal() unMarshalCommonResponseSuccess {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalcommonResponseSuccess{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalcommonResponseSuccess) ToProto() (*openapi.CommonResponseSuccess, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalcommonResponseSuccess) FromProto(msg *openapi.CommonResponseSuccess) (CommonResponseSuccess, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalcommonResponseSuccess) ToPbText() (string, error) {
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

func (m *unMarshalcommonResponseSuccess) FromPbText(value string) error {
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

func (m *marshalcommonResponseSuccess) ToYaml() (string, error) {
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

func (m *unMarshalcommonResponseSuccess) FromYaml(value string) error {
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

func (m *marshalcommonResponseSuccess) ToJson() (string, error) {
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

func (m *unMarshalcommonResponseSuccess) FromJson(value string) error {
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

func (obj *commonResponseSuccess) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *commonResponseSuccess) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *commonResponseSuccess) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *commonResponseSuccess) Clone() (CommonResponseSuccess, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCommonResponseSuccess()
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

// CommonResponseSuccess is description is TBD
type CommonResponseSuccess interface {
	Validation
	// msg marshals CommonResponseSuccess to protobuf object *openapi.CommonResponseSuccess
	// and doesn't set defaults
	msg() *openapi.CommonResponseSuccess
	// setMsg unmarshals CommonResponseSuccess from protobuf object *openapi.CommonResponseSuccess
	// and doesn't set defaults
	setMsg(*openapi.CommonResponseSuccess) CommonResponseSuccess
	// provides marshal interface
	Marshal() marshalCommonResponseSuccess
	// provides unmarshal interface
	Unmarshal() unMarshalCommonResponseSuccess
	// validate validates CommonResponseSuccess
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CommonResponseSuccess, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Message returns string, set in CommonResponseSuccess.
	Message() string
	// SetMessage assigns string provided by user to CommonResponseSuccess
	SetMessage(value string) CommonResponseSuccess
	// HasMessage checks if Message has been set in CommonResponseSuccess
	HasMessage() bool
}

// description is TBD
// Message returns a string
func (obj *commonResponseSuccess) Message() string {

	return *obj.obj.Message

}

// description is TBD
// Message returns a string
func (obj *commonResponseSuccess) HasMessage() bool {
	return obj.obj.Message != nil
}

// description is TBD
// SetMessage sets the string value in the CommonResponseSuccess object
func (obj *commonResponseSuccess) SetMessage(value string) CommonResponseSuccess {

	obj.obj.Message = &value
	return obj
}

func (obj *commonResponseSuccess) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *commonResponseSuccess) setDefault() {

}
