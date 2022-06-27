package pointers

import (
	"fmt"
)

func Pointers() {
	a := 5  // value
	b := &a // pointer to a

	fmt.Println(a, b)

	// read values from addresses
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointers
	*b = 10

	fmt.Println(a)
}
