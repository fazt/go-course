package conditionals

import "fmt"

func Else() {
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}
