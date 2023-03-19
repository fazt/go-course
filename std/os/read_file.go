package std

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFiles() {
	// this line, open the file and return a file object, if the file doesn't exist, create it
	// and put the file object in the variable file
	// the second parameter is the flag, if the file doesn't exist, create it
	// the third parameter is the permission
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
