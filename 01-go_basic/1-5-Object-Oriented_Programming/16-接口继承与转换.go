package main

import "fmt"

type Humaner interface {
	SayHello()
}

type Personer interface {
	Humaner
	Say()
}

type StudentDemo struct {
}

func (s *StudentDemo) SayHello() {
	fmt.Println("大家好")
}

func (s *StudentDemo) Say() {
	fmt.Println("你好")
}

func main() {
	var stuDemo StudentDemo
	var perDemo Personer
	perDemo = &stuDemo

	perDemo.Say()
	perDemo.SayHello() // 可以调用从父接口 Humaner 继承的方法

	var h Humaner
	h = &stuDemo
	h.SayHello()
}
