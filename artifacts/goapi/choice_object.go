package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChoiceObject *****
type choiceObject struct {
	validation
	obj          *openapi.ChoiceObject
	marshaller   marshalChoiceObject
	unMarshaller unMarshalChoiceObject
	eObjHolder   EObject
	fObjHolder   FObject
}

func NewChoiceObject() ChoiceObject {
	obj := choiceObject{obj: &openapi.ChoiceObject{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceObject) msg() *openapi.ChoiceObject {
	return obj.obj
}

func (obj *choiceObject) setMsg(msg *openapi.ChoiceObject) ChoiceObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchoiceObject struct {
	obj *choiceObject
}

type marshalChoiceObject interface {
	// ToProto marshals ChoiceObject to protobuf object *openapi.ChoiceObject
	ToProto() (*openapi.ChoiceObject, error)
	// ToPbText marshals ChoiceObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceObject to JSON text
	ToJson() (string, error)
}

type unMarshalchoiceObject struct {
	obj *choiceObject
}

type unMarshalChoiceObject interface {
	// FromProto unmarshals ChoiceObject from protobuf object *openapi.ChoiceObject
	FromProto(msg *openapi.ChoiceObject) (ChoiceObject, error)
	// FromPbText unmarshals ChoiceObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceObject from JSON text
	FromJson(value string) error
}

func (obj *choiceObject) Marshal() marshalChoiceObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchoiceObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *choiceObject) Unmarshal() unMarshalChoiceObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchoiceObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchoiceObject) ToProto() (*openapi.ChoiceObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchoiceObject) FromProto(msg *openapi.ChoiceObject) (ChoiceObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchoiceObject) ToPbText() (string, error) {
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

func (m *unMarshalchoiceObject) FromPbText(value string) error {
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

func (m *marshalchoiceObject) ToYaml() (string, error) {
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

func (m *unMarshalchoiceObject) FromYaml(value string) error {
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

func (m *marshalchoiceObject) ToJson() (string, error) {
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

func (m *unMarshalchoiceObject) FromJson(value string) error {
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

func (obj *choiceObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceObject) Clone() (ChoiceObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceObject()
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

func (obj *choiceObject) setNil() {
	obj.eObjHolder = nil
	obj.fObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChoiceObject is description is TBD
type ChoiceObject interface {
	Validation
	// msg marshals ChoiceObject to protobuf object *openapi.ChoiceObject
	// and doesn't set defaults
	msg() *openapi.ChoiceObject
	// setMsg unmarshals ChoiceObject from protobuf object *openapi.ChoiceObject
	// and doesn't set defaults
	setMsg(*openapi.ChoiceObject) ChoiceObject
	// provides marshal interface
	Marshal() marshalChoiceObject
	// provides unmarshal interface
	Unmarshal() unMarshalChoiceObject
	// validate validates ChoiceObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ChoiceObjectChoiceEnum, set in ChoiceObject
	Choice() ChoiceObjectChoiceEnum
	// setChoice assigns ChoiceObjectChoiceEnum provided by user to ChoiceObject
	setChoice(value ChoiceObjectChoiceEnum) ChoiceObject
	// HasChoice checks if Choice has been set in ChoiceObject
	HasChoice() bool
	// getter for NoObj to set choice.
	NoObj()
	// EObj returns EObject, set in ChoiceObject.
	// EObject is description is TBD
	EObj() EObject
	// SetEObj assigns EObject provided by user to ChoiceObject.
	// EObject is description is TBD
	SetEObj(value EObject) ChoiceObject
	// HasEObj checks if EObj has been set in ChoiceObject
	HasEObj() bool
	// FObj returns FObject, set in ChoiceObject.
	// FObject is description is TBD
	FObj() FObject
	// SetFObj assigns FObject provided by user to ChoiceObject.
	// FObject is description is TBD
	SetFObj(value FObject) ChoiceObject
	// HasFObj checks if FObj has been set in ChoiceObject
	HasFObj() bool
	setNil()
}

type ChoiceObjectChoiceEnum string

// Enum of Choice on ChoiceObject
var ChoiceObjectChoice = struct {
	E_OBJ  ChoiceObjectChoiceEnum
	F_OBJ  ChoiceObjectChoiceEnum
	NO_OBJ ChoiceObjectChoiceEnum
}{
	E_OBJ:  ChoiceObjectChoiceEnum("e_obj"),
	F_OBJ:  ChoiceObjectChoiceEnum("f_obj"),
	NO_OBJ: ChoiceObjectChoiceEnum("no_obj"),
}

func (obj *choiceObject) Choice() ChoiceObjectChoiceEnum {
	return ChoiceObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoObj to set choice
func (obj *choiceObject) NoObj() {
	obj.setChoice(ChoiceObjectChoice.NO_OBJ)
}

// description is TBD
// Choice returns a string
func (obj *choiceObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *choiceObject) setChoice(value ChoiceObjectChoiceEnum) ChoiceObject {
	intValue, ok := openapi.ChoiceObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ChoiceObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.ChoiceObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FObj = nil
	obj.fObjHolder = nil
	obj.obj.EObj = nil
	obj.eObjHolder = nil

	if value == ChoiceObjectChoice.E_OBJ {
		obj.obj.EObj = NewEObject().msg()
	}

	if value == ChoiceObjectChoice.F_OBJ {
		obj.obj.FObj = NewFObject().msg()
	}

	return obj
}

// description is TBD
// EObj returns a EObject
func (obj *choiceObject) EObj() EObject {
	if obj.obj.EObj == nil {
		obj.setChoice(ChoiceObjectChoice.E_OBJ)
	}
	if obj.eObjHolder == nil {
		obj.eObjHolder = &eObject{obj: obj.obj.EObj}
	}
	return obj.eObjHolder
}

// description is TBD
// EObj returns a EObject
func (obj *choiceObject) HasEObj() bool {
	return obj.obj.EObj != nil
}

// description is TBD
// SetEObj sets the EObject value in the ChoiceObject object
func (obj *choiceObject) SetEObj(value EObject) ChoiceObject {
	obj.setChoice(ChoiceObjectChoice.E_OBJ)
	obj.eObjHolder = nil
	obj.obj.EObj = value.msg()

	return obj
}

// description is TBD
// FObj returns a FObject
func (obj *choiceObject) FObj() FObject {
	if obj.obj.FObj == nil {
		obj.setChoice(ChoiceObjectChoice.F_OBJ)
	}
	if obj.fObjHolder == nil {
		obj.fObjHolder = &fObject{obj: obj.obj.FObj}
	}
	return obj.fObjHolder
}

// description is TBD
// FObj returns a FObject
func (obj *choiceObject) HasFObj() bool {
	return obj.obj.FObj != nil
}

// description is TBD
// SetFObj sets the FObject value in the ChoiceObject object
func (obj *choiceObject) SetFObj(value FObject) ChoiceObject {
	obj.setChoice(ChoiceObjectChoice.F_OBJ)
	obj.fObjHolder = nil
	obj.obj.FObj = value.msg()

	return obj
}

func (obj *choiceObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EObj != nil {

		obj.EObj().validateObj(vObj, set_default)
	}

	if obj.obj.FObj != nil {

		obj.FObj().validateObj(vObj, set_default)
	}

}

func (obj *choiceObject) setDefault() {
	var choices_set int = 0
	var choice ChoiceObjectChoiceEnum

	if obj.obj.EObj != nil {
		choices_set += 1
		choice = ChoiceObjectChoice.E_OBJ
	}

	if obj.obj.FObj != nil {
		choices_set += 1
		choice = ChoiceObjectChoice.F_OBJ
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(ChoiceObjectChoice.NO_OBJ)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ChoiceObject")
			}
		} else {
			intVal := openapi.ChoiceObject_Choice_Enum_value[string(choice)]
			enumValue := openapi.ChoiceObject_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
