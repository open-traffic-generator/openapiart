package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** KObject *****
type kObject struct {
	validation
	obj           *openapi.KObject
	marshaller    marshalKObject
	unMarshaller  unMarshalKObject
	eObjectHolder EObject
	fObjectHolder FObject
}

func NewKObject() KObject {
	obj := kObject{obj: &openapi.KObject{}}
	obj.setDefault()
	return &obj
}

func (obj *kObject) msg() *openapi.KObject {
	return obj.obj
}

func (obj *kObject) setMsg(msg *openapi.KObject) KObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalkObject struct {
	obj *kObject
}

type marshalKObject interface {
	// ToProto marshals KObject to protobuf object *openapi.KObject
	ToProto() (*openapi.KObject, error)
	// ToPbText marshals KObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals KObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals KObject to JSON text
	ToJson() (string, error)
}

type unMarshalkObject struct {
	obj *kObject
}

type unMarshalKObject interface {
	// FromProto unmarshals KObject from protobuf object *openapi.KObject
	FromProto(msg *openapi.KObject) (KObject, error)
	// FromPbText unmarshals KObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals KObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals KObject from JSON text
	FromJson(value string) error
}

func (obj *kObject) Marshal() marshalKObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalkObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *kObject) Unmarshal() unMarshalKObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalkObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalkObject) ToProto() (*openapi.KObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalkObject) FromProto(msg *openapi.KObject) (KObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalkObject) ToPbText() (string, error) {
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

func (m *unMarshalkObject) FromPbText(value string) error {
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

func (m *marshalkObject) ToYaml() (string, error) {
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

func (m *unMarshalkObject) FromYaml(value string) error {
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

func (m *marshalkObject) ToJson() (string, error) {
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

func (m *unMarshalkObject) FromJson(value string) error {
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

func (obj *kObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *kObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *kObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *kObject) Clone() (KObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewKObject()
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

func (obj *kObject) setNil() {
	obj.eObjectHolder = nil
	obj.fObjectHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// KObject is description is TBD
type KObject interface {
	Validation
	// msg marshals KObject to protobuf object *openapi.KObject
	// and doesn't set defaults
	msg() *openapi.KObject
	// setMsg unmarshals KObject from protobuf object *openapi.KObject
	// and doesn't set defaults
	setMsg(*openapi.KObject) KObject
	// provides marshal interface
	Marshal() marshalKObject
	// provides unmarshal interface
	Unmarshal() unMarshalKObject
	// validate validates KObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (KObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EObject returns EObject, set in KObject.
	// EObject is description is TBD
	EObject() EObject
	// SetEObject assigns EObject provided by user to KObject.
	// EObject is description is TBD
	SetEObject(value EObject) KObject
	// HasEObject checks if EObject has been set in KObject
	HasEObject() bool
	// FObject returns FObject, set in KObject.
	// FObject is description is TBD
	FObject() FObject
	// SetFObject assigns FObject provided by user to KObject.
	// FObject is description is TBD
	SetFObject(value FObject) KObject
	// HasFObject checks if FObject has been set in KObject
	HasFObject() bool
	setNil()
}

// description is TBD
// EObject returns a EObject
func (obj *kObject) EObject() EObject {
	if obj.obj.EObject == nil {
		obj.obj.EObject = NewEObject().msg()
	}
	if obj.eObjectHolder == nil {
		obj.eObjectHolder = &eObject{obj: obj.obj.EObject}
	}
	return obj.eObjectHolder
}

// description is TBD
// EObject returns a EObject
func (obj *kObject) HasEObject() bool {
	return obj.obj.EObject != nil
}

// description is TBD
// SetEObject sets the EObject value in the KObject object
func (obj *kObject) SetEObject(value EObject) KObject {

	obj.eObjectHolder = nil
	obj.obj.EObject = value.msg()

	return obj
}

// description is TBD
// FObject returns a FObject
func (obj *kObject) FObject() FObject {
	if obj.obj.FObject == nil {
		obj.obj.FObject = NewFObject().msg()
	}
	if obj.fObjectHolder == nil {
		obj.fObjectHolder = &fObject{obj: obj.obj.FObject}
	}
	return obj.fObjectHolder
}

// description is TBD
// FObject returns a FObject
func (obj *kObject) HasFObject() bool {
	return obj.obj.FObject != nil
}

// description is TBD
// SetFObject sets the FObject value in the KObject object
func (obj *kObject) SetFObject(value FObject) KObject {

	obj.fObjectHolder = nil
	obj.obj.FObject = value.msg()

	return obj
}

func (obj *kObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EObject != nil {

		obj.EObject().validateObj(vObj, set_default)
	}

	if obj.obj.FObject != nil {

		obj.FObject().validateObj(vObj, set_default)
	}

}

func (obj *kObject) setDefault() {

}
