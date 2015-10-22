package logger

import (
	"log"
	"encoding/json"
)

func Json(v interface{}) {
	b, _ := json.Marshal(v)
	log.Println(string(b[:]))
}
