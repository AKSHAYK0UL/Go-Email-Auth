package models

type OTPModel struct {
	UserEmail string `json:"useremail"`
	OTP       string `json:"otp"`
}
