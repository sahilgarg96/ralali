package redis

import (
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/config"
)

func RedisConnect() (client *redis.Client, err error) {
	redisConnectURI := config.Get("codeIndependent", "redisAConnectURI").String("localhost:6379")
	redisPassword := config.Get("codeIndependent", "redisAPassword").String("")
	redisDB := config.Get("codeIndependent", "redisADB").Int(0)

	client = redis.NewClient(&redis.Options{
		Addr:     redisConnectURI,
		Password: redisPassword,
		DB:       redisDB,
	})

	_, err = client.Ping().Result()

	return client, err
}
