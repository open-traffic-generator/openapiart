package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ExtendedFeatures *****
type extendedFeatures struct {
	validation
	obj                         *openapi.ExtendedFeatures
	marshaller                  marshalExtendedFeatures
	unMarshaller                unMarshalExtendedFeatures
	choiceValHolder             ChoiceVal
	choiceValNoPropertiesHolder ChoiceValWithNoProperties
	xStatusObjectHolder         XStatusObject
	xEnumObjectHolder           XEnumObject
	xFieldPatternObjectHolder   XFieldPatternObject
}

func NewExtendedFeatures() ExtendedFeatures {
	obj := extendedFeatures{obj: &openapi.ExtendedFeatures{}}
	obj.setDefault()
	return &obj
}

func (obj *extendedFeatures) msg() *openapi.ExtendedFeatures {
	return obj.obj
}

func (obj *extendedFeatures) setMsg(msg *openapi.ExtendedFeatures) ExtendedFeatures {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalextendedFeatures struct {
	obj *extendedFeatures
}

type marshalExtendedFeatures interface {
	// ToProto marshals ExtendedFeatures to protobuf object *openapi.ExtendedFeatures
	ToProto() (*openapi.ExtendedFeatures, error)
	// ToPbText marshals ExtendedFeatures to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ExtendedFeatures to YAML text
	ToYaml() (string, error)
	// ToJson marshals ExtendedFeatures to JSON text
	ToJson() (string, error)
}

type unMarshalextendedFeatures struct {
	obj *extendedFeatures
}

type unMarshalExtendedFeatures interface {
	// FromProto unmarshals ExtendedFeatures from protobuf object *openapi.ExtendedFeatures
	FromProto(msg *openapi.ExtendedFeatures) (ExtendedFeatures, error)
	// FromPbText unmarshals ExtendedFeatures from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ExtendedFeatures from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ExtendedFeatures from JSON text
	FromJson(value string) error
}

func (obj *extendedFeatures) Marshal() marshalExtendedFeatures {
	if obj.marshaller == nil {
		obj.marshaller = &marshalextendedFeatures{obj: obj}
	}
	return obj.marshaller
}

func (obj *extendedFeatures) Unmarshal() unMarshalExtendedFeatures {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalextendedFeatures{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalextendedFeatures) ToProto() (*openapi.ExtendedFeatures, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalextendedFeatures) FromProto(msg *openapi.ExtendedFeatures) (ExtendedFeatures, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalextendedFeatures) ToPbText() (string, error) {
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

func (m *unMarshalextendedFeatures) FromPbText(value string) error {
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

func (m *marshalextendedFeatures) ToYaml() (string, error) {
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

func (m *unMarshalextendedFeatures) FromYaml(value string) error {
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

func (m *marshalextendedFeatures) ToJson() (string, error) {
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

func (m *unMarshalextendedFeatures) FromJson(value string) error {
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

func (obj *extendedFeatures) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *extendedFeatures) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *extendedFeatures) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *extendedFeatures) Clone() (ExtendedFeatures, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewExtendedFeatures()
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

func (obj *extendedFeatures) setNil() {
	obj.choiceValHolder = nil
	obj.choiceValNoPropertiesHolder = nil
	obj.xStatusObjectHolder = nil
	obj.xEnumObjectHolder = nil
	obj.xFieldPatternObjectHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ExtendedFeatures is description is TBD
type ExtendedFeatures interface {
	Validation
	// msg marshals ExtendedFeatures to protobuf object *openapi.ExtendedFeatures
	// and doesn't set defaults
	msg() *openapi.ExtendedFeatures
	// setMsg unmarshals ExtendedFeatures from protobuf object *openapi.ExtendedFeatures
	// and doesn't set defaults
	setMsg(*openapi.ExtendedFeatures) ExtendedFeatures
	// provides marshal interface
	Marshal() marshalExtendedFeatures
	// provides unmarshal interface
	Unmarshal() unMarshalExtendedFeatures
	// validate validates ExtendedFeatures
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ExtendedFeatures, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ChoiceVal returns ChoiceVal, set in ExtendedFeatures.
	// ChoiceVal is description is TBD
	ChoiceVal() ChoiceVal
	// SetChoiceVal assigns ChoiceVal provided by user to ExtendedFeatures.
	// ChoiceVal is description is TBD
	SetChoiceVal(value ChoiceVal) ExtendedFeatures
	// HasChoiceVal checks if ChoiceVal has been set in ExtendedFeatures
	HasChoiceVal() bool
	// ChoiceValNoProperties returns ChoiceValWithNoProperties, set in ExtendedFeatures.
	// ChoiceValWithNoProperties is description is TBD
	ChoiceValNoProperties() ChoiceValWithNoProperties
	// SetChoiceValNoProperties assigns ChoiceValWithNoProperties provided by user to ExtendedFeatures.
	// ChoiceValWithNoProperties is description is TBD
	SetChoiceValNoProperties(value ChoiceValWithNoProperties) ExtendedFeatures
	// HasChoiceValNoProperties checks if ChoiceValNoProperties has been set in ExtendedFeatures
	HasChoiceValNoProperties() bool
	// XStatusObject returns XStatusObject, set in ExtendedFeatures.
	// XStatusObject is description is TBD
	XStatusObject() XStatusObject
	// SetXStatusObject assigns XStatusObject provided by user to ExtendedFeatures.
	// XStatusObject is description is TBD
	SetXStatusObject(value XStatusObject) ExtendedFeatures
	// HasXStatusObject checks if XStatusObject has been set in ExtendedFeatures
	HasXStatusObject() bool
	// XEnumObject returns XEnumObject, set in ExtendedFeatures.
	// XEnumObject is description is TBD
	XEnumObject() XEnumObject
	// SetXEnumObject assigns XEnumObject provided by user to ExtendedFeatures.
	// XEnumObject is description is TBD
	SetXEnumObject(value XEnumObject) ExtendedFeatures
	// HasXEnumObject checks if XEnumObject has been set in ExtendedFeatures
	HasXEnumObject() bool
	// XFieldPatternObject returns XFieldPatternObject, set in ExtendedFeatures.
	// XFieldPatternObject is description is TBD
	XFieldPatternObject() XFieldPatternObject
	// SetXFieldPatternObject assigns XFieldPatternObject provided by user to ExtendedFeatures.
	// XFieldPatternObject is description is TBD
	SetXFieldPatternObject(value XFieldPatternObject) ExtendedFeatures
	// HasXFieldPatternObject checks if XFieldPatternObject has been set in ExtendedFeatures
	HasXFieldPatternObject() bool
	setNil()
}

// description is TBD
// ChoiceVal returns a ChoiceVal
func (obj *extendedFeatures) ChoiceVal() ChoiceVal {
	if obj.obj.ChoiceVal == nil {
		obj.obj.ChoiceVal = NewChoiceVal().msg()
	}
	if obj.choiceValHolder == nil {
		obj.choiceValHolder = &choiceVal{obj: obj.obj.ChoiceVal}
	}
	return obj.choiceValHolder
}

// description is TBD
// ChoiceVal returns a ChoiceVal
func (obj *extendedFeatures) HasChoiceVal() bool {
	return obj.obj.ChoiceVal != nil
}

// description is TBD
// SetChoiceVal sets the ChoiceVal value in the ExtendedFeatures object
func (obj *extendedFeatures) SetChoiceVal(value ChoiceVal) ExtendedFeatures {

	obj.choiceValHolder = nil
	obj.obj.ChoiceVal = value.msg()

	return obj
}

// description is TBD
// ChoiceValNoProperties returns a ChoiceValWithNoProperties
func (obj *extendedFeatures) ChoiceValNoProperties() ChoiceValWithNoProperties {
	if obj.obj.ChoiceValNoProperties == nil {
		obj.obj.ChoiceValNoProperties = NewChoiceValWithNoProperties().msg()
	}
	if obj.choiceValNoPropertiesHolder == nil {
		obj.choiceValNoPropertiesHolder = &choiceValWithNoProperties{obj: obj.obj.ChoiceValNoProperties}
	}
	return obj.choiceValNoPropertiesHolder
}

// description is TBD
// ChoiceValNoProperties returns a ChoiceValWithNoProperties
func (obj *extendedFeatures) HasChoiceValNoProperties() bool {
	return obj.obj.ChoiceValNoProperties != nil
}

// description is TBD
// SetChoiceValNoProperties sets the ChoiceValWithNoProperties value in the ExtendedFeatures object
func (obj *extendedFeatures) SetChoiceValNoProperties(value ChoiceValWithNoProperties) ExtendedFeatures {

	obj.choiceValNoPropertiesHolder = nil
	obj.obj.ChoiceValNoProperties = value.msg()

	return obj
}

// Under Review: test under_review
//
// Description TBD
// XStatusObject returns a XStatusObject
func (obj *extendedFeatures) XStatusObject() XStatusObject {
	if obj.obj.XStatusObject == nil {
		obj.obj.XStatusObject = NewXStatusObject().msg()
	}
	if obj.xStatusObjectHolder == nil {
		obj.xStatusObjectHolder = &xStatusObject{obj: obj.obj.XStatusObject}
	}
	return obj.xStatusObjectHolder
}

// Under Review: test under_review
//
// Description TBD
// XStatusObject returns a XStatusObject
func (obj *extendedFeatures) HasXStatusObject() bool {
	return obj.obj.XStatusObject != nil
}

// Under Review: test under_review
//
// Description TBD
// SetXStatusObject sets the XStatusObject value in the ExtendedFeatures object
func (obj *extendedFeatures) SetXStatusObject(value XStatusObject) ExtendedFeatures {

	obj.xStatusObjectHolder = nil
	obj.obj.XStatusObject = value.msg()

	return obj
}

// description is TBD
// XEnumObject returns a XEnumObject
func (obj *extendedFeatures) XEnumObject() XEnumObject {
	if obj.obj.XEnumObject == nil {
		obj.obj.XEnumObject = NewXEnumObject().msg()
	}
	if obj.xEnumObjectHolder == nil {
		obj.xEnumObjectHolder = &xEnumObject{obj: obj.obj.XEnumObject}
	}
	return obj.xEnumObjectHolder
}

// description is TBD
// XEnumObject returns a XEnumObject
func (obj *extendedFeatures) HasXEnumObject() bool {
	return obj.obj.XEnumObject != nil
}

// description is TBD
// SetXEnumObject sets the XEnumObject value in the ExtendedFeatures object
func (obj *extendedFeatures) SetXEnumObject(value XEnumObject) ExtendedFeatures {

	obj.xEnumObjectHolder = nil
	obj.obj.XEnumObject = value.msg()

	return obj
}

// description is TBD
// XFieldPatternObject returns a XFieldPatternObject
func (obj *extendedFeatures) XFieldPatternObject() XFieldPatternObject {
	if obj.obj.XFieldPatternObject == nil {
		obj.obj.XFieldPatternObject = NewXFieldPatternObject().msg()
	}
	if obj.xFieldPatternObjectHolder == nil {
		obj.xFieldPatternObjectHolder = &xFieldPatternObject{obj: obj.obj.XFieldPatternObject}
	}
	return obj.xFieldPatternObjectHolder
}

// description is TBD
// XFieldPatternObject returns a XFieldPatternObject
func (obj *extendedFeatures) HasXFieldPatternObject() bool {
	return obj.obj.XFieldPatternObject != nil
}

// description is TBD
// SetXFieldPatternObject sets the XFieldPatternObject value in the ExtendedFeatures object
func (obj *extendedFeatures) SetXFieldPatternObject(value XFieldPatternObject) ExtendedFeatures {

	obj.xFieldPatternObjectHolder = nil
	obj.obj.XFieldPatternObject = value.msg()

	return obj
}

func (obj *extendedFeatures) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ChoiceVal != nil {

		obj.ChoiceVal().validateObj(vObj, set_default)
	}

	if obj.obj.ChoiceValNoProperties != nil {

		obj.ChoiceValNoProperties().validateObj(vObj, set_default)
	}

	if obj.obj.XStatusObject != nil {
		obj.addWarnings("XStatusObject property in schema ExtendedFeatures is under review, test under_review")
		obj.XStatusObject().validateObj(vObj, set_default)
	}

	if obj.obj.XEnumObject != nil {

		obj.XEnumObject().validateObj(vObj, set_default)
	}

	if obj.obj.XFieldPatternObject != nil {

		obj.XFieldPatternObject().validateObj(vObj, set_default)
	}

}

func (obj *extendedFeatures) setDefault() {

}
