package main

import "fmt"

func main() {
	// 让用户输入一个年份，如果是闰年，则输出 true，如果不是，则输出 false
	// 闰年的判定（符合下述条件之一即满足）
	// 1. 年份能被 400 整除（如：2000）
	// 2. 年份能被 4 整除但不能被 100 整除（如：2008）
	fmt.Println("请输入要判断的闰年的年份：")

	var year int
	_, err := fmt.Scan(&year)
	if err != nil {
		return
	}

	b := year%400 == 0 || year%4 == 0 && year%100 != 0
	fmt.Println(b)
}
