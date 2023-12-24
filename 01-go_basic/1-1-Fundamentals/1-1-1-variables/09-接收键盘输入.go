package main

import "fmt"

func main() {
	var age int
	var name string

	fmt.Println("请输入年龄：")
	_, err := fmt.Scanf("%d", &age) // 通过 Scanf() 函数将键盘输入的数据赋值给变量，但是变量前一定要加上 &(取地址符)
	if err != nil {
		return
	}
	fmt.Println("age = ", age)
	fmt.Println(&age)        // 打印 age 的内存地址
	fmt.Printf("%p\n", &age) // 使用 %p 来格式化打印出 age 的内存地址

	fmt.Println("请输入姓名：")
	_, err = fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println("name = ", name)
	fmt.Println(&name)
}
