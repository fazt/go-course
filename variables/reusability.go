package variables

import "fmt"

func Reusability() {
	var x int = 100

	fmt.Println("x =", x)
	fmt.Println("x + 100 =", x + 100)
	fmt.Println("x - 50 =", x - 50)
	fmt.Println("x * 2 =", x * 2)
	fmt.Println("x / 2 =", x / 2)

	fmt.Println("x =", x)

	// reassigning a variable
	var y int = 100

	fmt.Println("\ny =", y)
	y = y + 100
	y = y - 50
	y = y * 2
	y = y / 2
	fmt.Println("y =", y)
}
