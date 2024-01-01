package main

import "fmt"

type per struct {
	name string
	age  int
}

type stu struct {
	per
	score float64
}

func (p *per) printInfo() {
	fmt.Println("这是父类中的方法")
}

func (s *stu) printInfo() {
	fmt.Println("这是子类中的方法")
}

func main() {
	var s stu
	s.printInfo()     // 这是子类中的方法
	s.per.printInfo() // 这是父类中的方法
}
