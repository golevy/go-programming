package main

import "fmt"

func main() {
	type Human struct {
		id   int
		name string
		age  int
		addr string
	}

	// 1. 全部初始化
	var s1 = Human{
		id:   1,
		name: "Go",
		age:  14,
		addr: "666",
	}
	fmt.Println("全部初始化的结构体 s1 为：", s1) // {1 Go 14 666}

	// 2. 指定成员初始化
	var s2 = Human{
		name: "Mike",
		age:  18,
	}
	fmt.Println("指定成员初始化的结构体 s2 为：", s2) // {0 Mike 18 }

	// 3. 通过“结构体变量.成员”初始化
	var s3 Human
	s3.id = 2
	s3.name = "Python"
	s3.age = 28
	fmt.Println("通过“结构体变量.成员”初始化的结构体 s3 为：", s3) //  {2 Python 28 }
}
