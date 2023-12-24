package main

import "fmt"

func main() {
	// 请用户输入姓名、年龄，当用户输入完以后在屏幕上显示：您好XX，您的年龄是XX
	// 1. 定义变量
	var name string
	var age int
	// 2. 给出相应的录入提示
	fmt.Println("请输入姓名：")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println("请输入年龄：")
	_, err = fmt.Scan(&age)
	if err != nil {
		return
	}
	// 3. 完成输入
	// 4. 打印输出结果
	fmt.Printf("您好%s，您的年龄是%d", name, age) // 在这里 name 可以直接替换 %s 的位置，age 可以替换下一个 %d 的位置
	// 因为 %s 和 %d 格式符可以看作是占位符，按顺序依次接收后面提供的参数
}
