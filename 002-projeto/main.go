package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"log"
)

func InputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInputInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	
	firstInput := InputString("Digite uma das seguintes operações (+) (-) (*) (/): ")
	if firstInput != "+" && firstInput != "-" && firstInput != "*" && firstInput != "/" {fmt.Println("Opção inválida!") 
	return} 

	firstNum := ReadInputInt("Digite o primeiro número: ")

	secondNum := ReadInputInt("Digite o segundo número: ")

	switch firstInput{
	case "+":
		resul := firstNum + secondNum
		fmt.Println("O resultado é:", resul)
	case "-":
		resul := firstNum - secondNum
		fmt.Println("O resultado é:", resul)
	case "*":
		resul := firstNum * secondNum
		fmt.Println("O resultado é:", resul)
	case "/":
		resul := firstNum / secondNum
		fmt.Println("O resultado é:", resul)
	}

	
}