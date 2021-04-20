package main

import (
	"fmt"
)

// declared variable outside a function
var nickname string = "Fazt"

func main() {
	// declaring a variable
	var i int
	i = 100
	fmt.Println(i)

	// reassing a value
	i = 177
	fmt.Println(i)

	// create and assign
	var a int = 100
	fmt.Println(a)

	// reasign with another type
	// var b int = 200
	// b = "fazt"

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

	// variables
	var name string = "Fazt"

	// redefine the variable
	name = "Ryan Ray"

	// error if not use the variable
	// var year = 2020

	// type inference
	var lastname = "Ray"
	var age = 33
	var isAlive = true
	isAlive = false
	var score float32 = 152.30

	fmt.Println(name)
	fmt.Println(lastname)
	fmt.Println(age)

	// print format
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isAlive)
	fmt.Printf("%T\n", score)

	// other format
	fmt.Printf("%v %T\n", name, name)

	// declare variable outside a function
	fmt.Printf("%v %T\n", nickname, nickname)

}
