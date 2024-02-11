package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternSignedIntegerPatternIntegerCounter *****
type patternSignedIntegerPatternIntegerCounter struct {
	validation
	obj          *openapi.PatternSignedIntegerPatternIntegerCounter
	marshaller   marshalPatternSignedIntegerPatternIntegerCounter
	unMarshaller unMarshalPatternSignedIntegerPatternIntegerCounter
}

func NewPatternSignedIntegerPatternIntegerCounter() PatternSignedIntegerPatternIntegerCounter {
	obj := patternSignedIntegerPatternIntegerCounter{obj: &openapi.PatternSignedIntegerPatternIntegerCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternSignedIntegerPatternIntegerCounter) msg() *openapi.PatternSignedIntegerPatternIntegerCounter {
	return obj.obj
}

func (obj *patternSignedIntegerPatternIntegerCounter) setMsg(msg *openapi.PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternIntegerCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternSignedIntegerPatternIntegerCounter struct {
	obj *patternSignedIntegerPatternIntegerCounter
}

type marshalPatternSignedIntegerPatternIntegerCounter interface {
	// ToProto marshals PatternSignedIntegerPatternIntegerCounter to protobuf object *openapi.PatternSignedIntegerPatternIntegerCounter
	ToProto() (*openapi.PatternSignedIntegerPatternIntegerCounter, error)
	// ToPbText marshals PatternSignedIntegerPatternIntegerCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternSignedIntegerPatternIntegerCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternSignedIntegerPatternIntegerCounter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternSignedIntegerPatternIntegerCounter struct {
	obj *patternSignedIntegerPatternIntegerCounter
}

type unMarshalPatternSignedIntegerPatternIntegerCounter interface {
	// FromProto unmarshals PatternSignedIntegerPatternIntegerCounter from protobuf object *openapi.PatternSignedIntegerPatternIntegerCounter
	FromProto(msg *openapi.PatternSignedIntegerPatternIntegerCounter) (PatternSignedIntegerPatternIntegerCounter, error)
	// FromPbText unmarshals PatternSignedIntegerPatternIntegerCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternSignedIntegerPatternIntegerCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternSignedIntegerPatternIntegerCounter from JSON text
	FromJson(value string) error
}

func (obj *patternSignedIntegerPatternIntegerCounter) Marshal() marshalPatternSignedIntegerPatternIntegerCounter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternSignedIntegerPatternIntegerCounter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternSignedIntegerPatternIntegerCounter) Unmarshal() unMarshalPatternSignedIntegerPatternIntegerCounter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternSignedIntegerPatternIntegerCounter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternSignedIntegerPatternIntegerCounter) ToProto() (*openapi.PatternSignedIntegerPatternIntegerCounter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternSignedIntegerPatternIntegerCounter) FromProto(msg *openapi.PatternSignedIntegerPatternIntegerCounter) (PatternSignedIntegerPatternIntegerCounter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternSignedIntegerPatternIntegerCounter) ToPbText() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternIntegerCounter) FromPbText(value string) error {
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

func (m *marshalpatternSignedIntegerPatternIntegerCounter) ToYaml() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternIntegerCounter) FromYaml(value string) error {
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

func (m *marshalpatternSignedIntegerPatternIntegerCounter) ToJson() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternIntegerCounter) FromJson(value string) error {
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

func (obj *patternSignedIntegerPatternIntegerCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternSignedIntegerPatternIntegerCounter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternSignedIntegerPatternIntegerCounter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternSignedIntegerPatternIntegerCounter) Clone() (PatternSignedIntegerPatternIntegerCounter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternSignedIntegerPatternIntegerCounter()
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

// PatternSignedIntegerPatternIntegerCounter is integer counter pattern
type PatternSignedIntegerPatternIntegerCounter interface {
	Validation
	// msg marshals PatternSignedIntegerPatternIntegerCounter to protobuf object *openapi.PatternSignedIntegerPatternIntegerCounter
	// and doesn't set defaults
	msg() *openapi.PatternSignedIntegerPatternIntegerCounter
	// setMsg unmarshals PatternSignedIntegerPatternIntegerCounter from protobuf object *openapi.PatternSignedIntegerPatternIntegerCounter
	// and doesn't set defaults
	setMsg(*openapi.PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternIntegerCounter
	// provides marshal interface
	Marshal() marshalPatternSignedIntegerPatternIntegerCounter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternSignedIntegerPatternIntegerCounter
	// validate validates PatternSignedIntegerPatternIntegerCounter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternSignedIntegerPatternIntegerCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns int32, set in PatternSignedIntegerPatternIntegerCounter.
	Start() int32
	// SetStart assigns int32 provided by user to PatternSignedIntegerPatternIntegerCounter
	SetStart(value int32) PatternSignedIntegerPatternIntegerCounter
	// HasStart checks if Start has been set in PatternSignedIntegerPatternIntegerCounter
	HasStart() bool
	// Step returns int32, set in PatternSignedIntegerPatternIntegerCounter.
	Step() int32
	// SetStep assigns int32 provided by user to PatternSignedIntegerPatternIntegerCounter
	SetStep(value int32) PatternSignedIntegerPatternIntegerCounter
	// HasStep checks if Step has been set in PatternSignedIntegerPatternIntegerCounter
	HasStep() bool
	// Count returns int32, set in PatternSignedIntegerPatternIntegerCounter.
	Count() int32
	// SetCount assigns int32 provided by user to PatternSignedIntegerPatternIntegerCounter
	SetCount(value int32) PatternSignedIntegerPatternIntegerCounter
	// HasCount checks if Count has been set in PatternSignedIntegerPatternIntegerCounter
	HasCount() bool
}

// description is TBD
// Start returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) Start() int32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the int32 value in the PatternSignedIntegerPatternIntegerCounter object
func (obj *patternSignedIntegerPatternIntegerCounter) SetStart(value int32) PatternSignedIntegerPatternIntegerCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) Step() int32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the int32 value in the PatternSignedIntegerPatternIntegerCounter object
func (obj *patternSignedIntegerPatternIntegerCounter) SetStep(value int32) PatternSignedIntegerPatternIntegerCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) Count() int32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a int32
func (obj *patternSignedIntegerPatternIntegerCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the int32 value in the PatternSignedIntegerPatternIntegerCounter object
func (obj *patternSignedIntegerPatternIntegerCounter) SetCount(value int32) PatternSignedIntegerPatternIntegerCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternSignedIntegerPatternIntegerCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start < -128 || *obj.obj.Start > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-128 <= PatternSignedIntegerPatternIntegerCounter.Start <= 127 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step < -128 || *obj.obj.Step > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-128 <= PatternSignedIntegerPatternIntegerCounter.Step <= 127 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count < -128 || *obj.obj.Count > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-128 <= PatternSignedIntegerPatternIntegerCounter.Count <= 127 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternSignedIntegerPatternIntegerCounter) setDefault() {
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
