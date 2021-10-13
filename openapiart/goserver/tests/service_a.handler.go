package test

import (
	"localdev/art_go/models"
	"localdev/art_go/pkg/httpapi/controllers"
	"localdev/art_go/pkg/httpapi/interfaces"
	"net/http"
)

type httpApiTestHandler struct {
	controller interfaces.ApiTestController
}

func NewHttpApiTestHandler() interfaces.ApiTestServiceHandler {
	handler := new(httpApiTestHandler)
	handler.controller = controllers.NewHttpApiTestController(handler)
	return handler
}

func (h *httpApiTestHandler) GetController() interfaces.ApiTestController {
	return h.controller
}
func (h *httpApiTestHandler) GetRootResponse(r *http.Request) models.GetRootResponseResponse {
	return models.NewGetRootResponseResponse()
}
func (h *httpApiTestHandler) PostRootResponse(r *http.Request) models.PostRootResponseResponse {
	return models.NewPostRootResponseResponse()
}
