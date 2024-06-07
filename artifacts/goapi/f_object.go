package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FObject *****
type fObject struct {
	validation
	obj          *openapi.FObject
	marshaller   marshalFObject
	unMarshaller unMarshalFObject
}

func NewFObject() FObject {
	obj := fObject{obj: &openapi.FObject{}}
	obj.setDefault()
	return &obj
}

func (obj *fObject) msg() *openapi.FObject {
	return obj.obj
}

func (obj *fObject) setMsg(msg *openapi.FObject) FObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalfObject struct {
	obj *fObject
}

type marshalFObject interface {
	// ToProto marshals FObject to protobuf object *openapi.FObject
	ToProto() (*openapi.FObject, error)
	// ToPbText marshals FObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals FObject to JSON text
	ToJson() (string, error)
}

type unMarshalfObject struct {
	obj *fObject
}

type unMarshalFObject interface {
	// FromProto unmarshals FObject from protobuf object *openapi.FObject
	FromProto(msg *openapi.FObject) (FObject, error)
	// FromPbText unmarshals FObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FObject from JSON text
	FromJson(value string) error
}

func (obj *fObject) Marshal() marshalFObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalfObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *fObject) Unmarshal() unMarshalFObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalfObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalfObject) ToProto() (*openapi.FObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalfObject) FromProto(msg *openapi.FObject) (FObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalfObject) ToPbText() (string, error) {
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

func (m *unMarshalfObject) FromPbText(value string) error {
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

func (m *marshalfObject) ToYaml() (string, error) {
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

func (m *unMarshalfObject) FromYaml(value string) error {
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

func (m *marshalfObject) ToJson() (string, error) {
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

func (m *unMarshalfObject) FromJson(value string) error {
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

func (obj *fObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *fObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *fObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *fObject) Clone() (FObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFObject()
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

// FObject is description is TBD
type FObject interface {
	Validation
	// msg marshals FObject to protobuf object *openapi.FObject
	// and doesn't set defaults
	msg() *openapi.FObject
	// setMsg unmarshals FObject from protobuf object *openapi.FObject
	// and doesn't set defaults
	setMsg(*openapi.FObject) FObject
	// provides marshal interface
	Marshal() marshalFObject
	// provides unmarshal interface
	Unmarshal() unMarshalFObject
	// validate validates FObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FObjectChoiceEnum, set in FObject
	Choice() FObjectChoiceEnum
	// setChoice assigns FObjectChoiceEnum provided by user to FObject
	setChoice(value FObjectChoiceEnum) FObject
	// HasChoice checks if Choice has been set in FObject
	HasChoice() bool
	// getter for FC to set choice.
	FC()
	// FA returns string, set in FObject.
	FA() string
	// SetFA assigns string provided by user to FObject
	SetFA(value string) FObject
	// HasFA checks if FA has been set in FObject
	HasFA() bool
	// FB returns float64, set in FObject.
	FB() float64
	// SetFB assigns float64 provided by user to FObject
	SetFB(value float64) FObject
	// HasFB checks if FB has been set in FObject
	HasFB() bool
}

type FObjectChoiceEnum string

// Enum of Choice on FObject
var FObjectChoice = struct {
	F_A FObjectChoiceEnum
	F_B FObjectChoiceEnum
	F_C FObjectChoiceEnum
}{
	F_A: FObjectChoiceEnum("f_a"),
	F_B: FObjectChoiceEnum("f_b"),
	F_C: FObjectChoiceEnum("f_c"),
}

func (obj *fObject) Choice() FObjectChoiceEnum {
	return FObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for FC to set choice
func (obj *fObject) FC() {
	obj.setChoice(FObjectChoice.F_C)
}

// description is TBD
// Choice returns a string
func (obj *fObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *fObject) setChoice(value FObjectChoiceEnum) FObject {
	intValue, ok := openapi.FObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.FObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FB = nil
	obj.obj.FA = nil

	if value == FObjectChoice.F_A {
		defaultValue := "some string"
		obj.obj.FA = &defaultValue
	}

	if value == FObjectChoice.F_B {
		defaultValue := float64(3.0)
		obj.obj.FB = &defaultValue
	}

	return obj
}

// description is TBD
// FA returns a string
func (obj *fObject) FA() string {

	if obj.obj.FA == nil {
		obj.setChoice(FObjectChoice.F_A)
	}

	return *obj.obj.FA

}

// description is TBD
// FA returns a string
func (obj *fObject) HasFA() bool {
	return obj.obj.FA != nil
}

// description is TBD
// SetFA sets the string value in the FObject object
func (obj *fObject) SetFA(value string) FObject {
	obj.setChoice(FObjectChoice.F_A)
	obj.obj.FA = &value
	return obj
}

// description is TBD
// FB returns a float64
func (obj *fObject) FB() float64 {

	if obj.obj.FB == nil {
		obj.setChoice(FObjectChoice.F_B)
	}

	return *obj.obj.FB

}

// description is TBD
// FB returns a float64
func (obj *fObject) HasFB() bool {
	return obj.obj.FB != nil
}

// description is TBD
// SetFB sets the float64 value in the FObject object
func (obj *fObject) SetFB(value float64) FObject {
	obj.setChoice(FObjectChoice.F_B)
	obj.obj.FB = &value
	return obj
}

func (obj *fObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *fObject) setDefault() {
	var choices_set int = 0
	var choice FObjectChoiceEnum

	if obj.obj.FA != nil {
		choices_set += 1
		choice = FObjectChoice.F_A
	}

	if obj.obj.FB != nil {
		choices_set += 1
		choice = FObjectChoice.F_B
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FObjectChoice.F_A)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FObject")
			}
		} else {
			intVal := openapi.FObject_Choice_Enum_value[string(choice)]
			enumValue := openapi.FObject_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
