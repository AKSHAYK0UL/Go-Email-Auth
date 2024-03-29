package router

import (
	controllers "sendemail/Controllers"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/verify", controllers.VerifyOtp).Methods("POST")
	r.HandleFunc("/create", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/login", controllers.LoginToAccount).Methods("POST")

	return r
}
