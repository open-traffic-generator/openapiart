package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** XFieldPatternObject *****
type xFieldPatternObject struct {
	validation
	obj                   *openapi.XFieldPatternObject
	marshaller            marshalXFieldPatternObject
	unMarshaller          unMarshalXFieldPatternObject
	ipv4PatternHolder     Ipv4PatternObject
	ipv6PatternHolder     Ipv6PatternObject
	macPatternHolder      MacPatternObject
	integerPatternHolder  IntegerPatternObject
	checksumPatternHolder ChecksumPatternObject
}

func NewXFieldPatternObject() XFieldPatternObject {
	obj := xFieldPatternObject{obj: &openapi.XFieldPatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *xFieldPatternObject) msg() *openapi.XFieldPatternObject {
	return obj.obj
}

func (obj *xFieldPatternObject) setMsg(msg *openapi.XFieldPatternObject) XFieldPatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalxFieldPatternObject struct {
	obj *xFieldPatternObject
}

type marshalXFieldPatternObject interface {
	// ToProto marshals XFieldPatternObject to protobuf object *openapi.XFieldPatternObject
	ToProto() (*openapi.XFieldPatternObject, error)
	// ToPbText marshals XFieldPatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals XFieldPatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals XFieldPatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalxFieldPatternObject struct {
	obj *xFieldPatternObject
}

type unMarshalXFieldPatternObject interface {
	// FromProto unmarshals XFieldPatternObject from protobuf object *openapi.XFieldPatternObject
	FromProto(msg *openapi.XFieldPatternObject) (XFieldPatternObject, error)
	// FromPbText unmarshals XFieldPatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals XFieldPatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals XFieldPatternObject from JSON text
	FromJson(value string) error
}

func (obj *xFieldPatternObject) Marshal() marshalXFieldPatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalxFieldPatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *xFieldPatternObject) Unmarshal() unMarshalXFieldPatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalxFieldPatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalxFieldPatternObject) ToProto() (*openapi.XFieldPatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalxFieldPatternObject) FromProto(msg *openapi.XFieldPatternObject) (XFieldPatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalxFieldPatternObject) ToPbText() (string, error) {
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

func (m *unMarshalxFieldPatternObject) FromPbText(value string) error {
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

func (m *marshalxFieldPatternObject) ToYaml() (string, error) {
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

func (m *unMarshalxFieldPatternObject) FromYaml(value string) error {
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

func (m *marshalxFieldPatternObject) ToJson() (string, error) {
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

func (m *unMarshalxFieldPatternObject) FromJson(value string) error {
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

func (obj *xFieldPatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *xFieldPatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *xFieldPatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *xFieldPatternObject) Clone() (XFieldPatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewXFieldPatternObject()
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

func (obj *xFieldPatternObject) setNil() {
	obj.ipv4PatternHolder = nil
	obj.ipv6PatternHolder = nil
	obj.macPatternHolder = nil
	obj.integerPatternHolder = nil
	obj.checksumPatternHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// XFieldPatternObject is description is TBD
type XFieldPatternObject interface {
	Validation
	// msg marshals XFieldPatternObject to protobuf object *openapi.XFieldPatternObject
	// and doesn't set defaults
	msg() *openapi.XFieldPatternObject
	// setMsg unmarshals XFieldPatternObject from protobuf object *openapi.XFieldPatternObject
	// and doesn't set defaults
	setMsg(*openapi.XFieldPatternObject) XFieldPatternObject
	// provides marshal interface
	Marshal() marshalXFieldPatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalXFieldPatternObject
	// validate validates XFieldPatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (XFieldPatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Pattern returns Ipv4PatternObject, set in XFieldPatternObject.
	// Ipv4PatternObject is test ipv4 pattern
	Ipv4Pattern() Ipv4PatternObject
	// SetIpv4Pattern assigns Ipv4PatternObject provided by user to XFieldPatternObject.
	// Ipv4PatternObject is test ipv4 pattern
	SetIpv4Pattern(value Ipv4PatternObject) XFieldPatternObject
	// HasIpv4Pattern checks if Ipv4Pattern has been set in XFieldPatternObject
	HasIpv4Pattern() bool
	// Ipv6Pattern returns Ipv6PatternObject, set in XFieldPatternObject.
	// Ipv6PatternObject is test ipv6 pattern
	Ipv6Pattern() Ipv6PatternObject
	// SetIpv6Pattern assigns Ipv6PatternObject provided by user to XFieldPatternObject.
	// Ipv6PatternObject is test ipv6 pattern
	SetIpv6Pattern(value Ipv6PatternObject) XFieldPatternObject
	// HasIpv6Pattern checks if Ipv6Pattern has been set in XFieldPatternObject
	HasIpv6Pattern() bool
	// MacPattern returns MacPatternObject, set in XFieldPatternObject.
	// MacPatternObject is test mac pattern
	MacPattern() MacPatternObject
	// SetMacPattern assigns MacPatternObject provided by user to XFieldPatternObject.
	// MacPatternObject is test mac pattern
	SetMacPattern(value MacPatternObject) XFieldPatternObject
	// HasMacPattern checks if MacPattern has been set in XFieldPatternObject
	HasMacPattern() bool
	// IntegerPattern returns IntegerPatternObject, set in XFieldPatternObject.
	// IntegerPatternObject is test integer pattern
	IntegerPattern() IntegerPatternObject
	// SetIntegerPattern assigns IntegerPatternObject provided by user to XFieldPatternObject.
	// IntegerPatternObject is test integer pattern
	SetIntegerPattern(value IntegerPatternObject) XFieldPatternObject
	// HasIntegerPattern checks if IntegerPattern has been set in XFieldPatternObject
	HasIntegerPattern() bool
	// ChecksumPattern returns ChecksumPatternObject, set in XFieldPatternObject.
	// ChecksumPatternObject is test checksum pattern
	ChecksumPattern() ChecksumPatternObject
	// SetChecksumPattern assigns ChecksumPatternObject provided by user to XFieldPatternObject.
	// ChecksumPatternObject is test checksum pattern
	SetChecksumPattern(value ChecksumPatternObject) XFieldPatternObject
	// HasChecksumPattern checks if ChecksumPattern has been set in XFieldPatternObject
	HasChecksumPattern() bool
	setNil()
}

// description is TBD
// Ipv4Pattern returns a Ipv4PatternObject
func (obj *xFieldPatternObject) Ipv4Pattern() Ipv4PatternObject {
	if obj.obj.Ipv4Pattern == nil {
		obj.obj.Ipv4Pattern = NewIpv4PatternObject().msg()
	}
	if obj.ipv4PatternHolder == nil {
		obj.ipv4PatternHolder = &ipv4PatternObject{obj: obj.obj.Ipv4Pattern}
	}
	return obj.ipv4PatternHolder
}

// description is TBD
// Ipv4Pattern returns a Ipv4PatternObject
func (obj *xFieldPatternObject) HasIpv4Pattern() bool {
	return obj.obj.Ipv4Pattern != nil
}

// description is TBD
// SetIpv4Pattern sets the Ipv4PatternObject value in the XFieldPatternObject object
func (obj *xFieldPatternObject) SetIpv4Pattern(value Ipv4PatternObject) XFieldPatternObject {

	obj.ipv4PatternHolder = nil
	obj.obj.Ipv4Pattern = value.msg()

	return obj
}

// description is TBD
// Ipv6Pattern returns a Ipv6PatternObject
func (obj *xFieldPatternObject) Ipv6Pattern() Ipv6PatternObject {
	if obj.obj.Ipv6Pattern == nil {
		obj.obj.Ipv6Pattern = NewIpv6PatternObject().msg()
	}
	if obj.ipv6PatternHolder == nil {
		obj.ipv6PatternHolder = &ipv6PatternObject{obj: obj.obj.Ipv6Pattern}
	}
	return obj.ipv6PatternHolder
}

// description is TBD
// Ipv6Pattern returns a Ipv6PatternObject
func (obj *xFieldPatternObject) HasIpv6Pattern() bool {
	return obj.obj.Ipv6Pattern != nil
}

// description is TBD
// SetIpv6Pattern sets the Ipv6PatternObject value in the XFieldPatternObject object
func (obj *xFieldPatternObject) SetIpv6Pattern(value Ipv6PatternObject) XFieldPatternObject {

	obj.ipv6PatternHolder = nil
	obj.obj.Ipv6Pattern = value.msg()

	return obj
}

// description is TBD
// MacPattern returns a MacPatternObject
func (obj *xFieldPatternObject) MacPattern() MacPatternObject {
	if obj.obj.MacPattern == nil {
		obj.obj.MacPattern = NewMacPatternObject().msg()
	}
	if obj.macPatternHolder == nil {
		obj.macPatternHolder = &macPatternObject{obj: obj.obj.MacPattern}
	}
	return obj.macPatternHolder
}

// description is TBD
// MacPattern returns a MacPatternObject
func (obj *xFieldPatternObject) HasMacPattern() bool {
	return obj.obj.MacPattern != nil
}

// description is TBD
// SetMacPattern sets the MacPatternObject value in the XFieldPatternObject object
func (obj *xFieldPatternObject) SetMacPattern(value MacPatternObject) XFieldPatternObject {

	obj.macPatternHolder = nil
	obj.obj.MacPattern = value.msg()

	return obj
}

// description is TBD
// IntegerPattern returns a IntegerPatternObject
func (obj *xFieldPatternObject) IntegerPattern() IntegerPatternObject {
	if obj.obj.IntegerPattern == nil {
		obj.obj.IntegerPattern = NewIntegerPatternObject().msg()
	}
	if obj.integerPatternHolder == nil {
		obj.integerPatternHolder = &integerPatternObject{obj: obj.obj.IntegerPattern}
	}
	return obj.integerPatternHolder
}

// description is TBD
// IntegerPattern returns a IntegerPatternObject
func (obj *xFieldPatternObject) HasIntegerPattern() bool {
	return obj.obj.IntegerPattern != nil
}

// description is TBD
// SetIntegerPattern sets the IntegerPatternObject value in the XFieldPatternObject object
func (obj *xFieldPatternObject) SetIntegerPattern(value IntegerPatternObject) XFieldPatternObject {

	obj.integerPatternHolder = nil
	obj.obj.IntegerPattern = value.msg()

	return obj
}

// description is TBD
// ChecksumPattern returns a ChecksumPatternObject
func (obj *xFieldPatternObject) ChecksumPattern() ChecksumPatternObject {
	if obj.obj.ChecksumPattern == nil {
		obj.obj.ChecksumPattern = NewChecksumPatternObject().msg()
	}
	if obj.checksumPatternHolder == nil {
		obj.checksumPatternHolder = &checksumPatternObject{obj: obj.obj.ChecksumPattern}
	}
	return obj.checksumPatternHolder
}

// description is TBD
// ChecksumPattern returns a ChecksumPatternObject
func (obj *xFieldPatternObject) HasChecksumPattern() bool {
	return obj.obj.ChecksumPattern != nil
}

// description is TBD
// SetChecksumPattern sets the ChecksumPatternObject value in the XFieldPatternObject object
func (obj *xFieldPatternObject) SetChecksumPattern(value ChecksumPatternObject) XFieldPatternObject {

	obj.checksumPatternHolder = nil
	obj.obj.ChecksumPattern = value.msg()

	return obj
}

func (obj *xFieldPatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Pattern != nil {

		obj.Ipv4Pattern().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Pattern != nil {

		obj.Ipv6Pattern().validateObj(vObj, set_default)
	}

	if obj.obj.MacPattern != nil {

		obj.MacPattern().validateObj(vObj, set_default)
	}

	if obj.obj.IntegerPattern != nil {

		obj.IntegerPattern().validateObj(vObj, set_default)
	}

	if obj.obj.ChecksumPattern != nil {

		obj.ChecksumPattern().validateObj(vObj, set_default)
	}

}

func (obj *xFieldPatternObject) setDefault() {

}
