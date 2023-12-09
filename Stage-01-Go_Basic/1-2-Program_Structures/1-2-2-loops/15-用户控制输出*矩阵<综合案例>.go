package main

import "fmt"

func main() {
	// 要求用户输入一个数字 n，输出每边都有 n 个 * 的倒三角形
	var num int
	fmt.Println("请输入要每边要生成的 * 数：")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	for j := num; j > 0; j-- {
		for i := j; i > 0; i-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
