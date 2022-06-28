package variables

import (
	"fmt"
)

func DeclareVariables() {
	// declaring a variable
	var myName string = "Fazt"
	fmt.Println(myName)

	var age int = 30
	fmt.Println(age)

	// you can do this with any type
	var price float32 = 30.5
	fmt.Println(price)

	var isCool bool = true
	fmt.Println(isCool)

	// reassign a variable
	var year int = 2022
	year = 2023
	fmt.Println(year)

	var city string = "Londong"
	city = "Paris"
	fmt.Println(city)

	// reasign with another type
	// var b int = 200
	// b = "30"

	// type inference
	var name = "Ryan"
	var lastname = "Ray"
	var user_age = 33
	var isAlive = true
	var score float32 = 152.30

	isAlive = false

	fmt.Println(name)
	fmt.Println(lastname)
	fmt.Println(isAlive)
	fmt.Println(score)
	fmt.Println(user_age)

	// shorthands declaration varaibles for functions
	// this cannot be declared outside function
	// fullName := "Joe McMillan"
	firstName := "Joe"
	println(firstName)

	lastName := "McMillan"
	println(lastName)

	// multiple declarations
	// emailOne := "fazt@faztweb.com"
	// emailTwo := "fazttech@gmail.com"
	emailOne, emailTwo := "fazt@faztweb.com", "fazttech@gmail.com"

	// you can print multiple variables too
	println(emailOne, emailTwo)

	// error if not use the variable
	// var year = 2020

	// declare variable outside a function
	fmt.Printf("%v %T\n", nickname, nickname)

	// variables that are not used throw erro
	// var j int 30 // this trows an error

	// name convetions
	// variables that you

}
