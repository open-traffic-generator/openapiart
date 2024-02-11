package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** EObject *****
type eObject struct {
	validation
	obj          *openapi.EObject
	marshaller   marshalEObject
	unMarshaller unMarshalEObject
}

func NewEObject() EObject {
	obj := eObject{obj: &openapi.EObject{}}
	obj.setDefault()
	return &obj
}

func (obj *eObject) msg() *openapi.EObject {
	return obj.obj
}

func (obj *eObject) setMsg(msg *openapi.EObject) EObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaleObject struct {
	obj *eObject
}

type marshalEObject interface {
	// ToProto marshals EObject to protobuf object *openapi.EObject
	ToProto() (*openapi.EObject, error)
	// ToPbText marshals EObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals EObject to JSON text
	ToJson() (string, error)
}

type unMarshaleObject struct {
	obj *eObject
}

type unMarshalEObject interface {
	// FromProto unmarshals EObject from protobuf object *openapi.EObject
	FromProto(msg *openapi.EObject) (EObject, error)
	// FromPbText unmarshals EObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EObject from JSON text
	FromJson(value string) error
}

func (obj *eObject) Marshal() marshalEObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshaleObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *eObject) Unmarshal() unMarshalEObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaleObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaleObject) ToProto() (*openapi.EObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaleObject) FromProto(msg *openapi.EObject) (EObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaleObject) ToPbText() (string, error) {
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

func (m *unMarshaleObject) FromPbText(value string) error {
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

func (m *marshaleObject) ToYaml() (string, error) {
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

func (m *unMarshaleObject) FromYaml(value string) error {
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

func (m *marshaleObject) ToJson() (string, error) {
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

func (m *unMarshaleObject) FromJson(value string) error {
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

func (obj *eObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eObject) Clone() (EObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEObject()
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

// EObject is description is TBD
type EObject interface {
	Validation
	// msg marshals EObject to protobuf object *openapi.EObject
	// and doesn't set defaults
	msg() *openapi.EObject
	// setMsg unmarshals EObject from protobuf object *openapi.EObject
	// and doesn't set defaults
	setMsg(*openapi.EObject) EObject
	// provides marshal interface
	Marshal() marshalEObject
	// provides unmarshal interface
	Unmarshal() unMarshalEObject
	// validate validates EObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EA returns float32, set in EObject.
	EA() float32
	// SetEA assigns float32 provided by user to EObject
	SetEA(value float32) EObject
	// EB returns float64, set in EObject.
	EB() float64
	// SetEB assigns float64 provided by user to EObject
	SetEB(value float64) EObject
	// Name returns string, set in EObject.
	Name() string
	// SetName assigns string provided by user to EObject
	SetName(value string) EObject
	// HasName checks if Name has been set in EObject
	HasName() bool
	// MParam1 returns string, set in EObject.
	MParam1() string
	// SetMParam1 assigns string provided by user to EObject
	SetMParam1(value string) EObject
	// HasMParam1 checks if MParam1 has been set in EObject
	HasMParam1() bool
	// MParam2 returns string, set in EObject.
	MParam2() string
	// SetMParam2 assigns string provided by user to EObject
	SetMParam2(value string) EObject
	// HasMParam2 checks if MParam2 has been set in EObject
	HasMParam2() bool
}

// description is TBD
// EA returns a float32
func (obj *eObject) EA() float32 {

	return *obj.obj.EA

}

// description is TBD
// SetEA sets the float32 value in the EObject object
func (obj *eObject) SetEA(value float32) EObject {

	obj.obj.EA = &value
	return obj
}

// description is TBD
// EB returns a float64
func (obj *eObject) EB() float64 {

	return *obj.obj.EB

}

// description is TBD
// SetEB sets the float64 value in the EObject object
func (obj *eObject) SetEB(value float64) EObject {

	obj.obj.EB = &value
	return obj
}

// description is TBD
// Name returns a string
func (obj *eObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *eObject) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the EObject object
func (obj *eObject) SetName(value string) EObject {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// MParam1 returns a string
func (obj *eObject) MParam1() string {

	return *obj.obj.MParam1

}

// description is TBD
// MParam1 returns a string
func (obj *eObject) HasMParam1() bool {
	return obj.obj.MParam1 != nil
}

// description is TBD
// SetMParam1 sets the string value in the EObject object
func (obj *eObject) SetMParam1(value string) EObject {

	obj.obj.MParam1 = &value
	return obj
}

// description is TBD
// MParam2 returns a string
func (obj *eObject) MParam2() string {

	return *obj.obj.MParam2

}

// description is TBD
// MParam2 returns a string
func (obj *eObject) HasMParam2() bool {
	return obj.obj.MParam2 != nil
}

// description is TBD
// SetMParam2 sets the string value in the EObject object
func (obj *eObject) SetMParam2(value string) EObject {

	obj.obj.MParam2 = &value
	return obj
}

func (obj *eObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EA is required
	if obj.obj.EA == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EA is required field on interface EObject")
	}

	// EB is required
	if obj.obj.EB == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EB is required field on interface EObject")
	}
}

func (obj *eObject) setDefault() {

}
