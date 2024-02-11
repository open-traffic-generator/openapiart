package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LevelOne *****
type levelOne struct {
	validation
	obj          *openapi.LevelOne
	marshaller   marshalLevelOne
	unMarshaller unMarshalLevelOne
	l1P1Holder   LevelTwo
	l1P2Holder   LevelFour
}

func NewLevelOne() LevelOne {
	obj := levelOne{obj: &openapi.LevelOne{}}
	obj.setDefault()
	return &obj
}

func (obj *levelOne) msg() *openapi.LevelOne {
	return obj.obj
}

func (obj *levelOne) setMsg(msg *openapi.LevelOne) LevelOne {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallevelOne struct {
	obj *levelOne
}

type marshalLevelOne interface {
	// ToProto marshals LevelOne to protobuf object *openapi.LevelOne
	ToProto() (*openapi.LevelOne, error)
	// ToPbText marshals LevelOne to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelOne to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelOne to JSON text
	ToJson() (string, error)
}

type unMarshallevelOne struct {
	obj *levelOne
}

type unMarshalLevelOne interface {
	// FromProto unmarshals LevelOne from protobuf object *openapi.LevelOne
	FromProto(msg *openapi.LevelOne) (LevelOne, error)
	// FromPbText unmarshals LevelOne from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelOne from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelOne from JSON text
	FromJson(value string) error
}

func (obj *levelOne) Marshal() marshalLevelOne {
	if obj.marshaller == nil {
		obj.marshaller = &marshallevelOne{obj: obj}
	}
	return obj.marshaller
}

func (obj *levelOne) Unmarshal() unMarshalLevelOne {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallevelOne{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallevelOne) ToProto() (*openapi.LevelOne, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallevelOne) FromProto(msg *openapi.LevelOne) (LevelOne, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallevelOne) ToPbText() (string, error) {
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

func (m *unMarshallevelOne) FromPbText(value string) error {
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

func (m *marshallevelOne) ToYaml() (string, error) {
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

func (m *unMarshallevelOne) FromYaml(value string) error {
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

func (m *marshallevelOne) ToJson() (string, error) {
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

func (m *unMarshallevelOne) FromJson(value string) error {
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

func (obj *levelOne) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelOne) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelOne) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelOne) Clone() (LevelOne, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelOne()
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

func (obj *levelOne) setNil() {
	obj.l1P1Holder = nil
	obj.l1P2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelOne is to Test Multi level non-primitive types
type LevelOne interface {
	Validation
	// msg marshals LevelOne to protobuf object *openapi.LevelOne
	// and doesn't set defaults
	msg() *openapi.LevelOne
	// setMsg unmarshals LevelOne from protobuf object *openapi.LevelOne
	// and doesn't set defaults
	setMsg(*openapi.LevelOne) LevelOne
	// provides marshal interface
	Marshal() marshalLevelOne
	// provides unmarshal interface
	Unmarshal() unMarshalLevelOne
	// validate validates LevelOne
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelOne, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L1P1 returns LevelTwo, set in LevelOne.
	// LevelTwo is test Level 2
	L1P1() LevelTwo
	// SetL1P1 assigns LevelTwo provided by user to LevelOne.
	// LevelTwo is test Level 2
	SetL1P1(value LevelTwo) LevelOne
	// HasL1P1 checks if L1P1 has been set in LevelOne
	HasL1P1() bool
	// L1P2 returns LevelFour, set in LevelOne.
	// LevelFour is test level4 redundant junk testing
	L1P2() LevelFour
	// SetL1P2 assigns LevelFour provided by user to LevelOne.
	// LevelFour is test level4 redundant junk testing
	SetL1P2(value LevelFour) LevelOne
	// HasL1P2 checks if L1P2 has been set in LevelOne
	HasL1P2() bool
	setNil()
}

// Level one
// L1P1 returns a LevelTwo
func (obj *levelOne) L1P1() LevelTwo {
	if obj.obj.L1P1 == nil {
		obj.obj.L1P1 = NewLevelTwo().msg()
	}
	if obj.l1P1Holder == nil {
		obj.l1P1Holder = &levelTwo{obj: obj.obj.L1P1}
	}
	return obj.l1P1Holder
}

// Level one
// L1P1 returns a LevelTwo
func (obj *levelOne) HasL1P1() bool {
	return obj.obj.L1P1 != nil
}

// Level one
// SetL1P1 sets the LevelTwo value in the LevelOne object
func (obj *levelOne) SetL1P1(value LevelTwo) LevelOne {

	obj.l1P1Holder = nil
	obj.obj.L1P1 = value.msg()

	return obj
}

// Level one to four
// L1P2 returns a LevelFour
func (obj *levelOne) L1P2() LevelFour {
	if obj.obj.L1P2 == nil {
		obj.obj.L1P2 = NewLevelFour().msg()
	}
	if obj.l1P2Holder == nil {
		obj.l1P2Holder = &levelFour{obj: obj.obj.L1P2}
	}
	return obj.l1P2Holder
}

// Level one to four
// L1P2 returns a LevelFour
func (obj *levelOne) HasL1P2() bool {
	return obj.obj.L1P2 != nil
}

// Level one to four
// SetL1P2 sets the LevelFour value in the LevelOne object
func (obj *levelOne) SetL1P2(value LevelFour) LevelOne {

	obj.l1P2Holder = nil
	obj.obj.L1P2 = value.msg()

	return obj
}

func (obj *levelOne) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L1P1 != nil {

		obj.L1P1().validateObj(vObj, set_default)
	}

	if obj.obj.L1P2 != nil {

		obj.L1P2().validateObj(vObj, set_default)
	}

}

func (obj *levelOne) setDefault() {

}
