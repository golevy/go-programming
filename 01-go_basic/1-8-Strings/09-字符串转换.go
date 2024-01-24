package main

// https://studygolang.com/pkgdoc(Go语言标准库文档)

import (
	"fmt"
	"strconv"
)

func main() {
	// 将bool类型转换为字符串类型
	str1 := strconv.FormatBool(true)
	fmt.Println(str1)

	// 将int类型转换为字符串类型
	str2 := strconv.Itoa(123)
	fmt.Println(str2)

	// 将字符串转换为bool类型
	b, err := strconv.ParseBool("true") // ParseBool：解析（或转换）一个字符串，判断这个字符串是否代表布尔值（即真或假），并将其转换成程序中的布尔类型（bool）
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}

	// 将字符串转换为int类型
	num, err := strconv.Atoi("123") // Atoi："ASCII to Integer"，意即“从ASCII码转换到整数”
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
