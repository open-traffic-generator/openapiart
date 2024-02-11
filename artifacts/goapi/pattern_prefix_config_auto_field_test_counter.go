package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternPrefixConfigAutoFieldTestCounter *****
type patternPrefixConfigAutoFieldTestCounter struct {
	validation
	obj          *openapi.PatternPrefixConfigAutoFieldTestCounter
	marshaller   marshalPatternPrefixConfigAutoFieldTestCounter
	unMarshaller unMarshalPatternPrefixConfigAutoFieldTestCounter
}

func NewPatternPrefixConfigAutoFieldTestCounter() PatternPrefixConfigAutoFieldTestCounter {
	obj := patternPrefixConfigAutoFieldTestCounter{obj: &openapi.PatternPrefixConfigAutoFieldTestCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) msg() *openapi.PatternPrefixConfigAutoFieldTestCounter {
	return obj.obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) setMsg(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTestCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternPrefixConfigAutoFieldTestCounter struct {
	obj *patternPrefixConfigAutoFieldTestCounter
}

type marshalPatternPrefixConfigAutoFieldTestCounter interface {
	// ToProto marshals PatternPrefixConfigAutoFieldTestCounter to protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	ToProto() (*openapi.PatternPrefixConfigAutoFieldTestCounter, error)
	// ToPbText marshals PatternPrefixConfigAutoFieldTestCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigAutoFieldTestCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigAutoFieldTestCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternPrefixConfigAutoFieldTestCounter struct {
	obj *patternPrefixConfigAutoFieldTestCounter
}

type unMarshalPatternPrefixConfigAutoFieldTestCounter interface {
	// FromProto unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	FromProto(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) (PatternPrefixConfigAutoFieldTestCounter, error)
	// FromPbText unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigAutoFieldTestCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigAutoFieldTestCounter from JSON text
	FromJson(value string) error
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Marshal() marshalPatternPrefixConfigAutoFieldTestCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternPrefixConfigAutoFieldTestCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Unmarshal() unMarshalPatternPrefixConfigAutoFieldTestCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternPrefixConfigAutoFieldTestCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternPrefixConfigAutoFieldTestCounter) ToProto() (*openapi.PatternPrefixConfigAutoFieldTestCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternPrefixConfigAutoFieldTestCounter) FromProto(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) (PatternPrefixConfigAutoFieldTestCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternPrefixConfigAutoFieldTestCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTestCounter) FromPbText(value string) error {
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

func (m *marshalpatternPrefixConfigAutoFieldTestCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTestCounter) FromYaml(value string) error {
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

func (m *marshalpatternPrefixConfigAutoFieldTestCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTestCounter) FromJson(value string) error {
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

func (obj *patternPrefixConfigAutoFieldTestCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTestCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTestCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Clone() (PatternPrefixConfigAutoFieldTestCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigAutoFieldTestCounter()
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

// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
type PatternPrefixConfigAutoFieldTestCounter interface {
	Validation
	// msg marshals PatternPrefixConfigAutoFieldTestCounter to protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	// and doesn't set defaults
	msg() *openapi.PatternPrefixConfigAutoFieldTestCounter
	// setMsg unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTestCounter
	// provides marshal interface
	Marshal() marshalPatternPrefixConfigAutoFieldTestCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternPrefixConfigAutoFieldTestCounter
	// validate validates PatternPrefixConfigAutoFieldTestCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigAutoFieldTestCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetStart(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasStart checks if Start has been set in PatternPrefixConfigAutoFieldTestCounter
	HasStart() bool
	// Step returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetStep(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasStep checks if Step has been set in PatternPrefixConfigAutoFieldTestCounter
	HasStep() bool
	// Count returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetCount(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasCount checks if Count has been set in PatternPrefixConfigAutoFieldTestCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetStart(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetStep(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetCount(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternPrefixConfigAutoFieldTestCounter) setDefault() {
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
