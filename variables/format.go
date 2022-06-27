package variables

import (
	"fmt"
)

func FormatVariables() {

	name := "Fazt"
	age := 70
	isAlive := true
	score := 152.30

	// print the datatypes of variables
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isAlive)
	fmt.Printf("%T\n", score)

	// print the variables and its datatypes
	fmt.Printf("%v %T\n", name, name)
}
