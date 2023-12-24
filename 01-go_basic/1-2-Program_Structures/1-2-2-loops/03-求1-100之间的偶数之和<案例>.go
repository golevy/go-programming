package main

import "fmt"

func main() {
	// 求 1-100 之间的偶数之和
	sum := 0
	for i := 2; i <= 100; i += 2 {
		sum += i
	}
	fmt.Println("1-100 之间的偶数之和为：", sum)
}
