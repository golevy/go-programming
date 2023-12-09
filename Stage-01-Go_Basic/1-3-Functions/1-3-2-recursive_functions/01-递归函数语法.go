package main

import "fmt"

func Count(n int) {
	// 基本情况是当 n 等于 1 时，这是因为在这种情况下，函数不再进行递归调用
	// if n > 1 确保了当 n 等于 1 时，函数不会再调用自身，从而避免了无限递归
	if n > 1 {
		fmt.Println("前一排的排数是：", n-1)
		Count(n - 1)
	}
}

func main() {
	// 当用户输入一个数字（比如说5）后，程序首先打印出前一排的排数（即4），然后是前前一排的排数（即3），依次类推，直到第一排。
	var rowNum int
	fmt.Println("请输入你所在排数：")
	_, err := fmt.Scan(&rowNum)
	if err != nil {
		fmt.Println("输入错误！！！", err)
		return
	}

	Count(rowNum)
}
