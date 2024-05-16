package main

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8" // Package Redis untuk Golang
)

type ServRedis struct {
	Rdb *redis.Client
}

type ServRedisInt interface {
	SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	GetData(ctx context.Context, key string) (string, error)
}

func NewServRedis() *ServRedis {
	rdb := createRedisClient()
	return &ServRedis{
		Rdb: rdb,
	}
}

func createRedisClient() *redis.Client {
	// Mengganti "your-redis-address:port" dengan alamat dan port Redis Anda
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Ganti dengan alamat dan port Redis Anda
		Password: "",               // Jika Anda memiliki password Redis, isi di sini
		DB:       0,                // Indeks database, biasanya 0
	})

	// Pinging server Redis untuk memastikan koneksi berhasil
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

func (p *ServRedis) SetData(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	rds := p.Rdb.Set(context.Background(), key, value, ttl)

	return rds.Err()
}

func (p *ServRedis) GetData(ctx context.Context, key string) (string, error) {
	dataRedis, err := p.Rdb.Get(context.Background(), key).Result()

	return dataRedis, err
}
