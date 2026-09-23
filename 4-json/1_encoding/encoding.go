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
	m := Message{"Alice", "Hello", time.Date(2011, 1, 25, 0, 0, 0, 0, time.UTC)}

	b, err := json.Marshal(m)

	fmt.Println(b, err)
}
