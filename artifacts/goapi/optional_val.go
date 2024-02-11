package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** OptionalVal *****
type optionalVal struct {
	validation
	obj          *openapi.OptionalVal
	marshaller   marshalOptionalVal
	unMarshaller unMarshalOptionalVal
}

func NewOptionalVal() OptionalVal {
	obj := optionalVal{obj: &openapi.OptionalVal{}}
	obj.setDefault()
	return &obj
}

func (obj *optionalVal) msg() *openapi.OptionalVal {
	return obj.obj
}

func (obj *optionalVal) setMsg(msg *openapi.OptionalVal) OptionalVal {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaloptionalVal struct {
	obj *optionalVal
}

type marshalOptionalVal interface {
	// ToProto marshals OptionalVal to protobuf object *openapi.OptionalVal
	ToProto() (*openapi.OptionalVal, error)
	// ToPbText marshals OptionalVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals OptionalVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals OptionalVal to JSON text
	ToJson() (string, error)
}

type unMarshaloptionalVal struct {
	obj *optionalVal
}

type unMarshalOptionalVal interface {
	// FromProto unmarshals OptionalVal from protobuf object *openapi.OptionalVal
	FromProto(msg *openapi.OptionalVal) (OptionalVal, error)
	// FromPbText unmarshals OptionalVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals OptionalVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals OptionalVal from JSON text
	FromJson(value string) error
}

func (obj *optionalVal) Marshal() marshalOptionalVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshaloptionalVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *optionalVal) Unmarshal() unMarshalOptionalVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaloptionalVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaloptionalVal) ToProto() (*openapi.OptionalVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaloptionalVal) FromProto(msg *openapi.OptionalVal) (OptionalVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaloptionalVal) ToPbText() (string, error) {
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

func (m *unMarshaloptionalVal) FromPbText(value string) error {
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

func (m *marshaloptionalVal) ToYaml() (string, error) {
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

func (m *unMarshaloptionalVal) FromYaml(value string) error {
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

func (m *marshaloptionalVal) ToJson() (string, error) {
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

func (m *unMarshaloptionalVal) FromJson(value string) error {
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

func (obj *optionalVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *optionalVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *optionalVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *optionalVal) Clone() (OptionalVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOptionalVal()
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

// OptionalVal is description is TBD
type OptionalVal interface {
	Validation
	// msg marshals OptionalVal to protobuf object *openapi.OptionalVal
	// and doesn't set defaults
	msg() *openapi.OptionalVal
	// setMsg unmarshals OptionalVal from protobuf object *openapi.OptionalVal
	// and doesn't set defaults
	setMsg(*openapi.OptionalVal) OptionalVal
	// provides marshal interface
	Marshal() marshalOptionalVal
	// provides unmarshal interface
	Unmarshal() unMarshalOptionalVal
	// validate validates OptionalVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (OptionalVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVal returns int32, set in OptionalVal.
	IntVal() int32
	// SetIntVal assigns int32 provided by user to OptionalVal
	SetIntVal(value int32) OptionalVal
	// HasIntVal checks if IntVal has been set in OptionalVal
	HasIntVal() bool
	// NumVal returns float32, set in OptionalVal.
	NumVal() float32
	// SetNumVal assigns float32 provided by user to OptionalVal
	SetNumVal(value float32) OptionalVal
	// HasNumVal checks if NumVal has been set in OptionalVal
	HasNumVal() bool
	// StrVal returns string, set in OptionalVal.
	StrVal() string
	// SetStrVal assigns string provided by user to OptionalVal
	SetStrVal(value string) OptionalVal
	// HasStrVal checks if StrVal has been set in OptionalVal
	HasStrVal() bool
	// BoolVal returns bool, set in OptionalVal.
	BoolVal() bool
	// SetBoolVal assigns bool provided by user to OptionalVal
	SetBoolVal(value bool) OptionalVal
	// HasBoolVal checks if BoolVal has been set in OptionalVal
	HasBoolVal() bool
}

// description is TBD
// IntVal returns a int32
func (obj *optionalVal) IntVal() int32 {

	return *obj.obj.IntVal

}

// description is TBD
// IntVal returns a int32
func (obj *optionalVal) HasIntVal() bool {
	return obj.obj.IntVal != nil
}

// description is TBD
// SetIntVal sets the int32 value in the OptionalVal object
func (obj *optionalVal) SetIntVal(value int32) OptionalVal {

	obj.obj.IntVal = &value
	return obj
}

// description is TBD
// NumVal returns a float32
func (obj *optionalVal) NumVal() float32 {

	return *obj.obj.NumVal

}

// description is TBD
// NumVal returns a float32
func (obj *optionalVal) HasNumVal() bool {
	return obj.obj.NumVal != nil
}

// description is TBD
// SetNumVal sets the float32 value in the OptionalVal object
func (obj *optionalVal) SetNumVal(value float32) OptionalVal {

	obj.obj.NumVal = &value
	return obj
}

// description is TBD
// StrVal returns a string
func (obj *optionalVal) StrVal() string {

	return *obj.obj.StrVal

}

// description is TBD
// StrVal returns a string
func (obj *optionalVal) HasStrVal() bool {
	return obj.obj.StrVal != nil
}

// description is TBD
// SetStrVal sets the string value in the OptionalVal object
func (obj *optionalVal) SetStrVal(value string) OptionalVal {

	obj.obj.StrVal = &value
	return obj
}

// description is TBD
// BoolVal returns a bool
func (obj *optionalVal) BoolVal() bool {

	return *obj.obj.BoolVal

}

// description is TBD
// BoolVal returns a bool
func (obj *optionalVal) HasBoolVal() bool {
	return obj.obj.BoolVal != nil
}

// description is TBD
// SetBoolVal sets the bool value in the OptionalVal object
func (obj *optionalVal) SetBoolVal(value bool) OptionalVal {

	obj.obj.BoolVal = &value
	return obj
}

func (obj *optionalVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *optionalVal) setDefault() {
	if obj.obj.IntVal == nil {
		obj.SetIntVal(50)
	}
	if obj.obj.NumVal == nil {
		obj.SetNumVal(50.05)
	}
	if obj.obj.StrVal == nil {
		obj.SetStrVal("default_str_val")
	}
	if obj.obj.BoolVal == nil {
		obj.SetBoolVal(true)
	}

}
