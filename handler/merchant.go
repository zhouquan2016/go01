package handler

import (
	"go01/api"
	"go01/service"
	"go01/util"
	"net/http"
)

type merchantHandler struct {
}

var MerchantHandler = &merchantHandler{}

func (handler *merchantHandler) Handles() []RequestHandler {
	return []RequestHandler{
		{
			Path:         "/create",
			Method:       http.MethodPost,
			Handler:      handler.create,
			ResponseBody: true,
		},
	}
}

func (handler *merchantHandler) BasePath() string {
	return "/merchant"
}

func (handler *merchantHandler) create(_ http.ResponseWriter, request *http.Request) interface{} {
	query := &api.MerchantCreateQuery{}
	util.Body2Json(request, query)
	return service.MerchantService.Create(query)
}
