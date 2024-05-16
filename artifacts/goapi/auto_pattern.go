package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** AutoPattern *****
type autoPattern struct {
	validation
	obj          *openapi.AutoPattern
	marshaller   marshalAutoPattern
	unMarshaller unMarshalAutoPattern
	autoIpHolder PatternAutoPatternAutoIp
}

func NewAutoPattern() AutoPattern {
	obj := autoPattern{obj: &openapi.AutoPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *autoPattern) msg() *openapi.AutoPattern {
	return obj.obj
}

func (obj *autoPattern) setMsg(msg *openapi.AutoPattern) AutoPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalautoPattern struct {
	obj *autoPattern
}

type marshalAutoPattern interface {
	// ToProto marshals AutoPattern to protobuf object *openapi.AutoPattern
	ToProto() (*openapi.AutoPattern, error)
	// ToPbText marshals AutoPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals AutoPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals AutoPattern to JSON text
	ToJson() (string, error)
}

type unMarshalautoPattern struct {
	obj *autoPattern
}

type unMarshalAutoPattern interface {
	// FromProto unmarshals AutoPattern from protobuf object *openapi.AutoPattern
	FromProto(msg *openapi.AutoPattern) (AutoPattern, error)
	// FromPbText unmarshals AutoPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals AutoPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals AutoPattern from JSON text
	FromJson(value string) error
}

func (obj *autoPattern) Marshal() marshalAutoPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalautoPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *autoPattern) Unmarshal() unMarshalAutoPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalautoPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalautoPattern) ToProto() (*openapi.AutoPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalautoPattern) FromProto(msg *openapi.AutoPattern) (AutoPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalautoPattern) ToPbText() (string, error) {
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

func (m *unMarshalautoPattern) FromPbText(value string) error {
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

func (m *marshalautoPattern) ToYaml() (string, error) {
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

func (m *unMarshalautoPattern) FromYaml(value string) error {
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

func (m *marshalautoPattern) ToJson() (string, error) {
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

func (m *unMarshalautoPattern) FromJson(value string) error {
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

func (obj *autoPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *autoPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *autoPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *autoPattern) Clone() (AutoPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewAutoPattern()
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

func (obj *autoPattern) setNil() {
	obj.autoIpHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// AutoPattern is test auto pattern
type AutoPattern interface {
	Validation
	// msg marshals AutoPattern to protobuf object *openapi.AutoPattern
	// and doesn't set defaults
	msg() *openapi.AutoPattern
	// setMsg unmarshals AutoPattern from protobuf object *openapi.AutoPattern
	// and doesn't set defaults
	setMsg(*openapi.AutoPattern) AutoPattern
	// provides marshal interface
	Marshal() marshalAutoPattern
	// provides unmarshal interface
	Unmarshal() unMarshalAutoPattern
	// validate validates AutoPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (AutoPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AutoIp returns PatternAutoPatternAutoIp, set in AutoPattern.
	// PatternAutoPatternAutoIp is tBD
	AutoIp() PatternAutoPatternAutoIp
	// SetAutoIp assigns PatternAutoPatternAutoIp provided by user to AutoPattern.
	// PatternAutoPatternAutoIp is tBD
	SetAutoIp(value PatternAutoPatternAutoIp) AutoPattern
	// HasAutoIp checks if AutoIp has been set in AutoPattern
	HasAutoIp() bool
	setNil()
}

// description is TBD
// AutoIp returns a PatternAutoPatternAutoIp
func (obj *autoPattern) AutoIp() PatternAutoPatternAutoIp {
	if obj.obj.AutoIp == nil {
		obj.obj.AutoIp = NewPatternAutoPatternAutoIp().msg()
	}
	if obj.autoIpHolder == nil {
		obj.autoIpHolder = &patternAutoPatternAutoIp{obj: obj.obj.AutoIp}
	}
	return obj.autoIpHolder
}

// description is TBD
// AutoIp returns a PatternAutoPatternAutoIp
func (obj *autoPattern) HasAutoIp() bool {
	return obj.obj.AutoIp != nil
}

// description is TBD
// SetAutoIp sets the PatternAutoPatternAutoIp value in the AutoPattern object
func (obj *autoPattern) SetAutoIp(value PatternAutoPatternAutoIp) AutoPattern {

	obj.autoIpHolder = nil
	obj.obj.AutoIp = value.msg()

	return obj
}

func (obj *autoPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.AutoIp != nil {

		obj.AutoIp().validateObj(vObj, set_default)
	}

}

func (obj *autoPattern) setDefault() {

}
