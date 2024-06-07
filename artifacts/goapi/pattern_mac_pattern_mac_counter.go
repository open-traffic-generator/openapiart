package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternMacPatternMacCounter *****
type patternMacPatternMacCounter struct {
	validation
	obj          *openapi.PatternMacPatternMacCounter
	marshaller   marshalPatternMacPatternMacCounter
	unMarshaller unMarshalPatternMacPatternMacCounter
}

func NewPatternMacPatternMacCounter() PatternMacPatternMacCounter {
	obj := patternMacPatternMacCounter{obj: &openapi.PatternMacPatternMacCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternMacCounter) msg() *openapi.PatternMacPatternMacCounter {
	return obj.obj
}

func (obj *patternMacPatternMacCounter) setMsg(msg *openapi.PatternMacPatternMacCounter) PatternMacPatternMacCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternMacPatternMacCounter struct {
	obj *patternMacPatternMacCounter
}

type marshalPatternMacPatternMacCounter interface {
	// ToProto marshals PatternMacPatternMacCounter to protobuf object *openapi.PatternMacPatternMacCounter
	ToProto() (*openapi.PatternMacPatternMacCounter, error)
	// ToPbText marshals PatternMacPatternMacCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternMacCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternMacCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternMacPatternMacCounter struct {
	obj *patternMacPatternMacCounter
}

type unMarshalPatternMacPatternMacCounter interface {
	// FromProto unmarshals PatternMacPatternMacCounter from protobuf object *openapi.PatternMacPatternMacCounter
	FromProto(msg *openapi.PatternMacPatternMacCounter) (PatternMacPatternMacCounter, error)
	// FromPbText unmarshals PatternMacPatternMacCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternMacCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternMacCounter from JSON text
	FromJson(value string) error
}

func (obj *patternMacPatternMacCounter) Marshal() marshalPatternMacPatternMacCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternMacPatternMacCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternMacPatternMacCounter) Unmarshal() unMarshalPatternMacPatternMacCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternMacPatternMacCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternMacPatternMacCounter) ToProto() (*openapi.PatternMacPatternMacCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternMacPatternMacCounter) FromProto(msg *openapi.PatternMacPatternMacCounter) (PatternMacPatternMacCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternMacPatternMacCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternMacPatternMacCounter) FromPbText(value string) error {
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

func (m *marshalpatternMacPatternMacCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternMacPatternMacCounter) FromYaml(value string) error {
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

func (m *marshalpatternMacPatternMacCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternMacPatternMacCounter) FromJson(value string) error {
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

func (obj *patternMacPatternMacCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternMacCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternMacCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternMacCounter) Clone() (PatternMacPatternMacCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternMacCounter()
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

// PatternMacPatternMacCounter is mac counter pattern
type PatternMacPatternMacCounter interface {
	Validation
	// msg marshals PatternMacPatternMacCounter to protobuf object *openapi.PatternMacPatternMacCounter
	// and doesn't set defaults
	msg() *openapi.PatternMacPatternMacCounter
	// setMsg unmarshals PatternMacPatternMacCounter from protobuf object *openapi.PatternMacPatternMacCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternMacPatternMacCounter) PatternMacPatternMacCounter
	// provides marshal interface
	Marshal() marshalPatternMacPatternMacCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternMacPatternMacCounter
	// validate validates PatternMacPatternMacCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternMacCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternMacPatternMacCounter.
	Start() string
	// SetStart assigns string provided by user to PatternMacPatternMacCounter
	SetStart(value string) PatternMacPatternMacCounter
	// HasStart checks if Start has been set in PatternMacPatternMacCounter
	HasStart() bool
	// Step returns string, set in PatternMacPatternMacCounter.
	Step() string
	// SetStep assigns string provided by user to PatternMacPatternMacCounter
	SetStep(value string) PatternMacPatternMacCounter
	// HasStep checks if Step has been set in PatternMacPatternMacCounter
	HasStep() bool
	// Count returns uint32, set in PatternMacPatternMacCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternMacPatternMacCounter
	SetCount(value uint32) PatternMacPatternMacCounter
	// HasCount checks if Count has been set in PatternMacPatternMacCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternMacPatternMacCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternMacPatternMacCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetStart(value string) PatternMacPatternMacCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternMacPatternMacCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternMacPatternMacCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetStep(value string) PatternMacPatternMacCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternMacCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternMacCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetCount(value uint32) PatternMacPatternMacCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternMacPatternMacCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Step"))
		}

	}

}

func (obj *patternMacPatternMacCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("00:00:00:00:00:00")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
