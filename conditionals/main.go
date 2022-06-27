package ifelse

import "fmt"

func Switch() {

	color := "orange"

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
