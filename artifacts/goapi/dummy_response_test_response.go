package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DummyResponseTestResponse *****
type dummyResponseTestResponse struct {
	validation
	obj          *openapi.DummyResponseTestResponse
	marshaller   marshalDummyResponseTestResponse
	unMarshaller unMarshalDummyResponseTestResponse
}

func NewDummyResponseTestResponse() DummyResponseTestResponse {
	obj := dummyResponseTestResponse{obj: &openapi.DummyResponseTestResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *dummyResponseTestResponse) msg() *openapi.DummyResponseTestResponse {
	return obj.obj
}

func (obj *dummyResponseTestResponse) setMsg(msg *openapi.DummyResponseTestResponse) DummyResponseTestResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldummyResponseTestResponse struct {
	obj *dummyResponseTestResponse
}

type marshalDummyResponseTestResponse interface {
	// ToProto marshals DummyResponseTestResponse to protobuf object *openapi.DummyResponseTestResponse
	ToProto() (*openapi.DummyResponseTestResponse, error)
	// ToPbText marshals DummyResponseTestResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DummyResponseTestResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals DummyResponseTestResponse to JSON text
	ToJson() (string, error)
}

type unMarshaldummyResponseTestResponse struct {
	obj *dummyResponseTestResponse
}

type unMarshalDummyResponseTestResponse interface {
	// FromProto unmarshals DummyResponseTestResponse from protobuf object *openapi.DummyResponseTestResponse
	FromProto(msg *openapi.DummyResponseTestResponse) (DummyResponseTestResponse, error)
	// FromPbText unmarshals DummyResponseTestResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DummyResponseTestResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DummyResponseTestResponse from JSON text
	FromJson(value string) error
}

func (obj *dummyResponseTestResponse) Marshal() marshalDummyResponseTestResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldummyResponseTestResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *dummyResponseTestResponse) Unmarshal() unMarshalDummyResponseTestResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldummyResponseTestResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldummyResponseTestResponse) ToProto() (*openapi.DummyResponseTestResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldummyResponseTestResponse) FromProto(msg *openapi.DummyResponseTestResponse) (DummyResponseTestResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldummyResponseTestResponse) ToPbText() (string, error) {
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

func (m *unMarshaldummyResponseTestResponse) FromPbText(value string) error {
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

func (m *marshaldummyResponseTestResponse) ToYaml() (string, error) {
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

func (m *unMarshaldummyResponseTestResponse) FromYaml(value string) error {
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

func (m *marshaldummyResponseTestResponse) ToJson() (string, error) {
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

func (m *unMarshaldummyResponseTestResponse) FromJson(value string) error {
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

func (obj *dummyResponseTestResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dummyResponseTestResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dummyResponseTestResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dummyResponseTestResponse) Clone() (DummyResponseTestResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDummyResponseTestResponse()
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

// DummyResponseTestResponse is description is TBD
type DummyResponseTestResponse interface {
	Validation
	// msg marshals DummyResponseTestResponse to protobuf object *openapi.DummyResponseTestResponse
	// and doesn't set defaults
	msg() *openapi.DummyResponseTestResponse
	// setMsg unmarshals DummyResponseTestResponse from protobuf object *openapi.DummyResponseTestResponse
	// and doesn't set defaults
	setMsg(*openapi.DummyResponseTestResponse) DummyResponseTestResponse
	// provides marshal interface
	Marshal() marshalDummyResponseTestResponse
	// provides unmarshal interface
	Unmarshal() unMarshalDummyResponseTestResponse
	// validate validates DummyResponseTestResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DummyResponseTestResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseString returns string, set in DummyResponseTestResponse.
	ResponseString() string
	// SetResponseString assigns string provided by user to DummyResponseTestResponse
	SetResponseString(value string) DummyResponseTestResponse
	// HasResponseString checks if ResponseString has been set in DummyResponseTestResponse
	HasResponseString() bool
}

// description is TBD
// ResponseString returns a string
func (obj *dummyResponseTestResponse) ResponseString() string {
	return obj.obj.String_
}

// description is TBD
// ResponseString returns a string
func (obj *dummyResponseTestResponse) HasResponseString() bool {
	return obj.obj.String_ != ""
}

// description is TBD
// SetResponseString sets the string value in the DummyResponseTestResponse object
func (obj *dummyResponseTestResponse) SetResponseString(value string) DummyResponseTestResponse {
	obj.obj.String_ = value
	return obj
}

func (obj *dummyResponseTestResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dummyResponseTestResponse) setDefault() {

}
