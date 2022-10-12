package functions

import "fmt"

func sum(x int, y int) {
	fmt.Println(x + y)
}

func sum2(x, y int) {
	fmt.Println(x + y)
}

func Parameters() {
	sum(10, 20)
	sum2(40, 70)
}
