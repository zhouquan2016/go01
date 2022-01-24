package service

import (
	"encoding/json"
	"go01/api"
	"go01/config"
	"go01/dao"
	"go01/security"
	"go01/util"
	"net/http"
	"strings"
	"time"
)

type merchantService struct {
}

var MerchantService = &merchantService{}

func (*merchantService) Create(query *api.MerchantCreateQuery) (vo *api.MerchantCreateVO) {
	query = &api.MerchantCreateQuery{
		BusinessNo:  strings.TrimSpace(query.BusinessNo),
		CompanyName: strings.TrimSpace(query.CompanyName),
		UserName:    strings.TrimSpace(query.UserName),
		PhoneNo:     strings.TrimSpace(query.PhoneNo),
		IdentifyNo:  strings.TrimSpace(query.IdentifyNo),
	}
	checkCreateQuery(query)

	merchant := dao.Merchant{
		BusinessNo:  query.BusinessNo,
		CompanyName: query.CompanyName,
		UserName:    query.UserName,
		PhoneNo:     query.PhoneNo,
		IdentifyNo:  query.IdentifyNo,
		Disabled:    false,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}

	privateKey, publicKey, err := security.GenRsaKeys(2048)
	util.AssertError(err, "生成ras异常")
	merchatSecret := dao.MerchantSecret{
		SecretKey:  security.Md5(privateKey),
		PublicKey:  publicKey,
		ExpireTime: time.Now().AddDate(1, 0, 0),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	dao.MerchantDao.Create(&merchant, &merchatSecret)
	vo = &api.MerchantCreateVO{
		MerchantId: merchant.Id,
		SecretKey:  merchatSecret.SecretKey,
		PrivateKey: privateKey,
		PublicKey:  merchatSecret.PublicKey,
		ExpireTime: time.Now().AddDate(1, 0, 0).Format(api.DateTimeFormate),
	}
	return vo
}

func checkCreateQuery(query *api.MerchantCreateQuery) {
	util.ValidateError(query != nil, "参数为空")
	util.ValidateError(query.BusinessNo != "", "营业执照号码不能为空")
	util.ValidateError(query.CompanyName != "", "公司名称不能为空")
	util.ValidateError(query.UserName != "", "姓名不能为空")
	util.ValidateError(query.PhoneNo != "", "手机号不能为空")
	util.ValidateError(util.RegexMatch("^1\\d{10}$", query.PhoneNo), "手机号不合法")
	util.ValidateError(util.RegexMatch("^\\d{16}|\\d{18}$", query.IdentifyNo), "身份证号不合法")
	old := dao.MerchantDao.GetByBusssionNo(nil, query.BusinessNo)
	util.ValidateError(old == nil, "营业执照号码已经注册过商户了")
}

//
func (*merchantService) VerifySignAndParse(request *http.Request, val interface{}) {

	signQuery := new(api.SignQuery)
	util.Body2Json(request, signQuery)
	util.ValidateError(signQuery.EncryptData != "", "encrypt_data不能为空")

	decodeBytes, err := security.DecryptWithPrivateKey(config.Config.PrivateKey, signQuery.EncryptData)
	util.AssertError(err, "encrypt_data解密失败")

	err = json.Unmarshal(decodeBytes, val)
	util.AssertError(err, "encrypt_data转换json异常")
	//商户id和密钥
	merchantQuery := new(api.MerchantQuery)
	err = json.Unmarshal(decodeBytes, merchantQuery)
	util.AssertError(err, "从encrypt_data解码商户信息失败")

	checkMerchant(merchantQuery.MerchantId)
	checkSecret(signQuery.SignData, decodeBytes, merchantQuery)
}

func checkSecret(signData string, decodeBytes []byte, query *api.MerchantQuery) {
	merchantSecret := dao.MerchantSecretDao.GetByMerchantId(query.MerchantId)
	util.ValidateError(merchantSecret != nil, "商户的公钥信息未找到")
	util.ValidateError(merchantSecret.SecretKey == query.SecretKey, "商户密钥不对")
	util.ValidateError(merchantSecret.ExpireTime.After(time.Now()), "私钥对已过期")

	valid, err := security.VerifyWithPublicKey(signData, decodeBytes, merchantSecret.PublicKey)
	util.AssertError(err, "验签失败")
	util.ValidateError(valid, "验签失败")
}

func checkMerchant(merchantId int64) {
	util.ValidateError(merchantId > 0, "商户不存在")

	merchant := dao.MerchantDao.GetById(merchantId)
	util.ValidateError(merchant != nil, "商户不存在")
	util.ValidateError(!merchant.Disabled, "商户已被禁用")

}
