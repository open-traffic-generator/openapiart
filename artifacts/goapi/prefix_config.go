package goapi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** PrefixConfig *****
type prefixConfig struct {
	validation
	obj                         *openapi.PrefixConfig
	marshaller                  marshalPrefixConfig
	unMarshaller                unMarshalPrefixConfig
	requiredObjectHolder        EObject
	optionalObjectHolder        EObject
	eHolder                     EObject
	fHolder                     FObject
	gHolder                     PrefixConfigGObjectIter
	jHolder                     PrefixConfigJObjectIter
	kHolder                     KObject
	lHolder                     LObject
	levelHolder                 LevelOne
	mandatoryHolder             Mandate
	ipv4PatternHolder           Ipv4Pattern
	ipv6PatternHolder           Ipv6Pattern
	macPatternHolder            MacPattern
	integerPatternHolder        IntegerPattern
	checksumPatternHolder       ChecksumPattern
	caseHolder                  Layer1Ieee802X
	mObjectHolder               MObject
	headerChecksumHolder        PatternPrefixConfigHeaderChecksum
	autoFieldTestHolder         PatternPrefixConfigAutoFieldTest
	wListHolder                 PrefixConfigWObjectIter
	xListHolder                 PrefixConfigZObjectIter
	zObjectHolder               ZObject
	yObjectHolder               YObject
	choiceObjectHolder          PrefixConfigChoiceObjectIter
	requiredChoiceObjectHolder  RequiredChoiceParent
	g1Holder                    PrefixConfigGObjectIter
	g2Holder                    PrefixConfigGObjectIter
	choiceTestHolder            ChoiceTestObj
	signedIntegerPatternHolder  SignedIntegerPattern
	oidPatternHolder            OidPattern
	choiceDefaultHolder         ChoiceObject
	choiceRequiredDefaultHolder ChoiceRequiredAndDefault
	autoPatternHolder           AutoPattern
	autoPatternDefaultHolder    AutoPatternDefault
}

func NewPrefixConfig() PrefixConfig {
	obj := prefixConfig{obj: &openapi.PrefixConfig{}}
	obj.setDefault()
	return &obj
}

func (obj *prefixConfig) msg() *openapi.PrefixConfig {
	return obj.obj
}

func (obj *prefixConfig) setMsg(msg *openapi.PrefixConfig) PrefixConfig {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalprefixConfig struct {
	obj *prefixConfig
}

type marshalPrefixConfig interface {
	// ToProto marshals PrefixConfig to protobuf object *openapi.PrefixConfig
	ToProto() (*openapi.PrefixConfig, error)
	// ToPbText marshals PrefixConfig to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PrefixConfig to YAML text
	ToYaml() (string, error)
	// ToJson marshals PrefixConfig to JSON text
	ToJson() (string, error)
}

type unMarshalprefixConfig struct {
	obj *prefixConfig
}

type unMarshalPrefixConfig interface {
	// FromProto unmarshals PrefixConfig from protobuf object *openapi.PrefixConfig
	FromProto(msg *openapi.PrefixConfig) (PrefixConfig, error)
	// FromPbText unmarshals PrefixConfig from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PrefixConfig from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PrefixConfig from JSON text
	FromJson(value string) error
}

func (obj *prefixConfig) Marshal() marshalPrefixConfig {
	if obj.marshaller == nil {
		obj.marshaller = &marshalprefixConfig{obj: obj}
	}
	return obj.marshaller
}

func (obj *prefixConfig) Unmarshal() unMarshalPrefixConfig {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalprefixConfig{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalprefixConfig) ToProto() (*openapi.PrefixConfig, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalprefixConfig) FromProto(msg *openapi.PrefixConfig) (PrefixConfig, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalprefixConfig) ToPbText() (string, error) {
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

func (m *unMarshalprefixConfig) FromPbText(value string) error {
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

func (m *marshalprefixConfig) ToYaml() (string, error) {
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

func (m *unMarshalprefixConfig) FromYaml(value string) error {
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

func (m *marshalprefixConfig) ToJson() (string, error) {
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

func (m *unMarshalprefixConfig) FromJson(value string) error {
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

func (obj *prefixConfig) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *prefixConfig) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *prefixConfig) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *prefixConfig) Clone() (PrefixConfig, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPrefixConfig()
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

func (obj *prefixConfig) setNil() {
	obj.requiredObjectHolder = nil
	obj.optionalObjectHolder = nil
	obj.eHolder = nil
	obj.fHolder = nil
	obj.gHolder = nil
	obj.jHolder = nil
	obj.kHolder = nil
	obj.lHolder = nil
	obj.levelHolder = nil
	obj.mandatoryHolder = nil
	obj.ipv4PatternHolder = nil
	obj.ipv6PatternHolder = nil
	obj.macPatternHolder = nil
	obj.integerPatternHolder = nil
	obj.checksumPatternHolder = nil
	obj.caseHolder = nil
	obj.mObjectHolder = nil
	obj.headerChecksumHolder = nil
	obj.autoFieldTestHolder = nil
	obj.wListHolder = nil
	obj.xListHolder = nil
	obj.zObjectHolder = nil
	obj.yObjectHolder = nil
	obj.choiceObjectHolder = nil
	obj.requiredChoiceObjectHolder = nil
	obj.g1Holder = nil
	obj.g2Holder = nil
	obj.choiceTestHolder = nil
	obj.signedIntegerPatternHolder = nil
	obj.oidPatternHolder = nil
	obj.choiceDefaultHolder = nil
	obj.choiceRequiredDefaultHolder = nil
	obj.autoPatternHolder = nil
	obj.autoPatternDefaultHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PrefixConfig is container which retains the configuration
type PrefixConfig interface {
	Validation
	// msg marshals PrefixConfig to protobuf object *openapi.PrefixConfig
	// and doesn't set defaults
	msg() *openapi.PrefixConfig
	// setMsg unmarshals PrefixConfig from protobuf object *openapi.PrefixConfig
	// and doesn't set defaults
	setMsg(*openapi.PrefixConfig) PrefixConfig
	// provides marshal interface
	Marshal() marshalPrefixConfig
	// provides unmarshal interface
	Unmarshal() unMarshalPrefixConfig
	// validate validates PrefixConfig
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PrefixConfig, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequiredObject returns EObject, set in PrefixConfig.
	// EObject is description is TBD
	RequiredObject() EObject
	// SetRequiredObject assigns EObject provided by user to PrefixConfig.
	// EObject is description is TBD
	SetRequiredObject(value EObject) PrefixConfig
	// OptionalObject returns EObject, set in PrefixConfig.
	// EObject is description is TBD
	OptionalObject() EObject
	// SetOptionalObject assigns EObject provided by user to PrefixConfig.
	// EObject is description is TBD
	SetOptionalObject(value EObject) PrefixConfig
	// HasOptionalObject checks if OptionalObject has been set in PrefixConfig
	HasOptionalObject() bool
	// Ieee8021Qbb returns bool, set in PrefixConfig.
	Ieee8021Qbb() bool
	// SetIeee8021Qbb assigns bool provided by user to PrefixConfig
	SetIeee8021Qbb(value bool) PrefixConfig
	// HasIeee8021Qbb checks if Ieee8021Qbb has been set in PrefixConfig
	HasIeee8021Qbb() bool
	// Space1 returns int32, set in PrefixConfig.
	Space1() int32
	// SetSpace1 assigns int32 provided by user to PrefixConfig
	SetSpace1(value int32) PrefixConfig
	// HasSpace1 checks if Space1 has been set in PrefixConfig
	HasSpace1() bool
	// FullDuplex100Mb returns int64, set in PrefixConfig.
	FullDuplex100Mb() int64
	// SetFullDuplex100Mb assigns int64 provided by user to PrefixConfig
	SetFullDuplex100Mb(value int64) PrefixConfig
	// HasFullDuplex100Mb checks if FullDuplex100Mb has been set in PrefixConfig
	HasFullDuplex100Mb() bool
	// Response returns PrefixConfigResponseEnum, set in PrefixConfig
	Response() PrefixConfigResponseEnum
	// SetResponse assigns PrefixConfigResponseEnum provided by user to PrefixConfig
	SetResponse(value PrefixConfigResponseEnum) PrefixConfig
	// HasResponse checks if Response has been set in PrefixConfig
	HasResponse() bool
	// A returns string, set in PrefixConfig.
	A() string
	// SetA assigns string provided by user to PrefixConfig
	SetA(value string) PrefixConfig
	// B returns float32, set in PrefixConfig.
	B() float32
	// SetB assigns float32 provided by user to PrefixConfig
	SetB(value float32) PrefixConfig
	// C returns int32, set in PrefixConfig.
	C() int32
	// SetC assigns int32 provided by user to PrefixConfig
	SetC(value int32) PrefixConfig
	// DValues returns []PrefixConfigDValuesEnum, set in PrefixConfig
	DValues() []PrefixConfigDValuesEnum
	// SetDValues assigns []PrefixConfigDValuesEnum provided by user to PrefixConfig
	SetDValues(value []PrefixConfigDValuesEnum) PrefixConfig
	// E returns EObject, set in PrefixConfig.
	// EObject is description is TBD
	E() EObject
	// SetE assigns EObject provided by user to PrefixConfig.
	// EObject is description is TBD
	SetE(value EObject) PrefixConfig
	// HasE checks if E has been set in PrefixConfig
	HasE() bool
	// F returns FObject, set in PrefixConfig.
	// FObject is description is TBD
	F() FObject
	// SetF assigns FObject provided by user to PrefixConfig.
	// FObject is description is TBD
	SetF(value FObject) PrefixConfig
	// HasF checks if F has been set in PrefixConfig
	HasF() bool
	// G returns PrefixConfigGObjectIterIter, set in PrefixConfig
	G() PrefixConfigGObjectIter
	// H returns bool, set in PrefixConfig.
	H() bool
	// SetH assigns bool provided by user to PrefixConfig
	SetH(value bool) PrefixConfig
	// HasH checks if H has been set in PrefixConfig
	HasH() bool
	// I returns []byte, set in PrefixConfig.
	I() []byte
	// SetI assigns []byte provided by user to PrefixConfig
	SetI(value []byte) PrefixConfig
	// J returns PrefixConfigJObjectIterIter, set in PrefixConfig
	J() PrefixConfigJObjectIter
	// K returns KObject, set in PrefixConfig.
	// KObject is description is TBD
	K() KObject
	// SetK assigns KObject provided by user to PrefixConfig.
	// KObject is description is TBD
	SetK(value KObject) PrefixConfig
	// HasK checks if K has been set in PrefixConfig
	HasK() bool
	// L returns LObject, set in PrefixConfig.
	// LObject is format validation object
	L() LObject
	// SetL assigns LObject provided by user to PrefixConfig.
	// LObject is format validation object
	SetL(value LObject) PrefixConfig
	// HasL checks if L has been set in PrefixConfig
	HasL() bool
	// ListOfStringValues returns []string, set in PrefixConfig.
	ListOfStringValues() []string
	// SetListOfStringValues assigns []string provided by user to PrefixConfig
	SetListOfStringValues(value []string) PrefixConfig
	// ListOfIntegerValues returns []int32, set in PrefixConfig.
	ListOfIntegerValues() []int32
	// SetListOfIntegerValues assigns []int32 provided by user to PrefixConfig
	SetListOfIntegerValues(value []int32) PrefixConfig
	// Level returns LevelOne, set in PrefixConfig.
	// LevelOne is to Test Multi level non-primitive types
	Level() LevelOne
	// SetLevel assigns LevelOne provided by user to PrefixConfig.
	// LevelOne is to Test Multi level non-primitive types
	SetLevel(value LevelOne) PrefixConfig
	// HasLevel checks if Level has been set in PrefixConfig
	HasLevel() bool
	// Mandatory returns Mandate, set in PrefixConfig.
	// Mandate is object to Test required Parameter
	Mandatory() Mandate
	// SetMandatory assigns Mandate provided by user to PrefixConfig.
	// Mandate is object to Test required Parameter
	SetMandatory(value Mandate) PrefixConfig
	// HasMandatory checks if Mandatory has been set in PrefixConfig
	HasMandatory() bool
	// Ipv4Pattern returns Ipv4Pattern, set in PrefixConfig.
	// Ipv4Pattern is test ipv4 pattern
	Ipv4Pattern() Ipv4Pattern
	// SetIpv4Pattern assigns Ipv4Pattern provided by user to PrefixConfig.
	// Ipv4Pattern is test ipv4 pattern
	SetIpv4Pattern(value Ipv4Pattern) PrefixConfig
	// HasIpv4Pattern checks if Ipv4Pattern has been set in PrefixConfig
	HasIpv4Pattern() bool
	// Ipv6Pattern returns Ipv6Pattern, set in PrefixConfig.
	// Ipv6Pattern is test ipv6 pattern
	Ipv6Pattern() Ipv6Pattern
	// SetIpv6Pattern assigns Ipv6Pattern provided by user to PrefixConfig.
	// Ipv6Pattern is test ipv6 pattern
	SetIpv6Pattern(value Ipv6Pattern) PrefixConfig
	// HasIpv6Pattern checks if Ipv6Pattern has been set in PrefixConfig
	HasIpv6Pattern() bool
	// MacPattern returns MacPattern, set in PrefixConfig.
	// MacPattern is test mac pattern
	MacPattern() MacPattern
	// SetMacPattern assigns MacPattern provided by user to PrefixConfig.
	// MacPattern is test mac pattern
	SetMacPattern(value MacPattern) PrefixConfig
	// HasMacPattern checks if MacPattern has been set in PrefixConfig
	HasMacPattern() bool
	// IntegerPattern returns IntegerPattern, set in PrefixConfig.
	// IntegerPattern is test integer pattern
	IntegerPattern() IntegerPattern
	// SetIntegerPattern assigns IntegerPattern provided by user to PrefixConfig.
	// IntegerPattern is test integer pattern
	SetIntegerPattern(value IntegerPattern) PrefixConfig
	// HasIntegerPattern checks if IntegerPattern has been set in PrefixConfig
	HasIntegerPattern() bool
	// ChecksumPattern returns ChecksumPattern, set in PrefixConfig.
	// ChecksumPattern is test checksum pattern
	ChecksumPattern() ChecksumPattern
	// SetChecksumPattern assigns ChecksumPattern provided by user to PrefixConfig.
	// ChecksumPattern is test checksum pattern
	SetChecksumPattern(value ChecksumPattern) PrefixConfig
	// HasChecksumPattern checks if ChecksumPattern has been set in PrefixConfig
	HasChecksumPattern() bool
	// Case returns Layer1Ieee802X, set in PrefixConfig.
	Case() Layer1Ieee802X
	// SetCase assigns Layer1Ieee802X provided by user to PrefixConfig.
	SetCase(value Layer1Ieee802X) PrefixConfig
	// HasCase checks if Case has been set in PrefixConfig
	HasCase() bool
	// MObject returns MObject, set in PrefixConfig.
	// MObject is required format validation object
	MObject() MObject
	// SetMObject assigns MObject provided by user to PrefixConfig.
	// MObject is required format validation object
	SetMObject(value MObject) PrefixConfig
	// HasMObject checks if MObject has been set in PrefixConfig
	HasMObject() bool
	// Integer64 returns int64, set in PrefixConfig.
	Integer64() int64
	// SetInteger64 assigns int64 provided by user to PrefixConfig
	SetInteger64(value int64) PrefixConfig
	// HasInteger64 checks if Integer64 has been set in PrefixConfig
	HasInteger64() bool
	// Integer64List returns []int64, set in PrefixConfig.
	Integer64List() []int64
	// SetInteger64List assigns []int64 provided by user to PrefixConfig
	SetInteger64List(value []int64) PrefixConfig
	// HeaderChecksum returns PatternPrefixConfigHeaderChecksum, set in PrefixConfig.
	// PatternPrefixConfigHeaderChecksum is header checksum
	HeaderChecksum() PatternPrefixConfigHeaderChecksum
	// SetHeaderChecksum assigns PatternPrefixConfigHeaderChecksum provided by user to PrefixConfig.
	// PatternPrefixConfigHeaderChecksum is header checksum
	SetHeaderChecksum(value PatternPrefixConfigHeaderChecksum) PrefixConfig
	// HasHeaderChecksum checks if HeaderChecksum has been set in PrefixConfig
	HasHeaderChecksum() bool
	// StrLen returns string, set in PrefixConfig.
	StrLen() string
	// SetStrLen assigns string provided by user to PrefixConfig
	SetStrLen(value string) PrefixConfig
	// HasStrLen checks if StrLen has been set in PrefixConfig
	HasStrLen() bool
	// HexSlice returns []string, set in PrefixConfig.
	HexSlice() []string
	// SetHexSlice assigns []string provided by user to PrefixConfig
	SetHexSlice(value []string) PrefixConfig
	// AutoFieldTest returns PatternPrefixConfigAutoFieldTest, set in PrefixConfig.
	// PatternPrefixConfigAutoFieldTest is tBD
	AutoFieldTest() PatternPrefixConfigAutoFieldTest
	// SetAutoFieldTest assigns PatternPrefixConfigAutoFieldTest provided by user to PrefixConfig.
	// PatternPrefixConfigAutoFieldTest is tBD
	SetAutoFieldTest(value PatternPrefixConfigAutoFieldTest) PrefixConfig
	// HasAutoFieldTest checks if AutoFieldTest has been set in PrefixConfig
	HasAutoFieldTest() bool
	// Name returns string, set in PrefixConfig.
	Name() string
	// SetName assigns string provided by user to PrefixConfig
	SetName(value string) PrefixConfig
	// HasName checks if Name has been set in PrefixConfig
	HasName() bool
	// WList returns PrefixConfigWObjectIterIter, set in PrefixConfig
	WList() PrefixConfigWObjectIter
	// XList returns PrefixConfigZObjectIterIter, set in PrefixConfig
	XList() PrefixConfigZObjectIter
	// ZObject returns ZObject, set in PrefixConfig.
	// ZObject is description is TBD
	ZObject() ZObject
	// SetZObject assigns ZObject provided by user to PrefixConfig.
	// ZObject is description is TBD
	SetZObject(value ZObject) PrefixConfig
	// HasZObject checks if ZObject has been set in PrefixConfig
	HasZObject() bool
	// YObject returns YObject, set in PrefixConfig.
	// YObject is description is TBD
	YObject() YObject
	// SetYObject assigns YObject provided by user to PrefixConfig.
	// YObject is description is TBD
	SetYObject(value YObject) PrefixConfig
	// HasYObject checks if YObject has been set in PrefixConfig
	HasYObject() bool
	// ChoiceObject returns PrefixConfigChoiceObjectIterIter, set in PrefixConfig
	ChoiceObject() PrefixConfigChoiceObjectIter
	// RequiredChoiceObject returns RequiredChoiceParent, set in PrefixConfig.
	// RequiredChoiceParent is description is TBD
	RequiredChoiceObject() RequiredChoiceParent
	// SetRequiredChoiceObject assigns RequiredChoiceParent provided by user to PrefixConfig.
	// RequiredChoiceParent is description is TBD
	SetRequiredChoiceObject(value RequiredChoiceParent) PrefixConfig
	// HasRequiredChoiceObject checks if RequiredChoiceObject has been set in PrefixConfig
	HasRequiredChoiceObject() bool
	// G1 returns PrefixConfigGObjectIterIter, set in PrefixConfig
	G1() PrefixConfigGObjectIter
	// G2 returns PrefixConfigGObjectIterIter, set in PrefixConfig
	G2() PrefixConfigGObjectIter
	// Int32Param returns int32, set in PrefixConfig.
	Int32Param() int32
	// SetInt32Param assigns int32 provided by user to PrefixConfig
	SetInt32Param(value int32) PrefixConfig
	// HasInt32Param checks if Int32Param has been set in PrefixConfig
	HasInt32Param() bool
	// Int32ListParam returns []int32, set in PrefixConfig.
	Int32ListParam() []int32
	// SetInt32ListParam assigns []int32 provided by user to PrefixConfig
	SetInt32ListParam(value []int32) PrefixConfig
	// Uint32Param returns uint32, set in PrefixConfig.
	Uint32Param() uint32
	// SetUint32Param assigns uint32 provided by user to PrefixConfig
	SetUint32Param(value uint32) PrefixConfig
	// HasUint32Param checks if Uint32Param has been set in PrefixConfig
	HasUint32Param() bool
	// Uint32ListParam returns []uint32, set in PrefixConfig.
	Uint32ListParam() []uint32
	// SetUint32ListParam assigns []uint32 provided by user to PrefixConfig
	SetUint32ListParam(value []uint32) PrefixConfig
	// Uint64Param returns uint64, set in PrefixConfig.
	Uint64Param() uint64
	// SetUint64Param assigns uint64 provided by user to PrefixConfig
	SetUint64Param(value uint64) PrefixConfig
	// HasUint64Param checks if Uint64Param has been set in PrefixConfig
	HasUint64Param() bool
	// Uint64ListParam returns []uint64, set in PrefixConfig.
	Uint64ListParam() []uint64
	// SetUint64ListParam assigns []uint64 provided by user to PrefixConfig
	SetUint64ListParam(value []uint64) PrefixConfig
	// AutoInt32Param returns int32, set in PrefixConfig.
	AutoInt32Param() int32
	// SetAutoInt32Param assigns int32 provided by user to PrefixConfig
	SetAutoInt32Param(value int32) PrefixConfig
	// HasAutoInt32Param checks if AutoInt32Param has been set in PrefixConfig
	HasAutoInt32Param() bool
	// AutoInt32ListParam returns []int32, set in PrefixConfig.
	AutoInt32ListParam() []int32
	// SetAutoInt32ListParam assigns []int32 provided by user to PrefixConfig
	SetAutoInt32ListParam(value []int32) PrefixConfig
	// ChoiceTest returns ChoiceTestObj, set in PrefixConfig.
	// ChoiceTestObj is description is TBD
	ChoiceTest() ChoiceTestObj
	// SetChoiceTest assigns ChoiceTestObj provided by user to PrefixConfig.
	// ChoiceTestObj is description is TBD
	SetChoiceTest(value ChoiceTestObj) PrefixConfig
	// HasChoiceTest checks if ChoiceTest has been set in PrefixConfig
	HasChoiceTest() bool
	// SignedIntegerPattern returns SignedIntegerPattern, set in PrefixConfig.
	// SignedIntegerPattern is test signed integer pattern
	SignedIntegerPattern() SignedIntegerPattern
	// SetSignedIntegerPattern assigns SignedIntegerPattern provided by user to PrefixConfig.
	// SignedIntegerPattern is test signed integer pattern
	SetSignedIntegerPattern(value SignedIntegerPattern) PrefixConfig
	// HasSignedIntegerPattern checks if SignedIntegerPattern has been set in PrefixConfig
	HasSignedIntegerPattern() bool
	// OidPattern returns OidPattern, set in PrefixConfig.
	// OidPattern is test oid pattern
	OidPattern() OidPattern
	// SetOidPattern assigns OidPattern provided by user to PrefixConfig.
	// OidPattern is test oid pattern
	SetOidPattern(value OidPattern) PrefixConfig
	// HasOidPattern checks if OidPattern has been set in PrefixConfig
	HasOidPattern() bool
	// ChoiceDefault returns ChoiceObject, set in PrefixConfig.
	// ChoiceObject is description is TBD
	ChoiceDefault() ChoiceObject
	// SetChoiceDefault assigns ChoiceObject provided by user to PrefixConfig.
	// ChoiceObject is description is TBD
	SetChoiceDefault(value ChoiceObject) PrefixConfig
	// HasChoiceDefault checks if ChoiceDefault has been set in PrefixConfig
	HasChoiceDefault() bool
	// ChoiceRequiredDefault returns ChoiceRequiredAndDefault, set in PrefixConfig.
	// ChoiceRequiredAndDefault is description is TBD
	ChoiceRequiredDefault() ChoiceRequiredAndDefault
	// SetChoiceRequiredDefault assigns ChoiceRequiredAndDefault provided by user to PrefixConfig.
	// ChoiceRequiredAndDefault is description is TBD
	SetChoiceRequiredDefault(value ChoiceRequiredAndDefault) PrefixConfig
	// HasChoiceRequiredDefault checks if ChoiceRequiredDefault has been set in PrefixConfig
	HasChoiceRequiredDefault() bool
	// AutoPattern returns AutoPattern, set in PrefixConfig.
	// AutoPattern is test auto pattern
	AutoPattern() AutoPattern
	// SetAutoPattern assigns AutoPattern provided by user to PrefixConfig.
	// AutoPattern is test auto pattern
	SetAutoPattern(value AutoPattern) PrefixConfig
	// HasAutoPattern checks if AutoPattern has been set in PrefixConfig
	HasAutoPattern() bool
	// AutoPatternDefault returns AutoPatternDefault, set in PrefixConfig.
	// AutoPatternDefault is test auto pattern with default
	AutoPatternDefault() AutoPatternDefault
	// SetAutoPatternDefault assigns AutoPatternDefault provided by user to PrefixConfig.
	// AutoPatternDefault is test auto pattern with default
	SetAutoPatternDefault(value AutoPatternDefault) PrefixConfig
	// HasAutoPatternDefault checks if AutoPatternDefault has been set in PrefixConfig
	HasAutoPatternDefault() bool
	// NameEndingWithNumber234 returns string, set in PrefixConfig.
	NameEndingWithNumber234() string
	// SetNameEndingWithNumber234 assigns string provided by user to PrefixConfig
	SetNameEndingWithNumber234(value string) PrefixConfig
	// HasNameEndingWithNumber234 checks if NameEndingWithNumber234 has been set in PrefixConfig
	HasNameEndingWithNumber234() bool
	setNil()
}

// A required object that MUST be generated as such.
// RequiredObject returns a EObject
func (obj *prefixConfig) RequiredObject() EObject {
	if obj.obj.RequiredObject == nil {
		obj.obj.RequiredObject = NewEObject().msg()
	}
	if obj.requiredObjectHolder == nil {
		obj.requiredObjectHolder = &eObject{obj: obj.obj.RequiredObject}
	}
	return obj.requiredObjectHolder
}

// A required object that MUST be generated as such.
// SetRequiredObject sets the EObject value in the PrefixConfig object
func (obj *prefixConfig) SetRequiredObject(value EObject) PrefixConfig {

	obj.requiredObjectHolder = nil
	obj.obj.RequiredObject = value.msg()

	return obj
}

// An optional object that MUST be generated as such.
// OptionalObject returns a EObject
func (obj *prefixConfig) OptionalObject() EObject {
	if obj.obj.OptionalObject == nil {
		obj.obj.OptionalObject = NewEObject().msg()
	}
	if obj.optionalObjectHolder == nil {
		obj.optionalObjectHolder = &eObject{obj: obj.obj.OptionalObject}
	}
	return obj.optionalObjectHolder
}

// An optional object that MUST be generated as such.
// OptionalObject returns a EObject
func (obj *prefixConfig) HasOptionalObject() bool {
	return obj.obj.OptionalObject != nil
}

// An optional object that MUST be generated as such.
// SetOptionalObject sets the EObject value in the PrefixConfig object
func (obj *prefixConfig) SetOptionalObject(value EObject) PrefixConfig {

	obj.optionalObjectHolder = nil
	obj.obj.OptionalObject = value.msg()

	return obj
}

// description is TBD
// Ieee8021Qbb returns a bool
func (obj *prefixConfig) Ieee8021Qbb() bool {

	return *obj.obj.Ieee_802_1Qbb

}

// description is TBD
// Ieee8021Qbb returns a bool
func (obj *prefixConfig) HasIeee8021Qbb() bool {
	return obj.obj.Ieee_802_1Qbb != nil
}

// description is TBD
// SetIeee8021Qbb sets the bool value in the PrefixConfig object
func (obj *prefixConfig) SetIeee8021Qbb(value bool) PrefixConfig {

	obj.obj.Ieee_802_1Qbb = &value
	return obj
}

// Deprecated: Information TBD
//
// Description TBD
// Space1 returns a int32
func (obj *prefixConfig) Space1() int32 {

	return *obj.obj.Space_1

}

// Deprecated: Information TBD
//
// Description TBD
// Space1 returns a int32
func (obj *prefixConfig) HasSpace1() bool {
	return obj.obj.Space_1 != nil
}

// Deprecated: Information TBD
//
// Description TBD
// SetSpace1 sets the int32 value in the PrefixConfig object
func (obj *prefixConfig) SetSpace1(value int32) PrefixConfig {

	obj.obj.Space_1 = &value
	return obj
}

// description is TBD
// FullDuplex100Mb returns a int64
func (obj *prefixConfig) FullDuplex100Mb() int64 {

	return *obj.obj.FullDuplex_100Mb

}

// description is TBD
// FullDuplex100Mb returns a int64
func (obj *prefixConfig) HasFullDuplex100Mb() bool {
	return obj.obj.FullDuplex_100Mb != nil
}

// description is TBD
// SetFullDuplex100Mb sets the int64 value in the PrefixConfig object
func (obj *prefixConfig) SetFullDuplex100Mb(value int64) PrefixConfig {

	obj.obj.FullDuplex_100Mb = &value
	return obj
}

type PrefixConfigResponseEnum string

// Enum of Response on PrefixConfig
var PrefixConfigResponse = struct {
	STATUS_200 PrefixConfigResponseEnum
	STATUS_400 PrefixConfigResponseEnum
	STATUS_404 PrefixConfigResponseEnum
	STATUS_500 PrefixConfigResponseEnum
}{
	STATUS_200: PrefixConfigResponseEnum("status_200"),
	STATUS_400: PrefixConfigResponseEnum("status_400"),
	STATUS_404: PrefixConfigResponseEnum("status_404"),
	STATUS_500: PrefixConfigResponseEnum("status_500"),
}

func (obj *prefixConfig) Response() PrefixConfigResponseEnum {
	return PrefixConfigResponseEnum(obj.obj.Response.Enum().String())
}

// Indicate to the server what response should be returned
// Response returns a string
func (obj *prefixConfig) HasResponse() bool {
	return obj.obj.Response != nil
}

func (obj *prefixConfig) SetResponse(value PrefixConfigResponseEnum) PrefixConfig {
	intValue, ok := openapi.PrefixConfig_Response_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PrefixConfigResponseEnum", string(value)))
		return obj
	}
	enumValue := openapi.PrefixConfig_Response_Enum(intValue)
	obj.obj.Response = &enumValue

	return obj
}

// Under Review: Information TBD
//
// Small single line description
// A returns a string
func (obj *prefixConfig) A() string {

	return *obj.obj.A

}

// Under Review: Information TBD
//
// Small single line description
// SetA sets the string value in the PrefixConfig object
func (obj *prefixConfig) SetA(value string) PrefixConfig {

	obj.obj.A = &value
	return obj
}

// Longer multi-line description
// Second line is here
// Third line
// B returns a float32
func (obj *prefixConfig) B() float32 {

	return *obj.obj.B

}

// Longer multi-line description
// Second line is here
// Third line
// SetB sets the float32 value in the PrefixConfig object
func (obj *prefixConfig) SetB(value float32) PrefixConfig {

	obj.obj.B = &value
	return obj
}

// description is TBD
// C returns a int32
func (obj *prefixConfig) C() int32 {

	return *obj.obj.C

}

// description is TBD
// SetC sets the int32 value in the PrefixConfig object
func (obj *prefixConfig) SetC(value int32) PrefixConfig {

	obj.obj.C = &value
	return obj
}

type PrefixConfigDValuesEnum string

// Enum of DValues on PrefixConfig
var PrefixConfigDValues = struct {
	A PrefixConfigDValuesEnum
	B PrefixConfigDValuesEnum
	C PrefixConfigDValuesEnum
}{
	A: PrefixConfigDValuesEnum("a"),
	B: PrefixConfigDValuesEnum("b"),
	C: PrefixConfigDValuesEnum("c"),
}

func (obj *prefixConfig) DValues() []PrefixConfigDValuesEnum {
	items := []PrefixConfigDValuesEnum{}
	for _, item := range obj.obj.DValues {
		items = append(items, PrefixConfigDValuesEnum(item.String()))
	}
	return items
}

// Deprecated: Information TBD
//
// A list of enum values
// SetDValues sets the []string value in the PrefixConfig object
func (obj *prefixConfig) SetDValues(value []PrefixConfigDValuesEnum) PrefixConfig {

	items := []openapi.PrefixConfig_DValues_Enum{}
	for _, item := range value {
		intValue := openapi.PrefixConfig_DValues_Enum_value[string(item)]
		items = append(items, openapi.PrefixConfig_DValues_Enum(intValue))
	}
	obj.obj.DValues = items
	return obj
}

// Deprecated: Information TBD
//
// A child object
// E returns a EObject
func (obj *prefixConfig) E() EObject {
	if obj.obj.E == nil {
		obj.obj.E = NewEObject().msg()
	}
	if obj.eHolder == nil {
		obj.eHolder = &eObject{obj: obj.obj.E}
	}
	return obj.eHolder
}

// Deprecated: Information TBD
//
// A child object
// E returns a EObject
func (obj *prefixConfig) HasE() bool {
	return obj.obj.E != nil
}

// Deprecated: Information TBD
//
// A child object
// SetE sets the EObject value in the PrefixConfig object
func (obj *prefixConfig) SetE(value EObject) PrefixConfig {

	obj.eHolder = nil
	obj.obj.E = value.msg()

	return obj
}

// An object with only choice(s)
// F returns a FObject
func (obj *prefixConfig) F() FObject {
	if obj.obj.F == nil {
		obj.obj.F = NewFObject().msg()
	}
	if obj.fHolder == nil {
		obj.fHolder = &fObject{obj: obj.obj.F}
	}
	return obj.fHolder
}

// An object with only choice(s)
// F returns a FObject
func (obj *prefixConfig) HasF() bool {
	return obj.obj.F != nil
}

// An object with only choice(s)
// SetF sets the FObject value in the PrefixConfig object
func (obj *prefixConfig) SetF(value FObject) PrefixConfig {

	obj.fHolder = nil
	obj.obj.F = value.msg()

	return obj
}

// A list of objects with choice and properties
// G returns a []GObject
func (obj *prefixConfig) G() PrefixConfigGObjectIter {
	if len(obj.obj.G) == 0 {
		obj.obj.G = []*openapi.GObject{}
	}
	if obj.gHolder == nil {
		obj.gHolder = newPrefixConfigGObjectIter(&obj.obj.G).setMsg(obj)
	}
	return obj.gHolder
}

type prefixConfigGObjectIter struct {
	obj          *prefixConfig
	gObjectSlice []GObject
	fieldPtr     *[]*openapi.GObject
}

func newPrefixConfigGObjectIter(ptr *[]*openapi.GObject) PrefixConfigGObjectIter {
	return &prefixConfigGObjectIter{fieldPtr: ptr}
}

type PrefixConfigGObjectIter interface {
	setMsg(*prefixConfig) PrefixConfigGObjectIter
	Items() []GObject
	Add() GObject
	Append(items ...GObject) PrefixConfigGObjectIter
	Set(index int, newObj GObject) PrefixConfigGObjectIter
	Clear() PrefixConfigGObjectIter
	clearHolderSlice() PrefixConfigGObjectIter
	appendHolderSlice(item GObject) PrefixConfigGObjectIter
}

func (obj *prefixConfigGObjectIter) setMsg(msg *prefixConfig) PrefixConfigGObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&gObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *prefixConfigGObjectIter) Items() []GObject {
	return obj.gObjectSlice
}

func (obj *prefixConfigGObjectIter) Add() GObject {
	newObj := &openapi.GObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &gObject{obj: newObj}
	newLibObj.setDefault()
	obj.gObjectSlice = append(obj.gObjectSlice, newLibObj)
	return newLibObj
}

func (obj *prefixConfigGObjectIter) Append(items ...GObject) PrefixConfigGObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.gObjectSlice = append(obj.gObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigGObjectIter) Set(index int, newObj GObject) PrefixConfigGObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.gObjectSlice[index] = newObj
	return obj
}
func (obj *prefixConfigGObjectIter) Clear() PrefixConfigGObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.GObject{}
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *prefixConfigGObjectIter) clearHolderSlice() PrefixConfigGObjectIter {
	if len(obj.gObjectSlice) > 0 {
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *prefixConfigGObjectIter) appendHolderSlice(item GObject) PrefixConfigGObjectIter {
	obj.gObjectSlice = append(obj.gObjectSlice, item)
	return obj
}

// A boolean value
// H returns a bool
func (obj *prefixConfig) H() bool {

	return *obj.obj.H

}

// A boolean value
// H returns a bool
func (obj *prefixConfig) HasH() bool {
	return obj.obj.H != nil
}

// A boolean value
// SetH sets the bool value in the PrefixConfig object
func (obj *prefixConfig) SetH(value bool) PrefixConfig {

	obj.obj.H = &value
	return obj
}

// A byte string
// I returns a []byte
func (obj *prefixConfig) I() []byte {

	return obj.obj.I
}

// A byte string
// SetI sets the []byte value in the PrefixConfig object
func (obj *prefixConfig) SetI(value []byte) PrefixConfig {

	obj.obj.I = value
	return obj
}

// A list of objects with only choice
// J returns a []JObject
func (obj *prefixConfig) J() PrefixConfigJObjectIter {
	if len(obj.obj.J) == 0 {
		obj.obj.J = []*openapi.JObject{}
	}
	if obj.jHolder == nil {
		obj.jHolder = newPrefixConfigJObjectIter(&obj.obj.J).setMsg(obj)
	}
	return obj.jHolder
}

type prefixConfigJObjectIter struct {
	obj          *prefixConfig
	jObjectSlice []JObject
	fieldPtr     *[]*openapi.JObject
}

func newPrefixConfigJObjectIter(ptr *[]*openapi.JObject) PrefixConfigJObjectIter {
	return &prefixConfigJObjectIter{fieldPtr: ptr}
}

type PrefixConfigJObjectIter interface {
	setMsg(*prefixConfig) PrefixConfigJObjectIter
	Items() []JObject
	Add() JObject
	Append(items ...JObject) PrefixConfigJObjectIter
	Set(index int, newObj JObject) PrefixConfigJObjectIter
	Clear() PrefixConfigJObjectIter
	clearHolderSlice() PrefixConfigJObjectIter
	appendHolderSlice(item JObject) PrefixConfigJObjectIter
}

func (obj *prefixConfigJObjectIter) setMsg(msg *prefixConfig) PrefixConfigJObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&jObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *prefixConfigJObjectIter) Items() []JObject {
	return obj.jObjectSlice
}

func (obj *prefixConfigJObjectIter) Add() JObject {
	newObj := &openapi.JObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &jObject{obj: newObj}
	newLibObj.setDefault()
	obj.jObjectSlice = append(obj.jObjectSlice, newLibObj)
	return newLibObj
}

func (obj *prefixConfigJObjectIter) Append(items ...JObject) PrefixConfigJObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.jObjectSlice = append(obj.jObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigJObjectIter) Set(index int, newObj JObject) PrefixConfigJObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.jObjectSlice[index] = newObj
	return obj
}
func (obj *prefixConfigJObjectIter) Clear() PrefixConfigJObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.JObject{}
		obj.jObjectSlice = []JObject{}
	}
	return obj
}
func (obj *prefixConfigJObjectIter) clearHolderSlice() PrefixConfigJObjectIter {
	if len(obj.jObjectSlice) > 0 {
		obj.jObjectSlice = []JObject{}
	}
	return obj
}
func (obj *prefixConfigJObjectIter) appendHolderSlice(item JObject) PrefixConfigJObjectIter {
	obj.jObjectSlice = append(obj.jObjectSlice, item)
	return obj
}

// A nested object with only one property which is a choice object
// K returns a KObject
func (obj *prefixConfig) K() KObject {
	if obj.obj.K == nil {
		obj.obj.K = NewKObject().msg()
	}
	if obj.kHolder == nil {
		obj.kHolder = &kObject{obj: obj.obj.K}
	}
	return obj.kHolder
}

// A nested object with only one property which is a choice object
// K returns a KObject
func (obj *prefixConfig) HasK() bool {
	return obj.obj.K != nil
}

// A nested object with only one property which is a choice object
// SetK sets the KObject value in the PrefixConfig object
func (obj *prefixConfig) SetK(value KObject) PrefixConfig {

	obj.kHolder = nil
	obj.obj.K = value.msg()

	return obj
}

// description is TBD
// L returns a LObject
func (obj *prefixConfig) L() LObject {
	if obj.obj.L == nil {
		obj.obj.L = NewLObject().msg()
	}
	if obj.lHolder == nil {
		obj.lHolder = &lObject{obj: obj.obj.L}
	}
	return obj.lHolder
}

// description is TBD
// L returns a LObject
func (obj *prefixConfig) HasL() bool {
	return obj.obj.L != nil
}

// description is TBD
// SetL sets the LObject value in the PrefixConfig object
func (obj *prefixConfig) SetL(value LObject) PrefixConfig {

	obj.lHolder = nil
	obj.obj.L = value.msg()

	return obj
}

// A list of string values
// ListOfStringValues returns a []string
func (obj *prefixConfig) ListOfStringValues() []string {
	if obj.obj.ListOfStringValues == nil {
		obj.obj.ListOfStringValues = make([]string, 0)
	}
	return obj.obj.ListOfStringValues
}

// A list of string values
// SetListOfStringValues sets the []string value in the PrefixConfig object
func (obj *prefixConfig) SetListOfStringValues(value []string) PrefixConfig {

	if obj.obj.ListOfStringValues == nil {
		obj.obj.ListOfStringValues = make([]string, 0)
	}
	obj.obj.ListOfStringValues = value

	return obj
}

// A list of integer values
// ListOfIntegerValues returns a []int32
func (obj *prefixConfig) ListOfIntegerValues() []int32 {
	if obj.obj.ListOfIntegerValues == nil {
		obj.obj.ListOfIntegerValues = make([]int32, 0)
	}
	return obj.obj.ListOfIntegerValues
}

// A list of integer values
// SetListOfIntegerValues sets the []int32 value in the PrefixConfig object
func (obj *prefixConfig) SetListOfIntegerValues(value []int32) PrefixConfig {

	if obj.obj.ListOfIntegerValues == nil {
		obj.obj.ListOfIntegerValues = make([]int32, 0)
	}
	obj.obj.ListOfIntegerValues = value

	return obj
}

// description is TBD
// Level returns a LevelOne
func (obj *prefixConfig) Level() LevelOne {
	if obj.obj.Level == nil {
		obj.obj.Level = NewLevelOne().msg()
	}
	if obj.levelHolder == nil {
		obj.levelHolder = &levelOne{obj: obj.obj.Level}
	}
	return obj.levelHolder
}

// description is TBD
// Level returns a LevelOne
func (obj *prefixConfig) HasLevel() bool {
	return obj.obj.Level != nil
}

// description is TBD
// SetLevel sets the LevelOne value in the PrefixConfig object
func (obj *prefixConfig) SetLevel(value LevelOne) PrefixConfig {

	obj.levelHolder = nil
	obj.obj.Level = value.msg()

	return obj
}

// description is TBD
// Mandatory returns a Mandate
func (obj *prefixConfig) Mandatory() Mandate {
	if obj.obj.Mandatory == nil {
		obj.obj.Mandatory = NewMandate().msg()
	}
	if obj.mandatoryHolder == nil {
		obj.mandatoryHolder = &mandate{obj: obj.obj.Mandatory}
	}
	return obj.mandatoryHolder
}

// description is TBD
// Mandatory returns a Mandate
func (obj *prefixConfig) HasMandatory() bool {
	return obj.obj.Mandatory != nil
}

// description is TBD
// SetMandatory sets the Mandate value in the PrefixConfig object
func (obj *prefixConfig) SetMandatory(value Mandate) PrefixConfig {

	obj.mandatoryHolder = nil
	obj.obj.Mandatory = value.msg()

	return obj
}

// description is TBD
// Ipv4Pattern returns a Ipv4Pattern
func (obj *prefixConfig) Ipv4Pattern() Ipv4Pattern {
	if obj.obj.Ipv4Pattern == nil {
		obj.obj.Ipv4Pattern = NewIpv4Pattern().msg()
	}
	if obj.ipv4PatternHolder == nil {
		obj.ipv4PatternHolder = &ipv4Pattern{obj: obj.obj.Ipv4Pattern}
	}
	return obj.ipv4PatternHolder
}

// description is TBD
// Ipv4Pattern returns a Ipv4Pattern
func (obj *prefixConfig) HasIpv4Pattern() bool {
	return obj.obj.Ipv4Pattern != nil
}

// description is TBD
// SetIpv4Pattern sets the Ipv4Pattern value in the PrefixConfig object
func (obj *prefixConfig) SetIpv4Pattern(value Ipv4Pattern) PrefixConfig {

	obj.ipv4PatternHolder = nil
	obj.obj.Ipv4Pattern = value.msg()

	return obj
}

// description is TBD
// Ipv6Pattern returns a Ipv6Pattern
func (obj *prefixConfig) Ipv6Pattern() Ipv6Pattern {
	if obj.obj.Ipv6Pattern == nil {
		obj.obj.Ipv6Pattern = NewIpv6Pattern().msg()
	}
	if obj.ipv6PatternHolder == nil {
		obj.ipv6PatternHolder = &ipv6Pattern{obj: obj.obj.Ipv6Pattern}
	}
	return obj.ipv6PatternHolder
}

// description is TBD
// Ipv6Pattern returns a Ipv6Pattern
func (obj *prefixConfig) HasIpv6Pattern() bool {
	return obj.obj.Ipv6Pattern != nil
}

// description is TBD
// SetIpv6Pattern sets the Ipv6Pattern value in the PrefixConfig object
func (obj *prefixConfig) SetIpv6Pattern(value Ipv6Pattern) PrefixConfig {

	obj.ipv6PatternHolder = nil
	obj.obj.Ipv6Pattern = value.msg()

	return obj
}

// description is TBD
// MacPattern returns a MacPattern
func (obj *prefixConfig) MacPattern() MacPattern {
	if obj.obj.MacPattern == nil {
		obj.obj.MacPattern = NewMacPattern().msg()
	}
	if obj.macPatternHolder == nil {
		obj.macPatternHolder = &macPattern{obj: obj.obj.MacPattern}
	}
	return obj.macPatternHolder
}

// description is TBD
// MacPattern returns a MacPattern
func (obj *prefixConfig) HasMacPattern() bool {
	return obj.obj.MacPattern != nil
}

// description is TBD
// SetMacPattern sets the MacPattern value in the PrefixConfig object
func (obj *prefixConfig) SetMacPattern(value MacPattern) PrefixConfig {

	obj.macPatternHolder = nil
	obj.obj.MacPattern = value.msg()

	return obj
}

// description is TBD
// IntegerPattern returns a IntegerPattern
func (obj *prefixConfig) IntegerPattern() IntegerPattern {
	if obj.obj.IntegerPattern == nil {
		obj.obj.IntegerPattern = NewIntegerPattern().msg()
	}
	if obj.integerPatternHolder == nil {
		obj.integerPatternHolder = &integerPattern{obj: obj.obj.IntegerPattern}
	}
	return obj.integerPatternHolder
}

// description is TBD
// IntegerPattern returns a IntegerPattern
func (obj *prefixConfig) HasIntegerPattern() bool {
	return obj.obj.IntegerPattern != nil
}

// description is TBD
// SetIntegerPattern sets the IntegerPattern value in the PrefixConfig object
func (obj *prefixConfig) SetIntegerPattern(value IntegerPattern) PrefixConfig {

	obj.integerPatternHolder = nil
	obj.obj.IntegerPattern = value.msg()

	return obj
}

// description is TBD
// ChecksumPattern returns a ChecksumPattern
func (obj *prefixConfig) ChecksumPattern() ChecksumPattern {
	if obj.obj.ChecksumPattern == nil {
		obj.obj.ChecksumPattern = NewChecksumPattern().msg()
	}
	if obj.checksumPatternHolder == nil {
		obj.checksumPatternHolder = &checksumPattern{obj: obj.obj.ChecksumPattern}
	}
	return obj.checksumPatternHolder
}

// description is TBD
// ChecksumPattern returns a ChecksumPattern
func (obj *prefixConfig) HasChecksumPattern() bool {
	return obj.obj.ChecksumPattern != nil
}

// description is TBD
// SetChecksumPattern sets the ChecksumPattern value in the PrefixConfig object
func (obj *prefixConfig) SetChecksumPattern(value ChecksumPattern) PrefixConfig {

	obj.checksumPatternHolder = nil
	obj.obj.ChecksumPattern = value.msg()

	return obj
}

// description is TBD
// Case returns a Layer1Ieee802X
func (obj *prefixConfig) Case() Layer1Ieee802X {
	if obj.obj.Case == nil {
		obj.obj.Case = NewLayer1Ieee802X().msg()
	}
	if obj.caseHolder == nil {
		obj.caseHolder = &layer1Ieee802X{obj: obj.obj.Case}
	}
	return obj.caseHolder
}

// description is TBD
// Case returns a Layer1Ieee802X
func (obj *prefixConfig) HasCase() bool {
	return obj.obj.Case != nil
}

// description is TBD
// SetCase sets the Layer1Ieee802X value in the PrefixConfig object
func (obj *prefixConfig) SetCase(value Layer1Ieee802X) PrefixConfig {

	obj.caseHolder = nil
	obj.obj.Case = value.msg()

	return obj
}

// description is TBD
// MObject returns a MObject
func (obj *prefixConfig) MObject() MObject {
	if obj.obj.MObject == nil {
		obj.obj.MObject = NewMObject().msg()
	}
	if obj.mObjectHolder == nil {
		obj.mObjectHolder = &mObject{obj: obj.obj.MObject}
	}
	return obj.mObjectHolder
}

// description is TBD
// MObject returns a MObject
func (obj *prefixConfig) HasMObject() bool {
	return obj.obj.MObject != nil
}

// description is TBD
// SetMObject sets the MObject value in the PrefixConfig object
func (obj *prefixConfig) SetMObject(value MObject) PrefixConfig {

	obj.mObjectHolder = nil
	obj.obj.MObject = value.msg()

	return obj
}

// int64 type
// Integer64 returns a int64
func (obj *prefixConfig) Integer64() int64 {

	return *obj.obj.Integer64

}

// int64 type
// Integer64 returns a int64
func (obj *prefixConfig) HasInteger64() bool {
	return obj.obj.Integer64 != nil
}

// int64 type
// SetInteger64 sets the int64 value in the PrefixConfig object
func (obj *prefixConfig) SetInteger64(value int64) PrefixConfig {

	obj.obj.Integer64 = &value
	return obj
}

// int64 type list
// Integer64List returns a []int64
func (obj *prefixConfig) Integer64List() []int64 {
	if obj.obj.Integer64List == nil {
		obj.obj.Integer64List = make([]int64, 0)
	}
	return obj.obj.Integer64List
}

// int64 type list
// SetInteger64List sets the []int64 value in the PrefixConfig object
func (obj *prefixConfig) SetInteger64List(value []int64) PrefixConfig {

	if obj.obj.Integer64List == nil {
		obj.obj.Integer64List = make([]int64, 0)
	}
	obj.obj.Integer64List = value

	return obj
}

// description is TBD
// HeaderChecksum returns a PatternPrefixConfigHeaderChecksum
func (obj *prefixConfig) HeaderChecksum() PatternPrefixConfigHeaderChecksum {
	if obj.obj.HeaderChecksum == nil {
		obj.obj.HeaderChecksum = NewPatternPrefixConfigHeaderChecksum().msg()
	}
	if obj.headerChecksumHolder == nil {
		obj.headerChecksumHolder = &patternPrefixConfigHeaderChecksum{obj: obj.obj.HeaderChecksum}
	}
	return obj.headerChecksumHolder
}

// description is TBD
// HeaderChecksum returns a PatternPrefixConfigHeaderChecksum
func (obj *prefixConfig) HasHeaderChecksum() bool {
	return obj.obj.HeaderChecksum != nil
}

// description is TBD
// SetHeaderChecksum sets the PatternPrefixConfigHeaderChecksum value in the PrefixConfig object
func (obj *prefixConfig) SetHeaderChecksum(value PatternPrefixConfigHeaderChecksum) PrefixConfig {

	obj.headerChecksumHolder = nil
	obj.obj.HeaderChecksum = value.msg()

	return obj
}

// Under Review: Information TBD
//
// string minimum&maximum Length
// StrLen returns a string
func (obj *prefixConfig) StrLen() string {

	return *obj.obj.StrLen

}

// Under Review: Information TBD
//
// string minimum&maximum Length
// StrLen returns a string
func (obj *prefixConfig) HasStrLen() bool {
	return obj.obj.StrLen != nil
}

// Under Review: Information TBD
//
// string minimum&maximum Length
// SetStrLen sets the string value in the PrefixConfig object
func (obj *prefixConfig) SetStrLen(value string) PrefixConfig {

	obj.obj.StrLen = &value
	return obj
}

// Under Review: Information TBD
//
// Array of Hex
// HexSlice returns a []string
func (obj *prefixConfig) HexSlice() []string {
	if obj.obj.HexSlice == nil {
		obj.obj.HexSlice = make([]string, 0)
	}
	return obj.obj.HexSlice
}

// Under Review: Information TBD
//
// Array of Hex
// SetHexSlice sets the []string value in the PrefixConfig object
func (obj *prefixConfig) SetHexSlice(value []string) PrefixConfig {

	if obj.obj.HexSlice == nil {
		obj.obj.HexSlice = make([]string, 0)
	}
	obj.obj.HexSlice = value

	return obj
}

// description is TBD
// AutoFieldTest returns a PatternPrefixConfigAutoFieldTest
func (obj *prefixConfig) AutoFieldTest() PatternPrefixConfigAutoFieldTest {
	if obj.obj.AutoFieldTest == nil {
		obj.obj.AutoFieldTest = NewPatternPrefixConfigAutoFieldTest().msg()
	}
	if obj.autoFieldTestHolder == nil {
		obj.autoFieldTestHolder = &patternPrefixConfigAutoFieldTest{obj: obj.obj.AutoFieldTest}
	}
	return obj.autoFieldTestHolder
}

// description is TBD
// AutoFieldTest returns a PatternPrefixConfigAutoFieldTest
func (obj *prefixConfig) HasAutoFieldTest() bool {
	return obj.obj.AutoFieldTest != nil
}

// description is TBD
// SetAutoFieldTest sets the PatternPrefixConfigAutoFieldTest value in the PrefixConfig object
func (obj *prefixConfig) SetAutoFieldTest(value PatternPrefixConfigAutoFieldTest) PrefixConfig {

	obj.autoFieldTestHolder = nil
	obj.obj.AutoFieldTest = value.msg()

	return obj
}

// description is TBD
// Name returns a string
func (obj *prefixConfig) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *prefixConfig) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the PrefixConfig object
func (obj *prefixConfig) SetName(value string) PrefixConfig {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// WList returns a []WObject
func (obj *prefixConfig) WList() PrefixConfigWObjectIter {
	if len(obj.obj.WList) == 0 {
		obj.obj.WList = []*openapi.WObject{}
	}
	if obj.wListHolder == nil {
		obj.wListHolder = newPrefixConfigWObjectIter(&obj.obj.WList).setMsg(obj)
	}
	return obj.wListHolder
}

type prefixConfigWObjectIter struct {
	obj          *prefixConfig
	wObjectSlice []WObject
	fieldPtr     *[]*openapi.WObject
}

func newPrefixConfigWObjectIter(ptr *[]*openapi.WObject) PrefixConfigWObjectIter {
	return &prefixConfigWObjectIter{fieldPtr: ptr}
}

type PrefixConfigWObjectIter interface {
	setMsg(*prefixConfig) PrefixConfigWObjectIter
	Items() []WObject
	Add() WObject
	Append(items ...WObject) PrefixConfigWObjectIter
	Set(index int, newObj WObject) PrefixConfigWObjectIter
	Clear() PrefixConfigWObjectIter
	clearHolderSlice() PrefixConfigWObjectIter
	appendHolderSlice(item WObject) PrefixConfigWObjectIter
}

func (obj *prefixConfigWObjectIter) setMsg(msg *prefixConfig) PrefixConfigWObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&wObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *prefixConfigWObjectIter) Items() []WObject {
	return obj.wObjectSlice
}

func (obj *prefixConfigWObjectIter) Add() WObject {
	newObj := &openapi.WObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &wObject{obj: newObj}
	newLibObj.setDefault()
	obj.wObjectSlice = append(obj.wObjectSlice, newLibObj)
	return newLibObj
}

func (obj *prefixConfigWObjectIter) Append(items ...WObject) PrefixConfigWObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.wObjectSlice = append(obj.wObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigWObjectIter) Set(index int, newObj WObject) PrefixConfigWObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.wObjectSlice[index] = newObj
	return obj
}
func (obj *prefixConfigWObjectIter) Clear() PrefixConfigWObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.WObject{}
		obj.wObjectSlice = []WObject{}
	}
	return obj
}
func (obj *prefixConfigWObjectIter) clearHolderSlice() PrefixConfigWObjectIter {
	if len(obj.wObjectSlice) > 0 {
		obj.wObjectSlice = []WObject{}
	}
	return obj
}
func (obj *prefixConfigWObjectIter) appendHolderSlice(item WObject) PrefixConfigWObjectIter {
	obj.wObjectSlice = append(obj.wObjectSlice, item)
	return obj
}

// description is TBD
// XList returns a []ZObject
func (obj *prefixConfig) XList() PrefixConfigZObjectIter {
	if len(obj.obj.XList) == 0 {
		obj.obj.XList = []*openapi.ZObject{}
	}
	if obj.xListHolder == nil {
		obj.xListHolder = newPrefixConfigZObjectIter(&obj.obj.XList).setMsg(obj)
	}
	return obj.xListHolder
}

type prefixConfigZObjectIter struct {
	obj          *prefixConfig
	zObjectSlice []ZObject
	fieldPtr     *[]*openapi.ZObject
}

func newPrefixConfigZObjectIter(ptr *[]*openapi.ZObject) PrefixConfigZObjectIter {
	return &prefixConfigZObjectIter{fieldPtr: ptr}
}

type PrefixConfigZObjectIter interface {
	setMsg(*prefixConfig) PrefixConfigZObjectIter
	Items() []ZObject
	Add() ZObject
	Append(items ...ZObject) PrefixConfigZObjectIter
	Set(index int, newObj ZObject) PrefixConfigZObjectIter
	Clear() PrefixConfigZObjectIter
	clearHolderSlice() PrefixConfigZObjectIter
	appendHolderSlice(item ZObject) PrefixConfigZObjectIter
}

func (obj *prefixConfigZObjectIter) setMsg(msg *prefixConfig) PrefixConfigZObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&zObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *prefixConfigZObjectIter) Items() []ZObject {
	return obj.zObjectSlice
}

func (obj *prefixConfigZObjectIter) Add() ZObject {
	newObj := &openapi.ZObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &zObject{obj: newObj}
	newLibObj.setDefault()
	obj.zObjectSlice = append(obj.zObjectSlice, newLibObj)
	return newLibObj
}

func (obj *prefixConfigZObjectIter) Append(items ...ZObject) PrefixConfigZObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.zObjectSlice = append(obj.zObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigZObjectIter) Set(index int, newObj ZObject) PrefixConfigZObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.zObjectSlice[index] = newObj
	return obj
}
func (obj *prefixConfigZObjectIter) Clear() PrefixConfigZObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.ZObject{}
		obj.zObjectSlice = []ZObject{}
	}
	return obj
}
func (obj *prefixConfigZObjectIter) clearHolderSlice() PrefixConfigZObjectIter {
	if len(obj.zObjectSlice) > 0 {
		obj.zObjectSlice = []ZObject{}
	}
	return obj
}
func (obj *prefixConfigZObjectIter) appendHolderSlice(item ZObject) PrefixConfigZObjectIter {
	obj.zObjectSlice = append(obj.zObjectSlice, item)
	return obj
}

// description is TBD
// ZObject returns a ZObject
func (obj *prefixConfig) ZObject() ZObject {
	if obj.obj.ZObject == nil {
		obj.obj.ZObject = NewZObject().msg()
	}
	if obj.zObjectHolder == nil {
		obj.zObjectHolder = &zObject{obj: obj.obj.ZObject}
	}
	return obj.zObjectHolder
}

// description is TBD
// ZObject returns a ZObject
func (obj *prefixConfig) HasZObject() bool {
	return obj.obj.ZObject != nil
}

// description is TBD
// SetZObject sets the ZObject value in the PrefixConfig object
func (obj *prefixConfig) SetZObject(value ZObject) PrefixConfig {

	obj.zObjectHolder = nil
	obj.obj.ZObject = value.msg()

	return obj
}

// description is TBD
// YObject returns a YObject
func (obj *prefixConfig) YObject() YObject {
	if obj.obj.YObject == nil {
		obj.obj.YObject = NewYObject().msg()
	}
	if obj.yObjectHolder == nil {
		obj.yObjectHolder = &yObject{obj: obj.obj.YObject}
	}
	return obj.yObjectHolder
}

// description is TBD
// YObject returns a YObject
func (obj *prefixConfig) HasYObject() bool {
	return obj.obj.YObject != nil
}

// description is TBD
// SetYObject sets the YObject value in the PrefixConfig object
func (obj *prefixConfig) SetYObject(value YObject) PrefixConfig {

	obj.yObjectHolder = nil
	obj.obj.YObject = value.msg()

	return obj
}

// A list of objects with choice with and without properties
// ChoiceObject returns a []ChoiceObject
func (obj *prefixConfig) ChoiceObject() PrefixConfigChoiceObjectIter {
	if len(obj.obj.ChoiceObject) == 0 {
		obj.obj.ChoiceObject = []*openapi.ChoiceObject{}
	}
	if obj.choiceObjectHolder == nil {
		obj.choiceObjectHolder = newPrefixConfigChoiceObjectIter(&obj.obj.ChoiceObject).setMsg(obj)
	}
	return obj.choiceObjectHolder
}

type prefixConfigChoiceObjectIter struct {
	obj               *prefixConfig
	choiceObjectSlice []ChoiceObject
	fieldPtr          *[]*openapi.ChoiceObject
}

func newPrefixConfigChoiceObjectIter(ptr *[]*openapi.ChoiceObject) PrefixConfigChoiceObjectIter {
	return &prefixConfigChoiceObjectIter{fieldPtr: ptr}
}

type PrefixConfigChoiceObjectIter interface {
	setMsg(*prefixConfig) PrefixConfigChoiceObjectIter
	Items() []ChoiceObject
	Add() ChoiceObject
	Append(items ...ChoiceObject) PrefixConfigChoiceObjectIter
	Set(index int, newObj ChoiceObject) PrefixConfigChoiceObjectIter
	Clear() PrefixConfigChoiceObjectIter
	clearHolderSlice() PrefixConfigChoiceObjectIter
	appendHolderSlice(item ChoiceObject) PrefixConfigChoiceObjectIter
}

func (obj *prefixConfigChoiceObjectIter) setMsg(msg *prefixConfig) PrefixConfigChoiceObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&choiceObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *prefixConfigChoiceObjectIter) Items() []ChoiceObject {
	return obj.choiceObjectSlice
}

func (obj *prefixConfigChoiceObjectIter) Add() ChoiceObject {
	newObj := &openapi.ChoiceObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &choiceObject{obj: newObj}
	newLibObj.setDefault()
	obj.choiceObjectSlice = append(obj.choiceObjectSlice, newLibObj)
	return newLibObj
}

func (obj *prefixConfigChoiceObjectIter) Append(items ...ChoiceObject) PrefixConfigChoiceObjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.choiceObjectSlice = append(obj.choiceObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigChoiceObjectIter) Set(index int, newObj ChoiceObject) PrefixConfigChoiceObjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.choiceObjectSlice[index] = newObj
	return obj
}
func (obj *prefixConfigChoiceObjectIter) Clear() PrefixConfigChoiceObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.ChoiceObject{}
		obj.choiceObjectSlice = []ChoiceObject{}
	}
	return obj
}
func (obj *prefixConfigChoiceObjectIter) clearHolderSlice() PrefixConfigChoiceObjectIter {
	if len(obj.choiceObjectSlice) > 0 {
		obj.choiceObjectSlice = []ChoiceObject{}
	}
	return obj
}
func (obj *prefixConfigChoiceObjectIter) appendHolderSlice(item ChoiceObject) PrefixConfigChoiceObjectIter {
	obj.choiceObjectSlice = append(obj.choiceObjectSlice, item)
	return obj
}

// description is TBD
// RequiredChoiceObject returns a RequiredChoiceParent
func (obj *prefixConfig) RequiredChoiceObject() RequiredChoiceParent {
	if obj.obj.RequiredChoiceObject == nil {
		obj.obj.RequiredChoiceObject = NewRequiredChoiceParent().msg()
	}
	if obj.requiredChoiceObjectHolder == nil {
		obj.requiredChoiceObjectHolder = &requiredChoiceParent{obj: obj.obj.RequiredChoiceObject}
	}
	return obj.requiredChoiceObjectHolder
}

// description is TBD
// RequiredChoiceObject returns a RequiredChoiceParent
func (obj *prefixConfig) HasRequiredChoiceObject() bool {
	return obj.obj.RequiredChoiceObject != nil
}

// description is TBD
// SetRequiredChoiceObject sets the RequiredChoiceParent value in the PrefixConfig object
func (obj *prefixConfig) SetRequiredChoiceObject(value RequiredChoiceParent) PrefixConfig {

	obj.requiredChoiceObjectHolder = nil
	obj.obj.RequiredChoiceObject = value.msg()

	return obj
}

// A list of objects with choice and properties
// G1 returns a []GObject
func (obj *prefixConfig) G1() PrefixConfigGObjectIter {
	if len(obj.obj.G1) == 0 {
		obj.obj.G1 = []*openapi.GObject{}
	}
	if obj.g1Holder == nil {
		obj.g1Holder = newPrefixConfigGObjectIter(&obj.obj.G1).setMsg(obj)
	}
	return obj.g1Holder
}

// A list of objects with choice and properties
// G2 returns a []GObject
func (obj *prefixConfig) G2() PrefixConfigGObjectIter {
	if len(obj.obj.G2) == 0 {
		obj.obj.G2 = []*openapi.GObject{}
	}
	if obj.g2Holder == nil {
		obj.g2Holder = newPrefixConfigGObjectIter(&obj.obj.G2).setMsg(obj)
	}
	return obj.g2Holder
}

// int32 type
// Int32Param returns a int32
func (obj *prefixConfig) Int32Param() int32 {

	return *obj.obj.Int32Param

}

// int32 type
// Int32Param returns a int32
func (obj *prefixConfig) HasInt32Param() bool {
	return obj.obj.Int32Param != nil
}

// int32 type
// SetInt32Param sets the int32 value in the PrefixConfig object
func (obj *prefixConfig) SetInt32Param(value int32) PrefixConfig {

	obj.obj.Int32Param = &value
	return obj
}

// int32 type list
// Int32ListParam returns a []int32
func (obj *prefixConfig) Int32ListParam() []int32 {
	if obj.obj.Int32ListParam == nil {
		obj.obj.Int32ListParam = make([]int32, 0)
	}
	return obj.obj.Int32ListParam
}

// int32 type list
// SetInt32ListParam sets the []int32 value in the PrefixConfig object
func (obj *prefixConfig) SetInt32ListParam(value []int32) PrefixConfig {

	if obj.obj.Int32ListParam == nil {
		obj.obj.Int32ListParam = make([]int32, 0)
	}
	obj.obj.Int32ListParam = value

	return obj
}

// uint32 type
// Uint32Param returns a uint32
func (obj *prefixConfig) Uint32Param() uint32 {

	return *obj.obj.Uint32Param

}

// uint32 type
// Uint32Param returns a uint32
func (obj *prefixConfig) HasUint32Param() bool {
	return obj.obj.Uint32Param != nil
}

// uint32 type
// SetUint32Param sets the uint32 value in the PrefixConfig object
func (obj *prefixConfig) SetUint32Param(value uint32) PrefixConfig {

	obj.obj.Uint32Param = &value
	return obj
}

// uint32 type list
// Uint32ListParam returns a []uint32
func (obj *prefixConfig) Uint32ListParam() []uint32 {
	if obj.obj.Uint32ListParam == nil {
		obj.obj.Uint32ListParam = make([]uint32, 0)
	}
	return obj.obj.Uint32ListParam
}

// uint32 type list
// SetUint32ListParam sets the []uint32 value in the PrefixConfig object
func (obj *prefixConfig) SetUint32ListParam(value []uint32) PrefixConfig {

	if obj.obj.Uint32ListParam == nil {
		obj.obj.Uint32ListParam = make([]uint32, 0)
	}
	obj.obj.Uint32ListParam = value

	return obj
}

// uint64 type
// Uint64Param returns a uint64
func (obj *prefixConfig) Uint64Param() uint64 {

	return *obj.obj.Uint64Param

}

// uint64 type
// Uint64Param returns a uint64
func (obj *prefixConfig) HasUint64Param() bool {
	return obj.obj.Uint64Param != nil
}

// uint64 type
// SetUint64Param sets the uint64 value in the PrefixConfig object
func (obj *prefixConfig) SetUint64Param(value uint64) PrefixConfig {

	obj.obj.Uint64Param = &value
	return obj
}

// uint64 type list
// Uint64ListParam returns a []uint64
func (obj *prefixConfig) Uint64ListParam() []uint64 {
	if obj.obj.Uint64ListParam == nil {
		obj.obj.Uint64ListParam = make([]uint64, 0)
	}
	return obj.obj.Uint64ListParam
}

// uint64 type list
// SetUint64ListParam sets the []uint64 value in the PrefixConfig object
func (obj *prefixConfig) SetUint64ListParam(value []uint64) PrefixConfig {

	if obj.obj.Uint64ListParam == nil {
		obj.obj.Uint64ListParam = make([]uint64, 0)
	}
	obj.obj.Uint64ListParam = value

	return obj
}

// should automatically set type to int32
// AutoInt32Param returns a int32
func (obj *prefixConfig) AutoInt32Param() int32 {

	return *obj.obj.AutoInt32Param

}

// should automatically set type to int32
// AutoInt32Param returns a int32
func (obj *prefixConfig) HasAutoInt32Param() bool {
	return obj.obj.AutoInt32Param != nil
}

// should automatically set type to int32
// SetAutoInt32Param sets the int32 value in the PrefixConfig object
func (obj *prefixConfig) SetAutoInt32Param(value int32) PrefixConfig {

	obj.obj.AutoInt32Param = &value
	return obj
}

// should automatically set type to []int32
// AutoInt32ListParam returns a []int32
func (obj *prefixConfig) AutoInt32ListParam() []int32 {
	if obj.obj.AutoInt32ListParam == nil {
		obj.obj.AutoInt32ListParam = make([]int32, 0)
	}
	return obj.obj.AutoInt32ListParam
}

// should automatically set type to []int32
// SetAutoInt32ListParam sets the []int32 value in the PrefixConfig object
func (obj *prefixConfig) SetAutoInt32ListParam(value []int32) PrefixConfig {

	if obj.obj.AutoInt32ListParam == nil {
		obj.obj.AutoInt32ListParam = make([]int32, 0)
	}
	obj.obj.AutoInt32ListParam = value

	return obj
}

// description is TBD
// ChoiceTest returns a ChoiceTestObj
func (obj *prefixConfig) ChoiceTest() ChoiceTestObj {
	if obj.obj.ChoiceTest == nil {
		obj.obj.ChoiceTest = NewChoiceTestObj().msg()
	}
	if obj.choiceTestHolder == nil {
		obj.choiceTestHolder = &choiceTestObj{obj: obj.obj.ChoiceTest}
	}
	return obj.choiceTestHolder
}

// description is TBD
// ChoiceTest returns a ChoiceTestObj
func (obj *prefixConfig) HasChoiceTest() bool {
	return obj.obj.ChoiceTest != nil
}

// description is TBD
// SetChoiceTest sets the ChoiceTestObj value in the PrefixConfig object
func (obj *prefixConfig) SetChoiceTest(value ChoiceTestObj) PrefixConfig {

	obj.choiceTestHolder = nil
	obj.obj.ChoiceTest = value.msg()

	return obj
}

// description is TBD
// SignedIntegerPattern returns a SignedIntegerPattern
func (obj *prefixConfig) SignedIntegerPattern() SignedIntegerPattern {
	if obj.obj.SignedIntegerPattern == nil {
		obj.obj.SignedIntegerPattern = NewSignedIntegerPattern().msg()
	}
	if obj.signedIntegerPatternHolder == nil {
		obj.signedIntegerPatternHolder = &signedIntegerPattern{obj: obj.obj.SignedIntegerPattern}
	}
	return obj.signedIntegerPatternHolder
}

// description is TBD
// SignedIntegerPattern returns a SignedIntegerPattern
func (obj *prefixConfig) HasSignedIntegerPattern() bool {
	return obj.obj.SignedIntegerPattern != nil
}

// description is TBD
// SetSignedIntegerPattern sets the SignedIntegerPattern value in the PrefixConfig object
func (obj *prefixConfig) SetSignedIntegerPattern(value SignedIntegerPattern) PrefixConfig {

	obj.signedIntegerPatternHolder = nil
	obj.obj.SignedIntegerPattern = value.msg()

	return obj
}

// description is TBD
// OidPattern returns a OidPattern
func (obj *prefixConfig) OidPattern() OidPattern {
	if obj.obj.OidPattern == nil {
		obj.obj.OidPattern = NewOidPattern().msg()
	}
	if obj.oidPatternHolder == nil {
		obj.oidPatternHolder = &oidPattern{obj: obj.obj.OidPattern}
	}
	return obj.oidPatternHolder
}

// description is TBD
// OidPattern returns a OidPattern
func (obj *prefixConfig) HasOidPattern() bool {
	return obj.obj.OidPattern != nil
}

// description is TBD
// SetOidPattern sets the OidPattern value in the PrefixConfig object
func (obj *prefixConfig) SetOidPattern(value OidPattern) PrefixConfig {

	obj.oidPatternHolder = nil
	obj.obj.OidPattern = value.msg()

	return obj
}

// description is TBD
// ChoiceDefault returns a ChoiceObject
func (obj *prefixConfig) ChoiceDefault() ChoiceObject {
	if obj.obj.ChoiceDefault == nil {
		obj.obj.ChoiceDefault = NewChoiceObject().msg()
	}
	if obj.choiceDefaultHolder == nil {
		obj.choiceDefaultHolder = &choiceObject{obj: obj.obj.ChoiceDefault}
	}
	return obj.choiceDefaultHolder
}

// description is TBD
// ChoiceDefault returns a ChoiceObject
func (obj *prefixConfig) HasChoiceDefault() bool {
	return obj.obj.ChoiceDefault != nil
}

// description is TBD
// SetChoiceDefault sets the ChoiceObject value in the PrefixConfig object
func (obj *prefixConfig) SetChoiceDefault(value ChoiceObject) PrefixConfig {

	obj.choiceDefaultHolder = nil
	obj.obj.ChoiceDefault = value.msg()

	return obj
}

// description is TBD
// ChoiceRequiredDefault returns a ChoiceRequiredAndDefault
func (obj *prefixConfig) ChoiceRequiredDefault() ChoiceRequiredAndDefault {
	if obj.obj.ChoiceRequiredDefault == nil {
		obj.obj.ChoiceRequiredDefault = NewChoiceRequiredAndDefault().msg()
	}
	if obj.choiceRequiredDefaultHolder == nil {
		obj.choiceRequiredDefaultHolder = &choiceRequiredAndDefault{obj: obj.obj.ChoiceRequiredDefault}
	}
	return obj.choiceRequiredDefaultHolder
}

// description is TBD
// ChoiceRequiredDefault returns a ChoiceRequiredAndDefault
func (obj *prefixConfig) HasChoiceRequiredDefault() bool {
	return obj.obj.ChoiceRequiredDefault != nil
}

// description is TBD
// SetChoiceRequiredDefault sets the ChoiceRequiredAndDefault value in the PrefixConfig object
func (obj *prefixConfig) SetChoiceRequiredDefault(value ChoiceRequiredAndDefault) PrefixConfig {

	obj.choiceRequiredDefaultHolder = nil
	obj.obj.ChoiceRequiredDefault = value.msg()

	return obj
}

// description is TBD
// AutoPattern returns a AutoPattern
func (obj *prefixConfig) AutoPattern() AutoPattern {
	if obj.obj.AutoPattern == nil {
		obj.obj.AutoPattern = NewAutoPattern().msg()
	}
	if obj.autoPatternHolder == nil {
		obj.autoPatternHolder = &autoPattern{obj: obj.obj.AutoPattern}
	}
	return obj.autoPatternHolder
}

// description is TBD
// AutoPattern returns a AutoPattern
func (obj *prefixConfig) HasAutoPattern() bool {
	return obj.obj.AutoPattern != nil
}

// description is TBD
// SetAutoPattern sets the AutoPattern value in the PrefixConfig object
func (obj *prefixConfig) SetAutoPattern(value AutoPattern) PrefixConfig {

	obj.autoPatternHolder = nil
	obj.obj.AutoPattern = value.msg()

	return obj
}

// description is TBD
// AutoPatternDefault returns a AutoPatternDefault
func (obj *prefixConfig) AutoPatternDefault() AutoPatternDefault {
	if obj.obj.AutoPatternDefault == nil {
		obj.obj.AutoPatternDefault = NewAutoPatternDefault().msg()
	}
	if obj.autoPatternDefaultHolder == nil {
		obj.autoPatternDefaultHolder = &autoPatternDefault{obj: obj.obj.AutoPatternDefault}
	}
	return obj.autoPatternDefaultHolder
}

// description is TBD
// AutoPatternDefault returns a AutoPatternDefault
func (obj *prefixConfig) HasAutoPatternDefault() bool {
	return obj.obj.AutoPatternDefault != nil
}

// description is TBD
// SetAutoPatternDefault sets the AutoPatternDefault value in the PrefixConfig object
func (obj *prefixConfig) SetAutoPatternDefault(value AutoPatternDefault) PrefixConfig {

	obj.autoPatternDefaultHolder = nil
	obj.obj.AutoPatternDefault = value.msg()

	return obj
}

// description is TBD
// NameEndingWithNumber234 returns a string
func (obj *prefixConfig) NameEndingWithNumber234() string {

	return *obj.obj.NameEndingWithNumber_234

}

// description is TBD
// NameEndingWithNumber234 returns a string
func (obj *prefixConfig) HasNameEndingWithNumber234() bool {
	return obj.obj.NameEndingWithNumber_234 != nil
}

// description is TBD
// SetNameEndingWithNumber234 sets the string value in the PrefixConfig object
func (obj *prefixConfig) SetNameEndingWithNumber234(value string) PrefixConfig {

	obj.obj.NameEndingWithNumber_234 = &value
	return obj
}

func (obj *prefixConfig) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// RequiredObject is required
	if obj.obj.RequiredObject == nil {
		vObj.validationErrors = append(vObj.validationErrors, "RequiredObject is required field on interface PrefixConfig")
	}

	if obj.obj.RequiredObject != nil {

		obj.RequiredObject().validateObj(vObj, set_default)
	}

	if obj.obj.OptionalObject != nil {

		obj.OptionalObject().validateObj(vObj, set_default)
	}

	// Space_1 is deprecated
	if obj.obj.Space_1 != nil {
		obj.addWarnings("Space_1 property in schema PrefixConfig is deprecated, Information TBD")
	}

	if obj.obj.FullDuplex_100Mb != nil {

		if *obj.obj.FullDuplex_100Mb < -10 || *obj.obj.FullDuplex_100Mb > 4261412864 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-10 <= PrefixConfig.FullDuplex_100Mb <= 4261412864 but Got %d", *obj.obj.FullDuplex_100Mb))
		}

	}

	if obj.obj.Response.Number() == 3 {
		obj.addWarnings("STATUS_404 enum in property Response is deprecated, new code will be coming soon")
	}

	if obj.obj.Response.Number() == 4 {
		obj.addWarnings("STATUS_500 enum in property Response is under review, 500 can change to other values")
	}

	// A is required
	if obj.obj.A == nil {
		vObj.validationErrors = append(vObj.validationErrors, "A is required field on interface PrefixConfig")
	}

	// A is under_review
	if obj.obj.A != nil {
		obj.addWarnings("A property in schema PrefixConfig is under review, Information TBD")
	}

	// B is required
	if obj.obj.B == nil {
		vObj.validationErrors = append(vObj.validationErrors, "B is required field on interface PrefixConfig")
	}

	// C is required
	if obj.obj.C == nil {
		vObj.validationErrors = append(vObj.validationErrors, "C is required field on interface PrefixConfig")
	}

	// DValues is deprecated
	if obj.obj.DValues != nil {
		obj.addWarnings("DValues property in schema PrefixConfig is deprecated, Information TBD")
	}

	if obj.obj.E != nil {
		obj.addWarnings("E property in schema PrefixConfig is deprecated, Information TBD")
		obj.E().validateObj(vObj, set_default)
	}

	if obj.obj.F != nil {

		obj.F().validateObj(vObj, set_default)
	}

	if len(obj.obj.G) != 0 {

		if set_default {
			obj.G().clearHolderSlice()
			for _, item := range obj.obj.G {
				obj.G().appendHolderSlice(&gObject{obj: item})
			}
		}
		for _, item := range obj.G().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.J) != 0 {

		if set_default {
			obj.J().clearHolderSlice()
			for _, item := range obj.obj.J {
				obj.J().appendHolderSlice(&jObject{obj: item})
			}
		}
		for _, item := range obj.J().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.K != nil {

		obj.K().validateObj(vObj, set_default)
	}

	if obj.obj.L != nil {

		obj.L().validateObj(vObj, set_default)
	}

	if obj.obj.Level != nil {

		obj.Level().validateObj(vObj, set_default)
	}

	if obj.obj.Mandatory != nil {

		obj.Mandatory().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv4Pattern != nil {

		obj.Ipv4Pattern().validateObj(vObj, set_default)
	}

	if obj.obj.Ipv6Pattern != nil {

		obj.Ipv6Pattern().validateObj(vObj, set_default)
	}

	if obj.obj.MacPattern != nil {

		obj.MacPattern().validateObj(vObj, set_default)
	}

	if obj.obj.IntegerPattern != nil {

		obj.IntegerPattern().validateObj(vObj, set_default)
	}

	if obj.obj.ChecksumPattern != nil {

		obj.ChecksumPattern().validateObj(vObj, set_default)
	}

	if obj.obj.Case != nil {

		obj.Case().validateObj(vObj, set_default)
	}

	if obj.obj.MObject != nil {

		obj.MObject().validateObj(vObj, set_default)
	}

	if obj.obj.Integer64List != nil {

		for _, item := range obj.obj.Integer64List {
			if item < -12 || item > 4261412864 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("-12 <= PrefixConfig.Integer64List <= 4261412864 but Got %d", item))
			}

		}

	}

	if obj.obj.HeaderChecksum != nil {

		obj.HeaderChecksum().validateObj(vObj, set_default)
	}

	if obj.obj.StrLen != nil {
		obj.addWarnings("StrLen property in schema PrefixConfig is under review, Information TBD")
		if len(*obj.obj.StrLen) < 3 || len(*obj.obj.StrLen) > 6 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"3 <= length of PrefixConfig.StrLen <= 6 but Got %d",
					len(*obj.obj.StrLen)))
		}

	}

	if obj.obj.HexSlice != nil {
		obj.addWarnings("HexSlice property in schema PrefixConfig is under review, Information TBD")

		err := obj.validateHexSlice(obj.HexSlice())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PrefixConfig.HexSlice"))
		}

	}

	if obj.obj.AutoFieldTest != nil {

		obj.AutoFieldTest().validateObj(vObj, set_default)
	}

	if len(obj.obj.WList) != 0 {

		if set_default {
			obj.WList().clearHolderSlice()
			for _, item := range obj.obj.WList {
				obj.WList().appendHolderSlice(&wObject{obj: item})
			}
		}
		for _, item := range obj.WList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.XList) != 0 {

		if set_default {
			obj.XList().clearHolderSlice()
			for _, item := range obj.obj.XList {
				obj.XList().appendHolderSlice(&zObject{obj: item})
			}
		}
		for _, item := range obj.XList().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.ZObject != nil {

		obj.ZObject().validateObj(vObj, set_default)
	}

	if obj.obj.YObject != nil {

		obj.YObject().validateObj(vObj, set_default)
	}

	if len(obj.obj.ChoiceObject) != 0 {

		if set_default {
			obj.ChoiceObject().clearHolderSlice()
			for _, item := range obj.obj.ChoiceObject {
				obj.ChoiceObject().appendHolderSlice(&choiceObject{obj: item})
			}
		}
		for _, item := range obj.ChoiceObject().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.RequiredChoiceObject != nil {

		obj.RequiredChoiceObject().validateObj(vObj, set_default)
	}

	if len(obj.obj.G1) != 0 {

		if set_default {
			obj.G1().clearHolderSlice()
			for _, item := range obj.obj.G1 {
				obj.G1().appendHolderSlice(&gObject{obj: item})
			}
		}
		for _, item := range obj.G1().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.G2) != 0 {

		if set_default {
			obj.G2().clearHolderSlice()
			for _, item := range obj.obj.G2 {
				obj.G2().appendHolderSlice(&gObject{obj: item})
			}
		}
		for _, item := range obj.G2().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.Int32ListParam != nil {

		for _, item := range obj.obj.Int32ListParam {
			if item < -23456 || item > 23456 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("-23456 <= PrefixConfig.Int32ListParam <= 23456 but Got %d", item))
			}

		}

	}

	if obj.obj.Uint32ListParam != nil {

		for _, item := range obj.obj.Uint32ListParam {
			if item > 4294967293 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= PrefixConfig.Uint32ListParam <= 4294967293 but Got %d", item))
			}

		}

	}

	if obj.obj.AutoInt32Param != nil {

		if *obj.obj.AutoInt32Param < 64 || *obj.obj.AutoInt32Param > 9000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("64 <= PrefixConfig.AutoInt32Param <= 9000 but Got %d", *obj.obj.AutoInt32Param))
		}

	}

	if obj.obj.AutoInt32ListParam != nil {

		for _, item := range obj.obj.AutoInt32ListParam {
			if item < 64 || item > 9000 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("64 <= PrefixConfig.AutoInt32ListParam <= 9000 but Got %d", item))
			}

		}

	}

	if obj.obj.ChoiceTest != nil {

		obj.ChoiceTest().validateObj(vObj, set_default)
	}

	if obj.obj.SignedIntegerPattern != nil {

		obj.SignedIntegerPattern().validateObj(vObj, set_default)
	}

	if obj.obj.OidPattern != nil {

		obj.OidPattern().validateObj(vObj, set_default)
	}

	if obj.obj.ChoiceDefault != nil {

		obj.ChoiceDefault().validateObj(vObj, set_default)
	}

	if obj.obj.ChoiceRequiredDefault != nil {

		obj.ChoiceRequiredDefault().validateObj(vObj, set_default)
	}

	if obj.obj.AutoPattern != nil {

		obj.AutoPattern().validateObj(vObj, set_default)
	}

	if obj.obj.AutoPatternDefault != nil {

		obj.AutoPatternDefault().validateObj(vObj, set_default)
	}

	if obj.obj.NameEndingWithNumber_234 != nil {

		err := obj.validateIpv4(obj.NameEndingWithNumber234())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PrefixConfig.NameEndingWithNumber234"))
		}

	}

}

func (obj *prefixConfig) setDefault() {
	if obj.obj.Response == nil {
		obj.SetResponse(PrefixConfigResponse.STATUS_200)

	}
	if obj.obj.H == nil {
		obj.SetH(true)
	}

}
