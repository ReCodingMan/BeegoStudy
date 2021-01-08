package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbCon := "root:123@tcp(127.0.0.1:3306)/beego?charset=utf-8"
	dbCon := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf-8"

	err := orm.RegisterDataBase("default", driverName, dbCon)
	if err != nil {
		return
	}
}
