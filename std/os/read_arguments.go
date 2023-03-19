package std

import (
	"fmt"
	"os"
)

func OsArguments() {

	fmt.Println(os.Args)

	if len(os.Args) != 2 {
		fmt.Println("Please provide a name")
		return
	}

	argument := os.Args[1]

	fmt.Println("Hello", argument)

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
}
