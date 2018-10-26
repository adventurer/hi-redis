package models

import (
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	Connect()
}

func Connect() {
	Client = redis.NewClient(&redis.Options{
		Addr:     AppConfig.DBIp + ":" + AppConfig.DBPort,
		Password: AppConfig.DBAuth,
		DB:       0, // use default DB
	})
}
