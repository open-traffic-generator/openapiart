package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredVal *****
type requiredVal struct {
	validation
	obj          *openapi.RequiredVal
	marshaller   marshalRequiredVal
	unMarshaller unMarshalRequiredVal
}

func NewRequiredVal() RequiredVal {
	obj := requiredVal{obj: &openapi.RequiredVal{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredVal) msg() *openapi.RequiredVal {
	return obj.obj
}

func (obj *requiredVal) setMsg(msg *openapi.RequiredVal) RequiredVal {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredVal struct {
	obj *requiredVal
}

type marshalRequiredVal interface {
	// ToProto marshals RequiredVal to protobuf object *openapi.RequiredVal
	ToProto() (*openapi.RequiredVal, error)
	// ToPbText marshals RequiredVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredVal to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredVal struct {
	obj *requiredVal
}

type unMarshalRequiredVal interface {
	// FromProto unmarshals RequiredVal from protobuf object *openapi.RequiredVal
	FromProto(msg *openapi.RequiredVal) (RequiredVal, error)
	// FromPbText unmarshals RequiredVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredVal from JSON text
	FromJson(value string) error
}

func (obj *requiredVal) Marshal() marshalRequiredVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredVal) Unmarshal() unMarshalRequiredVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredVal) ToProto() (*openapi.RequiredVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredVal) FromProto(msg *openapi.RequiredVal) (RequiredVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredVal) ToPbText() (string, error) {
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

func (m *unMarshalrequiredVal) FromPbText(value string) error {
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

func (m *marshalrequiredVal) ToYaml() (string, error) {
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

func (m *unMarshalrequiredVal) FromYaml(value string) error {
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

func (m *marshalrequiredVal) ToJson() (string, error) {
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

func (m *unMarshalrequiredVal) FromJson(value string) error {
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

func (obj *requiredVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredVal) Clone() (RequiredVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredVal()
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

// RequiredVal is description is TBD
type RequiredVal interface {
	Validation
	// msg marshals RequiredVal to protobuf object *openapi.RequiredVal
	// and doesn't set defaults
	msg() *openapi.RequiredVal
	// setMsg unmarshals RequiredVal from protobuf object *openapi.RequiredVal
	// and doesn't set defaults
	setMsg(*openapi.RequiredVal) RequiredVal
	// provides marshal interface
	Marshal() marshalRequiredVal
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredVal
	// validate validates RequiredVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVal returns int32, set in RequiredVal.
	IntVal() int32
	// SetIntVal assigns int32 provided by user to RequiredVal
	SetIntVal(value int32) RequiredVal
	// NumVal returns float32, set in RequiredVal.
	NumVal() float32
	// SetNumVal assigns float32 provided by user to RequiredVal
	SetNumVal(value float32) RequiredVal
	// StrVal returns string, set in RequiredVal.
	StrVal() string
	// SetStrVal assigns string provided by user to RequiredVal
	SetStrVal(value string) RequiredVal
	// BoolVal returns bool, set in RequiredVal.
	BoolVal() bool
	// SetBoolVal assigns bool provided by user to RequiredVal
	SetBoolVal(value bool) RequiredVal
}

// description is TBD
// IntVal returns a int32
func (obj *requiredVal) IntVal() int32 {

	return *obj.obj.IntVal

}

// description is TBD
// SetIntVal sets the int32 value in the RequiredVal object
func (obj *requiredVal) SetIntVal(value int32) RequiredVal {

	obj.obj.IntVal = &value
	return obj
}

// description is TBD
// NumVal returns a float32
func (obj *requiredVal) NumVal() float32 {

	return *obj.obj.NumVal

}

// description is TBD
// SetNumVal sets the float32 value in the RequiredVal object
func (obj *requiredVal) SetNumVal(value float32) RequiredVal {

	obj.obj.NumVal = &value
	return obj
}

// description is TBD
// StrVal returns a string
func (obj *requiredVal) StrVal() string {

	return *obj.obj.StrVal

}

// description is TBD
// SetStrVal sets the string value in the RequiredVal object
func (obj *requiredVal) SetStrVal(value string) RequiredVal {

	obj.obj.StrVal = &value
	return obj
}

// description is TBD
// BoolVal returns a bool
func (obj *requiredVal) BoolVal() bool {

	return *obj.obj.BoolVal

}

// description is TBD
// SetBoolVal sets the bool value in the RequiredVal object
func (obj *requiredVal) SetBoolVal(value bool) RequiredVal {

	obj.obj.BoolVal = &value
	return obj
}

func (obj *requiredVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// IntVal is required
	if obj.obj.IntVal == nil {
		vObj.validationErrors = append(vObj.validationErrors, "IntVal is required field on interface RequiredVal")
	}

	// NumVal is required
	if obj.obj.NumVal == nil {
		vObj.validationErrors = append(vObj.validationErrors, "NumVal is required field on interface RequiredVal")
	}

	// StrVal is required
	if obj.obj.StrVal == nil {
		vObj.validationErrors = append(vObj.validationErrors, "StrVal is required field on interface RequiredVal")
	}

	// BoolVal is required
	if obj.obj.BoolVal == nil {
		vObj.validationErrors = append(vObj.validationErrors, "BoolVal is required field on interface RequiredVal")
	}
}

func (obj *requiredVal) setDefault() {

}
