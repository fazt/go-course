package functions

import "fmt"

var p = fmt.Println

func Alias() {
	p("Hello World")
	p("Hello World 2")
	p("Hello Go!")
}
