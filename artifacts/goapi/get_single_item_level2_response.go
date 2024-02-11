package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetSingleItemLevel2Response *****
type getSingleItemLevel2Response struct {
	validation
	obj                  *openapi.GetSingleItemLevel2Response
	marshaller           marshalGetSingleItemLevel2Response
	unMarshaller         unMarshalGetSingleItemLevel2Response
	serviceAbcItemHolder ServiceAbcItem
}

func NewGetSingleItemLevel2Response() GetSingleItemLevel2Response {
	obj := getSingleItemLevel2Response{obj: &openapi.GetSingleItemLevel2Response{}}
	obj.setDefault()
	return &obj
}

func (obj *getSingleItemLevel2Response) msg() *openapi.GetSingleItemLevel2Response {
	return obj.obj
}

func (obj *getSingleItemLevel2Response) setMsg(msg *openapi.GetSingleItemLevel2Response) GetSingleItemLevel2Response {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetSingleItemLevel2Response struct {
	obj *getSingleItemLevel2Response
}

type marshalGetSingleItemLevel2Response interface {
	// ToProto marshals GetSingleItemLevel2Response to protobuf object *openapi.GetSingleItemLevel2Response
	ToProto() (*openapi.GetSingleItemLevel2Response, error)
	// ToPbText marshals GetSingleItemLevel2Response to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetSingleItemLevel2Response to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetSingleItemLevel2Response to JSON text
	ToJson() (string, error)
}

type unMarshalgetSingleItemLevel2Response struct {
	obj *getSingleItemLevel2Response
}

type unMarshalGetSingleItemLevel2Response interface {
	// FromProto unmarshals GetSingleItemLevel2Response from protobuf object *openapi.GetSingleItemLevel2Response
	FromProto(msg *openapi.GetSingleItemLevel2Response) (GetSingleItemLevel2Response, error)
	// FromPbText unmarshals GetSingleItemLevel2Response from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetSingleItemLevel2Response from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetSingleItemLevel2Response from JSON text
	FromJson(value string) error
}

func (obj *getSingleItemLevel2Response) Marshal() marshalGetSingleItemLevel2Response {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetSingleItemLevel2Response{obj: obj}
	}
	return obj.marshaller
}

func (obj *getSingleItemLevel2Response) Unmarshal() unMarshalGetSingleItemLevel2Response {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetSingleItemLevel2Response{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetSingleItemLevel2Response) ToProto() (*openapi.GetSingleItemLevel2Response, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetSingleItemLevel2Response) FromProto(msg *openapi.GetSingleItemLevel2Response) (GetSingleItemLevel2Response, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetSingleItemLevel2Response) ToPbText() (string, error) {
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

func (m *unMarshalgetSingleItemLevel2Response) FromPbText(value string) error {
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

func (m *marshalgetSingleItemLevel2Response) ToYaml() (string, error) {
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

func (m *unMarshalgetSingleItemLevel2Response) FromYaml(value string) error {
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

func (m *marshalgetSingleItemLevel2Response) ToJson() (string, error) {
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

func (m *unMarshalgetSingleItemLevel2Response) FromJson(value string) error {
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

func (obj *getSingleItemLevel2Response) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getSingleItemLevel2Response) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getSingleItemLevel2Response) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getSingleItemLevel2Response) Clone() (GetSingleItemLevel2Response, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetSingleItemLevel2Response()
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

func (obj *getSingleItemLevel2Response) setNil() {
	obj.serviceAbcItemHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetSingleItemLevel2Response is description is TBD
type GetSingleItemLevel2Response interface {
	Validation
	// msg marshals GetSingleItemLevel2Response to protobuf object *openapi.GetSingleItemLevel2Response
	// and doesn't set defaults
	msg() *openapi.GetSingleItemLevel2Response
	// setMsg unmarshals GetSingleItemLevel2Response from protobuf object *openapi.GetSingleItemLevel2Response
	// and doesn't set defaults
	setMsg(*openapi.GetSingleItemLevel2Response) GetSingleItemLevel2Response
	// provides marshal interface
	Marshal() marshalGetSingleItemLevel2Response
	// provides unmarshal interface
	Unmarshal() unMarshalGetSingleItemLevel2Response
	// validate validates GetSingleItemLevel2Response
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetSingleItemLevel2Response, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItem returns ServiceAbcItem, set in GetSingleItemLevel2Response.
	// ServiceAbcItem is description is TBD
	ServiceAbcItem() ServiceAbcItem
	// SetServiceAbcItem assigns ServiceAbcItem provided by user to GetSingleItemLevel2Response.
	// ServiceAbcItem is description is TBD
	SetServiceAbcItem(value ServiceAbcItem) GetSingleItemLevel2Response
	// HasServiceAbcItem checks if ServiceAbcItem has been set in GetSingleItemLevel2Response
	HasServiceAbcItem() bool
	setNil()
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemLevel2Response) ServiceAbcItem() ServiceAbcItem {
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
func (obj *getSingleItemLevel2Response) HasServiceAbcItem() bool {
	return obj.obj.ServiceAbcItem != nil
}

// description is TBD
// SetServiceAbcItem sets the ServiceAbcItem value in the GetSingleItemLevel2Response object
func (obj *getSingleItemLevel2Response) SetServiceAbcItem(value ServiceAbcItem) GetSingleItemLevel2Response {

	obj.serviceAbcItemHolder = nil
	obj.obj.ServiceAbcItem = value.msg()

	return obj
}

func (obj *getSingleItemLevel2Response) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItem != nil {

		obj.ServiceAbcItem().validateObj(vObj, set_default)
	}

}

func (obj *getSingleItemLevel2Response) setDefault() {

}
