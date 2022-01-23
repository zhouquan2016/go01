package handler

import (
	"go01/security"
	"go01/util"
	"net/http"
	"strconv"
	"time"
)

type GenHandler struct {
}

func (genHandler *GenHandler) Handles() []RequestHandler {
	return []RequestHandler{
		{
			Path:         "/rsa",
			Method:       http.MethodGet,
			Handler:      genRsa,
			ResponseBody: true,
		},
	}
}

func (genHandler *GenHandler) BasePath() string {
	return "/gen"
}

type RsaKey struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	ExpireAt   string `json:"expireAt"`
}

func genRsa(_ http.ResponseWriter, request *http.Request) interface{} {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	parseInt, err := strconv.ParseInt(request.FormValue("bits"), 10, 32)
	if err != nil {
		panic(util.NewServiceError("参数不合法"))
	}
	privateKey, publicKey, err := security.GenRsaKeys(int(parseInt))
	if err != nil {
		panic(util.NewServiceError("生成秘钥失败"))
	}
	return &RsaKey{privateKey, publicKey, time.Now().AddDate(1, 0, 0).Format("2006-01-02 15:04:05")}
}
