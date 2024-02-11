package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LeafVal *****
type leafVal struct {
	validation
	obj          *openapi.LeafVal
	marshaller   marshalLeafVal
	unMarshaller unMarshalLeafVal
}

func NewLeafVal() LeafVal {
	obj := leafVal{obj: &openapi.LeafVal{}}
	obj.setDefault()
	return &obj
}

func (obj *leafVal) msg() *openapi.LeafVal {
	return obj.obj
}

func (obj *leafVal) setMsg(msg *openapi.LeafVal) LeafVal {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalleafVal struct {
	obj *leafVal
}

type marshalLeafVal interface {
	// ToProto marshals LeafVal to protobuf object *openapi.LeafVal
	ToProto() (*openapi.LeafVal, error)
	// ToPbText marshals LeafVal to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LeafVal to YAML text
	ToYaml() (string, error)
	// ToJson marshals LeafVal to JSON text
	ToJson() (string, error)
}

type unMarshalleafVal struct {
	obj *leafVal
}

type unMarshalLeafVal interface {
	// FromProto unmarshals LeafVal from protobuf object *openapi.LeafVal
	FromProto(msg *openapi.LeafVal) (LeafVal, error)
	// FromPbText unmarshals LeafVal from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LeafVal from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LeafVal from JSON text
	FromJson(value string) error
}

func (obj *leafVal) Marshal() marshalLeafVal {
	if obj.marshaller == nil {
		obj.marshaller = &marshalleafVal{obj: obj}
	}
	return obj.marshaller
}

func (obj *leafVal) Unmarshal() unMarshalLeafVal {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalleafVal{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalleafVal) ToProto() (*openapi.LeafVal, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalleafVal) FromProto(msg *openapi.LeafVal) (LeafVal, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalleafVal) ToPbText() (string, error) {
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

func (m *unMarshalleafVal) FromPbText(value string) error {
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

func (m *marshalleafVal) ToYaml() (string, error) {
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

func (m *unMarshalleafVal) FromYaml(value string) error {
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

func (m *marshalleafVal) ToJson() (string, error) {
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

func (m *unMarshalleafVal) FromJson(value string) error {
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

func (obj *leafVal) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *leafVal) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *leafVal) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *leafVal) Clone() (LeafVal, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLeafVal()
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

// LeafVal is description is TBD
type LeafVal interface {
	Validation
	// msg marshals LeafVal to protobuf object *openapi.LeafVal
	// and doesn't set defaults
	msg() *openapi.LeafVal
	// setMsg unmarshals LeafVal from protobuf object *openapi.LeafVal
	// and doesn't set defaults
	setMsg(*openapi.LeafVal) LeafVal
	// provides marshal interface
	Marshal() marshalLeafVal
	// provides unmarshal interface
	Unmarshal() unMarshalLeafVal
	// validate validates LeafVal
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LeafVal, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in LeafVal.
	Name() string
	// SetName assigns string provided by user to LeafVal
	SetName(value string) LeafVal
	// HasName checks if Name has been set in LeafVal
	HasName() bool
	// Value returns int32, set in LeafVal.
	Value() int32
	// SetValue assigns int32 provided by user to LeafVal
	SetValue(value int32) LeafVal
	// HasValue checks if Value has been set in LeafVal
	HasValue() bool
}

// description is TBD
// Name returns a string
func (obj *leafVal) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *leafVal) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the LeafVal object
func (obj *leafVal) SetName(value string) LeafVal {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// Value returns a int32
func (obj *leafVal) Value() int32 {

	return *obj.obj.Value

}

// description is TBD
// Value returns a int32
func (obj *leafVal) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the int32 value in the LeafVal object
func (obj *leafVal) SetValue(value int32) LeafVal {

	obj.obj.Value = &value
	return obj
}

func (obj *leafVal) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *leafVal) setDefault() {

}
