package pointers

import (
	"fmt"
)

func Pointers() {
	a := 5  // value
	b := &a // pointer to a

	fmt.Println(a, b) // 5 0xc0000b4008

	// read values from addresses
	fmt.Println(*b)  // 5
	fmt.Println(*&a) // 5

	// change value with pointers
	*b = 10

	fmt.Println(a)
}
