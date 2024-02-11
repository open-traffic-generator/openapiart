package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv6PatternIpv6 *****
type patternIpv6PatternIpv6 struct {
	validation
	obj             *openapi.PatternIpv6PatternIpv6
	marshaller      marshalPatternIpv6PatternIpv6
	unMarshaller    unMarshalPatternIpv6PatternIpv6
	incrementHolder PatternIpv6PatternIpv6Counter
	decrementHolder PatternIpv6PatternIpv6Counter
}

func NewPatternIpv6PatternIpv6() PatternIpv6PatternIpv6 {
	obj := patternIpv6PatternIpv6{obj: &openapi.PatternIpv6PatternIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv6PatternIpv6) msg() *openapi.PatternIpv6PatternIpv6 {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6) setMsg(msg *openapi.PatternIpv6PatternIpv6) PatternIpv6PatternIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv6PatternIpv6 struct {
	obj *patternIpv6PatternIpv6
}

type marshalPatternIpv6PatternIpv6 interface {
	// ToProto marshals PatternIpv6PatternIpv6 to protobuf object *openapi.PatternIpv6PatternIpv6
	ToProto() (*openapi.PatternIpv6PatternIpv6, error)
	// ToPbText marshals PatternIpv6PatternIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv6PatternIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv6PatternIpv6 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv6PatternIpv6 struct {
	obj *patternIpv6PatternIpv6
}

type unMarshalPatternIpv6PatternIpv6 interface {
	// FromProto unmarshals PatternIpv6PatternIpv6 from protobuf object *openapi.PatternIpv6PatternIpv6
	FromProto(msg *openapi.PatternIpv6PatternIpv6) (PatternIpv6PatternIpv6, error)
	// FromPbText unmarshals PatternIpv6PatternIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv6PatternIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv6PatternIpv6 from JSON text
	FromJson(value string) error
}

func (obj *patternIpv6PatternIpv6) Marshal() marshalPatternIpv6PatternIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv6PatternIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv6PatternIpv6) Unmarshal() unMarshalPatternIpv6PatternIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv6PatternIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv6PatternIpv6) ToProto() (*openapi.PatternIpv6PatternIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv6PatternIpv6) FromProto(msg *openapi.PatternIpv6PatternIpv6) (PatternIpv6PatternIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv6PatternIpv6) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalpatternIpv6PatternIpv6) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalpatternIpv6PatternIpv6) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv6PatternIpv6) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv6PatternIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv6PatternIpv6) Clone() (PatternIpv6PatternIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv6PatternIpv6()
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

func (obj *patternIpv6PatternIpv6) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv6PatternIpv6 is tBD
type PatternIpv6PatternIpv6 interface {
	Validation
	// msg marshals PatternIpv6PatternIpv6 to protobuf object *openapi.PatternIpv6PatternIpv6
	// and doesn't set defaults
	msg() *openapi.PatternIpv6PatternIpv6
	// setMsg unmarshals PatternIpv6PatternIpv6 from protobuf object *openapi.PatternIpv6PatternIpv6
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv6PatternIpv6) PatternIpv6PatternIpv6
	// provides marshal interface
	Marshal() marshalPatternIpv6PatternIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv6PatternIpv6
	// validate validates PatternIpv6PatternIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv6PatternIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv6PatternIpv6ChoiceEnum, set in PatternIpv6PatternIpv6
	Choice() PatternIpv6PatternIpv6ChoiceEnum
	// setChoice assigns PatternIpv6PatternIpv6ChoiceEnum provided by user to PatternIpv6PatternIpv6
	setChoice(value PatternIpv6PatternIpv6ChoiceEnum) PatternIpv6PatternIpv6
	// HasChoice checks if Choice has been set in PatternIpv6PatternIpv6
	HasChoice() bool
	// Value returns string, set in PatternIpv6PatternIpv6.
	Value() string
	// SetValue assigns string provided by user to PatternIpv6PatternIpv6
	SetValue(value string) PatternIpv6PatternIpv6
	// HasValue checks if Value has been set in PatternIpv6PatternIpv6
	HasValue() bool
	// Values returns []string, set in PatternIpv6PatternIpv6.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv6PatternIpv6
	SetValues(value []string) PatternIpv6PatternIpv6
	// Increment returns PatternIpv6PatternIpv6Counter, set in PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	Increment() PatternIpv6PatternIpv6Counter
	// SetIncrement assigns PatternIpv6PatternIpv6Counter provided by user to PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	SetIncrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6
	// HasIncrement checks if Increment has been set in PatternIpv6PatternIpv6
	HasIncrement() bool
	// Decrement returns PatternIpv6PatternIpv6Counter, set in PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	Decrement() PatternIpv6PatternIpv6Counter
	// SetDecrement assigns PatternIpv6PatternIpv6Counter provided by user to PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	SetDecrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6
	// HasDecrement checks if Decrement has been set in PatternIpv6PatternIpv6
	HasDecrement() bool
	setNil()
}

type PatternIpv6PatternIpv6ChoiceEnum string

// Enum of Choice on PatternIpv6PatternIpv6
var PatternIpv6PatternIpv6Choice = struct {
	VALUE     PatternIpv6PatternIpv6ChoiceEnum
	VALUES    PatternIpv6PatternIpv6ChoiceEnum
	INCREMENT PatternIpv6PatternIpv6ChoiceEnum
	DECREMENT PatternIpv6PatternIpv6ChoiceEnum
}{
	VALUE:     PatternIpv6PatternIpv6ChoiceEnum("value"),
	VALUES:    PatternIpv6PatternIpv6ChoiceEnum("values"),
	INCREMENT: PatternIpv6PatternIpv6ChoiceEnum("increment"),
	DECREMENT: PatternIpv6PatternIpv6ChoiceEnum("decrement"),
}

func (obj *patternIpv6PatternIpv6) Choice() PatternIpv6PatternIpv6ChoiceEnum {
	return PatternIpv6PatternIpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv6PatternIpv6) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv6PatternIpv6) setChoice(value PatternIpv6PatternIpv6ChoiceEnum) PatternIpv6PatternIpv6 {
	intValue, ok := openapi.PatternIpv6PatternIpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv6PatternIpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv6PatternIpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv6PatternIpv6Choice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv6PatternIpv6Choice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv6PatternIpv6Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv6PatternIpv6Counter().msg()
	}

	if value == PatternIpv6PatternIpv6Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv6PatternIpv6Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternIpv6) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIpv6PatternIpv6Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternIpv6) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetValue(value string) PatternIpv6PatternIpv6 {
	obj.setChoice(PatternIpv6PatternIpv6Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv6PatternIpv6) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetValues(value []string) PatternIpv6PatternIpv6 {
	obj.setChoice(PatternIpv6PatternIpv6Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) Increment() PatternIpv6PatternIpv6Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIpv6PatternIpv6Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv6PatternIpv6Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv6PatternIpv6Counter value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetIncrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6 {
	obj.setChoice(PatternIpv6PatternIpv6Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) Decrement() PatternIpv6PatternIpv6Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIpv6PatternIpv6Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv6PatternIpv6Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv6PatternIpv6Counter value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetDecrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6 {
	obj.setChoice(PatternIpv6PatternIpv6Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIpv6PatternIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv6PatternIpv6) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternIpv6PatternIpv6Choice.VALUE)

	}

}
