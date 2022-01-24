package dao

import (
	"database/sql"
	"go01/util"
)

type merchantDao struct {
}

var MerchantDao = &merchantDao{}

func (dao *merchantDao) Create(merchant *Merchant, secret *MerchantSecret) {
	tx := OpenTx()
	defer EndTx(tx)

	old := dao.GetByBusssionNo(tx, merchant.BusinessNo)
	util.ValidateError(old == nil, "营业执照号码已经注册过商户")
	resultSet, err := tx.Exec("insert into merchant(business_no, company_name, user_name, phone_no, identify_no, disabled, create_time, update_time) VALUE (?,?,?,?,?, ?, ?,?)",
		merchant.BusinessNo, merchant.CompanyName, merchant.UserName, merchant.PhoneNo, merchant.IdentifyNo, merchant.Disabled, merchant.CreateTime, merchant.UpdateTime)
	util.AssertError(err, "新增商户异常")
	id, err := resultSet.LastInsertId()
	util.AssertError(err, "新增商户id获取异常")
	merchant.Id = id

	MerchantSecretDao.Add(tx, secret)

}

//根据营业执照号码查找,并上读锁
//tx不为空加读锁，为空不加读锁
func (*merchantDao) GetByBusssionNo(tx *sql.Tx, businessNo string) *Merchant {
	querySql := "select business_no, company_name, user_name, phone_no, identify_no, disabled, create_time, update_time from merchant where business_no = ?"
	row := queryRow(tx, querySql, businessNo)
	merchant := new(Merchant)
	err := row.Scan(&merchant.BusinessNo, &merchant.CompanyName, &merchant.UserName, &merchant.PhoneNo, &merchant.IdentifyNo, &merchant.Disabled, &merchant.CreateTime, &merchant.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return merchant
}

func (dao *merchantDao) GetById(id int64) *Merchant {
	querySql := "select business_no, company_name, user_name, phone_no, identify_no, disabled, create_time, update_time from merchant where id = ?"
	row := queryRow(nil, querySql, id)
	return dao.getOne(row)
}

func (dao *merchantDao) getOne(row *sql.Row) *Merchant {
	merchant := new(Merchant)
	err := row.Scan(&merchant.BusinessNo, &merchant.CompanyName, &merchant.UserName, &merchant.PhoneNo, &merchant.IdentifyNo, &merchant.Disabled, &merchant.CreateTime, &merchant.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return merchant
}
