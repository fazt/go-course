package main

import "fmt"

func helloWorld() string {
	return "Hello World"
}

func main() {
	// Main TYpes

	// string
	fmt.Println("Hello World")
	fmt.Printf("%T\n", "Hello World")

	// bool
	fmt.Println(true)
	fmt.Println(false)

	fmt.Printf("%T\n", true)

	// int
	// int int8 int16 int32 int64
	fmt.Println(33)
	fmt.Printf("%T\n", 33)

	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for uint32
	// https://golang.org/pkg/builtin/#int16

	// float32 float 64
	fmt.Println(100.44)
	fmt.Printf("%T\n", 100.44)

	// complex64 complex128

	// Derivated Types

	// Functions
	fmt.Println(helloWorld())

}
