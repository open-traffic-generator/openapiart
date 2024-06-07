package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowMetric *****
type flowMetric struct {
	validation
	obj          *openapi.FlowMetric
	marshaller   marshalFlowMetric
	unMarshaller unMarshalFlowMetric
}

func NewFlowMetric() FlowMetric {
	obj := flowMetric{obj: &openapi.FlowMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetric) msg() *openapi.FlowMetric {
	return obj.obj
}

func (obj *flowMetric) setMsg(msg *openapi.FlowMetric) FlowMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowMetric struct {
	obj *flowMetric
}

type marshalFlowMetric interface {
	// ToProto marshals FlowMetric to protobuf object *openapi.FlowMetric
	ToProto() (*openapi.FlowMetric, error)
	// ToPbText marshals FlowMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetric to JSON text
	ToJson() (string, error)
}

type unMarshalflowMetric struct {
	obj *flowMetric
}

type unMarshalFlowMetric interface {
	// FromProto unmarshals FlowMetric from protobuf object *openapi.FlowMetric
	FromProto(msg *openapi.FlowMetric) (FlowMetric, error)
	// FromPbText unmarshals FlowMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetric from JSON text
	FromJson(value string) error
}

func (obj *flowMetric) Marshal() marshalFlowMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowMetric) Unmarshal() unMarshalFlowMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowMetric) ToProto() (*openapi.FlowMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowMetric) FromProto(msg *openapi.FlowMetric) (FlowMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowMetric) ToPbText() (string, error) {
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

func (m *unMarshalflowMetric) FromPbText(value string) error {
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

func (m *marshalflowMetric) ToYaml() (string, error) {
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

func (m *unMarshalflowMetric) FromYaml(value string) error {
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

func (m *marshalflowMetric) ToJson() (string, error) {
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

func (m *unMarshalflowMetric) FromJson(value string) error {
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

func (obj *flowMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetric) Clone() (FlowMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetric()
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

// FlowMetric is description is TBD
type FlowMetric interface {
	Validation
	// msg marshals FlowMetric to protobuf object *openapi.FlowMetric
	// and doesn't set defaults
	msg() *openapi.FlowMetric
	// setMsg unmarshals FlowMetric from protobuf object *openapi.FlowMetric
	// and doesn't set defaults
	setMsg(*openapi.FlowMetric) FlowMetric
	// provides marshal interface
	Marshal() marshalFlowMetric
	// provides unmarshal interface
	Unmarshal() unMarshalFlowMetric
	// validate validates FlowMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowMetric.
	Name() string
	// SetName assigns string provided by user to FlowMetric
	SetName(value string) FlowMetric
	// TxFrames returns float64, set in FlowMetric.
	TxFrames() float64
	// SetTxFrames assigns float64 provided by user to FlowMetric
	SetTxFrames(value float64) FlowMetric
	// RxFrames returns float64, set in FlowMetric.
	RxFrames() float64
	// SetRxFrames assigns float64 provided by user to FlowMetric
	SetRxFrames(value float64) FlowMetric
}

// description is TBD
// Name returns a string
func (obj *flowMetric) Name() string {

	return *obj.obj.Name

}

// description is TBD
// SetName sets the string value in the FlowMetric object
func (obj *flowMetric) SetName(value string) FlowMetric {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// TxFrames returns a float64
func (obj *flowMetric) TxFrames() float64 {

	return *obj.obj.TxFrames

}

// description is TBD
// SetTxFrames sets the float64 value in the FlowMetric object
func (obj *flowMetric) SetTxFrames(value float64) FlowMetric {

	obj.obj.TxFrames = &value
	return obj
}

// description is TBD
// RxFrames returns a float64
func (obj *flowMetric) RxFrames() float64 {

	return *obj.obj.RxFrames

}

// description is TBD
// SetRxFrames sets the float64 value in the FlowMetric object
func (obj *flowMetric) SetRxFrames(value float64) FlowMetric {

	obj.obj.RxFrames = &value
	return obj
}

func (obj *flowMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface FlowMetric")
	}

	// TxFrames is required
	if obj.obj.TxFrames == nil {
		vObj.validationErrors = append(vObj.validationErrors, "TxFrames is required field on interface FlowMetric")
	}

	// RxFrames is required
	if obj.obj.RxFrames == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RxFrames is required field on interface FlowMetric")
	}
}

func (obj *flowMetric) setDefault() {

}
