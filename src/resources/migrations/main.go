package main

import (
	"github.com/crypto-monitor/src/infrastructure/database"
	"github.com/crypto-monitor/src/users/domain/models"
)

func main () {
	db := database.Init()	
	db.AutoMigrate(&userModel.User{})
}