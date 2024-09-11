package appRedis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

type appRedis struct {
	Ctx    context.Context
	Client *redis.Client
}

var AppRedis appRedis

func New() *appRedis {
	url := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		DB:       0, // use default DB
	})

	return &appRedis{Ctx: context.Background(), Client: rdb}
}

func (r *appRedis) Get(key, field string) (string, error) {
	return r.Client.HGet(r.Ctx, key, field).Result()
}

func (r *appRedis) Set(key string, field string, val any) error {
	return r.Client.HSet(r.Ctx, key, field, val).Err()
}

func (r *appRedis) Del(key, field string) error {
	return r.Client.HDel(r.Ctx, key, field).Err()
}

func (r *appRedis) Count(key, pattern string) (int64, error) {
	var count int64
	iter := r.Client.HScan(r.Ctx, key, 0, pattern, 0).Iterator()
	for iter.Next(r.Ctx) {
		count++
	}
	return count, iter.Err()
}

func (r *appRedis) Exists(key, field string) (bool, error) {
	return r.Client.HExists(r.Ctx, key, field).Result()
}
