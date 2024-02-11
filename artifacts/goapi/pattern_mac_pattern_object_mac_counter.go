package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternMacPatternObjectMacCounter *****
type patternMacPatternObjectMacCounter struct {
	validation
	obj          *openapi.PatternMacPatternObjectMacCounter
	marshaller   marshalPatternMacPatternObjectMacCounter
	unMarshaller unMarshalPatternMacPatternObjectMacCounter
}

func NewPatternMacPatternObjectMacCounter() PatternMacPatternObjectMacCounter {
	obj := patternMacPatternObjectMacCounter{obj: &openapi.PatternMacPatternObjectMacCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternObjectMacCounter) msg() *openapi.PatternMacPatternObjectMacCounter {
	return obj.obj
}

func (obj *patternMacPatternObjectMacCounter) setMsg(msg *openapi.PatternMacPatternObjectMacCounter) PatternMacPatternObjectMacCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternMacPatternObjectMacCounter struct {
	obj *patternMacPatternObjectMacCounter
}

type marshalPatternMacPatternObjectMacCounter interface {
	// ToProto marshals PatternMacPatternObjectMacCounter to protobuf object *openapi.PatternMacPatternObjectMacCounter
	ToProto() (*openapi.PatternMacPatternObjectMacCounter, error)
	// ToPbText marshals PatternMacPatternObjectMacCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternObjectMacCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternObjectMacCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternMacPatternObjectMacCounter struct {
	obj *patternMacPatternObjectMacCounter
}

type unMarshalPatternMacPatternObjectMacCounter interface {
	// FromProto unmarshals PatternMacPatternObjectMacCounter from protobuf object *openapi.PatternMacPatternObjectMacCounter
	FromProto(msg *openapi.PatternMacPatternObjectMacCounter) (PatternMacPatternObjectMacCounter, error)
	// FromPbText unmarshals PatternMacPatternObjectMacCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternObjectMacCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternObjectMacCounter from JSON text
	FromJson(value string) error
}

func (obj *patternMacPatternObjectMacCounter) Marshal() marshalPatternMacPatternObjectMacCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternMacPatternObjectMacCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternMacPatternObjectMacCounter) Unmarshal() unMarshalPatternMacPatternObjectMacCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternMacPatternObjectMacCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternMacPatternObjectMacCounter) ToProto() (*openapi.PatternMacPatternObjectMacCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternMacPatternObjectMacCounter) FromProto(msg *openapi.PatternMacPatternObjectMacCounter) (PatternMacPatternObjectMacCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternMacPatternObjectMacCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMacCounter) FromPbText(value string) error {
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

func (m *marshalpatternMacPatternObjectMacCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMacCounter) FromYaml(value string) error {
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

func (m *marshalpatternMacPatternObjectMacCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMacCounter) FromJson(value string) error {
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

func (obj *patternMacPatternObjectMacCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternObjectMacCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternObjectMacCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternObjectMacCounter) Clone() (PatternMacPatternObjectMacCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternObjectMacCounter()
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

// PatternMacPatternObjectMacCounter is mac counter pattern
type PatternMacPatternObjectMacCounter interface {
	Validation
	// msg marshals PatternMacPatternObjectMacCounter to protobuf object *openapi.PatternMacPatternObjectMacCounter
	// and doesn't set defaults
	msg() *openapi.PatternMacPatternObjectMacCounter
	// setMsg unmarshals PatternMacPatternObjectMacCounter from protobuf object *openapi.PatternMacPatternObjectMacCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternMacPatternObjectMacCounter) PatternMacPatternObjectMacCounter
	// provides marshal interface
	Marshal() marshalPatternMacPatternObjectMacCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternMacPatternObjectMacCounter
	// validate validates PatternMacPatternObjectMacCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternObjectMacCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternMacPatternObjectMacCounter.
	Start() string
	// SetStart assigns string provided by user to PatternMacPatternObjectMacCounter
	SetStart(value string) PatternMacPatternObjectMacCounter
	// HasStart checks if Start has been set in PatternMacPatternObjectMacCounter
	HasStart() bool
	// Step returns string, set in PatternMacPatternObjectMacCounter.
	Step() string
	// SetStep assigns string provided by user to PatternMacPatternObjectMacCounter
	SetStep(value string) PatternMacPatternObjectMacCounter
	// HasStep checks if Step has been set in PatternMacPatternObjectMacCounter
	HasStep() bool
	// Count returns uint32, set in PatternMacPatternObjectMacCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternMacPatternObjectMacCounter
	SetCount(value uint32) PatternMacPatternObjectMacCounter
	// HasCount checks if Count has been set in PatternMacPatternObjectMacCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternMacPatternObjectMacCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternMacPatternObjectMacCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternMacPatternObjectMacCounter object
func (obj *patternMacPatternObjectMacCounter) SetStart(value string) PatternMacPatternObjectMacCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternMacPatternObjectMacCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternMacPatternObjectMacCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternMacPatternObjectMacCounter object
func (obj *patternMacPatternObjectMacCounter) SetStep(value string) PatternMacPatternObjectMacCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternObjectMacCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternObjectMacCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternMacPatternObjectMacCounter object
func (obj *patternMacPatternObjectMacCounter) SetCount(value uint32) PatternMacPatternObjectMacCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternMacPatternObjectMacCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternObjectMacCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternObjectMacCounter.Step"))
		}

	}

}

func (obj *patternMacPatternObjectMacCounter) setDefault() {
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
