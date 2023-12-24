package main

import "fmt"

type Person struct {
	id   int
	name string
	age  int
}

type Student struct {
	Person // 使用 Person 类型作为匿名字段，只有类型，没有成员的名字
	score  float64
}

type Teacher struct {
	Person // 匿名字段，只有类型，没有成员的名字
	salary float64
}

func main() {
	// 全部初始化
	var student = Student{Person{id: 1, name: "Levy", age: 23}, 100}
	fmt.Println("父类成员和子类成员全部初始化的结果为：", student) // {{1 Levy 23} 100}

	// 部分初始化
	var teacher = Teacher{Person{id: 2}, 5000}
	fmt.Println("父类成员部分初始化和子类成员全部初始化的结果为：", teacher) //  {{2  0} 5000}
}
