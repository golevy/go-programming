package main

import "fmt"

func main() {
	var num1 = 10
	num1 += 10 // num1 = num1 + 10
	num1 -= 5  // num1 = num1 - 5
	num1 *= 3  // num1 = num1 * 3
	num1 /= 3  // num1 = num1 / 3
	num1 %= 2  // num1 = num1 % 2
	fmt.Println("num1 = ", num1)

	var num2 = 10
	num2 %= 2 + 3 // num2 &= 5 ===> num = num % 5，这里算术运算符的优先级是高于赋值运算符的，所以先计算 2 + 3，最后再取余
	fmt.Println("num2 = ", num2)
}
