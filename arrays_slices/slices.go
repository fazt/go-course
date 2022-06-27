package arraysSlices

import (
	"fmt"
)

func Slices() {
	// Slice
	tasksSlice := []string{"I have to run", "Buy milk", "walking the dog", "I have to sleep"}

	// length
	fmt.Println(len(tasksSlice))
	fmt.Println(tasksSlice[0])
	fmt.Println(tasksSlice[1])
	fmt.Println(tasksSlice[2])

	// range
	fmt.Println(tasksSlice[1:3])
}
