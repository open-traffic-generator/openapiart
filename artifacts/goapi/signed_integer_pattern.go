package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** SignedIntegerPattern *****
type signedIntegerPattern struct {
	validation
	obj           *openapi.SignedIntegerPattern
	marshaller    marshalSignedIntegerPattern
	unMarshaller  unMarshalSignedIntegerPattern
	integerHolder PatternSignedIntegerPatternInteger
}

func NewSignedIntegerPattern() SignedIntegerPattern {
	obj := signedIntegerPattern{obj: &openapi.SignedIntegerPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *signedIntegerPattern) msg() *openapi.SignedIntegerPattern {
	return obj.obj
}

func (obj *signedIntegerPattern) setMsg(msg *openapi.SignedIntegerPattern) SignedIntegerPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalsignedIntegerPattern struct {
	obj *signedIntegerPattern
}

type marshalSignedIntegerPattern interface {
	// ToProto marshals SignedIntegerPattern to protobuf object *openapi.SignedIntegerPattern
	ToProto() (*openapi.SignedIntegerPattern, error)
	// ToPbText marshals SignedIntegerPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SignedIntegerPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals SignedIntegerPattern to JSON text
	ToJson() (string, error)
}

type unMarshalsignedIntegerPattern struct {
	obj *signedIntegerPattern
}

type unMarshalSignedIntegerPattern interface {
	// FromProto unmarshals SignedIntegerPattern from protobuf object *openapi.SignedIntegerPattern
	FromProto(msg *openapi.SignedIntegerPattern) (SignedIntegerPattern, error)
	// FromPbText unmarshals SignedIntegerPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SignedIntegerPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SignedIntegerPattern from JSON text
	FromJson(value string) error
}

func (obj *signedIntegerPattern) Marshal() marshalSignedIntegerPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalsignedIntegerPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *signedIntegerPattern) Unmarshal() unMarshalSignedIntegerPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalsignedIntegerPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalsignedIntegerPattern) ToProto() (*openapi.SignedIntegerPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalsignedIntegerPattern) FromProto(msg *openapi.SignedIntegerPattern) (SignedIntegerPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalsignedIntegerPattern) ToPbText() (string, error) {
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

func (m *unMarshalsignedIntegerPattern) FromPbText(value string) error {
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

func (m *marshalsignedIntegerPattern) ToYaml() (string, error) {
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

func (m *unMarshalsignedIntegerPattern) FromYaml(value string) error {
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

func (m *marshalsignedIntegerPattern) ToJson() (string, error) {
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

func (m *unMarshalsignedIntegerPattern) FromJson(value string) error {
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

func (obj *signedIntegerPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *signedIntegerPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *signedIntegerPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *signedIntegerPattern) Clone() (SignedIntegerPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSignedIntegerPattern()
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

func (obj *signedIntegerPattern) setNil() {
	obj.integerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// SignedIntegerPattern is test signed integer pattern
type SignedIntegerPattern interface {
	Validation
	// msg marshals SignedIntegerPattern to protobuf object *openapi.SignedIntegerPattern
	// and doesn't set defaults
	msg() *openapi.SignedIntegerPattern
	// setMsg unmarshals SignedIntegerPattern from protobuf object *openapi.SignedIntegerPattern
	// and doesn't set defaults
	setMsg(*openapi.SignedIntegerPattern) SignedIntegerPattern
	// provides marshal interface
	Marshal() marshalSignedIntegerPattern
	// provides unmarshal interface
	Unmarshal() unMarshalSignedIntegerPattern
	// validate validates SignedIntegerPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SignedIntegerPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Integer returns PatternSignedIntegerPatternInteger, set in SignedIntegerPattern.
	// PatternSignedIntegerPatternInteger is tBD
	Integer() PatternSignedIntegerPatternInteger
	// SetInteger assigns PatternSignedIntegerPatternInteger provided by user to SignedIntegerPattern.
	// PatternSignedIntegerPatternInteger is tBD
	SetInteger(value PatternSignedIntegerPatternInteger) SignedIntegerPattern
	// HasInteger checks if Integer has been set in SignedIntegerPattern
	HasInteger() bool
	setNil()
}

// description is TBD
// Integer returns a PatternSignedIntegerPatternInteger
func (obj *signedIntegerPattern) Integer() PatternSignedIntegerPatternInteger {
	if obj.obj.Integer == nil {
		obj.obj.Integer = NewPatternSignedIntegerPatternInteger().msg()
	}
	if obj.integerHolder == nil {
		obj.integerHolder = &patternSignedIntegerPatternInteger{obj: obj.obj.Integer}
	}
	return obj.integerHolder
}

// description is TBD
// Integer returns a PatternSignedIntegerPatternInteger
func (obj *signedIntegerPattern) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the PatternSignedIntegerPatternInteger value in the SignedIntegerPattern object
func (obj *signedIntegerPattern) SetInteger(value PatternSignedIntegerPatternInteger) SignedIntegerPattern {

	obj.integerHolder = nil
	obj.obj.Integer = value.msg()

	return obj
}

func (obj *signedIntegerPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		obj.Integer().validateObj(vObj, set_default)
	}

}

func (obj *signedIntegerPattern) setDefault() {

}
