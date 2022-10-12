package functions

import "fmt"

func helloWorld() string {
	return "Hello World"
}

// return int
func add(num1 int, num2 int) int {
	return num1 + num2
}

// return string with parameter
func greeting(name string) string {
	return "Hello " + name
}

// return multiple parameters
func swap(a, b string) (string, string) {
	return b, a
}

func operateTotal(n int) (x int, y int){
	x = n + 100	
	y = n + 500
	return
}

func Return() {
	fmt.Println(helloWorld())
	fmt.Println(greeting("fazt"))
	fmt.Println(add(10, 15))

	fmt.Println(swap("hello", "world"))

	lastname, name := swap("Joe", "Harper")
	fmt.Println(lastname, name)

	// named return
	fmt.Println(operateTotal(10))
}
