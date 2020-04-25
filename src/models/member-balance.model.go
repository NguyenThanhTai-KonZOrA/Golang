package models

import (
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB

type MemberBalance struct {
	gorm.Model
	//Id          string `json:"id"`
	MemberId string `gorm:""json:"name"`
	Balance  int    `json:"balance"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&MemberBalance{})
// }
