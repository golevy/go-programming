package main

import (
	"fmt"
	"os"
)

func main() {
	// Create file
	file, err := os.Create("/Users/levylv/Desktop/a.txt")

	if err != nil {
		fmt.Println(err)
	}

	// Schedule the closure to close the file at the end of the main function
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file) // Immediately call the anonymous function with 'file' as the argument
}
