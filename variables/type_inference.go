package variables

import "fmt"	

func TypeInference() {
	// type inference
	// var name = "Ryan"
	// var lastname = "Ray"
	// var user_age = 33
	// var isAlive = true
	// var score float32 = 152.30

	name := "Ryan"
	user_age := 33
	isAlive := true
	score := 152.30

	fmt.Println(name)
	fmt.Println(isAlive)
	fmt.Println(score)
	fmt.Println(user_age)

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", isAlive)
	fmt.Printf("%T\n", score)
	fmt.Printf("%T\n", user_age)

}
