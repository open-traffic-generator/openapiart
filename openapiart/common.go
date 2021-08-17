import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

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
	HasGrpcTransport() bool
	NewHttpTransport() HttpTransport
	HasHttpTransport() bool
}

// NewGrpcTransport sets the underlying transport of the Api as grpc
func (api *api) NewGrpcTransport() GrpcTransport {
	api.grpc = &grpcTransport{
		location:       "127.0.0.1:5050",
		requestTimeout: time.Duration(10),
	}
	api.http = nil
	return api.grpc
}

func (api *api) HasGrpcTransport() bool {
	return api.grpc != nil
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

func (api *api) HasHttpTransport() bool {
	return api.http != nil
}

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client HttpRequestDoer
	ctx    context.Context
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
