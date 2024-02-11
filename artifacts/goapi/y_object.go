package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** YObject *****
type yObject struct {
	validation
	obj          *openapi.YObject
	marshaller   marshalYObject
	unMarshaller unMarshalYObject
}

func NewYObject() YObject {
	obj := yObject{obj: &openapi.YObject{}}
	obj.setDefault()
	return &obj
}

func (obj *yObject) msg() *openapi.YObject {
	return obj.obj
}

func (obj *yObject) setMsg(msg *openapi.YObject) YObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalyObject struct {
	obj *yObject
}

type marshalYObject interface {
	// ToProto marshals YObject to protobuf object *openapi.YObject
	ToProto() (*openapi.YObject, error)
	// ToPbText marshals YObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals YObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals YObject to JSON text
	ToJson() (string, error)
}

type unMarshalyObject struct {
	obj *yObject
}

type unMarshalYObject interface {
	// FromProto unmarshals YObject from protobuf object *openapi.YObject
	FromProto(msg *openapi.YObject) (YObject, error)
	// FromPbText unmarshals YObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals YObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals YObject from JSON text
	FromJson(value string) error
}

func (obj *yObject) Marshal() marshalYObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalyObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *yObject) Unmarshal() unMarshalYObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalyObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalyObject) ToProto() (*openapi.YObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalyObject) FromProto(msg *openapi.YObject) (YObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalyObject) ToPbText() (string, error) {
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

func (m *unMarshalyObject) FromPbText(value string) error {
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

func (m *marshalyObject) ToYaml() (string, error) {
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

func (m *unMarshalyObject) FromYaml(value string) error {
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

func (m *marshalyObject) ToJson() (string, error) {
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

func (m *unMarshalyObject) FromJson(value string) error {
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

func (obj *yObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *yObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *yObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *yObject) Clone() (YObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewYObject()
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

// YObject is description is TBD
type YObject interface {
	Validation
	// msg marshals YObject to protobuf object *openapi.YObject
	// and doesn't set defaults
	msg() *openapi.YObject
	// setMsg unmarshals YObject from protobuf object *openapi.YObject
	// and doesn't set defaults
	setMsg(*openapi.YObject) YObject
	// provides marshal interface
	Marshal() marshalYObject
	// provides unmarshal interface
	Unmarshal() unMarshalYObject
	// validate validates YObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (YObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// YName returns string, set in YObject.
	YName() string
	// SetYName assigns string provided by user to YObject
	SetYName(value string) YObject
	// HasYName checks if YName has been set in YObject
	HasYName() bool
}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// YName returns a string
func (obj *yObject) YName() string {

	return *obj.obj.YName

}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// YName returns a string
func (obj *yObject) HasYName() bool {
	return obj.obj.YName != nil
}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// SetYName sets the string value in the YObject object
func (obj *yObject) SetYName(value string) YObject {

	obj.obj.YName = &value
	return obj
}

func (obj *yObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *yObject) setDefault() {

}
