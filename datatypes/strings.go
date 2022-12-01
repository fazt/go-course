package datatypes

import "fmt"

func Strings() {

	// a simple string in go
	fmt.Println("Hello World")

	// we can store in variables
	var courseName string = "Go Course"
	fmt.Println(courseName)

	// you can see the type of variable with %T
	fmt.Printf("%T\n", courseName)

	// in go we can use the backtick ` to create multiline strings
	var multiline string = `
	Hello World
	from Go Course
	`
	fmt.Println(multiline)

	// length of string
	fmt.Println(len(courseName))

	// access to a character
	fmt.Println(courseName[0])         // 71, the ascii code of G
	fmt.Println(string(courseName[0])) // G

	fmt.Println(courseName[0:2]) // Go
	fmt.Println(courseName[:2])  // Go
	fmt.Println(courseName[3:])  // Course
	fmt.Println(courseName[:])   // Go Course

	// we can use the + operator to concatenate strings
	fmt.Println("Hello " + "World")    // Hello World
	fmt.Println("Hello " + courseName) // Hello Go Course

	// we can use the += operator to concatenate strings
	var hello string = "Hello "
	hello += "World"
	hello += " !"
	hello += " from Go Course"
	fmt.Println(hello) // Hello World ! from Go Course

	// interpolation
	fmt.Printf("Welcome to %s!\n", courseName) // Welcome to Go Course!

	// interpolation with numbers
	fmt.Printf("The number is %d\n", 10) // The number is 10

	// interpolation with floats
	fmt.Printf("The float is %f\n", 10.5) // The float is 10.500000

	// interpolation with booleans
	fmt.Printf("The boolean is %t\n", true) // The boolean is true

	// interpolation with characters
	fmt.Printf("The character is %c\n", 71) // The character is G

	// interpolate multiple values
	var age = 30
	var name = "John"
	var married = false

	fmt.Printf("%s is %d years old and it is %t that he is married\n", name, age, married)

	// utf-8
	fmt.Println("Hello \xF0\x9F\x98\x8E")
	fmt.Println("Hello \U0001F60E")
	fmt.Println("Hello \U0001F915")
	fmt.Println("Hello 😴")

}
