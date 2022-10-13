package lists

import (
	"fmt"
)

func Arrays() {
	// Array
	var users [2]string
	fmt.Println(users)

	// assign values
	var favoriteFruits [3]string
	favoriteFruits[0] = "apple"
	favoriteFruits[1] = "orange"
	favoriteFruits[2] = "pineapple"
	fmt.Println(favoriteFruits)

	// Declare and assign
	programmingLanguages := [3]string{"Javascript", "Go", "Python"}
	fmt.Println(programmingLanguages)

	// length
	fmt.Println(len(programmingLanguages))

	// short declaration
	friends := [3]string{"fazt", "jesus", "joe"}
	fmt.Println(friends)

	// ellipsis
	brothers := [...]string{"fazt", "jesus", "joe", "jose"}
	fmt.Println(brothers)

	// access values
	cities := [3]string{"New York", "London", "Tokyo"}
	fmt.Println(cities[0])
	fmt.Println(cities[1])
	fmt.Println(cities[2])

	// range
	notes := [...]string{"do", "re", "mi", "fa", "sol", "la", "si"}
	fmt.Println(notes[2:4])
}
