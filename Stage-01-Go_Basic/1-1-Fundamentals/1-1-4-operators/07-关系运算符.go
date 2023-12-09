package main

import "fmt"

func main() {
	var num1 = 10
	var num2 = 20

	fmt.Println("num1 == num2", num1 == num2)
	fmt.Println("num1 != num2", num1 != num2)
	fmt.Println("num1 > num2", num1 > num2)
	fmt.Println("num1 < num2", num1 < num2)
	fmt.Println("num1 >= num2", num1 >= num2)
	fmt.Println("num1 <= num2", num1 <= num2)

	sum := num1+15 >= num2 // 算术运算符的优先级高于关系运算符
	fmt.Println(sum)
}
