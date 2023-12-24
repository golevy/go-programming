package main

import "fmt"

type person struct {
	id   int
	name string
	age  int
}

type student struct {
	*person
	score float64
}

func main() {
	var stu = student{&person{id: 001, name: "Mike", age: 20}, 99}
	fmt.Println(stu.name)
}
