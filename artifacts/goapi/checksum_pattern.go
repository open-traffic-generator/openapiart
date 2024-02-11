package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ChecksumPattern *****
type checksumPattern struct {
	validation
	obj            *openapi.ChecksumPattern
	marshaller     marshalChecksumPattern
	unMarshaller   unMarshalChecksumPattern
	checksumHolder PatternChecksumPatternChecksum
}

func NewChecksumPattern() ChecksumPattern {
	obj := checksumPattern{obj: &openapi.ChecksumPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *checksumPattern) msg() *openapi.ChecksumPattern {
	return obj.obj
}

func (obj *checksumPattern) setMsg(msg *openapi.ChecksumPattern) ChecksumPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalchecksumPattern struct {
	obj *checksumPattern
}

type marshalChecksumPattern interface {
	// ToProto marshals ChecksumPattern to protobuf object *openapi.ChecksumPattern
	ToProto() (*openapi.ChecksumPattern, error)
	// ToPbText marshals ChecksumPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChecksumPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChecksumPattern to JSON text
	ToJson() (string, error)
}

type unMarshalchecksumPattern struct {
	obj *checksumPattern
}

type unMarshalChecksumPattern interface {
	// FromProto unmarshals ChecksumPattern from protobuf object *openapi.ChecksumPattern
	FromProto(msg *openapi.ChecksumPattern) (ChecksumPattern, error)
	// FromPbText unmarshals ChecksumPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChecksumPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChecksumPattern from JSON text
	FromJson(value string) error
}

func (obj *checksumPattern) Marshal() marshalChecksumPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalchecksumPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *checksumPattern) Unmarshal() unMarshalChecksumPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalchecksumPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalchecksumPattern) ToProto() (*openapi.ChecksumPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalchecksumPattern) FromProto(msg *openapi.ChecksumPattern) (ChecksumPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalchecksumPattern) ToPbText() (string, error) {
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

func (m *unMarshalchecksumPattern) FromPbText(value string) error {
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

func (m *marshalchecksumPattern) ToYaml() (string, error) {
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

func (m *unMarshalchecksumPattern) FromYaml(value string) error {
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

func (m *marshalchecksumPattern) ToJson() (string, error) {
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

func (m *unMarshalchecksumPattern) FromJson(value string) error {
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

func (obj *checksumPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *checksumPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *checksumPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *checksumPattern) Clone() (ChecksumPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChecksumPattern()
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

func (obj *checksumPattern) setNil() {
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChecksumPattern is test checksum pattern
type ChecksumPattern interface {
	Validation
	// msg marshals ChecksumPattern to protobuf object *openapi.ChecksumPattern
	// and doesn't set defaults
	msg() *openapi.ChecksumPattern
	// setMsg unmarshals ChecksumPattern from protobuf object *openapi.ChecksumPattern
	// and doesn't set defaults
	setMsg(*openapi.ChecksumPattern) ChecksumPattern
	// provides marshal interface
	Marshal() marshalChecksumPattern
	// provides unmarshal interface
	Unmarshal() unMarshalChecksumPattern
	// validate validates ChecksumPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChecksumPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Checksum returns PatternChecksumPatternChecksum, set in ChecksumPattern.
	// PatternChecksumPatternChecksum is tBD
	Checksum() PatternChecksumPatternChecksum
	// SetChecksum assigns PatternChecksumPatternChecksum provided by user to ChecksumPattern.
	// PatternChecksumPatternChecksum is tBD
	SetChecksum(value PatternChecksumPatternChecksum) ChecksumPattern
	// HasChecksum checks if Checksum has been set in ChecksumPattern
	HasChecksum() bool
	setNil()
}

// description is TBD
// Checksum returns a PatternChecksumPatternChecksum
func (obj *checksumPattern) Checksum() PatternChecksumPatternChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternChecksumPatternChecksum().msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternChecksumPatternChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternChecksumPatternChecksum
func (obj *checksumPattern) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternChecksumPatternChecksum value in the ChecksumPattern object
func (obj *checksumPattern) SetChecksum(value PatternChecksumPatternChecksum) ChecksumPattern {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.msg()

	return obj
}

func (obj *checksumPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *checksumPattern) setDefault() {

}
