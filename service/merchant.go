package service

import (
	"go01/api"
	"go01/dao"
	"go01/security"
	"go01/util"
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
		MerchantId: 0,
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
