package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIntegerPatternInteger *****
type patternIntegerPatternInteger struct {
	validation
	obj             *openapi.PatternIntegerPatternInteger
	marshaller      marshalPatternIntegerPatternInteger
	unMarshaller    unMarshalPatternIntegerPatternInteger
	incrementHolder PatternIntegerPatternIntegerCounter
	decrementHolder PatternIntegerPatternIntegerCounter
}

func NewPatternIntegerPatternInteger() PatternIntegerPatternInteger {
	obj := patternIntegerPatternInteger{obj: &openapi.PatternIntegerPatternInteger{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternInteger) msg() *openapi.PatternIntegerPatternInteger {
	return obj.obj
}

func (obj *patternIntegerPatternInteger) setMsg(msg *openapi.PatternIntegerPatternInteger) PatternIntegerPatternInteger {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIntegerPatternInteger struct {
	obj *patternIntegerPatternInteger
}

type marshalPatternIntegerPatternInteger interface {
	// ToProto marshals PatternIntegerPatternInteger to protobuf object *openapi.PatternIntegerPatternInteger
	ToProto() (*openapi.PatternIntegerPatternInteger, error)
	// ToPbText marshals PatternIntegerPatternInteger to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternInteger to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternInteger to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIntegerPatternInteger struct {
	obj *patternIntegerPatternInteger
}

type unMarshalPatternIntegerPatternInteger interface {
	// FromProto unmarshals PatternIntegerPatternInteger from protobuf object *openapi.PatternIntegerPatternInteger
	FromProto(msg *openapi.PatternIntegerPatternInteger) (PatternIntegerPatternInteger, error)
	// FromPbText unmarshals PatternIntegerPatternInteger from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternInteger from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternInteger from JSON text
	FromJson(value string) error
}

func (obj *patternIntegerPatternInteger) Marshal() marshalPatternIntegerPatternInteger {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIntegerPatternInteger{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIntegerPatternInteger) Unmarshal() unMarshalPatternIntegerPatternInteger {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIntegerPatternInteger{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIntegerPatternInteger) ToProto() (*openapi.PatternIntegerPatternInteger, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIntegerPatternInteger) FromProto(msg *openapi.PatternIntegerPatternInteger) (PatternIntegerPatternInteger, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIntegerPatternInteger) ToPbText() (string, error) {
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

func (m *unMarshalpatternIntegerPatternInteger) FromPbText(value string) error {
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

func (m *marshalpatternIntegerPatternInteger) ToYaml() (string, error) {
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

func (m *unMarshalpatternIntegerPatternInteger) FromYaml(value string) error {
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

func (m *marshalpatternIntegerPatternInteger) ToJson() (string, error) {
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

func (m *unMarshalpatternIntegerPatternInteger) FromJson(value string) error {
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

func (obj *patternIntegerPatternInteger) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternInteger) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternInteger) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternInteger) Clone() (PatternIntegerPatternInteger, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternInteger()
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

func (obj *patternIntegerPatternInteger) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIntegerPatternInteger is tBD
type PatternIntegerPatternInteger interface {
	Validation
	// msg marshals PatternIntegerPatternInteger to protobuf object *openapi.PatternIntegerPatternInteger
	// and doesn't set defaults
	msg() *openapi.PatternIntegerPatternInteger
	// setMsg unmarshals PatternIntegerPatternInteger from protobuf object *openapi.PatternIntegerPatternInteger
	// and doesn't set defaults
	setMsg(*openapi.PatternIntegerPatternInteger) PatternIntegerPatternInteger
	// provides marshal interface
	Marshal() marshalPatternIntegerPatternInteger
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIntegerPatternInteger
	// validate validates PatternIntegerPatternInteger
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternInteger, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIntegerPatternIntegerChoiceEnum, set in PatternIntegerPatternInteger
	Choice() PatternIntegerPatternIntegerChoiceEnum
	// setChoice assigns PatternIntegerPatternIntegerChoiceEnum provided by user to PatternIntegerPatternInteger
	setChoice(value PatternIntegerPatternIntegerChoiceEnum) PatternIntegerPatternInteger
	// HasChoice checks if Choice has been set in PatternIntegerPatternInteger
	HasChoice() bool
	// Value returns uint32, set in PatternIntegerPatternInteger.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternIntegerPatternInteger
	SetValue(value uint32) PatternIntegerPatternInteger
	// HasValue checks if Value has been set in PatternIntegerPatternInteger
	HasValue() bool
	// Values returns []uint32, set in PatternIntegerPatternInteger.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternIntegerPatternInteger
	SetValues(value []uint32) PatternIntegerPatternInteger
	// Increment returns PatternIntegerPatternIntegerCounter, set in PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	Increment() PatternIntegerPatternIntegerCounter
	// SetIncrement assigns PatternIntegerPatternIntegerCounter provided by user to PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	SetIncrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger
	// HasIncrement checks if Increment has been set in PatternIntegerPatternInteger
	HasIncrement() bool
	// Decrement returns PatternIntegerPatternIntegerCounter, set in PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	Decrement() PatternIntegerPatternIntegerCounter
	// SetDecrement assigns PatternIntegerPatternIntegerCounter provided by user to PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	SetDecrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger
	// HasDecrement checks if Decrement has been set in PatternIntegerPatternInteger
	HasDecrement() bool
	setNil()
}

type PatternIntegerPatternIntegerChoiceEnum string

// Enum of Choice on PatternIntegerPatternInteger
var PatternIntegerPatternIntegerChoice = struct {
	VALUE     PatternIntegerPatternIntegerChoiceEnum
	VALUES    PatternIntegerPatternIntegerChoiceEnum
	INCREMENT PatternIntegerPatternIntegerChoiceEnum
	DECREMENT PatternIntegerPatternIntegerChoiceEnum
}{
	VALUE:     PatternIntegerPatternIntegerChoiceEnum("value"),
	VALUES:    PatternIntegerPatternIntegerChoiceEnum("values"),
	INCREMENT: PatternIntegerPatternIntegerChoiceEnum("increment"),
	DECREMENT: PatternIntegerPatternIntegerChoiceEnum("decrement"),
}

func (obj *patternIntegerPatternInteger) Choice() PatternIntegerPatternIntegerChoiceEnum {
	return PatternIntegerPatternIntegerChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIntegerPatternInteger) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIntegerPatternInteger) setChoice(value PatternIntegerPatternIntegerChoiceEnum) PatternIntegerPatternInteger {
	intValue, ok := openapi.PatternIntegerPatternInteger_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIntegerPatternIntegerChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIntegerPatternInteger_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIntegerPatternIntegerChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternIntegerPatternIntegerChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternIntegerPatternIntegerChoice.INCREMENT {
		obj.obj.Increment = NewPatternIntegerPatternIntegerCounter().msg()
	}

	if value == PatternIntegerPatternIntegerChoice.DECREMENT {
		obj.obj.Decrement = NewPatternIntegerPatternIntegerCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternInteger) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIntegerPatternIntegerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternInteger) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetValue(value uint32) PatternIntegerPatternInteger {
	obj.setChoice(PatternIntegerPatternIntegerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternIntegerPatternInteger) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetValues(value []uint32) PatternIntegerPatternInteger {
	obj.setChoice(PatternIntegerPatternIntegerChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) Increment() PatternIntegerPatternIntegerCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIntegerPatternIntegerChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIntegerPatternIntegerCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIntegerPatternIntegerCounter value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetIncrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger {
	obj.setChoice(PatternIntegerPatternIntegerChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) Decrement() PatternIntegerPatternIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIntegerPatternIntegerChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIntegerPatternIntegerCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIntegerPatternIntegerCounter value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetDecrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger {
	obj.setChoice(PatternIntegerPatternIntegerChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIntegerPatternInteger) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternInteger.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternIntegerPatternInteger.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIntegerPatternInteger) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternIntegerPatternIntegerChoice.VALUE)

	}

}
