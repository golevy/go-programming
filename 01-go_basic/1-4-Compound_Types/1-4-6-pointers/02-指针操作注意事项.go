package main

import "fmt"

func main() {
	// 1. 不要使用空指针
	var p1 *int
	fmt.Println(p1) // <nil>

	// 2. 不要操作没有合法指向的内存
	// var p2 *int
	// *p2 = 56
	// fmt.Println(p2) // panic: runtime error: invalid memory address or nil pointer dereference

	// 3. new() 函数的使用
	var p3 *int
	p3 = new(int) // 开辟数据类型 (int) 对应的内存空间，返回值为该数据类型 (int) 的指针
	*p3 = 57
	fmt.Println(*p3)
}
