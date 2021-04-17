package main

import "fmt"

func main() {
	// Main TYpes

	// sting

	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float 64
	// complex64 complex128

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

	// shorthands declaration varaibles for functions
	fullName := "Joe McMillan"
	println(fullName)

	// multiple declarations
	// emailOne := "fazt@faztweb.com"
	// emailTwo := "fazttech@gmail.com"
	emailOne, emailTwo := "fazt@faztweb.com", "fazttech@gmail.com"
	println(emailOne, emailTwo)

}

// it cannot be declared outside function
// fullName := "Joe McMillan"
