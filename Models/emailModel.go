package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	id          primitive.ObjectID `json:"_id"`
	ToEmail     string             `json:"toemail"`
	FromEmail   string             `json:"fromemail"`
	AppPassword string             `json:"apppassword"`
	DeviceToken string             `json:"devicetoken"`
	Host        string             `json:"host"`
	Port        string             `json:"port"`
}
