package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternAutoPatternDefaultAutoIpDefaultCounter *****
type patternAutoPatternDefaultAutoIpDefaultCounter struct {
	validation
	obj          *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	marshaller   marshalPatternAutoPatternDefaultAutoIpDefaultCounter
	unMarshaller unMarshalPatternAutoPatternDefaultAutoIpDefaultCounter
}

func NewPatternAutoPatternDefaultAutoIpDefaultCounter() PatternAutoPatternDefaultAutoIpDefaultCounter {
	obj := patternAutoPatternDefaultAutoIpDefaultCounter{obj: &openapi.PatternAutoPatternDefaultAutoIpDefaultCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) msg() *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter {
	return obj.obj
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) setMsg(msg *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefaultCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternAutoPatternDefaultAutoIpDefaultCounter struct {
	obj *patternAutoPatternDefaultAutoIpDefaultCounter
}

type marshalPatternAutoPatternDefaultAutoIpDefaultCounter interface {
	// ToProto marshals PatternAutoPatternDefaultAutoIpDefaultCounter to protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	ToProto() (*openapi.PatternAutoPatternDefaultAutoIpDefaultCounter, error)
	// ToPbText marshals PatternAutoPatternDefaultAutoIpDefaultCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternAutoPatternDefaultAutoIpDefaultCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternAutoPatternDefaultAutoIpDefaultCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter struct {
	obj *patternAutoPatternDefaultAutoIpDefaultCounter
}

type unMarshalPatternAutoPatternDefaultAutoIpDefaultCounter interface {
	// FromProto unmarshals PatternAutoPatternDefaultAutoIpDefaultCounter from protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	FromProto(msg *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter) (PatternAutoPatternDefaultAutoIpDefaultCounter, error)
	// FromPbText unmarshals PatternAutoPatternDefaultAutoIpDefaultCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternAutoPatternDefaultAutoIpDefaultCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternAutoPatternDefaultAutoIpDefaultCounter from JSON text
	FromJson(value string) error
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Marshal() marshalPatternAutoPatternDefaultAutoIpDefaultCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternAutoPatternDefaultAutoIpDefaultCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Unmarshal() unMarshalPatternAutoPatternDefaultAutoIpDefaultCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternAutoPatternDefaultAutoIpDefaultCounter) ToProto() (*openapi.PatternAutoPatternDefaultAutoIpDefaultCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter) FromProto(msg *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter) (PatternAutoPatternDefaultAutoIpDefaultCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternAutoPatternDefaultAutoIpDefaultCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter) FromPbText(value string) error {
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

func (m *marshalpatternAutoPatternDefaultAutoIpDefaultCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter) FromYaml(value string) error {
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

func (m *marshalpatternAutoPatternDefaultAutoIpDefaultCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefaultCounter) FromJson(value string) error {
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

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Clone() (PatternAutoPatternDefaultAutoIpDefaultCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternAutoPatternDefaultAutoIpDefaultCounter()
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

// PatternAutoPatternDefaultAutoIpDefaultCounter is ipv4 counter pattern
type PatternAutoPatternDefaultAutoIpDefaultCounter interface {
	Validation
	// msg marshals PatternAutoPatternDefaultAutoIpDefaultCounter to protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	// and doesn't set defaults
	msg() *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	// setMsg unmarshals PatternAutoPatternDefaultAutoIpDefaultCounter from protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefaultCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefaultCounter
	// provides marshal interface
	Marshal() marshalPatternAutoPatternDefaultAutoIpDefaultCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternAutoPatternDefaultAutoIpDefaultCounter
	// validate validates PatternAutoPatternDefaultAutoIpDefaultCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternAutoPatternDefaultAutoIpDefaultCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternAutoPatternDefaultAutoIpDefaultCounter.
	Start() string
	// SetStart assigns string provided by user to PatternAutoPatternDefaultAutoIpDefaultCounter
	SetStart(value string) PatternAutoPatternDefaultAutoIpDefaultCounter
	// HasStart checks if Start has been set in PatternAutoPatternDefaultAutoIpDefaultCounter
	HasStart() bool
	// Step returns string, set in PatternAutoPatternDefaultAutoIpDefaultCounter.
	Step() string
	// SetStep assigns string provided by user to PatternAutoPatternDefaultAutoIpDefaultCounter
	SetStep(value string) PatternAutoPatternDefaultAutoIpDefaultCounter
	// HasStep checks if Step has been set in PatternAutoPatternDefaultAutoIpDefaultCounter
	HasStep() bool
	// Count returns uint32, set in PatternAutoPatternDefaultAutoIpDefaultCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternAutoPatternDefaultAutoIpDefaultCounter
	SetCount(value uint32) PatternAutoPatternDefaultAutoIpDefaultCounter
	// HasCount checks if Count has been set in PatternAutoPatternDefaultAutoIpDefaultCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternAutoPatternDefaultAutoIpDefaultCounter object
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) SetStart(value string) PatternAutoPatternDefaultAutoIpDefaultCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternAutoPatternDefaultAutoIpDefaultCounter object
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) SetStep(value string) PatternAutoPatternDefaultAutoIpDefaultCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternAutoPatternDefaultAutoIpDefaultCounter object
func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) SetCount(value uint32) PatternAutoPatternDefaultAutoIpDefaultCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternDefaultAutoIpDefaultCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternDefaultAutoIpDefaultCounter.Step"))
		}

	}

}

func (obj *patternAutoPatternDefaultAutoIpDefaultCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
