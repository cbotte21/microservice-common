package test

import (
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"testing"
)

func TestFindMongo(t *testing.T) {
	enviroment.VerifyEnvVariable("mongo_uri")
	user := schema.User{Email: "cbotte21@gmail.com"}

	client := datastore.MongoClient[schema.User]{}
	err := client.Init()
	if err != nil {
		t.Fatalf(err.Error())
	}
	res, err := client.Find(user)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if res.Email != user.Email {
		t.Fatalf("received incorrect data!")
	}
}
