package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetTestConfigResponse *****
type getTestConfigResponse struct {
	validation
	obj              *openapi.GetTestConfigResponse
	marshaller       marshalGetTestConfigResponse
	unMarshaller     unMarshalGetTestConfigResponse
	testConfigHolder TestConfig
}

func NewGetTestConfigResponse() GetTestConfigResponse {
	obj := getTestConfigResponse{obj: &openapi.GetTestConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getTestConfigResponse) msg() *openapi.GetTestConfigResponse {
	return obj.obj
}

func (obj *getTestConfigResponse) setMsg(msg *openapi.GetTestConfigResponse) GetTestConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetTestConfigResponse struct {
	obj *getTestConfigResponse
}

type marshalGetTestConfigResponse interface {
	// ToProto marshals GetTestConfigResponse to protobuf object *openapi.GetTestConfigResponse
	ToProto() (*openapi.GetTestConfigResponse, error)
	// ToPbText marshals GetTestConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetTestConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetTestConfigResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetTestConfigResponse struct {
	obj *getTestConfigResponse
}

type unMarshalGetTestConfigResponse interface {
	// FromProto unmarshals GetTestConfigResponse from protobuf object *openapi.GetTestConfigResponse
	FromProto(msg *openapi.GetTestConfigResponse) (GetTestConfigResponse, error)
	// FromPbText unmarshals GetTestConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetTestConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetTestConfigResponse from JSON text
	FromJson(value string) error
}

func (obj *getTestConfigResponse) Marshal() marshalGetTestConfigResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetTestConfigResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getTestConfigResponse) Unmarshal() unMarshalGetTestConfigResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetTestConfigResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetTestConfigResponse) ToProto() (*openapi.GetTestConfigResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetTestConfigResponse) FromProto(msg *openapi.GetTestConfigResponse) (GetTestConfigResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetTestConfigResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetTestConfigResponse) FromPbText(value string) error {
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

func (m *marshalgetTestConfigResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetTestConfigResponse) FromYaml(value string) error {
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

func (m *marshalgetTestConfigResponse) ToJson() (string, error) {
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

func (m *unMarshalgetTestConfigResponse) FromJson(value string) error {
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

func (obj *getTestConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getTestConfigResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getTestConfigResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getTestConfigResponse) Clone() (GetTestConfigResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetTestConfigResponse()
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

func (obj *getTestConfigResponse) setNil() {
	obj.testConfigHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetTestConfigResponse is description is TBD
type GetTestConfigResponse interface {
	Validation
	// msg marshals GetTestConfigResponse to protobuf object *openapi.GetTestConfigResponse
	// and doesn't set defaults
	msg() *openapi.GetTestConfigResponse
	// setMsg unmarshals GetTestConfigResponse from protobuf object *openapi.GetTestConfigResponse
	// and doesn't set defaults
	setMsg(*openapi.GetTestConfigResponse) GetTestConfigResponse
	// provides marshal interface
	Marshal() marshalGetTestConfigResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetTestConfigResponse
	// validate validates GetTestConfigResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetTestConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TestConfig returns TestConfig, set in GetTestConfigResponse.
	// TestConfig is under Review: the whole schema is being reviewed
	//
	// Description TBD
	TestConfig() TestConfig
	// SetTestConfig assigns TestConfig provided by user to GetTestConfigResponse.
	// TestConfig is under Review: the whole schema is being reviewed
	//
	// Description TBD
	SetTestConfig(value TestConfig) GetTestConfigResponse
	// HasTestConfig checks if TestConfig has been set in GetTestConfigResponse
	HasTestConfig() bool
	setNil()
}

// description is TBD
// TestConfig returns a TestConfig
func (obj *getTestConfigResponse) TestConfig() TestConfig {
	if obj.obj.TestConfig == nil {
		obj.obj.TestConfig = NewTestConfig().msg()
	}
	if obj.testConfigHolder == nil {
		obj.testConfigHolder = &testConfig{obj: obj.obj.TestConfig}
	}
	return obj.testConfigHolder
}

// description is TBD
// TestConfig returns a TestConfig
func (obj *getTestConfigResponse) HasTestConfig() bool {
	return obj.obj.TestConfig != nil
}

// description is TBD
// SetTestConfig sets the TestConfig value in the GetTestConfigResponse object
func (obj *getTestConfigResponse) SetTestConfig(value TestConfig) GetTestConfigResponse {

	obj.testConfigHolder = nil
	obj.obj.TestConfig = value.msg()

	return obj
}

func (obj *getTestConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.TestConfig != nil {

		obj.TestConfig().validateObj(vObj, set_default)
	}

}

func (obj *getTestConfigResponse) setDefault() {

}
