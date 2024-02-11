package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PatternChecksumPatternObjectChecksum *****
type patternChecksumPatternObjectChecksum struct {
	validation
	obj          *openapi.PatternChecksumPatternObjectChecksum
	marshaller   marshalPatternChecksumPatternObjectChecksum
	unMarshaller unMarshalPatternChecksumPatternObjectChecksum
}

func NewPatternChecksumPatternObjectChecksum() PatternChecksumPatternObjectChecksum {
	obj := patternChecksumPatternObjectChecksum{obj: &openapi.PatternChecksumPatternObjectChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternChecksumPatternObjectChecksum) msg() *openapi.PatternChecksumPatternObjectChecksum {
	return obj.obj
}

func (obj *patternChecksumPatternObjectChecksum) setMsg(msg *openapi.PatternChecksumPatternObjectChecksum) PatternChecksumPatternObjectChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalpatternChecksumPatternObjectChecksum struct {
	obj *patternChecksumPatternObjectChecksum
}

type marshalPatternChecksumPatternObjectChecksum interface {
	// ToProto marshals PatternChecksumPatternObjectChecksum to protobuf object *openapi.PatternChecksumPatternObjectChecksum
	ToProto() (*openapi.PatternChecksumPatternObjectChecksum, error)
	// ToPbText marshals PatternChecksumPatternObjectChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternChecksumPatternObjectChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternChecksumPatternObjectChecksum to JSON text
	ToJson() (string, error)
}

type unMarshalpatternChecksumPatternObjectChecksum struct {
	obj *patternChecksumPatternObjectChecksum
}

type unMarshalPatternChecksumPatternObjectChecksum interface {
	// FromProto unmarshals PatternChecksumPatternObjectChecksum from protobuf object *openapi.PatternChecksumPatternObjectChecksum
	FromProto(msg *openapi.PatternChecksumPatternObjectChecksum) (PatternChecksumPatternObjectChecksum, error)
	// FromPbText unmarshals PatternChecksumPatternObjectChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternChecksumPatternObjectChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternChecksumPatternObjectChecksum from JSON text
	FromJson(value string) error
}

func (obj *patternChecksumPatternObjectChecksum) Marshal() marshalPatternChecksumPatternObjectChecksum {
	if obj.marshaller == nil {
		obj.marshaller = &marshalpatternChecksumPatternObjectChecksum{obj: obj}
	}
	return obj.marshaller
}

func (obj *patternChecksumPatternObjectChecksum) Unmarshal() unMarshalPatternChecksumPatternObjectChecksum {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalpatternChecksumPatternObjectChecksum{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalpatternChecksumPatternObjectChecksum) ToProto() (*openapi.PatternChecksumPatternObjectChecksum, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalpatternChecksumPatternObjectChecksum) FromProto(msg *openapi.PatternChecksumPatternObjectChecksum) (PatternChecksumPatternObjectChecksum, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalpatternChecksumPatternObjectChecksum) ToPbText() (string, error) {
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

func (m *unMarshalpatternChecksumPatternObjectChecksum) FromPbText(value string) error {
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

func (m *marshalpatternChecksumPatternObjectChecksum) ToYaml() (string, error) {
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

func (m *unMarshalpatternChecksumPatternObjectChecksum) FromYaml(value string) error {
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

func (m *marshalpatternChecksumPatternObjectChecksum) ToJson() (string, error) {
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

func (m *unMarshalpatternChecksumPatternObjectChecksum) FromJson(value string) error {
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

func (obj *patternChecksumPatternObjectChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternChecksumPatternObjectChecksum) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternChecksumPatternObjectChecksum) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternChecksumPatternObjectChecksum) Clone() (PatternChecksumPatternObjectChecksum, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternChecksumPatternObjectChecksum()
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

// PatternChecksumPatternObjectChecksum is tBD
type PatternChecksumPatternObjectChecksum interface {
	Validation
	// msg marshals PatternChecksumPatternObjectChecksum to protobuf object *openapi.PatternChecksumPatternObjectChecksum
	// and doesn't set defaults
	msg() *openapi.PatternChecksumPatternObjectChecksum
	// setMsg unmarshals PatternChecksumPatternObjectChecksum from protobuf object *openapi.PatternChecksumPatternObjectChecksum
	// and doesn't set defaults
	setMsg(*openapi.PatternChecksumPatternObjectChecksum) PatternChecksumPatternObjectChecksum
	// provides marshal interface
	Marshal() marshalPatternChecksumPatternObjectChecksum
	// provides unmarshal interface
	Unmarshal() unMarshalPatternChecksumPatternObjectChecksum
	// validate validates PatternChecksumPatternObjectChecksum
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternChecksumPatternObjectChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternChecksumPatternObjectChecksumChoiceEnum, set in PatternChecksumPatternObjectChecksum
	Choice() PatternChecksumPatternObjectChecksumChoiceEnum
	// setChoice assigns PatternChecksumPatternObjectChecksumChoiceEnum provided by user to PatternChecksumPatternObjectChecksum
	setChoice(value PatternChecksumPatternObjectChecksumChoiceEnum) PatternChecksumPatternObjectChecksum
	// HasChoice checks if Choice has been set in PatternChecksumPatternObjectChecksum
	HasChoice() bool
	// Generated returns PatternChecksumPatternObjectChecksumGeneratedEnum, set in PatternChecksumPatternObjectChecksum
	Generated() PatternChecksumPatternObjectChecksumGeneratedEnum
	// SetGenerated assigns PatternChecksumPatternObjectChecksumGeneratedEnum provided by user to PatternChecksumPatternObjectChecksum
	SetGenerated(value PatternChecksumPatternObjectChecksumGeneratedEnum) PatternChecksumPatternObjectChecksum
	// HasGenerated checks if Generated has been set in PatternChecksumPatternObjectChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternChecksumPatternObjectChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternChecksumPatternObjectChecksum
	SetCustom(value uint32) PatternChecksumPatternObjectChecksum
	// HasCustom checks if Custom has been set in PatternChecksumPatternObjectChecksum
	HasCustom() bool
}

type PatternChecksumPatternObjectChecksumChoiceEnum string

// Enum of Choice on PatternChecksumPatternObjectChecksum
var PatternChecksumPatternObjectChecksumChoice = struct {
	GENERATED PatternChecksumPatternObjectChecksumChoiceEnum
	CUSTOM    PatternChecksumPatternObjectChecksumChoiceEnum
}{
	GENERATED: PatternChecksumPatternObjectChecksumChoiceEnum("generated"),
	CUSTOM:    PatternChecksumPatternObjectChecksumChoiceEnum("custom"),
}

func (obj *patternChecksumPatternObjectChecksum) Choice() PatternChecksumPatternObjectChecksumChoiceEnum {
	return PatternChecksumPatternObjectChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternChecksumPatternObjectChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternChecksumPatternObjectChecksum) setChoice(value PatternChecksumPatternObjectChecksumChoiceEnum) PatternChecksumPatternObjectChecksum {
	intValue, ok := openapi.PatternChecksumPatternObjectChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternObjectChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternObjectChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = openapi.PatternChecksumPatternObjectChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternChecksumPatternObjectChecksumGeneratedEnum string

// Enum of Generated on PatternChecksumPatternObjectChecksum
var PatternChecksumPatternObjectChecksumGenerated = struct {
	GOOD PatternChecksumPatternObjectChecksumGeneratedEnum
	BAD  PatternChecksumPatternObjectChecksumGeneratedEnum
}{
	GOOD: PatternChecksumPatternObjectChecksumGeneratedEnum("good"),
	BAD:  PatternChecksumPatternObjectChecksumGeneratedEnum("bad"),
}

func (obj *patternChecksumPatternObjectChecksum) Generated() PatternChecksumPatternObjectChecksumGeneratedEnum {
	return PatternChecksumPatternObjectChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternChecksumPatternObjectChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternChecksumPatternObjectChecksum) SetGenerated(value PatternChecksumPatternObjectChecksumGeneratedEnum) PatternChecksumPatternObjectChecksum {
	intValue, ok := openapi.PatternChecksumPatternObjectChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternObjectChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternObjectChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternObjectChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.setChoice(PatternChecksumPatternObjectChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternObjectChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternChecksumPatternObjectChecksum object
func (obj *patternChecksumPatternObjectChecksum) SetCustom(value uint32) PatternChecksumPatternObjectChecksum {
	obj.setChoice(PatternChecksumPatternObjectChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternChecksumPatternObjectChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternChecksumPatternObjectChecksum.Custom <= 255 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternChecksumPatternObjectChecksum) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(PatternChecksumPatternObjectChecksumChoice.GENERATED)
		if obj.obj.Generated.Number() == 0 {
			obj.SetGenerated(PatternChecksumPatternObjectChecksumGenerated.GOOD)

		}

	}

}
