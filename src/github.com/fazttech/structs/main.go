package main

import (
	"fmt"
	"strconv"
)

// defint a person struct
type Person struct {
	// firstName string
	// lastName  string
	// gender    string

	firstName, lastName, gender string
	age                         int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I have: " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (value receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// init person using struct
	person1 := Person{firstName: "jesus", lastName: "McMillan", age: 30, gender: "m"}

	// second person
	person2 := Person{"Jane", "McMillan", "f", 40}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.firstName)

	person1.age++
	fmt.Println(person1)

	// fmt.Println(greet(person1))
	fmt.Println(person1.greet())

	// use pointer method
	person1.hasBirthday()

	fmt.Println(person1.age)

	person2.getMarried("Williams")
	fmt.Println(person2.greet())

}
