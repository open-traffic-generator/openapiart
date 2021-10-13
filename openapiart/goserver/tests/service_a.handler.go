package test

import (
	"io/ioutil"
	"localdev/art_go/models"
	"localdev/art_go/pkg/httpapi/controllers"
	"localdev/art_go/pkg/httpapi/interfaces"
	"net/http"
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
func (h *apiTestHandler) GetRootResponse(r *http.Request) models.GetRootResponseResponse {
	result := models.NewGetRootResponseResponse()
	result.StatusCode200().SetMessage("from GetRootResponse")
	return result
}
func (h *apiTestHandler) PostRootResponse(r *http.Request) models.PostRootResponseResponse {
	var item models.ApiTestInputBody = nil
	if r.Body != nil {
		body, _ := ioutil.ReadAll(r.Body)
		if body != nil {
			item = models.NewApiTestInputBody()
			item.FromJson(string(body))
		}
	}
	result := models.NewPostRootResponseResponse()
	if item != nil {
		result.StatusCode200().SetMessage(item.SomeString())
		return result
	}
	result.StatusCode500().SetMessage("missing input")
	return result
}
