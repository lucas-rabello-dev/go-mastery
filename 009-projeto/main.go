package main

import (
	"fmt"
	s "strings"
)	

func main() {
	words := s.Fields("meu nome é lucas") // método que separa cada palavra e retorna um slice com cada palvra
	var count int
	for _, word := range words {
		count++
		fmt.Printf("palavra: %s\n", word)
	}
	fmt.Println("número total de palvras:", count)
}