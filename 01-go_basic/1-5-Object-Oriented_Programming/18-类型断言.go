package main

import "fmt"

func main() {
	var i interface{}
	i = 123

	value, ok := i.(int) // 通过类型断言来判断空接口中存储的数据类型
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("类型推断错误")
	}
}
