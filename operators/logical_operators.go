package operators

import "fmt"

func LogicalOperator() {
	// logical operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// logical operators with relational operators
	fmt.Println(10 == 10 && 10 > 5)  // true
	fmt.Println(10 == 10 || 10 > 11) // true
}
