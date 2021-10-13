package test

import (
	"localdev/art_go/models"
	"localdev/art_go/pkg/httpapi/controllers"
	"localdev/art_go/pkg/httpapi/interfaces"
	"net/http"
)

type httpServiceBHandler struct {
	controller interfaces.ServiceBController
}

func NewHttpServiceBHandler() interfaces.ServiceBServiceHandler {
	handler := new(httpServiceBHandler)
	handler.controller = controllers.NewHttpServiceBController(handler)
	return handler
}

func (h *httpServiceBHandler) GetController() interfaces.ServiceBController {
	return h.controller
}
func (h *httpServiceBHandler) GetAllItems(r *http.Request) models.GetAllItemsResponse {
	return models.NewGetAllItemsResponse()
}
func (h *httpServiceBHandler) GetSingleItem(r *http.Request) models.GetSingleItemResponse {
	return models.NewGetSingleItemResponse()
}
func (h *httpServiceBHandler) GetSingleItemLevel2(r *http.Request) models.GetSingleItemLevel2Response {
	return models.NewGetSingleItemLevel2Response()
}
