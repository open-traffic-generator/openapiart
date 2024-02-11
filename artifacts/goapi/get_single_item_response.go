package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetSingleItemResponse *****
type getSingleItemResponse struct {
	validation
	obj                  *openapi.GetSingleItemResponse
	marshaller           marshalGetSingleItemResponse
	unMarshaller         unMarshalGetSingleItemResponse
	serviceAbcItemHolder ServiceAbcItem
}

func NewGetSingleItemResponse() GetSingleItemResponse {
	obj := getSingleItemResponse{obj: &openapi.GetSingleItemResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getSingleItemResponse) msg() *openapi.GetSingleItemResponse {
	return obj.obj
}

func (obj *getSingleItemResponse) setMsg(msg *openapi.GetSingleItemResponse) GetSingleItemResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetSingleItemResponse struct {
	obj *getSingleItemResponse
}

type marshalGetSingleItemResponse interface {
	// ToProto marshals GetSingleItemResponse to protobuf object *openapi.GetSingleItemResponse
	ToProto() (*openapi.GetSingleItemResponse, error)
	// ToPbText marshals GetSingleItemResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetSingleItemResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetSingleItemResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetSingleItemResponse struct {
	obj *getSingleItemResponse
}

type unMarshalGetSingleItemResponse interface {
	// FromProto unmarshals GetSingleItemResponse from protobuf object *openapi.GetSingleItemResponse
	FromProto(msg *openapi.GetSingleItemResponse) (GetSingleItemResponse, error)
	// FromPbText unmarshals GetSingleItemResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetSingleItemResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetSingleItemResponse from JSON text
	FromJson(value string) error
}

func (obj *getSingleItemResponse) Marshal() marshalGetSingleItemResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetSingleItemResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getSingleItemResponse) Unmarshal() unMarshalGetSingleItemResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetSingleItemResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetSingleItemResponse) ToProto() (*openapi.GetSingleItemResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetSingleItemResponse) FromProto(msg *openapi.GetSingleItemResponse) (GetSingleItemResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetSingleItemResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetSingleItemResponse) FromPbText(value string) error {
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

func (m *marshalgetSingleItemResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetSingleItemResponse) FromYaml(value string) error {
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

func (m *marshalgetSingleItemResponse) ToJson() (string, error) {
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

func (m *unMarshalgetSingleItemResponse) FromJson(value string) error {
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

func (obj *getSingleItemResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getSingleItemResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getSingleItemResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getSingleItemResponse) Clone() (GetSingleItemResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetSingleItemResponse()
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

func (obj *getSingleItemResponse) setNil() {
	obj.serviceAbcItemHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetSingleItemResponse is description is TBD
type GetSingleItemResponse interface {
	Validation
	// msg marshals GetSingleItemResponse to protobuf object *openapi.GetSingleItemResponse
	// and doesn't set defaults
	msg() *openapi.GetSingleItemResponse
	// setMsg unmarshals GetSingleItemResponse from protobuf object *openapi.GetSingleItemResponse
	// and doesn't set defaults
	setMsg(*openapi.GetSingleItemResponse) GetSingleItemResponse
	// provides marshal interface
	Marshal() marshalGetSingleItemResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetSingleItemResponse
	// validate validates GetSingleItemResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetSingleItemResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItem returns ServiceAbcItem, set in GetSingleItemResponse.
	// ServiceAbcItem is description is TBD
	ServiceAbcItem() ServiceAbcItem
	// SetServiceAbcItem assigns ServiceAbcItem provided by user to GetSingleItemResponse.
	// ServiceAbcItem is description is TBD
	SetServiceAbcItem(value ServiceAbcItem) GetSingleItemResponse
	// HasServiceAbcItem checks if ServiceAbcItem has been set in GetSingleItemResponse
	HasServiceAbcItem() bool
	setNil()
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemResponse) ServiceAbcItem() ServiceAbcItem {
	if obj.obj.ServiceAbcItem == nil {
		obj.obj.ServiceAbcItem = NewServiceAbcItem().msg()
	}
	if obj.serviceAbcItemHolder == nil {
		obj.serviceAbcItemHolder = &serviceAbcItem{obj: obj.obj.ServiceAbcItem}
	}
	return obj.serviceAbcItemHolder
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemResponse) HasServiceAbcItem() bool {
	return obj.obj.ServiceAbcItem != nil
}

// description is TBD
// SetServiceAbcItem sets the ServiceAbcItem value in the GetSingleItemResponse object
func (obj *getSingleItemResponse) SetServiceAbcItem(value ServiceAbcItem) GetSingleItemResponse {

	obj.serviceAbcItemHolder = nil
	obj.obj.ServiceAbcItem = value.msg()

	return obj
}

func (obj *getSingleItemResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItem != nil {

		obj.ServiceAbcItem().validateObj(vObj, set_default)
	}

}

func (obj *getSingleItemResponse) setDefault() {

}
