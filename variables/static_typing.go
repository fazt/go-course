package variables

import "fmt"

func StaticTyping() {
	// var b int = 200
	// b = "fazt" // error
	// fmt.Println(b)

	var name string = "Fazt"
	var age int = 30
	var score float32 = 30.5
	var isAlive bool = true

	// print the datatypes of variables
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isAlive)
	fmt.Printf("%T\n", score)

	// print the variables and its datatypes
	fmt.Printf("%v %T\n", name, name)
	fmt.Printf("%v %T\n", age, age)
	fmt.Printf("%v %T\n", isAlive, isAlive)
	fmt.Printf("%v %T\n", score, score)
}
