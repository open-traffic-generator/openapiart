package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** WObject *****
type wObject struct {
	validation
	obj          *openapi.WObject
	marshaller   marshalWObject
	unMarshaller unMarshalWObject
}

func NewWObject() WObject {
	obj := wObject{obj: &openapi.WObject{}}
	obj.setDefault()
	return &obj
}

func (obj *wObject) msg() *openapi.WObject {
	return obj.obj
}

func (obj *wObject) setMsg(msg *openapi.WObject) WObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalwObject struct {
	obj *wObject
}

type marshalWObject interface {
	// ToProto marshals WObject to protobuf object *openapi.WObject
	ToProto() (*openapi.WObject, error)
	// ToPbText marshals WObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals WObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals WObject to JSON text
	ToJson() (string, error)
}

type unMarshalwObject struct {
	obj *wObject
}

type unMarshalWObject interface {
	// FromProto unmarshals WObject from protobuf object *openapi.WObject
	FromProto(msg *openapi.WObject) (WObject, error)
	// FromPbText unmarshals WObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals WObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals WObject from JSON text
	FromJson(value string) error
}

func (obj *wObject) Marshal() marshalWObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalwObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *wObject) Unmarshal() unMarshalWObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalwObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalwObject) ToProto() (*openapi.WObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalwObject) FromProto(msg *openapi.WObject) (WObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalwObject) ToPbText() (string, error) {
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

func (m *unMarshalwObject) FromPbText(value string) error {
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

func (m *marshalwObject) ToYaml() (string, error) {
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

func (m *unMarshalwObject) FromYaml(value string) error {
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

func (m *marshalwObject) ToJson() (string, error) {
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

func (m *unMarshalwObject) FromJson(value string) error {
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

func (obj *wObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *wObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *wObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *wObject) Clone() (WObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewWObject()
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

// WObject is description is TBD
type WObject interface {
	Validation
	// msg marshals WObject to protobuf object *openapi.WObject
	// and doesn't set defaults
	msg() *openapi.WObject
	// setMsg unmarshals WObject from protobuf object *openapi.WObject
	// and doesn't set defaults
	setMsg(*openapi.WObject) WObject
	// provides marshal interface
	Marshal() marshalWObject
	// provides unmarshal interface
	Unmarshal() unMarshalWObject
	// validate validates WObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (WObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// WName returns string, set in WObject.
	WName() string
	// SetWName assigns string provided by user to WObject
	SetWName(value string) WObject
}

// description is TBD
// WName returns a string
func (obj *wObject) WName() string {

	return *obj.obj.WName

}

// description is TBD
// SetWName sets the string value in the WObject object
func (obj *wObject) SetWName(value string) WObject {

	obj.obj.WName = &value
	return obj
}

func (obj *wObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// WName is required
	if obj.obj.WName == nil {
		vObj.validationErrors = append(vObj.validationErrors, "WName is required field on interface WObject")
	}
}

func (obj *wObject) setDefault() {

}
