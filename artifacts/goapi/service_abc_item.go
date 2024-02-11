package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ServiceAbcItem *****
type serviceAbcItem struct {
	validation
	obj          *openapi.ServiceAbcItem
	marshaller   marshalServiceAbcItem
	unMarshaller unMarshalServiceAbcItem
}

func NewServiceAbcItem() ServiceAbcItem {
	obj := serviceAbcItem{obj: &openapi.ServiceAbcItem{}}
	obj.setDefault()
	return &obj
}

func (obj *serviceAbcItem) msg() *openapi.ServiceAbcItem {
	return obj.obj
}

func (obj *serviceAbcItem) setMsg(msg *openapi.ServiceAbcItem) ServiceAbcItem {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalserviceAbcItem struct {
	obj *serviceAbcItem
}

type marshalServiceAbcItem interface {
	// ToProto marshals ServiceAbcItem to protobuf object *openapi.ServiceAbcItem
	ToProto() (*openapi.ServiceAbcItem, error)
	// ToPbText marshals ServiceAbcItem to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ServiceAbcItem to YAML text
	ToYaml() (string, error)
	// ToJson marshals ServiceAbcItem to JSON text
	ToJson() (string, error)
}

type unMarshalserviceAbcItem struct {
	obj *serviceAbcItem
}

type unMarshalServiceAbcItem interface {
	// FromProto unmarshals ServiceAbcItem from protobuf object *openapi.ServiceAbcItem
	FromProto(msg *openapi.ServiceAbcItem) (ServiceAbcItem, error)
	// FromPbText unmarshals ServiceAbcItem from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ServiceAbcItem from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ServiceAbcItem from JSON text
	FromJson(value string) error
}

func (obj *serviceAbcItem) Marshal() marshalServiceAbcItem {
	if obj.marshaller == nil {
		obj.marshaller = &marshalserviceAbcItem{obj: obj}
	}
	return obj.marshaller
}

func (obj *serviceAbcItem) Unmarshal() unMarshalServiceAbcItem {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalserviceAbcItem{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalserviceAbcItem) ToProto() (*openapi.ServiceAbcItem, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalserviceAbcItem) FromProto(msg *openapi.ServiceAbcItem) (ServiceAbcItem, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalserviceAbcItem) ToPbText() (string, error) {
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

func (m *unMarshalserviceAbcItem) FromPbText(value string) error {
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

func (m *marshalserviceAbcItem) ToYaml() (string, error) {
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

func (m *unMarshalserviceAbcItem) FromYaml(value string) error {
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

func (m *marshalserviceAbcItem) ToJson() (string, error) {
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

func (m *unMarshalserviceAbcItem) FromJson(value string) error {
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

func (obj *serviceAbcItem) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *serviceAbcItem) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *serviceAbcItem) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *serviceAbcItem) Clone() (ServiceAbcItem, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewServiceAbcItem()
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

// ServiceAbcItem is description is TBD
type ServiceAbcItem interface {
	Validation
	// msg marshals ServiceAbcItem to protobuf object *openapi.ServiceAbcItem
	// and doesn't set defaults
	msg() *openapi.ServiceAbcItem
	// setMsg unmarshals ServiceAbcItem from protobuf object *openapi.ServiceAbcItem
	// and doesn't set defaults
	setMsg(*openapi.ServiceAbcItem) ServiceAbcItem
	// provides marshal interface
	Marshal() marshalServiceAbcItem
	// provides unmarshal interface
	Unmarshal() unMarshalServiceAbcItem
	// validate validates ServiceAbcItem
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ServiceAbcItem, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SomeId returns string, set in ServiceAbcItem.
	SomeId() string
	// SetSomeId assigns string provided by user to ServiceAbcItem
	SetSomeId(value string) ServiceAbcItem
	// HasSomeId checks if SomeId has been set in ServiceAbcItem
	HasSomeId() bool
	// SomeString returns string, set in ServiceAbcItem.
	SomeString() string
	// SetSomeString assigns string provided by user to ServiceAbcItem
	SetSomeString(value string) ServiceAbcItem
	// HasSomeString checks if SomeString has been set in ServiceAbcItem
	HasSomeString() bool
	// PathId returns string, set in ServiceAbcItem.
	PathId() string
	// SetPathId assigns string provided by user to ServiceAbcItem
	SetPathId(value string) ServiceAbcItem
	// HasPathId checks if PathId has been set in ServiceAbcItem
	HasPathId() bool
	// Level2 returns string, set in ServiceAbcItem.
	Level2() string
	// SetLevel2 assigns string provided by user to ServiceAbcItem
	SetLevel2(value string) ServiceAbcItem
	// HasLevel2 checks if Level2 has been set in ServiceAbcItem
	HasLevel2() bool
}

// description is TBD
// SomeId returns a string
func (obj *serviceAbcItem) SomeId() string {

	return *obj.obj.SomeId

}

// description is TBD
// SomeId returns a string
func (obj *serviceAbcItem) HasSomeId() bool {
	return obj.obj.SomeId != nil
}

// description is TBD
// SetSomeId sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetSomeId(value string) ServiceAbcItem {

	obj.obj.SomeId = &value
	return obj
}

// description is TBD
// SomeString returns a string
func (obj *serviceAbcItem) SomeString() string {

	return *obj.obj.SomeString

}

// description is TBD
// SomeString returns a string
func (obj *serviceAbcItem) HasSomeString() bool {
	return obj.obj.SomeString != nil
}

// description is TBD
// SetSomeString sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetSomeString(value string) ServiceAbcItem {

	obj.obj.SomeString = &value
	return obj
}

// description is TBD
// PathId returns a string
func (obj *serviceAbcItem) PathId() string {

	return *obj.obj.PathId

}

// description is TBD
// PathId returns a string
func (obj *serviceAbcItem) HasPathId() bool {
	return obj.obj.PathId != nil
}

// description is TBD
// SetPathId sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetPathId(value string) ServiceAbcItem {

	obj.obj.PathId = &value
	return obj
}

// description is TBD
// Level2 returns a string
func (obj *serviceAbcItem) Level2() string {

	return *obj.obj.Level_2

}

// description is TBD
// Level2 returns a string
func (obj *serviceAbcItem) HasLevel2() bool {
	return obj.obj.Level_2 != nil
}

// description is TBD
// SetLevel2 sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetLevel2(value string) ServiceAbcItem {

	obj.obj.Level_2 = &value
	return obj
}

func (obj *serviceAbcItem) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *serviceAbcItem) setDefault() {

}
