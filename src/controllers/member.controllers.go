package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NguyenThanhTai-KonZOrA/Golang/src/models"
	"github.com/NguyenThanhTai-KonZOrA/Golang/src/utils"
	"github.com/gorilla/mux"
)

var NewMember models.Member

func CreateMember(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("2================")
	CreateMember := &models.Member{}
	utils.ParseBody(r, CreateMember)
	b := CreateMember.CreateMember()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMember(w http.ResponseWriter, r *http.Request) {
	members := models.GetAllMembers()
	res, _ := json.Marshal(members)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMemberById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	member, _ := models.GetMemberById(ID)
	res, _ := json.Marshal(member)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateMember(w http.ResponseWriter, r *http.Request) {
	var updateMember = &models.Member{}
	utils.ParseBody(r, updateMember)
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	member, db := models.GetMemberById(ID)
	if updateMember.MemberName != "" {
		member.MemberName = updateMember.MemberName
	}
	if updateMember.PhoneNumber != "" {
		member.PhoneNumber = updateMember.PhoneNumber
	}
	if updateMember.Email != "" {
		member.Email = updateMember.Email
	}
	db.Save(&member)
	res, _ := json.Marshal(member)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMember(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberId := vars["memberId"]
	ID, err := strconv.ParseInt(memberId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	member := models.DeleteMember(ID)
	res, _ := json.Marshal(member)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
