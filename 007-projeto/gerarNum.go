package main

import (
	"fmt"
	"math/rand"
)
/*
int8 | 8 bits  | -128 to 127
int1 | 16 bits  | -32,768 to 32,767
int32 | 32 bits  | -2,147,483,648 to 2,147,483,647
int64 | 64 bits  | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
*/

/*
uint8 | 8 bits  | 0 to 255
uint1 | 16 bits  | 0 to 65,535
uint32 | 32 bits  | 0 to 4,294,967,295
uint64 | 64 bits  | 0 to 18,446,744,073,709,551,615
*/

// No parametres
func randInt32() int32 {
	return rand.Int31()
}
// No parametres 
func randInt64() int64 {
	return rand.Int63()
}

// Parameters
func randInt32n(num int32) int32 {
	return rand.Int31n(num)
}

// Parameters
func randInt64n(num int64) int64 {
	return rand.Int63n(num)	
}


// No Parameters
func randUInt32() uint32 {
	return rand.Uint32()
}

// No Parameters
func randUInt64() uint64 {
	return rand.Uint64()	
}

// Use more memory
func randIntNoSize(num int) int {
	return rand.Intn(num)
}

/*
float32  32 bits | 4 bytes | ~6–7 decimal digits

flota64  64 bits | 8 bytes | ~15–16 decimal digits
*/

// No parameters
func randFloat32() float32 {
	return rand.Float32()
}

// No parameters
func randFloat64() float64 {
	return rand.Float64()
}

// Interval 0 - n and n numbers 
func randSliceInt(num int) []int {
	return rand.Perm(num)
}

func main() {
	
	/*
	os inputs de string são convertidos para 
	*/

	var ( 
		input int8 // For use lass memory
		parameterInt int
		parameterInt32 int32
		parameterInt64 int64
	)
	for {
		fmt.Println(
		` 
		1- Int32                 | No parameters   | Range = 0 to 2,147,483,647 
		2- Int64                 | No parameters   | Range = 0 to 9,223,372,036,854,775,807
		3- Int32                 | Parameters      | Range = 0 to Parameter
		4- Int64                 | Parameters      | Range = 0 to Parameter
		5- uInt32                | No parameters   | Range = 0 to 4,294,967,295
		6- uInt64                | No parameters   | Range = 0 to 18,446,744,073,709,551,615
		7- Int (only)            | Parameters      | Range = 0 a Parameter
		8- Float32               | No parameters   | Range = ~6–7 decimal digits
		9- Float64               | No parameters   | Range = ~15–16 decimal digits
		10- Slice of numbers int | Parameters      | Range = 0 to (parameter - 1) 
		11- Exit
		`)
		fmt.Print("Escolha uma das opções: ")
		fmt.Scan(&input)
		
		switch input {
		case 1:
			fmt.Println("-------------------------------",randInt32(),"-------------------------------")
		case 2:
			fmt.Println("-------------------------------",randInt64(),"-------------------------------")
		case 3:
			fmt.Print("Enter the range: ")
			fmt.Scan(&parameterInt32)
			fmt.Println("-------------------------------",randInt32n(parameterInt32),"-------------------------------")
		case 4:
			fmt.Print("Enter the range: ")
			fmt.Scan(&parameterInt64)
			fmt.Println("-------------------------------",randInt64n(parameterInt64),"-------------------------------")
		case 5:
			fmt.Println("-------------------------------",randUInt32(),"-------------------------------")
		case 6:
			fmt.Println("-------------------------------",randUInt64(),"-------------------------------")
		case 7:
			fmt.Print("Enter the range: ")
			fmt.Scan(&parameterInt)
			fmt.Println("-------------------------------",randIntNoSize(parameterInt),"-------------------------------")
		case 8:
			fmt.Println("-------------------------------",randFloat32(),"-------------------------------")
		case 9:
			fmt.Println("-------------------------------",randFloat64(),"-------------------------------")
		case 10:
			fmt.Print("Enter the range: ")
			fmt.Scan(&parameterInt)
			fmt.Println("-------------------------------",randSliceInt(parameterInt),"-------------------------------")
		case 11:
			return
		default:
			fmt.Println("Enter a valited option!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		}
	}
}