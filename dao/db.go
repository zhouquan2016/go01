package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go01/config"
	"go01/util"
	"log"
)

var Db *sql.DB

func initDb() {
	ds := config.Config.Datasource
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", *ds.Username, *ds.Password, *ds.Host, *ds.Port, *ds.Database, *ds.Query)

	db, err := sql.Open(*ds.Type, dbDSN)

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Db = db
}

func init() {
	initDb()
}

func OpenTx() (tx *sql.Tx) {
	var err error
	tx, err = Db.Begin()
	util.AssertError(err, "打开事务失败")
	return tx
}

func EndTx(tx *sql.Tx) {
	if tx == nil {
		return
	}
	err := recover()
	if err != nil {
		_ = tx.Rollback()
		log.Println("tx rollback!")
	} else {
		commitErr := tx.Commit()
		util.AssertError(commitErr, "事务提交失败")
		log.Println("tx commit!")
		err = nil
	}
	if err != nil {
		panic(err)
	}

}

//tx不为空加读锁，为空不加读锁
func queryRow(tx *sql.Tx, querySql string, params ...interface{}) *sql.Row {
	if tx != nil {
		return tx.QueryRow(querySql+" LOCK IN SHARE MODE", params...)
	}
	return Db.QueryRow(querySql, params...)
}
