package test

import (
	"github.com/open-traffic-generator/goapi/pkg/httpapi"

	"github.com/gorilla/mux"
)

func setup() *mux.Router {
	// This setup is done on main in the actual app
	handler1 := NewApiTestHandler()
	handler2 := NewServiceBHandler()
	controllers := []httpapi.HttpController{
		handler1.GetController(),
		handler2.GetController(),
	}
	return httpapi.AppendRoutes(nil, controllers...)
}
