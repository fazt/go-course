package arraysSlices

import (
	"fmt"
)

func Arrays() {
	// Array
	var users [2]string

	// assign values
	users[0] = "John"
	users[1] = "Jane"
	// users[2] = "Jack" // this will cause an error

	fmt.Println(users)
	fmt.Println(users[0])
	fmt.Println(users[1])

	// Declare and assign
	languagesArr := [3]string{"Javascript", "Go", "Python"}

	fmt.Println(languagesArr)
	fmt.Println(languagesArr[0])
	fmt.Println(languagesArr[1])
	fmt.Println(languagesArr[2])

	fmt.Println(len(languagesArr))
}
