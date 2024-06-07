package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ZObject *****
type zObject struct {
	validation
	obj          *openapi.ZObject
	marshaller   marshalZObject
	unMarshaller unMarshalZObject
}

func NewZObject() ZObject {
	obj := zObject{obj: &openapi.ZObject{}}
	obj.setDefault()
	return &obj
}

func (obj *zObject) msg() *openapi.ZObject {
	return obj.obj
}

func (obj *zObject) setMsg(msg *openapi.ZObject) ZObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalzObject struct {
	obj *zObject
}

type marshalZObject interface {
	// ToProto marshals ZObject to protobuf object *openapi.ZObject
	ToProto() (*openapi.ZObject, error)
	// ToPbText marshals ZObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ZObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals ZObject to JSON text
	ToJson() (string, error)
}

type unMarshalzObject struct {
	obj *zObject
}

type unMarshalZObject interface {
	// FromProto unmarshals ZObject from protobuf object *openapi.ZObject
	FromProto(msg *openapi.ZObject) (ZObject, error)
	// FromPbText unmarshals ZObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ZObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ZObject from JSON text
	FromJson(value string) error
}

func (obj *zObject) Marshal() marshalZObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalzObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *zObject) Unmarshal() unMarshalZObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalzObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalzObject) ToProto() (*openapi.ZObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalzObject) FromProto(msg *openapi.ZObject) (ZObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalzObject) ToPbText() (string, error) {
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

func (m *unMarshalzObject) FromPbText(value string) error {
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

func (m *marshalzObject) ToYaml() (string, error) {
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

func (m *unMarshalzObject) FromYaml(value string) error {
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

func (m *marshalzObject) ToJson() (string, error) {
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

func (m *unMarshalzObject) FromJson(value string) error {
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

func (obj *zObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *zObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *zObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *zObject) Clone() (ZObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewZObject()
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

// ZObject is description is TBD
type ZObject interface {
	Validation
	// msg marshals ZObject to protobuf object *openapi.ZObject
	// and doesn't set defaults
	msg() *openapi.ZObject
	// setMsg unmarshals ZObject from protobuf object *openapi.ZObject
	// and doesn't set defaults
	setMsg(*openapi.ZObject) ZObject
	// provides marshal interface
	Marshal() marshalZObject
	// provides unmarshal interface
	Unmarshal() unMarshalZObject
	// validate validates ZObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ZObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in ZObject.
	Name() string
	// SetName assigns string provided by user to ZObject
	SetName(value string) ZObject
}

// description is TBD
// Name returns a string
func (obj *zObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// SetName sets the string value in the ZObject object
func (obj *zObject) SetName(value string) ZObject {

	obj.obj.Name = &value
	return obj
}

func (obj *zObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface ZObject")
	}
}

func (obj *zObject) setDefault() {

}
