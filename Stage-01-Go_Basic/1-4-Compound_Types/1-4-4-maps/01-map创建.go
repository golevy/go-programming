package main

import "fmt"

func main() {
	// 1. var 方式
	var map1 map[int]string // [int] 是键的类型，string 是值的类型
	fmt.Println("var 方式创建的 map 为：", map1)

	// 2. 自动推导类型方式
	map2 := map[int]string{} // [int] 是键的类型，string 是值的类型
	fmt.Println("自动推荐类型创建的 map 为：", map2)

	// 3. make() 函数的方式
	map3 := make(map[string]int, 10) // [string] 是键的类型，int 是值的类型
	fmt.Println("make() 函数创建的 map 为：", map3)
	fmt.Println("make() 函数创建的 map 中已有的键值对个数为：", len(map3)) // len() 函数返回的是 map 中已有键值对的个数
}
