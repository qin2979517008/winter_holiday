package model

import (
	"fmt"
	"winter_holiday/data"
	"winter_holiday/repository"
)

//根据用户名查询用户是否存在
func SearchbyName(name string)bool{
	db ,err := repository.OpenDatabase()
	defer db.Close()
	if err ==false{
		return false
		fmt.Println(err)
	}
	if repository.Isexist(db ,name){
		return false//该用户已经纯在，不能再进行注册
	}
	return true//该用户没有纯在,可以进行注册
}

//验证用户信息是否正确
func SearchUser(student *data.User)bool{
	db,_ := repository.OpenDatabase()
	defer db.Close()
	if repository.SelectDB(db, student) {
		return true//用户名密码匹配
	} else {
		return false
	}
}
