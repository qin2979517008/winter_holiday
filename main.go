package main

import (
	"winter_holiday/router"
)

func main() {
	router := router.Router()
	router.Run(":8080")
}