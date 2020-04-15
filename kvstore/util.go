package kvstore

import (
	"encoding/json"
	"log"
)

func StringifyJSON(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return string(b[:])
}
