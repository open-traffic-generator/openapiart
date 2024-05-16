package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternAutoPatternAutoIpCounter *****
type patternAutoPatternAutoIpCounter struct {
	validation
	obj          *openapi.PatternAutoPatternAutoIpCounter
	marshaller   marshalPatternAutoPatternAutoIpCounter
	unMarshaller unMarshalPatternAutoPatternAutoIpCounter
}

func NewPatternAutoPatternAutoIpCounter() PatternAutoPatternAutoIpCounter {
	obj := patternAutoPatternAutoIpCounter{obj: &openapi.PatternAutoPatternAutoIpCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternAutoPatternAutoIpCounter) msg() *openapi.PatternAutoPatternAutoIpCounter {
	return obj.obj
}

func (obj *patternAutoPatternAutoIpCounter) setMsg(msg *openapi.PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIpCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternAutoPatternAutoIpCounter struct {
	obj *patternAutoPatternAutoIpCounter
}

type marshalPatternAutoPatternAutoIpCounter interface {
	// ToProto marshals PatternAutoPatternAutoIpCounter to protobuf object *openapi.PatternAutoPatternAutoIpCounter
	ToProto() (*openapi.PatternAutoPatternAutoIpCounter, error)
	// ToPbText marshals PatternAutoPatternAutoIpCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternAutoPatternAutoIpCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternAutoPatternAutoIpCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternAutoPatternAutoIpCounter struct {
	obj *patternAutoPatternAutoIpCounter
}

type unMarshalPatternAutoPatternAutoIpCounter interface {
	// FromProto unmarshals PatternAutoPatternAutoIpCounter from protobuf object *openapi.PatternAutoPatternAutoIpCounter
	FromProto(msg *openapi.PatternAutoPatternAutoIpCounter) (PatternAutoPatternAutoIpCounter, error)
	// FromPbText unmarshals PatternAutoPatternAutoIpCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternAutoPatternAutoIpCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternAutoPatternAutoIpCounter from JSON text
	FromJson(value string) error
}

func (obj *patternAutoPatternAutoIpCounter) Marshal() marshalPatternAutoPatternAutoIpCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternAutoPatternAutoIpCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternAutoPatternAutoIpCounter) Unmarshal() unMarshalPatternAutoPatternAutoIpCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternAutoPatternAutoIpCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternAutoPatternAutoIpCounter) ToProto() (*openapi.PatternAutoPatternAutoIpCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternAutoPatternAutoIpCounter) FromProto(msg *openapi.PatternAutoPatternAutoIpCounter) (PatternAutoPatternAutoIpCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternAutoPatternAutoIpCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIpCounter) FromPbText(value string) error {
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

func (m *marshalpatternAutoPatternAutoIpCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIpCounter) FromYaml(value string) error {
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

func (m *marshalpatternAutoPatternAutoIpCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIpCounter) FromJson(value string) error {
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

func (obj *patternAutoPatternAutoIpCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternAutoPatternAutoIpCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternAutoPatternAutoIpCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternAutoPatternAutoIpCounter) Clone() (PatternAutoPatternAutoIpCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternAutoPatternAutoIpCounter()
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

// PatternAutoPatternAutoIpCounter is ipv4 counter pattern
type PatternAutoPatternAutoIpCounter interface {
	Validation
	// msg marshals PatternAutoPatternAutoIpCounter to protobuf object *openapi.PatternAutoPatternAutoIpCounter
	// and doesn't set defaults
	msg() *openapi.PatternAutoPatternAutoIpCounter
	// setMsg unmarshals PatternAutoPatternAutoIpCounter from protobuf object *openapi.PatternAutoPatternAutoIpCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIpCounter
	// provides marshal interface
	Marshal() marshalPatternAutoPatternAutoIpCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternAutoPatternAutoIpCounter
	// validate validates PatternAutoPatternAutoIpCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternAutoPatternAutoIpCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternAutoPatternAutoIpCounter.
	Start() string
	// SetStart assigns string provided by user to PatternAutoPatternAutoIpCounter
	SetStart(value string) PatternAutoPatternAutoIpCounter
	// HasStart checks if Start has been set in PatternAutoPatternAutoIpCounter
	HasStart() bool
	// Step returns string, set in PatternAutoPatternAutoIpCounter.
	Step() string
	// SetStep assigns string provided by user to PatternAutoPatternAutoIpCounter
	SetStep(value string) PatternAutoPatternAutoIpCounter
	// HasStep checks if Step has been set in PatternAutoPatternAutoIpCounter
	HasStep() bool
	// Count returns uint32, set in PatternAutoPatternAutoIpCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternAutoPatternAutoIpCounter
	SetCount(value uint32) PatternAutoPatternAutoIpCounter
	// HasCount checks if Count has been set in PatternAutoPatternAutoIpCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternAutoPatternAutoIpCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternAutoPatternAutoIpCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternAutoPatternAutoIpCounter object
func (obj *patternAutoPatternAutoIpCounter) SetStart(value string) PatternAutoPatternAutoIpCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternAutoPatternAutoIpCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternAutoPatternAutoIpCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternAutoPatternAutoIpCounter object
func (obj *patternAutoPatternAutoIpCounter) SetStep(value string) PatternAutoPatternAutoIpCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternAutoPatternAutoIpCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternAutoPatternAutoIpCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternAutoPatternAutoIpCounter object
func (obj *patternAutoPatternAutoIpCounter) SetCount(value uint32) PatternAutoPatternAutoIpCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternAutoPatternAutoIpCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternAutoIpCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternAutoIpCounter.Step"))
		}

	}

}

func (obj *patternAutoPatternAutoIpCounter) setDefault() {
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
