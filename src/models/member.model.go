package models

import (
	"fmt"

	"github.com/NguyenThanhTai-KonZOrA/Golang/src/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Member struct {
	gorm.Model
	//Id          string `json:"id"`
	MemberName  string `json:"memberName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	VoucherId   string `json:"voucherId"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Member{})
}

func (member *Member) CreateMember() *Member {
	fmt.Printf("asdasdasdasdasd")

	db.NewRecord(member)
	db.Create(&member)
	return member
}

func GetAllMembers() []Member {
	var Members []Member
	db.Find(&Members)
	return Members
}

func GetMemberById(Id int64) (*Member, *gorm.DB) {
	var member Member
	db := db.Where("ID = ?", Id).Find(&member)
	return &member, db
}

func DeleteMember(ID int64) Member {
	var member Member
	db.Where("ID = ?", ID).Delete(member)
	return member
}
