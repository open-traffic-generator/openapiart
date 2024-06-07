package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** JObject *****
type jObject struct {
	validation
	obj          *openapi.JObject
	marshaller   marshalJObject
	unMarshaller unMarshalJObject
	jAHolder     EObject
	jBHolder     FObject
}

func NewJObject() JObject {
	obj := jObject{obj: &openapi.JObject{}}
	obj.setDefault()
	return &obj
}

func (obj *jObject) msg() *openapi.JObject {
	return obj.obj
}

func (obj *jObject) setMsg(msg *openapi.JObject) JObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaljObject struct {
	obj *jObject
}

type marshalJObject interface {
	// ToProto marshals JObject to protobuf object *openapi.JObject
	ToProto() (*openapi.JObject, error)
	// ToPbText marshals JObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals JObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals JObject to JSON text
	ToJson() (string, error)
}

type unMarshaljObject struct {
	obj *jObject
}

type unMarshalJObject interface {
	// FromProto unmarshals JObject from protobuf object *openapi.JObject
	FromProto(msg *openapi.JObject) (JObject, error)
	// FromPbText unmarshals JObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals JObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals JObject from JSON text
	FromJson(value string) error
}

func (obj *jObject) Marshal() marshalJObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshaljObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *jObject) Unmarshal() unMarshalJObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaljObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaljObject) ToProto() (*openapi.JObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaljObject) FromProto(msg *openapi.JObject) (JObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaljObject) ToPbText() (string, error) {
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

func (m *unMarshaljObject) FromPbText(value string) error {
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

func (m *marshaljObject) ToYaml() (string, error) {
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

func (m *unMarshaljObject) FromYaml(value string) error {
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

func (m *marshaljObject) ToJson() (string, error) {
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

func (m *unMarshaljObject) FromJson(value string) error {
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

func (obj *jObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *jObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *jObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *jObject) Clone() (JObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewJObject()
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

func (obj *jObject) setNil() {
	obj.jAHolder = nil
	obj.jBHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// JObject is description is TBD
type JObject interface {
	Validation
	// msg marshals JObject to protobuf object *openapi.JObject
	// and doesn't set defaults
	msg() *openapi.JObject
	// setMsg unmarshals JObject from protobuf object *openapi.JObject
	// and doesn't set defaults
	setMsg(*openapi.JObject) JObject
	// provides marshal interface
	Marshal() marshalJObject
	// provides unmarshal interface
	Unmarshal() unMarshalJObject
	// validate validates JObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (JObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns JObjectChoiceEnum, set in JObject
	Choice() JObjectChoiceEnum
	// setChoice assigns JObjectChoiceEnum provided by user to JObject
	setChoice(value JObjectChoiceEnum) JObject
	// HasChoice checks if Choice has been set in JObject
	HasChoice() bool
	// JA returns EObject, set in JObject.
	// EObject is description is TBD
	JA() EObject
	// SetJA assigns EObject provided by user to JObject.
	// EObject is description is TBD
	SetJA(value EObject) JObject
	// HasJA checks if JA has been set in JObject
	HasJA() bool
	// JB returns FObject, set in JObject.
	// FObject is description is TBD
	JB() FObject
	// SetJB assigns FObject provided by user to JObject.
	// FObject is description is TBD
	SetJB(value FObject) JObject
	// HasJB checks if JB has been set in JObject
	HasJB() bool
	setNil()
}

type JObjectChoiceEnum string

// Enum of Choice on JObject
var JObjectChoice = struct {
	J_A JObjectChoiceEnum
	J_B JObjectChoiceEnum
}{
	J_A: JObjectChoiceEnum("j_a"),
	J_B: JObjectChoiceEnum("j_b"),
}

func (obj *jObject) Choice() JObjectChoiceEnum {
	return JObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *jObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *jObject) setChoice(value JObjectChoiceEnum) JObject {
	intValue, ok := openapi.JObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on JObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.JObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.JB = nil
	obj.jBHolder = nil
	obj.obj.JA = nil
	obj.jAHolder = nil

	if value == JObjectChoice.J_A {
		obj.obj.JA = NewEObject().msg()
	}

	if value == JObjectChoice.J_B {
		obj.obj.JB = NewFObject().msg()
	}

	return obj
}

// description is TBD
// JA returns a EObject
func (obj *jObject) JA() EObject {
	if obj.obj.JA == nil {
		obj.setChoice(JObjectChoice.J_A)
	}
	if obj.jAHolder == nil {
		obj.jAHolder = &eObject{obj: obj.obj.JA}
	}
	return obj.jAHolder
}

// description is TBD
// JA returns a EObject
func (obj *jObject) HasJA() bool {
	return obj.obj.JA != nil
}

// description is TBD
// SetJA sets the EObject value in the JObject object
func (obj *jObject) SetJA(value EObject) JObject {
	obj.setChoice(JObjectChoice.J_A)
	obj.jAHolder = nil
	obj.obj.JA = value.msg()

	return obj
}

// description is TBD
// JB returns a FObject
func (obj *jObject) JB() FObject {
	if obj.obj.JB == nil {
		obj.setChoice(JObjectChoice.J_B)
	}
	if obj.jBHolder == nil {
		obj.jBHolder = &fObject{obj: obj.obj.JB}
	}
	return obj.jBHolder
}

// description is TBD
// JB returns a FObject
func (obj *jObject) HasJB() bool {
	return obj.obj.JB != nil
}

// description is TBD
// SetJB sets the FObject value in the JObject object
func (obj *jObject) SetJB(value FObject) JObject {
	obj.setChoice(JObjectChoice.J_B)
	obj.jBHolder = nil
	obj.obj.JB = value.msg()

	return obj
}

func (obj *jObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Choice.Number() == 2 {
		obj.addWarnings("J_B enum in property Choice is deprecated, use j_a instead")
	}

	if obj.obj.JA != nil {

		obj.JA().validateObj(vObj, set_default)
	}

	if obj.obj.JB != nil {

		obj.JB().validateObj(vObj, set_default)
	}

}

func (obj *jObject) setDefault() {
	var choices_set int = 0
	var choice JObjectChoiceEnum

	if obj.obj.JA != nil {
		choices_set += 1
		choice = JObjectChoice.J_A
	}

	if obj.obj.JB != nil {
		choices_set += 1
		choice = JObjectChoice.J_B
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(JObjectChoice.J_A)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in JObject")
			}
		} else {
			intVal := openapi.JObject_Choice_Enum_value[string(choice)]
			enumValue := openapi.JObject_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
