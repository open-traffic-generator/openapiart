package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BoundaryVal *****
type boundaryVal struct {
	validation
	obj          *openapi.BoundaryVal
	marshaller   marshalBoundaryVal
	unMarshaller unMarshalBoundaryVal
}

func NewBoundaryVal() BoundaryVal {
	obj := boundaryVal{obj: &openapi.BoundaryVal{}}
	obj.setDefault()
	return &obj
}

func (obj *boundaryVal) msg() *openapi.BoundaryVal {
	return obj.obj
}

func (obj *boundaryVal) setMsg(msg *openapi.BoundaryVal) BoundaryVal {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalboundaryVal struct {
	obj *boundaryVal
}

type marshalBoundaryVal interface {
	// ToProto marshals BoundaryVal to protobuf object *openapi.BoundaryVal
	ToProto() (*openapi.BoundaryVal, error)
	// ToPbText marshals BoundaryVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BoundaryVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals BoundaryVal to JSON text
	ToJson() (string, error)
}

type unMarshalboundaryVal struct {
	obj *boundaryVal
}

type unMarshalBoundaryVal interface {
	// FromProto unmarshals BoundaryVal from protobuf object *openapi.BoundaryVal
	FromProto(msg *openapi.BoundaryVal) (BoundaryVal, error)
	// FromPbText unmarshals BoundaryVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BoundaryVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BoundaryVal from JSON text
	FromJson(value string) error
}

func (obj *boundaryVal) Marshal() marshalBoundaryVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalboundaryVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *boundaryVal) Unmarshal() unMarshalBoundaryVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalboundaryVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalboundaryVal) ToProto() (*openapi.BoundaryVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalboundaryVal) FromProto(msg *openapi.BoundaryVal) (BoundaryVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalboundaryVal) ToPbText() (string, error) {
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

func (m *unMarshalboundaryVal) FromPbText(value string) error {
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

func (m *marshalboundaryVal) ToYaml() (string, error) {
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

func (m *unMarshalboundaryVal) FromYaml(value string) error {
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

func (m *marshalboundaryVal) ToJson() (string, error) {
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

func (m *unMarshalboundaryVal) FromJson(value string) error {
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

func (obj *boundaryVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *boundaryVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *boundaryVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *boundaryVal) Clone() (BoundaryVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBoundaryVal()
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

// BoundaryVal is description is TBD
type BoundaryVal interface {
	Validation
	// msg marshals BoundaryVal to protobuf object *openapi.BoundaryVal
	// and doesn't set defaults
	msg() *openapi.BoundaryVal
	// setMsg unmarshals BoundaryVal from protobuf object *openapi.BoundaryVal
	// and doesn't set defaults
	setMsg(*openapi.BoundaryVal) BoundaryVal
	// provides marshal interface
	Marshal() marshalBoundaryVal
	// provides unmarshal interface
	Unmarshal() unMarshalBoundaryVal
	// validate validates BoundaryVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BoundaryVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IntVal returns int32, set in BoundaryVal.
	IntVal() int32
	// SetIntVal assigns int32 provided by user to BoundaryVal
	SetIntVal(value int32) BoundaryVal
	// HasIntVal checks if IntVal has been set in BoundaryVal
	HasIntVal() bool
	// NumVal returns float32, set in BoundaryVal.
	NumVal() float32
	// SetNumVal assigns float32 provided by user to BoundaryVal
	SetNumVal(value float32) BoundaryVal
	// HasNumVal checks if NumVal has been set in BoundaryVal
	HasNumVal() bool
}

// description is TBD
// IntVal returns a int32
func (obj *boundaryVal) IntVal() int32 {

	return *obj.obj.IntVal

}

// description is TBD
// IntVal returns a int32
func (obj *boundaryVal) HasIntVal() bool {
	return obj.obj.IntVal != nil
}

// description is TBD
// SetIntVal sets the int32 value in the BoundaryVal object
func (obj *boundaryVal) SetIntVal(value int32) BoundaryVal {

	obj.obj.IntVal = &value
	return obj
}

// description is TBD
// NumVal returns a float32
func (obj *boundaryVal) NumVal() float32 {

	return *obj.obj.NumVal

}

// description is TBD
// NumVal returns a float32
func (obj *boundaryVal) HasNumVal() bool {
	return obj.obj.NumVal != nil
}

// description is TBD
// SetNumVal sets the float32 value in the BoundaryVal object
func (obj *boundaryVal) SetNumVal(value float32) BoundaryVal {

	obj.obj.NumVal = &value
	return obj
}

func (obj *boundaryVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.IntVal != nil {

		if *obj.obj.IntVal < 5 || *obj.obj.IntVal > 100 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5 <= BoundaryVal.IntVal <= 100 but Got %d", *obj.obj.IntVal))
		}

	}

	if obj.obj.NumVal != nil {

		if *obj.obj.NumVal < 5.0 || *obj.obj.NumVal > 100.0 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("5.0 <= BoundaryVal.NumVal <= 100.0 but Got %f", *obj.obj.NumVal))
		}

	}

}

func (obj *boundaryVal) setDefault() {
	if obj.obj.IntVal == nil {
		obj.SetIntVal(50)
	}
	if obj.obj.NumVal == nil {
		obj.SetNumVal(50.05)
	}

}
