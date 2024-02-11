package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** NativeFeatures *****
type nativeFeatures struct {
	validation
	obj                    *openapi.NativeFeatures
	marshaller             marshalNativeFeatures
	unMarshaller           unMarshalNativeFeatures
	requiredValHolder      RequiredVal
	optionalValHolder      OptionalVal
	boundaryValHolder      BoundaryVal
	requiredValArrayHolder RequiredValArray
	optionalValArrayHolder OptionalValArray
	boundaryValArrayHolder BoundaryValArray
	nestedRefObjectHolder  NestedRefObject
	mixedObjectHolder      MixedObject
	numberTypeObjectHolder NumberTypeObject
	iterObjectHolder       NativeFeaturesMixedObjectIter
}

func NewNativeFeatures() NativeFeatures {
	obj := nativeFeatures{obj: &openapi.NativeFeatures{}}
	obj.setDefault()
	return &obj
}

func (obj *nativeFeatures) msg() *openapi.NativeFeatures {
	return obj.obj
}

func (obj *nativeFeatures) setMsg(msg *openapi.NativeFeatures) NativeFeatures {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalnativeFeatures struct {
	obj *nativeFeatures
}

type marshalNativeFeatures interface {
	// ToProto marshals NativeFeatures to protobuf object *openapi.NativeFeatures
	ToProto() (*openapi.NativeFeatures, error)
	// ToPbText marshals NativeFeatures to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals NativeFeatures to YAML text
	ToYaml() (string, error)
	// ToJson marshals NativeFeatures to JSON text
	ToJson() (string, error)
}

type unMarshalnativeFeatures struct {
	obj *nativeFeatures
}

type unMarshalNativeFeatures interface {
	// FromProto unmarshals NativeFeatures from protobuf object *openapi.NativeFeatures
	FromProto(msg *openapi.NativeFeatures) (NativeFeatures, error)
	// FromPbText unmarshals NativeFeatures from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals NativeFeatures from YAML text
	FromYaml(value string) error
	// FromJson unmarshals NativeFeatures from JSON text
	FromJson(value string) error
}

func (obj *nativeFeatures) Marshal() marshalNativeFeatures {
	if obj.marshaller == nil {
		obj.marshaller = &marshalnativeFeatures{obj: obj}
	}
	return obj.marshaller
}

func (obj *nativeFeatures) Unmarshal() unMarshalNativeFeatures {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalnativeFeatures{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalnativeFeatures) ToProto() (*openapi.NativeFeatures, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalnativeFeatures) FromProto(msg *openapi.NativeFeatures) (NativeFeatures, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalnativeFeatures) ToPbText() (string, error) {
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

func (m *unMarshalnativeFeatures) FromPbText(value string) error {
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

func (m *marshalnativeFeatures) ToYaml() (string, error) {
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

func (m *unMarshalnativeFeatures) FromYaml(value string) error {
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

func (m *marshalnativeFeatures) ToJson() (string, error) {
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

func (m *unMarshalnativeFeatures) FromJson(value string) error {
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

func (obj *nativeFeatures) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *nativeFeatures) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *nativeFeatures) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *nativeFeatures) Clone() (NativeFeatures, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNativeFeatures()
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

func (obj *nativeFeatures) setNil() {
	obj.requiredValHolder = nil
	obj.optionalValHolder = nil
	obj.boundaryValHolder = nil
	obj.requiredValArrayHolder = nil
	obj.optionalValArrayHolder = nil
	obj.boundaryValArrayHolder = nil
	obj.nestedRefObjectHolder = nil
	obj.mixedObjectHolder = nil
	obj.numberTypeObjectHolder = nil
	obj.iterObjectHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// NativeFeatures is description is TBD
type NativeFeatures interface {
	Validation
	// msg marshals NativeFeatures to protobuf object *openapi.NativeFeatures
	// and doesn't set defaults
	msg() *openapi.NativeFeatures
	// setMsg unmarshals NativeFeatures from protobuf object *openapi.NativeFeatures
	// and doesn't set defaults
	setMsg(*openapi.NativeFeatures) NativeFeatures
	// provides marshal interface
	Marshal() marshalNativeFeatures
	// provides unmarshal interface
	Unmarshal() unMarshalNativeFeatures
	// validate validates NativeFeatures
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (NativeFeatures, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequiredVal returns RequiredVal, set in NativeFeatures.
	// RequiredVal is description is TBD
	RequiredVal() RequiredVal
	// SetRequiredVal assigns RequiredVal provided by user to NativeFeatures.
	// RequiredVal is description is TBD
	SetRequiredVal(value RequiredVal) NativeFeatures
	// HasRequiredVal checks if RequiredVal has been set in NativeFeatures
	HasRequiredVal() bool
	// OptionalVal returns OptionalVal, set in NativeFeatures.
	// OptionalVal is description is TBD
	OptionalVal() OptionalVal
	// SetOptionalVal assigns OptionalVal provided by user to NativeFeatures.
	// OptionalVal is description is TBD
	SetOptionalVal(value OptionalVal) NativeFeatures
	// HasOptionalVal checks if OptionalVal has been set in NativeFeatures
	HasOptionalVal() bool
	// BoundaryVal returns BoundaryVal, set in NativeFeatures.
	// BoundaryVal is description is TBD
	BoundaryVal() BoundaryVal
	// SetBoundaryVal assigns BoundaryVal provided by user to NativeFeatures.
	// BoundaryVal is description is TBD
	SetBoundaryVal(value BoundaryVal) NativeFeatures
	// HasBoundaryVal checks if BoundaryVal has been set in NativeFeatures
	HasBoundaryVal() bool
	// RequiredValArray returns RequiredValArray, set in NativeFeatures.
	// RequiredValArray is description is TBD
	RequiredValArray() RequiredValArray
	// SetRequiredValArray assigns RequiredValArray provided by user to NativeFeatures.
	// RequiredValArray is description is TBD
	SetRequiredValArray(value RequiredValArray) NativeFeatures
	// HasRequiredValArray checks if RequiredValArray has been set in NativeFeatures
	HasRequiredValArray() bool
	// OptionalValArray returns OptionalValArray, set in NativeFeatures.
	// OptionalValArray is description is TBD
	OptionalValArray() OptionalValArray
	// SetOptionalValArray assigns OptionalValArray provided by user to NativeFeatures.
	// OptionalValArray is description is TBD
	SetOptionalValArray(value OptionalValArray) NativeFeatures
	// HasOptionalValArray checks if OptionalValArray has been set in NativeFeatures
	HasOptionalValArray() bool
	// BoundaryValArray returns BoundaryValArray, set in NativeFeatures.
	// BoundaryValArray is description is TBD
	BoundaryValArray() BoundaryValArray
	// SetBoundaryValArray assigns BoundaryValArray provided by user to NativeFeatures.
	// BoundaryValArray is description is TBD
	SetBoundaryValArray(value BoundaryValArray) NativeFeatures
	// HasBoundaryValArray checks if BoundaryValArray has been set in NativeFeatures
	HasBoundaryValArray() bool
	// NestedRefObject returns NestedRefObject, set in NativeFeatures.
	// NestedRefObject is description is TBD
	NestedRefObject() NestedRefObject
	// SetNestedRefObject assigns NestedRefObject provided by user to NativeFeatures.
	// NestedRefObject is description is TBD
	SetNestedRefObject(value NestedRefObject) NativeFeatures
	// HasNestedRefObject checks if NestedRefObject has been set in NativeFeatures
	HasNestedRefObject() bool
	// MixedObject returns MixedObject, set in NativeFeatures.
	// MixedObject is format validation object
	MixedObject() MixedObject
	// SetMixedObject assigns MixedObject provided by user to NativeFeatures.
	// MixedObject is format validation object
	SetMixedObject(value MixedObject) NativeFeatures
	// HasMixedObject checks if MixedObject has been set in NativeFeatures
	HasMixedObject() bool
	// NumberTypeObject returns NumberTypeObject, set in NativeFeatures.
	// NumberTypeObject is description is TBD
	NumberTypeObject() NumberTypeObject
	// SetNumberTypeObject assigns NumberTypeObject provided by user to NativeFeatures.
	// NumberTypeObject is description is TBD
	SetNumberTypeObject(value NumberTypeObject) NativeFeatures
	// HasNumberTypeObject checks if NumberTypeObject has been set in NativeFeatures
	HasNumberTypeObject() bool
	// IterObject returns NativeFeaturesMixedObjectIterIter, set in NativeFeatures
	IterObject() NativeFeaturesMixedObjectIter
	setNil()
}

// description is TBD
// RequiredVal returns a RequiredVal
func (obj *nativeFeatures) RequiredVal() RequiredVal {
	if obj.obj.RequiredVal == nil {
		obj.obj.RequiredVal = NewRequiredVal().msg()
	}
	if obj.requiredValHolder == nil {
		obj.requiredValHolder = &requiredVal{obj: obj.obj.RequiredVal}
	}
	return obj.requiredValHolder
}

// description is TBD
// RequiredVal returns a RequiredVal
func (obj *nativeFeatures) HasRequiredVal() bool {
	return obj.obj.RequiredVal != nil
}

// description is TBD
// SetRequiredVal sets the RequiredVal value in the NativeFeatures object
func (obj *nativeFeatures) SetRequiredVal(value RequiredVal) NativeFeatures {

	obj.requiredValHolder = nil
	obj.obj.RequiredVal = value.msg()

	return obj
}

// description is TBD
// OptionalVal returns a OptionalVal
func (obj *nativeFeatures) OptionalVal() OptionalVal {
	if obj.obj.OptionalVal == nil {
		obj.obj.OptionalVal = NewOptionalVal().msg()
	}
	if obj.optionalValHolder == nil {
		obj.optionalValHolder = &optionalVal{obj: obj.obj.OptionalVal}
	}
	return obj.optionalValHolder
}

// description is TBD
// OptionalVal returns a OptionalVal
func (obj *nativeFeatures) HasOptionalVal() bool {
	return obj.obj.OptionalVal != nil
}

// description is TBD
// SetOptionalVal sets the OptionalVal value in the NativeFeatures object
func (obj *nativeFeatures) SetOptionalVal(value OptionalVal) NativeFeatures {

	obj.optionalValHolder = nil
	obj.obj.OptionalVal = value.msg()

	return obj
}

// description is TBD
// BoundaryVal returns a BoundaryVal
func (obj *nativeFeatures) BoundaryVal() BoundaryVal {
	if obj.obj.BoundaryVal == nil {
		obj.obj.BoundaryVal = NewBoundaryVal().msg()
	}
	if obj.boundaryValHolder == nil {
		obj.boundaryValHolder = &boundaryVal{obj: obj.obj.BoundaryVal}
	}
	return obj.boundaryValHolder
}

// description is TBD
// BoundaryVal returns a BoundaryVal
func (obj *nativeFeatures) HasBoundaryVal() bool {
	return obj.obj.BoundaryVal != nil
}

// description is TBD
// SetBoundaryVal sets the BoundaryVal value in the NativeFeatures object
func (obj *nativeFeatures) SetBoundaryVal(value BoundaryVal) NativeFeatures {

	obj.boundaryValHolder = nil
	obj.obj.BoundaryVal = value.msg()

	return obj
}

// description is TBD
// RequiredValArray returns a RequiredValArray
func (obj *nativeFeatures) RequiredValArray() RequiredValArray {
	if obj.obj.RequiredValArray == nil {
		obj.obj.RequiredValArray = NewRequiredValArray().msg()
	}
	if obj.requiredValArrayHolder == nil {
		obj.requiredValArrayHolder = &requiredValArray{obj: obj.obj.RequiredValArray}
	}
	return obj.requiredValArrayHolder
}

// description is TBD
// RequiredValArray returns a RequiredValArray
func (obj *nativeFeatures) HasRequiredValArray() bool {
	return obj.obj.RequiredValArray != nil
}

// description is TBD
// SetRequiredValArray sets the RequiredValArray value in the NativeFeatures object
func (obj *nativeFeatures) SetRequiredValArray(value RequiredValArray) NativeFeatures {

	obj.requiredValArrayHolder = nil
	obj.obj.RequiredValArray = value.msg()

	return obj
}

// description is TBD
// OptionalValArray returns a OptionalValArray
func (obj *nativeFeatures) OptionalValArray() OptionalValArray {
	if obj.obj.OptionalValArray == nil {
		obj.obj.OptionalValArray = NewOptionalValArray().msg()
	}
	if obj.optionalValArrayHolder == nil {
		obj.optionalValArrayHolder = &optionalValArray{obj: obj.obj.OptionalValArray}
	}
	return obj.optionalValArrayHolder
}

// description is TBD
// OptionalValArray returns a OptionalValArray
func (obj *nativeFeatures) HasOptionalValArray() bool {
	return obj.obj.OptionalValArray != nil
}

// description is TBD
// SetOptionalValArray sets the OptionalValArray value in the NativeFeatures object
func (obj *nativeFeatures) SetOptionalValArray(value OptionalValArray) NativeFeatures {

	obj.optionalValArrayHolder = nil
	obj.obj.OptionalValArray = value.msg()

	return obj
}

// description is TBD
// BoundaryValArray returns a BoundaryValArray
func (obj *nativeFeatures) BoundaryValArray() BoundaryValArray {
	if obj.obj.BoundaryValArray == nil {
		obj.obj.BoundaryValArray = NewBoundaryValArray().msg()
	}
	if obj.boundaryValArrayHolder == nil {
		obj.boundaryValArrayHolder = &boundaryValArray{obj: obj.obj.BoundaryValArray}
	}
	return obj.boundaryValArrayHolder
}

// description is TBD
// BoundaryValArray returns a BoundaryValArray
func (obj *nativeFeatures) HasBoundaryValArray() bool {
	return obj.obj.BoundaryValArray != nil
}

// description is TBD
// SetBoundaryValArray sets the BoundaryValArray value in the NativeFeatures object
func (obj *nativeFeatures) SetBoundaryValArray(value BoundaryValArray) NativeFeatures {

	obj.boundaryValArrayHolder = nil
	obj.obj.BoundaryValArray = value.msg()

	return obj
}

// description is TBD
// NestedRefObject returns a NestedRefObject
func (obj *nativeFeatures) NestedRefObject() NestedRefObject {
	if obj.obj.NestedRefObject == nil {
		obj.obj.NestedRefObject = NewNestedRefObject().msg()
	}
	if obj.nestedRefObjectHolder == nil {
		obj.nestedRefObjectHolder = &nestedRefObject{obj: obj.obj.NestedRefObject}
	}
	return obj.nestedRefObjectHolder
}

// description is TBD
// NestedRefObject returns a NestedRefObject
func (obj *nativeFeatures) HasNestedRefObject() bool {
	return obj.obj.NestedRefObject != nil
}

// description is TBD
// SetNestedRefObject sets the NestedRefObject value in the NativeFeatures object
func (obj *nativeFeatures) SetNestedRefObject(value NestedRefObject) NativeFeatures {

	obj.nestedRefObjectHolder = nil
	obj.obj.NestedRefObject = value.msg()

	return obj
}

// description is TBD
// MixedObject returns a MixedObject
func (obj *nativeFeatures) MixedObject() MixedObject {
	if obj.obj.MixedObject == nil {
		obj.obj.MixedObject = NewMixedObject().msg()
	}
	if obj.mixedObjectHolder == nil {
		obj.mixedObjectHolder = &mixedObject{obj: obj.obj.MixedObject}
	}
	return obj.mixedObjectHolder
}

// description is TBD
// MixedObject returns a MixedObject
func (obj *nativeFeatures) HasMixedObject() bool {
	return obj.obj.MixedObject != nil
}

// description is TBD
// SetMixedObject sets the MixedObject value in the NativeFeatures object
func (obj *nativeFeatures) SetMixedObject(value MixedObject) NativeFeatures {

	obj.mixedObjectHolder = nil
	obj.obj.MixedObject = value.msg()

	return obj
}

// description is TBD
// NumberTypeObject returns a NumberTypeObject
func (obj *nativeFeatures) NumberTypeObject() NumberTypeObject {
	if obj.obj.NumberTypeObject == nil {
		obj.obj.NumberTypeObject = NewNumberTypeObject().msg()
	}
	if obj.numberTypeObjectHolder == nil {
		obj.numberTypeObjectHolder = &numberTypeObject{obj: obj.obj.NumberTypeObject}
	}
	return obj.numberTypeObjectHolder
}

// description is TBD
// NumberTypeObject returns a NumberTypeObject
func (obj *nativeFeatures) HasNumberTypeObject() bool {
	return obj.obj.NumberTypeObject != nil
}

// description is TBD
// SetNumberTypeObject sets the NumberTypeObject value in the NativeFeatures object
func (obj *nativeFeatures) SetNumberTypeObject(value NumberTypeObject) NativeFeatures {

	obj.numberTypeObjectHolder = nil
	obj.obj.NumberTypeObject = value.msg()

	return obj
}

// description is TBD
// IterObject returns a []MixedObject
func (obj *nativeFeatures) IterObject() NativeFeaturesMixedObjectIter {
	if len(obj.obj.IterObject) == 0 {
		obj.obj.IterObject = []*openapi.MixedObject{}
	}
	if obj.iterObjectHolder == nil {
		obj.iterObjectHolder = newNativeFeaturesMixedObjectIter(&obj.obj.IterObject).setMsg(obj)
	}
	return obj.iterObjectHolder
}

type nativeFeaturesMixedObjectIter struct {
	obj              *nativeFeatures
	mixedObjectSlice []MixedObject
	fieldPtr         *[]*openapi.MixedObject
}

func newNativeFeaturesMixedObjectIter(ptr *[]*openapi.MixedObject) NativeFeaturesMixedObjectIter {
	return &nativeFeaturesMixedObjectIter{fieldPtr: ptr}
}

type NativeFeaturesMixedObjectIter interface {
	setMsg(*nativeFeatures) NativeFeaturesMixedObjectIter
	Items() []MixedObject
	Add() MixedObject
	Append(items ...MixedObject) NativeFeaturesMixedObjectIter
	Set(index int, newObj MixedObject) NativeFeaturesMixedObjectIter
	Clear() NativeFeaturesMixedObjectIter
	clearHolderSlice() NativeFeaturesMixedObjectIter
	appendHolderSlice(item MixedObject) NativeFeaturesMixedObjectIter
}

func (obj *nativeFeaturesMixedObjectIter) setMsg(msg *nativeFeatures) NativeFeaturesMixedObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&mixedObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *nativeFeaturesMixedObjectIter) Items() []MixedObject {
	return obj.mixedObjectSlice
}

func (obj *nativeFeaturesMixedObjectIter) Add() MixedObject {
	newObj := &openapi.MixedObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &mixedObject{obj: newObj}
	newLibObj.setDefault()
	obj.mixedObjectSlice = append(obj.mixedObjectSlice, newLibObj)
	return newLibObj
}

func (obj *nativeFeaturesMixedObjectIter) Append(items ...MixedObject) NativeFeaturesMixedObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.mixedObjectSlice = append(obj.mixedObjectSlice, item)
	}
	return obj
}

func (obj *nativeFeaturesMixedObjectIter) Set(index int, newObj MixedObject) NativeFeaturesMixedObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.mixedObjectSlice[index] = newObj
	return obj
}
func (obj *nativeFeaturesMixedObjectIter) Clear() NativeFeaturesMixedObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.MixedObject{}
		obj.mixedObjectSlice = []MixedObject{}
	}
	return obj
}
func (obj *nativeFeaturesMixedObjectIter) clearHolderSlice() NativeFeaturesMixedObjectIter {
	if len(obj.mixedObjectSlice) > 0 {
		obj.mixedObjectSlice = []MixedObject{}
	}
	return obj
}
func (obj *nativeFeaturesMixedObjectIter) appendHolderSlice(item MixedObject) NativeFeaturesMixedObjectIter {
	obj.mixedObjectSlice = append(obj.mixedObjectSlice, item)
	return obj
}

func (obj *nativeFeatures) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RequiredVal != nil {

		obj.RequiredVal().validateObj(vObj, set_default)
	}

	if obj.obj.OptionalVal != nil {

		obj.OptionalVal().validateObj(vObj, set_default)
	}

	if obj.obj.BoundaryVal != nil {

		obj.BoundaryVal().validateObj(vObj, set_default)
	}

	if obj.obj.RequiredValArray != nil {

		obj.RequiredValArray().validateObj(vObj, set_default)
	}

	if obj.obj.OptionalValArray != nil {

		obj.OptionalValArray().validateObj(vObj, set_default)
	}

	if obj.obj.BoundaryValArray != nil {

		obj.BoundaryValArray().validateObj(vObj, set_default)
	}

	if obj.obj.NestedRefObject != nil {

		obj.NestedRefObject().validateObj(vObj, set_default)
	}

	if obj.obj.MixedObject != nil {

		obj.MixedObject().validateObj(vObj, set_default)
	}

	if obj.obj.NumberTypeObject != nil {

		obj.NumberTypeObject().validateObj(vObj, set_default)
	}

	if len(obj.obj.IterObject) != 0 {

		if set_default {
			obj.IterObject().clearHolderSlice()
			for _, item := range obj.obj.IterObject {
				obj.IterObject().appendHolderSlice(&mixedObject{obj: item})
			}
		}
		for _, item := range obj.IterObject().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *nativeFeatures) setDefault() {

}
