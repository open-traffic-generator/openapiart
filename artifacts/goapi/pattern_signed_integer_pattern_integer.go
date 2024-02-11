package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternSignedIntegerPatternInteger *****
type patternSignedIntegerPatternInteger struct {
	validation
	obj             *openapi.PatternSignedIntegerPatternInteger
	marshaller      marshalPatternSignedIntegerPatternInteger
	unMarshaller    unMarshalPatternSignedIntegerPatternInteger
	incrementHolder PatternSignedIntegerPatternIntegerCounter
	decrementHolder PatternSignedIntegerPatternIntegerCounter
}

func NewPatternSignedIntegerPatternInteger() PatternSignedIntegerPatternInteger {
	obj := patternSignedIntegerPatternInteger{obj: &openapi.PatternSignedIntegerPatternInteger{}}
	obj.setDefault()
	return &obj
}

func (obj *patternSignedIntegerPatternInteger) msg() *openapi.PatternSignedIntegerPatternInteger {
	return obj.obj
}

func (obj *patternSignedIntegerPatternInteger) setMsg(msg *openapi.PatternSignedIntegerPatternInteger) PatternSignedIntegerPatternInteger {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternSignedIntegerPatternInteger struct {
	obj *patternSignedIntegerPatternInteger
}

type marshalPatternSignedIntegerPatternInteger interface {
	// ToProto marshals PatternSignedIntegerPatternInteger to protobuf object *openapi.PatternSignedIntegerPatternInteger
	ToProto() (*openapi.PatternSignedIntegerPatternInteger, error)
	// ToPbText marshals PatternSignedIntegerPatternInteger to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternSignedIntegerPatternInteger to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternSignedIntegerPatternInteger to JSON text
	ToJson() (string, error)
}

type unMarshalpatternSignedIntegerPatternInteger struct {
	obj *patternSignedIntegerPatternInteger
}

type unMarshalPatternSignedIntegerPatternInteger interface {
	// FromProto unmarshals PatternSignedIntegerPatternInteger from protobuf object *openapi.PatternSignedIntegerPatternInteger
	FromProto(msg *openapi.PatternSignedIntegerPatternInteger) (PatternSignedIntegerPatternInteger, error)
	// FromPbText unmarshals PatternSignedIntegerPatternInteger from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternSignedIntegerPatternInteger from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternSignedIntegerPatternInteger from JSON text
	FromJson(value string) error
}

func (obj *patternSignedIntegerPatternInteger) Marshal() marshalPatternSignedIntegerPatternInteger {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternSignedIntegerPatternInteger{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternSignedIntegerPatternInteger) Unmarshal() unMarshalPatternSignedIntegerPatternInteger {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternSignedIntegerPatternInteger{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternSignedIntegerPatternInteger) ToProto() (*openapi.PatternSignedIntegerPatternInteger, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternSignedIntegerPatternInteger) FromProto(msg *openapi.PatternSignedIntegerPatternInteger) (PatternSignedIntegerPatternInteger, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternSignedIntegerPatternInteger) ToPbText() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternInteger) FromPbText(value string) error {
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

func (m *marshalpatternSignedIntegerPatternInteger) ToYaml() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternInteger) FromYaml(value string) error {
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

func (m *marshalpatternSignedIntegerPatternInteger) ToJson() (string, error) {
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

func (m *unMarshalpatternSignedIntegerPatternInteger) FromJson(value string) error {
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

func (obj *patternSignedIntegerPatternInteger) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternSignedIntegerPatternInteger) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternSignedIntegerPatternInteger) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternSignedIntegerPatternInteger) Clone() (PatternSignedIntegerPatternInteger, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternSignedIntegerPatternInteger()
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

func (obj *patternSignedIntegerPatternInteger) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternSignedIntegerPatternInteger is tBD
type PatternSignedIntegerPatternInteger interface {
	Validation
	// msg marshals PatternSignedIntegerPatternInteger to protobuf object *openapi.PatternSignedIntegerPatternInteger
	// and doesn't set defaults
	msg() *openapi.PatternSignedIntegerPatternInteger
	// setMsg unmarshals PatternSignedIntegerPatternInteger from protobuf object *openapi.PatternSignedIntegerPatternInteger
	// and doesn't set defaults
	setMsg(*openapi.PatternSignedIntegerPatternInteger) PatternSignedIntegerPatternInteger
	// provides marshal interface
	Marshal() marshalPatternSignedIntegerPatternInteger
	// provides unmarshal interface
	Unmarshal() unMarshalPatternSignedIntegerPatternInteger
	// validate validates PatternSignedIntegerPatternInteger
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternSignedIntegerPatternInteger, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternSignedIntegerPatternIntegerChoiceEnum, set in PatternSignedIntegerPatternInteger
	Choice() PatternSignedIntegerPatternIntegerChoiceEnum
	// setChoice assigns PatternSignedIntegerPatternIntegerChoiceEnum provided by user to PatternSignedIntegerPatternInteger
	setChoice(value PatternSignedIntegerPatternIntegerChoiceEnum) PatternSignedIntegerPatternInteger
	// HasChoice checks if Choice has been set in PatternSignedIntegerPatternInteger
	HasChoice() bool
	// Value returns int32, set in PatternSignedIntegerPatternInteger.
	Value() int32
	// SetValue assigns int32 provided by user to PatternSignedIntegerPatternInteger
	SetValue(value int32) PatternSignedIntegerPatternInteger
	// HasValue checks if Value has been set in PatternSignedIntegerPatternInteger
	HasValue() bool
	// Values returns []int32, set in PatternSignedIntegerPatternInteger.
	Values() []int32
	// SetValues assigns []int32 provided by user to PatternSignedIntegerPatternInteger
	SetValues(value []int32) PatternSignedIntegerPatternInteger
	// Increment returns PatternSignedIntegerPatternIntegerCounter, set in PatternSignedIntegerPatternInteger.
	// PatternSignedIntegerPatternIntegerCounter is integer counter pattern
	Increment() PatternSignedIntegerPatternIntegerCounter
	// SetIncrement assigns PatternSignedIntegerPatternIntegerCounter provided by user to PatternSignedIntegerPatternInteger.
	// PatternSignedIntegerPatternIntegerCounter is integer counter pattern
	SetIncrement(value PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternInteger
	// HasIncrement checks if Increment has been set in PatternSignedIntegerPatternInteger
	HasIncrement() bool
	// Decrement returns PatternSignedIntegerPatternIntegerCounter, set in PatternSignedIntegerPatternInteger.
	// PatternSignedIntegerPatternIntegerCounter is integer counter pattern
	Decrement() PatternSignedIntegerPatternIntegerCounter
	// SetDecrement assigns PatternSignedIntegerPatternIntegerCounter provided by user to PatternSignedIntegerPatternInteger.
	// PatternSignedIntegerPatternIntegerCounter is integer counter pattern
	SetDecrement(value PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternInteger
	// HasDecrement checks if Decrement has been set in PatternSignedIntegerPatternInteger
	HasDecrement() bool
	setNil()
}

type PatternSignedIntegerPatternIntegerChoiceEnum string

// Enum of Choice on PatternSignedIntegerPatternInteger
var PatternSignedIntegerPatternIntegerChoice = struct {
	VALUE     PatternSignedIntegerPatternIntegerChoiceEnum
	VALUES    PatternSignedIntegerPatternIntegerChoiceEnum
	INCREMENT PatternSignedIntegerPatternIntegerChoiceEnum
	DECREMENT PatternSignedIntegerPatternIntegerChoiceEnum
}{
	VALUE:     PatternSignedIntegerPatternIntegerChoiceEnum("value"),
	VALUES:    PatternSignedIntegerPatternIntegerChoiceEnum("values"),
	INCREMENT: PatternSignedIntegerPatternIntegerChoiceEnum("increment"),
	DECREMENT: PatternSignedIntegerPatternIntegerChoiceEnum("decrement"),
}

func (obj *patternSignedIntegerPatternInteger) Choice() PatternSignedIntegerPatternIntegerChoiceEnum {
	return PatternSignedIntegerPatternIntegerChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternSignedIntegerPatternInteger) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternSignedIntegerPatternInteger) setChoice(value PatternSignedIntegerPatternIntegerChoiceEnum) PatternSignedIntegerPatternInteger {
	intValue, ok := openapi.PatternSignedIntegerPatternInteger_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternSignedIntegerPatternIntegerChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternSignedIntegerPatternInteger_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternSignedIntegerPatternIntegerChoice.VALUE {
		defaultValue := int32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternSignedIntegerPatternIntegerChoice.VALUES {
		defaultValue := []int32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternSignedIntegerPatternIntegerChoice.INCREMENT {
		obj.obj.Increment = NewPatternSignedIntegerPatternIntegerCounter().msg()
	}

	if value == PatternSignedIntegerPatternIntegerChoice.DECREMENT {
		obj.obj.Decrement = NewPatternSignedIntegerPatternIntegerCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a int32
func (obj *patternSignedIntegerPatternInteger) Value() int32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternSignedIntegerPatternIntegerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a int32
func (obj *patternSignedIntegerPatternInteger) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the int32 value in the PatternSignedIntegerPatternInteger object
func (obj *patternSignedIntegerPatternInteger) SetValue(value int32) PatternSignedIntegerPatternInteger {
	obj.setChoice(PatternSignedIntegerPatternIntegerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []int32
func (obj *patternSignedIntegerPatternInteger) Values() []int32 {
	if obj.obj.Values == nil {
		obj.SetValues([]int32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []int32 value in the PatternSignedIntegerPatternInteger object
func (obj *patternSignedIntegerPatternInteger) SetValues(value []int32) PatternSignedIntegerPatternInteger {
	obj.setChoice(PatternSignedIntegerPatternIntegerChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]int32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternSignedIntegerPatternIntegerCounter
func (obj *patternSignedIntegerPatternInteger) Increment() PatternSignedIntegerPatternIntegerCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternSignedIntegerPatternIntegerChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternSignedIntegerPatternIntegerCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternSignedIntegerPatternIntegerCounter
func (obj *patternSignedIntegerPatternInteger) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternSignedIntegerPatternIntegerCounter value in the PatternSignedIntegerPatternInteger object
func (obj *patternSignedIntegerPatternInteger) SetIncrement(value PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternInteger {
	obj.setChoice(PatternSignedIntegerPatternIntegerChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternSignedIntegerPatternIntegerCounter
func (obj *patternSignedIntegerPatternInteger) Decrement() PatternSignedIntegerPatternIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternSignedIntegerPatternIntegerChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternSignedIntegerPatternIntegerCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternSignedIntegerPatternIntegerCounter
func (obj *patternSignedIntegerPatternInteger) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternSignedIntegerPatternIntegerCounter value in the PatternSignedIntegerPatternInteger object
func (obj *patternSignedIntegerPatternInteger) SetDecrement(value PatternSignedIntegerPatternIntegerCounter) PatternSignedIntegerPatternInteger {
	obj.setChoice(PatternSignedIntegerPatternIntegerChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternSignedIntegerPatternInteger) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value < -128 || *obj.obj.Value > 127 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-128 <= PatternSignedIntegerPatternInteger.Value <= 127 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item < -128 || item > 127 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("-128 <= PatternSignedIntegerPatternInteger.Values <= 127 but Got %d", item))
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

func (obj *patternSignedIntegerPatternInteger) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternSignedIntegerPatternIntegerChoice.VALUE)

	}

}
