package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv4PatternObjectIpv4 *****
type patternIpv4PatternObjectIpv4 struct {
	validation
	obj             *openapi.PatternIpv4PatternObjectIpv4
	marshaller      marshalPatternIpv4PatternObjectIpv4
	unMarshaller    unMarshalPatternIpv4PatternObjectIpv4
	incrementHolder PatternIpv4PatternObjectIpv4Counter
	decrementHolder PatternIpv4PatternObjectIpv4Counter
}

func NewPatternIpv4PatternObjectIpv4() PatternIpv4PatternObjectIpv4 {
	obj := patternIpv4PatternObjectIpv4{obj: &openapi.PatternIpv4PatternObjectIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternObjectIpv4) msg() *openapi.PatternIpv4PatternObjectIpv4 {
	return obj.obj
}

func (obj *patternIpv4PatternObjectIpv4) setMsg(msg *openapi.PatternIpv4PatternObjectIpv4) PatternIpv4PatternObjectIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv4PatternObjectIpv4 struct {
	obj *patternIpv4PatternObjectIpv4
}

type marshalPatternIpv4PatternObjectIpv4 interface {
	// ToProto marshals PatternIpv4PatternObjectIpv4 to protobuf object *openapi.PatternIpv4PatternObjectIpv4
	ToProto() (*openapi.PatternIpv4PatternObjectIpv4, error)
	// ToPbText marshals PatternIpv4PatternObjectIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternObjectIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternObjectIpv4 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv4PatternObjectIpv4 struct {
	obj *patternIpv4PatternObjectIpv4
}

type unMarshalPatternIpv4PatternObjectIpv4 interface {
	// FromProto unmarshals PatternIpv4PatternObjectIpv4 from protobuf object *openapi.PatternIpv4PatternObjectIpv4
	FromProto(msg *openapi.PatternIpv4PatternObjectIpv4) (PatternIpv4PatternObjectIpv4, error)
	// FromPbText unmarshals PatternIpv4PatternObjectIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternObjectIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternObjectIpv4 from JSON text
	FromJson(value string) error
}

func (obj *patternIpv4PatternObjectIpv4) Marshal() marshalPatternIpv4PatternObjectIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv4PatternObjectIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv4PatternObjectIpv4) Unmarshal() unMarshalPatternIpv4PatternObjectIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv4PatternObjectIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv4PatternObjectIpv4) ToProto() (*openapi.PatternIpv4PatternObjectIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv4PatternObjectIpv4) FromProto(msg *openapi.PatternIpv4PatternObjectIpv4) (PatternIpv4PatternObjectIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv4PatternObjectIpv4) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4) FromPbText(value string) error {
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

func (m *marshalpatternIpv4PatternObjectIpv4) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4) FromYaml(value string) error {
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

func (m *marshalpatternIpv4PatternObjectIpv4) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv4PatternObjectIpv4) FromJson(value string) error {
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

func (obj *patternIpv4PatternObjectIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternObjectIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternObjectIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternObjectIpv4) Clone() (PatternIpv4PatternObjectIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternObjectIpv4()
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

func (obj *patternIpv4PatternObjectIpv4) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv4PatternObjectIpv4 is tBD
type PatternIpv4PatternObjectIpv4 interface {
	Validation
	// msg marshals PatternIpv4PatternObjectIpv4 to protobuf object *openapi.PatternIpv4PatternObjectIpv4
	// and doesn't set defaults
	msg() *openapi.PatternIpv4PatternObjectIpv4
	// setMsg unmarshals PatternIpv4PatternObjectIpv4 from protobuf object *openapi.PatternIpv4PatternObjectIpv4
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv4PatternObjectIpv4) PatternIpv4PatternObjectIpv4
	// provides marshal interface
	Marshal() marshalPatternIpv4PatternObjectIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv4PatternObjectIpv4
	// validate validates PatternIpv4PatternObjectIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternObjectIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv4PatternObjectIpv4ChoiceEnum, set in PatternIpv4PatternObjectIpv4
	Choice() PatternIpv4PatternObjectIpv4ChoiceEnum
	// setChoice assigns PatternIpv4PatternObjectIpv4ChoiceEnum provided by user to PatternIpv4PatternObjectIpv4
	setChoice(value PatternIpv4PatternObjectIpv4ChoiceEnum) PatternIpv4PatternObjectIpv4
	// HasChoice checks if Choice has been set in PatternIpv4PatternObjectIpv4
	HasChoice() bool
	// Value returns string, set in PatternIpv4PatternObjectIpv4.
	Value() string
	// SetValue assigns string provided by user to PatternIpv4PatternObjectIpv4
	SetValue(value string) PatternIpv4PatternObjectIpv4
	// HasValue checks if Value has been set in PatternIpv4PatternObjectIpv4
	HasValue() bool
	// Values returns []string, set in PatternIpv4PatternObjectIpv4.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv4PatternObjectIpv4
	SetValues(value []string) PatternIpv4PatternObjectIpv4
	// Increment returns PatternIpv4PatternObjectIpv4Counter, set in PatternIpv4PatternObjectIpv4.
	// PatternIpv4PatternObjectIpv4Counter is ipv4 counter pattern
	Increment() PatternIpv4PatternObjectIpv4Counter
	// SetIncrement assigns PatternIpv4PatternObjectIpv4Counter provided by user to PatternIpv4PatternObjectIpv4.
	// PatternIpv4PatternObjectIpv4Counter is ipv4 counter pattern
	SetIncrement(value PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4
	// HasIncrement checks if Increment has been set in PatternIpv4PatternObjectIpv4
	HasIncrement() bool
	// Decrement returns PatternIpv4PatternObjectIpv4Counter, set in PatternIpv4PatternObjectIpv4.
	// PatternIpv4PatternObjectIpv4Counter is ipv4 counter pattern
	Decrement() PatternIpv4PatternObjectIpv4Counter
	// SetDecrement assigns PatternIpv4PatternObjectIpv4Counter provided by user to PatternIpv4PatternObjectIpv4.
	// PatternIpv4PatternObjectIpv4Counter is ipv4 counter pattern
	SetDecrement(value PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4
	// HasDecrement checks if Decrement has been set in PatternIpv4PatternObjectIpv4
	HasDecrement() bool
	setNil()
}

type PatternIpv4PatternObjectIpv4ChoiceEnum string

// Enum of Choice on PatternIpv4PatternObjectIpv4
var PatternIpv4PatternObjectIpv4Choice = struct {
	VALUE     PatternIpv4PatternObjectIpv4ChoiceEnum
	VALUES    PatternIpv4PatternObjectIpv4ChoiceEnum
	INCREMENT PatternIpv4PatternObjectIpv4ChoiceEnum
	DECREMENT PatternIpv4PatternObjectIpv4ChoiceEnum
}{
	VALUE:     PatternIpv4PatternObjectIpv4ChoiceEnum("value"),
	VALUES:    PatternIpv4PatternObjectIpv4ChoiceEnum("values"),
	INCREMENT: PatternIpv4PatternObjectIpv4ChoiceEnum("increment"),
	DECREMENT: PatternIpv4PatternObjectIpv4ChoiceEnum("decrement"),
}

func (obj *patternIpv4PatternObjectIpv4) Choice() PatternIpv4PatternObjectIpv4ChoiceEnum {
	return PatternIpv4PatternObjectIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv4PatternObjectIpv4) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv4PatternObjectIpv4) setChoice(value PatternIpv4PatternObjectIpv4ChoiceEnum) PatternIpv4PatternObjectIpv4 {
	intValue, ok := openapi.PatternIpv4PatternObjectIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv4PatternObjectIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv4PatternObjectIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv4PatternObjectIpv4Choice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv4PatternObjectIpv4Choice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv4PatternObjectIpv4Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv4PatternObjectIpv4Counter().msg()
	}

	if value == PatternIpv4PatternObjectIpv4Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv4PatternObjectIpv4Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternObjectIpv4) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIpv4PatternObjectIpv4Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternObjectIpv4) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv4PatternObjectIpv4 object
func (obj *patternIpv4PatternObjectIpv4) SetValue(value string) PatternIpv4PatternObjectIpv4 {
	obj.setChoice(PatternIpv4PatternObjectIpv4Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv4PatternObjectIpv4) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv4PatternObjectIpv4 object
func (obj *patternIpv4PatternObjectIpv4) SetValues(value []string) PatternIpv4PatternObjectIpv4 {
	obj.setChoice(PatternIpv4PatternObjectIpv4Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv4PatternObjectIpv4Counter
func (obj *patternIpv4PatternObjectIpv4) Increment() PatternIpv4PatternObjectIpv4Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIpv4PatternObjectIpv4Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv4PatternObjectIpv4Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv4PatternObjectIpv4Counter
func (obj *patternIpv4PatternObjectIpv4) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv4PatternObjectIpv4Counter value in the PatternIpv4PatternObjectIpv4 object
func (obj *patternIpv4PatternObjectIpv4) SetIncrement(value PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4 {
	obj.setChoice(PatternIpv4PatternObjectIpv4Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv4PatternObjectIpv4Counter
func (obj *patternIpv4PatternObjectIpv4) Decrement() PatternIpv4PatternObjectIpv4Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIpv4PatternObjectIpv4Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv4PatternObjectIpv4Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv4PatternObjectIpv4Counter
func (obj *patternIpv4PatternObjectIpv4) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv4PatternObjectIpv4Counter value in the PatternIpv4PatternObjectIpv4 object
func (obj *patternIpv4PatternObjectIpv4) SetDecrement(value PatternIpv4PatternObjectIpv4Counter) PatternIpv4PatternObjectIpv4 {
	obj.setChoice(PatternIpv4PatternObjectIpv4Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIpv4PatternObjectIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternObjectIpv4.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternObjectIpv4.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv4PatternObjectIpv4) setDefault() {
	var choices_set int = 0
	var choice PatternIpv4PatternObjectIpv4ChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternIpv4PatternObjectIpv4Choice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternIpv4PatternObjectIpv4Choice.VALUES
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternIpv4PatternObjectIpv4Choice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternIpv4PatternObjectIpv4Choice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternIpv4PatternObjectIpv4Choice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternIpv4PatternObjectIpv4")
			}
		} else {
			intVal := openapi.PatternIpv4PatternObjectIpv4_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternIpv4PatternObjectIpv4_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
