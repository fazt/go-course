package operators

import "fmt"

// AssignmentOperators ...
func AssignmentOperators() {
	var x = 30

	x = x + 10
	fmt.Println(x)

	x += 10
	fmt.Println(x)

	x -= 10
	fmt.Println(x)

	x *= 10
	fmt.Println(x)

	x /= 10
	fmt.Println(x)

	x %= 10
	fmt.Println(x)
}
