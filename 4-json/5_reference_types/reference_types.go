package main

import (
	"encoding/json/v2"
	"fmt"
)

type FamilyMember struct {
    Name    string
    Age     int
    Parents []string
}

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var m FamilyMember
	err := json.Unmarshal(b, &m)

	fmt.Println(m, err)
}