package redis

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

var lock = &sync.Mutex{}

type singleton struct {
	client *redis.Client
}

var instance *singleton

func GetInstance() *redis.Client {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &singleton{
				client: GetClient(),
			}
		}
	}
	return instance.client
}

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("SERVER - Error connecting to redis")
	}
	log.Print("SERVER - Connected to redis")
	return client
}

func Publish(channel string, message string) {
	GetInstance().Publish(channel, message)
}
