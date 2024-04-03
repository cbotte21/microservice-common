package datastore

import (
	"context"
	"github.com/cbotte21/microservice-common/pkg/environment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoClient[T schema.Schema[any]] struct {
	schema T
	*mongo.Client
}

func (client *MongoClient[T]) Init() {
	if client.Client == nil {
		//connect
		var err error
		client.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(environment.GetEnvVariable("mongo_addr")))

		//error check
		if err != nil {
			log.Fatalf(err.Error())
		}

		//ping
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			log.Fatalf(err.Error())
		}
	}
}

func (client *MongoClient[T]) Find(schema T) (T, error) {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	err := collection.FindOne(context.TODO(), schema).Decode(&schema)
	return schema, err
}

func (client *MongoClient[T]) FindAll() ([]T, error) {
	collection := client.Database(client.schema.Database()).Collection(client.schema.Collection())
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var elements []T
	for cursor.Next(context.Background()) {
		var found T
		if err := cursor.Decode(&found); err != nil {
			return nil, err
		}
		elements = append(elements, found)
	}

	return elements, err
}

func (client *MongoClient[T]) Create(schema T) error {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err := collection.InsertOne(context.TODO(), schema)
	return err
}

func (client *MongoClient[T]) Update(X, Y T) error {
	collection := client.Database(X.Database()).Collection(X.Collection())
	_, err := collection.UpdateOne(context.TODO(), X, bson.M{"$set": Y})
	return err
}

func (client *MongoClient[T]) Delete(schema T) error {
	collection := client.Database(schema.Database()).Collection(schema.Collection())
	_, err := collection.DeleteOne(context.TODO(), schema)
	return err
}
