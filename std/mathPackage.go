package std

import (
	"fmt"
	"math/rand"
	"time"
)

func MathPackage() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))      // random number between 0 and 10
	fmt.Println(rand.Float64())     // random float64 between 0 and 1
	fmt.Println(rand.Float64() * 5) // random float64 between 0 and 5
}
