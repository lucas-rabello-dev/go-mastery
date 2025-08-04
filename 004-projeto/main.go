package main

import (
	"fmt"
)

func ConverterCF(celcius float64) float64 {
	return (celcius * 1.8) + 32
}

func ConverterFC(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.8
}

func main() {

	celcius := 20.00

	fahrenheit := 68.00

	var result1 float64 = ConverterCF(celcius)
	fmt.Printf("O resultado aproximado de celcius para fahrenheit: %.2f \n", result1)

	var result2 float64 = ConverterFC(fahrenheit)
	fmt.Printf("O resultado aproximado de fahrenheit para celcius: %.2f \n", result2)
}