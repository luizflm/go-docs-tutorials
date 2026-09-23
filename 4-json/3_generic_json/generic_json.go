package main

import (
	"fmt"
	"math"
)

func main() {
	var a any
	// a = "a string"
	// a = 2011
	a = 2.777

	r := a.(float64)
	fmt.Println("the circle's area", math.Pi*r*r)

	switch v := a.(type) {
	case int:
		fmt.Println("twice a is", v*2)
	case float64:
		fmt.Println("the reciprocal of a is", 1/v)
	case string:
		h := len(v) / 2
		fmt.Println("a swapped by halves is", v[h:]+v[:h])
	default:
		// a isn't one of the types above
	}
}

