package datatypes

import (
	"fmt"
	"reflect"
	"strconv"
)

func TypeConversion() {
	var a int = 42
	fmt.Printf("%v %T\n", a, a) // 42 int

	// use the float32 function to convert
	var b float32
	b = float32(a)
	fmt.Printf("%v %T\n", b, b) // 42 float32

	// but if convert from float to int, you can lose data
	var f float32 = 455.78
	fmt.Printf("%v %T\n", f, f)

	var g int = int(f)
	fmt.Printf("%v %T\n", g, g)

	// convert int to string
	var h int = 70
	fmt.Printf("%v %T\n", h, h)

	// this convert to the unicode equivalent
	var i string
	i = string(h)
	fmt.Printf("%v %T\n", i, i) // F string

	// using the strconv package
	var j int = 100
	// var k string = string(j)
	var k string = strconv.Itoa(j)
	fmt.Printf("%v %T\n", k, k) // 100 string

	var l = "10033"
	var m, err = strconv.Atoi(l)
	fmt.Println(m, err, reflect.TypeOf(m)) // 10033 <nil> int

}
