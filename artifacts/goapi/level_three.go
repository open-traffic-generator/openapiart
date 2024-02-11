package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LevelThree *****
type levelThree struct {
	validation
	obj          *openapi.LevelThree
	marshaller   marshalLevelThree
	unMarshaller unMarshalLevelThree
}

func NewLevelThree() LevelThree {
	obj := levelThree{obj: &openapi.LevelThree{}}
	obj.setDefault()
	return &obj
}

func (obj *levelThree) msg() *openapi.LevelThree {
	return obj.obj
}

func (obj *levelThree) setMsg(msg *openapi.LevelThree) LevelThree {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallevelThree struct {
	obj *levelThree
}

type marshalLevelThree interface {
	// ToProto marshals LevelThree to protobuf object *openapi.LevelThree
	ToProto() (*openapi.LevelThree, error)
	// ToPbText marshals LevelThree to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelThree to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelThree to JSON text
	ToJson() (string, error)
}

type unMarshallevelThree struct {
	obj *levelThree
}

type unMarshalLevelThree interface {
	// FromProto unmarshals LevelThree from protobuf object *openapi.LevelThree
	FromProto(msg *openapi.LevelThree) (LevelThree, error)
	// FromPbText unmarshals LevelThree from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelThree from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelThree from JSON text
	FromJson(value string) error
}

func (obj *levelThree) Marshal() marshalLevelThree {
	if obj.marshaller == nil {
		obj.marshaller = &marshallevelThree{obj: obj}
	}
	return obj.marshaller
}

func (obj *levelThree) Unmarshal() unMarshalLevelThree {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallevelThree{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallevelThree) ToProto() (*openapi.LevelThree, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallevelThree) FromProto(msg *openapi.LevelThree) (LevelThree, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallevelThree) ToPbText() (string, error) {
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

func (m *unMarshallevelThree) FromPbText(value string) error {
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

func (m *marshallevelThree) ToYaml() (string, error) {
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

func (m *unMarshallevelThree) FromYaml(value string) error {
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

func (m *marshallevelThree) ToJson() (string, error) {
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

func (m *unMarshallevelThree) FromJson(value string) error {
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

func (obj *levelThree) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelThree) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelThree) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelThree) Clone() (LevelThree, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelThree()
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

// LevelThree is test Level3
type LevelThree interface {
	Validation
	// msg marshals LevelThree to protobuf object *openapi.LevelThree
	// and doesn't set defaults
	msg() *openapi.LevelThree
	// setMsg unmarshals LevelThree from protobuf object *openapi.LevelThree
	// and doesn't set defaults
	setMsg(*openapi.LevelThree) LevelThree
	// provides marshal interface
	Marshal() marshalLevelThree
	// provides unmarshal interface
	Unmarshal() unMarshalLevelThree
	// validate validates LevelThree
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelThree, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L3P1 returns string, set in LevelThree.
	L3P1() string
	// SetL3P1 assigns string provided by user to LevelThree
	SetL3P1(value string) LevelThree
	// HasL3P1 checks if L3P1 has been set in LevelThree
	HasL3P1() bool
}

// Set value at Level 3
// L3P1 returns a string
func (obj *levelThree) L3P1() string {

	return *obj.obj.L3P1

}

// Set value at Level 3
// L3P1 returns a string
func (obj *levelThree) HasL3P1() bool {
	return obj.obj.L3P1 != nil
}

// Set value at Level 3
// SetL3P1 sets the string value in the LevelThree object
func (obj *levelThree) SetL3P1(value string) LevelThree {

	obj.obj.L3P1 = &value
	return obj
}

func (obj *levelThree) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *levelThree) setDefault() {

}
