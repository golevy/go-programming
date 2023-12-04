package main

import "fmt"

func ReturnResult1(num1 int, num2 int) int { // 这里最后的 int 是用来声明返回值的类型
	var sum int
	sum = num1 + num2
	return sum
}

func ReturnResult2(num1 int, num2 int) (sum int) { // 这里最后直接定义了返回值及其类型
	sum = num1 + num2
	return sum
}

func ReturnResult3(num1 int, num2 int) (sum int) {
	sum = num1 + num2
	return // 直接 return 是因为 返回值 sum 及其类型已经被定义，无需再 return sum
}

func main() {
	var result int // 定义一个变量接收函数的返回值

	fmt.Println("这是第一种方法的测试结果：")
	result = ReturnResult1(1, 5)
	fmt.Println(result / 3)

	fmt.Println("这是第二种方法的测试结果：")
	result = ReturnResult2(1, 5)
	fmt.Println(result / 3)

	fmt.Println("这是第三种方法的测试结果：")
	result = ReturnResult3(1, 5)
	fmt.Println(result / 3)
}
