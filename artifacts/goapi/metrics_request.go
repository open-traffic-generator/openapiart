package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** MetricsRequest *****
type metricsRequest struct {
	validation
	obj          *openapi.MetricsRequest
	marshaller   marshalMetricsRequest
	unMarshaller unMarshalMetricsRequest
}

func NewMetricsRequest() MetricsRequest {
	obj := metricsRequest{obj: &openapi.MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *metricsRequest) msg() *openapi.MetricsRequest {
	return obj.obj
}

func (obj *metricsRequest) setMsg(msg *openapi.MetricsRequest) MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalmetricsRequest struct {
	obj *metricsRequest
}

type marshalMetricsRequest interface {
	// ToProto marshals MetricsRequest to protobuf object *openapi.MetricsRequest
	ToProto() (*openapi.MetricsRequest, error)
	// ToPbText marshals MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricsRequest to JSON text
	ToJson() (string, error)
}

type unMarshalmetricsRequest struct {
	obj *metricsRequest
}

type unMarshalMetricsRequest interface {
	// FromProto unmarshals MetricsRequest from protobuf object *openapi.MetricsRequest
	FromProto(msg *openapi.MetricsRequest) (MetricsRequest, error)
	// FromPbText unmarshals MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricsRequest from JSON text
	FromJson(value string) error
}

func (obj *metricsRequest) Marshal() marshalMetricsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshalmetricsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *metricsRequest) Unmarshal() unMarshalMetricsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalmetricsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalmetricsRequest) ToProto() (*openapi.MetricsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalmetricsRequest) FromProto(msg *openapi.MetricsRequest) (MetricsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalmetricsRequest) ToPbText() (string, error) {
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

func (m *unMarshalmetricsRequest) FromPbText(value string) error {
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

func (m *marshalmetricsRequest) ToYaml() (string, error) {
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

func (m *unMarshalmetricsRequest) FromYaml(value string) error {
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

func (m *marshalmetricsRequest) ToJson() (string, error) {
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

func (m *unMarshalmetricsRequest) FromJson(value string) error {
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

func (obj *metricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricsRequest) Clone() (MetricsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricsRequest()
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

// MetricsRequest is description is TBD
type MetricsRequest interface {
	Validation
	// msg marshals MetricsRequest to protobuf object *openapi.MetricsRequest
	// and doesn't set defaults
	msg() *openapi.MetricsRequest
	// setMsg unmarshals MetricsRequest from protobuf object *openapi.MetricsRequest
	// and doesn't set defaults
	setMsg(*openapi.MetricsRequest) MetricsRequest
	// provides marshal interface
	Marshal() marshalMetricsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalMetricsRequest
	// validate validates MetricsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsRequestChoiceEnum, set in MetricsRequest
	Choice() MetricsRequestChoiceEnum
	// setChoice assigns MetricsRequestChoiceEnum provided by user to MetricsRequest
	setChoice(value MetricsRequestChoiceEnum) MetricsRequest
	// HasChoice checks if Choice has been set in MetricsRequest
	HasChoice() bool
	// Port returns string, set in MetricsRequest.
	Port() string
	// SetPort assigns string provided by user to MetricsRequest
	SetPort(value string) MetricsRequest
	// HasPort checks if Port has been set in MetricsRequest
	HasPort() bool
	// Flow returns string, set in MetricsRequest.
	Flow() string
	// SetFlow assigns string provided by user to MetricsRequest
	SetFlow(value string) MetricsRequest
	// HasFlow checks if Flow has been set in MetricsRequest
	HasFlow() bool
}

type MetricsRequestChoiceEnum string

// Enum of Choice on MetricsRequest
var MetricsRequestChoice = struct {
	PORT MetricsRequestChoiceEnum
	FLOW MetricsRequestChoiceEnum
}{
	PORT: MetricsRequestChoiceEnum("port"),
	FLOW: MetricsRequestChoiceEnum("flow"),
}

func (obj *metricsRequest) Choice() MetricsRequestChoiceEnum {
	return MetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metricsRequest) setChoice(value MetricsRequestChoiceEnum) MetricsRequest {
	intValue, ok := openapi.MetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.MetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flow = nil
	obj.obj.Port = nil
	return obj
}

// description is TBD
// Port returns a string
func (obj *metricsRequest) Port() string {

	if obj.obj.Port == nil {
		obj.setChoice(MetricsRequestChoice.PORT)
	}

	return *obj.obj.Port

}

// description is TBD
// Port returns a string
func (obj *metricsRequest) HasPort() bool {
	return obj.obj.Port != nil
}

// description is TBD
// SetPort sets the string value in the MetricsRequest object
func (obj *metricsRequest) SetPort(value string) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.PORT)
	obj.obj.Port = &value
	return obj
}

// description is TBD
// Flow returns a string
func (obj *metricsRequest) Flow() string {

	if obj.obj.Flow == nil {
		obj.setChoice(MetricsRequestChoice.FLOW)
	}

	return *obj.obj.Flow

}

// description is TBD
// Flow returns a string
func (obj *metricsRequest) HasFlow() bool {
	return obj.obj.Flow != nil
}

// description is TBD
// SetFlow sets the string value in the MetricsRequest object
func (obj *metricsRequest) SetFlow(value string) MetricsRequest {
	obj.setChoice(MetricsRequestChoice.FLOW)
	obj.obj.Flow = &value
	return obj
}

func (obj *metricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *metricsRequest) setDefault() {
	var choices_set int = 0
	var choice MetricsRequestChoiceEnum

	if obj.obj.Port != nil {
		choices_set += 1
		choice = MetricsRequestChoice.PORT
	}

	if obj.obj.Flow != nil {
		choices_set += 1
		choice = MetricsRequestChoice.FLOW
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(MetricsRequestChoice.PORT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in MetricsRequest")
			}
		} else {
			intVal := openapi.MetricsRequest_Choice_Enum_value[string(choice)]
			enumValue := openapi.MetricsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
