package api

//商户详情
type MerchantCreateVO struct {
	//商户id
	MerchantId int64 `json:"merchant_id"`
	//请求密令(商户自己持有，切勿泄露给第三方，以免给自己的资产造成损失)
	SecretKey string `json:"secret_key"`
	//私钥
	PrivateKey string `json:"private_key"`
	//公钥
	PublicKey string `json:"public_key"`
	//过期日期
	ExpireTime string `json:"expire_time"`
}
