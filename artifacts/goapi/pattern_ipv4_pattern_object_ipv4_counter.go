package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv4PatternObjectIpv4Counter *****
type patternIpv4PatternObjectIpv4Counter struct {
	validation
	obj          *openapi.PatternIpv4PatternObjectIpv4Counter
	marshaller   marshalPatternIpv4PatternObjectIpv4Counter
	unMarshaller unMarshalPatternIpv4PatternObjectIpv4Counter
}

func NewPatternIpv4PatternObjectIpv4Counter() PatternIpv4PatternObjectIpv4Counter {
	obj := patternIpv4PatternObjectIpv4Counter{obj: &openapi.PatternIpv4PatternObjectIpv4Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternObjectIpv4Counter) msg() *openapi.PatternIpv4PatternObjectIpv4Counter {
	return obj.obj
}

func (obj *patternIpv4PatternObjectIpv4Counter) setMsg(msg *openapi.PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv4PatternObjectIpv4Counter struct {
	obj *patternIpv4PatternObjectIpv4Counter
}

type marshalPatternIpv4PatternObjectIpv4Counter interface {
	// ToProto marshals PatternIpv4PatternObjectIpv4Counter to protobuf object *openapi.PatternIpv4PatternObjectIpv4Counter
	ToProto() (*openapi.PatternIpv4PatternObjectIpv4Counter, error)
	// ToPbText marshals PatternIpv4PatternObjectIpv4Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternObjectIpv4Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternObjectIpv4Counter to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv4PatternObjectIpv4Counter struct {
	obj *patternIpv4PatternObjectIpv4Counter
}

type unMarshalPatternIpv4PatternObjectIpv4Counter interface {
	// FromProto unmarshals PatternIpv4PatternObjectIpv4Counter from protobuf object *openapi.PatternIpv4PatternObjectIpv4Counter
	FromProto(msg *openapi.PatternIpv4PatternObjectIpv4Counter) (PatternIpv4PatternObjectIpv4Counter, error)
	// FromPbText unmarshals PatternIpv4PatternObjectIpv4Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternObjectIpv4Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternObjectIpv4Counter from JSON text
	FromJson(value string) error
}

func (obj *patternIpv4PatternObjectIpv4Counter) Marshal() marshalPatternIpv4PatternObjectIpv4Counter {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv4PatternObjectIpv4Counter{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv4PatternObjectIpv4Counter) Unmarshal() unMarshalPatternIpv4PatternObjectIpv4Counter {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv4PatternObjectIpv4Counter{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv4PatternObjectIpv4Counter) ToProto() (*openapi.PatternIpv4PatternObjectIpv4Counter, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv4PatternObjectIpv4Counter) FromProto(msg *openapi.PatternIpv4PatternObjectIpv4Counter) (PatternIpv4PatternObjectIpv4Counter, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv4PatternObjectIpv4Counter) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4Counter) FromPbText(value string) error {
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

func (m *marshalpatternIpv4PatternObjectIpv4Counter) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4Counter) FromYaml(value string) error {
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

func (m *marshalpatternIpv4PatternObjectIpv4Counter) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4Counter) FromJson(value string) error {
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

func (obj *patternIpv4PatternObjectIpv4Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternObjectIpv4Counter) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternObjectIpv4Counter) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternObjectIpv4Counter) Clone() (PatternIpv4PatternObjectIpv4Counter, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternObjectIpv4Counter()
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

// PatternIpv4PatternObjectIpv4Counter is ipv4 counter pattern
type PatternIpv4PatternObjectIpv4Counter interface {
	Validation
	// msg marshals PatternIpv4PatternObjectIpv4Counter to protobuf object *openapi.PatternIpv4PatternObjectIpv4Counter
	// and doesn't set defaults
	msg() *openapi.PatternIpv4PatternObjectIpv4Counter
	// setMsg unmarshals PatternIpv4PatternObjectIpv4Counter from protobuf object *openapi.PatternIpv4PatternObjectIpv4Counter
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4Counter
	// provides marshal interface
	Marshal() marshalPatternIpv4PatternObjectIpv4Counter
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv4PatternObjectIpv4Counter
	// validate validates PatternIpv4PatternObjectIpv4Counter
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternObjectIpv4Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternIpv4PatternObjectIpv4Counter.
	Start() string
	// SetStart assigns string provided by user to PatternIpv4PatternObjectIpv4Counter
	SetStart(value string) PatternIpv4PatternObjectIpv4Counter
	// HasStart checks if Start has been set in PatternIpv4PatternObjectIpv4Counter
	HasStart() bool
	// Step returns string, set in PatternIpv4PatternObjectIpv4Counter.
	Step() string
	// SetStep assigns string provided by user to PatternIpv4PatternObjectIpv4Counter
	SetStep(value string) PatternIpv4PatternObjectIpv4Counter
	// HasStep checks if Step has been set in PatternIpv4PatternObjectIpv4Counter
	HasStep() bool
	// Count returns uint32, set in PatternIpv4PatternObjectIpv4Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIpv4PatternObjectIpv4Counter
	SetCount(value uint32) PatternIpv4PatternObjectIpv4Counter
	// HasCount checks if Count has been set in PatternIpv4PatternObjectIpv4Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternObjectIpv4Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternObjectIpv4Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternIpv4PatternObjectIpv4Counter object
func (obj *patternIpv4PatternObjectIpv4Counter) SetStart(value string) PatternIpv4PatternObjectIpv4Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternObjectIpv4Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternObjectIpv4Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternIpv4PatternObjectIpv4Counter object
func (obj *patternIpv4PatternObjectIpv4Counter) SetStep(value string) PatternIpv4PatternObjectIpv4Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternObjectIpv4Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternObjectIpv4Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIpv4PatternObjectIpv4Counter object
func (obj *patternIpv4PatternObjectIpv4Counter) SetCount(value uint32) PatternIpv4PatternObjectIpv4Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIpv4PatternObjectIpv4Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternObjectIpv4Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternObjectIpv4Counter.Step"))
		}

	}

}

func (obj *patternIpv4PatternObjectIpv4Counter) setDefault() {
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
