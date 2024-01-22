package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. 使用 os.Open() 以只读的形式打开文件（实际上 Open() 方法底层就是以只读的方式调用 OpenFile() 方法）
	file, err := os.Open("/Users/levylv/Desktop/c.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 2. 定义一个字符类型切片，存储从文件中读取的数据
	buffer := make([]byte, 1024*2) // 创建一个长度为 2KB（2048 字节）的字节切片，意味着这个切片可以存储最多 2KB 的数据
	n, err := file.Read(buffer)    // file.Read 方法会将从文件中读取的数据写入到 buffer 切片中
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(buffer) // 打印出来的都是字节数据（例如：[72 101 108 108 111 32 87 111 114 108]）
	// 如果这些字节对应于可打印的字符，在控制台上可以看到这些字符
	// 如果字节不对应于可打印字符，会看到一些看似随机或乱码的输出

	fmt.Println(string(buffer[:n]))
	fmt.Println("读取的数据长度为；", n)

	// 3. 关闭文件
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
}
