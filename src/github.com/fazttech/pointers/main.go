package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)

	// read values from addresses
	fmt.Println(*b)
	fmt.Println(*&a)

	// change value with pointers
	*b = 10

	fmt.Println(a)

}
