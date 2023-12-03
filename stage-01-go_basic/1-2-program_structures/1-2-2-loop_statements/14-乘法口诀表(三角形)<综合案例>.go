package main

import "fmt"

func main() {
	// 以三角形的方式打印九九乘法表
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			// 1*1=1
			// 2*1=2 2*2=4
			// 3*1=3 3*2=6 3*3=9
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
