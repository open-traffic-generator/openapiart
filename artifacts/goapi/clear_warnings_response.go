package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ClearWarningsResponse *****
type clearWarningsResponse struct {
	validation
	obj          *openapi.ClearWarningsResponse
	marshaller   marshalClearWarningsResponse
	unMarshaller unMarshalClearWarningsResponse
}

func NewClearWarningsResponse() ClearWarningsResponse {
	obj := clearWarningsResponse{obj: &openapi.ClearWarningsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *clearWarningsResponse) msg() *openapi.ClearWarningsResponse {
	return obj.obj
}

func (obj *clearWarningsResponse) setMsg(msg *openapi.ClearWarningsResponse) ClearWarningsResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalclearWarningsResponse struct {
	obj *clearWarningsResponse
}

type marshalClearWarningsResponse interface {
	// ToProto marshals ClearWarningsResponse to protobuf object *openapi.ClearWarningsResponse
	ToProto() (*openapi.ClearWarningsResponse, error)
	// ToPbText marshals ClearWarningsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ClearWarningsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ClearWarningsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalclearWarningsResponse struct {
	obj *clearWarningsResponse
}

type unMarshalClearWarningsResponse interface {
	// FromProto unmarshals ClearWarningsResponse from protobuf object *openapi.ClearWarningsResponse
	FromProto(msg *openapi.ClearWarningsResponse) (ClearWarningsResponse, error)
	// FromPbText unmarshals ClearWarningsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ClearWarningsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ClearWarningsResponse from JSON text
	FromJson(value string) error
}

func (obj *clearWarningsResponse) Marshal() marshalClearWarningsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalclearWarningsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *clearWarningsResponse) Unmarshal() unMarshalClearWarningsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalclearWarningsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalclearWarningsResponse) ToProto() (*openapi.ClearWarningsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalclearWarningsResponse) FromProto(msg *openapi.ClearWarningsResponse) (ClearWarningsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalclearWarningsResponse) ToPbText() (string, error) {
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

func (m *unMarshalclearWarningsResponse) FromPbText(value string) error {
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

func (m *marshalclearWarningsResponse) ToYaml() (string, error) {
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

func (m *unMarshalclearWarningsResponse) FromYaml(value string) error {
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

func (m *marshalclearWarningsResponse) ToJson() (string, error) {
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

func (m *unMarshalclearWarningsResponse) FromJson(value string) error {
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

func (obj *clearWarningsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *clearWarningsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *clearWarningsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *clearWarningsResponse) Clone() (ClearWarningsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewClearWarningsResponse()
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

// ClearWarningsResponse is description is TBD
type ClearWarningsResponse interface {
	Validation
	// msg marshals ClearWarningsResponse to protobuf object *openapi.ClearWarningsResponse
	// and doesn't set defaults
	msg() *openapi.ClearWarningsResponse
	// setMsg unmarshals ClearWarningsResponse from protobuf object *openapi.ClearWarningsResponse
	// and doesn't set defaults
	setMsg(*openapi.ClearWarningsResponse) ClearWarningsResponse
	// provides marshal interface
	Marshal() marshalClearWarningsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalClearWarningsResponse
	// validate validates ClearWarningsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ClearWarningsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseString returns string, set in ClearWarningsResponse.
	ResponseString() string
	// SetResponseString assigns string provided by user to ClearWarningsResponse
	SetResponseString(value string) ClearWarningsResponse
	// HasResponseString checks if ResponseString has been set in ClearWarningsResponse
	HasResponseString() bool
}

// description is TBD
// ResponseString returns a string
func (obj *clearWarningsResponse) ResponseString() string {
	return obj.obj.String_
}

// description is TBD
// ResponseString returns a string
func (obj *clearWarningsResponse) HasResponseString() bool {
	return obj.obj.String_ != ""
}

// description is TBD
// SetResponseString sets the string value in the ClearWarningsResponse object
func (obj *clearWarningsResponse) SetResponseString(value string) ClearWarningsResponse {
	obj.obj.String_ = value
	return obj
}

func (obj *clearWarningsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *clearWarningsResponse) setDefault() {

}
