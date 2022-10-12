package functions

import "fmt"

func hello() {
	fmt.Println("Hello")
}

// func operateTotal(total int) (x int, y int) {
// 	x = total + 20
// 	y = total - 50
// 	return
// }

func Functions() {
	// First function
	hello()
	hello()
	hello()
	hello()
	hello()

	// named return
	fmt.Println(operateTotal(10))
}
