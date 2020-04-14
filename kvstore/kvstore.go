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

func Exists(hashKey string, field string) (bool, error) {
	client := Dial()
	defer client.Close()
	return redis.Bool(client.Do("HEXISTS", hashKey, field))
}

func Set(hashKey string, field string, value string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HSET", hashKey, field, value)
}

func Get(hashKey string, field string) (string, error) {
	client := Dial()
	defer client.Close()
	return redis.String(client.Do("HGET", hashKey, field))
}

func Delete(hashKey string, field string) (interface{}, error) {
	client := Dial()
	defer client.Close()
	return client.Do("HDEL", hashKey, field)
}
