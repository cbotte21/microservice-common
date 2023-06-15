package test

import (
	"context"
	"fmt"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"testing"
)

func TestPubSubRedis(t *testing.T) {
	client := datastore.RedisClient[schema.User]{}

	err := client.Init()
	if err != nil {
		t.Fatalf(err.Error())
	}

	sub := client.Subscribe("test", "test2")
	pub := client.Publish("test",
		"hello")
	pub2 := client.Publish("test2", "hello")

	if sub == nil || pub != nil || pub2 != nil {
		if sub == nil {
			t.Fatalf("Subscription was null")
		}

		if pub != nil {
			t.Fatalf(pub.Error())
		}

		t.Fatalf(pub2.Error())
	}

	msgCount := 0

	for {
		_, err := sub.ReceiveMessage(context.Background())
		if err != nil {
			t.Fatalf(err.Error())
		}

		msgCount++

		if msgCount == 2 {
			break
		}
	}

	if msgCount != 2 {
		t.Fatalf("Message count was incorrect")
	}
}

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

	err = client.Delete(user1)
	if err != nil {
		return
	}
	err = client.Delete(user2)
	if err != nil {
		return
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

	fmt.Println(user)
	fmt.Println(res)
	if res.Email != user.Email {
		t.Fatalf("received incorrect data!")
	}
}
