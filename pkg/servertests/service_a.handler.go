package test

import (
	"net/http"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/open-traffic-generator/openapiart/pkg/httpapi/controllers"
	"github.com/open-traffic-generator/openapiart/pkg/httpapi/interfaces"
)

type apiTestHandler struct {
	controller interfaces.ApiTestController
}

func NewApiTestHandler() interfaces.ApiTestHandler {
	handler := new(apiTestHandler)
	handler.controller = controllers.NewHttpApiTestController(handler)
	return handler
}

func (h *apiTestHandler) GetController() interfaces.ApiTestController {
	return h.controller
}
func (h *apiTestHandler) GetRootResponse(r *http.Request) openapiart.GetRootResponseResponse {
	result := openapiart.NewGetRootResponseResponse()
	result.CommonResponseSuccess().SetMessage("from GetRootResponse")
	return result
}

func (h *apiTestHandler) PostRootResponse(requestbody openapiart.ApiTestInputBody, r *http.Request) openapiart.PostRootResponseResponse {
	result := openapiart.NewPostRootResponseResponse()
	if requestbody != nil {
		result.CommonResponseSuccess().SetMessage(requestbody.SomeString())
		return result
	}
	// result.StatusCode500().SetMessage("missing input")
	return result
}
func (h *apiTestHandler) DummyResponseTest(r *http.Request) openapiart.DummyResponseTestResponse {
	result := openapiart.NewDummyResponseTestResponse()
	result.SetResponseString("this is a string response")
	return result
}
