package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PortMetric *****
type portMetric struct {
	validation
	obj          *openapi.PortMetric
	marshaller   marshalPortMetric
	unMarshaller unMarshalPortMetric
}

func NewPortMetric() PortMetric {
	obj := portMetric{obj: &openapi.PortMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *portMetric) msg() *openapi.PortMetric {
	return obj.obj
}

func (obj *portMetric) setMsg(msg *openapi.PortMetric) PortMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalportMetric struct {
	obj *portMetric
}

type marshalPortMetric interface {
	// ToProto marshals PortMetric to protobuf object *openapi.PortMetric
	ToProto() (*openapi.PortMetric, error)
	// ToPbText marshals PortMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortMetric to JSON text
	ToJson() (string, error)
}

type unMarshalportMetric struct {
	obj *portMetric
}

type unMarshalPortMetric interface {
	// FromProto unmarshals PortMetric from protobuf object *openapi.PortMetric
	FromProto(msg *openapi.PortMetric) (PortMetric, error)
	// FromPbText unmarshals PortMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortMetric from JSON text
	FromJson(value string) error
}

func (obj *portMetric) Marshal() marshalPortMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalportMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *portMetric) Unmarshal() unMarshalPortMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalportMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalportMetric) ToProto() (*openapi.PortMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalportMetric) FromProto(msg *openapi.PortMetric) (PortMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalportMetric) ToPbText() (string, error) {
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

func (m *unMarshalportMetric) FromPbText(value string) error {
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

func (m *marshalportMetric) ToYaml() (string, error) {
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

func (m *unMarshalportMetric) FromYaml(value string) error {
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

func (m *marshalportMetric) ToJson() (string, error) {
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

func (m *unMarshalportMetric) FromJson(value string) error {
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

func (obj *portMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portMetric) Clone() (PortMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortMetric()
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

// PortMetric is description is TBD
type PortMetric interface {
	Validation
	// msg marshals PortMetric to protobuf object *openapi.PortMetric
	// and doesn't set defaults
	msg() *openapi.PortMetric
	// setMsg unmarshals PortMetric from protobuf object *openapi.PortMetric
	// and doesn't set defaults
	setMsg(*openapi.PortMetric) PortMetric
	// provides marshal interface
	Marshal() marshalPortMetric
	// provides unmarshal interface
	Unmarshal() unMarshalPortMetric
	// validate validates PortMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PortMetric.
	Name() string
	// SetName assigns string provided by user to PortMetric
	SetName(value string) PortMetric
	// TxFrames returns float64, set in PortMetric.
	TxFrames() float64
	// SetTxFrames assigns float64 provided by user to PortMetric
	SetTxFrames(value float64) PortMetric
	// RxFrames returns float64, set in PortMetric.
	RxFrames() float64
	// SetRxFrames assigns float64 provided by user to PortMetric
	SetRxFrames(value float64) PortMetric
}

// description is TBD
// Name returns a string
func (obj *portMetric) Name() string {

	return *obj.obj.Name

}

// description is TBD
// SetName sets the string value in the PortMetric object
func (obj *portMetric) SetName(value string) PortMetric {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// TxFrames returns a float64
func (obj *portMetric) TxFrames() float64 {

	return *obj.obj.TxFrames

}

// description is TBD
// SetTxFrames sets the float64 value in the PortMetric object
func (obj *portMetric) SetTxFrames(value float64) PortMetric {

	obj.obj.TxFrames = &value
	return obj
}

// description is TBD
// RxFrames returns a float64
func (obj *portMetric) RxFrames() float64 {

	return *obj.obj.RxFrames

}

// description is TBD
// SetRxFrames sets the float64 value in the PortMetric object
func (obj *portMetric) SetRxFrames(value float64) PortMetric {

	obj.obj.RxFrames = &value
	return obj
}

func (obj *portMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PortMetric")
	}

	// TxFrames is required
	if obj.obj.TxFrames == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxFrames is required field on interface PortMetric")
	}

	// RxFrames is required
	if obj.obj.RxFrames == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RxFrames is required field on interface PortMetric")
	}
}

func (obj *portMetric) setDefault() {

}
