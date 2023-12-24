package main

import "fmt"

func main() {
	var num1 float64 = 3.15
	var num2 int = 20

	fmt.Printf("%d\n", int(num1))     // 使用 int() 可以将 num1 从 float64 强制转换为 int32
	fmt.Printf("%f\n", float64(num2)) // 使用 float64 可以将 num2 从 int32 强制转换为 float64

	var num3 float32 = 5.8
	var num4 float64 = 6.6
	fmt.Printf("%f\n", float64(num3)+num4) // 使用 float64() 将 num3 从 float32 低类型转换成 float64 高类型，保证数据精度的同时满足相加的条件

	var num5 int = 1234
	fmt.Printf("%d", int8(num5)) // 使用 int8() 将 num5 从 int32 高类型转换为 int8 低类型，导致数据溢出，因为 int8 的范围是 -127~128，除此之外，还有可能会丢失精度
}
