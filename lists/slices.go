package lists

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

	// append
	tasksSlice = append(tasksSlice, "I have to eat")
	fmt.Println(tasksSlice)

	// delete the index 1
	tasksSlice = append(tasksSlice[:1], tasksSlice[2:]...)
	fmt.Println(tasksSlice)

	// make
	tasksSlice2 := make([]string, 3)

	tasksSlice2[0] = "I have to run"
	tasksSlice2[1] = "Buy milk"
	tasksSlice2[2] = "walking the dog"
	tasksSlice2 = append(tasksSlice2, "I have to sleep")

	fmt.Println(tasksSlice2)
}
