package main

import "fmt"

// 在 Go 语言中，接口命名通常遵循一种非正式的约定，即在接口名称的末尾加上 "-er" 后缀 (Sayer)
// 这种命名方式旨在描述实现该接口的对象的行为或能力

type Sayer interface {
	SayHi()
}

type StudentF struct {
	id   int
	name string
}

type TeacherB struct {
	id   int
	name string
}

func (s *StudentF) SayHi() {
	fmt.Println("Hi, I'm a student.")
}

func (t *TeacherB) SayHi() {
	fmt.Println("Hi, I'm a teacher.")
}

func WhoSay(p Sayer) { // 在 Go 语言中实现多态的核心在于定义一个接口，并让不同的类型实现这个接口
	p.SayHi()
}

func main() {
	var studentF StudentF
	var teacherB TeacherB

	WhoSay(&studentF)
	WhoSay(&teacherB)
}
