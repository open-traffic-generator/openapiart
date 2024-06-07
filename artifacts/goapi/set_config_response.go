package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SetConfigResponse *****
type setConfigResponse struct {
	validation
	obj          *openapi.SetConfigResponse
	marshaller   marshalSetConfigResponse
	unMarshaller unMarshalSetConfigResponse
}

func NewSetConfigResponse() SetConfigResponse {
	obj := setConfigResponse{obj: &openapi.SetConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *setConfigResponse) msg() *openapi.SetConfigResponse {
	return obj.obj
}

func (obj *setConfigResponse) setMsg(msg *openapi.SetConfigResponse) SetConfigResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsetConfigResponse struct {
	obj *setConfigResponse
}

type marshalSetConfigResponse interface {
	// ToProto marshals SetConfigResponse to protobuf object *openapi.SetConfigResponse
	ToProto() (*openapi.SetConfigResponse, error)
	// ToPbText marshals SetConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SetConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals SetConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalsetConfigResponse struct {
	obj *setConfigResponse
}

type unMarshalSetConfigResponse interface {
	// FromProto unmarshals SetConfigResponse from protobuf object *openapi.SetConfigResponse
	FromProto(msg *openapi.SetConfigResponse) (SetConfigResponse, error)
	// FromPbText unmarshals SetConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SetConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SetConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *setConfigResponse) Marshal() marshalSetConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsetConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *setConfigResponse) Unmarshal() unMarshalSetConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsetConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsetConfigResponse) ToProto() (*openapi.SetConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsetConfigResponse) FromProto(msg *openapi.SetConfigResponse) (SetConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsetConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalsetConfigResponse) FromPbText(value string) error {
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

func (m *marshalsetConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalsetConfigResponse) FromYaml(value string) error {
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

func (m *marshalsetConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalsetConfigResponse) FromJson(value string) error {
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

func (obj *setConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *setConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *setConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *setConfigResponse) Clone() (SetConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSetConfigResponse()
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

// SetConfigResponse is description is TBD
type SetConfigResponse interface {
	Validation
	// msg marshals SetConfigResponse to protobuf object *openapi.SetConfigResponse
	// and doesn't set defaults
	msg() *openapi.SetConfigResponse
	// setMsg unmarshals SetConfigResponse from protobuf object *openapi.SetConfigResponse
	// and doesn't set defaults
	setMsg(*openapi.SetConfigResponse) SetConfigResponse
	// provides marshal interface
	Marshal() marshalSetConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalSetConfigResponse
	// validate validates SetConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SetConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseBytes returns []byte, set in SetConfigResponse.
	ResponseBytes() []byte
	// SetResponseBytes assigns []byte provided by user to SetConfigResponse
	SetResponseBytes(value []byte) SetConfigResponse
	// HasResponseBytes checks if ResponseBytes has been set in SetConfigResponse
	HasResponseBytes() bool
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setConfigResponse) ResponseBytes() []byte {

	return obj.obj.ResponseBytes
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setConfigResponse) HasResponseBytes() bool {
	return obj.obj.ResponseBytes != nil
}

// description is TBD
// SetResponseBytes sets the []byte value in the SetConfigResponse object
func (obj *setConfigResponse) SetResponseBytes(value []byte) SetConfigResponse {

	obj.obj.ResponseBytes = value
	return obj
}

func (obj *setConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *setConfigResponse) setDefault() {

}
