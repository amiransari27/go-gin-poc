package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
	UserId     string             `json:"userId" bson:"userId"`
	NoteId     string             `json:"noteId" bson:"noteId"`
	Title      string             `json:"title" bson:"title"`
	Content    string             `json:"content" bson:"content" `
	SharedWith []string           `json:"sharedWith" bson:"sharedWith"`
}

func (o *Note) MarshalBSON() ([]byte, error) {
	if o.CreatedAt.IsZero() {
		o.CreatedAt = time.Now()
	}
	o.UpdatedAt = time.Now()

	type tmp Note
	return bson.Marshal((*tmp)(o))
}
