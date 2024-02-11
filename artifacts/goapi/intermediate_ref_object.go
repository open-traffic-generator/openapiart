package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IntermediateRefObject *****
type intermediateRefObject struct {
	validation
	obj            *openapi.IntermediateRefObject
	marshaller     marshalIntermediateRefObject
	unMarshaller   unMarshalIntermediateRefObject
	leafNodeHolder LeafVal
}

func NewIntermediateRefObject() IntermediateRefObject {
	obj := intermediateRefObject{obj: &openapi.IntermediateRefObject{}}
	obj.setDefault()
	return &obj
}

func (obj *intermediateRefObject) msg() *openapi.IntermediateRefObject {
	return obj.obj
}

func (obj *intermediateRefObject) setMsg(msg *openapi.IntermediateRefObject) IntermediateRefObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalintermediateRefObject struct {
	obj *intermediateRefObject
}

type marshalIntermediateRefObject interface {
	// ToProto marshals IntermediateRefObject to protobuf object *openapi.IntermediateRefObject
	ToProto() (*openapi.IntermediateRefObject, error)
	// ToPbText marshals IntermediateRefObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IntermediateRefObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals IntermediateRefObject to JSON text
	ToJson() (string, error)
}

type unMarshalintermediateRefObject struct {
	obj *intermediateRefObject
}

type unMarshalIntermediateRefObject interface {
	// FromProto unmarshals IntermediateRefObject from protobuf object *openapi.IntermediateRefObject
	FromProto(msg *openapi.IntermediateRefObject) (IntermediateRefObject, error)
	// FromPbText unmarshals IntermediateRefObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IntermediateRefObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IntermediateRefObject from JSON text
	FromJson(value string) error
}

func (obj *intermediateRefObject) Marshal() marshalIntermediateRefObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalintermediateRefObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *intermediateRefObject) Unmarshal() unMarshalIntermediateRefObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalintermediateRefObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalintermediateRefObject) ToProto() (*openapi.IntermediateRefObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalintermediateRefObject) FromProto(msg *openapi.IntermediateRefObject) (IntermediateRefObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalintermediateRefObject) ToPbText() (string, error) {
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

func (m *unMarshalintermediateRefObject) FromPbText(value string) error {
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

func (m *marshalintermediateRefObject) ToYaml() (string, error) {
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

func (m *unMarshalintermediateRefObject) FromYaml(value string) error {
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

func (m *marshalintermediateRefObject) ToJson() (string, error) {
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

func (m *unMarshalintermediateRefObject) FromJson(value string) error {
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

func (obj *intermediateRefObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *intermediateRefObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *intermediateRefObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *intermediateRefObject) Clone() (IntermediateRefObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIntermediateRefObject()
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

func (obj *intermediateRefObject) setNil() {
	obj.leafNodeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IntermediateRefObject is description is TBD
type IntermediateRefObject interface {
	Validation
	// msg marshals IntermediateRefObject to protobuf object *openapi.IntermediateRefObject
	// and doesn't set defaults
	msg() *openapi.IntermediateRefObject
	// setMsg unmarshals IntermediateRefObject from protobuf object *openapi.IntermediateRefObject
	// and doesn't set defaults
	setMsg(*openapi.IntermediateRefObject) IntermediateRefObject
	// provides marshal interface
	Marshal() marshalIntermediateRefObject
	// provides unmarshal interface
	Unmarshal() unMarshalIntermediateRefObject
	// validate validates IntermediateRefObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IntermediateRefObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in IntermediateRefObject.
	Name() string
	// SetName assigns string provided by user to IntermediateRefObject
	SetName(value string) IntermediateRefObject
	// HasName checks if Name has been set in IntermediateRefObject
	HasName() bool
	// LeafNode returns LeafVal, set in IntermediateRefObject.
	// LeafVal is description is TBD
	LeafNode() LeafVal
	// SetLeafNode assigns LeafVal provided by user to IntermediateRefObject.
	// LeafVal is description is TBD
	SetLeafNode(value LeafVal) IntermediateRefObject
	// HasLeafNode checks if LeafNode has been set in IntermediateRefObject
	HasLeafNode() bool
	setNil()
}

// description is TBD
// Name returns a string
func (obj *intermediateRefObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *intermediateRefObject) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the IntermediateRefObject object
func (obj *intermediateRefObject) SetName(value string) IntermediateRefObject {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// LeafNode returns a LeafVal
func (obj *intermediateRefObject) LeafNode() LeafVal {
	if obj.obj.LeafNode == nil {
		obj.obj.LeafNode = NewLeafVal().msg()
	}
	if obj.leafNodeHolder == nil {
		obj.leafNodeHolder = &leafVal{obj: obj.obj.LeafNode}
	}
	return obj.leafNodeHolder
}

// description is TBD
// LeafNode returns a LeafVal
func (obj *intermediateRefObject) HasLeafNode() bool {
	return obj.obj.LeafNode != nil
}

// description is TBD
// SetLeafNode sets the LeafVal value in the IntermediateRefObject object
func (obj *intermediateRefObject) SetLeafNode(value LeafVal) IntermediateRefObject {

	obj.leafNodeHolder = nil
	obj.obj.LeafNode = value.msg()

	return obj
}

func (obj *intermediateRefObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LeafNode != nil {

		obj.LeafNode().validateObj(vObj, set_default)
	}

}

func (obj *intermediateRefObject) setDefault() {

}
