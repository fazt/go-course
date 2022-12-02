package conditionals

import "fmt"

func If() {
	busy := true

	if busy {
		fmt.Println("I am sad \u2639")
	}

	if !busy {
		fmt.Println("I am not married \u263A")
	}

	status := "slpeeping"

	if status == "working" {
		fmt.Println("Iam working")
	}

	if status == "free" {
		fmt.Println("talk to me")
	}

	if status == "sleeping" {
		fmt.Println("Do not disturb")
	}

}
