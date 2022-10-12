package variables

import (
	"fmt"
)


func shadowing() {
	fmt.Println(i)
	// var i int = 70
	i := 70
	fmt.Println(i)
}
