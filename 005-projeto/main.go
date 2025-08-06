package main

import (
    "fmt"
)

func main() {
    var input int
    fmt.Print("Digite um número: ")
    fmt.Scan(&input)
    for i := 1; i <= 10; i++ {
        var resul int = input * i
        fmt.Printf("%d x %d = %d \n", input, i,resul)
    }
}
