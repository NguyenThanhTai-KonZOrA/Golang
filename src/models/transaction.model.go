package models

import (
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB

type Transaction struct {
	gorm.Model
	//Id          string `json:"id"`
	MemberId string `gorm:""json:"memberId"`
}

// func init() {
// 	config.Connect()
// 	db = config.GetDB()
// 	db.AutoMigrate(&Transaction{})
// }

func (transaction *Transaction) CreateTransaction() *Transaction {
	db.NewRecord(transaction)
	db.Create(&transaction)
	return transaction
}
