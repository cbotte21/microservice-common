package datastore

import (
	"context"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client = nil

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		//connect
		mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(enviroment.GetEnvVariable("mongo_uri")))

		//error check
		if err != nil {
			return nil, err
		}

		//ping
		if err := mongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
			return nil, err
		}
	}
	return mongoClient, nil
}
