package api

//支付基本参数
type MerchantQuery struct {
	//商户id
	MerchantId int64 `json:"merchant_id"`
	//请求密令(商户自己持有，切勿泄露给第三方，以免给自己的资产造成损失)
	SecretKey string `json:"secret_key"`
}

//交易请求
type PayQuery struct {
	MerchantQuery
	//交易id
	TradeId string `json:"trade_id"`
	//支付的金额(分)
	PayMoney int64
	//交易日期
	TradeDate string
}

//创建商户请求
type MerchantCreateQuery struct {
	//营业执照号码
	BusinessNo string `json:"business_no"`
	//公司名称
	CompanyName string `json:"company_name"`
	//姓名
	UserName string `json:"user_name"`
	//手机号码
	PhoneNo string `json:"phone_no"`
	//身份证号码
	IdentifyNo string `json:"identify_no"`
}

//加签的请求参数
type SignQuery struct {
	//签文
	SignData string `json:"sign_data"`
	//明文
	PlanData string `json:"plan_data"`
}
