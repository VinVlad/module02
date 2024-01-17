package main

import "fmt"

func contains(a []string, x string) bool {

	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func getMax(a ...int) int {
	x := 0
	for _, v := range a {
		if v > x {
			x = v
		}
	}
	return x
}

func main() {

	a := []string{"cлово", "буква", "число"}
	x := "число"

	fmt.Println(contains(a, x))

	fmt.Println(getMax(1, 2, 7, 5, 9, 4, -2, 2, 3))
}
