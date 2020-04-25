package routes

import (
	"github.com/NguyenThanhTai-KonZOrA/Golang/src/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	//Router Member
	router.HandleFunc("/member", controllers.CreateMember).Methods("POST")
	router.HandleFunc("/member", controllers.GetMember).Methods("GET")
	router.HandleFunc("/member/{memberId}", controllers.GetMemberById).Methods("GET")
	router.HandleFunc("/member/{memberId}", controllers.UpdateMember).Methods("PUT")
	router.HandleFunc("/member/{memberId}", controllers.DeleteMember).Methods("DELETE")
}
