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

	openapi "github.com/open-traffic-generator/goapi/pkg/openapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// function related to error handling
func FromError(err error) (Error, bool) {
	if rErr, ok := err.(Error); ok {
		return rErr, true
	}

	rErr := NewError()
	if err := rErr.Unmarshal().FromJson(err.Error()); err == nil {
		return rErr, true
	}

	return fromGrpcError(err)
}

func setResponseErr(obj Error, code int32, message string) {
	errors := []string{}
	errors = append(errors, message)
	obj.msg().Code = &code
	obj.msg().Errors = errors
}

// parses and return errors for grpc response
func fromGrpcError(err error) (Error, bool) {
	st, ok := status.FromError(err)
	if ok {
		rErr := NewError()
		if err := rErr.Unmarshal().FromJson(st.Message()); err == nil {
			var code = int32(st.Code())
			rErr.msg().Code = &code
			return rErr, true
		}

		setResponseErr(rErr, int32(st.Code()), st.Message())
		return rErr, true
	}

	return nil, false
}

// parses and return errors for http responses
func fromHttpError(statusCode int, body []byte) Error {
	rErr := NewError()
	bStr := string(body)
	if err := rErr.Unmarshal().FromJson(bStr); err == nil {
		return rErr
	}

	setResponseErr(rErr, int32(statusCode), bStr)

	return rErr
}

type versionMeta struct {
	checkVersion  bool
	localVersion  Version
	remoteVersion Version
	checkError    error
}
type goapiApi struct {
	apiSt
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
func NewApi() Api {
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
type Api interface {
	api
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
	// GetTestConfig get the new restructured unit test config.
	GetTestConfig() (TestConfig, error)
	// SetTestConfig sets the new restructured unit test configuration.
	SetTestConfig(testConfig TestConfig) ([]byte, error)
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

	if err := prefixConfig.validate(); err != nil {
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
	request := openapi.SetConfigRequest{PrefixConfig: prefixConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
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

	if err := updateConfig.validate(); err != nil {
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
	request := openapi.UpdateConfigurationRequest{UpdateConfig: updateConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.UpdateConfiguration(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewPrefixConfig()
	if resp.GetPrefixConfig() != nil {
		return ret.setMsg(resp.GetPrefixConfig()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewPrefixConfig()
	if resp.GetPrefixConfig() != nil {
		return ret.setMsg(resp.GetPrefixConfig()), nil
	}

	return ret, nil
}

func (api *goapiApi) GetMetrics(metricsRequest MetricsRequest) (Metrics, error) {

	if err := metricsRequest.validate(); err != nil {
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
	request := openapi.GetMetricsRequest{MetricsRequest: metricsRequest.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetMetrics(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewMetrics()
	if resp.GetMetrics() != nil {
		return ret.setMsg(resp.GetMetrics()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewWarningDetails()
	if resp.GetWarningDetails() != nil {
		return ret.setMsg(resp.GetWarningDetails()), nil
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
		if er, ok := fromGrpcError(err); ok {
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

func (api *goapiApi) GetTestConfig() (TestConfig, error) {

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpGetTestConfig()
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := emptypb.Empty{}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.GetTestConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewTestConfig()
	if resp.GetTestConfig() != nil {
		return ret.setMsg(resp.GetTestConfig()), nil
	}

	return ret, nil
}

func (api *goapiApi) SetTestConfig(testConfig TestConfig) ([]byte, error) {

	if err := testConfig.validate(); err != nil {
		return nil, err
	}

	if err := api.checkLocalRemoteVersionCompatibilityOnce(); err != nil {
		return nil, err
	}
	if api.hasHttpTransport() {
		return api.httpSetTestConfig(testConfig)
	}
	if err := api.grpcConnect(); err != nil {
		return nil, err
	}
	request := openapi.SetTestConfigRequest{TestConfig: testConfig.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.SetTestConfig(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	if resp.ResponseBytes != nil {
		return resp.ResponseBytes, nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewCommonResponseSuccess()
	if resp.GetCommonResponseSuccess() != nil {
		return ret.setMsg(resp.GetCommonResponseSuccess()), nil
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
		if er, ok := fromGrpcError(err); ok {
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

	if err := apiTestInputBody.validate(); err != nil {
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
	request := openapi.PostRootResponseRequest{ApiTestInputBody: apiTestInputBody.msg()}
	ctx, cancelFunc := context.WithTimeout(context.Background(), api.grpc.requestTimeout)
	defer cancelFunc()
	resp, err := api.grpcClient.PostRootResponse(ctx, &request)
	if err != nil {
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewCommonResponseSuccess()
	if resp.GetCommonResponseSuccess() != nil {
		return ret.setMsg(resp.GetCommonResponseSuccess()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItemList()
	if resp.GetServiceAbcItemList() != nil {
		return ret.setMsg(resp.GetServiceAbcItemList()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItem()
	if resp.GetServiceAbcItem() != nil {
		return ret.setMsg(resp.GetServiceAbcItem()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewServiceAbcItem()
	if resp.GetServiceAbcItem() != nil {
		return ret.setMsg(resp.GetServiceAbcItem()), nil
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
		if er, ok := fromGrpcError(err); ok {
			return nil, er
		}
		return nil, err
	}
	ret := NewVersion()
	if resp.GetVersion() != nil {
		return ret.setMsg(resp.GetVersion()), nil
	}

	return ret, nil
}

func (api *goapiApi) httpSetConfig(prefixConfig PrefixConfig) ([]byte, error) {
	prefixConfigJson, err := prefixConfig.Marshal().ToJson()
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
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpUpdateConfiguration(updateConfig UpdateConfig) (PrefixConfig, error) {
	updateConfigJson, err := updateConfig.Marshal().ToJson()
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
		obj := NewUpdateConfigurationResponse().PrefixConfig()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetConfigResponse().PrefixConfig()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetMetrics(metricsRequest MetricsRequest) (Metrics, error) {
	metricsRequestJson, err := metricsRequest.Marshal().ToJson()
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
		obj := NewGetMetricsResponse().Metrics()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetWarningsResponse().WarningDetails()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpGetTestConfig() (TestConfig, error) {
	resp, err := api.httpSendRecv("api/test", "", "GET")
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == 200 {
		obj := NewGetTestConfigResponse().TestConfig()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpSetTestConfig(testConfig TestConfig) ([]byte, error) {
	testConfigJson, err := testConfig.Marshal().ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := api.httpSendRecv("api/test", testConfigJson, "POST")

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
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetRootResponseResponse().CommonResponseSuccess()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}

func (api *goapiApi) httpPostRootResponse(apiTestInputBody ApiTestInputBody) (CommonResponseSuccess, error) {
	apiTestInputBodyJson, err := apiTestInputBody.Marshal().ToJson()
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
		obj := NewPostRootResponseResponse().CommonResponseSuccess()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetAllItemsResponse().ServiceAbcItemList()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetSingleItemResponse().ServiceAbcItem()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetSingleItemLevel2Response().ServiceAbcItem()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
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
		obj := NewGetVersionResponse().Version()
		if err := obj.Unmarshal().FromJson(string(bodyBytes)); err != nil {
			return nil, err
		}
		return obj, nil
	} else {
		return nil, fromHttpError(resp.StatusCode, bodyBytes)
	}
}
