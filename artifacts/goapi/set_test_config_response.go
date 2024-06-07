package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SetTestConfigResponse *****
type setTestConfigResponse struct {
	validation
	obj          *openapi.SetTestConfigResponse
	marshaller   marshalSetTestConfigResponse
	unMarshaller unMarshalSetTestConfigResponse
}

func NewSetTestConfigResponse() SetTestConfigResponse {
	obj := setTestConfigResponse{obj: &openapi.SetTestConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *setTestConfigResponse) msg() *openapi.SetTestConfigResponse {
	return obj.obj
}

func (obj *setTestConfigResponse) setMsg(msg *openapi.SetTestConfigResponse) SetTestConfigResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsetTestConfigResponse struct {
	obj *setTestConfigResponse
}

type marshalSetTestConfigResponse interface {
	// ToProto marshals SetTestConfigResponse to protobuf object *openapi.SetTestConfigResponse
	ToProto() (*openapi.SetTestConfigResponse, error)
	// ToPbText marshals SetTestConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SetTestConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals SetTestConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalsetTestConfigResponse struct {
	obj *setTestConfigResponse
}

type unMarshalSetTestConfigResponse interface {
	// FromProto unmarshals SetTestConfigResponse from protobuf object *openapi.SetTestConfigResponse
	FromProto(msg *openapi.SetTestConfigResponse) (SetTestConfigResponse, error)
	// FromPbText unmarshals SetTestConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SetTestConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SetTestConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *setTestConfigResponse) Marshal() marshalSetTestConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsetTestConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *setTestConfigResponse) Unmarshal() unMarshalSetTestConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsetTestConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsetTestConfigResponse) ToProto() (*openapi.SetTestConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsetTestConfigResponse) FromProto(msg *openapi.SetTestConfigResponse) (SetTestConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsetTestConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalsetTestConfigResponse) FromPbText(value string) error {
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

func (m *marshalsetTestConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalsetTestConfigResponse) FromYaml(value string) error {
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

func (m *marshalsetTestConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalsetTestConfigResponse) FromJson(value string) error {
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

func (obj *setTestConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *setTestConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *setTestConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *setTestConfigResponse) Clone() (SetTestConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSetTestConfigResponse()
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

// SetTestConfigResponse is description is TBD
type SetTestConfigResponse interface {
	Validation
	// msg marshals SetTestConfigResponse to protobuf object *openapi.SetTestConfigResponse
	// and doesn't set defaults
	msg() *openapi.SetTestConfigResponse
	// setMsg unmarshals SetTestConfigResponse from protobuf object *openapi.SetTestConfigResponse
	// and doesn't set defaults
	setMsg(*openapi.SetTestConfigResponse) SetTestConfigResponse
	// provides marshal interface
	Marshal() marshalSetTestConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalSetTestConfigResponse
	// validate validates SetTestConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SetTestConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseBytes returns []byte, set in SetTestConfigResponse.
	ResponseBytes() []byte
	// SetResponseBytes assigns []byte provided by user to SetTestConfigResponse
	SetResponseBytes(value []byte) SetTestConfigResponse
	// HasResponseBytes checks if ResponseBytes has been set in SetTestConfigResponse
	HasResponseBytes() bool
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setTestConfigResponse) ResponseBytes() []byte {

	return obj.obj.ResponseBytes
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setTestConfigResponse) HasResponseBytes() bool {
	return obj.obj.ResponseBytes != nil
}

// description is TBD
// SetResponseBytes sets the []byte value in the SetTestConfigResponse object
func (obj *setTestConfigResponse) SetResponseBytes(value []byte) SetTestConfigResponse {

	obj.obj.ResponseBytes = value
	return obj
}

func (obj *setTestConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *setTestConfigResponse) setDefault() {

}
