package rangePackage

import (
	"fmt"
)

func RangeIntro() {
	ids := []int{50, 100, 77, 10}

	// loop
	for i, id := range ids {
		fmt.Printf("id %d: %d\n", i, id)
	}

	// without using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sum ids
	sum := 0
	for _, id := range ids {
		sum += id
	}

	println(sum)

	// range with maps
	emails := map[string]string{"fazt": "fazt@faztweb.com", "jesus": "jesus@gmail.com", "ryan": "ryan@gmail.com"}

	// loop the keys
	for k := range emails {
		fmt.Println("Name:" + k)
	}

	// loop the values
	for _, v := range emails {
		fmt.Println("email:" + v)
	}

	// loop the keys and values
	for k, v := range emails {
		fmt.Printf("%s: \t\t%s\n", k, v)
	}

}
