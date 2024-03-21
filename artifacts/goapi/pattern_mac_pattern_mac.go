package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternMacPatternMac *****
type patternMacPatternMac struct {
	validation
	obj             *openapi.PatternMacPatternMac
	marshaller      marshalPatternMacPatternMac
	unMarshaller    unMarshalPatternMacPatternMac
	incrementHolder PatternMacPatternMacCounter
	decrementHolder PatternMacPatternMacCounter
}

func NewPatternMacPatternMac() PatternMacPatternMac {
	obj := patternMacPatternMac{obj: &openapi.PatternMacPatternMac{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternMac) msg() *openapi.PatternMacPatternMac {
	return obj.obj
}

func (obj *patternMacPatternMac) setMsg(msg *openapi.PatternMacPatternMac) PatternMacPatternMac {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternMacPatternMac struct {
	obj *patternMacPatternMac
}

type marshalPatternMacPatternMac interface {
	// ToProto marshals PatternMacPatternMac to protobuf object *openapi.PatternMacPatternMac
	ToProto() (*openapi.PatternMacPatternMac, error)
	// ToPbText marshals PatternMacPatternMac to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternMac to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternMac to JSON text
	ToJson() (string, error)
}

type unMarshalpatternMacPatternMac struct {
	obj *patternMacPatternMac
}

type unMarshalPatternMacPatternMac interface {
	// FromProto unmarshals PatternMacPatternMac from protobuf object *openapi.PatternMacPatternMac
	FromProto(msg *openapi.PatternMacPatternMac) (PatternMacPatternMac, error)
	// FromPbText unmarshals PatternMacPatternMac from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternMac from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternMac from JSON text
	FromJson(value string) error
}

func (obj *patternMacPatternMac) Marshal() marshalPatternMacPatternMac {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternMacPatternMac{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternMacPatternMac) Unmarshal() unMarshalPatternMacPatternMac {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternMacPatternMac{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternMacPatternMac) ToProto() (*openapi.PatternMacPatternMac, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternMacPatternMac) FromProto(msg *openapi.PatternMacPatternMac) (PatternMacPatternMac, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternMacPatternMac) ToPbText() (string, error) {
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

func (m *unMarshalpatternMacPatternMac) FromPbText(value string) error {
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

func (m *marshalpatternMacPatternMac) ToYaml() (string, error) {
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

func (m *unMarshalpatternMacPatternMac) FromYaml(value string) error {
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

func (m *marshalpatternMacPatternMac) ToJson() (string, error) {
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

func (m *unMarshalpatternMacPatternMac) FromJson(value string) error {
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

func (obj *patternMacPatternMac) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternMac) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternMac) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternMac) Clone() (PatternMacPatternMac, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternMac()
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

func (obj *patternMacPatternMac) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternMacPatternMac is tBD
type PatternMacPatternMac interface {
	Validation
	// msg marshals PatternMacPatternMac to protobuf object *openapi.PatternMacPatternMac
	// and doesn't set defaults
	msg() *openapi.PatternMacPatternMac
	// setMsg unmarshals PatternMacPatternMac from protobuf object *openapi.PatternMacPatternMac
	// and doesn't set defaults
	setMsg(*openapi.PatternMacPatternMac) PatternMacPatternMac
	// provides marshal interface
	Marshal() marshalPatternMacPatternMac
	// provides unmarshal interface
	Unmarshal() unMarshalPatternMacPatternMac
	// validate validates PatternMacPatternMac
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternMac, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternMacPatternMacChoiceEnum, set in PatternMacPatternMac
	Choice() PatternMacPatternMacChoiceEnum
	// setChoice assigns PatternMacPatternMacChoiceEnum provided by user to PatternMacPatternMac
	setChoice(value PatternMacPatternMacChoiceEnum) PatternMacPatternMac
	// HasChoice checks if Choice has been set in PatternMacPatternMac
	HasChoice() bool
	// Value returns string, set in PatternMacPatternMac.
	Value() string
	// SetValue assigns string provided by user to PatternMacPatternMac
	SetValue(value string) PatternMacPatternMac
	// HasValue checks if Value has been set in PatternMacPatternMac
	HasValue() bool
	// Values returns []string, set in PatternMacPatternMac.
	Values() []string
	// SetValues assigns []string provided by user to PatternMacPatternMac
	SetValues(value []string) PatternMacPatternMac
	// Auto returns string, set in PatternMacPatternMac.
	Auto() string
	// HasAuto checks if Auto has been set in PatternMacPatternMac
	HasAuto() bool
	// Increment returns PatternMacPatternMacCounter, set in PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	Increment() PatternMacPatternMacCounter
	// SetIncrement assigns PatternMacPatternMacCounter provided by user to PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	SetIncrement(value PatternMacPatternMacCounter) PatternMacPatternMac
	// HasIncrement checks if Increment has been set in PatternMacPatternMac
	HasIncrement() bool
	// Decrement returns PatternMacPatternMacCounter, set in PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	Decrement() PatternMacPatternMacCounter
	// SetDecrement assigns PatternMacPatternMacCounter provided by user to PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	SetDecrement(value PatternMacPatternMacCounter) PatternMacPatternMac
	// HasDecrement checks if Decrement has been set in PatternMacPatternMac
	HasDecrement() bool
	setNil()
}

type PatternMacPatternMacChoiceEnum string

// Enum of Choice on PatternMacPatternMac
var PatternMacPatternMacChoice = struct {
	VALUE     PatternMacPatternMacChoiceEnum
	VALUES    PatternMacPatternMacChoiceEnum
	AUTO      PatternMacPatternMacChoiceEnum
	INCREMENT PatternMacPatternMacChoiceEnum
	DECREMENT PatternMacPatternMacChoiceEnum
}{
	VALUE:     PatternMacPatternMacChoiceEnum("value"),
	VALUES:    PatternMacPatternMacChoiceEnum("values"),
	AUTO:      PatternMacPatternMacChoiceEnum("auto"),
	INCREMENT: PatternMacPatternMacChoiceEnum("increment"),
	DECREMENT: PatternMacPatternMacChoiceEnum("decrement"),
}

func (obj *patternMacPatternMac) Choice() PatternMacPatternMacChoiceEnum {
	return PatternMacPatternMacChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternMacPatternMac) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternMacPatternMac) setChoice(value PatternMacPatternMacChoiceEnum) PatternMacPatternMac {
	intValue, ok := openapi.PatternMacPatternMac_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternMacPatternMacChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternMacPatternMac_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternMacPatternMacChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternMacPatternMacChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternMacPatternMacChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == PatternMacPatternMacChoice.INCREMENT {
		obj.obj.Increment = NewPatternMacPatternMacCounter().msg()
	}

	if value == PatternMacPatternMacChoice.DECREMENT {
		obj.obj.Decrement = NewPatternMacPatternMacCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternMacPatternMac) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternMacPatternMacChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternMacPatternMac) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetValue(value string) PatternMacPatternMac {
	obj.setChoice(PatternMacPatternMacChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternMacPatternMac) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetValues(value []string) PatternMacPatternMac {
	obj.setChoice(PatternMacPatternMacChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternMacPatternMac) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternMacPatternMacChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternMacPatternMac) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) Increment() PatternMacPatternMacCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternMacPatternMacChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternMacPatternMacCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternMacPatternMacCounter value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetIncrement(value PatternMacPatternMacCounter) PatternMacPatternMac {
	obj.setChoice(PatternMacPatternMacChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) Decrement() PatternMacPatternMacCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternMacPatternMacChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternMacPatternMacCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternMacPatternMacCounter value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetDecrement(value PatternMacPatternMacCounter) PatternMacPatternMac {
	obj.setChoice(PatternMacPatternMacChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternMacPatternMac) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Auto"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternMacPatternMac) setDefault() {
	var choices_set int = 0
	var choice PatternMacPatternMacChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternMacPatternMacChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternMacPatternMacChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternMacPatternMacChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternMacPatternMacChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternMacPatternMacChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternMacPatternMacChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternMacPatternMac")
			}
		} else {
			intVal := openapi.PatternMacPatternMac_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternMacPatternMac_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
