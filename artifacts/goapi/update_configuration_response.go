package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateConfigurationResponse *****
type updateConfigurationResponse struct {
	validation
	obj                *openapi.UpdateConfigurationResponse
	marshaller         marshalUpdateConfigurationResponse
	unMarshaller       unMarshalUpdateConfigurationResponse
	prefixConfigHolder PrefixConfig
}

func NewUpdateConfigurationResponse() UpdateConfigurationResponse {
	obj := updateConfigurationResponse{obj: &openapi.UpdateConfigurationResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *updateConfigurationResponse) msg() *openapi.UpdateConfigurationResponse {
	return obj.obj
}

func (obj *updateConfigurationResponse) setMsg(msg *openapi.UpdateConfigurationResponse) UpdateConfigurationResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateConfigurationResponse struct {
	obj *updateConfigurationResponse
}

type marshalUpdateConfigurationResponse interface {
	// ToProto marshals UpdateConfigurationResponse to protobuf object *openapi.UpdateConfigurationResponse
	ToProto() (*openapi.UpdateConfigurationResponse, error)
	// ToPbText marshals UpdateConfigurationResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateConfigurationResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateConfigurationResponse to JSON text
	ToJson() (string, error)
}

type unMarshalupdateConfigurationResponse struct {
	obj *updateConfigurationResponse
}

type unMarshalUpdateConfigurationResponse interface {
	// FromProto unmarshals UpdateConfigurationResponse from protobuf object *openapi.UpdateConfigurationResponse
	FromProto(msg *openapi.UpdateConfigurationResponse) (UpdateConfigurationResponse, error)
	// FromPbText unmarshals UpdateConfigurationResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateConfigurationResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateConfigurationResponse from JSON text
	FromJson(value string) error
}

func (obj *updateConfigurationResponse) Marshal() marshalUpdateConfigurationResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateConfigurationResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateConfigurationResponse) Unmarshal() unMarshalUpdateConfigurationResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateConfigurationResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateConfigurationResponse) ToProto() (*openapi.UpdateConfigurationResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateConfigurationResponse) FromProto(msg *openapi.UpdateConfigurationResponse) (UpdateConfigurationResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateConfigurationResponse) ToPbText() (string, error) {
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

func (m *unMarshalupdateConfigurationResponse) FromPbText(value string) error {
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

func (m *marshalupdateConfigurationResponse) ToYaml() (string, error) {
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

func (m *unMarshalupdateConfigurationResponse) FromYaml(value string) error {
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

func (m *marshalupdateConfigurationResponse) ToJson() (string, error) {
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

func (m *unMarshalupdateConfigurationResponse) FromJson(value string) error {
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

func (obj *updateConfigurationResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateConfigurationResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateConfigurationResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateConfigurationResponse) Clone() (UpdateConfigurationResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateConfigurationResponse()
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

func (obj *updateConfigurationResponse) setNil() {
	obj.prefixConfigHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateConfigurationResponse is description is TBD
type UpdateConfigurationResponse interface {
	Validation
	// msg marshals UpdateConfigurationResponse to protobuf object *openapi.UpdateConfigurationResponse
	// and doesn't set defaults
	msg() *openapi.UpdateConfigurationResponse
	// setMsg unmarshals UpdateConfigurationResponse from protobuf object *openapi.UpdateConfigurationResponse
	// and doesn't set defaults
	setMsg(*openapi.UpdateConfigurationResponse) UpdateConfigurationResponse
	// provides marshal interface
	Marshal() marshalUpdateConfigurationResponse
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateConfigurationResponse
	// validate validates UpdateConfigurationResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateConfigurationResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrefixConfig returns PrefixConfig, set in UpdateConfigurationResponse.
	// PrefixConfig is container which retains the configuration
	PrefixConfig() PrefixConfig
	// SetPrefixConfig assigns PrefixConfig provided by user to UpdateConfigurationResponse.
	// PrefixConfig is container which retains the configuration
	SetPrefixConfig(value PrefixConfig) UpdateConfigurationResponse
	// HasPrefixConfig checks if PrefixConfig has been set in UpdateConfigurationResponse
	HasPrefixConfig() bool
	setNil()
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *updateConfigurationResponse) PrefixConfig() PrefixConfig {
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
func (obj *updateConfigurationResponse) HasPrefixConfig() bool {
	return obj.obj.PrefixConfig != nil
}

// description is TBD
// SetPrefixConfig sets the PrefixConfig value in the UpdateConfigurationResponse object
func (obj *updateConfigurationResponse) SetPrefixConfig(value PrefixConfig) UpdateConfigurationResponse {

	obj.prefixConfigHolder = nil
	obj.obj.PrefixConfig = value.msg()

	return obj
}

func (obj *updateConfigurationResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixConfig != nil {

		obj.PrefixConfig().validateObj(vObj, set_default)
	}

}

func (obj *updateConfigurationResponse) setDefault() {

}
