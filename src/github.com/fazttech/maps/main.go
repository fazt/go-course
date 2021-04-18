package main

import "fmt"

func main() {
	// define a map
	emails := make(map[string]string)

	// assign values
	emails["fazt"] = "fazt@faztweb.com"
	emails["jesus"] = "jesus@faztweb.com"
	emails["joe"] = "joe@faztweb.com"

	fmt.Println(emails)

	// len
	fmt.Println(len(emails))

	// access values
	fmt.Println(emails["fazt"])

	// delete from map
	delete(emails, "jesus")
	fmt.Println(emails)

	// declare and assign values
	scores := map[string]int{"fazt": 50, "jesus": 100, "joe": 70}

	fmt.Println(scores)

}
