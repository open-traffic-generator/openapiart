package test

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/open-traffic-generator/openapiart/pkg/httpapi/controllers"
	"github.com/open-traffic-generator/openapiart/pkg/httpapi/interfaces"
)

type serviceBHandler struct {
	controller interfaces.ServiceBController
}

func NewServiceBHandler() interfaces.ServiceBHandler {
	handler := new(serviceBHandler)
	handler.controller = controllers.NewHttpServiceBController(handler)
	return handler
}

func (h *serviceBHandler) GetController() interfaces.ServiceBController {
	return h.controller
}
func (h *serviceBHandler) GetAllItems(r *http.Request) openapiart.GetAllItemsResponse {
	items := h.getItems()
	result := openapiart.NewGetAllItemsResponse()
	result.StatusCode200().Items().Append(items...)
	return result
}
func (h *serviceBHandler) GetSingleItem(r *http.Request) openapiart.GetSingleItemResponse {
	vars := mux.Vars(r)
	id := vars[interfaces.ServiceBItemId]
	items := h.getItems()
	var item openapiart.ServiceBItem
	for _, i := range items {
		if i.SomeId() == id {
			item = i
			break
		}
	}
	result := openapiart.NewGetSingleItemResponse()
	if item != nil {
		result.SetStatusCode200(item)
	} else {
		result.StatusCode400().SetMessage(fmt.Sprintf("not found: id '%s'", id))
	}
	return result
}
func (h *serviceBHandler) GetSingleItemLevel2(r *http.Request) openapiart.GetSingleItemLevel2Response {
	vars := mux.Vars(r)
	id1 := vars[interfaces.ServiceBItemId]
	id2 := vars[interfaces.ServiceBLevel2]
	result := openapiart.NewGetSingleItemLevel2Response()
	result.StatusCode200().SetPathId(id1).SetLevel2(id2)
	return result
}

func (h *serviceBHandler) getItems() []openapiart.ServiceBItem {
	item1 := openapiart.NewServiceBItem()
	item1.SetSomeId("1")
	item1.SetSomeString("item 1")
	item2 := openapiart.NewServiceBItem()
	item2.SetSomeId("2")
	item2.SetSomeString("item 2")
	return []openapiart.ServiceBItem{
		item1,
		item2,
	}
}
