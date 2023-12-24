package main

import "fmt"

func main() {
	// 1.基本赋值
	var age, num int
	age = 10
	num = 20
	fmt.Println(age, num)

	// 2.变量给变量赋值
	var c int = 10
	var d int
	d = c
	fmt.Println(d)

	var e int = 10
	var f int = 20
	f = e
	fmt.Println(f)
}
