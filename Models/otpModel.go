package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OTPModel struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	UserEmail string             `json:"useremail"`
	OTP       string             `json:"otp"`
}
