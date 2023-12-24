package main

import "fmt"

func main() {
	// 用“*”打印一个长宽均为 6 的矩形
	for j := 0; j < 6; j++ {
		for i := 0; i < 6; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
