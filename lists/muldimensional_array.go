package lists

import "fmt"

func main() {

	// multidimensional array
	var matrix [2][2]int

	matrix[0][0] = 1
	matrix[0][1] = 2

	matrix[1][0] = 3
	matrix[1][1] = 4

	fmt.Println(matrix)

	// declare and assign
	matrix2 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(matrix2)

	// short declaration
	matrix3 := [...][2]int{{1, 2}, {3, 4}}

	fmt.Println(matrix3)

	// range
	for i := 0; i < len(matrix3); i++ {
		for j := 0; j < len(matrix3[i]); j++ {
			fmt.Println(matrix3[i][j])
		}
	}

	// range
	for _, row := range matrix3 {
		for _, value := range row {
			fmt.Println(value)
		}
	}
}
