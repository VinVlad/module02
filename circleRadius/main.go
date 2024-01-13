package main

import (
	"fmt"
	"math"
)

func main() {
	var l float64 = 35
	R := new(float64)

	r := (l / (math.Pi * 2))
	R = &r

	fmt.Println(*R)

	S := (math.Pi) * math.Pow(*R, 2)

	s := fmt.Sprintf("S = %.2f", S)
	fmt.Println(s)
}
