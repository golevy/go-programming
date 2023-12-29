package main

import "fmt"

type StudentC struct {
	id   int
	name string
	age  int
}

type StudentD struct {
	id   int
	name string
	age  int
}

func (s1 *StudentC) PrintInfo() {
	fmt.Println(s1)
}

func (s2 *StudentD) PrintInfo() { // 接收者类型不同，即使方法名相同也是不同的方法
	fmt.Println(s2)
}

func main() {
	stu1 := StudentC{id: 001, name: "Nike", age: 20}
	(&stu1).PrintInfo()
	stu1.PrintInfo() // 相当于 (&stu1).PrintInfo()
	stu2 := StudentD{id: 002, name: "Mike", age: 22}
	stu2.PrintInfo() // 接收者类型不同，即使方法名相同也是不同的方法
}
