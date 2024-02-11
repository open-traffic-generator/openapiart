package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Layer1Ieee802X *****
type layer1Ieee802X struct {
	validation
	obj          *openapi.Layer1Ieee802X
	marshaller   marshalLayer1Ieee802X
	unMarshaller unMarshalLayer1Ieee802X
}

func NewLayer1Ieee802X() Layer1Ieee802X {
	obj := layer1Ieee802X{obj: &openapi.Layer1Ieee802X{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1Ieee802X) msg() *openapi.Layer1Ieee802X {
	return obj.obj
}

func (obj *layer1Ieee802X) setMsg(msg *openapi.Layer1Ieee802X) Layer1Ieee802X {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallayer1Ieee802X struct {
	obj *layer1Ieee802X
}

type marshalLayer1Ieee802X interface {
	// ToProto marshals Layer1Ieee802X to protobuf object *openapi.Layer1Ieee802X
	ToProto() (*openapi.Layer1Ieee802X, error)
	// ToPbText marshals Layer1Ieee802X to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1Ieee802X to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1Ieee802X to JSON text
	ToJson() (string, error)
}

type unMarshallayer1Ieee802X struct {
	obj *layer1Ieee802X
}

type unMarshalLayer1Ieee802X interface {
	// FromProto unmarshals Layer1Ieee802X from protobuf object *openapi.Layer1Ieee802X
	FromProto(msg *openapi.Layer1Ieee802X) (Layer1Ieee802X, error)
	// FromPbText unmarshals Layer1Ieee802X from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1Ieee802X from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1Ieee802X from JSON text
	FromJson(value string) error
}

func (obj *layer1Ieee802X) Marshal() marshalLayer1Ieee802X {
	if obj.marshaller == nil {
		obj.marshaller = &marshallayer1Ieee802X{obj: obj}
	}
	return obj.marshaller
}

func (obj *layer1Ieee802X) Unmarshal() unMarshalLayer1Ieee802X {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallayer1Ieee802X{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallayer1Ieee802X) ToProto() (*openapi.Layer1Ieee802X, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallayer1Ieee802X) FromProto(msg *openapi.Layer1Ieee802X) (Layer1Ieee802X, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallayer1Ieee802X) ToPbText() (string, error) {
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

func (m *unMarshallayer1Ieee802X) FromPbText(value string) error {
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

func (m *marshallayer1Ieee802X) ToYaml() (string, error) {
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

func (m *unMarshallayer1Ieee802X) FromYaml(value string) error {
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

func (m *marshallayer1Ieee802X) ToJson() (string, error) {
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

func (m *unMarshallayer1Ieee802X) FromJson(value string) error {
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

func (obj *layer1Ieee802X) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1Ieee802X) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1Ieee802X) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1Ieee802X) Clone() (Layer1Ieee802X, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1Ieee802X()
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

// Layer1Ieee802X is description is TBD
type Layer1Ieee802X interface {
	Validation
	// msg marshals Layer1Ieee802X to protobuf object *openapi.Layer1Ieee802X
	// and doesn't set defaults
	msg() *openapi.Layer1Ieee802X
	// setMsg unmarshals Layer1Ieee802X from protobuf object *openapi.Layer1Ieee802X
	// and doesn't set defaults
	setMsg(*openapi.Layer1Ieee802X) Layer1Ieee802X
	// provides marshal interface
	Marshal() marshalLayer1Ieee802X
	// provides unmarshal interface
	Unmarshal() unMarshalLayer1Ieee802X
	// validate validates Layer1Ieee802X
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1Ieee802X, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowControl returns bool, set in Layer1Ieee802X.
	FlowControl() bool
	// SetFlowControl assigns bool provided by user to Layer1Ieee802X
	SetFlowControl(value bool) Layer1Ieee802X
	// HasFlowControl checks if FlowControl has been set in Layer1Ieee802X
	HasFlowControl() bool
}

// description is TBD
// FlowControl returns a bool
func (obj *layer1Ieee802X) FlowControl() bool {

	return *obj.obj.FlowControl

}

// description is TBD
// FlowControl returns a bool
func (obj *layer1Ieee802X) HasFlowControl() bool {
	return obj.obj.FlowControl != nil
}

// description is TBD
// SetFlowControl sets the bool value in the Layer1Ieee802X object
func (obj *layer1Ieee802X) SetFlowControl(value bool) Layer1Ieee802X {

	obj.obj.FlowControl = &value
	return obj
}

func (obj *layer1Ieee802X) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *layer1Ieee802X) setDefault() {

}
