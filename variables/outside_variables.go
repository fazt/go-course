package variables

import "fmt"

// variable in package level
var i int = 30

// or declared variable outside a function
var username string = "Fazt"

// grouping variables
var (
	characterName string = "Sherlock Holmes"
	proffession   string = "detective"
	age           int    = 30
	city              = "London" // type inference
)

// var theUrl string = "http://google.com"
var theURL string = "http://google.com"
var myHTTP string = ""

var (
	counter int = 0
)

func OutsideVariables() {
	fmt.Printf("%v %T\n", i, i)
	fmt.Printf("%v %T\n", username, username)
	counter++
	counter++
	counter++

	fmt.Println(characterName, proffession, age, city, theURL, myHTTP, counter)
}
