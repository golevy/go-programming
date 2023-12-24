package main

import "fmt"

func main() {
	var num1 float64
	var num2 float64
	num1 = 123.45
	num2 = 123.4567
	fmt.Printf("%f\n", num1)   // %f 表示输出浮点型的格式符
	fmt.Printf("%.2f\n", num2) // %.2f 表示保留两位小数，同时进行了四舍五入的格式符

	num3 := 11.12
	fmt.Printf("%T", num3) // %T 表示的是输出变量类型的格式符
}
