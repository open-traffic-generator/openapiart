package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChecksumPatternObject *****
type checksumPatternObject struct {
	validation
	obj            *openapi.ChecksumPatternObject
	marshaller     marshalChecksumPatternObject
	unMarshaller   unMarshalChecksumPatternObject
	checksumHolder PatternChecksumPatternObjectChecksum
}

func NewChecksumPatternObject() ChecksumPatternObject {
	obj := checksumPatternObject{obj: &openapi.ChecksumPatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *checksumPatternObject) msg() *openapi.ChecksumPatternObject {
	return obj.obj
}

func (obj *checksumPatternObject) setMsg(msg *openapi.ChecksumPatternObject) ChecksumPatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchecksumPatternObject struct {
	obj *checksumPatternObject
}

type marshalChecksumPatternObject interface {
	// ToProto marshals ChecksumPatternObject to protobuf object *openapi.ChecksumPatternObject
	ToProto() (*openapi.ChecksumPatternObject, error)
	// ToPbText marshals ChecksumPatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChecksumPatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChecksumPatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalchecksumPatternObject struct {
	obj *checksumPatternObject
}

type unMarshalChecksumPatternObject interface {
	// FromProto unmarshals ChecksumPatternObject from protobuf object *openapi.ChecksumPatternObject
	FromProto(msg *openapi.ChecksumPatternObject) (ChecksumPatternObject, error)
	// FromPbText unmarshals ChecksumPatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChecksumPatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChecksumPatternObject from JSON text
	FromJson(value string) error
}

func (obj *checksumPatternObject) Marshal() marshalChecksumPatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchecksumPatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *checksumPatternObject) Unmarshal() unMarshalChecksumPatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchecksumPatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchecksumPatternObject) ToProto() (*openapi.ChecksumPatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchecksumPatternObject) FromProto(msg *openapi.ChecksumPatternObject) (ChecksumPatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchecksumPatternObject) ToPbText() (string, error) {
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

func (m *unMarshalchecksumPatternObject) FromPbText(value string) error {
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

func (m *marshalchecksumPatternObject) ToYaml() (string, error) {
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

func (m *unMarshalchecksumPatternObject) FromYaml(value string) error {
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

func (m *marshalchecksumPatternObject) ToJson() (string, error) {
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

func (m *unMarshalchecksumPatternObject) FromJson(value string) error {
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

func (obj *checksumPatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *checksumPatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *checksumPatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *checksumPatternObject) Clone() (ChecksumPatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChecksumPatternObject()
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

func (obj *checksumPatternObject) setNil() {
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChecksumPatternObject is test checksum pattern
type ChecksumPatternObject interface {
	Validation
	// msg marshals ChecksumPatternObject to protobuf object *openapi.ChecksumPatternObject
	// and doesn't set defaults
	msg() *openapi.ChecksumPatternObject
	// setMsg unmarshals ChecksumPatternObject from protobuf object *openapi.ChecksumPatternObject
	// and doesn't set defaults
	setMsg(*openapi.ChecksumPatternObject) ChecksumPatternObject
	// provides marshal interface
	Marshal() marshalChecksumPatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalChecksumPatternObject
	// validate validates ChecksumPatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChecksumPatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Checksum returns PatternChecksumPatternObjectChecksum, set in ChecksumPatternObject.
	// PatternChecksumPatternObjectChecksum is tBD
	Checksum() PatternChecksumPatternObjectChecksum
	// SetChecksum assigns PatternChecksumPatternObjectChecksum provided by user to ChecksumPatternObject.
	// PatternChecksumPatternObjectChecksum is tBD
	SetChecksum(value PatternChecksumPatternObjectChecksum) ChecksumPatternObject
	// HasChecksum checks if Checksum has been set in ChecksumPatternObject
	HasChecksum() bool
	setNil()
}

// description is TBD
// Checksum returns a PatternChecksumPatternObjectChecksum
func (obj *checksumPatternObject) Checksum() PatternChecksumPatternObjectChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternChecksumPatternObjectChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternChecksumPatternObjectChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternChecksumPatternObjectChecksum
func (obj *checksumPatternObject) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternChecksumPatternObjectChecksum value in the ChecksumPatternObject object
func (obj *checksumPatternObject) SetChecksum(value PatternChecksumPatternObjectChecksum) ChecksumPatternObject {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

func (obj *checksumPatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *checksumPatternObject) setDefault() {

}
