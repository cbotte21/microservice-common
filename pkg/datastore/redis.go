package datastore

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cbotte21/microservice-common/pkg/environment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
)

type RedisClient[T schema.Schema[any]] struct {
	GoRedis *redis.Client
	Schema  T
	ctx     context.Context
}

func (client *RedisClient[T]) Init() {
	address := environment.GetEnvVariable("redis_addr") // Ex) "127.0.0.1:6379"
	client.GoRedis = redis.NewClient(&redis.Options{Addr: address, DB: 0})
	client.InitClient(client.GoRedis)
}

func (client *RedisClient[T]) InitClient(_client *redis.Client) {
	client.GoRedis = _client
	client.ctx = context.Background()
}

func (client *RedisClient[T]) MockInit() redismock.ClientMock {
	db, mock := redismock.NewClientMock()
	client.GoRedis = db
	client.ctx = context.Background()
	return mock
}

func (client *RedisClient[T]) Find(schema T) (T, error) {
	val, err := client.GoRedis.Get(client.ctx, schema.Key()).Result()
	if err != nil {
		return schema, err
	}
	err = json.Unmarshal([]byte(val), &schema)
	return schema, err
}

func (client *RedisClient[T]) Create(schema T) error {
	setCmd := client.GoRedis.Set(client.ctx, schema.Key(), schema, 0)
	_, err := setCmd.Result()
	return err
}

func (client *RedisClient[T]) Update(_, schema T) error {
	return client.Create(schema)
}

func (client *RedisClient[T]) Delete(schema T) error {
	_, err := client.GoRedis.Del(client.ctx, schema.Key()).Result()
	return err
}

// PubSub

func (client *RedisClient[T]) Subscribe(channels ...string) *redis.PubSub {
	return client.GoRedis.Subscribe(client.ctx, channels...)
}

func (client *RedisClient[T]) Publish(channel string, message any) error {
	return client.GoRedis.Publish(client.ctx, channel, message).Err()
}

// Queue

func (client *RedisClient[T]) Enqueue(schema T) error {
	res := client.GoRedis.LPush(context.Background(), schema.Key(), schema)
	return res.Err()
}

func (client *RedisClient[T]) Dequeue() (T, error) {
	resInArr, err := client.GoRedis.BRPop(context.Background(), 0, client.Schema.Key()).Result()
	if err != nil {
		return client.Schema, err
	}
	if len(resInArr) == 0 {
		return client.Schema, errors.New("queue is empty")
	}
	res := resInArr[1]
	var parsed T
	err = json.Unmarshal([]byte(res), &parsed)
	return parsed, err
}
