package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MixedVal *****
type mixedVal struct {
	validation
	obj          *openapi.MixedVal
	marshaller   marshalMixedVal
	unMarshaller unMarshalMixedVal
}

func NewMixedVal() MixedVal {
	obj := mixedVal{obj: &openapi.MixedVal{}}
	obj.setDefault()
	return &obj
}

func (obj *mixedVal) msg() *openapi.MixedVal {
	return obj.obj
}

func (obj *mixedVal) setMsg(msg *openapi.MixedVal) MixedVal {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmixedVal struct {
	obj *mixedVal
}

type marshalMixedVal interface {
	// ToProto marshals MixedVal to protobuf object *openapi.MixedVal
	ToProto() (*openapi.MixedVal, error)
	// ToPbText marshals MixedVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MixedVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals MixedVal to JSON text
	ToJson() (string, error)
}

type unMarshalmixedVal struct {
	obj *mixedVal
}

type unMarshalMixedVal interface {
	// FromProto unmarshals MixedVal from protobuf object *openapi.MixedVal
	FromProto(msg *openapi.MixedVal) (MixedVal, error)
	// FromPbText unmarshals MixedVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MixedVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MixedVal from JSON text
	FromJson(value string) error
}

func (obj *mixedVal) Marshal() marshalMixedVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmixedVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *mixedVal) Unmarshal() unMarshalMixedVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmixedVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmixedVal) ToProto() (*openapi.MixedVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmixedVal) FromProto(msg *openapi.MixedVal) (MixedVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmixedVal) ToPbText() (string, error) {
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

func (m *unMarshalmixedVal) FromPbText(value string) error {
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

func (m *marshalmixedVal) ToYaml() (string, error) {
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

func (m *unMarshalmixedVal) FromYaml(value string) error {
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

func (m *marshalmixedVal) ToJson() (string, error) {
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

func (m *unMarshalmixedVal) FromJson(value string) error {
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

func (obj *mixedVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mixedVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mixedVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mixedVal) Clone() (MixedVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMixedVal()
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

// MixedVal is description is TBD
type MixedVal interface {
	Validation
	// msg marshals MixedVal to protobuf object *openapi.MixedVal
	// and doesn't set defaults
	msg() *openapi.MixedVal
	// setMsg unmarshals MixedVal from protobuf object *openapi.MixedVal
	// and doesn't set defaults
	setMsg(*openapi.MixedVal) MixedVal
	// provides marshal interface
	Marshal() marshalMixedVal
	// provides unmarshal interface
	Unmarshal() unMarshalMixedVal
	// validate validates MixedVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MixedVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MixedValChoiceEnum, set in MixedVal
	Choice() MixedValChoiceEnum
	// setChoice assigns MixedValChoiceEnum provided by user to MixedVal
	setChoice(value MixedValChoiceEnum) MixedVal
	// HasChoice checks if Choice has been set in MixedVal
	HasChoice() bool
	// IntVal returns int32, set in MixedVal.
	IntVal() int32
	// SetIntVal assigns int32 provided by user to MixedVal
	SetIntVal(value int32) MixedVal
	// HasIntVal checks if IntVal has been set in MixedVal
	HasIntVal() bool
	// NumVal returns float32, set in MixedVal.
	NumVal() float32
	// SetNumVal assigns float32 provided by user to MixedVal
	SetNumVal(value float32) MixedVal
	// HasNumVal checks if NumVal has been set in MixedVal
	HasNumVal() bool
	// StrVal returns string, set in MixedVal.
	StrVal() string
	// SetStrVal assigns string provided by user to MixedVal
	SetStrVal(value string) MixedVal
	// HasStrVal checks if StrVal has been set in MixedVal
	HasStrVal() bool
	// BoolVal returns bool, set in MixedVal.
	BoolVal() bool
	// SetBoolVal assigns bool provided by user to MixedVal
	SetBoolVal(value bool) MixedVal
	// HasBoolVal checks if BoolVal has been set in MixedVal
	HasBoolVal() bool
}

type MixedValChoiceEnum string

// Enum of Choice on MixedVal
var MixedValChoice = struct {
	INT_VAL  MixedValChoiceEnum
	NUM_VAL  MixedValChoiceEnum
	STR_VAL  MixedValChoiceEnum
	BOOL_VAL MixedValChoiceEnum
}{
	INT_VAL:  MixedValChoiceEnum("int_val"),
	NUM_VAL:  MixedValChoiceEnum("num_val"),
	STR_VAL:  MixedValChoiceEnum("str_val"),
	BOOL_VAL: MixedValChoiceEnum("bool_val"),
}

func (obj *mixedVal) Choice() MixedValChoiceEnum {
	return MixedValChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *mixedVal) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *mixedVal) setChoice(value MixedValChoiceEnum) MixedVal {
	intValue, ok := openapi.MixedVal_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MixedValChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.MixedVal_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.BoolVal = nil
	obj.obj.StrVal = nil
	obj.obj.NumVal = nil
	obj.obj.IntVal = nil

	if value == MixedValChoice.INT_VAL {
		defaultValue := int32(50)
		obj.obj.IntVal = &defaultValue
	}

	if value == MixedValChoice.NUM_VAL {
		defaultValue := float32(50.05)
		obj.obj.NumVal = &defaultValue
	}

	if value == MixedValChoice.STR_VAL {
		defaultValue := "default_str_val"
		obj.obj.StrVal = &defaultValue
	}

	if value == MixedValChoice.BOOL_VAL {
		defaultValue := bool(true)
		obj.obj.BoolVal = &defaultValue
	}

	return obj
}

// description is TBD
// IntVal returns a int32
func (obj *mixedVal) IntVal() int32 {

	if obj.obj.IntVal == nil {
		obj.setChoice(MixedValChoice.INT_VAL)
	}

	return *obj.obj.IntVal

}

// description is TBD
// IntVal returns a int32
func (obj *mixedVal) HasIntVal() bool {
	return obj.obj.IntVal != nil
}

// description is TBD
// SetIntVal sets the int32 value in the MixedVal object
func (obj *mixedVal) SetIntVal(value int32) MixedVal {
	obj.setChoice(MixedValChoice.INT_VAL)
	obj.obj.IntVal = &value
	return obj
}

// description is TBD
// NumVal returns a float32
func (obj *mixedVal) NumVal() float32 {

	if obj.obj.NumVal == nil {
		obj.setChoice(MixedValChoice.NUM_VAL)
	}

	return *obj.obj.NumVal

}

// description is TBD
// NumVal returns a float32
func (obj *mixedVal) HasNumVal() bool {
	return obj.obj.NumVal != nil
}

// description is TBD
// SetNumVal sets the float32 value in the MixedVal object
func (obj *mixedVal) SetNumVal(value float32) MixedVal {
	obj.setChoice(MixedValChoice.NUM_VAL)
	obj.obj.NumVal = &value
	return obj
}

// description is TBD
// StrVal returns a string
func (obj *mixedVal) StrVal() string {

	if obj.obj.StrVal == nil {
		obj.setChoice(MixedValChoice.STR_VAL)
	}

	return *obj.obj.StrVal

}

// description is TBD
// StrVal returns a string
func (obj *mixedVal) HasStrVal() bool {
	return obj.obj.StrVal != nil
}

// description is TBD
// SetStrVal sets the string value in the MixedVal object
func (obj *mixedVal) SetStrVal(value string) MixedVal {
	obj.setChoice(MixedValChoice.STR_VAL)
	obj.obj.StrVal = &value
	return obj
}

// description is TBD
// BoolVal returns a bool
func (obj *mixedVal) BoolVal() bool {

	if obj.obj.BoolVal == nil {
		obj.setChoice(MixedValChoice.BOOL_VAL)
	}

	return *obj.obj.BoolVal

}

// description is TBD
// BoolVal returns a bool
func (obj *mixedVal) HasBoolVal() bool {
	return obj.obj.BoolVal != nil
}

// description is TBD
// SetBoolVal sets the bool value in the MixedVal object
func (obj *mixedVal) SetBoolVal(value bool) MixedVal {
	obj.setChoice(MixedValChoice.BOOL_VAL)
	obj.obj.BoolVal = &value
	return obj
}

func (obj *mixedVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *mixedVal) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(MixedValChoice.INT_VAL)

	}

}
