package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BoundaryValArray *****
type boundaryValArray struct {
	validation
	obj          *openapi.BoundaryValArray
	marshaller   marshalBoundaryValArray
	unMarshaller unMarshalBoundaryValArray
}

func NewBoundaryValArray() BoundaryValArray {
	obj := boundaryValArray{obj: &openapi.BoundaryValArray{}}
	obj.setDefault()
	return &obj
}

func (obj *boundaryValArray) msg() *openapi.BoundaryValArray {
	return obj.obj
}

func (obj *boundaryValArray) setMsg(msg *openapi.BoundaryValArray) BoundaryValArray {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalboundaryValArray struct {
	obj *boundaryValArray
}

type marshalBoundaryValArray interface {
	// ToProto marshals BoundaryValArray to protobuf object *openapi.BoundaryValArray
	ToProto() (*openapi.BoundaryValArray, error)
	// ToPbText marshals BoundaryValArray to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BoundaryValArray to YAML text
	ToYaml() (string, error)
	// ToJson marshals BoundaryValArray to JSON text
	ToJson() (string, error)
}

type unMarshalboundaryValArray struct {
	obj *boundaryValArray
}

type unMarshalBoundaryValArray interface {
	// FromProto unmarshals BoundaryValArray from protobuf object *openapi.BoundaryValArray
	FromProto(msg *openapi.BoundaryValArray) (BoundaryValArray, error)
	// FromPbText unmarshals BoundaryValArray from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BoundaryValArray from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BoundaryValArray from JSON text
	FromJson(value string) error
}

func (obj *boundaryValArray) Marshal() marshalBoundaryValArray {
	if obj.marshaller == nil {
		obj.marshaller = &marshalboundaryValArray{obj: obj}
	}
	return obj.marshaller
}

func (obj *boundaryValArray) Unmarshal() unMarshalBoundaryValArray {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalboundaryValArray{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalboundaryValArray) ToProto() (*openapi.BoundaryValArray, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalboundaryValArray) FromProto(msg *openapi.BoundaryValArray) (BoundaryValArray, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalboundaryValArray) ToPbText() (string, error) {
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

func (m *unMarshalboundaryValArray) FromPbText(value string) error {
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

func (m *marshalboundaryValArray) ToYaml() (string, error) {
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

func (m *unMarshalboundaryValArray) FromYaml(value string) error {
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

func (m *marshalboundaryValArray) ToJson() (string, error) {
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

func (m *unMarshalboundaryValArray) FromJson(value string) error {
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

func (obj *boundaryValArray) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *boundaryValArray) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *boundaryValArray) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *boundaryValArray) Clone() (BoundaryValArray, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBoundaryValArray()
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

// BoundaryValArray is description is TBD
type BoundaryValArray interface {
	Validation
	// msg marshals BoundaryValArray to protobuf object *openapi.BoundaryValArray
	// and doesn't set defaults
	msg() *openapi.BoundaryValArray
	// setMsg unmarshals BoundaryValArray from protobuf object *openapi.BoundaryValArray
	// and doesn't set defaults
	setMsg(*openapi.BoundaryValArray) BoundaryValArray
	// provides marshal interface
	Marshal() marshalBoundaryValArray
	// provides unmarshal interface
	Unmarshal() unMarshalBoundaryValArray
	// validate validates BoundaryValArray
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BoundaryValArray, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVals returns []int32, set in BoundaryValArray.
	IntVals() []int32
	// SetIntVals assigns []int32 provided by user to BoundaryValArray
	SetIntVals(value []int32) BoundaryValArray
	// NumVals returns []float32, set in BoundaryValArray.
	NumVals() []float32
	// SetNumVals assigns []float32 provided by user to BoundaryValArray
	SetNumVals(value []float32) BoundaryValArray
}

// description is TBD
// IntVals returns a []int32
func (obj *boundaryValArray) IntVals() []int32 {
	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	return obj.obj.IntVals
}

// description is TBD
// SetIntVals sets the []int32 value in the BoundaryValArray object
func (obj *boundaryValArray) SetIntVals(value []int32) BoundaryValArray {

	if obj.obj.IntVals == nil {
		obj.obj.IntVals = make([]int32, 0)
	}
	obj.obj.IntVals = value

	return obj
}

// description is TBD
// NumVals returns a []float32
func (obj *boundaryValArray) NumVals() []float32 {
	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	return obj.obj.NumVals
}

// description is TBD
// SetNumVals sets the []float32 value in the BoundaryValArray object
func (obj *boundaryValArray) SetNumVals(value []float32) BoundaryValArray {

	if obj.obj.NumVals == nil {
		obj.obj.NumVals = make([]float32, 0)
	}
	obj.obj.NumVals = value

	return obj
}

func (obj *boundaryValArray) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntVals != nil {

		for _, item := range obj.obj.IntVals {
			if item < 5 || item > 100 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("5 <= BoundaryValArray.IntVals <= 100 but Got %d", item))
			}

		}

	}

	if obj.obj.NumVals != nil {

		for _, item := range obj.obj.NumVals {
			if item < 5.0 || item > 100.0 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("5.0 <= BoundaryValArray.NumVals <= 100.0 but Got %f", item))
			}

		}

	}

}

func (obj *boundaryValArray) setDefault() {

}
