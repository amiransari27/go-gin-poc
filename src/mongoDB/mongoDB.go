package mongoDB

import (
	"context"
	"fmt"
	"go-gin-api/src/config"
	"log"
	"time"

	logger "github.com/openscriptsin/go-logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientDatabase struct {
	*mongo.Client
	*mongo.Database
}

func NewMongoConnection(logger logger.ILogrus) MongoClientDatabase {

	mongoCred := config.GetConfig().Mongo
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", mongoCred.DBUser, mongoCred.DBPassword, mongoCred.DBHost)
	// connectionString := "mongodb://localhost:27017/?"
	dbName := mongoCred.DBName

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		client = nil
		log.Fatal(err)
	}
	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		client = nil
		log.Fatal(err)
	}
	logger.Info(nil, "Connected to MongoDB")

	database := client.Database(dbName)

	return MongoClientDatabase{client, database}
}
