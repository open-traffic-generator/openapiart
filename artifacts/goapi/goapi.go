/* OpenAPIArt Test API 0.0.1
 * Demonstrates the features of the OpenAPIArt package
 * License: NO-LICENSE-PRESENT */

package goapi

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/ghodss/yaml"
	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type versionMeta struct {
	checkVersion  bool
	localVersion  Version
	remoteVersion Version
	checkError    error
}
type goapiApi struct {
	api
	grpcClient  openapi.OpenapiClient
	httpClient  httpClient
	versionMeta *versionMeta
}

// grpcConnect builds up a grpc connection
func (api *goapiApi) grpcConnect() error {
	if api.grpcClient == nil {
		if api.grpc.clientConnection == nil {
			ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.dialTimeout)
			defer cancelFunc()
			conn, err := grpc.DialContext(ctx, api.grpc.location, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			api.grpcClient = openapi.NewOpenapiClient(conn)
			api.grpc.clientConnection = conn
		} else {
			api.grpcClient = openapi.NewOpenapiClient(api.grpc.clientConnection)
		}
	}
	return nil
}

func (api *goapiApi) grpcClose() error {
	if api.grpc != nil {
		if api.grpc.clientConnection != nil {
			err := api.grpc.clientConnection.Close()
			if err != nil {
				return err
			}
		}
	}
	api.grpcClient = nil
	api.grpc = nil
	return nil
}

func (api *goapiApi) Close() error {
	if api.hasGrpcTransport() {
		err := api.grpcClose()
		return err
	}
	if api.hasHttpTransport() {
		err := api.http.conn.(*net.TCPConn).SetLinger(0)
		api.http.conn.Close()
		api.http.conn = nil
		api.http = nil
		api.httpClient.client = nil
		return err
	}
	return nil
}

// NewApi returns a new instance of the top level interface hierarchy
func NewApi() GoapiApi {
	api := goapiApi{}
	api.versionMeta = &versionMeta{checkVersion: false}
	return &api
}

// httpConnect builds up a http connection
func (api *goapiApi) httpConnect() error {
	if api.httpClient.client == nil {
		tr := http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				tcpConn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
				if err != nil {
					return nil, err
				}
				tlsConn := tls.Client(tcpConn, &tls.Config{InsecureSkipVerify: !api.http.verify})
				err = tlsConn.Handshake()
				if err != nil {
					return nil, err
				}
				api.http.conn = tcpConn
				return tlsConn, nil
			},
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				tcpConn, err := (&net.Dialer{}).DialContext(ctx, network, addr)
				if err != nil {
					return nil, err
				}
				api.http.conn = tcpConn
				return tcpConn, nil
			},
		}
		client := httpClient{
			client: &http.Client{
				Transport: &tr,
			},
			ctx: context.Background(),
		}
		api.httpClient = client
	}
	return nil
}

func (api *goapiApi) httpSendRecv(urlPath string, jsonBody string, method string) (*http.Response, error) {
	err := api.httpConnect()
	if err != nil {
		return nil, err
	}
	httpClient := api.httpClient
	var bodyReader = bytes.NewReader([]byte(jsonBody))
	queryUrl, err := url.Parse(api.http.location)
	if err != nil {
		return nil, err
	}
	queryUrl, _ = queryUrl.Parse(urlPath)
	req, _ := http.NewRequest(method, queryUrl.String(), bodyReader)
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(httpClient.ctx)
	response, err := httpClient.client.Do(req)
	return response, err
}

// GoapiApi demonstrates the features of the OpenAPIArt package
type GoapiApi interface {
	Api
	// PrefixConfig is container which retains the configuration
	// NewPrefixConfig returns a new instance of PrefixConfig.
	NewPrefixConfig() PrefixConfig
	// UpdateConfig is under Review: the whole schema is being reviewed
	//
	// Object to Test required Parameter
	// NewUpdateConfig returns a new instance of UpdateConfig.
	NewUpdateConfig() UpdateConfig
	// MetricsRequest is description is TBD
	// NewMetricsRequest returns a new instance of MetricsRequest.
	NewMetricsRequest() MetricsRequest
	// ApiTestInputBody is description is TBD
	// NewApiTestInputBody returns a new instance of ApiTestInputBody.
	NewApiTestInputBody() ApiTestInputBody
	// SetConfigResponse is description is TBD
	// NewSetConfigResponse returns a new instance of SetConfigResponse.
	NewSetConfigResponse() SetConfigResponse
	// UpdateConfigurationResponse is description is TBD
	// NewUpdateConfigurationResponse returns a new instance of UpdateConfigurationResponse.
	NewUpdateConfigurationResponse() UpdateConfigurationResponse
	// GetConfigResponse is description is TBD
	// NewGetConfigResponse returns a new instance of GetConfigResponse.
	NewGetConfigResponse() GetConfigResponse
	// GetMetricsResponse is description is TBD
	// NewGetMetricsResponse returns a new instance of GetMetricsResponse.
	NewGetMetricsResponse() GetMetricsResponse
	// GetWarningsResponse is description is TBD
	// NewGetWarningsResponse returns a new instance of GetWarningsResponse.
	NewGetWarningsResponse() GetWarningsResponse
	// ClearWarningsResponse is description is TBD
	// NewClearWarningsResponse returns a new instance of ClearWarningsResponse.
	NewClearWarningsResponse() ClearWarningsResponse
	// GetRootResponseResponse is description is TBD
	// NewGetRootResponseResponse returns a new instance of GetRootResponseResponse.
	NewGetRootResponseResponse() GetRootResponseResponse
	// DummyResponseTestResponse is description is TBD
	// NewDummyResponseTestResponse returns a new instance of DummyResponseTestResponse.
	NewDummyResponseTestResponse() DummyResponseTestResponse
	// PostRootResponseResponse is description is TBD
	// NewPostRootResponseResponse returns a new instance of PostRootResponseResponse.
	NewPostRootResponseResponse() PostRootResponseResponse
	// GetAllItemsResponse is description is TBD
	// NewGetAllItemsResponse returns a new instance of GetAllItemsResponse.
	NewGetAllItemsResponse() GetAllItemsResponse
	// GetSingleItemResponse is description is TBD
	// NewGetSingleItemResponse returns a new instance of GetSingleItemResponse.
	NewGetSingleItemResponse() GetSingleItemResponse
	// GetSingleItemLevel2Response is description is TBD
	// NewGetSingleItemLevel2Response returns a new instance of GetSingleItemLevel2Response.
	NewGetSingleItemLevel2Response() GetSingleItemLevel2Response
	// GetVersionResponse is description is TBD
	// NewGetVersionResponse returns a new instance of GetVersionResponse.
	NewGetVersionResponse() GetVersionResponse
	// SetConfig sets configuration resources.
	SetConfig(prefixConfig PrefixConfig) ([]byte, error)
	// UpdateConfiguration deprecated: please use post instead
	//
	// Sets configuration resources.
	UpdateConfiguration(updateConfig UpdateConfig) (PrefixConfig, error)
	// GetConfig gets the configuration resources.
	GetConfig() (PrefixConfig, error)
	// GetMetrics gets metrics.
	GetMetrics(metricsRequest MetricsRequest) (Metrics, error)
	// GetWarnings gets warnings.
	GetWarnings() (WarningDetails, error)
	// ClearWarnings clears warnings.
	ClearWarnings() (*string, error)
	// GetRootResponse simple GET api with single return type
	GetRootResponse() (CommonResponseSuccess, error)
	// DummyResponseTest description is TBD
	DummyResponseTest() (*string, error)
	// PostRootResponse simple POST api with single return type
	PostRootResponse(apiTestInputBody ApiTestInputBody) (CommonResponseSuccess, error)
	// GetAllItems return list of some items
	GetAllItems() (ServiceAbcItemList, error)
	// GetSingleItem return single item
	GetSingleItem() (ServiceAbcItem, error)
	// GetSingleItemLevel2 return single item
	GetSingleItemLevel2() (ServiceAbcItem, error)
	// GetVersion description is TBD
	GetVersion() (Version, error)
	// GetLocalVersion provides version details of local client
	GetLocalVersion() Version
	// GetRemoteVersion provides version details received from remote server
	GetRemoteVersion() (Version, error)
	// SetVersionCompatibilityCheck allows enabling or disabling automatic version
	// compatibility check between client and server API spec version upon API call
	SetVersionCompatibilityCheck(bool)
	// CheckVersionCompatibility compares API spec version for local client and remote server,
	// and returns an error if they are not compatible according to Semantic Versioning 2.0.0
	CheckVersionCompatibility() error
}

func (api *goapiApi) NewPrefixConfig() PrefixConfig {
	return NewPrefixConfig()
}

func (api *goapiApi) NewUpdateConfig() UpdateConfig {
	return NewUpdateConfig()
}

func (api *goapiApi) NewMetricsRequest() MetricsRequest {
	return NewMetricsRequest()
}

func (api *goapiApi) NewApiTestInputBody() ApiTestInputBody {
	return NewApiTestInputBody()
}

func (api *goapiApi) NewSetConfigResponse() SetConfigResponse {
	return NewSetConfigResponse()
}

func (api *goapiApi) NewUpdateConfigurationResponse() UpdateConfigurationResponse {
	return NewUpdateConfigurationResponse()
}

func (api *goapiApi) NewGetConfigResponse() GetConfigResponse {
	return NewGetConfigResponse()
}

func (api *goapiApi) NewGetMetricsResponse() GetMetricsResponse {
	return NewGetMetricsResponse()
}

func (api *goapiApi) NewGetWarningsResponse() GetWarningsResponse {
	return NewGetWarningsResponse()
}

func (api *goapiApi) NewClearWarningsResponse() ClearWarningsResponse {
	return NewClearWarningsResponse()
}

func (api *goapiApi) NewGetRootResponseResponse() GetRootResponseResponse {
	return NewGetRootResponseResponse()
}

func (api *goapiApi) NewDummyResponseTestResponse() DummyResponseTestResponse {
	return NewDummyResponseTestResponse()
}

func (api *goapiApi) NewPostRootResponseResponse() PostRootResponseResponse {
	return NewPostRootResponseResponse()
}

func (api *goapiApi) NewGetAllItemsResponse() GetAllItemsResponse {
	return NewGetAllItemsResponse()
}

func (api *goapiApi) NewGetSingleItemResponse() GetSingleItemResponse {
	return NewGetSingleItemResponse()
}

func (api *goapiApi) NewGetSingleItemLevel2Response() GetSingleItemLevel2Response {
	return NewGetSingleItemLevel2Response()
}

func (api *goapiApi) NewGetVersionResponse() GetVersionResponse {
	return NewGetVersionResponse()
}

func (api *goapiApi) GetLocalVersion() Version {
	if api.versionMeta.localVersion == nil {
		api.versionMeta.localVersion = NewVersion().SetApiSpecVersion("0.0.1").SetSdkVersion("0.0.1")
	}

	return api.versionMeta.localVersion
}

func (api *goapiApi) GetRemoteVersion() (Version, error) {
	if api.versionMeta.remoteVersion == nil {
		v, err := api.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("could not fetch remote version: %v", err)
		}

		api.versionMeta.remoteVersion = v
	}

	return api.versionMeta.remoteVersion, nil
}

func (api *goapiApi) SetVersionCompatibilityCheck(v bool) {
	api.versionMeta.checkVersion = v
}

func (api *goapiApi) checkLocalRemoteVersionCompatibility() (error, error) {
	localVer := api.GetLocalVersion()
	remoteVer, err := api.GetRemoteVersion()
	if err != nil {
		return nil, err
	}
	err = checkClientServerVersionCompatibility(localVer.ApiSpecVersion(), remoteVer.ApiSpecVersion(), "API spec")
	if err != nil {
		return fmt.Errorf(
			"client SDK version '%s' is not compatible with server SDK version '%s': %v",
			localVer.SdkVersion(), remoteVer.SdkVersion(), err,
		), nil
	}

	return nil, nil
}

func (api *goapiApi) checkLocalRemoteVersionCompatibilityOnce() error {
	if !api.versionMeta.checkVersion {
		return nil
	}

	if api.versionMeta.checkError != nil {
		return api.versionMeta.checkError
	}

	compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
	if compatErr != nil {
		api.versionMeta.checkError = compatErr
		return compatErr
	}
	if apiErr != nil {
		api.versionMeta.checkError = nil
		return apiErr
	}

	api.versionMeta.checkVersion = false
	api.versionMeta.checkError = nil
	return nil
}

func (api *goapiApi) CheckVersionCompatibility() error {
	compatErr, apiErr := api.checkLocalRemoteVersionCompatibility()
	if compatErr != nil {
		return fmt.Errorf("version error: %v", compatErr)
	}
	if apiErr != nil {
		return apiErr
	}

	return nil
}

func (api *goapiApi) SetConfig(prefixConfig PrefixConfig) ([]byte, error) {

	if err := prefixConfig.Validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpSetConfig(prefixConfig)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := openapi.SetConfigRequest{PrefixConfig: prefixConfig.Msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	if resp.ResponseBytes != nil {
		return resp.ResponseBytes, nil
	}
	return nil, nil
}

func (api *goapiApi) UpdateConfiguration(updateConfig UpdateConfig) (PrefixConfig, error) {
	api.addWarnings("UpdateConfiguration api is deprecated, please use post instead")

	if err := updateConfig.Validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpUpdateConfiguration(updateConfig)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := openapi.UpdateConfigurationRequest{UpdateConfig: updateConfig.Msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.UpdateConfiguration(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewPrefixConfig()
	if resp.GetPrefixConfig() != nil {
		return ret.SetMsg(resp.GetPrefixConfig()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetConfig() (PrefixConfig, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetConfig()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetConfig(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewPrefixConfig()
	if resp.GetPrefixConfig() != nil {
		return ret.SetMsg(resp.GetPrefixConfig()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetMetrics(metricsRequest MetricsRequest) (Metrics, error) {

	if err := metricsRequest.Validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetMetrics(metricsRequest)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := openapi.GetMetricsRequest{MetricsRequest: metricsRequest.Msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetMetrics(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewMetrics()
	if resp.GetMetrics() != nil {
		return ret.SetMsg(resp.GetMetrics()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetWarnings() (WarningDetails, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetWarnings()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetWarnings(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarningDetails()
	if resp.GetWarningDetails() != nil {
		return ret.SetMsg(resp.GetWarningDetails()), nil
	}

	return ret, nil
}

func (api *goapiApi) ClearWarnings() (*string, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpClearWarnings()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.ClearWarnings(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	if resp.GetString_() != "" {
		status_code_value := resp.GetString_()
		return &status_code_value, nil
	}
	return nil, nil
}

func (api *goapiApi) GetRootResponse() (CommonResponseSuccess, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetRootResponse()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetRootResponse(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewCommonResponseSuccess()
	if resp.GetCommonResponseSuccess() != nil {
		return ret.SetMsg(resp.GetCommonResponseSuccess()), nil
	}

	return ret, nil
}

func (api *goapiApi) DummyResponseTest() (*string, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpDummyResponseTest()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.DummyResponseTest(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	if resp.GetString_() != "" {
		status_code_value := resp.GetString_()
		return &status_code_value, nil
	}
	return nil, nil
}

func (api *goapiApi) PostRootResponse(apiTestInputBody ApiTestInputBody) (CommonResponseSuccess, error) {

	if err := apiTestInputBody.Validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpPostRootResponse(apiTestInputBody)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := openapi.PostRootResponseRequest{ApiTestInputBody: apiTestInputBody.Msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.PostRootResponse(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewCommonResponseSuccess()
	if resp.GetCommonResponseSuccess() != nil {
		return ret.SetMsg(resp.GetCommonResponseSuccess()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetAllItems() (ServiceAbcItemList, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetAllItems()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetAllItems(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItemList()
	if resp.GetServiceAbcItemList() != nil {
		return ret.SetMsg(resp.GetServiceAbcItemList()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetSingleItem() (ServiceAbcItem, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetSingleItem()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetSingleItem(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItem()
	if resp.GetServiceAbcItem() != nil {
		return ret.SetMsg(resp.GetServiceAbcItem()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetSingleItemLevel2() (ServiceAbcItem, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetSingleItemLevel2()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetSingleItemLevel2(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItem()
	if resp.GetServiceAbcItem() != nil {
		return ret.SetMsg(resp.GetServiceAbcItem()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetVersion() (Version, error) {

	if api.hasHttpTransport() {
		return api.httpGetVersion()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetVersion(ctx, &request)
	if err != nil {
		if er, ok := api.fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewVersion()
	if resp.GetVersion() != nil {
		return ret.SetMsg(resp.GetVersion()), nil
	}

	return ret, nil
}

func (api *goapiApi) httpSetConfig(prefixConfig PrefixConfig) ([]byte, error) {
	prefixConfigJson, err := prefixConfig.ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("api/config", prefixConfigJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		return bodyBytes, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpUpdateConfiguration(updateConfig UpdateConfig) (PrefixConfig, error) {
	updateConfigJson, err := updateConfig.ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("api/config", updateConfigJson, "PATCH")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewUpdateConfigurationResponse().PrefixConfig()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetConfig() (PrefixConfig, error) {
	resp, err := api.httpSendRecv("api/config", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetConfigResponse().PrefixConfig()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetMetrics(metricsRequest MetricsRequest) (Metrics, error) {
	metricsRequestJson, err := metricsRequest.ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("api/metrics", metricsRequestJson, "GET")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetMetricsResponse().Metrics()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetWarnings() (WarningDetails, error) {
	resp, err := api.httpSendRecv("api/warnings", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetWarningsResponse().WarningDetails()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpClearWarnings() (*string, error) {
	resp, err := api.httpSendRecv("api/warnings", "", "DELETE")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		bodyString := string(bodyBytes)
		return &bodyString, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetRootResponse() (CommonResponseSuccess, error) {
	resp, err := api.httpSendRecv("api/apitest", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetRootResponseResponse().CommonResponseSuccess()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpDummyResponseTest() (*string, error) {
	resp, err := api.httpSendRecv("api/apitest", "", "DELETE")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		bodyString := string(bodyBytes)
		return &bodyString, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpPostRootResponse(apiTestInputBody ApiTestInputBody) (CommonResponseSuccess, error) {
	apiTestInputBodyJson, err := apiTestInputBody.ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("api/apitest", apiTestInputBodyJson, "POST")

	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewPostRootResponseResponse().CommonResponseSuccess()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetAllItems() (ServiceAbcItemList, error) {
	resp, err := api.httpSendRecv("api/serviceb", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetAllItemsResponse().ServiceAbcItemList()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetSingleItem() (ServiceAbcItem, error) {
	resp, err := api.httpSendRecv("api/serviceb/{item_id}", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetSingleItemResponse().ServiceAbcItem()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetSingleItemLevel2() (ServiceAbcItem, error) {
	resp, err := api.httpSendRecv("api/serviceb/{item_id}/{level_2}", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetSingleItemLevel2Response().ServiceAbcItem()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetVersion() (Version, error) {
	resp, err := api.httpSendRecv("api/capabilities/version", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := api.NewGetVersionResponse().Version()
		if err := obj.FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, api.fromHttpError(resp.StatusCode, bodyBytes)
	}
}

// ***** PrefixConfig *****
type prefixConfig struct {
	validation
	obj                        *openapi.PrefixConfig
	requiredObjectHolder       EObject
	optionalObjectHolder       EObject
	eHolder                    EObject
	fHolder                    FObject
	gHolder                    PrefixConfigGObjectIter
	jHolder                    PrefixConfigJObjectIter
	kHolder                    KObject
	lHolder                    LObject
	levelHolder                LevelOne
	mandatoryHolder            Mandate
	ipv4PatternHolder          Ipv4Pattern
	ipv6PatternHolder          Ipv6Pattern
	macPatternHolder           MacPattern
	integerPatternHolder       IntegerPattern
	checksumPatternHolder      ChecksumPattern
	caseHolder                 Layer1Ieee802X
	mObjectHolder              MObject
	headerChecksumHolder       PatternPrefixConfigHeaderChecksum
	autoFieldTestHolder        PatternPrefixConfigAutoFieldTest
	wListHolder                PrefixConfigWObjectIter
	xListHolder                PrefixConfigZObjectIter
	zObjectHolder              ZObject
	yObjectHolder              YObject
	choiceObjectHolder         PrefixConfigChoiceObjectIter
	requiredChoiceObjectHolder RequiredChoiceParent
	g1Holder                   PrefixConfigGObjectIter
	g2Holder                   PrefixConfigGObjectIter
}

func NewPrefixConfig() PrefixConfig {
	obj := prefixConfig{obj: &openapi.PrefixConfig{}}
	obj.setDefault()
	return &obj
}

func (obj *prefixConfig) Msg() *openapi.PrefixConfig {
	return obj.obj
}

func (obj *prefixConfig) SetMsg(msg *openapi.PrefixConfig) PrefixConfig {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *prefixConfig) ToProto() (*openapi.PrefixConfig, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *prefixConfig) FromProto(msg *openapi.PrefixConfig) (PrefixConfig, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *prefixConfig) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *prefixConfig) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *prefixConfig) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *prefixConfig) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *prefixConfig) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *prefixConfig) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
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

func (obj *prefixConfig) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *prefixConfig) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *prefixConfig) Clone() (PrefixConfig, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPrefixConfig()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
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
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PrefixConfig is container which retains the configuration
type PrefixConfig interface {
	Validation
	// Msg marshals PrefixConfig to protobuf object *openapi.PrefixConfig
	// and doesn't set defaults
	Msg() *openapi.PrefixConfig
	// SetMsg unmarshals PrefixConfig from protobuf object *openapi.PrefixConfig
	// and doesn't set defaults
	SetMsg(*openapi.PrefixConfig) PrefixConfig
	// ToProto marshals PrefixConfig to protobuf object *openapi.PrefixConfig
	ToProto() (*openapi.PrefixConfig, error)
	// ToPbText marshals PrefixConfig to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PrefixConfig to YAML text
	ToYaml() (string, error)
	// ToJson marshals PrefixConfig to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PrefixConfig from protobuf object *openapi.PrefixConfig
	FromProto(msg *openapi.PrefixConfig) (PrefixConfig, error)
	// FromPbText unmarshals PrefixConfig from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PrefixConfig from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PrefixConfig from JSON text
	FromJson(value string) error
	// Validate validates PrefixConfig
	Validate() error
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
	setNil()
}

// A required object that MUST be generated as such.
// RequiredObject returns a EObject
func (obj *prefixConfig) RequiredObject() EObject {
	if obj.obj.RequiredObject == nil {
		obj.obj.RequiredObject = NewEObject().Msg()
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
	obj.obj.RequiredObject = value.Msg()

	return obj
}

// An optional object that MUST be generated as such.
// OptionalObject returns a EObject
func (obj *prefixConfig) OptionalObject() EObject {
	if obj.obj.OptionalObject == nil {
		obj.obj.OptionalObject = NewEObject().Msg()
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
	obj.obj.OptionalObject = value.Msg()

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

	return obj.obj.A
}

// Under Review: Information TBD
//
// Small single line description
// SetA sets the string value in the PrefixConfig object
func (obj *prefixConfig) SetA(value string) PrefixConfig {

	obj.obj.A = value
	return obj
}

// Longer multi-line description
// Second line is here
// Third line
// B returns a float32
func (obj *prefixConfig) B() float32 {

	return obj.obj.B
}

// Longer multi-line description
// Second line is here
// Third line
// SetB sets the float32 value in the PrefixConfig object
func (obj *prefixConfig) SetB(value float32) PrefixConfig {

	obj.obj.B = value
	return obj
}

// description is TBD
// C returns a int32
func (obj *prefixConfig) C() int32 {

	return obj.obj.C
}

// description is TBD
// SetC sets the int32 value in the PrefixConfig object
func (obj *prefixConfig) SetC(value int32) PrefixConfig {

	obj.obj.C = value
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
		obj.obj.E = NewEObject().Msg()
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
	obj.obj.E = value.Msg()

	return obj
}

// An object with only choice(s)
// F returns a FObject
func (obj *prefixConfig) F() FObject {
	if obj.obj.F == nil {
		obj.obj.F = NewFObject().Msg()
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
	obj.obj.F = value.Msg()

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
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.gObjectSlice = append(obj.gObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigGObjectIter) Set(index int, newObj GObject) PrefixConfigGObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
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
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.jObjectSlice = append(obj.jObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigJObjectIter) Set(index int, newObj JObject) PrefixConfigJObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
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
		obj.obj.K = NewKObject().Msg()
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
	obj.obj.K = value.Msg()

	return obj
}

// description is TBD
// L returns a LObject
func (obj *prefixConfig) L() LObject {
	if obj.obj.L == nil {
		obj.obj.L = NewLObject().Msg()
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
	obj.obj.L = value.Msg()

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
		obj.obj.Level = NewLevelOne().Msg()
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
	obj.obj.Level = value.Msg()

	return obj
}

// description is TBD
// Mandatory returns a Mandate
func (obj *prefixConfig) Mandatory() Mandate {
	if obj.obj.Mandatory == nil {
		obj.obj.Mandatory = NewMandate().Msg()
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
	obj.obj.Mandatory = value.Msg()

	return obj
}

// description is TBD
// Ipv4Pattern returns a Ipv4Pattern
func (obj *prefixConfig) Ipv4Pattern() Ipv4Pattern {
	if obj.obj.Ipv4Pattern == nil {
		obj.obj.Ipv4Pattern = NewIpv4Pattern().Msg()
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
	obj.obj.Ipv4Pattern = value.Msg()

	return obj
}

// description is TBD
// Ipv6Pattern returns a Ipv6Pattern
func (obj *prefixConfig) Ipv6Pattern() Ipv6Pattern {
	if obj.obj.Ipv6Pattern == nil {
		obj.obj.Ipv6Pattern = NewIpv6Pattern().Msg()
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
	obj.obj.Ipv6Pattern = value.Msg()

	return obj
}

// description is TBD
// MacPattern returns a MacPattern
func (obj *prefixConfig) MacPattern() MacPattern {
	if obj.obj.MacPattern == nil {
		obj.obj.MacPattern = NewMacPattern().Msg()
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
	obj.obj.MacPattern = value.Msg()

	return obj
}

// description is TBD
// IntegerPattern returns a IntegerPattern
func (obj *prefixConfig) IntegerPattern() IntegerPattern {
	if obj.obj.IntegerPattern == nil {
		obj.obj.IntegerPattern = NewIntegerPattern().Msg()
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
	obj.obj.IntegerPattern = value.Msg()

	return obj
}

// description is TBD
// ChecksumPattern returns a ChecksumPattern
func (obj *prefixConfig) ChecksumPattern() ChecksumPattern {
	if obj.obj.ChecksumPattern == nil {
		obj.obj.ChecksumPattern = NewChecksumPattern().Msg()
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
	obj.obj.ChecksumPattern = value.Msg()

	return obj
}

// description is TBD
// Case returns a Layer1Ieee802X
func (obj *prefixConfig) Case() Layer1Ieee802X {
	if obj.obj.Case == nil {
		obj.obj.Case = NewLayer1Ieee802X().Msg()
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
	obj.obj.Case = value.Msg()

	return obj
}

// description is TBD
// MObject returns a MObject
func (obj *prefixConfig) MObject() MObject {
	if obj.obj.MObject == nil {
		obj.obj.MObject = NewMObject().Msg()
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
	obj.obj.MObject = value.Msg()

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
		obj.obj.HeaderChecksum = NewPatternPrefixConfigHeaderChecksum().Msg()
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
	obj.obj.HeaderChecksum = value.Msg()

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
		obj.obj.AutoFieldTest = NewPatternPrefixConfigAutoFieldTest().Msg()
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
	obj.obj.AutoFieldTest = value.Msg()

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
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.wObjectSlice = append(obj.wObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigWObjectIter) Set(index int, newObj WObject) PrefixConfigWObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
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
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.zObjectSlice = append(obj.zObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigZObjectIter) Set(index int, newObj ZObject) PrefixConfigZObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
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
		obj.obj.ZObject = NewZObject().Msg()
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
	obj.obj.ZObject = value.Msg()

	return obj
}

// description is TBD
// YObject returns a YObject
func (obj *prefixConfig) YObject() YObject {
	if obj.obj.YObject == nil {
		obj.obj.YObject = NewYObject().Msg()
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
	obj.obj.YObject = value.Msg()

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
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.choiceObjectSlice = append(obj.choiceObjectSlice, item)
	}
	return obj
}

func (obj *prefixConfigChoiceObjectIter) Set(index int, newObj ChoiceObject) PrefixConfigChoiceObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
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
		obj.obj.RequiredChoiceObject = NewRequiredChoiceParent().Msg()
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
	obj.obj.RequiredChoiceObject = value.Msg()

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
	if obj.obj.A == "" {
		vObj.validationErrors = append(vObj.validationErrors, "A is required field on interface PrefixConfig")
	}

	// A is under_review
	if obj.obj.A != "" {
		obj.addWarnings("A property in schema PrefixConfig is under review, Information TBD")
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

}

func (obj *prefixConfig) setDefault() {
	if obj.obj.Response == nil {
		obj.SetResponse(PrefixConfigResponse.STATUS_200)

	}
	if obj.obj.H == nil {
		obj.SetH(true)
	}

}

// ***** UpdateConfig *****
type updateConfig struct {
	validation
	obj     *openapi.UpdateConfig
	gHolder UpdateConfigGObjectIter
}

func NewUpdateConfig() UpdateConfig {
	obj := updateConfig{obj: &openapi.UpdateConfig{}}
	obj.setDefault()
	return &obj
}

func (obj *updateConfig) Msg() *openapi.UpdateConfig {
	return obj.obj
}

func (obj *updateConfig) SetMsg(msg *openapi.UpdateConfig) UpdateConfig {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *updateConfig) ToProto() (*openapi.UpdateConfig, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *updateConfig) FromProto(msg *openapi.UpdateConfig) (UpdateConfig, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *updateConfig) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *updateConfig) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *updateConfig) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *updateConfig) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *updateConfig) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *updateConfig) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *updateConfig) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateConfig) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateConfig) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateConfig) Clone() (UpdateConfig, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateConfig()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *updateConfig) setNil() {
	obj.gHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateConfig is under Review: the whole schema is being reviewed
//
// Object to Test required Parameter
type UpdateConfig interface {
	Validation
	// Msg marshals UpdateConfig to protobuf object *openapi.UpdateConfig
	// and doesn't set defaults
	Msg() *openapi.UpdateConfig
	// SetMsg unmarshals UpdateConfig from protobuf object *openapi.UpdateConfig
	// and doesn't set defaults
	SetMsg(*openapi.UpdateConfig) UpdateConfig
	// ToProto marshals UpdateConfig to protobuf object *openapi.UpdateConfig
	ToProto() (*openapi.UpdateConfig, error)
	// ToPbText marshals UpdateConfig to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateConfig to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateConfig to JSON text
	ToJson() (string, error)
	// FromProto unmarshals UpdateConfig from protobuf object *openapi.UpdateConfig
	FromProto(msg *openapi.UpdateConfig) (UpdateConfig, error)
	// FromPbText unmarshals UpdateConfig from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateConfig from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateConfig from JSON text
	FromJson(value string) error
	// Validate validates UpdateConfig
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateConfig, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// G returns UpdateConfigGObjectIterIter, set in UpdateConfig
	G() UpdateConfigGObjectIter
	setNil()
}

// A list of objects with choice and properties
// G returns a []GObject
func (obj *updateConfig) G() UpdateConfigGObjectIter {
	if len(obj.obj.G) == 0 {
		obj.obj.G = []*openapi.GObject{}
	}
	if obj.gHolder == nil {
		obj.gHolder = newUpdateConfigGObjectIter(&obj.obj.G).setMsg(obj)
	}
	return obj.gHolder
}

type updateConfigGObjectIter struct {
	obj          *updateConfig
	gObjectSlice []GObject
	fieldPtr     *[]*openapi.GObject
}

func newUpdateConfigGObjectIter(ptr *[]*openapi.GObject) UpdateConfigGObjectIter {
	return &updateConfigGObjectIter{fieldPtr: ptr}
}

type UpdateConfigGObjectIter interface {
	setMsg(*updateConfig) UpdateConfigGObjectIter
	Items() []GObject
	Add() GObject
	Append(items ...GObject) UpdateConfigGObjectIter
	Set(index int, newObj GObject) UpdateConfigGObjectIter
	Clear() UpdateConfigGObjectIter
	clearHolderSlice() UpdateConfigGObjectIter
	appendHolderSlice(item GObject) UpdateConfigGObjectIter
}

func (obj *updateConfigGObjectIter) setMsg(msg *updateConfig) UpdateConfigGObjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&gObject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *updateConfigGObjectIter) Items() []GObject {
	return obj.gObjectSlice
}

func (obj *updateConfigGObjectIter) Add() GObject {
	newObj := &openapi.GObject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &gObject{obj: newObj}
	newLibObj.setDefault()
	obj.gObjectSlice = append(obj.gObjectSlice, newLibObj)
	return newLibObj
}

func (obj *updateConfigGObjectIter) Append(items ...GObject) UpdateConfigGObjectIter {
	for _, item := range items {
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.gObjectSlice = append(obj.gObjectSlice, item)
	}
	return obj
}

func (obj *updateConfigGObjectIter) Set(index int, newObj GObject) UpdateConfigGObjectIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
	obj.gObjectSlice[index] = newObj
	return obj
}
func (obj *updateConfigGObjectIter) Clear() UpdateConfigGObjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.GObject{}
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *updateConfigGObjectIter) clearHolderSlice() UpdateConfigGObjectIter {
	if len(obj.gObjectSlice) > 0 {
		obj.gObjectSlice = []GObject{}
	}
	return obj
}
func (obj *updateConfigGObjectIter) appendHolderSlice(item GObject) UpdateConfigGObjectIter {
	obj.gObjectSlice = append(obj.gObjectSlice, item)
	return obj
}

func (obj *updateConfig) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	obj.addWarnings("UpdateConfig is under review, the whole schema is being reviewed")

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

}

func (obj *updateConfig) setDefault() {

}

// ***** MetricsRequest *****
type metricsRequest struct {
	validation
	obj *openapi.MetricsRequest
}

func NewMetricsRequest() MetricsRequest {
	obj := metricsRequest{obj: &openapi.MetricsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *metricsRequest) Msg() *openapi.MetricsRequest {
	return obj.obj
}

func (obj *metricsRequest) SetMsg(msg *openapi.MetricsRequest) MetricsRequest {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *metricsRequest) ToProto() (*openapi.MetricsRequest, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *metricsRequest) FromProto(msg *openapi.MetricsRequest) (MetricsRequest, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *metricsRequest) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *metricsRequest) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *metricsRequest) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *metricsRequest) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *metricsRequest) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *metricsRequest) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *metricsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metricsRequest) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metricsRequest) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metricsRequest) Clone() (MetricsRequest, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetricsRequest()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// MetricsRequest is description is TBD
type MetricsRequest interface {
	Validation
	// Msg marshals MetricsRequest to protobuf object *openapi.MetricsRequest
	// and doesn't set defaults
	Msg() *openapi.MetricsRequest
	// SetMsg unmarshals MetricsRequest from protobuf object *openapi.MetricsRequest
	// and doesn't set defaults
	SetMsg(*openapi.MetricsRequest) MetricsRequest
	// ToProto marshals MetricsRequest to protobuf object *openapi.MetricsRequest
	ToProto() (*openapi.MetricsRequest, error)
	// ToPbText marshals MetricsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MetricsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals MetricsRequest to JSON text
	ToJson() (string, error)
	// FromProto unmarshals MetricsRequest from protobuf object *openapi.MetricsRequest
	FromProto(msg *openapi.MetricsRequest) (MetricsRequest, error)
	// FromPbText unmarshals MetricsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MetricsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MetricsRequest from JSON text
	FromJson(value string) error
	// Validate validates MetricsRequest
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MetricsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsRequestChoiceEnum, set in MetricsRequest
	Choice() MetricsRequestChoiceEnum
	// SetChoice assigns MetricsRequestChoiceEnum provided by user to MetricsRequest
	SetChoice(value MetricsRequestChoiceEnum) MetricsRequest
	// HasChoice checks if Choice has been set in MetricsRequest
	HasChoice() bool
	// Port returns string, set in MetricsRequest.
	Port() string
	// SetPort assigns string provided by user to MetricsRequest
	SetPort(value string) MetricsRequest
	// HasPort checks if Port has been set in MetricsRequest
	HasPort() bool
	// Flow returns string, set in MetricsRequest.
	Flow() string
	// SetFlow assigns string provided by user to MetricsRequest
	SetFlow(value string) MetricsRequest
	// HasFlow checks if Flow has been set in MetricsRequest
	HasFlow() bool
}

type MetricsRequestChoiceEnum string

// Enum of Choice on MetricsRequest
var MetricsRequestChoice = struct {
	PORT MetricsRequestChoiceEnum
	FLOW MetricsRequestChoiceEnum
}{
	PORT: MetricsRequestChoiceEnum("port"),
	FLOW: MetricsRequestChoiceEnum("flow"),
}

func (obj *metricsRequest) Choice() MetricsRequestChoiceEnum {
	return MetricsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metricsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metricsRequest) SetChoice(value MetricsRequestChoiceEnum) MetricsRequest {
	intValue, ok := openapi.MetricsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.MetricsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flow = nil
	obj.obj.Port = nil
	return obj
}

// description is TBD
// Port returns a string
func (obj *metricsRequest) Port() string {

	if obj.obj.Port == nil {
		obj.SetChoice(MetricsRequestChoice.PORT)
	}

	return *obj.obj.Port

}

// description is TBD
// Port returns a string
func (obj *metricsRequest) HasPort() bool {
	return obj.obj.Port != nil
}

// description is TBD
// SetPort sets the string value in the MetricsRequest object
func (obj *metricsRequest) SetPort(value string) MetricsRequest {
	obj.SetChoice(MetricsRequestChoice.PORT)
	obj.obj.Port = &value
	return obj
}

// description is TBD
// Flow returns a string
func (obj *metricsRequest) Flow() string {

	if obj.obj.Flow == nil {
		obj.SetChoice(MetricsRequestChoice.FLOW)
	}

	return *obj.obj.Flow

}

// description is TBD
// Flow returns a string
func (obj *metricsRequest) HasFlow() bool {
	return obj.obj.Flow != nil
}

// description is TBD
// SetFlow sets the string value in the MetricsRequest object
func (obj *metricsRequest) SetFlow(value string) MetricsRequest {
	obj.SetChoice(MetricsRequestChoice.FLOW)
	obj.obj.Flow = &value
	return obj
}

func (obj *metricsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *metricsRequest) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(MetricsRequestChoice.PORT)

	}

}

// ***** ApiTestInputBody *****
type apiTestInputBody struct {
	validation
	obj *openapi.ApiTestInputBody
}

func NewApiTestInputBody() ApiTestInputBody {
	obj := apiTestInputBody{obj: &openapi.ApiTestInputBody{}}
	obj.setDefault()
	return &obj
}

func (obj *apiTestInputBody) Msg() *openapi.ApiTestInputBody {
	return obj.obj
}

func (obj *apiTestInputBody) SetMsg(msg *openapi.ApiTestInputBody) ApiTestInputBody {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *apiTestInputBody) ToProto() (*openapi.ApiTestInputBody, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *apiTestInputBody) FromProto(msg *openapi.ApiTestInputBody) (ApiTestInputBody, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *apiTestInputBody) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *apiTestInputBody) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *apiTestInputBody) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *apiTestInputBody) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *apiTestInputBody) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *apiTestInputBody) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *apiTestInputBody) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *apiTestInputBody) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *apiTestInputBody) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *apiTestInputBody) Clone() (ApiTestInputBody, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewApiTestInputBody()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// ApiTestInputBody is description is TBD
type ApiTestInputBody interface {
	Validation
	// Msg marshals ApiTestInputBody to protobuf object *openapi.ApiTestInputBody
	// and doesn't set defaults
	Msg() *openapi.ApiTestInputBody
	// SetMsg unmarshals ApiTestInputBody from protobuf object *openapi.ApiTestInputBody
	// and doesn't set defaults
	SetMsg(*openapi.ApiTestInputBody) ApiTestInputBody
	// ToProto marshals ApiTestInputBody to protobuf object *openapi.ApiTestInputBody
	ToProto() (*openapi.ApiTestInputBody, error)
	// ToPbText marshals ApiTestInputBody to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ApiTestInputBody to YAML text
	ToYaml() (string, error)
	// ToJson marshals ApiTestInputBody to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ApiTestInputBody from protobuf object *openapi.ApiTestInputBody
	FromProto(msg *openapi.ApiTestInputBody) (ApiTestInputBody, error)
	// FromPbText unmarshals ApiTestInputBody from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ApiTestInputBody from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ApiTestInputBody from JSON text
	FromJson(value string) error
	// Validate validates ApiTestInputBody
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ApiTestInputBody, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SomeString returns string, set in ApiTestInputBody.
	SomeString() string
	// SetSomeString assigns string provided by user to ApiTestInputBody
	SetSomeString(value string) ApiTestInputBody
	// HasSomeString checks if SomeString has been set in ApiTestInputBody
	HasSomeString() bool
}

// description is TBD
// SomeString returns a string
func (obj *apiTestInputBody) SomeString() string {

	return *obj.obj.SomeString

}

// description is TBD
// SomeString returns a string
func (obj *apiTestInputBody) HasSomeString() bool {
	return obj.obj.SomeString != nil
}

// description is TBD
// SetSomeString sets the string value in the ApiTestInputBody object
func (obj *apiTestInputBody) SetSomeString(value string) ApiTestInputBody {

	obj.obj.SomeString = &value
	return obj
}

func (obj *apiTestInputBody) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *apiTestInputBody) setDefault() {

}

// ***** SetConfigResponse *****
type setConfigResponse struct {
	validation
	obj *openapi.SetConfigResponse
}

func NewSetConfigResponse() SetConfigResponse {
	obj := setConfigResponse{obj: &openapi.SetConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *setConfigResponse) Msg() *openapi.SetConfigResponse {
	return obj.obj
}

func (obj *setConfigResponse) SetMsg(msg *openapi.SetConfigResponse) SetConfigResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *setConfigResponse) ToProto() (*openapi.SetConfigResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *setConfigResponse) FromProto(msg *openapi.SetConfigResponse) (SetConfigResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *setConfigResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *setConfigResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *setConfigResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *setConfigResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *setConfigResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *setConfigResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *setConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *setConfigResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *setConfigResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *setConfigResponse) Clone() (SetConfigResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewSetConfigResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// SetConfigResponse is description is TBD
type SetConfigResponse interface {
	Validation
	// Msg marshals SetConfigResponse to protobuf object *openapi.SetConfigResponse
	// and doesn't set defaults
	Msg() *openapi.SetConfigResponse
	// SetMsg unmarshals SetConfigResponse from protobuf object *openapi.SetConfigResponse
	// and doesn't set defaults
	SetMsg(*openapi.SetConfigResponse) SetConfigResponse
	// ToProto marshals SetConfigResponse to protobuf object *openapi.SetConfigResponse
	ToProto() (*openapi.SetConfigResponse, error)
	// ToPbText marshals SetConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals SetConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals SetConfigResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals SetConfigResponse from protobuf object *openapi.SetConfigResponse
	FromProto(msg *openapi.SetConfigResponse) (SetConfigResponse, error)
	// FromPbText unmarshals SetConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals SetConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals SetConfigResponse from JSON text
	FromJson(value string) error
	// Validate validates SetConfigResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (SetConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseBytes returns []byte, set in SetConfigResponse.
	ResponseBytes() []byte
	// SetResponseBytes assigns []byte provided by user to SetConfigResponse
	SetResponseBytes(value []byte) SetConfigResponse
	// HasResponseBytes checks if ResponseBytes has been set in SetConfigResponse
	HasResponseBytes() bool
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setConfigResponse) ResponseBytes() []byte {

	return obj.obj.ResponseBytes
}

// description is TBD
// ResponseBytes returns a []byte
func (obj *setConfigResponse) HasResponseBytes() bool {
	return obj.obj.ResponseBytes != nil
}

// description is TBD
// SetResponseBytes sets the []byte value in the SetConfigResponse object
func (obj *setConfigResponse) SetResponseBytes(value []byte) SetConfigResponse {

	obj.obj.ResponseBytes = value
	return obj
}

func (obj *setConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *setConfigResponse) setDefault() {

}

// ***** UpdateConfigurationResponse *****
type updateConfigurationResponse struct {
	validation
	obj                *openapi.UpdateConfigurationResponse
	prefixConfigHolder PrefixConfig
}

func NewUpdateConfigurationResponse() UpdateConfigurationResponse {
	obj := updateConfigurationResponse{obj: &openapi.UpdateConfigurationResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *updateConfigurationResponse) Msg() *openapi.UpdateConfigurationResponse {
	return obj.obj
}

func (obj *updateConfigurationResponse) SetMsg(msg *openapi.UpdateConfigurationResponse) UpdateConfigurationResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *updateConfigurationResponse) ToProto() (*openapi.UpdateConfigurationResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *updateConfigurationResponse) FromProto(msg *openapi.UpdateConfigurationResponse) (UpdateConfigurationResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *updateConfigurationResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *updateConfigurationResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *updateConfigurationResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *updateConfigurationResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *updateConfigurationResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *updateConfigurationResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *updateConfigurationResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateConfigurationResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateConfigurationResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateConfigurationResponse) Clone() (UpdateConfigurationResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateConfigurationResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *updateConfigurationResponse) setNil() {
	obj.prefixConfigHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateConfigurationResponse is description is TBD
type UpdateConfigurationResponse interface {
	Validation
	// Msg marshals UpdateConfigurationResponse to protobuf object *openapi.UpdateConfigurationResponse
	// and doesn't set defaults
	Msg() *openapi.UpdateConfigurationResponse
	// SetMsg unmarshals UpdateConfigurationResponse from protobuf object *openapi.UpdateConfigurationResponse
	// and doesn't set defaults
	SetMsg(*openapi.UpdateConfigurationResponse) UpdateConfigurationResponse
	// ToProto marshals UpdateConfigurationResponse to protobuf object *openapi.UpdateConfigurationResponse
	ToProto() (*openapi.UpdateConfigurationResponse, error)
	// ToPbText marshals UpdateConfigurationResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateConfigurationResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateConfigurationResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals UpdateConfigurationResponse from protobuf object *openapi.UpdateConfigurationResponse
	FromProto(msg *openapi.UpdateConfigurationResponse) (UpdateConfigurationResponse, error)
	// FromPbText unmarshals UpdateConfigurationResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateConfigurationResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateConfigurationResponse from JSON text
	FromJson(value string) error
	// Validate validates UpdateConfigurationResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateConfigurationResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrefixConfig returns PrefixConfig, set in UpdateConfigurationResponse.
	// PrefixConfig is container which retains the configuration
	PrefixConfig() PrefixConfig
	// SetPrefixConfig assigns PrefixConfig provided by user to UpdateConfigurationResponse.
	// PrefixConfig is container which retains the configuration
	SetPrefixConfig(value PrefixConfig) UpdateConfigurationResponse
	// HasPrefixConfig checks if PrefixConfig has been set in UpdateConfigurationResponse
	HasPrefixConfig() bool
	setNil()
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *updateConfigurationResponse) PrefixConfig() PrefixConfig {
	if obj.obj.PrefixConfig == nil {
		obj.obj.PrefixConfig = NewPrefixConfig().Msg()
	}
	if obj.prefixConfigHolder == nil {
		obj.prefixConfigHolder = &prefixConfig{obj: obj.obj.PrefixConfig}
	}
	return obj.prefixConfigHolder
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *updateConfigurationResponse) HasPrefixConfig() bool {
	return obj.obj.PrefixConfig != nil
}

// description is TBD
// SetPrefixConfig sets the PrefixConfig value in the UpdateConfigurationResponse object
func (obj *updateConfigurationResponse) SetPrefixConfig(value PrefixConfig) UpdateConfigurationResponse {

	obj.prefixConfigHolder = nil
	obj.obj.PrefixConfig = value.Msg()

	return obj
}

func (obj *updateConfigurationResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixConfig != nil {

		obj.PrefixConfig().validateObj(vObj, set_default)
	}

}

func (obj *updateConfigurationResponse) setDefault() {

}

// ***** GetConfigResponse *****
type getConfigResponse struct {
	validation
	obj                *openapi.GetConfigResponse
	prefixConfigHolder PrefixConfig
}

func NewGetConfigResponse() GetConfigResponse {
	obj := getConfigResponse{obj: &openapi.GetConfigResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getConfigResponse) Msg() *openapi.GetConfigResponse {
	return obj.obj
}

func (obj *getConfigResponse) SetMsg(msg *openapi.GetConfigResponse) GetConfigResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getConfigResponse) ToProto() (*openapi.GetConfigResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getConfigResponse) FromProto(msg *openapi.GetConfigResponse) (GetConfigResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getConfigResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getConfigResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getConfigResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getConfigResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getConfigResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getConfigResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getConfigResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getConfigResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getConfigResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getConfigResponse) Clone() (GetConfigResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetConfigResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getConfigResponse) setNil() {
	obj.prefixConfigHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetConfigResponse is description is TBD
type GetConfigResponse interface {
	Validation
	// Msg marshals GetConfigResponse to protobuf object *openapi.GetConfigResponse
	// and doesn't set defaults
	Msg() *openapi.GetConfigResponse
	// SetMsg unmarshals GetConfigResponse from protobuf object *openapi.GetConfigResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetConfigResponse) GetConfigResponse
	// ToProto marshals GetConfigResponse to protobuf object *openapi.GetConfigResponse
	ToProto() (*openapi.GetConfigResponse, error)
	// ToPbText marshals GetConfigResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetConfigResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetConfigResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetConfigResponse from protobuf object *openapi.GetConfigResponse
	FromProto(msg *openapi.GetConfigResponse) (GetConfigResponse, error)
	// FromPbText unmarshals GetConfigResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetConfigResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetConfigResponse from JSON text
	FromJson(value string) error
	// Validate validates GetConfigResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetConfigResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrefixConfig returns PrefixConfig, set in GetConfigResponse.
	// PrefixConfig is container which retains the configuration
	PrefixConfig() PrefixConfig
	// SetPrefixConfig assigns PrefixConfig provided by user to GetConfigResponse.
	// PrefixConfig is container which retains the configuration
	SetPrefixConfig(value PrefixConfig) GetConfigResponse
	// HasPrefixConfig checks if PrefixConfig has been set in GetConfigResponse
	HasPrefixConfig() bool
	setNil()
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *getConfigResponse) PrefixConfig() PrefixConfig {
	if obj.obj.PrefixConfig == nil {
		obj.obj.PrefixConfig = NewPrefixConfig().Msg()
	}
	if obj.prefixConfigHolder == nil {
		obj.prefixConfigHolder = &prefixConfig{obj: obj.obj.PrefixConfig}
	}
	return obj.prefixConfigHolder
}

// description is TBD
// PrefixConfig returns a PrefixConfig
func (obj *getConfigResponse) HasPrefixConfig() bool {
	return obj.obj.PrefixConfig != nil
}

// description is TBD
// SetPrefixConfig sets the PrefixConfig value in the GetConfigResponse object
func (obj *getConfigResponse) SetPrefixConfig(value PrefixConfig) GetConfigResponse {

	obj.prefixConfigHolder = nil
	obj.obj.PrefixConfig = value.Msg()

	return obj
}

func (obj *getConfigResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixConfig != nil {

		obj.PrefixConfig().validateObj(vObj, set_default)
	}

}

func (obj *getConfigResponse) setDefault() {

}

// ***** GetMetricsResponse *****
type getMetricsResponse struct {
	validation
	obj           *openapi.GetMetricsResponse
	metricsHolder Metrics
}

func NewGetMetricsResponse() GetMetricsResponse {
	obj := getMetricsResponse{obj: &openapi.GetMetricsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getMetricsResponse) Msg() *openapi.GetMetricsResponse {
	return obj.obj
}

func (obj *getMetricsResponse) SetMsg(msg *openapi.GetMetricsResponse) GetMetricsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getMetricsResponse) ToProto() (*openapi.GetMetricsResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getMetricsResponse) FromProto(msg *openapi.GetMetricsResponse) (GetMetricsResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getMetricsResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getMetricsResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getMetricsResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getMetricsResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getMetricsResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getMetricsResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getMetricsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getMetricsResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getMetricsResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getMetricsResponse) Clone() (GetMetricsResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetMetricsResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getMetricsResponse) setNil() {
	obj.metricsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetMetricsResponse is description is TBD
type GetMetricsResponse interface {
	Validation
	// Msg marshals GetMetricsResponse to protobuf object *openapi.GetMetricsResponse
	// and doesn't set defaults
	Msg() *openapi.GetMetricsResponse
	// SetMsg unmarshals GetMetricsResponse from protobuf object *openapi.GetMetricsResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetMetricsResponse) GetMetricsResponse
	// ToProto marshals GetMetricsResponse to protobuf object *openapi.GetMetricsResponse
	ToProto() (*openapi.GetMetricsResponse, error)
	// ToPbText marshals GetMetricsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetMetricsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetMetricsResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetMetricsResponse from protobuf object *openapi.GetMetricsResponse
	FromProto(msg *openapi.GetMetricsResponse) (GetMetricsResponse, error)
	// FromPbText unmarshals GetMetricsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetMetricsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetMetricsResponse from JSON text
	FromJson(value string) error
	// Validate validates GetMetricsResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetMetricsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Metrics returns Metrics, set in GetMetricsResponse.
	// Metrics is description is TBD
	Metrics() Metrics
	// SetMetrics assigns Metrics provided by user to GetMetricsResponse.
	// Metrics is description is TBD
	SetMetrics(value Metrics) GetMetricsResponse
	// HasMetrics checks if Metrics has been set in GetMetricsResponse
	HasMetrics() bool
	setNil()
}

// description is TBD
// Metrics returns a Metrics
func (obj *getMetricsResponse) Metrics() Metrics {
	if obj.obj.Metrics == nil {
		obj.obj.Metrics = NewMetrics().Msg()
	}
	if obj.metricsHolder == nil {
		obj.metricsHolder = &metrics{obj: obj.obj.Metrics}
	}
	return obj.metricsHolder
}

// description is TBD
// Metrics returns a Metrics
func (obj *getMetricsResponse) HasMetrics() bool {
	return obj.obj.Metrics != nil
}

// description is TBD
// SetMetrics sets the Metrics value in the GetMetricsResponse object
func (obj *getMetricsResponse) SetMetrics(value Metrics) GetMetricsResponse {

	obj.metricsHolder = nil
	obj.obj.Metrics = value.Msg()

	return obj
}

func (obj *getMetricsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Metrics != nil {

		obj.Metrics().validateObj(vObj, set_default)
	}

}

func (obj *getMetricsResponse) setDefault() {

}

// ***** GetWarningsResponse *****
type getWarningsResponse struct {
	validation
	obj                  *openapi.GetWarningsResponse
	warningDetailsHolder WarningDetails
}

func NewGetWarningsResponse() GetWarningsResponse {
	obj := getWarningsResponse{obj: &openapi.GetWarningsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getWarningsResponse) Msg() *openapi.GetWarningsResponse {
	return obj.obj
}

func (obj *getWarningsResponse) SetMsg(msg *openapi.GetWarningsResponse) GetWarningsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getWarningsResponse) ToProto() (*openapi.GetWarningsResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getWarningsResponse) FromProto(msg *openapi.GetWarningsResponse) (GetWarningsResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getWarningsResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getWarningsResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getWarningsResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getWarningsResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getWarningsResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getWarningsResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getWarningsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getWarningsResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getWarningsResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getWarningsResponse) Clone() (GetWarningsResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetWarningsResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getWarningsResponse) setNil() {
	obj.warningDetailsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetWarningsResponse is description is TBD
type GetWarningsResponse interface {
	Validation
	// Msg marshals GetWarningsResponse to protobuf object *openapi.GetWarningsResponse
	// and doesn't set defaults
	Msg() *openapi.GetWarningsResponse
	// SetMsg unmarshals GetWarningsResponse from protobuf object *openapi.GetWarningsResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetWarningsResponse) GetWarningsResponse
	// ToProto marshals GetWarningsResponse to protobuf object *openapi.GetWarningsResponse
	ToProto() (*openapi.GetWarningsResponse, error)
	// ToPbText marshals GetWarningsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetWarningsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetWarningsResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetWarningsResponse from protobuf object *openapi.GetWarningsResponse
	FromProto(msg *openapi.GetWarningsResponse) (GetWarningsResponse, error)
	// FromPbText unmarshals GetWarningsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetWarningsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetWarningsResponse from JSON text
	FromJson(value string) error
	// Validate validates GetWarningsResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetWarningsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// WarningDetails returns WarningDetails, set in GetWarningsResponse.
	// WarningDetails is description is TBD
	WarningDetails() WarningDetails
	// SetWarningDetails assigns WarningDetails provided by user to GetWarningsResponse.
	// WarningDetails is description is TBD
	SetWarningDetails(value WarningDetails) GetWarningsResponse
	// HasWarningDetails checks if WarningDetails has been set in GetWarningsResponse
	HasWarningDetails() bool
	setNil()
}

// description is TBD
// WarningDetails returns a WarningDetails
func (obj *getWarningsResponse) WarningDetails() WarningDetails {
	if obj.obj.WarningDetails == nil {
		obj.obj.WarningDetails = NewWarningDetails().Msg()
	}
	if obj.warningDetailsHolder == nil {
		obj.warningDetailsHolder = &warningDetails{obj: obj.obj.WarningDetails}
	}
	return obj.warningDetailsHolder
}

// description is TBD
// WarningDetails returns a WarningDetails
func (obj *getWarningsResponse) HasWarningDetails() bool {
	return obj.obj.WarningDetails != nil
}

// description is TBD
// SetWarningDetails sets the WarningDetails value in the GetWarningsResponse object
func (obj *getWarningsResponse) SetWarningDetails(value WarningDetails) GetWarningsResponse {

	obj.warningDetailsHolder = nil
	obj.obj.WarningDetails = value.Msg()

	return obj
}

func (obj *getWarningsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.WarningDetails != nil {

		obj.WarningDetails().validateObj(vObj, set_default)
	}

}

func (obj *getWarningsResponse) setDefault() {

}

// ***** ClearWarningsResponse *****
type clearWarningsResponse struct {
	validation
	obj *openapi.ClearWarningsResponse
}

func NewClearWarningsResponse() ClearWarningsResponse {
	obj := clearWarningsResponse{obj: &openapi.ClearWarningsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *clearWarningsResponse) Msg() *openapi.ClearWarningsResponse {
	return obj.obj
}

func (obj *clearWarningsResponse) SetMsg(msg *openapi.ClearWarningsResponse) ClearWarningsResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *clearWarningsResponse) ToProto() (*openapi.ClearWarningsResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *clearWarningsResponse) FromProto(msg *openapi.ClearWarningsResponse) (ClearWarningsResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *clearWarningsResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *clearWarningsResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *clearWarningsResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *clearWarningsResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *clearWarningsResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *clearWarningsResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *clearWarningsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *clearWarningsResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *clearWarningsResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *clearWarningsResponse) Clone() (ClearWarningsResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewClearWarningsResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// ClearWarningsResponse is description is TBD
type ClearWarningsResponse interface {
	Validation
	// Msg marshals ClearWarningsResponse to protobuf object *openapi.ClearWarningsResponse
	// and doesn't set defaults
	Msg() *openapi.ClearWarningsResponse
	// SetMsg unmarshals ClearWarningsResponse from protobuf object *openapi.ClearWarningsResponse
	// and doesn't set defaults
	SetMsg(*openapi.ClearWarningsResponse) ClearWarningsResponse
	// ToProto marshals ClearWarningsResponse to protobuf object *openapi.ClearWarningsResponse
	ToProto() (*openapi.ClearWarningsResponse, error)
	// ToPbText marshals ClearWarningsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ClearWarningsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals ClearWarningsResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ClearWarningsResponse from protobuf object *openapi.ClearWarningsResponse
	FromProto(msg *openapi.ClearWarningsResponse) (ClearWarningsResponse, error)
	// FromPbText unmarshals ClearWarningsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ClearWarningsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ClearWarningsResponse from JSON text
	FromJson(value string) error
	// Validate validates ClearWarningsResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ClearWarningsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseString returns string, set in ClearWarningsResponse.
	ResponseString() string
	// SetResponseString assigns string provided by user to ClearWarningsResponse
	SetResponseString(value string) ClearWarningsResponse
	// HasResponseString checks if ResponseString has been set in ClearWarningsResponse
	HasResponseString() bool
}

// description is TBD
// ResponseString returns a string
func (obj *clearWarningsResponse) ResponseString() string {
	return obj.obj.String_
}

// description is TBD
// ResponseString returns a string
func (obj *clearWarningsResponse) HasResponseString() bool {
	return obj.obj.String_ != ""
}

// description is TBD
// SetResponseString sets the string value in the ClearWarningsResponse object
func (obj *clearWarningsResponse) SetResponseString(value string) ClearWarningsResponse {
	obj.obj.String_ = value
	return obj
}

func (obj *clearWarningsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *clearWarningsResponse) setDefault() {

}

// ***** GetRootResponseResponse *****
type getRootResponseResponse struct {
	validation
	obj                         *openapi.GetRootResponseResponse
	commonResponseSuccessHolder CommonResponseSuccess
}

func NewGetRootResponseResponse() GetRootResponseResponse {
	obj := getRootResponseResponse{obj: &openapi.GetRootResponseResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getRootResponseResponse) Msg() *openapi.GetRootResponseResponse {
	return obj.obj
}

func (obj *getRootResponseResponse) SetMsg(msg *openapi.GetRootResponseResponse) GetRootResponseResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getRootResponseResponse) ToProto() (*openapi.GetRootResponseResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getRootResponseResponse) FromProto(msg *openapi.GetRootResponseResponse) (GetRootResponseResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getRootResponseResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getRootResponseResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getRootResponseResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getRootResponseResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getRootResponseResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getRootResponseResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getRootResponseResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getRootResponseResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getRootResponseResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getRootResponseResponse) Clone() (GetRootResponseResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetRootResponseResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getRootResponseResponse) setNil() {
	obj.commonResponseSuccessHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetRootResponseResponse is description is TBD
type GetRootResponseResponse interface {
	Validation
	// Msg marshals GetRootResponseResponse to protobuf object *openapi.GetRootResponseResponse
	// and doesn't set defaults
	Msg() *openapi.GetRootResponseResponse
	// SetMsg unmarshals GetRootResponseResponse from protobuf object *openapi.GetRootResponseResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetRootResponseResponse) GetRootResponseResponse
	// ToProto marshals GetRootResponseResponse to protobuf object *openapi.GetRootResponseResponse
	ToProto() (*openapi.GetRootResponseResponse, error)
	// ToPbText marshals GetRootResponseResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetRootResponseResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetRootResponseResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetRootResponseResponse from protobuf object *openapi.GetRootResponseResponse
	FromProto(msg *openapi.GetRootResponseResponse) (GetRootResponseResponse, error)
	// FromPbText unmarshals GetRootResponseResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetRootResponseResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetRootResponseResponse from JSON text
	FromJson(value string) error
	// Validate validates GetRootResponseResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetRootResponseResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonResponseSuccess returns CommonResponseSuccess, set in GetRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	CommonResponseSuccess() CommonResponseSuccess
	// SetCommonResponseSuccess assigns CommonResponseSuccess provided by user to GetRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	SetCommonResponseSuccess(value CommonResponseSuccess) GetRootResponseResponse
	// HasCommonResponseSuccess checks if CommonResponseSuccess has been set in GetRootResponseResponse
	HasCommonResponseSuccess() bool
	setNil()
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *getRootResponseResponse) CommonResponseSuccess() CommonResponseSuccess {
	if obj.obj.CommonResponseSuccess == nil {
		obj.obj.CommonResponseSuccess = NewCommonResponseSuccess().Msg()
	}
	if obj.commonResponseSuccessHolder == nil {
		obj.commonResponseSuccessHolder = &commonResponseSuccess{obj: obj.obj.CommonResponseSuccess}
	}
	return obj.commonResponseSuccessHolder
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *getRootResponseResponse) HasCommonResponseSuccess() bool {
	return obj.obj.CommonResponseSuccess != nil
}

// description is TBD
// SetCommonResponseSuccess sets the CommonResponseSuccess value in the GetRootResponseResponse object
func (obj *getRootResponseResponse) SetCommonResponseSuccess(value CommonResponseSuccess) GetRootResponseResponse {

	obj.commonResponseSuccessHolder = nil
	obj.obj.CommonResponseSuccess = value.Msg()

	return obj
}

func (obj *getRootResponseResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonResponseSuccess != nil {

		obj.CommonResponseSuccess().validateObj(vObj, set_default)
	}

}

func (obj *getRootResponseResponse) setDefault() {

}

// ***** DummyResponseTestResponse *****
type dummyResponseTestResponse struct {
	validation
	obj *openapi.DummyResponseTestResponse
}

func NewDummyResponseTestResponse() DummyResponseTestResponse {
	obj := dummyResponseTestResponse{obj: &openapi.DummyResponseTestResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *dummyResponseTestResponse) Msg() *openapi.DummyResponseTestResponse {
	return obj.obj
}

func (obj *dummyResponseTestResponse) SetMsg(msg *openapi.DummyResponseTestResponse) DummyResponseTestResponse {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *dummyResponseTestResponse) ToProto() (*openapi.DummyResponseTestResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *dummyResponseTestResponse) FromProto(msg *openapi.DummyResponseTestResponse) (DummyResponseTestResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *dummyResponseTestResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *dummyResponseTestResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *dummyResponseTestResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *dummyResponseTestResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *dummyResponseTestResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *dummyResponseTestResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *dummyResponseTestResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dummyResponseTestResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dummyResponseTestResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dummyResponseTestResponse) Clone() (DummyResponseTestResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDummyResponseTestResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// DummyResponseTestResponse is description is TBD
type DummyResponseTestResponse interface {
	Validation
	// Msg marshals DummyResponseTestResponse to protobuf object *openapi.DummyResponseTestResponse
	// and doesn't set defaults
	Msg() *openapi.DummyResponseTestResponse
	// SetMsg unmarshals DummyResponseTestResponse from protobuf object *openapi.DummyResponseTestResponse
	// and doesn't set defaults
	SetMsg(*openapi.DummyResponseTestResponse) DummyResponseTestResponse
	// ToProto marshals DummyResponseTestResponse to protobuf object *openapi.DummyResponseTestResponse
	ToProto() (*openapi.DummyResponseTestResponse, error)
	// ToPbText marshals DummyResponseTestResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DummyResponseTestResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals DummyResponseTestResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals DummyResponseTestResponse from protobuf object *openapi.DummyResponseTestResponse
	FromProto(msg *openapi.DummyResponseTestResponse) (DummyResponseTestResponse, error)
	// FromPbText unmarshals DummyResponseTestResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DummyResponseTestResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DummyResponseTestResponse from JSON text
	FromJson(value string) error
	// Validate validates DummyResponseTestResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DummyResponseTestResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ResponseString returns string, set in DummyResponseTestResponse.
	ResponseString() string
	// SetResponseString assigns string provided by user to DummyResponseTestResponse
	SetResponseString(value string) DummyResponseTestResponse
	// HasResponseString checks if ResponseString has been set in DummyResponseTestResponse
	HasResponseString() bool
}

// description is TBD
// ResponseString returns a string
func (obj *dummyResponseTestResponse) ResponseString() string {
	return obj.obj.String_
}

// description is TBD
// ResponseString returns a string
func (obj *dummyResponseTestResponse) HasResponseString() bool {
	return obj.obj.String_ != ""
}

// description is TBD
// SetResponseString sets the string value in the DummyResponseTestResponse object
func (obj *dummyResponseTestResponse) SetResponseString(value string) DummyResponseTestResponse {
	obj.obj.String_ = value
	return obj
}

func (obj *dummyResponseTestResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *dummyResponseTestResponse) setDefault() {

}

// ***** PostRootResponseResponse *****
type postRootResponseResponse struct {
	validation
	obj                         *openapi.PostRootResponseResponse
	commonResponseSuccessHolder CommonResponseSuccess
}

func NewPostRootResponseResponse() PostRootResponseResponse {
	obj := postRootResponseResponse{obj: &openapi.PostRootResponseResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *postRootResponseResponse) Msg() *openapi.PostRootResponseResponse {
	return obj.obj
}

func (obj *postRootResponseResponse) SetMsg(msg *openapi.PostRootResponseResponse) PostRootResponseResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *postRootResponseResponse) ToProto() (*openapi.PostRootResponseResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *postRootResponseResponse) FromProto(msg *openapi.PostRootResponseResponse) (PostRootResponseResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *postRootResponseResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *postRootResponseResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *postRootResponseResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *postRootResponseResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *postRootResponseResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *postRootResponseResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *postRootResponseResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *postRootResponseResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *postRootResponseResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *postRootResponseResponse) Clone() (PostRootResponseResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPostRootResponseResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *postRootResponseResponse) setNil() {
	obj.commonResponseSuccessHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PostRootResponseResponse is description is TBD
type PostRootResponseResponse interface {
	Validation
	// Msg marshals PostRootResponseResponse to protobuf object *openapi.PostRootResponseResponse
	// and doesn't set defaults
	Msg() *openapi.PostRootResponseResponse
	// SetMsg unmarshals PostRootResponseResponse from protobuf object *openapi.PostRootResponseResponse
	// and doesn't set defaults
	SetMsg(*openapi.PostRootResponseResponse) PostRootResponseResponse
	// ToProto marshals PostRootResponseResponse to protobuf object *openapi.PostRootResponseResponse
	ToProto() (*openapi.PostRootResponseResponse, error)
	// ToPbText marshals PostRootResponseResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PostRootResponseResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals PostRootResponseResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PostRootResponseResponse from protobuf object *openapi.PostRootResponseResponse
	FromProto(msg *openapi.PostRootResponseResponse) (PostRootResponseResponse, error)
	// FromPbText unmarshals PostRootResponseResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PostRootResponseResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PostRootResponseResponse from JSON text
	FromJson(value string) error
	// Validate validates PostRootResponseResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PostRootResponseResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// CommonResponseSuccess returns CommonResponseSuccess, set in PostRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	CommonResponseSuccess() CommonResponseSuccess
	// SetCommonResponseSuccess assigns CommonResponseSuccess provided by user to PostRootResponseResponse.
	// CommonResponseSuccess is description is TBD
	SetCommonResponseSuccess(value CommonResponseSuccess) PostRootResponseResponse
	// HasCommonResponseSuccess checks if CommonResponseSuccess has been set in PostRootResponseResponse
	HasCommonResponseSuccess() bool
	setNil()
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *postRootResponseResponse) CommonResponseSuccess() CommonResponseSuccess {
	if obj.obj.CommonResponseSuccess == nil {
		obj.obj.CommonResponseSuccess = NewCommonResponseSuccess().Msg()
	}
	if obj.commonResponseSuccessHolder == nil {
		obj.commonResponseSuccessHolder = &commonResponseSuccess{obj: obj.obj.CommonResponseSuccess}
	}
	return obj.commonResponseSuccessHolder
}

// description is TBD
// CommonResponseSuccess returns a CommonResponseSuccess
func (obj *postRootResponseResponse) HasCommonResponseSuccess() bool {
	return obj.obj.CommonResponseSuccess != nil
}

// description is TBD
// SetCommonResponseSuccess sets the CommonResponseSuccess value in the PostRootResponseResponse object
func (obj *postRootResponseResponse) SetCommonResponseSuccess(value CommonResponseSuccess) PostRootResponseResponse {

	obj.commonResponseSuccessHolder = nil
	obj.obj.CommonResponseSuccess = value.Msg()

	return obj
}

func (obj *postRootResponseResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CommonResponseSuccess != nil {

		obj.CommonResponseSuccess().validateObj(vObj, set_default)
	}

}

func (obj *postRootResponseResponse) setDefault() {

}

// ***** GetAllItemsResponse *****
type getAllItemsResponse struct {
	validation
	obj                      *openapi.GetAllItemsResponse
	serviceAbcItemListHolder ServiceAbcItemList
}

func NewGetAllItemsResponse() GetAllItemsResponse {
	obj := getAllItemsResponse{obj: &openapi.GetAllItemsResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getAllItemsResponse) Msg() *openapi.GetAllItemsResponse {
	return obj.obj
}

func (obj *getAllItemsResponse) SetMsg(msg *openapi.GetAllItemsResponse) GetAllItemsResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getAllItemsResponse) ToProto() (*openapi.GetAllItemsResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getAllItemsResponse) FromProto(msg *openapi.GetAllItemsResponse) (GetAllItemsResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getAllItemsResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getAllItemsResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getAllItemsResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getAllItemsResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getAllItemsResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getAllItemsResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getAllItemsResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getAllItemsResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getAllItemsResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getAllItemsResponse) Clone() (GetAllItemsResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetAllItemsResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getAllItemsResponse) setNil() {
	obj.serviceAbcItemListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetAllItemsResponse is description is TBD
type GetAllItemsResponse interface {
	Validation
	// Msg marshals GetAllItemsResponse to protobuf object *openapi.GetAllItemsResponse
	// and doesn't set defaults
	Msg() *openapi.GetAllItemsResponse
	// SetMsg unmarshals GetAllItemsResponse from protobuf object *openapi.GetAllItemsResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetAllItemsResponse) GetAllItemsResponse
	// ToProto marshals GetAllItemsResponse to protobuf object *openapi.GetAllItemsResponse
	ToProto() (*openapi.GetAllItemsResponse, error)
	// ToPbText marshals GetAllItemsResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetAllItemsResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetAllItemsResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetAllItemsResponse from protobuf object *openapi.GetAllItemsResponse
	FromProto(msg *openapi.GetAllItemsResponse) (GetAllItemsResponse, error)
	// FromPbText unmarshals GetAllItemsResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetAllItemsResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetAllItemsResponse from JSON text
	FromJson(value string) error
	// Validate validates GetAllItemsResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetAllItemsResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItemList returns ServiceAbcItemList, set in GetAllItemsResponse.
	// ServiceAbcItemList is description is TBD
	ServiceAbcItemList() ServiceAbcItemList
	// SetServiceAbcItemList assigns ServiceAbcItemList provided by user to GetAllItemsResponse.
	// ServiceAbcItemList is description is TBD
	SetServiceAbcItemList(value ServiceAbcItemList) GetAllItemsResponse
	// HasServiceAbcItemList checks if ServiceAbcItemList has been set in GetAllItemsResponse
	HasServiceAbcItemList() bool
	setNil()
}

// description is TBD
// ServiceAbcItemList returns a ServiceAbcItemList
func (obj *getAllItemsResponse) ServiceAbcItemList() ServiceAbcItemList {
	if obj.obj.ServiceAbcItemList == nil {
		obj.obj.ServiceAbcItemList = NewServiceAbcItemList().Msg()
	}
	if obj.serviceAbcItemListHolder == nil {
		obj.serviceAbcItemListHolder = &serviceAbcItemList{obj: obj.obj.ServiceAbcItemList}
	}
	return obj.serviceAbcItemListHolder
}

// description is TBD
// ServiceAbcItemList returns a ServiceAbcItemList
func (obj *getAllItemsResponse) HasServiceAbcItemList() bool {
	return obj.obj.ServiceAbcItemList != nil
}

// description is TBD
// SetServiceAbcItemList sets the ServiceAbcItemList value in the GetAllItemsResponse object
func (obj *getAllItemsResponse) SetServiceAbcItemList(value ServiceAbcItemList) GetAllItemsResponse {

	obj.serviceAbcItemListHolder = nil
	obj.obj.ServiceAbcItemList = value.Msg()

	return obj
}

func (obj *getAllItemsResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItemList != nil {

		obj.ServiceAbcItemList().validateObj(vObj, set_default)
	}

}

func (obj *getAllItemsResponse) setDefault() {

}

// ***** GetSingleItemResponse *****
type getSingleItemResponse struct {
	validation
	obj                  *openapi.GetSingleItemResponse
	serviceAbcItemHolder ServiceAbcItem
}

func NewGetSingleItemResponse() GetSingleItemResponse {
	obj := getSingleItemResponse{obj: &openapi.GetSingleItemResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getSingleItemResponse) Msg() *openapi.GetSingleItemResponse {
	return obj.obj
}

func (obj *getSingleItemResponse) SetMsg(msg *openapi.GetSingleItemResponse) GetSingleItemResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getSingleItemResponse) ToProto() (*openapi.GetSingleItemResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getSingleItemResponse) FromProto(msg *openapi.GetSingleItemResponse) (GetSingleItemResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getSingleItemResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getSingleItemResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getSingleItemResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getSingleItemResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getSingleItemResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getSingleItemResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getSingleItemResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getSingleItemResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getSingleItemResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getSingleItemResponse) Clone() (GetSingleItemResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetSingleItemResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getSingleItemResponse) setNil() {
	obj.serviceAbcItemHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetSingleItemResponse is description is TBD
type GetSingleItemResponse interface {
	Validation
	// Msg marshals GetSingleItemResponse to protobuf object *openapi.GetSingleItemResponse
	// and doesn't set defaults
	Msg() *openapi.GetSingleItemResponse
	// SetMsg unmarshals GetSingleItemResponse from protobuf object *openapi.GetSingleItemResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetSingleItemResponse) GetSingleItemResponse
	// ToProto marshals GetSingleItemResponse to protobuf object *openapi.GetSingleItemResponse
	ToProto() (*openapi.GetSingleItemResponse, error)
	// ToPbText marshals GetSingleItemResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetSingleItemResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetSingleItemResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetSingleItemResponse from protobuf object *openapi.GetSingleItemResponse
	FromProto(msg *openapi.GetSingleItemResponse) (GetSingleItemResponse, error)
	// FromPbText unmarshals GetSingleItemResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetSingleItemResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetSingleItemResponse from JSON text
	FromJson(value string) error
	// Validate validates GetSingleItemResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetSingleItemResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItem returns ServiceAbcItem, set in GetSingleItemResponse.
	// ServiceAbcItem is description is TBD
	ServiceAbcItem() ServiceAbcItem
	// SetServiceAbcItem assigns ServiceAbcItem provided by user to GetSingleItemResponse.
	// ServiceAbcItem is description is TBD
	SetServiceAbcItem(value ServiceAbcItem) GetSingleItemResponse
	// HasServiceAbcItem checks if ServiceAbcItem has been set in GetSingleItemResponse
	HasServiceAbcItem() bool
	setNil()
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemResponse) ServiceAbcItem() ServiceAbcItem {
	if obj.obj.ServiceAbcItem == nil {
		obj.obj.ServiceAbcItem = NewServiceAbcItem().Msg()
	}
	if obj.serviceAbcItemHolder == nil {
		obj.serviceAbcItemHolder = &serviceAbcItem{obj: obj.obj.ServiceAbcItem}
	}
	return obj.serviceAbcItemHolder
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemResponse) HasServiceAbcItem() bool {
	return obj.obj.ServiceAbcItem != nil
}

// description is TBD
// SetServiceAbcItem sets the ServiceAbcItem value in the GetSingleItemResponse object
func (obj *getSingleItemResponse) SetServiceAbcItem(value ServiceAbcItem) GetSingleItemResponse {

	obj.serviceAbcItemHolder = nil
	obj.obj.ServiceAbcItem = value.Msg()

	return obj
}

func (obj *getSingleItemResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItem != nil {

		obj.ServiceAbcItem().validateObj(vObj, set_default)
	}

}

func (obj *getSingleItemResponse) setDefault() {

}

// ***** GetSingleItemLevel2Response *****
type getSingleItemLevel2Response struct {
	validation
	obj                  *openapi.GetSingleItemLevel2Response
	serviceAbcItemHolder ServiceAbcItem
}

func NewGetSingleItemLevel2Response() GetSingleItemLevel2Response {
	obj := getSingleItemLevel2Response{obj: &openapi.GetSingleItemLevel2Response{}}
	obj.setDefault()
	return &obj
}

func (obj *getSingleItemLevel2Response) Msg() *openapi.GetSingleItemLevel2Response {
	return obj.obj
}

func (obj *getSingleItemLevel2Response) SetMsg(msg *openapi.GetSingleItemLevel2Response) GetSingleItemLevel2Response {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getSingleItemLevel2Response) ToProto() (*openapi.GetSingleItemLevel2Response, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getSingleItemLevel2Response) FromProto(msg *openapi.GetSingleItemLevel2Response) (GetSingleItemLevel2Response, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getSingleItemLevel2Response) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getSingleItemLevel2Response) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getSingleItemLevel2Response) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getSingleItemLevel2Response) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getSingleItemLevel2Response) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getSingleItemLevel2Response) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getSingleItemLevel2Response) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getSingleItemLevel2Response) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getSingleItemLevel2Response) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getSingleItemLevel2Response) Clone() (GetSingleItemLevel2Response, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetSingleItemLevel2Response()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getSingleItemLevel2Response) setNil() {
	obj.serviceAbcItemHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetSingleItemLevel2Response is description is TBD
type GetSingleItemLevel2Response interface {
	Validation
	// Msg marshals GetSingleItemLevel2Response to protobuf object *openapi.GetSingleItemLevel2Response
	// and doesn't set defaults
	Msg() *openapi.GetSingleItemLevel2Response
	// SetMsg unmarshals GetSingleItemLevel2Response from protobuf object *openapi.GetSingleItemLevel2Response
	// and doesn't set defaults
	SetMsg(*openapi.GetSingleItemLevel2Response) GetSingleItemLevel2Response
	// ToProto marshals GetSingleItemLevel2Response to protobuf object *openapi.GetSingleItemLevel2Response
	ToProto() (*openapi.GetSingleItemLevel2Response, error)
	// ToPbText marshals GetSingleItemLevel2Response to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetSingleItemLevel2Response to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetSingleItemLevel2Response to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetSingleItemLevel2Response from protobuf object *openapi.GetSingleItemLevel2Response
	FromProto(msg *openapi.GetSingleItemLevel2Response) (GetSingleItemLevel2Response, error)
	// FromPbText unmarshals GetSingleItemLevel2Response from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetSingleItemLevel2Response from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetSingleItemLevel2Response from JSON text
	FromJson(value string) error
	// Validate validates GetSingleItemLevel2Response
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetSingleItemLevel2Response, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ServiceAbcItem returns ServiceAbcItem, set in GetSingleItemLevel2Response.
	// ServiceAbcItem is description is TBD
	ServiceAbcItem() ServiceAbcItem
	// SetServiceAbcItem assigns ServiceAbcItem provided by user to GetSingleItemLevel2Response.
	// ServiceAbcItem is description is TBD
	SetServiceAbcItem(value ServiceAbcItem) GetSingleItemLevel2Response
	// HasServiceAbcItem checks if ServiceAbcItem has been set in GetSingleItemLevel2Response
	HasServiceAbcItem() bool
	setNil()
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemLevel2Response) ServiceAbcItem() ServiceAbcItem {
	if obj.obj.ServiceAbcItem == nil {
		obj.obj.ServiceAbcItem = NewServiceAbcItem().Msg()
	}
	if obj.serviceAbcItemHolder == nil {
		obj.serviceAbcItemHolder = &serviceAbcItem{obj: obj.obj.ServiceAbcItem}
	}
	return obj.serviceAbcItemHolder
}

// description is TBD
// ServiceAbcItem returns a ServiceAbcItem
func (obj *getSingleItemLevel2Response) HasServiceAbcItem() bool {
	return obj.obj.ServiceAbcItem != nil
}

// description is TBD
// SetServiceAbcItem sets the ServiceAbcItem value in the GetSingleItemLevel2Response object
func (obj *getSingleItemLevel2Response) SetServiceAbcItem(value ServiceAbcItem) GetSingleItemLevel2Response {

	obj.serviceAbcItemHolder = nil
	obj.obj.ServiceAbcItem = value.Msg()

	return obj
}

func (obj *getSingleItemLevel2Response) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ServiceAbcItem != nil {

		obj.ServiceAbcItem().validateObj(vObj, set_default)
	}

}

func (obj *getSingleItemLevel2Response) setDefault() {

}

// ***** GetVersionResponse *****
type getVersionResponse struct {
	validation
	obj           *openapi.GetVersionResponse
	versionHolder Version
}

func NewGetVersionResponse() GetVersionResponse {
	obj := getVersionResponse{obj: &openapi.GetVersionResponse{}}
	obj.setDefault()
	return &obj
}

func (obj *getVersionResponse) Msg() *openapi.GetVersionResponse {
	return obj.obj
}

func (obj *getVersionResponse) SetMsg(msg *openapi.GetVersionResponse) GetVersionResponse {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *getVersionResponse) ToProto() (*openapi.GetVersionResponse, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *getVersionResponse) FromProto(msg *openapi.GetVersionResponse) (GetVersionResponse, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *getVersionResponse) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *getVersionResponse) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *getVersionResponse) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getVersionResponse) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *getVersionResponse) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *getVersionResponse) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *getVersionResponse) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *getVersionResponse) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *getVersionResponse) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *getVersionResponse) Clone() (GetVersionResponse, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGetVersionResponse()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *getVersionResponse) setNil() {
	obj.versionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// GetVersionResponse is description is TBD
type GetVersionResponse interface {
	Validation
	// Msg marshals GetVersionResponse to protobuf object *openapi.GetVersionResponse
	// and doesn't set defaults
	Msg() *openapi.GetVersionResponse
	// SetMsg unmarshals GetVersionResponse from protobuf object *openapi.GetVersionResponse
	// and doesn't set defaults
	SetMsg(*openapi.GetVersionResponse) GetVersionResponse
	// ToProto marshals GetVersionResponse to protobuf object *openapi.GetVersionResponse
	ToProto() (*openapi.GetVersionResponse, error)
	// ToPbText marshals GetVersionResponse to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GetVersionResponse to YAML text
	ToYaml() (string, error)
	// ToJson marshals GetVersionResponse to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GetVersionResponse from protobuf object *openapi.GetVersionResponse
	FromProto(msg *openapi.GetVersionResponse) (GetVersionResponse, error)
	// FromPbText unmarshals GetVersionResponse from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GetVersionResponse from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GetVersionResponse from JSON text
	FromJson(value string) error
	// Validate validates GetVersionResponse
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (GetVersionResponse, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Version returns Version, set in GetVersionResponse.
	// Version is version details
	Version() Version
	// SetVersion assigns Version provided by user to GetVersionResponse.
	// Version is version details
	SetVersion(value Version) GetVersionResponse
	// HasVersion checks if Version has been set in GetVersionResponse
	HasVersion() bool
	setNil()
}

// description is TBD
// Version returns a Version
func (obj *getVersionResponse) Version() Version {
	if obj.obj.Version == nil {
		obj.obj.Version = NewVersion().Msg()
	}
	if obj.versionHolder == nil {
		obj.versionHolder = &version{obj: obj.obj.Version}
	}
	return obj.versionHolder
}

// description is TBD
// Version returns a Version
func (obj *getVersionResponse) HasVersion() bool {
	return obj.obj.Version != nil
}

// description is TBD
// SetVersion sets the Version value in the GetVersionResponse object
func (obj *getVersionResponse) SetVersion(value Version) GetVersionResponse {

	obj.versionHolder = nil
	obj.obj.Version = value.Msg()

	return obj
}

func (obj *getVersionResponse) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Version != nil {

		obj.Version().validateObj(vObj, set_default)
	}

}

func (obj *getVersionResponse) setDefault() {

}

// ***** EObject *****
type eObject struct {
	validation
	obj *openapi.EObject
}

func NewEObject() EObject {
	obj := eObject{obj: &openapi.EObject{}}
	obj.setDefault()
	return &obj
}

func (obj *eObject) Msg() *openapi.EObject {
	return obj.obj
}

func (obj *eObject) SetMsg(msg *openapi.EObject) EObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *eObject) ToProto() (*openapi.EObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *eObject) FromProto(msg *openapi.EObject) (EObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *eObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *eObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *eObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *eObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *eObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *eObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *eObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *eObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *eObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *eObject) Clone() (EObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewEObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// EObject is description is TBD
type EObject interface {
	Validation
	// Msg marshals EObject to protobuf object *openapi.EObject
	// and doesn't set defaults
	Msg() *openapi.EObject
	// SetMsg unmarshals EObject from protobuf object *openapi.EObject
	// and doesn't set defaults
	SetMsg(*openapi.EObject) EObject
	// ToProto marshals EObject to protobuf object *openapi.EObject
	ToProto() (*openapi.EObject, error)
	// ToPbText marshals EObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals EObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals EObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals EObject from protobuf object *openapi.EObject
	FromProto(msg *openapi.EObject) (EObject, error)
	// FromPbText unmarshals EObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals EObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals EObject from JSON text
	FromJson(value string) error
	// Validate validates EObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (EObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EA returns float32, set in EObject.
	EA() float32
	// SetEA assigns float32 provided by user to EObject
	SetEA(value float32) EObject
	// EB returns float64, set in EObject.
	EB() float64
	// SetEB assigns float64 provided by user to EObject
	SetEB(value float64) EObject
	// Name returns string, set in EObject.
	Name() string
	// SetName assigns string provided by user to EObject
	SetName(value string) EObject
	// HasName checks if Name has been set in EObject
	HasName() bool
	// MParam1 returns string, set in EObject.
	MParam1() string
	// SetMParam1 assigns string provided by user to EObject
	SetMParam1(value string) EObject
	// HasMParam1 checks if MParam1 has been set in EObject
	HasMParam1() bool
	// MParam2 returns string, set in EObject.
	MParam2() string
	// SetMParam2 assigns string provided by user to EObject
	SetMParam2(value string) EObject
	// HasMParam2 checks if MParam2 has been set in EObject
	HasMParam2() bool
}

// description is TBD
// EA returns a float32
func (obj *eObject) EA() float32 {

	return obj.obj.EA
}

// description is TBD
// SetEA sets the float32 value in the EObject object
func (obj *eObject) SetEA(value float32) EObject {

	obj.obj.EA = value
	return obj
}

// description is TBD
// EB returns a float64
func (obj *eObject) EB() float64 {

	return obj.obj.EB
}

// description is TBD
// SetEB sets the float64 value in the EObject object
func (obj *eObject) SetEB(value float64) EObject {

	obj.obj.EB = value
	return obj
}

// description is TBD
// Name returns a string
func (obj *eObject) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *eObject) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the EObject object
func (obj *eObject) SetName(value string) EObject {

	obj.obj.Name = &value
	return obj
}

// description is TBD
// MParam1 returns a string
func (obj *eObject) MParam1() string {

	return *obj.obj.MParam1

}

// description is TBD
// MParam1 returns a string
func (obj *eObject) HasMParam1() bool {
	return obj.obj.MParam1 != nil
}

// description is TBD
// SetMParam1 sets the string value in the EObject object
func (obj *eObject) SetMParam1(value string) EObject {

	obj.obj.MParam1 = &value
	return obj
}

// description is TBD
// MParam2 returns a string
func (obj *eObject) MParam2() string {

	return *obj.obj.MParam2

}

// description is TBD
// MParam2 returns a string
func (obj *eObject) HasMParam2() bool {
	return obj.obj.MParam2 != nil
}

// description is TBD
// SetMParam2 sets the string value in the EObject object
func (obj *eObject) SetMParam2(value string) EObject {

	obj.obj.MParam2 = &value
	return obj
}

func (obj *eObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *eObject) setDefault() {

}

// ***** FObject *****
type fObject struct {
	validation
	obj *openapi.FObject
}

func NewFObject() FObject {
	obj := fObject{obj: &openapi.FObject{}}
	obj.setDefault()
	return &obj
}

func (obj *fObject) Msg() *openapi.FObject {
	return obj.obj
}

func (obj *fObject) SetMsg(msg *openapi.FObject) FObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *fObject) ToProto() (*openapi.FObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *fObject) FromProto(msg *openapi.FObject) (FObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *fObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *fObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *fObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *fObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *fObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *fObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *fObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *fObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *fObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *fObject) Clone() (FObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// FObject is description is TBD
type FObject interface {
	Validation
	// Msg marshals FObject to protobuf object *openapi.FObject
	// and doesn't set defaults
	Msg() *openapi.FObject
	// SetMsg unmarshals FObject from protobuf object *openapi.FObject
	// and doesn't set defaults
	SetMsg(*openapi.FObject) FObject
	// ToProto marshals FObject to protobuf object *openapi.FObject
	ToProto() (*openapi.FObject, error)
	// ToPbText marshals FObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals FObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals FObject from protobuf object *openapi.FObject
	FromProto(msg *openapi.FObject) (FObject, error)
	// FromPbText unmarshals FObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FObject from JSON text
	FromJson(value string) error
	// Validate validates FObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FObjectChoiceEnum, set in FObject
	Choice() FObjectChoiceEnum
	// SetChoice assigns FObjectChoiceEnum provided by user to FObject
	SetChoice(value FObjectChoiceEnum) FObject
	// HasChoice checks if Choice has been set in FObject
	HasChoice() bool
	// FA returns string, set in FObject.
	FA() string
	// SetFA assigns string provided by user to FObject
	SetFA(value string) FObject
	// HasFA checks if FA has been set in FObject
	HasFA() bool
	// FB returns float64, set in FObject.
	FB() float64
	// SetFB assigns float64 provided by user to FObject
	SetFB(value float64) FObject
	// HasFB checks if FB has been set in FObject
	HasFB() bool
}

type FObjectChoiceEnum string

// Enum of Choice on FObject
var FObjectChoice = struct {
	F_A FObjectChoiceEnum
	F_B FObjectChoiceEnum
	F_C FObjectChoiceEnum
}{
	F_A: FObjectChoiceEnum("f_a"),
	F_B: FObjectChoiceEnum("f_b"),
	F_C: FObjectChoiceEnum("f_c"),
}

func (obj *fObject) Choice() FObjectChoiceEnum {
	return FObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *fObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *fObject) SetChoice(value FObjectChoiceEnum) FObject {
	intValue, ok := openapi.FObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.FObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FB = nil
	obj.obj.FA = nil

	if value == FObjectChoice.F_A {
		defaultValue := "some string"
		obj.obj.FA = &defaultValue
	}

	if value == FObjectChoice.F_B {
		defaultValue := float64(3.0)
		obj.obj.FB = &defaultValue
	}

	return obj
}

// description is TBD
// FA returns a string
func (obj *fObject) FA() string {

	if obj.obj.FA == nil {
		obj.SetChoice(FObjectChoice.F_A)
	}

	return *obj.obj.FA

}

// description is TBD
// FA returns a string
func (obj *fObject) HasFA() bool {
	return obj.obj.FA != nil
}

// description is TBD
// SetFA sets the string value in the FObject object
func (obj *fObject) SetFA(value string) FObject {
	obj.SetChoice(FObjectChoice.F_A)
	obj.obj.FA = &value
	return obj
}

// description is TBD
// FB returns a float64
func (obj *fObject) FB() float64 {

	if obj.obj.FB == nil {
		obj.SetChoice(FObjectChoice.F_B)
	}

	return *obj.obj.FB

}

// description is TBD
// FB returns a float64
func (obj *fObject) HasFB() bool {
	return obj.obj.FB != nil
}

// description is TBD
// SetFB sets the float64 value in the FObject object
func (obj *fObject) SetFB(value float64) FObject {
	obj.SetChoice(FObjectChoice.F_B)
	obj.obj.FB = &value
	return obj
}

func (obj *fObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *fObject) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(FObjectChoice.F_A)

	}

}

// ***** GObject *****
type gObject struct {
	validation
	obj *openapi.GObject
}

func NewGObject() GObject {
	obj := gObject{obj: &openapi.GObject{}}
	obj.setDefault()
	return &obj
}

func (obj *gObject) Msg() *openapi.GObject {
	return obj.obj
}

func (obj *gObject) SetMsg(msg *openapi.GObject) GObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *gObject) ToProto() (*openapi.GObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *gObject) FromProto(msg *openapi.GObject) (GObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *gObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *gObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *gObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *gObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *gObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *gObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
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

func (obj *gObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *gObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *gObject) Clone() (GObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewGObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
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
	// Msg marshals GObject to protobuf object *openapi.GObject
	// and doesn't set defaults
	Msg() *openapi.GObject
	// SetMsg unmarshals GObject from protobuf object *openapi.GObject
	// and doesn't set defaults
	SetMsg(*openapi.GObject) GObject
	// ToProto marshals GObject to protobuf object *openapi.GObject
	ToProto() (*openapi.GObject, error)
	// ToPbText marshals GObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals GObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals GObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals GObject from protobuf object *openapi.GObject
	FromProto(msg *openapi.GObject) (GObject, error)
	// FromPbText unmarshals GObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals GObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals GObject from JSON text
	FromJson(value string) error
	// Validate validates GObject
	Validate() error
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
	// SetChoice assigns GObjectChoiceEnum provided by user to GObject
	SetChoice(value GObjectChoiceEnum) GObject
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

func (obj *gObject) SetChoice(value GObjectChoiceEnum) GObject {
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
		obj.SetChoice(GObjectChoice.G_D)
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
	obj.SetChoice(GObjectChoice.G_D)
	obj.obj.GD = &value
	return obj
}

// description is TBD
// GE returns a float64
func (obj *gObject) GE() float64 {

	if obj.obj.GE == nil {
		obj.SetChoice(GObjectChoice.G_E)
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
	obj.SetChoice(GObjectChoice.G_E)
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
		obj.SetChoice(GObjectChoice.G_D)

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

// ***** JObject *****
type jObject struct {
	validation
	obj      *openapi.JObject
	jAHolder EObject
	jBHolder FObject
}

func NewJObject() JObject {
	obj := jObject{obj: &openapi.JObject{}}
	obj.setDefault()
	return &obj
}

func (obj *jObject) Msg() *openapi.JObject {
	return obj.obj
}

func (obj *jObject) SetMsg(msg *openapi.JObject) JObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *jObject) ToProto() (*openapi.JObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *jObject) FromProto(msg *openapi.JObject) (JObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *jObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *jObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *jObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *jObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *jObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *jObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *jObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *jObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *jObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *jObject) Clone() (JObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewJObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *jObject) setNil() {
	obj.jAHolder = nil
	obj.jBHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// JObject is description is TBD
type JObject interface {
	Validation
	// Msg marshals JObject to protobuf object *openapi.JObject
	// and doesn't set defaults
	Msg() *openapi.JObject
	// SetMsg unmarshals JObject from protobuf object *openapi.JObject
	// and doesn't set defaults
	SetMsg(*openapi.JObject) JObject
	// ToProto marshals JObject to protobuf object *openapi.JObject
	ToProto() (*openapi.JObject, error)
	// ToPbText marshals JObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals JObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals JObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals JObject from protobuf object *openapi.JObject
	FromProto(msg *openapi.JObject) (JObject, error)
	// FromPbText unmarshals JObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals JObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals JObject from JSON text
	FromJson(value string) error
	// Validate validates JObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (JObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns JObjectChoiceEnum, set in JObject
	Choice() JObjectChoiceEnum
	// SetChoice assigns JObjectChoiceEnum provided by user to JObject
	SetChoice(value JObjectChoiceEnum) JObject
	// HasChoice checks if Choice has been set in JObject
	HasChoice() bool
	// JA returns EObject, set in JObject.
	// EObject is description is TBD
	JA() EObject
	// SetJA assigns EObject provided by user to JObject.
	// EObject is description is TBD
	SetJA(value EObject) JObject
	// HasJA checks if JA has been set in JObject
	HasJA() bool
	// JB returns FObject, set in JObject.
	// FObject is description is TBD
	JB() FObject
	// SetJB assigns FObject provided by user to JObject.
	// FObject is description is TBD
	SetJB(value FObject) JObject
	// HasJB checks if JB has been set in JObject
	HasJB() bool
	setNil()
}

type JObjectChoiceEnum string

// Enum of Choice on JObject
var JObjectChoice = struct {
	J_A JObjectChoiceEnum
	J_B JObjectChoiceEnum
}{
	J_A: JObjectChoiceEnum("j_a"),
	J_B: JObjectChoiceEnum("j_b"),
}

func (obj *jObject) Choice() JObjectChoiceEnum {
	return JObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *jObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *jObject) SetChoice(value JObjectChoiceEnum) JObject {
	intValue, ok := openapi.JObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on JObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.JObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.JB = nil
	obj.jBHolder = nil
	obj.obj.JA = nil
	obj.jAHolder = nil

	if value == JObjectChoice.J_A {
		obj.obj.JA = NewEObject().Msg()
	}

	if value == JObjectChoice.J_B {
		obj.obj.JB = NewFObject().Msg()
	}

	return obj
}

// description is TBD
// JA returns a EObject
func (obj *jObject) JA() EObject {
	if obj.obj.JA == nil {
		obj.SetChoice(JObjectChoice.J_A)
	}
	if obj.jAHolder == nil {
		obj.jAHolder = &eObject{obj: obj.obj.JA}
	}
	return obj.jAHolder
}

// description is TBD
// JA returns a EObject
func (obj *jObject) HasJA() bool {
	return obj.obj.JA != nil
}

// description is TBD
// SetJA sets the EObject value in the JObject object
func (obj *jObject) SetJA(value EObject) JObject {
	obj.SetChoice(JObjectChoice.J_A)
	obj.jAHolder = nil
	obj.obj.JA = value.Msg()

	return obj
}

// description is TBD
// JB returns a FObject
func (obj *jObject) JB() FObject {
	if obj.obj.JB == nil {
		obj.SetChoice(JObjectChoice.J_B)
	}
	if obj.jBHolder == nil {
		obj.jBHolder = &fObject{obj: obj.obj.JB}
	}
	return obj.jBHolder
}

// description is TBD
// JB returns a FObject
func (obj *jObject) HasJB() bool {
	return obj.obj.JB != nil
}

// description is TBD
// SetJB sets the FObject value in the JObject object
func (obj *jObject) SetJB(value FObject) JObject {
	obj.SetChoice(JObjectChoice.J_B)
	obj.jBHolder = nil
	obj.obj.JB = value.Msg()

	return obj
}

func (obj *jObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Choice.Number() == 2 {
		obj.addWarnings("J_B enum in property Choice is deprecated, use j_a instead")
	}

	if obj.obj.JA != nil {

		obj.JA().validateObj(vObj, set_default)
	}

	if obj.obj.JB != nil {

		obj.JB().validateObj(vObj, set_default)
	}

}

func (obj *jObject) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(JObjectChoice.J_A)

	}

}

// ***** KObject *****
type kObject struct {
	validation
	obj           *openapi.KObject
	eObjectHolder EObject
	fObjectHolder FObject
}

func NewKObject() KObject {
	obj := kObject{obj: &openapi.KObject{}}
	obj.setDefault()
	return &obj
}

func (obj *kObject) Msg() *openapi.KObject {
	return obj.obj
}

func (obj *kObject) SetMsg(msg *openapi.KObject) KObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *kObject) ToProto() (*openapi.KObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *kObject) FromProto(msg *openapi.KObject) (KObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *kObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *kObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *kObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *kObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *kObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *kObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
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

func (obj *kObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *kObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *kObject) Clone() (KObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewKObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
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
	// Msg marshals KObject to protobuf object *openapi.KObject
	// and doesn't set defaults
	Msg() *openapi.KObject
	// SetMsg unmarshals KObject from protobuf object *openapi.KObject
	// and doesn't set defaults
	SetMsg(*openapi.KObject) KObject
	// ToProto marshals KObject to protobuf object *openapi.KObject
	ToProto() (*openapi.KObject, error)
	// ToPbText marshals KObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals KObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals KObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals KObject from protobuf object *openapi.KObject
	FromProto(msg *openapi.KObject) (KObject, error)
	// FromPbText unmarshals KObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals KObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals KObject from JSON text
	FromJson(value string) error
	// Validate validates KObject
	Validate() error
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
		obj.obj.EObject = NewEObject().Msg()
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
	obj.obj.EObject = value.Msg()

	return obj
}

// description is TBD
// FObject returns a FObject
func (obj *kObject) FObject() FObject {
	if obj.obj.FObject == nil {
		obj.obj.FObject = NewFObject().Msg()
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
	obj.obj.FObject = value.Msg()

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

// ***** LObject *****
type lObject struct {
	validation
	obj *openapi.LObject
}

func NewLObject() LObject {
	obj := lObject{obj: &openapi.LObject{}}
	obj.setDefault()
	return &obj
}

func (obj *lObject) Msg() *openapi.LObject {
	return obj.obj
}

func (obj *lObject) SetMsg(msg *openapi.LObject) LObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *lObject) ToProto() (*openapi.LObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *lObject) FromProto(msg *openapi.LObject) (LObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *lObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *lObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *lObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *lObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *lObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *lObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *lObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lObject) Clone() (LObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// LObject is format validation object
type LObject interface {
	Validation
	// Msg marshals LObject to protobuf object *openapi.LObject
	// and doesn't set defaults
	Msg() *openapi.LObject
	// SetMsg unmarshals LObject from protobuf object *openapi.LObject
	// and doesn't set defaults
	SetMsg(*openapi.LObject) LObject
	// ToProto marshals LObject to protobuf object *openapi.LObject
	ToProto() (*openapi.LObject, error)
	// ToPbText marshals LObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals LObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals LObject from protobuf object *openapi.LObject
	FromProto(msg *openapi.LObject) (LObject, error)
	// FromPbText unmarshals LObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LObject from JSON text
	FromJson(value string) error
	// Validate validates LObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StringParam returns string, set in LObject.
	StringParam() string
	// SetStringParam assigns string provided by user to LObject
	SetStringParam(value string) LObject
	// HasStringParam checks if StringParam has been set in LObject
	HasStringParam() bool
	// Integer returns int32, set in LObject.
	Integer() int32
	// SetInteger assigns int32 provided by user to LObject
	SetInteger(value int32) LObject
	// HasInteger checks if Integer has been set in LObject
	HasInteger() bool
	// Float returns float32, set in LObject.
	Float() float32
	// SetFloat assigns float32 provided by user to LObject
	SetFloat(value float32) LObject
	// HasFloat checks if Float has been set in LObject
	HasFloat() bool
	// Double returns float64, set in LObject.
	Double() float64
	// SetDouble assigns float64 provided by user to LObject
	SetDouble(value float64) LObject
	// HasDouble checks if Double has been set in LObject
	HasDouble() bool
	// Mac returns string, set in LObject.
	Mac() string
	// SetMac assigns string provided by user to LObject
	SetMac(value string) LObject
	// HasMac checks if Mac has been set in LObject
	HasMac() bool
	// Ipv4 returns string, set in LObject.
	Ipv4() string
	// SetIpv4 assigns string provided by user to LObject
	SetIpv4(value string) LObject
	// HasIpv4 checks if Ipv4 has been set in LObject
	HasIpv4() bool
	// Ipv6 returns string, set in LObject.
	Ipv6() string
	// SetIpv6 assigns string provided by user to LObject
	SetIpv6(value string) LObject
	// HasIpv6 checks if Ipv6 has been set in LObject
	HasIpv6() bool
	// Hex returns string, set in LObject.
	Hex() string
	// SetHex assigns string provided by user to LObject
	SetHex(value string) LObject
	// HasHex checks if Hex has been set in LObject
	HasHex() bool
}

// description is TBD
// StringParam returns a string
func (obj *lObject) StringParam() string {

	return *obj.obj.StringParam

}

// description is TBD
// StringParam returns a string
func (obj *lObject) HasStringParam() bool {
	return obj.obj.StringParam != nil
}

// description is TBD
// SetStringParam sets the string value in the LObject object
func (obj *lObject) SetStringParam(value string) LObject {

	obj.obj.StringParam = &value
	return obj
}

// description is TBD
// Integer returns a int32
func (obj *lObject) Integer() int32 {

	return *obj.obj.Integer

}

// description is TBD
// Integer returns a int32
func (obj *lObject) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the int32 value in the LObject object
func (obj *lObject) SetInteger(value int32) LObject {

	obj.obj.Integer = &value
	return obj
}

// description is TBD
// Float returns a float32
func (obj *lObject) Float() float32 {

	return *obj.obj.Float

}

// description is TBD
// Float returns a float32
func (obj *lObject) HasFloat() bool {
	return obj.obj.Float != nil
}

// description is TBD
// SetFloat sets the float32 value in the LObject object
func (obj *lObject) SetFloat(value float32) LObject {

	obj.obj.Float = &value
	return obj
}

// description is TBD
// Double returns a float64
func (obj *lObject) Double() float64 {

	return *obj.obj.Double

}

// description is TBD
// Double returns a float64
func (obj *lObject) HasDouble() bool {
	return obj.obj.Double != nil
}

// description is TBD
// SetDouble sets the float64 value in the LObject object
func (obj *lObject) SetDouble(value float64) LObject {

	obj.obj.Double = &value
	return obj
}

// description is TBD
// Mac returns a string
func (obj *lObject) Mac() string {

	return *obj.obj.Mac

}

// description is TBD
// Mac returns a string
func (obj *lObject) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the string value in the LObject object
func (obj *lObject) SetMac(value string) LObject {

	obj.obj.Mac = &value
	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *lObject) Ipv4() string {

	return *obj.obj.Ipv4

}

// description is TBD
// Ipv4 returns a string
func (obj *lObject) HasIpv4() bool {
	return obj.obj.Ipv4 != nil
}

// description is TBD
// SetIpv4 sets the string value in the LObject object
func (obj *lObject) SetIpv4(value string) LObject {

	obj.obj.Ipv4 = &value
	return obj
}

// description is TBD
// Ipv6 returns a string
func (obj *lObject) Ipv6() string {

	return *obj.obj.Ipv6

}

// description is TBD
// Ipv6 returns a string
func (obj *lObject) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the string value in the LObject object
func (obj *lObject) SetIpv6(value string) LObject {

	obj.obj.Ipv6 = &value
	return obj
}

// description is TBD
// Hex returns a string
func (obj *lObject) Hex() string {

	return *obj.obj.Hex

}

// description is TBD
// Hex returns a string
func (obj *lObject) HasHex() bool {
	return obj.obj.Hex != nil
}

// description is TBD
// SetHex sets the string value in the LObject object
func (obj *lObject) SetHex(value string) LObject {

	obj.obj.Hex = &value
	return obj
}

func (obj *lObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		if *obj.obj.Integer < -10 || *obj.obj.Integer > 90 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-10 <= LObject.Integer <= 90 but Got %d", *obj.obj.Integer))
		}

	}

	if obj.obj.Mac != nil {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Mac"))
		}

	}

	if obj.obj.Ipv4 != nil {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv4"))
		}

	}

	if obj.obj.Ipv6 != nil {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Ipv6"))
		}

	}

	if obj.obj.Hex != nil {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on LObject.Hex"))
		}

	}

}

func (obj *lObject) setDefault() {

}

// ***** LevelOne *****
type levelOne struct {
	validation
	obj        *openapi.LevelOne
	l1P1Holder LevelTwo
	l1P2Holder LevelFour
}

func NewLevelOne() LevelOne {
	obj := levelOne{obj: &openapi.LevelOne{}}
	obj.setDefault()
	return &obj
}

func (obj *levelOne) Msg() *openapi.LevelOne {
	return obj.obj
}

func (obj *levelOne) SetMsg(msg *openapi.LevelOne) LevelOne {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *levelOne) ToProto() (*openapi.LevelOne, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *levelOne) FromProto(msg *openapi.LevelOne) (LevelOne, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *levelOne) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *levelOne) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *levelOne) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelOne) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *levelOne) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelOne) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *levelOne) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelOne) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelOne) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelOne) Clone() (LevelOne, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelOne()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *levelOne) setNil() {
	obj.l1P1Holder = nil
	obj.l1P2Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelOne is to Test Multi level non-primitive types
type LevelOne interface {
	Validation
	// Msg marshals LevelOne to protobuf object *openapi.LevelOne
	// and doesn't set defaults
	Msg() *openapi.LevelOne
	// SetMsg unmarshals LevelOne from protobuf object *openapi.LevelOne
	// and doesn't set defaults
	SetMsg(*openapi.LevelOne) LevelOne
	// ToProto marshals LevelOne to protobuf object *openapi.LevelOne
	ToProto() (*openapi.LevelOne, error)
	// ToPbText marshals LevelOne to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelOne to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelOne to JSON text
	ToJson() (string, error)
	// FromProto unmarshals LevelOne from protobuf object *openapi.LevelOne
	FromProto(msg *openapi.LevelOne) (LevelOne, error)
	// FromPbText unmarshals LevelOne from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelOne from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelOne from JSON text
	FromJson(value string) error
	// Validate validates LevelOne
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelOne, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L1P1 returns LevelTwo, set in LevelOne.
	// LevelTwo is test Level 2
	L1P1() LevelTwo
	// SetL1P1 assigns LevelTwo provided by user to LevelOne.
	// LevelTwo is test Level 2
	SetL1P1(value LevelTwo) LevelOne
	// HasL1P1 checks if L1P1 has been set in LevelOne
	HasL1P1() bool
	// L1P2 returns LevelFour, set in LevelOne.
	// LevelFour is test level4 redundant junk testing
	L1P2() LevelFour
	// SetL1P2 assigns LevelFour provided by user to LevelOne.
	// LevelFour is test level4 redundant junk testing
	SetL1P2(value LevelFour) LevelOne
	// HasL1P2 checks if L1P2 has been set in LevelOne
	HasL1P2() bool
	setNil()
}

// Level one
// L1P1 returns a LevelTwo
func (obj *levelOne) L1P1() LevelTwo {
	if obj.obj.L1P1 == nil {
		obj.obj.L1P1 = NewLevelTwo().Msg()
	}
	if obj.l1P1Holder == nil {
		obj.l1P1Holder = &levelTwo{obj: obj.obj.L1P1}
	}
	return obj.l1P1Holder
}

// Level one
// L1P1 returns a LevelTwo
func (obj *levelOne) HasL1P1() bool {
	return obj.obj.L1P1 != nil
}

// Level one
// SetL1P1 sets the LevelTwo value in the LevelOne object
func (obj *levelOne) SetL1P1(value LevelTwo) LevelOne {

	obj.l1P1Holder = nil
	obj.obj.L1P1 = value.Msg()

	return obj
}

// Level one to four
// L1P2 returns a LevelFour
func (obj *levelOne) L1P2() LevelFour {
	if obj.obj.L1P2 == nil {
		obj.obj.L1P2 = NewLevelFour().Msg()
	}
	if obj.l1P2Holder == nil {
		obj.l1P2Holder = &levelFour{obj: obj.obj.L1P2}
	}
	return obj.l1P2Holder
}

// Level one to four
// L1P2 returns a LevelFour
func (obj *levelOne) HasL1P2() bool {
	return obj.obj.L1P2 != nil
}

// Level one to four
// SetL1P2 sets the LevelFour value in the LevelOne object
func (obj *levelOne) SetL1P2(value LevelFour) LevelOne {

	obj.l1P2Holder = nil
	obj.obj.L1P2 = value.Msg()

	return obj
}

func (obj *levelOne) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L1P1 != nil {

		obj.L1P1().validateObj(vObj, set_default)
	}

	if obj.obj.L1P2 != nil {

		obj.L1P2().validateObj(vObj, set_default)
	}

}

func (obj *levelOne) setDefault() {

}

// ***** Mandate *****
type mandate struct {
	validation
	obj *openapi.Mandate
}

func NewMandate() Mandate {
	obj := mandate{obj: &openapi.Mandate{}}
	obj.setDefault()
	return &obj
}

func (obj *mandate) Msg() *openapi.Mandate {
	return obj.obj
}

func (obj *mandate) SetMsg(msg *openapi.Mandate) Mandate {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *mandate) ToProto() (*openapi.Mandate, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *mandate) FromProto(msg *openapi.Mandate) (Mandate, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *mandate) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *mandate) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *mandate) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *mandate) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *mandate) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *mandate) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *mandate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mandate) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mandate) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mandate) Clone() (Mandate, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMandate()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Mandate is object to Test required Parameter
type Mandate interface {
	Validation
	// Msg marshals Mandate to protobuf object *openapi.Mandate
	// and doesn't set defaults
	Msg() *openapi.Mandate
	// SetMsg unmarshals Mandate from protobuf object *openapi.Mandate
	// and doesn't set defaults
	SetMsg(*openapi.Mandate) Mandate
	// ToProto marshals Mandate to protobuf object *openapi.Mandate
	ToProto() (*openapi.Mandate, error)
	// ToPbText marshals Mandate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Mandate to YAML text
	ToYaml() (string, error)
	// ToJson marshals Mandate to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Mandate from protobuf object *openapi.Mandate
	FromProto(msg *openapi.Mandate) (Mandate, error)
	// FromPbText unmarshals Mandate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Mandate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Mandate from JSON text
	FromJson(value string) error
	// Validate validates Mandate
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Mandate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RequiredParam returns string, set in Mandate.
	RequiredParam() string
	// SetRequiredParam assigns string provided by user to Mandate
	SetRequiredParam(value string) Mandate
}

// description is TBD
// RequiredParam returns a string
func (obj *mandate) RequiredParam() string {

	return obj.obj.RequiredParam
}

// description is TBD
// SetRequiredParam sets the string value in the Mandate object
func (obj *mandate) SetRequiredParam(value string) Mandate {

	obj.obj.RequiredParam = value
	return obj
}

func (obj *mandate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// RequiredParam is required
	if obj.obj.RequiredParam == "" {
		vObj.validationErrors = append(vObj.validationErrors, "RequiredParam is required field on interface Mandate")
	}
}

func (obj *mandate) setDefault() {

}

// ***** Ipv4Pattern *****
type ipv4Pattern struct {
	validation
	obj        *openapi.Ipv4Pattern
	ipv4Holder PatternIpv4PatternIpv4
}

func NewIpv4Pattern() Ipv4Pattern {
	obj := ipv4Pattern{obj: &openapi.Ipv4Pattern{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv4Pattern) Msg() *openapi.Ipv4Pattern {
	return obj.obj
}

func (obj *ipv4Pattern) SetMsg(msg *openapi.Ipv4Pattern) Ipv4Pattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *ipv4Pattern) ToProto() (*openapi.Ipv4Pattern, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *ipv4Pattern) FromProto(msg *openapi.Ipv4Pattern) (Ipv4Pattern, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *ipv4Pattern) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *ipv4Pattern) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *ipv4Pattern) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *ipv4Pattern) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *ipv4Pattern) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *ipv4Pattern) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
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

func (obj *ipv4Pattern) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv4Pattern) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv4Pattern) Clone() (Ipv4Pattern, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv4Pattern()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
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
	// Msg marshals Ipv4Pattern to protobuf object *openapi.Ipv4Pattern
	// and doesn't set defaults
	Msg() *openapi.Ipv4Pattern
	// SetMsg unmarshals Ipv4Pattern from protobuf object *openapi.Ipv4Pattern
	// and doesn't set defaults
	SetMsg(*openapi.Ipv4Pattern) Ipv4Pattern
	// ToProto marshals Ipv4Pattern to protobuf object *openapi.Ipv4Pattern
	ToProto() (*openapi.Ipv4Pattern, error)
	// ToPbText marshals Ipv4Pattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv4Pattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv4Pattern to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Ipv4Pattern from protobuf object *openapi.Ipv4Pattern
	FromProto(msg *openapi.Ipv4Pattern) (Ipv4Pattern, error)
	// FromPbText unmarshals Ipv4Pattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv4Pattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv4Pattern from JSON text
	FromJson(value string) error
	// Validate validates Ipv4Pattern
	Validate() error
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
		obj.obj.Ipv4 = NewPatternIpv4PatternIpv4().Msg()
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
	obj.obj.Ipv4 = value.Msg()

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

// ***** Ipv6Pattern *****
type ipv6Pattern struct {
	validation
	obj        *openapi.Ipv6Pattern
	ipv6Holder PatternIpv6PatternIpv6
}

func NewIpv6Pattern() Ipv6Pattern {
	obj := ipv6Pattern{obj: &openapi.Ipv6Pattern{}}
	obj.setDefault()
	return &obj
}

func (obj *ipv6Pattern) Msg() *openapi.Ipv6Pattern {
	return obj.obj
}

func (obj *ipv6Pattern) SetMsg(msg *openapi.Ipv6Pattern) Ipv6Pattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *ipv6Pattern) ToProto() (*openapi.Ipv6Pattern, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *ipv6Pattern) FromProto(msg *openapi.Ipv6Pattern) (Ipv6Pattern, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *ipv6Pattern) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *ipv6Pattern) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *ipv6Pattern) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *ipv6Pattern) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *ipv6Pattern) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *ipv6Pattern) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *ipv6Pattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ipv6Pattern) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ipv6Pattern) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ipv6Pattern) Clone() (Ipv6Pattern, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIpv6Pattern()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *ipv6Pattern) setNil() {
	obj.ipv6Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ipv6Pattern is test ipv6 pattern
type Ipv6Pattern interface {
	Validation
	// Msg marshals Ipv6Pattern to protobuf object *openapi.Ipv6Pattern
	// and doesn't set defaults
	Msg() *openapi.Ipv6Pattern
	// SetMsg unmarshals Ipv6Pattern from protobuf object *openapi.Ipv6Pattern
	// and doesn't set defaults
	SetMsg(*openapi.Ipv6Pattern) Ipv6Pattern
	// ToProto marshals Ipv6Pattern to protobuf object *openapi.Ipv6Pattern
	ToProto() (*openapi.Ipv6Pattern, error)
	// ToPbText marshals Ipv6Pattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ipv6Pattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ipv6Pattern to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Ipv6Pattern from protobuf object *openapi.Ipv6Pattern
	FromProto(msg *openapi.Ipv6Pattern) (Ipv6Pattern, error)
	// FromPbText unmarshals Ipv6Pattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ipv6Pattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ipv6Pattern from JSON text
	FromJson(value string) error
	// Validate validates Ipv6Pattern
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ipv6Pattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv6 returns PatternIpv6PatternIpv6, set in Ipv6Pattern.
	// PatternIpv6PatternIpv6 is tBD
	Ipv6() PatternIpv6PatternIpv6
	// SetIpv6 assigns PatternIpv6PatternIpv6 provided by user to Ipv6Pattern.
	// PatternIpv6PatternIpv6 is tBD
	SetIpv6(value PatternIpv6PatternIpv6) Ipv6Pattern
	// HasIpv6 checks if Ipv6 has been set in Ipv6Pattern
	HasIpv6() bool
	setNil()
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternIpv6
func (obj *ipv6Pattern) Ipv6() PatternIpv6PatternIpv6 {
	if obj.obj.Ipv6 == nil {
		obj.obj.Ipv6 = NewPatternIpv6PatternIpv6().Msg()
	}
	if obj.ipv6Holder == nil {
		obj.ipv6Holder = &patternIpv6PatternIpv6{obj: obj.obj.Ipv6}
	}
	return obj.ipv6Holder
}

// description is TBD
// Ipv6 returns a PatternIpv6PatternIpv6
func (obj *ipv6Pattern) HasIpv6() bool {
	return obj.obj.Ipv6 != nil
}

// description is TBD
// SetIpv6 sets the PatternIpv6PatternIpv6 value in the Ipv6Pattern object
func (obj *ipv6Pattern) SetIpv6(value PatternIpv6PatternIpv6) Ipv6Pattern {

	obj.ipv6Holder = nil
	obj.obj.Ipv6 = value.Msg()

	return obj
}

func (obj *ipv6Pattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv6 != nil {

		obj.Ipv6().validateObj(vObj, set_default)
	}

}

func (obj *ipv6Pattern) setDefault() {

}

// ***** MacPattern *****
type macPattern struct {
	validation
	obj       *openapi.MacPattern
	macHolder PatternMacPatternMac
}

func NewMacPattern() MacPattern {
	obj := macPattern{obj: &openapi.MacPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *macPattern) Msg() *openapi.MacPattern {
	return obj.obj
}

func (obj *macPattern) SetMsg(msg *openapi.MacPattern) MacPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *macPattern) ToProto() (*openapi.MacPattern, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *macPattern) FromProto(msg *openapi.MacPattern) (MacPattern, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *macPattern) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *macPattern) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *macPattern) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *macPattern) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *macPattern) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *macPattern) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *macPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *macPattern) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *macPattern) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *macPattern) Clone() (MacPattern, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMacPattern()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *macPattern) setNil() {
	obj.macHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// MacPattern is test mac pattern
type MacPattern interface {
	Validation
	// Msg marshals MacPattern to protobuf object *openapi.MacPattern
	// and doesn't set defaults
	Msg() *openapi.MacPattern
	// SetMsg unmarshals MacPattern from protobuf object *openapi.MacPattern
	// and doesn't set defaults
	SetMsg(*openapi.MacPattern) MacPattern
	// ToProto marshals MacPattern to protobuf object *openapi.MacPattern
	ToProto() (*openapi.MacPattern, error)
	// ToPbText marshals MacPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MacPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals MacPattern to JSON text
	ToJson() (string, error)
	// FromProto unmarshals MacPattern from protobuf object *openapi.MacPattern
	FromProto(msg *openapi.MacPattern) (MacPattern, error)
	// FromPbText unmarshals MacPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MacPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MacPattern from JSON text
	FromJson(value string) error
	// Validate validates MacPattern
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MacPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Mac returns PatternMacPatternMac, set in MacPattern.
	// PatternMacPatternMac is tBD
	Mac() PatternMacPatternMac
	// SetMac assigns PatternMacPatternMac provided by user to MacPattern.
	// PatternMacPatternMac is tBD
	SetMac(value PatternMacPatternMac) MacPattern
	// HasMac checks if Mac has been set in MacPattern
	HasMac() bool
	setNil()
}

// description is TBD
// Mac returns a PatternMacPatternMac
func (obj *macPattern) Mac() PatternMacPatternMac {
	if obj.obj.Mac == nil {
		obj.obj.Mac = NewPatternMacPatternMac().Msg()
	}
	if obj.macHolder == nil {
		obj.macHolder = &patternMacPatternMac{obj: obj.obj.Mac}
	}
	return obj.macHolder
}

// description is TBD
// Mac returns a PatternMacPatternMac
func (obj *macPattern) HasMac() bool {
	return obj.obj.Mac != nil
}

// description is TBD
// SetMac sets the PatternMacPatternMac value in the MacPattern object
func (obj *macPattern) SetMac(value PatternMacPatternMac) MacPattern {

	obj.macHolder = nil
	obj.obj.Mac = value.Msg()

	return obj
}

func (obj *macPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mac != nil {

		obj.Mac().validateObj(vObj, set_default)
	}

}

func (obj *macPattern) setDefault() {

}

// ***** IntegerPattern *****
type integerPattern struct {
	validation
	obj           *openapi.IntegerPattern
	integerHolder PatternIntegerPatternInteger
}

func NewIntegerPattern() IntegerPattern {
	obj := integerPattern{obj: &openapi.IntegerPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *integerPattern) Msg() *openapi.IntegerPattern {
	return obj.obj
}

func (obj *integerPattern) SetMsg(msg *openapi.IntegerPattern) IntegerPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *integerPattern) ToProto() (*openapi.IntegerPattern, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *integerPattern) FromProto(msg *openapi.IntegerPattern) (IntegerPattern, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *integerPattern) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *integerPattern) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *integerPattern) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *integerPattern) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *integerPattern) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *integerPattern) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *integerPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *integerPattern) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *integerPattern) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *integerPattern) Clone() (IntegerPattern, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIntegerPattern()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *integerPattern) setNil() {
	obj.integerHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IntegerPattern is test integer pattern
type IntegerPattern interface {
	Validation
	// Msg marshals IntegerPattern to protobuf object *openapi.IntegerPattern
	// and doesn't set defaults
	Msg() *openapi.IntegerPattern
	// SetMsg unmarshals IntegerPattern from protobuf object *openapi.IntegerPattern
	// and doesn't set defaults
	SetMsg(*openapi.IntegerPattern) IntegerPattern
	// ToProto marshals IntegerPattern to protobuf object *openapi.IntegerPattern
	ToProto() (*openapi.IntegerPattern, error)
	// ToPbText marshals IntegerPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IntegerPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals IntegerPattern to JSON text
	ToJson() (string, error)
	// FromProto unmarshals IntegerPattern from protobuf object *openapi.IntegerPattern
	FromProto(msg *openapi.IntegerPattern) (IntegerPattern, error)
	// FromPbText unmarshals IntegerPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IntegerPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IntegerPattern from JSON text
	FromJson(value string) error
	// Validate validates IntegerPattern
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IntegerPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Integer returns PatternIntegerPatternInteger, set in IntegerPattern.
	// PatternIntegerPatternInteger is tBD
	Integer() PatternIntegerPatternInteger
	// SetInteger assigns PatternIntegerPatternInteger provided by user to IntegerPattern.
	// PatternIntegerPatternInteger is tBD
	SetInteger(value PatternIntegerPatternInteger) IntegerPattern
	// HasInteger checks if Integer has been set in IntegerPattern
	HasInteger() bool
	setNil()
}

// description is TBD
// Integer returns a PatternIntegerPatternInteger
func (obj *integerPattern) Integer() PatternIntegerPatternInteger {
	if obj.obj.Integer == nil {
		obj.obj.Integer = NewPatternIntegerPatternInteger().Msg()
	}
	if obj.integerHolder == nil {
		obj.integerHolder = &patternIntegerPatternInteger{obj: obj.obj.Integer}
	}
	return obj.integerHolder
}

// description is TBD
// Integer returns a PatternIntegerPatternInteger
func (obj *integerPattern) HasInteger() bool {
	return obj.obj.Integer != nil
}

// description is TBD
// SetInteger sets the PatternIntegerPatternInteger value in the IntegerPattern object
func (obj *integerPattern) SetInteger(value PatternIntegerPatternInteger) IntegerPattern {

	obj.integerHolder = nil
	obj.obj.Integer = value.Msg()

	return obj
}

func (obj *integerPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Integer != nil {

		obj.Integer().validateObj(vObj, set_default)
	}

}

func (obj *integerPattern) setDefault() {

}

// ***** ChecksumPattern *****
type checksumPattern struct {
	validation
	obj            *openapi.ChecksumPattern
	checksumHolder PatternChecksumPatternChecksum
}

func NewChecksumPattern() ChecksumPattern {
	obj := checksumPattern{obj: &openapi.ChecksumPattern{}}
	obj.setDefault()
	return &obj
}

func (obj *checksumPattern) Msg() *openapi.ChecksumPattern {
	return obj.obj
}

func (obj *checksumPattern) SetMsg(msg *openapi.ChecksumPattern) ChecksumPattern {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *checksumPattern) ToProto() (*openapi.ChecksumPattern, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *checksumPattern) FromProto(msg *openapi.ChecksumPattern) (ChecksumPattern, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *checksumPattern) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *checksumPattern) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *checksumPattern) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *checksumPattern) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *checksumPattern) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *checksumPattern) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *checksumPattern) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *checksumPattern) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *checksumPattern) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *checksumPattern) Clone() (ChecksumPattern, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChecksumPattern()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *checksumPattern) setNil() {
	obj.checksumHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChecksumPattern is test checksum pattern
type ChecksumPattern interface {
	Validation
	// Msg marshals ChecksumPattern to protobuf object *openapi.ChecksumPattern
	// and doesn't set defaults
	Msg() *openapi.ChecksumPattern
	// SetMsg unmarshals ChecksumPattern from protobuf object *openapi.ChecksumPattern
	// and doesn't set defaults
	SetMsg(*openapi.ChecksumPattern) ChecksumPattern
	// ToProto marshals ChecksumPattern to protobuf object *openapi.ChecksumPattern
	ToProto() (*openapi.ChecksumPattern, error)
	// ToPbText marshals ChecksumPattern to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChecksumPattern to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChecksumPattern to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ChecksumPattern from protobuf object *openapi.ChecksumPattern
	FromProto(msg *openapi.ChecksumPattern) (ChecksumPattern, error)
	// FromPbText unmarshals ChecksumPattern from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChecksumPattern from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChecksumPattern from JSON text
	FromJson(value string) error
	// Validate validates ChecksumPattern
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChecksumPattern, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Checksum returns PatternChecksumPatternChecksum, set in ChecksumPattern.
	// PatternChecksumPatternChecksum is tBD
	Checksum() PatternChecksumPatternChecksum
	// SetChecksum assigns PatternChecksumPatternChecksum provided by user to ChecksumPattern.
	// PatternChecksumPatternChecksum is tBD
	SetChecksum(value PatternChecksumPatternChecksum) ChecksumPattern
	// HasChecksum checks if Checksum has been set in ChecksumPattern
	HasChecksum() bool
	setNil()
}

// description is TBD
// Checksum returns a PatternChecksumPatternChecksum
func (obj *checksumPattern) Checksum() PatternChecksumPatternChecksum {
	if obj.obj.Checksum == nil {
		obj.obj.Checksum = NewPatternChecksumPatternChecksum().Msg()
	}
	if obj.checksumHolder == nil {
		obj.checksumHolder = &patternChecksumPatternChecksum{obj: obj.obj.Checksum}
	}
	return obj.checksumHolder
}

// description is TBD
// Checksum returns a PatternChecksumPatternChecksum
func (obj *checksumPattern) HasChecksum() bool {
	return obj.obj.Checksum != nil
}

// description is TBD
// SetChecksum sets the PatternChecksumPatternChecksum value in the ChecksumPattern object
func (obj *checksumPattern) SetChecksum(value PatternChecksumPatternChecksum) ChecksumPattern {

	obj.checksumHolder = nil
	obj.obj.Checksum = value.Msg()

	return obj
}

func (obj *checksumPattern) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Checksum != nil {

		obj.Checksum().validateObj(vObj, set_default)
	}

}

func (obj *checksumPattern) setDefault() {

}

// ***** Layer1Ieee802X *****
type layer1Ieee802X struct {
	validation
	obj *openapi.Layer1Ieee802X
}

func NewLayer1Ieee802X() Layer1Ieee802X {
	obj := layer1Ieee802X{obj: &openapi.Layer1Ieee802X{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1Ieee802X) Msg() *openapi.Layer1Ieee802X {
	return obj.obj
}

func (obj *layer1Ieee802X) SetMsg(msg *openapi.Layer1Ieee802X) Layer1Ieee802X {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *layer1Ieee802X) ToProto() (*openapi.Layer1Ieee802X, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *layer1Ieee802X) FromProto(msg *openapi.Layer1Ieee802X) (Layer1Ieee802X, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *layer1Ieee802X) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *layer1Ieee802X) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *layer1Ieee802X) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *layer1Ieee802X) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *layer1Ieee802X) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *layer1Ieee802X) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *layer1Ieee802X) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1Ieee802X) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1Ieee802X) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1Ieee802X) Clone() (Layer1Ieee802X, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1Ieee802X()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Layer1Ieee802X is description is TBD
type Layer1Ieee802X interface {
	Validation
	// Msg marshals Layer1Ieee802X to protobuf object *openapi.Layer1Ieee802X
	// and doesn't set defaults
	Msg() *openapi.Layer1Ieee802X
	// SetMsg unmarshals Layer1Ieee802X from protobuf object *openapi.Layer1Ieee802X
	// and doesn't set defaults
	SetMsg(*openapi.Layer1Ieee802X) Layer1Ieee802X
	// ToProto marshals Layer1Ieee802X to protobuf object *openapi.Layer1Ieee802X
	ToProto() (*openapi.Layer1Ieee802X, error)
	// ToPbText marshals Layer1Ieee802X to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1Ieee802X to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1Ieee802X to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Layer1Ieee802X from protobuf object *openapi.Layer1Ieee802X
	FromProto(msg *openapi.Layer1Ieee802X) (Layer1Ieee802X, error)
	// FromPbText unmarshals Layer1Ieee802X from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1Ieee802X from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1Ieee802X from JSON text
	FromJson(value string) error
	// Validate validates Layer1Ieee802X
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1Ieee802X, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowControl returns bool, set in Layer1Ieee802X.
	FlowControl() bool
	// SetFlowControl assigns bool provided by user to Layer1Ieee802X
	SetFlowControl(value bool) Layer1Ieee802X
	// HasFlowControl checks if FlowControl has been set in Layer1Ieee802X
	HasFlowControl() bool
}

// description is TBD
// FlowControl returns a bool
func (obj *layer1Ieee802X) FlowControl() bool {

	return *obj.obj.FlowControl

}

// description is TBD
// FlowControl returns a bool
func (obj *layer1Ieee802X) HasFlowControl() bool {
	return obj.obj.FlowControl != nil
}

// description is TBD
// SetFlowControl sets the bool value in the Layer1Ieee802X object
func (obj *layer1Ieee802X) SetFlowControl(value bool) Layer1Ieee802X {

	obj.obj.FlowControl = &value
	return obj
}

func (obj *layer1Ieee802X) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *layer1Ieee802X) setDefault() {

}

// ***** MObject *****
type mObject struct {
	validation
	obj *openapi.MObject
}

func NewMObject() MObject {
	obj := mObject{obj: &openapi.MObject{}}
	obj.setDefault()
	return &obj
}

func (obj *mObject) Msg() *openapi.MObject {
	return obj.obj
}

func (obj *mObject) SetMsg(msg *openapi.MObject) MObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *mObject) ToProto() (*openapi.MObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *mObject) FromProto(msg *openapi.MObject) (MObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *mObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *mObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *mObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *mObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *mObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *mObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *mObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *mObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *mObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *mObject) Clone() (MObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// MObject is required format validation object
type MObject interface {
	Validation
	// Msg marshals MObject to protobuf object *openapi.MObject
	// and doesn't set defaults
	Msg() *openapi.MObject
	// SetMsg unmarshals MObject from protobuf object *openapi.MObject
	// and doesn't set defaults
	SetMsg(*openapi.MObject) MObject
	// ToProto marshals MObject to protobuf object *openapi.MObject
	ToProto() (*openapi.MObject, error)
	// ToPbText marshals MObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals MObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals MObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals MObject from protobuf object *openapi.MObject
	FromProto(msg *openapi.MObject) (MObject, error)
	// FromPbText unmarshals MObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals MObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals MObject from JSON text
	FromJson(value string) error
	// Validate validates MObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (MObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// StringParam returns string, set in MObject.
	StringParam() string
	// SetStringParam assigns string provided by user to MObject
	SetStringParam(value string) MObject
	// Integer returns int32, set in MObject.
	Integer() int32
	// SetInteger assigns int32 provided by user to MObject
	SetInteger(value int32) MObject
	// Float returns float32, set in MObject.
	Float() float32
	// SetFloat assigns float32 provided by user to MObject
	SetFloat(value float32) MObject
	// Double returns float64, set in MObject.
	Double() float64
	// SetDouble assigns float64 provided by user to MObject
	SetDouble(value float64) MObject
	// Mac returns string, set in MObject.
	Mac() string
	// SetMac assigns string provided by user to MObject
	SetMac(value string) MObject
	// Ipv4 returns string, set in MObject.
	Ipv4() string
	// SetIpv4 assigns string provided by user to MObject
	SetIpv4(value string) MObject
	// Ipv6 returns string, set in MObject.
	Ipv6() string
	// SetIpv6 assigns string provided by user to MObject
	SetIpv6(value string) MObject
	// Hex returns string, set in MObject.
	Hex() string
	// SetHex assigns string provided by user to MObject
	SetHex(value string) MObject
}

// description is TBD
// StringParam returns a string
func (obj *mObject) StringParam() string {

	return obj.obj.StringParam
}

// description is TBD
// SetStringParam sets the string value in the MObject object
func (obj *mObject) SetStringParam(value string) MObject {

	obj.obj.StringParam = value
	return obj
}

// description is TBD
// Integer returns a int32
func (obj *mObject) Integer() int32 {

	return obj.obj.Integer
}

// description is TBD
// SetInteger sets the int32 value in the MObject object
func (obj *mObject) SetInteger(value int32) MObject {

	obj.obj.Integer = value
	return obj
}

// description is TBD
// Float returns a float32
func (obj *mObject) Float() float32 {

	return obj.obj.Float
}

// description is TBD
// SetFloat sets the float32 value in the MObject object
func (obj *mObject) SetFloat(value float32) MObject {

	obj.obj.Float = value
	return obj
}

// description is TBD
// Double returns a float64
func (obj *mObject) Double() float64 {

	return obj.obj.Double
}

// description is TBD
// SetDouble sets the float64 value in the MObject object
func (obj *mObject) SetDouble(value float64) MObject {

	obj.obj.Double = value
	return obj
}

// description is TBD
// Mac returns a string
func (obj *mObject) Mac() string {

	return obj.obj.Mac
}

// description is TBD
// SetMac sets the string value in the MObject object
func (obj *mObject) SetMac(value string) MObject {

	obj.obj.Mac = value
	return obj
}

// description is TBD
// Ipv4 returns a string
func (obj *mObject) Ipv4() string {

	return obj.obj.Ipv4
}

// description is TBD
// SetIpv4 sets the string value in the MObject object
func (obj *mObject) SetIpv4(value string) MObject {

	obj.obj.Ipv4 = value
	return obj
}

// description is TBD
// Ipv6 returns a string
func (obj *mObject) Ipv6() string {

	return obj.obj.Ipv6
}

// description is TBD
// SetIpv6 sets the string value in the MObject object
func (obj *mObject) SetIpv6(value string) MObject {

	obj.obj.Ipv6 = value
	return obj
}

// description is TBD
// Hex returns a string
func (obj *mObject) Hex() string {

	return obj.obj.Hex
}

// description is TBD
// SetHex sets the string value in the MObject object
func (obj *mObject) SetHex(value string) MObject {

	obj.obj.Hex = value
	return obj
}

func (obj *mObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// StringParam is required
	if obj.obj.StringParam == "" {
		vObj.validationErrors = append(vObj.validationErrors, "StringParam is required field on interface MObject")
	}

	if obj.obj.Integer != 0 {

		if obj.obj.Integer < -10 || obj.obj.Integer > 90 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("-10 <= MObject.Integer <= 90 but Got %d", obj.obj.Integer))
		}

	}

	// Mac is required
	if obj.obj.Mac == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Mac is required field on interface MObject")
	}
	if obj.obj.Mac != "" {

		err := obj.validateMac(obj.Mac())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Mac"))
		}

	}

	// Ipv4 is required
	if obj.obj.Ipv4 == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv4 is required field on interface MObject")
	}
	if obj.obj.Ipv4 != "" {

		err := obj.validateIpv4(obj.Ipv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv4"))
		}

	}

	// Ipv6 is required
	if obj.obj.Ipv6 == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Ipv6 is required field on interface MObject")
	}
	if obj.obj.Ipv6 != "" {

		err := obj.validateIpv6(obj.Ipv6())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Ipv6"))
		}

	}

	// Hex is required
	if obj.obj.Hex == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Hex is required field on interface MObject")
	}
	if obj.obj.Hex != "" {

		err := obj.validateHex(obj.Hex())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on MObject.Hex"))
		}

	}

}

func (obj *mObject) setDefault() {

}

// ***** PatternPrefixConfigHeaderChecksum *****
type patternPrefixConfigHeaderChecksum struct {
	validation
	obj *openapi.PatternPrefixConfigHeaderChecksum
}

func NewPatternPrefixConfigHeaderChecksum() PatternPrefixConfigHeaderChecksum {
	obj := patternPrefixConfigHeaderChecksum{obj: &openapi.PatternPrefixConfigHeaderChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigHeaderChecksum) Msg() *openapi.PatternPrefixConfigHeaderChecksum {
	return obj.obj
}

func (obj *patternPrefixConfigHeaderChecksum) SetMsg(msg *openapi.PatternPrefixConfigHeaderChecksum) PatternPrefixConfigHeaderChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternPrefixConfigHeaderChecksum) ToProto() (*openapi.PatternPrefixConfigHeaderChecksum, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternPrefixConfigHeaderChecksum) FromProto(msg *openapi.PatternPrefixConfigHeaderChecksum) (PatternPrefixConfigHeaderChecksum, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternPrefixConfigHeaderChecksum) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternPrefixConfigHeaderChecksum) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternPrefixConfigHeaderChecksum) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigHeaderChecksum) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternPrefixConfigHeaderChecksum) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigHeaderChecksum) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternPrefixConfigHeaderChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigHeaderChecksum) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigHeaderChecksum) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigHeaderChecksum) Clone() (PatternPrefixConfigHeaderChecksum, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigHeaderChecksum()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternPrefixConfigHeaderChecksum is header checksum
type PatternPrefixConfigHeaderChecksum interface {
	Validation
	// Msg marshals PatternPrefixConfigHeaderChecksum to protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	// and doesn't set defaults
	Msg() *openapi.PatternPrefixConfigHeaderChecksum
	// SetMsg unmarshals PatternPrefixConfigHeaderChecksum from protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	// and doesn't set defaults
	SetMsg(*openapi.PatternPrefixConfigHeaderChecksum) PatternPrefixConfigHeaderChecksum
	// ToProto marshals PatternPrefixConfigHeaderChecksum to protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	ToProto() (*openapi.PatternPrefixConfigHeaderChecksum, error)
	// ToPbText marshals PatternPrefixConfigHeaderChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigHeaderChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigHeaderChecksum to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternPrefixConfigHeaderChecksum from protobuf object *openapi.PatternPrefixConfigHeaderChecksum
	FromProto(msg *openapi.PatternPrefixConfigHeaderChecksum) (PatternPrefixConfigHeaderChecksum, error)
	// FromPbText unmarshals PatternPrefixConfigHeaderChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigHeaderChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigHeaderChecksum from JSON text
	FromJson(value string) error
	// Validate validates PatternPrefixConfigHeaderChecksum
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigHeaderChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternPrefixConfigHeaderChecksumChoiceEnum, set in PatternPrefixConfigHeaderChecksum
	Choice() PatternPrefixConfigHeaderChecksumChoiceEnum
	// SetChoice assigns PatternPrefixConfigHeaderChecksumChoiceEnum provided by user to PatternPrefixConfigHeaderChecksum
	SetChoice(value PatternPrefixConfigHeaderChecksumChoiceEnum) PatternPrefixConfigHeaderChecksum
	// HasChoice checks if Choice has been set in PatternPrefixConfigHeaderChecksum
	HasChoice() bool
	// Generated returns PatternPrefixConfigHeaderChecksumGeneratedEnum, set in PatternPrefixConfigHeaderChecksum
	Generated() PatternPrefixConfigHeaderChecksumGeneratedEnum
	// SetGenerated assigns PatternPrefixConfigHeaderChecksumGeneratedEnum provided by user to PatternPrefixConfigHeaderChecksum
	SetGenerated(value PatternPrefixConfigHeaderChecksumGeneratedEnum) PatternPrefixConfigHeaderChecksum
	// HasGenerated checks if Generated has been set in PatternPrefixConfigHeaderChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternPrefixConfigHeaderChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternPrefixConfigHeaderChecksum
	SetCustom(value uint32) PatternPrefixConfigHeaderChecksum
	// HasCustom checks if Custom has been set in PatternPrefixConfigHeaderChecksum
	HasCustom() bool
}

type PatternPrefixConfigHeaderChecksumChoiceEnum string

// Enum of Choice on PatternPrefixConfigHeaderChecksum
var PatternPrefixConfigHeaderChecksumChoice = struct {
	GENERATED PatternPrefixConfigHeaderChecksumChoiceEnum
	CUSTOM    PatternPrefixConfigHeaderChecksumChoiceEnum
}{
	GENERATED: PatternPrefixConfigHeaderChecksumChoiceEnum("generated"),
	CUSTOM:    PatternPrefixConfigHeaderChecksumChoiceEnum("custom"),
}

func (obj *patternPrefixConfigHeaderChecksum) Choice() PatternPrefixConfigHeaderChecksumChoiceEnum {
	return PatternPrefixConfigHeaderChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternPrefixConfigHeaderChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternPrefixConfigHeaderChecksum) SetChoice(value PatternPrefixConfigHeaderChecksumChoiceEnum) PatternPrefixConfigHeaderChecksum {
	intValue, ok := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigHeaderChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigHeaderChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = openapi.PatternPrefixConfigHeaderChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternPrefixConfigHeaderChecksumGeneratedEnum string

// Enum of Generated on PatternPrefixConfigHeaderChecksum
var PatternPrefixConfigHeaderChecksumGenerated = struct {
	GOOD PatternPrefixConfigHeaderChecksumGeneratedEnum
	BAD  PatternPrefixConfigHeaderChecksumGeneratedEnum
}{
	GOOD: PatternPrefixConfigHeaderChecksumGeneratedEnum("good"),
	BAD:  PatternPrefixConfigHeaderChecksumGeneratedEnum("bad"),
}

func (obj *patternPrefixConfigHeaderChecksum) Generated() PatternPrefixConfigHeaderChecksumGeneratedEnum {
	return PatternPrefixConfigHeaderChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternPrefixConfigHeaderChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternPrefixConfigHeaderChecksum) SetGenerated(value PatternPrefixConfigHeaderChecksumGeneratedEnum) PatternPrefixConfigHeaderChecksum {
	intValue, ok := openapi.PatternPrefixConfigHeaderChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigHeaderChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigHeaderChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternPrefixConfigHeaderChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.SetChoice(PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternPrefixConfigHeaderChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternPrefixConfigHeaderChecksum object
func (obj *patternPrefixConfigHeaderChecksum) SetCustom(value uint32) PatternPrefixConfigHeaderChecksum {
	obj.SetChoice(PatternPrefixConfigHeaderChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternPrefixConfigHeaderChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigHeaderChecksum.Custom <= 65535 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternPrefixConfigHeaderChecksum) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternPrefixConfigHeaderChecksumChoice.GENERATED)
		if obj.obj.Generated.Number() == 0 {
			obj.SetGenerated(PatternPrefixConfigHeaderChecksumGenerated.GOOD)

		}

	}

}

// ***** PatternPrefixConfigAutoFieldTest *****
type patternPrefixConfigAutoFieldTest struct {
	validation
	obj             *openapi.PatternPrefixConfigAutoFieldTest
	incrementHolder PatternPrefixConfigAutoFieldTestCounter
	decrementHolder PatternPrefixConfigAutoFieldTestCounter
}

func NewPatternPrefixConfigAutoFieldTest() PatternPrefixConfigAutoFieldTest {
	obj := patternPrefixConfigAutoFieldTest{obj: &openapi.PatternPrefixConfigAutoFieldTest{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigAutoFieldTest) Msg() *openapi.PatternPrefixConfigAutoFieldTest {
	return obj.obj
}

func (obj *patternPrefixConfigAutoFieldTest) SetMsg(msg *openapi.PatternPrefixConfigAutoFieldTest) PatternPrefixConfigAutoFieldTest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternPrefixConfigAutoFieldTest) ToProto() (*openapi.PatternPrefixConfigAutoFieldTest, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternPrefixConfigAutoFieldTest) FromProto(msg *openapi.PatternPrefixConfigAutoFieldTest) (PatternPrefixConfigAutoFieldTest, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternPrefixConfigAutoFieldTest) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternPrefixConfigAutoFieldTest) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternPrefixConfigAutoFieldTest) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigAutoFieldTest) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternPrefixConfigAutoFieldTest) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigAutoFieldTest) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternPrefixConfigAutoFieldTest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTest) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTest) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigAutoFieldTest) Clone() (PatternPrefixConfigAutoFieldTest, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigAutoFieldTest()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *patternPrefixConfigAutoFieldTest) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternPrefixConfigAutoFieldTest is tBD
type PatternPrefixConfigAutoFieldTest interface {
	Validation
	// Msg marshals PatternPrefixConfigAutoFieldTest to protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	// and doesn't set defaults
	Msg() *openapi.PatternPrefixConfigAutoFieldTest
	// SetMsg unmarshals PatternPrefixConfigAutoFieldTest from protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	// and doesn't set defaults
	SetMsg(*openapi.PatternPrefixConfigAutoFieldTest) PatternPrefixConfigAutoFieldTest
	// ToProto marshals PatternPrefixConfigAutoFieldTest to protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	ToProto() (*openapi.PatternPrefixConfigAutoFieldTest, error)
	// ToPbText marshals PatternPrefixConfigAutoFieldTest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigAutoFieldTest to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigAutoFieldTest to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternPrefixConfigAutoFieldTest from protobuf object *openapi.PatternPrefixConfigAutoFieldTest
	FromProto(msg *openapi.PatternPrefixConfigAutoFieldTest) (PatternPrefixConfigAutoFieldTest, error)
	// FromPbText unmarshals PatternPrefixConfigAutoFieldTest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigAutoFieldTest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigAutoFieldTest from JSON text
	FromJson(value string) error
	// Validate validates PatternPrefixConfigAutoFieldTest
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigAutoFieldTest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternPrefixConfigAutoFieldTestChoiceEnum, set in PatternPrefixConfigAutoFieldTest
	Choice() PatternPrefixConfigAutoFieldTestChoiceEnum
	// SetChoice assigns PatternPrefixConfigAutoFieldTestChoiceEnum provided by user to PatternPrefixConfigAutoFieldTest
	SetChoice(value PatternPrefixConfigAutoFieldTestChoiceEnum) PatternPrefixConfigAutoFieldTest
	// HasChoice checks if Choice has been set in PatternPrefixConfigAutoFieldTest
	HasChoice() bool
	// Value returns uint32, set in PatternPrefixConfigAutoFieldTest.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternPrefixConfigAutoFieldTest
	SetValue(value uint32) PatternPrefixConfigAutoFieldTest
	// HasValue checks if Value has been set in PatternPrefixConfigAutoFieldTest
	HasValue() bool
	// Values returns []uint32, set in PatternPrefixConfigAutoFieldTest.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternPrefixConfigAutoFieldTest
	SetValues(value []uint32) PatternPrefixConfigAutoFieldTest
	// Auto returns uint32, set in PatternPrefixConfigAutoFieldTest.
	Auto() uint32
	// HasAuto checks if Auto has been set in PatternPrefixConfigAutoFieldTest
	HasAuto() bool
	// Increment returns PatternPrefixConfigAutoFieldTestCounter, set in PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	Increment() PatternPrefixConfigAutoFieldTestCounter
	// SetIncrement assigns PatternPrefixConfigAutoFieldTestCounter provided by user to PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	SetIncrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest
	// HasIncrement checks if Increment has been set in PatternPrefixConfigAutoFieldTest
	HasIncrement() bool
	// Decrement returns PatternPrefixConfigAutoFieldTestCounter, set in PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	Decrement() PatternPrefixConfigAutoFieldTestCounter
	// SetDecrement assigns PatternPrefixConfigAutoFieldTestCounter provided by user to PatternPrefixConfigAutoFieldTest.
	// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
	SetDecrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest
	// HasDecrement checks if Decrement has been set in PatternPrefixConfigAutoFieldTest
	HasDecrement() bool
	setNil()
}

type PatternPrefixConfigAutoFieldTestChoiceEnum string

// Enum of Choice on PatternPrefixConfigAutoFieldTest
var PatternPrefixConfigAutoFieldTestChoice = struct {
	VALUE     PatternPrefixConfigAutoFieldTestChoiceEnum
	VALUES    PatternPrefixConfigAutoFieldTestChoiceEnum
	AUTO      PatternPrefixConfigAutoFieldTestChoiceEnum
	INCREMENT PatternPrefixConfigAutoFieldTestChoiceEnum
	DECREMENT PatternPrefixConfigAutoFieldTestChoiceEnum
}{
	VALUE:     PatternPrefixConfigAutoFieldTestChoiceEnum("value"),
	VALUES:    PatternPrefixConfigAutoFieldTestChoiceEnum("values"),
	AUTO:      PatternPrefixConfigAutoFieldTestChoiceEnum("auto"),
	INCREMENT: PatternPrefixConfigAutoFieldTestChoiceEnum("increment"),
	DECREMENT: PatternPrefixConfigAutoFieldTestChoiceEnum("decrement"),
}

func (obj *patternPrefixConfigAutoFieldTest) Choice() PatternPrefixConfigAutoFieldTestChoiceEnum {
	return PatternPrefixConfigAutoFieldTestChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternPrefixConfigAutoFieldTest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternPrefixConfigAutoFieldTest) SetChoice(value PatternPrefixConfigAutoFieldTestChoiceEnum) PatternPrefixConfigAutoFieldTest {
	intValue, ok := openapi.PatternPrefixConfigAutoFieldTest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternPrefixConfigAutoFieldTestChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternPrefixConfigAutoFieldTest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternPrefixConfigAutoFieldTestChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.AUTO {
		defaultValue := uint32(0)
		obj.obj.Auto = &defaultValue
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.INCREMENT {
		obj.obj.Increment = NewPatternPrefixConfigAutoFieldTestCounter().Msg()
	}

	if value == PatternPrefixConfigAutoFieldTestChoice.DECREMENT {
		obj.obj.Decrement = NewPatternPrefixConfigAutoFieldTestCounter().Msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) Value() uint32 {

	if obj.obj.Value == nil {
		obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetValue(value uint32) PatternPrefixConfigAutoFieldTest {
	obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternPrefixConfigAutoFieldTest) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetValues(value []uint32) PatternPrefixConfigAutoFieldTest {
	obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) Auto() uint32 {

	if obj.obj.Auto == nil {
		obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a uint32
func (obj *patternPrefixConfigAutoFieldTest) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) Increment() PatternPrefixConfigAutoFieldTestCounter {
	if obj.obj.Increment == nil {
		obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternPrefixConfigAutoFieldTestCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternPrefixConfigAutoFieldTestCounter value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetIncrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest {
	obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.Msg()

	return obj
}

// description is TBD
// Decrement returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) Decrement() PatternPrefixConfigAutoFieldTestCounter {
	if obj.obj.Decrement == nil {
		obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternPrefixConfigAutoFieldTestCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternPrefixConfigAutoFieldTestCounter
func (obj *patternPrefixConfigAutoFieldTest) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternPrefixConfigAutoFieldTestCounter value in the PatternPrefixConfigAutoFieldTest object
func (obj *patternPrefixConfigAutoFieldTest) SetDecrement(value PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTest {
	obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.Msg()

	return obj
}

func (obj *patternPrefixConfigAutoFieldTest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternPrefixConfigAutoFieldTest.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Auto != nil {

		if *obj.obj.Auto > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTest.Auto <= 255 but Got %d", *obj.obj.Auto))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternPrefixConfigAutoFieldTest) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternPrefixConfigAutoFieldTestChoice.AUTO)

	}

}

// ***** WObject *****
type wObject struct {
	validation
	obj *openapi.WObject
}

func NewWObject() WObject {
	obj := wObject{obj: &openapi.WObject{}}
	obj.setDefault()
	return &obj
}

func (obj *wObject) Msg() *openapi.WObject {
	return obj.obj
}

func (obj *wObject) SetMsg(msg *openapi.WObject) WObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *wObject) ToProto() (*openapi.WObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *wObject) FromProto(msg *openapi.WObject) (WObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *wObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *wObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *wObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *wObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *wObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *wObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *wObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *wObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *wObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *wObject) Clone() (WObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewWObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// WObject is description is TBD
type WObject interface {
	Validation
	// Msg marshals WObject to protobuf object *openapi.WObject
	// and doesn't set defaults
	Msg() *openapi.WObject
	// SetMsg unmarshals WObject from protobuf object *openapi.WObject
	// and doesn't set defaults
	SetMsg(*openapi.WObject) WObject
	// ToProto marshals WObject to protobuf object *openapi.WObject
	ToProto() (*openapi.WObject, error)
	// ToPbText marshals WObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals WObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals WObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals WObject from protobuf object *openapi.WObject
	FromProto(msg *openapi.WObject) (WObject, error)
	// FromPbText unmarshals WObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals WObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals WObject from JSON text
	FromJson(value string) error
	// Validate validates WObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (WObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// WName returns string, set in WObject.
	WName() string
	// SetWName assigns string provided by user to WObject
	SetWName(value string) WObject
}

// description is TBD
// WName returns a string
func (obj *wObject) WName() string {

	return obj.obj.WName
}

// description is TBD
// SetWName sets the string value in the WObject object
func (obj *wObject) SetWName(value string) WObject {

	obj.obj.WName = value
	return obj
}

func (obj *wObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// WName is required
	if obj.obj.WName == "" {
		vObj.validationErrors = append(vObj.validationErrors, "WName is required field on interface WObject")
	}
}

func (obj *wObject) setDefault() {

}

// ***** ZObject *****
type zObject struct {
	validation
	obj *openapi.ZObject
}

func NewZObject() ZObject {
	obj := zObject{obj: &openapi.ZObject{}}
	obj.setDefault()
	return &obj
}

func (obj *zObject) Msg() *openapi.ZObject {
	return obj.obj
}

func (obj *zObject) SetMsg(msg *openapi.ZObject) ZObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *zObject) ToProto() (*openapi.ZObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *zObject) FromProto(msg *openapi.ZObject) (ZObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *zObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *zObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *zObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *zObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *zObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *zObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *zObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *zObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *zObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *zObject) Clone() (ZObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewZObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// ZObject is description is TBD
type ZObject interface {
	Validation
	// Msg marshals ZObject to protobuf object *openapi.ZObject
	// and doesn't set defaults
	Msg() *openapi.ZObject
	// SetMsg unmarshals ZObject from protobuf object *openapi.ZObject
	// and doesn't set defaults
	SetMsg(*openapi.ZObject) ZObject
	// ToProto marshals ZObject to protobuf object *openapi.ZObject
	ToProto() (*openapi.ZObject, error)
	// ToPbText marshals ZObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ZObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals ZObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ZObject from protobuf object *openapi.ZObject
	FromProto(msg *openapi.ZObject) (ZObject, error)
	// FromPbText unmarshals ZObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ZObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ZObject from JSON text
	FromJson(value string) error
	// Validate validates ZObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ZObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in ZObject.
	Name() string
	// SetName assigns string provided by user to ZObject
	SetName(value string) ZObject
}

// description is TBD
// Name returns a string
func (obj *zObject) Name() string {

	return obj.obj.Name
}

// description is TBD
// SetName sets the string value in the ZObject object
func (obj *zObject) SetName(value string) ZObject {

	obj.obj.Name = value
	return obj
}

func (obj *zObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface ZObject")
	}
}

func (obj *zObject) setDefault() {

}

// ***** YObject *****
type yObject struct {
	validation
	obj *openapi.YObject
}

func NewYObject() YObject {
	obj := yObject{obj: &openapi.YObject{}}
	obj.setDefault()
	return &obj
}

func (obj *yObject) Msg() *openapi.YObject {
	return obj.obj
}

func (obj *yObject) SetMsg(msg *openapi.YObject) YObject {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *yObject) ToProto() (*openapi.YObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *yObject) FromProto(msg *openapi.YObject) (YObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *yObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *yObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *yObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *yObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *yObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *yObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *yObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *yObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *yObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *yObject) Clone() (YObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewYObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// YObject is description is TBD
type YObject interface {
	Validation
	// Msg marshals YObject to protobuf object *openapi.YObject
	// and doesn't set defaults
	Msg() *openapi.YObject
	// SetMsg unmarshals YObject from protobuf object *openapi.YObject
	// and doesn't set defaults
	SetMsg(*openapi.YObject) YObject
	// ToProto marshals YObject to protobuf object *openapi.YObject
	ToProto() (*openapi.YObject, error)
	// ToPbText marshals YObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals YObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals YObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals YObject from protobuf object *openapi.YObject
	FromProto(msg *openapi.YObject) (YObject, error)
	// FromPbText unmarshals YObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals YObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals YObject from JSON text
	FromJson(value string) error
	// Validate validates YObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (YObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// YName returns string, set in YObject.
	YName() string
	// SetYName assigns string provided by user to YObject
	SetYName(value string) YObject
	// HasYName checks if YName has been set in YObject
	HasYName() bool
}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// YName returns a string
func (obj *yObject) YName() string {

	return *obj.obj.YName

}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// YName returns a string
func (obj *yObject) HasYName() bool {
	return obj.obj.YName != nil
}

// TBD
//
// x-constraint:
// - /components/schemas/ZObject/properties/name
// - /components/schemas/WObject/properties/w_name
//
// SetYName sets the string value in the YObject object
func (obj *yObject) SetYName(value string) YObject {

	obj.obj.YName = &value
	return obj
}

func (obj *yObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *yObject) setDefault() {

}

// ***** ChoiceObject *****
type choiceObject struct {
	validation
	obj        *openapi.ChoiceObject
	eObjHolder EObject
	fObjHolder FObject
}

func NewChoiceObject() ChoiceObject {
	obj := choiceObject{obj: &openapi.ChoiceObject{}}
	obj.setDefault()
	return &obj
}

func (obj *choiceObject) Msg() *openapi.ChoiceObject {
	return obj.obj
}

func (obj *choiceObject) SetMsg(msg *openapi.ChoiceObject) ChoiceObject {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *choiceObject) ToProto() (*openapi.ChoiceObject, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *choiceObject) FromProto(msg *openapi.ChoiceObject) (ChoiceObject, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *choiceObject) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *choiceObject) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *choiceObject) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *choiceObject) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *choiceObject) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *choiceObject) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *choiceObject) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *choiceObject) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *choiceObject) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *choiceObject) Clone() (ChoiceObject, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewChoiceObject()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *choiceObject) setNil() {
	obj.eObjHolder = nil
	obj.fObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ChoiceObject is description is TBD
type ChoiceObject interface {
	Validation
	// Msg marshals ChoiceObject to protobuf object *openapi.ChoiceObject
	// and doesn't set defaults
	Msg() *openapi.ChoiceObject
	// SetMsg unmarshals ChoiceObject from protobuf object *openapi.ChoiceObject
	// and doesn't set defaults
	SetMsg(*openapi.ChoiceObject) ChoiceObject
	// ToProto marshals ChoiceObject to protobuf object *openapi.ChoiceObject
	ToProto() (*openapi.ChoiceObject, error)
	// ToPbText marshals ChoiceObject to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ChoiceObject to YAML text
	ToYaml() (string, error)
	// ToJson marshals ChoiceObject to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ChoiceObject from protobuf object *openapi.ChoiceObject
	FromProto(msg *openapi.ChoiceObject) (ChoiceObject, error)
	// FromPbText unmarshals ChoiceObject from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ChoiceObject from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ChoiceObject from JSON text
	FromJson(value string) error
	// Validate validates ChoiceObject
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ChoiceObject, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ChoiceObjectChoiceEnum, set in ChoiceObject
	Choice() ChoiceObjectChoiceEnum
	// SetChoice assigns ChoiceObjectChoiceEnum provided by user to ChoiceObject
	SetChoice(value ChoiceObjectChoiceEnum) ChoiceObject
	// HasChoice checks if Choice has been set in ChoiceObject
	HasChoice() bool
	// EObj returns EObject, set in ChoiceObject.
	// EObject is description is TBD
	EObj() EObject
	// SetEObj assigns EObject provided by user to ChoiceObject.
	// EObject is description is TBD
	SetEObj(value EObject) ChoiceObject
	// HasEObj checks if EObj has been set in ChoiceObject
	HasEObj() bool
	// FObj returns FObject, set in ChoiceObject.
	// FObject is description is TBD
	FObj() FObject
	// SetFObj assigns FObject provided by user to ChoiceObject.
	// FObject is description is TBD
	SetFObj(value FObject) ChoiceObject
	// HasFObj checks if FObj has been set in ChoiceObject
	HasFObj() bool
	setNil()
}

type ChoiceObjectChoiceEnum string

// Enum of Choice on ChoiceObject
var ChoiceObjectChoice = struct {
	E_OBJ  ChoiceObjectChoiceEnum
	F_OBJ  ChoiceObjectChoiceEnum
	NO_OBJ ChoiceObjectChoiceEnum
}{
	E_OBJ:  ChoiceObjectChoiceEnum("e_obj"),
	F_OBJ:  ChoiceObjectChoiceEnum("f_obj"),
	NO_OBJ: ChoiceObjectChoiceEnum("no_obj"),
}

func (obj *choiceObject) Choice() ChoiceObjectChoiceEnum {
	return ChoiceObjectChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *choiceObject) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *choiceObject) SetChoice(value ChoiceObjectChoiceEnum) ChoiceObject {
	intValue, ok := openapi.ChoiceObject_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ChoiceObjectChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.ChoiceObject_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.FObj = nil
	obj.fObjHolder = nil
	obj.obj.EObj = nil
	obj.eObjHolder = nil

	if value == ChoiceObjectChoice.E_OBJ {
		obj.obj.EObj = NewEObject().Msg()
	}

	if value == ChoiceObjectChoice.F_OBJ {
		obj.obj.FObj = NewFObject().Msg()
	}

	return obj
}

// description is TBD
// EObj returns a EObject
func (obj *choiceObject) EObj() EObject {
	if obj.obj.EObj == nil {
		obj.SetChoice(ChoiceObjectChoice.E_OBJ)
	}
	if obj.eObjHolder == nil {
		obj.eObjHolder = &eObject{obj: obj.obj.EObj}
	}
	return obj.eObjHolder
}

// description is TBD
// EObj returns a EObject
func (obj *choiceObject) HasEObj() bool {
	return obj.obj.EObj != nil
}

// description is TBD
// SetEObj sets the EObject value in the ChoiceObject object
func (obj *choiceObject) SetEObj(value EObject) ChoiceObject {
	obj.SetChoice(ChoiceObjectChoice.E_OBJ)
	obj.eObjHolder = nil
	obj.obj.EObj = value.Msg()

	return obj
}

// description is TBD
// FObj returns a FObject
func (obj *choiceObject) FObj() FObject {
	if obj.obj.FObj == nil {
		obj.SetChoice(ChoiceObjectChoice.F_OBJ)
	}
	if obj.fObjHolder == nil {
		obj.fObjHolder = &fObject{obj: obj.obj.FObj}
	}
	return obj.fObjHolder
}

// description is TBD
// FObj returns a FObject
func (obj *choiceObject) HasFObj() bool {
	return obj.obj.FObj != nil
}

// description is TBD
// SetFObj sets the FObject value in the ChoiceObject object
func (obj *choiceObject) SetFObj(value FObject) ChoiceObject {
	obj.SetChoice(ChoiceObjectChoice.F_OBJ)
	obj.fObjHolder = nil
	obj.obj.FObj = value.Msg()

	return obj
}

func (obj *choiceObject) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.EObj != nil {

		obj.EObj().validateObj(vObj, set_default)
	}

	if obj.obj.FObj != nil {

		obj.FObj().validateObj(vObj, set_default)
	}

}

func (obj *choiceObject) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(ChoiceObjectChoice.NO_OBJ)

	}

}

// ***** RequiredChoiceParent *****
type requiredChoiceParent struct {
	validation
	obj                   *openapi.RequiredChoiceParent
	intermediateObjHolder RequiredChoiceIntermediate
}

func NewRequiredChoiceParent() RequiredChoiceParent {
	obj := requiredChoiceParent{obj: &openapi.RequiredChoiceParent{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceParent) Msg() *openapi.RequiredChoiceParent {
	return obj.obj
}

func (obj *requiredChoiceParent) SetMsg(msg *openapi.RequiredChoiceParent) RequiredChoiceParent {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *requiredChoiceParent) ToProto() (*openapi.RequiredChoiceParent, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *requiredChoiceParent) FromProto(msg *openapi.RequiredChoiceParent) (RequiredChoiceParent, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *requiredChoiceParent) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *requiredChoiceParent) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *requiredChoiceParent) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceParent) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *requiredChoiceParent) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceParent) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *requiredChoiceParent) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceParent) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceParent) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceParent) Clone() (RequiredChoiceParent, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceParent()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *requiredChoiceParent) setNil() {
	obj.intermediateObjHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RequiredChoiceParent is description is TBD
type RequiredChoiceParent interface {
	Validation
	// Msg marshals RequiredChoiceParent to protobuf object *openapi.RequiredChoiceParent
	// and doesn't set defaults
	Msg() *openapi.RequiredChoiceParent
	// SetMsg unmarshals RequiredChoiceParent from protobuf object *openapi.RequiredChoiceParent
	// and doesn't set defaults
	SetMsg(*openapi.RequiredChoiceParent) RequiredChoiceParent
	// ToProto marshals RequiredChoiceParent to protobuf object *openapi.RequiredChoiceParent
	ToProto() (*openapi.RequiredChoiceParent, error)
	// ToPbText marshals RequiredChoiceParent to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceParent to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceParent to JSON text
	ToJson() (string, error)
	// FromProto unmarshals RequiredChoiceParent from protobuf object *openapi.RequiredChoiceParent
	FromProto(msg *openapi.RequiredChoiceParent) (RequiredChoiceParent, error)
	// FromPbText unmarshals RequiredChoiceParent from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceParent from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceParent from JSON text
	FromJson(value string) error
	// Validate validates RequiredChoiceParent
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceParent, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns RequiredChoiceParentChoiceEnum, set in RequiredChoiceParent
	Choice() RequiredChoiceParentChoiceEnum
	// SetChoice assigns RequiredChoiceParentChoiceEnum provided by user to RequiredChoiceParent
	SetChoice(value RequiredChoiceParentChoiceEnum) RequiredChoiceParent
	// IntermediateObj returns RequiredChoiceIntermediate, set in RequiredChoiceParent.
	// RequiredChoiceIntermediate is description is TBD
	IntermediateObj() RequiredChoiceIntermediate
	// SetIntermediateObj assigns RequiredChoiceIntermediate provided by user to RequiredChoiceParent.
	// RequiredChoiceIntermediate is description is TBD
	SetIntermediateObj(value RequiredChoiceIntermediate) RequiredChoiceParent
	// HasIntermediateObj checks if IntermediateObj has been set in RequiredChoiceParent
	HasIntermediateObj() bool
	setNil()
}

type RequiredChoiceParentChoiceEnum string

// Enum of Choice on RequiredChoiceParent
var RequiredChoiceParentChoice = struct {
	INTERMEDIATE_OBJ RequiredChoiceParentChoiceEnum
	NO_OBJ           RequiredChoiceParentChoiceEnum
}{
	INTERMEDIATE_OBJ: RequiredChoiceParentChoiceEnum("intermediate_obj"),
	NO_OBJ:           RequiredChoiceParentChoiceEnum("no_obj"),
}

func (obj *requiredChoiceParent) Choice() RequiredChoiceParentChoiceEnum {
	return RequiredChoiceParentChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *requiredChoiceParent) SetChoice(value RequiredChoiceParentChoiceEnum) RequiredChoiceParent {
	intValue, ok := openapi.RequiredChoiceParent_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RequiredChoiceParentChoiceEnum", string(value)))
		return obj
	}
	obj.obj.Choice = openapi.RequiredChoiceParent_Choice_Enum(intValue)
	obj.obj.IntermediateObj = nil
	obj.intermediateObjHolder = nil

	if value == RequiredChoiceParentChoice.INTERMEDIATE_OBJ {
		obj.obj.IntermediateObj = NewRequiredChoiceIntermediate().Msg()
	}

	return obj
}

// description is TBD
// IntermediateObj returns a RequiredChoiceIntermediate
func (obj *requiredChoiceParent) IntermediateObj() RequiredChoiceIntermediate {
	if obj.obj.IntermediateObj == nil {
		obj.SetChoice(RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	}
	if obj.intermediateObjHolder == nil {
		obj.intermediateObjHolder = &requiredChoiceIntermediate{obj: obj.obj.IntermediateObj}
	}
	return obj.intermediateObjHolder
}

// description is TBD
// IntermediateObj returns a RequiredChoiceIntermediate
func (obj *requiredChoiceParent) HasIntermediateObj() bool {
	return obj.obj.IntermediateObj != nil
}

// description is TBD
// SetIntermediateObj sets the RequiredChoiceIntermediate value in the RequiredChoiceParent object
func (obj *requiredChoiceParent) SetIntermediateObj(value RequiredChoiceIntermediate) RequiredChoiceParent {
	obj.SetChoice(RequiredChoiceParentChoice.INTERMEDIATE_OBJ)
	obj.intermediateObjHolder = nil
	obj.obj.IntermediateObj = value.Msg()

	return obj
}

func (obj *requiredChoiceParent) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice.Number() == 0 {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface RequiredChoiceParent")
	}

	if obj.obj.IntermediateObj != nil {

		obj.IntermediateObj().validateObj(vObj, set_default)
	}

}

func (obj *requiredChoiceParent) setDefault() {

}

// ***** Error *****
type _error struct {
	validation
	obj *openapi.Error
}

func NewError() Error {
	obj := _error{obj: &openapi.Error{}}
	obj.setDefault()
	return &obj
}

func (obj *_error) Msg() *openapi.Error {
	return obj.obj
}

func (obj *_error) SetMsg(msg *openapi.Error) Error {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *_error) ToProto() (*openapi.Error, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *_error) FromProto(msg *openapi.Error) (Error, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *_error) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *_error) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *_error) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *_error) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *_error) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *_error) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *_error) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *_error) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *_error) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *_error) Clone() (Error, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewError()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Error is error response generated while serving API request.
type Error interface {
	Validation
	// Msg marshals Error to protobuf object *openapi.Error
	// and doesn't set defaults
	Msg() *openapi.Error
	// SetMsg unmarshals Error from protobuf object *openapi.Error
	// and doesn't set defaults
	SetMsg(*openapi.Error) Error
	// ToProto marshals Error to protobuf object *openapi.Error
	ToProto() (*openapi.Error, error)
	// ToPbText marshals Error to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Error to YAML text
	ToYaml() (string, error)
	// ToJson marshals Error to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Error from protobuf object *openapi.Error
	FromProto(msg *openapi.Error) (Error, error)
	// FromPbText unmarshals Error from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Error from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Error from JSON text
	FromJson(value string) error
	// Validate validates Error
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Error, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Code returns int32, set in Error.
	Code() int32
	// SetCode assigns int32 provided by user to Error
	SetCode(value int32) Error
	// Kind returns ErrorKindEnum, set in Error
	Kind() ErrorKindEnum
	// SetKind assigns ErrorKindEnum provided by user to Error
	SetKind(value ErrorKindEnum) Error
	// HasKind checks if Kind has been set in Error
	HasKind() bool
	// Errors returns []string, set in Error.
	Errors() []string
	// SetErrors assigns []string provided by user to Error
	SetErrors(value []string) Error
	// implement Error function for implementingnative Error Interface.
	Error() string
}

func (obj *_error) Error() string {
	json, err := obj.ToJson()
	if err != nil {
		return fmt.Sprintf("could not convert Error to JSON: %v", err)
	}
	return json
}

// Numeric status code based on underlying transport being used.
// Code returns a int32
func (obj *_error) Code() int32 {

	return obj.obj.Code
}

// Numeric status code based on underlying transport being used.
// SetCode sets the int32 value in the Error object
func (obj *_error) SetCode(value int32) Error {

	obj.obj.Code = value
	return obj
}

type ErrorKindEnum string

// Enum of Kind on Error
var ErrorKind = struct {
	TRANSPORT  ErrorKindEnum
	VALIDATION ErrorKindEnum
	INTERNAL   ErrorKindEnum
}{
	TRANSPORT:  ErrorKindEnum("transport"),
	VALIDATION: ErrorKindEnum("validation"),
	INTERNAL:   ErrorKindEnum("internal"),
}

func (obj *_error) Kind() ErrorKindEnum {
	return ErrorKindEnum(obj.obj.Kind.Enum().String())
}

// Kind of error message.
// Kind returns a string
func (obj *_error) HasKind() bool {
	return obj.obj.Kind != nil
}

func (obj *_error) SetKind(value ErrorKindEnum) Error {
	intValue, ok := openapi.Error_Kind_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ErrorKindEnum", string(value)))
		return obj
	}
	enumValue := openapi.Error_Kind_Enum(intValue)
	obj.obj.Kind = &enumValue

	return obj
}

// List of error messages generated while serving API request.
// Errors returns a []string
func (obj *_error) Errors() []string {
	if obj.obj.Errors == nil {
		obj.obj.Errors = make([]string, 0)
	}
	return obj.obj.Errors
}

// List of error messages generated while serving API request.
// SetErrors sets the []string value in the Error object
func (obj *_error) SetErrors(value []string) Error {

	if obj.obj.Errors == nil {
		obj.obj.Errors = make([]string, 0)
	}
	obj.obj.Errors = value

	return obj
}

func (obj *_error) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Errors is required
	if obj.obj.Errors == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Errors is required field on interface Error")
	}
}

func (obj *_error) setDefault() {

}

// ***** Metrics *****
type metrics struct {
	validation
	obj         *openapi.Metrics
	portsHolder MetricsPortMetricIter
	flowsHolder MetricsFlowMetricIter
}

func NewMetrics() Metrics {
	obj := metrics{obj: &openapi.Metrics{}}
	obj.setDefault()
	return &obj
}

func (obj *metrics) Msg() *openapi.Metrics {
	return obj.obj
}

func (obj *metrics) SetMsg(msg *openapi.Metrics) Metrics {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *metrics) ToProto() (*openapi.Metrics, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *metrics) FromProto(msg *openapi.Metrics) (Metrics, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *metrics) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *metrics) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *metrics) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *metrics) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *metrics) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *metrics) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *metrics) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *metrics) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *metrics) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *metrics) Clone() (Metrics, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewMetrics()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *metrics) setNil() {
	obj.portsHolder = nil
	obj.flowsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Metrics is description is TBD
type Metrics interface {
	Validation
	// Msg marshals Metrics to protobuf object *openapi.Metrics
	// and doesn't set defaults
	Msg() *openapi.Metrics
	// SetMsg unmarshals Metrics from protobuf object *openapi.Metrics
	// and doesn't set defaults
	SetMsg(*openapi.Metrics) Metrics
	// ToProto marshals Metrics to protobuf object *openapi.Metrics
	ToProto() (*openapi.Metrics, error)
	// ToPbText marshals Metrics to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Metrics to YAML text
	ToYaml() (string, error)
	// ToJson marshals Metrics to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Metrics from protobuf object *openapi.Metrics
	FromProto(msg *openapi.Metrics) (Metrics, error)
	// FromPbText unmarshals Metrics from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Metrics from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Metrics from JSON text
	FromJson(value string) error
	// Validate validates Metrics
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Metrics, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns MetricsChoiceEnum, set in Metrics
	Choice() MetricsChoiceEnum
	// SetChoice assigns MetricsChoiceEnum provided by user to Metrics
	SetChoice(value MetricsChoiceEnum) Metrics
	// HasChoice checks if Choice has been set in Metrics
	HasChoice() bool
	// Ports returns MetricsPortMetricIterIter, set in Metrics
	Ports() MetricsPortMetricIter
	// Flows returns MetricsFlowMetricIterIter, set in Metrics
	Flows() MetricsFlowMetricIter
	setNil()
}

type MetricsChoiceEnum string

// Enum of Choice on Metrics
var MetricsChoice = struct {
	PORTS MetricsChoiceEnum
	FLOWS MetricsChoiceEnum
}{
	PORTS: MetricsChoiceEnum("ports"),
	FLOWS: MetricsChoiceEnum("flows"),
}

func (obj *metrics) Choice() MetricsChoiceEnum {
	return MetricsChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *metrics) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *metrics) SetChoice(value MetricsChoiceEnum) Metrics {
	intValue, ok := openapi.Metrics_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on MetricsChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.Metrics_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Flows = nil
	obj.flowsHolder = nil
	obj.obj.Ports = nil
	obj.portsHolder = nil

	if value == MetricsChoice.PORTS {
		obj.obj.Ports = []*openapi.PortMetric{}
	}

	if value == MetricsChoice.FLOWS {
		obj.obj.Flows = []*openapi.FlowMetric{}
	}

	return obj
}

// description is TBD
// Ports returns a []PortMetric
func (obj *metrics) Ports() MetricsPortMetricIter {
	if len(obj.obj.Ports) == 0 {
		obj.SetChoice(MetricsChoice.PORTS)
	}
	if obj.portsHolder == nil {
		obj.portsHolder = newMetricsPortMetricIter(&obj.obj.Ports).setMsg(obj)
	}
	return obj.portsHolder
}

type metricsPortMetricIter struct {
	obj             *metrics
	portMetricSlice []PortMetric
	fieldPtr        *[]*openapi.PortMetric
}

func newMetricsPortMetricIter(ptr *[]*openapi.PortMetric) MetricsPortMetricIter {
	return &metricsPortMetricIter{fieldPtr: ptr}
}

type MetricsPortMetricIter interface {
	setMsg(*metrics) MetricsPortMetricIter
	Items() []PortMetric
	Add() PortMetric
	Append(items ...PortMetric) MetricsPortMetricIter
	Set(index int, newObj PortMetric) MetricsPortMetricIter
	Clear() MetricsPortMetricIter
	clearHolderSlice() MetricsPortMetricIter
	appendHolderSlice(item PortMetric) MetricsPortMetricIter
}

func (obj *metricsPortMetricIter) setMsg(msg *metrics) MetricsPortMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&portMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsPortMetricIter) Items() []PortMetric {
	return obj.portMetricSlice
}

func (obj *metricsPortMetricIter) Add() PortMetric {
	newObj := &openapi.PortMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &portMetric{obj: newObj}
	newLibObj.setDefault()
	obj.portMetricSlice = append(obj.portMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsPortMetricIter) Append(items ...PortMetric) MetricsPortMetricIter {
	for _, item := range items {
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.portMetricSlice = append(obj.portMetricSlice, item)
	}
	return obj
}

func (obj *metricsPortMetricIter) Set(index int, newObj PortMetric) MetricsPortMetricIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
	obj.portMetricSlice[index] = newObj
	return obj
}
func (obj *metricsPortMetricIter) Clear() MetricsPortMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.PortMetric{}
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsPortMetricIter) clearHolderSlice() MetricsPortMetricIter {
	if len(obj.portMetricSlice) > 0 {
		obj.portMetricSlice = []PortMetric{}
	}
	return obj
}
func (obj *metricsPortMetricIter) appendHolderSlice(item PortMetric) MetricsPortMetricIter {
	obj.portMetricSlice = append(obj.portMetricSlice, item)
	return obj
}

// description is TBD
// Flows returns a []FlowMetric
func (obj *metrics) Flows() MetricsFlowMetricIter {
	if len(obj.obj.Flows) == 0 {
		obj.SetChoice(MetricsChoice.FLOWS)
	}
	if obj.flowsHolder == nil {
		obj.flowsHolder = newMetricsFlowMetricIter(&obj.obj.Flows).setMsg(obj)
	}
	return obj.flowsHolder
}

type metricsFlowMetricIter struct {
	obj             *metrics
	flowMetricSlice []FlowMetric
	fieldPtr        *[]*openapi.FlowMetric
}

func newMetricsFlowMetricIter(ptr *[]*openapi.FlowMetric) MetricsFlowMetricIter {
	return &metricsFlowMetricIter{fieldPtr: ptr}
}

type MetricsFlowMetricIter interface {
	setMsg(*metrics) MetricsFlowMetricIter
	Items() []FlowMetric
	Add() FlowMetric
	Append(items ...FlowMetric) MetricsFlowMetricIter
	Set(index int, newObj FlowMetric) MetricsFlowMetricIter
	Clear() MetricsFlowMetricIter
	clearHolderSlice() MetricsFlowMetricIter
	appendHolderSlice(item FlowMetric) MetricsFlowMetricIter
}

func (obj *metricsFlowMetricIter) setMsg(msg *metrics) MetricsFlowMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *metricsFlowMetricIter) Items() []FlowMetric {
	return obj.flowMetricSlice
}

func (obj *metricsFlowMetricIter) Add() FlowMetric {
	newObj := &openapi.FlowMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowMetric{obj: newObj}
	newLibObj.setDefault()
	obj.flowMetricSlice = append(obj.flowMetricSlice, newLibObj)
	return newLibObj
}

func (obj *metricsFlowMetricIter) Append(items ...FlowMetric) MetricsFlowMetricIter {
	for _, item := range items {
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	}
	return obj
}

func (obj *metricsFlowMetricIter) Set(index int, newObj FlowMetric) MetricsFlowMetricIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
	obj.flowMetricSlice[index] = newObj
	return obj
}
func (obj *metricsFlowMetricIter) Clear() MetricsFlowMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.FlowMetric{}
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsFlowMetricIter) clearHolderSlice() MetricsFlowMetricIter {
	if len(obj.flowMetricSlice) > 0 {
		obj.flowMetricSlice = []FlowMetric{}
	}
	return obj
}
func (obj *metricsFlowMetricIter) appendHolderSlice(item FlowMetric) MetricsFlowMetricIter {
	obj.flowMetricSlice = append(obj.flowMetricSlice, item)
	return obj
}

func (obj *metrics) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Ports) != 0 {

		if set_default {
			obj.Ports().clearHolderSlice()
			for _, item := range obj.obj.Ports {
				obj.Ports().appendHolderSlice(&portMetric{obj: item})
			}
		}
		for _, item := range obj.Ports().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Flows) != 0 {

		if set_default {
			obj.Flows().clearHolderSlice()
			for _, item := range obj.obj.Flows {
				obj.Flows().appendHolderSlice(&flowMetric{obj: item})
			}
		}
		for _, item := range obj.Flows().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *metrics) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(MetricsChoice.PORTS)

	}

}

// ***** WarningDetails *****
type warningDetails struct {
	validation
	obj *openapi.WarningDetails
}

func NewWarningDetails() WarningDetails {
	obj := warningDetails{obj: &openapi.WarningDetails{}}
	obj.setDefault()
	return &obj
}

func (obj *warningDetails) Msg() *openapi.WarningDetails {
	return obj.obj
}

func (obj *warningDetails) SetMsg(msg *openapi.WarningDetails) WarningDetails {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *warningDetails) ToProto() (*openapi.WarningDetails, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *warningDetails) FromProto(msg *openapi.WarningDetails) (WarningDetails, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *warningDetails) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *warningDetails) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *warningDetails) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *warningDetails) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *warningDetails) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *warningDetails) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *warningDetails) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *warningDetails) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *warningDetails) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *warningDetails) Clone() (WarningDetails, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewWarningDetails()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// WarningDetails is description is TBD
type WarningDetails interface {
	Validation
	// Msg marshals WarningDetails to protobuf object *openapi.WarningDetails
	// and doesn't set defaults
	Msg() *openapi.WarningDetails
	// SetMsg unmarshals WarningDetails from protobuf object *openapi.WarningDetails
	// and doesn't set defaults
	SetMsg(*openapi.WarningDetails) WarningDetails
	// ToProto marshals WarningDetails to protobuf object *openapi.WarningDetails
	ToProto() (*openapi.WarningDetails, error)
	// ToPbText marshals WarningDetails to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals WarningDetails to YAML text
	ToYaml() (string, error)
	// ToJson marshals WarningDetails to JSON text
	ToJson() (string, error)
	// FromProto unmarshals WarningDetails from protobuf object *openapi.WarningDetails
	FromProto(msg *openapi.WarningDetails) (WarningDetails, error)
	// FromPbText unmarshals WarningDetails from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals WarningDetails from YAML text
	FromYaml(value string) error
	// FromJson unmarshals WarningDetails from JSON text
	FromJson(value string) error
	// Validate validates WarningDetails
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (WarningDetails, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Warnings returns []string, set in WarningDetails.
	Warnings() []string
	// SetWarnings assigns []string provided by user to WarningDetails
	SetWarnings(value []string) WarningDetails
}

// description is TBD
// Warnings returns a []string
func (obj *warningDetails) Warnings() []string {
	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	return obj.obj.Warnings
}

// description is TBD
// SetWarnings sets the []string value in the WarningDetails object
func (obj *warningDetails) SetWarnings(value []string) WarningDetails {

	if obj.obj.Warnings == nil {
		obj.obj.Warnings = make([]string, 0)
	}
	obj.obj.Warnings = value

	return obj
}

func (obj *warningDetails) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *warningDetails) setDefault() {

}

// ***** CommonResponseSuccess *****
type commonResponseSuccess struct {
	validation
	obj *openapi.CommonResponseSuccess
}

func NewCommonResponseSuccess() CommonResponseSuccess {
	obj := commonResponseSuccess{obj: &openapi.CommonResponseSuccess{}}
	obj.setDefault()
	return &obj
}

func (obj *commonResponseSuccess) Msg() *openapi.CommonResponseSuccess {
	return obj.obj
}

func (obj *commonResponseSuccess) SetMsg(msg *openapi.CommonResponseSuccess) CommonResponseSuccess {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *commonResponseSuccess) ToProto() (*openapi.CommonResponseSuccess, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *commonResponseSuccess) FromProto(msg *openapi.CommonResponseSuccess) (CommonResponseSuccess, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *commonResponseSuccess) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *commonResponseSuccess) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *commonResponseSuccess) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *commonResponseSuccess) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *commonResponseSuccess) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *commonResponseSuccess) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *commonResponseSuccess) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *commonResponseSuccess) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *commonResponseSuccess) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *commonResponseSuccess) Clone() (CommonResponseSuccess, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewCommonResponseSuccess()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// CommonResponseSuccess is description is TBD
type CommonResponseSuccess interface {
	Validation
	// Msg marshals CommonResponseSuccess to protobuf object *openapi.CommonResponseSuccess
	// and doesn't set defaults
	Msg() *openapi.CommonResponseSuccess
	// SetMsg unmarshals CommonResponseSuccess from protobuf object *openapi.CommonResponseSuccess
	// and doesn't set defaults
	SetMsg(*openapi.CommonResponseSuccess) CommonResponseSuccess
	// ToProto marshals CommonResponseSuccess to protobuf object *openapi.CommonResponseSuccess
	ToProto() (*openapi.CommonResponseSuccess, error)
	// ToPbText marshals CommonResponseSuccess to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals CommonResponseSuccess to YAML text
	ToYaml() (string, error)
	// ToJson marshals CommonResponseSuccess to JSON text
	ToJson() (string, error)
	// FromProto unmarshals CommonResponseSuccess from protobuf object *openapi.CommonResponseSuccess
	FromProto(msg *openapi.CommonResponseSuccess) (CommonResponseSuccess, error)
	// FromPbText unmarshals CommonResponseSuccess from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals CommonResponseSuccess from YAML text
	FromYaml(value string) error
	// FromJson unmarshals CommonResponseSuccess from JSON text
	FromJson(value string) error
	// Validate validates CommonResponseSuccess
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (CommonResponseSuccess, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Message returns string, set in CommonResponseSuccess.
	Message() string
	// SetMessage assigns string provided by user to CommonResponseSuccess
	SetMessage(value string) CommonResponseSuccess
	// HasMessage checks if Message has been set in CommonResponseSuccess
	HasMessage() bool
}

// description is TBD
// Message returns a string
func (obj *commonResponseSuccess) Message() string {

	return *obj.obj.Message

}

// description is TBD
// Message returns a string
func (obj *commonResponseSuccess) HasMessage() bool {
	return obj.obj.Message != nil
}

// description is TBD
// SetMessage sets the string value in the CommonResponseSuccess object
func (obj *commonResponseSuccess) SetMessage(value string) CommonResponseSuccess {

	obj.obj.Message = &value
	return obj
}

func (obj *commonResponseSuccess) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *commonResponseSuccess) setDefault() {

}

// ***** ServiceAbcItemList *****
type serviceAbcItemList struct {
	validation
	obj         *openapi.ServiceAbcItemList
	itemsHolder ServiceAbcItemListServiceAbcItemIter
}

func NewServiceAbcItemList() ServiceAbcItemList {
	obj := serviceAbcItemList{obj: &openapi.ServiceAbcItemList{}}
	obj.setDefault()
	return &obj
}

func (obj *serviceAbcItemList) Msg() *openapi.ServiceAbcItemList {
	return obj.obj
}

func (obj *serviceAbcItemList) SetMsg(msg *openapi.ServiceAbcItemList) ServiceAbcItemList {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *serviceAbcItemList) ToProto() (*openapi.ServiceAbcItemList, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *serviceAbcItemList) FromProto(msg *openapi.ServiceAbcItemList) (ServiceAbcItemList, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *serviceAbcItemList) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *serviceAbcItemList) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *serviceAbcItemList) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *serviceAbcItemList) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *serviceAbcItemList) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *serviceAbcItemList) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *serviceAbcItemList) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *serviceAbcItemList) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *serviceAbcItemList) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *serviceAbcItemList) Clone() (ServiceAbcItemList, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewServiceAbcItemList()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *serviceAbcItemList) setNil() {
	obj.itemsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ServiceAbcItemList is description is TBD
type ServiceAbcItemList interface {
	Validation
	// Msg marshals ServiceAbcItemList to protobuf object *openapi.ServiceAbcItemList
	// and doesn't set defaults
	Msg() *openapi.ServiceAbcItemList
	// SetMsg unmarshals ServiceAbcItemList from protobuf object *openapi.ServiceAbcItemList
	// and doesn't set defaults
	SetMsg(*openapi.ServiceAbcItemList) ServiceAbcItemList
	// ToProto marshals ServiceAbcItemList to protobuf object *openapi.ServiceAbcItemList
	ToProto() (*openapi.ServiceAbcItemList, error)
	// ToPbText marshals ServiceAbcItemList to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ServiceAbcItemList to YAML text
	ToYaml() (string, error)
	// ToJson marshals ServiceAbcItemList to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ServiceAbcItemList from protobuf object *openapi.ServiceAbcItemList
	FromProto(msg *openapi.ServiceAbcItemList) (ServiceAbcItemList, error)
	// FromPbText unmarshals ServiceAbcItemList from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ServiceAbcItemList from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ServiceAbcItemList from JSON text
	FromJson(value string) error
	// Validate validates ServiceAbcItemList
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ServiceAbcItemList, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Items returns ServiceAbcItemListServiceAbcItemIterIter, set in ServiceAbcItemList
	Items() ServiceAbcItemListServiceAbcItemIter
	setNil()
}

// description is TBD
// Items returns a []ServiceAbcItem
func (obj *serviceAbcItemList) Items() ServiceAbcItemListServiceAbcItemIter {
	if len(obj.obj.Items) == 0 {
		obj.obj.Items = []*openapi.ServiceAbcItem{}
	}
	if obj.itemsHolder == nil {
		obj.itemsHolder = newServiceAbcItemListServiceAbcItemIter(&obj.obj.Items).setMsg(obj)
	}
	return obj.itemsHolder
}

type serviceAbcItemListServiceAbcItemIter struct {
	obj                 *serviceAbcItemList
	serviceAbcItemSlice []ServiceAbcItem
	fieldPtr            *[]*openapi.ServiceAbcItem
}

func newServiceAbcItemListServiceAbcItemIter(ptr *[]*openapi.ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	return &serviceAbcItemListServiceAbcItemIter{fieldPtr: ptr}
}

type ServiceAbcItemListServiceAbcItemIter interface {
	setMsg(*serviceAbcItemList) ServiceAbcItemListServiceAbcItemIter
	Items() []ServiceAbcItem
	Add() ServiceAbcItem
	Append(items ...ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
	Set(index int, newObj ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
	Clear() ServiceAbcItemListServiceAbcItemIter
	clearHolderSlice() ServiceAbcItemListServiceAbcItemIter
	appendHolderSlice(item ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter
}

func (obj *serviceAbcItemListServiceAbcItemIter) setMsg(msg *serviceAbcItemList) ServiceAbcItemListServiceAbcItemIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&serviceAbcItem{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Items() []ServiceAbcItem {
	return obj.serviceAbcItemSlice
}

func (obj *serviceAbcItemListServiceAbcItemIter) Add() ServiceAbcItem {
	newObj := &openapi.ServiceAbcItem{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &serviceAbcItem{obj: newObj}
	newLibObj.setDefault()
	obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, newLibObj)
	return newLibObj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Append(items ...ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	for _, item := range items {
		newObj := item.Msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, item)
	}
	return obj
}

func (obj *serviceAbcItemListServiceAbcItemIter) Set(index int, newObj ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	(*obj.fieldPtr)[index] = newObj.Msg()
	obj.serviceAbcItemSlice[index] = newObj
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) Clear() ServiceAbcItemListServiceAbcItemIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*openapi.ServiceAbcItem{}
		obj.serviceAbcItemSlice = []ServiceAbcItem{}
	}
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) clearHolderSlice() ServiceAbcItemListServiceAbcItemIter {
	if len(obj.serviceAbcItemSlice) > 0 {
		obj.serviceAbcItemSlice = []ServiceAbcItem{}
	}
	return obj
}
func (obj *serviceAbcItemListServiceAbcItemIter) appendHolderSlice(item ServiceAbcItem) ServiceAbcItemListServiceAbcItemIter {
	obj.serviceAbcItemSlice = append(obj.serviceAbcItemSlice, item)
	return obj
}

func (obj *serviceAbcItemList) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Items) != 0 {

		if set_default {
			obj.Items().clearHolderSlice()
			for _, item := range obj.obj.Items {
				obj.Items().appendHolderSlice(&serviceAbcItem{obj: item})
			}
		}
		for _, item := range obj.Items().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *serviceAbcItemList) setDefault() {

}

// ***** ServiceAbcItem *****
type serviceAbcItem struct {
	validation
	obj *openapi.ServiceAbcItem
}

func NewServiceAbcItem() ServiceAbcItem {
	obj := serviceAbcItem{obj: &openapi.ServiceAbcItem{}}
	obj.setDefault()
	return &obj
}

func (obj *serviceAbcItem) Msg() *openapi.ServiceAbcItem {
	return obj.obj
}

func (obj *serviceAbcItem) SetMsg(msg *openapi.ServiceAbcItem) ServiceAbcItem {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *serviceAbcItem) ToProto() (*openapi.ServiceAbcItem, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *serviceAbcItem) FromProto(msg *openapi.ServiceAbcItem) (ServiceAbcItem, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *serviceAbcItem) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *serviceAbcItem) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *serviceAbcItem) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *serviceAbcItem) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *serviceAbcItem) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *serviceAbcItem) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *serviceAbcItem) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *serviceAbcItem) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *serviceAbcItem) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *serviceAbcItem) Clone() (ServiceAbcItem, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewServiceAbcItem()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// ServiceAbcItem is description is TBD
type ServiceAbcItem interface {
	Validation
	// Msg marshals ServiceAbcItem to protobuf object *openapi.ServiceAbcItem
	// and doesn't set defaults
	Msg() *openapi.ServiceAbcItem
	// SetMsg unmarshals ServiceAbcItem from protobuf object *openapi.ServiceAbcItem
	// and doesn't set defaults
	SetMsg(*openapi.ServiceAbcItem) ServiceAbcItem
	// ToProto marshals ServiceAbcItem to protobuf object *openapi.ServiceAbcItem
	ToProto() (*openapi.ServiceAbcItem, error)
	// ToPbText marshals ServiceAbcItem to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ServiceAbcItem to YAML text
	ToYaml() (string, error)
	// ToJson marshals ServiceAbcItem to JSON text
	ToJson() (string, error)
	// FromProto unmarshals ServiceAbcItem from protobuf object *openapi.ServiceAbcItem
	FromProto(msg *openapi.ServiceAbcItem) (ServiceAbcItem, error)
	// FromPbText unmarshals ServiceAbcItem from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ServiceAbcItem from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ServiceAbcItem from JSON text
	FromJson(value string) error
	// Validate validates ServiceAbcItem
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ServiceAbcItem, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SomeId returns string, set in ServiceAbcItem.
	SomeId() string
	// SetSomeId assigns string provided by user to ServiceAbcItem
	SetSomeId(value string) ServiceAbcItem
	// HasSomeId checks if SomeId has been set in ServiceAbcItem
	HasSomeId() bool
	// SomeString returns string, set in ServiceAbcItem.
	SomeString() string
	// SetSomeString assigns string provided by user to ServiceAbcItem
	SetSomeString(value string) ServiceAbcItem
	// HasSomeString checks if SomeString has been set in ServiceAbcItem
	HasSomeString() bool
	// PathId returns string, set in ServiceAbcItem.
	PathId() string
	// SetPathId assigns string provided by user to ServiceAbcItem
	SetPathId(value string) ServiceAbcItem
	// HasPathId checks if PathId has been set in ServiceAbcItem
	HasPathId() bool
	// Level2 returns string, set in ServiceAbcItem.
	Level2() string
	// SetLevel2 assigns string provided by user to ServiceAbcItem
	SetLevel2(value string) ServiceAbcItem
	// HasLevel2 checks if Level2 has been set in ServiceAbcItem
	HasLevel2() bool
}

// description is TBD
// SomeId returns a string
func (obj *serviceAbcItem) SomeId() string {

	return *obj.obj.SomeId

}

// description is TBD
// SomeId returns a string
func (obj *serviceAbcItem) HasSomeId() bool {
	return obj.obj.SomeId != nil
}

// description is TBD
// SetSomeId sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetSomeId(value string) ServiceAbcItem {

	obj.obj.SomeId = &value
	return obj
}

// description is TBD
// SomeString returns a string
func (obj *serviceAbcItem) SomeString() string {

	return *obj.obj.SomeString

}

// description is TBD
// SomeString returns a string
func (obj *serviceAbcItem) HasSomeString() bool {
	return obj.obj.SomeString != nil
}

// description is TBD
// SetSomeString sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetSomeString(value string) ServiceAbcItem {

	obj.obj.SomeString = &value
	return obj
}

// description is TBD
// PathId returns a string
func (obj *serviceAbcItem) PathId() string {

	return *obj.obj.PathId

}

// description is TBD
// PathId returns a string
func (obj *serviceAbcItem) HasPathId() bool {
	return obj.obj.PathId != nil
}

// description is TBD
// SetPathId sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetPathId(value string) ServiceAbcItem {

	obj.obj.PathId = &value
	return obj
}

// description is TBD
// Level2 returns a string
func (obj *serviceAbcItem) Level2() string {

	return *obj.obj.Level_2

}

// description is TBD
// Level2 returns a string
func (obj *serviceAbcItem) HasLevel2() bool {
	return obj.obj.Level_2 != nil
}

// description is TBD
// SetLevel2 sets the string value in the ServiceAbcItem object
func (obj *serviceAbcItem) SetLevel2(value string) ServiceAbcItem {

	obj.obj.Level_2 = &value
	return obj
}

func (obj *serviceAbcItem) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *serviceAbcItem) setDefault() {

}

// ***** Version *****
type version struct {
	validation
	obj *openapi.Version
}

func NewVersion() Version {
	obj := version{obj: &openapi.Version{}}
	obj.setDefault()
	return &obj
}

func (obj *version) Msg() *openapi.Version {
	return obj.obj
}

func (obj *version) SetMsg(msg *openapi.Version) Version {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *version) ToProto() (*openapi.Version, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *version) FromProto(msg *openapi.Version) (Version, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *version) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *version) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *version) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *version) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *version) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *version) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *version) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *version) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *version) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *version) Clone() (Version, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVersion()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// Version is version details
type Version interface {
	Validation
	// Msg marshals Version to protobuf object *openapi.Version
	// and doesn't set defaults
	Msg() *openapi.Version
	// SetMsg unmarshals Version from protobuf object *openapi.Version
	// and doesn't set defaults
	SetMsg(*openapi.Version) Version
	// ToProto marshals Version to protobuf object *openapi.Version
	ToProto() (*openapi.Version, error)
	// ToPbText marshals Version to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Version to YAML text
	ToYaml() (string, error)
	// ToJson marshals Version to JSON text
	ToJson() (string, error)
	// FromProto unmarshals Version from protobuf object *openapi.Version
	FromProto(msg *openapi.Version) (Version, error)
	// FromPbText unmarshals Version from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Version from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Version from JSON text
	FromJson(value string) error
	// Validate validates Version
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Version, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ApiSpecVersion returns string, set in Version.
	ApiSpecVersion() string
	// SetApiSpecVersion assigns string provided by user to Version
	SetApiSpecVersion(value string) Version
	// HasApiSpecVersion checks if ApiSpecVersion has been set in Version
	HasApiSpecVersion() bool
	// SdkVersion returns string, set in Version.
	SdkVersion() string
	// SetSdkVersion assigns string provided by user to Version
	SetSdkVersion(value string) Version
	// HasSdkVersion checks if SdkVersion has been set in Version
	HasSdkVersion() bool
	// AppVersion returns string, set in Version.
	AppVersion() string
	// SetAppVersion assigns string provided by user to Version
	SetAppVersion(value string) Version
	// HasAppVersion checks if AppVersion has been set in Version
	HasAppVersion() bool
}

// Version of API specification
// ApiSpecVersion returns a string
func (obj *version) ApiSpecVersion() string {

	return *obj.obj.ApiSpecVersion

}

// Version of API specification
// ApiSpecVersion returns a string
func (obj *version) HasApiSpecVersion() bool {
	return obj.obj.ApiSpecVersion != nil
}

// Version of API specification
// SetApiSpecVersion sets the string value in the Version object
func (obj *version) SetApiSpecVersion(value string) Version {

	obj.obj.ApiSpecVersion = &value
	return obj
}

// Version of SDK generated from API specification
// SdkVersion returns a string
func (obj *version) SdkVersion() string {

	return *obj.obj.SdkVersion

}

// Version of SDK generated from API specification
// SdkVersion returns a string
func (obj *version) HasSdkVersion() bool {
	return obj.obj.SdkVersion != nil
}

// Version of SDK generated from API specification
// SetSdkVersion sets the string value in the Version object
func (obj *version) SetSdkVersion(value string) Version {

	obj.obj.SdkVersion = &value
	return obj
}

// Version of application consuming or serving the API
// AppVersion returns a string
func (obj *version) AppVersion() string {

	return *obj.obj.AppVersion

}

// Version of application consuming or serving the API
// AppVersion returns a string
func (obj *version) HasAppVersion() bool {
	return obj.obj.AppVersion != nil
}

// Version of application consuming or serving the API
// SetAppVersion sets the string value in the Version object
func (obj *version) SetAppVersion(value string) Version {

	obj.obj.AppVersion = &value
	return obj
}

func (obj *version) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *version) setDefault() {
	if obj.obj.ApiSpecVersion == nil {
		obj.SetApiSpecVersion("")
	}
	if obj.obj.SdkVersion == nil {
		obj.SetSdkVersion("")
	}
	if obj.obj.AppVersion == nil {
		obj.SetAppVersion("")
	}

}

// ***** LevelTwo *****
type levelTwo struct {
	validation
	obj        *openapi.LevelTwo
	l2P1Holder LevelThree
}

func NewLevelTwo() LevelTwo {
	obj := levelTwo{obj: &openapi.LevelTwo{}}
	obj.setDefault()
	return &obj
}

func (obj *levelTwo) Msg() *openapi.LevelTwo {
	return obj.obj
}

func (obj *levelTwo) SetMsg(msg *openapi.LevelTwo) LevelTwo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *levelTwo) ToProto() (*openapi.LevelTwo, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *levelTwo) FromProto(msg *openapi.LevelTwo) (LevelTwo, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *levelTwo) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *levelTwo) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *levelTwo) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelTwo) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *levelTwo) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelTwo) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *levelTwo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelTwo) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelTwo) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelTwo) Clone() (LevelTwo, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelTwo()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *levelTwo) setNil() {
	obj.l2P1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelTwo is test Level 2
type LevelTwo interface {
	Validation
	// Msg marshals LevelTwo to protobuf object *openapi.LevelTwo
	// and doesn't set defaults
	Msg() *openapi.LevelTwo
	// SetMsg unmarshals LevelTwo from protobuf object *openapi.LevelTwo
	// and doesn't set defaults
	SetMsg(*openapi.LevelTwo) LevelTwo
	// ToProto marshals LevelTwo to protobuf object *openapi.LevelTwo
	ToProto() (*openapi.LevelTwo, error)
	// ToPbText marshals LevelTwo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelTwo to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelTwo to JSON text
	ToJson() (string, error)
	// FromProto unmarshals LevelTwo from protobuf object *openapi.LevelTwo
	FromProto(msg *openapi.LevelTwo) (LevelTwo, error)
	// FromPbText unmarshals LevelTwo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelTwo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelTwo from JSON text
	FromJson(value string) error
	// Validate validates LevelTwo
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelTwo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L2P1 returns LevelThree, set in LevelTwo.
	// LevelThree is test Level3
	L2P1() LevelThree
	// SetL2P1 assigns LevelThree provided by user to LevelTwo.
	// LevelThree is test Level3
	SetL2P1(value LevelThree) LevelTwo
	// HasL2P1 checks if L2P1 has been set in LevelTwo
	HasL2P1() bool
	setNil()
}

// Level Two
// L2P1 returns a LevelThree
func (obj *levelTwo) L2P1() LevelThree {
	if obj.obj.L2P1 == nil {
		obj.obj.L2P1 = NewLevelThree().Msg()
	}
	if obj.l2P1Holder == nil {
		obj.l2P1Holder = &levelThree{obj: obj.obj.L2P1}
	}
	return obj.l2P1Holder
}

// Level Two
// L2P1 returns a LevelThree
func (obj *levelTwo) HasL2P1() bool {
	return obj.obj.L2P1 != nil
}

// Level Two
// SetL2P1 sets the LevelThree value in the LevelTwo object
func (obj *levelTwo) SetL2P1(value LevelThree) LevelTwo {

	obj.l2P1Holder = nil
	obj.obj.L2P1 = value.Msg()

	return obj
}

func (obj *levelTwo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L2P1 != nil {

		obj.L2P1().validateObj(vObj, set_default)
	}

}

func (obj *levelTwo) setDefault() {

}

// ***** LevelFour *****
type levelFour struct {
	validation
	obj        *openapi.LevelFour
	l4P1Holder LevelOne
}

func NewLevelFour() LevelFour {
	obj := levelFour{obj: &openapi.LevelFour{}}
	obj.setDefault()
	return &obj
}

func (obj *levelFour) Msg() *openapi.LevelFour {
	return obj.obj
}

func (obj *levelFour) SetMsg(msg *openapi.LevelFour) LevelFour {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *levelFour) ToProto() (*openapi.LevelFour, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *levelFour) FromProto(msg *openapi.LevelFour) (LevelFour, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *levelFour) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *levelFour) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *levelFour) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelFour) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *levelFour) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelFour) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *levelFour) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelFour) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelFour) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelFour) Clone() (LevelFour, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelFour()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *levelFour) setNil() {
	obj.l4P1Holder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// LevelFour is test level4 redundant junk testing
type LevelFour interface {
	Validation
	// Msg marshals LevelFour to protobuf object *openapi.LevelFour
	// and doesn't set defaults
	Msg() *openapi.LevelFour
	// SetMsg unmarshals LevelFour from protobuf object *openapi.LevelFour
	// and doesn't set defaults
	SetMsg(*openapi.LevelFour) LevelFour
	// ToProto marshals LevelFour to protobuf object *openapi.LevelFour
	ToProto() (*openapi.LevelFour, error)
	// ToPbText marshals LevelFour to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelFour to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelFour to JSON text
	ToJson() (string, error)
	// FromProto unmarshals LevelFour from protobuf object *openapi.LevelFour
	FromProto(msg *openapi.LevelFour) (LevelFour, error)
	// FromPbText unmarshals LevelFour from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelFour from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelFour from JSON text
	FromJson(value string) error
	// Validate validates LevelFour
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelFour, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L4P1 returns LevelOne, set in LevelFour.
	// LevelOne is to Test Multi level non-primitive types
	L4P1() LevelOne
	// SetL4P1 assigns LevelOne provided by user to LevelFour.
	// LevelOne is to Test Multi level non-primitive types
	SetL4P1(value LevelOne) LevelFour
	// HasL4P1 checks if L4P1 has been set in LevelFour
	HasL4P1() bool
	setNil()
}

// loop over level 1
// L4P1 returns a LevelOne
func (obj *levelFour) L4P1() LevelOne {
	if obj.obj.L4P1 == nil {
		obj.obj.L4P1 = NewLevelOne().Msg()
	}
	if obj.l4P1Holder == nil {
		obj.l4P1Holder = &levelOne{obj: obj.obj.L4P1}
	}
	return obj.l4P1Holder
}

// loop over level 1
// L4P1 returns a LevelOne
func (obj *levelFour) HasL4P1() bool {
	return obj.obj.L4P1 != nil
}

// loop over level 1
// SetL4P1 sets the LevelOne value in the LevelFour object
func (obj *levelFour) SetL4P1(value LevelOne) LevelFour {

	obj.l4P1Holder = nil
	obj.obj.L4P1 = value.Msg()

	return obj
}

func (obj *levelFour) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.L4P1 != nil {

		obj.L4P1().validateObj(vObj, set_default)
	}

}

func (obj *levelFour) setDefault() {

}

// ***** PatternIpv4PatternIpv4 *****
type patternIpv4PatternIpv4 struct {
	validation
	obj             *openapi.PatternIpv4PatternIpv4
	incrementHolder PatternIpv4PatternIpv4Counter
	decrementHolder PatternIpv4PatternIpv4Counter
}

func NewPatternIpv4PatternIpv4() PatternIpv4PatternIpv4 {
	obj := patternIpv4PatternIpv4{obj: &openapi.PatternIpv4PatternIpv4{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternIpv4) Msg() *openapi.PatternIpv4PatternIpv4 {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4) SetMsg(msg *openapi.PatternIpv4PatternIpv4) PatternIpv4PatternIpv4 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIpv4PatternIpv4) ToProto() (*openapi.PatternIpv4PatternIpv4, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIpv4PatternIpv4) FromProto(msg *openapi.PatternIpv4PatternIpv4) (PatternIpv4PatternIpv4, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIpv4PatternIpv4) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIpv4PatternIpv4) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIpv4PatternIpv4) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv4PatternIpv4) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIpv4PatternIpv4) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv4PatternIpv4) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv4PatternIpv4) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternIpv4) Clone() (PatternIpv4PatternIpv4, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternIpv4()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *patternIpv4PatternIpv4) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv4PatternIpv4 is tBD
type PatternIpv4PatternIpv4 interface {
	Validation
	// Msg marshals PatternIpv4PatternIpv4 to protobuf object *openapi.PatternIpv4PatternIpv4
	// and doesn't set defaults
	Msg() *openapi.PatternIpv4PatternIpv4
	// SetMsg unmarshals PatternIpv4PatternIpv4 from protobuf object *openapi.PatternIpv4PatternIpv4
	// and doesn't set defaults
	SetMsg(*openapi.PatternIpv4PatternIpv4) PatternIpv4PatternIpv4
	// ToProto marshals PatternIpv4PatternIpv4 to protobuf object *openapi.PatternIpv4PatternIpv4
	ToProto() (*openapi.PatternIpv4PatternIpv4, error)
	// ToPbText marshals PatternIpv4PatternIpv4 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternIpv4 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternIpv4 to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIpv4PatternIpv4 from protobuf object *openapi.PatternIpv4PatternIpv4
	FromProto(msg *openapi.PatternIpv4PatternIpv4) (PatternIpv4PatternIpv4, error)
	// FromPbText unmarshals PatternIpv4PatternIpv4 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternIpv4 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternIpv4 from JSON text
	FromJson(value string) error
	// Validate validates PatternIpv4PatternIpv4
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternIpv4, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv4PatternIpv4ChoiceEnum, set in PatternIpv4PatternIpv4
	Choice() PatternIpv4PatternIpv4ChoiceEnum
	// SetChoice assigns PatternIpv4PatternIpv4ChoiceEnum provided by user to PatternIpv4PatternIpv4
	SetChoice(value PatternIpv4PatternIpv4ChoiceEnum) PatternIpv4PatternIpv4
	// HasChoice checks if Choice has been set in PatternIpv4PatternIpv4
	HasChoice() bool
	// Value returns string, set in PatternIpv4PatternIpv4.
	Value() string
	// SetValue assigns string provided by user to PatternIpv4PatternIpv4
	SetValue(value string) PatternIpv4PatternIpv4
	// HasValue checks if Value has been set in PatternIpv4PatternIpv4
	HasValue() bool
	// Values returns []string, set in PatternIpv4PatternIpv4.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv4PatternIpv4
	SetValues(value []string) PatternIpv4PatternIpv4
	// Increment returns PatternIpv4PatternIpv4Counter, set in PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	Increment() PatternIpv4PatternIpv4Counter
	// SetIncrement assigns PatternIpv4PatternIpv4Counter provided by user to PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	SetIncrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4
	// HasIncrement checks if Increment has been set in PatternIpv4PatternIpv4
	HasIncrement() bool
	// Decrement returns PatternIpv4PatternIpv4Counter, set in PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	Decrement() PatternIpv4PatternIpv4Counter
	// SetDecrement assigns PatternIpv4PatternIpv4Counter provided by user to PatternIpv4PatternIpv4.
	// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
	SetDecrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4
	// HasDecrement checks if Decrement has been set in PatternIpv4PatternIpv4
	HasDecrement() bool
	setNil()
}

type PatternIpv4PatternIpv4ChoiceEnum string

// Enum of Choice on PatternIpv4PatternIpv4
var PatternIpv4PatternIpv4Choice = struct {
	VALUE     PatternIpv4PatternIpv4ChoiceEnum
	VALUES    PatternIpv4PatternIpv4ChoiceEnum
	INCREMENT PatternIpv4PatternIpv4ChoiceEnum
	DECREMENT PatternIpv4PatternIpv4ChoiceEnum
}{
	VALUE:     PatternIpv4PatternIpv4ChoiceEnum("value"),
	VALUES:    PatternIpv4PatternIpv4ChoiceEnum("values"),
	INCREMENT: PatternIpv4PatternIpv4ChoiceEnum("increment"),
	DECREMENT: PatternIpv4PatternIpv4ChoiceEnum("decrement"),
}

func (obj *patternIpv4PatternIpv4) Choice() PatternIpv4PatternIpv4ChoiceEnum {
	return PatternIpv4PatternIpv4ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv4PatternIpv4) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv4PatternIpv4) SetChoice(value PatternIpv4PatternIpv4ChoiceEnum) PatternIpv4PatternIpv4 {
	intValue, ok := openapi.PatternIpv4PatternIpv4_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv4PatternIpv4ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv4PatternIpv4_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv4PatternIpv4Choice.VALUE {
		defaultValue := "0.0.0.0"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv4PatternIpv4Choice.VALUES {
		defaultValue := []string{"0.0.0.0"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv4PatternIpv4Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv4PatternIpv4Counter().Msg()
	}

	if value == PatternIpv4PatternIpv4Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv4PatternIpv4Counter().Msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternIpv4) Value() string {

	if obj.obj.Value == nil {
		obj.SetChoice(PatternIpv4PatternIpv4Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv4PatternIpv4) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetValue(value string) PatternIpv4PatternIpv4 {
	obj.SetChoice(PatternIpv4PatternIpv4Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv4PatternIpv4) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"0.0.0.0"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetValues(value []string) PatternIpv4PatternIpv4 {
	obj.SetChoice(PatternIpv4PatternIpv4Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) Increment() PatternIpv4PatternIpv4Counter {
	if obj.obj.Increment == nil {
		obj.SetChoice(PatternIpv4PatternIpv4Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv4PatternIpv4Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv4PatternIpv4Counter value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetIncrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4 {
	obj.SetChoice(PatternIpv4PatternIpv4Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.Msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) Decrement() PatternIpv4PatternIpv4Counter {
	if obj.obj.Decrement == nil {
		obj.SetChoice(PatternIpv4PatternIpv4Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv4PatternIpv4Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv4PatternIpv4Counter
func (obj *patternIpv4PatternIpv4) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv4PatternIpv4Counter value in the PatternIpv4PatternIpv4 object
func (obj *patternIpv4PatternIpv4) SetDecrement(value PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4 {
	obj.SetChoice(PatternIpv4PatternIpv4Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.Msg()

	return obj
}

func (obj *patternIpv4PatternIpv4) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv4(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv4Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv4PatternIpv4) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternIpv4PatternIpv4Choice.VALUE)

	}

}

// ***** PatternIpv6PatternIpv6 *****
type patternIpv6PatternIpv6 struct {
	validation
	obj             *openapi.PatternIpv6PatternIpv6
	incrementHolder PatternIpv6PatternIpv6Counter
	decrementHolder PatternIpv6PatternIpv6Counter
}

func NewPatternIpv6PatternIpv6() PatternIpv6PatternIpv6 {
	obj := patternIpv6PatternIpv6{obj: &openapi.PatternIpv6PatternIpv6{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv6PatternIpv6) Msg() *openapi.PatternIpv6PatternIpv6 {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6) SetMsg(msg *openapi.PatternIpv6PatternIpv6) PatternIpv6PatternIpv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIpv6PatternIpv6) ToProto() (*openapi.PatternIpv6PatternIpv6, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIpv6PatternIpv6) FromProto(msg *openapi.PatternIpv6PatternIpv6) (PatternIpv6PatternIpv6, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIpv6PatternIpv6) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIpv6PatternIpv6) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIpv6PatternIpv6) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv6PatternIpv6) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIpv6PatternIpv6) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv6PatternIpv6) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv6PatternIpv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv6PatternIpv6) Clone() (PatternIpv6PatternIpv6, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv6PatternIpv6()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *patternIpv6PatternIpv6) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIpv6PatternIpv6 is tBD
type PatternIpv6PatternIpv6 interface {
	Validation
	// Msg marshals PatternIpv6PatternIpv6 to protobuf object *openapi.PatternIpv6PatternIpv6
	// and doesn't set defaults
	Msg() *openapi.PatternIpv6PatternIpv6
	// SetMsg unmarshals PatternIpv6PatternIpv6 from protobuf object *openapi.PatternIpv6PatternIpv6
	// and doesn't set defaults
	SetMsg(*openapi.PatternIpv6PatternIpv6) PatternIpv6PatternIpv6
	// ToProto marshals PatternIpv6PatternIpv6 to protobuf object *openapi.PatternIpv6PatternIpv6
	ToProto() (*openapi.PatternIpv6PatternIpv6, error)
	// ToPbText marshals PatternIpv6PatternIpv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv6PatternIpv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv6PatternIpv6 to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIpv6PatternIpv6 from protobuf object *openapi.PatternIpv6PatternIpv6
	FromProto(msg *openapi.PatternIpv6PatternIpv6) (PatternIpv6PatternIpv6, error)
	// FromPbText unmarshals PatternIpv6PatternIpv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv6PatternIpv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv6PatternIpv6 from JSON text
	FromJson(value string) error
	// Validate validates PatternIpv6PatternIpv6
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv6PatternIpv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIpv6PatternIpv6ChoiceEnum, set in PatternIpv6PatternIpv6
	Choice() PatternIpv6PatternIpv6ChoiceEnum
	// SetChoice assigns PatternIpv6PatternIpv6ChoiceEnum provided by user to PatternIpv6PatternIpv6
	SetChoice(value PatternIpv6PatternIpv6ChoiceEnum) PatternIpv6PatternIpv6
	// HasChoice checks if Choice has been set in PatternIpv6PatternIpv6
	HasChoice() bool
	// Value returns string, set in PatternIpv6PatternIpv6.
	Value() string
	// SetValue assigns string provided by user to PatternIpv6PatternIpv6
	SetValue(value string) PatternIpv6PatternIpv6
	// HasValue checks if Value has been set in PatternIpv6PatternIpv6
	HasValue() bool
	// Values returns []string, set in PatternIpv6PatternIpv6.
	Values() []string
	// SetValues assigns []string provided by user to PatternIpv6PatternIpv6
	SetValues(value []string) PatternIpv6PatternIpv6
	// Increment returns PatternIpv6PatternIpv6Counter, set in PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	Increment() PatternIpv6PatternIpv6Counter
	// SetIncrement assigns PatternIpv6PatternIpv6Counter provided by user to PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	SetIncrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6
	// HasIncrement checks if Increment has been set in PatternIpv6PatternIpv6
	HasIncrement() bool
	// Decrement returns PatternIpv6PatternIpv6Counter, set in PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	Decrement() PatternIpv6PatternIpv6Counter
	// SetDecrement assigns PatternIpv6PatternIpv6Counter provided by user to PatternIpv6PatternIpv6.
	// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
	SetDecrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6
	// HasDecrement checks if Decrement has been set in PatternIpv6PatternIpv6
	HasDecrement() bool
	setNil()
}

type PatternIpv6PatternIpv6ChoiceEnum string

// Enum of Choice on PatternIpv6PatternIpv6
var PatternIpv6PatternIpv6Choice = struct {
	VALUE     PatternIpv6PatternIpv6ChoiceEnum
	VALUES    PatternIpv6PatternIpv6ChoiceEnum
	INCREMENT PatternIpv6PatternIpv6ChoiceEnum
	DECREMENT PatternIpv6PatternIpv6ChoiceEnum
}{
	VALUE:     PatternIpv6PatternIpv6ChoiceEnum("value"),
	VALUES:    PatternIpv6PatternIpv6ChoiceEnum("values"),
	INCREMENT: PatternIpv6PatternIpv6ChoiceEnum("increment"),
	DECREMENT: PatternIpv6PatternIpv6ChoiceEnum("decrement"),
}

func (obj *patternIpv6PatternIpv6) Choice() PatternIpv6PatternIpv6ChoiceEnum {
	return PatternIpv6PatternIpv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIpv6PatternIpv6) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIpv6PatternIpv6) SetChoice(value PatternIpv6PatternIpv6ChoiceEnum) PatternIpv6PatternIpv6 {
	intValue, ok := openapi.PatternIpv6PatternIpv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIpv6PatternIpv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIpv6PatternIpv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIpv6PatternIpv6Choice.VALUE {
		defaultValue := "::"
		obj.obj.Value = &defaultValue
	}

	if value == PatternIpv6PatternIpv6Choice.VALUES {
		defaultValue := []string{"::"}
		obj.obj.Values = defaultValue
	}

	if value == PatternIpv6PatternIpv6Choice.INCREMENT {
		obj.obj.Increment = NewPatternIpv6PatternIpv6Counter().Msg()
	}

	if value == PatternIpv6PatternIpv6Choice.DECREMENT {
		obj.obj.Decrement = NewPatternIpv6PatternIpv6Counter().Msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternIpv6) Value() string {

	if obj.obj.Value == nil {
		obj.SetChoice(PatternIpv6PatternIpv6Choice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternIpv6PatternIpv6) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetValue(value string) PatternIpv6PatternIpv6 {
	obj.SetChoice(PatternIpv6PatternIpv6Choice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternIpv6PatternIpv6) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"::"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetValues(value []string) PatternIpv6PatternIpv6 {
	obj.SetChoice(PatternIpv6PatternIpv6Choice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) Increment() PatternIpv6PatternIpv6Counter {
	if obj.obj.Increment == nil {
		obj.SetChoice(PatternIpv6PatternIpv6Choice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIpv6PatternIpv6Counter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIpv6PatternIpv6Counter value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetIncrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6 {
	obj.SetChoice(PatternIpv6PatternIpv6Choice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.Msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) Decrement() PatternIpv6PatternIpv6Counter {
	if obj.obj.Decrement == nil {
		obj.SetChoice(PatternIpv6PatternIpv6Choice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIpv6PatternIpv6Counter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIpv6PatternIpv6Counter
func (obj *patternIpv6PatternIpv6) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIpv6PatternIpv6Counter value in the PatternIpv6PatternIpv6 object
func (obj *patternIpv6PatternIpv6) SetDecrement(value PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6 {
	obj.SetChoice(PatternIpv6PatternIpv6Choice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.Msg()

	return obj
}

func (obj *patternIpv6PatternIpv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateIpv6(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateIpv6Slice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6.Values"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIpv6PatternIpv6) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternIpv6PatternIpv6Choice.VALUE)

	}

}

// ***** PatternMacPatternMac *****
type patternMacPatternMac struct {
	validation
	obj             *openapi.PatternMacPatternMac
	incrementHolder PatternMacPatternMacCounter
	decrementHolder PatternMacPatternMacCounter
}

func NewPatternMacPatternMac() PatternMacPatternMac {
	obj := patternMacPatternMac{obj: &openapi.PatternMacPatternMac{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternMac) Msg() *openapi.PatternMacPatternMac {
	return obj.obj
}

func (obj *patternMacPatternMac) SetMsg(msg *openapi.PatternMacPatternMac) PatternMacPatternMac {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternMacPatternMac) ToProto() (*openapi.PatternMacPatternMac, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternMacPatternMac) FromProto(msg *openapi.PatternMacPatternMac) (PatternMacPatternMac, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternMacPatternMac) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternMacPatternMac) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternMacPatternMac) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternMacPatternMac) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternMacPatternMac) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternMacPatternMac) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternMacPatternMac) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternMac) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternMac) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternMac) Clone() (PatternMacPatternMac, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternMac()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *patternMacPatternMac) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternMacPatternMac is tBD
type PatternMacPatternMac interface {
	Validation
	// Msg marshals PatternMacPatternMac to protobuf object *openapi.PatternMacPatternMac
	// and doesn't set defaults
	Msg() *openapi.PatternMacPatternMac
	// SetMsg unmarshals PatternMacPatternMac from protobuf object *openapi.PatternMacPatternMac
	// and doesn't set defaults
	SetMsg(*openapi.PatternMacPatternMac) PatternMacPatternMac
	// ToProto marshals PatternMacPatternMac to protobuf object *openapi.PatternMacPatternMac
	ToProto() (*openapi.PatternMacPatternMac, error)
	// ToPbText marshals PatternMacPatternMac to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternMac to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternMac to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternMacPatternMac from protobuf object *openapi.PatternMacPatternMac
	FromProto(msg *openapi.PatternMacPatternMac) (PatternMacPatternMac, error)
	// FromPbText unmarshals PatternMacPatternMac from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternMac from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternMac from JSON text
	FromJson(value string) error
	// Validate validates PatternMacPatternMac
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternMac, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternMacPatternMacChoiceEnum, set in PatternMacPatternMac
	Choice() PatternMacPatternMacChoiceEnum
	// SetChoice assigns PatternMacPatternMacChoiceEnum provided by user to PatternMacPatternMac
	SetChoice(value PatternMacPatternMacChoiceEnum) PatternMacPatternMac
	// HasChoice checks if Choice has been set in PatternMacPatternMac
	HasChoice() bool
	// Value returns string, set in PatternMacPatternMac.
	Value() string
	// SetValue assigns string provided by user to PatternMacPatternMac
	SetValue(value string) PatternMacPatternMac
	// HasValue checks if Value has been set in PatternMacPatternMac
	HasValue() bool
	// Values returns []string, set in PatternMacPatternMac.
	Values() []string
	// SetValues assigns []string provided by user to PatternMacPatternMac
	SetValues(value []string) PatternMacPatternMac
	// Auto returns string, set in PatternMacPatternMac.
	Auto() string
	// HasAuto checks if Auto has been set in PatternMacPatternMac
	HasAuto() bool
	// Increment returns PatternMacPatternMacCounter, set in PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	Increment() PatternMacPatternMacCounter
	// SetIncrement assigns PatternMacPatternMacCounter provided by user to PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	SetIncrement(value PatternMacPatternMacCounter) PatternMacPatternMac
	// HasIncrement checks if Increment has been set in PatternMacPatternMac
	HasIncrement() bool
	// Decrement returns PatternMacPatternMacCounter, set in PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	Decrement() PatternMacPatternMacCounter
	// SetDecrement assigns PatternMacPatternMacCounter provided by user to PatternMacPatternMac.
	// PatternMacPatternMacCounter is mac counter pattern
	SetDecrement(value PatternMacPatternMacCounter) PatternMacPatternMac
	// HasDecrement checks if Decrement has been set in PatternMacPatternMac
	HasDecrement() bool
	setNil()
}

type PatternMacPatternMacChoiceEnum string

// Enum of Choice on PatternMacPatternMac
var PatternMacPatternMacChoice = struct {
	VALUE     PatternMacPatternMacChoiceEnum
	VALUES    PatternMacPatternMacChoiceEnum
	AUTO      PatternMacPatternMacChoiceEnum
	INCREMENT PatternMacPatternMacChoiceEnum
	DECREMENT PatternMacPatternMacChoiceEnum
}{
	VALUE:     PatternMacPatternMacChoiceEnum("value"),
	VALUES:    PatternMacPatternMacChoiceEnum("values"),
	AUTO:      PatternMacPatternMacChoiceEnum("auto"),
	INCREMENT: PatternMacPatternMacChoiceEnum("increment"),
	DECREMENT: PatternMacPatternMacChoiceEnum("decrement"),
}

func (obj *patternMacPatternMac) Choice() PatternMacPatternMacChoiceEnum {
	return PatternMacPatternMacChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternMacPatternMac) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternMacPatternMac) SetChoice(value PatternMacPatternMacChoiceEnum) PatternMacPatternMac {
	intValue, ok := openapi.PatternMacPatternMac_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternMacPatternMacChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternMacPatternMac_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Auto = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternMacPatternMacChoice.VALUE {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Value = &defaultValue
	}

	if value == PatternMacPatternMacChoice.VALUES {
		defaultValue := []string{"00:00:00:00:00:00"}
		obj.obj.Values = defaultValue
	}

	if value == PatternMacPatternMacChoice.AUTO {
		defaultValue := "00:00:00:00:00:00"
		obj.obj.Auto = &defaultValue
	}

	if value == PatternMacPatternMacChoice.INCREMENT {
		obj.obj.Increment = NewPatternMacPatternMacCounter().Msg()
	}

	if value == PatternMacPatternMacChoice.DECREMENT {
		obj.obj.Decrement = NewPatternMacPatternMacCounter().Msg()
	}

	return obj
}

// description is TBD
// Value returns a string
func (obj *patternMacPatternMac) Value() string {

	if obj.obj.Value == nil {
		obj.SetChoice(PatternMacPatternMacChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a string
func (obj *patternMacPatternMac) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the string value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetValue(value string) PatternMacPatternMac {
	obj.SetChoice(PatternMacPatternMacChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []string
func (obj *patternMacPatternMac) Values() []string {
	if obj.obj.Values == nil {
		obj.SetValues([]string{"00:00:00:00:00:00"})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []string value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetValues(value []string) PatternMacPatternMac {
	obj.SetChoice(PatternMacPatternMacChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]string, 0)
	}
	obj.obj.Values = value

	return obj
}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternMacPatternMac) Auto() string {

	if obj.obj.Auto == nil {
		obj.SetChoice(PatternMacPatternMacChoice.AUTO)
	}

	return *obj.obj.Auto

}

// The OTG implementation can provide a system generated
// value for this property. If the OTG is unable to generate a value
// the default value must be used.
// Auto returns a string
func (obj *patternMacPatternMac) HasAuto() bool {
	return obj.obj.Auto != nil
}

// description is TBD
// Increment returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) Increment() PatternMacPatternMacCounter {
	if obj.obj.Increment == nil {
		obj.SetChoice(PatternMacPatternMacChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternMacPatternMacCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternMacPatternMacCounter value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetIncrement(value PatternMacPatternMacCounter) PatternMacPatternMac {
	obj.SetChoice(PatternMacPatternMacChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.Msg()

	return obj
}

// description is TBD
// Decrement returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) Decrement() PatternMacPatternMacCounter {
	if obj.obj.Decrement == nil {
		obj.SetChoice(PatternMacPatternMacChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternMacPatternMacCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternMacPatternMacCounter
func (obj *patternMacPatternMac) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternMacPatternMacCounter value in the PatternMacPatternMac object
func (obj *patternMacPatternMac) SetDecrement(value PatternMacPatternMacCounter) PatternMacPatternMac {
	obj.SetChoice(PatternMacPatternMacChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.Msg()

	return obj
}

func (obj *patternMacPatternMac) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		err := obj.validateMac(obj.Value())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Value"))
		}

	}

	if obj.obj.Values != nil {

		err := obj.validateMacSlice(obj.Values())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Values"))
		}

	}

	if obj.obj.Auto != nil {

		err := obj.validateMac(obj.Auto())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMac.Auto"))
		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternMacPatternMac) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternMacPatternMacChoice.AUTO)

	}

}

// ***** PatternIntegerPatternInteger *****
type patternIntegerPatternInteger struct {
	validation
	obj             *openapi.PatternIntegerPatternInteger
	incrementHolder PatternIntegerPatternIntegerCounter
	decrementHolder PatternIntegerPatternIntegerCounter
}

func NewPatternIntegerPatternInteger() PatternIntegerPatternInteger {
	obj := patternIntegerPatternInteger{obj: &openapi.PatternIntegerPatternInteger{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternInteger) Msg() *openapi.PatternIntegerPatternInteger {
	return obj.obj
}

func (obj *patternIntegerPatternInteger) SetMsg(msg *openapi.PatternIntegerPatternInteger) PatternIntegerPatternInteger {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIntegerPatternInteger) ToProto() (*openapi.PatternIntegerPatternInteger, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIntegerPatternInteger) FromProto(msg *openapi.PatternIntegerPatternInteger) (PatternIntegerPatternInteger, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIntegerPatternInteger) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIntegerPatternInteger) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIntegerPatternInteger) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIntegerPatternInteger) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIntegerPatternInteger) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIntegerPatternInteger) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIntegerPatternInteger) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternInteger) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternInteger) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternInteger) Clone() (PatternIntegerPatternInteger, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternInteger()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *patternIntegerPatternInteger) setNil() {
	obj.incrementHolder = nil
	obj.decrementHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// PatternIntegerPatternInteger is tBD
type PatternIntegerPatternInteger interface {
	Validation
	// Msg marshals PatternIntegerPatternInteger to protobuf object *openapi.PatternIntegerPatternInteger
	// and doesn't set defaults
	Msg() *openapi.PatternIntegerPatternInteger
	// SetMsg unmarshals PatternIntegerPatternInteger from protobuf object *openapi.PatternIntegerPatternInteger
	// and doesn't set defaults
	SetMsg(*openapi.PatternIntegerPatternInteger) PatternIntegerPatternInteger
	// ToProto marshals PatternIntegerPatternInteger to protobuf object *openapi.PatternIntegerPatternInteger
	ToProto() (*openapi.PatternIntegerPatternInteger, error)
	// ToPbText marshals PatternIntegerPatternInteger to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternInteger to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternInteger to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIntegerPatternInteger from protobuf object *openapi.PatternIntegerPatternInteger
	FromProto(msg *openapi.PatternIntegerPatternInteger) (PatternIntegerPatternInteger, error)
	// FromPbText unmarshals PatternIntegerPatternInteger from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternInteger from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternInteger from JSON text
	FromJson(value string) error
	// Validate validates PatternIntegerPatternInteger
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternInteger, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternIntegerPatternIntegerChoiceEnum, set in PatternIntegerPatternInteger
	Choice() PatternIntegerPatternIntegerChoiceEnum
	// SetChoice assigns PatternIntegerPatternIntegerChoiceEnum provided by user to PatternIntegerPatternInteger
	SetChoice(value PatternIntegerPatternIntegerChoiceEnum) PatternIntegerPatternInteger
	// HasChoice checks if Choice has been set in PatternIntegerPatternInteger
	HasChoice() bool
	// Value returns uint32, set in PatternIntegerPatternInteger.
	Value() uint32
	// SetValue assigns uint32 provided by user to PatternIntegerPatternInteger
	SetValue(value uint32) PatternIntegerPatternInteger
	// HasValue checks if Value has been set in PatternIntegerPatternInteger
	HasValue() bool
	// Values returns []uint32, set in PatternIntegerPatternInteger.
	Values() []uint32
	// SetValues assigns []uint32 provided by user to PatternIntegerPatternInteger
	SetValues(value []uint32) PatternIntegerPatternInteger
	// Increment returns PatternIntegerPatternIntegerCounter, set in PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	Increment() PatternIntegerPatternIntegerCounter
	// SetIncrement assigns PatternIntegerPatternIntegerCounter provided by user to PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	SetIncrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger
	// HasIncrement checks if Increment has been set in PatternIntegerPatternInteger
	HasIncrement() bool
	// Decrement returns PatternIntegerPatternIntegerCounter, set in PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	Decrement() PatternIntegerPatternIntegerCounter
	// SetDecrement assigns PatternIntegerPatternIntegerCounter provided by user to PatternIntegerPatternInteger.
	// PatternIntegerPatternIntegerCounter is integer counter pattern
	SetDecrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger
	// HasDecrement checks if Decrement has been set in PatternIntegerPatternInteger
	HasDecrement() bool
	setNil()
}

type PatternIntegerPatternIntegerChoiceEnum string

// Enum of Choice on PatternIntegerPatternInteger
var PatternIntegerPatternIntegerChoice = struct {
	VALUE     PatternIntegerPatternIntegerChoiceEnum
	VALUES    PatternIntegerPatternIntegerChoiceEnum
	INCREMENT PatternIntegerPatternIntegerChoiceEnum
	DECREMENT PatternIntegerPatternIntegerChoiceEnum
}{
	VALUE:     PatternIntegerPatternIntegerChoiceEnum("value"),
	VALUES:    PatternIntegerPatternIntegerChoiceEnum("values"),
	INCREMENT: PatternIntegerPatternIntegerChoiceEnum("increment"),
	DECREMENT: PatternIntegerPatternIntegerChoiceEnum("decrement"),
}

func (obj *patternIntegerPatternInteger) Choice() PatternIntegerPatternIntegerChoiceEnum {
	return PatternIntegerPatternIntegerChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *patternIntegerPatternInteger) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternIntegerPatternInteger) SetChoice(value PatternIntegerPatternIntegerChoiceEnum) PatternIntegerPatternInteger {
	intValue, ok := openapi.PatternIntegerPatternInteger_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternIntegerPatternIntegerChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternIntegerPatternInteger_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Decrement = nil
	obj.decrementHolder = nil
	obj.obj.Increment = nil
	obj.incrementHolder = nil
	obj.obj.Values = nil
	obj.obj.Value = nil

	if value == PatternIntegerPatternIntegerChoice.VALUE {
		defaultValue := uint32(0)
		obj.obj.Value = &defaultValue
	}

	if value == PatternIntegerPatternIntegerChoice.VALUES {
		defaultValue := []uint32{0}
		obj.obj.Values = defaultValue
	}

	if value == PatternIntegerPatternIntegerChoice.INCREMENT {
		obj.obj.Increment = NewPatternIntegerPatternIntegerCounter().Msg()
	}

	if value == PatternIntegerPatternIntegerChoice.DECREMENT {
		obj.obj.Decrement = NewPatternIntegerPatternIntegerCounter().Msg()
	}

	return obj
}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternInteger) Value() uint32 {

	if obj.obj.Value == nil {
		obj.SetChoice(PatternIntegerPatternIntegerChoice.VALUE)
	}

	return *obj.obj.Value

}

// description is TBD
// Value returns a uint32
func (obj *patternIntegerPatternInteger) HasValue() bool {
	return obj.obj.Value != nil
}

// description is TBD
// SetValue sets the uint32 value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetValue(value uint32) PatternIntegerPatternInteger {
	obj.SetChoice(PatternIntegerPatternIntegerChoice.VALUE)
	obj.obj.Value = &value
	return obj
}

// description is TBD
// Values returns a []uint32
func (obj *patternIntegerPatternInteger) Values() []uint32 {
	if obj.obj.Values == nil {
		obj.SetValues([]uint32{0})
	}
	return obj.obj.Values
}

// description is TBD
// SetValues sets the []uint32 value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetValues(value []uint32) PatternIntegerPatternInteger {
	obj.SetChoice(PatternIntegerPatternIntegerChoice.VALUES)
	if obj.obj.Values == nil {
		obj.obj.Values = make([]uint32, 0)
	}
	obj.obj.Values = value

	return obj
}

// description is TBD
// Increment returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) Increment() PatternIntegerPatternIntegerCounter {
	if obj.obj.Increment == nil {
		obj.SetChoice(PatternIntegerPatternIntegerChoice.INCREMENT)
	}
	if obj.incrementHolder == nil {
		obj.incrementHolder = &patternIntegerPatternIntegerCounter{obj: obj.obj.Increment}
	}
	return obj.incrementHolder
}

// description is TBD
// Increment returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) HasIncrement() bool {
	return obj.obj.Increment != nil
}

// description is TBD
// SetIncrement sets the PatternIntegerPatternIntegerCounter value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetIncrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger {
	obj.SetChoice(PatternIntegerPatternIntegerChoice.INCREMENT)
	obj.incrementHolder = nil
	obj.obj.Increment = value.Msg()

	return obj
}

// description is TBD
// Decrement returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) Decrement() PatternIntegerPatternIntegerCounter {
	if obj.obj.Decrement == nil {
		obj.SetChoice(PatternIntegerPatternIntegerChoice.DECREMENT)
	}
	if obj.decrementHolder == nil {
		obj.decrementHolder = &patternIntegerPatternIntegerCounter{obj: obj.obj.Decrement}
	}
	return obj.decrementHolder
}

// description is TBD
// Decrement returns a PatternIntegerPatternIntegerCounter
func (obj *patternIntegerPatternInteger) HasDecrement() bool {
	return obj.obj.Decrement != nil
}

// description is TBD
// SetDecrement sets the PatternIntegerPatternIntegerCounter value in the PatternIntegerPatternInteger object
func (obj *patternIntegerPatternInteger) SetDecrement(value PatternIntegerPatternIntegerCounter) PatternIntegerPatternInteger {
	obj.SetChoice(PatternIntegerPatternIntegerChoice.DECREMENT)
	obj.decrementHolder = nil
	obj.obj.Decrement = value.Msg()

	return obj
}

func (obj *patternIntegerPatternInteger) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Value != nil {

		if *obj.obj.Value > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternInteger.Value <= 255 but Got %d", *obj.obj.Value))
		}

	}

	if obj.obj.Values != nil {

		for _, item := range obj.obj.Values {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("min(uint32) <= PatternIntegerPatternInteger.Values <= 255 but Got %d", item))
			}

		}

	}

	if obj.obj.Increment != nil {

		obj.Increment().validateObj(vObj, set_default)
	}

	if obj.obj.Decrement != nil {

		obj.Decrement().validateObj(vObj, set_default)
	}

}

func (obj *patternIntegerPatternInteger) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternIntegerPatternIntegerChoice.VALUE)

	}

}

// ***** PatternChecksumPatternChecksum *****
type patternChecksumPatternChecksum struct {
	validation
	obj *openapi.PatternChecksumPatternChecksum
}

func NewPatternChecksumPatternChecksum() PatternChecksumPatternChecksum {
	obj := patternChecksumPatternChecksum{obj: &openapi.PatternChecksumPatternChecksum{}}
	obj.setDefault()
	return &obj
}

func (obj *patternChecksumPatternChecksum) Msg() *openapi.PatternChecksumPatternChecksum {
	return obj.obj
}

func (obj *patternChecksumPatternChecksum) SetMsg(msg *openapi.PatternChecksumPatternChecksum) PatternChecksumPatternChecksum {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternChecksumPatternChecksum) ToProto() (*openapi.PatternChecksumPatternChecksum, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternChecksumPatternChecksum) FromProto(msg *openapi.PatternChecksumPatternChecksum) (PatternChecksumPatternChecksum, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternChecksumPatternChecksum) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternChecksumPatternChecksum) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternChecksumPatternChecksum) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternChecksumPatternChecksum) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternChecksumPatternChecksum) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternChecksumPatternChecksum) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternChecksumPatternChecksum) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternChecksumPatternChecksum) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternChecksumPatternChecksum) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternChecksumPatternChecksum) Clone() (PatternChecksumPatternChecksum, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternChecksumPatternChecksum()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternChecksumPatternChecksum is tBD
type PatternChecksumPatternChecksum interface {
	Validation
	// Msg marshals PatternChecksumPatternChecksum to protobuf object *openapi.PatternChecksumPatternChecksum
	// and doesn't set defaults
	Msg() *openapi.PatternChecksumPatternChecksum
	// SetMsg unmarshals PatternChecksumPatternChecksum from protobuf object *openapi.PatternChecksumPatternChecksum
	// and doesn't set defaults
	SetMsg(*openapi.PatternChecksumPatternChecksum) PatternChecksumPatternChecksum
	// ToProto marshals PatternChecksumPatternChecksum to protobuf object *openapi.PatternChecksumPatternChecksum
	ToProto() (*openapi.PatternChecksumPatternChecksum, error)
	// ToPbText marshals PatternChecksumPatternChecksum to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternChecksumPatternChecksum to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternChecksumPatternChecksum to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternChecksumPatternChecksum from protobuf object *openapi.PatternChecksumPatternChecksum
	FromProto(msg *openapi.PatternChecksumPatternChecksum) (PatternChecksumPatternChecksum, error)
	// FromPbText unmarshals PatternChecksumPatternChecksum from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternChecksumPatternChecksum from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternChecksumPatternChecksum from JSON text
	FromJson(value string) error
	// Validate validates PatternChecksumPatternChecksum
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternChecksumPatternChecksum, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns PatternChecksumPatternChecksumChoiceEnum, set in PatternChecksumPatternChecksum
	Choice() PatternChecksumPatternChecksumChoiceEnum
	// SetChoice assigns PatternChecksumPatternChecksumChoiceEnum provided by user to PatternChecksumPatternChecksum
	SetChoice(value PatternChecksumPatternChecksumChoiceEnum) PatternChecksumPatternChecksum
	// HasChoice checks if Choice has been set in PatternChecksumPatternChecksum
	HasChoice() bool
	// Generated returns PatternChecksumPatternChecksumGeneratedEnum, set in PatternChecksumPatternChecksum
	Generated() PatternChecksumPatternChecksumGeneratedEnum
	// SetGenerated assigns PatternChecksumPatternChecksumGeneratedEnum provided by user to PatternChecksumPatternChecksum
	SetGenerated(value PatternChecksumPatternChecksumGeneratedEnum) PatternChecksumPatternChecksum
	// HasGenerated checks if Generated has been set in PatternChecksumPatternChecksum
	HasGenerated() bool
	// Custom returns uint32, set in PatternChecksumPatternChecksum.
	Custom() uint32
	// SetCustom assigns uint32 provided by user to PatternChecksumPatternChecksum
	SetCustom(value uint32) PatternChecksumPatternChecksum
	// HasCustom checks if Custom has been set in PatternChecksumPatternChecksum
	HasCustom() bool
}

type PatternChecksumPatternChecksumChoiceEnum string

// Enum of Choice on PatternChecksumPatternChecksum
var PatternChecksumPatternChecksumChoice = struct {
	GENERATED PatternChecksumPatternChecksumChoiceEnum
	CUSTOM    PatternChecksumPatternChecksumChoiceEnum
}{
	GENERATED: PatternChecksumPatternChecksumChoiceEnum("generated"),
	CUSTOM:    PatternChecksumPatternChecksumChoiceEnum("custom"),
}

func (obj *patternChecksumPatternChecksum) Choice() PatternChecksumPatternChecksumChoiceEnum {
	return PatternChecksumPatternChecksumChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of checksum
// Choice returns a string
func (obj *patternChecksumPatternChecksum) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *patternChecksumPatternChecksum) SetChoice(value PatternChecksumPatternChecksumChoiceEnum) PatternChecksumPatternChecksum {
	intValue, ok := openapi.PatternChecksumPatternChecksum_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternChecksumChoiceEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternChecksum_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.obj.Generated = openapi.PatternChecksumPatternChecksum_Generated_unspecified.Enum()
	return obj
}

type PatternChecksumPatternChecksumGeneratedEnum string

// Enum of Generated on PatternChecksumPatternChecksum
var PatternChecksumPatternChecksumGenerated = struct {
	GOOD PatternChecksumPatternChecksumGeneratedEnum
	BAD  PatternChecksumPatternChecksumGeneratedEnum
}{
	GOOD: PatternChecksumPatternChecksumGeneratedEnum("good"),
	BAD:  PatternChecksumPatternChecksumGeneratedEnum("bad"),
}

func (obj *patternChecksumPatternChecksum) Generated() PatternChecksumPatternChecksumGeneratedEnum {
	return PatternChecksumPatternChecksumGeneratedEnum(obj.obj.Generated.Enum().String())
}

// A system generated checksum value
// Generated returns a string
func (obj *patternChecksumPatternChecksum) HasGenerated() bool {
	return obj.obj.Generated != nil
}

func (obj *patternChecksumPatternChecksum) SetGenerated(value PatternChecksumPatternChecksumGeneratedEnum) PatternChecksumPatternChecksum {
	intValue, ok := openapi.PatternChecksumPatternChecksum_Generated_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on PatternChecksumPatternChecksumGeneratedEnum", string(value)))
		return obj
	}
	enumValue := openapi.PatternChecksumPatternChecksum_Generated_Enum(intValue)
	obj.obj.Generated = &enumValue

	return obj
}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternChecksum) Custom() uint32 {

	if obj.obj.Custom == nil {
		obj.SetChoice(PatternChecksumPatternChecksumChoice.CUSTOM)
	}

	return *obj.obj.Custom

}

// A custom checksum value
// Custom returns a uint32
func (obj *patternChecksumPatternChecksum) HasCustom() bool {
	return obj.obj.Custom != nil
}

// A custom checksum value
// SetCustom sets the uint32 value in the PatternChecksumPatternChecksum object
func (obj *patternChecksumPatternChecksum) SetCustom(value uint32) PatternChecksumPatternChecksum {
	obj.SetChoice(PatternChecksumPatternChecksumChoice.CUSTOM)
	obj.obj.Custom = &value
	return obj
}

func (obj *patternChecksumPatternChecksum) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		if *obj.obj.Custom > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternChecksumPatternChecksum.Custom <= 255 but Got %d", *obj.obj.Custom))
		}

	}

}

func (obj *patternChecksumPatternChecksum) setDefault() {
	if obj.obj.Choice == nil {
		obj.SetChoice(PatternChecksumPatternChecksumChoice.GENERATED)
		if obj.obj.Generated.Number() == 0 {
			obj.SetGenerated(PatternChecksumPatternChecksumGenerated.GOOD)

		}

	}

}

// ***** PatternPrefixConfigAutoFieldTestCounter *****
type patternPrefixConfigAutoFieldTestCounter struct {
	validation
	obj *openapi.PatternPrefixConfigAutoFieldTestCounter
}

func NewPatternPrefixConfigAutoFieldTestCounter() PatternPrefixConfigAutoFieldTestCounter {
	obj := patternPrefixConfigAutoFieldTestCounter{obj: &openapi.PatternPrefixConfigAutoFieldTestCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Msg() *openapi.PatternPrefixConfigAutoFieldTestCounter {
	return obj.obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) SetMsg(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTestCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) ToProto() (*openapi.PatternPrefixConfigAutoFieldTestCounter, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) FromProto(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) (PatternPrefixConfigAutoFieldTestCounter, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternPrefixConfigAutoFieldTestCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternPrefixConfigAutoFieldTestCounter) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternPrefixConfigAutoFieldTestCounter) Clone() (PatternPrefixConfigAutoFieldTestCounter, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternPrefixConfigAutoFieldTestCounter()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternPrefixConfigAutoFieldTestCounter is integer counter pattern
type PatternPrefixConfigAutoFieldTestCounter interface {
	Validation
	// Msg marshals PatternPrefixConfigAutoFieldTestCounter to protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	// and doesn't set defaults
	Msg() *openapi.PatternPrefixConfigAutoFieldTestCounter
	// SetMsg unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	// and doesn't set defaults
	SetMsg(*openapi.PatternPrefixConfigAutoFieldTestCounter) PatternPrefixConfigAutoFieldTestCounter
	// ToProto marshals PatternPrefixConfigAutoFieldTestCounter to protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	ToProto() (*openapi.PatternPrefixConfigAutoFieldTestCounter, error)
	// ToPbText marshals PatternPrefixConfigAutoFieldTestCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternPrefixConfigAutoFieldTestCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternPrefixConfigAutoFieldTestCounter to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf object *openapi.PatternPrefixConfigAutoFieldTestCounter
	FromProto(msg *openapi.PatternPrefixConfigAutoFieldTestCounter) (PatternPrefixConfigAutoFieldTestCounter, error)
	// FromPbText unmarshals PatternPrefixConfigAutoFieldTestCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternPrefixConfigAutoFieldTestCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternPrefixConfigAutoFieldTestCounter from JSON text
	FromJson(value string) error
	// Validate validates PatternPrefixConfigAutoFieldTestCounter
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternPrefixConfigAutoFieldTestCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetStart(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasStart checks if Start has been set in PatternPrefixConfigAutoFieldTestCounter
	HasStart() bool
	// Step returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetStep(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasStep checks if Step has been set in PatternPrefixConfigAutoFieldTestCounter
	HasStep() bool
	// Count returns uint32, set in PatternPrefixConfigAutoFieldTestCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternPrefixConfigAutoFieldTestCounter
	SetCount(value uint32) PatternPrefixConfigAutoFieldTestCounter
	// HasCount checks if Count has been set in PatternPrefixConfigAutoFieldTestCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetStart(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetStep(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternPrefixConfigAutoFieldTestCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternPrefixConfigAutoFieldTestCounter object
func (obj *patternPrefixConfigAutoFieldTestCounter) SetCount(value uint32) PatternPrefixConfigAutoFieldTestCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternPrefixConfigAutoFieldTestCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternPrefixConfigAutoFieldTestCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternPrefixConfigAutoFieldTestCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

// ***** RequiredChoiceIntermediate *****
type requiredChoiceIntermediate struct {
	validation
	obj        *openapi.RequiredChoiceIntermediate
	leafHolder RequiredChoiceIntermeLeaf
}

func NewRequiredChoiceIntermediate() RequiredChoiceIntermediate {
	obj := requiredChoiceIntermediate{obj: &openapi.RequiredChoiceIntermediate{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceIntermediate) Msg() *openapi.RequiredChoiceIntermediate {
	return obj.obj
}

func (obj *requiredChoiceIntermediate) SetMsg(msg *openapi.RequiredChoiceIntermediate) RequiredChoiceIntermediate {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *requiredChoiceIntermediate) ToProto() (*openapi.RequiredChoiceIntermediate, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *requiredChoiceIntermediate) FromProto(msg *openapi.RequiredChoiceIntermediate) (RequiredChoiceIntermediate, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *requiredChoiceIntermediate) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *requiredChoiceIntermediate) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *requiredChoiceIntermediate) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceIntermediate) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *requiredChoiceIntermediate) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceIntermediate) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}
	obj.setNil()
	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *requiredChoiceIntermediate) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermediate) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermediate) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceIntermediate) Clone() (RequiredChoiceIntermediate, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceIntermediate()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

func (obj *requiredChoiceIntermediate) setNil() {
	obj.leafHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RequiredChoiceIntermediate is description is TBD
type RequiredChoiceIntermediate interface {
	Validation
	// Msg marshals RequiredChoiceIntermediate to protobuf object *openapi.RequiredChoiceIntermediate
	// and doesn't set defaults
	Msg() *openapi.RequiredChoiceIntermediate
	// SetMsg unmarshals RequiredChoiceIntermediate from protobuf object *openapi.RequiredChoiceIntermediate
	// and doesn't set defaults
	SetMsg(*openapi.RequiredChoiceIntermediate) RequiredChoiceIntermediate
	// ToProto marshals RequiredChoiceIntermediate to protobuf object *openapi.RequiredChoiceIntermediate
	ToProto() (*openapi.RequiredChoiceIntermediate, error)
	// ToPbText marshals RequiredChoiceIntermediate to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceIntermediate to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceIntermediate to JSON text
	ToJson() (string, error)
	// FromProto unmarshals RequiredChoiceIntermediate from protobuf object *openapi.RequiredChoiceIntermediate
	FromProto(msg *openapi.RequiredChoiceIntermediate) (RequiredChoiceIntermediate, error)
	// FromPbText unmarshals RequiredChoiceIntermediate from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceIntermediate from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceIntermediate from JSON text
	FromJson(value string) error
	// Validate validates RequiredChoiceIntermediate
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceIntermediate, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns RequiredChoiceIntermediateChoiceEnum, set in RequiredChoiceIntermediate
	Choice() RequiredChoiceIntermediateChoiceEnum
	// SetChoice assigns RequiredChoiceIntermediateChoiceEnum provided by user to RequiredChoiceIntermediate
	SetChoice(value RequiredChoiceIntermediateChoiceEnum) RequiredChoiceIntermediate
	// FA returns string, set in RequiredChoiceIntermediate.
	FA() string
	// SetFA assigns string provided by user to RequiredChoiceIntermediate
	SetFA(value string) RequiredChoiceIntermediate
	// HasFA checks if FA has been set in RequiredChoiceIntermediate
	HasFA() bool
	// Leaf returns RequiredChoiceIntermeLeaf, set in RequiredChoiceIntermediate.
	// RequiredChoiceIntermeLeaf is description is TBD
	Leaf() RequiredChoiceIntermeLeaf
	// SetLeaf assigns RequiredChoiceIntermeLeaf provided by user to RequiredChoiceIntermediate.
	// RequiredChoiceIntermeLeaf is description is TBD
	SetLeaf(value RequiredChoiceIntermeLeaf) RequiredChoiceIntermediate
	// HasLeaf checks if Leaf has been set in RequiredChoiceIntermediate
	HasLeaf() bool
	setNil()
}

type RequiredChoiceIntermediateChoiceEnum string

// Enum of Choice on RequiredChoiceIntermediate
var RequiredChoiceIntermediateChoice = struct {
	F_A  RequiredChoiceIntermediateChoiceEnum
	LEAF RequiredChoiceIntermediateChoiceEnum
}{
	F_A:  RequiredChoiceIntermediateChoiceEnum("f_a"),
	LEAF: RequiredChoiceIntermediateChoiceEnum("leaf"),
}

func (obj *requiredChoiceIntermediate) Choice() RequiredChoiceIntermediateChoiceEnum {
	return RequiredChoiceIntermediateChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *requiredChoiceIntermediate) SetChoice(value RequiredChoiceIntermediateChoiceEnum) RequiredChoiceIntermediate {
	intValue, ok := openapi.RequiredChoiceIntermediate_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RequiredChoiceIntermediateChoiceEnum", string(value)))
		return obj
	}
	obj.obj.Choice = openapi.RequiredChoiceIntermediate_Choice_Enum(intValue)
	obj.obj.Leaf = nil
	obj.leafHolder = nil
	obj.obj.FA = nil

	if value == RequiredChoiceIntermediateChoice.F_A {
		defaultValue := "some string"
		obj.obj.FA = &defaultValue
	}

	if value == RequiredChoiceIntermediateChoice.LEAF {
		obj.obj.Leaf = NewRequiredChoiceIntermeLeaf().Msg()
	}

	return obj
}

// description is TBD
// FA returns a string
func (obj *requiredChoiceIntermediate) FA() string {

	if obj.obj.FA == nil {
		obj.SetChoice(RequiredChoiceIntermediateChoice.F_A)
	}

	return *obj.obj.FA

}

// description is TBD
// FA returns a string
func (obj *requiredChoiceIntermediate) HasFA() bool {
	return obj.obj.FA != nil
}

// description is TBD
// SetFA sets the string value in the RequiredChoiceIntermediate object
func (obj *requiredChoiceIntermediate) SetFA(value string) RequiredChoiceIntermediate {
	obj.SetChoice(RequiredChoiceIntermediateChoice.F_A)
	obj.obj.FA = &value
	return obj
}

// description is TBD
// Leaf returns a RequiredChoiceIntermeLeaf
func (obj *requiredChoiceIntermediate) Leaf() RequiredChoiceIntermeLeaf {
	if obj.obj.Leaf == nil {
		obj.SetChoice(RequiredChoiceIntermediateChoice.LEAF)
	}
	if obj.leafHolder == nil {
		obj.leafHolder = &requiredChoiceIntermeLeaf{obj: obj.obj.Leaf}
	}
	return obj.leafHolder
}

// description is TBD
// Leaf returns a RequiredChoiceIntermeLeaf
func (obj *requiredChoiceIntermediate) HasLeaf() bool {
	return obj.obj.Leaf != nil
}

// description is TBD
// SetLeaf sets the RequiredChoiceIntermeLeaf value in the RequiredChoiceIntermediate object
func (obj *requiredChoiceIntermediate) SetLeaf(value RequiredChoiceIntermeLeaf) RequiredChoiceIntermediate {
	obj.SetChoice(RequiredChoiceIntermediateChoice.LEAF)
	obj.leafHolder = nil
	obj.obj.Leaf = value.Msg()

	return obj
}

func (obj *requiredChoiceIntermediate) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice.Number() == 0 {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface RequiredChoiceIntermediate")
	}

	if obj.obj.Leaf != nil {

		obj.Leaf().validateObj(vObj, set_default)
	}

}

func (obj *requiredChoiceIntermediate) setDefault() {
	if obj.obj.FA == nil {
		obj.SetFA("some string")
	}

}

// ***** PortMetric *****
type portMetric struct {
	validation
	obj *openapi.PortMetric
}

func NewPortMetric() PortMetric {
	obj := portMetric{obj: &openapi.PortMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *portMetric) Msg() *openapi.PortMetric {
	return obj.obj
}

func (obj *portMetric) SetMsg(msg *openapi.PortMetric) PortMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *portMetric) ToProto() (*openapi.PortMetric, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *portMetric) FromProto(msg *openapi.PortMetric) (PortMetric, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *portMetric) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *portMetric) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *portMetric) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *portMetric) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *portMetric) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *portMetric) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *portMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *portMetric) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *portMetric) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *portMetric) Clone() (PortMetric, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPortMetric()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PortMetric is description is TBD
type PortMetric interface {
	Validation
	// Msg marshals PortMetric to protobuf object *openapi.PortMetric
	// and doesn't set defaults
	Msg() *openapi.PortMetric
	// SetMsg unmarshals PortMetric from protobuf object *openapi.PortMetric
	// and doesn't set defaults
	SetMsg(*openapi.PortMetric) PortMetric
	// ToProto marshals PortMetric to protobuf object *openapi.PortMetric
	ToProto() (*openapi.PortMetric, error)
	// ToPbText marshals PortMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PortMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals PortMetric to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PortMetric from protobuf object *openapi.PortMetric
	FromProto(msg *openapi.PortMetric) (PortMetric, error)
	// FromPbText unmarshals PortMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PortMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PortMetric from JSON text
	FromJson(value string) error
	// Validate validates PortMetric
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PortMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in PortMetric.
	Name() string
	// SetName assigns string provided by user to PortMetric
	SetName(value string) PortMetric
	// TxFrames returns float64, set in PortMetric.
	TxFrames() float64
	// SetTxFrames assigns float64 provided by user to PortMetric
	SetTxFrames(value float64) PortMetric
	// RxFrames returns float64, set in PortMetric.
	RxFrames() float64
	// SetRxFrames assigns float64 provided by user to PortMetric
	SetRxFrames(value float64) PortMetric
}

// description is TBD
// Name returns a string
func (obj *portMetric) Name() string {

	return obj.obj.Name
}

// description is TBD
// SetName sets the string value in the PortMetric object
func (obj *portMetric) SetName(value string) PortMetric {

	obj.obj.Name = value
	return obj
}

// description is TBD
// TxFrames returns a float64
func (obj *portMetric) TxFrames() float64 {

	return obj.obj.TxFrames
}

// description is TBD
// SetTxFrames sets the float64 value in the PortMetric object
func (obj *portMetric) SetTxFrames(value float64) PortMetric {

	obj.obj.TxFrames = value
	return obj
}

// description is TBD
// RxFrames returns a float64
func (obj *portMetric) RxFrames() float64 {

	return obj.obj.RxFrames
}

// description is TBD
// SetRxFrames sets the float64 value in the PortMetric object
func (obj *portMetric) SetRxFrames(value float64) PortMetric {

	obj.obj.RxFrames = value
	return obj
}

func (obj *portMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface PortMetric")
	}
}

func (obj *portMetric) setDefault() {

}

// ***** FlowMetric *****
type flowMetric struct {
	validation
	obj *openapi.FlowMetric
}

func NewFlowMetric() FlowMetric {
	obj := flowMetric{obj: &openapi.FlowMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *flowMetric) Msg() *openapi.FlowMetric {
	return obj.obj
}

func (obj *flowMetric) SetMsg(msg *openapi.FlowMetric) FlowMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *flowMetric) ToProto() (*openapi.FlowMetric, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *flowMetric) FromProto(msg *openapi.FlowMetric) (FlowMetric, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *flowMetric) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *flowMetric) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *flowMetric) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *flowMetric) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *flowMetric) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *flowMetric) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *flowMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowMetric) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowMetric) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowMetric) Clone() (FlowMetric, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowMetric()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// FlowMetric is description is TBD
type FlowMetric interface {
	Validation
	// Msg marshals FlowMetric to protobuf object *openapi.FlowMetric
	// and doesn't set defaults
	Msg() *openapi.FlowMetric
	// SetMsg unmarshals FlowMetric from protobuf object *openapi.FlowMetric
	// and doesn't set defaults
	SetMsg(*openapi.FlowMetric) FlowMetric
	// ToProto marshals FlowMetric to protobuf object *openapi.FlowMetric
	ToProto() (*openapi.FlowMetric, error)
	// ToPbText marshals FlowMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowMetric to JSON text
	ToJson() (string, error)
	// FromProto unmarshals FlowMetric from protobuf object *openapi.FlowMetric
	FromProto(msg *openapi.FlowMetric) (FlowMetric, error)
	// FromPbText unmarshals FlowMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowMetric from JSON text
	FromJson(value string) error
	// Validate validates FlowMetric
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in FlowMetric.
	Name() string
	// SetName assigns string provided by user to FlowMetric
	SetName(value string) FlowMetric
	// TxFrames returns float64, set in FlowMetric.
	TxFrames() float64
	// SetTxFrames assigns float64 provided by user to FlowMetric
	SetTxFrames(value float64) FlowMetric
	// RxFrames returns float64, set in FlowMetric.
	RxFrames() float64
	// SetRxFrames assigns float64 provided by user to FlowMetric
	SetRxFrames(value float64) FlowMetric
}

// description is TBD
// Name returns a string
func (obj *flowMetric) Name() string {

	return obj.obj.Name
}

// description is TBD
// SetName sets the string value in the FlowMetric object
func (obj *flowMetric) SetName(value string) FlowMetric {

	obj.obj.Name = value
	return obj
}

// description is TBD
// TxFrames returns a float64
func (obj *flowMetric) TxFrames() float64 {

	return obj.obj.TxFrames
}

// description is TBD
// SetTxFrames sets the float64 value in the FlowMetric object
func (obj *flowMetric) SetTxFrames(value float64) FlowMetric {

	obj.obj.TxFrames = value
	return obj
}

// description is TBD
// RxFrames returns a float64
func (obj *flowMetric) RxFrames() float64 {

	return obj.obj.RxFrames
}

// description is TBD
// SetRxFrames sets the float64 value in the FlowMetric object
func (obj *flowMetric) SetRxFrames(value float64) FlowMetric {

	obj.obj.RxFrames = value
	return obj
}

func (obj *flowMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == "" {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface FlowMetric")
	}
}

func (obj *flowMetric) setDefault() {

}

// ***** LevelThree *****
type levelThree struct {
	validation
	obj *openapi.LevelThree
}

func NewLevelThree() LevelThree {
	obj := levelThree{obj: &openapi.LevelThree{}}
	obj.setDefault()
	return &obj
}

func (obj *levelThree) Msg() *openapi.LevelThree {
	return obj.obj
}

func (obj *levelThree) SetMsg(msg *openapi.LevelThree) LevelThree {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *levelThree) ToProto() (*openapi.LevelThree, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *levelThree) FromProto(msg *openapi.LevelThree) (LevelThree, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *levelThree) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *levelThree) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *levelThree) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelThree) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *levelThree) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *levelThree) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *levelThree) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *levelThree) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *levelThree) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *levelThree) Clone() (LevelThree, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLevelThree()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// LevelThree is test Level3
type LevelThree interface {
	Validation
	// Msg marshals LevelThree to protobuf object *openapi.LevelThree
	// and doesn't set defaults
	Msg() *openapi.LevelThree
	// SetMsg unmarshals LevelThree from protobuf object *openapi.LevelThree
	// and doesn't set defaults
	SetMsg(*openapi.LevelThree) LevelThree
	// ToProto marshals LevelThree to protobuf object *openapi.LevelThree
	ToProto() (*openapi.LevelThree, error)
	// ToPbText marshals LevelThree to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LevelThree to YAML text
	ToYaml() (string, error)
	// ToJson marshals LevelThree to JSON text
	ToJson() (string, error)
	// FromProto unmarshals LevelThree from protobuf object *openapi.LevelThree
	FromProto(msg *openapi.LevelThree) (LevelThree, error)
	// FromPbText unmarshals LevelThree from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LevelThree from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LevelThree from JSON text
	FromJson(value string) error
	// Validate validates LevelThree
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LevelThree, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// L3P1 returns string, set in LevelThree.
	L3P1() string
	// SetL3P1 assigns string provided by user to LevelThree
	SetL3P1(value string) LevelThree
	// HasL3P1 checks if L3P1 has been set in LevelThree
	HasL3P1() bool
}

// Set value at Level 3
// L3P1 returns a string
func (obj *levelThree) L3P1() string {

	return *obj.obj.L3P1

}

// Set value at Level 3
// L3P1 returns a string
func (obj *levelThree) HasL3P1() bool {
	return obj.obj.L3P1 != nil
}

// Set value at Level 3
// SetL3P1 sets the string value in the LevelThree object
func (obj *levelThree) SetL3P1(value string) LevelThree {

	obj.obj.L3P1 = &value
	return obj
}

func (obj *levelThree) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *levelThree) setDefault() {

}

// ***** PatternIpv4PatternIpv4Counter *****
type patternIpv4PatternIpv4Counter struct {
	validation
	obj *openapi.PatternIpv4PatternIpv4Counter
}

func NewPatternIpv4PatternIpv4Counter() PatternIpv4PatternIpv4Counter {
	obj := patternIpv4PatternIpv4Counter{obj: &openapi.PatternIpv4PatternIpv4Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv4PatternIpv4Counter) Msg() *openapi.PatternIpv4PatternIpv4Counter {
	return obj.obj
}

func (obj *patternIpv4PatternIpv4Counter) SetMsg(msg *openapi.PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIpv4PatternIpv4Counter) ToProto() (*openapi.PatternIpv4PatternIpv4Counter, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIpv4PatternIpv4Counter) FromProto(msg *openapi.PatternIpv4PatternIpv4Counter) (PatternIpv4PatternIpv4Counter, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIpv4PatternIpv4Counter) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIpv4PatternIpv4Counter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIpv4PatternIpv4Counter) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv4PatternIpv4Counter) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIpv4PatternIpv4Counter) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv4PatternIpv4Counter) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv4PatternIpv4Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4Counter) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv4PatternIpv4Counter) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv4PatternIpv4Counter) Clone() (PatternIpv4PatternIpv4Counter, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv4PatternIpv4Counter()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternIpv4PatternIpv4Counter is ipv4 counter pattern
type PatternIpv4PatternIpv4Counter interface {
	Validation
	// Msg marshals PatternIpv4PatternIpv4Counter to protobuf object *openapi.PatternIpv4PatternIpv4Counter
	// and doesn't set defaults
	Msg() *openapi.PatternIpv4PatternIpv4Counter
	// SetMsg unmarshals PatternIpv4PatternIpv4Counter from protobuf object *openapi.PatternIpv4PatternIpv4Counter
	// and doesn't set defaults
	SetMsg(*openapi.PatternIpv4PatternIpv4Counter) PatternIpv4PatternIpv4Counter
	// ToProto marshals PatternIpv4PatternIpv4Counter to protobuf object *openapi.PatternIpv4PatternIpv4Counter
	ToProto() (*openapi.PatternIpv4PatternIpv4Counter, error)
	// ToPbText marshals PatternIpv4PatternIpv4Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv4PatternIpv4Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv4PatternIpv4Counter to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIpv4PatternIpv4Counter from protobuf object *openapi.PatternIpv4PatternIpv4Counter
	FromProto(msg *openapi.PatternIpv4PatternIpv4Counter) (PatternIpv4PatternIpv4Counter, error)
	// FromPbText unmarshals PatternIpv4PatternIpv4Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv4PatternIpv4Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv4PatternIpv4Counter from JSON text
	FromJson(value string) error
	// Validate validates PatternIpv4PatternIpv4Counter
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv4PatternIpv4Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternIpv4PatternIpv4Counter.
	Start() string
	// SetStart assigns string provided by user to PatternIpv4PatternIpv4Counter
	SetStart(value string) PatternIpv4PatternIpv4Counter
	// HasStart checks if Start has been set in PatternIpv4PatternIpv4Counter
	HasStart() bool
	// Step returns string, set in PatternIpv4PatternIpv4Counter.
	Step() string
	// SetStep assigns string provided by user to PatternIpv4PatternIpv4Counter
	SetStep(value string) PatternIpv4PatternIpv4Counter
	// HasStep checks if Step has been set in PatternIpv4PatternIpv4Counter
	HasStep() bool
	// Count returns uint32, set in PatternIpv4PatternIpv4Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIpv4PatternIpv4Counter
	SetCount(value uint32) PatternIpv4PatternIpv4Counter
	// HasCount checks if Count has been set in PatternIpv4PatternIpv4Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternIpv4Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternIpv4PatternIpv4Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetStart(value string) PatternIpv4PatternIpv4Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternIpv4Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternIpv4PatternIpv4Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetStep(value string) PatternIpv4PatternIpv4Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternIpv4Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIpv4PatternIpv4Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIpv4PatternIpv4Counter object
func (obj *patternIpv4PatternIpv4Counter) SetCount(value uint32) PatternIpv4PatternIpv4Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIpv4PatternIpv4Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv4(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv4(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv4PatternIpv4Counter.Step"))
		}

	}

}

func (obj *patternIpv4PatternIpv4Counter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("0.0.0.0")
	}
	if obj.obj.Step == nil {
		obj.SetStep("0.0.0.1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

// ***** PatternIpv6PatternIpv6Counter *****
type patternIpv6PatternIpv6Counter struct {
	validation
	obj *openapi.PatternIpv6PatternIpv6Counter
}

func NewPatternIpv6PatternIpv6Counter() PatternIpv6PatternIpv6Counter {
	obj := patternIpv6PatternIpv6Counter{obj: &openapi.PatternIpv6PatternIpv6Counter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIpv6PatternIpv6Counter) Msg() *openapi.PatternIpv6PatternIpv6Counter {
	return obj.obj
}

func (obj *patternIpv6PatternIpv6Counter) SetMsg(msg *openapi.PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6Counter {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIpv6PatternIpv6Counter) ToProto() (*openapi.PatternIpv6PatternIpv6Counter, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIpv6PatternIpv6Counter) FromProto(msg *openapi.PatternIpv6PatternIpv6Counter) (PatternIpv6PatternIpv6Counter, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIpv6PatternIpv6Counter) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIpv6PatternIpv6Counter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIpv6PatternIpv6Counter) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv6PatternIpv6Counter) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIpv6PatternIpv6Counter) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIpv6PatternIpv6Counter) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIpv6PatternIpv6Counter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6Counter) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIpv6PatternIpv6Counter) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIpv6PatternIpv6Counter) Clone() (PatternIpv6PatternIpv6Counter, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIpv6PatternIpv6Counter()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternIpv6PatternIpv6Counter is ipv6 counter pattern
type PatternIpv6PatternIpv6Counter interface {
	Validation
	// Msg marshals PatternIpv6PatternIpv6Counter to protobuf object *openapi.PatternIpv6PatternIpv6Counter
	// and doesn't set defaults
	Msg() *openapi.PatternIpv6PatternIpv6Counter
	// SetMsg unmarshals PatternIpv6PatternIpv6Counter from protobuf object *openapi.PatternIpv6PatternIpv6Counter
	// and doesn't set defaults
	SetMsg(*openapi.PatternIpv6PatternIpv6Counter) PatternIpv6PatternIpv6Counter
	// ToProto marshals PatternIpv6PatternIpv6Counter to protobuf object *openapi.PatternIpv6PatternIpv6Counter
	ToProto() (*openapi.PatternIpv6PatternIpv6Counter, error)
	// ToPbText marshals PatternIpv6PatternIpv6Counter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIpv6PatternIpv6Counter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIpv6PatternIpv6Counter to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIpv6PatternIpv6Counter from protobuf object *openapi.PatternIpv6PatternIpv6Counter
	FromProto(msg *openapi.PatternIpv6PatternIpv6Counter) (PatternIpv6PatternIpv6Counter, error)
	// FromPbText unmarshals PatternIpv6PatternIpv6Counter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIpv6PatternIpv6Counter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIpv6PatternIpv6Counter from JSON text
	FromJson(value string) error
	// Validate validates PatternIpv6PatternIpv6Counter
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIpv6PatternIpv6Counter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternIpv6PatternIpv6Counter.
	Start() string
	// SetStart assigns string provided by user to PatternIpv6PatternIpv6Counter
	SetStart(value string) PatternIpv6PatternIpv6Counter
	// HasStart checks if Start has been set in PatternIpv6PatternIpv6Counter
	HasStart() bool
	// Step returns string, set in PatternIpv6PatternIpv6Counter.
	Step() string
	// SetStep assigns string provided by user to PatternIpv6PatternIpv6Counter
	SetStep(value string) PatternIpv6PatternIpv6Counter
	// HasStep checks if Step has been set in PatternIpv6PatternIpv6Counter
	HasStep() bool
	// Count returns uint32, set in PatternIpv6PatternIpv6Counter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIpv6PatternIpv6Counter
	SetCount(value uint32) PatternIpv6PatternIpv6Counter
	// HasCount checks if Count has been set in PatternIpv6PatternIpv6Counter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternIpv6PatternIpv6Counter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternIpv6PatternIpv6Counter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetStart(value string) PatternIpv6PatternIpv6Counter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternIpv6PatternIpv6Counter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternIpv6PatternIpv6Counter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetStep(value string) PatternIpv6PatternIpv6Counter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIpv6PatternIpv6Counter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIpv6PatternIpv6Counter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIpv6PatternIpv6Counter object
func (obj *patternIpv6PatternIpv6Counter) SetCount(value uint32) PatternIpv6PatternIpv6Counter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIpv6PatternIpv6Counter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateIpv6(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateIpv6(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternIpv6PatternIpv6Counter.Step"))
		}

	}

}

func (obj *patternIpv6PatternIpv6Counter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("::")
	}
	if obj.obj.Step == nil {
		obj.SetStep("::1")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

// ***** PatternMacPatternMacCounter *****
type patternMacPatternMacCounter struct {
	validation
	obj *openapi.PatternMacPatternMacCounter
}

func NewPatternMacPatternMacCounter() PatternMacPatternMacCounter {
	obj := patternMacPatternMacCounter{obj: &openapi.PatternMacPatternMacCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternMacPatternMacCounter) Msg() *openapi.PatternMacPatternMacCounter {
	return obj.obj
}

func (obj *patternMacPatternMacCounter) SetMsg(msg *openapi.PatternMacPatternMacCounter) PatternMacPatternMacCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternMacPatternMacCounter) ToProto() (*openapi.PatternMacPatternMacCounter, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternMacPatternMacCounter) FromProto(msg *openapi.PatternMacPatternMacCounter) (PatternMacPatternMacCounter, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternMacPatternMacCounter) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternMacPatternMacCounter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternMacPatternMacCounter) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternMacPatternMacCounter) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternMacPatternMacCounter) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternMacPatternMacCounter) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternMacPatternMacCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternMacPatternMacCounter) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternMacPatternMacCounter) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternMacPatternMacCounter) Clone() (PatternMacPatternMacCounter, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternMacPatternMacCounter()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternMacPatternMacCounter is mac counter pattern
type PatternMacPatternMacCounter interface {
	Validation
	// Msg marshals PatternMacPatternMacCounter to protobuf object *openapi.PatternMacPatternMacCounter
	// and doesn't set defaults
	Msg() *openapi.PatternMacPatternMacCounter
	// SetMsg unmarshals PatternMacPatternMacCounter from protobuf object *openapi.PatternMacPatternMacCounter
	// and doesn't set defaults
	SetMsg(*openapi.PatternMacPatternMacCounter) PatternMacPatternMacCounter
	// ToProto marshals PatternMacPatternMacCounter to protobuf object *openapi.PatternMacPatternMacCounter
	ToProto() (*openapi.PatternMacPatternMacCounter, error)
	// ToPbText marshals PatternMacPatternMacCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternMacPatternMacCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternMacPatternMacCounter to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternMacPatternMacCounter from protobuf object *openapi.PatternMacPatternMacCounter
	FromProto(msg *openapi.PatternMacPatternMacCounter) (PatternMacPatternMacCounter, error)
	// FromPbText unmarshals PatternMacPatternMacCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternMacPatternMacCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternMacPatternMacCounter from JSON text
	FromJson(value string) error
	// Validate validates PatternMacPatternMacCounter
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternMacPatternMacCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns string, set in PatternMacPatternMacCounter.
	Start() string
	// SetStart assigns string provided by user to PatternMacPatternMacCounter
	SetStart(value string) PatternMacPatternMacCounter
	// HasStart checks if Start has been set in PatternMacPatternMacCounter
	HasStart() bool
	// Step returns string, set in PatternMacPatternMacCounter.
	Step() string
	// SetStep assigns string provided by user to PatternMacPatternMacCounter
	SetStep(value string) PatternMacPatternMacCounter
	// HasStep checks if Step has been set in PatternMacPatternMacCounter
	HasStep() bool
	// Count returns uint32, set in PatternMacPatternMacCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternMacPatternMacCounter
	SetCount(value uint32) PatternMacPatternMacCounter
	// HasCount checks if Count has been set in PatternMacPatternMacCounter
	HasCount() bool
}

// description is TBD
// Start returns a string
func (obj *patternMacPatternMacCounter) Start() string {

	return *obj.obj.Start

}

// description is TBD
// Start returns a string
func (obj *patternMacPatternMacCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the string value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetStart(value string) PatternMacPatternMacCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a string
func (obj *patternMacPatternMacCounter) Step() string {

	return *obj.obj.Step

}

// description is TBD
// Step returns a string
func (obj *patternMacPatternMacCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the string value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetStep(value string) PatternMacPatternMacCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternMacCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternMacPatternMacCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternMacPatternMacCounter object
func (obj *patternMacPatternMacCounter) SetCount(value uint32) PatternMacPatternMacCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternMacPatternMacCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		err := obj.validateMac(obj.Start())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Start"))
		}

	}

	if obj.obj.Step != nil {

		err := obj.validateMac(obj.Step())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on PatternMacPatternMacCounter.Step"))
		}

	}

}

func (obj *patternMacPatternMacCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart("00:00:00:00:00:00")
	}
	if obj.obj.Step == nil {
		obj.SetStep("00:00:00:00:00:01")
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

// ***** PatternIntegerPatternIntegerCounter *****
type patternIntegerPatternIntegerCounter struct {
	validation
	obj *openapi.PatternIntegerPatternIntegerCounter
}

func NewPatternIntegerPatternIntegerCounter() PatternIntegerPatternIntegerCounter {
	obj := patternIntegerPatternIntegerCounter{obj: &openapi.PatternIntegerPatternIntegerCounter{}}
	obj.setDefault()
	return &obj
}

func (obj *patternIntegerPatternIntegerCounter) Msg() *openapi.PatternIntegerPatternIntegerCounter {
	return obj.obj
}

func (obj *patternIntegerPatternIntegerCounter) SetMsg(msg *openapi.PatternIntegerPatternIntegerCounter) PatternIntegerPatternIntegerCounter {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *patternIntegerPatternIntegerCounter) ToProto() (*openapi.PatternIntegerPatternIntegerCounter, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *patternIntegerPatternIntegerCounter) FromProto(msg *openapi.PatternIntegerPatternIntegerCounter) (PatternIntegerPatternIntegerCounter, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *patternIntegerPatternIntegerCounter) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *patternIntegerPatternIntegerCounter) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *patternIntegerPatternIntegerCounter) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIntegerPatternIntegerCounter) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *patternIntegerPatternIntegerCounter) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *patternIntegerPatternIntegerCounter) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *patternIntegerPatternIntegerCounter) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *patternIntegerPatternIntegerCounter) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *patternIntegerPatternIntegerCounter) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *patternIntegerPatternIntegerCounter) Clone() (PatternIntegerPatternIntegerCounter, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewPatternIntegerPatternIntegerCounter()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// PatternIntegerPatternIntegerCounter is integer counter pattern
type PatternIntegerPatternIntegerCounter interface {
	Validation
	// Msg marshals PatternIntegerPatternIntegerCounter to protobuf object *openapi.PatternIntegerPatternIntegerCounter
	// and doesn't set defaults
	Msg() *openapi.PatternIntegerPatternIntegerCounter
	// SetMsg unmarshals PatternIntegerPatternIntegerCounter from protobuf object *openapi.PatternIntegerPatternIntegerCounter
	// and doesn't set defaults
	SetMsg(*openapi.PatternIntegerPatternIntegerCounter) PatternIntegerPatternIntegerCounter
	// ToProto marshals PatternIntegerPatternIntegerCounter to protobuf object *openapi.PatternIntegerPatternIntegerCounter
	ToProto() (*openapi.PatternIntegerPatternIntegerCounter, error)
	// ToPbText marshals PatternIntegerPatternIntegerCounter to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals PatternIntegerPatternIntegerCounter to YAML text
	ToYaml() (string, error)
	// ToJson marshals PatternIntegerPatternIntegerCounter to JSON text
	ToJson() (string, error)
	// FromProto unmarshals PatternIntegerPatternIntegerCounter from protobuf object *openapi.PatternIntegerPatternIntegerCounter
	FromProto(msg *openapi.PatternIntegerPatternIntegerCounter) (PatternIntegerPatternIntegerCounter, error)
	// FromPbText unmarshals PatternIntegerPatternIntegerCounter from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals PatternIntegerPatternIntegerCounter from YAML text
	FromYaml(value string) error
	// FromJson unmarshals PatternIntegerPatternIntegerCounter from JSON text
	FromJson(value string) error
	// Validate validates PatternIntegerPatternIntegerCounter
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (PatternIntegerPatternIntegerCounter, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Start returns uint32, set in PatternIntegerPatternIntegerCounter.
	Start() uint32
	// SetStart assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetStart(value uint32) PatternIntegerPatternIntegerCounter
	// HasStart checks if Start has been set in PatternIntegerPatternIntegerCounter
	HasStart() bool
	// Step returns uint32, set in PatternIntegerPatternIntegerCounter.
	Step() uint32
	// SetStep assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetStep(value uint32) PatternIntegerPatternIntegerCounter
	// HasStep checks if Step has been set in PatternIntegerPatternIntegerCounter
	HasStep() bool
	// Count returns uint32, set in PatternIntegerPatternIntegerCounter.
	Count() uint32
	// SetCount assigns uint32 provided by user to PatternIntegerPatternIntegerCounter
	SetCount(value uint32) PatternIntegerPatternIntegerCounter
	// HasCount checks if Count has been set in PatternIntegerPatternIntegerCounter
	HasCount() bool
}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Start() uint32 {

	return *obj.obj.Start

}

// description is TBD
// Start returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasStart() bool {
	return obj.obj.Start != nil
}

// description is TBD
// SetStart sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetStart(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Start = &value
	return obj
}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Step() uint32 {

	return *obj.obj.Step

}

// description is TBD
// Step returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasStep() bool {
	return obj.obj.Step != nil
}

// description is TBD
// SetStep sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetStep(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Step = &value
	return obj
}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternIntegerCounter) Count() uint32 {

	return *obj.obj.Count

}

// description is TBD
// Count returns a uint32
func (obj *patternIntegerPatternIntegerCounter) HasCount() bool {
	return obj.obj.Count != nil
}

// description is TBD
// SetCount sets the uint32 value in the PatternIntegerPatternIntegerCounter object
func (obj *patternIntegerPatternIntegerCounter) SetCount(value uint32) PatternIntegerPatternIntegerCounter {

	obj.obj.Count = &value
	return obj
}

func (obj *patternIntegerPatternIntegerCounter) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Start != nil {

		if *obj.obj.Start > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Start <= 255 but Got %d", *obj.obj.Start))
		}

	}

	if obj.obj.Step != nil {

		if *obj.obj.Step > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Step <= 255 but Got %d", *obj.obj.Step))
		}

	}

	if obj.obj.Count != nil {

		if *obj.obj.Count > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= PatternIntegerPatternIntegerCounter.Count <= 255 but Got %d", *obj.obj.Count))
		}

	}

}

func (obj *patternIntegerPatternIntegerCounter) setDefault() {
	if obj.obj.Start == nil {
		obj.SetStart(0)
	}
	if obj.obj.Step == nil {
		obj.SetStep(1)
	}
	if obj.obj.Count == nil {
		obj.SetCount(1)
	}

}

// ***** RequiredChoiceIntermeLeaf *****
type requiredChoiceIntermeLeaf struct {
	validation
	obj *openapi.RequiredChoiceIntermeLeaf
}

func NewRequiredChoiceIntermeLeaf() RequiredChoiceIntermeLeaf {
	obj := requiredChoiceIntermeLeaf{obj: &openapi.RequiredChoiceIntermeLeaf{}}
	obj.setDefault()
	return &obj
}

func (obj *requiredChoiceIntermeLeaf) Msg() *openapi.RequiredChoiceIntermeLeaf {
	return obj.obj
}

func (obj *requiredChoiceIntermeLeaf) SetMsg(msg *openapi.RequiredChoiceIntermeLeaf) RequiredChoiceIntermeLeaf {

	proto.Merge(obj.obj, msg)
	return obj
}

func (obj *requiredChoiceIntermeLeaf) ToProto() (*openapi.RequiredChoiceIntermeLeaf, error) {
	err := obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return obj.Msg(), nil
}

func (obj *requiredChoiceIntermeLeaf) FromProto(msg *openapi.RequiredChoiceIntermeLeaf) (RequiredChoiceIntermeLeaf, error) {
	newObj := obj.SetMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (obj *requiredChoiceIntermeLeaf) ToPbText() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	protoMarshal, err := proto.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(protoMarshal), nil
}

func (obj *requiredChoiceIntermeLeaf) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), obj.Msg())
	if retObj != nil {
		return retObj
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (obj *requiredChoiceIntermeLeaf) ToYaml() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	data, err = yaml.JSONToYAML(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceIntermeLeaf) FromYaml(value string) error {
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
	uError := opts.Unmarshal([]byte(data), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (obj *requiredChoiceIntermeLeaf) ToJson() (string, error) {
	vErr := obj.validateToAndFrom()
	if vErr != nil {
		return "", vErr
	}
	opts := protojson.MarshalOptions{
		UseProtoNames:   true,
		AllowPartial:    true,
		EmitUnpopulated: false,
		Indent:          "  ",
	}
	data, err := opts.Marshal(obj.Msg())
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (obj *requiredChoiceIntermeLeaf) FromJson(value string) error {
	opts := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: false,
	}
	if value == "" {
		value = "{}"
	}
	uError := opts.Unmarshal([]byte(value), obj.Msg())
	if uError != nil {
		return fmt.Errorf("unmarshal error %s", strings.Replace(
			uError.Error(), "\u00a0", " ", -1)[7:])
	}

	err := obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *requiredChoiceIntermeLeaf) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermeLeaf) Validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *requiredChoiceIntermeLeaf) String() string {
	str, err := obj.ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *requiredChoiceIntermeLeaf) Clone() (RequiredChoiceIntermeLeaf, error) {
	vErr := obj.Validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRequiredChoiceIntermeLeaf()
	data, err := proto.Marshal(obj.Msg())
	if err != nil {
		return nil, err
	}
	pbErr := proto.Unmarshal(data, newObj.Msg())
	if pbErr != nil {
		return nil, pbErr
	}
	return newObj, nil
}

// RequiredChoiceIntermeLeaf is description is TBD
type RequiredChoiceIntermeLeaf interface {
	Validation
	// Msg marshals RequiredChoiceIntermeLeaf to protobuf object *openapi.RequiredChoiceIntermeLeaf
	// and doesn't set defaults
	Msg() *openapi.RequiredChoiceIntermeLeaf
	// SetMsg unmarshals RequiredChoiceIntermeLeaf from protobuf object *openapi.RequiredChoiceIntermeLeaf
	// and doesn't set defaults
	SetMsg(*openapi.RequiredChoiceIntermeLeaf) RequiredChoiceIntermeLeaf
	// ToProto marshals RequiredChoiceIntermeLeaf to protobuf object *openapi.RequiredChoiceIntermeLeaf
	ToProto() (*openapi.RequiredChoiceIntermeLeaf, error)
	// ToPbText marshals RequiredChoiceIntermeLeaf to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RequiredChoiceIntermeLeaf to YAML text
	ToYaml() (string, error)
	// ToJson marshals RequiredChoiceIntermeLeaf to JSON text
	ToJson() (string, error)
	// FromProto unmarshals RequiredChoiceIntermeLeaf from protobuf object *openapi.RequiredChoiceIntermeLeaf
	FromProto(msg *openapi.RequiredChoiceIntermeLeaf) (RequiredChoiceIntermeLeaf, error)
	// FromPbText unmarshals RequiredChoiceIntermeLeaf from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RequiredChoiceIntermeLeaf from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RequiredChoiceIntermeLeaf from JSON text
	FromJson(value string) error
	// Validate validates RequiredChoiceIntermeLeaf
	Validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RequiredChoiceIntermeLeaf, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in RequiredChoiceIntermeLeaf.
	Name() string
	// SetName assigns string provided by user to RequiredChoiceIntermeLeaf
	SetName(value string) RequiredChoiceIntermeLeaf
	// HasName checks if Name has been set in RequiredChoiceIntermeLeaf
	HasName() bool
}

// description is TBD
// Name returns a string
func (obj *requiredChoiceIntermeLeaf) Name() string {

	return *obj.obj.Name

}

// description is TBD
// Name returns a string
func (obj *requiredChoiceIntermeLeaf) HasName() bool {
	return obj.obj.Name != nil
}

// description is TBD
// SetName sets the string value in the RequiredChoiceIntermeLeaf object
func (obj *requiredChoiceIntermeLeaf) SetName(value string) RequiredChoiceIntermeLeaf {

	obj.obj.Name = &value
	return obj
}

func (obj *requiredChoiceIntermeLeaf) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *requiredChoiceIntermeLeaf) setDefault() {

}
