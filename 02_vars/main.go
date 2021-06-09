package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// var - Don't technically need to add string/int
	var name string = "codefluorite"
	var number int32 = 100
	var isCool = true
	isCool = false

	// shorthand

	flamingo_height := 4.5

	flamingo, email := "flora", "flora@codefluorite.com"

	fmt.Println(name, number, isCool, flamingo, email)
	fmt.Printf("%T\n", flamingo_height)
}
