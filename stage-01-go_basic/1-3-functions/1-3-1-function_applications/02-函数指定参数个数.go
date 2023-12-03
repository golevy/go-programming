package main

import "fmt"

var num1 int
var num2 int

func SumAdd(num1 int, num2 int) {
	// 计算 num1 到 num2 之间数字之和
	sum := 0
	for i := num1; i <= num2; i++ {
		sum += i
	}
	fmt.Printf("%d-%d 数字之和为：%d\n", num1, num2, sum)
}

func main() {
	fmt.Println("请输入需要求和的起始整数：")
	_, err := fmt.Scan(&num1)
	if err != nil {
		return
	}
	fmt.Println("请输入需要求和的结束整数：")
	_, err = fmt.Scan(&num2)
	if err != nil {
		return
	}

	SumAdd(num1, num2)
}
