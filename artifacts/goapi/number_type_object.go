package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** NumberTypeObject *****
type numberTypeObject struct {
	validation
	obj          *openapi.NumberTypeObject
	marshaller   marshalNumberTypeObject
	unMarshaller unMarshalNumberTypeObject
}

func NewNumberTypeObject() NumberTypeObject {
	obj := numberTypeObject{obj: &openapi.NumberTypeObject{}}
	obj.setDefault()
	return &obj
}

func (obj *numberTypeObject) msg() *openapi.NumberTypeObject {
	return obj.obj
}

func (obj *numberTypeObject) setMsg(msg *openapi.NumberTypeObject) NumberTypeObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalnumberTypeObject struct {
	obj *numberTypeObject
}

type marshalNumberTypeObject interface {
	// ToProto marshals NumberTypeObject to protobuf object *openapi.NumberTypeObject
	ToProto() (*openapi.NumberTypeObject, error)
	// ToPbText marshals NumberTypeObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals NumberTypeObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals NumberTypeObject to JSON text
	ToJson() (string, error)
}

type unMarshalnumberTypeObject struct {
	obj *numberTypeObject
}

type unMarshalNumberTypeObject interface {
	// FromProto unmarshals NumberTypeObject from protobuf object *openapi.NumberTypeObject
	FromProto(msg *openapi.NumberTypeObject) (NumberTypeObject, error)
	// FromPbText unmarshals NumberTypeObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals NumberTypeObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals NumberTypeObject from JSON text
	FromJson(value string) error
}

func (obj *numberTypeObject) Marshal() marshalNumberTypeObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalnumberTypeObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *numberTypeObject) Unmarshal() unMarshalNumberTypeObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalnumberTypeObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalnumberTypeObject) ToProto() (*openapi.NumberTypeObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalnumberTypeObject) FromProto(msg *openapi.NumberTypeObject) (NumberTypeObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalnumberTypeObject) ToPbText() (string, error) {
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

func (m *unMarshalnumberTypeObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalnumberTypeObject) ToYaml() (string, error) {
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

func (m *unMarshalnumberTypeObject) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalnumberTypeObject) ToJson() (string, error) {
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

func (m *unMarshalnumberTypeObject) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *numberTypeObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *numberTypeObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *numberTypeObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *numberTypeObject) Clone() (NumberTypeObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewNumberTypeObject()
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

// NumberTypeObject is description is TBD
type NumberTypeObject interface {
	Validation
	// msg marshals NumberTypeObject to protobuf object *openapi.NumberTypeObject
	// and doesn't set defaults
	msg() *openapi.NumberTypeObject
	// setMsg unmarshals NumberTypeObject from protobuf object *openapi.NumberTypeObject
	// and doesn't set defaults
	setMsg(*openapi.NumberTypeObject) NumberTypeObject
	// provides marshal interface
	Marshal() marshalNumberTypeObject
	// provides unmarshal interface
	Unmarshal() unMarshalNumberTypeObject
	// validate validates NumberTypeObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (NumberTypeObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ValidateUint321 returns uint32, set in NumberTypeObject.
	ValidateUint321() uint32
	// SetValidateUint321 assigns uint32 provided by user to NumberTypeObject
	SetValidateUint321(value uint32) NumberTypeObject
	// HasValidateUint321 checks if ValidateUint321 has been set in NumberTypeObject
	HasValidateUint321() bool
	// ValidateUint322 returns uint32, set in NumberTypeObject.
	ValidateUint322() uint32
	// SetValidateUint322 assigns uint32 provided by user to NumberTypeObject
	SetValidateUint322(value uint32) NumberTypeObject
	// HasValidateUint322 checks if ValidateUint322 has been set in NumberTypeObject
	HasValidateUint322() bool
	// ValidateUint641 returns uint64, set in NumberTypeObject.
	ValidateUint641() uint64
	// SetValidateUint641 assigns uint64 provided by user to NumberTypeObject
	SetValidateUint641(value uint64) NumberTypeObject
	// HasValidateUint641 checks if ValidateUint641 has been set in NumberTypeObject
	HasValidateUint641() bool
	// ValidateUint642 returns uint64, set in NumberTypeObject.
	ValidateUint642() uint64
	// SetValidateUint642 assigns uint64 provided by user to NumberTypeObject
	SetValidateUint642(value uint64) NumberTypeObject
	// HasValidateUint642 checks if ValidateUint642 has been set in NumberTypeObject
	HasValidateUint642() bool
	// ValidateInt321 returns int32, set in NumberTypeObject.
	ValidateInt321() int32
	// SetValidateInt321 assigns int32 provided by user to NumberTypeObject
	SetValidateInt321(value int32) NumberTypeObject
	// HasValidateInt321 checks if ValidateInt321 has been set in NumberTypeObject
	HasValidateInt321() bool
	// ValidateInt322 returns int32, set in NumberTypeObject.
	ValidateInt322() int32
	// SetValidateInt322 assigns int32 provided by user to NumberTypeObject
	SetValidateInt322(value int32) NumberTypeObject
	// HasValidateInt322 checks if ValidateInt322 has been set in NumberTypeObject
	HasValidateInt322() bool
	// ValidateInt641 returns int64, set in NumberTypeObject.
	ValidateInt641() int64
	// SetValidateInt641 assigns int64 provided by user to NumberTypeObject
	SetValidateInt641(value int64) NumberTypeObject
	// HasValidateInt641 checks if ValidateInt641 has been set in NumberTypeObject
	HasValidateInt641() bool
	// ValidateInt642 returns int64, set in NumberTypeObject.
	ValidateInt642() int64
	// SetValidateInt642 assigns int64 provided by user to NumberTypeObject
	SetValidateInt642(value int64) NumberTypeObject
	// HasValidateInt642 checks if ValidateInt642 has been set in NumberTypeObject
	HasValidateInt642() bool
}

// description is TBD
// ValidateUint321 returns a uint32
func (obj *numberTypeObject) ValidateUint321() uint32 {

	return *obj.obj.ValidateUint32_1

}

// description is TBD
// ValidateUint321 returns a uint32
func (obj *numberTypeObject) HasValidateUint321() bool {
	return obj.obj.ValidateUint32_1 != nil
}

// description is TBD
// SetValidateUint321 sets the uint32 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateUint321(value uint32) NumberTypeObject {

	obj.obj.ValidateUint32_1 = &value
	return obj
}

// description is TBD
// ValidateUint322 returns a uint32
func (obj *numberTypeObject) ValidateUint322() uint32 {

	return *obj.obj.ValidateUint32_2

}

// description is TBD
// ValidateUint322 returns a uint32
func (obj *numberTypeObject) HasValidateUint322() bool {
	return obj.obj.ValidateUint32_2 != nil
}

// description is TBD
// SetValidateUint322 sets the uint32 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateUint322(value uint32) NumberTypeObject {

	obj.obj.ValidateUint32_2 = &value
	return obj
}

// description is TBD
// ValidateUint641 returns a uint64
func (obj *numberTypeObject) ValidateUint641() uint64 {

	return *obj.obj.ValidateUint64_1

}

// description is TBD
// ValidateUint641 returns a uint64
func (obj *numberTypeObject) HasValidateUint641() bool {
	return obj.obj.ValidateUint64_1 != nil
}

// description is TBD
// SetValidateUint641 sets the uint64 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateUint641(value uint64) NumberTypeObject {

	obj.obj.ValidateUint64_1 = &value
	return obj
}

// description is TBD
// ValidateUint642 returns a uint64
func (obj *numberTypeObject) ValidateUint642() uint64 {

	return *obj.obj.ValidateUint64_2

}

// description is TBD
// ValidateUint642 returns a uint64
func (obj *numberTypeObject) HasValidateUint642() bool {
	return obj.obj.ValidateUint64_2 != nil
}

// description is TBD
// SetValidateUint642 sets the uint64 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateUint642(value uint64) NumberTypeObject {

	obj.obj.ValidateUint64_2 = &value
	return obj
}

// description is TBD
// ValidateInt321 returns a int32
func (obj *numberTypeObject) ValidateInt321() int32 {

	return *obj.obj.ValidateInt32_1

}

// description is TBD
// ValidateInt321 returns a int32
func (obj *numberTypeObject) HasValidateInt321() bool {
	return obj.obj.ValidateInt32_1 != nil
}

// description is TBD
// SetValidateInt321 sets the int32 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateInt321(value int32) NumberTypeObject {

	obj.obj.ValidateInt32_1 = &value
	return obj
}

// description is TBD
// ValidateInt322 returns a int32
func (obj *numberTypeObject) ValidateInt322() int32 {

	return *obj.obj.ValidateInt32_2

}

// description is TBD
// ValidateInt322 returns a int32
func (obj *numberTypeObject) HasValidateInt322() bool {
	return obj.obj.ValidateInt32_2 != nil
}

// description is TBD
// SetValidateInt322 sets the int32 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateInt322(value int32) NumberTypeObject {

	obj.obj.ValidateInt32_2 = &value
	return obj
}

// description is TBD
// ValidateInt641 returns a int64
func (obj *numberTypeObject) ValidateInt641() int64 {

	return *obj.obj.ValidateInt64_1

}

// description is TBD
// ValidateInt641 returns a int64
func (obj *numberTypeObject) HasValidateInt641() bool {
	return obj.obj.ValidateInt64_1 != nil
}

// description is TBD
// SetValidateInt641 sets the int64 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateInt641(value int64) NumberTypeObject {

	obj.obj.ValidateInt64_1 = &value
	return obj
}

// description is TBD
// ValidateInt642 returns a int64
func (obj *numberTypeObject) ValidateInt642() int64 {

	return *obj.obj.ValidateInt64_2

}

// description is TBD
// ValidateInt642 returns a int64
func (obj *numberTypeObject) HasValidateInt642() bool {
	return obj.obj.ValidateInt64_2 != nil
}

// description is TBD
// SetValidateInt642 sets the int64 value in the NumberTypeObject object
func (obj *numberTypeObject) SetValidateInt642(value int64) NumberTypeObject {

	obj.obj.ValidateInt64_2 = &value
	return obj
}

func (obj *numberTypeObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ValidateUint32_1 != nil {

		if *obj.obj.ValidateUint32_1 > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= NumberTypeObject.ValidateUint32_1 <= 4294967295 but Got %d", *obj.obj.ValidateUint32_1))
		}

	}

	if obj.obj.ValidateUint64_1 != nil {

		if *obj.obj.ValidateUint64_1 > 9223372036854775807 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= NumberTypeObject.ValidateUint64_1 <= 9223372036854775807 but Got %d", *obj.obj.ValidateUint64_1))
		}

	}

	if obj.obj.ValidateInt32_1 != nil {

		if *obj.obj.ValidateInt32_1 < 0 || *obj.obj.ValidateInt32_1 > 2147483646 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= NumberTypeObject.ValidateInt32_1 <= 2147483646 but Got %d", *obj.obj.ValidateInt32_1))
		}

	}

	if obj.obj.ValidateInt64_1 != nil {

		if *obj.obj.ValidateInt64_1 < 0 || *obj.obj.ValidateInt64_1 > 9223372036854775807 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= NumberTypeObject.ValidateInt64_1 <= 9223372036854775807 but Got %d", *obj.obj.ValidateInt64_1))
		}

	}

}

func (obj *numberTypeObject) setDefault() {

}
