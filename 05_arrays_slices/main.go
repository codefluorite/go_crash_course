package main

import "fmt"

func main() {
	// Arrays - must be a fixed length
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Pear"
	// fruitArr[1] = "Watermelon"

	// Declare and assign

	// fruitArr := [2]string{"Pear", "Watermelon"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Pear", "Watermelon", "Dragonfruit"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}
