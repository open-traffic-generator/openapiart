package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternPrefixConfigHeaderChecksum *****
type patternPrefixConfigHeaderChecksum struct {
	validation
	obj          *openapi.PatternPrefixConfigHeaderChecksum
	marshaller   marshalPatternPrefixConfigHeaderChecksum
	unMarshaller unMarshalPatternPrefixConfigHeaderChecksum
}

func NewPatternPrefixConfigHeaderChecksum() PatternPrefixConfigHeaderChecksum {
	obj := patternPrefixConfigHeaderChecksum{obj: &openapi.PatternPrefixConfigHeaderChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigHeaderChecksum) msg() *openapi.PatternPrefixConfigHeaderChecksum {
	return obj.obj
}

func (obj *patternPrefixConfigHeaderChecksum) setMsg(msg *openapi.PatternPrefixConfigHeaderChecksum) PatternPrefixConfigHeaderChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternPrefixConfigHeaderChecksum struct {
	obj *patternPrefixConfigHeaderChecksum
}

type marshalPatternPrefixConfigHeaderChecksum interface {
	// ToProto marshals PatternPrefixConfigHeaderChecksum to protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	ToProto() (*openapi.PatternPrefixConfigHeaderChecksum, error)
	// ToPbText marshals PatternPrefixConfigHeaderChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigHeaderChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigHeaderChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternPrefixConfigHeaderChecksum struct {
	obj *patternPrefixConfigHeaderChecksum
}

type unMarshalPatternPrefixConfigHeaderChecksum interface {
	// FromProto unmarshals PatternPrefixConfigHeaderChecksum from protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	FromProto(msg *openapi.PatternPrefixConfigHeaderChecksum) (PatternPrefixConfigHeaderChecksum, error)
	// FromPbText unmarshals PatternPrefixConfigHeaderChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigHeaderChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigHeaderChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternPrefixConfigHeaderChecksum) Marshal() marshalPatternPrefixConfigHeaderChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternPrefixConfigHeaderChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternPrefixConfigHeaderChecksum) Unmarshal() unMarshalPatternPrefixConfigHeaderChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternPrefixConfigHeaderChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternPrefixConfigHeaderChecksum) ToProto() (*openapi.PatternPrefixConfigHeaderChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternPrefixConfigHeaderChecksum) FromProto(msg *openapi.PatternPrefixConfigHeaderChecksum) (PatternPrefixConfigHeaderChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternPrefixConfigHeaderChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternPrefixConfigHeaderChecksum) FromPbText(value string) error {
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

func (m *marshalpatternPrefixConfigHeaderChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternPrefixConfigHeaderChecksum) FromYaml(value string) error {
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

func (m *marshalpatternPrefixConfigHeaderChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternPrefixConfigHeaderChecksum) FromJson(value string) error {
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

func (obj *patternPrefixConfigHeaderChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigHeaderChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigHeaderChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigHeaderChecksum) Clone() (PatternPrefixConfigHeaderChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigHeaderChecksum()
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

// PatternPrefixConfigHeaderChecksum is header checksum
type PatternPrefixConfigHeaderChecksum interface {
	Validation
	// msg marshals PatternPrefixConfigHeaderChecksum to protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	// and doesn't set defaults
	msg() *openapi.PatternPrefixConfigHeaderChecksum
	// setMsg unmarshals PatternPrefixConfigHeaderChecksum from protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	// and doesn't set defaults
	setMsg(*openapi.PatternPrefixConfigHeaderChecksum) PatternPrefixConfigHeaderChecksum
	// provides marshal interface
	Marshal() marshalPatternPrefixConfigHeaderChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternPrefixConfigHeaderChecksum
	// validate validates PatternPrefixConfigHeaderChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigHeaderChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternPrefixConfigHeaderChecksumChoiceEnum, set in PatternPrefixConfigHeaderChecksum
	Choice() PatternPrefixConfigHeaderChecksumChoiceEnum
	// setChoice assigns PatternPrefixConfigHeaderChecksumChoiceEnum provided by user to PatternPrefixConfigHeaderChecksum
	setChoice(value PatternPrefixConfigHeaderChecksumChoiceEnum) PatternPrefixConfigHeaderChecksum
	// HasChoice checks if Choice has been set in PatternPrefixConfigHeaderChecksum
	HasChoice() bool
	// Generated returns PatternPrefixConfigHeaderChecksumGeneratedEnum, set in PatternPrefixConfigHeaderChecksum
	Generated() PatternPrefixConfigHeaderChecksumGeneratedEnum
	// SetGenerated assigns PatternPrefixConfigHeaderChecksumGeneratedEnum provided by user to PatternPrefixConfigHeaderChecksum
	SetGenerated(value PatternPrefixConfigHeaderChecksumGeneratedEnum) PatternPrefixConfigHeaderChecksum
	// HasGenerated checks if Generated has been set in PatternPrefixConfigHeaderChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternPrefixConfigHeaderChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternPrefixConfigHeaderChecksum
	SetCustom(value uint32) PatternPrefixConfigHeaderChecksum
	// HasCustom checks if Custom has been set in PatternPrefixConfigHeaderChecksum
	HasCustom() bool
}

type PatternPrefixConfigHeaderChecksumChoiceEnum string

// Enum of Choice on PatternPrefixConfigHeaderChecksum
var PatternPrefixConfigHeaderChecksumChoice = struct {
	GENERATED PatternPrefixConfigHeaderChecksumChoiceEnum
	CUSTOM    PatternPrefixConfigHeaderChecksumChoiceEnum
}{
	GENERATED: PatternPrefixConfigHeaderChecksumChoiceEnum("generated"),
	CUSTOM:    PatternPrefixConfigHeaderChecksumChoiceEnum("custom"),
}

func (obj *patternPrefixConfigHeaderChecksum) Choice() PatternPrefixConfigHeaderChecksumChoiceEnum {
	return PatternPrefixConfigHeaderChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternPrefixConfigHeaderChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternPrefixConfigHeaderChecksum) setChoice(value PatternPrefixConfigHeaderChecksumChoiceEnum) PatternPrefixConfigHeaderChecksum {
	intValue, ok := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigHeaderChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = openapi.PatternPrefixConfigHeaderChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternPrefixConfigHeaderChecksumGeneratedEnum string

// Enum of Generated on PatternPrefixConfigHeaderChecksum
var PatternPrefixConfigHeaderChecksumGenerated = struct {
	GOOD PatternPrefixConfigHeaderChecksumGeneratedEnum
	BAD  PatternPrefixConfigHeaderChecksumGeneratedEnum
}{
	GOOD: PatternPrefixConfigHeaderChecksumGeneratedEnum("good"),
	BAD:  PatternPrefixConfigHeaderChecksumGeneratedEnum("bad"),
}

func (obj *patternPrefixConfigHeaderChecksum) Generated() PatternPrefixConfigHeaderChecksumGeneratedEnum {
	return PatternPrefixConfigHeaderChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternPrefixConfigHeaderChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternPrefixConfigHeaderChecksum) SetGenerated(value PatternPrefixConfigHeaderChecksumGeneratedEnum) PatternPrefixConfigHeaderChecksum {
	intValue, ok := openapi.PatternPrefixConfigHeaderChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigHeaderChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigHeaderChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternPrefixConfigHeaderChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternPrefixConfigHeaderChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternPrefixConfigHeaderChecksum object
func (obj *patternPrefixConfigHeaderChecksum) SetCustom(value uint32) PatternPrefixConfigHeaderChecksum {
	obj.setChoice(PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternPrefixConfigHeaderChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigHeaderChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternPrefixConfigHeaderChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternPrefixConfigHeaderChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternPrefixConfigHeaderChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternPrefixConfigHeaderChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternPrefixConfigHeaderChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternPrefixConfigHeaderChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternPrefixConfigHeaderChecksum")
			}
		} else {
			intVal := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
