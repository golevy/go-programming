package main

import "fmt"

type Demo struct {
	id   int
	name string
	addr string
}

func PrintStruct(s Demo) {
	fmt.Println(s)
}

func RefactorStruct(s Demo) {
	s.name = "None" // 在函数中修改结构体成员的值，不会影响到原结构体
	// 结构体默认是按值传递的
	// 这意味着当你将一个结构体传递给一个函数时，实际上传递的是该结构体的一个副本
	// 因此，如果在函数内部修改了这个副本的字段，原始结构体并不会受到影响
}

func main() {
	s := Demo{
		id:   1,
		name: "Mike",
		addr: "US",
	}

	PrintStruct(s) // {1 Mike US}

	RefactorStruct(s)
	fmt.Println("修改结构体 s 的结果为：", s) // {1 Mike US}
}
