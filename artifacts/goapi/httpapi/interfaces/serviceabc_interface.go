// This file is autogenerated. Do not modify
package interfaces

import (
	"net/http"

	openapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/open-traffic-generator/goapi/pkg/httpapi"
)

const (
	ServiceAbcItemId = "item_id"
	ServiceAbcLevel2 = "level_2"
)

type ServiceAbcController interface {
	Routes() []httpapi.Route
	GetAllItems(w http.ResponseWriter, r *http.Request)
	GetSingleItem(w http.ResponseWriter, r *http.Request)
	GetSingleItemLevel2(w http.ResponseWriter, r *http.Request)
}

type ServiceAbcHandler interface {
	GetController() ServiceAbcController
	/*
		GetAllItems: GET /api/serviceb
		Description: return list of some items
	*/
	GetAllItems(r *http.Request) (openapi.GetAllItemsResponse, error)
	/*
		GetSingleItem: GET /api/serviceb/{item_id}
		Description: return single item
	*/
	GetSingleItem(r *http.Request) (openapi.GetSingleItemResponse, error)
	/*
		GetSingleItemLevel2: GET /api/serviceb/{item_id}/{level_2}
		Description: return single item
	*/
	GetSingleItemLevel2(r *http.Request) (openapi.GetSingleItemLevel2Response, error)
}