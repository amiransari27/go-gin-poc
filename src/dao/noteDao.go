package dao

import (
	"context"
	"fmt"
	"go-gin-api/src/model"
	"go-gin-api/src/mongoDB"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const notesCollectionName string = "notes"

type INoteDao interface {
	Find(bson.M) ([]*model.Note, error)
	FindOne(bson.M) (*model.Note, error)
	Save(*model.Note) (interface{}, error)
	FindOneAndUpdate(bson.M, bson.M) (*model.Note, error)
}

type noteDao struct {
	coll *mongo.Collection
}

func NewNoteDao(mcd mongoDB.MongoClientDatabase) INoteDao {

	noteCollection := mcd.Database.Collection(notesCollectionName)
	obj := &noteDao{
		coll: noteCollection,
	}
	obj.createIndex()
	return obj
}

func (nd *noteDao) createIndex() {
	//create index and unique key
	uniqueIndex := mongo.IndexModel{
		Keys: bson.M{
			"noteId": -1, // index in decending order
		}, Options: options.Index().SetUnique(true),
	}

	_, err := nd.coll.Indexes().CreateOne(context.Background(), uniqueIndex)
	if err != nil {
		log.Fatal(err)
	}

	index := mongo.IndexModel{
		Keys: bson.M{
			"userId": -1, // index in decending order
		}, Options: nil,
	}

	_, err = nd.coll.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatal(err)
	}

	index = mongo.IndexModel{
		Keys: bson.M{
			"createdAt": -1, // index in decending order
		}, Options: nil,
	}

	_, err = nd.coll.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success index creation for note collection")
}

func (nd *noteDao) Save(noteObj *model.Note) (interface{}, error) {
	// var user model.User
	inserted, err := nd.coll.InsertOne(context.Background(), noteObj)

	if err != nil {
		return nil, err
	}

	return inserted.InsertedID, err
}

func (nd *noteDao) Find(cond bson.M) ([]*model.Note, error) {
	var notes []*model.Note
	cursor, err := nd.coll.Find(context.Background(), cond) // find multiple

	if err != nil {
		return nil, err
	}

	err = cursor.All(context.Background(), &notes)

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (nd *noteDao) FindOne(cond bson.M) (*model.Note, error) {
	var note model.Note
	obj := nd.coll.FindOne(context.Background(), cond) // find

	obj.Decode(&note)

	return &note, nil
}

func (nd *noteDao) FindOneAndUpdate(filter bson.M, updatedContent bson.M) (*model.Note, error) {
	var note model.Note
	opt := options.FindOneAndUpdate().SetReturnDocument(options.After)
	obj := nd.coll.FindOneAndUpdate(context.Background(), filter, updatedContent, opt)

	obj.Decode(&note)

	return &note, nil
}
