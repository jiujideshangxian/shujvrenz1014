package db_mysql

import (
	"database/sql"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect(){
	config:=beego.AppConfig
	dbDriver:=config.String("db_driver")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIp:=config.String("db_ip")
    db_Name:=config.String("db_name")

	connUrl:=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+db_Name+"?charset+utf8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err != nil {
		panic("错误")
	}
	Db=db
}
