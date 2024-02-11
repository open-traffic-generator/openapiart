package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RequiredValArray *****
type requiredValArray struct {
	validation
	obj          *openapi.RequiredValArray
	marshaller   marshalRequiredValArray
	unMarshaller unMarshalRequiredValArray
}

func NewRequiredValArray() RequiredValArray {
	obj := requiredValArray{obj: &openapi.RequiredValArray{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredValArray) msg() *openapi.RequiredValArray {
	return obj.obj
}

func (obj *requiredValArray) setMsg(msg *openapi.RequiredValArray) RequiredValArray {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrequiredValArray struct {
	obj *requiredValArray
}

type marshalRequiredValArray interface {
	// ToProto marshals RequiredValArray to protobuf object *openapi.RequiredValArray
	ToProto() (*openapi.RequiredValArray, error)
	// ToPbText marshals RequiredValArray to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredValArray to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredValArray to JSON text
	ToJson() (string, error)
}

type unMarshalrequiredValArray struct {
	obj *requiredValArray
}

type unMarshalRequiredValArray interface {
	// FromProto unmarshals RequiredValArray from protobuf object *openapi.RequiredValArray
	FromProto(msg *openapi.RequiredValArray) (RequiredValArray, error)
	// FromPbText unmarshals RequiredValArray from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredValArray from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredValArray from JSON text
	FromJson(value string) error
}

func (obj *requiredValArray) Marshal() marshalRequiredValArray {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrequiredValArray{obj: obj}
	}
	return obj.marshaller
}

func (obj *requiredValArray) Unmarshal() unMarshalRequiredValArray {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrequiredValArray{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrequiredValArray) ToProto() (*openapi.RequiredValArray, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrequiredValArray) FromProto(msg *openapi.RequiredValArray) (RequiredValArray, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrequiredValArray) ToPbText() (string, error) {
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

func (m *unMarshalrequiredValArray) FromPbText(value string) error {
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

func (m *marshalrequiredValArray) ToYaml() (string, error) {
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

func (m *unMarshalrequiredValArray) FromYaml(value string) error {
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

func (m *marshalrequiredValArray) ToJson() (string, error) {
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

func (m *unMarshalrequiredValArray) FromJson(value string) error {
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

func (obj *requiredValArray) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredValArray) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredValArray) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredValArray) Clone() (RequiredValArray, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredValArray()
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

// RequiredValArray is description is TBD
type RequiredValArray interface {
	Validation
	// msg marshals RequiredValArray to protobuf object *openapi.RequiredValArray
	// and doesn't set defaults
	msg() *openapi.RequiredValArray
	// setMsg unmarshals RequiredValArray from protobuf object *openapi.RequiredValArray
	// and doesn't set defaults
	setMsg(*openapi.RequiredValArray) RequiredValArray
	// provides marshal interface
	Marshal() marshalRequiredValArray
	// provides unmarshal interface
	Unmarshal() unMarshalRequiredValArray
	// validate validates RequiredValArray
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredValArray, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVals returns []int32, set in RequiredValArray.
	IntVals() []int32
	// SetIntVals assigns []int32 provided by user to RequiredValArray
	SetIntVals(value []int32) RequiredValArray
	// NumVals returns []float32, set in RequiredValArray.
	NumVals() []float32
	// SetNumVals assigns []float32 provided by user to RequiredValArray
	SetNumVals(value []float32) RequiredValArray
	// StrVals returns []string, set in RequiredValArray.
	StrVals() []string
	// SetStrVals assigns []string provided by user to RequiredValArray
	SetStrVals(value []string) RequiredValArray
	// BoolVals returns []bool, set in RequiredValArray.
	BoolVals() []bool
	// SetBoolVals assigns []bool provided by user to RequiredValArray
	SetBoolVals(value []bool) RequiredValArray
}

// description is TBD
// IntVals returns a []int32
func (obj *requiredValArray) IntVals() []int32 {
	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	return obj.obj.IntVals
}

// description is TBD
// SetIntVals sets the []int32 value in the RequiredValArray object
func (obj *requiredValArray) SetIntVals(value []int32) RequiredValArray {

	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	obj.obj.IntVals = value

	return obj
}

// description is TBD
// NumVals returns a []float32
func (obj *requiredValArray) NumVals() []float32 {
	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	return obj.obj.NumVals
}

// description is TBD
// SetNumVals sets the []float32 value in the RequiredValArray object
func (obj *requiredValArray) SetNumVals(value []float32) RequiredValArray {

	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	obj.obj.NumVals = value

	return obj
}

// description is TBD
// StrVals returns a []string
func (obj *requiredValArray) StrVals() []string {
	if obj.obj.StrVals == nil {
		obj.obj.StrVals = make([]string, 0)
	}
	return obj.obj.StrVals
}

// description is TBD
// SetStrVals sets the []string value in the RequiredValArray object
func (obj *requiredValArray) SetStrVals(value []string) RequiredValArray {

	if obj.obj.StrVals == nil {
		obj.obj.StrVals = make([]string, 0)
	}
	obj.obj.StrVals = value

	return obj
}

// description is TBD
// BoolVals returns a []bool
func (obj *requiredValArray) BoolVals() []bool {
	if obj.obj.BoolVals == nil {
		obj.obj.BoolVals = make([]bool, 0)
	}
	return obj.obj.BoolVals
}

// description is TBD
// SetBoolVals sets the []bool value in the RequiredValArray object
func (obj *requiredValArray) SetBoolVals(value []bool) RequiredValArray {

	if obj.obj.BoolVals == nil {
		obj.obj.BoolVals = make([]bool, 0)
	}
	obj.obj.BoolVals = value

	return obj
}

func (obj *requiredValArray) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *requiredValArray) setDefault() {

}
