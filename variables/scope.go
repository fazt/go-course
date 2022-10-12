package variables

import "fmt"

var nickname = "Fazt"

func main() {
	// declare variable outside a function
	fmt.Printf("%v %T\n", nickname, nickname)
}
