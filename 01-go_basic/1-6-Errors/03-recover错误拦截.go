package main

import "fmt"

func TestRecover() {
	fmt.Println(recover()) // runtime error: index out of range [10] with length 10
}

func Test(n int) {
	defer TestRecover() // recover() 只有在 defer 调用的函数中有效
	var num [10]int
	num[n] = 12
	fmt.Println(num)
}

func main() {
	Test(10)
}
