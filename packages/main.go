package packagesTutorial

import (
	"fmt"

	"github.com/fazttech/go-course/packages/greet"
)

func SomeMainPackage() {
	fmt.Println(greet.Greet("Fazt"))
}
