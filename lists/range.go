package lists

import (
	"fmt"
)

func Range() {

	ids := []int{50, 100, 77, 10}
	for i, id := range ids {
		fmt.Printf("id %d: %d\n", i, id)
	}

	// range a slice without index
	points := []int{15, 20, 25, 30}

	for _, point := range points {
		fmt.Println(point)
	}

	// range and calculate sum
	expenses := []int{14, 20, 100}
	sum := 0
	for _, num := range expenses {
		sum += num
	}
	fmt.Println("total expenses:", sum)

	// range with condition
	ages := []int{35, 20, 15, 30}
	for i, age := range ages {
		if age <= 15 {
			fmt.Println(i, ": The age", age, "is not allowed")
		}
		fmt.Println(i, ": The age", age, "is allowed")
	}

	// range a map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}
	// range
	// for key, value := range students{
	// 	fmt.Println(key, value)
	// }

	// range a string
	for i, c := range "go" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}
