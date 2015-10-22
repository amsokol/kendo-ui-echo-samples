package logger

import (
	"encoding/json"
	"log"
)

func Json(v interface{}) {
	b, _ := json.Marshal(v)
	log.Println(string(b[:]))
}
