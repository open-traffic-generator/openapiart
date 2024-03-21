package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv6PatternObjectIpv6 *****
type patternIpv6PatternObjectIpv6 struct {
	validation
	obj             *openapi.PatternIpv6PatternObjectIpv6
	marshaller      marshalPatternIpv6PatternObjectIpv6
	unMarshaller    unMarshalPatternIpv6PatternObjectIpv6
	incrementHolder PatternIpv6PatternObjectIpv6Counter
	decrementHolder PatternIpv6PatternObjectIpv6Counter
}

func NewPatternIpv6PatternObjectIpv6() PatternIpv6PatternObjectIpv6 {
	obj := patternIpv6PatternObjectIpv6{obj: &openapi.PatternIpv6PatternObjectIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv6PatternObjectIpv6) msg() *openapi.PatternIpv6PatternObjectIpv6 {
	return obj.obj
}

func (obj *patternIpv6PatternObjectIpv6) setMsg(msg *openapi.PatternIpv6PatternObjectIpv6) PatternIpv6PatternObjectIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv6PatternObjectIpv6 struct {
	obj *patternIpv6PatternObjectIpv6
}

type marshalPatternIpv6PatternObjectIpv6 interface {
	// ToProto marshals PatternIpv6PatternObjectIpv6 to protobuf object *openapi.PatternIpv6PatternObjectIpv6
	ToProto() (*openapi.PatternIpv6PatternObjectIpv6, error)
	// ToPbText marshals PatternIpv6PatternObjectIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv6PatternObjectIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv6PatternObjectIpv6 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv6PatternObjectIpv6 struct {
	obj *patternIpv6PatternObjectIpv6
}

type unMarshalPatternIpv6PatternObjectIpv6 interface {
	// FromProto unmarshals PatternIpv6PatternObjectIpv6 from protobuf object *openapi.PatternIpv6PatternObjectIpv6
	FromProto(msg *openapi.PatternIpv6PatternObjectIpv6) (PatternIpv6PatternObjectIpv6, error)
	// FromPbText unmarshals PatternIpv6PatternObjectIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv6PatternObjectIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv6PatternObjectIpv6 from JSON text
	FromJson(value string) error
}

func (obj *patternIpv6PatternObjectIpv6) Marshal() marshalPatternIpv6PatternObjectIpv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv6PatternObjectIpv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv6PatternObjectIpv6) Unmarshal() unMarshalPatternIpv6PatternObjectIpv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv6PatternObjectIpv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv6PatternObjectIpv6) ToProto() (*openapi.PatternIpv6PatternObjectIpv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv6PatternObjectIpv6) FromProto(msg *openapi.PatternIpv6PatternObjectIpv6) (PatternIpv6PatternObjectIpv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv6PatternObjectIpv6) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv6PatternObjectIpv6) FromPbText(value string) error {
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

func (m *marshalpatternIpv6PatternObjectIpv6) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv6PatternObjectIpv6) FromYaml(value string) error {
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

func (m *marshalpatternIpv6PatternObjectIpv6) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv6PatternObjectIpv6) FromJson(value string) error {
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

func (obj *patternIpv6PatternObjectIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv6PatternObjectIpv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv6PatternObjectIpv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv6PatternObjectIpv6) Clone() (PatternIpv6PatternObjectIpv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv6PatternObjectIpv6()
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

func (obj *patternIpv6PatternObjectIpv6) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv6PatternObjectIpv6 is tBD
type PatternIpv6PatternObjectIpv6 interface {
	Validation
	// msg marshals PatternIpv6PatternObjectIpv6 to protobuf object *openapi.PatternIpv6PatternObjectIpv6
	// and doesn't set defaults
	msg() *openapi.PatternIpv6PatternObjectIpv6
	// setMsg unmarshals PatternIpv6PatternObjectIpv6 from protobuf object *openapi.PatternIpv6PatternObjectIpv6
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv6PatternObjectIpv6) PatternIpv6PatternObjectIpv6
	// provides marshal interface
	Marshal() marshalPatternIpv6PatternObjectIpv6
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv6PatternObjectIpv6
	// validate validates PatternIpv6PatternObjectIpv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv6PatternObjectIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv6PatternObjectIpv6ChoiceEnum, set in PatternIpv6PatternObjectIpv6
	Choice() PatternIpv6PatternObjectIpv6ChoiceEnum
	// setChoice assigns PatternIpv6PatternObjectIpv6ChoiceEnum provided by user to PatternIpv6PatternObjectIpv6
	setChoice(value PatternIpv6PatternObjectIpv6ChoiceEnum) PatternIpv6PatternObjectIpv6
	// HasChoice checks if Choice has been set in PatternIpv6PatternObjectIpv6
	HasChoice() bool
	// Value returns string, set in PatternIpv6PatternObjectIpv6.
	Value() string
	// SetValue assigns string provided by user to PatternIpv6PatternObjectIpv6
	SetValue(value string) PatternIpv6PatternObjectIpv6
	// HasValue checks if Value has been set in PatternIpv6PatternObjectIpv6
	HasValue() bool
	// Values returns []string, set in PatternIpv6PatternObjectIpv6.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv6PatternObjectIpv6
	SetValues(value []string) PatternIpv6PatternObjectIpv6
	// Increment returns PatternIpv6PatternObjectIpv6Counter, set in PatternIpv6PatternObjectIpv6.
	// PatternIpv6PatternObjectIpv6Counter is ipv6 counter pattern
	Increment() PatternIpv6PatternObjectIpv6Counter
	// SetIncrement assigns PatternIpv6PatternObjectIpv6Counter provided by user to PatternIpv6PatternObjectIpv6.
	// PatternIpv6PatternObjectIpv6Counter is ipv6 counter pattern
	SetIncrement(value PatternIpv6PatternObjectIpv6Counter) PatternIpv6PatternObjectIpv6
	// HasIncrement checks if Increment has been set in PatternIpv6PatternObjectIpv6
	HasIncrement() bool
	// Decrement returns PatternIpv6PatternObjectIpv6Counter, set in PatternIpv6PatternObjectIpv6.
	// PatternIpv6PatternObjectIpv6Counter is ipv6 counter pattern
	Decrement() PatternIpv6PatternObjectIpv6Counter
	// SetDecrement assigns PatternIpv6PatternObjectIpv6Counter provided by user to PatternIpv6PatternObjectIpv6.
	// PatternIpv6PatternObjectIpv6Counter is ipv6 counter pattern
	SetDecrement(value PatternIpv6PatternObjectIpv6Counter) PatternIpv6PatternObjectIpv6
	// HasDecrement checks if Decrement has been set in PatternIpv6PatternObjectIpv6
	HasDecrement() bool
	setNil()
}

type PatternIpv6PatternObjectIpv6ChoiceEnum string

// Enum of Choice on PatternIpv6PatternObjectIpv6
var PatternIpv6PatternObjectIpv6Choice = struct {
	VALUE     PatternIpv6PatternObjectIpv6ChoiceEnum
	VALUES    PatternIpv6PatternObjectIpv6ChoiceEnum
	INCREMENT PatternIpv6PatternObjectIpv6ChoiceEnum
	DECREMENT PatternIpv6PatternObjectIpv6ChoiceEnum
}{
	VALUE:     PatternIpv6PatternObjectIpv6ChoiceEnum("value"),
	VALUES:    PatternIpv6PatternObjectIpv6ChoiceEnum("values"),
	INCREMENT: PatternIpv6PatternObjectIpv6ChoiceEnum("increment"),
	DECREMENT: PatternIpv6PatternObjectIpv6ChoiceEnum("decrement"),
}

func (obj *patternIpv6PatternObjectIpv6) Choice() PatternIpv6PatternObjectIpv6ChoiceEnum {
	return PatternIpv6PatternObjectIpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv6PatternObjectIpv6) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv6PatternObjectIpv6) setChoice(value PatternIpv6PatternObjectIpv6ChoiceEnum) PatternIpv6PatternObjectIpv6 {
	intValue, ok := openapi.PatternIpv6PatternObjectIpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv6PatternObjectIpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv6PatternObjectIpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv6PatternObjectIpv6Choice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv6PatternObjectIpv6Choice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv6PatternObjectIpv6Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv6PatternObjectIpv6Counter().msg()
	}

	if value == PatternIpv6PatternObjectIpv6Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv6PatternObjectIpv6Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternObjectIpv6) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIpv6PatternObjectIpv6Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternObjectIpv6) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv6PatternObjectIpv6 object
func (obj *patternIpv6PatternObjectIpv6) SetValue(value string) PatternIpv6PatternObjectIpv6 {
	obj.setChoice(PatternIpv6PatternObjectIpv6Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv6PatternObjectIpv6) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv6PatternObjectIpv6 object
func (obj *patternIpv6PatternObjectIpv6) SetValues(value []string) PatternIpv6PatternObjectIpv6 {
	obj.setChoice(PatternIpv6PatternObjectIpv6Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv6PatternObjectIpv6Counter
func (obj *patternIpv6PatternObjectIpv6) Increment() PatternIpv6PatternObjectIpv6Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIpv6PatternObjectIpv6Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv6PatternObjectIpv6Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv6PatternObjectIpv6Counter
func (obj *patternIpv6PatternObjectIpv6) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv6PatternObjectIpv6Counter value in the PatternIpv6PatternObjectIpv6 object
func (obj *patternIpv6PatternObjectIpv6) SetIncrement(value PatternIpv6PatternObjectIpv6Counter) PatternIpv6PatternObjectIpv6 {
	obj.setChoice(PatternIpv6PatternObjectIpv6Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv6PatternObjectIpv6Counter
func (obj *patternIpv6PatternObjectIpv6) Decrement() PatternIpv6PatternObjectIpv6Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIpv6PatternObjectIpv6Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv6PatternObjectIpv6Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv6PatternObjectIpv6Counter
func (obj *patternIpv6PatternObjectIpv6) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv6PatternObjectIpv6Counter value in the PatternIpv6PatternObjectIpv6 object
func (obj *patternIpv6PatternObjectIpv6) SetDecrement(value PatternIpv6PatternObjectIpv6Counter) PatternIpv6PatternObjectIpv6 {
	obj.setChoice(PatternIpv6PatternObjectIpv6Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIpv6PatternObjectIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternObjectIpv6.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternObjectIpv6.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv6PatternObjectIpv6) setDefault() {
	var choices_set int = 0
	var choice PatternIpv6PatternObjectIpv6ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternIpv6PatternObjectIpv6Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternIpv6PatternObjectIpv6Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternIpv6PatternObjectIpv6Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternIpv6PatternObjectIpv6Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternIpv6PatternObjectIpv6Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternIpv6PatternObjectIpv6")
			}
		} else {
			intVal := openapi.PatternIpv6PatternObjectIpv6_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternIpv6PatternObjectIpv6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
