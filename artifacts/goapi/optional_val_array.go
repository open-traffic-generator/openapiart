package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** OptionalValArray *****
type optionalValArray struct {
	validation
	obj          *openapi.OptionalValArray
	marshaller   marshalOptionalValArray
	unMarshaller unMarshalOptionalValArray
}

func NewOptionalValArray() OptionalValArray {
	obj := optionalValArray{obj: &openapi.OptionalValArray{}}
	obj.setDefault()
	return &obj
}

func (obj *optionalValArray) msg() *openapi.OptionalValArray {
	return obj.obj
}

func (obj *optionalValArray) setMsg(msg *openapi.OptionalValArray) OptionalValArray {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshaloptionalValArray struct {
	obj *optionalValArray
}

type marshalOptionalValArray interface {
	// ToProto marshals OptionalValArray to protobuf object *openapi.OptionalValArray
	ToProto() (*openapi.OptionalValArray, error)
	// ToPbText marshals OptionalValArray to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals OptionalValArray to YAML text
	ToYaml() (string, error)
	// ToJson marshals OptionalValArray to JSON text
	ToJson() (string, error)
}

type unMarshaloptionalValArray struct {
	obj *optionalValArray
}

type unMarshalOptionalValArray interface {
	// FromProto unmarshals OptionalValArray from protobuf object *openapi.OptionalValArray
	FromProto(msg *openapi.OptionalValArray) (OptionalValArray, error)
	// FromPbText unmarshals OptionalValArray from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals OptionalValArray from YAML text
	FromYaml(value string) error
	// FromJson unmarshals OptionalValArray from JSON text
	FromJson(value string) error
}

func (obj *optionalValArray) Marshal() marshalOptionalValArray {
	if obj.marshaller == nil {
		obj.marshaller = &marshaloptionalValArray{obj: obj}
	}
	return obj.marshaller
}

func (obj *optionalValArray) Unmarshal() unMarshalOptionalValArray {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaloptionalValArray{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaloptionalValArray) ToProto() (*openapi.OptionalValArray, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaloptionalValArray) FromProto(msg *openapi.OptionalValArray) (OptionalValArray, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaloptionalValArray) ToPbText() (string, error) {
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

func (m *unMarshaloptionalValArray) FromPbText(value string) error {
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

func (m *marshaloptionalValArray) ToYaml() (string, error) {
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

func (m *unMarshaloptionalValArray) FromYaml(value string) error {
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

func (m *marshaloptionalValArray) ToJson() (string, error) {
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

func (m *unMarshaloptionalValArray) FromJson(value string) error {
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

func (obj *optionalValArray) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *optionalValArray) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *optionalValArray) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *optionalValArray) Clone() (OptionalValArray, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOptionalValArray()
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

// OptionalValArray is description is TBD
type OptionalValArray interface {
	Validation
	// msg marshals OptionalValArray to protobuf object *openapi.OptionalValArray
	// and doesn't set defaults
	msg() *openapi.OptionalValArray
	// setMsg unmarshals OptionalValArray from protobuf object *openapi.OptionalValArray
	// and doesn't set defaults
	setMsg(*openapi.OptionalValArray) OptionalValArray
	// provides marshal interface
	Marshal() marshalOptionalValArray
	// provides unmarshal interface
	Unmarshal() unMarshalOptionalValArray
	// validate validates OptionalValArray
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (OptionalValArray, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVals returns []int32, set in OptionalValArray.
	IntVals() []int32
	// SetIntVals assigns []int32 provided by user to OptionalValArray
	SetIntVals(value []int32) OptionalValArray
	// NumVals returns []float32, set in OptionalValArray.
	NumVals() []float32
	// SetNumVals assigns []float32 provided by user to OptionalValArray
	SetNumVals(value []float32) OptionalValArray
	// StrVals returns []string, set in OptionalValArray.
	StrVals() []string
	// SetStrVals assigns []string provided by user to OptionalValArray
	SetStrVals(value []string) OptionalValArray
	// BoolVals returns []bool, set in OptionalValArray.
	BoolVals() []bool
	// SetBoolVals assigns []bool provided by user to OptionalValArray
	SetBoolVals(value []bool) OptionalValArray
}

// description is TBD
// IntVals returns a []int32
func (obj *optionalValArray) IntVals() []int32 {
	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	return obj.obj.IntVals
}

// description is TBD
// SetIntVals sets the []int32 value in the OptionalValArray object
func (obj *optionalValArray) SetIntVals(value []int32) OptionalValArray {

	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	obj.obj.IntVals = value

	return obj
}

// description is TBD
// NumVals returns a []float32
func (obj *optionalValArray) NumVals() []float32 {
	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	return obj.obj.NumVals
}

// description is TBD
// SetNumVals sets the []float32 value in the OptionalValArray object
func (obj *optionalValArray) SetNumVals(value []float32) OptionalValArray {

	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	obj.obj.NumVals = value

	return obj
}

// description is TBD
// StrVals returns a []string
func (obj *optionalValArray) StrVals() []string {
	if obj.obj.StrVals == nil {
		obj.obj.StrVals = make([]string, 0)
	}
	return obj.obj.StrVals
}

// description is TBD
// SetStrVals sets the []string value in the OptionalValArray object
func (obj *optionalValArray) SetStrVals(value []string) OptionalValArray {

	if obj.obj.StrVals == nil {
		obj.obj.StrVals = make([]string, 0)
	}
	obj.obj.StrVals = value

	return obj
}

// description is TBD
// BoolVals returns a []bool
func (obj *optionalValArray) BoolVals() []bool {
	if obj.obj.BoolVals == nil {
		obj.obj.BoolVals = make([]bool, 0)
	}
	return obj.obj.BoolVals
}

// description is TBD
// SetBoolVals sets the []bool value in the OptionalValArray object
func (obj *optionalValArray) SetBoolVals(value []bool) OptionalValArray {

	if obj.obj.BoolVals == nil {
		obj.obj.BoolVals = make([]bool, 0)
	}
	obj.obj.BoolVals = value

	return obj
}

func (obj *optionalValArray) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *optionalValArray) setDefault() {
	if obj.obj.IntVals == nil {
		obj.SetIntVals([]int32{10, 20})
	}
	if obj.obj.NumVals == nil {
		obj.SetNumVals([]float32{10.01, 20.02})
	}
	if obj.obj.StrVals == nil {
		obj.SetStrVals([]string{"first_str", "second_str"})
	}

}
