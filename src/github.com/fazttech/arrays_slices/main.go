package main

import (
	"fmt"
)

func main() {
	// Array
	var fruitArr [2]string

	// assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	// Declare and assign

	languagesArr := [3]string{"Javascript", "Go", "Python"}
	fmt.Println(languagesArr)
	fmt.Println(languagesArr[0])
	fmt.Println(languagesArr[1])
	fmt.Println(languagesArr[2])

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	// length
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[0])
	fmt.Println(fruitSlice[1])
	fmt.Println(fruitSlice[2])

	// range
	fmt.Println(fruitSlice[1:3])

}
