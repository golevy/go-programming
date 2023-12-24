package main

import "fmt"

func main() {
	var num1 = 10
	var num2 = 20

	num1++ // 将 num1 变量中的值取出来然后进行加 1 运算，运算完成后重新赋值给变量 num，在 Go 语言中只存在后自增 num++，而不存在前自增 ++num
	num2-- // 将 num2 变量中的值取出来然后进行减 1 运算，运算完成后重新赋值给变量 num，在 Go 语言中只存在后自减 num--，而不存在前自减 --num

	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)
}
