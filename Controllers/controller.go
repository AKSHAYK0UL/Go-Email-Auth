package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	helpers "sendemail/Helpers"
	models "sendemail/Models"
)

// Send Email
func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var maildata models.Email
	err := json.NewDecoder(r.Body).Decode(&maildata)
	if err != nil {
		log.Fatal(err)

	}
	val := helpers.SendEmailToUser(maildata)
	if val == "Vcode Send" {
		json.NewEncoder(w).Encode("VERIFICATION CODE IS SEND ON YOUR EMAIL")

	} else if val == "Email Already Exist" {
		json.NewEncoder(w).Encode("EMAIL ALREADY EXIST")

	} else if val == "UserName Already Exist" {
		json.NewEncoder(w).Encode("USERNAME ALREADY EXIST")

	} else {

		json.NewEncoder(w).Encode("USERNAME AND EMAIL ALREADY EXIST")

	}
}

// Verify Verification code
func VerifyCode(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userVcode models.Email
	json.NewDecoder(r.Body).Decode(&userVcode)
	val := helpers.CheckVCode(userVcode)
	if val {
		json.NewEncoder(w).Encode("1")
	} else {
		json.NewEncoder(w).Encode("0")
	}
}

// Login
func LoginToAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	val := helpers.Login(user.Email, user.Password)
	if val {
		json.NewEncoder(w).Encode("You are logged in")

	} else {
		json.NewEncoder(w).Encode("Invalid email or password")
	}

}

// ResetPasswword
func Resetpasswword(w http.ResponseWriter, r *http.Request) {
	var userEmail models.Email
	json.NewDecoder(r.Body).Decode(&userEmail)
	exist := helpers.ResetPasssword(userEmail)
	if exist {
		json.NewEncoder(w).Encode("USER EXIST")

	} else {
		json.NewEncoder(w).Encode("USER NOT FOUND")
	}
}

// Verify and Upddate
func VerifyandUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userVcode models.Email
	json.NewDecoder(r.Body).Decode(&userVcode)
	val := helpers.CheckResetVCode(userVcode)
	if val {
		json.NewEncoder(w).Encode("1")
	} else {
		json.NewEncoder(w).Encode("0")
	}
}
