package datatypes

import "fmt"

func ShowTypes() {
	// Literals
	fmt.Println("Hello World") // string
	fmt.Println(true) 				// boolean
	fmt.Println(33) 					// integer
	fmt.Println(100.44) 			// float

	fmt.Println()

	// data types
	fmt.Printf("%T\n", "Hello World")
	fmt.Printf("%T\n", true)
	fmt.Printf("%T\n", 33)
	fmt.Printf("%T\n", 100.44)
}
