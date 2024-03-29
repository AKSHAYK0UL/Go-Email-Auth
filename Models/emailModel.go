package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	ToEmail     string             `json:"toemail"`
	FromEmail   string             `json:"fromemail"`
	AppPassword string             `json:"apppassword"`
	Host        string             `json:"host"`
	Port        string             `json:"port"`
}
