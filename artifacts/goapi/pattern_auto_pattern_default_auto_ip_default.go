package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternAutoPatternDefaultAutoIpDefault *****
type patternAutoPatternDefaultAutoIpDefault struct {
	validation
	obj             *openapi.PatternAutoPatternDefaultAutoIpDefault
	marshaller      marshalPatternAutoPatternDefaultAutoIpDefault
	unMarshaller    unMarshalPatternAutoPatternDefaultAutoIpDefault
	autoHolder      AutoIpDefault
	incrementHolder PatternAutoPatternDefaultAutoIpDefaultCounter
	decrementHolder PatternAutoPatternDefaultAutoIpDefaultCounter
}

func NewPatternAutoPatternDefaultAutoIpDefault() PatternAutoPatternDefaultAutoIpDefault {
	obj := patternAutoPatternDefaultAutoIpDefault{obj: &openapi.PatternAutoPatternDefaultAutoIpDefault{}}
	obj.setDefault()
	return &obj
}

func (obj *patternAutoPatternDefaultAutoIpDefault) msg() *openapi.PatternAutoPatternDefaultAutoIpDefault {
	return obj.obj
}

func (obj *patternAutoPatternDefaultAutoIpDefault) setMsg(msg *openapi.PatternAutoPatternDefaultAutoIpDefault) PatternAutoPatternDefaultAutoIpDefault {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternAutoPatternDefaultAutoIpDefault struct {
	obj *patternAutoPatternDefaultAutoIpDefault
}

type marshalPatternAutoPatternDefaultAutoIpDefault interface {
	// ToProto marshals PatternAutoPatternDefaultAutoIpDefault to protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefault
	ToProto() (*openapi.PatternAutoPatternDefaultAutoIpDefault, error)
	// ToPbText marshals PatternAutoPatternDefaultAutoIpDefault to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternAutoPatternDefaultAutoIpDefault to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternAutoPatternDefaultAutoIpDefault to JSON text
	ToJson() (string, error)
}

type unMarshalpatternAutoPatternDefaultAutoIpDefault struct {
	obj *patternAutoPatternDefaultAutoIpDefault
}

type unMarshalPatternAutoPatternDefaultAutoIpDefault interface {
	// FromProto unmarshals PatternAutoPatternDefaultAutoIpDefault from protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefault
	FromProto(msg *openapi.PatternAutoPatternDefaultAutoIpDefault) (PatternAutoPatternDefaultAutoIpDefault, error)
	// FromPbText unmarshals PatternAutoPatternDefaultAutoIpDefault from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternAutoPatternDefaultAutoIpDefault from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternAutoPatternDefaultAutoIpDefault from JSON text
	FromJson(value string) error
}

func (obj *patternAutoPatternDefaultAutoIpDefault) Marshal() marshalPatternAutoPatternDefaultAutoIpDefault {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternAutoPatternDefaultAutoIpDefault{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternAutoPatternDefaultAutoIpDefault) Unmarshal() unMarshalPatternAutoPatternDefaultAutoIpDefault {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternAutoPatternDefaultAutoIpDefault{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternAutoPatternDefaultAutoIpDefault) ToProto() (*openapi.PatternAutoPatternDefaultAutoIpDefault, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefault) FromProto(msg *openapi.PatternAutoPatternDefaultAutoIpDefault) (PatternAutoPatternDefaultAutoIpDefault, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternAutoPatternDefaultAutoIpDefault) ToPbText() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefault) FromPbText(value string) error {
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

func (m *marshalpatternAutoPatternDefaultAutoIpDefault) ToYaml() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefault) FromYaml(value string) error {
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

func (m *marshalpatternAutoPatternDefaultAutoIpDefault) ToJson() (string, error) {
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

func (m *unMarshalpatternAutoPatternDefaultAutoIpDefault) FromJson(value string) error {
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

func (obj *patternAutoPatternDefaultAutoIpDefault) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternAutoPatternDefaultAutoIpDefault) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternAutoPatternDefaultAutoIpDefault) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternAutoPatternDefaultAutoIpDefault) Clone() (PatternAutoPatternDefaultAutoIpDefault, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternAutoPatternDefaultAutoIpDefault()
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

func (obj *patternAutoPatternDefaultAutoIpDefault) setNil() {
	obj.autoHolder = nil
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternAutoPatternDefaultAutoIpDefault is tBD
type PatternAutoPatternDefaultAutoIpDefault interface {
	Validation
	// msg marshals PatternAutoPatternDefaultAutoIpDefault to protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefault
	// and doesn't set defaults
	msg() *openapi.PatternAutoPatternDefaultAutoIpDefault
	// setMsg unmarshals PatternAutoPatternDefaultAutoIpDefault from protobuf object *openapi.PatternAutoPatternDefaultAutoIpDefault
	// and doesn't set defaults
	setMsg(*openapi.PatternAutoPatternDefaultAutoIpDefault) PatternAutoPatternDefaultAutoIpDefault
	// provides marshal interface
	Marshal() marshalPatternAutoPatternDefaultAutoIpDefault
	// provides unmarshal interface
	Unmarshal() unMarshalPatternAutoPatternDefaultAutoIpDefault
	// validate validates PatternAutoPatternDefaultAutoIpDefault
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternAutoPatternDefaultAutoIpDefault, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternAutoPatternDefaultAutoIpDefaultChoiceEnum, set in PatternAutoPatternDefaultAutoIpDefault
	Choice() PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
	// setChoice assigns PatternAutoPatternDefaultAutoIpDefaultChoiceEnum provided by user to PatternAutoPatternDefaultAutoIpDefault
	setChoice(value PatternAutoPatternDefaultAutoIpDefaultChoiceEnum) PatternAutoPatternDefaultAutoIpDefault
	// HasChoice checks if Choice has been set in PatternAutoPatternDefaultAutoIpDefault
	HasChoice() bool
	// Value returns string, set in PatternAutoPatternDefaultAutoIpDefault.
	Value() string
	// SetValue assigns string provided by user to PatternAutoPatternDefaultAutoIpDefault
	SetValue(value string) PatternAutoPatternDefaultAutoIpDefault
	// HasValue checks if Value has been set in PatternAutoPatternDefaultAutoIpDefault
	HasValue() bool
	// Values returns []string, set in PatternAutoPatternDefaultAutoIpDefault.
	Values() []string
	// SetValues assigns []string provided by user to PatternAutoPatternDefaultAutoIpDefault
	SetValues(value []string) PatternAutoPatternDefaultAutoIpDefault
	// Auto returns AutoIpDefault, set in PatternAutoPatternDefaultAutoIpDefault.
	// AutoIpDefault is the OTG implementation can provide a system generated,
	// value for this property. If the OTG is unable to generate a value,
	// the default value must be used.
	Auto() AutoIpDefault
	// HasAuto checks if Auto has been set in PatternAutoPatternDefaultAutoIpDefault
	HasAuto() bool
	// Increment returns PatternAutoPatternDefaultAutoIpDefaultCounter, set in PatternAutoPatternDefaultAutoIpDefault.
	// PatternAutoPatternDefaultAutoIpDefaultCounter is ipv4 counter pattern
	Increment() PatternAutoPatternDefaultAutoIpDefaultCounter
	// SetIncrement assigns PatternAutoPatternDefaultAutoIpDefaultCounter provided by user to PatternAutoPatternDefaultAutoIpDefault.
	// PatternAutoPatternDefaultAutoIpDefaultCounter is ipv4 counter pattern
	SetIncrement(value PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefault
	// HasIncrement checks if Increment has been set in PatternAutoPatternDefaultAutoIpDefault
	HasIncrement() bool
	// Decrement returns PatternAutoPatternDefaultAutoIpDefaultCounter, set in PatternAutoPatternDefaultAutoIpDefault.
	// PatternAutoPatternDefaultAutoIpDefaultCounter is ipv4 counter pattern
	Decrement() PatternAutoPatternDefaultAutoIpDefaultCounter
	// SetDecrement assigns PatternAutoPatternDefaultAutoIpDefaultCounter provided by user to PatternAutoPatternDefaultAutoIpDefault.
	// PatternAutoPatternDefaultAutoIpDefaultCounter is ipv4 counter pattern
	SetDecrement(value PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefault
	// HasDecrement checks if Decrement has been set in PatternAutoPatternDefaultAutoIpDefault
	HasDecrement() bool
	setNil()
}

type PatternAutoPatternDefaultAutoIpDefaultChoiceEnum string

// Enum of Choice on PatternAutoPatternDefaultAutoIpDefault
var PatternAutoPatternDefaultAutoIpDefaultChoice = struct {
	VALUE     PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
	VALUES    PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
	AUTO      PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
	INCREMENT PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
	DECREMENT PatternAutoPatternDefaultAutoIpDefaultChoiceEnum
}{
	VALUE:     PatternAutoPatternDefaultAutoIpDefaultChoiceEnum("value"),
	VALUES:    PatternAutoPatternDefaultAutoIpDefaultChoiceEnum("values"),
	AUTO:      PatternAutoPatternDefaultAutoIpDefaultChoiceEnum("auto"),
	INCREMENT: PatternAutoPatternDefaultAutoIpDefaultChoiceEnum("increment"),
	DECREMENT: PatternAutoPatternDefaultAutoIpDefaultChoiceEnum("decrement"),
}

func (obj *patternAutoPatternDefaultAutoIpDefault) Choice() PatternAutoPatternDefaultAutoIpDefaultChoiceEnum {
	return PatternAutoPatternDefaultAutoIpDefaultChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternAutoPatternDefaultAutoIpDefault) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternAutoPatternDefaultAutoIpDefault) setChoice(value PatternAutoPatternDefaultAutoIpDefaultChoiceEnum) PatternAutoPatternDefaultAutoIpDefault {
	intValue, ok := openapi.PatternAutoPatternDefaultAutoIpDefault_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternAutoPatternDefaultAutoIpDefaultChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternAutoPatternDefaultAutoIpDefault_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternAutoPatternDefaultAutoIpDefaultChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternAutoPatternDefaultAutoIpDefaultChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternAutoPatternDefaultAutoIpDefaultChoice.AUTO {
		obj.obj.Auto = NewAutoIpDefault().msg()
	}

	if value == PatternAutoPatternDefaultAutoIpDefaultChoice.INCREMENT {
		obj.obj.Increment = NewPatternAutoPatternDefaultAutoIpDefaultCounter().msg()
	}

	if value == PatternAutoPatternDefaultAutoIpDefaultChoice.DECREMENT {
		obj.obj.Decrement = NewPatternAutoPatternDefaultAutoIpDefaultCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternAutoPatternDefaultAutoIpDefault) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternAutoPatternDefaultAutoIpDefault) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternAutoPatternDefaultAutoIpDefault object
func (obj *patternAutoPatternDefaultAutoIpDefault) SetValue(value string) PatternAutoPatternDefaultAutoIpDefault {
	obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternAutoPatternDefaultAutoIpDefault) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternAutoPatternDefaultAutoIpDefault object
func (obj *patternAutoPatternDefaultAutoIpDefault) SetValues(value []string) PatternAutoPatternDefaultAutoIpDefault {
	obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Auto returns a AutoIpDefault
func (obj *patternAutoPatternDefaultAutoIpDefault) Auto() AutoIpDefault {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &autoIpDefault{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a AutoIpDefault
func (obj *patternAutoPatternDefaultAutoIpDefault) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternAutoPatternDefaultAutoIpDefaultCounter
func (obj *patternAutoPatternDefaultAutoIpDefault) Increment() PatternAutoPatternDefaultAutoIpDefaultCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternAutoPatternDefaultAutoIpDefaultCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternAutoPatternDefaultAutoIpDefaultCounter
func (obj *patternAutoPatternDefaultAutoIpDefault) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternAutoPatternDefaultAutoIpDefaultCounter value in the PatternAutoPatternDefaultAutoIpDefault object
func (obj *patternAutoPatternDefaultAutoIpDefault) SetIncrement(value PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefault {
	obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternAutoPatternDefaultAutoIpDefaultCounter
func (obj *patternAutoPatternDefaultAutoIpDefault) Decrement() PatternAutoPatternDefaultAutoIpDefaultCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternAutoPatternDefaultAutoIpDefaultCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternAutoPatternDefaultAutoIpDefaultCounter
func (obj *patternAutoPatternDefaultAutoIpDefault) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternAutoPatternDefaultAutoIpDefaultCounter value in the PatternAutoPatternDefaultAutoIpDefault object
func (obj *patternAutoPatternDefaultAutoIpDefault) SetDecrement(value PatternAutoPatternDefaultAutoIpDefaultCounter) PatternAutoPatternDefaultAutoIpDefault {
	obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternAutoPatternDefaultAutoIpDefault) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternDefaultAutoIpDefault.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternDefaultAutoIpDefault.Values"))
		}

	}

	if obj.obj.Auto != nil {

		obj.Auto().validateObj(vObj, set_default)
	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternAutoPatternDefaultAutoIpDefault) setDefault() {
	var choices_set int = 0
	var choice PatternAutoPatternDefaultAutoIpDefaultChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternAutoPatternDefaultAutoIpDefaultChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternAutoPatternDefaultAutoIpDefaultChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternAutoPatternDefaultAutoIpDefaultChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternAutoPatternDefaultAutoIpDefaultChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternAutoPatternDefaultAutoIpDefaultChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternAutoPatternDefaultAutoIpDefaultChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternAutoPatternDefaultAutoIpDefault")
			}
		} else {
			intVal := openapi.PatternAutoPatternDefaultAutoIpDefault_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternAutoPatternDefaultAutoIpDefault_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
