package main

import "fmt"

func main() {
	// 计算 1 到 100（包含100）之间的所有整数的和，能被 7 整除的整数除外
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			continue
		}
		sum += i
	}
	fmt.Println("符合条件的所有整数和为：", sum)
}
