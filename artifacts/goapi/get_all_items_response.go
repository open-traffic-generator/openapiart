package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetAllItemsResponse *****
type getAllItemsResponse struct {
	validation
	obj                      *openapi.GetAllItemsResponse
	marshaller               marshalGetAllItemsResponse
	unMarshaller             unMarshalGetAllItemsResponse
	serviceAbcItemListHolder ServiceAbcItemList
}

func NewGetAllItemsResponse() GetAllItemsResponse {
	obj := getAllItemsResponse{obj: &openapi.GetAllItemsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getAllItemsResponse) msg() *openapi.GetAllItemsResponse {
	return obj.obj
}

func (obj *getAllItemsResponse) setMsg(msg *openapi.GetAllItemsResponse) GetAllItemsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetAllItemsResponse struct {
	obj *getAllItemsResponse
}

type marshalGetAllItemsResponse interface {
	// ToProto marshals GetAllItemsResponse to protobuf object *openapi.GetAllItemsResponse
	ToProto() (*openapi.GetAllItemsResponse, error)
	// ToPbText marshals GetAllItemsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetAllItemsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetAllItemsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetAllItemsResponse struct {
	obj *getAllItemsResponse
}

type unMarshalGetAllItemsResponse interface {
	// FromProto unmarshals GetAllItemsResponse from protobuf object *openapi.GetAllItemsResponse
	FromProto(msg *openapi.GetAllItemsResponse) (GetAllItemsResponse, error)
	// FromPbText unmarshals GetAllItemsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetAllItemsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetAllItemsResponse from JSON text
	FromJson(value string) error
}

func (obj *getAllItemsResponse) Marshal() marshalGetAllItemsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetAllItemsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getAllItemsResponse) Unmarshal() unMarshalGetAllItemsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetAllItemsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetAllItemsResponse) ToProto() (*openapi.GetAllItemsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetAllItemsResponse) FromProto(msg *openapi.GetAllItemsResponse) (GetAllItemsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetAllItemsResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetAllItemsResponse) FromPbText(value string) error {
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

func (m *marshalgetAllItemsResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetAllItemsResponse) FromYaml(value string) error {
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

func (m *marshalgetAllItemsResponse) ToJson() (string, error) {
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

func (m *unMarshalgetAllItemsResponse) FromJson(value string) error {
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

func (obj *getAllItemsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getAllItemsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getAllItemsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getAllItemsResponse) Clone() (GetAllItemsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetAllItemsResponse()
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

func (obj *getAllItemsResponse) setNil() {
	obj.serviceAbcItemListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetAllItemsResponse is description is TBD
type GetAllItemsResponse interface {
	Validation
	// msg marshals GetAllItemsResponse to protobuf object *openapi.GetAllItemsResponse
	// and doesn't set defaults
	msg() *openapi.GetAllItemsResponse
	// setMsg unmarshals GetAllItemsResponse from protobuf object *openapi.GetAllItemsResponse
	// and doesn't set defaults
	setMsg(*openapi.GetAllItemsResponse) GetAllItemsResponse
	// provides marshal interface
	Marshal() marshalGetAllItemsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetAllItemsResponse
	// validate validates GetAllItemsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetAllItemsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItemList returns ServiceAbcItemList, set in GetAllItemsResponse.
	// ServiceAbcItemList is description is TBD
	ServiceAbcItemList() ServiceAbcItemList
	// SetServiceAbcItemList assigns ServiceAbcItemList provided by user to GetAllItemsResponse.
	// ServiceAbcItemList is description is TBD
	SetServiceAbcItemList(value ServiceAbcItemList) GetAllItemsResponse
	// HasServiceAbcItemList checks if ServiceAbcItemList has been set in GetAllItemsResponse
	HasServiceAbcItemList() bool
	setNil()
}

// description is TBD
// ServiceAbcItemList returns a ServiceAbcItemList
func (obj *getAllItemsResponse) ServiceAbcItemList() ServiceAbcItemList {
	if obj.obj.ServiceAbcItemList == nil {
		obj.obj.ServiceAbcItemList = NewServiceAbcItemList().msg()
	}
	if obj.serviceAbcItemListHolder == nil {
		obj.serviceAbcItemListHolder = &serviceAbcItemList{obj: obj.obj.ServiceAbcItemList}
	}
	return obj.serviceAbcItemListHolder
}

// description is TBD
// ServiceAbcItemList returns a ServiceAbcItemList
func (obj *getAllItemsResponse) HasServiceAbcItemList() bool {
	return obj.obj.ServiceAbcItemList != nil
}

// description is TBD
// SetServiceAbcItemList sets the ServiceAbcItemList value in the GetAllItemsResponse object
func (obj *getAllItemsResponse) SetServiceAbcItemList(value ServiceAbcItemList) GetAllItemsResponse {

	obj.serviceAbcItemListHolder = nil
	obj.obj.ServiceAbcItemList = value.msg()

	return obj
}

func (obj *getAllItemsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItemList != nil {

		obj.ServiceAbcItemList().validateObj(vObj, set_default)
	}

}

func (obj *getAllItemsResponse) setDefault() {

}
