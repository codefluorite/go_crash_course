package main

import "fmt"

func main() {
	x := 15
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if

	color := "green"

	if color == "purple" {
		fmt.Println("colour is purple")
	} else if color == "blue" {
		fmt.Println("colour is blue")
	} else {
		fmt.Println("colour is NOT blue or red")
	}

	// Switch

	switch color {
	case "purple":
		fmt.Println("colour is purple")
	case "blue":
		fmt.Println("colour is blue")
	default:
		fmt.Println("colour is NOT purple or blue")
	}
}
