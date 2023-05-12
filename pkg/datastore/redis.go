package datastore

import (
	"encoding/json"
	"github.com/cbotte21/microservice-common/pkg/enviroment"
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/go-redis/redis/v8"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
)

type RedisClient[T schema.Schema[any]] struct {
	*rejson.Handler
}

func (client *RedisClient[T]) Init() error {
	client.Handler = rejson.NewReJSONHandler()
	enviroment.VerifyEnvVariable("redis_addr")
	client.SetGoRedisClient(redis.NewClient(&redis.Options{Addr: enviroment.GetEnvVariable("redis_addr"), DB: 0}))

	return nil
}

func (client *RedisClient[T]) Find(schema T) (T, error) {
	res, err := client.JSONGet(schema.Key(), ".")
	bytes, err := redigo.Bytes(res, err)
	if err != nil {
		return schema, err
	}
	err = json.Unmarshal(bytes, schema)
	return schema, err
}

func (client *RedisClient[T]) Create(schema T) error {
	_, err := client.JSONSet(schema.Key(), ".", schema)
	return err
}

func (client *RedisClient[T]) Update(_, Y T) error {
	return client.Create(Y)
}

func (client *RedisClient[T]) Delete(schema T) error {
	_, err := client.JSONDel(schema.Key(), ".")
	return err
}
