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
	helpers.SendEmailToUser(maildata)

	json.NewEncoder(w).Encode("OTP IS SEND ON YOUR EMAIL")

}

// Verify OTP
func VerifyOtp(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var userotp models.OTPModel
	json.NewDecoder(r.Body).Decode(&userotp)
	val := helpers.CheckOTP(userotp)
	if val {
		json.NewEncoder(w).Encode("1")
	} else {
		json.NewEncoder(w).Encode("0")
	}
}

// Create New User
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	helpers.CreateAndSaveUser(user)
	json.NewEncoder(w).Encode("User Created")
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
