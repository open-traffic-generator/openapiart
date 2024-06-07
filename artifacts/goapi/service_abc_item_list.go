package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ServiceAbcItemList *****
type serviceAbcItemList struct {
	validation
	obj          *openapi.ServiceAbcItemList
	marshaller   marshalServiceAbcItemList
	unMarshaller unMarshalServiceAbcItemList
	itemsHolder  ServiceAbcItemListServiceAbcItemIter
}

func NewServiceAbcItemList() ServiceAbcItemList {
	obj := serviceAbcItemList{obj: &openapi.ServiceAbcItemList{}}
	obj.setDefault()
	return &obj
}

func (obj *serviceAbcItemList) msg() *openapi.ServiceAbcItemList {
	return obj.obj
}

func (obj *serviceAbcItemList) setMsg(msg *openapi.ServiceAbcItemList) ServiceAbcItemList {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalserviceAbcItemList struct {
	obj *serviceAbcItemList
}

type marshalServiceAbcItemList interface {
	// ToProto marshals ServiceAbcItemList to protobuf object *openapi.ServiceAbcItemList
	ToProto() (*openapi.ServiceAbcItemList, error)
	// ToPbText marshals ServiceAbcItemList to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ServiceAbcItemList to YAML text
	ToYaml() (string, error)
	// ToJson marshals ServiceAbcItemList to JSON text
	ToJson() (string, error)
}

type unMarshalserviceAbcItemList struct {
	obj *serviceAbcItemList
}

type unMarshalServiceAbcItemList interface {
	// FromProto unmarshals ServiceAbcItemList from protobuf object *openapi.ServiceAbcItemList
	FromProto(msg *openapi.ServiceAbcItemList) (ServiceAbcItemList, error)
	// FromPbText unmarshals ServiceAbcItemList from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ServiceAbcItemList from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ServiceAbcItemList from JSON text
	FromJson(value string) error
}

func (obj *serviceAbcItemList) Marshal() marshalServiceAbcItemList {
	if obj.marshaller == nil {
		obj.marshaller = &marshalserviceAbcItemList{obj: obj}
	}
	return obj.marshaller
}

func (obj *serviceAbcItemList) Unmarshal() unMarshalServiceAbcItemList {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalserviceAbcItemList{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalserviceAbcItemList) ToProto() (*openapi.ServiceAbcItemList, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalserviceAbcItemList) FromProto(msg *openapi.ServiceAbcItemList) (ServiceAbcItemList, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalserviceAbcItemList) ToPbText() (string, error) {
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

func (m *unMarshalserviceAbcItemList) FromPbText(value string) error {
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

func (m *marshalserviceAbcItemList) ToYaml() (string, error) {
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

func (m *unMarshalserviceAbcItemList) FromYaml(value string) error {
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

func (m *marshalserviceAbcItemList) ToJson() (string, error) {
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

func (m *unMarshalserviceAbcItemList) FromJson(value string) error {
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

func (obj *serviceAbcItemList) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *serviceAbcItemList) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *serviceAbcItemList) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *serviceAbcItemList) Clone() (ServiceAbcItemList, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewServiceAbcItemList()
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

func (obj *serviceAbcItemList) setNil() {
	obj.itemsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ServiceAbcItemList is description is TBD
type ServiceAbcItemList interface {
	Validation
	// msg marshals ServiceAbcItemList to protobuf object *openapi.ServiceAbcItemList
	// and doesn't set defaults
	msg() *openapi.ServiceAbcItemList
	// setMsg unmarshals ServiceAbcItemList from protobuf object *openapi.ServiceAbcItemList
	// and doesn't set defaults
	setMsg(*openapi.ServiceAbcItemList) ServiceAbcItemList
	// provides marshal interface
	Marshal() marshalServiceAbcItemList
	// provides unmarshal interface
	Unmarshal() unMarshalServiceAbcItemList
	// validate validates ServiceAbcItemList
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ServiceAbcItemList, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Items returns ServiceAbcItemListServiceAbcItemIterIter, set in ServiceAbcItemList
	Items() ServiceAbcItemListServiceAbcItemIter
	setNil()
}

// description is TBD
// Items returns a []ServiceAbcItem
func (obj *serviceAbcItemList) Items() ServiceAbcItemListServiceAbcItemIter {
	if len(obj.obj.Items) == 0 {
		obj.obj.Items = []*openapi.ServiceAbcItem{}
	}
	if obj.itemsHolder == nil {
		obj.itemsHolder = newServiceAbcItemListServiceAbcItemIter(&obj.obj.Items).setMsg(obj)
	}
	return obj.itemsHolder
}

type serviceAbcItemListServiceAbcItemIter struct {
	obj                 *serviceAbcItemList
	serviceAbcItemSlice []ServiceAbcItem
	fieldPtr            *[]*openapi.ServiceAbcItem
}

func newServiceAbcItemListServiceAbcItemIter(ptr *[]*openapi.ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	return &serviceAbcItemListServiceAbcItemIter{fieldPtr: ptr}
}

type ServiceAbcItemListServiceAbcItemIter interface {
	setMsg(*serviceAbcItemList) ServiceAbcItemListServiceAbcItemIter
	Items() []ServiceAbcItem
	Add() ServiceAbcItem
	Append(items ...ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
	Set(index int, newObj ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
	Clear() ServiceAbcItemListServiceAbcItemIter
	clearHolderSlice() ServiceAbcItemListServiceAbcItemIter
	appendHolderSlice(item ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
}

func (obj *serviceAbcItemListServiceAbcItemIter) setMsg(msg *serviceAbcItemList) ServiceAbcItemListServiceAbcItemIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&serviceAbcItem{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Items() []ServiceAbcItem {
	return obj.serviceAbcItemSlice
}

func (obj *serviceAbcItemListServiceAbcItemIter) Add() ServiceAbcItem {
	newObj := &openapi.ServiceAbcItem{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &serviceAbcItem{obj: newObj}
	newLibObj.setDefault()
	obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, newLibObj)
	return newLibObj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Append(items ...ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, item)
	}
	return obj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Set(index int, newObj ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.serviceAbcItemSlice[index] = newObj
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) Clear() ServiceAbcItemListServiceAbcItemIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.ServiceAbcItem{}
		obj.serviceAbcItemSlice = []ServiceAbcItem{}
	}
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) clearHolderSlice() ServiceAbcItemListServiceAbcItemIter {
	if len(obj.serviceAbcItemSlice) > 0 {
		obj.serviceAbcItemSlice = []ServiceAbcItem{}
	}
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) appendHolderSlice(item ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, item)
	return obj
}

func (obj *serviceAbcItemList) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Items) != 0 {

		if set_default {
			obj.Items().clearHolderSlice()
			for _, item := range obj.obj.Items {
				obj.Items().appendHolderSlice(&serviceAbcItem{obj: item})
			}
		}
		for _, item := range obj.Items().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *serviceAbcItemList) setDefault() {

}
