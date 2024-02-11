package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetConfigResponse *****
type getConfigResponse struct {
	validation
	obj                *openapi.GetConfigResponse
	marshaller         marshalGetConfigResponse
	unMarshaller       unMarshalGetConfigResponse
	prefixConfigHolder PrefixConfig
}

func NewGetConfigResponse() GetConfigResponse {
	obj := getConfigResponse{obj: &openapi.GetConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getConfigResponse) msg() *openapi.GetConfigResponse {
	return obj.obj
}

func (obj *getConfigResponse) setMsg(msg *openapi.GetConfigResponse) GetConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetConfigResponse struct {
	obj *getConfigResponse
}

type marshalGetConfigResponse interface {
	// ToProto marshals GetConfigResponse to protobuf object *openapi.GetConfigResponse
	ToProto() (*openapi.GetConfigResponse, error)
	// ToPbText marshals GetConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetConfigResponse struct {
	obj *getConfigResponse
}

type unMarshalGetConfigResponse interface {
	// FromProto unmarshals GetConfigResponse from protobuf object *openapi.GetConfigResponse
	FromProto(msg *openapi.GetConfigResponse) (GetConfigResponse, error)
	// FromPbText unmarshals GetConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *getConfigResponse) Marshal() marshalGetConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getConfigResponse) Unmarshal() unMarshalGetConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetConfigResponse) ToProto() (*openapi.GetConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetConfigResponse) FromProto(msg *openapi.GetConfigResponse) (GetConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromPbText(value string) error {
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

func (m *marshalgetConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromYaml(value string) error {
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

func (m *marshalgetConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalgetConfigResponse) FromJson(value string) error {
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

func (obj *getConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getConfigResponse) Clone() (GetConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetConfigResponse()
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

func (obj *getConfigResponse) setNil() {
	obj.prefixConfigHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetConfigResponse is description is TBD
type GetConfigResponse interface {
	Validation
	// msg marshals GetConfigResponse to protobuf object *openapi.GetConfigResponse
	// and doesn't set defaults
	msg() *openapi.GetConfigResponse
	// setMsg unmarshals GetConfigResponse from protobuf object *openapi.GetConfigResponse
	// and doesn't set defaults
	setMsg(*openapi.GetConfigResponse) GetConfigResponse
	// provides marshal interface
	Marshal() marshalGetConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetConfigResponse
	// validate validates GetConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrefixConfig returns PrefixConfig, set in GetConfigResponse.
	// PrefixConfig is container which retains the configuration
	PrefixConfig() PrefixConfig
	// SetPrefixConfig assigns PrefixConfig provided by user to GetConfigResponse.
	// PrefixConfig is container which retains the configuration
	SetPrefixConfig(value PrefixConfig) GetConfigResponse
	// HasPrefixConfig checks if PrefixConfig has been set in GetConfigResponse
	HasPrefixConfig() bool
	setNil()
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *getConfigResponse) PrefixConfig() PrefixConfig {
	if obj.obj.PrefixConfig == nil {
		obj.obj.PrefixConfig = NewPrefixConfig().msg()
	}
	if obj.prefixConfigHolder == nil {
		obj.prefixConfigHolder = &prefixConfig{obj: obj.obj.PrefixConfig}
	}
	return obj.prefixConfigHolder
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *getConfigResponse) HasPrefixConfig() bool {
	return obj.obj.PrefixConfig != nil
}

// description is TBD
// SetPrefixConfig sets the PrefixConfig value in the GetConfigResponse object
func (obj *getConfigResponse) SetPrefixConfig(value PrefixConfig) GetConfigResponse {

	obj.prefixConfigHolder = nil
	obj.obj.PrefixConfig = value.msg()

	return obj
}

func (obj *getConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixConfig != nil {

		obj.PrefixConfig().validateObj(vObj, set_default)
	}

}

func (obj *getConfigResponse) setDefault() {

}
