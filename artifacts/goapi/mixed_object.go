package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MixedObject *****
type mixedObject struct {
	validation
	obj          *openapi.MixedObject
	marshaller   marshalMixedObject
	unMarshaller unMarshalMixedObject
}

func NewMixedObject() MixedObject {
	obj := mixedObject{obj: &openapi.MixedObject{}}
	obj.setDefault()
	return &obj
}

func (obj *mixedObject) msg() *openapi.MixedObject {
	return obj.obj
}

func (obj *mixedObject) setMsg(msg *openapi.MixedObject) MixedObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmixedObject struct {
	obj *mixedObject
}

type marshalMixedObject interface {
	// ToProto marshals MixedObject to protobuf object *openapi.MixedObject
	ToProto() (*openapi.MixedObject, error)
	// ToPbText marshals MixedObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MixedObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals MixedObject to JSON text
	ToJson() (string, error)
}

type unMarshalmixedObject struct {
	obj *mixedObject
}

type unMarshalMixedObject interface {
	// FromProto unmarshals MixedObject from protobuf object *openapi.MixedObject
	FromProto(msg *openapi.MixedObject) (MixedObject, error)
	// FromPbText unmarshals MixedObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MixedObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MixedObject from JSON text
	FromJson(value string) error
}

func (obj *mixedObject) Marshal() marshalMixedObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmixedObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *mixedObject) Unmarshal() unMarshalMixedObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmixedObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmixedObject) ToProto() (*openapi.MixedObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmixedObject) FromProto(msg *openapi.MixedObject) (MixedObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmixedObject) ToPbText() (string, error) {
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

func (m *unMarshalmixedObject) FromPbText(value string) error {
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

func (m *marshalmixedObject) ToYaml() (string, error) {
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

func (m *unMarshalmixedObject) FromYaml(value string) error {
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

func (m *marshalmixedObject) ToJson() (string, error) {
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

func (m *unMarshalmixedObject) FromJson(value string) error {
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

func (obj *mixedObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mixedObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mixedObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mixedObject) Clone() (MixedObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMixedObject()
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

// MixedObject is format validation object
type MixedObject interface {
	Validation
	// msg marshals MixedObject to protobuf object *openapi.MixedObject
	// and doesn't set defaults
	msg() *openapi.MixedObject
	// setMsg unmarshals MixedObject from protobuf object *openapi.MixedObject
	// and doesn't set defaults
	setMsg(*openapi.MixedObject) MixedObject
	// provides marshal interface
	Marshal() marshalMixedObject
	// provides unmarshal interface
	Unmarshal() unMarshalMixedObject
	// validate validates MixedObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MixedObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StringParam returns string, set in MixedObject.
	StringParam() string
	// SetStringParam assigns string provided by user to MixedObject
	SetStringParam(value string) MixedObject
	// HasStringParam checks if StringParam has been set in MixedObject
	HasStringParam() bool
	// Integer returns int32, set in MixedObject.
	Integer() int32
	// SetInteger assigns int32 provided by user to MixedObject
	SetInteger(value int32) MixedObject
	// HasInteger checks if Integer has been set in MixedObject
	HasInteger() bool
	// Float returns float32, set in MixedObject.
	Float() float32
	// SetFloat assigns float32 provided by user to MixedObject
	SetFloat(value float32) MixedObject
	// HasFloat checks if Float has been set in MixedObject
	HasFloat() bool
	// Double returns float64, set in MixedObject.
	Double() float64
	// SetDouble assigns float64 provided by user to MixedObject
	SetDouble(value float64) MixedObject
	// HasDouble checks if Double has been set in MixedObject
	HasDouble() bool
	// Mac returns string, set in MixedObject.
	Mac() string
	// SetMac assigns string provided by user to MixedObject
	SetMac(value string) MixedObject
	// HasMac checks if Mac has been set in MixedObject
	HasMac() bool
	// Ipv4 returns string, set in MixedObject.
	Ipv4() string
	// SetIpv4 assigns string provided by user to MixedObject
	SetIpv4(value string) MixedObject
	// HasIpv4 checks if Ipv4 has been set in MixedObject
	HasIpv4() bool
	// Ipv6 returns string, set in MixedObject.
	Ipv6() string
	// SetIpv6 assigns string provided by user to MixedObject
	SetIpv6(value string) MixedObject
	// HasIpv6 checks if Ipv6 has been set in MixedObject
	HasIpv6() bool
	// Hex returns string, set in MixedObject.
	Hex() string
	// SetHex assigns string provided by user to MixedObject
	SetHex(value string) MixedObject
	// HasHex checks if Hex has been set in MixedObject
	HasHex() bool
	// StrLen returns string, set in MixedObject.
	StrLen() string
	// SetStrLen assigns string provided by user to MixedObject
	SetStrLen(value string) MixedObject
	// HasStrLen checks if StrLen has been set in MixedObject
	HasStrLen() bool
	// Integer641 returns int64, set in MixedObject.
	Integer641() int64
	// SetInteger641 assigns int64 provided by user to MixedObject
	SetInteger641(value int64) MixedObject
	// HasInteger641 checks if Integer641 has been set in MixedObject
	HasInteger641() bool
	// Integer642 returns int64, set in MixedObject.
	Integer642() int64
	// SetInteger642 assigns int64 provided by user to MixedObject
	SetInteger642(value int64) MixedObject
	// HasInteger642 checks if Integer642 has been set in MixedObject
	HasInteger642() bool
	// Integer64List returns []int64, set in MixedObject.
	Integer64List() []int64
	// SetInteger64List assigns []int64 provided by user to MixedObject
	SetInteger64List(value []int64) MixedObject
}

// description is TBD
// StringParam returns a string
func (obj *mixedObject) StringParam() string {

	return *obj.obj.StringParam

}

// description is TBD
// StringParam returns a string
func (obj *mixedObject) HasStringParam() bool {
	return obj.obj.StringParam != nil
}

// description is TBD
// SetStringParam sets the string value in the MixedObject object
func (obj *mixedObject) SetStringParam(value string) MixedObject {

	obj.obj.StringParam = &value
	return obj
}

// description is TBD
// Integer returns a int32
func (obj *mixedObject) Integer() int32 {

	return *obj.obj.Integer

}

// description is TBD
// Integer returns a int32
func (obj *mixedObject) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the int32 value in the MixedObject object
func (obj *mixedObject) SetInteger(value int32) MixedObject {

	obj.obj.Integer = &value
	return obj
}

// description is TBD
// Float returns a float32
func (obj *mixedObject) Float() float32 {

	return *obj.obj.Float

}

// description is TBD
// Float returns a float32
func (obj *mixedObject) HasFloat() bool {
	return obj.obj.Float != nil
}

// description is TBD
// SetFloat sets the float32 value in the MixedObject object
func (obj *mixedObject) SetFloat(value float32) MixedObject {

	obj.obj.Float = &value
	return obj
}

// description is TBD
// Double returns a float64
func (obj *mixedObject) Double() float64 {

	return *obj.obj.Double

}

// description is TBD
// Double returns a float64
func (obj *mixedObject) HasDouble() bool {
	return obj.obj.Double != nil
}

// description is TBD
// SetDouble sets the float64 value in the MixedObject object
func (obj *mixedObject) SetDouble(value float64) MixedObject {

	obj.obj.Double = &value
	return obj
}

// description is TBD
// Mac returns a string
func (obj *mixedObject) Mac() string {

	return *obj.obj.Mac

}

// description is TBD
// Mac returns a string
func (obj *mixedObject) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the string value in the MixedObject object
func (obj *mixedObject) SetMac(value string) MixedObject {

	obj.obj.Mac = &value
	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *mixedObject) Ipv4() string {

	return *obj.obj.Ipv4

}

// description is TBD
// Ipv4 returns a string
func (obj *mixedObject) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the string value in the MixedObject object
func (obj *mixedObject) SetIpv4(value string) MixedObject {

	obj.obj.Ipv4 = &value
	return obj
}

// description is TBD
// Ipv6 returns a string
func (obj *mixedObject) Ipv6() string {

	return *obj.obj.Ipv6

}

// description is TBD
// Ipv6 returns a string
func (obj *mixedObject) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the string value in the MixedObject object
func (obj *mixedObject) SetIpv6(value string) MixedObject {

	obj.obj.Ipv6 = &value
	return obj
}

// description is TBD
// Hex returns a string
func (obj *mixedObject) Hex() string {

	return *obj.obj.Hex

}

// description is TBD
// Hex returns a string
func (obj *mixedObject) HasHex() bool {
	return obj.obj.Hex != nil
}

// description is TBD
// SetHex sets the string value in the MixedObject object
func (obj *mixedObject) SetHex(value string) MixedObject {

	obj.obj.Hex = &value
	return obj
}

// Under Review: Information TBD
//
// Description TBD
// StrLen returns a string
func (obj *mixedObject) StrLen() string {

	return *obj.obj.StrLen

}

// Under Review: Information TBD
//
// Description TBD
// StrLen returns a string
func (obj *mixedObject) HasStrLen() bool {
	return obj.obj.StrLen != nil
}

// Under Review: Information TBD
//
// Description TBD
// SetStrLen sets the string value in the MixedObject object
func (obj *mixedObject) SetStrLen(value string) MixedObject {

	obj.obj.StrLen = &value
	return obj
}

// int64 type
// Integer641 returns a int64
func (obj *mixedObject) Integer641() int64 {

	return *obj.obj.Integer64_1

}

// int64 type
// Integer641 returns a int64
func (obj *mixedObject) HasInteger641() bool {
	return obj.obj.Integer64_1 != nil
}

// int64 type
// SetInteger641 sets the int64 value in the MixedObject object
func (obj *mixedObject) SetInteger641(value int64) MixedObject {

	obj.obj.Integer64_1 = &value
	return obj
}

// description is TBD
// Integer642 returns a int64
func (obj *mixedObject) Integer642() int64 {

	return *obj.obj.Integer64_2

}

// description is TBD
// Integer642 returns a int64
func (obj *mixedObject) HasInteger642() bool {
	return obj.obj.Integer64_2 != nil
}

// description is TBD
// SetInteger642 sets the int64 value in the MixedObject object
func (obj *mixedObject) SetInteger642(value int64) MixedObject {

	obj.obj.Integer64_2 = &value
	return obj
}

// int64 type list
// Integer64List returns a []int64
func (obj *mixedObject) Integer64List() []int64 {
	if obj.obj.Integer64List == nil {
		obj.obj.Integer64List = make([]int64, 0)
	}
	return obj.obj.Integer64List
}

// int64 type list
// SetInteger64List sets the []int64 value in the MixedObject object
func (obj *mixedObject) SetInteger64List(value []int64) MixedObject {

	if obj.obj.Integer64List == nil {
		obj.obj.Integer64List = make([]int64, 0)
	}
	obj.obj.Integer64List = value

	return obj
}

func (obj *mixedObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		if *obj.obj.Integer < 10 || *obj.obj.Integer > 90 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= MixedObject.Integer <= 90 but Got %d", *obj.obj.Integer))
		}

	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MixedObject.Mac"))
		}

	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MixedObject.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MixedObject.Ipv6"))
		}

	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MixedObject.Hex"))
		}

	}

	if obj.obj.StrLen != nil {
		obj.addWarnings("StrLen property in schema MixedObject is under review, Information TBD")
		if len(*obj.obj.StrLen) < 3 || len(*obj.obj.StrLen) > 6 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"3 <= length of MixedObject.StrLen <= 6 but Got %d",
					len(*obj.obj.StrLen)))
		}

	}

	if obj.obj.Integer64_2 != nil {

		if *obj.obj.Integer64_2 < 0 || *obj.obj.Integer64_2 > 4261412864 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= MixedObject.Integer64_2 <= 4261412864 but Got %d", *obj.obj.Integer64_2))
		}

	}

	if obj.obj.Integer64List != nil {

		for _, item := range obj.obj.Integer64List {
			if item < 0 || item > 4261412864 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= MixedObject.Integer64List <= 4261412864 but Got %d", item))
			}

		}

	}

}

func (obj *mixedObject) setDefault() {
	if obj.obj.StringParam == nil {
		obj.SetStringParam("asdf")
	}
	if obj.obj.Integer == nil {
		obj.SetInteger(88)
	}
	if obj.obj.Float == nil {
		obj.SetFloat(22.2)
	}
	if obj.obj.Double == nil {
		obj.SetDouble(2342.22)
	}
	if obj.obj.Mac == nil {
		obj.SetMac("00:00:fa:ce:fa:ce")
	}
	if obj.obj.Ipv4 == nil {
		obj.SetIpv4("1.1.1.1")
	}
	if obj.obj.Ipv6 == nil {
		obj.SetIpv6("::02")
	}
	if obj.obj.Hex == nil {
		obj.SetHex("0102030405060708090a0b0c0d0e0f")
	}

}
