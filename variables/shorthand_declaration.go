package variables

func Shorthand() {
	// shorthands declaration varaibles for functions
	// this cannot be declared outside function
	// fullName := "Joe McMillan"
	firstName := "Joe"
	println(firstName)

	lastName := "McMillan"
	println(lastName)

	// multiple declarations
	// emailOne := "fazt@faztweb.com"
	// emailTwo := "fazttech@gmail.com"
	emailOne, emailTwo := "fazt@faztweb.com", "fazttech@gmail.com"

	// you can print multiple variables too
	println(emailOne, emailTwo)
}
