package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** AutoPatternDefault *****
type autoPatternDefault struct {
	validation
	obj                 *openapi.AutoPatternDefault
	marshaller          marshalAutoPatternDefault
	unMarshaller        unMarshalAutoPatternDefault
	autoIpDefaultHolder PatternAutoPatternDefaultAutoIpDefault
}

func NewAutoPatternDefault() AutoPatternDefault {
	obj := autoPatternDefault{obj: &openapi.AutoPatternDefault{}}
	obj.setDefault()
	return &obj
}

func (obj *autoPatternDefault) msg() *openapi.AutoPatternDefault {
	return obj.obj
}

func (obj *autoPatternDefault) setMsg(msg *openapi.AutoPatternDefault) AutoPatternDefault {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalautoPatternDefault struct {
	obj *autoPatternDefault
}

type marshalAutoPatternDefault interface {
	// ToProto marshals AutoPatternDefault to protobuf object *openapi.AutoPatternDefault
	ToProto() (*openapi.AutoPatternDefault, error)
	// ToPbText marshals AutoPatternDefault to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals AutoPatternDefault to YAML text
	ToYaml() (string, error)
	// ToJson marshals AutoPatternDefault to JSON text
	ToJson() (string, error)
}

type unMarshalautoPatternDefault struct {
	obj *autoPatternDefault
}

type unMarshalAutoPatternDefault interface {
	// FromProto unmarshals AutoPatternDefault from protobuf object *openapi.AutoPatternDefault
	FromProto(msg *openapi.AutoPatternDefault) (AutoPatternDefault, error)
	// FromPbText unmarshals AutoPatternDefault from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals AutoPatternDefault from YAML text
	FromYaml(value string) error
	// FromJson unmarshals AutoPatternDefault from JSON text
	FromJson(value string) error
}

func (obj *autoPatternDefault) Marshal() marshalAutoPatternDefault {
	if obj.marshaller == nil {
		obj.marshaller = &marshalautoPatternDefault{obj: obj}
	}
	return obj.marshaller
}

func (obj *autoPatternDefault) Unmarshal() unMarshalAutoPatternDefault {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalautoPatternDefault{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalautoPatternDefault) ToProto() (*openapi.AutoPatternDefault, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalautoPatternDefault) FromProto(msg *openapi.AutoPatternDefault) (AutoPatternDefault, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalautoPatternDefault) ToPbText() (string, error) {
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

func (m *unMarshalautoPatternDefault) FromPbText(value string) error {
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

func (m *marshalautoPatternDefault) ToYaml() (string, error) {
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

func (m *unMarshalautoPatternDefault) FromYaml(value string) error {
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

func (m *marshalautoPatternDefault) ToJson() (string, error) {
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

func (m *unMarshalautoPatternDefault) FromJson(value string) error {
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

func (obj *autoPatternDefault) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *autoPatternDefault) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *autoPatternDefault) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *autoPatternDefault) Clone() (AutoPatternDefault, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewAutoPatternDefault()
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

func (obj *autoPatternDefault) setNil() {
	obj.autoIpDefaultHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// AutoPatternDefault is test auto pattern with default
type AutoPatternDefault interface {
	Validation
	// msg marshals AutoPatternDefault to protobuf object *openapi.AutoPatternDefault
	// and doesn't set defaults
	msg() *openapi.AutoPatternDefault
	// setMsg unmarshals AutoPatternDefault from protobuf object *openapi.AutoPatternDefault
	// and doesn't set defaults
	setMsg(*openapi.AutoPatternDefault) AutoPatternDefault
	// provides marshal interface
	Marshal() marshalAutoPatternDefault
	// provides unmarshal interface
	Unmarshal() unMarshalAutoPatternDefault
	// validate validates AutoPatternDefault
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (AutoPatternDefault, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AutoIpDefault returns PatternAutoPatternDefaultAutoIpDefault, set in AutoPatternDefault.
	// PatternAutoPatternDefaultAutoIpDefault is tBD
	AutoIpDefault() PatternAutoPatternDefaultAutoIpDefault
	// SetAutoIpDefault assigns PatternAutoPatternDefaultAutoIpDefault provided by user to AutoPatternDefault.
	// PatternAutoPatternDefaultAutoIpDefault is tBD
	SetAutoIpDefault(value PatternAutoPatternDefaultAutoIpDefault) AutoPatternDefault
	// HasAutoIpDefault checks if AutoIpDefault has been set in AutoPatternDefault
	HasAutoIpDefault() bool
	setNil()
}

// description is TBD
// AutoIpDefault returns a PatternAutoPatternDefaultAutoIpDefault
func (obj *autoPatternDefault) AutoIpDefault() PatternAutoPatternDefaultAutoIpDefault {
	if obj.obj.AutoIpDefault == nil {
		obj.obj.AutoIpDefault = NewPatternAutoPatternDefaultAutoIpDefault().msg()
	}
	if obj.autoIpDefaultHolder == nil {
		obj.autoIpDefaultHolder = &patternAutoPatternDefaultAutoIpDefault{obj: obj.obj.AutoIpDefault}
	}
	return obj.autoIpDefaultHolder
}

// description is TBD
// AutoIpDefault returns a PatternAutoPatternDefaultAutoIpDefault
func (obj *autoPatternDefault) HasAutoIpDefault() bool {
	return obj.obj.AutoIpDefault != nil
}

// description is TBD
// SetAutoIpDefault sets the PatternAutoPatternDefaultAutoIpDefault value in the AutoPatternDefault object
func (obj *autoPatternDefault) SetAutoIpDefault(value PatternAutoPatternDefaultAutoIpDefault) AutoPatternDefault {

	obj.autoIpDefaultHolder = nil
	obj.obj.AutoIpDefault = value.msg()

	return obj
}

func (obj *autoPatternDefault) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AutoIpDefault != nil {

		obj.AutoIpDefault().validateObj(vObj, set_default)
	}

}

func (obj *autoPatternDefault) setDefault() {

}
