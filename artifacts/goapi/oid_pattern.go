package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** OidPattern *****
type oidPattern struct {
	validation
	obj          *openapi.OidPattern
	marshaller   marshalOidPattern
	unMarshaller unMarshalOidPattern
	oidHolder    PatternOidPatternOid
}

func NewOidPattern() OidPattern {
	obj := oidPattern{obj: &openapi.OidPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *oidPattern) msg() *openapi.OidPattern {
	return obj.obj
}

func (obj *oidPattern) setMsg(msg *openapi.OidPattern) OidPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaloidPattern struct {
	obj *oidPattern
}

type marshalOidPattern interface {
	// ToProto marshals OidPattern to protobuf object *openapi.OidPattern
	ToProto() (*openapi.OidPattern, error)
	// ToPbText marshals OidPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals OidPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals OidPattern to JSON text
	ToJson() (string, error)
}

type unMarshaloidPattern struct {
	obj *oidPattern
}

type unMarshalOidPattern interface {
	// FromProto unmarshals OidPattern from protobuf object *openapi.OidPattern
	FromProto(msg *openapi.OidPattern) (OidPattern, error)
	// FromPbText unmarshals OidPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals OidPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals OidPattern from JSON text
	FromJson(value string) error
}

func (obj *oidPattern) Marshal() marshalOidPattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshaloidPattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *oidPattern) Unmarshal() unMarshalOidPattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaloidPattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaloidPattern) ToProto() (*openapi.OidPattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaloidPattern) FromProto(msg *openapi.OidPattern) (OidPattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaloidPattern) ToPbText() (string, error) {
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

func (m *unMarshaloidPattern) FromPbText(value string) error {
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

func (m *marshaloidPattern) ToYaml() (string, error) {
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

func (m *unMarshaloidPattern) FromYaml(value string) error {
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

func (m *marshaloidPattern) ToJson() (string, error) {
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

func (m *unMarshaloidPattern) FromJson(value string) error {
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

func (obj *oidPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *oidPattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *oidPattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *oidPattern) Clone() (OidPattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOidPattern()
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

func (obj *oidPattern) setNil() {
	obj.oidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// OidPattern is test oid pattern
type OidPattern interface {
	Validation
	// msg marshals OidPattern to protobuf object *openapi.OidPattern
	// and doesn't set defaults
	msg() *openapi.OidPattern
	// setMsg unmarshals OidPattern from protobuf object *openapi.OidPattern
	// and doesn't set defaults
	setMsg(*openapi.OidPattern) OidPattern
	// provides marshal interface
	Marshal() marshalOidPattern
	// provides unmarshal interface
	Unmarshal() unMarshalOidPattern
	// validate validates OidPattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (OidPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Oid returns PatternOidPatternOid, set in OidPattern.
	// PatternOidPatternOid is tBD
	Oid() PatternOidPatternOid
	// SetOid assigns PatternOidPatternOid provided by user to OidPattern.
	// PatternOidPatternOid is tBD
	SetOid(value PatternOidPatternOid) OidPattern
	// HasOid checks if Oid has been set in OidPattern
	HasOid() bool
	setNil()
}

// description is TBD
// Oid returns a PatternOidPatternOid
func (obj *oidPattern) Oid() PatternOidPatternOid {
	if obj.obj.Oid == nil {
		obj.obj.Oid = NewPatternOidPatternOid().msg()
	}
	if obj.oidHolder == nil {
		obj.oidHolder = &patternOidPatternOid{obj: obj.obj.Oid}
	}
	return obj.oidHolder
}

// description is TBD
// Oid returns a PatternOidPatternOid
func (obj *oidPattern) HasOid() bool {
	return obj.obj.Oid != nil
}

// description is TBD
// SetOid sets the PatternOidPatternOid value in the OidPattern object
func (obj *oidPattern) SetOid(value PatternOidPatternOid) OidPattern {

	obj.oidHolder = nil
	obj.obj.Oid = value.msg()

	return obj
}

func (obj *oidPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Oid != nil {

		obj.Oid().validateObj(vObj, set_default)
	}

}

func (obj *oidPattern) setDefault() {

}
