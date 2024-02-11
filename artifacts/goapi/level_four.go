package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LevelFour *****
type levelFour struct {
	validation
	obj          *openapi.LevelFour
	marshaller   marshalLevelFour
	unMarshaller unMarshalLevelFour
	l4P1Holder   LevelOne
}

func NewLevelFour() LevelFour {
	obj := levelFour{obj: &openapi.LevelFour{}}
	obj.setDefault()
	return &obj
}

func (obj *levelFour) msg() *openapi.LevelFour {
	return obj.obj
}

func (obj *levelFour) setMsg(msg *openapi.LevelFour) LevelFour {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallevelFour struct {
	obj *levelFour
}

type marshalLevelFour interface {
	// ToProto marshals LevelFour to protobuf object *openapi.LevelFour
	ToProto() (*openapi.LevelFour, error)
	// ToPbText marshals LevelFour to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelFour to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelFour to JSON text
	ToJson() (string, error)
}

type unMarshallevelFour struct {
	obj *levelFour
}

type unMarshalLevelFour interface {
	// FromProto unmarshals LevelFour from protobuf object *openapi.LevelFour
	FromProto(msg *openapi.LevelFour) (LevelFour, error)
	// FromPbText unmarshals LevelFour from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelFour from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelFour from JSON text
	FromJson(value string) error
}

func (obj *levelFour) Marshal() marshalLevelFour {
	if obj.marshaller == nil {
		obj.marshaller = &marshallevelFour{obj: obj}
	}
	return obj.marshaller
}

func (obj *levelFour) Unmarshal() unMarshalLevelFour {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallevelFour{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallevelFour) ToProto() (*openapi.LevelFour, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallevelFour) FromProto(msg *openapi.LevelFour) (LevelFour, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallevelFour) ToPbText() (string, error) {
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

func (m *unMarshallevelFour) FromPbText(value string) error {
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

func (m *marshallevelFour) ToYaml() (string, error) {
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

func (m *unMarshallevelFour) FromYaml(value string) error {
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

func (m *marshallevelFour) ToJson() (string, error) {
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

func (m *unMarshallevelFour) FromJson(value string) error {
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

func (obj *levelFour) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelFour) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelFour) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelFour) Clone() (LevelFour, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelFour()
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

func (obj *levelFour) setNil() {
	obj.l4P1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelFour is test level4 redundant junk testing
type LevelFour interface {
	Validation
	// msg marshals LevelFour to protobuf object *openapi.LevelFour
	// and doesn't set defaults
	msg() *openapi.LevelFour
	// setMsg unmarshals LevelFour from protobuf object *openapi.LevelFour
	// and doesn't set defaults
	setMsg(*openapi.LevelFour) LevelFour
	// provides marshal interface
	Marshal() marshalLevelFour
	// provides unmarshal interface
	Unmarshal() unMarshalLevelFour
	// validate validates LevelFour
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelFour, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L4P1 returns LevelOne, set in LevelFour.
	// LevelOne is to Test Multi level non-primitive types
	L4P1() LevelOne
	// SetL4P1 assigns LevelOne provided by user to LevelFour.
	// LevelOne is to Test Multi level non-primitive types
	SetL4P1(value LevelOne) LevelFour
	// HasL4P1 checks if L4P1 has been set in LevelFour
	HasL4P1() bool
	setNil()
}

// loop over level 1
// L4P1 returns a LevelOne
func (obj *levelFour) L4P1() LevelOne {
	if obj.obj.L4P1 == nil {
		obj.obj.L4P1 = NewLevelOne().msg()
	}
	if obj.l4P1Holder == nil {
		obj.l4P1Holder = &levelOne{obj: obj.obj.L4P1}
	}
	return obj.l4P1Holder
}

// loop over level 1
// L4P1 returns a LevelOne
func (obj *levelFour) HasL4P1() bool {
	return obj.obj.L4P1 != nil
}

// loop over level 1
// SetL4P1 sets the LevelOne value in the LevelFour object
func (obj *levelFour) SetL4P1(value LevelOne) LevelFour {

	obj.l4P1Holder = nil
	obj.obj.L4P1 = value.msg()

	return obj
}

func (obj *levelFour) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L4P1 != nil {

		obj.L4P1().validateObj(vObj, set_default)
	}

}

func (obj *levelFour) setDefault() {

}
