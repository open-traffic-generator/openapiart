package openapiart_test

import (
	"log"
	"net/http"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	httpapi "github.com/open-traffic-generator/openapiart/pkg/httpapi"
	controllers "github.com/open-traffic-generator/openapiart/pkg/httpapi/controllers"
	interfaces "github.com/open-traffic-generator/openapiart/pkg/httpapi/interfaces"
)

func StartHttpServer() {
	handler := bundlerHandler{}
	handler.controller = controllers.NewHttpBundlerController(&handler)
	controllers := []httpapi.HttpController{
		handler.GetController(),
	}
	router := httpapi.AppendRoutes(nil, controllers...)
	go func() {
		log.Println("Generated Http Server serving incoming HTTP requests on 127.0.0.1:50071.")
		if err := http.ListenAndServe("127.0.0.1:50071", router); err != nil {
			log.Fatal("Generated Http Server failed to serve incoming HTTP request.")
		}
	}()
}

type bundlerHandler struct {
	controller interfaces.BundlerController
}

func (h *bundlerHandler) GetController() interfaces.BundlerController {
	return h.controller
}

func (h *bundlerHandler) SetConfig(r *http.Request) openapiart.SetConfigResponse {
	result := openapiart.NewSetConfigResponse()
	result.SetStatusCode200([]byte("set config was a success"))
	return result
}

func (h *bundlerHandler) UpdateConfig(r *http.Request) openapiart.UpdateConfigResponse {
	result := openapiart.NewUpdateConfigResponse()
	result.SetStatusCode200(openapiart.NewPrefixConfig())
	return result
}

func (h *bundlerHandler) GetConfig(r *http.Request) openapiart.GetConfigResponse {
	result := openapiart.NewGetConfigResponse()
	result.SetStatusCode200(openapiart.NewPrefixConfig())
	return result
}
