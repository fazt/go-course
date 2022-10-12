package conditionals

import "fmt"

func ElseIf() {

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
}
