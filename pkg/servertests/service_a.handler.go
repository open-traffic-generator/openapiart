package test

import (
	"net/http"

	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/open-traffic-generator/goapi/pkg/httpapi/controllers"
	"github.com/open-traffic-generator/goapi/pkg/httpapi/interfaces"
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
func (h *apiTestHandler) GetRootResponse(r *http.Request) (goapi.GetRootResponseResponse, error) {
	result := goapi.NewGetRootResponseResponse()
	result.CommonResponseSuccess().SetMessage("from GetRootResponse")
	return result, nil
}

func (h *apiTestHandler) PostRootResponse(requestbody goapi.ApiTestInputBody, r *http.Request) (goapi.PostRootResponseResponse, error) {
	result := goapi.NewPostRootResponseResponse()
	if requestbody != nil {
		result.CommonResponseSuccess().SetMessage(requestbody.SomeString())
		return result, nil
	}
	// result.StatusCode500().SetMessage("missing input")
	return result, nil
}
func (h *apiTestHandler) DummyResponseTest(r *http.Request) (goapi.DummyResponseTestResponse, error) {
	result := goapi.NewDummyResponseTestResponse()
	result.SetResponseString("this is a string response")
	return result, nil
}
