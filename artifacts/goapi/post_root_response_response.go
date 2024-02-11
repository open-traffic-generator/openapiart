package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PostRootResponseResponse *****
type postRootResponseResponse struct {
	validation
	obj                         *openapi.PostRootResponseResponse
	marshaller                  marshalPostRootResponseResponse
	unMarshaller                unMarshalPostRootResponseResponse
	commonResponseSuccessHolder CommonResponseSuccess
}

func NewPostRootResponseResponse() PostRootResponseResponse {
	obj := postRootResponseResponse{obj: &openapi.PostRootResponseResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *postRootResponseResponse) msg() *openapi.PostRootResponseResponse {
	return obj.obj
}

func (obj *postRootResponseResponse) setMsg(msg *openapi.PostRootResponseResponse) PostRootResponseResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpostRootResponseResponse struct {
	obj *postRootResponseResponse
}

type marshalPostRootResponseResponse interface {
	// ToProto marshals PostRootResponseResponse to protobuf object *openapi.PostRootResponseResponse
	ToProto() (*openapi.PostRootResponseResponse, error)
	// ToPbText marshals PostRootResponseResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PostRootResponseResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals PostRootResponseResponse to JSON text
	ToJson() (string, error)
}

type unMarshalpostRootResponseResponse struct {
	obj *postRootResponseResponse
}

type unMarshalPostRootResponseResponse interface {
	// FromProto unmarshals PostRootResponseResponse from protobuf object *openapi.PostRootResponseResponse
	FromProto(msg *openapi.PostRootResponseResponse) (PostRootResponseResponse, error)
	// FromPbText unmarshals PostRootResponseResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PostRootResponseResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PostRootResponseResponse from JSON text
	FromJson(value string) error
}

func (obj *postRootResponseResponse) Marshal() marshalPostRootResponseResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpostRootResponseResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *postRootResponseResponse) Unmarshal() unMarshalPostRootResponseResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpostRootResponseResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpostRootResponseResponse) ToProto() (*openapi.PostRootResponseResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpostRootResponseResponse) FromProto(msg *openapi.PostRootResponseResponse) (PostRootResponseResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpostRootResponseResponse) ToPbText() (string, error) {
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

func (m *unMarshalpostRootResponseResponse) FromPbText(value string) error {
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

func (m *marshalpostRootResponseResponse) ToYaml() (string, error) {
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

func (m *unMarshalpostRootResponseResponse) FromYaml(value string) error {
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

func (m *marshalpostRootResponseResponse) ToJson() (string, error) {
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

func (m *unMarshalpostRootResponseResponse) FromJson(value string) error {
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

func (obj *postRootResponseResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *postRootResponseResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *postRootResponseResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *postRootResponseResponse) Clone() (PostRootResponseResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPostRootResponseResponse()
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

func (obj *postRootResponseResponse) setNil() {
	obj.commonResponseSuccessHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PostRootResponseResponse is description is TBD
type PostRootResponseResponse interface {
	Validation
	// msg marshals PostRootResponseResponse to protobuf object *openapi.PostRootResponseResponse
	// and doesn't set defaults
	msg() *openapi.PostRootResponseResponse
	// setMsg unmarshals PostRootResponseResponse from protobuf object *openapi.PostRootResponseResponse
	// and doesn't set defaults
	setMsg(*openapi.PostRootResponseResponse) PostRootResponseResponse
	// provides marshal interface
	Marshal() marshalPostRootResponseResponse
	// provides unmarshal interface
	Unmarshal() unMarshalPostRootResponseResponse
	// validate validates PostRootResponseResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PostRootResponseResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonResponseSuccess returns CommonResponseSuccess, set in PostRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	CommonResponseSuccess() CommonResponseSuccess
	// SetCommonResponseSuccess assigns CommonResponseSuccess provided by user to PostRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	SetCommonResponseSuccess(value CommonResponseSuccess) PostRootResponseResponse
	// HasCommonResponseSuccess checks if CommonResponseSuccess has been set in PostRootResponseResponse
	HasCommonResponseSuccess() bool
	setNil()
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *postRootResponseResponse) CommonResponseSuccess() CommonResponseSuccess {
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
func (obj *postRootResponseResponse) HasCommonResponseSuccess() bool {
	return obj.obj.CommonResponseSuccess != nil
}

// description is TBD
// SetCommonResponseSuccess sets the CommonResponseSuccess value in the PostRootResponseResponse object
func (obj *postRootResponseResponse) SetCommonResponseSuccess(value CommonResponseSuccess) PostRootResponseResponse {

	obj.commonResponseSuccessHolder = nil
	obj.obj.CommonResponseSuccess = value.msg()

	return obj
}

func (obj *postRootResponseResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonResponseSuccess != nil {

		obj.CommonResponseSuccess().validateObj(vObj, set_default)
	}

}

func (obj *postRootResponseResponse) setDefault() {

}
