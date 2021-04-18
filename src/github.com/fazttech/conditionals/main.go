package main

import "fmt"

func main() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if

	color := "orange"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else if color == "yellow" {
		fmt.Println("color is yellow")
	} else {
		fmt.Println("I don't know this color")
	}

	// Switch

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	case "yellow":
		fmt.Println("color is yellow")
	default:
		fmt.Println("I don't know this color")
	}

}
