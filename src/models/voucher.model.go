package models

import (
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB

type Voucher struct {
	gorm.Model
	//Id          string `json:"id"`
	VoucherName string `gorm:""json:"memberId"`
	Quantity    int    `json:"quantity"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&Voucher{})
// }
