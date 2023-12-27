package main

import "fmt"

type StudentB struct {
	id   int
	name string
	age  int
}

func (s *StudentB) EditInfo() { // s 是 StudentB 类型的一个实例，它是一个接收器（receiver）
	s.age = 18 // 可以使用 s 来访问 StudentB 实例的属性和方法
}

func main() {
	var stu = StudentB{id: 001, name: "test", age: 20}
	stu.EditInfo()
	fmt.Println("调用结构体方法修改后的结果为：", stu)
}
