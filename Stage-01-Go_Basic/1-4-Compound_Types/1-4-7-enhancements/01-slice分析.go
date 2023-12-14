package main

import (
	"fmt"
	"unsafe"
)

func ModifySlice(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
}

func main() {
	s := make([]int, 5)
	fmt.Println(s)

	ModifySlice(s)
	fmt.Println(s)
}

// 在 Go SDK src/runtime/slice.go 给出了切片的定义
// 在 package unsafe 里 type Pointer *ArbitraryType 定义的是一个指针
// 当你在函数中修改切片的内容时，实际上是在修改切片所引用的底层数组
// 所以就不难理解为什么在函数内对传过来的切片进行修改后，原来的切片也跟着被修改了
// 下面是关键代码
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
