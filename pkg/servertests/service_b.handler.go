package test

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	goapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/open-traffic-generator/goapi/pkg/httpapi/controllers"
	"github.com/open-traffic-generator/goapi/pkg/httpapi/interfaces"
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
func (h *serviceBHandler) GetAllItems(r *http.Request) (goapi.GetAllItemsResponse, error) {
	items := h.getItems()
	result := goapi.NewGetAllItemsResponse()
	result.ServiceAbcItemList().Items().Append(items...)
	return result, nil
}
func (h *serviceBHandler) GetSingleItem(r *http.Request) (goapi.GetSingleItemResponse, error) {
	vars := mux.Vars(r)
	id := vars[interfaces.ServiceAbcItemId]
	items := h.getItems()
	var item goapi.ServiceAbcItem
	for _, i := range items {
		if i.SomeId() == id {
			item = i
			break
		}
	}
	result := goapi.NewGetSingleItemResponse()
	if item != nil {
		result.SetServiceAbcItem(item)
	} else {
		err := goapi.NewError()
		var code int32 = 500
		_ = err.SetCode(code)
		_ = err.SetErrors([]string{fmt.Sprintf("not found: id '%s'", id)})
		jsonStr, _ := err.Marshal().ToJson()
		return nil, fmt.Errorf(jsonStr)
	}
	return result, nil
}
func (h *serviceBHandler) GetSingleItemLevel2(r *http.Request) (goapi.GetSingleItemLevel2Response, error) {
	vars := mux.Vars(r)
	id1 := vars[interfaces.ServiceAbcItemId]
	id2 := vars[interfaces.ServiceAbcLevel2]
	result := goapi.NewGetSingleItemLevel2Response()
	result.ServiceAbcItem().SetPathId(id1).SetLevel2(id2)
	return result, nil
}

func (h *serviceBHandler) getItems() []goapi.ServiceAbcItem {
	item1 := goapi.NewServiceAbcItem()
	item1.SetSomeId("1")
	item1.SetSomeString("item 1")
	item2 := goapi.NewServiceAbcItem()
	item2.SetSomeId("2")
	item2.SetSomeString("item 2")
	return []goapi.ServiceAbcItem{
		item1,
		item2,
	}
}
