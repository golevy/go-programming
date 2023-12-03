package main

import "fmt"

func main() {
	// 打印一个边长为 6 的等边三角形
	for j := 0; j < 6; j++ {
		for i := 0; i < j; i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
