package main

import "fmt"

func main() {
	// Converts Fahrenheit into Celsius
	fmt.Println("Enter temperature in Fahrenheit")
	var temperature_F float64
	fmt.Scanf("%f", &temperature_F)
	temperature_C := (temperature_F - 32) * 5 / 9
	fmt.Println("Temperature in Celsius", temperature_C)
}
