package std

import "fmt"

func UserInput() {
	var name string

	fmt.Print("What is your name?: ")
	fmt.Scan(&name)

	fmt.Println(name)
}
