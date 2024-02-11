package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacPattern *****
type macPattern struct {
	validation
	obj          *openapi.MacPattern
	marshaller   marshalMacPattern
	unMarshaller unMarshalMacPattern
	macHolder    PatternMacPatternMac
}

func NewMacPattern() MacPattern {
	obj := macPattern{obj: &openapi.MacPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *macPattern) msg() *openapi.MacPattern {
	return obj.obj
}

func (obj *macPattern) setMsg(msg *openapi.MacPattern) MacPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacPattern struct {
	obj *macPattern
}

type marshalMacPattern interface {
	// ToProto marshals MacPattern to protobuf object *openapi.MacPattern
	ToProto() (*openapi.MacPattern, error)
	// ToPbText marshals MacPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacPattern to JSON text
	ToJson() (string, error)
}

type unMarshalmacPattern struct {
	obj *macPattern
}

type unMarshalMacPattern interface {
	// FromProto unmarshals MacPattern from protobuf object *openapi.MacPattern
	FromProto(msg *openapi.MacPattern) (MacPattern, error)
	// FromPbText unmarshals MacPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacPattern from JSON text
	FromJson(value string) error
}

func (obj *macPattern) Marshal() marshalMacPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *macPattern) Unmarshal() unMarshalMacPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacPattern) ToProto() (*openapi.MacPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacPattern) FromProto(msg *openapi.MacPattern) (MacPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacPattern) ToPbText() (string, error) {
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

func (m *unMarshalmacPattern) FromPbText(value string) error {
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

func (m *marshalmacPattern) ToYaml() (string, error) {
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

func (m *unMarshalmacPattern) FromYaml(value string) error {
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

func (m *marshalmacPattern) ToJson() (string, error) {
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

func (m *unMarshalmacPattern) FromJson(value string) error {
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

func (obj *macPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macPattern) Clone() (MacPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacPattern()
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

func (obj *macPattern) setNil() {
	obj.macHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacPattern is test mac pattern
type MacPattern interface {
	Validation
	// msg marshals MacPattern to protobuf object *openapi.MacPattern
	// and doesn't set defaults
	msg() *openapi.MacPattern
	// setMsg unmarshals MacPattern from protobuf object *openapi.MacPattern
	// and doesn't set defaults
	setMsg(*openapi.MacPattern) MacPattern
	// provides marshal interface
	Marshal() marshalMacPattern
	// provides unmarshal interface
	Unmarshal() unMarshalMacPattern
	// validate validates MacPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mac returns PatternMacPatternMac, set in MacPattern.
	// PatternMacPatternMac is tBD
	Mac() PatternMacPatternMac
	// SetMac assigns PatternMacPatternMac provided by user to MacPattern.
	// PatternMacPatternMac is tBD
	SetMac(value PatternMacPatternMac) MacPattern
	// HasMac checks if Mac has been set in MacPattern
	HasMac() bool
	setNil()
}

// description is TBD
// Mac returns a PatternMacPatternMac
func (obj *macPattern) Mac() PatternMacPatternMac {
	if obj.obj.Mac == nil {
		obj.obj.Mac = NewPatternMacPatternMac().msg()
	}
	if obj.macHolder == nil {
		obj.macHolder = &patternMacPatternMac{obj: obj.obj.Mac}
	}
	return obj.macHolder
}

// description is TBD
// Mac returns a PatternMacPatternMac
func (obj *macPattern) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the PatternMacPatternMac value in the MacPattern object
func (obj *macPattern) SetMac(value PatternMacPatternMac) MacPattern {

	obj.macHolder = nil
	obj.obj.Mac = value.msg()

	return obj
}

func (obj *macPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		obj.Mac().validateObj(vObj, set_default)
	}

}

func (obj *macPattern) setDefault() {

}
