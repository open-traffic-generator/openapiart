// This file is autogenerated. Do not modify
package controllers

import (
	"errors"
	"log"
	"net/http"

	openapi "github.com/open-traffic-generator/goapi/pkg"
	"github.com/open-traffic-generator/goapi/pkg/httpapi"
	"github.com/open-traffic-generator/goapi/pkg/httpapi/interfaces"
)

type capabilitiesController struct {
	handler interfaces.CapabilitiesHandler
}

func NewHttpCapabilitiesController(handler interfaces.CapabilitiesHandler) interfaces.CapabilitiesController {
	return &capabilitiesController{handler}
}

func (ctrl *capabilitiesController) Routes() []httpapi.Route {
	return []httpapi.Route{
		{Path: "/api/capabilities/version", Method: "GET", Name: "GetVersion", Handler: ctrl.GetVersion},
	}
}

/*
GetVersion: GET /api/capabilities/version
Description:
*/
func (ctrl *capabilitiesController) GetVersion(w http.ResponseWriter, r *http.Request) {
	result, err := ctrl.handler.GetVersion(r)
	if err != nil {
		ctrl.responseGetVersionError(w, "internal", err)
		return
	}

	if result.HasVersion() {
		if _, err := httpapi.WriteJSONResponse(w, 200, result.Version().Marshal()); err != nil {
			log.Print(err.Error())
		}
		return
	}
	ctrl.responseGetVersionError(w, "internal", errors.New("Unknown error"))
}

func (ctrl *capabilitiesController) responseGetVersionError(w http.ResponseWriter, errorKind openapi.ErrorKindEnum, rsp_err error) {
	var result openapi.Error
	var statusCode int32
	if errorKind == "validation" {
		statusCode = 400
	} else if errorKind == "internal" {
		statusCode = 500
	}

	if rErr, ok := rsp_err.(openapi.Error); ok {
		result = rErr
	} else {
		result = openapi.NewError()
		err := result.Unmarshal().FromJson(rsp_err.Error())
		if err != nil {
			_ = result.SetCode(statusCode)
			err = result.SetKind(errorKind)
			if err != nil {
				log.Print(err.Error())
			}
			_ = result.SetErrors([]string{rsp_err.Error()})
		}
	}

	if _, err := httpapi.WriteJSONResponse(w, int(result.Code()), result.Marshal()); err != nil {
		log.Print(err.Error())
	}
}