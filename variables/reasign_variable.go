package variables

import "fmt"

func reasignVariable() {
	// reassign a variable
	var year int = 2022
	year = 2023
	fmt.Println(year)

	var city string = "Londong"
	city = "Paris"
	fmt.Println(city)

	// reasign with another type
	// var b int = 200
	// b = "30" // error
}
