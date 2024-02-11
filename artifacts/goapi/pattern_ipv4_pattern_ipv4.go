package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIpv4PatternIpv4 *****
type patternIpv4PatternIpv4 struct {
	validation
	obj             *openapi.PatternIpv4PatternIpv4
	marshaller      marshalPatternIpv4PatternIpv4
	unMarshaller    unMarshalPatternIpv4PatternIpv4
	incrementHolder PatternIpv4PatternIpv4Counter
	decrementHolder PatternIpv4PatternIpv4Counter
}

func NewPatternIpv4PatternIpv4() PatternIpv4PatternIpv4 {
	obj := patternIpv4PatternIpv4{obj: &openapi.PatternIpv4PatternIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternIpv4) msg() *openapi.PatternIpv4PatternIpv4 {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4) setMsg(msg *openapi.PatternIpv4PatternIpv4) PatternIpv4PatternIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIpv4PatternIpv4 struct {
	obj *patternIpv4PatternIpv4
}

type marshalPatternIpv4PatternIpv4 interface {
	// ToProto marshals PatternIpv4PatternIpv4 to protobuf object *openapi.PatternIpv4PatternIpv4
	ToProto() (*openapi.PatternIpv4PatternIpv4, error)
	// ToPbText marshals PatternIpv4PatternIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternIpv4 to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIpv4PatternIpv4 struct {
	obj *patternIpv4PatternIpv4
}

type unMarshalPatternIpv4PatternIpv4 interface {
	// FromProto unmarshals PatternIpv4PatternIpv4 from protobuf object *openapi.PatternIpv4PatternIpv4
	FromProto(msg *openapi.PatternIpv4PatternIpv4) (PatternIpv4PatternIpv4, error)
	// FromPbText unmarshals PatternIpv4PatternIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternIpv4 from JSON text
	FromJson(value string) error
}

func (obj *patternIpv4PatternIpv4) Marshal() marshalPatternIpv4PatternIpv4 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIpv4PatternIpv4{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIpv4PatternIpv4) Unmarshal() unMarshalPatternIpv4PatternIpv4 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIpv4PatternIpv4{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIpv4PatternIpv4) ToProto() (*openapi.PatternIpv4PatternIpv4, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIpv4PatternIpv4) FromProto(msg *openapi.PatternIpv4PatternIpv4) (PatternIpv4PatternIpv4, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIpv4PatternIpv4) ToPbText() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4) FromPbText(value string) error {
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

func (m *marshalpatternIpv4PatternIpv4) ToYaml() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4) FromYaml(value string) error {
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

func (m *marshalpatternIpv4PatternIpv4) ToJson() (string, error) {
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

func (m *unMarshalpatternIpv4PatternIpv4) FromJson(value string) error {
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

func (obj *patternIpv4PatternIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternIpv4) Clone() (PatternIpv4PatternIpv4, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternIpv4()
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

func (obj *patternIpv4PatternIpv4) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv4PatternIpv4 is tBD
type PatternIpv4PatternIpv4 interface {
	Validation
	// msg marshals PatternIpv4PatternIpv4 to protobuf object *openapi.PatternIpv4PatternIpv4
	// and doesn't set defaults
	msg() *openapi.PatternIpv4PatternIpv4
	// setMsg unmarshals PatternIpv4PatternIpv4 from protobuf object *openapi.PatternIpv4PatternIpv4
	// and doesn't set defaults
	setMsg(*openapi.PatternIpv4PatternIpv4) PatternIpv4PatternIpv4
	// provides marshal interface
	Marshal() marshalPatternIpv4PatternIpv4
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIpv4PatternIpv4
	// validate validates PatternIpv4PatternIpv4
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv4PatternIpv4ChoiceEnum, set in PatternIpv4PatternIpv4
	Choice() PatternIpv4PatternIpv4ChoiceEnum
	// setChoice assigns PatternIpv4PatternIpv4ChoiceEnum provided by user to PatternIpv4PatternIpv4
	setChoice(value PatternIpv4PatternIpv4ChoiceEnum) PatternIpv4PatternIpv4
	// HasChoice checks if Choice has been set in PatternIpv4PatternIpv4
	HasChoice() bool
	// Value returns string, set in PatternIpv4PatternIpv4.
	Value() string
	// SetValue assigns string provided by user to PatternIpv4PatternIpv4
	SetValue(value string) PatternIpv4PatternIpv4
	// HasValue checks if Value has been set in PatternIpv4PatternIpv4
	HasValue() bool
	// Values returns []string, set in PatternIpv4PatternIpv4.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv4PatternIpv4
	SetValues(value []string) PatternIpv4PatternIpv4
	// Increment returns PatternIpv4PatternIpv4Counter, set in PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	Increment() PatternIpv4PatternIpv4Counter
	// SetIncrement assigns PatternIpv4PatternIpv4Counter provided by user to PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	SetIncrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4
	// HasIncrement checks if Increment has been set in PatternIpv4PatternIpv4
	HasIncrement() bool
	// Decrement returns PatternIpv4PatternIpv4Counter, set in PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	Decrement() PatternIpv4PatternIpv4Counter
	// SetDecrement assigns PatternIpv4PatternIpv4Counter provided by user to PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	SetDecrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4
	// HasDecrement checks if Decrement has been set in PatternIpv4PatternIpv4
	HasDecrement() bool
	setNil()
}

type PatternIpv4PatternIpv4ChoiceEnum string

// Enum of Choice on PatternIpv4PatternIpv4
var PatternIpv4PatternIpv4Choice = struct {
	VALUE     PatternIpv4PatternIpv4ChoiceEnum
	VALUES    PatternIpv4PatternIpv4ChoiceEnum
	INCREMENT PatternIpv4PatternIpv4ChoiceEnum
	DECREMENT PatternIpv4PatternIpv4ChoiceEnum
}{
	VALUE:     PatternIpv4PatternIpv4ChoiceEnum("value"),
	VALUES:    PatternIpv4PatternIpv4ChoiceEnum("values"),
	INCREMENT: PatternIpv4PatternIpv4ChoiceEnum("increment"),
	DECREMENT: PatternIpv4PatternIpv4ChoiceEnum("decrement"),
}

func (obj *patternIpv4PatternIpv4) Choice() PatternIpv4PatternIpv4ChoiceEnum {
	return PatternIpv4PatternIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv4PatternIpv4) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv4PatternIpv4) setChoice(value PatternIpv4PatternIpv4ChoiceEnum) PatternIpv4PatternIpv4 {
	intValue, ok := openapi.PatternIpv4PatternIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv4PatternIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv4PatternIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv4PatternIpv4Choice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv4PatternIpv4Choice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv4PatternIpv4Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv4PatternIpv4Counter().msg()
	}

	if value == PatternIpv4PatternIpv4Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv4PatternIpv4Counter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternIpv4) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIpv4PatternIpv4Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternIpv4) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetValue(value string) PatternIpv4PatternIpv4 {
	obj.setChoice(PatternIpv4PatternIpv4Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv4PatternIpv4) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetValues(value []string) PatternIpv4PatternIpv4 {
	obj.setChoice(PatternIpv4PatternIpv4Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) Increment() PatternIpv4PatternIpv4Counter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIpv4PatternIpv4Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv4PatternIpv4Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv4PatternIpv4Counter value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetIncrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4 {
	obj.setChoice(PatternIpv4PatternIpv4Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) Decrement() PatternIpv4PatternIpv4Counter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIpv4PatternIpv4Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv4PatternIpv4Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv4PatternIpv4Counter value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetDecrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4 {
	obj.setChoice(PatternIpv4PatternIpv4Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIpv4PatternIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv4PatternIpv4) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternIpv4PatternIpv4Choice.VALUE)

	}

}
