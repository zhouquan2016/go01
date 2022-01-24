package handler

import (
	"go01/api"
	"go01/service"
	"log"
	"net/http"
)

type payHandler struct {
}

var PayHandler = &payHandler{}

func (handler *payHandler) Handles() []RequestHandler {
	return []RequestHandler{
		{
			Path:         "",
			Method:       http.MethodPost,
			Handler:      handler.pay,
			ResponseBody: true,
		},
	}
}

func (handler *payHandler) BasePath() string {
	return "/pay"
}

func (handler *payHandler) pay(_ http.ResponseWriter, request *http.Request) interface{} {
	query := new(api.PayQuery)
	service.MerchantService.VerifySignAndParse(request, query)
	log.Println(query)
	return nil
}
