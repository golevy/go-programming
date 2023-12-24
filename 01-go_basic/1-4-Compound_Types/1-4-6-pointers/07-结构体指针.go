package main

import "fmt"

type Student struct {
	id   int
	name string
	age  int
}

func UpdateStu(stu *Student) {
	stu.age = 19
}

func main() {
	stu := Student{id: 1, name: "Go", age: 20}

	// 结构体指针的定义
	var p *Student
	p = &stu
	fmt.Println("结构体指针对应结构体的初始值为：", *p) // {1 Go 20}
	fmt.Println("结构体指针对应结构体的 name 为：", (*p).name)
	fmt.Println("结构体指针对应结构体的 age 为：", p.age) // 除了切片指针之外，数组指针和结构体指针都能使用 p 来快捷访问其中的值
	println()

	p.age = 23
	fmt.Println("经修改后结构体指针对应结构体的初始值为", *p)

	UpdateStu(p)
	fmt.Println("UpdateStu 后结构体指针对应结构体的初始值为：", *p)
}
