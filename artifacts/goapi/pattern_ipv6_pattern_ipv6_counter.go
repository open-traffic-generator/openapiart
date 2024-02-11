package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv6PatternIpv6Counter *****
type patternIpv6PatternIpv6Counter struct {
	validation
	obj          *openapi.PatternIpv6PatternIpv6Counter
	marshaller   marshalPatternIpv6PatternIpv6Counter
	unMarshaller unMarshalPatternIpv6PatternIpv6Counter
}

func NewPatternIpv6PatternIpv6Counter() PatternIpv6PatternIpv6Counter {
	obj := patternIpv6PatternIpv6Counter{obj: &openapi.PatternIpv6PatternIpv6Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv6PatternIpv6Counter) msg() *openapi.PatternIpv6PatternIpv6Counter {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6Counter) setMsg(msg *openapi.PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv6PatternIpv6Counter struct {
	obj *patternIpv6PatternIpv6Counter
}

type marshalPatternIpv6PatternIpv6Counter interface {
	// ToProto marshals PatternIpv6PatternIpv6Counter to protobuf object *openapi.PatternIpv6PatternIpv6Counter
	ToProto() (*openapi.PatternIpv6PatternIpv6Counter, error)
	// ToPbText marshals PatternIpv6PatternIpv6Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv6PatternIpv6Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv6PatternIpv6Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv6PatternIpv6Counter struct {
	obj *patternIpv6PatternIpv6Counter
}

type unMarshalPatternIpv6PatternIpv6Counter interface {
	// FromProto unmarshals PatternIpv6PatternIpv6Counter from protobuf object *openapi.PatternIpv6PatternIpv6Counter
	FromProto(msg *openapi.PatternIpv6PatternIpv6Counter) (PatternIpv6PatternIpv6Counter, error)
	// FromPbText unmarshals PatternIpv6PatternIpv6Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv6PatternIpv6Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv6PatternIpv6Counter from JSON text
	FromJson(value string) error
}

func (obj *patternIpv6PatternIpv6Counter) Marshal() marshalPatternIpv6PatternIpv6Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv6PatternIpv6Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv6PatternIpv6Counter) Unmarshal() unMarshalPatternIpv6PatternIpv6Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv6PatternIpv6Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv6PatternIpv6Counter) ToProto() (*openapi.PatternIpv6PatternIpv6Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv6PatternIpv6Counter) FromProto(msg *openapi.PatternIpv6PatternIpv6Counter) (PatternIpv6PatternIpv6Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv6PatternIpv6Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6Counter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalpatternIpv6PatternIpv6Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6Counter) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalpatternIpv6PatternIpv6Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6Counter) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv6PatternIpv6Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv6PatternIpv6Counter) Clone() (PatternIpv6PatternIpv6Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv6PatternIpv6Counter()
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

// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
type PatternIpv6PatternIpv6Counter interface {
	Validation
	// msg marshals PatternIpv6PatternIpv6Counter to protobuf object *openapi.PatternIpv6PatternIpv6Counter
	// and doesn't set defaults
	msg() *openapi.PatternIpv6PatternIpv6Counter
	// setMsg unmarshals PatternIpv6PatternIpv6Counter from protobuf object *openapi.PatternIpv6PatternIpv6Counter
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6Counter
	// provides marshal interface
	Marshal() marshalPatternIpv6PatternIpv6Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv6PatternIpv6Counter
	// validate validates PatternIpv6PatternIpv6Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv6PatternIpv6Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternIpv6PatternIpv6Counter.
	Start() string
	// SetStart assigns string provided by user to PatternIpv6PatternIpv6Counter
	SetStart(value string) PatternIpv6PatternIpv6Counter
	// HasStart checks if Start has been set in PatternIpv6PatternIpv6Counter
	HasStart() bool
	// Step returns string, set in PatternIpv6PatternIpv6Counter.
	Step() string
	// SetStep assigns string provided by user to PatternIpv6PatternIpv6Counter
	SetStep(value string) PatternIpv6PatternIpv6Counter
	// HasStep checks if Step has been set in PatternIpv6PatternIpv6Counter
	HasStep() bool
	// Count returns uint32, set in PatternIpv6PatternIpv6Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIpv6PatternIpv6Counter
	SetCount(value uint32) PatternIpv6PatternIpv6Counter
	// HasCount checks if Count has been set in PatternIpv6PatternIpv6Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternIpv6PatternIpv6Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternIpv6PatternIpv6Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetStart(value string) PatternIpv6PatternIpv6Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternIpv6PatternIpv6Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternIpv6PatternIpv6Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetStep(value string) PatternIpv6PatternIpv6Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIpv6PatternIpv6Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIpv6PatternIpv6Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetCount(value uint32) PatternIpv6PatternIpv6Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIpv6PatternIpv6Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Step"))
		}

	}

}

func (obj *patternIpv6PatternIpv6Counter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
