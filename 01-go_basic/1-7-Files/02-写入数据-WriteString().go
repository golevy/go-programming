package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("/Users/levylv/desktop/b.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()

		if err != nil {
			fmt.Println(err)
		}
	}(file)

	n, err := file.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("写入的数据长度为 %d\n", n) // 11
}
