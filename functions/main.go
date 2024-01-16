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

func main() {

	a := []string{"cлово", "буква", "число"}
	x := "число"

	fmt.Println(contains(a, x))
}
