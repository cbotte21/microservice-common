package datastore

import (
	"context"
	"encoding/json"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

type RedisClient[T schema.Schema[any]] struct {
	GoRedisClient *redis.Client
	ReJsonHandler *rejson.Handler
}

func (client *RedisClient[T]) Init() error {
	client.ReJsonHandler = rejson.NewReJSONHandler()

	address := enviroment.GetEnvVariable("redis_addr")

	if enviroment.GetEnvVariable("redis_addr") == "" {
		address = "127.0.0.1:6379"
	}

	client.GoRedisClient = redis.NewClient(&redis.Options{Addr: address, DB: 0})
	client.ReJsonHandler.SetGoRedisClient(client.GoRedisClient)
	return nil
}

func (client *RedisClient[T]) InitTest() {
	client.ReJsonHandler = rejson.NewReJSONHandler()
	db, _ := redismock.NewClientMock()
	client.GoRedisClient = db
	client.ReJsonHandler.SetGoRedisClient(client.GoRedisClient)
}

func (client *RedisClient[T]) Find(schema T) (T, error) {
	res, err := client.ReJsonHandler.JSONGet(schema.Key(), ".")
	bytes, err := redigo.Bytes(res, err)
	if err != nil {
		return schema, err
	}
	err = json.Unmarshal(bytes, &schema)
	return schema, err
}

func (client *RedisClient[T]) Create(schema T) error {
	_, err := client.ReJsonHandler.JSONSet(schema.Key(), ".", schema)
	return err
}

func (client *RedisClient[T]) Update(_, schema T) error {
	return client.Create(schema)
}

func (client *RedisClient[T]) Delete(schema T) error {
	_, err := client.ReJsonHandler.JSONDel(schema.Key(), ".")
	return err
}

func (client *RedisClient[T]) Subscribe(channels ...string) *redis.PubSub {
	ctx := context.Background()
	subscriber := client.GoRedisClient.Subscribe(ctx, channels...)
	return subscriber
}

func (client *RedisClient[T]) Publish(channel string, message any) error {
	return client.GoRedisClient.Publish(context.Background(), channel, message).Err()
}
