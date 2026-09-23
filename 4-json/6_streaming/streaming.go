package main

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {
	m := FamilyMember{
		Name:    "Wednesday",
		Age:     6,
		Parents: []string{"Gomez", "Morticia"},
	}

	if err := json.MarshalWrite(os.Stdout, m); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
