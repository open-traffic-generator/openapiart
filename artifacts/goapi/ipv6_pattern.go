package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ipv6Pattern *****
type ipv6Pattern struct {
	validation
	obj          *openapi.Ipv6Pattern
	marshaller   marshalIpv6Pattern
	unMarshaller unMarshalIpv6Pattern
	ipv6Holder   PatternIpv6PatternIpv6
}

func NewIpv6Pattern() Ipv6Pattern {
	obj := ipv6Pattern{obj: &openapi.Ipv6Pattern{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv6Pattern) msg() *openapi.Ipv6Pattern {
	return obj.obj
}

func (obj *ipv6Pattern) setMsg(msg *openapi.Ipv6Pattern) Ipv6Pattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalipv6Pattern struct {
	obj *ipv6Pattern
}

type marshalIpv6Pattern interface {
	// ToProto marshals Ipv6Pattern to protobuf object *openapi.Ipv6Pattern
	ToProto() (*openapi.Ipv6Pattern, error)
	// ToPbText marshals Ipv6Pattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv6Pattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv6Pattern to JSON text
	ToJson() (string, error)
}

type unMarshalipv6Pattern struct {
	obj *ipv6Pattern
}

type unMarshalIpv6Pattern interface {
	// FromProto unmarshals Ipv6Pattern from protobuf object *openapi.Ipv6Pattern
	FromProto(msg *openapi.Ipv6Pattern) (Ipv6Pattern, error)
	// FromPbText unmarshals Ipv6Pattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv6Pattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv6Pattern from JSON text
	FromJson(value string) error
}

func (obj *ipv6Pattern) Marshal() marshalIpv6Pattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalipv6Pattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *ipv6Pattern) Unmarshal() unMarshalIpv6Pattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalipv6Pattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalipv6Pattern) ToProto() (*openapi.Ipv6Pattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalipv6Pattern) FromProto(msg *openapi.Ipv6Pattern) (Ipv6Pattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalipv6Pattern) ToPbText() (string, error) {
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

func (m *unMarshalipv6Pattern) FromPbText(value string) error {
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

func (m *marshalipv6Pattern) ToYaml() (string, error) {
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

func (m *unMarshalipv6Pattern) FromYaml(value string) error {
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

func (m *marshalipv6Pattern) ToJson() (string, error) {
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

func (m *unMarshalipv6Pattern) FromJson(value string) error {
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

func (obj *ipv6Pattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ipv6Pattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv6Pattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv6Pattern) Clone() (Ipv6Pattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv6Pattern()
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

func (obj *ipv6Pattern) setNil() {
	obj.ipv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ipv6Pattern is test ipv6 pattern
type Ipv6Pattern interface {
	Validation
	// msg marshals Ipv6Pattern to protobuf object *openapi.Ipv6Pattern
	// and doesn't set defaults
	msg() *openapi.Ipv6Pattern
	// setMsg unmarshals Ipv6Pattern from protobuf object *openapi.Ipv6Pattern
	// and doesn't set defaults
	setMsg(*openapi.Ipv6Pattern) Ipv6Pattern
	// provides marshal interface
	Marshal() marshalIpv6Pattern
	// provides unmarshal interface
	Unmarshal() unMarshalIpv6Pattern
	// validate validates Ipv6Pattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ipv6Pattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6 returns PatternIpv6PatternIpv6, set in Ipv6Pattern.
	// PatternIpv6PatternIpv6 is tBD
	Ipv6() PatternIpv6PatternIpv6
	// SetIpv6 assigns PatternIpv6PatternIpv6 provided by user to Ipv6Pattern.
	// PatternIpv6PatternIpv6 is tBD
	SetIpv6(value PatternIpv6PatternIpv6) Ipv6Pattern
	// HasIpv6 checks if Ipv6 has been set in Ipv6Pattern
	HasIpv6() bool
	setNil()
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternIpv6
func (obj *ipv6Pattern) Ipv6() PatternIpv6PatternIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = NewPatternIpv6PatternIpv6().msg()
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &patternIpv6PatternIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternIpv6
func (obj *ipv6Pattern) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the PatternIpv6PatternIpv6 value in the Ipv6Pattern object
func (obj *ipv6Pattern) SetIpv6(value PatternIpv6PatternIpv6) Ipv6Pattern {

	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

func (obj *ipv6Pattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

}

func (obj *ipv6Pattern) setDefault() {

}
