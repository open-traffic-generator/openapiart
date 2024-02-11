package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** XEnumObject *****
type xEnumObject struct {
	validation
	obj          *openapi.XEnumObject
	marshaller   marshalXEnumObject
	unMarshaller unMarshalXEnumObject
}

func NewXEnumObject() XEnumObject {
	obj := xEnumObject{obj: &openapi.XEnumObject{}}
	obj.setDefault()
	return &obj
}

func (obj *xEnumObject) msg() *openapi.XEnumObject {
	return obj.obj
}

func (obj *xEnumObject) setMsg(msg *openapi.XEnumObject) XEnumObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalxEnumObject struct {
	obj *xEnumObject
}

type marshalXEnumObject interface {
	// ToProto marshals XEnumObject to protobuf object *openapi.XEnumObject
	ToProto() (*openapi.XEnumObject, error)
	// ToPbText marshals XEnumObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals XEnumObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals XEnumObject to JSON text
	ToJson() (string, error)
}

type unMarshalxEnumObject struct {
	obj *xEnumObject
}

type unMarshalXEnumObject interface {
	// FromProto unmarshals XEnumObject from protobuf object *openapi.XEnumObject
	FromProto(msg *openapi.XEnumObject) (XEnumObject, error)
	// FromPbText unmarshals XEnumObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals XEnumObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals XEnumObject from JSON text
	FromJson(value string) error
}

func (obj *xEnumObject) Marshal() marshalXEnumObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalxEnumObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *xEnumObject) Unmarshal() unMarshalXEnumObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalxEnumObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalxEnumObject) ToProto() (*openapi.XEnumObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalxEnumObject) FromProto(msg *openapi.XEnumObject) (XEnumObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalxEnumObject) ToPbText() (string, error) {
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

func (m *unMarshalxEnumObject) FromPbText(value string) error {
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

func (m *marshalxEnumObject) ToYaml() (string, error) {
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

func (m *unMarshalxEnumObject) FromYaml(value string) error {
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

func (m *marshalxEnumObject) ToJson() (string, error) {
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

func (m *unMarshalxEnumObject) FromJson(value string) error {
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

func (obj *xEnumObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *xEnumObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *xEnumObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *xEnumObject) Clone() (XEnumObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewXEnumObject()
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

// XEnumObject is description is TBD
type XEnumObject interface {
	Validation
	// msg marshals XEnumObject to protobuf object *openapi.XEnumObject
	// and doesn't set defaults
	msg() *openapi.XEnumObject
	// setMsg unmarshals XEnumObject from protobuf object *openapi.XEnumObject
	// and doesn't set defaults
	setMsg(*openapi.XEnumObject) XEnumObject
	// provides marshal interface
	Marshal() marshalXEnumObject
	// provides unmarshal interface
	Unmarshal() unMarshalXEnumObject
	// validate validates XEnumObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (XEnumObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// XEnumVal returns XEnumObjectXEnumValEnum, set in XEnumObject
	XEnumVal() XEnumObjectXEnumValEnum
	// SetXEnumVal assigns XEnumObjectXEnumValEnum provided by user to XEnumObject
	SetXEnumVal(value XEnumObjectXEnumValEnum) XEnumObject
	// HasXEnumVal checks if XEnumVal has been set in XEnumObject
	HasXEnumVal() bool
}

type XEnumObjectXEnumValEnum string

// Enum of XEnumVal on XEnumObject
var XEnumObjectXEnumVal = struct {
	FIRST_VAL  XEnumObjectXEnumValEnum
	SECOND_VAL XEnumObjectXEnumValEnum
	THIRD_VAL  XEnumObjectXEnumValEnum
	FOURTH_VAL XEnumObjectXEnumValEnum
}{
	FIRST_VAL:  XEnumObjectXEnumValEnum("first_val"),
	SECOND_VAL: XEnumObjectXEnumValEnum("second_val"),
	THIRD_VAL:  XEnumObjectXEnumValEnum("third_val"),
	FOURTH_VAL: XEnumObjectXEnumValEnum("fourth_val"),
}

func (obj *xEnumObject) XEnumVal() XEnumObjectXEnumValEnum {
	return XEnumObjectXEnumValEnum(obj.obj.XEnumVal.Enum().String())
}

// A property to showcase x-enum feature
// XEnumVal returns a string
func (obj *xEnumObject) HasXEnumVal() bool {
	return obj.obj.XEnumVal != nil
}

func (obj *xEnumObject) SetXEnumVal(value XEnumObjectXEnumValEnum) XEnumObject {
	intValue, ok := openapi.XEnumObject_XEnumVal_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on XEnumObjectXEnumValEnum", string(value)))
		return obj
	}
	enumValue := openapi.XEnumObject_XEnumVal_Enum(intValue)
	obj.obj.XEnumVal = &enumValue

	return obj
}

func (obj *xEnumObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *xEnumObject) setDefault() {
	if obj.obj.XEnumVal == nil {
		obj.SetXEnumVal(XEnumObjectXEnumVal.THIRD_VAL)

	}

}
