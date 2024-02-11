package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** TestConfig *****
type testConfig struct {
	validation
	obj                    *openapi.TestConfig
	marshaller             marshalTestConfig
	unMarshaller           unMarshalTestConfig
	nativeFeaturesHolder   NativeFeatures
	extendedFeaturesHolder ExtendedFeatures
}

func NewTestConfig() TestConfig {
	obj := testConfig{obj: &openapi.TestConfig{}}
	obj.setDefault()
	return &obj
}

func (obj *testConfig) msg() *openapi.TestConfig {
	return obj.obj
}

func (obj *testConfig) setMsg(msg *openapi.TestConfig) TestConfig {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaltestConfig struct {
	obj *testConfig
}

type marshalTestConfig interface {
	// ToProto marshals TestConfig to protobuf object *openapi.TestConfig
	ToProto() (*openapi.TestConfig, error)
	// ToPbText marshals TestConfig to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals TestConfig to YAML text
	ToYaml() (string, error)
	// ToJson marshals TestConfig to JSON text
	ToJson() (string, error)
}

type unMarshaltestConfig struct {
	obj *testConfig
}

type unMarshalTestConfig interface {
	// FromProto unmarshals TestConfig from protobuf object *openapi.TestConfig
	FromProto(msg *openapi.TestConfig) (TestConfig, error)
	// FromPbText unmarshals TestConfig from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals TestConfig from YAML text
	FromYaml(value string) error
	// FromJson unmarshals TestConfig from JSON text
	FromJson(value string) error
}

func (obj *testConfig) Marshal() marshalTestConfig {
	if obj.marshaller == nil {
		obj.marshaller = &marshaltestConfig{obj: obj}
	}
	return obj.marshaller
}

func (obj *testConfig) Unmarshal() unMarshalTestConfig {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaltestConfig{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaltestConfig) ToProto() (*openapi.TestConfig, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaltestConfig) FromProto(msg *openapi.TestConfig) (TestConfig, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaltestConfig) ToPbText() (string, error) {
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

func (m *unMarshaltestConfig) FromPbText(value string) error {
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

func (m *marshaltestConfig) ToYaml() (string, error) {
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

func (m *unMarshaltestConfig) FromYaml(value string) error {
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

func (m *marshaltestConfig) ToJson() (string, error) {
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

func (m *unMarshaltestConfig) FromJson(value string) error {
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

func (obj *testConfig) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *testConfig) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *testConfig) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *testConfig) Clone() (TestConfig, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewTestConfig()
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

func (obj *testConfig) setNil() {
	obj.nativeFeaturesHolder = nil
	obj.extendedFeaturesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// TestConfig is under Review: the whole schema is being reviewed
//
// Description TBD
type TestConfig interface {
	Validation
	// msg marshals TestConfig to protobuf object *openapi.TestConfig
	// and doesn't set defaults
	msg() *openapi.TestConfig
	// setMsg unmarshals TestConfig from protobuf object *openapi.TestConfig
	// and doesn't set defaults
	setMsg(*openapi.TestConfig) TestConfig
	// provides marshal interface
	Marshal() marshalTestConfig
	// provides unmarshal interface
	Unmarshal() unMarshalTestConfig
	// validate validates TestConfig
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (TestConfig, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// NativeFeatures returns NativeFeatures, set in TestConfig.
	// NativeFeatures is description is TBD
	NativeFeatures() NativeFeatures
	// SetNativeFeatures assigns NativeFeatures provided by user to TestConfig.
	// NativeFeatures is description is TBD
	SetNativeFeatures(value NativeFeatures) TestConfig
	// HasNativeFeatures checks if NativeFeatures has been set in TestConfig
	HasNativeFeatures() bool
	// ExtendedFeatures returns ExtendedFeatures, set in TestConfig.
	// ExtendedFeatures is description is TBD
	ExtendedFeatures() ExtendedFeatures
	// SetExtendedFeatures assigns ExtendedFeatures provided by user to TestConfig.
	// ExtendedFeatures is description is TBD
	SetExtendedFeatures(value ExtendedFeatures) TestConfig
	// HasExtendedFeatures checks if ExtendedFeatures has been set in TestConfig
	HasExtendedFeatures() bool
	setNil()
}

// description is TBD
// NativeFeatures returns a NativeFeatures
func (obj *testConfig) NativeFeatures() NativeFeatures {
	if obj.obj.NativeFeatures == nil {
		obj.obj.NativeFeatures = NewNativeFeatures().msg()
	}
	if obj.nativeFeaturesHolder == nil {
		obj.nativeFeaturesHolder = &nativeFeatures{obj: obj.obj.NativeFeatures}
	}
	return obj.nativeFeaturesHolder
}

// description is TBD
// NativeFeatures returns a NativeFeatures
func (obj *testConfig) HasNativeFeatures() bool {
	return obj.obj.NativeFeatures != nil
}

// description is TBD
// SetNativeFeatures sets the NativeFeatures value in the TestConfig object
func (obj *testConfig) SetNativeFeatures(value NativeFeatures) TestConfig {

	obj.nativeFeaturesHolder = nil
	obj.obj.NativeFeatures = value.msg()

	return obj
}

// description is TBD
// ExtendedFeatures returns a ExtendedFeatures
func (obj *testConfig) ExtendedFeatures() ExtendedFeatures {
	if obj.obj.ExtendedFeatures == nil {
		obj.obj.ExtendedFeatures = NewExtendedFeatures().msg()
	}
	if obj.extendedFeaturesHolder == nil {
		obj.extendedFeaturesHolder = &extendedFeatures{obj: obj.obj.ExtendedFeatures}
	}
	return obj.extendedFeaturesHolder
}

// description is TBD
// ExtendedFeatures returns a ExtendedFeatures
func (obj *testConfig) HasExtendedFeatures() bool {
	return obj.obj.ExtendedFeatures != nil
}

// description is TBD
// SetExtendedFeatures sets the ExtendedFeatures value in the TestConfig object
func (obj *testConfig) SetExtendedFeatures(value ExtendedFeatures) TestConfig {

	obj.extendedFeaturesHolder = nil
	obj.obj.ExtendedFeatures = value.msg()

	return obj
}

func (obj *testConfig) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("TestConfig is under review, the whole schema is being reviewed")

	if obj.obj.NativeFeatures != nil {

		obj.NativeFeatures().validateObj(vObj, set_default)
	}

	if obj.obj.ExtendedFeatures != nil {

		obj.ExtendedFeatures().validateObj(vObj, set_default)
	}

}

func (obj *testConfig) setDefault() {

}
