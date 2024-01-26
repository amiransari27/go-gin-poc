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

const collectionName string = "users"

type IUserDao interface {
	Find(bson.M) (*model.User, error)
	FindOne(bson.M) (*model.User, error)
	Save(*model.User) (interface{}, error)
}

type userDao struct {
	coll *mongo.Collection
}

func NewUserDao(mcd mongoDB.MongoClientDatabase) IUserDao {

	userCollection := mcd.Database.Collection(collectionName)

	obj := &userDao{
		coll: userCollection,
	}
	obj.createIndex()

	return obj
}

func (ud *userDao) createIndex() {
	//create index and unique key
	uniqueIndex := mongo.IndexModel{
		Keys: bson.M{
			"username": -1, // index in decending order
		}, Options: options.Index().SetUnique(true),
	}

	_, err := ud.coll.Indexes().CreateOne(context.Background(), uniqueIndex)
	if err != nil {
		log.Fatal(err)
	}

	uniqueIndex = mongo.IndexModel{
		Keys: bson.M{
			"userId": -1, // index in decending order
		}, Options: options.Index().SetUnique(true),
	}

	_, err = ud.coll.Indexes().CreateOne(context.Background(), uniqueIndex)
	if err != nil {
		log.Fatal(err)
	}

	index := mongo.IndexModel{
		Keys: bson.M{
			"createdAt": -1, // index in decending order
		}, Options: nil,
	}

	_, err = ud.coll.Indexes().CreateOne(context.Background(), index)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("success index creation for user collection")
}

func (ud *userDao) Save(userObj *model.User) (interface{}, error) {
	// var user model.User
	inserted, err := ud.coll.InsertOne(context.Background(), userObj)

	if err != nil {
		return nil, err
	}

	return inserted.InsertedID, err
}

func (ud *userDao) Find(cond bson.M) (*model.User, error) {
	var user []model.User
	userCursor, err := ud.coll.Find(context.Background(), cond) // find multiple

	if err != nil {
		return nil, err
	}

	err = userCursor.All(context.Background(), &user)

	if err != nil {
		return nil, err
	}

	return &user[0], nil
}

func (ud *userDao) FindOne(cond bson.M) (*model.User, error) {
	var user model.User
	userObj := ud.coll.FindOne(context.Background(), cond) // find

	userObj.Decode(&user)

	return &user, nil
}
