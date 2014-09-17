package main

import "fmt"

func main() {
	fmt.Print("Enter Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * (5.0 / 9.0)

	fmt.Printf("Celsius: %f\n", celsius)
}
