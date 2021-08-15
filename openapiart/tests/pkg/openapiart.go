// OpenAPIArt Test API 0.0.1
// MIT License
//
// Copyright (c) 2021 https://github.com/open-traffic-generator/openapiart
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package openapiart

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type ApiTransportEnum string

var ApiTransport = struct {
	GRPC ApiTransportEnum
	HTTP ApiTransportEnum
}{
	GRPC: "grpc",
	HTTP: "http",
}

type api struct {
	transport          ApiTransportEnum
	grpcLocation       string
	grpcRequestTimeout time.Duration
}

type Api interface {
	SetTransport(value ApiTransportEnum) *api
	SetGrpcLocation(value string) *api
	SetGrpcRequestTimeout(value time.Duration) *api
}

// Transport returns the active transport
func (api *api) Transport() string {
	return string(api.transport)
}

// SetTransport sets the active type of transport to be used
func (api *api) SetTransport(value ApiTransportEnum) *api {
	api.transport = value
	return api
}

func (api *api) GrpcLocation() string {
	return api.grpcLocation
}

// SetGrpcLocation
func (api *api) SetGrpcLocation(value string) *api {
	api.grpcLocation = value
	return api
}

func (api *api) GrpcRequestTimeout() time.Duration {
	return api.grpcRequestTimeout
}

// SetGrpcRequestTimeout contains the timeout value in seconds for a grpc request
func (api *api) SetGrpcRequestTimeout(value int) *api {
	api.grpcRequestTimeout = time.Duration(value) * time.Second
	return api
}

// All methods that perform validation will add errors here
// All api rpcs MUST call Validate
var validation []string

func Validate() {
	if len(validation) > 0 {
		for _, item := range validation {
			fmt.Println(item)
		}
		panic("validation errors")
	}
}

type openapiartApi struct {
	api
	grpcClient sanity.OpenapiClient
}

// grpcConnect builds up a grpc connection
func (api *openapiartApi) grpcConnect() error {
	if api.grpcClient == nil {
		conn, err := grpc.Dial(api.grpcLocation, grpc.WithInsecure())
		if err != nil {
			return err
		}
		api.grpcClient = sanity.NewOpenapiClient(conn)
	}
	return nil
}

// NewApi returns a new instance of the top level interface hierarchy
func NewApi() *openapiartApi {
	api := openapiartApi{}
	api.transport = ApiTransport.GRPC
	api.grpcLocation = "127.0.0.1:5050"
	api.grpcRequestTimeout = 10 * time.Second
	api.grpcClient = nil
	return &api
}

type OpenapiartApi interface {
	Api
	NewPrefixConfig() PrefixConfig
	SetConfig(prefixConfig PrefixConfig) error
}

func (api *openapiartApi) NewPrefixConfig() PrefixConfig {
	return &prefixConfig{obj: &sanity.PrefixConfig{}}
}

func (api *openapiartApi) SetConfig(prefixConfig PrefixConfig) error {
	if err := api.grpcConnect(); err != nil {
		return err
	}
	request := sanity.SetConfigRequest{PrefixConfig: prefixConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpcRequestTimeout)
	defer cancelFunc()
	client, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		return err
	}
	resp, _ := client.Recv()
	if resp.GetStatusCode_200() == nil {
		return fmt.Errorf("fail")
	}
	return nil
}

type prefixConfig struct {
	obj *sanity.PrefixConfig
}

func (obj *prefixConfig) msg() *sanity.PrefixConfig {
	return obj.obj
}

func (obj *prefixConfig) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *prefixConfig) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PrefixConfig interface {
	msg() *sanity.PrefixConfig
	Yaml() string
	Json() string
	A() string
	SetA(value string) PrefixConfig
	B() float32
	SetB(value float32) PrefixConfig
	C() int32
	SetC(value int32) PrefixConfig
	E() EObject
	F() FObject
	H() bool
	SetH(value bool) PrefixConfig
	I() []byte
	SetI(value []byte) PrefixConfig
	K() KObject
	L() LObject
	Level() LevelOne
	Mandatory() Mandate
	Ipv4Pattern() Ipv4Pattern
	Ipv6Pattern() Ipv6Pattern
	MacPattern() MacPattern
	IntegerPattern() IntegerPattern
	ChecksumPattern() ChecksumPattern
	Name() string
	SetName(value string) PrefixConfig
}

func (obj *prefixConfig) A() string {
	return obj.obj.A
}

func (obj *prefixConfig) SetA(value string) PrefixConfig {
	obj.obj.A = value
	return obj
}

func (obj *prefixConfig) B() float32 {
	return obj.obj.B
}

func (obj *prefixConfig) SetB(value float32) PrefixConfig {
	obj.obj.B = value
	return obj
}

func (obj *prefixConfig) C() int32 {
	return obj.obj.C
}

func (obj *prefixConfig) SetC(value int32) PrefixConfig {
	obj.obj.C = value
	return obj
}

func (obj *prefixConfig) E() EObject {
	if obj.obj.E == nil {
		obj.obj.E = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.E}

}

func (obj *prefixConfig) F() FObject {
	if obj.obj.F == nil {
		obj.obj.F = &sanity.FObject{}
	}
	return &fObject{obj: obj.obj.F}

}

func (obj *prefixConfig) H() bool {
	return *obj.obj.H
}

func (obj *prefixConfig) SetH(value bool) PrefixConfig {
	obj.obj.H = &value
	return obj
}

func (obj *prefixConfig) I() []byte {
	return obj.obj.I
}

func (obj *prefixConfig) SetI(value []byte) PrefixConfig {
	obj.obj.I = value
	return obj
}

func (obj *prefixConfig) K() KObject {
	if obj.obj.K == nil {
		obj.obj.K = &sanity.KObject{}
	}
	return &kObject{obj: obj.obj.K}

}

func (obj *prefixConfig) L() LObject {
	if obj.obj.L == nil {
		obj.obj.L = &sanity.LObject{}
	}
	return &lObject{obj: obj.obj.L}

}

func (obj *prefixConfig) Level() LevelOne {
	if obj.obj.Level == nil {
		obj.obj.Level = &sanity.LevelOne{}
	}
	return &levelOne{obj: obj.obj.Level}

}

func (obj *prefixConfig) Mandatory() Mandate {
	if obj.obj.Mandatory == nil {
		obj.obj.Mandatory = &sanity.Mandate{}
	}
	return &mandate{obj: obj.obj.Mandatory}

}

func (obj *prefixConfig) Ipv4Pattern() Ipv4Pattern {
	if obj.obj.Ipv4Pattern == nil {
		obj.obj.Ipv4Pattern = &sanity.Ipv4Pattern{}
	}
	return &ipv4Pattern{obj: obj.obj.Ipv4Pattern}

}

func (obj *prefixConfig) Ipv6Pattern() Ipv6Pattern {
	if obj.obj.Ipv6Pattern == nil {
		obj.obj.Ipv6Pattern = &sanity.Ipv6Pattern{}
	}
	return &ipv6Pattern{obj: obj.obj.Ipv6Pattern}

}

func (obj *prefixConfig) MacPattern() MacPattern {
	if obj.obj.MacPattern == nil {
		obj.obj.MacPattern = &sanity.MacPattern{}
	}
	return &macPattern{obj: obj.obj.MacPattern}

}

func (obj *prefixConfig) IntegerPattern() IntegerPattern {
	if obj.obj.IntegerPattern == nil {
		obj.obj.IntegerPattern = &sanity.IntegerPattern{}
	}
	return &integerPattern{obj: obj.obj.IntegerPattern}

}

func (obj *prefixConfig) ChecksumPattern() ChecksumPattern {
	if obj.obj.ChecksumPattern == nil {
		obj.obj.ChecksumPattern = &sanity.ChecksumPattern{}
	}
	return &checksumPattern{obj: obj.obj.ChecksumPattern}

}

func (obj *prefixConfig) Name() string {
	return *obj.obj.Name
}

func (obj *prefixConfig) SetName(value string) PrefixConfig {
	obj.obj.Name = &value
	return obj
}

type eObject struct {
	obj *sanity.EObject
}

func (obj *eObject) msg() *sanity.EObject {
	return obj.obj
}

func (obj *eObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *eObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type EObject interface {
	msg() *sanity.EObject
	Yaml() string
	Json() string
	EA() float32
	SetEA(value float32) EObject
	EB() float64
	SetEB(value float64) EObject
	Name() string
	SetName(value string) EObject
	MParam1() string
	SetMParam1(value string) EObject
	MParam2() string
	SetMParam2(value string) EObject
}

func (obj *eObject) EA() float32 {
	return obj.obj.EA
}

func (obj *eObject) SetEA(value float32) EObject {
	obj.obj.EA = value
	return obj
}

func (obj *eObject) EB() float64 {
	return obj.obj.EB
}

func (obj *eObject) SetEB(value float64) EObject {
	obj.obj.EB = value
	return obj
}

func (obj *eObject) Name() string {
	return *obj.obj.Name
}

func (obj *eObject) SetName(value string) EObject {
	obj.obj.Name = &value
	return obj
}

func (obj *eObject) MParam1() string {
	return *obj.obj.MParam1
}

func (obj *eObject) SetMParam1(value string) EObject {
	obj.obj.MParam1 = &value
	return obj
}

func (obj *eObject) MParam2() string {
	return *obj.obj.MParam2
}

func (obj *eObject) SetMParam2(value string) EObject {
	obj.obj.MParam2 = &value
	return obj
}

type fObject struct {
	obj *sanity.FObject
}

func (obj *fObject) msg() *sanity.FObject {
	return obj.obj
}

func (obj *fObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *fObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type FObject interface {
	msg() *sanity.FObject
	Yaml() string
	Json() string
	FA() string
	SetFA(value string) FObject
	FB() float64
	SetFB(value float64) FObject
}

func (obj *fObject) FA() string {
	return *obj.obj.FA
}

func (obj *fObject) SetFA(value string) FObject {
	obj.obj.FA = &value
	return obj
}

func (obj *fObject) FB() float64 {
	return *obj.obj.FB
}

func (obj *fObject) SetFB(value float64) FObject {
	obj.obj.FB = &value
	return obj
}

type gObject struct {
	obj *sanity.GObject
}

func (obj *gObject) msg() *sanity.GObject {
	return obj.obj
}

func (obj *gObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *gObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type GObject interface {
	msg() *sanity.GObject
	Yaml() string
	Json() string
	GA() string
	SetGA(value string) GObject
	GB() int32
	SetGB(value int32) GObject
	GC() float32
	SetGC(value float32) GObject
	GD() string
	SetGD(value string) GObject
	GE() float64
	SetGE(value float64) GObject
	Name() string
	SetName(value string) GObject
}

func (obj *gObject) GA() string {
	return *obj.obj.GA
}

func (obj *gObject) SetGA(value string) GObject {
	obj.obj.GA = &value
	return obj
}

func (obj *gObject) GB() int32 {
	return *obj.obj.GB
}

func (obj *gObject) SetGB(value int32) GObject {
	obj.obj.GB = &value
	return obj
}

func (obj *gObject) GC() float32 {
	return *obj.obj.GC
}

func (obj *gObject) SetGC(value float32) GObject {
	obj.obj.GC = &value
	return obj
}

func (obj *gObject) GD() string {
	return *obj.obj.GD
}

func (obj *gObject) SetGD(value string) GObject {
	obj.obj.GD = &value
	return obj
}

func (obj *gObject) GE() float64 {
	return *obj.obj.GE
}

func (obj *gObject) SetGE(value float64) GObject {
	obj.obj.GE = &value
	return obj
}

func (obj *gObject) Name() string {
	return *obj.obj.Name
}

func (obj *gObject) SetName(value string) GObject {
	obj.obj.Name = &value
	return obj
}

type jObject struct {
	obj *sanity.JObject
}

func (obj *jObject) msg() *sanity.JObject {
	return obj.obj
}

func (obj *jObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *jObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type JObject interface {
	msg() *sanity.JObject
	Yaml() string
	Json() string
	JA() EObject
	JB() FObject
}

func (obj *jObject) JA() EObject {
	if obj.obj.JA == nil {
		obj.obj.JA = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.JA}

}

func (obj *jObject) JB() FObject {
	if obj.obj.JB == nil {
		obj.obj.JB = &sanity.FObject{}
	}
	return &fObject{obj: obj.obj.JB}

}

type kObject struct {
	obj *sanity.KObject
}

func (obj *kObject) msg() *sanity.KObject {
	return obj.obj
}

func (obj *kObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *kObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type KObject interface {
	msg() *sanity.KObject
	Yaml() string
	Json() string
	EObject() EObject
	FObject() FObject
}

func (obj *kObject) EObject() EObject {
	if obj.obj.EObject == nil {
		obj.obj.EObject = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.EObject}

}

func (obj *kObject) FObject() FObject {
	if obj.obj.FObject == nil {
		obj.obj.FObject = &sanity.FObject{}
	}
	return &fObject{obj: obj.obj.FObject}

}

type lObject struct {
	obj *sanity.LObject
}

func (obj *lObject) msg() *sanity.LObject {
	return obj.obj
}

func (obj *lObject) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lObject) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LObject interface {
	msg() *sanity.LObject
	Yaml() string
	Json() string
	String_() string
	SetString_(value string) LObject
	Integer() int32
	SetInteger(value int32) LObject
	Float() float32
	SetFloat(value float32) LObject
	Double() float64
	SetDouble(value float64) LObject
	Mac() string
	SetMac(value string) LObject
	Ipv4() string
	SetIpv4(value string) LObject
	Ipv6() string
	SetIpv6(value string) LObject
	Hex() string
	SetHex(value string) LObject
}

func (obj *lObject) String_() string {
	return *obj.obj.String_
}

func (obj *lObject) SetString_(value string) LObject {
	obj.obj.String_ = &value
	return obj
}

func (obj *lObject) Integer() int32 {
	return *obj.obj.Integer
}

func (obj *lObject) SetInteger(value int32) LObject {
	obj.obj.Integer = &value
	return obj
}

func (obj *lObject) Float() float32 {
	return *obj.obj.Float
}

func (obj *lObject) SetFloat(value float32) LObject {
	obj.obj.Float = &value
	return obj
}

func (obj *lObject) Double() float64 {
	return *obj.obj.Double
}

func (obj *lObject) SetDouble(value float64) LObject {
	obj.obj.Double = &value
	return obj
}

func (obj *lObject) Mac() string {
	return *obj.obj.Mac
}

func (obj *lObject) SetMac(value string) LObject {
	obj.obj.Mac = &value
	return obj
}

func (obj *lObject) Ipv4() string {
	return *obj.obj.Ipv4
}

func (obj *lObject) SetIpv4(value string) LObject {
	obj.obj.Ipv4 = &value
	return obj
}

func (obj *lObject) Ipv6() string {
	return *obj.obj.Ipv6
}

func (obj *lObject) SetIpv6(value string) LObject {
	obj.obj.Ipv6 = &value
	return obj
}

func (obj *lObject) Hex() string {
	return *obj.obj.Hex
}

func (obj *lObject) SetHex(value string) LObject {
	obj.obj.Hex = &value
	return obj
}

type levelOne struct {
	obj *sanity.LevelOne
}

func (obj *levelOne) msg() *sanity.LevelOne {
	return obj.obj
}

func (obj *levelOne) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelOne) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LevelOne interface {
	msg() *sanity.LevelOne
	Yaml() string
	Json() string
	L1P1() LevelTwo
	L1P2() LevelFour
}

func (obj *levelOne) L1P1() LevelTwo {
	if obj.obj.L1P1 == nil {
		obj.obj.L1P1 = &sanity.LevelTwo{}
	}
	return &levelTwo{obj: obj.obj.L1P1}

}

func (obj *levelOne) L1P2() LevelFour {
	if obj.obj.L1P2 == nil {
		obj.obj.L1P2 = &sanity.LevelFour{}
	}
	return &levelFour{obj: obj.obj.L1P2}

}

type mandate struct {
	obj *sanity.Mandate
}

func (obj *mandate) msg() *sanity.Mandate {
	return obj.obj
}

func (obj *mandate) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *mandate) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Mandate interface {
	msg() *sanity.Mandate
	Yaml() string
	Json() string
	RequiredParam() string
	SetRequiredParam(value string) Mandate
}

func (obj *mandate) RequiredParam() string {
	return obj.obj.RequiredParam
}

func (obj *mandate) SetRequiredParam(value string) Mandate {
	obj.obj.RequiredParam = value
	return obj
}

type ipv4Pattern struct {
	obj *sanity.Ipv4Pattern
}

func (obj *ipv4Pattern) msg() *sanity.Ipv4Pattern {
	return obj.obj
}

func (obj *ipv4Pattern) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv4Pattern) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Ipv4Pattern interface {
	msg() *sanity.Ipv4Pattern
	Yaml() string
	Json() string
	Ipv4() PatternIpv4PatternIpv4
}

func (obj *ipv4Pattern) Ipv4() PatternIpv4PatternIpv4 {
	if obj.obj.Ipv4 == nil {
		obj.obj.Ipv4 = &sanity.PatternIpv4PatternIpv4{}
	}
	return &patternIpv4PatternIpv4{obj: obj.obj.Ipv4}

}

type ipv6Pattern struct {
	obj *sanity.Ipv6Pattern
}

func (obj *ipv6Pattern) msg() *sanity.Ipv6Pattern {
	return obj.obj
}

func (obj *ipv6Pattern) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv6Pattern) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type Ipv6Pattern interface {
	msg() *sanity.Ipv6Pattern
	Yaml() string
	Json() string
	Ipv6() PatternIpv6PatternIpv6
}

func (obj *ipv6Pattern) Ipv6() PatternIpv6PatternIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = &sanity.PatternIpv6PatternIpv6{}
	}
	return &patternIpv6PatternIpv6{obj: obj.obj.Ipv6}

}

type macPattern struct {
	obj *sanity.MacPattern
}

func (obj *macPattern) msg() *sanity.MacPattern {
	return obj.obj
}

func (obj *macPattern) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *macPattern) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type MacPattern interface {
	msg() *sanity.MacPattern
	Yaml() string
	Json() string
	Mac() PatternMacPatternMac
}

func (obj *macPattern) Mac() PatternMacPatternMac {
	if obj.obj.Mac == nil {
		obj.obj.Mac = &sanity.PatternMacPatternMac{}
	}
	return &patternMacPatternMac{obj: obj.obj.Mac}

}

type integerPattern struct {
	obj *sanity.IntegerPattern
}

func (obj *integerPattern) msg() *sanity.IntegerPattern {
	return obj.obj
}

func (obj *integerPattern) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *integerPattern) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type IntegerPattern interface {
	msg() *sanity.IntegerPattern
	Yaml() string
	Json() string
	Integer() PatternIntegerPatternInteger
}

func (obj *integerPattern) Integer() PatternIntegerPatternInteger {
	if obj.obj.Integer == nil {
		obj.obj.Integer = &sanity.PatternIntegerPatternInteger{}
	}
	return &patternIntegerPatternInteger{obj: obj.obj.Integer}

}

type checksumPattern struct {
	obj *sanity.ChecksumPattern
}

func (obj *checksumPattern) msg() *sanity.ChecksumPattern {
	return obj.obj
}

func (obj *checksumPattern) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *checksumPattern) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type ChecksumPattern interface {
	msg() *sanity.ChecksumPattern
	Yaml() string
	Json() string
	Checksum() PatternChecksumPatternChecksum
}

func (obj *checksumPattern) Checksum() PatternChecksumPatternChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = &sanity.PatternChecksumPatternChecksum{}
	}
	return &patternChecksumPatternChecksum{obj: obj.obj.Checksum}

}

type levelTwo struct {
	obj *sanity.LevelTwo
}

func (obj *levelTwo) msg() *sanity.LevelTwo {
	return obj.obj
}

func (obj *levelTwo) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelTwo) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LevelTwo interface {
	msg() *sanity.LevelTwo
	Yaml() string
	Json() string
	L2P1() LevelThree
}

func (obj *levelTwo) L2P1() LevelThree {
	if obj.obj.L2P1 == nil {
		obj.obj.L2P1 = &sanity.LevelThree{}
	}
	return &levelThree{obj: obj.obj.L2P1}

}

type levelFour struct {
	obj *sanity.LevelFour
}

func (obj *levelFour) msg() *sanity.LevelFour {
	return obj.obj
}

func (obj *levelFour) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelFour) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LevelFour interface {
	msg() *sanity.LevelFour
	Yaml() string
	Json() string
	L4P1() LevelOne
}

func (obj *levelFour) L4P1() LevelOne {
	if obj.obj.L4P1 == nil {
		obj.obj.L4P1 = &sanity.LevelOne{}
	}
	return &levelOne{obj: obj.obj.L4P1}

}

type patternIpv4PatternIpv4 struct {
	obj *sanity.PatternIpv4PatternIpv4
}

func (obj *patternIpv4PatternIpv4) msg() *sanity.PatternIpv4PatternIpv4 {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIpv4PatternIpv4 interface {
	msg() *sanity.PatternIpv4PatternIpv4
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternIpv4PatternIpv4
	Increment() PatternIpv4PatternIpv4Counter
	Decrement() PatternIpv4PatternIpv4Counter
}

func (obj *patternIpv4PatternIpv4) Value() string {
	return *obj.obj.Value
}

func (obj *patternIpv4PatternIpv4) SetValue(value string) PatternIpv4PatternIpv4 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternIpv4PatternIpv4) Increment() PatternIpv4PatternIpv4Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIpv4PatternIpv4Counter{}
	}
	return &patternIpv4PatternIpv4Counter{obj: obj.obj.Increment}

}

func (obj *patternIpv4PatternIpv4) Decrement() PatternIpv4PatternIpv4Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &sanity.PatternIpv4PatternIpv4Counter{}
	}
	return &patternIpv4PatternIpv4Counter{obj: obj.obj.Decrement}

}

type patternIpv6PatternIpv6 struct {
	obj *sanity.PatternIpv6PatternIpv6
}

func (obj *patternIpv6PatternIpv6) msg() *sanity.PatternIpv6PatternIpv6 {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIpv6PatternIpv6 interface {
	msg() *sanity.PatternIpv6PatternIpv6
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternIpv6PatternIpv6
	Increment() PatternIpv6PatternIpv6Counter
	Decrement() PatternIpv6PatternIpv6Counter
}

func (obj *patternIpv6PatternIpv6) Value() string {
	return *obj.obj.Value
}

func (obj *patternIpv6PatternIpv6) SetValue(value string) PatternIpv6PatternIpv6 {
	obj.obj.Value = &value
	return obj
}

func (obj *patternIpv6PatternIpv6) Increment() PatternIpv6PatternIpv6Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIpv6PatternIpv6Counter{}
	}
	return &patternIpv6PatternIpv6Counter{obj: obj.obj.Increment}

}

func (obj *patternIpv6PatternIpv6) Decrement() PatternIpv6PatternIpv6Counter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &sanity.PatternIpv6PatternIpv6Counter{}
	}
	return &patternIpv6PatternIpv6Counter{obj: obj.obj.Decrement}

}

type patternMacPatternMac struct {
	obj *sanity.PatternMacPatternMac
}

func (obj *patternMacPatternMac) msg() *sanity.PatternMacPatternMac {
	return obj.obj
}

func (obj *patternMacPatternMac) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMac) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternMacPatternMac interface {
	msg() *sanity.PatternMacPatternMac
	Yaml() string
	Json() string
	Value() string
	SetValue(value string) PatternMacPatternMac
	Increment() PatternMacPatternMacCounter
	Decrement() PatternMacPatternMacCounter
}

func (obj *patternMacPatternMac) Value() string {
	return *obj.obj.Value
}

func (obj *patternMacPatternMac) SetValue(value string) PatternMacPatternMac {
	obj.obj.Value = &value
	return obj
}

func (obj *patternMacPatternMac) Increment() PatternMacPatternMacCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternMacPatternMacCounter{}
	}
	return &patternMacPatternMacCounter{obj: obj.obj.Increment}

}

func (obj *patternMacPatternMac) Decrement() PatternMacPatternMacCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &sanity.PatternMacPatternMacCounter{}
	}
	return &patternMacPatternMacCounter{obj: obj.obj.Decrement}

}

type patternIntegerPatternInteger struct {
	obj *sanity.PatternIntegerPatternInteger
}

func (obj *patternIntegerPatternInteger) msg() *sanity.PatternIntegerPatternInteger {
	return obj.obj
}

func (obj *patternIntegerPatternInteger) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternInteger) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIntegerPatternInteger interface {
	msg() *sanity.PatternIntegerPatternInteger
	Yaml() string
	Json() string
	Value() int32
	SetValue(value int32) PatternIntegerPatternInteger
	Increment() PatternIntegerPatternIntegerCounter
	Decrement() PatternIntegerPatternIntegerCounter
}

func (obj *patternIntegerPatternInteger) Value() int32 {
	return *obj.obj.Value
}

func (obj *patternIntegerPatternInteger) SetValue(value int32) PatternIntegerPatternInteger {
	obj.obj.Value = &value
	return obj
}

func (obj *patternIntegerPatternInteger) Increment() PatternIntegerPatternIntegerCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIntegerPatternIntegerCounter{}
	}
	return &patternIntegerPatternIntegerCounter{obj: obj.obj.Increment}

}

func (obj *patternIntegerPatternInteger) Decrement() PatternIntegerPatternIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.obj.Decrement = &sanity.PatternIntegerPatternIntegerCounter{}
	}
	return &patternIntegerPatternIntegerCounter{obj: obj.obj.Decrement}

}

type patternChecksumPatternChecksum struct {
	obj *sanity.PatternChecksumPatternChecksum
}

func (obj *patternChecksumPatternChecksum) msg() *sanity.PatternChecksumPatternChecksum {
	return obj.obj
}

func (obj *patternChecksumPatternChecksum) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternChecksumPatternChecksum) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternChecksumPatternChecksum interface {
	msg() *sanity.PatternChecksumPatternChecksum
	Yaml() string
	Json() string
	Custom() int32
	SetCustom(value int32) PatternChecksumPatternChecksum
}

func (obj *patternChecksumPatternChecksum) Custom() int32 {
	return *obj.obj.Custom
}

func (obj *patternChecksumPatternChecksum) SetCustom(value int32) PatternChecksumPatternChecksum {
	obj.obj.Custom = &value
	return obj
}

type levelThree struct {
	obj *sanity.LevelThree
}

func (obj *levelThree) msg() *sanity.LevelThree {
	return obj.obj
}

func (obj *levelThree) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelThree) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type LevelThree interface {
	msg() *sanity.LevelThree
	Yaml() string
	Json() string
	L3P1() string
	SetL3P1(value string) LevelThree
}

func (obj *levelThree) L3P1() string {
	return *obj.obj.L3P1
}

func (obj *levelThree) SetL3P1(value string) LevelThree {
	obj.obj.L3P1 = &value
	return obj
}

type patternIpv4PatternIpv4Counter struct {
	obj *sanity.PatternIpv4PatternIpv4Counter
}

func (obj *patternIpv4PatternIpv4Counter) msg() *sanity.PatternIpv4PatternIpv4Counter {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIpv4PatternIpv4Counter interface {
	msg() *sanity.PatternIpv4PatternIpv4Counter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternIpv4PatternIpv4Counter
	Step() string
	SetStep(value string) PatternIpv4PatternIpv4Counter
	Count() int32
	SetCount(value int32) PatternIpv4PatternIpv4Counter
}

func (obj *patternIpv4PatternIpv4Counter) Start() string {
	return *obj.obj.Start
}

func (obj *patternIpv4PatternIpv4Counter) SetStart(value string) PatternIpv4PatternIpv4Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternIpv4PatternIpv4Counter) Step() string {
	return *obj.obj.Step
}

func (obj *patternIpv4PatternIpv4Counter) SetStep(value string) PatternIpv4PatternIpv4Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternIpv4PatternIpv4Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternIpv4PatternIpv4Counter) SetCount(value int32) PatternIpv4PatternIpv4Counter {
	obj.obj.Count = &value
	return obj
}

type patternIpv6PatternIpv6Counter struct {
	obj *sanity.PatternIpv6PatternIpv6Counter
}

func (obj *patternIpv6PatternIpv6Counter) msg() *sanity.PatternIpv6PatternIpv6Counter {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6Counter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6Counter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIpv6PatternIpv6Counter interface {
	msg() *sanity.PatternIpv6PatternIpv6Counter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternIpv6PatternIpv6Counter
	Step() string
	SetStep(value string) PatternIpv6PatternIpv6Counter
	Count() int32
	SetCount(value int32) PatternIpv6PatternIpv6Counter
}

func (obj *patternIpv6PatternIpv6Counter) Start() string {
	return *obj.obj.Start
}

func (obj *patternIpv6PatternIpv6Counter) SetStart(value string) PatternIpv6PatternIpv6Counter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternIpv6PatternIpv6Counter) Step() string {
	return *obj.obj.Step
}

func (obj *patternIpv6PatternIpv6Counter) SetStep(value string) PatternIpv6PatternIpv6Counter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternIpv6PatternIpv6Counter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternIpv6PatternIpv6Counter) SetCount(value int32) PatternIpv6PatternIpv6Counter {
	obj.obj.Count = &value
	return obj
}

type patternMacPatternMacCounter struct {
	obj *sanity.PatternMacPatternMacCounter
}

func (obj *patternMacPatternMacCounter) msg() *sanity.PatternMacPatternMacCounter {
	return obj.obj
}

func (obj *patternMacPatternMacCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMacCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternMacPatternMacCounter interface {
	msg() *sanity.PatternMacPatternMacCounter
	Yaml() string
	Json() string
	Start() string
	SetStart(value string) PatternMacPatternMacCounter
	Step() string
	SetStep(value string) PatternMacPatternMacCounter
	Count() int32
	SetCount(value int32) PatternMacPatternMacCounter
}

func (obj *patternMacPatternMacCounter) Start() string {
	return *obj.obj.Start
}

func (obj *patternMacPatternMacCounter) SetStart(value string) PatternMacPatternMacCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternMacPatternMacCounter) Step() string {
	return *obj.obj.Step
}

func (obj *patternMacPatternMacCounter) SetStep(value string) PatternMacPatternMacCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternMacPatternMacCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternMacPatternMacCounter) SetCount(value int32) PatternMacPatternMacCounter {
	obj.obj.Count = &value
	return obj
}

type patternIntegerPatternIntegerCounter struct {
	obj *sanity.PatternIntegerPatternIntegerCounter
}

func (obj *patternIntegerPatternIntegerCounter) msg() *sanity.PatternIntegerPatternIntegerCounter {
	return obj.obj
}

func (obj *patternIntegerPatternIntegerCounter) Yaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternIntegerCounter) Json() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

type PatternIntegerPatternIntegerCounter interface {
	msg() *sanity.PatternIntegerPatternIntegerCounter
	Yaml() string
	Json() string
	Start() int32
	SetStart(value int32) PatternIntegerPatternIntegerCounter
	Step() int32
	SetStep(value int32) PatternIntegerPatternIntegerCounter
	Count() int32
	SetCount(value int32) PatternIntegerPatternIntegerCounter
}

func (obj *patternIntegerPatternIntegerCounter) Start() int32 {
	return *obj.obj.Start
}

func (obj *patternIntegerPatternIntegerCounter) SetStart(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Start = &value
	return obj
}

func (obj *patternIntegerPatternIntegerCounter) Step() int32 {
	return *obj.obj.Step
}

func (obj *patternIntegerPatternIntegerCounter) SetStep(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Step = &value
	return obj
}

func (obj *patternIntegerPatternIntegerCounter) Count() int32 {
	return *obj.obj.Count
}

func (obj *patternIntegerPatternIntegerCounter) SetCount(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Count = &value
	return obj
}
