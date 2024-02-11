package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIntegerPatternIntegerCounter *****
type patternIntegerPatternIntegerCounter struct {
	validation
	obj          *openapi.PatternIntegerPatternIntegerCounter
	marshaller   marshalPatternIntegerPatternIntegerCounter
	unMarshaller unMarshalPatternIntegerPatternIntegerCounter
}

func NewPatternIntegerPatternIntegerCounter() PatternIntegerPatternIntegerCounter {
	obj := patternIntegerPatternIntegerCounter{obj: &openapi.PatternIntegerPatternIntegerCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternIntegerCounter) msg() *openapi.PatternIntegerPatternIntegerCounter {
	return obj.obj
}

func (obj *patternIntegerPatternIntegerCounter) setMsg(msg *openapi.PatternIntegerPatternIntegerCounter) PatternIntegerPatternIntegerCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIntegerPatternIntegerCounter struct {
	obj *patternIntegerPatternIntegerCounter
}

type marshalPatternIntegerPatternIntegerCounter interface {
	// ToProto marshals PatternIntegerPatternIntegerCounter to protobuf object *openapi.PatternIntegerPatternIntegerCounter
	ToProto() (*openapi.PatternIntegerPatternIntegerCounter, error)
	// ToPbText marshals PatternIntegerPatternIntegerCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternIntegerCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternIntegerCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIntegerPatternIntegerCounter struct {
	obj *patternIntegerPatternIntegerCounter
}

type unMarshalPatternIntegerPatternIntegerCounter interface {
	// FromProto unmarshals PatternIntegerPatternIntegerCounter from protobuf object *openapi.PatternIntegerPatternIntegerCounter
	FromProto(msg *openapi.PatternIntegerPatternIntegerCounter) (PatternIntegerPatternIntegerCounter, error)
	// FromPbText unmarshals PatternIntegerPatternIntegerCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternIntegerCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternIntegerCounter from JSON text
	FromJson(value string) error
}

func (obj *patternIntegerPatternIntegerCounter) Marshal() marshalPatternIntegerPatternIntegerCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIntegerPatternIntegerCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIntegerPatternIntegerCounter) Unmarshal() unMarshalPatternIntegerPatternIntegerCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIntegerPatternIntegerCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIntegerPatternIntegerCounter) ToProto() (*openapi.PatternIntegerPatternIntegerCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIntegerPatternIntegerCounter) FromProto(msg *openapi.PatternIntegerPatternIntegerCounter) (PatternIntegerPatternIntegerCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIntegerPatternIntegerCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternIntegerPatternIntegerCounter) FromPbText(value string) error {
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

func (m *marshalpatternIntegerPatternIntegerCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternIntegerPatternIntegerCounter) FromYaml(value string) error {
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

func (m *marshalpatternIntegerPatternIntegerCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternIntegerPatternIntegerCounter) FromJson(value string) error {
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

func (obj *patternIntegerPatternIntegerCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternIntegerCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternIntegerCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternIntegerCounter) Clone() (PatternIntegerPatternIntegerCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternIntegerCounter()
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

// PatternIntegerPatternIntegerCounter is integer counter pattern
type PatternIntegerPatternIntegerCounter interface {
	Validation
	// msg marshals PatternIntegerPatternIntegerCounter to protobuf object *openapi.PatternIntegerPatternIntegerCounter
	// and doesn't set defaults
	msg() *openapi.PatternIntegerPatternIntegerCounter
	// setMsg unmarshals PatternIntegerPatternIntegerCounter from protobuf object *openapi.PatternIntegerPatternIntegerCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternIntegerPatternIntegerCounter) PatternIntegerPatternIntegerCounter
	// provides marshal interface
	Marshal() marshalPatternIntegerPatternIntegerCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIntegerPatternIntegerCounter
	// validate validates PatternIntegerPatternIntegerCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternIntegerCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternIntegerPatternIntegerCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetStart(value uint32) PatternIntegerPatternIntegerCounter
	// HasStart checks if Start has been set in PatternIntegerPatternIntegerCounter
	HasStart() bool
	// Step returns uint32, set in PatternIntegerPatternIntegerCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetStep(value uint32) PatternIntegerPatternIntegerCounter
	// HasStep checks if Step has been set in PatternIntegerPatternIntegerCounter
	HasStep() bool
	// Count returns uint32, set in PatternIntegerPatternIntegerCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetCount(value uint32) PatternIntegerPatternIntegerCounter
	// HasCount checks if Count has been set in PatternIntegerPatternIntegerCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetStart(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetStep(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetCount(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIntegerPatternIntegerCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternIntegerPatternIntegerCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}
