package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** WarningDetails *****
type warningDetails struct {
	validation
	obj          *openapi.WarningDetails
	marshaller   marshalWarningDetails
	unMarshaller unMarshalWarningDetails
}

func NewWarningDetails() WarningDetails {
	obj := warningDetails{obj: &openapi.WarningDetails{}}
	obj.setDefault()
	return &obj
}

func (obj *warningDetails) msg() *openapi.WarningDetails {
	return obj.obj
}

func (obj *warningDetails) setMsg(msg *openapi.WarningDetails) WarningDetails {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalwarningDetails struct {
	obj *warningDetails
}

type marshalWarningDetails interface {
	// ToProto marshals WarningDetails to protobuf object *openapi.WarningDetails
	ToProto() (*openapi.WarningDetails, error)
	// ToPbText marshals WarningDetails to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals WarningDetails to YAML text
	ToYaml() (string, error)
	// ToJson marshals WarningDetails to JSON text
	ToJson() (string, error)
}

type unMarshalwarningDetails struct {
	obj *warningDetails
}

type unMarshalWarningDetails interface {
	// FromProto unmarshals WarningDetails from protobuf object *openapi.WarningDetails
	FromProto(msg *openapi.WarningDetails) (WarningDetails, error)
	// FromPbText unmarshals WarningDetails from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals WarningDetails from YAML text
	FromYaml(value string) error
	// FromJson unmarshals WarningDetails from JSON text
	FromJson(value string) error
}

func (obj *warningDetails) Marshal() marshalWarningDetails {
	if obj.marshaller == nil {
		obj.marshaller = &marshalwarningDetails{obj: obj}
	}
	return obj.marshaller
}

func (obj *warningDetails) Unmarshal() unMarshalWarningDetails {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalwarningDetails{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalwarningDetails) ToProto() (*openapi.WarningDetails, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalwarningDetails) FromProto(msg *openapi.WarningDetails) (WarningDetails, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalwarningDetails) ToPbText() (string, error) {
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

func (m *unMarshalwarningDetails) FromPbText(value string) error {
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

func (m *marshalwarningDetails) ToYaml() (string, error) {
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

func (m *unMarshalwarningDetails) FromYaml(value string) error {
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

func (m *marshalwarningDetails) ToJson() (string, error) {
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

func (m *unMarshalwarningDetails) FromJson(value string) error {
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

func (obj *warningDetails) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *warningDetails) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *warningDetails) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *warningDetails) Clone() (WarningDetails, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewWarningDetails()
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

// WarningDetails is description is TBD
type WarningDetails interface {
	Validation
	// msg marshals WarningDetails to protobuf object *openapi.WarningDetails
	// and doesn't set defaults
	msg() *openapi.WarningDetails
	// setMsg unmarshals WarningDetails from protobuf object *openapi.WarningDetails
	// and doesn't set defaults
	setMsg(*openapi.WarningDetails) WarningDetails
	// provides marshal interface
	Marshal() marshalWarningDetails
	// provides unmarshal interface
	Unmarshal() unMarshalWarningDetails
	// validate validates WarningDetails
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (WarningDetails, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warnings returns []string, set in WarningDetails.
	Warnings() []string
	// SetWarnings assigns []string provided by user to WarningDetails
	SetWarnings(value []string) WarningDetails
}

// description is TBD
// Warnings returns a []string
func (obj *warningDetails) Warnings() []string {
	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	return obj.obj.Warnings
}

// description is TBD
// SetWarnings sets the []string value in the WarningDetails object
func (obj *warningDetails) SetWarnings(value []string) WarningDetails {

	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	obj.obj.Warnings = value

	return obj
}

func (obj *warningDetails) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *warningDetails) setDefault() {

}
