package datastore

import (
	"context"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient[T schema.Schema[any]] struct {
	*mongo.Client
}

func (client *MongoClient[T]) Init() error {
	enviroment.VerifyEnvVariable("mongo_uri")
	if client.Client == nil {
		//connect
		var err error
		client.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(enviroment.GetEnvVariable("mongo_uri")))

		//error check
		if err != nil {
			panic(err)
			return err
		}

		//ping
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
			return err
		}
	}
	return nil
}

func (client *MongoClient[T]) Find(schema T) (T, error) {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	err := collection.FindOne(context.TODO(), schema).Decode(&schema)
	return schema, err
}

func (client *MongoClient[T]) Create(schema T) error {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err := collection.InsertOne(context.TODO(), schema)
	return err
}

func (client *MongoClient[T]) Update(X, Y T) error {
	collection := client.Database(X.Database()).Collection(X.Collection())
	_, err := collection.UpdateOne(context.TODO(), X, Y)
	return err
}

func (client *MongoClient[T]) Delete(schema T) error {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err := collection.DeleteOne(context.TODO(), schema)
	return err
}
