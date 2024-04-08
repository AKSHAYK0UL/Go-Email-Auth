package router

import (
	controllers "sendemail/Controllers"

	"github.com/gorilla/mux"
)

func Routers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	r.HandleFunc("/verify", controllers.VerifyCode).Methods("POST")
	r.HandleFunc("/login", controllers.LoginToAccount).Methods("POST")
	r.HandleFunc("/reset", controllers.Resetpasswword).Methods("POST")
	r.HandleFunc("/resetverify", controllers.VerifyandUpdate).Methods("POST")

	return r
}
