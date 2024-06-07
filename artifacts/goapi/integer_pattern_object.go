package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IntegerPatternObject *****
type integerPatternObject struct {
	validation
	obj           *openapi.IntegerPatternObject
	marshaller    marshalIntegerPatternObject
	unMarshaller  unMarshalIntegerPatternObject
	integerHolder PatternIntegerPatternObjectInteger
}

func NewIntegerPatternObject() IntegerPatternObject {
	obj := integerPatternObject{obj: &openapi.IntegerPatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *integerPatternObject) msg() *openapi.IntegerPatternObject {
	return obj.obj
}

func (obj *integerPatternObject) setMsg(msg *openapi.IntegerPatternObject) IntegerPatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalintegerPatternObject struct {
	obj *integerPatternObject
}

type marshalIntegerPatternObject interface {
	// ToProto marshals IntegerPatternObject to protobuf object *openapi.IntegerPatternObject
	ToProto() (*openapi.IntegerPatternObject, error)
	// ToPbText marshals IntegerPatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IntegerPatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals IntegerPatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalintegerPatternObject struct {
	obj *integerPatternObject
}

type unMarshalIntegerPatternObject interface {
	// FromProto unmarshals IntegerPatternObject from protobuf object *openapi.IntegerPatternObject
	FromProto(msg *openapi.IntegerPatternObject) (IntegerPatternObject, error)
	// FromPbText unmarshals IntegerPatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IntegerPatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IntegerPatternObject from JSON text
	FromJson(value string) error
}

func (obj *integerPatternObject) Marshal() marshalIntegerPatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalintegerPatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *integerPatternObject) Unmarshal() unMarshalIntegerPatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalintegerPatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalintegerPatternObject) ToProto() (*openapi.IntegerPatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalintegerPatternObject) FromProto(msg *openapi.IntegerPatternObject) (IntegerPatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalintegerPatternObject) ToPbText() (string, error) {
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

func (m *unMarshalintegerPatternObject) FromPbText(value string) error {
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

func (m *marshalintegerPatternObject) ToYaml() (string, error) {
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

func (m *unMarshalintegerPatternObject) FromYaml(value string) error {
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

func (m *marshalintegerPatternObject) ToJson() (string, error) {
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

func (m *unMarshalintegerPatternObject) FromJson(value string) error {
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

func (obj *integerPatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *integerPatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *integerPatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *integerPatternObject) Clone() (IntegerPatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIntegerPatternObject()
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

func (obj *integerPatternObject) setNil() {
	obj.integerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IntegerPatternObject is test integer pattern
type IntegerPatternObject interface {
	Validation
	// msg marshals IntegerPatternObject to protobuf object *openapi.IntegerPatternObject
	// and doesn't set defaults
	msg() *openapi.IntegerPatternObject
	// setMsg unmarshals IntegerPatternObject from protobuf object *openapi.IntegerPatternObject
	// and doesn't set defaults
	setMsg(*openapi.IntegerPatternObject) IntegerPatternObject
	// provides marshal interface
	Marshal() marshalIntegerPatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalIntegerPatternObject
	// validate validates IntegerPatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IntegerPatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Integer returns PatternIntegerPatternObjectInteger, set in IntegerPatternObject.
	// PatternIntegerPatternObjectInteger is tBD
	Integer() PatternIntegerPatternObjectInteger
	// SetInteger assigns PatternIntegerPatternObjectInteger provided by user to IntegerPatternObject.
	// PatternIntegerPatternObjectInteger is tBD
	SetInteger(value PatternIntegerPatternObjectInteger) IntegerPatternObject
	// HasInteger checks if Integer has been set in IntegerPatternObject
	HasInteger() bool
	setNil()
}

// description is TBD
// Integer returns a PatternIntegerPatternObjectInteger
func (obj *integerPatternObject) Integer() PatternIntegerPatternObjectInteger {
	if obj.obj.Integer == nil {
		obj.obj.Integer = NewPatternIntegerPatternObjectInteger().msg()
	}
	if obj.integerHolder == nil {
		obj.integerHolder = &patternIntegerPatternObjectInteger{obj: obj.obj.Integer}
	}
	return obj.integerHolder
}

// description is TBD
// Integer returns a PatternIntegerPatternObjectInteger
func (obj *integerPatternObject) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the PatternIntegerPatternObjectInteger value in the IntegerPatternObject object
func (obj *integerPatternObject) SetInteger(value PatternIntegerPatternObjectInteger) IntegerPatternObject {

	obj.integerHolder = nil
	obj.obj.Integer = value.msg()

	return obj
}

func (obj *integerPatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		obj.Integer().validateObj(vObj, set_default)
	}

}

func (obj *integerPatternObject) setDefault() {

}
