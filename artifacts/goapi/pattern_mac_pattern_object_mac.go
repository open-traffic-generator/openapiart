package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternMacPatternObjectMac *****
type patternMacPatternObjectMac struct {
	validation
	obj             *openapi.PatternMacPatternObjectMac
	marshaller      marshalPatternMacPatternObjectMac
	unMarshaller    unMarshalPatternMacPatternObjectMac
	incrementHolder PatternMacPatternObjectMacCounter
	decrementHolder PatternMacPatternObjectMacCounter
}

func NewPatternMacPatternObjectMac() PatternMacPatternObjectMac {
	obj := patternMacPatternObjectMac{obj: &openapi.PatternMacPatternObjectMac{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternObjectMac) msg() *openapi.PatternMacPatternObjectMac {
	return obj.obj
}

func (obj *patternMacPatternObjectMac) setMsg(msg *openapi.PatternMacPatternObjectMac) PatternMacPatternObjectMac {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternMacPatternObjectMac struct {
	obj *patternMacPatternObjectMac
}

type marshalPatternMacPatternObjectMac interface {
	// ToProto marshals PatternMacPatternObjectMac to protobuf object *openapi.PatternMacPatternObjectMac
	ToProto() (*openapi.PatternMacPatternObjectMac, error)
	// ToPbText marshals PatternMacPatternObjectMac to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternObjectMac to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternObjectMac to JSON text
	ToJson() (string, error)
}

type unMarshalpatternMacPatternObjectMac struct {
	obj *patternMacPatternObjectMac
}

type unMarshalPatternMacPatternObjectMac interface {
	// FromProto unmarshals PatternMacPatternObjectMac from protobuf object *openapi.PatternMacPatternObjectMac
	FromProto(msg *openapi.PatternMacPatternObjectMac) (PatternMacPatternObjectMac, error)
	// FromPbText unmarshals PatternMacPatternObjectMac from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternObjectMac from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternObjectMac from JSON text
	FromJson(value string) error
}

func (obj *patternMacPatternObjectMac) Marshal() marshalPatternMacPatternObjectMac {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternMacPatternObjectMac{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternMacPatternObjectMac) Unmarshal() unMarshalPatternMacPatternObjectMac {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternMacPatternObjectMac{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternMacPatternObjectMac) ToProto() (*openapi.PatternMacPatternObjectMac, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternMacPatternObjectMac) FromProto(msg *openapi.PatternMacPatternObjectMac) (PatternMacPatternObjectMac, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternMacPatternObjectMac) ToPbText() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMac) FromPbText(value string) error {
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

func (m *marshalpatternMacPatternObjectMac) ToYaml() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMac) FromYaml(value string) error {
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

func (m *marshalpatternMacPatternObjectMac) ToJson() (string, error) {
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

func (m *unMarshalpatternMacPatternObjectMac) FromJson(value string) error {
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

func (obj *patternMacPatternObjectMac) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternObjectMac) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternObjectMac) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternObjectMac) Clone() (PatternMacPatternObjectMac, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternObjectMac()
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

func (obj *patternMacPatternObjectMac) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternMacPatternObjectMac is tBD
type PatternMacPatternObjectMac interface {
	Validation
	// msg marshals PatternMacPatternObjectMac to protobuf object *openapi.PatternMacPatternObjectMac
	// and doesn't set defaults
	msg() *openapi.PatternMacPatternObjectMac
	// setMsg unmarshals PatternMacPatternObjectMac from protobuf object *openapi.PatternMacPatternObjectMac
	// and doesn't set defaults
	setMsg(*openapi.PatternMacPatternObjectMac) PatternMacPatternObjectMac
	// provides marshal interface
	Marshal() marshalPatternMacPatternObjectMac
	// provides unmarshal interface
	Unmarshal() unMarshalPatternMacPatternObjectMac
	// validate validates PatternMacPatternObjectMac
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternObjectMac, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternMacPatternObjectMacChoiceEnum, set in PatternMacPatternObjectMac
	Choice() PatternMacPatternObjectMacChoiceEnum
	// setChoice assigns PatternMacPatternObjectMacChoiceEnum provided by user to PatternMacPatternObjectMac
	setChoice(value PatternMacPatternObjectMacChoiceEnum) PatternMacPatternObjectMac
	// HasChoice checks if Choice has been set in PatternMacPatternObjectMac
	HasChoice() bool
	// Value returns string, set in PatternMacPatternObjectMac.
	Value() string
	// SetValue assigns string provided by user to PatternMacPatternObjectMac
	SetValue(value string) PatternMacPatternObjectMac
	// HasValue checks if Value has been set in PatternMacPatternObjectMac
	HasValue() bool
	// Values returns []string, set in PatternMacPatternObjectMac.
	Values() []string
	// SetValues assigns []string provided by user to PatternMacPatternObjectMac
	SetValues(value []string) PatternMacPatternObjectMac
	// Auto returns string, set in PatternMacPatternObjectMac.
	Auto() string
	// HasAuto checks if Auto has been set in PatternMacPatternObjectMac
	HasAuto() bool
	// Increment returns PatternMacPatternObjectMacCounter, set in PatternMacPatternObjectMac.
	// PatternMacPatternObjectMacCounter is mac counter pattern
	Increment() PatternMacPatternObjectMacCounter
	// SetIncrement assigns PatternMacPatternObjectMacCounter provided by user to PatternMacPatternObjectMac.
	// PatternMacPatternObjectMacCounter is mac counter pattern
	SetIncrement(value PatternMacPatternObjectMacCounter) PatternMacPatternObjectMac
	// HasIncrement checks if Increment has been set in PatternMacPatternObjectMac
	HasIncrement() bool
	// Decrement returns PatternMacPatternObjectMacCounter, set in PatternMacPatternObjectMac.
	// PatternMacPatternObjectMacCounter is mac counter pattern
	Decrement() PatternMacPatternObjectMacCounter
	// SetDecrement assigns PatternMacPatternObjectMacCounter provided by user to PatternMacPatternObjectMac.
	// PatternMacPatternObjectMacCounter is mac counter pattern
	SetDecrement(value PatternMacPatternObjectMacCounter) PatternMacPatternObjectMac
	// HasDecrement checks if Decrement has been set in PatternMacPatternObjectMac
	HasDecrement() bool
	setNil()
}

type PatternMacPatternObjectMacChoiceEnum string

// Enum of Choice on PatternMacPatternObjectMac
var PatternMacPatternObjectMacChoice = struct {
	VALUE     PatternMacPatternObjectMacChoiceEnum
	VALUES    PatternMacPatternObjectMacChoiceEnum
	AUTO      PatternMacPatternObjectMacChoiceEnum
	INCREMENT PatternMacPatternObjectMacChoiceEnum
	DECREMENT PatternMacPatternObjectMacChoiceEnum
}{
	VALUE:     PatternMacPatternObjectMacChoiceEnum("value"),
	VALUES:    PatternMacPatternObjectMacChoiceEnum("values"),
	AUTO:      PatternMacPatternObjectMacChoiceEnum("auto"),
	INCREMENT: PatternMacPatternObjectMacChoiceEnum("increment"),
	DECREMENT: PatternMacPatternObjectMacChoiceEnum("decrement"),
}

func (obj *patternMacPatternObjectMac) Choice() PatternMacPatternObjectMacChoiceEnum {
	return PatternMacPatternObjectMacChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternMacPatternObjectMac) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternMacPatternObjectMac) setChoice(value PatternMacPatternObjectMacChoiceEnum) PatternMacPatternObjectMac {
	intValue, ok := openapi.PatternMacPatternObjectMac_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternMacPatternObjectMacChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternMacPatternObjectMac_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternMacPatternObjectMacChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternMacPatternObjectMacChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternMacPatternObjectMacChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == PatternMacPatternObjectMacChoice.INCREMENT {
		obj.obj.Increment = NewPatternMacPatternObjectMacCounter().msg()
	}

	if value == PatternMacPatternObjectMacChoice.DECREMENT {
		obj.obj.Decrement = NewPatternMacPatternObjectMacCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternMacPatternObjectMac) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternMacPatternObjectMacChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternMacPatternObjectMac) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternMacPatternObjectMac object
func (obj *patternMacPatternObjectMac) SetValue(value string) PatternMacPatternObjectMac {
	obj.setChoice(PatternMacPatternObjectMacChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternMacPatternObjectMac) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternMacPatternObjectMac object
func (obj *patternMacPatternObjectMac) SetValues(value []string) PatternMacPatternObjectMac {
	obj.setChoice(PatternMacPatternObjectMacChoice.VALUES)
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
func (obj *patternMacPatternObjectMac) Auto() string {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternMacPatternObjectMacChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternMacPatternObjectMac) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternMacPatternObjectMacCounter
func (obj *patternMacPatternObjectMac) Increment() PatternMacPatternObjectMacCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternMacPatternObjectMacChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternMacPatternObjectMacCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternMacPatternObjectMacCounter
func (obj *patternMacPatternObjectMac) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternMacPatternObjectMacCounter value in the PatternMacPatternObjectMac object
func (obj *patternMacPatternObjectMac) SetIncrement(value PatternMacPatternObjectMacCounter) PatternMacPatternObjectMac {
	obj.setChoice(PatternMacPatternObjectMacChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternMacPatternObjectMacCounter
func (obj *patternMacPatternObjectMac) Decrement() PatternMacPatternObjectMacCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternMacPatternObjectMacChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternMacPatternObjectMacCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternMacPatternObjectMacCounter
func (obj *patternMacPatternObjectMac) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternMacPatternObjectMacCounter value in the PatternMacPatternObjectMac object
func (obj *patternMacPatternObjectMac) SetDecrement(value PatternMacPatternObjectMacCounter) PatternMacPatternObjectMac {
	obj.setChoice(PatternMacPatternObjectMacChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternMacPatternObjectMac) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternObjectMac.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternObjectMac.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternObjectMac.Auto"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternMacPatternObjectMac) setDefault() {
	var choices_set int = 0
	var choice PatternMacPatternObjectMacChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternMacPatternObjectMacChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternMacPatternObjectMacChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternMacPatternObjectMacChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternMacPatternObjectMacChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternMacPatternObjectMacChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternMacPatternObjectMacChoice.AUTO)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternMacPatternObjectMac")
			}
		} else {
			intVal := openapi.PatternMacPatternObjectMac_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternMacPatternObjectMac_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
