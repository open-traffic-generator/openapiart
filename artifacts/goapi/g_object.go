package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** GObject *****
type gObject struct {
	validation
	obj          *openapi.GObject
	marshaller   marshalGObject
	unMarshaller unMarshalGObject
}

func NewGObject() GObject {
	obj := gObject{obj: &openapi.GObject{}}
	obj.setDefault()
	return &obj
}

func (obj *gObject) msg() *openapi.GObject {
	return obj.obj
}

func (obj *gObject) setMsg(msg *openapi.GObject) GObject {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalgObject struct {
	obj *gObject
}

type marshalGObject interface {
	// ToProto marshals GObject to protobuf object *openapi.GObject
	ToProto() (*openapi.GObject, error)
	// ToPbText marshals GObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals GObject to JSON text
	ToJson() (string, error)
}

type unMarshalgObject struct {
	obj *gObject
}

type unMarshalGObject interface {
	// FromProto unmarshals GObject from protobuf object *openapi.GObject
	FromProto(msg *openapi.GObject) (GObject, error)
	// FromPbText unmarshals GObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GObject from JSON text
	FromJson(value string) error
}

func (obj *gObject) Marshal() marshalGObject {
	if obj.marshaller == nil {
		obj.marshaller = &marshalgObject{obj: obj}
	}
	return obj.marshaller
}

func (obj *gObject) Unmarshal() unMarshalGObject {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalgObject{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalgObject) ToProto() (*openapi.GObject, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalgObject) FromProto(msg *openapi.GObject) (GObject, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalgObject) ToPbText() (string, error) {
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

func (m *unMarshalgObject) FromPbText(value string) error {
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

func (m *marshalgObject) ToYaml() (string, error) {
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

func (m *unMarshalgObject) FromYaml(value string) error {
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

func (m *marshalgObject) ToJson() (string, error) {
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

func (m *unMarshalgObject) FromJson(value string) error {
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

func (obj *gObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *gObject) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *gObject) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *gObject) Clone() (GObject, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGObject()
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

// GObject is deprecated: new schema Jobject to be used
//
// Description TBD
type GObject interface {
	Validation
	// msg marshals GObject to protobuf object *openapi.GObject
	// and doesn't set defaults
	msg() *openapi.GObject
	// setMsg unmarshals GObject from protobuf object *openapi.GObject
	// and doesn't set defaults
	setMsg(*openapi.GObject) GObject
	// provides marshal interface
	Marshal() marshalGObject
	// provides unmarshal interface
	Unmarshal() unMarshalGObject
	// validate validates GObject
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// GA returns string, set in GObject.
	GA() string
	// SetGA assigns string provided by user to GObject
	SetGA(value string) GObject
	// HasGA checks if GA has been set in GObject
	HasGA() bool
	// GB returns int32, set in GObject.
	GB() int32
	// SetGB assigns int32 provided by user to GObject
	SetGB(value int32) GObject
	// HasGB checks if GB has been set in GObject
	HasGB() bool
	// GC returns float32, set in GObject.
	GC() float32
	// SetGC assigns float32 provided by user to GObject
	SetGC(value float32) GObject
	// HasGC checks if GC has been set in GObject
	HasGC() bool
	// Choice returns GObjectChoiceEnum, set in GObject
	Choice() GObjectChoiceEnum
	// setChoice assigns GObjectChoiceEnum provided by user to GObject
	setChoice(value GObjectChoiceEnum) GObject
	// HasChoice checks if Choice has been set in GObject
	HasChoice() bool
	// GD returns string, set in GObject.
	GD() string
	// SetGD assigns string provided by user to GObject
	SetGD(value string) GObject
	// HasGD checks if GD has been set in GObject
	HasGD() bool
	// GE returns float64, set in GObject.
	GE() float64
	// SetGE assigns float64 provided by user to GObject
	SetGE(value float64) GObject
	// HasGE checks if GE has been set in GObject
	HasGE() bool
	// GF returns GObjectGFEnum, set in GObject
	GF() GObjectGFEnum
	// SetGF assigns GObjectGFEnum provided by user to GObject
	SetGF(value GObjectGFEnum) GObject
	// HasGF checks if GF has been set in GObject
	HasGF() bool
	// Name returns string, set in GObject.
	Name() string
	// SetName assigns string provided by user to GObject
	SetName(value string) GObject
	// HasName checks if Name has been set in GObject
	HasName() bool
}

// description is TBD
// GA returns a string
func (obj *gObject) GA() string {

	return *obj.obj.GA

}

// description is TBD
// GA returns a string
func (obj *gObject) HasGA() bool {
	return obj.obj.GA != nil
}

// description is TBD
// SetGA sets the string value in the GObject object
func (obj *gObject) SetGA(value string) GObject {

	obj.obj.GA = &value
	return obj
}

// description is TBD
// GB returns a int32
func (obj *gObject) GB() int32 {

	return *obj.obj.GB

}

// description is TBD
// GB returns a int32
func (obj *gObject) HasGB() bool {
	return obj.obj.GB != nil
}

// description is TBD
// SetGB sets the int32 value in the GObject object
func (obj *gObject) SetGB(value int32) GObject {

	obj.obj.GB = &value
	return obj
}

// Deprecated: Information TBD
//
// Description TBD
// GC returns a float32
func (obj *gObject) GC() float32 {

	return *obj.obj.GC

}

// Deprecated: Information TBD
//
// Description TBD
// GC returns a float32
func (obj *gObject) HasGC() bool {
	return obj.obj.GC != nil
}

// Deprecated: Information TBD
//
// Description TBD
// SetGC sets the float32 value in the GObject object
func (obj *gObject) SetGC(value float32) GObject {

	obj.obj.GC = &value
	return obj
}

type GObjectChoiceEnum string

// Enum of Choice on GObject
var GObjectChoice = struct {
	G_D GObjectChoiceEnum
	G_E GObjectChoiceEnum
}{
	G_D: GObjectChoiceEnum("g_d"),
	G_E: GObjectChoiceEnum("g_e"),
}

func (obj *gObject) Choice() GObjectChoiceEnum {
	return GObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *gObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *gObject) setChoice(value GObjectChoiceEnum) GObject {
	intValue, ok := openapi.GObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on GObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.GObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.GE = nil
	obj.obj.GD = nil

	if value == GObjectChoice.G_D {
		defaultValue := "some string"
		obj.obj.GD = &defaultValue
	}

	if value == GObjectChoice.G_E {
		defaultValue := float64(3.0)
		obj.obj.GE = &defaultValue
	}

	return obj
}

// description is TBD
// GD returns a string
func (obj *gObject) GD() string {

	if obj.obj.GD == nil {
		obj.setChoice(GObjectChoice.G_D)
	}

	return *obj.obj.GD

}

// description is TBD
// GD returns a string
func (obj *gObject) HasGD() bool {
	return obj.obj.GD != nil
}

// description is TBD
// SetGD sets the string value in the GObject object
func (obj *gObject) SetGD(value string) GObject {
	obj.setChoice(GObjectChoice.G_D)
	obj.obj.GD = &value
	return obj
}

// description is TBD
// GE returns a float64
func (obj *gObject) GE() float64 {

	if obj.obj.GE == nil {
		obj.setChoice(GObjectChoice.G_E)
	}

	return *obj.obj.GE

}

// description is TBD
// GE returns a float64
func (obj *gObject) HasGE() bool {
	return obj.obj.GE != nil
}

// description is TBD
// SetGE sets the float64 value in the GObject object
func (obj *gObject) SetGE(value float64) GObject {
	obj.setChoice(GObjectChoice.G_E)
	obj.obj.GE = &value
	return obj
}

type GObjectGFEnum string

// Enum of GF on GObject
var GObjectGF = struct {
	A GObjectGFEnum
	B GObjectGFEnum
	C GObjectGFEnum
}{
	A: GObjectGFEnum("a"),
	B: GObjectGFEnum("b"),
	C: GObjectGFEnum("c"),
}

func (obj *gObject) GF() GObjectGFEnum {
	return GObjectGFEnum(obj.obj.GF.Enum().String())
}

// Another enum to test protbuf enum generation
// GF returns a string
func (obj *gObject) HasGF() bool {
	return obj.obj.GF != nil
}

func (obj *gObject) SetGF(value GObjectGFEnum) GObject {
	intValue, ok := openapi.GObject_GF_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on GObjectGFEnum", string(value)))
		return obj
	}
	enumValue := openapi.GObject_GF_Enum(intValue)
	obj.obj.GF = &enumValue

	return obj
}

// description is TBD
// Name returns a string
func (obj *gObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *gObject) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the GObject object
func (obj *gObject) SetName(value string) GObject {

	obj.obj.Name = &value
	return obj
}

func (obj *gObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("GObject is deprecated, new schema Jobject to be used")

	// GC is deprecated
	if obj.obj.GC != nil {
		obj.addWarnings("GC property in schema GObject is deprecated, Information TBD")
	}

}

func (obj *gObject) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(GObjectChoice.G_D)

	}
	if obj.obj.GA == nil {
		obj.SetGA("asdf")
	}
	if obj.obj.GB == nil {
		obj.SetGB(6)
	}
	if obj.obj.GC == nil {
		obj.SetGC(77.7)
	}
	if obj.obj.GF == nil {
		obj.SetGF(GObjectGF.A)

	}

}
