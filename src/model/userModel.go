package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Username  string             `json:"username" `
	Password  string             `json:"-" `
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
}