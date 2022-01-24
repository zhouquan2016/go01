package handler

import (
	"go01/security"
	"go01/util"
	"net/http"
	"strconv"
	"time"
)

type genHandler struct {
}

var GenHandler = &genHandler{}

func (genHandler *genHandler) Handles() []RequestHandler {
	return []RequestHandler{
		{
			Path:         "/rsa",
			Method:       http.MethodGet,
			Handler:      genHandler.genRsa,
			ResponseBody: true,
		},
	}
}

func (genHandler *genHandler) BasePath() string {
	return "/gen"
}

type RsaKey struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	ExpireAt   string `json:"expireAt"`
}

func (genHandler *genHandler) genRsa(_ http.ResponseWriter, request *http.Request) interface{} {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}
	parseInt, err := strconv.ParseInt(request.FormValue("bits"), 10, 32)
	util.AssertError(err, "参数不合法")
	privateKey, publicKey, err := security.GenRsaKeys(int(parseInt))
	util.AssertError(err, "生成秘钥失败")
	return &RsaKey{privateKey, publicKey, time.Now().AddDate(1, 0, 0).Format("2006-01-02 15:04:05")}
}
