package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LevelTwo *****
type levelTwo struct {
	validation
	obj          *openapi.LevelTwo
	marshaller   marshalLevelTwo
	unMarshaller unMarshalLevelTwo
	l2P1Holder   LevelThree
}

func NewLevelTwo() LevelTwo {
	obj := levelTwo{obj: &openapi.LevelTwo{}}
	obj.setDefault()
	return &obj
}

func (obj *levelTwo) msg() *openapi.LevelTwo {
	return obj.obj
}

func (obj *levelTwo) setMsg(msg *openapi.LevelTwo) LevelTwo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallevelTwo struct {
	obj *levelTwo
}

type marshalLevelTwo interface {
	// ToProto marshals LevelTwo to protobuf object *openapi.LevelTwo
	ToProto() (*openapi.LevelTwo, error)
	// ToPbText marshals LevelTwo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelTwo to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelTwo to JSON text
	ToJson() (string, error)
}

type unMarshallevelTwo struct {
	obj *levelTwo
}

type unMarshalLevelTwo interface {
	// FromProto unmarshals LevelTwo from protobuf object *openapi.LevelTwo
	FromProto(msg *openapi.LevelTwo) (LevelTwo, error)
	// FromPbText unmarshals LevelTwo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelTwo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelTwo from JSON text
	FromJson(value string) error
}

func (obj *levelTwo) Marshal() marshalLevelTwo {
	if obj.marshaller == nil {
		obj.marshaller = &marshallevelTwo{obj: obj}
	}
	return obj.marshaller
}

func (obj *levelTwo) Unmarshal() unMarshalLevelTwo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallevelTwo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallevelTwo) ToProto() (*openapi.LevelTwo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallevelTwo) FromProto(msg *openapi.LevelTwo) (LevelTwo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallevelTwo) ToPbText() (string, error) {
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

func (m *unMarshallevelTwo) FromPbText(value string) error {
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

func (m *marshallevelTwo) ToYaml() (string, error) {
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

func (m *unMarshallevelTwo) FromYaml(value string) error {
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

func (m *marshallevelTwo) ToJson() (string, error) {
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

func (m *unMarshallevelTwo) FromJson(value string) error {
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

func (obj *levelTwo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelTwo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelTwo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelTwo) Clone() (LevelTwo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelTwo()
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

func (obj *levelTwo) setNil() {
	obj.l2P1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelTwo is test Level 2
type LevelTwo interface {
	Validation
	// msg marshals LevelTwo to protobuf object *openapi.LevelTwo
	// and doesn't set defaults
	msg() *openapi.LevelTwo
	// setMsg unmarshals LevelTwo from protobuf object *openapi.LevelTwo
	// and doesn't set defaults
	setMsg(*openapi.LevelTwo) LevelTwo
	// provides marshal interface
	Marshal() marshalLevelTwo
	// provides unmarshal interface
	Unmarshal() unMarshalLevelTwo
	// validate validates LevelTwo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelTwo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L2P1 returns LevelThree, set in LevelTwo.
	// LevelThree is test Level3
	L2P1() LevelThree
	// SetL2P1 assigns LevelThree provided by user to LevelTwo.
	// LevelThree is test Level3
	SetL2P1(value LevelThree) LevelTwo
	// HasL2P1 checks if L2P1 has been set in LevelTwo
	HasL2P1() bool
	setNil()
}

// Level Two
// L2P1 returns a LevelThree
func (obj *levelTwo) L2P1() LevelThree {
	if obj.obj.L2P1 == nil {
		obj.obj.L2P1 = NewLevelThree().msg()
	}
	if obj.l2P1Holder == nil {
		obj.l2P1Holder = &levelThree{obj: obj.obj.L2P1}
	}
	return obj.l2P1Holder
}

// Level Two
// L2P1 returns a LevelThree
func (obj *levelTwo) HasL2P1() bool {
	return obj.obj.L2P1 != nil
}

// Level Two
// SetL2P1 sets the LevelThree value in the LevelTwo object
func (obj *levelTwo) SetL2P1(value LevelThree) LevelTwo {

	obj.l2P1Holder = nil
	obj.obj.L2P1 = value.msg()

	return obj
}

func (obj *levelTwo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L2P1 != nil {

		obj.L2P1().validateObj(vObj, set_default)
	}

}

func (obj *levelTwo) setDefault() {

}
