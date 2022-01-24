package dao

import (
	"database/sql"
	"go01/util"
)

type merchantSecretDao struct {
}

var MerchantSecretDao = new(merchantSecretDao)

func (*merchantSecretDao) Add(tx *sql.Tx, secret *MerchantSecret) {
	resultSet, err := tx.Exec("insert into merchant_secret(merchant_id, secret_key, public_key, expire_time, create_time, update_time) values (?,?,?,?,?,?)",
		secret.MerchantId, secret.SecretKey, secret.PublicKey, secret.ExpireTime, secret.CreateTime, secret.UpdateTime)
	util.AssertError(err, "新增商户密钥失败")
	lastId, err := resultSet.LastInsertId()
	util.AssertError(err, "新增商户密钥失败")
	secret.Id = lastId
}

func (d *merchantSecretDao) GetByMerchantId(merchantId int64) *MerchantSecret {
	querySql := "select id, merchant_id, secret_key, public_key, expire_time, create_time, update_time where merchant_id=?"
	row := queryRow(nil, querySql, merchantId)
	return d.getOne(row)

}

func (d *merchantSecretDao) getOne(row *sql.Row) *MerchantSecret {
	secret := new(MerchantSecret)
	err := row.Scan(&secret.Id, &secret.MerchantId, &secret.SecretKey, &secret.PublicKey, &secret.ExpireTime, &secret.CreateTime, &secret.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		panic(err)
	}
	return secret
}
