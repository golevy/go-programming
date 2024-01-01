package main

import "fmt"

type personInfo struct {
	name string
	age  int
}

func (p *personInfo) printInfo() {
	fmt.Println(*p)
}

func main() {
	per := personInfo{name: "Mike", age: 18}

	// 1. 对象名.方法名
	per.printInfo() // {Mike 18}

	// 2.方法值
	f1 := per.printInfo
	fmt.Printf("%T\n", f1) // func()
	f1()                   // {Mike 18}

	// 3. 方法表达式
	f2 := (*personInfo).printInfo
	f2(&per) // {Mike 18}
}
