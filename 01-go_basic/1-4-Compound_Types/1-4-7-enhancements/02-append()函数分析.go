package main

import "fmt"

func Modify(s []int) {
	for i := 0; i < 5; i++ {
		s = append(s, i+1)
	}
}

func main() {
	s := make([]int, 5)
	fmt.Println("这是初始的切片：", s) // [0 0 0 0 0]

	Modify(s)
	fmt.Println("这是修改后的切片：", s) // [0 0 0 0 0]
}

// 当切片容量不足的时候 append() 函数会动态分配一个新的内存
// 所以使用 append() 函数操作原有切片后，并不会对原有的切片产生影响
