package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetWarningsResponse *****
type getWarningsResponse struct {
	validation
	obj                  *openapi.GetWarningsResponse
	marshaller           marshalGetWarningsResponse
	unMarshaller         unMarshalGetWarningsResponse
	warningDetailsHolder WarningDetails
}

func NewGetWarningsResponse() GetWarningsResponse {
	obj := getWarningsResponse{obj: &openapi.GetWarningsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getWarningsResponse) msg() *openapi.GetWarningsResponse {
	return obj.obj
}

func (obj *getWarningsResponse) setMsg(msg *openapi.GetWarningsResponse) GetWarningsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetWarningsResponse struct {
	obj *getWarningsResponse
}

type marshalGetWarningsResponse interface {
	// ToProto marshals GetWarningsResponse to protobuf object *openapi.GetWarningsResponse
	ToProto() (*openapi.GetWarningsResponse, error)
	// ToPbText marshals GetWarningsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetWarningsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetWarningsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetWarningsResponse struct {
	obj *getWarningsResponse
}

type unMarshalGetWarningsResponse interface {
	// FromProto unmarshals GetWarningsResponse from protobuf object *openapi.GetWarningsResponse
	FromProto(msg *openapi.GetWarningsResponse) (GetWarningsResponse, error)
	// FromPbText unmarshals GetWarningsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetWarningsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetWarningsResponse from JSON text
	FromJson(value string) error
}

func (obj *getWarningsResponse) Marshal() marshalGetWarningsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetWarningsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getWarningsResponse) Unmarshal() unMarshalGetWarningsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetWarningsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetWarningsResponse) ToProto() (*openapi.GetWarningsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetWarningsResponse) FromProto(msg *openapi.GetWarningsResponse) (GetWarningsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetWarningsResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetWarningsResponse) FromPbText(value string) error {
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

func (m *marshalgetWarningsResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetWarningsResponse) FromYaml(value string) error {
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

func (m *marshalgetWarningsResponse) ToJson() (string, error) {
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

func (m *unMarshalgetWarningsResponse) FromJson(value string) error {
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

func (obj *getWarningsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getWarningsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getWarningsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getWarningsResponse) Clone() (GetWarningsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetWarningsResponse()
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

func (obj *getWarningsResponse) setNil() {
	obj.warningDetailsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetWarningsResponse is description is TBD
type GetWarningsResponse interface {
	Validation
	// msg marshals GetWarningsResponse to protobuf object *openapi.GetWarningsResponse
	// and doesn't set defaults
	msg() *openapi.GetWarningsResponse
	// setMsg unmarshals GetWarningsResponse from protobuf object *openapi.GetWarningsResponse
	// and doesn't set defaults
	setMsg(*openapi.GetWarningsResponse) GetWarningsResponse
	// provides marshal interface
	Marshal() marshalGetWarningsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetWarningsResponse
	// validate validates GetWarningsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetWarningsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// WarningDetails returns WarningDetails, set in GetWarningsResponse.
	// WarningDetails is description is TBD
	WarningDetails() WarningDetails
	// SetWarningDetails assigns WarningDetails provided by user to GetWarningsResponse.
	// WarningDetails is description is TBD
	SetWarningDetails(value WarningDetails) GetWarningsResponse
	// HasWarningDetails checks if WarningDetails has been set in GetWarningsResponse
	HasWarningDetails() bool
	setNil()
}

// description is TBD
// WarningDetails returns a WarningDetails
func (obj *getWarningsResponse) WarningDetails() WarningDetails {
	if obj.obj.WarningDetails == nil {
		obj.obj.WarningDetails = NewWarningDetails().msg()
	}
	if obj.warningDetailsHolder == nil {
		obj.warningDetailsHolder = &warningDetails{obj: obj.obj.WarningDetails}
	}
	return obj.warningDetailsHolder
}

// description is TBD
// WarningDetails returns a WarningDetails
func (obj *getWarningsResponse) HasWarningDetails() bool {
	return obj.obj.WarningDetails != nil
}

// description is TBD
// SetWarningDetails sets the WarningDetails value in the GetWarningsResponse object
func (obj *getWarningsResponse) SetWarningDetails(value WarningDetails) GetWarningsResponse {

	obj.warningDetailsHolder = nil
	obj.obj.WarningDetails = value.msg()

	return obj
}

func (obj *getWarningsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.WarningDetails != nil {

		obj.WarningDetails().validateObj(vObj, set_default)
	}

}

func (obj *getWarningsResponse) setDefault() {

}
