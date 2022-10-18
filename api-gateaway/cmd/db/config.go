package db

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

var Client redis.Client

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(pong)

	Client = *client
}
