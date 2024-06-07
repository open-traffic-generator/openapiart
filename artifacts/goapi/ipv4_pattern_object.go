package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ipv4PatternObject *****
type ipv4PatternObject struct {
	validation
	obj          *openapi.Ipv4PatternObject
	marshaller   marshalIpv4PatternObject
	unMarshaller unMarshalIpv4PatternObject
	ipv4Holder   PatternIpv4PatternObjectIpv4
}

func NewIpv4PatternObject() Ipv4PatternObject {
	obj := ipv4PatternObject{obj: &openapi.Ipv4PatternObject{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv4PatternObject) msg() *openapi.Ipv4PatternObject {
	return obj.obj
}

func (obj *ipv4PatternObject) setMsg(msg *openapi.Ipv4PatternObject) Ipv4PatternObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalipv4PatternObject struct {
	obj *ipv4PatternObject
}

type marshalIpv4PatternObject interface {
	// ToProto marshals Ipv4PatternObject to protobuf object *openapi.Ipv4PatternObject
	ToProto() (*openapi.Ipv4PatternObject, error)
	// ToPbText marshals Ipv4PatternObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv4PatternObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv4PatternObject to JSON text
	ToJson() (string, error)
}

type unMarshalipv4PatternObject struct {
	obj *ipv4PatternObject
}

type unMarshalIpv4PatternObject interface {
	// FromProto unmarshals Ipv4PatternObject from protobuf object *openapi.Ipv4PatternObject
	FromProto(msg *openapi.Ipv4PatternObject) (Ipv4PatternObject, error)
	// FromPbText unmarshals Ipv4PatternObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv4PatternObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv4PatternObject from JSON text
	FromJson(value string) error
}

func (obj *ipv4PatternObject) Marshal() marshalIpv4PatternObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalipv4PatternObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *ipv4PatternObject) Unmarshal() unMarshalIpv4PatternObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalipv4PatternObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalipv4PatternObject) ToProto() (*openapi.Ipv4PatternObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalipv4PatternObject) FromProto(msg *openapi.Ipv4PatternObject) (Ipv4PatternObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalipv4PatternObject) ToPbText() (string, error) {
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

func (m *unMarshalipv4PatternObject) FromPbText(value string) error {
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

func (m *marshalipv4PatternObject) ToYaml() (string, error) {
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

func (m *unMarshalipv4PatternObject) FromYaml(value string) error {
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

func (m *marshalipv4PatternObject) ToJson() (string, error) {
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

func (m *unMarshalipv4PatternObject) FromJson(value string) error {
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

func (obj *ipv4PatternObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ipv4PatternObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv4PatternObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv4PatternObject) Clone() (Ipv4PatternObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv4PatternObject()
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

func (obj *ipv4PatternObject) setNil() {
	obj.ipv4Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ipv4PatternObject is test ipv4 pattern
type Ipv4PatternObject interface {
	Validation
	// msg marshals Ipv4PatternObject to protobuf object *openapi.Ipv4PatternObject
	// and doesn't set defaults
	msg() *openapi.Ipv4PatternObject
	// setMsg unmarshals Ipv4PatternObject from protobuf object *openapi.Ipv4PatternObject
	// and doesn't set defaults
	setMsg(*openapi.Ipv4PatternObject) Ipv4PatternObject
	// provides marshal interface
	Marshal() marshalIpv4PatternObject
	// provides unmarshal interface
	Unmarshal() unMarshalIpv4PatternObject
	// validate validates Ipv4PatternObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ipv4PatternObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4 returns PatternIpv4PatternObjectIpv4, set in Ipv4PatternObject.
	// PatternIpv4PatternObjectIpv4 is tBD
	Ipv4() PatternIpv4PatternObjectIpv4
	// SetIpv4 assigns PatternIpv4PatternObjectIpv4 provided by user to Ipv4PatternObject.
	// PatternIpv4PatternObjectIpv4 is tBD
	SetIpv4(value PatternIpv4PatternObjectIpv4) Ipv4PatternObject
	// HasIpv4 checks if Ipv4 has been set in Ipv4PatternObject
	HasIpv4() bool
	setNil()
}

// description is TBD
// Ipv4 returns a PatternIpv4PatternObjectIpv4
func (obj *ipv4PatternObject) Ipv4() PatternIpv4PatternObjectIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = NewPatternIpv4PatternObjectIpv4().msg()
	}
	if obj.ipv4Holder == nil {
		obj.ipv4Holder = &patternIpv4PatternObjectIpv4{obj: obj.obj.Ipv4}
	}
	return obj.ipv4Holder
}

// description is TBD
// Ipv4 returns a PatternIpv4PatternObjectIpv4
func (obj *ipv4PatternObject) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the PatternIpv4PatternObjectIpv4 value in the Ipv4PatternObject object
func (obj *ipv4PatternObject) SetIpv4(value PatternIpv4PatternObjectIpv4) Ipv4PatternObject {

	obj.ipv4Holder = nil
	obj.obj.Ipv4 = value.msg()

	return obj
}

func (obj *ipv4PatternObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4 != nil {

		obj.Ipv4().validateObj(vObj, set_default)
	}

}

func (obj *ipv4PatternObject) setDefault() {

}
