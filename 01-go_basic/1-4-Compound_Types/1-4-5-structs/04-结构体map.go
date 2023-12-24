package main

import "fmt"

type People struct {
	id   int
	name string
	age  int
}

func main() {
	// 结构体map的定义与初始化
	m := make(map[int]People)
	m[1] = People{
		id:   101,
		name: "One",
		age:  18,
	}
	m[2] = People{
		id:   102,
		name: "Two",
		age:  20,
	}
	fmt.Println("结构体map的初始化结果为：", m) // map[1:{101 One 18} 2:{102 Two 20}]
	fmt.Println(m[1])                // {101 One 18}
	fmt.Println(m[1].name)           // One
	println()

	// 循环遍历
	for k, v := range m {
		fmt.Println("键：", k)
		fmt.Println("值：", v)
	}
	println()

	// 删除结构体map中的成员
	delete(m, 1)
	fmt.Println("delete() 函数删除后的结构体map为：", m) // map[2:{102 Two 20}]
}
