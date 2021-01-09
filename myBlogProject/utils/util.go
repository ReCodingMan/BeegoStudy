package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
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

	db1, err := sql.Open(driverName, dbCon)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
	}
}

//创建用户表
func CreateTableWithUser() {
	sql := "CREATE TABLE IF NOT EXISTS users(" +
		"id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL," +
		"username VARCHAR(64)," +
		"password VARCHAR(64)," +
		"status INT(4)," +
		"createtime INT(11)" +
		")"
	ModifyDB(sql)
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
