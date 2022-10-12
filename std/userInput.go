package main

import "fmt"

func main() {
	var name string
	fmt.Print("What is your name?: ")
	fmt.Scan(&name)
	fmt.Println(name)
}
