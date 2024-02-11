package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateConfig *****
type updateConfig struct {
	validation
	obj          *openapi.UpdateConfig
	marshaller   marshalUpdateConfig
	unMarshaller unMarshalUpdateConfig
	gHolder      UpdateConfigGObjectIter
}

func NewUpdateConfig() UpdateConfig {
	obj := updateConfig{obj: &openapi.UpdateConfig{}}
	obj.setDefault()
	return &obj
}

func (obj *updateConfig) msg() *openapi.UpdateConfig {
	return obj.obj
}

func (obj *updateConfig) setMsg(msg *openapi.UpdateConfig) UpdateConfig {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateConfig struct {
	obj *updateConfig
}

type marshalUpdateConfig interface {
	// ToProto marshals UpdateConfig to protobuf object *openapi.UpdateConfig
	ToProto() (*openapi.UpdateConfig, error)
	// ToPbText marshals UpdateConfig to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateConfig to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateConfig to JSON text
	ToJson() (string, error)
}

type unMarshalupdateConfig struct {
	obj *updateConfig
}

type unMarshalUpdateConfig interface {
	// FromProto unmarshals UpdateConfig from protobuf object *openapi.UpdateConfig
	FromProto(msg *openapi.UpdateConfig) (UpdateConfig, error)
	// FromPbText unmarshals UpdateConfig from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateConfig from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateConfig from JSON text
	FromJson(value string) error
}

func (obj *updateConfig) Marshal() marshalUpdateConfig {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateConfig{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateConfig) Unmarshal() unMarshalUpdateConfig {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateConfig{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateConfig) ToProto() (*openapi.UpdateConfig, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateConfig) FromProto(msg *openapi.UpdateConfig) (UpdateConfig, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateConfig) ToPbText() (string, error) {
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

func (m *unMarshalupdateConfig) FromPbText(value string) error {
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

func (m *marshalupdateConfig) ToYaml() (string, error) {
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

func (m *unMarshalupdateConfig) FromYaml(value string) error {
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

func (m *marshalupdateConfig) ToJson() (string, error) {
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

func (m *unMarshalupdateConfig) FromJson(value string) error {
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

func (obj *updateConfig) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateConfig) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateConfig) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateConfig) Clone() (UpdateConfig, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateConfig()
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

func (obj *updateConfig) setNil() {
	obj.gHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateConfig is under Review: the whole schema is being reviewed
//
// Object to Test required Parameter
type UpdateConfig interface {
	Validation
	// msg marshals UpdateConfig to protobuf object *openapi.UpdateConfig
	// and doesn't set defaults
	msg() *openapi.UpdateConfig
	// setMsg unmarshals UpdateConfig from protobuf object *openapi.UpdateConfig
	// and doesn't set defaults
	setMsg(*openapi.UpdateConfig) UpdateConfig
	// provides marshal interface
	Marshal() marshalUpdateConfig
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateConfig
	// validate validates UpdateConfig
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateConfig, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// G returns UpdateConfigGObjectIterIter, set in UpdateConfig
	G() UpdateConfigGObjectIter
	setNil()
}

// A list of objects with choice and properties
// G returns a []GObject
func (obj *updateConfig) G() UpdateConfigGObjectIter {
	if len(obj.obj.G) == 0 {
		obj.obj.G = []*openapi.GObject{}
	}
	if obj.gHolder == nil {
		obj.gHolder = newUpdateConfigGObjectIter(&obj.obj.G).setMsg(obj)
	}
	return obj.gHolder
}

type updateConfigGObjectIter struct {
	obj          *updateConfig
	gObjectSlice []GObject
	fieldPtr     *[]*openapi.GObject
}

func newUpdateConfigGObjectIter(ptr *[]*openapi.GObject) UpdateConfigGObjectIter {
	return &updateConfigGObjectIter{fieldPtr: ptr}
}

type UpdateConfigGObjectIter interface {
	setMsg(*updateConfig) UpdateConfigGObjectIter
	Items() []GObject
	Add() GObject
	Append(items ...GObject) UpdateConfigGObjectIter
	Set(index int, newObj GObject) UpdateConfigGObjectIter
	Clear() UpdateConfigGObjectIter
	clearHolderSlice() UpdateConfigGObjectIter
	appendHolderSlice(item GObject) UpdateConfigGObjectIter
}

func (obj *updateConfigGObjectIter) setMsg(msg *updateConfig) UpdateConfigGObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&gObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateConfigGObjectIter) Items() []GObject {
	return obj.gObjectSlice
}

func (obj *updateConfigGObjectIter) Add() GObject {
	newObj := &openapi.GObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &gObject{obj: newObj}
	newLibObj.setDefault()
	obj.gObjectSlice = append(obj.gObjectSlice, newLibObj)
	return newLibObj
}

func (obj *updateConfigGObjectIter) Append(items ...GObject) UpdateConfigGObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.gObjectSlice = append(obj.gObjectSlice, item)
	}
	return obj
}

func (obj *updateConfigGObjectIter) Set(index int, newObj GObject) UpdateConfigGObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.gObjectSlice[index] = newObj
	return obj
}
func (obj *updateConfigGObjectIter) Clear() UpdateConfigGObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.GObject{}
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *updateConfigGObjectIter) clearHolderSlice() UpdateConfigGObjectIter {
	if len(obj.gObjectSlice) > 0 {
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *updateConfigGObjectIter) appendHolderSlice(item GObject) UpdateConfigGObjectIter {
	obj.gObjectSlice = append(obj.gObjectSlice, item)
	return obj
}

func (obj *updateConfig) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("UpdateConfig is under review, the whole schema is being reviewed")

	if len(obj.obj.G) != 0 {

		if set_default {
			obj.G().clearHolderSlice()
			for _, item := range obj.obj.G {
				obj.G().appendHolderSlice(&gObject{obj: item})
			}
		}
		for _, item := range obj.G().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *updateConfig) setDefault() {

}
