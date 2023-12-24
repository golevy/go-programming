package main

import "fmt"

func main() {
	var a = 10

	// 指针定义
	var p *int
	p = &a                               // & 是一种获取变量内存地址的符号
	fmt.Printf("变量 a 的内存地址为：%p\n", &a)   // 0x14000110018
	fmt.Printf("指针 p 存储的内存地址为：%p\n", &a) // 0x14000110018
	fmt.Printf("指针 p 自身的内存地址为：%p\n", &p) // 0x14000116018

	// 通过指针间接修改变量的值
	*p = 222
	fmt.Println("通过指针修改后变量 a 的值为：", a)
}
