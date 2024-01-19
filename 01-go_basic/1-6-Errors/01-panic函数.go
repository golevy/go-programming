package main

import "fmt"

func test(n int) {
	var num [3]int
	num[n] = 3 // 引发异常，从而强制终止整个程序的执行
	// panic: runtime error: index out of range [3] with length 3
	fmt.Println(num)
	panic("This is a panic") // func panic(v any)
}

func main() {
	test(3)
}
