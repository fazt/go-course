package lists

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
	users := map[string]int{"fazt": 50, "jesus": 100, "joe": 70}

	fmt.Println(users)

	// check if key exists
	_, ok := users["fazt"]

	if ok {
		fmt.Println("user exists")
	} else {
		fmt.Println("user doesn't exists")
	}

	// check if key exists
	if _, ok := users["jose"]; ok {
		fmt.Println("Key exists")
	} else {
		fmt.Println("Key doesn't exists")
	}

	// check if key does not exists
	if _, ok := users["jose"]; !ok {
		fmt.Println("Key doesn't exists")
	} else {
		fmt.Println("Key exists")
	}
}
