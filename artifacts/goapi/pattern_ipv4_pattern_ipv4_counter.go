package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv4PatternIpv4Counter *****
type patternIpv4PatternIpv4Counter struct {
	validation
	obj          *openapi.PatternIpv4PatternIpv4Counter
	marshaller   marshalPatternIpv4PatternIpv4Counter
	unMarshaller unMarshalPatternIpv4PatternIpv4Counter
}

func NewPatternIpv4PatternIpv4Counter() PatternIpv4PatternIpv4Counter {
	obj := patternIpv4PatternIpv4Counter{obj: &openapi.PatternIpv4PatternIpv4Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternIpv4Counter) msg() *openapi.PatternIpv4PatternIpv4Counter {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4Counter) setMsg(msg *openapi.PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv4PatternIpv4Counter struct {
	obj *patternIpv4PatternIpv4Counter
}

type marshalPatternIpv4PatternIpv4Counter interface {
	// ToProto marshals PatternIpv4PatternIpv4Counter to protobuf object *openapi.PatternIpv4PatternIpv4Counter
	ToProto() (*openapi.PatternIpv4PatternIpv4Counter, error)
	// ToPbText marshals PatternIpv4PatternIpv4Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternIpv4Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternIpv4Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv4PatternIpv4Counter struct {
	obj *patternIpv4PatternIpv4Counter
}

type unMarshalPatternIpv4PatternIpv4Counter interface {
	// FromProto unmarshals PatternIpv4PatternIpv4Counter from protobuf object *openapi.PatternIpv4PatternIpv4Counter
	FromProto(msg *openapi.PatternIpv4PatternIpv4Counter) (PatternIpv4PatternIpv4Counter, error)
	// FromPbText unmarshals PatternIpv4PatternIpv4Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternIpv4Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternIpv4Counter from JSON text
	FromJson(value string) error
}

func (obj *patternIpv4PatternIpv4Counter) Marshal() marshalPatternIpv4PatternIpv4Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv4PatternIpv4Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv4PatternIpv4Counter) Unmarshal() unMarshalPatternIpv4PatternIpv4Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv4PatternIpv4Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv4PatternIpv4Counter) ToProto() (*openapi.PatternIpv4PatternIpv4Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv4PatternIpv4Counter) FromProto(msg *openapi.PatternIpv4PatternIpv4Counter) (PatternIpv4PatternIpv4Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv4PatternIpv4Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4Counter) FromPbText(value string) error {
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

func (m *marshalpatternIpv4PatternIpv4Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4Counter) FromYaml(value string) error {
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

func (m *marshalpatternIpv4PatternIpv4Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4Counter) FromJson(value string) error {
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

func (obj *patternIpv4PatternIpv4Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternIpv4Counter) Clone() (PatternIpv4PatternIpv4Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternIpv4Counter()
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

// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
type PatternIpv4PatternIpv4Counter interface {
	Validation
	// msg marshals PatternIpv4PatternIpv4Counter to protobuf object *openapi.PatternIpv4PatternIpv4Counter
	// and doesn't set defaults
	msg() *openapi.PatternIpv4PatternIpv4Counter
	// setMsg unmarshals PatternIpv4PatternIpv4Counter from protobuf object *openapi.PatternIpv4PatternIpv4Counter
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4Counter
	// provides marshal interface
	Marshal() marshalPatternIpv4PatternIpv4Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv4PatternIpv4Counter
	// validate validates PatternIpv4PatternIpv4Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternIpv4Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternIpv4PatternIpv4Counter.
	Start() string
	// SetStart assigns string provided by user to PatternIpv4PatternIpv4Counter
	SetStart(value string) PatternIpv4PatternIpv4Counter
	// HasStart checks if Start has been set in PatternIpv4PatternIpv4Counter
	HasStart() bool
	// Step returns string, set in PatternIpv4PatternIpv4Counter.
	Step() string
	// SetStep assigns string provided by user to PatternIpv4PatternIpv4Counter
	SetStep(value string) PatternIpv4PatternIpv4Counter
	// HasStep checks if Step has been set in PatternIpv4PatternIpv4Counter
	HasStep() bool
	// Count returns uint32, set in PatternIpv4PatternIpv4Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIpv4PatternIpv4Counter
	SetCount(value uint32) PatternIpv4PatternIpv4Counter
	// HasCount checks if Count has been set in PatternIpv4PatternIpv4Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternIpv4Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternIpv4Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetStart(value string) PatternIpv4PatternIpv4Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternIpv4Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternIpv4Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetStep(value string) PatternIpv4PatternIpv4Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternIpv4Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternIpv4Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetCount(value uint32) PatternIpv4PatternIpv4Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIpv4PatternIpv4Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Step"))
		}

	}

}

func (obj *patternIpv4PatternIpv4Counter) setDefault() {
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
