package main

import "fmt"

func main() {
	num1 := 10
	num2 := 20
	num3 := 30

	fmt.Print("num1 = ", num1)                                        // 默认打印不换行
	fmt.Println("num1 = ", num1, ",num2 = ", num2, ",num3 = ", num3)  // Println 默认打印换行
	fmt.Printf("num1 = %d, num2 = %d, num3 = %d\n", num1, num2, num3) // %d 表示输出的是一个整型标量中的值，\n 表示换行符
}
