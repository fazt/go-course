package variables

import (
	"fmt"
	"strconv"
)

func otherVars() {
	var i int = 42
	fmt.Printf("%v %T\n", i, i)

	// use the float32 function to convert
	var j float32
	j = float32(i)
	fmt.Printf("%v %T\n", j, j)

	// but if convert from float to int, you can lose data
	var f float32 = 455.78
	fmt.Printf("%v %T\n", f, f)

	var g int = int(f)
	fmt.Printf("%v %T\n", g, g)

	// convert int to string
	var h int = 70
	fmt.Printf("%v %T\n", h, h)

	// this convert to the unicode equivalent
	var k string
	k = string(h)
	fmt.Printf("%v %T\n", k, k)

	// using the strconv package
	var l string
	l = strconv.Itoa(h)
	fmt.Printf("%v %T\n", l, l)

}
