package main

import "fmt"

func main() {
	// 以矩形的方式打印九九乘法表
	for j := 1; j <= 9; j++ {
		for i := 1; i <= 9; i++ {
			// 2*3=6
			fmt.Printf("%d*%d=%d\t", j, i, j*i) // \t 相当于一个 tab 的缩进
		}
		fmt.Println()
	}
}
