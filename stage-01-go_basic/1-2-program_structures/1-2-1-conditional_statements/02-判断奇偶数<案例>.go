package main

import "fmt"

func main() {
	// 输入一个数字，判断是否为偶数，如果是输出“该数是偶数”，否则输出“该数是奇数”。
	// 1. 定义一个整型变量接收用户输入数字
	var num int
	fmt.Println("请输入一个数字")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	// 2. 进行判断，能被 2 整除的数为偶数，并打印输出结果
	if num%2 == 0 {
		fmt.Println("输入的数字为偶数")
	} else {
		fmt.Println("输入的数字为奇数")
	}
}
