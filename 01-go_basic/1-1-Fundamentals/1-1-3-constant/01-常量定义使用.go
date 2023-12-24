package main

import "fmt"

func main() {
	const PI1 float64 = 3.14 // 显示声明常量的类型
	const PI2 = 3.14         // 可以自动推导常量的类型，注意这里区别于变量的 :=

	//PI1 = 3.15 // 不允许修改常量的值
	//fmt.Printf("%p\n", PI2) // 不允许打印常量的地址

	fmt.Printf("%.2f\n", PI1)
	fmt.Printf("%.2f", PI2)
}
