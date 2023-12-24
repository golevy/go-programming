package main

import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func main() {
	/*
		计算一个数的阶乘
		n! = 0 * 1 * 2 * 3 * ... * n
	*/
	var num int
	fmt.Println("请输入需要计算阶乘的数：")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("输入出错！！！", err)
		return
	}

	result := Factorial(num)
	fmt.Printf("%d 的阶乘是：%d", num, result)
}
