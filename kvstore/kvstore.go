package kvstore

import (
	"log"
	// "strings"

	"github.com/gomodule/redigo/redis"
)

// Dial connects to redis and returns the connected client
func Dial() redis.Conn {
	client, err := redis.Dial(connProto, connURL)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func HExists(hashKey string, field string) (bool, error) {
	client := Dial()
	defer client.Close()
	return redis.Bool(client.Do("HEXISTS", hashKey, field))
}

func HSet(hashKey string, field string, value string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HSET", hashKey, field, value)
}

func HGet(hashKey string, field string) (string, error) {
	client := Dial()
	defer client.Close()
	return redis.String(client.Do("HGET", hashKey, field))
}

func HDelete(hashKey string, field string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HDEL", hashKey, field)
}

func Set(key string, value string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("SET", key, value)
}

func Get(key string) (string, error) {
	client := Dial()
	defer client.Close()
	return redis.String(client.Do("GET", key))
}
