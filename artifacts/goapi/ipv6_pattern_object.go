package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ipv6PatternObject *****
type ipv6PatternObject struct {
	validation
	obj          *openapi.Ipv6PatternObject
	marshaller   marshalIpv6PatternObject
	unMarshaller unMarshalIpv6PatternObject
	ipv6Holder   PatternIpv6PatternObjectIpv6
}

func NewIpv6PatternObject() Ipv6PatternObject {
	obj := ipv6PatternObject{obj: &openapi.Ipv6PatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv6PatternObject) msg() *openapi.Ipv6PatternObject {
	return obj.obj
}

func (obj *ipv6PatternObject) setMsg(msg *openapi.Ipv6PatternObject) Ipv6PatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalipv6PatternObject struct {
	obj *ipv6PatternObject
}

type marshalIpv6PatternObject interface {
	// ToProto marshals Ipv6PatternObject to protobuf object *openapi.Ipv6PatternObject
	ToProto() (*openapi.Ipv6PatternObject, error)
	// ToPbText marshals Ipv6PatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv6PatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv6PatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalipv6PatternObject struct {
	obj *ipv6PatternObject
}

type unMarshalIpv6PatternObject interface {
	// FromProto unmarshals Ipv6PatternObject from protobuf object *openapi.Ipv6PatternObject
	FromProto(msg *openapi.Ipv6PatternObject) (Ipv6PatternObject, error)
	// FromPbText unmarshals Ipv6PatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv6PatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv6PatternObject from JSON text
	FromJson(value string) error
}

func (obj *ipv6PatternObject) Marshal() marshalIpv6PatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalipv6PatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *ipv6PatternObject) Unmarshal() unMarshalIpv6PatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalipv6PatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalipv6PatternObject) ToProto() (*openapi.Ipv6PatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalipv6PatternObject) FromProto(msg *openapi.Ipv6PatternObject) (Ipv6PatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalipv6PatternObject) ToPbText() (string, error) {
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

func (m *unMarshalipv6PatternObject) FromPbText(value string) error {
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

func (m *marshalipv6PatternObject) ToYaml() (string, error) {
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

func (m *unMarshalipv6PatternObject) FromYaml(value string) error {
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

func (m *marshalipv6PatternObject) ToJson() (string, error) {
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

func (m *unMarshalipv6PatternObject) FromJson(value string) error {
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

func (obj *ipv6PatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ipv6PatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv6PatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv6PatternObject) Clone() (Ipv6PatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv6PatternObject()
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

func (obj *ipv6PatternObject) setNil() {
	obj.ipv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ipv6PatternObject is test ipv6 pattern
type Ipv6PatternObject interface {
	Validation
	// msg marshals Ipv6PatternObject to protobuf object *openapi.Ipv6PatternObject
	// and doesn't set defaults
	msg() *openapi.Ipv6PatternObject
	// setMsg unmarshals Ipv6PatternObject from protobuf object *openapi.Ipv6PatternObject
	// and doesn't set defaults
	setMsg(*openapi.Ipv6PatternObject) Ipv6PatternObject
	// provides marshal interface
	Marshal() marshalIpv6PatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalIpv6PatternObject
	// validate validates Ipv6PatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ipv6PatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6 returns PatternIpv6PatternObjectIpv6, set in Ipv6PatternObject.
	// PatternIpv6PatternObjectIpv6 is tBD
	Ipv6() PatternIpv6PatternObjectIpv6
	// SetIpv6 assigns PatternIpv6PatternObjectIpv6 provided by user to Ipv6PatternObject.
	// PatternIpv6PatternObjectIpv6 is tBD
	SetIpv6(value PatternIpv6PatternObjectIpv6) Ipv6PatternObject
	// HasIpv6 checks if Ipv6 has been set in Ipv6PatternObject
	HasIpv6() bool
	setNil()
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternObjectIpv6
func (obj *ipv6PatternObject) Ipv6() PatternIpv6PatternObjectIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = NewPatternIpv6PatternObjectIpv6().msg()
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &patternIpv6PatternObjectIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternObjectIpv6
func (obj *ipv6PatternObject) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the PatternIpv6PatternObjectIpv6 value in the Ipv6PatternObject object
func (obj *ipv6PatternObject) SetIpv6(value PatternIpv6PatternObjectIpv6) Ipv6PatternObject {

	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.msg()

	return obj
}

func (obj *ipv6PatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

}

func (obj *ipv6PatternObject) setDefault() {

}
