package main

import "fmt"

type Greeter interface {
	SayHello()
	//  Say()  // Cannot use '&student' (type *StudentE) as the type Greeter Type does not implement 'Greeter' as some methods are missing: Say()
	//	如果一个类型要实现一个接口，它必须实现该接口中定义的所有方法
}

type StudentE struct {
	id   int
	name string
}

type TeacherA struct {
	id   int
	name string
}

func (s *StudentE) SayHello() {
	fmt.Println("Hello, I'm a student.")
}

func (t *TeacherA) SayHello() {
	fmt.Println("Hello, I'm a teacher.")
}

func main() {
	var student StudentE
	var teacher TeacherA
	var person Greeter

	person = &student
	person.SayHello()

	person = &teacher
	person.SayHello()
}
