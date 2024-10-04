package mongodb

import (
	"context"
	"fmt"
	"log"
	// "project1/models"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo(){
	uri := "mongodb://localhost:27017"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Panic("Error: Creating mongoDB client", err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Panic("Could not ping mongoDB", err)
	}

	MongoClient = client
	fmt.Println("MongoDB connection established")
}

func InsertOneDocument(collectionName string , document interface{}) (interface {}, error){
	fmt.Println("###############",document)
	collection := MongoClient.Database("testdb").Collection(collectionName)

	result, err := collection.InsertOne(context.Background(), document)
    if err != nil {
		fmt.Println(err)
        return nil, fmt.Errorf("error inserting document: %v", err)
    }
    return result.InsertedID, nil
}

func GetDocumentById(collectionName string, )