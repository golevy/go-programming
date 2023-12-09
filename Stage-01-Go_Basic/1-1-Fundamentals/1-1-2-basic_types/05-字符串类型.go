package main

import "fmt"

func main() {
	var name1 string = "levy"
	var name2 string = "莱维"
	str := "a" // 字符串类型的变量一定要用 "" 双引号表示，区别于字符类型的 '' 单引号，并且字符串包含一个或多个字符

	fmt.Printf("%s\n", name1) // %s 表示输出字符串类型变量的格式符
	fmt.Printf("%T\n", str)   // "" 双引号括起来的变量自动推导类型为 string
	fmt.Println(len(name1))   // 使用 len() 方法可以获得字符串的长度，一个英文字母的长度为 1
	fmt.Println(len(name2))   // 注意在 Go 语言中，一个汉字的长度相当于 3 个英文字母的长度

	// 注意：字符串隐藏了一个结束符 '\0'，并且用 len() 函数测试字符串长度时，不会包含 '\0'
}
