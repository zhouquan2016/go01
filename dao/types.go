package dao

import "time"

type Merchant struct {
	Id int64 `db:"id"`
	//营业执照号码
	BusinessNo string `db:"business_no"`
	//公司名称
	CompanyName string `db:"company_name"`
	//用户姓名
	UserName string `db:"user_name"`
	//手机号码
	PhoneNo string `db:"phone_no"`
	//身份证号码
	IdentifyNo string `db:"identify_no"`
	//是否被禁用
	Disabled bool `db:"disabled"`
	//创建时间
	CreateTime time.Time `db:"create_time"`
	//更新时间
	UpdateTime time.Time `db:"update_time"`
}

//商户密钥信息
type MerchantSecret struct {
	Id int64 `db:"id"`
	//商户id
	MerchantId int64 `db:"merchant_id"`
	//请求密钥
	SecretKey string `db:"secret_key"`
	//公钥
	PublicKey string `db:"public_key"`
	//公私钥过期时间
	ExpireTime time.Time `db:"expire_time"`
	//创建时间
	CreateTime time.Time `db:"create_time"`
	//更新时间
	UpdateTime time.Time `db:"update_time"`
}
