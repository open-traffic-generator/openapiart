package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Metrics *****
type metrics struct {
	validation
	obj          *openapi.Metrics
	marshaller   marshalMetrics
	unMarshaller unMarshalMetrics
	portsHolder  MetricsPortMetricIter
	flowsHolder  MetricsFlowMetricIter
}

func NewMetrics() Metrics {
	obj := metrics{obj: &openapi.Metrics{}}
	obj.setDefault()
	return &obj
}

func (obj *metrics) msg() *openapi.Metrics {
	return obj.obj
}

func (obj *metrics) setMsg(msg *openapi.Metrics) Metrics {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetrics struct {
	obj *metrics
}

type marshalMetrics interface {
	// ToProto marshals Metrics to protobuf object *openapi.Metrics
	ToProto() (*openapi.Metrics, error)
	// ToPbText marshals Metrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Metrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals Metrics to JSON text
	ToJson() (string, error)
}

type unMarshalmetrics struct {
	obj *metrics
}

type unMarshalMetrics interface {
	// FromProto unmarshals Metrics from protobuf object *openapi.Metrics
	FromProto(msg *openapi.Metrics) (Metrics, error)
	// FromPbText unmarshals Metrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Metrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Metrics from JSON text
	FromJson(value string) error
}

func (obj *metrics) Marshal() marshalMetrics {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetrics{obj: obj}
	}
	return obj.marshaller
}

func (obj *metrics) Unmarshal() unMarshalMetrics {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetrics{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetrics) ToProto() (*openapi.Metrics, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetrics) FromProto(msg *openapi.Metrics) (Metrics, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetrics) ToPbText() (string, error) {
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

func (m *unMarshalmetrics) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalmetrics) ToYaml() (string, error) {
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

func (m *unMarshalmetrics) FromYaml(value string) error {
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
	m.obj.setNil()
	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalmetrics) ToJson() (string, error) {
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

func (m *unMarshalmetrics) FromJson(value string) error {
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
	m.obj.setNil()
	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *metrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metrics) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metrics) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metrics) Clone() (Metrics, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetrics()
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

func (obj *metrics) setNil() {
	obj.portsHolder = nil
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Metrics is description is TBD
type Metrics interface {
	Validation
	// msg marshals Metrics to protobuf object *openapi.Metrics
	// and doesn't set defaults
	msg() *openapi.Metrics
	// setMsg unmarshals Metrics from protobuf object *openapi.Metrics
	// and doesn't set defaults
	setMsg(*openapi.Metrics) Metrics
	// provides marshal interface
	Marshal() marshalMetrics
	// provides unmarshal interface
	Unmarshal() unMarshalMetrics
	// validate validates Metrics
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Metrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsChoiceEnum, set in Metrics
	Choice() MetricsChoiceEnum
	// setChoice assigns MetricsChoiceEnum provided by user to Metrics
	setChoice(value MetricsChoiceEnum) Metrics
	// HasChoice checks if Choice has been set in Metrics
	HasChoice() bool
	// Ports returns MetricsPortMetricIterIter, set in Metrics
	Ports() MetricsPortMetricIter
	// Flows returns MetricsFlowMetricIterIter, set in Metrics
	Flows() MetricsFlowMetricIter
	setNil()
}

type MetricsChoiceEnum string

// Enum of Choice on Metrics
var MetricsChoice = struct {
	PORTS MetricsChoiceEnum
	FLOWS MetricsChoiceEnum
}{
	PORTS: MetricsChoiceEnum("ports"),
	FLOWS: MetricsChoiceEnum("flows"),
}

func (obj *metrics) Choice() MetricsChoiceEnum {
	return MetricsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metrics) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metrics) setChoice(value MetricsChoiceEnum) Metrics {
	intValue, ok := openapi.Metrics_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.Metrics_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	obj.flowsHolder = nil
	obj.obj.Ports = nil
	obj.portsHolder = nil

	if value == MetricsChoice.PORTS {
		obj.obj.Ports = []*openapi.PortMetric{}
	}

	if value == MetricsChoice.FLOWS {
		obj.obj.Flows = []*openapi.FlowMetric{}
	}

	return obj
}

// description is TBD
// Ports returns a []PortMetric
func (obj *metrics) Ports() MetricsPortMetricIter {
	if len(obj.obj.Ports) == 0 {
		obj.setChoice(MetricsChoice.PORTS)
	}
	if obj.portsHolder == nil {
		obj.portsHolder = newMetricsPortMetricIter(&obj.obj.Ports).setMsg(obj)
	}
	return obj.portsHolder
}

type metricsPortMetricIter struct {
	obj             *metrics
	portMetricSlice []PortMetric
	fieldPtr        *[]*openapi.PortMetric
}

func newMetricsPortMetricIter(ptr *[]*openapi.PortMetric) MetricsPortMetricIter {
	return &metricsPortMetricIter{fieldPtr: ptr}
}

type MetricsPortMetricIter interface {
	setMsg(*metrics) MetricsPortMetricIter
	Items() []PortMetric
	Add() PortMetric
	Append(items ...PortMetric) MetricsPortMetricIter
	Set(index int, newObj PortMetric) MetricsPortMetricIter
	Clear() MetricsPortMetricIter
	clearHolderSlice() MetricsPortMetricIter
	appendHolderSlice(item PortMetric) MetricsPortMetricIter
}

func (obj *metricsPortMetricIter) setMsg(msg *metrics) MetricsPortMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&portMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsPortMetricIter) Items() []PortMetric {
	return obj.portMetricSlice
}

func (obj *metricsPortMetricIter) Add() PortMetric {
	newObj := &openapi.PortMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &portMetric{obj: newObj}
	newLibObj.setDefault()
	obj.portMetricSlice = append(obj.portMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsPortMetricIter) Append(items ...PortMetric) MetricsPortMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.portMetricSlice = append(obj.portMetricSlice, item)
	}
	return obj
}

func (obj *metricsPortMetricIter) Set(index int, newObj PortMetric) MetricsPortMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.portMetricSlice[index] = newObj
	return obj
}
func (obj *metricsPortMetricIter) Clear() MetricsPortMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.PortMetric{}
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsPortMetricIter) clearHolderSlice() MetricsPortMetricIter {
	if len(obj.portMetricSlice) > 0 {
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsPortMetricIter) appendHolderSlice(item PortMetric) MetricsPortMetricIter {
	obj.portMetricSlice = append(obj.portMetricSlice, item)
	return obj
}

// description is TBD
// Flows returns a []FlowMetric
func (obj *metrics) Flows() MetricsFlowMetricIter {
	if len(obj.obj.Flows) == 0 {
		obj.setChoice(MetricsChoice.FLOWS)
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newMetricsFlowMetricIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type metricsFlowMetricIter struct {
	obj             *metrics
	flowMetricSlice []FlowMetric
	fieldPtr        *[]*openapi.FlowMetric
}

func newMetricsFlowMetricIter(ptr *[]*openapi.FlowMetric) MetricsFlowMetricIter {
	return &metricsFlowMetricIter{fieldPtr: ptr}
}

type MetricsFlowMetricIter interface {
	setMsg(*metrics) MetricsFlowMetricIter
	Items() []FlowMetric
	Add() FlowMetric
	Append(items ...FlowMetric) MetricsFlowMetricIter
	Set(index int, newObj FlowMetric) MetricsFlowMetricIter
	Clear() MetricsFlowMetricIter
	clearHolderSlice() MetricsFlowMetricIter
	appendHolderSlice(item FlowMetric) MetricsFlowMetricIter
}

func (obj *metricsFlowMetricIter) setMsg(msg *metrics) MetricsFlowMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsFlowMetricIter) Items() []FlowMetric {
	return obj.flowMetricSlice
}

func (obj *metricsFlowMetricIter) Add() FlowMetric {
	newObj := &openapi.FlowMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowMetric{obj: newObj}
	newLibObj.setDefault()
	obj.flowMetricSlice = append(obj.flowMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsFlowMetricIter) Append(items ...FlowMetric) MetricsFlowMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	}
	return obj
}

func (obj *metricsFlowMetricIter) Set(index int, newObj FlowMetric) MetricsFlowMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowMetricSlice[index] = newObj
	return obj
}
func (obj *metricsFlowMetricIter) Clear() MetricsFlowMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.FlowMetric{}
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsFlowMetricIter) clearHolderSlice() MetricsFlowMetricIter {
	if len(obj.flowMetricSlice) > 0 {
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsFlowMetricIter) appendHolderSlice(item FlowMetric) MetricsFlowMetricIter {
	obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	return obj
}

func (obj *metrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if set_default {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				obj.Ports().appendHolderSlice(&portMetric{obj: item})
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&flowMetric{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *metrics) setDefault() {
	var choices_set int = 0
	var choice MetricsChoiceEnum

	if len(obj.obj.Ports) > 0 {
		choices_set += 1
		choice = MetricsChoice.PORTS
	}

	if len(obj.obj.Flows) > 0 {
		choices_set += 1
		choice = MetricsChoice.FLOWS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MetricsChoice.PORTS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Metrics")
			}
		} else {
			intVal := openapi.Metrics_Choice_Enum_value[string(choice)]
			enumValue := openapi.Metrics_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
