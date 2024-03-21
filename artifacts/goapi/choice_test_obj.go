package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChoiceTestObj *****
type choiceTestObj struct {
	validation
	obj          *openapi.ChoiceTestObj
	marshaller   marshalChoiceTestObj
	unMarshaller unMarshalChoiceTestObj
	eObjHolder   EObject
	fObjHolder   FObject
}

func NewChoiceTestObj() ChoiceTestObj {
	obj := choiceTestObj{obj: &openapi.ChoiceTestObj{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceTestObj) msg() *openapi.ChoiceTestObj {
	return obj.obj
}

func (obj *choiceTestObj) setMsg(msg *openapi.ChoiceTestObj) ChoiceTestObj {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchoiceTestObj struct {
	obj *choiceTestObj
}

type marshalChoiceTestObj interface {
	// ToProto marshals ChoiceTestObj to protobuf object *openapi.ChoiceTestObj
	ToProto() (*openapi.ChoiceTestObj, error)
	// ToPbText marshals ChoiceTestObj to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceTestObj to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceTestObj to JSON text
	ToJson() (string, error)
}

type unMarshalchoiceTestObj struct {
	obj *choiceTestObj
}

type unMarshalChoiceTestObj interface {
	// FromProto unmarshals ChoiceTestObj from protobuf object *openapi.ChoiceTestObj
	FromProto(msg *openapi.ChoiceTestObj) (ChoiceTestObj, error)
	// FromPbText unmarshals ChoiceTestObj from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceTestObj from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceTestObj from JSON text
	FromJson(value string) error
}

func (obj *choiceTestObj) Marshal() marshalChoiceTestObj {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchoiceTestObj{obj: obj}
	}
	return obj.marshaller
}

func (obj *choiceTestObj) Unmarshal() unMarshalChoiceTestObj {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchoiceTestObj{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchoiceTestObj) ToProto() (*openapi.ChoiceTestObj, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchoiceTestObj) FromProto(msg *openapi.ChoiceTestObj) (ChoiceTestObj, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchoiceTestObj) ToPbText() (string, error) {
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

func (m *unMarshalchoiceTestObj) FromPbText(value string) error {
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

func (m *marshalchoiceTestObj) ToYaml() (string, error) {
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

func (m *unMarshalchoiceTestObj) FromYaml(value string) error {
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

func (m *marshalchoiceTestObj) ToJson() (string, error) {
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

func (m *unMarshalchoiceTestObj) FromJson(value string) error {
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

func (obj *choiceTestObj) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceTestObj) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceTestObj) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceTestObj) Clone() (ChoiceTestObj, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceTestObj()
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

func (obj *choiceTestObj) setNil() {
	obj.eObjHolder = nil
	obj.fObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChoiceTestObj is description is TBD
type ChoiceTestObj interface {
	Validation
	// msg marshals ChoiceTestObj to protobuf object *openapi.ChoiceTestObj
	// and doesn't set defaults
	msg() *openapi.ChoiceTestObj
	// setMsg unmarshals ChoiceTestObj from protobuf object *openapi.ChoiceTestObj
	// and doesn't set defaults
	setMsg(*openapi.ChoiceTestObj) ChoiceTestObj
	// provides marshal interface
	Marshal() marshalChoiceTestObj
	// provides unmarshal interface
	Unmarshal() unMarshalChoiceTestObj
	// validate validates ChoiceTestObj
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceTestObj, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ChoiceTestObjChoiceEnum, set in ChoiceTestObj
	Choice() ChoiceTestObjChoiceEnum
	// setChoice assigns ChoiceTestObjChoiceEnum provided by user to ChoiceTestObj
	setChoice(value ChoiceTestObjChoiceEnum) ChoiceTestObj
	// HasChoice checks if Choice has been set in ChoiceTestObj
	HasChoice() bool
	// getter for NoObj to set choice.
	NoObj()
	// EObj returns EObject, set in ChoiceTestObj.
	// EObject is description is TBD
	EObj() EObject
	// SetEObj assigns EObject provided by user to ChoiceTestObj.
	// EObject is description is TBD
	SetEObj(value EObject) ChoiceTestObj
	// HasEObj checks if EObj has been set in ChoiceTestObj
	HasEObj() bool
	// FObj returns FObject, set in ChoiceTestObj.
	// FObject is description is TBD
	FObj() FObject
	// SetFObj assigns FObject provided by user to ChoiceTestObj.
	// FObject is description is TBD
	SetFObj(value FObject) ChoiceTestObj
	// HasFObj checks if FObj has been set in ChoiceTestObj
	HasFObj() bool
	// Ieee8021Qbb returns string, set in ChoiceTestObj.
	Ieee8021Qbb() string
	// SetIeee8021Qbb assigns string provided by user to ChoiceTestObj
	SetIeee8021Qbb(value string) ChoiceTestObj
	// HasIeee8021Qbb checks if Ieee8021Qbb has been set in ChoiceTestObj
	HasIeee8021Qbb() bool
	// Ieee8023X returns string, set in ChoiceTestObj.
	Ieee8023X() string
	// SetIeee8023X assigns string provided by user to ChoiceTestObj
	SetIeee8023X(value string) ChoiceTestObj
	// HasIeee8023X checks if Ieee8023X has been set in ChoiceTestObj
	HasIeee8023X() bool
	setNil()
}

type ChoiceTestObjChoiceEnum string

// Enum of Choice on ChoiceTestObj
var ChoiceTestObjChoice = struct {
	E_OBJ         ChoiceTestObjChoiceEnum
	F_OBJ         ChoiceTestObjChoiceEnum
	NO_OBJ        ChoiceTestObjChoiceEnum
	IEEE_802_1QBB ChoiceTestObjChoiceEnum
	IEEE_802_3X   ChoiceTestObjChoiceEnum
}{
	E_OBJ:         ChoiceTestObjChoiceEnum("e_obj"),
	F_OBJ:         ChoiceTestObjChoiceEnum("f_obj"),
	NO_OBJ:        ChoiceTestObjChoiceEnum("no_obj"),
	IEEE_802_1QBB: ChoiceTestObjChoiceEnum("ieee_802_1qbb"),
	IEEE_802_3X:   ChoiceTestObjChoiceEnum("ieee_802_3x"),
}

func (obj *choiceTestObj) Choice() ChoiceTestObjChoiceEnum {
	return ChoiceTestObjChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NoObj to set choice
func (obj *choiceTestObj) NoObj() {
	obj.setChoice(ChoiceTestObjChoice.NO_OBJ)
}

// description is TBD
// Choice returns a string
func (obj *choiceTestObj) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *choiceTestObj) setChoice(value ChoiceTestObjChoiceEnum) ChoiceTestObj {
	intValue, ok := openapi.ChoiceTestObj_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ChoiceTestObjChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.ChoiceTestObj_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ieee_802_3X = nil
	obj.obj.Ieee_802_1Qbb = nil
	obj.obj.FObj = nil
	obj.fObjHolder = nil
	obj.obj.EObj = nil
	obj.eObjHolder = nil

	if value == ChoiceTestObjChoice.E_OBJ {
		obj.obj.EObj = NewEObject().msg()
	}

	if value == ChoiceTestObjChoice.F_OBJ {
		obj.obj.FObj = NewFObject().msg()
	}

	return obj
}

// description is TBD
// EObj returns a EObject
func (obj *choiceTestObj) EObj() EObject {
	if obj.obj.EObj == nil {
		obj.setChoice(ChoiceTestObjChoice.E_OBJ)
	}
	if obj.eObjHolder == nil {
		obj.eObjHolder = &eObject{obj: obj.obj.EObj}
	}
	return obj.eObjHolder
}

// description is TBD
// EObj returns a EObject
func (obj *choiceTestObj) HasEObj() bool {
	return obj.obj.EObj != nil
}

// description is TBD
// SetEObj sets the EObject value in the ChoiceTestObj object
func (obj *choiceTestObj) SetEObj(value EObject) ChoiceTestObj {
	obj.setChoice(ChoiceTestObjChoice.E_OBJ)
	obj.eObjHolder = nil
	obj.obj.EObj = value.msg()

	return obj
}

// description is TBD
// FObj returns a FObject
func (obj *choiceTestObj) FObj() FObject {
	if obj.obj.FObj == nil {
		obj.setChoice(ChoiceTestObjChoice.F_OBJ)
	}
	if obj.fObjHolder == nil {
		obj.fObjHolder = &fObject{obj: obj.obj.FObj}
	}
	return obj.fObjHolder
}

// description is TBD
// FObj returns a FObject
func (obj *choiceTestObj) HasFObj() bool {
	return obj.obj.FObj != nil
}

// description is TBD
// SetFObj sets the FObject value in the ChoiceTestObj object
func (obj *choiceTestObj) SetFObj(value FObject) ChoiceTestObj {
	obj.setChoice(ChoiceTestObjChoice.F_OBJ)
	obj.fObjHolder = nil
	obj.obj.FObj = value.msg()

	return obj
}

// description is TBD
// Ieee8021Qbb returns a string
func (obj *choiceTestObj) Ieee8021Qbb() string {

	if obj.obj.Ieee_802_1Qbb == nil {
		obj.setChoice(ChoiceTestObjChoice.IEEE_802_1QBB)
	}

	return *obj.obj.Ieee_802_1Qbb

}

// description is TBD
// Ieee8021Qbb returns a string
func (obj *choiceTestObj) HasIeee8021Qbb() bool {
	return obj.obj.Ieee_802_1Qbb != nil
}

// description is TBD
// SetIeee8021Qbb sets the string value in the ChoiceTestObj object
func (obj *choiceTestObj) SetIeee8021Qbb(value string) ChoiceTestObj {
	obj.setChoice(ChoiceTestObjChoice.IEEE_802_1QBB)
	obj.obj.Ieee_802_1Qbb = &value
	return obj
}

// description is TBD
// Ieee8023X returns a string
func (obj *choiceTestObj) Ieee8023X() string {

	if obj.obj.Ieee_802_3X == nil {
		obj.setChoice(ChoiceTestObjChoice.IEEE_802_3X)
	}

	return *obj.obj.Ieee_802_3X

}

// description is TBD
// Ieee8023X returns a string
func (obj *choiceTestObj) HasIeee8023X() bool {
	return obj.obj.Ieee_802_3X != nil
}

// description is TBD
// SetIeee8023X sets the string value in the ChoiceTestObj object
func (obj *choiceTestObj) SetIeee8023X(value string) ChoiceTestObj {
	obj.setChoice(ChoiceTestObjChoice.IEEE_802_3X)
	obj.obj.Ieee_802_3X = &value
	return obj
}

func (obj *choiceTestObj) validateObj(vObj *validation, set_default bool) {
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

func (obj *choiceTestObj) setDefault() {
	var choices_set int = 0
	var choice ChoiceTestObjChoiceEnum

	if obj.obj.EObj != nil {
		choices_set += 1
		choice = ChoiceTestObjChoice.E_OBJ
	}

	if obj.obj.FObj != nil {
		choices_set += 1
		choice = ChoiceTestObjChoice.F_OBJ
	}

	if obj.obj.Ieee_802_1Qbb != nil {
		choices_set += 1
		choice = ChoiceTestObjChoice.IEEE_802_1QBB
	}

	if obj.obj.Ieee_802_3X != nil {
		choices_set += 1
		choice = ChoiceTestObjChoice.IEEE_802_3X
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ChoiceTestObj")
			}
		} else {
			intVal := openapi.ChoiceTestObj_Choice_Enum_value[string(choice)]
			enumValue := openapi.ChoiceTestObj_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
