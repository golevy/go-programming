package main

import "fmt"

// 使用面向对象方式，实现一个支持两数加法和减法的计算器

type object struct {
	num1 int
	num2 int
}

type Add struct {
	object
}

type Sub struct {
	object
}

func (n *Add) GetResult(a int, b int) int {
	n.num1 = a
	n.num2 = b
	return n.num1 + n.num2
}

func (n *Sub) GetResult(a int, b int) int {
	n.num1 = a
	n.num2 = b
	return n.num1 - n.num2
}

func main() {
	var add Add
	r1 := add.GetResult(1, 2)
	fmt.Println(r1) // 3

	var sub Sub
	r2 := sub.GetResult(3, 2)
	fmt.Println(r2) // 1
}
