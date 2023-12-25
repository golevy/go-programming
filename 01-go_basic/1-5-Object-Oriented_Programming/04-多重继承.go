package main

import "fmt"

type Object struct {
	id int
}

type PersonA struct {
	Object
	name string
	age  int
}

type StudentA struct {
	PersonA
	score float64
}

func main() {
	var stu StudentA
	stu.age = 18 // 等价于 stu.PersonA.age
	fmt.Println(stu.PersonA.age)
	stu.id = 100 // 等价于 stu.PersonA.Object.id
	fmt.Println(stu.PersonA.Object.id)
}
