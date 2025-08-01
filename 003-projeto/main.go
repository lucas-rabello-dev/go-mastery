package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	
	fmt.Println("Digite um número: ")
	fmt.Scan(&input)
	num, err := strconv.Atoi(input) // A conversão de string foi feita para identificar pelo strconv.Atoi() se o input é um número válido
	if err != nil {
		fmt.Println("Você deve digitar um número válido!")
		return
	}
	// Verifica se o número é impar ou par
	if num % 2 == 0 {
		fmt.Printf("O número %d é par! \n", num)
		return
	}
	fmt.Printf("O número %d é impar! \n", num)
}