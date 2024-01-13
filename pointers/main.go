package main

import "fmt"

func main() {
	A := new(int)
	B := 22

	A = &B

	fmt.Println(*A)

	B = 15

	fmt.Println(*A)

}
