package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//
type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	DeviceToken string             `json:"devicetoken"`
	Islogin     bool               `json:"islogin"`
	HostName    string             `json:"hostname"`
}
