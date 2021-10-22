package db

import (
	"afterSaleSys/util"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	Db *sqlx.DB
)

func InitDb() (*sqlx.DB, error) {
	keys := []string{"url", "user", "passwd"}
	iniMap, err := util.ReadIni("mysql", keys)
	if err != nil {
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/aftersale?charset=utf8&parseTime=True", iniMap["user"], iniMap["passwd"], iniMap["url"])
	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}
	Db.SetMaxOpenConns(200)
	Db.SetMaxIdleConns(100)
	return Db, nil
}
