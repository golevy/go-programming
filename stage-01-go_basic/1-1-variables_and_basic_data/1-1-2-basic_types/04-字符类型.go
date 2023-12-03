package main

import "fmt"

func main() {
	var ch byte = 'A' // 字符型变量一定要用 '' 单引号表示，并且包含一个字符

	fmt.Printf("%c\n", ch) // %c 表示输出字符类型变量的值
	fmt.Printf("%d", ch)   // %d 表示输出整型变量的值，这里直接就输出字符类型变量所对应的 ASCII 的值
}
