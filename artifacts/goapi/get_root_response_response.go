package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetRootResponseResponse *****
type getRootResponseResponse struct {
	validation
	obj                         *openapi.GetRootResponseResponse
	marshaller                  marshalGetRootResponseResponse
	unMarshaller                unMarshalGetRootResponseResponse
	commonResponseSuccessHolder CommonResponseSuccess
}

func NewGetRootResponseResponse() GetRootResponseResponse {
	obj := getRootResponseResponse{obj: &openapi.GetRootResponseResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getRootResponseResponse) msg() *openapi.GetRootResponseResponse {
	return obj.obj
}

func (obj *getRootResponseResponse) setMsg(msg *openapi.GetRootResponseResponse) GetRootResponseResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetRootResponseResponse struct {
	obj *getRootResponseResponse
}

type marshalGetRootResponseResponse interface {
	// ToProto marshals GetRootResponseResponse to protobuf object *openapi.GetRootResponseResponse
	ToProto() (*openapi.GetRootResponseResponse, error)
	// ToPbText marshals GetRootResponseResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetRootResponseResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetRootResponseResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetRootResponseResponse struct {
	obj *getRootResponseResponse
}

type unMarshalGetRootResponseResponse interface {
	// FromProto unmarshals GetRootResponseResponse from protobuf object *openapi.GetRootResponseResponse
	FromProto(msg *openapi.GetRootResponseResponse) (GetRootResponseResponse, error)
	// FromPbText unmarshals GetRootResponseResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetRootResponseResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetRootResponseResponse from JSON text
	FromJson(value string) error
}

func (obj *getRootResponseResponse) Marshal() marshalGetRootResponseResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetRootResponseResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getRootResponseResponse) Unmarshal() unMarshalGetRootResponseResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetRootResponseResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetRootResponseResponse) ToProto() (*openapi.GetRootResponseResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetRootResponseResponse) FromProto(msg *openapi.GetRootResponseResponse) (GetRootResponseResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetRootResponseResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetRootResponseResponse) FromPbText(value string) error {
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

func (m *marshalgetRootResponseResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetRootResponseResponse) FromYaml(value string) error {
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

func (m *marshalgetRootResponseResponse) ToJson() (string, error) {
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

func (m *unMarshalgetRootResponseResponse) FromJson(value string) error {
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

func (obj *getRootResponseResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getRootResponseResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getRootResponseResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getRootResponseResponse) Clone() (GetRootResponseResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetRootResponseResponse()
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

func (obj *getRootResponseResponse) setNil() {
	obj.commonResponseSuccessHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetRootResponseResponse is description is TBD
type GetRootResponseResponse interface {
	Validation
	// msg marshals GetRootResponseResponse to protobuf object *openapi.GetRootResponseResponse
	// and doesn't set defaults
	msg() *openapi.GetRootResponseResponse
	// setMsg unmarshals GetRootResponseResponse from protobuf object *openapi.GetRootResponseResponse
	// and doesn't set defaults
	setMsg(*openapi.GetRootResponseResponse) GetRootResponseResponse
	// provides marshal interface
	Marshal() marshalGetRootResponseResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetRootResponseResponse
	// validate validates GetRootResponseResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetRootResponseResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonResponseSuccess returns CommonResponseSuccess, set in GetRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	CommonResponseSuccess() CommonResponseSuccess
	// SetCommonResponseSuccess assigns CommonResponseSuccess provided by user to GetRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	SetCommonResponseSuccess(value CommonResponseSuccess) GetRootResponseResponse
	// HasCommonResponseSuccess checks if CommonResponseSuccess has been set in GetRootResponseResponse
	HasCommonResponseSuccess() bool
	setNil()
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *getRootResponseResponse) CommonResponseSuccess() CommonResponseSuccess {
	if obj.obj.CommonResponseSuccess == nil {
		obj.obj.CommonResponseSuccess = NewCommonResponseSuccess().msg()
	}
	if obj.commonResponseSuccessHolder == nil {
		obj.commonResponseSuccessHolder = &commonResponseSuccess{obj: obj.obj.CommonResponseSuccess}
	}
	return obj.commonResponseSuccessHolder
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *getRootResponseResponse) HasCommonResponseSuccess() bool {
	return obj.obj.CommonResponseSuccess != nil
}

// description is TBD
// SetCommonResponseSuccess sets the CommonResponseSuccess value in the GetRootResponseResponse object
func (obj *getRootResponseResponse) SetCommonResponseSuccess(value CommonResponseSuccess) GetRootResponseResponse {

	obj.commonResponseSuccessHolder = nil
	obj.obj.CommonResponseSuccess = value.msg()

	return obj
}

func (obj *getRootResponseResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonResponseSuccess != nil {

		obj.CommonResponseSuccess().validateObj(vObj, set_default)
	}

}

func (obj *getRootResponseResponse) setDefault() {

}
