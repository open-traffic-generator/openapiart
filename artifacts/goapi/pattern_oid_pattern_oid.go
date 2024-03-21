package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternOidPatternOid *****
type patternOidPatternOid struct {
	validation
	obj          *openapi.PatternOidPatternOid
	marshaller   marshalPatternOidPatternOid
	unMarshaller unMarshalPatternOidPatternOid
}

func NewPatternOidPatternOid() PatternOidPatternOid {
	obj := patternOidPatternOid{obj: &openapi.PatternOidPatternOid{}}
	obj.setDefault()
	return &obj
}

func (obj *patternOidPatternOid) msg() *openapi.PatternOidPatternOid {
	return obj.obj
}

func (obj *patternOidPatternOid) setMsg(msg *openapi.PatternOidPatternOid) PatternOidPatternOid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternOidPatternOid struct {
	obj *patternOidPatternOid
}

type marshalPatternOidPatternOid interface {
	// ToProto marshals PatternOidPatternOid to protobuf object *openapi.PatternOidPatternOid
	ToProto() (*openapi.PatternOidPatternOid, error)
	// ToPbText marshals PatternOidPatternOid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternOidPatternOid to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternOidPatternOid to JSON text
	ToJson() (string, error)
}

type unMarshalpatternOidPatternOid struct {
	obj *patternOidPatternOid
}

type unMarshalPatternOidPatternOid interface {
	// FromProto unmarshals PatternOidPatternOid from protobuf object *openapi.PatternOidPatternOid
	FromProto(msg *openapi.PatternOidPatternOid) (PatternOidPatternOid, error)
	// FromPbText unmarshals PatternOidPatternOid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternOidPatternOid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternOidPatternOid from JSON text
	FromJson(value string) error
}

func (obj *patternOidPatternOid) Marshal() marshalPatternOidPatternOid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternOidPatternOid{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternOidPatternOid) Unmarshal() unMarshalPatternOidPatternOid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternOidPatternOid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternOidPatternOid) ToProto() (*openapi.PatternOidPatternOid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternOidPatternOid) FromProto(msg *openapi.PatternOidPatternOid) (PatternOidPatternOid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternOidPatternOid) ToPbText() (string, error) {
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

func (m *unMarshalpatternOidPatternOid) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalpatternOidPatternOid) ToYaml() (string, error) {
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

func (m *unMarshalpatternOidPatternOid) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalpatternOidPatternOid) ToJson() (string, error) {
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

func (m *unMarshalpatternOidPatternOid) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternOidPatternOid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternOidPatternOid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternOidPatternOid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternOidPatternOid) Clone() (PatternOidPatternOid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternOidPatternOid()
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

// PatternOidPatternOid is tBD
type PatternOidPatternOid interface {
	Validation
	// msg marshals PatternOidPatternOid to protobuf object *openapi.PatternOidPatternOid
	// and doesn't set defaults
	msg() *openapi.PatternOidPatternOid
	// setMsg unmarshals PatternOidPatternOid from protobuf object *openapi.PatternOidPatternOid
	// and doesn't set defaults
	setMsg(*openapi.PatternOidPatternOid) PatternOidPatternOid
	// provides marshal interface
	Marshal() marshalPatternOidPatternOid
	// provides unmarshal interface
	Unmarshal() unMarshalPatternOidPatternOid
	// validate validates PatternOidPatternOid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternOidPatternOid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternOidPatternOidChoiceEnum, set in PatternOidPatternOid
	Choice() PatternOidPatternOidChoiceEnum
	// setChoice assigns PatternOidPatternOidChoiceEnum provided by user to PatternOidPatternOid
	setChoice(value PatternOidPatternOidChoiceEnum) PatternOidPatternOid
	// HasChoice checks if Choice has been set in PatternOidPatternOid
	HasChoice() bool
	// Value returns string, set in PatternOidPatternOid.
	Value() string
	// SetValue assigns string provided by user to PatternOidPatternOid
	SetValue(value string) PatternOidPatternOid
	// HasValue checks if Value has been set in PatternOidPatternOid
	HasValue() bool
	// Values returns []string, set in PatternOidPatternOid.
	Values() []string
	// SetValues assigns []string provided by user to PatternOidPatternOid
	SetValues(value []string) PatternOidPatternOid
}

type PatternOidPatternOidChoiceEnum string

// Enum of Choice on PatternOidPatternOid
var PatternOidPatternOidChoice = struct {
	VALUE  PatternOidPatternOidChoiceEnum
	VALUES PatternOidPatternOidChoiceEnum
}{
	VALUE:  PatternOidPatternOidChoiceEnum("value"),
	VALUES: PatternOidPatternOidChoiceEnum("values"),
}

func (obj *patternOidPatternOid) Choice() PatternOidPatternOidChoiceEnum {
	return PatternOidPatternOidChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternOidPatternOid) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternOidPatternOid) setChoice(value PatternOidPatternOidChoiceEnum) PatternOidPatternOid {
	intValue, ok := openapi.PatternOidPatternOid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternOidPatternOidChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternOidPatternOid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternOidPatternOidChoice.VALUE {
		defaultValue := "0.1"
		obj.obj.Value = &defaultValue
	}

	if value == PatternOidPatternOidChoice.VALUES {
		defaultValue := []string{"0.1"}
		obj.obj.Values = defaultValue
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternOidPatternOid) Value() string {

	if obj.obj.Value == nil {
		obj.setChoice(PatternOidPatternOidChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternOidPatternOid) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternOidPatternOid object
func (obj *patternOidPatternOid) SetValue(value string) PatternOidPatternOid {
	obj.setChoice(PatternOidPatternOidChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternOidPatternOid) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.1"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternOidPatternOid object
func (obj *patternOidPatternOid) SetValues(value []string) PatternOidPatternOid {
	obj.setChoice(PatternOidPatternOidChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

func (obj *patternOidPatternOid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateOid(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternOidPatternOid.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateOidSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternOidPatternOid.Values"))
		}

	}

}

func (obj *patternOidPatternOid) setDefault() {
	var choices_set int = 0
	var choice PatternOidPatternOidChoiceEnum

	if obj.obj.Value != nil {
		choices_set += 1
		choice = PatternOidPatternOidChoice.VALUE
	}

	if len(obj.obj.Values) > 0 {
		choices_set += 1
		choice = PatternOidPatternOidChoice.VALUES
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternOidPatternOidChoice.VALUE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternOidPatternOid")
			}
		} else {
			intVal := openapi.PatternOidPatternOid_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternOidPatternOid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
