package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LObject *****
type lObject struct {
	validation
	obj          *openapi.LObject
	marshaller   marshalLObject
	unMarshaller unMarshalLObject
}

func NewLObject() LObject {
	obj := lObject{obj: &openapi.LObject{}}
	obj.setDefault()
	return &obj
}

func (obj *lObject) msg() *openapi.LObject {
	return obj.obj
}

func (obj *lObject) setMsg(msg *openapi.LObject) LObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallObject struct {
	obj *lObject
}

type marshalLObject interface {
	// ToProto marshals LObject to protobuf object *openapi.LObject
	ToProto() (*openapi.LObject, error)
	// ToPbText marshals LObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals LObject to JSON text
	ToJson() (string, error)
}

type unMarshallObject struct {
	obj *lObject
}

type unMarshalLObject interface {
	// FromProto unmarshals LObject from protobuf object *openapi.LObject
	FromProto(msg *openapi.LObject) (LObject, error)
	// FromPbText unmarshals LObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LObject from JSON text
	FromJson(value string) error
}

func (obj *lObject) Marshal() marshalLObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshallObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *lObject) Unmarshal() unMarshalLObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallObject) ToProto() (*openapi.LObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallObject) FromProto(msg *openapi.LObject) (LObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallObject) ToPbText() (string, error) {
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

func (m *unMarshallObject) FromPbText(value string) error {
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

func (m *marshallObject) ToYaml() (string, error) {
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

func (m *unMarshallObject) FromYaml(value string) error {
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

func (m *marshallObject) ToJson() (string, error) {
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

func (m *unMarshallObject) FromJson(value string) error {
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

func (obj *lObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lObject) Clone() (LObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLObject()
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

// LObject is format validation object
type LObject interface {
	Validation
	// msg marshals LObject to protobuf object *openapi.LObject
	// and doesn't set defaults
	msg() *openapi.LObject
	// setMsg unmarshals LObject from protobuf object *openapi.LObject
	// and doesn't set defaults
	setMsg(*openapi.LObject) LObject
	// provides marshal interface
	Marshal() marshalLObject
	// provides unmarshal interface
	Unmarshal() unMarshalLObject
	// validate validates LObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StringParam returns string, set in LObject.
	StringParam() string
	// SetStringParam assigns string provided by user to LObject
	SetStringParam(value string) LObject
	// HasStringParam checks if StringParam has been set in LObject
	HasStringParam() bool
	// Integer returns int32, set in LObject.
	Integer() int32
	// SetInteger assigns int32 provided by user to LObject
	SetInteger(value int32) LObject
	// HasInteger checks if Integer has been set in LObject
	HasInteger() bool
	// Float returns float32, set in LObject.
	Float() float32
	// SetFloat assigns float32 provided by user to LObject
	SetFloat(value float32) LObject
	// HasFloat checks if Float has been set in LObject
	HasFloat() bool
	// Double returns float64, set in LObject.
	Double() float64
	// SetDouble assigns float64 provided by user to LObject
	SetDouble(value float64) LObject
	// HasDouble checks if Double has been set in LObject
	HasDouble() bool
	// Mac returns string, set in LObject.
	Mac() string
	// SetMac assigns string provided by user to LObject
	SetMac(value string) LObject
	// HasMac checks if Mac has been set in LObject
	HasMac() bool
	// Ipv4 returns string, set in LObject.
	Ipv4() string
	// SetIpv4 assigns string provided by user to LObject
	SetIpv4(value string) LObject
	// HasIpv4 checks if Ipv4 has been set in LObject
	HasIpv4() bool
	// Ipv6 returns string, set in LObject.
	Ipv6() string
	// SetIpv6 assigns string provided by user to LObject
	SetIpv6(value string) LObject
	// HasIpv6 checks if Ipv6 has been set in LObject
	HasIpv6() bool
	// Hex returns string, set in LObject.
	Hex() string
	// SetHex assigns string provided by user to LObject
	SetHex(value string) LObject
	// HasHex checks if Hex has been set in LObject
	HasHex() bool
}

// description is TBD
// StringParam returns a string
func (obj *lObject) StringParam() string {

	return *obj.obj.StringParam

}

// description is TBD
// StringParam returns a string
func (obj *lObject) HasStringParam() bool {
	return obj.obj.StringParam != nil
}

// description is TBD
// SetStringParam sets the string value in the LObject object
func (obj *lObject) SetStringParam(value string) LObject {

	obj.obj.StringParam = &value
	return obj
}

// description is TBD
// Integer returns a int32
func (obj *lObject) Integer() int32 {

	return *obj.obj.Integer

}

// description is TBD
// Integer returns a int32
func (obj *lObject) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the int32 value in the LObject object
func (obj *lObject) SetInteger(value int32) LObject {

	obj.obj.Integer = &value
	return obj
}

// description is TBD
// Float returns a float32
func (obj *lObject) Float() float32 {

	return *obj.obj.Float

}

// description is TBD
// Float returns a float32
func (obj *lObject) HasFloat() bool {
	return obj.obj.Float != nil
}

// description is TBD
// SetFloat sets the float32 value in the LObject object
func (obj *lObject) SetFloat(value float32) LObject {

	obj.obj.Float = &value
	return obj
}

// description is TBD
// Double returns a float64
func (obj *lObject) Double() float64 {

	return *obj.obj.Double

}

// description is TBD
// Double returns a float64
func (obj *lObject) HasDouble() bool {
	return obj.obj.Double != nil
}

// description is TBD
// SetDouble sets the float64 value in the LObject object
func (obj *lObject) SetDouble(value float64) LObject {

	obj.obj.Double = &value
	return obj
}

// description is TBD
// Mac returns a string
func (obj *lObject) Mac() string {

	return *obj.obj.Mac

}

// description is TBD
// Mac returns a string
func (obj *lObject) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the string value in the LObject object
func (obj *lObject) SetMac(value string) LObject {

	obj.obj.Mac = &value
	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *lObject) Ipv4() string {

	return *obj.obj.Ipv4

}

// description is TBD
// Ipv4 returns a string
func (obj *lObject) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the string value in the LObject object
func (obj *lObject) SetIpv4(value string) LObject {

	obj.obj.Ipv4 = &value
	return obj
}

// description is TBD
// Ipv6 returns a string
func (obj *lObject) Ipv6() string {

	return *obj.obj.Ipv6

}

// description is TBD
// Ipv6 returns a string
func (obj *lObject) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the string value in the LObject object
func (obj *lObject) SetIpv6(value string) LObject {

	obj.obj.Ipv6 = &value
	return obj
}

// description is TBD
// Hex returns a string
func (obj *lObject) Hex() string {

	return *obj.obj.Hex

}

// description is TBD
// Hex returns a string
func (obj *lObject) HasHex() bool {
	return obj.obj.Hex != nil
}

// description is TBD
// SetHex sets the string value in the LObject object
func (obj *lObject) SetHex(value string) LObject {

	obj.obj.Hex = &value
	return obj
}

func (obj *lObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		if *obj.obj.Integer < -10 || *obj.obj.Integer > 90 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-10 <= LObject.Integer <= 90 but Got %d", *obj.obj.Integer))
		}

	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Mac"))
		}

	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv6"))
		}

	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Hex"))
		}

	}

}

func (obj *lObject) setDefault() {

}
