package main

import "fmt"

func main() {
	// 求 1-100 之间所有数字之和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1-100 之间所有数字之和为：", sum)
}
