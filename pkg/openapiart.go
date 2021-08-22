// OpenAPIArt Test API 0.0.1
// License: MIT

package openapiart

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	sanity "github.com/open-traffic-generator/openapiart/pkg/sanity"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v3"
)

type grpcTransport struct {
	location       string
	requestTimeout time.Duration
}

type GrpcTransport interface {
	SetLocation(value string) GrpcTransport
	Location() string
	SetRequestTimeout(value int) GrpcTransport
	RequestTimeout() int
}

// Location
func (obj *grpcTransport) Location() string {
	return obj.location
}

// SetLocation
func (obj *grpcTransport) SetLocation(value string) GrpcTransport {
	obj.location = value
	return obj
}

// RequestTimeout returns the grpc request timeout in seconds
func (obj *grpcTransport) RequestTimeout() int {
	return int(obj.requestTimeout / time.Second)
}

// SetRequestTimeout contains the timeout value in seconds for a grpc request
func (obj *grpcTransport) SetRequestTimeout(value int) GrpcTransport {
	obj.requestTimeout = time.Duration(value) * time.Second
	return obj
}

type httpTransport struct {
	location string
	verify   bool
}

type HttpTransport interface {
	SetLocation(value string) HttpTransport
	Location() string
	SetVerify(value bool) HttpTransport
	Verify() bool
}

// Location
func (obj *httpTransport) Location() string {
	return obj.location
}

// SetLocation
func (obj *httpTransport) SetLocation(value string) HttpTransport {
	obj.location = value
	return obj
}

// Verify returns whether or not TLS certificates will be verified by the server
func (obj *httpTransport) Verify() bool {
	return obj.verify
}

// SetVerify determines whether or not TLS certificates will be verified by the server
func (obj *httpTransport) SetVerify(value bool) HttpTransport {
	obj.verify = value
	return obj
}

type api struct {
	grpc *grpcTransport
	http *httpTransport
}

type Api interface {
	NewGrpcTransport() GrpcTransport
	NewHttpTransport() HttpTransport
}

// NewGrpcTransport sets the underlying transport of the Api as grpc
func (api *api) NewGrpcTransport() GrpcTransport {
	api.grpc = &grpcTransport{
		location:       "127.0.0.1:5050",
		requestTimeout: 10 * time.Second,
	}
	api.http = nil
	return api.grpc
}

// NewHttpTransport sets the underlying transport of the Api as http
func (api *api) NewHttpTransport() HttpTransport {
	api.http = &httpTransport{
		location: "https://127.0.0.1:443",
		verify:   false,
	}
	api.grpc = nil
	return api.http
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
		conn, err := grpc.Dial(api.grpc.location, grpc.WithInsecure())
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
	return &api
}

type OpenapiartApi interface {
	Api
	NewPrefixConfig() PrefixConfig
	NewUpdateConfig() UpdateConfig
	SetConfig(prefixConfig PrefixConfig) (SetConfigResponse_StatusCode200, error)
	UpdateConfig(updateConfig UpdateConfig) (UpdateConfigResponse_StatusCode200, error)
	GetConfig() (GetConfigResponse_StatusCode200, error)
}

func (api *openapiartApi) NewPrefixConfig() PrefixConfig {
	return &prefixConfig{obj: &sanity.PrefixConfig{}}
}

func (api *openapiartApi) NewUpdateConfig() UpdateConfig {
	return &updateConfig{obj: &sanity.UpdateConfig{}}
}

func (api *openapiartApi) SetConfig(prefixConfig PrefixConfig) (SetConfigResponse_StatusCode200, error) {
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := sanity.SetConfigRequest{PrefixConfig: prefixConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode_200() != nil {
		return &setConfigResponseStatusCode200{obj: resp.GetStatusCode_200()}, nil
	}
	if resp.GetStatusCode_400() != nil {
		data, _ := yaml.Marshal(resp.GetStatusCode_400())
		return nil, fmt.Errorf(string(data))
	}
	if resp.GetStatusCode_500() != nil {
		data, _ := yaml.Marshal(resp.GetStatusCode_400())
		return nil, fmt.Errorf(string(data))
	}
	return nil, fmt.Errorf("Response not implemented")
}

func (api *openapiartApi) UpdateConfig(updateConfig UpdateConfig) (UpdateConfigResponse_StatusCode200, error) {
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := sanity.UpdateConfigRequest{UpdateConfig: updateConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.UpdateConfig(ctx, &request)
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode_200() != nil {
		return &updateConfigResponseStatusCode200{obj: resp.GetStatusCode_200()}, nil
	}
	if resp.GetStatusCode_400() != nil {
		data, _ := yaml.Marshal(resp.GetStatusCode_400())
		return nil, fmt.Errorf(string(data))
	}
	if resp.GetStatusCode_500() != nil {
		data, _ := yaml.Marshal(resp.GetStatusCode_400())
		return nil, fmt.Errorf(string(data))
	}
	return nil, fmt.Errorf("Response not implemented")
}

func (api *openapiartApi) GetConfig() (GetConfigResponse_StatusCode200, error) {
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetConfig(ctx, &request)
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode_200() != nil {
		return &getConfigResponseStatusCode200{obj: resp.GetStatusCode_200()}, nil
	}
	return nil, fmt.Errorf("Response not implemented")
}

type prefixConfig struct {
	obj *sanity.PrefixConfig
}

func (obj *prefixConfig) msg() *sanity.PrefixConfig {
	return obj.obj
}

func (obj *prefixConfig) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *prefixConfig) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *prefixConfig) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *prefixConfig) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PrefixConfig interface {
	msg() *sanity.PrefixConfig
	ToYaml() string
	ToJson() string
	Ieee8021Qbb() bool
	SetIeee8021Qbb(value bool) PrefixConfig
	Space1() int32
	SetSpace1(value int32) PrefixConfig
	FullDuplex100Mb() int32
	SetFullDuplex100Mb(value int32) PrefixConfig
	A() string
	SetA(value string) PrefixConfig
	B() float32
	SetB(value float32) PrefixConfig
	C() int32
	SetC(value int32) PrefixConfig
	E() EObject
	F() FObject
	G() PrefixConfigGObjectIter
	H() bool
	SetH(value bool) PrefixConfig
	I() []byte
	SetI(value []byte) PrefixConfig
	J() PrefixConfigJObjectIter
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

// Ieee_802_1Qbb returns a bool
//  description is TBD
func (obj *prefixConfig) Ieee8021Qbb() bool {
	return *obj.obj.Ieee_802_1Qbb
}

// SetIeee_802_1Qbb sets the bool value in the None object
//  description is TBD
func (obj *prefixConfig) SetIeee8021Qbb(value bool) PrefixConfig {
	obj.obj.Ieee_802_1Qbb = &value
	return obj
}

// Space_1 returns a int32
//  description is TBD
func (obj *prefixConfig) Space1() int32 {
	return *obj.obj.Space_1
}

// SetSpace_1 sets the int32 value in the None object
//  description is TBD
func (obj *prefixConfig) SetSpace1(value int32) PrefixConfig {
	obj.obj.Space_1 = &value
	return obj
}

// FullDuplex_100Mb returns a int32
//  description is TBD
func (obj *prefixConfig) FullDuplex100Mb() int32 {
	return *obj.obj.FullDuplex_100Mb
}

// SetFullDuplex_100Mb sets the int32 value in the None object
//  description is TBD
func (obj *prefixConfig) SetFullDuplex100Mb(value int32) PrefixConfig {
	obj.obj.FullDuplex_100Mb = &value
	return obj
}

// A returns a string
//  Small single line description
func (obj *prefixConfig) A() string {
	return obj.obj.A
}

// SetA sets the string value in the None object
//  Small single line description
func (obj *prefixConfig) SetA(value string) PrefixConfig {
	obj.obj.A = value
	return obj
}

// B returns a float32
//  Longer multi-line description
//  Second line is here
//  Third line
func (obj *prefixConfig) B() float32 {
	return obj.obj.B
}

// SetB sets the float32 value in the None object
//  Longer multi-line description
//  Second line is here
//  Third line
func (obj *prefixConfig) SetB(value float32) PrefixConfig {
	obj.obj.B = value
	return obj
}

// C returns a int32
//  description is TBD
func (obj *prefixConfig) C() int32 {
	return obj.obj.C
}

// SetC sets the int32 value in the None object
//  description is TBD
func (obj *prefixConfig) SetC(value int32) PrefixConfig {
	obj.obj.C = value
	return obj
}

// E returns a EObject
//  A child object
func (obj *prefixConfig) E() EObject {
	if obj.obj.E == nil {
		obj.obj.E = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.E}

}

// F returns a FObject
//  An object with only choice(s)
func (obj *prefixConfig) F() FObject {
	if obj.obj.F == nil {
		obj.obj.F = &sanity.FObject{}
	}
	return &fObject{obj: obj.obj.F}

}

// G returns a []GObject
//  A list of objects with choice and properties
func (obj *prefixConfig) G() PrefixConfigGObjectIter {
	if obj.obj.G == nil {
		obj.obj.G = []*sanity.GObject{}
	}
	return &prefixConfigGObjectIter{obj: obj}

}

type prefixConfigGObjectIter struct {
	obj *prefixConfig
}

type PrefixConfigGObjectIter interface {
	Add() GObject
	Items() []GObject
}

func (obj *prefixConfigGObjectIter) Add() GObject {
	newObj := &sanity.GObject{}
	obj.obj.obj.G = append(obj.obj.obj.G, newObj)
	return &gObject{obj: newObj}
}

func (obj *prefixConfigGObjectIter) Items() []GObject {
	slice := []GObject{}
	for _, item := range obj.obj.obj.G {
		slice = append(slice, &gObject{obj: item})
	}
	return slice
}

// H returns a bool
//  A boolean value
func (obj *prefixConfig) H() bool {
	return *obj.obj.H
}

// SetH sets the bool value in the None object
//  A boolean value
func (obj *prefixConfig) SetH(value bool) PrefixConfig {
	obj.obj.H = &value
	return obj
}

// I returns a []byte
//  A byte string
func (obj *prefixConfig) I() []byte {
	return obj.obj.I
}

// SetI sets the []byte value in the None object
//  A byte string
func (obj *prefixConfig) SetI(value []byte) PrefixConfig {
	obj.obj.I = value
	return obj
}

// J returns a []JObject
//  A list of objects with only choice
func (obj *prefixConfig) J() PrefixConfigJObjectIter {
	if obj.obj.J == nil {
		obj.obj.J = []*sanity.JObject{}
	}
	return &prefixConfigJObjectIter{obj: obj}

}

type prefixConfigJObjectIter struct {
	obj *prefixConfig
}

type PrefixConfigJObjectIter interface {
	Add() JObject
	Items() []JObject
}

func (obj *prefixConfigJObjectIter) Add() JObject {
	newObj := &sanity.JObject{}
	obj.obj.obj.J = append(obj.obj.obj.J, newObj)
	return &jObject{obj: newObj}
}

func (obj *prefixConfigJObjectIter) Items() []JObject {
	slice := []JObject{}
	for _, item := range obj.obj.obj.J {
		slice = append(slice, &jObject{obj: item})
	}
	return slice
}

// K returns a KObject
//  A nested object with only one property which is a choice object
func (obj *prefixConfig) K() KObject {
	if obj.obj.K == nil {
		obj.obj.K = &sanity.KObject{}
	}
	return &kObject{obj: obj.obj.K}

}

// L returns a LObject
//  description is TBD
func (obj *prefixConfig) L() LObject {
	if obj.obj.L == nil {
		obj.obj.L = &sanity.LObject{}
	}
	return &lObject{obj: obj.obj.L}

}

// Level returns a LevelOne
//  description is TBD
func (obj *prefixConfig) Level() LevelOne {
	if obj.obj.Level == nil {
		obj.obj.Level = &sanity.LevelOne{}
	}
	return &levelOne{obj: obj.obj.Level}

}

// Mandatory returns a Mandate
//  description is TBD
func (obj *prefixConfig) Mandatory() Mandate {
	if obj.obj.Mandatory == nil {
		obj.obj.Mandatory = &sanity.Mandate{}
	}
	return &mandate{obj: obj.obj.Mandatory}

}

// Ipv4Pattern returns a Ipv4Pattern
//  description is TBD
func (obj *prefixConfig) Ipv4Pattern() Ipv4Pattern {
	if obj.obj.Ipv4Pattern == nil {
		obj.obj.Ipv4Pattern = &sanity.Ipv4Pattern{}
	}
	return &ipv4Pattern{obj: obj.obj.Ipv4Pattern}

}

// Ipv6Pattern returns a Ipv6Pattern
//  description is TBD
func (obj *prefixConfig) Ipv6Pattern() Ipv6Pattern {
	if obj.obj.Ipv6Pattern == nil {
		obj.obj.Ipv6Pattern = &sanity.Ipv6Pattern{}
	}
	return &ipv6Pattern{obj: obj.obj.Ipv6Pattern}

}

// MacPattern returns a MacPattern
//  description is TBD
func (obj *prefixConfig) MacPattern() MacPattern {
	if obj.obj.MacPattern == nil {
		obj.obj.MacPattern = &sanity.MacPattern{}
	}
	return &macPattern{obj: obj.obj.MacPattern}

}

// IntegerPattern returns a IntegerPattern
//  description is TBD
func (obj *prefixConfig) IntegerPattern() IntegerPattern {
	if obj.obj.IntegerPattern == nil {
		obj.obj.IntegerPattern = &sanity.IntegerPattern{}
	}
	return &integerPattern{obj: obj.obj.IntegerPattern}

}

// ChecksumPattern returns a ChecksumPattern
//  description is TBD
func (obj *prefixConfig) ChecksumPattern() ChecksumPattern {
	if obj.obj.ChecksumPattern == nil {
		obj.obj.ChecksumPattern = &sanity.ChecksumPattern{}
	}
	return &checksumPattern{obj: obj.obj.ChecksumPattern}

}

// Name returns a string
//  description is TBD
func (obj *prefixConfig) Name() string {
	return *obj.obj.Name
}

// SetName sets the string value in the None object
//  description is TBD
func (obj *prefixConfig) SetName(value string) PrefixConfig {
	obj.obj.Name = &value
	return obj
}

type updateConfig struct {
	obj *sanity.UpdateConfig
}

func (obj *updateConfig) msg() *sanity.UpdateConfig {
	return obj.obj
}

func (obj *updateConfig) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *updateConfig) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *updateConfig) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *updateConfig) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type UpdateConfig interface {
	msg() *sanity.UpdateConfig
	ToYaml() string
	ToJson() string
	G() UpdateConfigGObjectIter
}

// G returns a []GObject
//  A list of objects with choice and properties
func (obj *updateConfig) G() UpdateConfigGObjectIter {
	if obj.obj.G == nil {
		obj.obj.G = []*sanity.GObject{}
	}
	return &updateConfigGObjectIter{obj: obj}

}

type updateConfigGObjectIter struct {
	obj *updateConfig
}

type UpdateConfigGObjectIter interface {
	Add() GObject
	Items() []GObject
}

func (obj *updateConfigGObjectIter) Add() GObject {
	newObj := &sanity.GObject{}
	obj.obj.obj.G = append(obj.obj.obj.G, newObj)
	return &gObject{obj: newObj}
}

func (obj *updateConfigGObjectIter) Items() []GObject {
	slice := []GObject{}
	for _, item := range obj.obj.obj.G {
		slice = append(slice, &gObject{obj: item})
	}
	return slice
}

type eObject struct {
	obj *sanity.EObject
}

func (obj *eObject) msg() *sanity.EObject {
	return obj.obj
}

func (obj *eObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *eObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *eObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *eObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type EObject interface {
	msg() *sanity.EObject
	ToYaml() string
	ToJson() string
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

// EA returns a float32
//  description is TBD
func (obj *eObject) EA() float32 {
	return obj.obj.EA
}

// SetEA sets the float32 value in the None object
//  description is TBD
func (obj *eObject) SetEA(value float32) EObject {
	obj.obj.EA = value
	return obj
}

// EB returns a float64
//  description is TBD
func (obj *eObject) EB() float64 {
	return obj.obj.EB
}

// SetEB sets the float64 value in the None object
//  description is TBD
func (obj *eObject) SetEB(value float64) EObject {
	obj.obj.EB = value
	return obj
}

// Name returns a string
//  description is TBD
func (obj *eObject) Name() string {
	return *obj.obj.Name
}

// SetName sets the string value in the None object
//  description is TBD
func (obj *eObject) SetName(value string) EObject {
	obj.obj.Name = &value
	return obj
}

// MParam1 returns a string
//  description is TBD
func (obj *eObject) MParam1() string {
	return *obj.obj.MParam1
}

// SetMParam1 sets the string value in the None object
//  description is TBD
func (obj *eObject) SetMParam1(value string) EObject {
	obj.obj.MParam1 = &value
	return obj
}

// MParam2 returns a string
//  description is TBD
func (obj *eObject) MParam2() string {
	return *obj.obj.MParam2
}

// SetMParam2 sets the string value in the None object
//  description is TBD
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

func (obj *fObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *fObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *fObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *fObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type FObject interface {
	msg() *sanity.FObject
	ToYaml() string
	ToJson() string
	FA() string
	SetFA(value string) FObject
	FB() float64
	SetFB(value float64) FObject
}

// FA returns a string
//  description is TBD
func (obj *fObject) FA() string {
	return *obj.obj.FA
}

// SetFA sets the string value in the None object
//  description is TBD
func (obj *fObject) SetFA(value string) FObject {
	obj.obj.FA = &value
	return obj
}

// FB returns a float64
//  description is TBD
func (obj *fObject) FB() float64 {
	return *obj.obj.FB
}

// SetFB sets the float64 value in the None object
//  description is TBD
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

func (obj *gObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *gObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *gObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *gObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type GObject interface {
	msg() *sanity.GObject
	ToYaml() string
	ToJson() string
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

// GA returns a string
//  description is TBD
func (obj *gObject) GA() string {
	return *obj.obj.GA
}

// SetGA sets the string value in the None object
//  description is TBD
func (obj *gObject) SetGA(value string) GObject {
	obj.obj.GA = &value
	return obj
}

// GB returns a int32
//  description is TBD
func (obj *gObject) GB() int32 {
	return *obj.obj.GB
}

// SetGB sets the int32 value in the None object
//  description is TBD
func (obj *gObject) SetGB(value int32) GObject {
	obj.obj.GB = &value
	return obj
}

// GC returns a float32
//  description is TBD
func (obj *gObject) GC() float32 {
	return *obj.obj.GC
}

// SetGC sets the float32 value in the None object
//  description is TBD
func (obj *gObject) SetGC(value float32) GObject {
	obj.obj.GC = &value
	return obj
}

// GD returns a string
//  description is TBD
func (obj *gObject) GD() string {
	return *obj.obj.GD
}

// SetGD sets the string value in the None object
//  description is TBD
func (obj *gObject) SetGD(value string) GObject {
	obj.obj.GD = &value
	return obj
}

// GE returns a float64
//  description is TBD
func (obj *gObject) GE() float64 {
	return *obj.obj.GE
}

// SetGE sets the float64 value in the None object
//  description is TBD
func (obj *gObject) SetGE(value float64) GObject {
	obj.obj.GE = &value
	return obj
}

// Name returns a string
//  description is TBD
func (obj *gObject) Name() string {
	return *obj.obj.Name
}

// SetName sets the string value in the None object
//  description is TBD
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

func (obj *jObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *jObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *jObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *jObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type JObject interface {
	msg() *sanity.JObject
	ToYaml() string
	ToJson() string
	JA() EObject
	JB() FObject
}

// JA returns a EObject
//  description is TBD
func (obj *jObject) JA() EObject {
	if obj.obj.JA == nil {
		obj.obj.JA = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.JA}

}

// JB returns a FObject
//  description is TBD
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

func (obj *kObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *kObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *kObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *kObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type KObject interface {
	msg() *sanity.KObject
	ToYaml() string
	ToJson() string
	EObject() EObject
	FObject() FObject
}

// EObject returns a EObject
//  description is TBD
func (obj *kObject) EObject() EObject {
	if obj.obj.EObject == nil {
		obj.obj.EObject = &sanity.EObject{}
	}
	return &eObject{obj: obj.obj.EObject}

}

// FObject returns a FObject
//  description is TBD
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

func (obj *lObject) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *lObject) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *lObject) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *lObject) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type LObject interface {
	msg() *sanity.LObject
	ToYaml() string
	ToJson() string
	String() string
	SetString(value string) LObject
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

// String_ returns a string
//  description is TBD
func (obj *lObject) String() string {
	return *obj.obj.String_
}

// SetString_ sets the string value in the None object
//  description is TBD
func (obj *lObject) SetString(value string) LObject {
	obj.obj.String_ = &value
	return obj
}

// Integer returns a int32
//  description is TBD
func (obj *lObject) Integer() int32 {
	return *obj.obj.Integer
}

// SetInteger sets the int32 value in the None object
//  description is TBD
func (obj *lObject) SetInteger(value int32) LObject {
	obj.obj.Integer = &value
	return obj
}

// Float returns a float32
//  description is TBD
func (obj *lObject) Float() float32 {
	return *obj.obj.Float
}

// SetFloat sets the float32 value in the None object
//  description is TBD
func (obj *lObject) SetFloat(value float32) LObject {
	obj.obj.Float = &value
	return obj
}

// Double returns a float64
//  description is TBD
func (obj *lObject) Double() float64 {
	return *obj.obj.Double
}

// SetDouble sets the float64 value in the None object
//  description is TBD
func (obj *lObject) SetDouble(value float64) LObject {
	obj.obj.Double = &value
	return obj
}

// Mac returns a string
//  description is TBD
func (obj *lObject) Mac() string {
	return *obj.obj.Mac
}

// SetMac sets the string value in the None object
//  description is TBD
func (obj *lObject) SetMac(value string) LObject {
	obj.obj.Mac = &value
	return obj
}

// Ipv4 returns a string
//  description is TBD
func (obj *lObject) Ipv4() string {
	return *obj.obj.Ipv4
}

// SetIpv4 sets the string value in the None object
//  description is TBD
func (obj *lObject) SetIpv4(value string) LObject {
	obj.obj.Ipv4 = &value
	return obj
}

// Ipv6 returns a string
//  description is TBD
func (obj *lObject) Ipv6() string {
	return *obj.obj.Ipv6
}

// SetIpv6 sets the string value in the None object
//  description is TBD
func (obj *lObject) SetIpv6(value string) LObject {
	obj.obj.Ipv6 = &value
	return obj
}

// Hex returns a string
//  description is TBD
func (obj *lObject) Hex() string {
	return *obj.obj.Hex
}

// SetHex sets the string value in the None object
//  description is TBD
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

func (obj *levelOne) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelOne) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *levelOne) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *levelOne) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type LevelOne interface {
	msg() *sanity.LevelOne
	ToYaml() string
	ToJson() string
	L1P1() LevelTwo
	L1P2() LevelFour
}

// L1P1 returns a LevelTwo
//  Level one
func (obj *levelOne) L1P1() LevelTwo {
	if obj.obj.L1P1 == nil {
		obj.obj.L1P1 = &sanity.LevelTwo{}
	}
	return &levelTwo{obj: obj.obj.L1P1}

}

// L1P2 returns a LevelFour
//  Level one to four
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

func (obj *mandate) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *mandate) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *mandate) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *mandate) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type Mandate interface {
	msg() *sanity.Mandate
	ToYaml() string
	ToJson() string
	RequiredParam() string
	SetRequiredParam(value string) Mandate
}

// RequiredParam returns a string
//  description is TBD
func (obj *mandate) RequiredParam() string {
	return obj.obj.RequiredParam
}

// SetRequiredParam sets the string value in the None object
//  description is TBD
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

func (obj *ipv4Pattern) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv4Pattern) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *ipv4Pattern) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv4Pattern) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type Ipv4Pattern interface {
	msg() *sanity.Ipv4Pattern
	ToYaml() string
	ToJson() string
	Ipv4() PatternIpv4PatternIpv4
}

// Ipv4 returns a PatternIpv4PatternIpv4
//  description is TBD
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

func (obj *ipv6Pattern) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv6Pattern) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *ipv6Pattern) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *ipv6Pattern) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type Ipv6Pattern interface {
	msg() *sanity.Ipv6Pattern
	ToYaml() string
	ToJson() string
	Ipv6() PatternIpv6PatternIpv6
}

// Ipv6 returns a PatternIpv6PatternIpv6
//  description is TBD
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

func (obj *macPattern) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *macPattern) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *macPattern) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *macPattern) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type MacPattern interface {
	msg() *sanity.MacPattern
	ToYaml() string
	ToJson() string
	Mac() PatternMacPatternMac
}

// Mac returns a PatternMacPatternMac
//  description is TBD
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

func (obj *integerPattern) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *integerPattern) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *integerPattern) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *integerPattern) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type IntegerPattern interface {
	msg() *sanity.IntegerPattern
	ToYaml() string
	ToJson() string
	Integer() PatternIntegerPatternInteger
}

// Integer returns a PatternIntegerPatternInteger
//  description is TBD
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

func (obj *checksumPattern) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *checksumPattern) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *checksumPattern) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *checksumPattern) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type ChecksumPattern interface {
	msg() *sanity.ChecksumPattern
	ToYaml() string
	ToJson() string
	Checksum() PatternChecksumPatternChecksum
}

// Checksum returns a PatternChecksumPatternChecksum
//  description is TBD
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

func (obj *levelTwo) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelTwo) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *levelTwo) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *levelTwo) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type LevelTwo interface {
	msg() *sanity.LevelTwo
	ToYaml() string
	ToJson() string
	L2P1() LevelThree
}

// L2P1 returns a LevelThree
//  Level Two
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

func (obj *levelFour) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelFour) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *levelFour) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *levelFour) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type LevelFour interface {
	msg() *sanity.LevelFour
	ToYaml() string
	ToJson() string
	L4P1() LevelOne
}

// L4P1 returns a LevelOne
//  loop over level 1
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

func (obj *patternIpv4PatternIpv4) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIpv4PatternIpv4) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIpv4PatternIpv4 interface {
	msg() *sanity.PatternIpv4PatternIpv4
	ToYaml() string
	ToJson() string
	Value() string
	SetValue(value string) PatternIpv4PatternIpv4
	Values() []string
	SetValues(value []string) PatternIpv4PatternIpv4
	Increment() PatternIpv4PatternIpv4Counter
	Decrement() PatternIpv4PatternIpv4Counter
}

// Value returns a string
//  description is TBD
func (obj *patternIpv4PatternIpv4) Value() string {
	return *obj.obj.Value
}

// SetValue sets the string value in the None object
//  description is TBD
func (obj *patternIpv4PatternIpv4) SetValue(value string) PatternIpv4PatternIpv4 {
	obj.obj.Value = &value
	return obj
}

// Values returns a []string
//  description is TBD
func (obj *patternIpv4PatternIpv4) Values() []string {
	return obj.obj.Values
}

// SetValues sets the []string value in the None object
//  description is TBD
func (obj *patternIpv4PatternIpv4) SetValues(value []string) PatternIpv4PatternIpv4 {
	obj.obj.Values = value
	return obj
}

// Increment returns a PatternIpv4PatternIpv4Counter
//  description is TBD
func (obj *patternIpv4PatternIpv4) Increment() PatternIpv4PatternIpv4Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIpv4PatternIpv4Counter{}
	}
	return &patternIpv4PatternIpv4Counter{obj: obj.obj.Increment}

}

// Decrement returns a PatternIpv4PatternIpv4Counter
//  description is TBD
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

func (obj *patternIpv6PatternIpv6) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIpv6PatternIpv6) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIpv6PatternIpv6 interface {
	msg() *sanity.PatternIpv6PatternIpv6
	ToYaml() string
	ToJson() string
	Value() string
	SetValue(value string) PatternIpv6PatternIpv6
	Values() []string
	SetValues(value []string) PatternIpv6PatternIpv6
	Increment() PatternIpv6PatternIpv6Counter
	Decrement() PatternIpv6PatternIpv6Counter
}

// Value returns a string
//  description is TBD
func (obj *patternIpv6PatternIpv6) Value() string {
	return *obj.obj.Value
}

// SetValue sets the string value in the None object
//  description is TBD
func (obj *patternIpv6PatternIpv6) SetValue(value string) PatternIpv6PatternIpv6 {
	obj.obj.Value = &value
	return obj
}

// Values returns a []string
//  description is TBD
func (obj *patternIpv6PatternIpv6) Values() []string {
	return obj.obj.Values
}

// SetValues sets the []string value in the None object
//  description is TBD
func (obj *patternIpv6PatternIpv6) SetValues(value []string) PatternIpv6PatternIpv6 {
	obj.obj.Values = value
	return obj
}

// Increment returns a PatternIpv6PatternIpv6Counter
//  description is TBD
func (obj *patternIpv6PatternIpv6) Increment() PatternIpv6PatternIpv6Counter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIpv6PatternIpv6Counter{}
	}
	return &patternIpv6PatternIpv6Counter{obj: obj.obj.Increment}

}

// Decrement returns a PatternIpv6PatternIpv6Counter
//  description is TBD
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

func (obj *patternMacPatternMac) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMac) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternMacPatternMac) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMac) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternMacPatternMac interface {
	msg() *sanity.PatternMacPatternMac
	ToYaml() string
	ToJson() string
	Value() string
	SetValue(value string) PatternMacPatternMac
	Values() []string
	SetValues(value []string) PatternMacPatternMac
	Increment() PatternMacPatternMacCounter
	Decrement() PatternMacPatternMacCounter
}

// Value returns a string
//  description is TBD
func (obj *patternMacPatternMac) Value() string {
	return *obj.obj.Value
}

// SetValue sets the string value in the None object
//  description is TBD
func (obj *patternMacPatternMac) SetValue(value string) PatternMacPatternMac {
	obj.obj.Value = &value
	return obj
}

// Values returns a []string
//  description is TBD
func (obj *patternMacPatternMac) Values() []string {
	return obj.obj.Values
}

// SetValues sets the []string value in the None object
//  description is TBD
func (obj *patternMacPatternMac) SetValues(value []string) PatternMacPatternMac {
	obj.obj.Values = value
	return obj
}

// Increment returns a PatternMacPatternMacCounter
//  description is TBD
func (obj *patternMacPatternMac) Increment() PatternMacPatternMacCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternMacPatternMacCounter{}
	}
	return &patternMacPatternMacCounter{obj: obj.obj.Increment}

}

// Decrement returns a PatternMacPatternMacCounter
//  description is TBD
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

func (obj *patternIntegerPatternInteger) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternInteger) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIntegerPatternInteger) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternInteger) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIntegerPatternInteger interface {
	msg() *sanity.PatternIntegerPatternInteger
	ToYaml() string
	ToJson() string
	Value() int32
	SetValue(value int32) PatternIntegerPatternInteger
	Values() []int32
	SetValues(value []int32) PatternIntegerPatternInteger
	Increment() PatternIntegerPatternIntegerCounter
	Decrement() PatternIntegerPatternIntegerCounter
}

// Value returns a int32
//  description is TBD
func (obj *patternIntegerPatternInteger) Value() int32 {
	return *obj.obj.Value
}

// SetValue sets the int32 value in the None object
//  description is TBD
func (obj *patternIntegerPatternInteger) SetValue(value int32) PatternIntegerPatternInteger {
	obj.obj.Value = &value
	return obj
}

// Values returns a []int32
//  description is TBD
func (obj *patternIntegerPatternInteger) Values() []int32 {
	return obj.obj.Values
}

// SetValues sets the []int32 value in the None object
//  description is TBD
func (obj *patternIntegerPatternInteger) SetValues(value []int32) PatternIntegerPatternInteger {
	obj.obj.Values = value
	return obj
}

// Increment returns a PatternIntegerPatternIntegerCounter
//  description is TBD
func (obj *patternIntegerPatternInteger) Increment() PatternIntegerPatternIntegerCounter {
	if obj.obj.Increment == nil {
		obj.obj.Increment = &sanity.PatternIntegerPatternIntegerCounter{}
	}
	return &patternIntegerPatternIntegerCounter{obj: obj.obj.Increment}

}

// Decrement returns a PatternIntegerPatternIntegerCounter
//  description is TBD
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

func (obj *patternChecksumPatternChecksum) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternChecksumPatternChecksum) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternChecksumPatternChecksum) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternChecksumPatternChecksum) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternChecksumPatternChecksum interface {
	msg() *sanity.PatternChecksumPatternChecksum
	ToYaml() string
	ToJson() string
	Custom() int32
	SetCustom(value int32) PatternChecksumPatternChecksum
}

// Custom returns a int32
//  A custom checksum value
func (obj *patternChecksumPatternChecksum) Custom() int32 {
	return *obj.obj.Custom
}

// SetCustom sets the int32 value in the None object
//  A custom checksum value
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

func (obj *levelThree) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *levelThree) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *levelThree) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *levelThree) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type LevelThree interface {
	msg() *sanity.LevelThree
	ToYaml() string
	ToJson() string
	L3P1() string
	SetL3P1(value string) LevelThree
}

// L3P1 returns a string
//  Set value at Level 3
func (obj *levelThree) L3P1() string {
	return *obj.obj.L3P1
}

// SetL3P1 sets the string value in the None object
//  Set value at Level 3
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

func (obj *patternIpv4PatternIpv4Counter) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4Counter) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIpv4PatternIpv4Counter) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv4PatternIpv4Counter) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIpv4PatternIpv4Counter interface {
	msg() *sanity.PatternIpv4PatternIpv4Counter
	ToYaml() string
	ToJson() string
	Start() string
	SetStart(value string) PatternIpv4PatternIpv4Counter
	Step() string
	SetStep(value string) PatternIpv4PatternIpv4Counter
	Count() int32
	SetCount(value int32) PatternIpv4PatternIpv4Counter
}

// Start returns a string
//  description is TBD
func (obj *patternIpv4PatternIpv4Counter) Start() string {
	return *obj.obj.Start
}

// SetStart sets the string value in the None object
//  description is TBD
func (obj *patternIpv4PatternIpv4Counter) SetStart(value string) PatternIpv4PatternIpv4Counter {
	obj.obj.Start = &value
	return obj
}

// Step returns a string
//  description is TBD
func (obj *patternIpv4PatternIpv4Counter) Step() string {
	return *obj.obj.Step
}

// SetStep sets the string value in the None object
//  description is TBD
func (obj *patternIpv4PatternIpv4Counter) SetStep(value string) PatternIpv4PatternIpv4Counter {
	obj.obj.Step = &value
	return obj
}

// Count returns a int32
//  description is TBD
func (obj *patternIpv4PatternIpv4Counter) Count() int32 {
	return *obj.obj.Count
}

// SetCount sets the int32 value in the None object
//  description is TBD
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

func (obj *patternIpv6PatternIpv6Counter) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6Counter) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIpv6PatternIpv6Counter) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIpv6PatternIpv6Counter) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIpv6PatternIpv6Counter interface {
	msg() *sanity.PatternIpv6PatternIpv6Counter
	ToYaml() string
	ToJson() string
	Start() string
	SetStart(value string) PatternIpv6PatternIpv6Counter
	Step() string
	SetStep(value string) PatternIpv6PatternIpv6Counter
	Count() int32
	SetCount(value int32) PatternIpv6PatternIpv6Counter
}

// Start returns a string
//  description is TBD
func (obj *patternIpv6PatternIpv6Counter) Start() string {
	return *obj.obj.Start
}

// SetStart sets the string value in the None object
//  description is TBD
func (obj *patternIpv6PatternIpv6Counter) SetStart(value string) PatternIpv6PatternIpv6Counter {
	obj.obj.Start = &value
	return obj
}

// Step returns a string
//  description is TBD
func (obj *patternIpv6PatternIpv6Counter) Step() string {
	return *obj.obj.Step
}

// SetStep sets the string value in the None object
//  description is TBD
func (obj *patternIpv6PatternIpv6Counter) SetStep(value string) PatternIpv6PatternIpv6Counter {
	obj.obj.Step = &value
	return obj
}

// Count returns a int32
//  description is TBD
func (obj *patternIpv6PatternIpv6Counter) Count() int32 {
	return *obj.obj.Count
}

// SetCount sets the int32 value in the None object
//  description is TBD
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

func (obj *patternMacPatternMacCounter) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMacCounter) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternMacPatternMacCounter) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternMacPatternMacCounter) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternMacPatternMacCounter interface {
	msg() *sanity.PatternMacPatternMacCounter
	ToYaml() string
	ToJson() string
	Start() string
	SetStart(value string) PatternMacPatternMacCounter
	Step() string
	SetStep(value string) PatternMacPatternMacCounter
	Count() int32
	SetCount(value int32) PatternMacPatternMacCounter
}

// Start returns a string
//  description is TBD
func (obj *patternMacPatternMacCounter) Start() string {
	return *obj.obj.Start
}

// SetStart sets the string value in the None object
//  description is TBD
func (obj *patternMacPatternMacCounter) SetStart(value string) PatternMacPatternMacCounter {
	obj.obj.Start = &value
	return obj
}

// Step returns a string
//  description is TBD
func (obj *patternMacPatternMacCounter) Step() string {
	return *obj.obj.Step
}

// SetStep sets the string value in the None object
//  description is TBD
func (obj *patternMacPatternMacCounter) SetStep(value string) PatternMacPatternMacCounter {
	obj.obj.Step = &value
	return obj
}

// Count returns a int32
//  description is TBD
func (obj *patternMacPatternMacCounter) Count() int32 {
	return *obj.obj.Count
}

// SetCount sets the int32 value in the None object
//  description is TBD
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

func (obj *patternIntegerPatternIntegerCounter) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternIntegerCounter) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *patternIntegerPatternIntegerCounter) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *patternIntegerPatternIntegerCounter) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type PatternIntegerPatternIntegerCounter interface {
	msg() *sanity.PatternIntegerPatternIntegerCounter
	ToYaml() string
	ToJson() string
	Start() int32
	SetStart(value int32) PatternIntegerPatternIntegerCounter
	Step() int32
	SetStep(value int32) PatternIntegerPatternIntegerCounter
	Count() int32
	SetCount(value int32) PatternIntegerPatternIntegerCounter
}

// Start returns a int32
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) Start() int32 {
	return *obj.obj.Start
}

// SetStart sets the int32 value in the None object
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) SetStart(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Start = &value
	return obj
}

// Step returns a int32
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) Step() int32 {
	return *obj.obj.Step
}

// SetStep sets the int32 value in the None object
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) SetStep(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Step = &value
	return obj
}

// Count returns a int32
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) Count() int32 {
	return *obj.obj.Count
}

// SetCount sets the int32 value in the None object
//  description is TBD
func (obj *patternIntegerPatternIntegerCounter) SetCount(value int32) PatternIntegerPatternIntegerCounter {
	obj.obj.Count = &value
	return obj
}

type setConfigResponseStatusCode200 struct {
	obj *sanity.SetConfigResponse_StatusCode200
}

func (obj *setConfigResponseStatusCode200) msg() *sanity.SetConfigResponse_StatusCode200 {
	return obj.obj
}

func (obj *setConfigResponseStatusCode200) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *setConfigResponseStatusCode200) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *setConfigResponseStatusCode200) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *setConfigResponseStatusCode200) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type SetConfigResponse_StatusCode200 interface {
	msg() *sanity.SetConfigResponse_StatusCode200
	ToYaml() string
	ToJson() string
}

type updateConfigResponseStatusCode200 struct {
	obj *sanity.UpdateConfigResponse_StatusCode200
}

func (obj *updateConfigResponseStatusCode200) msg() *sanity.UpdateConfigResponse_StatusCode200 {
	return obj.obj
}

func (obj *updateConfigResponseStatusCode200) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *updateConfigResponseStatusCode200) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *updateConfigResponseStatusCode200) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *updateConfigResponseStatusCode200) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type UpdateConfigResponse_StatusCode200 interface {
	msg() *sanity.UpdateConfigResponse_StatusCode200
	ToYaml() string
	ToJson() string
}

type getConfigResponseStatusCode200 struct {
	obj *sanity.GetConfigResponse_StatusCode200
}

func (obj *getConfigResponseStatusCode200) msg() *sanity.GetConfigResponse_StatusCode200 {
	return obj.obj
}

func (obj *getConfigResponseStatusCode200) ToYaml() string {
	data, _ := yaml.Marshal(obj.msg())
	return string(data)
}

func (obj *getConfigResponseStatusCode200) FromYaml(value string) error {
	return yaml.Unmarshal([]byte(value), obj.msg())
}

func (obj *getConfigResponseStatusCode200) ToJson() string {
	data, _ := json.Marshal(obj.msg())
	return string(data)
}

func (obj *getConfigResponseStatusCode200) FromJson(value string) error {
	return json.Unmarshal([]byte(value), obj.msg())
}

type GetConfigResponse_StatusCode200 interface {
	msg() *sanity.GetConfigResponse_StatusCode200
	ToYaml() string
	ToJson() string
}
