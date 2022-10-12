package variables

import (
	"fmt"
)

func DeclareVariables() {
	// declare a variable
	var message string
	message = "My secret message"
	fmt.Println(message)

	// declaring and assigning a variable
	var myName string = "Fazt"
	fmt.Println(myName)

	// you can do this with any type
	var price float32 = 30.5
	fmt.Println(price)

	var isCool bool = true
	fmt.Println(isCool)

	var year int = 2022
	fmt.Println(year)

	// error if not use the variable
	// var n = 2020
	// var j int = 10

	// multiple declarations
	var name, lastname, age = "Joe", "Harper", 33
	fmt.Println(name, lastname, age)

	var (
		username = "fazt"
		email    = "fazt@mail.com"
		password = "123456"
	)
	fmt.Println(username, email, password)

}
