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
	controller interfaces.ServiceAbcController
}

func NewServiceBHandler() interfaces.ServiceAbcHandler {
	handler := new(serviceBHandler)
	handler.controller = controllers.NewHttpServiceAbcController(handler)
	return handler
}

func (h *serviceBHandler) GetController() interfaces.ServiceAbcController {
	return h.controller
}
func (h *serviceBHandler) GetAllItems(r *http.Request) openapiart.GetAllItemsResponse {
	items := h.getItems()
	result := openapiart.NewGetAllItemsResponse()
	result.ServiceAbcItemList().Items().Append(items...)
	return result
}
func (h *serviceBHandler) GetSingleItem(r *http.Request) openapiart.GetSingleItemResponse {
	vars := mux.Vars(r)
	id := vars[interfaces.ServiceAbcItemId]
	items := h.getItems()
	var item openapiart.ServiceAbcItem
	for _, i := range items {
		if i.SomeId() == id {
			item = i
			break
		}
	}
	result := openapiart.NewGetSingleItemResponse()
	if item != nil {
		result.SetServiceAbcItem(item)
	} else {
		fmt.Println("commenting out for now")
		// result.StatusCode400().SetMessage(fmt.Sprintf("not found: id '%s'", id))
	}
	return result
}
func (h *serviceBHandler) GetSingleItemLevel2(r *http.Request) openapiart.GetSingleItemLevel2Response {
	vars := mux.Vars(r)
	id1 := vars[interfaces.ServiceAbcItemId]
	id2 := vars[interfaces.ServiceAbcLevel2]
	result := openapiart.NewGetSingleItemLevel2Response()
	result.ServiceAbcItem().SetPathId(id1).SetLevel2(id2)
	return result
}

func (h *serviceBHandler) getItems() []openapiart.ServiceAbcItem {
	item1 := openapiart.NewServiceAbcItem()
	item1.SetSomeId("1")
	item1.SetSomeString("item 1")
	item2 := openapiart.NewServiceAbcItem()
	item2.SetSomeId("2")
	item2.SetSomeString("item 2")
	return []openapiart.ServiceAbcItem{
		item1,
		item2,
	}
}
