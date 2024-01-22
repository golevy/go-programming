package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/Users/levylv/Desktop/c.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	buffer := make([]byte, 10)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF { // 表示达到文件末尾了，意味着文件已经被完全读取
			break
		}

		fmt.Printf(string(buffer[:n]))
	}
}
