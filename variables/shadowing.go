package variables

import (
	"fmt"
)

// variable in package level
var i int = 30

func shadowing() {
	fmt.Println(i)
	// var i int = 70
	i := 70
	fmt.Println(i)
}
