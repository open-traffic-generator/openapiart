package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IntegerPattern *****
type integerPattern struct {
	validation
	obj           *openapi.IntegerPattern
	marshaller    marshalIntegerPattern
	unMarshaller  unMarshalIntegerPattern
	integerHolder PatternIntegerPatternInteger
}

func NewIntegerPattern() IntegerPattern {
	obj := integerPattern{obj: &openapi.IntegerPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *integerPattern) msg() *openapi.IntegerPattern {
	return obj.obj
}

func (obj *integerPattern) setMsg(msg *openapi.IntegerPattern) IntegerPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalintegerPattern struct {
	obj *integerPattern
}

type marshalIntegerPattern interface {
	// ToProto marshals IntegerPattern to protobuf object *openapi.IntegerPattern
	ToProto() (*openapi.IntegerPattern, error)
	// ToPbText marshals IntegerPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IntegerPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals IntegerPattern to JSON text
	ToJson() (string, error)
}

type unMarshalintegerPattern struct {
	obj *integerPattern
}

type unMarshalIntegerPattern interface {
	// FromProto unmarshals IntegerPattern from protobuf object *openapi.IntegerPattern
	FromProto(msg *openapi.IntegerPattern) (IntegerPattern, error)
	// FromPbText unmarshals IntegerPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IntegerPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IntegerPattern from JSON text
	FromJson(value string) error
}

func (obj *integerPattern) Marshal() marshalIntegerPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalintegerPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *integerPattern) Unmarshal() unMarshalIntegerPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalintegerPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalintegerPattern) ToProto() (*openapi.IntegerPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalintegerPattern) FromProto(msg *openapi.IntegerPattern) (IntegerPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalintegerPattern) ToPbText() (string, error) {
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

func (m *unMarshalintegerPattern) FromPbText(value string) error {
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

func (m *marshalintegerPattern) ToYaml() (string, error) {
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

func (m *unMarshalintegerPattern) FromYaml(value string) error {
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

func (m *marshalintegerPattern) ToJson() (string, error) {
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

func (m *unMarshalintegerPattern) FromJson(value string) error {
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

func (obj *integerPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *integerPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *integerPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *integerPattern) Clone() (IntegerPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIntegerPattern()
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

func (obj *integerPattern) setNil() {
	obj.integerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IntegerPattern is test integer pattern
type IntegerPattern interface {
	Validation
	// msg marshals IntegerPattern to protobuf object *openapi.IntegerPattern
	// and doesn't set defaults
	msg() *openapi.IntegerPattern
	// setMsg unmarshals IntegerPattern from protobuf object *openapi.IntegerPattern
	// and doesn't set defaults
	setMsg(*openapi.IntegerPattern) IntegerPattern
	// provides marshal interface
	Marshal() marshalIntegerPattern
	// provides unmarshal interface
	Unmarshal() unMarshalIntegerPattern
	// validate validates IntegerPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IntegerPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Integer returns PatternIntegerPatternInteger, set in IntegerPattern.
	// PatternIntegerPatternInteger is tBD
	Integer() PatternIntegerPatternInteger
	// SetInteger assigns PatternIntegerPatternInteger provided by user to IntegerPattern.
	// PatternIntegerPatternInteger is tBD
	SetInteger(value PatternIntegerPatternInteger) IntegerPattern
	// HasInteger checks if Integer has been set in IntegerPattern
	HasInteger() bool
	setNil()
}

// description is TBD
// Integer returns a PatternIntegerPatternInteger
func (obj *integerPattern) Integer() PatternIntegerPatternInteger {
	if obj.obj.Integer == nil {
		obj.obj.Integer = NewPatternIntegerPatternInteger().msg()
	}
	if obj.integerHolder == nil {
		obj.integerHolder = &patternIntegerPatternInteger{obj: obj.obj.Integer}
	}
	return obj.integerHolder
}

// description is TBD
// Integer returns a PatternIntegerPatternInteger
func (obj *integerPattern) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the PatternIntegerPatternInteger value in the IntegerPattern object
func (obj *integerPattern) SetInteger(value PatternIntegerPatternInteger) IntegerPattern {

	obj.integerHolder = nil
	obj.obj.Integer = value.msg()

	return obj
}

func (obj *integerPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		obj.Integer().validateObj(vObj, set_default)
	}

}

func (obj *integerPattern) setDefault() {

}
