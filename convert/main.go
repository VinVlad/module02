package main

import (
	"fmt"
	"strconv"
)

func main() {

	var strValue string = "104"
	var digitNumber int = 35

	strValueConv, _ := strconv.Atoi(strValue)
	digitNumberConv := strconv.Itoa(digitNumber)

	fmt.Println(strValueConv)
	fmt.Println(digitNumberConv)

}
