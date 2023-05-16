package test

import (
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPubSubRedis(t *testing.T) {
	client := datastore.RedisClient[schema.User]{}

	mock := client.MockInit()
	mock.MatchExpectationsInOrder(false)
	_ = client.Publish("test", "hello")
	_ = client.Publish("test2", "hello")

	mock.ExpectPublish("test", "hello")
	mock.ExpectPublish("test2", "hello")
}

func TestMockCreateFind(t *testing.T) {
	client := datastore.RedisClient[schema.User]{}

	user1 := schema.User{Id: "1", Email: "john@redis.com", Password: "Asdfasdf1"}
	user2 := schema.User{Id: "2", Email: "cbotte21@gmail.com", Password: "Asdfasdf1"}

	mock := client.MockInit()
	mock.MatchExpectationsInOrder(false)
	mock.ExpectSet(user1.Id, user1, 10*time.Second).SetVal("OK")
	mock.ExpectSet(user2.Id, user2, 10*time.Second).SetVal("OK")
	mock.ExpectGet(user1.Id).SetVal(Marshal(t, user1))
	mock.ExpectGet(user2.Id).SetVal(Marshal(t, user2))

	err := client.Create(user1)
	assert.NoError(t, err)
	err = client.Create(user2)
	assert.NoError(t, err)

	_, err = client.Find(schema.User{Id: "1"})
	assert.NoError(t, err)
	_, err = client.Find(schema.User{Id: "2"})
	assert.NoError(t, err)
}

func TestPubSub(t *testing.T) {
	client, err := GetRedis()
	assert.NoError(t, err)
	ps := client.Subscribe("test")
	go func() {
		for msg := range ps.Channel() {
			if len(msg.Payload) > 0 {
				break
			}
		}
	}()
	err = client.Publish("test", "test")
	assert.NoError(t, err)
}

func TestCreateFind(t *testing.T) {
	client, err := GetRedis()
	assert.NoError(t, err)

	user1 := schema.User{Id: "1", Email: "john@redis.com", Password: "Asdfasdf1"}
	user2 := schema.User{Id: "2", Email: "cbotte21@gmail.com", Password: "Asdfasdf1"}

	err = client.Create(user1)
	assert.NoError(t, err)
	err = client.Create(user2)
	assert.NoError(t, err)

	res, err := client.Find(user1)
	assert.NoError(t, err)
	assert.Equal(t, user1, res)
	res, err = client.Find(user2)
	assert.NoError(t, err)
	assert.Equal(t, user2, res)
}
