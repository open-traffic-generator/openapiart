package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternChecksumPatternChecksum *****
type patternChecksumPatternChecksum struct {
	validation
	obj          *openapi.PatternChecksumPatternChecksum
	marshaller   marshalPatternChecksumPatternChecksum
	unMarshaller unMarshalPatternChecksumPatternChecksum
}

func NewPatternChecksumPatternChecksum() PatternChecksumPatternChecksum {
	obj := patternChecksumPatternChecksum{obj: &openapi.PatternChecksumPatternChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternChecksumPatternChecksum) msg() *openapi.PatternChecksumPatternChecksum {
	return obj.obj
}

func (obj *patternChecksumPatternChecksum) setMsg(msg *openapi.PatternChecksumPatternChecksum) PatternChecksumPatternChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternChecksumPatternChecksum struct {
	obj *patternChecksumPatternChecksum
}

type marshalPatternChecksumPatternChecksum interface {
	// ToProto marshals PatternChecksumPatternChecksum to protobuf object *openapi.PatternChecksumPatternChecksum
	ToProto() (*openapi.PatternChecksumPatternChecksum, error)
	// ToPbText marshals PatternChecksumPatternChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternChecksumPatternChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternChecksumPatternChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternChecksumPatternChecksum struct {
	obj *patternChecksumPatternChecksum
}

type unMarshalPatternChecksumPatternChecksum interface {
	// FromProto unmarshals PatternChecksumPatternChecksum from protobuf object *openapi.PatternChecksumPatternChecksum
	FromProto(msg *openapi.PatternChecksumPatternChecksum) (PatternChecksumPatternChecksum, error)
	// FromPbText unmarshals PatternChecksumPatternChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternChecksumPatternChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternChecksumPatternChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternChecksumPatternChecksum) Marshal() marshalPatternChecksumPatternChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternChecksumPatternChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternChecksumPatternChecksum) Unmarshal() unMarshalPatternChecksumPatternChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternChecksumPatternChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternChecksumPatternChecksum) ToProto() (*openapi.PatternChecksumPatternChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternChecksumPatternChecksum) FromProto(msg *openapi.PatternChecksumPatternChecksum) (PatternChecksumPatternChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternChecksumPatternChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternChecksumPatternChecksum) FromPbText(value string) error {
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

func (m *marshalpatternChecksumPatternChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternChecksumPatternChecksum) FromYaml(value string) error {
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

func (m *marshalpatternChecksumPatternChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternChecksumPatternChecksum) FromJson(value string) error {
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

func (obj *patternChecksumPatternChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternChecksumPatternChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternChecksumPatternChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternChecksumPatternChecksum) Clone() (PatternChecksumPatternChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternChecksumPatternChecksum()
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

// PatternChecksumPatternChecksum is tBD
type PatternChecksumPatternChecksum interface {
	Validation
	// msg marshals PatternChecksumPatternChecksum to protobuf object *openapi.PatternChecksumPatternChecksum
	// and doesn't set defaults
	msg() *openapi.PatternChecksumPatternChecksum
	// setMsg unmarshals PatternChecksumPatternChecksum from protobuf object *openapi.PatternChecksumPatternChecksum
	// and doesn't set defaults
	setMsg(*openapi.PatternChecksumPatternChecksum) PatternChecksumPatternChecksum
	// provides marshal interface
	Marshal() marshalPatternChecksumPatternChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternChecksumPatternChecksum
	// validate validates PatternChecksumPatternChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternChecksumPatternChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternChecksumPatternChecksumChoiceEnum, set in PatternChecksumPatternChecksum
	Choice() PatternChecksumPatternChecksumChoiceEnum
	// setChoice assigns PatternChecksumPatternChecksumChoiceEnum provided by user to PatternChecksumPatternChecksum
	setChoice(value PatternChecksumPatternChecksumChoiceEnum) PatternChecksumPatternChecksum
	// HasChoice checks if Choice has been set in PatternChecksumPatternChecksum
	HasChoice() bool
	// Generated returns PatternChecksumPatternChecksumGeneratedEnum, set in PatternChecksumPatternChecksum
	Generated() PatternChecksumPatternChecksumGeneratedEnum
	// SetGenerated assigns PatternChecksumPatternChecksumGeneratedEnum provided by user to PatternChecksumPatternChecksum
	SetGenerated(value PatternChecksumPatternChecksumGeneratedEnum) PatternChecksumPatternChecksum
	// HasGenerated checks if Generated has been set in PatternChecksumPatternChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternChecksumPatternChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternChecksumPatternChecksum
	SetCustom(value uint32) PatternChecksumPatternChecksum
	// HasCustom checks if Custom has been set in PatternChecksumPatternChecksum
	HasCustom() bool
}

type PatternChecksumPatternChecksumChoiceEnum string

// Enum of Choice on PatternChecksumPatternChecksum
var PatternChecksumPatternChecksumChoice = struct {
	GENERATED PatternChecksumPatternChecksumChoiceEnum
	CUSTOM    PatternChecksumPatternChecksumChoiceEnum
}{
	GENERATED: PatternChecksumPatternChecksumChoiceEnum("generated"),
	CUSTOM:    PatternChecksumPatternChecksumChoiceEnum("custom"),
}

func (obj *patternChecksumPatternChecksum) Choice() PatternChecksumPatternChecksumChoiceEnum {
	return PatternChecksumPatternChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternChecksumPatternChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternChecksumPatternChecksum) setChoice(value PatternChecksumPatternChecksumChoiceEnum) PatternChecksumPatternChecksum {
	intValue, ok := openapi.PatternChecksumPatternChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = openapi.PatternChecksumPatternChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternChecksumPatternChecksumGeneratedEnum string

// Enum of Generated on PatternChecksumPatternChecksum
var PatternChecksumPatternChecksumGenerated = struct {
	GOOD PatternChecksumPatternChecksumGeneratedEnum
	BAD  PatternChecksumPatternChecksumGeneratedEnum
}{
	GOOD: PatternChecksumPatternChecksumGeneratedEnum("good"),
	BAD:  PatternChecksumPatternChecksumGeneratedEnum("bad"),
}

func (obj *patternChecksumPatternChecksum) Generated() PatternChecksumPatternChecksumGeneratedEnum {
	return PatternChecksumPatternChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternChecksumPatternChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternChecksumPatternChecksum) SetGenerated(value PatternChecksumPatternChecksumGeneratedEnum) PatternChecksumPatternChecksum {
	intValue, ok := openapi.PatternChecksumPatternChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternChecksumPatternChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternChecksumPatternChecksum object
func (obj *patternChecksumPatternChecksum) SetCustom(value uint32) PatternChecksumPatternChecksum {
	obj.setChoice(PatternChecksumPatternChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternChecksumPatternChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternChecksumPatternChecksum.Custom <= 255 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternChecksumPatternChecksum) setDefault() {
	var choices_set int = 0
	var choice PatternChecksumPatternChecksumChoiceEnum

	if obj.obj.Generated != nil && obj.obj.Generated.Number() != 0 {
		choices_set += 1
		choice = PatternChecksumPatternChecksumChoice.GENERATED
	}

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = PatternChecksumPatternChecksumChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(PatternChecksumPatternChecksumChoice.GENERATED)
			if obj.obj.Generated.Number() == 0 {
				obj.SetGenerated(PatternChecksumPatternChecksumGenerated.GOOD)

			}

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in PatternChecksumPatternChecksum")
			}
		} else {
			intVal := openapi.PatternChecksumPatternChecksum_Choice_Enum_value[string(choice)]
			enumValue := openapi.PatternChecksumPatternChecksum_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
