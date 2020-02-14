package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"winter_holiday/data"
)

func OpenDatabase() (db *sql.DB, a bool) {
	db, err := sql.Open("mysql", "root:123456@tcp(Localhost:3306)/Zhihu?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return db, false
	}
	return db, true
}

func InsertDB(db *sql.DB, temp *data.User) bool {
	stmt, err := db.Prepare("insert into user(username,password,phonenumber) values (?,?,?)")
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		stmt.Exec(temp.Name, temp.Password, temp.Phonenumber)
		fmt.Println("添加用户成功")
		return true
	}
}

func SelectDB(db *sql.DB, temp *data.User) bool {
	stmt, err := db.Query("select * from user where username =?;", temp.Name)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	for stmt.Next() {
		var name string
		var password string
		var phonenumber string
		err := stmt.Scan(&name, &password, &phonenumber)
		if err != nil {
			fmt.Println(err)
		}
		if name == temp.Name && password == temp.Password {
			fmt.Println("验证通过，")
			return true
		} else if phonenumber == temp.Phonenumber && password == temp.Password {
			fmt.Println("验证通过，")
			return true
		}
	}
	fmt.Println("没有查询到该用户")
	return false
}

func DeleteDB(db *sql.DB, temp *data.User) {
	stmt, err := db.Prepare("delete from user  where username =  (?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	stmt.Exec(temp.Name)
	fmt.Println("删除了用户", temp.Name)
}

func UpdateDB(db *sql.DB, temp *data.User) {
	stmt, err := db.Prepare("UPDATE user SET password = (?),phonenumber=(?) where username = (?)")
	if err != nil {
		fmt.Println(err)
	}
	stmt.Exec(temp.Password, temp.Phonenumber, temp.Name)
}

//检测用户名是否在数据库中已经纯在
func Isexist(db *sql.DB, Name string) bool {
	stmt, err := db.Query("select username from user where username = ?;", Name)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()
	var name string
	for stmt.Next() {
		err := stmt.Scan(&name)
		if err != nil {
			fmt.Println(err)
		}
		if Name == name {
			return true //已经纯在
		}
	}
	return false //没有检测到该用户
}
