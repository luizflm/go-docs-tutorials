package main

import (
	"encoding/json/v2"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time time.Time
}

func main() {
	// 1:
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":"2011-01-25T00:00:00Z"}`)
	var m Message
	err := json.Unmarshal(b, &m)

	// 2 (structure of the JSON data doesn’t exactly match the Go type):
	// b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	// var m Message
	// err := json.Unmarshal(b, &m)

	fmt.Println(m, err)
}
