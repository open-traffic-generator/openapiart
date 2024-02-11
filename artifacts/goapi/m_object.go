package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MObject *****
type mObject struct {
	validation
	obj          *openapi.MObject
	marshaller   marshalMObject
	unMarshaller unMarshalMObject
}

func NewMObject() MObject {
	obj := mObject{obj: &openapi.MObject{}}
	obj.setDefault()
	return &obj
}

func (obj *mObject) msg() *openapi.MObject {
	return obj.obj
}

func (obj *mObject) setMsg(msg *openapi.MObject) MObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmObject struct {
	obj *mObject
}

type marshalMObject interface {
	// ToProto marshals MObject to protobuf object *openapi.MObject
	ToProto() (*openapi.MObject, error)
	// ToPbText marshals MObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals MObject to JSON text
	ToJson() (string, error)
}

type unMarshalmObject struct {
	obj *mObject
}

type unMarshalMObject interface {
	// FromProto unmarshals MObject from protobuf object *openapi.MObject
	FromProto(msg *openapi.MObject) (MObject, error)
	// FromPbText unmarshals MObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MObject from JSON text
	FromJson(value string) error
}

func (obj *mObject) Marshal() marshalMObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *mObject) Unmarshal() unMarshalMObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmObject) ToProto() (*openapi.MObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmObject) FromProto(msg *openapi.MObject) (MObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmObject) ToPbText() (string, error) {
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

func (m *unMarshalmObject) FromPbText(value string) error {
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

func (m *marshalmObject) ToYaml() (string, error) {
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

func (m *unMarshalmObject) FromYaml(value string) error {
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

func (m *marshalmObject) ToJson() (string, error) {
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

func (m *unMarshalmObject) FromJson(value string) error {
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

func (obj *mObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mObject) Clone() (MObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMObject()
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

// MObject is required format validation object
type MObject interface {
	Validation
	// msg marshals MObject to protobuf object *openapi.MObject
	// and doesn't set defaults
	msg() *openapi.MObject
	// setMsg unmarshals MObject from protobuf object *openapi.MObject
	// and doesn't set defaults
	setMsg(*openapi.MObject) MObject
	// provides marshal interface
	Marshal() marshalMObject
	// provides unmarshal interface
	Unmarshal() unMarshalMObject
	// validate validates MObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StringParam returns string, set in MObject.
	StringParam() string
	// SetStringParam assigns string provided by user to MObject
	SetStringParam(value string) MObject
	// Integer returns int32, set in MObject.
	Integer() int32
	// SetInteger assigns int32 provided by user to MObject
	SetInteger(value int32) MObject
	// Float returns float32, set in MObject.
	Float() float32
	// SetFloat assigns float32 provided by user to MObject
	SetFloat(value float32) MObject
	// Double returns float64, set in MObject.
	Double() float64
	// SetDouble assigns float64 provided by user to MObject
	SetDouble(value float64) MObject
	// Mac returns string, set in MObject.
	Mac() string
	// SetMac assigns string provided by user to MObject
	SetMac(value string) MObject
	// Ipv4 returns string, set in MObject.
	Ipv4() string
	// SetIpv4 assigns string provided by user to MObject
	SetIpv4(value string) MObject
	// Ipv6 returns string, set in MObject.
	Ipv6() string
	// SetIpv6 assigns string provided by user to MObject
	SetIpv6(value string) MObject
	// Hex returns string, set in MObject.
	Hex() string
	// SetHex assigns string provided by user to MObject
	SetHex(value string) MObject
	// Oid returns string, set in MObject.
	Oid() string
	// SetOid assigns string provided by user to MObject
	SetOid(value string) MObject
	// HasOid checks if Oid has been set in MObject
	HasOid() bool
}

// description is TBD
// StringParam returns a string
func (obj *mObject) StringParam() string {

	return *obj.obj.StringParam

}

// description is TBD
// SetStringParam sets the string value in the MObject object
func (obj *mObject) SetStringParam(value string) MObject {

	obj.obj.StringParam = &value
	return obj
}

// description is TBD
// Integer returns a int32
func (obj *mObject) Integer() int32 {

	return *obj.obj.Integer

}

// description is TBD
// SetInteger sets the int32 value in the MObject object
func (obj *mObject) SetInteger(value int32) MObject {

	obj.obj.Integer = &value
	return obj
}

// description is TBD
// Float returns a float32
func (obj *mObject) Float() float32 {

	return *obj.obj.Float

}

// description is TBD
// SetFloat sets the float32 value in the MObject object
func (obj *mObject) SetFloat(value float32) MObject {

	obj.obj.Float = &value
	return obj
}

// description is TBD
// Double returns a float64
func (obj *mObject) Double() float64 {

	return *obj.obj.Double

}

// description is TBD
// SetDouble sets the float64 value in the MObject object
func (obj *mObject) SetDouble(value float64) MObject {

	obj.obj.Double = &value
	return obj
}

// description is TBD
// Mac returns a string
func (obj *mObject) Mac() string {

	return *obj.obj.Mac

}

// description is TBD
// SetMac sets the string value in the MObject object
func (obj *mObject) SetMac(value string) MObject {

	obj.obj.Mac = &value
	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *mObject) Ipv4() string {

	return *obj.obj.Ipv4

}

// description is TBD
// SetIpv4 sets the string value in the MObject object
func (obj *mObject) SetIpv4(value string) MObject {

	obj.obj.Ipv4 = &value
	return obj
}

// description is TBD
// Ipv6 returns a string
func (obj *mObject) Ipv6() string {

	return *obj.obj.Ipv6

}

// description is TBD
// SetIpv6 sets the string value in the MObject object
func (obj *mObject) SetIpv6(value string) MObject {

	obj.obj.Ipv6 = &value
	return obj
}

// description is TBD
// Hex returns a string
func (obj *mObject) Hex() string {

	return *obj.obj.Hex

}

// description is TBD
// SetHex sets the string value in the MObject object
func (obj *mObject) SetHex(value string) MObject {

	obj.obj.Hex = &value
	return obj
}

// description is TBD
// Oid returns a string
func (obj *mObject) Oid() string {

	return *obj.obj.Oid

}

// description is TBD
// Oid returns a string
func (obj *mObject) HasOid() bool {
	return obj.obj.Oid != nil
}

// description is TBD
// SetOid sets the string value in the MObject object
func (obj *mObject) SetOid(value string) MObject {

	obj.obj.Oid = &value
	return obj
}

func (obj *mObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// StringParam is required
	if obj.obj.StringParam == nil {
		vObj.validationErrors = append(vObj.validationErrors, "StringParam is required field on interface MObject")
	}

	// Integer is required
	if obj.obj.Integer == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Integer is required field on interface MObject")
	}
	if obj.obj.Integer != nil {

		if *obj.obj.Integer < -10 || *obj.obj.Integer > 90 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-10 <= MObject.Integer <= 90 but Got %d", *obj.obj.Integer))
		}

	}

	// Float is required
	if obj.obj.Float == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Float is required field on interface MObject")
	}

	// Double is required
	if obj.obj.Double == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Double is required field on interface MObject")
	}

	// Mac is required
	if obj.obj.Mac == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Mac is required field on interface MObject")
	}
	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Mac"))
		}

	}

	// Ipv4 is required
	if obj.obj.Ipv4 == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4 is required field on interface MObject")
	}
	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv4"))
		}

	}

	// Ipv6 is required
	if obj.obj.Ipv6 == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6 is required field on interface MObject")
	}
	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv6"))
		}

	}

	// Hex is required
	if obj.obj.Hex == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Hex is required field on interface MObject")
	}
	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Hex"))
		}

	}

	if obj.obj.Oid != nil {

		err := obj.validateOid(obj.Oid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Oid"))
		}

	}

}

func (obj *mObject) setDefault() {

}
