package openapiart_test

import (
	"fmt"
	"log"
	"net/http"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	httpapi "github.com/open-traffic-generator/openapiart/pkg/httpapi"
	controllers "github.com/open-traffic-generator/openapiart/pkg/httpapi/controllers"
	interfaces "github.com/open-traffic-generator/openapiart/pkg/httpapi/interfaces"
)

// 	Common struct

type HttpServer struct {
	serverLocation string
	Location       string
	Config         openapiart.PrefixConfig
}

var (
	httpServer HttpServer = HttpServer{
		serverLocation: "127.0.0.1:8444",
	}
)

// 	Add route and strat HTTP server

func StartMockHttpServer() {
	bundlerHandler := NewBundlerHandler()
	metricsHandler := NewMetricsHandler()
	capabilitiesHandler := NewCapabilitiesHandler()
	controllers := []httpapi.HttpController{
		bundlerHandler.GetController(),
		metricsHandler.GetController(),
		capabilitiesHandler.GetController(),
	}
	router := httpapi.AppendRoutes(nil, controllers...)
	httpServer.Location = fmt.Sprintf("http://%s", httpServer.serverLocation)
	go func() {
		log.Println("Generated Http Server serving incoming HTTP requests on ", httpServer.serverLocation)
		if err := http.ListenAndServe(httpServer.serverLocation, router); err != nil {
			log.Fatal("Generated Http Server failed to serve incoming HTTP request.")
		}
	}()
}

// 	Defined bundler interface

type bundlerHandler struct {
	controller interfaces.BundlerController
}

type capabilitiesHandler struct {
	controller interfaces.CapabilitiesController
}

func NewBundlerHandler() interfaces.BundlerHandler {
	handler := new(bundlerHandler)
	handler.controller = controllers.NewHttpBundlerController(handler)
	return handler
}

func NewCapabilitiesHandler() interfaces.CapabilitiesHandler {
	handler := new(capabilitiesHandler)
	handler.controller = controllers.NewHttpCapabilitiesController(handler)
	return handler
}

func (h *bundlerHandler) GetController() interfaces.BundlerController {
	return h.controller
}

func (h *capabilitiesHandler) GetController() interfaces.CapabilitiesController {
	return h.controller
}

func (h *bundlerHandler) SetConfig(rbody openapiart.PrefixConfig, r *http.Request) (openapiart.SetConfigResponse, error) {
	httpServer.Config = rbody
	response := openapiart.NewSetConfigResponse()
	switch httpServer.Config.Response() {
	case openapiart.PrefixConfigResponse.STATUS_200:
		response.SetResponseBytes([]byte("Successful set config operation"))
	case openapiart.PrefixConfigResponse.STATUS_400:
		return nil, fmt.Errorf("client error !!!!")
	case openapiart.PrefixConfigResponse.STATUS_500:
		err := openapiart.NewError()
		var code int32 = 500
		_ = err.SetCode(code)
		e := err.SetKind("internal")
		fmt.Println(e)
		_ = err.SetErrors([]string{"internal err 1", "internal err 2", "internal err 3"})
		jsonStr, _ := err.Marshal().ToJson()
		return nil, fmt.Errorf(jsonStr)
	}
	return response, nil
}

func (h *bundlerHandler) UpdateConfiguration(rbody openapiart.UpdateConfig, r *http.Request) (openapiart.UpdateConfigurationResponse, error) {
	response := openapiart.NewUpdateConfigurationResponse()
	data, _ := httpServer.Config.Marshal().ToJson()
	err := response.PrefixConfig().Unmarshal().FromJson(data)
	if err != nil {
		log.Print(err.Error())
	}
	return response, nil
}

func (h *bundlerHandler) GetConfig(r *http.Request) (openapiart.GetConfigResponse, error) {
	response := openapiart.NewGetConfigResponse()
	response.SetPrefixConfig(httpServer.Config)
	return response, nil
}

func (h *capabilitiesHandler) GetVersion(r *http.Request) (openapiart.GetVersionResponse, error) {
	response := openapiart.NewGetVersionResponse()
	response.SetVersion(openapiart.NewApi().GetLocalVersion())
	return response, nil
}

// Defined Metrics interface

type metricsHandler struct {
	controller interfaces.MetricsController
}

func NewMetricsHandler() interfaces.MetricsHandler {
	handler := new(metricsHandler)
	handler.controller = controllers.NewHttpMetricsController(handler)
	return handler
}

func (h *metricsHandler) GetController() interfaces.MetricsController {
	return h.controller
}

func (h *metricsHandler) GetMetrics(req openapiart.MetricsRequest, r *http.Request) (openapiart.GetMetricsResponse, error) {
	ch, _ := req.Marshal().ToProto()
	choice := ch.GetChoice().String()
	switch choice {
	case "port":
		response := openapiart.NewGetMetricsResponse()
		response.Metrics().Ports().Add().SetName("p1").SetTxFrames(2000).SetRxFrames(1777)
		response.Metrics().Ports().Add().SetName("p2").SetTxFrames(3000).SetRxFrames(2999)
		return response, nil
	case "flow":
		response := openapiart.NewGetMetricsResponse()
		response.Metrics().Flows().Add().SetName("f1").SetTxFrames(2000).SetRxFrames(1777)
		response.Metrics().Flows().Add().SetName("f2").SetTxFrames(3000).SetRxFrames(2999)
		return response, nil
	default:
		return openapiart.NewGetMetricsResponse(), nil
	}
}

func (h *metricsHandler) GetWarnings(r *http.Request) (openapiart.GetWarningsResponse, error) {
	response := openapiart.NewGetWarningsResponse()
	response.WarningDetails().SetWarnings([]string{"This is your first warning", "Your last warning"})
	return response, nil
}

func (h *metricsHandler) ClearWarnings(r *http.Request) (openapiart.ClearWarningsResponse, error) {
	response := openapiart.NewClearWarningsResponse()
	response.SetResponseString("success")
	return response, nil
}
