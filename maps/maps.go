package maps

import (
	"fmt"
)

func Maps() {
	// define a map
	students := make(map[string]string)

	// assign values
	students["fazt"] = "fazt@faztweb.com"
	students["jesus"] = "jesus@faztweb.com"
	students["joe"] = "joe@faztweb.com"
	fmt.Println(students)

	scores := make(map[string]int)

	scores["fazt"] = 100
	scores["jesus"] = 90
	scores["joe"] = 80
	fmt.Println(scores)

	// len
	fmt.Println(len(students))

	// access values
	fmt.Println(scores["fazt"])

	// delete from map
	delete(students, "jesus")
	fmt.Println(students)

	// declare and assign values
	scores2 := map[string]int{"fazt": 50, "jesus": 100, "joe": 70}

	fmt.Println(scores2)

}
