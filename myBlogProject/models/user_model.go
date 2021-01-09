package models

import (
	"fmt"
	"myBlogProject/utils"
)

type User struct {
	Id 			int
	Username	string
	Password	string
	Status		int
	Createtime	int64
}

//插入
func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username, password, status, createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}