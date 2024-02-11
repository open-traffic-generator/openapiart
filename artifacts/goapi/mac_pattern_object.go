package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MacPatternObject *****
type macPatternObject struct {
	validation
	obj          *openapi.MacPatternObject
	marshaller   marshalMacPatternObject
	unMarshaller unMarshalMacPatternObject
	macHolder    PatternMacPatternObjectMac
}

func NewMacPatternObject() MacPatternObject {
	obj := macPatternObject{obj: &openapi.MacPatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *macPatternObject) msg() *openapi.MacPatternObject {
	return obj.obj
}

func (obj *macPatternObject) setMsg(msg *openapi.MacPatternObject) MacPatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmacPatternObject struct {
	obj *macPatternObject
}

type marshalMacPatternObject interface {
	// ToProto marshals MacPatternObject to protobuf object *openapi.MacPatternObject
	ToProto() (*openapi.MacPatternObject, error)
	// ToPbText marshals MacPatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacPatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacPatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalmacPatternObject struct {
	obj *macPatternObject
}

type unMarshalMacPatternObject interface {
	// FromProto unmarshals MacPatternObject from protobuf object *openapi.MacPatternObject
	FromProto(msg *openapi.MacPatternObject) (MacPatternObject, error)
	// FromPbText unmarshals MacPatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacPatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacPatternObject from JSON text
	FromJson(value string) error
}

func (obj *macPatternObject) Marshal() marshalMacPatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmacPatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *macPatternObject) Unmarshal() unMarshalMacPatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmacPatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmacPatternObject) ToProto() (*openapi.MacPatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmacPatternObject) FromProto(msg *openapi.MacPatternObject) (MacPatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmacPatternObject) ToPbText() (string, error) {
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

func (m *unMarshalmacPatternObject) FromPbText(value string) error {
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

func (m *marshalmacPatternObject) ToYaml() (string, error) {
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

func (m *unMarshalmacPatternObject) FromYaml(value string) error {
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

func (m *marshalmacPatternObject) ToJson() (string, error) {
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

func (m *unMarshalmacPatternObject) FromJson(value string) error {
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

func (obj *macPatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macPatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macPatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macPatternObject) Clone() (MacPatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacPatternObject()
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

func (obj *macPatternObject) setNil() {
	obj.macHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacPatternObject is test mac pattern
type MacPatternObject interface {
	Validation
	// msg marshals MacPatternObject to protobuf object *openapi.MacPatternObject
	// and doesn't set defaults
	msg() *openapi.MacPatternObject
	// setMsg unmarshals MacPatternObject from protobuf object *openapi.MacPatternObject
	// and doesn't set defaults
	setMsg(*openapi.MacPatternObject) MacPatternObject
	// provides marshal interface
	Marshal() marshalMacPatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalMacPatternObject
	// validate validates MacPatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacPatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mac returns PatternMacPatternObjectMac, set in MacPatternObject.
	// PatternMacPatternObjectMac is tBD
	Mac() PatternMacPatternObjectMac
	// SetMac assigns PatternMacPatternObjectMac provided by user to MacPatternObject.
	// PatternMacPatternObjectMac is tBD
	SetMac(value PatternMacPatternObjectMac) MacPatternObject
	// HasMac checks if Mac has been set in MacPatternObject
	HasMac() bool
	setNil()
}

// description is TBD
// Mac returns a PatternMacPatternObjectMac
func (obj *macPatternObject) Mac() PatternMacPatternObjectMac {
	if obj.obj.Mac == nil {
		obj.obj.Mac = NewPatternMacPatternObjectMac().msg()
	}
	if obj.macHolder == nil {
		obj.macHolder = &patternMacPatternObjectMac{obj: obj.obj.Mac}
	}
	return obj.macHolder
}

// description is TBD
// Mac returns a PatternMacPatternObjectMac
func (obj *macPatternObject) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the PatternMacPatternObjectMac value in the MacPatternObject object
func (obj *macPatternObject) SetMac(value PatternMacPatternObjectMac) MacPatternObject {

	obj.macHolder = nil
	obj.obj.Mac = value.msg()

	return obj
}

func (obj *macPatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		obj.Mac().validateObj(vObj, set_default)
	}

}

func (obj *macPatternObject) setDefault() {

}
