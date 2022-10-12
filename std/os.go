package main

import (
	"fmt"
	"os"
)


func main() {

	fmt.Println(os.Args)

	if (len(os.Args) != 2) {
		fmt.Println("Please provide a name")
		return
	}

	argument := os.Args[1]

	fmt.Println("Hello", argument)
}
