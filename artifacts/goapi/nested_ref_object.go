package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** NestedRefObject *****
type nestedRefObject struct {
	validation
	obj                    *openapi.NestedRefObject
	marshaller             marshalNestedRefObject
	unMarshaller           unMarshalNestedRefObject
	intermediateNodeHolder IntermediateRefObject
}

func NewNestedRefObject() NestedRefObject {
	obj := nestedRefObject{obj: &openapi.NestedRefObject{}}
	obj.setDefault()
	return &obj
}

func (obj *nestedRefObject) msg() *openapi.NestedRefObject {
	return obj.obj
}

func (obj *nestedRefObject) setMsg(msg *openapi.NestedRefObject) NestedRefObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalnestedRefObject struct {
	obj *nestedRefObject
}

type marshalNestedRefObject interface {
	// ToProto marshals NestedRefObject to protobuf object *openapi.NestedRefObject
	ToProto() (*openapi.NestedRefObject, error)
	// ToPbText marshals NestedRefObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals NestedRefObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals NestedRefObject to JSON text
	ToJson() (string, error)
}

type unMarshalnestedRefObject struct {
	obj *nestedRefObject
}

type unMarshalNestedRefObject interface {
	// FromProto unmarshals NestedRefObject from protobuf object *openapi.NestedRefObject
	FromProto(msg *openapi.NestedRefObject) (NestedRefObject, error)
	// FromPbText unmarshals NestedRefObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals NestedRefObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals NestedRefObject from JSON text
	FromJson(value string) error
}

func (obj *nestedRefObject) Marshal() marshalNestedRefObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalnestedRefObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *nestedRefObject) Unmarshal() unMarshalNestedRefObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalnestedRefObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalnestedRefObject) ToProto() (*openapi.NestedRefObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalnestedRefObject) FromProto(msg *openapi.NestedRefObject) (NestedRefObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalnestedRefObject) ToPbText() (string, error) {
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

func (m *unMarshalnestedRefObject) FromPbText(value string) error {
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

func (m *marshalnestedRefObject) ToYaml() (string, error) {
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

func (m *unMarshalnestedRefObject) FromYaml(value string) error {
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

func (m *marshalnestedRefObject) ToJson() (string, error) {
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

func (m *unMarshalnestedRefObject) FromJson(value string) error {
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

func (obj *nestedRefObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *nestedRefObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *nestedRefObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *nestedRefObject) Clone() (NestedRefObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNestedRefObject()
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

func (obj *nestedRefObject) setNil() {
	obj.intermediateNodeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// NestedRefObject is description is TBD
type NestedRefObject interface {
	Validation
	// msg marshals NestedRefObject to protobuf object *openapi.NestedRefObject
	// and doesn't set defaults
	msg() *openapi.NestedRefObject
	// setMsg unmarshals NestedRefObject from protobuf object *openapi.NestedRefObject
	// and doesn't set defaults
	setMsg(*openapi.NestedRefObject) NestedRefObject
	// provides marshal interface
	Marshal() marshalNestedRefObject
	// provides unmarshal interface
	Unmarshal() unMarshalNestedRefObject
	// validate validates NestedRefObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (NestedRefObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in NestedRefObject.
	Name() string
	// SetName assigns string provided by user to NestedRefObject
	SetName(value string) NestedRefObject
	// HasName checks if Name has been set in NestedRefObject
	HasName() bool
	// IntermediateNode returns IntermediateRefObject, set in NestedRefObject.
	// IntermediateRefObject is description is TBD
	IntermediateNode() IntermediateRefObject
	// SetIntermediateNode assigns IntermediateRefObject provided by user to NestedRefObject.
	// IntermediateRefObject is description is TBD
	SetIntermediateNode(value IntermediateRefObject) NestedRefObject
	// HasIntermediateNode checks if IntermediateNode has been set in NestedRefObject
	HasIntermediateNode() bool
	setNil()
}

// description is TBD
// Name returns a string
func (obj *nestedRefObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *nestedRefObject) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the NestedRefObject object
func (obj *nestedRefObject) SetName(value string) NestedRefObject {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// IntermediateNode returns a IntermediateRefObject
func (obj *nestedRefObject) IntermediateNode() IntermediateRefObject {
	if obj.obj.IntermediateNode == nil {
		obj.obj.IntermediateNode = NewIntermediateRefObject().msg()
	}
	if obj.intermediateNodeHolder == nil {
		obj.intermediateNodeHolder = &intermediateRefObject{obj: obj.obj.IntermediateNode}
	}
	return obj.intermediateNodeHolder
}

// description is TBD
// IntermediateNode returns a IntermediateRefObject
func (obj *nestedRefObject) HasIntermediateNode() bool {
	return obj.obj.IntermediateNode != nil
}

// description is TBD
// SetIntermediateNode sets the IntermediateRefObject value in the NestedRefObject object
func (obj *nestedRefObject) SetIntermediateNode(value IntermediateRefObject) NestedRefObject {

	obj.intermediateNodeHolder = nil
	obj.obj.IntermediateNode = value.msg()

	return obj
}

func (obj *nestedRefObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntermediateNode != nil {

		obj.IntermediateNode().validateObj(vObj, set_default)
	}

}

func (obj *nestedRefObject) setDefault() {

}
