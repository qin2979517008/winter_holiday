package model

import (
	"fmt"
	"winter_holiday/data"
	"winter_holiday/repository"
)

func  InserUser(temp *data.User)bool{
	db, err := repository.OpenDatabase()
	defer db.Close()
	if err == false {
		return false
		fmt.Println(err)
	}
	if repository.InsertDB(db , temp){
		return true
	}else{
		return false
	}
}