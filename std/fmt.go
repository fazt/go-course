package std

import "fmt"

func Fmt() {
	name := "Joe"
	fmt.Println("Hello", name, "Welcome")

	nickname := "Fazt"
	fmt.Printf("Hello %s, Welcome to my program!\n", nickname)

	// print format data with Struct
	type CategoryProduct struct {
		name  string
		price float32
	}

	laptop := CategoryProduct{
		name:  "Macbook Pro",
		price: 1299.99,
	}

	fmt.Printf("%v\n", laptop)
	fmt.Printf("%+v\n", laptop)

	// Print datatypes of variables
	username := "Fazt"
	age := 70
	isAlive := true
	score := 152.30

	// print the datatypes of variables
	fmt.Printf("%T\n", username)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isAlive)
	fmt.Printf("%T\n", score)

	// print the value and its datatypes
	fmt.Printf("%v %T\n", username, username)

}
