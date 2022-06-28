package functions

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func sum(x int, y int) {
	fmt.Println(x + y)
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func greeting(name string) string {
	return "Hello " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func swap(a, b string) (string, string) {
	return b, a
}

func operateTotal(total int) (x int, y int) {
	x = total + 20
	y = total - 50
	return
}

func Functions() {
	// First function
	hello()

	// Parameters
	sum(10, 20)
	sum2(40, 70)

	// return
	fmt.Println(greeting("fazt"))
	fmt.Println(add(10, 15))

	// multiple return
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// named return
	fmt.Println(operateTotal(10))
}
