package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ApiTestInputBody *****
type apiTestInputBody struct {
	validation
	obj          *openapi.ApiTestInputBody
	marshaller   marshalApiTestInputBody
	unMarshaller unMarshalApiTestInputBody
}

func NewApiTestInputBody() ApiTestInputBody {
	obj := apiTestInputBody{obj: &openapi.ApiTestInputBody{}}
	obj.setDefault()
	return &obj
}

func (obj *apiTestInputBody) msg() *openapi.ApiTestInputBody {
	return obj.obj
}

func (obj *apiTestInputBody) setMsg(msg *openapi.ApiTestInputBody) ApiTestInputBody {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalapiTestInputBody struct {
	obj *apiTestInputBody
}

type marshalApiTestInputBody interface {
	// ToProto marshals ApiTestInputBody to protobuf object *openapi.ApiTestInputBody
	ToProto() (*openapi.ApiTestInputBody, error)
	// ToPbText marshals ApiTestInputBody to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ApiTestInputBody to YAML text
	ToYaml() (string, error)
	// ToJson marshals ApiTestInputBody to JSON text
	ToJson() (string, error)
}

type unMarshalapiTestInputBody struct {
	obj *apiTestInputBody
}

type unMarshalApiTestInputBody interface {
	// FromProto unmarshals ApiTestInputBody from protobuf object *openapi.ApiTestInputBody
	FromProto(msg *openapi.ApiTestInputBody) (ApiTestInputBody, error)
	// FromPbText unmarshals ApiTestInputBody from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ApiTestInputBody from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ApiTestInputBody from JSON text
	FromJson(value string) error
}

func (obj *apiTestInputBody) Marshal() marshalApiTestInputBody {
	if obj.marshaller == nil {
		obj.marshaller = &marshalapiTestInputBody{obj: obj}
	}
	return obj.marshaller
}

func (obj *apiTestInputBody) Unmarshal() unMarshalApiTestInputBody {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalapiTestInputBody{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalapiTestInputBody) ToProto() (*openapi.ApiTestInputBody, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalapiTestInputBody) FromProto(msg *openapi.ApiTestInputBody) (ApiTestInputBody, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalapiTestInputBody) ToPbText() (string, error) {
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

func (m *unMarshalapiTestInputBody) FromPbText(value string) error {
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

func (m *marshalapiTestInputBody) ToYaml() (string, error) {
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

func (m *unMarshalapiTestInputBody) FromYaml(value string) error {
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

func (m *marshalapiTestInputBody) ToJson() (string, error) {
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

func (m *unMarshalapiTestInputBody) FromJson(value string) error {
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

func (obj *apiTestInputBody) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *apiTestInputBody) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *apiTestInputBody) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *apiTestInputBody) Clone() (ApiTestInputBody, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewApiTestInputBody()
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

// ApiTestInputBody is description is TBD
type ApiTestInputBody interface {
	Validation
	// msg marshals ApiTestInputBody to protobuf object *openapi.ApiTestInputBody
	// and doesn't set defaults
	msg() *openapi.ApiTestInputBody
	// setMsg unmarshals ApiTestInputBody from protobuf object *openapi.ApiTestInputBody
	// and doesn't set defaults
	setMsg(*openapi.ApiTestInputBody) ApiTestInputBody
	// provides marshal interface
	Marshal() marshalApiTestInputBody
	// provides unmarshal interface
	Unmarshal() unMarshalApiTestInputBody
	// validate validates ApiTestInputBody
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ApiTestInputBody, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SomeString returns string, set in ApiTestInputBody.
	SomeString() string
	// SetSomeString assigns string provided by user to ApiTestInputBody
	SetSomeString(value string) ApiTestInputBody
	// HasSomeString checks if SomeString has been set in ApiTestInputBody
	HasSomeString() bool
}

// description is TBD
// SomeString returns a string
func (obj *apiTestInputBody) SomeString() string {

	return *obj.obj.SomeString

}

// description is TBD
// SomeString returns a string
func (obj *apiTestInputBody) HasSomeString() bool {
	return obj.obj.SomeString != nil
}

// description is TBD
// SetSomeString sets the string value in the ApiTestInputBody object
func (obj *apiTestInputBody) SetSomeString(value string) ApiTestInputBody {

	obj.obj.SomeString = &value
	return obj
}

func (obj *apiTestInputBody) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *apiTestInputBody) setDefault() {

}
