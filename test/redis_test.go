package test

import (
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"testing"
)

func TestCreateRedis(t *testing.T) {
	client := datastore.RedisClient[schema.User]{}

	user1 := schema.User{Id: "1", Email: "john@redis.com", Password: "Asdfasdf1"}
	user2 := schema.User{Id: "2", Email: "cbotte21@gmail.com", Password: "Asdfasdf1"}

	err := client.Init()
	if err != nil {
		t.Fatalf(err.Error())
	}

	err = client.Create(user1)
	if err != nil {
		t.Fatalf(err.Error())
	}
	err = client.Create(user2)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestFindRedis(t *testing.T) {
	client := datastore.RedisClient[schema.User]{}

	user := schema.User{Id: "1", Email: "cody@redis.com", Password: "Asdfasdf1"}

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
