package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ipv4Pattern *****
type ipv4Pattern struct {
	validation
	obj          *openapi.Ipv4Pattern
	marshaller   marshalIpv4Pattern
	unMarshaller unMarshalIpv4Pattern
	ipv4Holder   PatternIpv4PatternIpv4
}

func NewIpv4Pattern() Ipv4Pattern {
	obj := ipv4Pattern{obj: &openapi.Ipv4Pattern{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv4Pattern) msg() *openapi.Ipv4Pattern {
	return obj.obj
}

func (obj *ipv4Pattern) setMsg(msg *openapi.Ipv4Pattern) Ipv4Pattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalipv4Pattern struct {
	obj *ipv4Pattern
}

type marshalIpv4Pattern interface {
	// ToProto marshals Ipv4Pattern to protobuf object *openapi.Ipv4Pattern
	ToProto() (*openapi.Ipv4Pattern, error)
	// ToPbText marshals Ipv4Pattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv4Pattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv4Pattern to JSON text
	ToJson() (string, error)
}

type unMarshalipv4Pattern struct {
	obj *ipv4Pattern
}

type unMarshalIpv4Pattern interface {
	// FromProto unmarshals Ipv4Pattern from protobuf object *openapi.Ipv4Pattern
	FromProto(msg *openapi.Ipv4Pattern) (Ipv4Pattern, error)
	// FromPbText unmarshals Ipv4Pattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv4Pattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv4Pattern from JSON text
	FromJson(value string) error
}

func (obj *ipv4Pattern) Marshal() marshalIpv4Pattern {
	if obj.marshaller == nil {
		obj.marshaller = &marshalipv4Pattern{obj: obj}
	}
	return obj.marshaller
}

func (obj *ipv4Pattern) Unmarshal() unMarshalIpv4Pattern {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalipv4Pattern{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalipv4Pattern) ToProto() (*openapi.Ipv4Pattern, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalipv4Pattern) FromProto(msg *openapi.Ipv4Pattern) (Ipv4Pattern, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalipv4Pattern) ToPbText() (string, error) {
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

func (m *unMarshalipv4Pattern) FromPbText(value string) error {
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

func (m *marshalipv4Pattern) ToYaml() (string, error) {
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

func (m *unMarshalipv4Pattern) FromYaml(value string) error {
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

func (m *marshalipv4Pattern) ToJson() (string, error) {
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

func (m *unMarshalipv4Pattern) FromJson(value string) error {
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

func (obj *ipv4Pattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ipv4Pattern) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv4Pattern) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv4Pattern) Clone() (Ipv4Pattern, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv4Pattern()
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

func (obj *ipv4Pattern) setNil() {
	obj.ipv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ipv4Pattern is test ipv4 pattern
type Ipv4Pattern interface {
	Validation
	// msg marshals Ipv4Pattern to protobuf object *openapi.Ipv4Pattern
	// and doesn't set defaults
	msg() *openapi.Ipv4Pattern
	// setMsg unmarshals Ipv4Pattern from protobuf object *openapi.Ipv4Pattern
	// and doesn't set defaults
	setMsg(*openapi.Ipv4Pattern) Ipv4Pattern
	// provides marshal interface
	Marshal() marshalIpv4Pattern
	// provides unmarshal interface
	Unmarshal() unMarshalIpv4Pattern
	// validate validates Ipv4Pattern
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ipv4Pattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4 returns PatternIpv4PatternIpv4, set in Ipv4Pattern.
	// PatternIpv4PatternIpv4 is tBD
	Ipv4() PatternIpv4PatternIpv4
	// SetIpv4 assigns PatternIpv4PatternIpv4 provided by user to Ipv4Pattern.
	// PatternIpv4PatternIpv4 is tBD
	SetIpv4(value PatternIpv4PatternIpv4) Ipv4Pattern
	// HasIpv4 checks if Ipv4 has been set in Ipv4Pattern
	HasIpv4() bool
	setNil()
}

// description is TBD
// Ipv4 returns a PatternIpv4PatternIpv4
func (obj *ipv4Pattern) Ipv4() PatternIpv4PatternIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = NewPatternIpv4PatternIpv4().msg()
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &patternIpv4PatternIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a PatternIpv4PatternIpv4
func (obj *ipv4Pattern) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the PatternIpv4PatternIpv4 value in the Ipv4Pattern object
func (obj *ipv4Pattern) SetIpv4(value PatternIpv4PatternIpv4) Ipv4Pattern {

	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

func (obj *ipv4Pattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

}

func (obj *ipv4Pattern) setDefault() {

}
