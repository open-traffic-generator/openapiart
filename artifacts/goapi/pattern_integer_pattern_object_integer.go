package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternIntegerPatternObjectInteger *****
type patternIntegerPatternObjectInteger struct {
	validation
	obj             *openapi.PatternIntegerPatternObjectInteger
	marshaller      marshalPatternIntegerPatternObjectInteger
	unMarshaller    unMarshalPatternIntegerPatternObjectInteger
	incrementHolder PatternIntegerPatternObjectIntegerCounter
	decrementHolder PatternIntegerPatternObjectIntegerCounter
}

func NewPatternIntegerPatternObjectInteger() PatternIntegerPatternObjectInteger {
	obj := patternIntegerPatternObjectInteger{obj: &openapi.PatternIntegerPatternObjectInteger{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternObjectInteger) msg() *openapi.PatternIntegerPatternObjectInteger {
	return obj.obj
}

func (obj *patternIntegerPatternObjectInteger) setMsg(msg *openapi.PatternIntegerPatternObjectInteger) PatternIntegerPatternObjectInteger {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternIntegerPatternObjectInteger struct {
	obj *patternIntegerPatternObjectInteger
}

type marshalPatternIntegerPatternObjectInteger interface {
	// ToProto marshals PatternIntegerPatternObjectInteger to protobuf object *openapi.PatternIntegerPatternObjectInteger
	ToProto() (*openapi.PatternIntegerPatternObjectInteger, error)
	// ToPbText marshals PatternIntegerPatternObjectInteger to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternObjectInteger to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternObjectInteger to JSON text
	ToJson() (string, error)
}

type unMarshalpatternIntegerPatternObjectInteger struct {
	obj *patternIntegerPatternObjectInteger
}

type unMarshalPatternIntegerPatternObjectInteger interface {
	// FromProto unmarshals PatternIntegerPatternObjectInteger from protobuf object *openapi.PatternIntegerPatternObjectInteger
	FromProto(msg *openapi.PatternIntegerPatternObjectInteger) (PatternIntegerPatternObjectInteger, error)
	// FromPbText unmarshals PatternIntegerPatternObjectInteger from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternObjectInteger from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternObjectInteger from JSON text
	FromJson(value string) error
}

func (obj *patternIntegerPatternObjectInteger) Marshal() marshalPatternIntegerPatternObjectInteger {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternIntegerPatternObjectInteger{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternIntegerPatternObjectInteger) Unmarshal() unMarshalPatternIntegerPatternObjectInteger {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternIntegerPatternObjectInteger{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternIntegerPatternObjectInteger) ToProto() (*openapi.PatternIntegerPatternObjectInteger, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternIntegerPatternObjectInteger) FromProto(msg *openapi.PatternIntegerPatternObjectInteger) (PatternIntegerPatternObjectInteger, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternIntegerPatternObjectInteger) ToPbText() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectInteger) FromPbText(value string) error {
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

func (m *marshalpatternIntegerPatternObjectInteger) ToYaml() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectInteger) FromYaml(value string) error {
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

func (m *marshalpatternIntegerPatternObjectInteger) ToJson() (string, error) {
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

func (m *unMarshalpatternIntegerPatternObjectInteger) FromJson(value string) error {
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

func (obj *patternIntegerPatternObjectInteger) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternObjectInteger) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternObjectInteger) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternObjectInteger) Clone() (PatternIntegerPatternObjectInteger, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternObjectInteger()
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

func (obj *patternIntegerPatternObjectInteger) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIntegerPatternObjectInteger is tBD
type PatternIntegerPatternObjectInteger interface {
	Validation
	// msg marshals PatternIntegerPatternObjectInteger to protobuf object *openapi.PatternIntegerPatternObjectInteger
	// and doesn't set defaults
	msg() *openapi.PatternIntegerPatternObjectInteger
	// setMsg unmarshals PatternIntegerPatternObjectInteger from protobuf object *openapi.PatternIntegerPatternObjectInteger
	// and doesn't set defaults
	setMsg(*openapi.PatternIntegerPatternObjectInteger) PatternIntegerPatternObjectInteger
	// provides marshal interface
	Marshal() marshalPatternIntegerPatternObjectInteger
	// provides unmarshal interface
	Unmarshal() unMarshalPatternIntegerPatternObjectInteger
	// validate validates PatternIntegerPatternObjectInteger
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternObjectInteger, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIntegerPatternObjectIntegerChoiceEnum, set in PatternIntegerPatternObjectInteger
	Choice() PatternIntegerPatternObjectIntegerChoiceEnum
	// setChoice assigns PatternIntegerPatternObjectIntegerChoiceEnum provided by user to PatternIntegerPatternObjectInteger
	setChoice(value PatternIntegerPatternObjectIntegerChoiceEnum) PatternIntegerPatternObjectInteger
	// HasChoice checks if Choice has been set in PatternIntegerPatternObjectInteger
	HasChoice() bool
	// Value returns uint32, set in PatternIntegerPatternObjectInteger.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternIntegerPatternObjectInteger
	SetValue(value uint32) PatternIntegerPatternObjectInteger
	// HasValue checks if Value has been set in PatternIntegerPatternObjectInteger
	HasValue() bool
	// Values returns []uint32, set in PatternIntegerPatternObjectInteger.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternIntegerPatternObjectInteger
	SetValues(value []uint32) PatternIntegerPatternObjectInteger
	// Increment returns PatternIntegerPatternObjectIntegerCounter, set in PatternIntegerPatternObjectInteger.
	// PatternIntegerPatternObjectIntegerCounter is integer counter pattern
	Increment() PatternIntegerPatternObjectIntegerCounter
	// SetIncrement assigns PatternIntegerPatternObjectIntegerCounter provided by user to PatternIntegerPatternObjectInteger.
	// PatternIntegerPatternObjectIntegerCounter is integer counter pattern
	SetIncrement(value PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectInteger
	// HasIncrement checks if Increment has been set in PatternIntegerPatternObjectInteger
	HasIncrement() bool
	// Decrement returns PatternIntegerPatternObjectIntegerCounter, set in PatternIntegerPatternObjectInteger.
	// PatternIntegerPatternObjectIntegerCounter is integer counter pattern
	Decrement() PatternIntegerPatternObjectIntegerCounter
	// SetDecrement assigns PatternIntegerPatternObjectIntegerCounter provided by user to PatternIntegerPatternObjectInteger.
	// PatternIntegerPatternObjectIntegerCounter is integer counter pattern
	SetDecrement(value PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectInteger
	// HasDecrement checks if Decrement has been set in PatternIntegerPatternObjectInteger
	HasDecrement() bool
	setNil()
}

type PatternIntegerPatternObjectIntegerChoiceEnum string

// Enum of Choice on PatternIntegerPatternObjectInteger
var PatternIntegerPatternObjectIntegerChoice = struct {
	VALUE     PatternIntegerPatternObjectIntegerChoiceEnum
	VALUES    PatternIntegerPatternObjectIntegerChoiceEnum
	INCREMENT PatternIntegerPatternObjectIntegerChoiceEnum
	DECREMENT PatternIntegerPatternObjectIntegerChoiceEnum
}{
	VALUE:     PatternIntegerPatternObjectIntegerChoiceEnum("value"),
	VALUES:    PatternIntegerPatternObjectIntegerChoiceEnum("values"),
	INCREMENT: PatternIntegerPatternObjectIntegerChoiceEnum("increment"),
	DECREMENT: PatternIntegerPatternObjectIntegerChoiceEnum("decrement"),
}

func (obj *patternIntegerPatternObjectInteger) Choice() PatternIntegerPatternObjectIntegerChoiceEnum {
	return PatternIntegerPatternObjectIntegerChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIntegerPatternObjectInteger) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIntegerPatternObjectInteger) setChoice(value PatternIntegerPatternObjectIntegerChoiceEnum) PatternIntegerPatternObjectInteger {
	intValue, ok := openapi.PatternIntegerPatternObjectInteger_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIntegerPatternObjectIntegerChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIntegerPatternObjectInteger_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIntegerPatternObjectIntegerChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternIntegerPatternObjectIntegerChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternIntegerPatternObjectIntegerChoice.INCREMENT {
		obj.obj.Increment = NewPatternIntegerPatternObjectIntegerCounter().msg()
	}

	if value == PatternIntegerPatternObjectIntegerChoice.DECREMENT {
		obj.obj.Decrement = NewPatternIntegerPatternObjectIntegerCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternObjectInteger) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternIntegerPatternObjectIntegerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternObjectInteger) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternIntegerPatternObjectInteger object
func (obj *patternIntegerPatternObjectInteger) SetValue(value uint32) PatternIntegerPatternObjectInteger {
	obj.setChoice(PatternIntegerPatternObjectIntegerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternIntegerPatternObjectInteger) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternIntegerPatternObjectInteger object
func (obj *patternIntegerPatternObjectInteger) SetValues(value []uint32) PatternIntegerPatternObjectInteger {
	obj.setChoice(PatternIntegerPatternObjectIntegerChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIntegerPatternObjectIntegerCounter
func (obj *patternIntegerPatternObjectInteger) Increment() PatternIntegerPatternObjectIntegerCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternIntegerPatternObjectIntegerChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIntegerPatternObjectIntegerCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIntegerPatternObjectIntegerCounter
func (obj *patternIntegerPatternObjectInteger) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIntegerPatternObjectIntegerCounter value in the PatternIntegerPatternObjectInteger object
func (obj *patternIntegerPatternObjectInteger) SetIncrement(value PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectInteger {
	obj.setChoice(PatternIntegerPatternObjectIntegerChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIntegerPatternObjectIntegerCounter
func (obj *patternIntegerPatternObjectInteger) Decrement() PatternIntegerPatternObjectIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternIntegerPatternObjectIntegerChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIntegerPatternObjectIntegerCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIntegerPatternObjectIntegerCounter
func (obj *patternIntegerPatternObjectInteger) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIntegerPatternObjectIntegerCounter value in the PatternIntegerPatternObjectInteger object
func (obj *patternIntegerPatternObjectInteger) SetDecrement(value PatternIntegerPatternObjectIntegerCounter) PatternIntegerPatternObjectInteger {
	obj.setChoice(PatternIntegerPatternObjectIntegerChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternIntegerPatternObjectInteger) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternObjectInteger.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternIntegerPatternObjectInteger.Values <= 255 but Got %d", item))
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

func (obj *patternIntegerPatternObjectInteger) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternIntegerPatternObjectIntegerChoice.VALUE)

	}

}
