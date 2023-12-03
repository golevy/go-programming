package main

import "fmt"

func main() {
	defer fmt.Println("我爱 Golang！！！")
	defer fmt.Println("Learn Go Happily!!!") // 如果一个函数中有多个 defer 语句，它们会以 LIFO（后进先出）的顺序执行
	fmt.Println("Levy 666!!!")
}
