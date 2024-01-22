package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 1. 打开原有文件 (file.txt)
	file1, err := os.Open("/Users/levylv/Desktop/file1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. 创建一个新的文件
	file2, err := os.Create("/Users/levylv/Desktop/file2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 关闭文件
	defer func(file1 *os.File) {
		err := file1.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file1)

	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(file2)

	// 3. 将原有文件中的内容读取出来，然后写入到新的文件中
	buffer := make([]byte, 1024*2)

	for {
		n, err := file1.Read(buffer)
		if err == io.EOF {
			break
		}

		_, err = file2.Write(buffer[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
