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

	buffer := make([]byte, 10) // "buffer"（缓冲区）是一个常用术语，用来指代一个临时的存储区，它用于在数据在两个地方之间传输时暂存这些数据

	for {
		n, err := file.Read(buffer)
		if err == io.EOF { // 表示达到文件末尾了，意味着文件已经被完全读取
			break
		}

		fmt.Printf(string(buffer[:n]))
	}
}
