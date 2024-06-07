package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GetMetricsResponse *****
type getMetricsResponse struct {
	validation
	obj           *openapi.GetMetricsResponse
	marshaller    marshalGetMetricsResponse
	unMarshaller  unMarshalGetMetricsResponse
	metricsHolder Metrics
}

func NewGetMetricsResponse() GetMetricsResponse {
	obj := getMetricsResponse{obj: &openapi.GetMetricsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getMetricsResponse) msg() *openapi.GetMetricsResponse {
	return obj.obj
}

func (obj *getMetricsResponse) setMsg(msg *openapi.GetMetricsResponse) GetMetricsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgetMetricsResponse struct {
	obj *getMetricsResponse
}

type marshalGetMetricsResponse interface {
	// ToProto marshals GetMetricsResponse to protobuf object *openapi.GetMetricsResponse
	ToProto() (*openapi.GetMetricsResponse, error)
	// ToPbText marshals GetMetricsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetMetricsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetMetricsResponse to JSON text
	ToJson() (string, error)
}

type unMarshalgetMetricsResponse struct {
	obj *getMetricsResponse
}

type unMarshalGetMetricsResponse interface {
	// FromProto unmarshals GetMetricsResponse from protobuf object *openapi.GetMetricsResponse
	FromProto(msg *openapi.GetMetricsResponse) (GetMetricsResponse, error)
	// FromPbText unmarshals GetMetricsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetMetricsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetMetricsResponse from JSON text
	FromJson(value string) error
}

func (obj *getMetricsResponse) Marshal() marshalGetMetricsResponse {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgetMetricsResponse{obj: obj}
	}
	return obj.marshaller
}

func (obj *getMetricsResponse) Unmarshal() unMarshalGetMetricsResponse {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgetMetricsResponse{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgetMetricsResponse) ToProto() (*openapi.GetMetricsResponse, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgetMetricsResponse) FromProto(msg *openapi.GetMetricsResponse) (GetMetricsResponse, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgetMetricsResponse) ToPbText() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromPbText(value string) error {
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

func (m *marshalgetMetricsResponse) ToYaml() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromYaml(value string) error {
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

func (m *marshalgetMetricsResponse) ToJson() (string, error) {
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

func (m *unMarshalgetMetricsResponse) FromJson(value string) error {
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

func (obj *getMetricsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getMetricsResponse) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getMetricsResponse) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getMetricsResponse) Clone() (GetMetricsResponse, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetMetricsResponse()
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

func (obj *getMetricsResponse) setNil() {
	obj.metricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetMetricsResponse is description is TBD
type GetMetricsResponse interface {
	Validation
	// msg marshals GetMetricsResponse to protobuf object *openapi.GetMetricsResponse
	// and doesn't set defaults
	msg() *openapi.GetMetricsResponse
	// setMsg unmarshals GetMetricsResponse from protobuf object *openapi.GetMetricsResponse
	// and doesn't set defaults
	setMsg(*openapi.GetMetricsResponse) GetMetricsResponse
	// provides marshal interface
	Marshal() marshalGetMetricsResponse
	// provides unmarshal interface
	Unmarshal() unMarshalGetMetricsResponse
	// validate validates GetMetricsResponse
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetMetricsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Metrics returns Metrics, set in GetMetricsResponse.
	// Metrics is description is TBD
	Metrics() Metrics
	// SetMetrics assigns Metrics provided by user to GetMetricsResponse.
	// Metrics is description is TBD
	SetMetrics(value Metrics) GetMetricsResponse
	// HasMetrics checks if Metrics has been set in GetMetricsResponse
	HasMetrics() bool
	setNil()
}

// description is TBD
// Metrics returns a Metrics
func (obj *getMetricsResponse) Metrics() Metrics {
	if obj.obj.Metrics == nil {
		obj.obj.Metrics = NewMetrics().msg()
	}
	if obj.metricsHolder == nil {
		obj.metricsHolder = &metrics{obj: obj.obj.Metrics}
	}
	return obj.metricsHolder
}

// description is TBD
// Metrics returns a Metrics
func (obj *getMetricsResponse) HasMetrics() bool {
	return obj.obj.Metrics != nil
}

// description is TBD
// SetMetrics sets the Metrics value in the GetMetricsResponse object
func (obj *getMetricsResponse) SetMetrics(value Metrics) GetMetricsResponse {

	obj.metricsHolder = nil
	obj.obj.Metrics = value.msg()

	return obj
}

func (obj *getMetricsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metrics != nil {

		obj.Metrics().validateObj(vObj, set_default)
	}

}

func (obj *getMetricsResponse) setDefault() {

}
