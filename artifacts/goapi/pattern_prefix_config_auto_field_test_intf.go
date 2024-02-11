package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternPrefixConfigAutoFieldTest *****
type patternPrefixConfigAutoFieldTest struct {
	validation
	obj             *openapi.PatternPrefixConfigAutoFieldTest
	marshaller      marshalPatternPrefixConfigAutoFieldTest
	unMarshaller    unMarshalPatternPrefixConfigAutoFieldTest
	incrementHolder PatternPrefixConfigAutoFieldTestCounter
	decrementHolder PatternPrefixConfigAutoFieldTestCounter
}

func NewPatternPrefixConfigAutoFieldTest() PatternPrefixConfigAutoFieldTest {
	obj := patternPrefixConfigAutoFieldTest{obj: &openapi.PatternPrefixConfigAutoFieldTest{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigAutoFieldTest) msg() *openapi.PatternPrefixConfigAutoFieldTest {
	return obj.obj
}

func (obj *patternPrefixConfigAutoFieldTest) setMsg(msg *openapi.PatternPrefixConfigAutoFieldTest) PatternPrefixConfigAutoFieldTest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternPrefixConfigAutoFieldTest struct {
	obj *patternPrefixConfigAutoFieldTest
}

type marshalPatternPrefixConfigAutoFieldTest interface {
	// ToProto marshals PatternPrefixConfigAutoFieldTest to protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	ToProto() (*openapi.PatternPrefixConfigAutoFieldTest, error)
	// ToPbText marshals PatternPrefixConfigAutoFieldTest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigAutoFieldTest to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigAutoFieldTest to JSON text
	ToJson() (string, error)
}

type unMarshalpatternPrefixConfigAutoFieldTest struct {
	obj *patternPrefixConfigAutoFieldTest
}

type unMarshalPatternPrefixConfigAutoFieldTest interface {
	// FromProto unmarshals PatternPrefixConfigAutoFieldTest from protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	FromProto(msg *openapi.PatternPrefixConfigAutoFieldTest) (PatternPrefixConfigAutoFieldTest, error)
	// FromPbText unmarshals PatternPrefixConfigAutoFieldTest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigAutoFieldTest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigAutoFieldTest from JSON text
	FromJson(value string) error
}

func (obj *patternPrefixConfigAutoFieldTest) Marshal() marshalPatternPrefixConfigAutoFieldTest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternPrefixConfigAutoFieldTest{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternPrefixConfigAutoFieldTest) Unmarshal() unMarshalPatternPrefixConfigAutoFieldTest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternPrefixConfigAutoFieldTest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternPrefixConfigAutoFieldTest) ToProto() (*openapi.PatternPrefixConfigAutoFieldTest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternPrefixConfigAutoFieldTest) FromProto(msg *openapi.PatternPrefixConfigAutoFieldTest) (PatternPrefixConfigAutoFieldTest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternPrefixConfigAutoFieldTest) ToPbText() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTest) FromPbText(value string) error {
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

func (m *marshalpatternPrefixConfigAutoFieldTest) ToYaml() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTest) FromYaml(value string) error {
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

func (m *marshalpatternPrefixConfigAutoFieldTest) ToJson() (string, error) {
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

func (m *unMarshalpatternPrefixConfigAutoFieldTest) FromJson(value string) error {
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

func (obj *patternPrefixConfigAutoFieldTest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigAutoFieldTest) Clone() (PatternPrefixConfigAutoFieldTest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigAutoFieldTest()
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

func (obj *patternPrefixConfigAutoFieldTest) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternPrefixConfigAutoFieldTest is tBD
type PatternPrefixConfigAutoFieldTest interface {
	Validation
	// msg marshals PatternPrefixConfigAutoFieldTest to protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	// and doesn't set defaults
	msg() *openapi.PatternPrefixConfigAutoFieldTest
	// setMsg unmarshals PatternPrefixConfigAutoFieldTest from protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	// and doesn't set defaults
	setMsg(*openapi.PatternPrefixConfigAutoFieldTest) PatternPrefixConfigAutoFieldTest
	// provides marshal interface
	Marshal() marshalPatternPrefixConfigAutoFieldTest
	// provides unmarshal interface
	Unmarshal() unMarshalPatternPrefixConfigAutoFieldTest
	// validate validates PatternPrefixConfigAutoFieldTest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigAutoFieldTest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternPrefixConfigAutoFieldTestChoiceEnum, set in PatternPrefixConfigAutoFieldTest
	Choice() PatternPrefixConfigAutoFieldTestChoiceEnum
	// setChoice assigns PatternPrefixConfigAutoFieldTestChoiceEnum provided by user to PatternPrefixConfigAutoFieldTest
	setChoice(value PatternPrefixConfigAutoFieldTestChoiceEnum) PatternPrefixConfigAutoFieldTest
	// HasChoice checks if Choice has been set in PatternPrefixConfigAutoFieldTest
	HasChoice() bool
	// Value returns uint32, set in PatternPrefixConfigAutoFieldTest.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternPrefixConfigAutoFieldTest
	SetValue(value uint32) PatternPrefixConfigAutoFieldTest
	// HasValue checks if Value has been set in PatternPrefixConfigAutoFieldTest
	HasValue() bool
	// Values returns []uint32, set in PatternPrefixConfigAutoFieldTest.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternPrefixConfigAutoFieldTest
	SetValues(value []uint32) PatternPrefixConfigAutoFieldTest
	// Auto returns uint32, set in PatternPrefixConfigAutoFieldTest.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternPrefixConfigAutoFieldTest
	HasAuto() bool
	// Increment returns PatternPrefixConfigAutoFieldTestCounter, set in PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	Increment() PatternPrefixConfigAutoFieldTestCounter
	// SetIncrement assigns PatternPrefixConfigAutoFieldTestCounter provided by user to PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	SetIncrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest
	// HasIncrement checks if Increment has been set in PatternPrefixConfigAutoFieldTest
	HasIncrement() bool
	// Decrement returns PatternPrefixConfigAutoFieldTestCounter, set in PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	Decrement() PatternPrefixConfigAutoFieldTestCounter
	// SetDecrement assigns PatternPrefixConfigAutoFieldTestCounter provided by user to PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	SetDecrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest
	// HasDecrement checks if Decrement has been set in PatternPrefixConfigAutoFieldTest
	HasDecrement() bool
	setNil()
}

type PatternPrefixConfigAutoFieldTestChoiceEnum string

// Enum of Choice on PatternPrefixConfigAutoFieldTest
var PatternPrefixConfigAutoFieldTestChoice = struct {
	VALUE     PatternPrefixConfigAutoFieldTestChoiceEnum
	VALUES    PatternPrefixConfigAutoFieldTestChoiceEnum
	AUTO      PatternPrefixConfigAutoFieldTestChoiceEnum
	INCREMENT PatternPrefixConfigAutoFieldTestChoiceEnum
	DECREMENT PatternPrefixConfigAutoFieldTestChoiceEnum
}{
	VALUE:     PatternPrefixConfigAutoFieldTestChoiceEnum("value"),
	VALUES:    PatternPrefixConfigAutoFieldTestChoiceEnum("values"),
	AUTO:      PatternPrefixConfigAutoFieldTestChoiceEnum("auto"),
	INCREMENT: PatternPrefixConfigAutoFieldTestChoiceEnum("increment"),
	DECREMENT: PatternPrefixConfigAutoFieldTestChoiceEnum("decrement"),
}

func (obj *patternPrefixConfigAutoFieldTest) Choice() PatternPrefixConfigAutoFieldTestChoiceEnum {
	return PatternPrefixConfigAutoFieldTestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternPrefixConfigAutoFieldTest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternPrefixConfigAutoFieldTest) setChoice(value PatternPrefixConfigAutoFieldTestChoiceEnum) PatternPrefixConfigAutoFieldTest {
	intValue, ok := openapi.PatternPrefixConfigAutoFieldTest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigAutoFieldTestChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigAutoFieldTest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternPrefixConfigAutoFieldTestChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.INCREMENT {
		obj.obj.Increment = NewPatternPrefixConfigAutoFieldTestCounter().msg()
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.DECREMENT {
		obj.obj.Decrement = NewPatternPrefixConfigAutoFieldTestCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) Value() uint32 {

	if obj.obj.Value == nil {
		obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetValue(value uint32) PatternPrefixConfigAutoFieldTest {
	obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternPrefixConfigAutoFieldTest) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetValues(value []uint32) PatternPrefixConfigAutoFieldTest {
	obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) Increment() PatternPrefixConfigAutoFieldTestCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternPrefixConfigAutoFieldTestCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternPrefixConfigAutoFieldTestCounter value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetIncrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest {
	obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) Decrement() PatternPrefixConfigAutoFieldTestCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternPrefixConfigAutoFieldTestCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternPrefixConfigAutoFieldTestCounter value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetDecrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest {
	obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternPrefixConfigAutoFieldTest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternPrefixConfigAutoFieldTest.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternPrefixConfigAutoFieldTest) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternPrefixConfigAutoFieldTestChoice.AUTO)

	}

}
