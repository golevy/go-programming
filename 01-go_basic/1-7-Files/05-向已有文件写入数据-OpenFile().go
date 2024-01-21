package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("/Users/levylv/Desktop/a.txt", os.O_APPEND|os.O_WRONLY, 0666)
	// 路径，追加模式，读权限与写权限
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	n, err := file.WriteString("Good Job")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("写入的数据长度为：", n)
}
