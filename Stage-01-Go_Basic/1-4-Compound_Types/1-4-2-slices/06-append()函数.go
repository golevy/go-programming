package main

import "fmt"

func main() {
	slice := make([]int, 5, 8)
	fmt.Println("slice 的初始结果为：", slice)
	fmt.Println("slice 的初始长度为：", len(slice))
	fmt.Println("slice 的初始容量为：", cap(slice))

	// append() 函数的作用就是向切片末尾追加数据
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4) // 超出原有容量后，切片就会自动扩容
	fmt.Println("slice 的数据追加结果为：", slice)
	fmt.Println("slice 的长度为：", len(slice))
	fmt.Println("slice 的容量为：", cap(slice))
}
