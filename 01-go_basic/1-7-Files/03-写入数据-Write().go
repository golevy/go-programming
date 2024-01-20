package main

import (
	"fmt"
	"os"
)

// WriteString() 调用的就是 Write() 方法

//func (f *File) WriteString(s string) (n int, err error) {
//	b := unsafe.Slice(unsafe.StringData(s), len(s))
//	return f.Write(b)
//}

func main() {
	file, err := os.Create("/Users/levylv/desktop/c.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	var str = "Hello World"
	n, err := file.Write([]byte(str)) // 需要将字符串转换成字节切片
	if err != nil {
		return
	}
	fmt.Printf("写入的数据长度为 %d\n", n)

}
