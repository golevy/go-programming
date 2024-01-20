package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("/Users/levylv/desktop/d..txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	_, err = file.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
	}

	var str = "!!"
	num, _ := file.Seek(0, io.SeekEnd) // 将文件指针移动到文件末尾
	// 0 是偏移量(offset)，表示相对于参考点的移动距离，在这个特定的调用中，偏移量设置为 0，意味着不需要从参考点移动
	// io.SeekEnd 是 Seek 方法的一个参考点(whence)，表示文件的末尾，这意味着偏移量会从文件末尾开始计算。
	fmt.Println("num= ", num) // 打印当前文件指针的位置

	n, err := file.WriteAt([]byte(str), num) // 第二个参数表示在指定位置写入数据
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
}
