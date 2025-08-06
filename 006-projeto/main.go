package main

import (
    "fmt"
    "strconv"
)

// function to reverse a string 
func reverseString(s string) string {
    runes := []rune(s)             // convert string to bytes
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i] // invert
    }
    return string(runes)
}

// func para capturar erro de tipo
func ErrorTypeInt(num int)  error {
    test_convert := strconv.Itoa(num)

	// converto novamente para int para tentar capturar o erro usando uma função já pronta
	_, err  := strconv.Atoi(test_convert)
	if err != nil {
		return err
	}
	 // Só para que o Go não reclame de eu não usar a variavel 'convert'
	return nil
}

func Palindromo(n int) (bool, error) {
    var num string = strconv.Itoa(n)

    realInvert := reverseString(num)

    invertNum, err := strconv.Atoi(realInvert)
    if err != nil {
        return false, err
    }

    if  invertNum == n {
        return true, nil
    }
    return false, nil
}

func main() {
    var input int
    fmt.Print("Digite um número: ")
    fmt.Scan(&input)
    // Tentativa de captura de erro caso o input não seja um número válido
    err := ErrorTypeInt(input)
    if err != nil {
        panic(err)
    }

    result, err := Palindromo(input)
    if err != nil {
        panic(err)
    }

    if result == true {
		fmt.Println("é um palindromo!")
    	return
    }
    fmt.Println("não é palindromo!")

}