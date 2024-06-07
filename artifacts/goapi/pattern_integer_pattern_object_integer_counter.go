package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIntegerPatternObjectIntegerCounter *****
type patternIntegerPatternObjectIntegerCounter struct {
	validation
	obj          *openapi.PatternIntegerPatternObjectIntegerCounter
	marshaller   marshalPatternIntegerPatternObjectIntegerCounter
	unMarshaller unMarshalPatternIntegerPatternObjectIntegerCounter
}

func NewPatternIntegerPatternObjectIntegerCounter() PatternIntegerPatternObjectIntegerCounter {
	obj := patternIntegerPatternObjectIntegerCounter{obj: &openapi.PatternIntegerPatternObjectIntegerCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternObjectIntegerCounter) msg() *openapi.PatternIntegerPatternObjectIntegerCounter {
	return obj.obj
}

func (obj *patternIntegerPatternObjectIntegerCounter) setMsg(msg *openapi.PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectIntegerCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIntegerPatternObjectIntegerCounter struct {
	obj *patternIntegerPatternObjectIntegerCounter
}

type marshalPatternIntegerPatternObjectIntegerCounter interface {
	// ToProto marshals PatternIntegerPatternObjectIntegerCounter to protobuf object *openapi.PatternIntegerPatternObjectIntegerCounter
	ToProto() (*openapi.PatternIntegerPatternObjectIntegerCounter, error)
	// ToPbText marshals PatternIntegerPatternObjectIntegerCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternObjectIntegerCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternObjectIntegerCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIntegerPatternObjectIntegerCounter struct {
	obj *patternIntegerPatternObjectIntegerCounter
}

type unMarshalPatternIntegerPatternObjectIntegerCounter interface {
	// FromProto unmarshals PatternIntegerPatternObjectIntegerCounter from protobuf object *openapi.PatternIntegerPatternObjectIntegerCounter
	FromProto(msg *openapi.PatternIntegerPatternObjectIntegerCounter) (PatternIntegerPatternObjectIntegerCounter, error)
	// FromPbText unmarshals PatternIntegerPatternObjectIntegerCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternObjectIntegerCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternObjectIntegerCounter from JSON text
	FromJson(value string) error
}

func (obj *patternIntegerPatternObjectIntegerCounter) Marshal() marshalPatternIntegerPatternObjectIntegerCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIntegerPatternObjectIntegerCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIntegerPatternObjectIntegerCounter) Unmarshal() unMarshalPatternIntegerPatternObjectIntegerCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIntegerPatternObjectIntegerCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIntegerPatternObjectIntegerCounter) ToProto() (*openapi.PatternIntegerPatternObjectIntegerCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIntegerPatternObjectIntegerCounter) FromProto(msg *openapi.PatternIntegerPatternObjectIntegerCounter) (PatternIntegerPatternObjectIntegerCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIntegerPatternObjectIntegerCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectIntegerCounter) FromPbText(value string) error {
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

func (m *marshalpatternIntegerPatternObjectIntegerCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectIntegerCounter) FromYaml(value string) error {
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

func (m *marshalpatternIntegerPatternObjectIntegerCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectIntegerCounter) FromJson(value string) error {
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

func (obj *patternIntegerPatternObjectIntegerCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternObjectIntegerCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternObjectIntegerCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternObjectIntegerCounter) Clone() (PatternIntegerPatternObjectIntegerCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternObjectIntegerCounter()
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

// PatternIntegerPatternObjectIntegerCounter is integer counter pattern
type PatternIntegerPatternObjectIntegerCounter interface {
	Validation
	// msg marshals PatternIntegerPatternObjectIntegerCounter to protobuf object *openapi.PatternIntegerPatternObjectIntegerCounter
	// and doesn't set defaults
	msg() *openapi.PatternIntegerPatternObjectIntegerCounter
	// setMsg unmarshals PatternIntegerPatternObjectIntegerCounter from protobuf object *openapi.PatternIntegerPatternObjectIntegerCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectIntegerCounter
	// provides marshal interface
	Marshal() marshalPatternIntegerPatternObjectIntegerCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIntegerPatternObjectIntegerCounter
	// validate validates PatternIntegerPatternObjectIntegerCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternObjectIntegerCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternIntegerPatternObjectIntegerCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternIntegerPatternObjectIntegerCounter
	SetStart(value uint32) PatternIntegerPatternObjectIntegerCounter
	// HasStart checks if Start has been set in PatternIntegerPatternObjectIntegerCounter
	HasStart() bool
	// Step returns uint32, set in PatternIntegerPatternObjectIntegerCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternIntegerPatternObjectIntegerCounter
	SetStep(value uint32) PatternIntegerPatternObjectIntegerCounter
	// HasStep checks if Step has been set in PatternIntegerPatternObjectIntegerCounter
	HasStep() bool
	// Count returns uint32, set in PatternIntegerPatternObjectIntegerCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIntegerPatternObjectIntegerCounter
	SetCount(value uint32) PatternIntegerPatternObjectIntegerCounter
	// HasCount checks if Count has been set in PatternIntegerPatternObjectIntegerCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternIntegerPatternObjectIntegerCounter object
func (obj *patternIntegerPatternObjectIntegerCounter) SetStart(value uint32) PatternIntegerPatternObjectIntegerCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternIntegerPatternObjectIntegerCounter object
func (obj *patternIntegerPatternObjectIntegerCounter) SetStep(value uint32) PatternIntegerPatternObjectIntegerCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternObjectIntegerCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIntegerPatternObjectIntegerCounter object
func (obj *patternIntegerPatternObjectIntegerCounter) SetCount(value uint32) PatternIntegerPatternObjectIntegerCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIntegerPatternObjectIntegerCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternObjectIntegerCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternObjectIntegerCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternObjectIntegerCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternIntegerPatternObjectIntegerCounter) setDefault() {
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
