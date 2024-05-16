package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternAutoPatternAutoIp *****
type patternAutoPatternAutoIp struct {
	validation
	obj             *openapi.PatternAutoPatternAutoIp
	marshaller      marshalPatternAutoPatternAutoIp
	unMarshaller    unMarshalPatternAutoPatternAutoIp
	autoHolder      AutoIpOptions
	incrementHolder PatternAutoPatternAutoIpCounter
	decrementHolder PatternAutoPatternAutoIpCounter
}

func NewPatternAutoPatternAutoIp() PatternAutoPatternAutoIp {
	obj := patternAutoPatternAutoIp{obj: &openapi.PatternAutoPatternAutoIp{}}
	obj.setDefault()
	return &obj
}

func (obj *patternAutoPatternAutoIp) msg() *openapi.PatternAutoPatternAutoIp {
	return obj.obj
}

func (obj *patternAutoPatternAutoIp) setMsg(msg *openapi.PatternAutoPatternAutoIp) PatternAutoPatternAutoIp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternAutoPatternAutoIp struct {
	obj *patternAutoPatternAutoIp
}

type marshalPatternAutoPatternAutoIp interface {
	// ToProto marshals PatternAutoPatternAutoIp to protobuf object *openapi.PatternAutoPatternAutoIp
	ToProto() (*openapi.PatternAutoPatternAutoIp, error)
	// ToPbText marshals PatternAutoPatternAutoIp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternAutoPatternAutoIp to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternAutoPatternAutoIp to JSON text
	ToJson() (string, error)
}

type unMarshalpatternAutoPatternAutoIp struct {
	obj *patternAutoPatternAutoIp
}

type unMarshalPatternAutoPatternAutoIp interface {
	// FromProto unmarshals PatternAutoPatternAutoIp from protobuf object *openapi.PatternAutoPatternAutoIp
	FromProto(msg *openapi.PatternAutoPatternAutoIp) (PatternAutoPatternAutoIp, error)
	// FromPbText unmarshals PatternAutoPatternAutoIp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternAutoPatternAutoIp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternAutoPatternAutoIp from JSON text
	FromJson(value string) error
}

func (obj *patternAutoPatternAutoIp) Marshal() marshalPatternAutoPatternAutoIp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternAutoPatternAutoIp{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternAutoPatternAutoIp) Unmarshal() unMarshalPatternAutoPatternAutoIp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternAutoPatternAutoIp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternAutoPatternAutoIp) ToProto() (*openapi.PatternAutoPatternAutoIp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternAutoPatternAutoIp) FromProto(msg *openapi.PatternAutoPatternAutoIp) (PatternAutoPatternAutoIp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternAutoPatternAutoIp) ToPbText() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIp) FromPbText(value string) error {
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

func (m *marshalpatternAutoPatternAutoIp) ToYaml() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIp) FromYaml(value string) error {
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

func (m *marshalpatternAutoPatternAutoIp) ToJson() (string, error) {
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

func (m *unMarshalpatternAutoPatternAutoIp) FromJson(value string) error {
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

func (obj *patternAutoPatternAutoIp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternAutoPatternAutoIp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternAutoPatternAutoIp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternAutoPatternAutoIp) Clone() (PatternAutoPatternAutoIp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternAutoPatternAutoIp()
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

func (obj *patternAutoPatternAutoIp) setNil() {
	obj.autoHolder = nil
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternAutoPatternAutoIp is tBD
type PatternAutoPatternAutoIp interface {
	Validation
	// msg marshals PatternAutoPatternAutoIp to protobuf object *openapi.PatternAutoPatternAutoIp
	// and doesn't set defaults
	msg() *openapi.PatternAutoPatternAutoIp
	// setMsg unmarshals PatternAutoPatternAutoIp from protobuf object *openapi.PatternAutoPatternAutoIp
	// and doesn't set defaults
	setMsg(*openapi.PatternAutoPatternAutoIp) PatternAutoPatternAutoIp
	// provides marshal interface
	Marshal() marshalPatternAutoPatternAutoIp
	// provides unmarshal interface
	Unmarshal() unMarshalPatternAutoPatternAutoIp
	// validate validates PatternAutoPatternAutoIp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternAutoPatternAutoIp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternAutoPatternAutoIpChoiceEnum, set in PatternAutoPatternAutoIp
	Choice() PatternAutoPatternAutoIpChoiceEnum
	// setChoice assigns PatternAutoPatternAutoIpChoiceEnum provided by user to PatternAutoPatternAutoIp
	setChoice(value PatternAutoPatternAutoIpChoiceEnum) PatternAutoPatternAutoIp
	// HasChoice checks if Choice has been set in PatternAutoPatternAutoIp
	HasChoice() bool
	// Value returns string, set in PatternAutoPatternAutoIp.
	Value() string
	// SetValue assigns string provided by user to PatternAutoPatternAutoIp
	SetValue(value string) PatternAutoPatternAutoIp
	// HasValue checks if Value has been set in PatternAutoPatternAutoIp
	HasValue() bool
	// Values returns []string, set in PatternAutoPatternAutoIp.
	Values() []string
	// SetValues assigns []string provided by user to PatternAutoPatternAutoIp
	SetValues(value []string) PatternAutoPatternAutoIp
	// Auto returns AutoIpOptions, set in PatternAutoPatternAutoIp.
	// AutoIpOptions is the OTG implementation can provide a system generated,
	// value for this property. If the OTG is unable to generate a value,
	// the default value must be used.
	Auto() AutoIpOptions
	// HasAuto checks if Auto has been set in PatternAutoPatternAutoIp
	HasAuto() bool
	// Increment returns PatternAutoPatternAutoIpCounter, set in PatternAutoPatternAutoIp.
	// PatternAutoPatternAutoIpCounter is ipv4 counter pattern
	Increment() PatternAutoPatternAutoIpCounter
	// SetIncrement assigns PatternAutoPatternAutoIpCounter provided by user to PatternAutoPatternAutoIp.
	// PatternAutoPatternAutoIpCounter is ipv4 counter pattern
	SetIncrement(value PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIp
	// HasIncrement checks if Increment has been set in PatternAutoPatternAutoIp
	HasIncrement() bool
	// Decrement returns PatternAutoPatternAutoIpCounter, set in PatternAutoPatternAutoIp.
	// PatternAutoPatternAutoIpCounter is ipv4 counter pattern
	Decrement() PatternAutoPatternAutoIpCounter
	// SetDecrement assigns PatternAutoPatternAutoIpCounter provided by user to PatternAutoPatternAutoIp.
	// PatternAutoPatternAutoIpCounter is ipv4 counter pattern
	SetDecrement(value PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIp
	// HasDecrement checks if Decrement has been set in PatternAutoPatternAutoIp
	HasDecrement() bool
	setNil()
}

type PatternAutoPatternAutoIpChoiceEnum string

// Enum of Choice on PatternAutoPatternAutoIp
var PatternAutoPatternAutoIpChoice = struct {
	VALUE     PatternAutoPatternAutoIpChoiceEnum
	VALUES    PatternAutoPatternAutoIpChoiceEnum
	AUTO      PatternAutoPatternAutoIpChoiceEnum
	INCREMENT PatternAutoPatternAutoIpChoiceEnum
	DECREMENT PatternAutoPatternAutoIpChoiceEnum
}{
	VALUE:     PatternAutoPatternAutoIpChoiceEnum("value"),
	VALUES:    PatternAutoPatternAutoIpChoiceEnum("values"),
	AUTO:      PatternAutoPatternAutoIpChoiceEnum("auto"),
	INCREMENT: PatternAutoPatternAutoIpChoiceEnum("increment"),
	DECREMENT: PatternAutoPatternAutoIpChoiceEnum("decrement"),
}

func (obj *patternAutoPatternAutoIp) Choice() PatternAutoPatternAutoIpChoiceEnum {
	return PatternAutoPatternAutoIpChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternAutoPatternAutoIp) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternAutoPatternAutoIp) setChoice(value PatternAutoPatternAutoIpChoiceEnum) PatternAutoPatternAutoIp {
	intValue, ok := openapi.PatternAutoPatternAutoIp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternAutoPatternAutoIpChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternAutoPatternAutoIp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.autoHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternAutoPatternAutoIpChoice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternAutoPatternAutoIpChoice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternAutoPatternAutoIpChoice.AUTO {
		obj.obj.Auto = NewAutoIpOptions().msg()
	}

	if value == PatternAutoPatternAutoIpChoice.INCREMENT {
		obj.obj.Increment = NewPatternAutoPatternAutoIpCounter().msg()
	}

	if value == PatternAutoPatternAutoIpChoice.DECREMENT {
		obj.obj.Decrement = NewPatternAutoPatternAutoIpCounter().msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternAutoPatternAutoIp) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternAutoPatternAutoIpChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternAutoPatternAutoIp) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternAutoPatternAutoIp object
func (obj *patternAutoPatternAutoIp) SetValue(value string) PatternAutoPatternAutoIp {
	obj.setChoice(PatternAutoPatternAutoIpChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternAutoPatternAutoIp) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternAutoPatternAutoIp object
func (obj *patternAutoPatternAutoIp) SetValues(value []string) PatternAutoPatternAutoIp {
	obj.setChoice(PatternAutoPatternAutoIpChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Auto returns a AutoIpOptions
func (obj *patternAutoPatternAutoIp) Auto() AutoIpOptions {
	if obj.obj.Auto == nil {
		obj.setChoice(PatternAutoPatternAutoIpChoice.AUTO)
	}
	if obj.autoHolder == nil {
		obj.autoHolder = &autoIpOptions{obj: obj.obj.Auto}
	}
	return obj.autoHolder
}

// description is TBD
// Auto returns a AutoIpOptions
func (obj *patternAutoPatternAutoIp) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternAutoPatternAutoIpCounter
func (obj *patternAutoPatternAutoIp) Increment() PatternAutoPatternAutoIpCounter {
	if obj.obj.Increment == nil {
		obj.setChoice(PatternAutoPatternAutoIpChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternAutoPatternAutoIpCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternAutoPatternAutoIpCounter
func (obj *patternAutoPatternAutoIp) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternAutoPatternAutoIpCounter value in the PatternAutoPatternAutoIp object
func (obj *patternAutoPatternAutoIp) SetIncrement(value PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIp {
	obj.setChoice(PatternAutoPatternAutoIpChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.msg()

	return obj
}

// description is TBD
// Decrement returns a PatternAutoPatternAutoIpCounter
func (obj *patternAutoPatternAutoIp) Decrement() PatternAutoPatternAutoIpCounter {
	if obj.obj.Decrement == nil {
		obj.setChoice(PatternAutoPatternAutoIpChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternAutoPatternAutoIpCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternAutoPatternAutoIpCounter
func (obj *patternAutoPatternAutoIp) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternAutoPatternAutoIpCounter value in the PatternAutoPatternAutoIp object
func (obj *patternAutoPatternAutoIp) SetDecrement(value PatternAutoPatternAutoIpCounter) PatternAutoPatternAutoIp {
	obj.setChoice(PatternAutoPatternAutoIpChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.msg()

	return obj
}

func (obj *patternAutoPatternAutoIp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternAutoIp.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternAutoPatternAutoIp.Values"))
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

func (obj *patternAutoPatternAutoIp) setDefault() {
	var choices_set int = 0
	var choice PatternAutoPatternAutoIpChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternAutoPatternAutoIpChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternAutoPatternAutoIpChoice.VALUES
	}

	if obj.obj.Auto != nil {
		choices_set += 1
		choice = PatternAutoPatternAutoIpChoice.AUTO
	}

	if obj.obj.Increment != nil {
		choices_set += 1
		choice = PatternAutoPatternAutoIpChoice.INCREMENT
	}

	if obj.obj.Decrement != nil {
		choices_set += 1
		choice = PatternAutoPatternAutoIpChoice.DECREMENT
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternAutoPatternAutoIpChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternAutoPatternAutoIp")
			}
		} else {
			intVal := openapi.PatternAutoPatternAutoIp_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternAutoPatternAutoIp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
