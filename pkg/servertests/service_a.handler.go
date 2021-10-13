package test

import (
	"io/ioutil"
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
	result.StatusCode200().SetMessage("from GetRootResponse")
	return result
}
func (h *apiTestHandler) PostRootResponse(r *http.Request) openapiart.PostRootResponseResponse {
	var item openapiart.ApiTestInputBody = nil
	if r.Body != nil {
		body, _ := ioutil.ReadAll(r.Body)
		if body != nil {
			item = openapiart.NewApiTestInputBody()
			item.FromJson(string(body))
		}
	}
	result := openapiart.NewPostRootResponseResponse()
	if item != nil {
		result.StatusCode200().SetMessage(item.SomeString())
		return result
	}
	result.StatusCode500().SetMessage("missing input")
	return result
}
func (h *apiTestHandler) DummyResponseTest(r *http.Request) openapiart.DummyResponseTestResponse {
	result := openapiart.NewDummyResponseTestResponse()
	result.SetStatusCode200("this is a string response")
	return result
}
