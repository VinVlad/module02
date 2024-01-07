package main

import "fmt"

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	var EuroSpeed EuropeanVelocity = 120.4 * 3.6
	var AmericanSpeed AmericanVelocity = 130 * 2.237

	fmt.Println(EuroSpeed)
	fmt.Println(AmericanSpeed)

}
