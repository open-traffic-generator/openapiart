// This file is autogenerated. Do not modify
package interfaces

import (
	"net/http"

	openapiart "github.com/open-traffic-generator/openapiart/pkg"
	"github.com/open-traffic-generator/openapiart/pkg/httpapi"
)

type BundlerController interface {
	Routes() []httpapi.Route
	SetConfig(w http.ResponseWriter, r *http.Request)
	UpdateConfiguration(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)
}

type BundlerHandler interface {
	GetController() BundlerController
	/*
		SetConfig: POST /api/config
		Description: Sets configuration resources.
	*/
	SetConfig(rbody openapiart.PrefixConfig, r *http.Request) (openapiart.SetConfigResponse, error)
	/*
			UpdateConfiguration: PATCH /api/config
			Description: Deprecated: please use post instead

		Sets configuration resources.
	*/
	UpdateConfiguration(rbody openapiart.UpdateConfig, r *http.Request) (openapiart.UpdateConfigurationResponse, error)
	/*
		GetConfig: GET /api/config
		Description: Gets the configuration resources.
	*/
	GetConfig(r *http.Request) (openapiart.GetConfigResponse, error)
}
