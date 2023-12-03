package main

import "fmt"

func GetMultiResult1() (int, int) {
	var num1 = 10
	var num2 = 20
	return num1, num2
}

func GetMultiResult2() (num1 int, num2 int) {
	num1 = 10
	num2 = 20
	return num1, num2
}

func GetMultiResult3() (num1 int, num2 int) { // 定义好返回值的变量及其类型后，即可直接在函数里处理完数据后 return
	num1 = 10
	num2 = 20
	return
}

func main() {
	var result1, result2 int

	fmt.Println("第一种方法测试的结果：")
	result1, result2 = GetMultiResult1()
	fmt.Println(result1, result2)

	fmt.Println("第二种方法测试的结果：")
	result1, result2 = GetMultiResult2()
	fmt.Println(result1, result2)

	fmt.Println("第三种方法测试的结果：")
	result1, result2 = GetMultiResult3()
	fmt.Println(result1, result2)
}
