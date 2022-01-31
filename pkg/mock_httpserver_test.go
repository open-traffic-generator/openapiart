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
		serverLocation: "127.0.0.1:50051",
	}
)

// 	Add route and strat HTTP server

func StartMockHttpServer() {
	bundlerHandler := NewBundlerHandler()
	metricsHandler := NewMetricsHandler()
	controllers := []httpapi.HttpController{
		bundlerHandler.GetController(),
		metricsHandler.GetController(),
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

func NewBundlerHandler() interfaces.BundlerHandler {
	handler := new(bundlerHandler)
	handler.controller = controllers.NewHttpBundlerController(handler)
	return handler
}

func (h *bundlerHandler) GetController() interfaces.BundlerController {
	return h.controller
}

func (h *bundlerHandler) SetConfig(rbody openapiart.PrefixConfig, r *http.Request) openapiart.SetConfigResponse {
	httpServer.Config = rbody
	response := openapiart.NewSetConfigResponse()
	switch httpServer.Config.Response() {
	case openapiart.PrefixConfigResponse.STATUS_200:
		response.SetStatusCode200([]byte("Successful set config operation"))
	case openapiart.PrefixConfigResponse.STATUS_400:
		response.StatusCode400().SetErrors([]string{"A 400 error has occurred"})
	case openapiart.PrefixConfigResponse.STATUS_500:
		response.StatusCode500().SetErrors([]string{"A 500 error has occurred"})
	}
	return response
}

func (h *bundlerHandler) UpdateConfiguration(rbody openapiart.UpdateConfig, r *http.Request) openapiart.UpdateConfigurationResponse {
	response := openapiart.NewUpdateConfigurationResponse()
	data, _ := httpServer.Config.ToJson()
	err := response.StatusCode200().FromJson(data)
	if err != nil {
		log.Print(err.Error())
	}
	return response
}

func (h *bundlerHandler) GetConfig(r *http.Request) openapiart.GetConfigResponse {
	response := openapiart.NewGetConfigResponse()
	response.SetStatusCode200(httpServer.Config)
	return response
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

func (h *metricsHandler) GetMetrics(r *http.Request) openapiart.GetMetricsResponse {
	response := openapiart.NewGetMetricsResponse()
	response.StatusCode200().Ports().Add().SetName("p1").SetTxFrames(2000).SetRxFrames(1777)
	response.StatusCode200().Ports().Add().SetName("p2").SetTxFrames(3000).SetRxFrames(2999)
	return response
}

func (h *metricsHandler) GetWarnings(r *http.Request) openapiart.GetWarningsResponse {
	response := openapiart.NewGetWarningsResponse()
	response.StatusCode200().SetWarnings([]string{"This is your first warning", "Your last warning"})
	return response
}

func (h *metricsHandler) ClearWarnings(r *http.Request) openapiart.ClearWarningsResponse {
	response := openapiart.NewClearWarningsResponse()
	response.SetStatusCode200("success")
	return response
}
