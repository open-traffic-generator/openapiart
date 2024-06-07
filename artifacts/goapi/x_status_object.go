package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** XStatusObject *****
type xStatusObject struct {
	validation
	obj          *openapi.XStatusObject
	marshaller   marshalXStatusObject
	unMarshaller unMarshalXStatusObject
}

func NewXStatusObject() XStatusObject {
	obj := xStatusObject{obj: &openapi.XStatusObject{}}
	obj.setDefault()
	return &obj
}

func (obj *xStatusObject) msg() *openapi.XStatusObject {
	return obj.obj
}

func (obj *xStatusObject) setMsg(msg *openapi.XStatusObject) XStatusObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalxStatusObject struct {
	obj *xStatusObject
}

type marshalXStatusObject interface {
	// ToProto marshals XStatusObject to protobuf object *openapi.XStatusObject
	ToProto() (*openapi.XStatusObject, error)
	// ToPbText marshals XStatusObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals XStatusObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals XStatusObject to JSON text
	ToJson() (string, error)
}

type unMarshalxStatusObject struct {
	obj *xStatusObject
}

type unMarshalXStatusObject interface {
	// FromProto unmarshals XStatusObject from protobuf object *openapi.XStatusObject
	FromProto(msg *openapi.XStatusObject) (XStatusObject, error)
	// FromPbText unmarshals XStatusObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals XStatusObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals XStatusObject from JSON text
	FromJson(value string) error
}

func (obj *xStatusObject) Marshal() marshalXStatusObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalxStatusObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *xStatusObject) Unmarshal() unMarshalXStatusObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalxStatusObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalxStatusObject) ToProto() (*openapi.XStatusObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalxStatusObject) FromProto(msg *openapi.XStatusObject) (XStatusObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalxStatusObject) ToPbText() (string, error) {
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

func (m *unMarshalxStatusObject) FromPbText(value string) error {
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

func (m *marshalxStatusObject) ToYaml() (string, error) {
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

func (m *unMarshalxStatusObject) FromYaml(value string) error {
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

func (m *marshalxStatusObject) ToJson() (string, error) {
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

func (m *unMarshalxStatusObject) FromJson(value string) error {
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

func (obj *xStatusObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *xStatusObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *xStatusObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *xStatusObject) Clone() (XStatusObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewXStatusObject()
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

// XStatusObject is description is TBD
type XStatusObject interface {
	Validation
	// msg marshals XStatusObject to protobuf object *openapi.XStatusObject
	// and doesn't set defaults
	msg() *openapi.XStatusObject
	// setMsg unmarshals XStatusObject from protobuf object *openapi.XStatusObject
	// and doesn't set defaults
	setMsg(*openapi.XStatusObject) XStatusObject
	// provides marshal interface
	Marshal() marshalXStatusObject
	// provides unmarshal interface
	Unmarshal() unMarshalXStatusObject
	// validate validates XStatusObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (XStatusObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnumProperty returns XStatusObjectEnumPropertyEnum, set in XStatusObject
	EnumProperty() XStatusObjectEnumPropertyEnum
	// SetEnumProperty assigns XStatusObjectEnumPropertyEnum provided by user to XStatusObject
	SetEnumProperty(value XStatusObjectEnumPropertyEnum) XStatusObject
	// HasEnumProperty checks if EnumProperty has been set in XStatusObject
	HasEnumProperty() bool
	// DecprecatedProperty1 returns string, set in XStatusObject.
	DecprecatedProperty1() string
	// SetDecprecatedProperty1 assigns string provided by user to XStatusObject
	SetDecprecatedProperty1(value string) XStatusObject
	// HasDecprecatedProperty1 checks if DecprecatedProperty1 has been set in XStatusObject
	HasDecprecatedProperty1() bool
	// DecprecatedProperty2 returns int32, set in XStatusObject.
	DecprecatedProperty2() int32
	// SetDecprecatedProperty2 assigns int32 provided by user to XStatusObject
	SetDecprecatedProperty2(value int32) XStatusObject
	// HasDecprecatedProperty2 checks if DecprecatedProperty2 has been set in XStatusObject
	HasDecprecatedProperty2() bool
	// UnderReviewProperty1 returns string, set in XStatusObject.
	UnderReviewProperty1() string
	// SetUnderReviewProperty1 assigns string provided by user to XStatusObject
	SetUnderReviewProperty1(value string) XStatusObject
	// HasUnderReviewProperty1 checks if UnderReviewProperty1 has been set in XStatusObject
	HasUnderReviewProperty1() bool
	// UnderReviewProperty2 returns int32, set in XStatusObject.
	UnderReviewProperty2() int32
	// SetUnderReviewProperty2 assigns int32 provided by user to XStatusObject
	SetUnderReviewProperty2(value int32) XStatusObject
	// HasUnderReviewProperty2 checks if UnderReviewProperty2 has been set in XStatusObject
	HasUnderReviewProperty2() bool
	// Basic returns string, set in XStatusObject.
	Basic() string
	// SetBasic assigns string provided by user to XStatusObject
	SetBasic(value string) XStatusObject
	// HasBasic checks if Basic has been set in XStatusObject
	HasBasic() bool
}

type XStatusObjectEnumPropertyEnum string

// Enum of EnumProperty on XStatusObject
var XStatusObjectEnumProperty = struct {
	DECPRECATED_PROPERTY_1  XStatusObjectEnumPropertyEnum
	UNDER_REVIEW_PROPERTY_1 XStatusObjectEnumPropertyEnum
	BASIC                   XStatusObjectEnumPropertyEnum
}{
	DECPRECATED_PROPERTY_1:  XStatusObjectEnumPropertyEnum("decprecated_property_1"),
	UNDER_REVIEW_PROPERTY_1: XStatusObjectEnumPropertyEnum("under_review_property_1"),
	BASIC:                   XStatusObjectEnumPropertyEnum("basic"),
}

func (obj *xStatusObject) EnumProperty() XStatusObjectEnumPropertyEnum {
	return XStatusObjectEnumPropertyEnum(obj.obj.EnumProperty.Enum().String())
}

// description is TBD
// EnumProperty returns a string
func (obj *xStatusObject) HasEnumProperty() bool {
	return obj.obj.EnumProperty != nil
}

func (obj *xStatusObject) SetEnumProperty(value XStatusObjectEnumPropertyEnum) XStatusObject {
	intValue, ok := openapi.XStatusObject_EnumProperty_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on XStatusObjectEnumPropertyEnum", string(value)))
		return obj
	}
	enumValue := openapi.XStatusObject_EnumProperty_Enum(intValue)
	obj.obj.EnumProperty = &enumValue

	return obj
}

// description is TBD
// DecprecatedProperty1 returns a string
func (obj *xStatusObject) DecprecatedProperty1() string {

	return *obj.obj.DecprecatedProperty_1

}

// description is TBD
// DecprecatedProperty1 returns a string
func (obj *xStatusObject) HasDecprecatedProperty1() bool {
	return obj.obj.DecprecatedProperty_1 != nil
}

// description is TBD
// SetDecprecatedProperty1 sets the string value in the XStatusObject object
func (obj *xStatusObject) SetDecprecatedProperty1(value string) XStatusObject {

	obj.obj.DecprecatedProperty_1 = &value
	return obj
}

// Deprecated: test deprecated
//
// Description TBD
// DecprecatedProperty2 returns a int32
func (obj *xStatusObject) DecprecatedProperty2() int32 {

	return *obj.obj.DecprecatedProperty_2

}

// Deprecated: test deprecated
//
// Description TBD
// DecprecatedProperty2 returns a int32
func (obj *xStatusObject) HasDecprecatedProperty2() bool {
	return obj.obj.DecprecatedProperty_2 != nil
}

// Deprecated: test deprecated
//
// Description TBD
// SetDecprecatedProperty2 sets the int32 value in the XStatusObject object
func (obj *xStatusObject) SetDecprecatedProperty2(value int32) XStatusObject {

	obj.obj.DecprecatedProperty_2 = &value
	return obj
}

// description is TBD
// UnderReviewProperty1 returns a string
func (obj *xStatusObject) UnderReviewProperty1() string {

	return *obj.obj.UnderReviewProperty_1

}

// description is TBD
// UnderReviewProperty1 returns a string
func (obj *xStatusObject) HasUnderReviewProperty1() bool {
	return obj.obj.UnderReviewProperty_1 != nil
}

// description is TBD
// SetUnderReviewProperty1 sets the string value in the XStatusObject object
func (obj *xStatusObject) SetUnderReviewProperty1(value string) XStatusObject {

	obj.obj.UnderReviewProperty_1 = &value
	return obj
}

// Under Review: test under_review
//
// Description TBD
// UnderReviewProperty2 returns a int32
func (obj *xStatusObject) UnderReviewProperty2() int32 {

	return *obj.obj.UnderReviewProperty_2

}

// Under Review: test under_review
//
// Description TBD
// UnderReviewProperty2 returns a int32
func (obj *xStatusObject) HasUnderReviewProperty2() bool {
	return obj.obj.UnderReviewProperty_2 != nil
}

// Under Review: test under_review
//
// Description TBD
// SetUnderReviewProperty2 sets the int32 value in the XStatusObject object
func (obj *xStatusObject) SetUnderReviewProperty2(value int32) XStatusObject {

	obj.obj.UnderReviewProperty_2 = &value
	return obj
}

// description is TBD
// Basic returns a string
func (obj *xStatusObject) Basic() string {

	return *obj.obj.Basic

}

// description is TBD
// Basic returns a string
func (obj *xStatusObject) HasBasic() bool {
	return obj.obj.Basic != nil
}

// description is TBD
// SetBasic sets the string value in the XStatusObject object
func (obj *xStatusObject) SetBasic(value string) XStatusObject {

	obj.obj.Basic = &value
	return obj
}

func (obj *xStatusObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EnumProperty.Number() == 1 {
		obj.addWarnings("DECPRECATED_PROPERTY_1 enum in property EnumProperty is deprecated, test deprecated")
	}

	if obj.obj.EnumProperty.Number() == 2 {
		obj.addWarnings("UNDER_REVIEW_PROPERTY_1 enum in property EnumProperty is under review, test under_review")
	}

	// DecprecatedProperty_2 is deprecated
	if obj.obj.DecprecatedProperty_2 != nil {
		obj.addWarnings("DecprecatedProperty_2 property in schema XStatusObject is deprecated, test deprecated")
	}

	// UnderReviewProperty_2 is under_review
	if obj.obj.UnderReviewProperty_2 != nil {
		obj.addWarnings("UnderReviewProperty_2 property in schema XStatusObject is under review, test under_review")
	}

}

func (obj *xStatusObject) setDefault() {

}
