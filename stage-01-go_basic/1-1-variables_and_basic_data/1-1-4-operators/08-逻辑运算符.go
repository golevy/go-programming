package main

import "fmt"

func main() {
	var num1 = 10
	var num2 = 20
	isResult := num1 > num2
	fmt.Println("isResult", !isResult)

	// 逻辑非
	//fmt.Println(!num1 > num2) // 逻辑非后面的变量一定是布尔类型，逻辑非的运算优先级高于关系运算符
	fmt.Println(!(num1 > num2)) // 括号的运算优先级高于逻辑或

	// 逻辑与
	fmt.Println(num1 > num2 && num1 == 10) // 逻辑与运算优先级低于关系运算符

	// 逻辑或
	fmt.Println(num1 > num2 || num1 == 10) // 逻辑或运算优先级低于算术运算符

	fmt.Println(num1 < num2 || num1 > num2 && num1 != 10) // 逻辑与的优先级高于逻辑或
}
