package test

import (
	"context"
	"encoding/json"
	"github.com/alicebob/miniredis/v2"
	"github.com/cbotte21/microservice-common/pkg/datastore"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MarshalAuto(t *testing.T, user schema.User) []byte {
	marshal, err := json.Marshal(user)
	assert.Equal(t, err, nil)
	return marshal
}

func Marshal(t *testing.T, user schema.User) string {
	return string(MarshalAuto(t, user))
}

func MockRedis() *miniredis.Miniredis {
	miniRedis, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return miniRedis
}

func SetupRedis() (*miniredis.Miniredis, *redis.Client) {
	server := MockRedis()
	return server, redis.NewClient(&redis.Options{Addr: server.Addr()})
}

func GetRedis() (datastore.RedisClient[schema.User], error) {
	_, redisClient := SetupRedis()
	client := datastore.RedisClient[schema.User]{}
	client.InitClient(redisClient)

	err := redisClient.Ping(context.Background()).Err()
	return client, err
}
