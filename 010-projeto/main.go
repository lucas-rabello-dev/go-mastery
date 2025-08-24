package main

import "fmt"

func main() {
	
	var ( 
		nota1 = 5
		nota2 = 4
		nota3 = 8
		nota4 = 9
		media = (nota1 + nota2 + nota3 + nota4) / 4 
	)
	fmt.Printf(" A média de %d | %d | %d | %d é: %d \n", nota1, nota2, nota3, nota4, media)
}