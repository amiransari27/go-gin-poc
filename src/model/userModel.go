package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	UserId    string             `json:"userId" bson:"userId"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"-" bson:"password"`
	FirstName string             `json:"firstName" bson:"firstName"`
	LastName  string             `json:"lastName" bson:"lastName"`
}

func (o *User) MarshalBSON() ([]byte, error) {
	if o.CreatedAt.IsZero() {
		o.CreatedAt = time.Now()
	}
	o.UpdatedAt = time.Now()

	type tmp User
	return bson.Marshal((*tmp)(o))
}
