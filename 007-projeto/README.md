# Gerador de Números Aleatórios em Go

Este projeto em Go é uma aplicação de linha de comando (CLI) para geração de números aleatórios com base em diferentes tipos e tamanhos de dados, utilizando a biblioteca padrão `math/rand`.

## 📦 Funcionalidades

O programa permite a geração de:

- Números inteiros com e sem sinal:
  - `int32`, `int64`, `uint32`, `uint64`
- Números de ponto flutuante:
  - `float32`, `float64`
- Slices com permutação de inteiros (`[]int`)
- Suporte a geração com ou sem parâmetros (limites superiores)

---

## 📋 Menu Interativo

Ao executar o programa, um menu interativo será apresentado com as seguintes opções:

1 - Int32 | Sem parâmetro | Range: 0 até 2.147.483.647
2 - Int64 | Sem parâmetro | Range: 0 até 9.223.372.036.854.775.807
3 - Int32 | Com parâmetro | Range: 0 até parâmetro
4 - Int64 | Com parâmetro | Range: 0 até parâmetro
5 - UInt32 | Sem parâmetro | Range: 0 até 4.294.967.295
6 - UInt64 | Sem parâmetro | Range: 0 até 18.446.744.073.709.551.615
7 - Int | Com parâmetro | Range: 0 até parâmetro
8 - Float32 | Sem parâmetro | ~6–7 dígitos decimais
9 - Float64 | Sem parâmetro | ~15–16 dígitos decimais
10 - Slice de inteiros | Com parâmetro | Permutação de 0 até (parâmetro - 1)
11 - Sair

yaml
Copiar
Editar

---

## 🚀 Como executar

1. Certifique-se de ter o Go instalado ([instalar Go](https://golang.org/dl/)).

2. Clone o repositório ou salve o código em um arquivo `.go`, por exemplo `main.go`.

3. Compile e execute:

```bash
go run main.go
```

## 🧠 Conceitos utilizados
- Tipos numéricos em Go (int8, int32, uint64, etc.)

- Funções da biblioteca math/rand:

- rand.Int31(), rand.Int63n(n), rand.Float64(), etc.

- Funções com e sem parâmetros

- Interação via terminal com fmt.Scan

- Estruturas de controle: switch, for, case