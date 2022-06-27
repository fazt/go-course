package functions

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1, num2 int) int {
	return num1 + num2
}

func Functions() {
	fmt.Println(greeting("fazt"))
	fmt.Println(add(10, 15))
	fmt.Println(add2(20, 20))
}
